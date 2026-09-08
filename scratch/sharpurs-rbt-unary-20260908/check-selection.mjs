// Observe the same PBO callback input as Sharpurs, without generating F#.
import { writeFileSync, readFileSync, readdirSync } from "node:fs";
import { fileURLToPath, pathToFileURL } from "node:url";
import { dirname, resolve } from "node:path";
import { createHash } from "node:crypto";
import { execFileSync } from "node:child_process";

const audit = dirname(fileURLToPath(import.meta.url));
const benchmark = resolve(audit, "../..");
const backend = resolve(benchmark, "../sharpurs/sharpurs");
const pbo = resolve(benchmark, "../purescript-backend-optimizer-sharpurs");
process.chdir(audit); // Builder's .purmeta writes stay in this isolated directory.
const load = (name) => import(pathToFileURL(resolve(backend, "output", name, "index.js")));
const [Aff, Applicative, Either, Maybe, Set, Map, CoreFn, App, Builder, Foreign] =
  await Promise.all([
    "Effect.Aff", "Control.Applicative", "Data.Either", "Data.Maybe", "Data.Set",
    "Data.Map.Internal", "PureScript.Backend.Optimizer.CoreFn",
    "PureScript.Backend.Optimizer.App", "PureScript.Backend.Optimizer.Builder",
    "PureScript.Backend.Optimizer.Semantics.Foreign",
  ].map(load));
const pure = Applicative.pure(Aff.applicativeAff);
const run = (aff) => new Promise((resolve, reject) => {
  Aff.runAff((result) => () => {
    if (result instanceof Either.Left) reject(result.value0);
    else resolve(result.value0);
  })(aff)();
});
const encode = (value) => JSON.stringify(value, (_key, value) => {
  if (value && typeof value === "object" && !Array.isArray(value)) {
    const tag = value.__psTag || value.constructor?.name;
    if (tag && tag !== "Object") return { $tag: tag, ...value };
  }
  return value;
}, 2) + "\n";


const [CodeGen, Printer, FsAst, Foldable, ArrayModule, TupleModule, ListTypes] = await Promise.all([
  "Sharpurs.CodeGen", "Sharpurs.Printer", "Sharpurs.FsAst", "Data.Foldable", "Data.Array", "Data.Tuple", "Data.List.Types",
].map(load));
const modules = await run(App.coreFnModulesFromOutput(resolve(benchmark, "run/bak/sharp/output")));
const moduleArray = ArrayModule.fromFoldable(ListTypes.foldableList)(modules);
const entries = moduleArray.flatMap((module) => module.dataDecls.flatMap((decl) => decl.constructors.map((ctor) =>
  new TupleModule.Tuple(FsAst.sanitizeName(module.name.replaceAll(".", "_") + "_" + ctor.name), ctor.fields.length))));
const arities = Map.fromFoldable(CoreFn.ordIdent)(Foldable.foldableArray)(entries);
const checked = [], changed = [], selections = [];
const Kernel = await load("Sharpurs.AdtKernel");
let wrappers = Set.empty;
await run(Builder.buildModules(Aff.monadAff)({
  directives: await run(App.loadDirectives), rewriteLimit: 10000,
  analyzeCustom: (_) => (_) => Maybe.Nothing.value,
  foreignSemantics: Map.filterKeys(CoreFn.ordQualified(CoreFn.ordIdent))((qualified) =>
    !(qualified.value0 instanceof Maybe.Just) ||
    (!qualified.value0.value0.includes("Effect") && !qualified.value0.value0.includes("Control.Monad.ST"))
  )(Foreign.coreForeignSemantics),
  traceIdents: Set.empty,
  onPrepareModule: (_) => (module) => pure(module),
  onSkipModule: (_) => (_) => pure(Maybe.Nothing.value),
  onCodegenModule: (_) => (core) => (backendModule) => (_) => {
    const selected = Kernel.prepareUnary(core)(backendModule);
    if (selected instanceof Maybe.Just) {
      selections.push({module:backendModule.name, functions:selected.value0.nativeNames});
      for (const decl of selected.value0.layout.declarations) for (const ctor of decl.constructors) {
        wrappers=Set.insert(CoreFn.ordIdent)(backendModule.name.replaceAll(".", "_")+"_"+ctor.sourceName.value1)(wrappers);
      }
    }
    const code = Printer.printModule(CodeGen.translateOptimizedModuleWithAdts(wrappers)(selected)(arities)(backendModule)(core));
    const old = readFileSync(resolve(audit, "before", backendModule.name + ".fs"), "utf8");
    checked.push(backendModule.name);
    if (!old.endsWith(code + "\n")) changed.push(backendModule.name);
    return pure(undefined);
  },
})(modules));
writeFileSync(resolve(audit, "default-generation.json"), JSON.stringify({checkedModules: checked.length, changed, modules: checked}, null, 2) + "\n");
writeFileSync(resolve(audit,"selection.json"), JSON.stringify({selections,changed},null,2)+"\n");
console.log(`${checked.length} modules inspected for native unary ADTs.`);

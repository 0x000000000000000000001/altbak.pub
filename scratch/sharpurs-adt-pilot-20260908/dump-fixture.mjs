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

const generated = resolve(benchmark, "output/Main");
const hashGenerated = () => Object.fromEntries(readdirSync(generated).filter((name) =>
  /\.(fs|cs|fsproj|csproj|props)$/.test(name)).sort().map((name) => [
    name, createHash("sha256").update(readFileSync(resolve(generated, name))).digest("hex"),
  ]));
const generatedBefore = hashGenerated();
const Kernel = await load("Sharpurs.IntKernel");
const currentKernelAcceptance = {};
const modules = await run(App.coreFnModulesFromOutput(resolve(audit, "output")));
const directives = await run(App.loadDirectives);
let count = 0;
let found = false;
const captured = [];
await run(Builder.buildModules(Aff.monadAff)({
  directives,
  rewriteLimit: 10000,
  analyzeCustom: (_) => (_) => Maybe.Nothing.value,
  // Keep precisely Sharpurs Main's foreign-semantics filter.
  foreignSemantics: Map.filterKeys(CoreFn.ordQualified(CoreFn.ordIdent))((qualified) => {
    if (qualified.value0 instanceof Maybe.Just) {
      const name = qualified.value0.value0;
      return !name.includes("Effect") && !name.includes("Control.Monad.ST");
    }
    return true;
  })(Foreign.coreForeignSemantics),
  traceIdents: Set.empty,
  onPrepareModule: (_) => (module) => pure(module),
  onSkipModule: (_) => (_) => pure(Maybe.Nothing.value),
  onCodegenModule: (_) => (corefn) => (module) => (_) => {
    count++;
    if (module.name === "AdtPilot") {
      if (found) throw new Error("Duplicate AdtPilot module");
      found = true;
      const original = corefn.value0 ?? corefn;
      if (original.name !== module.name) throw new Error("Unexpected CoreFn module shape");
      writeFileSync(resolve(audit, "input.parsed.json"), encode({
        module: original.name, classDecls: original.classDecls,
        dataDecls: original.dataDecls, decls: original.decls,
      }));
      writeFileSync(resolve(audit, "module.metadata.json"), encode({
        module: module.name, classDecls: module.classDecls, dataDecls: module.dataDecls,
      }));
      for (const group of module.bindings) {
        for (const binding of group.bindings) {
          captured.push(binding.value0);
          currentKernelAcceptance[binding.value0] = Kernel.fromBinding(
            new CoreFn.Qualified(new Maybe.Just(module.name), binding.value0)
          )(binding.value1) instanceof Maybe.Just;
          writeFileSync(resolve(audit, `${binding.value0}.optimized.json`), encode({
            module: module.name, recursive: group.recursive,
            ident: binding.value0, expression: binding.value1,
          }));
          console.log(`Captured optimized AdtPilot.${binding.value0}`);
        }
      }
    }
    return pure(undefined);
  },
})(modules));
if (!found) throw new Error("Optimized AdtPilot not found");
for (const name of ["depth", "max", "empty", "singleton", "asymmetric", "singletonWith", "rootValue", "leftChild", "isRed", "rootColor"]) {
  if (!captured.includes(name)) throw new Error(`Missing ${name}`);
}

const generatedUnchanged = JSON.stringify(generatedBefore) === JSON.stringify(hashGenerated());
if (!generatedUnchanged) throw new Error("Generated files unexpectedly changed");
const inputs = [
  resolve(audit, "AdtPilot.purs"),
  resolve(backend, "output/PureScript.Backend.Optimizer.CoreFn/index.js"),
  resolve(audit, "output/AdtPilot/corefn.json"),
  resolve(backend, "src/Main.purs"),
  resolve(backend, "src/Sharpurs/IntKernel.purs"),
  resolve(backend, "output/Sharpurs.IntKernel/index.js"),
  ...["Builder", "App", "Convert", "Semantics", "Syntax", "Semantics.Foreign"]
    .map((name) => resolve(backend, `output/PureScript.Backend.Optimizer.${name}/index.js`)),
];
writeFileSync(resolve(audit, "metadata.json"), encode({
  node: process.version, optimizedModules: count, capturedBindings: captured,
  currentKernelAcceptance, generatedUnchanged, generatedFilesChecked: Object.keys(generatedBefore).length,
  commits: Object.fromEntries([benchmark, backend, pbo].map((repo) => [
    repo, execFileSync("git", ["-C", repo, "rev-parse", "HEAD"], { encoding: "utf8" }).trim(),
  ])),
  sha256: Object.fromEntries(inputs.map((path) => [
    path, createHash("sha256").update(readFileSync(path)).digest("hex"),
  ])),
}));
console.log(`Optimized ${count} modules; no code generation, no benchmark execution`);

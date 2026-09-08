// Observe the same PBO callback input as Sharpurs, without generating F#.
import { writeFileSync, readFileSync } from "node:fs";
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

const modules = await run(App.coreFnModulesFromOutput(resolve(benchmark, "output")));
const directives = await run(App.loadDirectives);
let count = 0;
let found = false;
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
  onCodegenModule: (_) => (_) => (module) => (_) => {
    count++;
    if (module.name === "Test.TCO") {
      for (const group of module.bindings) {
        for (const binding of group.bindings) {
          if (binding.value0 === "deepTailRec") {
            if (found) throw new Error("Duplicate deepTailRec binding");
            found = true;
            writeFileSync(resolve(audit, "deepTailRec.optimized.json"), encode({
              module: module.name, recursive: group.recursive,
              ident: binding.value0, expression: binding.value1,
            }));
            console.log("Captured optimized Test.TCO.deepTailRec from onCodegenModule");
          }
        }
      }
    }
    return pure(undefined);
  },
})(modules));
if (!found) throw new Error("Optimized deepTailRec not found");

const inputs = [
  resolve(benchmark, "output/Test.TCO/corefn.json"),
  resolve(backend, "src/Main.purs"),
  ...["Builder", "App", "Convert", "Semantics", "Syntax", "Semantics.Foreign"]
    .map((name) => resolve(backend, `output/PureScript.Backend.Optimizer.${name}/index.js`)),
];
writeFileSync(resolve(audit, "metadata.json"), encode({
  node: process.version, optimizedModules: count,
  commits: Object.fromEntries([benchmark, backend, pbo].map((repo) => [
    repo, execFileSync("git", ["-C", repo, "rev-parse", "HEAD"], { encoding: "utf8" }).trim(),
  ])),
  sha256: Object.fromEntries(inputs.map((path) => [
    path, createHash("sha256").update(readFileSync(path)).digest("hex"),
  ])),
}));
console.log(`Optimized ${count} modules; no code generation, no benchmark execution`);

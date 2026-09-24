// Observe the same PBO callback input as Sharpurs and report, for each
// Test.RBTree binding, whether the current Sharpurs converters accept the
// optimized IR. Also traces the first rejection of the Optimized envelope.
import { readFileSync, writeFileSync } from "node:fs";
import { fileURLToPath, pathToFileURL } from "node:url";
import { dirname, resolve } from "node:path";
import { createHash } from "node:crypto";

const audit = dirname(fileURLToPath(import.meta.url));
const benchmark = resolve(audit, "../..");
const backend = resolve(benchmark, "../sharpurs/sharpurs");
process.chdir(audit); // Builder's .purmeta writes stay in this isolated directory.

const load = (name) => import(pathToFileURL(resolve(backend, "output", name, "index.js")));
const [
  Aff, Applicative, Either, Maybe, Set, MapInternal, CoreFn, App, Builder, Foreign,
  Syntax, IntKernel, Optimized, ThunkKernel, AdtKernel, AdtLayout, DataArray, DataMap, DataFoldable, DataEq,
] = await Promise.all([
  "Effect.Aff", "Control.Applicative", "Data.Either", "Data.Maybe", "Data.Set",
  "Data.Map.Internal", "PureScript.Backend.Optimizer.CoreFn",
  "PureScript.Backend.Optimizer.App", "PureScript.Backend.Optimizer.Builder",
  "PureScript.Backend.Optimizer.Semantics.Foreign",
  "PureScript.Backend.Optimizer.Syntax", "Sharpurs.IntKernel", "Sharpurs.Optimized",
  "Sharpurs.ThunkKernel", "Sharpurs.AdtKernel", "Sharpurs.AdtLayout",
  "Data.Array", "Data.Map", "Data.Foldable", "Data.Eq",
].map(load));

const pure = Applicative.pure(Aff.applicativeAff);
const run = (aff) => new Promise((resolveRun, reject) => {
  Aff.runAff((result) => () => {
    if (result instanceof Either.Left) reject(result.value0);
    else resolveRun(result.value0);
  })(aff)();
});

const encode = (value) => JSON.stringify(value, (_key, value) => {
  if (value && typeof value === "object" && !Array.isArray(value)) {
    const tag = value.__psTag || value.constructor?.name;
    if (tag && tag !== "Object") return { $tag: tag, ...value };
  }
  return value;
}, 2) + "\n";

const Nothing = Maybe.Nothing.value;
const Just = (v) => new Maybe.Just(v);
const eqType = (a) => (b) => DataEq.eq(CoreFn.eqExprType)(a)(b);
const CoreFnProbe = await load("PureScript.Backend.Optimizer.CoreFn");

// ---------------------------------------------------------------------------
// Faithful mirror of Sharpurs.Optimized.lower with rejection tracing.
// ---------------------------------------------------------------------------
const fromLocal = IntKernel.fromLocal;
const fmtType = (t) => {
  if (t == null) return "null";
  if (typeof t === "string") return t;
  const tag = t?.constructor?.name ?? String(t);
  if (tag === "Just") return fmtType(t.value0);
  if (tag === "Nothing") return "Nothing";
  if (tag === "Func") return `[${t.value0.map(fmtType).join(", ")}] -> ${fmtType(t.value1)}`;
  if (tag === "TypeVar") return `a:${t.value0}`;
  if (tag === "ADT") return `ADT(${t.value1}, [${t.value2.map(fmtType).join(", ")}])`;
  if (tag === "Int" || tag === "String" || tag === "Boolean") return tag;
  return tag;
};
const showSyntax = (v) => (v?.constructor?.name ?? typeof v);
const tail = (v) => (v && v.constructor?.name === "Local" ? ` level=${v.value1}` : "");

function lowerTrace(scope, expected, v, depth = 0, path = "root") {
  // scope: native Map<number, ExprType> mirroring the PBO locals map.
  const isJustOf = (maybe, cls) => maybe instanceof Maybe.Just && maybe.value0 instanceof cls;
  const where = `${path} ${showSyntax(v)}${tail(v)} expected=${fmtType(expected)}`;
  const fail = (reason) => ({ ok: false, reason, where });
  const intLevels = [...scope.entries()]
    .filter(([, ty]) => eqType(ty)(CoreFn.Int.value))
    .map(([level]) => level)
    .sort((a, b) => a - b);
  const local = fromLocal(intLevels)(v);
  if (local instanceof Maybe.Just) {
    if (expected instanceof Maybe.Nothing || eqType(expected.value0)(CoreFn.Int.value)) return { ok: true, kernel: true };
    return fail("fromLocal accepted but expected type is not Int");
  }
  if (v instanceof Syntax.Typed) {
    if (isJustOf(expected, CoreFn.Func) && !eqType(expected.value0)(v.value0)) {
      return fail(`Typed annotation ${fmtType(v.value0)} disagrees with expected ${fmtType(expected)}`);
    }
    return lowerTrace(scope, Just(v.value0), v.value1, depth + 1, path + "/Typed");
  }
  if (v instanceof Syntax.TypeApp) {
    // isGlobalReference on the fn
    let fn = v.value0;
    let ok = false;
    for (;;) {
      if (fn instanceof Syntax.Var && fn.value0.value0 instanceof Maybe.Just) { ok = true; break; }
      if (fn instanceof Syntax.Typed) { fn = fn.value1; continue; }
      if (fn instanceof Syntax.TypeApp) { fn = fn.value0; continue; }
      break;
    }
    if (!ok) return fail("TypeApp head is not a global reference");
    return lowerTrace(scope, expected, v.value0, depth + 1, path + "/TypeApp");
  }
  if (v instanceof Syntax.Var) {
    if (v.value0.value0 instanceof Maybe.Just) return { ok: true };
    return fail("unqualified/local Var without level");
  }
  if (v instanceof Syntax.Local) {
    if (scope.has(v.value1)) return { ok: true };
    return fail(`Local level=${v.value1} was not introduced by an enclosing Abs`);
  }
  if (v instanceof Syntax.Lit) {
    const tag = v.value0.constructor.name;
    if (tag === "LitInt" || tag === "LitString" || tag === "LitBoolean") return { ok: true };
    return fail(`literal ${tag} is not in the Optimized subset`);
  }
  if (v instanceof Syntax.App) {
    const head = lowerTrace(scope, Nothing, v.value0, depth + 1, path + "/head");
    if (!head.ok) return { ...head, reason: `App head: ${head.reason}` };
    const args = [...v.value1];
    for (let i = 0; i < args.length; i++) {
      const arg = lowerTrace(scope, Nothing, args[i], depth + 1, `${path}/arg${i}`);
      if (!arg.ok) return { ...arg, reason: `App argument ${i}: ${arg.reason}` };
    }
    return { ok: true };
  }
  if (v instanceof Syntax.Abs) {
    if (!isJustOf(expected, CoreFn.Func)) {
      return fail(`Abs ${JSON.stringify([...v.value0].map((t) => t.value1))} reached with expected=${fmtType(expected)}`);
    }
    const signature = expected.value0;
    const levels = [...v.value0].map((t) => t.value1);
    if (levels.length > signature.value0.length) return fail("Abs has more binders than its signature");
    if (new globalThis.Set(levels.map(String)).size !== levels.length) return fail("Abs has duplicate binder levels");
    if (levels.some((l) => l < 0 || scope.has(l))) {
      return fail("Abs binder collides with an enclosing scope");
    }
    const inner = new globalThis.Map(scope);
    levels.forEach((level, i) => inner.set(level, signature.value0[i]));
    const remaining = signature.value0.slice(levels.length);
    const nextType = Just(remaining.length === 0 ? signature.value1 : new CoreFn.Func(remaining, signature.value1));
    return lowerTrace(inner, nextType, v.value1, depth + 1, path + "/Abs");
  }
  if (v instanceof Syntax.Let || v instanceof Syntax.LetRec) {
    return fail(`${showSyntax(v)} must be absorbed by IntKernel.fromLocal (whole-binding)`);
  }
  return fail(`unsupported node ${showSyntax(v)}`);
}

// ---------------------------------------------------------------------------
const modules = await run(App.coreFnModulesFromOutput(resolve(benchmark, "output")));
const directives = await run(App.loadDirectives);
let count = 0;
let captured = null;
await run(Builder.buildModules(Aff.monadAff)({
  directives,
  rewriteLimit: 10000,
  analyzeCustom: (_) => (_) => Nothing,
  foreignSemantics: MapInternal.filterKeys(CoreFn.ordQualified(CoreFn.ordIdent))((qualified) => {
    if (qualified.value0 instanceof Maybe.Just) {
      const name = qualified.value0.value0;
      return !name.includes("Effect") && !name.includes("Control.Monad.ST");
    }
    return true;
  })(Foreign.coreForeignSemantics),
  traceIdents: Set.empty,
  onPrepareModule: (_) => (module) => pure(module),
  onSkipModule: (_) => (_) => pure(Nothing),
  onCodegenModule: (_) => (corefn) => (module) => (_) => {
    count++;
    if (module.name === "Test.RBTree") {
      if (captured) throw new Error("Duplicate Test.RBTree module");
      captured = { corefn, module };
    }
    return pure(undefined);
  },
})(modules));
if (!captured) throw new Error("Optimized Test.RBTree not found");
console.log(`Optimized ${count} modules`);

const { corefn, module } = captured;
const report = { module: module.name, bindings: {} };
for (const group of module.bindings) {
  for (const binding of group.bindings) {
    const name = binding.value0;
    const expr = binding.value1;
    const optimized = Optimized.fromBinding(expr);
    const kernel = IntKernel.fromBinding(new CoreFn.Qualified(Just(module.name), name))(expr);
    const trace = lowerTrace(new globalThis.Map(), Nothing, expr);
    report.bindings[name] = {
      recursive: group.recursive,
      optimizedAccepted: optimized instanceof Maybe.Just,
      globalKernelAccepted: kernel instanceof Maybe.Just,
      trace: trace.ok ? { ok: true, note: trace.kernel ? "absorbed by local Int kernel" : "envelope accepted" }
        : { ok: false, reason: trace.reason, where: trace.where },
    };
    writeFileSync(resolve(audit, `${name}.optimized.json`), encode({
      module: module.name, recursive: group.recursive, ident: name, expression: expr,
    }));
    console.log(`${name}: optimized=${optimized instanceof Maybe.Just} kernel=${kernel instanceof Maybe.Just} ${trace.ok ? "accepted" : "REJECTED: " + trace.reason + " @ " + trace.where}`);
  }
}
const thunks = ThunkKernel.prepareModule(corefn.value0 ?? corefn)(module);
report.thunks = thunks instanceof Maybe.Just ? "accepted" : "rejected";
const adt = AdtKernel.prepareUnary(corefn.value0 ?? corefn)(module);
report.adt = adt instanceof Maybe.Just ? "accepted" : "rejected";
writeFileSync(resolve(audit, "report.json"), encode(report));
console.log(`thunks=${report.thunks} adt=${report.adt}`);
console.log(`Inputs sha256: output/Test.RBTree/corefn.json ${createHash("sha256").update(readFileSync(resolve(benchmark, "output/Test.RBTree/corefn.json"))).digest("hex").slice(0, 12)}`);

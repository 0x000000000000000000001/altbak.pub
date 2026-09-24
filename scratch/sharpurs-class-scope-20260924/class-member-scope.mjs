// Class member annotation scoping regression tests.
// After building PBO (or the backend that embeds it):
//   node test/class-member-scope.mjs [compiled-output-directory]
// The optional directory permits testing the exact PBO build used by a backend.
import assert from "node:assert/strict";
import { resolve } from "node:path";
import { fileURLToPath, pathToFileURL } from "node:url";

const output = process.argv[2]
  ? resolve(process.argv[2])
  : fileURLToPath(new URL("../output/", import.meta.url));
const load = (name) => import(pathToFileURL(resolve(output, name, "index.js")));
const [C, Convert, Maybe, Tuple] = await Promise.all([
  "PureScript.Backend.Optimizer.CoreFn",
  "PureScript.Backend.Optimizer.Convert",
  "Data.Maybe", "Data.Tuple",
].map(load));

const nothing = Maybe.Nothing.value;
const just = (value) => new Maybe.Just(value);
const pair = (left, right) => new Tuple.Tuple(left, right);
const variable = (name) => new C.TypeVar(name);
const func = (args, result) => new C.Func(args, result);
const forall = (vars, body) => new C.ForAll(vars, body);
const constrained = (entries, body) => new C.ConstrainedType(entries, body);
const int = C.Int.value;
const emptySpan = C.emptySpan;
const ann = (type, meta = nothing) => new C.Ann({ span: emptySpan, meta, type });
const accessor = (type) => new C.ExprVar(ann(just(type)), new C.Qualified(just("Fixture"), "accessor"));
const classDecl = (name, vars, methods) => ({
  name, vars, superclasses: [], methods,
});
const monklish = classDecl("Monoidish", ["a"], [pair("mempty_", variable("a"))]);
const className = ["Test", "Polymorphism", "Monoidish"];

// The class member's signature quantifies the scoped variable while the
// generated body annotations still use the class declaration's variable name.
const memberType = forall(["a$scope0"],
  constrained([pair(className, [variable("a$scope0")])], variable("a$scope0")));
const member = new C.NonRec(new C.Binding(ann(just(memberType)), "mempty_",
  new C.ExprAbs(ann(just(memberType)), "dict", accessor(variable("a")))));

// A plain polymorphic binding has no declared class member and must not change.
const plainType = forall(["a"], func([variable("a")], variable("a")));
const plain = new C.NonRec(new C.Binding(ann(just(plainType)), "identity",
  new C.ExprAbs(ann(just(plainType)), "value", accessor(variable("a")))));

// A constrained binding whose identifier is not a declared method stays as is.
const other = new C.NonRec(new C.Binding(ann(just(memberType)), "notAMethod",
  new C.ExprAbs(ann(just(memberType)), "dict", accessor(variable("a")))));

// A constraint argument that is not a quantified variable cannot be paired.
const concreteType = forall([], constrained([pair(className, [int])], int));
const concrete = new C.NonRec(new C.Binding(ann(just(concreteType)), "mempty_",
  accessor(variable("a"))));

let passed = 0;
let failed = 0;
const test = (name, run) => {
  try {
    run();
    passed++;
    console.log(`PASS ${name}`);
  } catch (error) {
    failed++;
    console.error(`FAIL ${name}\n${error.stack}`);
  }
};
const align = (bindings, classDecls = [monklish]) => {
  const [result] = Convert.alignClassMemberAnnotations(classDecls)(bindings);
  return result;
};
const bindingOf = (bind) => (bind instanceof C.NonRec ? bind.value0 : bind.value0[0]);

test("a declared class member uses its own quantified variable in body annotations", () => {
  const result = align([member]);
  const body = bindingOf(result).value2;
  assert.deepStrictEqual(C.exprAnn(body).type, just(memberType));
  assert.deepStrictEqual(C.exprAnn(body.value2).type.value0, variable("a$scope0"));
});

test("the member's nested annotations are aligned recursively", () => {
  const nested = new C.NonRec(new C.Binding(ann(just(memberType)), "mempty_",
    new C.ExprAbs(ann(just(memberType)), "dict",
      new C.ExprAbs(ann(just(func([variable("a")], variable("a")))), "value",
        accessor(variable("a"))))));
  const result = align([nested]);
  const lambda = bindingOf(result).value2;
  assert.deepStrictEqual(C.exprAnn(lambda).type, just(memberType));
  assert.deepStrictEqual(C.exprAnn(lambda.value2).type, just(func([variable("a$scope0")], variable("a$scope0"))));
  assert.deepStrictEqual(C.exprAnn(lambda.value2.value2).type.value0, variable("a$scope0"));
});

test("bindings without a declared class member keep every annotation", () => {
  const [plainResult, otherResult] = Convert.alignClassMemberAnnotations([monklish])([plain, other]);
  assert.deepStrictEqual(C.exprAnn(bindingOf(plainResult).value2.value2).type.value0, variable("a"));
  assert.deepStrictEqual(C.exprAnn(bindingOf(otherResult).value2.value2).type.value0, variable("a"));
});

test("a concrete constraint argument does not invent a variable substitution", () => {
  const result = align([concrete]);
  assert.deepStrictEqual(C.exprAnn(bindingOf(result).value2).type.value0, variable("a"));
});

test("recursive groups are aligned member by member", () => {
  const group = new C.Rec([bindingOf(member), bindingOf(plain)]);
  const [result] = Convert.alignClassMemberAnnotations([monklish])([group]);
  assert.ok(result instanceof C.Rec);
  assert.deepStrictEqual(C.exprAnn(result.value0[0].value2.value2).type.value0, variable("a$scope0"));
  assert.deepStrictEqual(C.exprAnn(result.value0[1].value2.value2).type.value0, variable("a"));
});

test("an unknown class leaves annotations untouched", () => {
  const foreignName = ["Other", "Class"];
  const type = forall(["a$scope0"], constrained([pair(foreignName, [variable("a$scope0")])], variable("a$scope0")));
  const binding = new C.NonRec(new C.Binding(ann(just(type)), "mempty_", accessor(variable("a"))));
  const result = align([binding]);
  assert.deepStrictEqual(C.exprAnn(bindingOf(result).value2).type.value0, variable("a"));
});

console.log(`Class member scope regression tests: ${passed} passed, ${failed} failed (${output})`);
if (failed) process.exitCode = 1;

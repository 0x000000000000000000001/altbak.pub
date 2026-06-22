import * as $runtime from "../runtime.js";
import * as Data$dShow from "../Data.Show/index.js";
import * as Effect$dConsole from "../Effect.Console/index.js";
const $Expr = (tag, _1, _2) => ({tag, _1, _2});
const Val = value0 => $Expr("Val", value0);
const Add = value0 => value1 => $Expr("Add", value0, value1);
const Mul = value0 => value1 => $Expr("Mul", value0, value1);
const Sub = value0 => value1 => $Expr("Sub", value0, value1);
const fib = v => {
  if (v === 0) { return 0; }
  if (v === 1) { return 1; }
  return fib(v - 1 | 0) + fib(v - 2 | 0) | 0;
};
const $$eval = v => {
  if (v.tag === "Val") { return v._1; }
  if (v.tag === "Add") { return $$eval(v._1) + $$eval(v._2) | 0; }
  if (v.tag === "Mul") { return $$eval(v._1) * $$eval(v._2) | 0; }
  if (v.tag === "Sub") { return $$eval(v._1) - $$eval(v._2) | 0; }
  $runtime.fail();
};
const buildTree = v => {
  if (v === 0) { return $Expr("Val", 1); }
  return $Expr("Add", $Expr("Mul", $Expr("Val", v), buildTree(v - 1 | 0)), $Expr("Sub", buildTree(v - 1 | 0), $Expr("Val", 1)));
};
const main = /* #__PURE__ */ (() => {
  const $0 = Effect$dConsole.log("=== AST Evaluation ===");
  return () => {
    $0();
    Effect$dConsole.log(Data$dShow.showIntImpl($$eval(buildTree(12))))();
    Effect$dConsole.log("=== Fibonacci ===")();
    return Effect$dConsole.log(Data$dShow.showIntImpl(fib(40)))();
  };
})();
export {$Expr, Add, Mul, Sub, Val, buildTree, $$eval as eval, fib, main};

export const runAstTreeFFI = function(depth) {
  return evaluate(buildTree(depth));
};

function buildTree(n) {
  if (n === 0) return { type: "Val", value: 1 };
  return { type: "Add",
    left: { type: "Mul", left: { type: "Val", value: n }, right: buildTree(n - 1) },
    right: { type: "Sub", left: buildTree(n - 1), right: { type: "Val", value: 1 } } };
}

function evaluate(tree) {
  switch (tree.type) {
    case "Val": return tree.value;
    case "Add": return evaluate(tree.left) + evaluate(tree.right);
    case "Mul": return evaluate(tree.left) * evaluate(tree.right);
    case "Sub": return evaluate(tree.left) - evaluate(tree.right);
    default: throw new Error("Invalid expression");
  }
}

export const runAstTreeFFICheatcode = function(limit) {
  let depth = Math.floor(limit);
  let tree = buildTree(depth);
  return evalTree(tree);
};

function buildTree(depth) {
  if (depth === 0) {
    return { type: "Literal", value: 1 };
  }
  return {
    type: "Add",
    left: buildTree(depth - 1),
    right: buildTree(depth - 1)
  };
}

function evalTree(node) {
  if (node.type === "Literal") {
    return node.value;
  }
  if (node.type === "Add") {
    return evalTree(node.left) + evalTree(node.right);
  }
  return 0;
}

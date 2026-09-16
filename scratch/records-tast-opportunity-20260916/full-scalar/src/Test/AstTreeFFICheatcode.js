export const runAstTreeFFICheatcode = function(limit) {
  function buildAst(n) {
    if (n === 0) return { typ: 0, value: 1 };
    return {
      typ: 1,
      left: { typ: 2, left: { typ: 0, value: n }, right: buildAst(n - 1) },
      right: { typ: 3, left: buildAst(n - 1), right: { typ: 0, value: 1 } }
    };
  }
  function evalAst(e) {
    switch (e.typ) {
      case 0: return e.value;
      case 1: return evalAst(e.left) + evalAst(e.right);
      case 2: return evalAst(e.left) * evalAst(e.right);
      case 3: return evalAst(e.left) - evalAst(e.right);
    }
    return 0;
  }
  return evalAst(buildAst(Math.floor(limit)));
};

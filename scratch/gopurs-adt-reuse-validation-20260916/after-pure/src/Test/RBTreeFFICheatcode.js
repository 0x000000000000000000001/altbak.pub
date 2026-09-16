
class RBNodeCheat {
  constructor(c, l, v, r) {
    this.color = c;
    this.left = l;
    this.value = v;
    this.right = r;
  }
}
export const runRBTreeFFICheatcode = function(limit) {
  function balance(c, a, x, b) {
    if (c === 1) { // black
        if (a !== null && a.color === 0) {
            if (a.left !== null && a.left.color === 0) {
                return new RBNodeCheat(0, new RBNodeCheat(1, a.left.left, a.left.value, a.left.right), a.value, new RBNodeCheat(1, a.right, x, b));
            }
            if (a.right !== null && a.right.color === 0) {
                return new RBNodeCheat(0, new RBNodeCheat(1, a.left, a.value, a.right.left), a.right.value, new RBNodeCheat(1, a.right.right, x, b));
            }
        }
        if (b !== null && b.color === 0) {
            if (b.left !== null && b.left.color === 0) {
                return new RBNodeCheat(0, new RBNodeCheat(1, a, x, b.left.left), b.left.value, new RBNodeCheat(1, b.left.right, b.value, b.right));
            }
            if (b.right !== null && b.right.color === 0) {
                return new RBNodeCheat(0, new RBNodeCheat(1, a, x, b.left), b.value, new RBNodeCheat(1, b.right.left, b.right.value, b.right.right));
            }
        }
    }
    return new RBNodeCheat(c, a, x, b);
  }

  function ins(x, t) {
      if (t === null) return new RBNodeCheat(0, null, x, null);
      if (x < t.value) return balance(t.color, ins(x, t.left), t.value, t.right);
      if (x > t.value) return balance(t.color, t.left, t.value, ins(x, t.right));
      return t;
  }

  function insert(x, t) {
      let res = ins(x, t);
      return new RBNodeCheat(1, res.left, res.value, res.right);
  }

  function depth(t) {
      if (t === null) return 0;
      let ld = depth(t.left);
      let rd = depth(t.right);
      return 1 + (ld > rd ? ld : rd);
  }

  let n = Math.floor(limit);
  let acc = null;
  for (let i = n; i > 0; i--) {
      acc = insert(i, acc);
  }
  return depth(acc);
};

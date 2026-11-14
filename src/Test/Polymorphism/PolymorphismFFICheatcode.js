class Monoidish {
  mempty_() { return 1; }
  mappend_(x, y) { return x + y; }
}
export const runPolymorphismFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let acc = 0;
  let m = new Monoidish();
  while (n > 0) {
    acc = m.mappend_(acc, m.mempty_());
    n--;
  }
  return acc;
};

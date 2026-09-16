export const runRecordsFFI = function(limit) {
  let n = Math.floor(limit);
  let r = { a: 0, b: { c: 0, d: { e: 0, f: 0 } } };
  while (n > 0) {
    r = {
      a: r.a + 1,
      b: {
        c: r.b.c + 2,
        d: {
          e: r.b.d.e + 3,
          f: r.b.d.f + (n % 5)
        }
      }
    };
    n--;
  }
  return r.b.d.f;
};

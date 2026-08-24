export const runRecordsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let r = { a: 0, b: { c: 0, d: { e: 0, f: 0 } } };
  while (n > 0) {
    r.a += 1;
    r.b.c += 2;
    r.b.d.e += 3;
    r.b.d.f += (n % 5);
    n--;
  }
  return r.b.d.f;
};

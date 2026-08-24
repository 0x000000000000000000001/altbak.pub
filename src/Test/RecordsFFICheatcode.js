export const runRecordsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let obj = { a: 1, b: { c: 2, d: 3 } };
  for (let i = 0; i < n; i++) {
    obj.b.c += 1;
  }
  return obj.b.c - 2;
};

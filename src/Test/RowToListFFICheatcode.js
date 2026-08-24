export const runRowToListFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  for (let i = 0; i < n; i++) {
    let rec = { a: 1, b: "hello", c: true };
    sum += rec.a;
  }
  return sum;
};

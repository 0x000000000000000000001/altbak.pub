export const runStateMonadFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  for (let i = 1; i <= n; i++) {
    sum += i;
  }
  return sum;
};

export const runArrayOpsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  for (let i = 1; i <= n; i++) {
    if (i % 2 === 0) sum += i;
  }
  return sum;
};

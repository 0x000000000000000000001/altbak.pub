export const runArrayOpsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  const step = n >= 1 ? 1 : -1;
  for (let i = 1; i !== n + step; i += step) {
    if (i % 2 === 0) sum += i;
  }
  return sum;
};

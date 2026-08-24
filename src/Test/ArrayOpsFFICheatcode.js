export const runArrayOpsFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let arr = [];
  for (let i = 1; i <= n; i++) {
    arr.push(i);
  }
  let sum = 0;
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] % 2 === 0) sum += arr[i];
  }
  return sum;
};

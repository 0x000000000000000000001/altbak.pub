export const runArrayOpsFFI = function(limit) {
  let n = Math.floor(limit);
  let arr = [];
  for (let i = 1; i <= n; i++) {
    arr.push(i);
  }
  
  let evens = arr.filter(function(x) { return x % 2 === 0; });
  
  let sum = evens.reduce(function(acc, x) { return acc + x; }, 0);
  
  return sum;
};

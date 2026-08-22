export const runPrimesFFI = function(limit) {
  let n = Math.floor(limit);
  let count = 0;
  for (let i = 2; count < n; i++) {
    let isPrime = true;
    for (let j = 2; j * j <= i; j++) {
      if (i % j === 0) {
        isPrime = false;
        break;
      }
    }
    if (isPrime) {
      count++;
    }
  }
  return count;
};

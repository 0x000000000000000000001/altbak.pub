export const runTCOFFI = function(limit) {
  let n = Math.floor(limit);
  let acc = 0;
  while (n > 0) {
    acc += (n % 3);
    n -= 1;
  }
  return acc;
};

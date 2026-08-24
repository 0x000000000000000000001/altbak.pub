export const runTCOFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let acc = 0;
  while (n > 0) {
    acc += n;
    n -= 1;
  }
  return acc;
};

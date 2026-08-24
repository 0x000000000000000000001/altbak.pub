export const runTCOFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let acc = 0;
  let i = n;
  while (i > 0) {
    acc += i;
    i--;
  }
  return acc;
};

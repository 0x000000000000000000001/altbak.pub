export const runChurchFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let acc = 0;
  for (let i = 0; i < n * 10000; i++) {
    acc++;
  }
  return acc;
};

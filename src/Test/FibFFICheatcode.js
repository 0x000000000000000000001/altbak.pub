export const runFibFFICheatcode = function(n) {
  if (n === 0) return 0;
  if (n === 1) return 1;
  return runFibFFICheatcode(n - 1) + runFibFFICheatcode(n - 2);
};

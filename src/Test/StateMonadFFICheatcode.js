export const runStateMonadFFICheatcode = function(limit) {
  let n = Math.floor(limit) * 20; // 60 * 20 = 1200
  let state = 0;
  for (let i = 0; i < n; i++) {
    state += 1;
  }
  return state;
};

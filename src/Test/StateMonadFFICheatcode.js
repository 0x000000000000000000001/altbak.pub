export const runStateMonadFFICheatcode = function(limit) {
  let state = 0;
  for (let i = 0; i < 60; i++) {
    for (let j = 0; j < 20; j++) {
      state += 1;
    }
  }
  return state;
};

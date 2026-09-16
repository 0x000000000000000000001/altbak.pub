export const runLazyEvaluationFFICheatcode = function(limit) {
  let n = Math.floor(limit) * 1000;
  let acc = 0;
  for (let i = 0; i < n; i++) {
    acc += 1;
  }
  return acc;
};

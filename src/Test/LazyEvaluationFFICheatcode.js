export const runLazyEvaluationFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  let acc = 0;
  for (let i = 0; i < n; i++) {
    acc += 1000;
  }
  return acc;
};

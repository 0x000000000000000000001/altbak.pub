export const runLazyEvaluationFFI = function(limit) {
  let n = Math.floor(limit);
  
  function force(lazy) {
    return lazy();
  }
  
  function defer(f) {
    return f;
  }
  
  function buildThunks(depth, acc) {
    if (depth === 0) return acc;
    return buildThunks(depth - 1, defer(() => force(acc) + 1));
  }
  
  function runManyTimes(times, acc) {
    let currTimes = times;
    let currAcc = acc;
    while (currTimes > 0) {
      currAcc += force(buildThunks(1000, defer(() => 0)));
      currTimes--;
    }
    return currAcc;
  }
  
  return runManyTimes(n, 0);
};

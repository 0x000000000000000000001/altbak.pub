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
    if (times === 0) return acc;
    return runManyTimes(times - 1, acc + force(buildThunks(1000, defer(() => 0))));
  }
  
  return runManyTimes(n, 0);
};

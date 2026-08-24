export const runStateMonadFFI = function(limit) {
  let dummy = Math.floor(limit);
  return runManyTimes(dummy, 0);
};

function runState(state) {
  return function(s) {
    return state(s);
  };
}

function bindState(state) {
  return function(g) {
    return function(s) {
      let r1 = state(s);
      let gPrime = g(r1.val);
      return gPrime(r1.state);
    };
  };
}

function pureState(a) {
  return function(s) {
    return { val: a, state: s };
  };
}

function get(s) {
  return { val: s, state: s };
}

function put(s) {
  return function(_s) {
    return { val: null, state: s };
  };
}

function modify(f) {
  return bindState(get)(function(s) {
    return put(f(s));
  });
}

function chainModifications(n) {
  if (n === 0) return pureState(null);
  return bindState(modify(function(x) { return x + 1; }))(function(_ignore) {
    return chainModifications(n - 1);
  });
}

function runManyTimes(n, acc) {
  if (n === 0) return acc;
  return runManyTimes(n - 1, acc + runState(chainModifications(60))(0).state);
}

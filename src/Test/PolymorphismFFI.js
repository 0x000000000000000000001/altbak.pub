export const runPolymorphismFFI = function(limit) {
  let dummy = Math.floor(limit);
  return polyLoop(intMonoidish)(dummy)(0);
};

const intMonoidish = {
  mempty_: 1,
  mappend_: function(x) {
    return function(y) {
      return x + y;
    };
  }
};

function polyLoop(dict) {
  return function(n_init) {
    return function(acc_init) {
      function go(n, acc) {
        if (n === 0) return acc;
        return go(n - 1, dict.mappend_(acc)(dict.mempty_));
      }
      return go(n_init, acc_init);
    };
  };
}

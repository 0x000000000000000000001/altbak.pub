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
        let currN = n;
        let currAcc = acc;
        while (currN > 0) {
          currAcc = dict.mappend_(currAcc)(dict.mempty_);
          currN--;
        }
        return currAcc;
      }
      return go(n_init, acc_init);
    };
  };
}

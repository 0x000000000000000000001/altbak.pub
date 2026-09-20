const intMonoidish = {
    mempty_: 1,
    mappend_: function(x) {
      return function(y) {
        return x + y;
      };
    }
};

function polyLoop(dict, n, acc) {
  while (n !== 0) {
    acc = dict.mappend_(acc)(dict.mempty_);
    n--;
  }
  return acc;
}

export const runPolymorphismFFI = function(limit) {
  return polyLoop(intMonoidish, Math.floor(limit), 0);
};

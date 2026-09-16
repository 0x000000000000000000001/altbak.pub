export const runPolymorphismFFI = function(limit) {
  let n = Math.floor(limit);
  let dict = {
    mempty_: 1,
    mappend_: function(x) {
      return function(y) {
        return x + y;
      };
    }
  };
  let acc = 0;
  while (n > 0) {
    acc = dict.mappend_(acc)(dict.mempty_);
    n--;
  }
  return acc;
};

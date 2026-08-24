export const runChurchFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  return unchurch(add(church(n))(church(n)));
};

function church(n) {
  return function(f) {
    return function(x) {
      let res = x;
      for (let i = 0; i < n; i++) {
        res = f(res);
      }
      return res;
    };
  };
}

function add(n) {
  return function(m) {
    return function(f) {
      return function(x) {
        return m(f)(n(f)(x));
      };
    };
  };
}

function unchurch(n) {
  return n(function(x) { return x + 1; })(0);
}

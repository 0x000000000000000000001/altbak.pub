export const runChurchFFI = function(limit) {
  let dummy = Math.floor(limit);
  return toInt(c100k(dummy));
};

function zeroC(f) {
  return function(x) {
    return x;
  };
}

function succC(n) {
  return function(f) {
    return function(x) {
      return f(n(f)(x));
    };
  };
}

function addC(m) {
  return function(n) {
    return function(f) {
      return function(x) {
        return m(f)(n(f)(x));
      };
    };
  };
}

function mulC(m) {
  return function(n) {
    return function(f) {
      return function(x) {
        return m(n(f))(x);
      };
    };
  };
}

function fromInt(n) {
  if (n === 0) return zeroC;
  return succC(fromInt(n - 1));
}

function toInt(n) {
  return n(function(x) { return x + 1; })(0);
}

function c10(n) {
  return fromInt(n);
}

function c100(n) {
  return mulC(c10(n))(c10(n));
}

function c10k(n) {
  return mulC(c100(n))(c100(n));
}

function c100k(n) {
  return mulC(c10k(n))(c10(n));
}

export const runPrimesFFI = function(limit) {
  let n = Math.floor(limit);
  return sumList(sieve(range(2)(n)));
};

function range(start) {
  return function(end) {
    function go(curr, acc) {
      if (curr < start) return acc;
      return go(curr - 1, { type: "Cons", value0: curr, value1: acc });
    }
    return go(end, { type: "Nil" });
  };
}

function filter(p) {
  return function(lst) {
    function go(list, acc) {
      if (list.type === "Nil") {
        function rev(l, a) {
          if (l.type === "Nil") return a;
          return rev(l.value1, { type: "Cons", value0: l.value0, value1: a });
        }
        return rev(acc, { type: "Nil" });
      }
      let x = list.value0;
      let xs = list.value1;
      if (p(x)) {
        return go(xs, { type: "Cons", value0: x, value1: acc });
      } else {
        return go(xs, acc);
      }
    }
    return go(lst, { type: "Nil" });
  };
}

function sieve(lst) {
  if (lst.type === "Nil") return { type: "Nil" };
  let p = lst.value0;
  let xs = lst.value1;
  return {
    type: "Cons",
    value0: p,
    value1: sieve(filter(function(x) { return x % p !== 0; })(xs))
  };
}

function sumList(lst) {
  function go(list, acc) {
    if (list.type === "Nil") return acc;
    return go(list.value1, acc + list.value0);
  }
  return go(lst, 0);
}

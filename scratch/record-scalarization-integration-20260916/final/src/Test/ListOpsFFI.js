export const runListOpsFFI = function(limit) {
  let n = Math.floor(limit);
  return foldl(function(acc) {
    return function(x) {
      return acc + x;
    };
  })(0)(filterEvens(range(1)(n)));
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

function filterEvens(lst) {
  function go(list, acc) {
    if (list.type === "Nil") return acc;
    let x = list.value0;
    let xs = list.value1;
    if (x % 2 === 0) {
      return go(xs, { type: "Cons", value0: x, value1: acc });
    } else {
      return go(xs, acc);
    }
  }
  return go(lst, { type: "Nil" });
}

function foldl(f) {
  return function(acc) {
    return function(lst) {
      function go(list, a) {
        if (list.type === "Nil") return a;
        return go(list.value1, f(a)(list.value0));
      }
      return go(lst, acc);
    };
  };
}

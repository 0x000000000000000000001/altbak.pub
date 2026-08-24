export const runPrimesFFICheatcode = function(limit) {
  let dummy = Math.floor(limit);
  return sumList(sieve(range(2, dummy)));
};

function range(start, end) {
  function go(curr, acc) {
    if (curr < start) return acc;
    return go(curr - 1, { type: "Cons", value0: curr, value1: acc });
  }
  return go(end, { type: "Nil" });
}

function filter(p, lst) {
  function go(list, acc) {
    if (list.type === "Nil") return reverse(acc);
    let x = list.value0;
    let xs = list.value1;
    if (p(x)) {
      return go(xs, { type: "Cons", value0: x, value1: acc });
    } else {
      return go(xs, acc);
    }
  }
  return go(lst, { type: "Nil" });
}

function reverse(lst) {
  function go(list, acc) {
    if (list.type === "Nil") return acc;
    return go(list.value1, { type: "Cons", value0: list.value0, value1: acc });
  }
  return go(lst, { type: "Nil" });
}

function sieve(lst) {
  if (lst.type === "Nil") return { type: "Nil" };
  let p = lst.value0;
  let xs = lst.value1;
  return {
    type: "Cons",
    value0: p,
    value1: sieve(filter((x) => x % p !== 0, xs))
  };
}

function sumList(lst) {
  function go(list, acc) {
    if (list.type === "Nil") return acc;
    return go(list.value1, acc + list.value0);
  }
  return go(lst, 0);
}

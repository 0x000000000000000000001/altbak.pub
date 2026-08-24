export const runRecordsFFI = function(limit) {
  let dummy = Math.floor(limit);
  let initial = { a: 0, b: { c: 0, d: { e: 0, f: 0 } } };
  return updateRec(dummy)(initial).b.d.f;
};

function updateRec(n) {
  return function(r) {
    if (n === 0) return r;
    
    // PureScript compiles record updates into Object.assign.
    // For a naive JS mirroring, we do immutable structural updates.
    let newD = {
      e: r.b.d.e + 3,
      f: r.b.d.f + (n % 5)
    };
    let newB = {
      c: r.b.c + 2,
      d: newD
    };
    let newR = {
      a: r.a + 1,
      b: newB
    };
    
    return updateRec(n - 1)(newR);
  };
}

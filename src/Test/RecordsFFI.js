export const runRecordsFFI = function(limit) {
  let dummy = Math.floor(limit);
  let initial = { a: 0, b: { c: 0, d: { e: 0, f: 0 } } };
  return updateRec(dummy, initial).b.d.f;
};

function updateRec(n, r) {
  let currN = n;
  let currR = r;
  
  while (currN > 0) {
    let newD = {
      e: currR.b.d.e + 3,
      f: currR.b.d.f + (currN % 5)
    };
    let newB = {
      c: currR.b.c + 2,
      d: newD
    };
    currR = {
      a: currR.a + 1,
      b: newB
    };
    
    currN = currN - 1;
  }
  
  return currR;
}

export const runRecordsFFI = function(limit) {
  let n = Math.floor(limit);
  let rec = { a: { b: { c: { d: { e: 0 } } } } };
  for (let i = 0; i < n; i++) {
    rec = {
      a: {
        b: {
          c: {
            d: {
              e: rec.a.b.c.d.e + 1
            }
          }
        }
      }
    };
  }
  return rec.a.b.c.d.e;
};

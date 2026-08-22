export const runPolymorphismFFI = function(limit) {
  let n = Math.floor(limit);
  let sum = 0;
  for (let i = 0; i < n; i++) {
    sum += computeLength("hello") + computeLength([1, 2, 3]);
  }
  return sum;
};

function computeLength(x) {
  return x.length;
}

// Imperative native implementation; preparation is outside the timed action.
const weightedChecksum = values => {
  let total = 0;
  for (let i = 0; i < values.length; i++) total += (1 + i % 2) * values[i];
  return total;
};
export const makeAction = size => () => {
  const left = Array.from({ length: size }, (_, index) => index + 1);
  const right = Array.from({ length: size }, (_, index) => size + index + 1);
  return count => {
    const result = new Array(2 * count);
    for (let i = 0; i < count; i++) { result[i] = left[i]; result[count + i] = right[i]; }
    return weightedChecksum(result);
  };
};

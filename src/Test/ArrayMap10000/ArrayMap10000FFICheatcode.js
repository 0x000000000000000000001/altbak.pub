// Imperative native implementation; preparation is outside the timed action.
const weightedChecksum = values => {
  let total = 0;
  for (let i = 0; i < values.length; i++) total += (1 + i % 2) * values[i];
  return total;
};
export const makeAction = size => () => {
  const input = Array.from({ length: size }, (_, index) => index + 1);
  return count => {
    const result = new Array(count);
    for (let i = 0; i < count; i++) result[i] = input[i] + 1;
    return weightedChecksum(result);
  };
};

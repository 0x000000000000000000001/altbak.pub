// Imperative native implementation; preparation is outside the timed action.
const weightedChecksum = values => {
  let total = 0;
  for (let i = 0; i < values.length; i++) total += (1 + i % 2) * values[i];
  return total;
};
export const makeAction = size => () => {
  const input = Array.from({ length: size }, (_, index) => size - index);
  return count => {
    const result = input.slice(0, count);
    result.sort((left, right) => left - right);
    return weightedChecksum(result);
  };
};

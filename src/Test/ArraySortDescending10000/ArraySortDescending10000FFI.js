// FP-style native implementation; preparation is outside the timed action.
const weightedChecksum = values => values.reduce(
  (total, value, index) => total + (1 + index % 2) * value, 0);
export const makeAction = size => () => {
  const input = Array.from({ length: size }, (_, index) => size - index);
  return count => {
    return weightedChecksum(input.slice(0, count).sort((left, right) => left - right));
  };
};

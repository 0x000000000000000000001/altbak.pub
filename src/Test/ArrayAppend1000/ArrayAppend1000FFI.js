// FP-style native implementation; preparation is outside the timed action.
const weightedChecksum = values => values.reduce(
  (total, value, index) => total + (1 + index % 2) * value, 0);
const prefix = (values, count) => count === values.length ? values : values.slice(0, count);
export const makeAction = size => () => {
  const left = Array.from({ length: size }, (_, index) => index + 1);
  const right = Array.from({ length: size }, (_, index) => size + index + 1);
  return count => {
    return weightedChecksum(prefix(left, count).concat(prefix(right, count)));
  };
};

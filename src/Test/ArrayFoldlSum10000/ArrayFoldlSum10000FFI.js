// FP-style native implementation; preparation is outside the timed action.
const prefix = (values, count) => count === values.length ? values : values.slice(0, count);
export const makeAction = size => () => {
  const input = Array.from({ length: size }, (_, index) => index + 1);
  return count => {
    return prefix(input, count).reduce((total, value) => total + value, 0);
  };
};

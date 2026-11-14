// FP-style native implementation; preparation is outside the timed action.
const prefix = (values, count) => count === values.length ? values : values.slice(0, count);
export const makeAction = size => () => {
  const left = Array.from({ length: size }, (_, index) => index + 1);
  const right = Array.from({ length: size }, (_, index) => index + 1);
  right[size - 1] = size + 1;
  return count => {
    return prefix(left, count).every((value, index) => value === right[index]) ? 1 : 0;
  };
};

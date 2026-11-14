// FP-style native implementation; preparation is outside the timed action.
const checksum = text => Array.from(text, char => char.codePointAt(0))
  .reduce((total, point) => total + point, 0);
const prefix = (values, count) => count === values.length ? values : values.slice(0, count);
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "ab");
  return count => {
    return checksum(prefix(input, count).reduce((text, part) => text + part, ""));
  };
};

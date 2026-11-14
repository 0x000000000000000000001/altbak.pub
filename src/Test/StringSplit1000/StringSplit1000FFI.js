// FP-style native implementation; preparation is outside the timed action.
const checksum = text => Array.from(text, char => char.codePointAt(0))
  .reduce((total, point) => total + point, 0);
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "abc").join(",");
  return count => {
    const text = count === size ? input : input.slice(0, count * 4 - 1);
    const parts = text.split(",");
    return parts.reduce((total, part) => total + checksum(part), parts.length);
  };
};

// FP-style native implementation; preparation is outside the timed action.
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "a😀é").join("");
  return count => {
    const text = count === size ? input : input.slice(0, count * 4);
    return Array.from(text, char => char.codePointAt(0)).length;
  };
};

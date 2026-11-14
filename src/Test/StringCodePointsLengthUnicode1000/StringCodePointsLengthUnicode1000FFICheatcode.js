// Imperative native implementation; preparation is outside the timed action.
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "a😀é").join("");
  return count => {
    const text = count === size ? input : input.slice(0, count * 4);
    let length = 0;
    for (const _char of text) length++;
    return length;
  };
};

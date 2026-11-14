// Imperative native implementation; preparation is outside the timed action.
const checksum = text => {
  let total = 0;
  for (const char of text) total += char.codePointAt(0);
  return total;
};
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "abc").join(",");
  return count => {
    const text = count === size ? input : input.slice(0, count * 4 - 1);
    const parts = text.split(",");
    let total = parts.length;
    for (let i = 0; i < parts.length; i++) total += checksum(parts[i]);
    return total;
  };
};

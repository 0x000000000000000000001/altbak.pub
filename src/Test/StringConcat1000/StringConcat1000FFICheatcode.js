// Imperative native implementation; preparation is outside the timed action.
const checksum = text => {
  let total = 0;
  for (const char of text) total += char.codePointAt(0);
  return total;
};
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "ab");
  return count => {
    let result = "";
    for (let i = 0; i < count; i++) result += input[i];
    return checksum(result);
  };
};

// Imperative native implementation; preparation is outside the timed action.
const checksum = text => {
  let total = 0;
  for (const char of text) total += char.codePointAt(0);
  return total;
};
export const makeAction = size => () => {
  const input = Array.from({ length: size }, () => "a😀é").join("");
  return count => {
    const take = count + Math.floor(count / 2);
    let result = "";
    const iterator = input[Symbol.iterator]();
    for (let i = 0; i < take; i++) {
      const next = iterator.next();
      if (next.done) break;
      result += next.value;
    }
    return checksum(result);
  };
};

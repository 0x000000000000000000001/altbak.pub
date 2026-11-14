// FP-style native implementation; preparation is outside the timed action.
const checksum = text => Array.from(text, char => char.codePointAt(0))
  .reduce((total, point) => total + point, 0);
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

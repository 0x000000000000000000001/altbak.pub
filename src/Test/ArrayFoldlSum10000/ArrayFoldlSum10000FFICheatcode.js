// Imperative native implementation; preparation is outside the timed action.
export const makeAction = size => () => {
  const input = Array.from({ length: size }, (_, index) => index + 1);
  return count => {
    let total = 0;
    for (let i = 0; i < count; i++) total += input[i];
    return total;
  };
};

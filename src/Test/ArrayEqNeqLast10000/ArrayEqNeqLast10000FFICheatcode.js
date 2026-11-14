// Imperative native implementation; preparation is outside the timed action.
export const makeAction = size => () => {
  const left = Array.from({ length: size }, (_, index) => index + 1);
  const right = Array.from({ length: size }, (_, index) => index + 1);
  right[size - 1] = size + 1;
  return count => {
    for (let i = 0; i < count; i++) if (left[i] !== right[i]) return 0;
    return 1;
  };
};

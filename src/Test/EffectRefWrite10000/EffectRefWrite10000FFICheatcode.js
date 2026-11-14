// Imperative native implementation; preparation is outside the timed action.
export const makeAction = size => () => {
  return count => {
    const ref = { value: 0 };
    for (let i = 1; i <= count; i++) ref.value = i;
    return ref.value;
  };
};

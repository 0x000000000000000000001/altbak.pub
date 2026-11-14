// Imperative native implementation; preparation is outside the timed action.
export const makeAction = size => () => {
  return count => {
    const ref = { value: 0 };
    for (let i = 0; i < count; i++) ref.value++;
    return ref.value;
  };
};

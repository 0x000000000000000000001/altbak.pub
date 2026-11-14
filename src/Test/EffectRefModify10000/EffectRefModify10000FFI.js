// FP-style native implementation; preparation is outside the timed action.
const modifyRef = (callback, ref) => { ref.value = callback(ref.value); };
export const makeAction = size => () => {
  return count => {
    const ref = { value: 0 };
    for (let i = 0; i < count; i++) modifyRef(value => value + 1, ref);
    return ref.value;
  };
};

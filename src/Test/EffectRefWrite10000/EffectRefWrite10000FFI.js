// FP-style native implementation; preparation is outside the timed action.
const writeRef = (value, ref) => { ref.value = value; };
export const makeAction = size => () => {
  return count => {
    const ref = { value: 0 };
    for (let i = 1; i <= count; i++) writeRef(i, ref);
    return ref.value;
  };
};

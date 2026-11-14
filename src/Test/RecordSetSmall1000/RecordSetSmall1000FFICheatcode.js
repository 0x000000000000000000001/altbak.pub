// Imperative native implementation; preparation is outside the timed action.
const sumRecord = value => value.a + value.b;
export const makeAction = size => () => {
  const input = { a: 1, b: 2 };
  return count => {
    const result = { ...input };
    for (let i = 0; i < count; i++) result.b++;
    return sumRecord(result) + sumRecord(input);
  };
};

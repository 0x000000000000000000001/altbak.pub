// FP-style native implementation; preparation is outside the timed action.
const sumRecord = value => value.a + value.b;
export const makeAction = size => () => {
  const input = { a: 1, b: 2 };
  return count => {
    let result = input;
    for (let i = 0; i < count; i++) result = { ...result, b: result.b + 1 };
    return sumRecord(result) + sumRecord(input);
  };
};

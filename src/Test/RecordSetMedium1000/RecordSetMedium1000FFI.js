// FP-style native implementation; preparation is outside the timed action.
const sumRecord = value => value.a + value.b + value.c + value.d + value.e + value.f + value.g + value.h;
export const makeAction = size => () => {
  const input = { a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8 };
  return count => {
    let result = input;
    for (let i = 0; i < count; i++) result = { ...result, h: result.h + 1 };
    return sumRecord(result) + sumRecord(input);
  };
};

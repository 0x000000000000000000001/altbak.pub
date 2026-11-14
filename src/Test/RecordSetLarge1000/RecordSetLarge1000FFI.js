// FP-style native implementation; preparation is outside the timed action.
const sumRecord = value => value.a + value.b + value.c + value.d + value.e + value.f + value.g + value.h + value.i + value.j + value.k + value.l + value.m + value.n + value.o + value.p;
export const makeAction = size => () => {
  const input = { a: 1, b: 2, c: 3, d: 4, e: 5, f: 6, g: 7, h: 8, i: 9, j: 10, k: 11, l: 12, m: 13, n: 14, o: 15, p: 16 };
  return count => {
    let result = input;
    for (let i = 0; i < count; i++) result = { ...result, p: result.p + 1 };
    return sumRecord(result) + sumRecord(input);
  };
};

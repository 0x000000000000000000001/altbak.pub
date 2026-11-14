// FP-style native implementation; preparation is outside the timed action.
const weightedChecksum = values => values.reduce(
  (total, value, index) => total + (1 + index % 2) * value, 0);
const prefix = (values, count) => count === values.length ? values : values.slice(0, count);
// Preserve the Array Traversable split/concat structure and eager callbacks.
const just = value => ({ value });
const mapMaybe = (f, value) => value === null ? null : just(f(value.value));
const applyMaybe = (f, value) => f === null || value === null ? null : just(f.value(value.value));
const traverseMaybe = (values, callback) => {
  const go = (bottom, top) => {
    switch (top - bottom) {
      case 0: return just([]);
      case 1: return mapMaybe(a => [a], callback(values[bottom]));
      case 2: return applyMaybe(mapMaybe(a => b => [a, b], callback(values[bottom])), callback(values[bottom + 1]));
      case 3: return applyMaybe(applyMaybe(mapMaybe(a => b => c => [a, b, c], callback(values[bottom])), callback(values[bottom + 1])), callback(values[bottom + 2]));
      default: {
        const pivot = bottom + Math.floor((top - bottom) / 4) * 2;
        return applyMaybe(mapMaybe(xs => ys => xs.concat(ys), go(bottom, pivot)), go(pivot, top));
      }
    }
  };
  return go(0, values.length);
};
export const makeAction = size => () => {
  const input = Array.from({ length: size }, (_, index) => index + 1);
  return count => {
    const rejected = null;
    const result = traverseMaybe(prefix(input, count), value => value === rejected ? null : just(value + 1));
    return result === null ? -1 : weightedChecksum(result.value);
  };
};

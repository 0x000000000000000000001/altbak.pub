// performance.now is monotonic and expressed in milliseconds. Do not fall back
// to the adjustable wall clock when the host lacks this API.
const clock = globalThis.performance;
if (!clock || typeof clock.now !== "function") {
  throw new Error("Benchmark requires a monotonic performance.now clock");
}
const origin = clock.now();
export const benchNow = () => (clock.now() - origin) * 1000;
// This is the language-level Effect boundary; JS has no portable no-inline API.
const inputCell = new Int32Array(new SharedArrayBuffer(4));
const resultCell = new Int32Array(new SharedArrayBuffer(4));
export const opaque = (a) => () => {
  if (typeof a !== "number" || !Number.isInteger(a)) return a;
  Atomics.store(inputCell, 0, a);
  return Atomics.load(inputCell, 0);
};
export const formatNumber = (n) => n.toFixed(6);
export const measureBatch = (iterations) => (expected) => (act) => () => {
  let result = 0;
  const start = benchNow();
  for (let i = 0; i < iterations; i++) {
    result = act();
    Atomics.store(resultCell, 0, result);
  }
  const elapsed = benchNow() - start;
  if (result !== expected) throw new Error(`Unstable benchmark result: ${result} != ${expected}`);
  return elapsed / iterations;
};

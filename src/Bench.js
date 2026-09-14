// performance.now is monotonic and expressed in milliseconds. Do not fall back
// to the adjustable wall clock when the host lacks this API.
const clock = globalThis.performance;
if (!clock || typeof clock.now !== "function") {
  throw new Error("Benchmark requires a monotonic performance.now clock");
}
const origin = clock.now();
export const benchNow = () => (clock.now() - origin) * 1000;
// This is the language-level Effect boundary; JS has no portable no-inline API.
export const opaque = (a) => () => a;
export const formatNumber = (n) => n.toFixed(2);

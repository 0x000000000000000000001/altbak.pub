export const consumeResult = expected => result => () => {
  globalThis.altbakExtendedResult = result;
  if (result !== expected) throw new Error("Unstable extended benchmark result");
};

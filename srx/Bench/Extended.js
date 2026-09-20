export const consumeResult = result => () => {
  globalThis.altbakExtendedResult = result;
};

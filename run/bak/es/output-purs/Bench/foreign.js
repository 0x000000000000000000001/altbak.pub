export const benchNow = function (unit) {
  return typeof performance !== "undefined" && performance.now
    ? performance.now() * 1000.0
    : Date.now() * 1000.0;
};
export const opaque = (a) => () => a;
export const formatNumber = (n) => n.toFixed(2);

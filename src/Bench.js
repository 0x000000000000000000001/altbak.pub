export const benchNow = function (unit) {
  return typeof performance !== "undefined" && performance.now
    ? performance.now()
    : Date.now();
};
export const opaque = (a) => () => a;
export const formatNumber = (n) => n.toFixed(2);

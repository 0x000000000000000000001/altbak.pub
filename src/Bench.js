export const benchNow = () => performance.now() * 1000.0;
export const opaque = (a) => () => a;
export const formatNumber = (n) => n.toFixed(2);
export const keepAlive = () => {
    setInterval(() => {}, 10000);
};

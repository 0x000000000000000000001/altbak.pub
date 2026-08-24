export const runAckermannFFICheatcode = function(ignore) {
  return ack(3, 4);
};

function ack(m, n) {
  if (m === 0) return n + 1;
  if (m > 0 && n === 0) return ack(m - 1, 1);
  return ack(m - 1, ack(m, n - 1));
}

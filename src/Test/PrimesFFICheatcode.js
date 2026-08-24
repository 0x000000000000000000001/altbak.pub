export const runPrimesFFICheatcode = function(limit) {
  let n = Math.floor(limit);
  if (n < 2) return 0;
  let sieve = new Uint8Array(n + 1);
  sieve.fill(1);
  let sum = 0;
  for (let p = 2; p * p <= n; p++) {
    if (sieve[p]) {
      for (let i = p * p; i <= n; i += p) sieve[i] = 0;
    }
  }
  for (let p = 2; p <= n; p++) {
    if (sieve[p]) sum += p;
  }
  return sum;
};

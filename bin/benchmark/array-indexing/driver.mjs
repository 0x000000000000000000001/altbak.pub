// Timing/input harness only: both indexed-read loops come from ArrayIndexing.purs.
import { nativeReads, boxedReads } from './output/Test.ArrayIndexing/index.js';

const [accesses = 1 << 23, batches = 10, seed = 5] = process.argv.slice(2).map(Number);
if (!Number.isInteger(accesses) || accesses < 16384 || accesses > 1 << 23 ||
    !Number.isInteger(batches) || batches < 3 || !Number.isInteger(seed) || seed < 0 || seed > 1000) {
  throw new Error('invalid accesses, batches or seed');
}
let sink = 0;
for (const size of [16, 1024, 16384]) {
  const source = Array.from({ length: size }, (_, i) => (i * 17 + seed * 31) % 251 + 1);
  const start = seed % size;
  const expected = source.reduce((sum, value) => sum + value, 0) * Math.floor(accesses / size) +
    Array.from({ length: accesses % size }, (_, i) => source[(start + i) % size]).reduce((sum, value) => sum + value, 0);
  for (const [representation, kernel] of [['native', nativeReads], ['boxed', boxedReads]]) {
    // JS has one array representation; these are the same two PureScript exports.
    const call = () => kernel(source)(size)(accesses)(start);
    for (let warm = 0; warm < 3; warm++) {
      sink = call();
      if (sink !== expected) throw new Error('warm-up checksum mismatch');
    }
    const nanoseconds = [];
    for (let batch = 0; batch < batches; batch++) {
      const begin = process.hrtime.bigint();
      sink = call();
      const elapsed = Number(process.hrtime.bigint() - begin);
      if (sink !== expected || elapsed <= 0) throw new Error('invalid measured result');
      nanoseconds.push(elapsed / accesses);
    }
    console.log(JSON.stringify({ runtime: 'js', representation, size, accesses, seed, checksum: sink,
      warmups: 3, ns_per_access: nanoseconds, bytes_per_access: null }));
  }
}

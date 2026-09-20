#!/usr/bin/env python3
"""Check timing boundaries, invocation counts, and rejection of wrong results."""
import json
from pathlib import Path
import shutil
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[1]
with tempfile.TemporaryDirectory(prefix='altbak-batch-contract-') as temporary:
    work = Path(temporary)
    shutil.copy2(ROOT / 'src/Bench.js', work / 'Bench.mjs')
    (work / 'check.mjs').write_text('''
import assert from 'node:assert/strict';
import {measureBatch, opaque} from './Bench.mjs';
let calls = 0;
const action = measureBatch(17)(42)(() => { calls++; return opaque(42)(); });
assert.equal(calls, 0, 'constructing the Effect must not run the calculation');
const first = action();
assert.equal(calls, 17);
assert(Number.isFinite(first) && first > 0);
action();
assert.equal(calls, 34, 'reusing the Effect must recompute every invocation');
assert.throws(() => measureBatch(3)(42)(() => 43)(), /Unstable benchmark result/);
assert.equal(opaque(-2147483648)(), -2147483648);
assert.equal(opaque(2147483647)(), 2147483647);
console.log('JS batch: deferred execution, exact counts, recomputation, oracle rejection passed');
''')
    subprocess.run(['node', 'check.mjs'], cwd=work, check=True)
    php = shutil.which('php')
    if php:
        (work / 'check.php').write_text('<?php\n$exports = [];\n$bench = require ' +
            json.dumps(str(ROOT / 'src/Bench.php')) + ''';
$calls = 0;
$act = $bench['measureBatch'](17, 42, function() use (&$calls) { $calls++; return 42; });
if ($calls !== 0) throw new Exception('Effect ran during construction');
$elapsed = $act();
if ($calls !== 17 || !is_finite($elapsed) || $elapsed <= 0) throw new Exception('Invalid batch');
$act();
if ($calls !== 34) throw new Exception('Cached result');
try { $bench['measureBatch'](3, 42, fn() => 43)(); }
catch (RuntimeException $error) { echo "PHP batch contract passed\\n"; exit(0); }
throw new Exception('Wrong result accepted');
''')
        subprocess.run([php, 'check.php'], cwd=work, check=True)

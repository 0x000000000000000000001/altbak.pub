#!/usr/bin/env python3
"""Check exact PHP FFI sources against shared independent values, without timing."""
from pathlib import Path
import json
import os
import runpy
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[3]
ORACLE = runpy.run_path(str(ROOT / 'test/native/oracle.py'))
CASES = ORACLE['cases']()
lines = ['<?php', 'error_reporting(E_ALL);', 'set_error_handler(function($level, $message, $file, $line) { throw new ErrorException($message, 0, $level, $file, $line); });', '$kernels = [];']
for name in ORACLE['NOMINAL']:
    for suffix in ['FFI', 'FFICheatcode']:
        stem = name + suffix
        path = ROOT / f'src/Test/{stem}.php'
        lines += ['$exports = [];', f'$module = require {json.dumps(str(path))};', f'$kernels[{json.dumps(stem)}] = $module[{json.dumps("run"+stem)}];']
for name, argument, expected in CASES:
    for suffix in ['FFI', 'FFICheatcode']:
        stem = name + suffix
        lines += [f'$actual = $kernels[{json.dumps(stem)}]({argument});',
                  f'if ($actual !== {expected}) throw new RuntimeException({json.dumps(stem+"("+str(argument)+"): expected "+str(expected)+", got ")} . var_export($actual, true));']
# Repeated invocations are essential: named functions declared inside callbacks
# used to fail on the second call, already during the benchmark warmup.
for stem in ['AstTreeFFICheatcode', 'FibFFICheatcode', 'AckermannFFICheatcode', 'RBTreeFFICheatcode']:
    name = stem.replace('FFICheatcode', '')
    argument, expected = ORACLE['NOMINAL'][name]
    for _ in range(2):
        lines.append(f'if ($kernels[{json.dumps(stem)}]({argument}) !== {expected}) throw new RuntimeException("repeat invocation failed");')
lines += ['$exports = [];', f'$bench = require {json.dumps(str(ROOT / "src/Bench.php"))};',
          '$before = hrtime(true) / 1000.0;', '$clock = $bench["benchNow"]();', '$after = hrtime(true) / 1000.0;',
          'if ($clock < $before - 1.0 || $clock > $after + 1.0) throw new RuntimeException("clock must use monotonic microseconds");',
          'foreach ([123, 1.25, "ok"] as $value) if ($bench["opaque"]($value)() !== $value) throw new RuntimeException("opaque changed its input");',
          f'echo "PASS native PHP: {2*len(CASES)} values, repeated calls, clock/opaque contracts; no benchmark timing\\n";']
with tempfile.TemporaryDirectory(prefix='altbak-native-php-') as directory:
    check = Path(directory) / 'check.php'
    check.write_text('\n'.join(lines)+'\n')
    subprocess.run([os.environ.get('PHP', 'php'), str(check)], check=True)

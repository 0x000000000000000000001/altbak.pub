#!/usr/bin/env python3
"""Check the actual Phpurs FFI adapter and generated core actions, without benchmarking."""
import argparse
import base64
import json
from pathlib import Path
import runpy
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[3]
NOMINAL = runpy.run_path(str(ROOT / 'test/native/oracle.py'))['NOMINAL']


def check(directory, mode, php):
    entry = {'pure': 'App', 'ffi': 'AppFFI', 'fficc': 'AppFFICheatcode', 'x': 'AppX'}[mode]
    main = directory / 'output' / entry / 'main.mod.php'
    source = main.read_text()
    call = f"$GLOBALS['{entry}_main']();"
    if source.count(call) != 1:
        raise RuntimeError(f'Cannot identify entry call in {main}')
    # Load dependencies in the compiler's order, stopping before benchmark main.
    bootstrap = source.split(call, 1)[0]
    encoded = base64.b64encode(str(main.parent).encode()).decode()
    bootstrap = bootstrap.replace('__DIR__', f"base64_decode('{encoded}')")
    bootstrap = bootstrap.replace('<?php', '''<?php
error_reporting(E_ALL);
set_error_handler(function($level, $message, $file, $line) {
    throw new ErrorException($message, 0, $level, $file, $line);
});''', 1)
    probes = [r'''
$calls = 0;
$action = function() use (&$calls) { $calls++; return 7; };
$batch = \Bench\majBench_measuremajBatch(3, 7, $action);
if ($calls !== 0) throw new RuntimeException('constructing a batch executed its Effect');
$elapsed = \Bench\phpurs_execute_effect($batch);
if (!is_float($elapsed) || !is_finite($elapsed) || $elapsed < 0 || $calls !== 3) {
    throw new RuntimeException('generated batch ABI: expected a float and exactly three action calls');
}
if ($GLOBALS['altbak_benchmark_result'] !== 7) throw new RuntimeException('batch result was not consumed');
\Bench\phpurs_execute_effect($batch);
if ($calls !== 6) throw new RuntimeException('reusing a batch must execute it again');
$partial = ($GLOBALS['Bench_measureBatch'])(2)(7)($action);
if ($calls !== 6) throw new RuntimeException('partial application executed its Effect');
$elapsed = \Bench\phpurs_execute_effect($partial);
if (!is_float($elapsed) || $calls !== 8) throw new RuntimeException('generated curried wrapper ABI');
''']
    if mode == 'x':
        probes.append(r'''
$GLOBALS['altbak_extended_result'] = null;
$consume = $GLOBALS['Bench_Extended_consumeResult'];
$effect = $consume('value', 'value');
if ($GLOBALS['altbak_extended_result'] !== null) throw new RuntimeException('extended sink ran before its Effect');
\Bench\phpurs_execute_effect($effect);
if ($GLOBALS['altbak_extended_result'] !== 'value') throw new RuntimeException('extended grouped wrapper ABI');
\Bench\phpurs_execute_effect($consume('other')('other'));
if ($GLOBALS['altbak_extended_result'] !== 'other') throw new RuntimeException('extended curried wrapper ABI');
$bad = $consume('expected', 'actual');
$rejected = false;
try { \Bench\phpurs_execute_effect($bad); } catch (RuntimeException $error) { $rejected = true; }
if (!$rejected) throw new RuntimeException('extended sink accepted an unstable result');
''')
        detail = 'extended String Effect ABI and unstable-result rejection'
    else:
        suffix = {'pure': '', 'ffi': 'FFI', 'fficc': 'FFICheatcode'}[mode]
        for name, (_argument, expected) in NOMINAL.items():
            key = f'Test_{name}{suffix}_act'
            probes += [f'$actual = \\Bench\\phpurs_execute_effect($GLOBALS[{json.dumps(key)}]);',
                       f'if ($actual !== {expected}) throw new RuntimeException({json.dumps(key)} . ": wrong result " . var_export($actual, true));']
        detail = '14 core results'
    probes.append(f'echo "PASS generated PHP {mode}: numeric Effect ABI, grouped/curried calls, sink, {detail}; no benchmark timing\\n";')
    with tempfile.TemporaryDirectory(prefix='altbak-generated-php-') as temporary:
        script = Path(temporary) / 'check.php'
        script.write_text(bootstrap + '\n' + '\n'.join(probes) + '\n')
        subprocess.run([php, '-d', 'xdebug.mode=off', str(script)], check=True)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('directory', type=Path)
    parser.add_argument('--mode', required=True, choices=['pure', 'ffi', 'fficc', 'x'])
    parser.add_argument('--php', default='php')
    args = parser.parse_args()
    check(args.directory.resolve(), args.mode, args.php)

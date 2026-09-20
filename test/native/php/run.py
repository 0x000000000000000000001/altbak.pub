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
PROBES = {
    'ListOpsFFI': ('    $rng = $range(1)($n);', '''
    $words = (object)['type' => 'Cons', 'value0' => 'ab', 'value1' => (object)['type' => 'Cons', 'value0' => 'c', 'value1' => (object)['type' => 'Nil']]];
    if ($foldl(function($n) { return function($s) use ($n) { return $n + strlen($s); }; }, 0, $words) !== 3) throw new RuntimeException('generic list fold');
    $ordered = $filter(function($x) { return $x % 2 === 0; })($range(1)(6));
    if ($foldl(function($s) { return function($x) use ($s) { return $s . $x; }; }, '', $ordered) !== '642') throw new RuntimeException('list filter order');
'''),
    'ArrayOpsFFI': ('    $rng = $range(1)($n);', '''
    $selected = $filter(function($s) { return strlen($s) === 1; })(['a', 'bb', 'c']);
    if ($foldl(function($n) { return function($s) use ($n) { return $n + strlen($s); }; }, 4, $selected) !== 6) throw new RuntimeException('generic array filter/fold');
'''),
    'PrimesFFI': ('    $rng = $range(2)($n);', '''
    $words = (object)['type' => 'Cons', 'value0' => 'a', 'value1' => (object)['type' => 'Cons', 'value0' => 'bb', 'value1' => (object)['type' => 'Cons', 'value0' => 'c', 'value1' => (object)['type' => 'Nil']]]];
    $selected = $filter(function($s) { return strlen($s) === 1; })($words);
    if ($selected->value0 !== 'a' || $selected->value1->value0 !== 'c' || $words->value1->value0 !== 'bb') throw new RuntimeException('generic persistent sieve filter');
'''),
    'PolymorphismFFI': ('    return $polyLoop($dict, $n, 0);', '''
    $strings = (object)['mempty_' => 'x', 'mappend_' => function($a) { return function($b) use ($a) { return $a . $b; }; }];
    if ($polyLoop($strings, 3, '!') !== '!xxx' || $polyLoop($strings, 0, '!') !== '!') throw new RuntimeException('generic dictionary');
'''),
    'ChurchFFI': ('    return $toInt($c100k((int)$limit));', '''
    $two = $succC($succC($zeroC));
    if ($mulC($two)($two)(function($s) { return $s . 'x'; })('!') !== '!xxxx') throw new RuntimeException('generic Church');
'''),
    'StateMonadFFI': ('    $total = 0;', '''
    $probe = $bind($get, function($s) use ($bind, $put, $pure) {
        return $bind($put($s . '!'), function($_) use ($pure, $s) { return $pure(strlen($s)); });
    });
    $result = $probe('ab');
    if ($result->value !== 2 || $result->state !== 'ab!') throw new RuntimeException('generic State');
'''),
    'RowToListFFI': ('    return $keys($dict, $record);', '''
    if ($keys($dictNil, (object)[]) !== 0 || $keys($dictCons($dictCons($dictNil)), (object)['name' => 'value', 'flag' => true]) !== 2) throw new RuntimeException('recursive row dictionary');
'''),
}
lines = ['<?php', 'error_reporting(E_ALL);', 'set_error_handler(function($level, $message, $file, $line) { throw new ErrorException($message, 0, $level, $file, $line); });', '$kernels = [];']
for name in ORACLE['NOMINAL']:
    for suffix in ['FFI', 'FFICheatcode']:
        stem = name + suffix
        path = ROOT / f'src/Test/{stem}.php'
        lines += ['$exports = [];', f'$module = require {json.dumps(str(path))};', f'$kernels[{json.dumps(stem)}] = $module[{json.dumps("run"+stem)}];']
        if stem in PROBES:
            # Instrument only this test copy, after validating the actual helper
            # definitions exist. No helper or test branch enters the benchmark.
            needle, probe = PROBES[stem]
            source = path.read_text()
            if source.count(needle) != 1:
                raise RuntimeError(f'Cannot locate generic helper probe in {path}')
            instrumented = source.replace(needle, probe + needle, 1).removeprefix('<?php')
            lines += ['$exports = [];', f'$probeModule = eval({json.dumps(instrumented)});',
                      f'$probeModule[{json.dumps("run"+stem)}](0);']
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
          f'echo "PASS native PHP: {2*len(CASES)} values, {len(PROBES)} generic/structural cases, repeated calls, clock/opaque contracts; no benchmark timing\\n";']
with tempfile.TemporaryDirectory(prefix='altbak-native-php-') as directory:
    check = Path(directory) / 'check.php'
    check.write_text('\n'.join(lines)+'\n')
    subprocess.run([os.environ.get('PHP', 'php'), str(check)], check=True)

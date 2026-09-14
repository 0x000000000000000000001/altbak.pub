#!/usr/bin/env python3
"""Check exact native F# modules in an isolated FSI process, without timing."""
from pathlib import Path
import os
import runpy
import shutil
import subprocess
import tempfile

ROOT = Path(__file__).resolve().parents[3]
ORACLE = runpy.run_path(str(ROOT / 'test/native/oracle.py'))
CASES = ORACLE['cases']()
dotnet = os.environ.get('DOTNET') or shutil.which('dotnet') or str(Path.home()/'.dotnet/dotnet')
lines = []
for name in ORACLE['NOMINAL']:
    for suffix in ['FFI', 'FFICheatcode']:
        lines.append('#load @"'+str(ROOT/f'src/Test/{name}{suffix}.fs')+'"')
lines += ['#load @"'+str(ROOT/'src/Bench.fs')+'"', 'let check label expected (actual: obj) =', '    let value = unbox<int> actual', '    if value <> expected then failwithf "%s: expected %d, got %d" label expected value']
for name, argument, expected in CASES:
    for suffix in ['FFI', 'FFICheatcode']:
        stem = name+suffix
        lines.append(f'check "{stem}({argument})" {expected} (Test.{stem}.run{stem} (box ({argument})))')
lines += ['let frequency = float System.Diagnostics.Stopwatch.Frequency',
          'let before = float (System.Diagnostics.Stopwatch.GetTimestamp()) / frequency * 1000000.0',
          'let clock = unbox<float> (Bench_FFI.benchNow null)',
          'let after = float (System.Diagnostics.Stopwatch.GetTimestamp()) / frequency * 1000000.0',
          'if clock < before - 1.0 || clock > after + 1.0 then failwith "clock must use monotonic microseconds"',
          'for value in [box 123; box 1.25; box "ok"] do',
          '    let action = unbox<obj -> obj> (Bench_FFI.opaque value)',
          '    if action null <> value then failwith "opaque changed its input"',
          f'printfn "PASS native F#: {2*len(CASES)} values and clock/opaque contracts; no benchmark timing"']
with tempfile.TemporaryDirectory(prefix='altbak-native-fsharp-') as directory:
    check = Path(directory)/'check.fsx'
    check.write_text('\n'.join(lines)+'\n')
    subprocess.run([dotnet, 'fsi', '--exec', str(check)], check=True)

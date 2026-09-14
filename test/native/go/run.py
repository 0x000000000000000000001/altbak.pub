#!/usr/bin/env python3
"""Compile exact native Go sources in a temporary module, without timing."""
import os
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True
sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from oracle import ROOT, NOMINAL, cases

with tempfile.TemporaryDirectory(prefix='altbak-native-go-') as temporary:
    folder = Path(temporary)
    (folder / 'go.mod').write_text('module nativecontracts\n\ngo 1.22\n')
    bench = folder / 'Bench'
    bench.mkdir()
    shutil.copy2(ROOT / 'src/Bench.go', bench / 'bench.go')
    lines = ['package nativecontracts', 'import ("testing"; Bench "nativecontracts/Bench"']
    for name in NOMINAL:
        for suffix in ['FFI', 'FFICheatcode']:
            module = name + suffix
            target = folder / module
            target.mkdir()
            shutil.copy2(ROOT / f'src/Test/{module}.go', target / 'kernel.go')
            lines.append(f'{module} "nativecontracts/{module}"')
    lines += [')', 'func TestNativeContracts(t *testing.T) {']
    count = 0
    for name, argument, expected in cases():
        for suffix in ['FFI', 'FFICheatcode']:
            module = name + suffix
            lines += [f't.Run("{module}_{argument}", func(t *testing.T) {{',
                      f'got := {module}.Run{module}({argument})',
                      f'if got != {expected} {{ t.Errorf("expected {expected}, got %d", got) }}', '})']
            count += 1
    lines += ['}', 'func TestClockAndBarrier(t *testing.T) {',
              'previous := Bench.BenchNow(); fractional := false',
              'for i:=0; i<100; i++ { current := Bench.BenchNow(); if current < previous { t.Fatal("clock went backwards") }; fractional = fractional || current != float64(int64(current)); previous=current }',
              'if !fractional { t.Fatal("clock truncated fractional microseconds") }',
              'if Bench.Opaque(123)().(int) != 123 { t.Fatal("opaque changed input") }', '}']
    (folder / 'contracts_test.go').write_text('\n'.join(lines) + '\n')
    environment = dict(os.environ, GOWORK='off', GOFLAGS='', GOCACHE=str(folder / 'cache'))
    subprocess.run(['go', 'test', '-count=1', '-pgo=off', '.'], cwd=folder, env=environment, check=True)
    print(f'PASS native Go: {count} values; no benchmark timing', flush=True)

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

PROBES = {
    'ListOps': '''
func TestGenericFoldAndOrder(t *testing.T) {
    words := ListOpsCons[string]{"a", ListOpsCons[string]{"bc", ListOpsNil[string]{}}}
    if got := foldl(func(n int) func(string) int { return func(s string) int { return n + len(s) } }, 0, words); got != 3 { t.Fatal(got) }
    if got := foldl(func(s string) func(int) string { return func(n int) string { return s + string(rune('0' + n)) } }, "", filterEvens(rangeListOps(1, 6))); got != "642" { t.Fatal(got) }
}''',
    'Primes': '''
func TestGenericFilterAndReverse(t *testing.T) {
    values := Cons[string]{"a", Cons[string]{"bb", Cons[string]{"c", Nil[string]{}}}}
    result := filter(func(s string) bool { return len(s) == 1 }, values).(Cons[string])
    if result.value0 != "a" || result.value1.(Cons[string]).value0 != "c" { t.Fatal(result) }
    if values.value1.(Cons[string]).value0 != "bb" { t.Fatal("input was mutated") }
}''',
    'ArrayOps': '''
func TestGenericArrayStages(t *testing.T) {
    values := []string{"a", "bb", "c"}
    selected := arrayFilter(func(s string) bool { return len(s) == 1 }, values)
    if got := arrayFoldl(func(n int) func(string) int { return func(s string) int { return n + len(s) } }, 4, selected); got != 6 { t.Fatal(got) }
    selected[0] = "changed"
    if values[0] != "a" { t.Fatal("filter must produce a fresh array") }
}''',
    'Polymorphism': '''
func TestDifferentMonoidish(t *testing.T) {
    dict := Monoidish[string]{Mempty: "x", Mappend: func(x string) func(string) string { return func(y string) string { return x + y } }}
    if got := polyLoop(dict, 3, "!"); got != "!xxx" { t.Fatal(got) }
    if got := polyLoop(dict, 0, "unchanged"); got != "unchanged" { t.Fatal(got) }
}''',
    'Church': '''
func TestGenericChurch(t *testing.T) {
    two := succC(succC(zeroC[string]()))
    if got := mulC(two, two)(func(s string) string { return s + "x" })("!"); got != "!xxxx" { t.Fatal(got) }
}''',
    'StateMonad': '''
func TestGenericState(t *testing.T) {
    action := bindState(get[string](), func(s string) State[string, int] {
        return bindState(put(s + "!"), func(_ struct{}) State[string, int] { return pureState[string](len(s)) })
    })
    got := runState(action, "ab")
    if got.val != 2 || got.state != "ab!" { t.Fatal(got) }
    if again := runState(action, "z"); again.val != 1 || again.state != "z!" { t.Fatal(again) }
}''',
    'LazyEvaluation': '''
func TestGenericNonMemoizingLazy(t *testing.T) {
    calls := 0
    thunk := deferFunc(func() string { calls++; return "value" })
    if force(thunk) != "value" || force(thunk) != "value" || calls != 2 { t.Fatal(calls) }
}''',
    'RowToList': '''
func TestTypedRowDictionary(t *testing.T) {
    if keys[rowNil](keysNil{}, rowNil{}) != 0 { t.Fatal("empty row") }
    type Tail = rowCons[bool, rowNil]
    type Row = rowCons[string, Tail]
    record := Row{"name", "value", Tail{"flag", true, rowNil{}}}
    dict := keysCons[string, Tail]{keysCons[bool, rowNil]{keysNil{}}}
    if keys[Row](dict, record) != 2 { t.Fatal("two-field row") }
}''',
}

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
            if suffix == 'FFI' and name in PROBES:
                (target / 'generic_test.go').write_text(
                    f'package Test_{module}\nimport "testing"\n' + PROBES[name] + '\n')
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
              'if Bench.Opaque(123)().(int) != 123 { t.Fatal("opaque changed input") }',
              'calls := 0; batch := Bench.MeasureBatch(3, 7, func() int { calls++; return 7 })',
              'if calls != 0 { t.Fatal("constructing a batch must not execute the Effect") }',
              'if elapsed := batch(); calls != 3 || elapsed < 0 { t.Fatal("batch did not execute exactly three calls") }', '}']
    (folder / 'contracts_test.go').write_text('\n'.join(lines) + '\n')
    environment = dict(os.environ, GOWORK='off', GOFLAGS='', GOCACHE=str(folder / 'cache'))
    subprocess.run(['go', 'test', '-count=1', '-pgo=off', './...'], cwd=folder, env=environment, check=True)
    print(f'PASS native Go: {count} values, {len(PROBES)} generic/structural cases; no benchmark timing', flush=True)

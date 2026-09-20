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
lines += '''
let words = Test.ListOpsFFI.Cons ("a", Test.ListOpsFFI.Cons ("b", Test.ListOpsFFI.Nil))
if Test.ListOpsFFI.foldl (+) "!" words <> "!ab" then failwith "generic list/fold"
let selected = Test.ListOpsFFI.filterEvens (Test.ListOpsFFI.range 1 6)
if Test.ListOpsFFI.foldl (fun acc x -> acc + string x) "" selected <> "642" then failwith "filter order"
let primesWords = Test.PrimesFFI.Cons ("a", Test.PrimesFFI.Cons ("skip", Test.PrimesFFI.Cons ("b", Test.PrimesFFI.Nil)))
let filteredWords = Test.PrimesFFI.filter ((<>) "skip") primesWords
if filteredWords <> Test.PrimesFFI.Cons ("a", Test.PrimesFFI.Cons ("b", Test.PrimesFFI.Nil)) then failwith "generic filter/reverse"
let dictionary : Test.PolymorphismFFI.Monoidish<string> = { mempty = "x"; mappend = (+) }
if Test.PolymorphismFFI.polyLoop dictionary 3 "s" <> "sxxx" then failwith "generic dictionary"
let two : Test.ChurchFFI.Church<string> = Test.ChurchFFI.successor (Test.ChurchFFI.successor Test.ChurchFFI.zero)
if Test.ChurchFFI.multiply two two (fun value -> value + "!") "s" <> "s!!!!" then failwith "generic Church"
let state : Test.StateMonadFFI.State<string, int> =
    Test.StateMonadFFI.bind Test.StateMonadFFI.get (fun initial ->
        Test.StateMonadFFI.bind (Test.StateMonadFFI.put (initial + "!")) (fun () -> Test.StateMonadFFI.pure' initial.Length))
let stateResult = Test.StateMonadFFI.runState state "abc"
if stateResult.value <> 3 || stateResult.state <> "abc!" then failwith "generic State"
let mutable forced = 0
let lazyWord = Test.LazyEvaluationFFI.defer (fun () -> forced <- forced + 1; "word")
if Test.LazyEvaluationFFI.force lazyWord <> "word" || Test.LazyEvaluationFFI.force lazyWord <> "word" || forced <> 2 then failwith "non-memoizing generic Lazy"
let empty = Test.RowToListFFI.RowNil
let single : Test.RowToListFFI.RowCons<string, Test.RowToListFFI.RowNil> = { Head = "one"; Tail = empty }
let pair : Test.RowToListFFI.RowCons<bool, _> = { Head = true; Tail = single }
if Test.RowToListFFI.keys Test.RowToListFFI.keysNil empty <> 0 then failwith "empty row"
if Test.RowToListFFI.keys (Test.RowToListFFI.keysCons Test.RowToListFFI.keysNil) single <> 1 then failwith "single row"
if Test.RowToListFFI.keys (Test.RowToListFFI.keysCons (Test.RowToListFFI.keysCons Test.RowToListFFI.keysNil)) pair <> 2 then failwith "heterogeneous row"
'''.strip().splitlines()
lines += ['let frequency = float System.Diagnostics.Stopwatch.Frequency',
          'let before = float (System.Diagnostics.Stopwatch.GetTimestamp()) / frequency * 1000000.0',
          'let clock = unbox<float> (Bench_FFI.benchNow null)',
          'let after = float (System.Diagnostics.Stopwatch.GetTimestamp()) / frequency * 1000000.0',
          'if clock < before - 1.0 || clock > after + 1.0 then failwith "clock must use monotonic microseconds"',
          'for value in [box 123; box 1.25; box "ok"] do',
          '    let action = unbox<obj -> obj> (Bench_FFI.opaque value)',
          '    if action null <> value then failwith "opaque changed its input"',
          f'printfn "PASS native F#: {2*len(CASES)} values, generic FP structures and clock/opaque contracts; no benchmark timing"']
with tempfile.TemporaryDirectory(prefix='altbak-native-fsharp-') as directory:
    check = Path(directory)/'check.fsx'
    check.write_text('\n'.join(lines)+'\n')
    subprocess.run([dotnet, 'fsi', '--exec', str(check)], check=True)

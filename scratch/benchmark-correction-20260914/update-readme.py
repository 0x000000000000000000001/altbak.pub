#!/usr/bin/env python3
"""Render only validated results.json measurements into the core README tables."""
import json
from pathlib import Path
import re

HERE=Path(__file__).resolve().parent
ROOT=HERE.parents[1]
RESULTS=json.loads((HERE/'results.json').read_text())
ROWS=['AST Evaluation','Fibonacci','List Processing','Tail Call Optimization','Deep Record Updates','Ackermann','Church Numerals','Prime Sieve','Red-Black Tree','Polymorphism','State Monad','Lazy Evaluation','Array Processing','RowToList']
GROUPS={
 'JavaScript':['js-pure','es-pure','js-ffi','js-fficc'],
 'Go':['go-pure','psgo-pure','go-ffi','go-fficc'],
 'Scheme':['scm-pure','scm-ffi','scm-fficc'],
 'Erlang':['erl-pure','erl-ffi','erl-fficc'],
 'PHP':['php-pure','php-ffi','php-fficc'],
 'Rust':['rust-pure','rust-ffi','rust-fficc'],
 'C++':['cpp-pure','cpp-ffi','cpp-fficc'],
 'F#/C#':['sharp-pure','sharp-ffi','sharp-fficc'],
 'Java':['java-pure','java-ffi','java-fficc']}
text=(ROOT/'README.md').read_text()
for title, keys in GROUPS.items():
 pattern=re.compile(r'(#### '+re.escape(title)+r'\n\n)([^\n]*Benchmark[^\n]*\n)(?:[^\n]*\|[^\n]*\n)+')
 match=pattern.search(text)
 if not match:raise RuntimeError('Missing table '+title)
 if any(len(RESULTS[key]['times_us'])!=14 for key in keys):raise RuntimeError('Incomplete '+title)
 lines=[match[1].rstrip(),'',match[2].strip(),' | '.join(['---']*(len(keys)+1))+' |']
 for index,row in enumerate(ROWS):
  lines.append(row+' | '+' | '.join(f"~ {RESULTS[key]['times_us'][index]:.2f} μs" for key in keys)+' |')
 lines.append('**Total Execution Time** | '+' | '.join(f"~ {RESULTS[key]['total_ms']:.2f} ms" for key in keys)+' |')
 text=text[:match.start()]+'\n'.join(lines)+'\n'+text[match.end():]
wasm=['#### WebAssembly GC','', 'Wasm Benchmark | Compiled Wasm GC (purs-wasm / Node) |','--- | --- |']
for index,row in enumerate(ROWS):wasm.append(f"{row} | ~ {RESULTS['wasm-pure']['times_us'][index]:.2f} μs |")
wasm.append(f"**Total Execution Time** | ~ {RESULTS['wasm-pure']['total_ms']:.2f} ms |")
text=text.replace('### Reading the three implementation columns','\n'.join(wasm)+'\n\n### Reading the three implementation columns',1)
anchor='Command: `./bin/run` (Runs all configured core backends, including Wasm). New tests will gradually be added.\n'
note='''
Core tables remeasured on 15 September 2026 after correcting the native workloads and runners. Each of the 30 columns passed all 14 expected outputs in three independent processes. Every process uses three global warmup suites, three warmups per test, then the best of ten measurements. Each table row is the median of those three process results; the total is the sum of its 14 displayed rows (before final rounding).

Rust uses O3/debug=false in all three columns (previous runner default: O1/debug=true); Go and psgo use GOGC=800 without PGO; C++ uses C++11/O3/NDEBUG without LTO. Exact runtimes, flags, source hashes and logs are archived in the [correction report](scratch/benchmark-correction-20260914/report.md), [validated measurements](scratch/benchmark-correction-20260914/results.json) and [reproduction commands](scratch/benchmark-correction-20260914/measurement-plan.json). The [previous README](scratch/benchmark-correction-20260914/previous-readme.md) is historical evidence; differences from invalid workloads or older profiles are not compiler speedups.
'''
text=text.replace(anchor,anchor+note,1)
(ROOT/'README.md').write_text(text)
# A compact total table is also kept in the French report.
rows=['## Totaux de la campagne validée','', '| Cible | Compilé | FP natif | Manuscrit natif |','|---|---:|---:|---:|']
for title,keys in GROUPS.items():
 compiled,fp,cc=keys[0],keys[-2],keys[-1]
 rows.append(f"| {title} | {RESULTS[compiled]['total_ms']:.2f} ms | {RESULTS[fp]['total_ms']:.2f} ms | {RESULTS[cc]['total_ms']:.2f} ms |")
rows.extend(['',f"Autres colonnes compilées : Arista ES **{RESULTS['es-pure']['total_ms']:.2f} ms**, psgo **{RESULTS['psgo-pure']['total_ms']:.2f} ms**, Wasm GC **{RESULTS['wasm-pure']['total_ms']:.2f} ms**.",''])
with (HERE/'report.md').open('a') as out:out.write('\n'+'\n'.join(rows))

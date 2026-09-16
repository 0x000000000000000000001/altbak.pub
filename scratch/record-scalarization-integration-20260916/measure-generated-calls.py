#!/usr/bin/env python3
"""Measure the actual compiler-generated scalar helpers, not handwritten workers."""
from pathlib import Path
import hashlib
import json
import random
import re
import statistics
import subprocess

BASE = Path(__file__).resolve().parent
ROOT = BASE.parents[1]
source = (BASE/'generated-call-fixture.rs').read_text()
deps = ROOT/'output/purust_output/target/release/deps'
mimalloc = next(deps.glob('libmimalloc-*.rlib'))
native = next((deps.parent/'build').glob('libmimalloc-sys-*/out/libmimalloc.a')).parent
harness = r'''
#[global_allocator] static BENCH_ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;
#[inline(never)]
fn reference_loop(mut n: i64, mut r: Value) -> Value {
    while n != 0 { r = ScalarRecords_nextFields(n, r); n -= 1; }
    r
}
fn main() {
    let args: Vec<String> = std::env::args().collect();
    let f: fn(i64, Value) -> Value = match args[1].as_str() {
        "reference" => reference_loop, "generated" => ScalarRecords_throughCall, _ => panic!()
    };
    let input = seed([0,0,0,0]);
    let oracle = expected(10000,[0,0,0,0]);
    for _ in 0..3 { drop(std::hint::black_box(f(std::hint::black_box(10000), std::hint::black_box(input.clone())))); }
    for _ in 0..10 {
        let start=std::time::Instant::now();
        let result=f(std::hint::black_box(10000),std::hint::black_box(input.clone()));
        let actual=std::hint::black_box(values(&result)); drop(result);
        let elapsed=start.elapsed().as_nanos();
        assert_eq!(actual,oracle); assert_eq!(values(&input),[0,0,0,0]);
        println!("{elapsed}");
    }
}
'''
commands = []
for mode in ['normal', 'noinline']:
    code = source
    if mode == 'noinline':
        code, count = re.subn(r'(?m)^(pub )?fn (ScalarRecords_nextFields(?:__purust_record_field_\d+_\d+)?\()',
                             r'#[inline(never)]\n\1fn \2', code)
        assert count == 5, count  # one public record helper + four generated scalar helpers
    file = BASE/f'calls-{mode}.rs'
    file.write_text(code+harness)
    command = ['rustc','--edition=2021','-Awarnings','-C','opt-level=3','-C','lto=off',str(file),
               '-L',f'dependency={deps}','-L',f'native={native}',
               '--extern',f'mimalloc={mimalloc}','-o',str(BASE/f'calls-{mode}')]
    subprocess.run(command,check=True,timeout=60)
    commands.append(command)

# All compilation has finished before any timed subprocess starts.
rng = random.Random(916265)
rows = []
cases = [(mode, name) for mode in ['normal','noinline'] for name in ['reference','generated']]
for rnd in range(21):
    order = cases.copy(); rng.shuffle(order); row = {'round':rnd,'order':order,'ns':{}}
    for mode, name in order:
        output = subprocess.check_output([str(BASE/f'calls-{mode}'),name],text=True,timeout=15)
        values_ns = [int(n) for n in output.splitlines()]
        assert len(values_ns) == 10 and min(values_ns) > 0
        row['ns'][mode+'/'+name] = values_ns
    rows.append(row)
summary = {}
for mode in ['normal','noinline']:
    refs = [min(r['ns'][mode+'/reference']) for r in rows]
    optimized = [min(r['ns'][mode+'/generated']) for r in rows]
    summary[mode] = {'reference_median_best_us':statistics.median(refs)/1000,
                     'generated_median_best_us':statistics.median(optimized)/1000,
                     'paired_change_pct':100*(statistics.median(b/a for a,b in zip(refs,optimized))-1),
                     'wins':sum(b<a for a,b in zip(refs,optimized))}
data = {'commands':commands,'source_sha256':hashlib.sha256(source.encode()).hexdigest(),
        'binary_sha256':{m:hashlib.sha256((BASE/f'calls-{m}').read_bytes()).hexdigest() for m in ['normal','noinline']},
        'rounds':rows,'summary':summary}
(BASE/'generated-calls-results.json').write_text(json.dumps(data,indent=2)+'\n')
print(json.dumps(summary,indent=2))

"""Extract emitted RBTree unchanged, then make only depth borrow its argument.

This is an experimental backend output transformation, not a TAST implementation.
Do not attribute this transformation to usageCount/escapes alone.
"""
from pathlib import Path
from collections import defaultdict
import hashlib
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
OLD = HERE.parent.parent / 'purust-fbip-20260916'
SOURCE = OLD / 'benchmark/output/purust_output/Purs_Test_RBTree/src/lib.rs'
source = SOURCE.read_text()
starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
excluded = {'R', 'B', 'E', 'T', 'describe', 'act'}
ranges = [(source.index('#[derive(Clone'), starts[0].start())]
for index, match in enumerate(starts):
    end = starts[index + 1].start() if index + 1 < len(starts) else len(source)
    if match[1].removeprefix('Test_RBTree_') not in excluded:
        ranges.append((match.start(), end))
lines = [''] * len(source.splitlines())
for start, end in ranges:
    line = source[:start].count('\n')
    part = source[start:end].splitlines()
    lines[line:line + len(part)] = part
baseline = '\n'.join(lines) + '\n'
(HERE / 'kernel-baseline.rs').write_text(baseline)

depth_start = baseline.index('pub fn Test_RBTree_depth(')
depth_end = baseline.index('pub fn Test_RBTree_balance(', depth_start)
depth = baseline[depth_start:depth_end]
assert depth.count('=> f.clone()') == 2
borrow_worker = depth.replace('Test_RBTree_depth', 'Test_RBTree_depth__borrow')
borrow_worker = borrow_worker.replace(
    'mut purs_local_0: std::rc::Rc<crate::Tree>',
    'purs_local_0: &std::rc::Rc<crate::Tree>')
borrow_worker = borrow_worker.replace('=> f.clone()', '=> f')
wrapper = '''// Owned ABI stays unchanged. The borrowed worker retains no input reference.
pub fn Test_RBTree_depth(purs_local_0: std::rc::Rc<crate::Tree>) -> i64 {
    Test_RBTree_depth__borrow(&purs_local_0)
}

'''
variant = baseline[:depth_start] + wrapper + borrow_worker + baseline[depth_end:]
assert baseline[:depth_start] == variant[:depth_start]
assert baseline[depth_end:] == variant[variant.index('pub fn Test_RBTree_balance('):]
(HERE / 'kernel-borrow.rs').write_text(variant)

# Counter runs are for logical operations and correctness, never timings.
(HERE / 'tracked_rc.rs').write_text((OLD / 'allocation-probe/tracked_rc.rs').read_text())
harness = (OLD / 'allocation-probe/count_harness.rs').read_text()
(HERE / 'count_harness.rs').write_text(harness)
results = {}
for name, kernel in [('baseline', baseline), ('borrow', variant)]:
    counted = kernel.replace('std::rc::Rc', 'tracked::Rc')
    counted = re.sub(r'^((?:pub )?fn (Test_RBTree_\w+)\([^\n]+?\{)',
        lambda match: match[1] + f' tracked::call("{match[2].removeprefix("Test_RBTree_")}");',
        counted, flags=re.M)
    (HERE / f'kernel-count-{name}.rs').write_text(counted)
    main = '#![allow(warnings)]\n#[path="tracked_rc.rs"] mod tracked;\n'
    main += f'include!("kernel-count-{name}.rs");\ninclude!("count_harness.rs");\n'
    (HERE / f'count-{name}.rs').write_text(main)
    binary = HERE / f'count-{name}'
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1',
        str(HERE / f'count-{name}.rs'), '-o', str(binary)], check=True)
    output = subprocess.check_output([str(binary)], text=True, timeout=120)
    (HERE / f'events-{name}.tsv').write_text(output)
    phases = defaultdict(lambda: defaultdict(int))
    for line in output.splitlines():
        phase, op, kind, site, count = line.split('\t')
        phases[phase][op] += int(count)
    totals = defaultdict(int)
    for phase in phases.values():
        for op, count in phase.items():
            totals[op] += count
    assert totals['new'] + totals['clone'] == sum(totals[op] for op in
        ['drop_last', 'drop_shared', 'unwrap_unique', 'unwrap_shared'])
    assert totals['new'] == totals['drop_last'] + totals['unwrap_unique']
    results[name] = dict(phases)
assert results['baseline']['unique_build'] == results['borrow']['unique_build']
assert results['baseline']['unique_depth_and_drop']['clone'] == 200000
assert results['borrow']['unique_depth_and_drop'].get('clone', 0) == 0
report = {
    'source': str(SOURCE),
    'source_sha256': hashlib.sha256(source.encode()).hexdigest(),
    'construction_change': 'none; all functions except depth byte-identical in extracted kernels',
    'transform': 'Original depth body kept; Rc input becomes borrowed; two recursive field clones become references; public owned wrapper retains original drop boundary.',
    'extra_fact_required': 'Transitive read-only, non-retaining argument contract, no FFI/effect/reference observation invalidating borrow; final backend analysis validates after optimization.',
    'validation': 'Both: 100000 nodes/depth22/BST/red-black invariants/four rotations/weak references/200 persistent versions/exact keys/balanced logical lifetimes passed',
    'timing': 'not run by this script',
    'counts': results,
}
(HERE / 'counts.json').write_text(json.dumps(report, indent=2) + '\n')
print(json.dumps(report, indent=2))

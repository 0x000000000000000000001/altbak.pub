"""Prepare optional mutable-slot insertion worker; run correctness, never timing."""
from pathlib import Path
from collections import defaultdict
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
worker = (HERE / 'ins_worker.rs').read_text()
results = {}
for name, original in [('borrow-ins', 'kernel-baseline.rs'),
                       ('borrow-ins-depth', 'kernel-borrow.rs')]:
    source = (HERE / original).read_text()
    # Keep the generated generic fallback verbatim, only rename its binding.
    # Its recursive invocations use the safe owned adapter above.
    assert source.count('pub fn Test_RBTree_ins(') == 1
    source = source.replace('pub fn Test_RBTree_ins(', 'pub fn Test_RBTree_ins__original(')
    kernel = source + '\n' + worker
    (HERE / f'kernel-{name}.rs').write_text(kernel)
    counted = kernel.replace('std::rc::Rc', 'tracked::Rc')
    counted = re.sub(r'^((?:pub )?fn (Test_RBTree_\w+)\([^\n]+?\{)',
        lambda match: match[1] + f' tracked::call("{match[2].removeprefix("Test_RBTree_")}");',
        counted, flags=re.M)
    (HERE / f'kernel-count-{name}.rs').write_text(counted)
    (HERE / f'count-{name}.rs').write_text('#![allow(warnings)]\n#[path="tracked_rc.rs"] mod tracked;\n'
        + f'include!("kernel-count-{name}.rs");\ninclude!("count_harness.rs");\n')
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
report = {
    'transform': 'Mutable-slot recursive ins; emitted read/rotation guards, field permutation, shared fallback and reuse worker retained. Owned public adapter unchanged.',
    'qualification': 'Backend output prototype, not automatic Haskell inference. Construction changes versus baseline; depth independently isolated/composed.',
    'extra_fact_required': 'Argument consumed and result returned to same slot; child access path exclusive under get_mut; non-retaining/reentrant behavior; sharing and weak fallback remain dynamic.',
    'validation': 'Both: 100000 nodes/depth22/BST/red-black invariants/four rotations/weak references/200 persistent versions/exact keys/balanced logical lifetimes passed',
    'timing': 'not run by this script',
    'counts': results,
}
(HERE / 'counts-ins.json').write_text(json.dumps(report, indent=2) + '\n')
print(json.dumps(report, indent=2))

"""Count fresh RBTree Rc operations separately from uninstrumented timings."""
from collections import defaultdict
import hashlib
import json
from pathlib import Path
import re
import subprocess

HERE = Path(__file__).resolve().parent
SOURCE = HERE.parent / 'benchmark/output/purust_output/Purs_Test_RBTree/src/lib.rs'
source = SOURCE.read_text()
starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
excluded = {'R', 'B', 'E', 'T', 'describe', 'act'}
ranges = [(source.index('#[derive(Clone'), starts[0].start())]
functions = []
for index, match in enumerate(starts):
    end = starts[index + 1].start() if index + 1 < len(starts) else len(source)
    name = match[1].removeprefix('Test_RBTree_')
    if name not in excluded:
        ranges.append((match.start(), end))
        functions.append((source[:match.start()].count('\n') + 1,
                          source[:end].count('\n') + 1, name))
lines = [''] * len(source.splitlines())
for start, end in ranges:
    line = source[:start].count('\n')
    part = source[start:end].splitlines()
    lines[line:line + len(part)] = part
kernel = '\n'.join(lines) + '\n'
kernel = kernel.replace('std::rc::Rc', 'tracked::Rc')
# Keep line numbers identical to the emitted Rust for operation attribution.
kernel = re.sub(r'^((?:pub )?fn (Test_RBTree_\w+)\([^\n]+?\{)',
                lambda match: match[1] + f' tracked::call("{match[2].removeprefix("Test_RBTree_")}");',
                kernel, flags=re.M)
(HERE / 'kernel.rs').write_text(kernel)
(HERE / 'main.rs').write_text('#![allow(warnings)]\n#[path="tracked_rc.rs"] mod tracked;\n'
                              'include!("kernel.rs");\ninclude!("count_harness.rs");\n')
binary = HERE / 'count'
subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(HERE / 'main.rs'),
                '-o', str(binary)], check=True)
output = subprocess.check_output([str(binary)], text=True, timeout=120)
(HERE / 'events.tsv').write_text(output)
phases = defaultdict(lambda: defaultdict(int))
by_function = defaultdict(lambda: defaultdict(lambda: defaultdict(int)))
events = []
for line in output.splitlines():
    phase, op, kind, site, count = line.split('\t')
    site, count = int(site), int(count)
    function = next((name for start, end, name in functions if start <= site < end), 'harness_or_drop')
    phases[phase][op] += count
    by_function[phase][function][op] += count
    events.append(dict(phase=phase, operation=op, kind=kind, line=site, count=count, function=function))
totals = defaultdict(int)
for phase in phases.values():
    for op, count in phase.items():
        totals[op] += count
assert totals['new'] + totals['clone'] == sum(totals[op] for op in
    ['drop_last', 'drop_shared', 'unwrap_unique', 'unwrap_shared']), dict(totals)
assert totals['new'] == totals['drop_last'] + totals['unwrap_unique'], dict(totals)
result = dict(source=str(SOURCE), source_sha256=hashlib.sha256(source.encode()).hexdigest(),
    method='Logical Rc operations using a one-word wrapper around native Rc; not a timing measurement.',
    validation='100000 nodes, depth 22, ordering/red-black invariants, four rotations, weak exclusion, '
               '200 retained exact-key snapshots and balanced reference/payload lifetimes passed.',
    phases=dict(phases), by_function=dict(by_function), events=events)
(HERE / 'results.json').write_text(json.dumps(result, indent=2) + '\n')
print(json.dumps({key: result[key] for key in ['source_sha256', 'validation', 'phases', 'by_function']}, indent=2))

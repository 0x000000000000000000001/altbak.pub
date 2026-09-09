"""Five alternating pairs of complete runner processes, with checked outputs."""
from pathlib import Path
import hashlib
import json
import re
import statistics
import subprocess
import sys

HERE = Path(__file__).resolve().parent
OUT = HERE / sys.argv[1] if len(sys.argv) > 1 else HERE
OUT.mkdir(exist_ok=True)
ROOT = HERE.parents[1]
GENERATED = ROOT / 'run/bak/rust/output/purust_output'
BINARIES = {'before': HERE / 'before-binary', 'after': HERE / 'after-binary'}
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000',
            '21536', '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)'
                     r'\s+\(Execution time - best of 10\)\s+([0-9.]+) μs')

def read_run(text):
    matches = PATTERN.findall(text)
    assert [value for _, value, _ in matches] == EXPECTED
    assert [name for name, _, _ in matches] == json.loads((HERE / 'metadata.json').read_text())['benchmark_names']
    return {name: float(us) for name, _, us in matches}

read_run((HERE / 'clean-run.log').read_text())
runs = {side: [] for side in BINARIES}
for pair in range(5):
    for side in (['before', 'after'] if pair % 2 == 0 else ['after', 'before']):
        text = subprocess.check_output([str(BINARIES[side])], text=True)
        (OUT / f'paired-{side}-{pair + 1}.log').write_text(text)
        values = read_run(text)
        runs[side].append(values)
        print(f'{pair + 1} {side}: {sum(values.values()) / 1000:.3f} ms; 14 outputs checked', flush=True)

medians = {side: {name: statistics.median(run[name] for run in values) for name in values[0]}
           for side, values in runs.items()}
result = {
    'method': '5 alternating process pairs; each process warms up and keeps best of 10 per benchmark; medians across 5 processes; totals sum per-benchmark medians; original mimalloc and release opt-level=1',
    'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
    'binary_sha256': {side: hashlib.sha256(path.read_bytes()).hexdigest() for side, path in BINARIES.items()},
    'runs_us': runs,
    'median_us': medians,
    'total_us': {side: sum(values.values()) for side, values in medians.items()},
}
(OUT / 'results.json').write_text(json.dumps(result, indent=2) + '\n')
for name in medians['before']:
    before, after = (medians[side][name] for side in BINARIES)
    print(name, before, '->', after)
print('Total:', result['total_us'])

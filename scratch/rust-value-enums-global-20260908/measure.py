"""Three serial runner processes per side, same profile and 14 checked outputs."""
from pathlib import Path
import hashlib
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
GENERATED = ROOT/'run/bak/rust/output/purust_output'
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000',
            '21536', '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)'
                     r'\s+\(Execution time - best of 10\)\s+([0-9.]+) μs')
assert [v for _, v, _ in PATTERN.findall((HERE/'clean-run.log').read_text())] == EXPECTED
results = {}
for side in ['before', 'after']:
    runs = []
    for index in range(1, 4):
        log = HERE/f'{side}-{index}.log'
        if side == 'after':
            print('After run', index, flush=True)
            run = subprocess.run([str(GENERATED/'target/release/purust_output')],
                                 stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True, check=True)
            log.write_text(run.stdout)
        matches = PATTERN.findall(log.read_text())
        assert [value for _, value, _ in matches] == EXPECTED, log
        runs.append({name: float(us) for name, _, us in matches})
    results[side] = {name: {'samples_us': [run[name] for run in runs],
                            'median_us': statistics.median(run[name] for run in runs)}
                     for name in runs[0]}
    results[side]['total_us'] = sum(item['median_us'] for item in results[side].values())
(HERE/'results.json').write_text(json.dumps(results, indent=2)+'\n')
metadata = {
    'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
    'method': 'Three serial processes before/after; median of per-benchmark best-of-ten; 14 outputs checked each run; mimalloc, release opt-level=1',
    'sha256': {name: hashlib.sha256(path.read_bytes()).hexdigest() for name, path in {
        'RBTree_before': HERE/'RBTree-before.rs',
        'RBTree_after': GENERATED/'Purs_Test_RBTree/src/lib.rs',
        'purust_before': HERE/'originals/bin/purust.js',
        'purust_after': ROOT/'../purust/purust/bin/purust.js',
    }.items()},
}
(HERE/'metadata.json').write_text(json.dumps(metadata, indent=2)+'\n')
for name in results['before']:
    before = results['before'][name] if name == 'total_us' else results['before'][name]['median_us']
    after = results['after'][name] if name == 'total_us' else results['after'][name]['median_us']
    print(name, before, '->', after, f'({(after/before-1)*100:+.1f}%)')

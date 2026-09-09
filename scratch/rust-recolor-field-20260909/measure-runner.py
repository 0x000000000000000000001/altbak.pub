"""Compare the complete baseline and integrated runners, without compilation."""
from pathlib import Path
import hashlib
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000',
            '21536', '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)'
                     r'\s+\(Execution time - best of 10\)\s+([0-9.]+) μs')
def parse(text):
    values = PATTERN.findall(text)
    assert [value for _, value, _ in values] == EXPECTED, values
    return {name: float(us) for name, _, us in values}

names = list(parse((HERE/'clean-run.log').read_text()))
runs = {side: [] for side in ['before', 'after']}
for pair in range(5):
    for side in (['before', 'after'] if pair % 2 == 0 else ['after', 'before']):
        output = subprocess.check_output([str(HERE/f'build/runner-{side}')], text=True)
        (HERE/f'build/paired-{side}-{pair+1}.log').write_text(output)
        values = parse(output)
        assert list(values) == names
        runs[side].append(values)
        print(pair+1, side, sum(values.values())/1000, 'ms; 14 outputs checked', flush=True)
medians = {side: {name: statistics.median(run[name] for run in values) for name in names}
           for side, values in runs.items()}
result = {
    'method': 'Five alternating pairs of full runner processes; original warmup and best of 10 '
              'per benchmark; medians across processes. Totals sum benchmark medians. O1/mimalloc; '
              'no concurrent build or instrumentation. All 14 outputs verified each run.',
    'binary_sha256': {side: hashlib.sha256((HERE/f'build/runner-{side}').read_bytes()).hexdigest() for side in runs},
    'runs_us': runs, 'median_us': medians,
    'total_us': {side: sum(values.values()) for side, values in medians.items()}}
(HERE/'runner-results.json').write_text(json.dumps(result, indent=2)+'\n')
for name in names: print(name, medians['before'][name], '->', medians['after'][name], flush=True)
print('Totals:', result['total_us'], flush=True)

"""Post-process existing timings only; never executes a benchmark or compiler."""
from pathlib import Path
import hashlib
import json
import statistics

HERE = Path(__file__).resolve().parent
source = HERE / 'timings.json'
timings = json.loads(source.read_text())
stages = list(timings['raw_ns'])
profiles = list(timings['raw_ns'][stages[0]])
blocks = {
    stage: {key: [statistics.median(row) for row in rows] for key, rows in cases.items()}
    for stage, cases in timings['raw_ns'].items()
}


def function_source(stage, function):
    text = (HERE / f'build/{stage}.rs').read_text()
    start = text.index('fn ' + function + '(')
    end = text.index('{', start) + 1
    depth = 1
    while depth:
        depth += (text[end] == '{') - (text[end] == '}')
        end += 1
    return text[start:end]


def comparison(before, after, key):
    left, right = blocks[before][key], blocks[after][key]
    return {
        'before_ms': statistics.median(left) / 1e6,
        'after_ms': statistics.median(right) / 1e6,
        'delta_percent_ratio_of_medians': (statistics.median(right) / statistics.median(left) - 1) * 100,
        'favorable_blocks': sum(b < a for a, b in zip(left, right)),
        'equal_blocks': sum(b == a for a, b in zip(left, right)),
        'blocks': len(left),
        'paired_block_deltas_percent': [(b / a - 1) * 100 for a, b in zip(left, right)],
    }


result = {
    'timings_sha256': hashlib.sha256(source.read_bytes()).hexdigest(),
    'method': 'Median of 3 per process, then median of 7 blocks; percent is ratio of these medians. Favorable blocks compare matched block medians.',
    'baseline': 'Synthetic naive IR, not current Purust output.',
    'adjacent': {},
    'final_vs_naive': {key: comparison(stages[0], stages[-1], key) for key in profiles},
}
for before, after in zip(stages, stages[1:]):
    result['adjacent'][after] = {
        'before': before,
        'profiles': {key: comparison(before, after, key) for key in profiles},
        'function_source_identical': {
            function: function_source(before, function) == function_source(after, function)
            for function in dict.fromkeys(key.split('/')[0] for key in profiles)
        },
    }
(HERE / 'timings-analysis.json').write_text(json.dumps(result, indent=2) + '\n')

modes = {'0': 'unique', '1': 'racine partagée', '2': 'mixte (enfant partagé)'}
lines = [
    '| Profil | naive | pushdown | precise | fusion | specialized_drop | reuse | retained_fields | Final / naive | Blocs favorables |',
    '| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |',
]
for key in profiles:
    function, mode = key.split('/')
    values = ' | '.join(f'{timings["median_ms"][stage][key]:.3f}' for stage in stages)
    delta = result['final_vs_naive'][key]
    lines.append(f'| `{function}` {modes[mode]} | {values} | {delta["delta_percent_ratio_of_medians"]:+.2f} % | {delta["favorable_blocks"]}/7 |')
lines += [
    '',
    'Écarts de chaque passage face au précédent. Chaque cellule indique **variation de durée / blocs favorables** ; un signe négatif signifie plus rapide. `=` indique que le corps Rust de cette fonction reste identique : son écart ne mesure pas une transformation locale de cette fonction.',
    '',
    '| Profil | pushdown | precise | fusion | specialized_drop | reuse | retained_fields |',
    '| --- | ---: | ---: | ---: | ---: | ---: | ---: |',
]
for key in profiles:
    function, mode = key.split('/')
    cells = []
    for stage in stages[1:]:
        data = result['adjacent'][stage]
        value = data['profiles'][key]
        marker = ' =' if data['function_source_identical'][function] else ''
        cells.append(f'{value["delta_percent_ratio_of_medians"]:+.2f} % / {value["favorable_blocks"]}{marker}')
    lines.append(f'| `{function}` {modes[mode]} | ' + ' | '.join(cells) + ' |')
(HERE / 'timing-tables.md').write_text('\n'.join(lines) + '\n')
print('Post-processing written: timings-analysis.json and timing-tables.md; no executions or clocks.')

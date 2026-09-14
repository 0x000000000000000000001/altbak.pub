"""Summarize the retained paired series without mixing their baselines."""
from pathlib import Path
import json
import statistics

HERE = Path(__file__).resolve().parent


def read(name):
    return json.loads((HERE / name).read_text())


def comparison(before, after):
    assert len(before) == len(after)
    b, a = statistics.median(before), statistics.median(after)
    return {'before_median': b, 'after_median': a, 'change_percent': 100 * (a / b - 1),
            'before_range': [min(before), max(before)], 'after_range': [min(after), max(after)],
            'favorable_blocks': sum(x < y for x, y in zip(after, before)), 'blocks': len(before)}


full = read('full-runner-results.json')
assert 'hybrid' in full['runs_us']
initial = read('full-runner-results-initial.json')
shared = read('b15/shared-timings.json')
sticky = read('sticky/timings.json')
tree = 'Red-Black Tree (100k Worst-Case Insertions):'
lists = 'List Processing (900 elements):'
before_tree = [row[tree] for row in full['runs_us']['before']]
result = {
    'full_tree_us': {name: comparison(before_tree, [row[tree] for row in full['runs_us'][name]])
                     for name in ['borrowed', 'consuming', 'hybrid']},
    'full_total_sum_of_medians_us': full['sum_of_benchmark_medians_us'],
    'list_us_initial_series': comparison([row[lists] for row in initial['runs_us']['before']],
                                       [row[lists] for row in initial['runs_us']['list_consuming']]),
    'retained_tree_ms': {name: comparison(shared['process_medians_ms']['before'], shared['process_medians_ms'][name])
                         for name in ['borrowed', 'consuming', 'hybrid']},
    'sticky_ms_per_million_iterations': {},
}
for scenario in sticky['summary_ns']:
    rows = sticky['rows']
    before = [row['elapsed_ns'] / 1e6 for row in rows if row['scenario'] == scenario and row['sticky'] == 'normal']
    after = [row['elapsed_ns'] / 1e6 for row in rows if row['scenario'] == scenario and row['sticky'] == 'forced-sticky']
    result['sticky_ms_per_million_iterations'][scenario] = comparison(before, after)
(HERE / 'summary.json').write_text(json.dumps(result, indent=2) + '\n')
print(json.dumps(result, indent=2))

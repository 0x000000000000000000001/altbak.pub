"""Summarize paired timings without pooling different experiments or sessions."""
from pathlib import Path
import json
import statistics

HERE = Path(__file__).resolve().parent


def read(name):
    return json.loads((HERE / name).read_text())


def compare(series, baseline):
    base = series[baseline]
    result = {}
    for name, values in series.items():
        assert len(values) == len(base)
        middle = statistics.median(values)
        result[name] = {
            'median': middle,
            'minimum': min(values), 'maximum': max(values),
            'delta_percent': 100 * (middle / statistics.median(base) - 1),
            'paired_favorable': sum(value < control for value, control in zip(values, base)),
            'blocks': len(values),
        }
    return result


def summarize():
    result = {}
    full = read('full-runner-results.json')
    result['full_runner_us'] = {
        label: compare({name: [row[label] for row in rows] for name, rows in full['runs_us'].items()}, 'before')
        for label in next(iter(full['runs_us'].values()))[0]
    }
    result['full_runner_observed_total_us'] = compare({name: [sum(row.values()) for row in rows] for name, rows in full['runs_us'].items()}, 'before')
    result['full_runner_sum_of_medians_us'] = full['sum_of_benchmark_medians_us']
    list_rows = read('list/sharing-timings.json')['runs']
    result['list_sharing_ns_per_operation'] = []
    for n in [900, 9000]:
        for scenario in ['unique', 'root-retained', 'interior-retained', 'weak']:
            series = {name: [statistics.median(row['samples_ns']) / row['iterations'] for row in sorted(list_rows, key=lambda row: row['round'])
                             if row['variant'] == name and row['n'] == n and row['scenario'] == scenario]
                      for name in ['baseline', 'lifetimes', 'consuming', 'reuse']}
            result['list_sharing_ns_per_operation'].append({'n': n, 'scenario': scenario, 'stages': compare(series, 'baseline')})
    fbip = read('fbip/timings.json')
    assert fbip['complete']
    result['fbip_ns'] = []
    rows = fbip['rows']
    keys = sorted({(row['scenario'], row['depth'], row['updates']) for row in rows})
    for scenario, depth, updates in keys:
        series = {name: [row['process_median_ns'] for row in sorted(rows, key=lambda row: row['block'])
                         if (row['scenario'], row['depth'], row['updates'], row['stage']) == (scenario, depth, updates, name)]
                  for name in ['persistent', 'reuse_rechecked', 'reuse_retained', 'fields_retained']}
        result['fbip_ns'].append({'scenario': scenario, 'depth': depth, 'updates': updates,
            'stages': compare(series, 'persistent'),
            'vs_reuse_rechecked': compare(series, 'reuse_rechecked'),
            'vs_reuse_retained': compare(series, 'reuse_retained')})
    rows = read('sticky/timings.json')['rows']
    result['sticky_ns'] = {
        scenario: compare({name: [row['elapsed_ns'] for row in sorted(rows, key=lambda row: row['block'])
                                 if row['scenario'] == scenario and row['mode'] == name]
                           for name in ['normal', 'checked', 'forced']}, 'normal')
        for scenario in dict.fromkeys(row['scenario'] for row in rows)
    }
    ir = read('ir/timings.json')
    stages = list(ir['raw_ns'])
    result['ir_batch_ns'] = {}
    for case in ir['raw_ns'][stages[0]]:
        series = {name: [statistics.median(samples) for samples in ir['raw_ns'][name][case]] for name in stages}
        result['ir_batch_ns'][case] = {'stages': compare(series, stages[0]),
            'adjacent': {name: compare({previous: series[previous], name: series[name]}, previous)[name]
                         for previous, name in zip(stages, stages[1:])}}
    (HERE / 'summary.json').write_text(json.dumps(result, indent=2) + '\n')
    return result


if __name__ == '__main__':
    print(json.dumps(summarize(), indent=2))

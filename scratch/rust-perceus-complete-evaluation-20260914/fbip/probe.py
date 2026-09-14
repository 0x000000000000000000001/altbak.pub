"""Synthetic, same-layout FBIP stages. Build/validate never read clocks."""
from pathlib import Path
import argparse
import hashlib
import json
import shutil
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
BUILD = HERE / 'build'
STAGES = ['persistent', 'reuse_rechecked', 'reuse_retained', 'fields_retained']
SCENARIOS = ['unique', 'old-root', 'mixed', 'diamond', 'weak-root', 'weak-mixed']


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def write_json(path, value):
    path.write_text(json.dumps(value, indent=2) + '\n')


def source_hashes():
    return {str(path.relative_to(HERE)): sha(path) for path in
            [HERE / 'src/main.rs', HERE / 'Cargo.toml', HERE / 'Cargo.lock', HERE / 'probe.py']}


def build():
    BUILD.mkdir(exist_ok=True)
    bins = {}
    for name in ['plain', 'counts']:
        command = ['cargo', 'build', '--release', '--offline', '--locked']
        if name == 'counts': command += ['--features', 'counts']
        with (BUILD / f'build-{name}.log').open('w') as log:
            subprocess.run(command, cwd=HERE, stdout=log, stderr=subprocess.STDOUT, check=True)
        binary = BUILD / name
        shutil.copy2(HERE / 'target/release/fbip_fields_probe', binary)
        bins[name] = {'path': str(binary), 'sha256': sha(binary), 'command': command}
    write_json(HERE / 'metadata.json', {
        'scope': 'Synthetic native two-child/two-i64 ADT; same representation and update plan, not compiler-emitted output.',
        'stages': STAGES, 'scenarios': SCENARIOS,
        'dimensions': {'depth': [7, 10], 'updates': [128, 2048]},
        'actions': 'Deterministic interleaving: leaf scalar add, two-scalar change on middle-depth node, child permutation on shallower node.',
        'profile': 'Cargo release opt-level=1 debug=true mimalloc, same Cargo.lock for all stages; counters compile only in counts executable.',
        'source_hashes': source_hashes(), 'binaries': bins,
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'counter_scope': 'Logical Rc cell allocations/payload destructions, Ptr clone/drop, uniqueness tests and payload operations. Empty cells are included. Weak control blocks are not separate allocation counters; all observers are dropped and strong payload expiration is checked.',
        'timing_scope': 'Plan and oracle construction outside timer. Native tree construction, ownership setup, updates, final checksum, all owners/weak observers destruction inside timer. One unmeasured verified warmup then seven samples, plain executable only.',
        'timings_run': False,
    })
    print('plain + counts compiled; no timing executed', flush=True)


def verify():
    metadata = json.loads((HERE / 'metadata.json').read_text())
    assert metadata['source_hashes'] == source_hashes(), 'source changed after build'
    for binary in metadata['binaries'].values():
        assert sha(Path(binary['path'])) == binary['sha256'], binary
    return metadata


def validate():
    metadata = verify()
    rows = {}
    for name in ['plain', 'counts']:
        output = subprocess.check_output([metadata['binaries'][name]['path'], 'validate'], text=True)
        (BUILD / f'validation-{name}.jsonl').write_text(output)
        rows[name] = [json.loads(line) for line in output.splitlines()]
        assert len(rows[name]) == 96
        print(name, '96/96 oracle and lifetime cases passed', flush=True)
    key = lambda row: tuple(row[field] for field in ['stage', 'scenario', 'depth', 'updates'])
    plain = {key(row): row for row in rows['plain']}
    for row in rows['counts']:
        assert row['checksum'] == plain[key(row)]['checksum']
        assert row['allocations'] == row['cell_drops']
        assert row['allocations'] + row['dups'] == row['drops']
        assert row['weak_created'] == row['weak_drops']
    write_json(HERE / 'validation.json', {
        'passed': True, 'cases_per_binary': 96, 'stage_count': 4, 'scenario_count': 6,
        'dimensions': metadata['dimensions'], 'plain_matches_instrumented': True,
        'oracle': 'Full structural equality with independent Box<Model> update implementation, including every retained old version/subtree; checked before final owner release.',
        'lifetimes': 'Every allocated payload, every strong pointer handle and every weak observer balanced. Weak::strong_count zero after final owners dropped.',
        'logs': ['build/validation-plain.jsonl', 'build/validation-counts.jsonl'],
    })
    write_json(HERE / 'counts.json', rows['counts'])


def measure(scenarios, depths, updates, blocks, authorized):
    assert authorized, 'time requires --coordinated'
    metadata = verify()
    collected = []
    orders = []
    completed = 0
    total = len(scenarios) * len(depths) * len(updates)
    def checkpoint():
        write_json(HERE / 'timings.json', {'orders': orders, 'rows': collected,
            'complete': completed == total, 'completed_combinations': completed,
            'total_combinations': total,
            'method': 'Rotations then reverse rotations, one unmeasured warmup and seven samples per separate process; median of process medians for comparisons. All stages same native representation, workload and allocator.',
            'metadata_sha256': sha(HERE / 'metadata.json')})
    for scenario in scenarios:
        for depth in depths:
            for n in updates:
                for block in range(blocks):
                    offset = block % len(STAGES)
                    order = STAGES[offset:] + STAGES[:offset]
                    if (block // len(STAGES)) % 2: order = list(reversed(order))
                    orders.append({'scenario': scenario, 'depth': depth, 'updates': n, 'block': block, 'order': order})
                    for stage in order:
                        command = [metadata['binaries']['plain']['path'], 'time', stage, scenario,
                                   str(depth), str(n), '--coordinated']
                        output = subprocess.check_output(command, text=True)
                        rows = [json.loads(line) for line in output.splitlines()]
                        assert len(rows) == 7 and len({row['checksum'] for row in rows}) == 1
                        collected.append({'stage': stage, 'scenario': scenario, 'depth': depth, 'updates': n,
                                          'block': block, 'samples_ns': [row['elapsed_ns'] for row in rows],
                                          'checksum': rows[0]['checksum'],
                                          'process_median_ns': statistics.median(row['elapsed_ns'] for row in rows)})
                completed += 1
                checkpoint()
                print(f'FBIP {completed}/{total}: {scenario}, depth={depth}, updates={n}, {blocks} blocks saved', flush=True)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('mode', choices=['build', 'validate', 'time'])
    parser.add_argument('--coordinated', action='store_true')
    parser.add_argument('--scenario', action='append', choices=SCENARIOS)
    parser.add_argument('--depth', action='append', type=int)
    parser.add_argument('--updates', action='append', type=int)
    parser.add_argument('--blocks', type=int, default=8)
    args = parser.parse_args()
    if args.mode == 'build': build()
    elif args.mode == 'validate': validate()
    else: measure(args.scenario or SCENARIOS, args.depth or [7, 10], args.updates or [128, 2048], args.blocks, args.coordinated)

"""Ablate sticky checks and exercise high fan-out without modifying Purust."""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
HERE = Path(__file__).resolve().parent
BUILD = HERE / 'build'
ROOT = HERE.parents[2]
PREVIOUS = HERE.parents[1] / 'rust-perceus-evaluation-20260914/sticky/probe.py'
spec = importlib.util.spec_from_file_location('original_sticky_probe', PREVIOUS)
original = importlib.util.module_from_spec(spec)
spec.loader.exec_module(original)
original.HERE, original.BUILD = HERE, BUILD
MODES = ['normal', 'checked', 'forced']
SCENARIOS = ['read-owned', 'read-borrowed', 'closure-read', 'update-unique', 'update-retained', 'fanout']


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def run(command):
    result = subprocess.run(list(map(str, command)), capture_output=True, text=True)
    if result.returncode:
        raise RuntimeError(str(command) + '\n' + result.stdout + result.stderr)
    return result.stdout


def build():
    original.prepare()
    current = (BUILD / 'local.rs').read_text()
    assert current.count('if count != Self::STICKY_COUNT {') == 2
    checked = current.replace('if count != Self::STICKY_COUNT {', '{')
    checked = checked.replace('count.saturating_add(1)', 'count.checked_add(1).expect("reference count overflow in ablation")')
    (BUILD / 'checked.rs').write_text(checked)
    deps = original.GENERATED / 'target/release/deps'
    regex = list(deps.glob('libfancy_regex-*.rlib'))
    assert len(regex) == 1
    metadata = {'runtime_source_sha256': sha(original.RUNTIME / 'local.rs'),
        'scope': 'Same payload, u32 layout, getter/update code and checksummed harness; checked removes only sticky guards and uses checked_add. Forced uses the unmodified saturating runtime with three initial counters manually set to MAX.',
        'profile': 'O1; overflow checks enabled; default System allocator; matches the prior isolated sticky experiment, not the full mimalloc runner.', 'binaries': {}}
    for runtime in ['normal', 'checked']:
        for counters in [False, True]:
            name = runtime + ('-counts' if counters else '-plain')
            directory = BUILD / name
            directory.mkdir(exist_ok=True)
            flags = ['--edition=2021', '-C', 'opt-level=1', '-C', 'overflow-checks=yes', '--cfg', 'sticky_probe']
            if counters: flags += ['--cfg', 'sticky_counts']
            pointer = directory / 'libperceus_ptr.rlib'
            core = directory / 'libpurust_core.rlib'
            source = BUILD / ('local.rs' if runtime == 'normal' else 'checked.rs')
            run(['rustc', *flags, '--crate-name=perceus_ptr', '--crate-type=rlib', source, '-o', pointer])
            run(['rustc', *flags, '--crate-name=purust_core', '--crate-type=rlib', BUILD / 'purust_core.rs', '-o', core,
                 '-L', f'dependency={directory}', '-L', f'dependency={deps}', '--extern', f'perceus_ptr={pointer}', '--extern', f'fancy_regex={regex[0]}'])
            run(['rustc', *flags, HERE / 'harness.rs', '-o', directory / 'probe', '-L', f'dependency={directory}', '-L', f'dependency={deps}',
                 '--extern', f'perceus_ptr={pointer}', '--extern', f'purust_core={core}'])
            metadata['binaries'][name] = {'path': str(directory / 'probe'), 'sha256': sha(directory / 'probe')}
            print(name, 'compiled, no timing', flush=True)
    metadata['files'] = {str(path.relative_to(HERE)): sha(path) for path in [HERE / 'harness.rs', BUILD / 'local.rs', BUILD / 'checked.rs', BUILD / 'records.rs', BUILD / 'purust_core.rs']}
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2) + '\n')


def execute(mode, scenario, n, command):
    runtime = 'checked' if mode == 'checked' else 'normal'
    binary = BUILD / (runtime + ('-counts' if command == 'count' else '-plain')) / 'probe'
    saturation = 'forced-sticky' if mode == 'forced' else 'normal'
    row = json.loads(run([binary, command, scenario, saturation, n]))
    row['mode'] = mode
    return row


def checks():
    rows = []
    for scenario in SCENARIOS:
        n = 100000 if scenario == 'fanout' else 10000
        for mode in MODES:
            row = execute(mode, scenario, n, 'count')
            plain = execute(mode, scenario, n, 'check')
            assert plain['checksum'] == row['checksum'] and plain['elapsed_ns'] == 0
            counts = row['counts']
            assert counts['allocations'] - counts['frees'] == (3 if mode == 'forced' else 0)
            if mode != 'forced': assert counts['sticky_clones'] == counts['sticky_drops'] == 0
            if scenario == 'fanout' and mode != 'forced': assert counts['max_normal_owners'] == n + 1
            rows.append(row)
        assert rows[-1]['checksum'] == rows[-2]['checksum'] == rows[-3]['checksum']
        assert rows[-2]['counts'] == rows[-3]['counts'], 'Removing saturation must preserve all normal RC operations'
    (HERE / 'counts.json').write_text(json.dumps({'rows': rows, 'no_timing': True, 'counter_scope': 'PerceusPtr cells; Vec and Rc closure allocations excluded'}, indent=2) + '\n')
    print('18 cases validated; checked/current counts identical; high fan-out reaches 100001 owners', flush=True)


def time(authorized):
    assert authorized, 'Coordinate timing first'
    metadata = json.loads((HERE / 'metadata.json').read_text())
    for path, digest in metadata['files'].items(): assert sha(HERE / path) == digest
    for record in metadata['binaries'].values(): assert sha(Path(record['path'])) == record['sha256']
    rows = []
    orders = []
    for scenario in SCENARIOS:
        for block in range(6):
            shift = block % 3
            order = MODES[shift:] + MODES[:shift]
            if block // 3: order.reverse()
            orders.append({'scenario': scenario, 'block': block, 'order': order})
            for mode in order:
                row = execute(mode, scenario, 1000000, 'time')
                row['block'] = block
                rows.append(row)
        print(scenario, {mode: round(statistics.median(row['elapsed_ns'] for row in rows if row['scenario'] == scenario and row['mode'] == mode) / 1e6, 3) for mode in MODES}, flush=True)
    result = {'rows': rows, 'order': orders, 'summary_ms': {scenario: {mode: statistics.median(row['elapsed_ns'] for row in rows if row['scenario'] == scenario and row['mode'] == mode) / 1e6 for mode in MODES} for scenario in SCENARIOS},
        'scope': 'Six balanced blocks; 1M iterations or simultaneously retained owners; 10k warmup; input setup outside timer; final releases inside; no counters.'}
    (HERE / 'timings.json').write_text(json.dumps(result, indent=2) + '\n')


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('command', choices=['build', 'check', 'time'])
    parser.add_argument('--coordinated', action='store_true')
    args = parser.parse_args()
    if args.command == 'build': build()
    elif args.command == 'check': checks()
    else: time(args.coordinated)

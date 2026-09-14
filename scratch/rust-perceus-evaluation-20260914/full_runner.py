"""Full runners from one frozen generated tree; no writes to live output.

Config: {"before": {}, "candidate": {"RBTree": "/absolute/prototype.rs",
                                      "ListOps": "/absolute/prototype.rs"}}
Only replacement modules vary. App and root are copied into disjoint variant
workspaces; all other Cargo path dependencies point to the frozen snapshot.
"""
from pathlib import Path
import argparse
from contextlib import contextmanager
import fcntl
import hashlib
import json
import os
import re
import shutil
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
BUILD = HERE / 'build'
LIVE = HERE.parents[1] / 'run/bak/rust/output/purust_output'
FROZEN = BUILD / 'frozen'
WORKSPACES = BUILD / 'full-runners'
TARGET = BUILD / 'cargo-target'
BINARIES = BUILD / 'full-binaries'
CONFIG = HERE / 'full-runner-variants.json'
SNAPSHOT = HERE / 'full-runner-snapshot.json'
METADATA = HERE / 'full-runner-build.json'
BUILD_LOCK = BUILD / 'baseline-Cargo.lock'
LOCK_METADATA = HERE / 'full-runner-lock.json'
LOCAL = {'Purs_App', 'Purs_Test_RBTree', 'Purs_Test_ListOps'}
MODULES = {'RBTree': 'Purs_Test_RBTree', 'ListOps': 'Purs_Test_ListOps'}
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000',
            '21536', '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+'
                     r'\(Execution time - best of 10\)\s+([0-9.]+) μs')


def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def write_json(path, value):
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2) + '\n')


def source_files(root):
    for directory, dirs, files in os.walk(root):
        dirs[:] = sorted(name for name in dirs if name not in {'target', '.git'})
        for name in sorted(files):
            path = Path(directory) / name
            if path.suffix in {'.rs', '.toml'} or name == 'Cargo.lock':
                yield path


def inputs(root):
    return {str(path.relative_to(root)): digest(path) for path in source_files(root)}


def verify_frozen():
    metadata = json.loads(SNAPSHOT.read_text())
    assert inputs(FROZEN) == metadata['inputs'], 'frozen inputs changed'
    return metadata


def ensure_build_lock():
    if LOCK_METADATA.exists():
        metadata = json.loads(LOCK_METADATA.read_text())
        assert digest(BUILD_LOCK) == metadata['runner_lock_sha256']
        return metadata
    previous = HERE.parent / 'rust-field-permutation-integration-20260914/build/full-output/Cargo.lock'
    assert previous.is_file(), 'the previous paired runner lock is required once'
    def packages(path):
        return {re.search(r'(?m)^name = "([^"]+)"', block).group(1): block.strip()
                for block in path.read_text().split('[[package]]')[1:]}
    original, effective = packages(FROZEN / 'Cargo.lock'), packages(previous)
    assert all(original.get(name) == block for name, block in effective.items()), 'locked package changed'
    shutil.copy2(previous, BUILD_LOCK)
    metadata = {'source': str(previous), 'original_lock_sha256': digest(FROZEN / 'Cargo.lock'),
                'runner_lock_sha256': digest(BUILD_LOCK), 'original_packages': len(original),
                'runner_packages': len(effective), 'added_or_changed_packages': [],
                'note': 'Same exact effective Cargo.lock as the preceding full-runner experiment. Removed entries only: unrelated workspace packages; all remaining complete package blocks are identical to the live baseline lock.'}
    write_json(LOCK_METADATA, metadata)
    return metadata


@contextmanager
def exclusive_build():
    BUILD.mkdir(parents=True, exist_ok=True)
    with (BUILD / 'full-runner.lock').open('w') as lock:
        fcntl.flock(lock, fcntl.LOCK_EX | fcntl.LOCK_NB)
        yield


def cargo_paths(text, original_dir, destination):
    def replace(match):
        original = (original_dir / match[1]).resolve()
        assert original.is_relative_to(FROZEN), original
        relative = original.relative_to(FROZEN)
        assert len(relative.parts) == 1, original
        target = destination / relative if original.name in LOCAL else original
        return 'path = ' + json.dumps(str(target))
    return re.sub(r'path\s*=\s*"([^"]+)"', replace, text)


def verify_ffi():
    console = FROZEN / 'Purs_Effect_Console/src/lib.rs'
    bench = FROZEN / 'Purs_Bench/src/lib.rs'
    console_text, bench_text = console.read_text(), bench.read_text()
    assert re.search(r'pub fn Effect_Console_log\(message: String\) -> UnknownType\s*\{\s*'
                     r'purust_console_message\(message, false\)', console_text)
    assert 'SystemTime::now()' in bench_text and 'std::hint::black_box(a0.clone())' in bench_text
    assert re.search(r'pub fn Bench_benchNow\(\) -> UnknownType\s*\{\s*crate::Value::Func1', bench_text)
    return {'Console': {'sha256': digest(console), 'verified': 'real purust_console_message'},
            'Bench': {'sha256': digest(bench), 'verified': 'real SystemTime clock and black_box'}}


def snapshot():
    if SNAPSHOT.exists():
        metadata = verify_frozen()
        print('existing frozen snapshot verified:', len(metadata['inputs']), 'files', flush=True)
        return
    assert LIVE.is_dir(), LIVE
    assert not FROZEN.exists(), 'refuse to overwrite an unrecorded snapshot'
    before = inputs(LIVE)
    for relative in before:
        destination = FROZEN / relative
        destination.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(LIVE / relative, destination)
    assert before == inputs(LIVE) == inputs(FROZEN), 'live source changed during snapshot'
    # Generated FFI is inlined. Reject an unnoticed external Cargo/build input.
    for manifest in FROZEN.rglob('Cargo.toml'):
        for raw in re.findall(r'path\s*=\s*"([^"]+)"', manifest.read_text()):
            assert (manifest.parent / raw).resolve().is_relative_to(FROZEN), (manifest, raw)
    for source in FROZEN.rglob('*.rs'):
        assert not re.search(r'\binclude(?:_str|_bytes)?!\s*\(', source.read_text()), source
    profile_text = re.search(r'(?ms)^\[profile\.release\]\s*\n(.*?)(?=^\[|\Z)',
                             (FROZEN / 'Cargo.toml').read_text()).group(1)
    assert re.fullmatch(r'\s*debug\s*=\s*true\s*\nopt-level\s*=\s*1\s*', profile_text), profile_text
    profile = {'debug': True, 'opt-level': 1}
    metadata = {'live_source_read_only': str(LIVE), 'frozen': str(FROZEN), 'inputs': before,
                'profile': profile, 'lock_sha256': digest(FROZEN / 'Cargo.lock'),
                'ffi': verify_ffi(), 'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
                'cargo': subprocess.check_output(['cargo', '--version'], text=True).strip(),
                'note': 'Only .rs/.toml/Cargo.lock copied. No include macros or external local paths; inlined FFI needs no JavaScript files. Live output is never changed.'}
    write_json(SNAPSHOT, metadata)
    if not CONFIG.exists():
        write_json(CONFIG, {'before': {}})
    print('frozen snapshot:', len(before), 'files; Console/Bench FFI verified', flush=True)


def validate_name(name):
    assert re.fullmatch(r'[a-z][a-z0-9_-]*', name), name


def prepare_variant(name, replacements):
    validate_name(name)
    assert set(replacements) <= set(MODULES), replacements
    destination = WORKSPACES / name
    destination.mkdir(parents=True, exist_ok=True)
    manifest, count = re.subn(r'(?s)^\[workspace\].*?(?=\[package\])',
        '[workspace]\nmembers = ' + json.dumps(sorted(LOCAL)) + '\n\n',
        (FROZEN / 'Cargo.toml').read_text())
    assert count == 1
    (destination / 'Cargo.toml').write_text(cargo_paths(manifest, FROZEN, destination))
    shutil.copytree(FROZEN / 'src', destination / 'src', dirs_exist_ok=True)
    for crate in sorted(LOCAL):
        shutil.copytree(FROZEN / crate, destination / crate, dirs_exist_ok=True)
        manifest_path = destination / crate / 'Cargo.toml'
        manifest_path.write_text(cargo_paths(manifest_path.read_text(), FROZEN / crate, destination))
    chosen = {}
    for key, crate in MODULES.items():
        source = Path(replacements[key]).resolve() if key in replacements else FROZEN / crate / 'src/lib.rs'
        assert source.is_file(), source
        saved = destination / crate / 'src/lib.rs'
        shutil.copy2(source, saved)
        assert digest(source) == digest(saved), 'prototype changed during copy'
        chosen[key] = {'input': str(source), 'source_sha256': digest(source),
                       'saved': str(saved), 'saved_sha256': digest(saved),
                       'differs_from_baseline': digest(source) != digest(FROZEN / crate / 'src/lib.rs')}
    shutil.copy2(BUILD_LOCK, destination / 'Cargo.lock')
    return destination, chosen


def build_variant(name, replacements):
    """Build a named variant from {'RBTree': path, 'ListOps': path}; omitted = frozen."""
    verify_frozen()
    lock_metadata = ensure_build_lock()
    destination, chosen = prepare_variant(name, replacements)
    log_path = BUILD / f'full-build-{name}.log'
    env = {**os.environ, 'CARGO_TARGET_DIR': str(TARGET)}
    env.pop('RUSTFLAGS', None)
    env.pop('CARGO_ENCODED_RUSTFLAGS', None)
    with log_path.open('w') as log:
        # Lock resolution itself is read-only: same package versions/profile
        # as the live baseline, with only local manifest locations changed.
        subprocess.run(['cargo', 'metadata', '--offline', '--locked', '--format-version', '1'],
                       cwd=destination, env=env, stdout=log, stderr=subprocess.STDOUT, check=True)
        subprocess.run(['cargo', 'clean', '--release', '--offline', '--locked', '-p', 'Purs_Test_RBTree', '-p', 'Purs_Test_ListOps',
                        '-p', 'Purs_App', '-p', 'purust_output'],
                       cwd=destination, env=env, stdout=log, stderr=subprocess.STDOUT, check=True)
        subprocess.run(['cargo', 'build', '--release', '--offline', '--locked'], cwd=destination,
                       env=env, stdout=log, stderr=subprocess.STDOUT, check=True)
    assert digest(destination / 'Cargo.lock') == lock_metadata['runner_lock_sha256']
    BINARIES.mkdir(parents=True, exist_ok=True)
    binary = BINARIES / name
    shutil.copy2(TARGET / 'release/purust_output', binary)
    result = {'workspace': str(destination), 'replacements': chosen, 'inputs': inputs(destination),
              'lock_sha256': digest(destination / 'Cargo.lock'), 'binary': str(binary),
              'binary_sha256': digest(binary), 'log': str(log_path)}
    metadata = json.loads(METADATA.read_text()) if METADATA.exists() else {
        'snapshot_sha256': digest(SNAPSHOT), 'variants': {},
        'method': 'Disjoint root/App/RBTree/ListOps sources, frozen shared dependencies, four packages cleaned before each offline locked O1/debug=true/mimalloc build.'}
    assert metadata['snapshot_sha256'] == digest(SNAPSHOT)
    metadata['variants'][name] = result
    verify_frozen()
    write_json(METADATA, metadata)
    print(name, 'built; frozen dependencies and baseline lock preserved', flush=True)
    return result


def verify_variant(name):
    verify_frozen()
    metadata = json.loads(METADATA.read_text())
    assert metadata['snapshot_sha256'] == digest(SNAPSHOT)
    variant = metadata['variants'][name]
    assert inputs(Path(variant['workspace'])) == variant['inputs'], name + ': workspace changed'
    assert digest(Path(variant['binary'])) == variant['binary_sha256'], name + ': binary changed'
    return variant


def execute(name, log_path):
    variant = verify_variant(name)
    output = subprocess.check_output([variant['binary']], text=True)
    log_path.write_text(output)
    rows = PATTERN.findall(output)
    assert [value for _, value, _ in rows] == EXPECTED, rows
    assert len({label for label, _, _ in rows}) == 14
    return rows


def smoke(name):
    log = BUILD / f'full-smoke-{name}.log'
    rows = execute(name, log)
    write_json(BUILD / f'full-smoke-{name}.json', {'name': name, 'values': [value for _, value, _ in rows],
        'passed': True, 'log': str(log), 'note': 'Smoke verification only; reported times are not a performance experiment.'})
    print(name, 'smoke: 14/14 values verified (no performance conclusion)', flush=True)


def measure(names, blocks):
    assert len(names) >= 2 and len(set(names)) == len(names)
    runs, orders = {name: [] for name in names}, []
    for block in range(blocks):
        if len(names) == 2:
            order = names if block % 2 == 0 else list(reversed(names))
        else:
            shift = block % len(names)
            order = names[shift:] + names[:shift]
            if (block // len(names)) % 2: order = list(reversed(order))
        orders.append(order)
        for name in order:
            rows = execute(name, BUILD / f'full-timing-{block + 1}-{name}.log')
            runs[name].append({label.strip(): float(us) for label, _, us in rows})
            print(block + 1, name, '14/14 values verified', flush=True)
    labels = list(runs[names[0]][0])
    medians = {name: {label: statistics.median(row[label] for row in rows) for label in labels}
               for name, rows in runs.items()}
    result = {'method': 'Coordinated alternating/rotating full runners; warmup/best-of-10 preserved. Only consume these as performance results after other builds/tests/instrumentation have stopped.',
              'order': orders, 'runs_us': runs, 'median_us': medians,
              'sum_of_benchmark_medians_us': {name: sum(rows.values()) for name, rows in medians.items()},
              'median_observed_total_us': {name: statistics.median(sum(row.values()) for row in rows) for name, rows in runs.items()},
              'snapshot_sha256': digest(SNAPSHOT), 'build_metadata_sha256': digest(METADATA)}
    write_json(HERE / 'full-runner-results.json', result)
    print(json.dumps(result['sum_of_benchmark_medians_us'], indent=2))


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('mode', choices=['snapshot', 'build', 'smoke', 'time'])
    parser.add_argument('--config', type=Path, default=CONFIG)
    parser.add_argument('--variant', action='append', help='restrict selected variants; repeatable')
    parser.add_argument('--blocks', type=int, default=5)
    parser.add_argument('--coordinated', action='store_true', help='explicitly confirm other task builds/tests are stopped')
    args = parser.parse_args()
    with exclusive_build():
        if args.mode == 'snapshot':
            snapshot()
            return
        config = json.loads(args.config.read_text())
        names = args.variant or list(config)
        assert names and all(name in config for name in names), names
        if args.mode == 'build':
            for name in names: build_variant(name, config[name])
        elif args.mode == 'smoke':
            for name in names: smoke(name)
        else:
            assert args.coordinated, 'time requires --coordinated after all other work stops'
            measure(names, args.blocks)


if __name__ == '__main__':
    main()

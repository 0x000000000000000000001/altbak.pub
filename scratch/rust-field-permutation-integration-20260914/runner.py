"""Compare actual emitted RBTree modules with identical full-runner dependencies."""
from pathlib import Path
import argparse
import hashlib
import json
import re
import shutil
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
BUILD = HERE / 'build'
GENERATED = BUILD / 'after'
FULL = BUILD / 'full-output'
NAMES = ['before', 'after']
LOCAL = {'Purs_App', 'Purs_Test_RBTree'}
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000',
            '21536', '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+'
                     r'\(Execution time - best of 10\)\s+([0-9.]+) μs')


def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def write_json(path, value):
    path.write_text(json.dumps(value, indent=2) + '\n')


def source(name):
    return BUILD / ('disabled' if name == 'before' else name) / 'Purs_Test_RBTree/src/lib.rs'


def inputs():
    return {str(p.relative_to(GENERATED)): digest(p)
            for p in sorted(GENERATED.rglob('*'))
            if p.is_file() and p.suffix in {'.rs', '.toml'}}


def paths(manifest, original_dir):
    def replace(match):
        original = (original_dir / match[1]).resolve()
        target = FULL / original.name if original.name in LOCAL else original
        assert target.parent == FULL or original.is_relative_to(GENERATED)
        return 'path = ' + json.dumps(str(target))
    return re.sub(r'path\s*=\s*"([^"]+)"', replace, manifest)


def build():
    FULL.mkdir(exist_ok=True)
    manifest, count = re.subn(r'(?s)^\[workspace\].*?(?=\[package\])',
        '[workspace]\nmembers = ["Purs_App", "Purs_Test_RBTree"]\n\n',
        (GENERATED / 'Cargo.toml').read_text())
    assert count == 1
    (FULL / 'Cargo.toml').write_text(paths(manifest, GENERATED))
    shutil.copytree(GENERATED / 'src', FULL / 'src', dirs_exist_ok=True)
    for name in LOCAL:
        shutil.copytree(GENERATED / name, FULL / name, dirs_exist_ok=True)
        manifest_path = FULL / name / 'Cargo.toml'
        manifest_path.write_text(paths(manifest_path.read_text(), GENERATED / name))
    shutil.copy2(BUILD / 'baseline-Cargo.lock', FULL / 'Cargo.lock')
    with (BUILD / 'runner-resolve.log').open('w') as log:
        subprocess.run(['cargo', 'metadata', '--offline', '--format-version', '1'],
                       cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
    lock_hash = digest(FULL / 'Cargo.lock')
    metadata = {'inputs': inputs(), 'lock_sha256': lock_hash, 'variants': {},
        'method': 'Only RBTree differs in the two built runners. The control module is emitted with the new proof disabled and equals the saved original baseline. Both runners share the same freshly generated after dependencies and allocator lock.'}
    for name in NAMES:
        shutil.copy2(source(name), FULL / 'Purs_Test_RBTree/src/lib.rs')
        with (BUILD / f'runner-build-{name}.log').open('w') as log:
            subprocess.run(['cargo', 'clean', '-p', 'Purs_Test_RBTree', '-p', 'Purs_App', '-p', 'purust_output', '--release'],
                           cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
            subprocess.run(['cargo', 'build', '--release', '--offline', '--locked'],
                           cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
        binary = BUILD / f'paired-{name}'
        shutil.copy2(FULL / 'target/release/purust_output', binary)
        assert digest(FULL / 'Cargo.lock') == lock_hash
        metadata['variants'][name] = {'source_sha256': digest(source(name)), 'binary_sha256': digest(binary)}
        print(name, 'runner built', flush=True)
    assert inputs() == metadata['inputs']
    write_json(HERE / 'runner-build.json', metadata)


def measure():
    metadata = json.loads((HERE / 'runner-build.json').read_text())
    assert inputs() == metadata['inputs']
    assert digest(FULL / 'Cargo.lock') == metadata['lock_sha256']
    for name in NAMES:
        assert digest(source(name)) == metadata['variants'][name]['source_sha256']
        assert digest(BUILD / f'paired-{name}') == metadata['variants'][name]['binary_sha256']
    runs = {name: [] for name in NAMES}
    orders = []
    for block in range(5):
        order = NAMES if block % 2 == 0 else list(reversed(NAMES))
        orders.append(order)
        for name in order:
            output = subprocess.check_output([str(BUILD / f'paired-{name}')], text=True)
            (BUILD / f'runner-{block + 1}-{name}.log').write_text(output)
            rows = PATTERN.findall(output)
            assert [value for _, value, _ in rows] == EXPECTED, rows
            values = {label.strip(): float(us) for label, _, us in rows}
            assert len(values) == 14
            runs[name].append(values)
            print(block + 1, name, round(sum(values.values()) / 1000, 3), 'ms; 14 results checked', flush=True)
    labels = list(runs['before'][0])
    medians = {name: {label: statistics.median(row[label] for row in rows) for label in labels}
               for name, rows in runs.items()}
    result = {'method': 'Five alternating pairs; full runner warmup/best-of-10 retained; construction, depth and destruction included; no simultaneous task builds/tests/instrumentation.',
        'order': orders, 'runs_us': runs, 'median_us': medians,
        'sum_of_benchmark_medians_us': {name: sum(rows.values()) for name, rows in medians.items()},
        'median_observed_total_us': {name: statistics.median(sum(row.values()) for row in rows) for name, rows in runs.items()},
        'lock_sha256': metadata['lock_sha256'], 'variants': metadata['variants']}
    write_json(HERE / 'runner-results.json', result)
    print(json.dumps({key: result[key] for key in ['median_us', 'sum_of_benchmark_medians_us']}, indent=2))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['build', 'time'])
    args = parser.parse_args()
    build() if args.mode == 'build' else measure()

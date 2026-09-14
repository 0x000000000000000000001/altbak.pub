"""Build full runners separately from their coordinated performance measurement."""
from pathlib import Path
import argparse
import json
import re
import shutil
import statistics
import subprocess
import sys

sys.dont_write_bytecode = True
import probe

HERE, BUILD, GENERATED = probe.HERE, probe.BUILD, probe.GENERATED
FULL = BUILD / 'full-output'
NAMES = ['before', 'rebuild', 'fields']
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000',
            '21536', '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(
    r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+'
    r'\(Execution time - best of 10\)\s+([0-9.]+) μs')


def write_json(path, value):
    path.write_text(json.dumps(value, indent=2) + '\n')


def paths(manifest, original_dir, local_names):
    def replace(match):
        original = (original_dir / match[1]).resolve()
        target = FULL / original.name if original.name in local_names else original
        assert target == FULL / original.name or original.is_relative_to(GENERATED)
        return 'path = ' + json.dumps(str(target))
    return re.sub(r'path\s*=\s*"([^"]+)"', replace, manifest)


def generated_inputs():
    """Hash source inputs, without traversing or copying Cargo build artifacts."""
    files = {GENERATED / 'Cargo.toml'}
    directories = [GENERATED / 'src']
    for directory in GENERATED.iterdir():
        if directory.is_dir() and (directory / 'Cargo.toml').is_file():
            files.add(directory / 'Cargo.toml')
            directories.append(directory / 'src')
            if (directory / 'build.rs').is_file():
                files.add(directory / 'build.rs')
    for directory in directories:
        files.update(path for path in directory.rglob('*') if path.is_file())
    hashes = {str(path.relative_to(GENERATED)): probe.digest(path.read_bytes())
              for path in sorted(files)}
    return {'file_count': len(hashes),
            'sha256': probe.digest(json.dumps(hashes, sort_keys=True).encode())}


def registry_packages(lockfile):
    # Read the small, regular Cargo.lock subset without requiring Python 3.11.
    packages = {}
    for section in lockfile.read_text().split('[[package]]')[1:]:
        fields = dict(re.findall(r'^(name|version|source|checksum) = "([^"]+)"$',
                                 section, re.M))
        if fields.get('source', '').startswith('registry+'):
            key = (fields['name'], fields['version'])
            packages[key] = (fields['source'], fields['checksum'])
    return packages


def check_seed_lock():
    seed = registry_packages(BUILD / 'Cargo.lock')
    actual = registry_packages(FULL / 'Cargo.lock')
    assert seed, 'Run probe.py prepare first to create the allocator lockfile'
    assert all(actual.get(key) == value for key, value in seed.items()), (seed, actual)
    seed_names = {name for name, _ in seed}
    assert {key for key in actual if key[0] in seed_names} == set(seed), (
        'Additional allocator/build-dependency versions appeared', actual)
    return [{'name': name, 'version': version, 'source': source, 'checksum': checksum}
            for (name, version), (source, checksum) in sorted(seed.items())]


def parse(output):
    rows = PATTERN.findall(output)
    assert [value for _, value, _ in rows] == EXPECTED, rows
    values = {name.strip(): float(us) for name, _, us in rows}
    assert len(values) == 14, rows
    assert all(value >= 0 for value in values.values()), values
    return values


def build():
    assert probe.SOURCE.read_bytes() == (BUILD / 'RBTree-original.rs').read_bytes()
    before_inputs = generated_inputs()
    FULL.mkdir(exist_ok=True)
    manifest = (GENERATED / 'Cargo.toml').read_text()
    manifest, replacements = re.subn(
        r'(?s)^\[workspace\].*?(?=\[package\])',
        '[workspace]\nmembers = ["Purs_App", "Purs_Test_RBTree"]\n\n', manifest)
    assert replacements == 1
    (FULL / 'Cargo.toml').write_text(paths(manifest, GENERATED, {'Purs_App'}))
    shutil.copytree(GENERATED / 'src', FULL / 'src', dirs_exist_ok=True)
    for name in ['Purs_App', 'Purs_Test_RBTree']:
        shutil.copytree(GENERATED / name, FULL / name, dirs_exist_ok=True)
        manifest_path = FULL / name / 'Cargo.toml'
        manifest_path.write_text(paths(
            manifest_path.read_text(), GENERATED / name,
            {'Purs_Test_RBTree'} if name == 'Purs_App' else set()))

    # No generated lockfile/target is assumed. Expand the already measured
    # allocator's lockfile offline, then freeze it for every variant.
    shutil.copy2(BUILD / 'Cargo.lock', FULL / 'Cargo.lock')
    with (BUILD / 'cargo-runner-resolve.log').open('w') as log:
        subprocess.run(['cargo', 'metadata', '--offline', '--format-version', '1'],
                       cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
    allocator_packages = check_seed_lock()
    lock_sha256 = probe.digest((FULL / 'Cargo.lock').read_bytes())
    variants = probe.variants(full=True)
    assert list(variants) == NAMES
    results = {}
    for name in NAMES:
        code = variants[name]
        (FULL / 'Purs_Test_RBTree/src/lib.rs').write_text(code)
        with (BUILD / f'cargo-runner-{name}.log').open('w') as log:
            # Keep one dependency cache, but force both the changed crate and
            # its callers to rebuild so no mtime fingerprint can mix variants.
            subprocess.run(
                ['cargo', 'clean', '--release', '--offline', '--locked',
                 '-p', 'Purs_Test_RBTree', '-p', 'Purs_App', '-p', 'purust_output'],
                cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
            subprocess.run(['cargo', 'build', '--release', '--offline', '--locked',
                            '-p', 'purust_output'],
                           cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
        assert probe.digest((FULL / 'Cargo.lock').read_bytes()) == lock_sha256
        binary = BUILD / f'runner-{name}'
        shutil.copy2(FULL / 'target/release/purust_output', binary)
        results[name] = {'binary_sha256': probe.digest(binary.read_bytes()),
                         'source_sha256': probe.digest(code.encode())}
        print(name, 'full runner built; not executed', flush=True)
    assert len({result['binary_sha256'] for result in results.values()}) == 3
    assert generated_inputs() == before_inputs, 'Generated source inputs changed during build'
    write_json(HERE / 'runner-build.json', {
        'profile': 'Original generated release O1/debug=true/mimalloc; one shared dependency cache; explicit clean of RBTree/App/root for each variant',
        'source_directory': str(GENERATED), 'generated_inputs': before_inputs,
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'cargo': subprocess.check_output(['cargo', '--version'], text=True).strip(),
        'lock_sha256': lock_sha256,
        'allocator_packages_identical_to_probe': allocator_packages,
        'variants': results, 'runners_executed': False})


def measure():
    # Root coordinates this invocation after other builds/tests have stopped.
    metadata = json.loads((HERE / 'runner-build.json').read_text())
    assert generated_inputs() == metadata['generated_inputs']
    assert probe.digest((FULL / 'Cargo.lock').read_bytes()) == metadata['lock_sha256']
    for name, code in probe.variants(full=True).items():
        assert probe.digest(code.encode()) == metadata['variants'][name]['source_sha256']
        assert probe.digest((BUILD / f'runner-{name}').read_bytes()) == metadata['variants'][name]['binary_sha256']
    runs = {name: [] for name in NAMES}
    orders = []
    labels = None
    for block in range(5):
        order = NAMES[block % 3:] + NAMES[:block % 3]
        if block % 2:
            order.reverse()
        orders.append(order)
        for name in order:
            output = subprocess.check_output([str(BUILD / f'runner-{name}')], text=True)
            (BUILD / f'runner-timing-{block + 1}-{name}.log').write_text(output)
            values = parse(output)
            if labels is None:
                labels = list(values)
            assert list(values) == labels
            runs[name].append(values)
            print(block + 1, name, round(sum(values.values()) / 1000, 3),
                  'ms; all 14 results checked', flush=True)
    medians = {name: {label: statistics.median(row[label] for row in rows)
                      for label in labels} for name, rows in runs.items()}
    assert generated_inputs() == metadata['generated_inputs']
    result = {
        'method': 'Five rotating/reversed blocks of three full runners; unchanged warmup/best-of-10 per benchmark; construction, depth and destruction remain inside RBTree; all 14 outputs verified in every run; coordinate to exclude simultaneous builds/tests/instrumentation',
        'order': orders, 'runs_us': runs, 'median_us': medians,
        'sum_of_benchmark_medians_us': {name: sum(values.values()) for name, values in medians.items()},
        'median_observed_total_us': {name: statistics.median(sum(row.values()) for row in rows)
                                     for name, rows in runs.items()},
        'binary_sha256': {name: data['binary_sha256'] for name, data in metadata['variants'].items()},
        'generated_inputs': metadata['generated_inputs'], 'lock_sha256': metadata['lock_sha256']}
    write_json(HERE / 'runner-results.json', result)
    print(json.dumps({key: result[key] for key in
                      ['median_us', 'sum_of_benchmark_medians_us', 'median_observed_total_us']}, indent=2))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['build', 'time'])
    args = parser.parse_args()
    build() if args.mode == 'build' else measure()

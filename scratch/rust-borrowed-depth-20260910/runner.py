"""Prepare full runners; timing requires a separate, coordinated invocation."""
import argparse
import json
import re
import shutil
import statistics
import subprocess
import probe

HERE, BUILD, GENERATED = probe.HERE, probe.BUILD, probe.GENERATED
FULL = BUILD / 'full-output'
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000', '21536', '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+\(Execution time - best of 10\)\s+([0-9.]+) μs')


def paths(text, keep):
    return re.sub(r'path = "(?:\.\./)?([A-Za-z_0-9]+)"',
        lambda match: match[0] if match[1] in keep else f'path = "{GENERATED / match[1]}"', text)


def parse(output):
    rows = PATTERN.findall(output)
    assert [value for _, value, _ in rows] == EXPECTED, rows
    return {name: float(us) for name, _, us in rows}


def build():
    FULL.mkdir(exist_ok=True)
    manifest = (GENERATED / 'Cargo.toml').read_text()
    manifest = re.sub(r'(?s)^\[workspace\].*?(?=\[package\])',
        '[workspace]\nmembers = ["Purs_App", "Purs_Test_RBTree"]\n\n', manifest)
    (FULL / 'Cargo.toml').write_text(paths(manifest, {'Purs_App'}))
    shutil.copy2(GENERATED / 'Cargo.lock', FULL / 'Cargo.lock')
    shutil.copytree(GENERATED / 'src', FULL / 'src', dirs_exist_ok=True)
    for name in ['Purs_App', 'Purs_Test_RBTree']:
        shutil.copytree(GENERATED / name, FULL / name, dirs_exist_ok=True)
        manifest_path = FULL / name / 'Cargo.toml'
        manifest_path.write_text(paths(manifest_path.read_text(), {'Purs_Test_RBTree'} if name == 'Purs_App' else set()))
    if not (FULL / 'target').exists():
        subprocess.run(['cp', '-cR', str(GENERATED / 'target'), str(FULL / 'target')], check=True)
    results = {}
    for name, code in probe.variants().items():
        (FULL / 'Purs_Test_RBTree/src/lib.rs').write_text(code)
        with (BUILD / f'cargo-{name}.log').open('w') as log:
            # Clear changed crates and their caller: copied mtime fingerprints
            # must never make Cargo reuse the other variant's RBTree object.
            subprocess.run(['cargo', 'clean', '--release', '-p', 'Purs_Test_RBTree', '-p', 'Purs_App', '-p', 'purust_output'],
                cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
            subprocess.run(['cargo', 'build', '--release', '--offline'],
                cwd=FULL, stdout=log, stderr=subprocess.STDOUT, check=True)
        binary = BUILD / f'runner-{name}'
        shutil.copy2(FULL / 'target/release/purust_output', binary)
        results[name] = {'binary_sha256': probe.digest(binary.read_bytes()), 'source_sha256': probe.digest(code.encode())}
        print(name, 'full runner built; not executed', flush=True)
    assert results['before']['binary_sha256'] != results['borrowed']['binary_sha256']
    assert probe.SOURCE.read_bytes() == (BUILD / 'RBTree-original.rs').read_bytes()
    (HERE / 'runner-build.json').write_text(json.dumps({'profile': 'release O1/debug=true/mimalloc; same lockfile; explicit clean of changed crates for each variant', 'variants': results, 'runners_executed': False}, indent=2) + '\n')


def measure():
    # Root explicitly coordinates this command after all other work has stopped.
    variants = ['before', 'borrowed']
    runs = {name: [] for name in variants}
    for pair in range(5):
        for name in (variants if pair % 2 == 0 else variants[::-1]):
            output = subprocess.check_output([str(BUILD / f'runner-{name}')], text=True)
            (BUILD / f'timing-{pair + 1}-{name}.log').write_text(output)
            values = parse(output)
            runs[name].append(values)
            print(pair + 1, name, round(sum(values.values()) / 1000, 3), 'ms; 14 results checked', flush=True)
    labels = list(runs['before'][0])
    medians = {name: {label: statistics.median(row[label] for row in rows) for label in labels} for name, rows in runs.items()}
    result = {'method': 'Five alternating pairs of full runners with unchanged warmup/best-of-10. Construction, depth and destruction all remain inside RBTree. All 14 results verified. Coordinate to exclude concurrent compilation, tests and instrumentation.',
        'runs_us': runs, 'median_us': medians, 'total_us': {name: sum(values.values()) for name, values in medians.items()},
        'binary_sha256': {name: probe.digest((BUILD / f'runner-{name}').read_bytes()) for name in variants}}
    (HERE / 'runner-results.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps({'median_us': medians, 'total_us': result['total_us']}, indent=2))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['build', 'time'])
    args = parser.parse_args()
    build() if args.mode == 'build' else measure()

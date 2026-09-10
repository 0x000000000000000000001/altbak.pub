"""Compare full runners generated from the same fresh TAST, without Rust patches."""
from pathlib import Path
import argparse
import hashlib
import json
import re
import shutil
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE / 'build'
AFTER = ROOT / 'run/bak/rust/output/purust_output'
BEFORE = BUILD / 'before-output'
EXPECTED = ['7', '55', '202950', '100000', '20000', '125', '100000', '21536',
            '22', '10000000', '1200', '1000000', '202950', '5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+'
                     r'\(Execution time - best of 10\)\s+([0-9.]+) μs')


def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def parse(output):
    rows = PATTERN.findall(output)
    assert [value for _, value, _ in rows] == EXPECTED, rows
    return {name: float(us) for name, _, us in rows}


def build():
    # The before bundle is saved before changing Purust. Both invocations use
    # the same source/output, frontend, directives, dependencies and runtime.
    assert (BEFORE / 'Cargo.toml').is_file()
    assert (AFTER / 'target/release/purust_output').is_file()
    shutil.copy2(AFTER / 'Cargo.lock', BEFORE / 'Cargo.lock')
    with (BUILD / 'cargo-before.log').open('w') as log:
        # A copied Cargo target can look fresh despite different source bytes
        # when the before sources predate the after build. Rebuild this isolated
        # reference from scratch instead of trusting those mtime fingerprints.
        subprocess.run(['cargo', 'clean', '--release'], cwd=BEFORE,
                       stdout=log, stderr=subprocess.STDOUT, check=True)
        subprocess.run(['cargo', 'build', '--release', '--offline'], cwd=BEFORE,
                       stdout=log, stderr=subprocess.STDOUT, check=True)
    for name, directory in [('before', BEFORE), ('after', AFTER)]:
        binary = BUILD / f'runner-{name}'
        shutil.copy2(directory / 'target/release/purust_output', binary)
        output = subprocess.check_output([str(binary)], text=True)
        parse(output)
        (BUILD / f'check-{name}.log').write_text(output)
        print(name, '14 results verified', flush=True)
    assert digest(BUILD / 'runner-before') != digest(BUILD / 'runner-after'), 'Identical binaries: invalid comparison'
    sources = {name: {str(path.relative_to(directory)): digest(path)
                     for path in directory.rglob('*.rs') if 'target' not in path.parts}
               for name, directory in [('before', BEFORE), ('after', AFTER)]}
    assert sources['before'].keys() == sources['after'].keys()
    changed = [path for path in sources['before'] if sources['before'][path] != sources['after'][path]]
    metadata = {
        'before': json.loads((HERE / 'before-metadata.json').read_text()),
        'after_bundle_sha256': digest(ROOT.parent / 'purust/purust/bin/purust.js'),
        'tast_sha256': {str(path.relative_to(ROOT / 'run/bak/rust/output')): digest(path)
                       for path in (ROOT / 'run/bak/rust/output').glob('*/corefn.json')},
        'sources_sha256': sources,
        'changed_rust_sources': changed,
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'profile': 'release, opt-level=1, debug=true, mimalloc 0.1.32 requirement; identical Cargo.lock',
        'readme_sha256': digest(ROOT.parent / 'altbak.pub/README.md'),
    }
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2) + '\n')
    print('Changed generated Rust:', changed, flush=True)


def measure():
    variants = ['before', 'after']
    runs = {name: [] for name in variants}
    for block in range(5):
        order = variants if block % 2 == 0 else variants[::-1]
        for name in order:
            output = subprocess.check_output([str(BUILD / f'runner-{name}')], text=True)
            (BUILD / f'timing-{block + 1}-{name}.log').write_text(output)
            values = parse(output)
            runs[name].append(values)
            print(block + 1, name, round(sum(values.values()) / 1000, 3), 'ms', flush=True)
    names = list(runs['before'][0])
    medians = {name: {bench: statistics.median(row[bench] for row in rows) for bench in names}
               for name, rows in runs.items()}
    result = {
        'method': '5 alternating pairs of uninstrumented full runner processes, identical fresh TAST; '
                  'unchanged warmup and best of 10 per benchmark; all 14 results checked each run; '
                  'no concurrent compilation, instrumentation or tests; sums of per-benchmark medians.',
        'runs_us': runs, 'median_us': medians,
        'total_us': {name: sum(values.values()) for name, values in medians.items()},
        'binary_sha256': {name: digest(BUILD / f'runner-{name}') for name in variants},
    }
    (HERE / 'runner-results.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps({'median_us': medians, 'total_us': result['total_us']}, indent=2))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['build', 'time'])
    arguments = parser.parse_args()
    build() if arguments.mode == 'build' else measure()

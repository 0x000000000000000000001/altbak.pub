"""Isolated borrowed record projections; preserve the runtime layout and setters."""
from pathlib import Path
import argparse
import ast
import hashlib
import json
import re
import shutil
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE / 'build'
LIVE = ROOT / 'run/bak/rust/output/purust_output'
WORK = BUILD / 'workspace'
RECORD = Path('Purs_Test_Records/src/lib.rs')
CORE = Path('purust_core/src/lib.rs')
NAMES = ['before', 'root_borrow', 'path_borrow']
EXPECTED = ['7','55','202950','100000','20000','125','100000','21536','22','10000000','1200','1000000','202950','5']
PATTERN = re.compile(r'\(Test\)\s+([^\n]+)\s+\(Output & Warm-up\)\s+(\S+)\s+\(Execution time - best of 10\)\s+([0-9.]+) μs')


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def parse(output):
    rows = PATTERN.findall(output)
    assert [v for _, v, _ in rows] == EXPECTED, rows
    return {n: float(t) for n, _, t in rows}


def sources():
    original = (LIVE / RECORD).read_text()
    start = original.index('    let _record_update_0 = ')
    end = original.index('    let mut _base = purs_local_1;', start)
    rhs = original[start:end]
    assert rhs.count('purs_local_1.clone()') == 4
    root_rhs = rhs.replace('purs_local_1.clone()', '&purs_local_1')
    path_rhs, count = re.subn(r'\.get_([a-f])\(\)', r'.get_\1_ref()', root_rhs)
    assert count == 9
    return {'before': original, 'root_borrow': original[:start] + root_rhs + original[end:],
            'path_borrow': original[:start] + path_rhs + original[end:]}


def borrowed_core():
    core = (LIVE / CORE).read_text()
    methods = []
    for field in 'abcdef':
        method, = re.findall(r'    pub fn get_' + field + r'\(&self\) -> UnknownType \{.*?^    \}', core, re.S | re.M)
        methods.append(method.replace(f'get_{field}(', f'get_{field}_ref(')
                       .replace('-> UnknownType', '-> &UnknownType')
                       .replace('.clone().unwrap()', '.as_ref().unwrap()'))
    return core + '\nimpl Value {\n' + '\n'.join(methods) + '\n}\n'


def build():
    BUILD.mkdir(exist_ok=True)
    assert not WORK.exists(), 'Use a fresh isolated workspace; no copied Cargo target cache'
    shutil.copytree(LIVE, WORK, ignore=shutil.ignore_patterns('target'))
    # All variants have the same extra unused/used borrowed getter methods.
    # They retain the original getters' visibility, dispatch and checks.
    (WORK / CORE).write_text(borrowed_core())
    nodes = ast.parse((ROOT / 'scratch/rust-records-audit-20260909/probe.py').read_text()).body
    helpers = {node.targets[0].id: ast.literal_eval(node.value) for node in nodes
               if isinstance(node, ast.Assign) and isinstance(node.targets[0], ast.Name)
               and node.targets[0].id in ['HEADER', 'HELPERS', 'COUNT', 'TIME']}
    check = (ROOT / 'scratch/rust-record-path-20260909/sharing-check.rs').read_text()
    validations = {}
    for name, source in sources().items():
        (WORK / RECORD).write_text(source)
        (BUILD / f'Records-{name}.rs').write_text(source)
        with (BUILD / f'cargo-{name}.log').open('w') as log:
            subprocess.run(['cargo', 'build', '--release', '--offline'], cwd=WORK,
                           stdout=log, stderr=subprocess.STDOUT, check=True)
        shutil.copy2(WORK / 'target/release/purust_output', BUILD / f'runner-{name}')
        deps = WORK / 'target/release/deps'
        libs = []
        for library in ['purust_core', 'perceus_ptr', 'mimalloc']:
            paths = list(deps.glob(f'lib{library}-*.rlib'))
            assert len(paths) == 1, (library, paths)
            libs += ['--extern', f'{library}={paths[0]}']
        kernel = source[source.index('pub fn Test_Records_updateRec('):source.index('pub fn Test_Records_describe(')]
        validations[name] = {}
        for mode, main in [('check', check), ('count', helpers['COUNT']), ('time', helpers['TIME'])]:
            path = BUILD / f'{mode}-{name}.rs'
            path.write_text(helpers['HEADER'] + kernel + helpers['HELPERS'] + main)
            subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path),
                            '-o', str(path.with_suffix('')), '-L', f'dependency={deps}', *libs], check=True)
            if mode != 'time':
                output = subprocess.check_output([str(path.with_suffix(''))], text=True)
                validations[name][mode] = output
                if mode == 'count':
                    assert output.splitlines()[-1] == '10000 3 3', output
        print(name, 'built; all four fields, sharing and 3 allocations/releases checked', flush=True)
    (HERE / 'validation.json').write_text(json.dumps(validations, indent=2) + '\n')
    (HERE / 'metadata.json').write_text(json.dumps({
        'source_sha256': {str(path): sha(LIVE / path) for path in [RECORD, CORE]},
        'binary_sha256': {name: sha(BUILD / f'runner-{name}') for name in NAMES},
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'profile': 'O1/debug=true/mimalloc; fresh isolated Cargo workspace, same added borrowed methods in all variants',
        'scope': 'Only four RHS projections change. Scalar values, all setters and runtime layouts remain unchanged.',
    }, indent=2) + '\n')


def measure():
    full = {name: [] for name in NAMES}
    kernel = {name: [] for name in NAMES}
    for block in range(5):
        order = NAMES[block % 3:] + NAMES[:block % 3]
        if block % 2:
            order = order[::-1]
        for name in order:
            output = subprocess.check_output([str(BUILD / f'runner-{name}')], text=True)
            (BUILD / f'runner-{block + 1}-{name}.log').write_text(output)
            full[name].append(parse(output))
            samples = list(map(int, subprocess.check_output([str(BUILD / f'time-{name}')], text=True).split()))
            assert len(samples) == 20
            kernel[name].append(samples)
            print(block + 1, name, 'suite', round(sum(full[name][-1].values()) / 1000, 3),
                  'ms, Records', full[name][-1]['Deep Record Updates (10k iterations):'], 'us', flush=True)
    labels = list(full['before'][0])
    medians = {name: {k: statistics.median(row[k] for row in rows) for k in labels}
               for name, rows in full.items()}
    result = {'method': 'Five rotating/reversing triples, uninstrumented O1/mimalloc; all 14 runner results verified. '
              'Separate kernels observe all four fields and include destruction; 20 samples after warmup per process. '
              'No concurrent builds, tests or instrumentation.',
              'runs_us': full, 'median_us': medians,
              'total_us': {name: sum(row.values()) for name, row in medians.items()},
              'kernel_samples_ns': kernel,
              'kernel_median_us': {name: statistics.median(sum(rows, [])) / 1000 for name, rows in kernel.items()}}
    (HERE / 'results.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps({k: result[k] for k in ['total_us', 'kernel_median_us']}, indent=2))


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['build', 'time'])
    args = parser.parse_args()
    build() if args.mode == 'build' else measure()

"""Validate and count frozen Rust actually emitted by both compiler versions.

No source rewrite implements the optimization: only counters are injected.
The native-Rc validation binaries contain the emitted kernel unchanged.
"""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import re
import subprocess
import sys

sys.dont_write_bytecode = True
HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE / 'build'
OUT = BUILD / 'instrumentation'
PROTOTYPE = ROOT / 'scratch/rust-rotation-fields-20260914'
PREVIOUS = ROOT / 'scratch/rust-perceus-counts-20260909'
EXTRA = ROOT / 'scratch/rust-child-field-20260909/validation-extra.rs'
NAMES = ['before', 'after']
HELPER = 'Test_RBTree_balance__purust_permute_fields'


def load_module(name, path):
    spec = importlib.util.spec_from_file_location(name, path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


prototype = load_module('rotation_prototype', PROTOTYPE / 'probe.py')
previous = load_module('tracked_probe', PREVIOUS / 'probe.py')


def digest(data):
    return hashlib.sha256(data).hexdigest()


def inputs():
    result = {}
    OUT.mkdir(parents=True, exist_ok=True)
    proof = json.loads((HERE / 'generation-proof.json').read_text())
    for name in NAMES:
        variant = 'disabled' if name == 'before' else name
        source = BUILD / variant / 'Purs_Test_RBTree/src/lib.rs'
        sha256 = digest(source.read_bytes())
        assert sha256 == proof[f'{name}_RBTree_sha256'], 'Frozen source differs from generation proof'
        full = source.read_text()
        code = prototype.kernel(full)
        assert ('fn ' + HELPER + '(' in code) == (name == 'after')
        result[name] = {'path': str(source), 'sha256': sha256, 'code': code}
    return result


def rust_binary(name, source):
    path = OUT / f'{name}.rs'
    binary = OUT / name
    path.write_text(source)
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
    output = subprocess.check_output([str(binary)], text=True)
    (OUT / f'{name}.log').write_text(output)
    return output


def validate():
    sources = inputs()
    base = (PREVIOUS / 'count_harness.rs').read_text()
    base = base[base.index('fn validate('):].replace('tracked::Rc', 'std::rc::Rc')
    base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n', '', base)
    base = base.replace('fn main() {', 'fn original_checks() {')
    extra = EXTRA.read_text()
    results = {}
    for name, source in sources.items():
        direct = ''
        call = ''
        if name == 'after':
            direct = (PROTOTYPE / 'validation-extra.rs').read_text().replace('probe_rotate_ll', HELPER)
            call = 'validation_extra();'
        code = '#![allow(warnings)]\n' + source['code'] + base + extra + direct
        code += '\nfn main() { original_checks(); extra_checks(); ' + call + ' println!("all checks passed"); }\n'
        output = rust_binary(f'validate-{name}', code)
        results[name] = {'source': source['path'], 'source_sha256': source['sha256'], 'output': output}
        print(name, output, flush=True)
    (HERE / 'validation.json').write_text(json.dumps(results, indent=2) + '\n')


COUNTERS = '''
thread_local! { static PROBE_COUNTS: std::cell::Cell<[u64;3]> = const { std::cell::Cell::new([0;3]) }; }
fn probe_event(i: usize) { PROBE_COUNTS.with(|s| { let mut a=s.get(); a[i]+=1; s.set(a); }); }
fn probe_reset() { PROBE_COUNTS.with(|s|s.set([0;3])); }
fn probe_print() { PROBE_COUNTS.with(|s| { print!("PROBE"); for n in s.get() { print!(" {n}"); } println!(); }); }
'''


def count():
    sources = inputs()
    harness = (PREVIOUS / 'count_harness.rs').read_text()
    harness = harness.replace('tracked::phase("unique_build");', 'tracked::phase("unique_build"); probe_reset();')
    harness = harness.replace('tracked::phase("unique_depth_and_drop");', 'probe_print(); tracked::phase("unique_depth_and_drop");')
    results = {}
    for name, source in sources.items():
        code = source['code']
        take = 'pub fn __purust_take(&mut self) -> std::option::Option<Self> {'
        rebuild = 'let payload = crate::Tree::T(a0, a1, a2, a3);'
        assert code.count(take) == code.count(rebuild) == 1
        code = code.replace(take, take + ' probe_event(0);')
        code = code.replace(rebuild, 'probe_event(1); ' + rebuild)
        if name == 'after':
            start = code.index('fn ' + HELPER + '(')
            end = code.index('\ntrue\n}', start)
            code = code[:end] + '\nprobe_event(2);' + code[end:]
        code = code.replace('std::rc::Rc', 'tracked::Rc')
        code = '#![allow(warnings)]\n#[path="' + str(PREVIOUS / 'tracked_rc.rs') + '"] mod tracked;\n' + code + COUNTERS + harness
        output = rust_binary(f'count-{name}', code)
        line, = [line for line in output.splitlines() if line.startswith('PROBE ')]
        summary, _ = previous.summarize('\n'.join(line for line in output.splitlines() if not line.startswith('PROBE ')))
        counts = dict(zip(['take', 'rebuild_helper', 'generated_permutations'], map(int, line.split()[1:])))
        results[name] = {'source': source['path'], 'source_sha256': source['sha256'], 'unique_build': counts, 'phases': summary['phases']}
        assert summary['phases']['unique_build']['new'] == 100001
        assert counts['generated_permutations'] == (0 if name == 'before' else 99978)
        assert counts['take'] == (299934 if name == 'before' else 0)
        print(name, json.dumps(counts), flush=True)
    (HERE / 'counts.json').write_text(json.dumps(results, indent=2) + '\n')


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['validate', 'count'])
    args = parser.parse_args()
    {'validate': validate, 'count': count}[args.mode]()

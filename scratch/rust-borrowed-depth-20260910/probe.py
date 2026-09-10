"""Isolated borrowed-depth prototype with the same owned public boundary."""
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
GENERATED = ROOT / 'run/bak/rust/output/purust_output'
SOURCE = GENERATED / 'Purs_Test_RBTree/src/lib.rs'
PREVIOUS = HERE.parent / 'rust-perceus-counts-20260909'
spec = importlib.util.spec_from_file_location('emitted_checks', HERE.parent / 'rust-post-call-child-integration-20260910/count.py')
emitted_checks = importlib.util.module_from_spec(spec)
spec.loader.exec_module(emitted_checks)

BORROWED = '''pub fn Test_RBTree_depth(owner: std::rc::Rc<crate::Tree>) -> i64 {
    let depth = Test_RBTree_depth__purust_borrowed(owner.as_ref());
    // Keep destruction inside the public call, including for the timed runner.
    drop(owner);
    depth
}

fn Test_RBTree_depth__purust_borrowed(node: &crate::Tree) -> i64 {
    match node {
        crate::Tree::E => 0,
        crate::Tree::T(_, left, _, right) => {
            let left_depth = Test_RBTree_depth__purust_borrowed(left.as_ref());
            let right_depth = Test_RBTree_depth__purust_borrowed(right.as_ref());
            1 + if left_depth > right_depth { left_depth } else { right_depth }
        }
    }
}

'''


def digest(data):
    return hashlib.sha256(data).hexdigest()


def variants():
    original = (BUILD / 'RBTree-original.rs').read_text()
    begin = original.index('pub fn Test_RBTree_depth(')
    end = original.index('pub fn Test_RBTree_balance(', begin)
    assert 'Test_RBTree_depth__purust_borrowed' not in original
    result = {'before': original, 'borrowed': original[:begin] + BORROWED + original[end:]}
    for name, code in result.items():
        (BUILD / f'RBTree-{name}.rs').write_text(code)
    return result


def prepare():
    BUILD.mkdir(parents=True, exist_ok=True)
    snapshot = BUILD / 'RBTree-original.rs'
    original = SOURCE.read_bytes()
    if snapshot.exists(): assert snapshot.read_bytes() == original, 'Live source changed after capture'
    else: snapshot.write_bytes(original)
    source_variants = variants()
    metadata = {'source': str(SOURCE), 'source_sha256': digest(original),
        'variants_sha256': {name: digest(code.encode()) for name, code in source_variants.items()},
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'change': 'Only the owned depth function is replaced by an owned wrapper, explicit final drop, and a worker borrowing &Tree; all construction/insertion code remains identical.',
        'timing_scope': 'Full runner construction + traversal + destruction, O1/debug=true/mimalloc. No traversal-only timing.',
        'timings_run': False}
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2) + '\n')


def count():
    harness = (PREVIOUS / 'count_harness.rs').read_text()
    results = {'method': 'Logical Rc operations through the existing one-word tracked wrapper; no timing and no insertion changes.', 'variants': {}}
    for name, source in variants().items():
        code = emitted_checks.kernel(source).replace('std::rc::Rc', 'tracked::Rc')
        path = BUILD / f'count-{name}.rs'
        path.write_text('#![allow(warnings)]\n' + f'#[path="{PREVIOUS}/tracked_rc.rs"] mod tracked;\n' + code + harness)
        binary = BUILD / f'count-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        (BUILD / f'events-{name}.tsv').write_text(output)
        summary, _ = emitted_checks.previous.summarize(output)
        phases = summary['phases']
        assert phases['unique_build']['new'] == 100001
        assert phases['unique_depth_and_drop']['drop_last'] == 100001
        assert phases['unique_depth_and_drop'].get('clone', 0) == (200000 if name == 'before' else 0)
        results['variants'][name] = {'source_sha256': digest(source.encode()), 'phases': phases,
            'global_handle_and_payload_balance_passed': True}
        print(name, 'depth and drop:', dict(phases['unique_depth_and_drop']), flush=True)
    assert results['variants']['before']['phases']['unique_build'] == results['variants']['borrowed']['phases']['unique_build']
    assert SOURCE.read_bytes() == (BUILD / 'RBTree-original.rs').read_bytes()
    (HERE / 'counts.json').write_text(json.dumps(results, indent=2) + '\n')


def validate():
    base = (PREVIOUS / 'count_harness.rs').read_text()
    base = base[base.index('fn validate('):].replace('tracked::Rc', 'std::rc::Rc')
    base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n', '', base)
    base = base.replace('fn main() {', 'fn original_checks() {')
    extra = (HERE.parent / 'rust-child-field-20260909/validation-extra.rs').read_text()
    depth = (HERE / 'validation-depth.rs').read_text()
    results = {'method': 'Native Rc O1 correctness without clocks: existing 100k-tree/rotation/persistence harness plus repeated depth and shared/weak destruction checks.', 'variants': {}}
    for name, source in variants().items():
        path = BUILD / f'validate-{name}.rs'
        path.write_text('#![allow(warnings)]\n' + emitted_checks.kernel(source) + base + extra + depth
            + '\nfn main() { original_checks(); extra_checks(); depth_checks(); }\n')
        binary = BUILD / f'validate-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        results['variants'][name] = {'source_sha256': digest(source.encode()), 'passed': True, 'output': output.strip()}
        print(name, output.strip(), flush=True)
    assert SOURCE.read_bytes() == (BUILD / 'RBTree-original.rs').read_bytes()
    (HERE / 'validation.json').write_text(json.dumps(results, indent=2) + '\n')


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['prepare', 'count', 'validate'])
    args = parser.parse_args()
    {'prepare': prepare, 'count': count, 'validate': validate}[args.mode]()

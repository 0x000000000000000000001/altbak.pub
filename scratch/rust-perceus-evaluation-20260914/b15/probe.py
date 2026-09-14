"""B15 consuming depth experiment; only scratch Rust changes, no timings here."""
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
ROOT = HERE.parents[2]
BUILD = HERE / 'build'
SOURCE = ROOT / 'run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs'
PREVIOUS = ROOT / 'scratch/rust-perceus-counts-20260909'
BORROWED_PREVIOUS = ROOT / 'scratch/rust-borrowed-depth-20260910'
PERMUTATION = ROOT / 'scratch/rust-rotation-fields-20260914'


def load_module(name, path):
    spec = importlib.util.spec_from_file_location(name, path)
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


previous = load_module('rc_counts', PREVIOUS / 'probe.py')
old_depth = load_module('borrowed_depth', BORROWED_PREVIOUS / 'probe.py')
extractor = load_module('kernel_extractor', PERMUTATION / 'probe.py')

CONSUMING = '''pub fn Test_RBTree_depth(owner: std::rc::Rc<crate::Tree>) -> i64 {
    // B15: last owner destroys only the parent cell and moves its fields.
    // Shared owners clone the payload, preserving independent retained trees.
    match std::rc::Rc::unwrap_or_clone(owner) {
        crate::Tree::E => 0,
        crate::Tree::T(_, left, _, right) => {
            let left_depth = Test_RBTree_depth(left);
            let right_depth = Test_RBTree_depth(right);
            1 + if left_depth > right_depth { left_depth } else { right_depth }
        }
    }
}

'''

HYBRID = '''pub fn Test_RBTree_depth(owner: std::rc::Rc<crate::Tree>) -> i64 {
    match std::rc::Rc::try_unwrap(owner) {
        Ok(crate::Tree::E) => 0,
        Ok(crate::Tree::T(_, left, _, right)) => {
            let left_depth = Test_RBTree_depth(left);
            let right_depth = Test_RBTree_depth(right);
            1 + if left_depth > right_depth { left_depth } else { right_depth }
        }
        Err(owner) => {
            // A shared subtree is read without cloning any of its fields.
            let depth = Test_RBTree_depth__purust_borrowed(owner.as_ref());
            drop(owner);
            depth
        }
    }
}

''' + old_depth.BORROWED[old_depth.BORROWED.index('fn Test_RBTree_depth__purust_borrowed('):]


def digest(data):
    return hashlib.sha256(data).hexdigest()


def variants():
    original = (BUILD / 'RBTree-original.rs').read_text()
    begin = original.index('pub fn Test_RBTree_depth(')
    end = original.index('pub fn Test_RBTree_balance(', begin)
    assert 'Test_RBTree_depth__purust_borrowed' not in original
    return {'before': original,
            'borrowed': original[:begin] + old_depth.BORROWED + original[end:],
            'consuming': original[:begin] + CONSUMING + original[end:],
            'hybrid': original[:begin] + HYBRID + original[end:]}


def prepare():
    BUILD.mkdir(parents=True, exist_ok=True)
    original = SOURCE.read_bytes()
    snapshot = BUILD / 'RBTree-original.rs'
    if snapshot.exists():
        assert snapshot.read_bytes() == original, 'Live source changed after capture'
    else:
        snapshot.write_bytes(original)
    sources = variants()
    if (HERE / 'metadata.json').exists():
        previous_metadata = json.loads((HERE / 'metadata.json').read_text())
        for name in ['before', 'borrowed', 'consuming']:
            assert digest(sources[name].encode()) == previous_metadata['variants_sha256'][name], 'An established variant changed'
    for name, code in sources.items():
        (BUILD / f'RBTree-{name}.rs').write_text(code)
    metadata = {
        'source': str(SOURCE), 'source_sha256': digest(original),
        'variants_sha256': {name: digest(code.encode()) for name, code in variants().items()},
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'change': 'Only depth replaced. Before is emitted Rust; borrowed is B14 control with final owned drop; consuming is B15 unwrap-or-clone with child moves on unique nodes; hybrid uses try_unwrap for unique nodes and the same borrowed worker on shared subtrees.',
        'scope': 'Native Rc Tree closed over Color Copy, Int, Rc<Tree>. No callback, opaque payload destructor or observable operation inside traversal. No threaded claim.',
        'timing_scope': 'Full runner must include construction, traversal and all destruction with O1/debug=true/mimalloc, no instrumentation. Parent coordinates clocks separately.',
        'timings_run': False,
    }
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2) + '\n')
    print(json.dumps(metadata, indent=2), flush=True)


def run_native(name, code):
    path = BUILD / f'{name}.rs'
    binary = BUILD / name
    path.write_text(code)
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
    output = subprocess.check_output([str(binary)], text=True)
    (BUILD / f'{name}.log').write_text(output)
    return output


def count():
    harness = (PREVIOUS / 'count_harness.rs').read_text()
    wrapper = (PREVIOUS / 'tracked_rc.rs').read_text()
    insertion = '''    #[track_caller]
    pub fn try_unwrap(mut this: Self) -> Result<T, Self> {
        let pointer = this.0.take().unwrap();
        let kind = pointer.kind();
        let line = Location::caller().line();
        match NativeRc::try_unwrap(pointer) {
            Ok(value) => {
                record("unwrap_unique", kind, line);
                Ok(value)
            }
            Err(pointer) => {
                // A failed try returns the same owner; no release or clone.
                record("try_unwrap_shared", kind, line);
                Err(Self(Some(pointer)))
            }
        }
    }
'''
    marker = '    pub fn downgrade(this: &Self) -> std::rc::Weak<T> {'
    assert wrapper.count(marker) == 1
    wrapper = wrapper.replace(marker, insertion + marker)
    wrapper_path = BUILD / 'tracked_rc.rs'
    wrapper_path.write_text(wrapper)
    results = {'method': 'Scratch copy of existing one-word native Rc wrapper, extended only with try_unwrap. Operations only, no clocks. Successful consumption=unwrap_unique; failed try_unwrap_shared returns ownership and is not a release. Destruction=drop_last+unwrap_unique. Global balances checked.', 'variants': {}}
    for name, source in variants().items():
        code = extractor.kernel(source).replace('std::rc::Rc', 'tracked::Rc')
        output = run_native(f'count-{name}', '#![allow(warnings)]\n' + f'#[path="{wrapper_path}"] mod tracked;\n' + code + harness)
        summary, _ = previous.summarize(output)
        phases = summary['phases']
        phase = phases['unique_depth_and_drop']
        assert phases['unique_build']['new'] == 100001
        assert phase.get('new', 0) == 0
        assert phase.get('drop_last', 0) + phase.get('unwrap_unique', 0) == 100001
        assert phase.get('clone', 0) == (200000 if name == 'before' else 0)
        results['variants'][name] = {'source_sha256': digest(source.encode()), 'phases': phases, 'global_handle_and_payload_balance_passed': True}
        print(name, 'depth and drop:', dict(phase), flush=True)
    build_counts = [row['phases']['unique_build'] for row in results['variants'].values()]
    assert all(counts == build_counts[0] for counts in build_counts)
    assert SOURCE.read_bytes() == (BUILD / 'RBTree-original.rs').read_bytes()
    (HERE / 'counts.json').write_text(json.dumps(results, indent=2) + '\n')


def validate():
    base = (PREVIOUS / 'count_harness.rs').read_text()
    base = base[base.index('fn validate('):].replace('tracked::Rc', 'std::rc::Rc')
    base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n', '', base)
    base = base.replace('fn main() {', 'fn original_checks() {')
    extra = (ROOT / 'scratch/rust-child-field-20260909/validation-extra.rs').read_text()
    depth = (BORROWED_PREVIOUS / 'validation-depth.rs').read_text()
    results = {'method': 'Native Rc O1, no clocks. Existing 100k rotations/invariants/persistence suite and owned depth/shared/Weak boundary checks.', 'variants': {}}
    for name, source in variants().items():
        output = run_native(f'validate-{name}', '#![allow(warnings)]\n' + extractor.kernel(source) + base + extra + depth
            + '\nfn main() { original_checks(); extra_checks(); depth_checks(); }\n')
        results['variants'][name] = {'source_sha256': digest(source.encode()), 'passed': True, 'output': output.strip()}
        print(name, output.strip(), flush=True)
    scope = run_native('scope-checks', (HERE / 'scope-checks.rs').read_text())
    results['scope_checks'] = {'passed': True, 'output': scope.strip(), 'meaning': 'Adversarial callbacks, destructor order and unwinding are counterexamples outside the eligible plain-tree total traversal; they are not claimed equivalent.'}
    assert SOURCE.read_bytes() == (BUILD / 'RBTree-original.rs').read_bytes()
    (HERE / 'validation.json').write_text(json.dumps(results, indent=2) + '\n')


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['prepare', 'count', 'validate'])
    args = parser.parse_args()
    {'prepare': prepare, 'count': count, 'validate': validate}[args.mode]()

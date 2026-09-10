"""Count ownership operations in freshly emitted Rust; never use for timings."""
from pathlib import Path
from collections import defaultdict
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
PREVIOUS = HERE.parent / 'rust-perceus-counts-20260909'
SOURCE = ROOT / 'run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs'
BEFORE = BUILD / 'before-output/Purs_Test_RBTree/src/lib.rs'
spec = importlib.util.spec_from_file_location('ownership_probe', PREVIOUS / 'probe.py')
previous = importlib.util.module_from_spec(spec)
spec.loader.exec_module(previous)


def sha256(data):
    return hashlib.sha256(data).hexdigest()


def inputs():
    BUILD.mkdir(exist_ok=True)
    result = {}
    for name, source in [('before', BEFORE), ('after', SOURCE)]:
        data = source.read_bytes()
        snapshot = BUILD / f'operations-input-{name}.rs'
        if snapshot.exists():
            assert snapshot.read_bytes() == data, f'{source} changed after capture'
        else:
            snapshot.write_bytes(data)
        result[name] = {'path': str(source), 'sha256': sha256(data),
            'snapshot': str(snapshot), 'code': data.decode()}
    return result


def kernel(source):
    """Copy kernel declarations verbatim, including every emitted helper."""
    starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
    result = [''] * len(source.splitlines())
    ranges = [(source.index('#[derive(Clone'), starts[0].start())]
    names = {'max', 'makeBlack', 'depth', 'balance', 'ins', 'insert', 'buildTree'}
    for index, match in enumerate(starts):
        name = match[1].removeprefix('Test_RBTree_')
        if name in names or '__purust_' in name:
            ranges.append((match.start(), starts[index + 1].start() if index + 1 < len(starts) else len(source)))
    for start, end in ranges:
        line = source[:start].count('\n')
        part = source[start:end].splitlines()
        result[line:line + len(part)] = part
    return '\n'.join(result) + '\n'


COUNTERS = '''
thread_local! { static POST_COUNTS: std::cell::Cell<[u64;7]> = const { std::cell::Cell::new([0;7]) }; }
fn post_event(i: usize) { POST_COUNTS.with(|s| { let mut a=s.get(); a[i]+=1; s.set(a); }); }
fn post_reset() { POST_COUNTS.with(|s|s.set([0;7])); }
fn post_print() { POST_COUNTS.with(|s| { print!("POST"); for n in s.get() { print!(" {n}"); } println!(); }); }
'''
COUNTER_NAMES = ['take', 'rebuild_helper', 'red_retained_fields', 'post_no_rotation',
    'post_rotation', 'post_unique', 'ins_calls']


def instrument(code, name):
    replacements = {}
    def replace(old, new, expected):
        nonlocal code
        count = code.count(old)
        assert count == expected, (name, old, count, expected)
        code = code.replace(old, new)
        replacements[old] = count
    take = 'pub fn __purust_take(&mut self) -> std::option::Option<Self> {'
    replace(take, take + ' post_event(0);', 1)
    rebuild = 'let payload = crate::Tree::T(a0, a1, a2, a3);'
    replace(rebuild, 'post_event(1); ' + rebuild, 1)
    red = 'if let std::option::Option::Some(_purust_child_slot) = std::rc::Rc::get_mut(&mut purs_local_1) {'
    replace(red, red + ' post_event(2);', 2)
    post = 'if let std::option::Option::Some(_purust_post_slot) = std::rc::Rc::get_mut(&mut purs_local_1) {'
    replace(post, post + ' post_event(5);', 2 if name == 'after' else 0)
    predicate = 'if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3) { purs_local_1 } else {'
    replace(predicate, predicate.replace('{ purs_local_1 } else {',
        '{ post_event(3); purs_local_1 } else { post_event(4);'), 2 if name == 'after' else 0)
    ins = 'pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {'
    replace(ins, ins + ' post_event(6);', 1)
    return code.replace('std::rc::Rc', 'tracked::Rc'), replacements


def main():
    sources = inputs()
    harness = (PREVIOUS / 'count_harness.rs').read_text()
    harness = harness.replace('tracked::phase("unique_build");', 'tracked::phase("unique_build"); post_reset();')
    harness = harness.replace('tracked::phase("unique_depth_and_drop");', 'post_print(); tracked::phase("unique_depth_and_drop");')
    result = {'method': 'Exact freshly emitted kernel declarations; only Rc API instrumentation and logical event counters added. No algorithm patch and no timings.',
        'checks': 'Four rotations, ordering/red-black invariants, exact keys of 200 retained versions, weak-reference exclusion, global handle and payload lifetime balances.',
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(), 'variants': {}}
    for name, source in sources.items():
        emitted = kernel(source['code'])
        code, replacements = instrument(emitted, name)
        path = BUILD / f'count-{name}.rs'
        path.write_text('#![allow(warnings)]\n' + f'#[path="{PREVIOUS}/tracked_rc.rs"] mod tracked;\n'
            + code + COUNTERS + harness)
        binary = BUILD / f'count-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        (BUILD / f'events-{name}.tsv').write_text(output)
        line, = [line for line in output.splitlines() if line.startswith('POST ')]
        counters = dict(zip(COUNTER_NAMES, map(int, line.split()[1:])))
        summary, _ = previous.summarize('\n'.join(line for line in output.splitlines() if not line.startswith('POST ')))
        totals = defaultdict(int)
        for phase in summary['phases'].values():
            for operation, count in phase.items(): totals[operation] += count
        assert summary['phases']['unique_build']['new'] == 100001
        if name == 'after':
            assert counters['post_no_rotation'] == 1368968
            assert counters['post_rotation'] == 99978
            assert counters['post_unique'] == counters['post_no_rotation'] + counters['post_rotation']
        result['variants'][name] = {key: value for key, value in source.items() if key != 'code'} | {
            'extracted_kernel_sha256': sha256(emitted.encode()), 'unique_build': counters,
            'phases': summary['phases'], 'lifetime_totals': dict(totals), 'instrumented_sites': replacements}
        print(name, counters, summary['phases']['unique_build'], flush=True)
    assert result['variants']['before']['unique_build']['ins_calls'] == result['variants']['after']['unique_build']['ins_calls']
    for source in sources.values():
        assert sha256(Path(source['path']).read_bytes()) == source['sha256'], 'Source changed during checks'
    (HERE / 'counts.json').write_text(json.dumps(result, indent=2) + '\n')


if __name__ == '__main__': main()

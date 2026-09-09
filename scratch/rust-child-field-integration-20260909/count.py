"""Count the actual generated variants separately from native-Rc timings."""
from pathlib import Path
import hashlib
import importlib.util
import json
import subprocess

HERE = Path(__file__).resolve().parent
BUILD = HERE / 'build'
PREVIOUS = HERE.parent / 'rust-perceus-counts-20260909'
spec = importlib.util.spec_from_file_location('previous', PREVIOUS / 'probe.py')
previous = importlib.util.module_from_spec(spec)
spec.loader.exec_module(previous)

COUNTERS = '''
thread_local! { static CHILD_COUNTS: std::cell::Cell<[u64;4]> = const { std::cell::Cell::new([0;4]) }; }
fn child_event(i: usize) { CHILD_COUNTS.with(|s| { let mut a=s.get(); a[i]+=1; s.set(a); }); }
fn child_reset() { CHILD_COUNTS.with(|s|s.set([0;4])); }
fn child_print() { CHILD_COUNTS.with(|s|println!("CHILD {} {} {} {}",s.get()[0],s.get()[1],s.get()[2],s.get()[3])); }
'''

if __name__ == '__main__':
    sources = {
        'before': HERE.parent / 'rust-child-field-20260909/build/RBTree-original.rs',
        'stage1': BUILD / 'RBTree-stage1.rs',
        'stage2': BUILD / 'RBTree-stage2.rs',
    }
    harness = (PREVIOUS / 'count_harness.rs').read_text()
    harness = harness.replace('tracked::phase("unique_build");', 'tracked::phase("unique_build"); child_reset();')
    harness = harness.replace('tracked::phase("unique_depth_and_drop");', 'child_print(); tracked::phase("unique_depth_and_drop");')
    results = {'method': 'Logical Rc API operations, separate from timings; actual generated kernels with the original validation harness.', 'variants': {}}
    for name, source in sources.items():
        code = previous.kernel(source.read_text())
        code = code.replace('pub fn __purust_take(&mut self) -> std::option::Option<Self> {',
            'pub fn __purust_take(&mut self) -> std::option::Option<Self> { child_event(0);')
        code = code.replace('let payload = crate::Tree::T(a0, a1, a2, a3);',
            'child_event(1); let payload = crate::Tree::T(a0, a1, a2, a3);')
        code = code.replace('let _purust_new_child = ', 'child_event(2); let _purust_new_child = ')
        code = code.replace('*_purust_child_slot = ', 'child_event(3); *_purust_child_slot = ')
        path = BUILD / f'count-{name}.rs'
        path.write_text('#![allow(warnings)]\n' + f'#[path="{PREVIOUS}/tracked_rc.rs"] mod tracked;\n'
            + code.replace('std::rc::Rc', 'tracked::Rc') + COUNTERS + harness)
        binary = BUILD / f'count-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        (BUILD / f'events-{name}.tsv').write_text(output)
        line, = [s for s in output.splitlines() if s.startswith('CHILD ')]
        counts = dict(zip(['take', 'rebuild_helper', 'child_call', 'full_parent_write'], map(int, line.split()[1:])))
        summary, _ = previous.summarize('\n'.join(s for s in output.splitlines() if not s.startswith('CHILD ')))
        results['variants'][name] = {'source_sha256': hashlib.sha256(source.read_bytes()).hexdigest(),
            'unique_build_child': counts, 'phases': summary['phases']}
        print(name, counts, summary['phases']['unique_build'], flush=True)
    (HERE / 'counts.json').write_text(json.dumps(results, indent=2) + '\n')

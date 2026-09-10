"""Observe ownership operations separately from timings in each prototype."""
import json
import subprocess
import probe

COUNTERS = '''
thread_local! { static PROBE_COUNTS: std::cell::Cell<[u64;5]> = const { std::cell::Cell::new([0;5]) }; }
fn probe_event(i: usize) { PROBE_COUNTS.with(|s| { let mut a=s.get(); a[i]+=1; s.set(a); }); }
fn probe_reset() { PROBE_COUNTS.with(|s|s.set([0;5])); }
fn probe_print() { PROBE_COUNTS.with(|s|println!("PROBE {} {} {} {} {}",s.get()[0],s.get()[1],s.get()[2],s.get()[3],s.get()[4])); }
'''

if __name__ == '__main__':
    harness = (probe.PREVIOUS / 'count_harness.rs').read_text()
    harness = harness.replace('tracked::phase("unique_build");', 'tracked::phase("unique_build"); probe_reset();')
    harness = harness.replace('tracked::phase("unique_depth_and_drop");', 'probe_print(); tracked::phase("unique_depth_and_drop");')
    results = {'method': 'Logical Rc API operations through the existing one-word wrapper, independent of timings. Four rotations, 200 persistent versions and global lifetime balances checked.', 'variants': {}}
    for name, code in probe.variants().items():
        code = code.replace('pub fn __purust_take(&mut self) -> std::option::Option<Self> {',
            'pub fn __purust_take(&mut self) -> std::option::Option<Self> { probe_event(0);')
        code = code.replace('let payload = crate::Tree::T(a0, a1, a2, a3);',
            'probe_event(1); let payload = crate::Tree::T(a0, a1, a2, a3);')
        code = code.replace('*probe_slot = Tree::T', 'probe_event(2); *probe_slot = Tree::T')
        code = code.replace('if !probe_rotation(probe_slot) {', 'if !probe_rotation(probe_slot) { probe_event(3);')
        code = code.replace('// Keep the existing rotation worker', 'probe_event(4); // Keep the existing rotation worker')
        path = probe.BUILD / f'count-{name}.rs'
        path.write_text('#![allow(warnings)]\n' + f'#[path="{probe.PREVIOUS}/tracked_rc.rs"] mod tracked;\n'
            + code.replace('std::rc::Rc', 'tracked::Rc') + COUNTERS + harness)
        binary = probe.BUILD / f'count-{name}'
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
        output = subprocess.check_output([str(binary)], text=True)
        (probe.BUILD / f'events-{name}.tsv').write_text(output)
        line, = [s for s in output.splitlines() if s.startswith('PROBE ')]
        counts = dict(zip(['take', 'rebuild_helper', 'direct_full_write', 'no_rotation', 'rotation'], map(int, line.split()[1:])))
        summary, _ = probe.previous.summarize('\n'.join(s for s in output.splitlines() if not s.startswith('PROBE ')))
        results['variants'][name] = {'unique_build': counts, 'phases': summary['phases']}
        print(name, counts, summary['phases']['unique_build'], flush=True)
    (probe.HERE / 'counts.json').write_text(json.dumps(results, indent=2) + '\n')

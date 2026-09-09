"""B13 experiment only: leave a unique red parent's unchanged fields in place."""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE/'build'
PREVIOUS = HERE.parent/'rust-perceus-counts-20260909'
spec = importlib.util.spec_from_file_location('previous', PREVIOUS/'probe.py')
previous = importlib.util.module_from_spec(spec)
spec.loader.exec_module(previous)
SOURCE = ROOT/'run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs'
DEPS = ROOT/'run/bak/rust/output/purust_output/target/release/deps'

GUARD = '''
    // A red parent cannot rotate in balance. Equal keys retain the old path.
    let b13_direction = match purs_local_1.as_ref() {
        crate::Tree::T(crate::Color::R, _, key, _) if purs_local_0 < *key => -1,
        crate::Tree::T(crate::Color::R, _, key, _) if purs_local_0 > *key => 1,
        _ => 0,
    };
    if b13_direction != 0 {
        if let Some(b13_node) = std::rc::Rc::get_mut(&mut purs_local_1) {
            /* B13 fast path */
            BODY
        }
    }
'''
FIELD = '''if let crate::Tree::T(_, left, _, right) = b13_node {
                // The sibling is an initialized placeholder, never the child
                // passed recursively: that child must keep its ownership.
                if b13_direction < 0 {
                    let child = std::mem::replace(left, right.clone());
                    *left = Test_RBTree_ins(purs_local_0, child);
                } else {
                    let child = std::mem::replace(right, left.clone());
                    *right = Test_RBTree_ins(purs_local_0, child);
                }
                return purs_local_1;
            } else { unreachable!(); }'''
REBUILD = '''let crate::Tree::T(color, left, key, right) = b13_node.__purust_take().unwrap() else { unreachable!(); };
            if b13_direction < 0 {
                let child = Test_RBTree_ins(purs_local_0, left);
                /* B13 full rebuild */
                *b13_node = crate::Tree::T(color, child, key, right);
            } else {
                let child = Test_RBTree_ins(purs_local_0, right);
                /* B13 full rebuild */
                *b13_node = crate::Tree::T(color, left, key, child);
            }
            return purs_local_1;'''

def variants():
    original = previous.kernel((BUILD/'RBTree-original.rs').read_text())
    signature = 'pub fn Test_RBTree_ins(mut purs_local_0: i64, mut purs_local_1: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {'
    assert original.count(signature) == 1
    return {'before': original,
        'red_rebuild': original.replace(signature, signature + GUARD.replace('BODY', REBUILD)),
        'red_field': original.replace(signature, signature + GUARD.replace('BODY', FIELD))}

HARNESS = '''
#[global_allocator] static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;
fn main() {
    for i in 0..16 {
        let start = std::time::Instant::now();
        let tree = Test_RBTree_buildTree(std::hint::black_box(100000), std::rc::Rc::new(Tree::E));
        assert_eq!(Test_RBTree_depth(tree), 22);
        let elapsed = start.elapsed().as_nanos();
        if i > 0 { println!("{}", elapsed); }
    }
}
'''

def prepare():
    BUILD.mkdir(exist_ok=True)
    if not (BUILD/'RBTree-original.rs').exists():
        (BUILD/'RBTree-original.rs').write_bytes(SOURCE.read_bytes())
    mimalloc, = DEPS.glob('libmimalloc-*.rlib')
    common = ['rustc', '--edition=2021', '-C', 'opt-level=1',
        '--extern', f'mimalloc={mimalloc}', '-L', f'dependency={DEPS}']
    for name, code in variants().items():
        path = BUILD/f'time-{name}.rs'
        path.write_text('#![allow(warnings)]\n'+code+HARNESS)
        subprocess.run(common+[str(path),'-o',str(BUILD/f'time-{name}')],check=True)
        print(name,'compiled',flush=True)
    metadata = {'source':str(SOURCE),
        'source_sha256':hashlib.sha256((BUILD/'RBTree-original.rs').read_bytes()).hexdigest(),
        'rustc':subprocess.check_output(['rustc','--version'],text=True).strip(),
        'binary_sha256':{n:hashlib.sha256((BUILD/f'time-{n}').read_bytes()).hexdigest() for n in variants()},
        'flags':common,'revisions':{str(p):subprocess.check_output(['git','-C',str(p),'rev-parse','HEAD'],text=True).strip() for p in [ROOT,ROOT.parent/'purust/purust']}}
    (HERE/'metadata.json').write_text(json.dumps(metadata,indent=2)+'\n')

def count():
    harness = (PREVIOUS/'count_harness.rs').read_text()
    harness = harness.replace('tracked::phase("unique_build");','tracked::phase("unique_build"); b13_reset();')
    harness = harness.replace('tracked::phase("unique_depth_and_drop");','b13_print(); tracked::phase("unique_depth_and_drop");')
    counters = '''
thread_local! { static B13_COUNTS: std::cell::Cell<[u64;4]> = const { std::cell::Cell::new([0;4]) }; }
fn b13_event(i: usize) { B13_COUNTS.with(|s| { let mut a=s.get(); a[i]+=1; s.set(a); }); }
fn b13_reset() { B13_COUNTS.with(|s|s.set([0;4])); }
fn b13_print() { B13_COUNTS.with(|s|println!("B13 {} {} {} {}",s.get()[0],s.get()[1],s.get()[2],s.get()[3])); }
'''
    data = {}
    for name, code in variants().items():
        code = code.replace('pub fn __purust_take(&mut self) -> std::option::Option<Self> {',
            'pub fn __purust_take(&mut self) -> std::option::Option<Self> { b13_event(0);')
        code = code.replace('let payload = crate::Tree::T(a0, a1, a2, a3);',
            'b13_event(1); let payload = crate::Tree::T(a0, a1, a2, a3);')
        code = code.replace('/* B13 fast path */','b13_event(2);')
        code = code.replace('/* B13 full rebuild */','b13_event(3);')
        path = BUILD/f'count-{name}.rs'
        path.write_text('#![allow(warnings)]\n'+f'#[path="{PREVIOUS}/tracked_rc.rs"] mod tracked;\n'+
            code.replace('std::rc::Rc','tracked::Rc')+counters+harness)
        subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(path),'-o',str(BUILD/f'count-{name}')],check=True)
        output = subprocess.check_output([str(BUILD/f'count-{name}')],text=True)
        (BUILD/f'events-{name}.tsv').write_text(output)
        events = '\n'.join(s for s in output.splitlines() if not s.startswith('B13 '))
        counters_line, = [s for s in output.splitlines() if s.startswith('B13 ')]
        counts = dict(zip(['take','rebuild_helper','red_fast_path','red_full_write'],map(int,counters_line.split()[1:])))
        summary, _ = previous.summarize(events)
        data[name] = {'unique_build_b13':counts,'phases':summary['phases']}
        print(name, counts, summary['phases']['unique_build'],flush=True)
    (HERE/'counts.json').write_text(json.dumps(data,indent=2)+'\n')

def measure(pairs, output):
    runs = {n:[] for n in variants()}
    for pair in range(pairs):
        # A rotating order gives every variant early and late positions.
        names = list(runs)
        names = names[pair%3:]+names[:pair%3]
        if pair % 2: names.reverse()
        for name in names:
            ns = list(map(int,subprocess.check_output([str(BUILD/f'time-{name}')],text=True).split()))
            assert len(ns)==15
            runs[name].append(ns)
            print(pair+1,name,statistics.median(ns)/1e6,'ms',flush=True)
    data = {'method':f'{pairs} rotating/reversed blocks of 3 processes; each 1 warmup + 15 samples; O1/mimalloc, native Rc, 100k insertions, depth and destruction included; no concurrent instrumentation or compilation.',
        'runs_ns':runs,'median_ms':{n:statistics.median(sum(samples,[]))/1e6 for n,samples in runs.items()},
        'process_median_ms':{n:[statistics.median(s)/1e6 for s in samples] for n,samples in runs.items()}}
    (HERE/output).write_text(json.dumps(data,indent=2)+'\n')
    print(json.dumps(data['median_ms'],indent=2),flush=True)

def validate_native():
    base = (PREVIOUS/'count_harness.rs').read_text()
    base = base[base.index('fn validate('):].replace('tracked::Rc','std::rc::Rc')
    base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n','',base)
    base = base.replace('fn main() {','fn original_checks() {')
    extra = (HERE/'validation-extra.rs').read_text()
    results = {}
    for name, code in variants().items():
        path = BUILD/f'validate-{name}.rs'
        path.write_text('#![allow(warnings)]\n'+code+base+extra+'\nfn main() { original_checks(); extra_checks(); }\n')
        subprocess.run(['rustc','--edition=2021','-C','opt-level=1',str(path),'-o',str(BUILD/f'validate-{name}')],check=True)
        output = subprocess.check_output([str(BUILD/f'validate-{name}')],text=True)
        results[name] = output.strip()
        print(name,output.strip(),flush=True)
    (HERE/'validation.json').write_text(json.dumps(results,indent=2)+'\n')

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode',choices=['prepare','count','validate','time'])
    parser.add_argument('--pairs',type=int,default=5)
    parser.add_argument('--output',default='timings.json')
    args=parser.parse_args()
    if args.mode=='prepare': prepare()
    elif args.mode=='count': count()
    elif args.mode=='validate': validate_native()
    else: measure(args.pairs,args.output)

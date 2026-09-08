"""Isolated generated-code experiments; no backend or generated project edits."""
from pathlib import Path
import hashlib, json, os, re, statistics, subprocess, time

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
GEN = ROOT / 'output/purust_output'
original = (GEN / 'Purs_Test_RBTree/src/lib.rs').read_text()
starts = list(re.finditer(r'^pub fn (Test_RBTree_\w+)\(', original, re.M))
functions = {m[1]: original[m.start():starts[i+1].start() if i+1 < len(starts) else len(original)] for i,m in enumerate(starts)}
enums = original[original.index('#[derive(Clone)]'):starts[0].start()]
names = ['max', 'makeBlack', 'depth', 'balance', 'ins', 'insert', 'buildTree']
body = enums + '\n'.join(functions['Test_RBTree_'+name] for name in names)

harness = r'''
use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicU64, Ordering};
static ALLOCS: AtomicU64 = AtomicU64::new(0);
static BYTES: AtomicU64 = AtomicU64::new(0);
static SMALL: AtomicU64 = AtomicU64::new(0);
struct Alloc;
unsafe impl GlobalAlloc for Alloc {
    unsafe fn alloc(&self, l: Layout) -> *mut u8 {
        #[cfg(audit_allocs)] {
            ALLOCS.fetch_add(1, Ordering::Relaxed);
            BYTES.fetch_add(l.size() as u64, Ordering::Relaxed);
            if l.size() == 24 { SMALL.fetch_add(1, Ordering::Relaxed); }
        }
        mimalloc::MiMalloc.alloc(l)
    }
    unsafe fn dealloc(&self, p: *mut u8, l: Layout) { mimalloc::MiMalloc.dealloc(p,l) }
}
#[global_allocator] static GLOBAL: Alloc = Alloc;

fn validate(t: &Tree, min: i64, max: i64) -> (usize, usize) {
    match t {
        Tree::E => (0, 1),
        Tree::T(c, a, x, b) => {
            assert!(min < *x && *x < max);
            if matches!(c.as_ref(), Color::R) {
                for child in [a, b] {
                    if let Tree::T(cc, ..) = child.as_ref() { assert!(matches!(cc.as_ref(), Color::B)); }
                }
            }
            let (na, ha) = validate(a.as_ref(), min, *x);
            let (nb, hb) = validate(b.as_ref(), *x, max);
            assert_eq!(ha, hb);
            (1 + na + nb, ha + if matches!(c.as_ref(), Color::B) {1} else {0})
        }
    }
}
fn correctness() {
    let mut keys: Vec<i64> = (1..=1000).collect();
    let mut seed = 37_u64;
    for i in (1..keys.len()).rev() {
        seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
        keys.swap(i, seed as usize % (i+1));
    }
    let mut tree = std::rc::Rc::new(Tree::E);
    for x in keys { tree = Test_RBTree_insert(x, tree); }
    assert_eq!(validate(tree.as_ref(), 0, 1001).0, 1000);
    let old = tree.clone();
    tree = Test_RBTree_insert(1001, tree);
    assert_eq!(validate(old.as_ref(), 0, 1001).0, 1000);
    assert_eq!(validate(tree.as_ref(), 0, 1002).0, 1001);
    tree = Test_RBTree_insert(501, tree);
    assert_eq!(validate(tree.as_ref(), 0, 1002).0, 1001);
}
fn run(n: i64) -> i64 {
    let tree = Test_RBTree_buildTree(std::hint::black_box(n), std::rc::Rc::new(Tree::E));
    std::hint::black_box(Test_RBTree_depth(tree))
}
fn main() {
    correctness();
    let repeats: usize = std::env::args().nth(1).unwrap_or("5".into()).parse().unwrap();
    assert_eq!(run(100000), 22);
    for _ in 0..repeats {
        ALLOCS.store(0, Ordering::Relaxed); BYTES.store(0, Ordering::Relaxed); SMALL.store(0, Ordering::Relaxed);
        let start = std::time::Instant::now();
        let value = run(100000);
        let elapsed = start.elapsed().as_secs_f64() * 1e6;
        assert_eq!(value, 22);
        println!("{:.3},{},{},{},{},{}", elapsed, ALLOCS.load(Ordering::Relaxed), BYTES.load(Ordering::Relaxed), SMALL.load(Ordering::Relaxed), std::mem::size_of::<Tree>(), std::mem::size_of::<Color>());
    }
}
'''

def color_value(s):
    s = s.replace('#[derive(Clone)]\npub enum Color', '#[derive(Clone, Copy)]\npub enum Color')
    s = s.replace('std::rc::Rc<crate::Color>', 'crate::Color')
    s = re.sub(r'std::rc::Rc::new\(crate::Color::([RB])\)', r'crate::Color::\1', s)
    return s + '\nimpl Color { fn as_ref(&self) -> &Self { self } }\n'

def borrow_tests(s):
    return re.sub(r'\((purs_local_\w+)\.clone\(\)\)\.as_ref\(\)', r'\1.as_ref()', s)

def root_reuse(s):
    for name,args,expr in [('makeBlack','mut tree: std::rc::Rc<crate::Tree>',''),
                            ('insert','x: i64, tree: std::rc::Rc<crate::Tree>','let mut tree = Test_RBTree_ins(x, tree);')]:
        old = functions['Test_RBTree_'+name]
        old = color_value(old).split('\nimpl Color')[0]
        replacement = f'''pub fn Test_RBTree_{name}({args}) -> std::rc::Rc<crate::Tree> {{
            {expr}
            if let Tree::T(ref mut color, ..) = std::rc::Rc::make_mut(&mut tree) {{ *color = Color::B; }}
            tree
        }}\n\n'''
        assert old in s, name
        s = s.replace(old, replacement)
    return s

variants = {
    'generated-o1': (body, 1),
    'generated-o3': (body, 3),
    'color-copy-o1': (color_value(body), 1),
    'color-copy-o3': (color_value(body), 3),
    'borrow-tests-o3': (borrow_tests(body), 3),
    'color-root-reuse-o3': (root_reuse(color_value(body)), 3),
    'color-borrow-o3': (borrow_tests(color_value(body)), 3),
    'compact-balance-o3': (body.replace(functions['Test_RBTree_balance'], (HERE/'compact-balance.rs').read_text()), 3),
    'color-compact-o3': (color_value(body.replace(functions['Test_RBTree_balance'], (HERE/'compact-balance.rs').read_text())), 3),
    'borrow-tests-o1': (borrow_tests(body), 1),
    'color-borrow-o1': (borrow_tests(color_value(body)), 1),
    'compact-balance-o1': (body.replace(functions['Test_RBTree_balance'], (HERE/'compact-balance.rs').read_text()), 1),
    'color-compact-o1': (color_value(body.replace(functions['Test_RBTree_balance'], (HERE/'compact-balance.rs').read_text())), 1),
}
if os.environ.get('RUST_AUDIT_VARIANTS'):
    variants = {name: variants[name] for name in os.environ['RUST_AUDIT_VARIANTS'].split(',')}
count_names = [name for name in ('generated-o3','color-copy-o3','color-root-reuse-o3','color-compact-o3') if name in variants]
deps = GEN / 'target/release/deps'
rlibs = list(deps.glob('libmimalloc-*.rlib'))
assert len(rlibs) == 1, rlibs
meta = {'generated_sha256': hashlib.sha256(original.encode()).hexdigest(),
        'rustc': subprocess.check_output(['rustc','--version'], text=True).strip(),
        'method': '100000 descending insertions + depth + drop; mimalloc; 3 interleaved processes per variant, 1 warm-up + 5 timed runs each; Instant; correctness/persistence outside timing'}
for name,(source,opt) in variants.items():
    (HERE/(name+'.csv')).write_text('')
    (HERE / (name+'.rs')).write_text('#![allow(warnings)]\n'+source+harness)
    command = ['rustc','--edition=2021','-C',f'opt-level={opt}','-C','debuginfo=1',
               '--extern',f'mimalloc={rlibs[0]}','-L',f'dependency={deps}',str(HERE/(name+'.rs')),'-o',str(HERE/name)]
    print('BUILD',name,flush=True)
    subprocess.run(command,check=True,stdout=subprocess.PIPE,stderr=subprocess.PIPE)
    if name in count_names:
        subprocess.run(command[:-1]+[str(HERE/(name+'-counts')),'--cfg','audit_allocs'],check=True,stdout=subprocess.PIPE,stderr=subprocess.PIPE)

samples = {name: [] for name in variants}
order = list(variants)
for trial in range(3):
    for name in order[trial:] + order[:trial]:
        print('MEASURE',trial,name,flush=True)
        raw = subprocess.check_output([str(HERE/name),'5'],text=True)
        (HERE/(name+'.csv')).open('a').write(raw)
        samples[name] += [float(row.split(',')[0]) for row in raw.splitlines()]
meta['results'] = {name: {'n':len(xs), 'min_us':min(xs), 'median_us':statistics.median(xs), 'max_us':max(xs), 'samples_us':xs} for name,xs in samples.items()}
meta['allocations'] = {}
for name in count_names:
    raw = subprocess.check_output([str(HERE/(name+'-counts')),'1'],text=True).strip()
    meta['allocations'][name] = dict(zip(['instrumented_us','allocations','requested_bytes','allocations_size_24','tree_size','color_size'],map(float,raw.split(','))))
(HERE/'results.json').write_text(json.dumps(meta,indent=2)+'\n')
print(json.dumps({k:v for k,v in meta.items() if k != 'results'},indent=2),flush=True)
for name, xs in samples.items(): print(name, 'min',round(min(xs),2),'median',round(statistics.median(xs),2),'max',round(max(xs),2),flush=True)

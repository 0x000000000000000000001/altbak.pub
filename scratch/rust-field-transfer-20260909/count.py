"""Count allocation requests and check persistence on the exact generated kernel."""
from pathlib import Path
import difflib
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
GENERATED = ROOT/'run/bak/rust/output/purust_output'
before = (HERE/'RBTree-before.rs').read_text()
after = (HERE/'RBTree-after.rs').read_text()
(HERE/'RBTree.diff').write_text(''.join(difflib.unified_diff(
    before.splitlines(True), after.splitlines(True), fromfile='RBTree-before.rs', tofile='RBTree-after.rs')))
assert '#[derive(Clone, Copy)]\npub enum Color' in after
assert 'Rc<crate::Color>' not in after
assert 'Rc::new(crate::Color::' not in after
assert 'std::rc::Rc<crate::Tree>' in after

harness = r'''
use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicUsize, Ordering};
static ALLOCS: AtomicUsize = AtomicUsize::new(0);
static FREES: AtomicUsize = AtomicUsize::new(0);
static BYTES: AtomicUsize = AtomicUsize::new(0);
static SMALL: AtomicUsize = AtomicUsize::new(0);
struct Counting;
unsafe impl GlobalAlloc for Counting {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        ALLOCS.fetch_add(1, Ordering::Relaxed);
        BYTES.fetch_add(layout.size(), Ordering::Relaxed);
        if layout.size() == 24 { SMALL.fetch_add(1, Ordering::Relaxed); }
        mimalloc::MiMalloc.alloc(layout)
    }
    unsafe fn dealloc(&self, ptr: *mut u8, layout: Layout) {
        FREES.fetch_add(1, Ordering::Relaxed);
        mimalloc::MiMalloc.dealloc(ptr, layout)
    }
}
#[global_allocator] static ALLOCATOR: Counting = Counting;
fn red(color: &Color) -> bool { matches!(color, Color::R) }
fn validate(tree: &Tree, min: i64, max: i64) -> (usize, usize) {
    match tree {
        Tree::E => (0, 1),
        Tree::T(color, left, key, right) => {
            assert!(min < *key && *key < max);
            if red(color) {
                for child in [left, right] {
                    if let Tree::T(c, ..) = child.as_ref() { assert!(!red(c)); }
                }
            }
            let (nl, hl) = validate(left, min, *key);
            let (nr, hr) = validate(right, *key, max);
            assert_eq!(hl, hr);
            (1 + nl + nr, hl + if red(color) { 0 } else { 1 })
        }
    }
}
fn checked_root(tree: &Tree, count: usize, max: i64) {
    if let Tree::T(color, ..) = tree { assert!(!red(color)); }
    assert_eq!(validate(tree, 0, max).0, count);
}
fn correctness() {
    // The four rotation shapes, then mixed insertion order and a retained root.
    for keys in [[3, 2, 1], [3, 1, 2], [1, 3, 2], [1, 2, 3]] {
        let mut tree = std::rc::Rc::new(Tree::E);
        for (i, key) in keys.into_iter().enumerate() {
            tree = Test_RBTree_insert(key, tree);
            checked_root(&tree, i + 1, 4);
        }
    }
    let mut keys: Vec<i64> = (1..=1000).collect();
    let mut seed = 37_u64;
    for i in (1..keys.len()).rev() {
        seed = seed.wrapping_mul(6364136223846793005).wrapping_add(1);
        keys.swap(i, seed as usize % (i + 1));
    }
    let mut tree = std::rc::Rc::new(Tree::E);
    for key in keys { tree = Test_RBTree_insert(key, tree); }
    checked_root(&tree, 1000, 1001);
    let old = tree.clone();
    tree = Test_RBTree_insert(1001, tree);
    checked_root(&tree, 1001, 1002);
    checked_root(&old, 1000, 1001);
    tree = Test_RBTree_insert(501, tree);
    checked_root(&tree, 1001, 1002);
    drop(tree);
    checked_root(&old, 1000, 1001);
}
fn keys(tree: &Tree) -> Vec<i64> {
    match tree { Tree::E => Vec::new(), Tree::T(_, left, key, right) => {
        let mut result = keys(left); result.push(*key); result.extend(keys(right)); result
    }}
}
fn retained_versions() {
    let mut order: Vec<i64> = (1..=200).collect();
    let mut seed = 41_u64;
    for i in (1..order.len()).rev() { seed=seed.wrapping_mul(6364136223846793005).wrapping_add(1); order.swap(i, seed as usize % (i+1)); }
    let mut root = std::rc::Rc::new(Tree::E);
    let mut expected = Vec::new();
    let mut snapshots = Vec::new();
    for key in order {
        snapshots.push((root.clone(),expected.clone()));
        root=Test_RBTree_insert(key,root);
        expected.push(key); expected.sort();
        checked_root(&root,expected.len(),201);
        assert_eq!(keys(&root),expected);
        for (old,expected) in &snapshots { checked_root(old,expected.len(),201); assert_eq!(&keys(old),expected); }
    }
    drop(root);
    for (old,expected) in &snapshots { checked_root(old,expected.len(),201); assert_eq!(&keys(old),expected); }
}
fn main() {
    correctness();
    retained_versions();
    for counter in [&ALLOCS, &FREES, &BYTES, &SMALL] { counter.store(0, Ordering::Relaxed); }
    let tree = Test_RBTree_buildTree(std::hint::black_box(100000), std::rc::Rc::new(Tree::E));
    assert_eq!(Test_RBTree_depth(tree), 22);
    let allocs = ALLOCS.load(Ordering::Relaxed);
    let frees = FREES.load(Ordering::Relaxed);
    let bytes = BYTES.load(Ordering::Relaxed);
    let small = SMALL.load(Ordering::Relaxed);
    assert_eq!(allocs, frees);
    println!("{} {} {} {} {}", allocs, frees, bytes, small, std::mem::size_of::<Tree>());
}
'''
deps = GENERATED/'target/release/deps'
mimalloc = list(deps.glob('libmimalloc-*.rlib'))
assert len(mimalloc) == 1, mimalloc
results = {}
BUILD = HERE/'build'
BUILD.mkdir(exist_ok=True)
for side, generated in [('before', before), ('after', after)]:
    starts = list(re.finditer(r'^pub fn (Test_RBTree_\w+)\(', generated, re.M))
    functions = {m[1]: generated[m.start():starts[i+1].start() if i+1 < len(starts) else len(generated)]
                 for i, m in enumerate(starts)}
    enums = generated[generated.index('#[derive(Clone'):starts[0].start()]
    names = ['max', 'makeBlack', 'depth', 'balance', 'ins', 'insert', 'buildTree']
    source = BUILD/f'rbtree-{side}-checks.rs'
    binary = BUILD/f'rbtree-{side}-checks'
    source.write_text('#![allow(warnings)]\n'+enums+'\n'.join(functions['Test_RBTree_'+n] for n in names)+harness)
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', '--extern', f'mimalloc={mimalloc[0]}',
                    '-L', f'dependency={deps}', str(source), '-o', str(binary)], check=True)
    output = subprocess.check_output([str(binary)], text=True).strip()
    values = list(map(int, output.split()))
    results[side] = dict(zip(['allocations', 'deallocations', 'requested_bytes', 'allocations_24_bytes', 'tree_size'], values))
    print(side, results[side], flush=True)
results['invariants'] = 'Four rotations, ordering, red/black invariants, duplicates, and the exact keys of 200 simultaneously retained versions checked before and after insertion/destruction on both generated kernels.'
(HERE/'allocations.json').write_text(json.dumps(results, indent=2)+'\n')
print(results['invariants'])

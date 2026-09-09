use std::rc::Rc;
use std::alloc::{GlobalAlloc, Layout, System};
use std::sync::atomic::{AtomicUsize, Ordering};
static ALLOCS: AtomicUsize = AtomicUsize::new(0);
static FREES: AtomicUsize = AtomicUsize::new(0);
struct Counting;
unsafe impl GlobalAlloc for Counting {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 { ALLOCS.fetch_add(1, Ordering::Relaxed); System.alloc(layout) }
    unsafe fn dealloc(&self, ptr: *mut u8, layout: Layout) { FREES.fetch_add(1, Ordering::Relaxed); System.dealloc(ptr, layout) }
}
#[global_allocator] static ALLOCATOR: Counting = Counting;
#[derive(Clone)] enum Tree { Empty, Node(Rc<Tree>, i64, Rc<Tree>) }
fn reconstruct(left: Rc<Tree>, key: i64, right: Rc<Tree>, slot: Option<Rc<Tree>>) -> Rc<Tree> {
    let payload = Tree::Node(left, key + 1, right);
    if let Some(mut slot) = slot {
        if let Some(target) = Rc::get_mut(&mut slot) { *target = payload; return slot; }
    }
    Rc::new(payload)
}
fn caller(mut tree: Rc<Tree>, reuse: bool) -> Rc<Tree> {
    let (payload, slot) = if reuse && Rc::get_mut(&mut tree).is_some() {
        (std::mem::replace(Rc::get_mut(&mut tree).unwrap(), Tree::Empty), Some(tree))
    } else { (Rc::unwrap_or_clone(tree), None) };
    let Tree::Node(left, key, right) = payload else { return Rc::new(Tree::Empty) };
    reconstruct(left, key, right, slot)
}
fn key(tree: &Tree) -> i64 { if let Tree::Node(_, key, _) = tree { *key } else { panic!() } }
fn measure(reuse: bool, shared: bool) -> usize {
    let mut tree = Rc::new(Tree::Node(Rc::new(Tree::Empty), 0, Rc::new(Tree::Empty)));
    let before = ALLOCS.load(Ordering::Relaxed);
    for i in 0..1000 {
        let old = shared.then(|| tree.clone());
        let addr = Rc::as_ptr(&tree);
        tree = caller(tree, reuse);
        assert_eq!(key(&tree), i+1);
        if reuse && !shared { assert_eq!(addr, Rc::as_ptr(&tree)); }
        if let Some(old) = old { assert_eq!(key(&old), i); }
    }
    ALLOCS.load(Ordering::Relaxed)-before
}
fn main() {
    for (reuse, shared) in [(false,false),(true,false),(true,true)] {
        let allocations = ALLOCS.load(Ordering::Relaxed);
        let frees = FREES.load(Ordering::Relaxed);
        let count = measure(reuse,shared);
        assert_eq!(ALLOCS.load(Ordering::Relaxed)-allocations,FREES.load(Ordering::Relaxed)-frees);
        assert_eq!(count, if reuse && !shared {0} else {1000});
        println!("reuse={reuse}, shared={shared}: {count} allocations / 1000 calls; values, address and release counts checked");
    }
}

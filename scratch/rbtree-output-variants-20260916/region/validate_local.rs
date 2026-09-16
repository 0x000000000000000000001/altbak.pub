#![allow(warnings)]
include!("kernel.rs");
#[path="region_local.rs"] mod region_alloc;
use std::rc::Rc;

fn same(a: &Tree, b: &Tree) {
    match (a,b) {
        (Tree::E,Tree::E) => (),
        (Tree::T(ac,al,ak,ar),Tree::T(bc,bl,bk,br)) => {
            assert_eq!(matches!(ac,Color::R),matches!(bc,Color::R));
            assert_eq!(ak,bk); same(al,bl); same(ar,br);
        },
        _ => panic!("shape mismatch"),
    }
}
fn invariants(t: &Tree, min: i64, max: i64) -> (usize,usize) {
    match t {
        Tree::E => (0,1),
        Tree::T(c,l,k,r) => {
            assert!(min < *k && *k < max);
            if matches!(c,Color::R) {
                assert!(!matches!(l.as_ref(),Tree::T(Color::R,..)));
                assert!(!matches!(r.as_ref(),Tree::T(Color::R,..)));
            }
            let (ln,lh)=invariants(l,min,*k); let (rn,rh)=invariants(r,*k,max);
            assert_eq!(lh,rh); (1+ln+rn,lh+usize::from(matches!(c,Color::B)))
        }
    }
}
fn build(keys: &[i64]) -> Rc<Tree> {
    keys.iter().fold(Rc::new(Tree::E), |tree,&key| Test_RBTree_insert(key,tree))
}
fn validate(keys: &[i64], capacity: usize) -> (usize,usize) {
    let baseline=build(keys);
    let region=unsafe { region_alloc::Region::new(capacity) };
    let actual=build(keys);
    same(&baseline,&actual);
    invariants(&actual,i64::MIN,i64::MAX);
    assert!(matches!(actual.as_ref(),Tree::E|Tree::T(Color::B,..)));
    let used=region.used(); let fallback=region.fallbacks();
    drop(actual); drop(region); drop(baseline);
    (used,fallback)
}
fn main() {
    for order in [[3,2,1],[3,1,2],[1,3,2],[1,2,3]] { validate(&order,1024); }
    let ascending: Vec<_>=(1..=100000).collect();
    let descending: Vec<_>=ascending.iter().copied().rev().collect();
    let mut random=ascending.clone(); let mut seed=91_u64;
    for i in (1..random.len()).rev() { seed=seed.wrapping_mul(6364136223846793005).wrapping_add(1); random.swap(i,(seed as usize)%(i+1)); }
    let mut duplicates: Vec<_>=(1..=10000).collect(); duplicates.extend_from_within(..);
    for (name, keys) in [("ascending",&ascending),("descending",&descending),("random",&random),("duplicates",&duplicates)] {
        let (used,fallback)=validate(keys,(keys.len()+1)*64);
        assert_eq!(fallback,0); println!("{name}: exact shape/colors/keys and RB invariants; used={used}, fallback={fallback}");
    }
    let (_,fallback)=validate(&ascending[..2000],64);
    assert!(fallback>0); println!("capacity overflow: system fallback exact and valid");
    for _ in 0..4 { assert_eq!(region_alloc::run(100000),22); }
    // Weak and persistent snapshots remain valid when all are retained inside region.
    let region=unsafe { region_alloc::Region::new(8*1024*1024) };
    let mut snapshots=Vec::new(); let mut root=Rc::new(Tree::E);
    for k in 1..=200 { root=Test_RBTree_insert(k,root); snapshots.push(root.clone()); }
    for (i,t) in snapshots.iter().enumerate() { assert_eq!(invariants(t,i64::MIN,i64::MAX).0,i+1); }
    let weak=Rc::downgrade(&root); drop(root); assert!(weak.upgrade().is_some());
    drop(snapshots); assert!(weak.upgrade().is_none()); drop(weak); drop(region);
    // Exercise nontrivial requested alignment without using allocator recursion.
    let region=unsafe { region_alloc::Region::new(16384) };
    for align in [1,8,64,128,4096] { let l=std::alloc::Layout::from_size_align(17,align).unwrap(); unsafe { let p=std::alloc::alloc(l); assert!(!p.is_null()); assert_eq!(p as usize%align,0); std::alloc::dealloc(p,l); } }
    drop(region);
    println!("4 fresh region reinitializations; retained snapshots; Weak; alignment: passed");
}

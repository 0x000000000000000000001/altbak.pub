// Include after an extracted kernel in the crate root.
mod reference { include!("reference.rs"); }
fn fingerprint(t: &Tree, out: &mut Vec<(i64,u8)>) {
    match t {
        Tree::E => out.push((0, 2)),
        Tree::T(c,l,k,r) => {
            out.push((*k, if matches!(c,Color::B) {1} else {0}));
            fingerprint(l,out); fingerprint(r,out);
        }
    }
}
fn reference_fingerprint(t: &reference::Tree, out: &mut Vec<(i64,u8)>) {
    match t {
        reference::Tree::E => out.push((0,2)),
        reference::Tree::T(c,l,k,r) => {
            out.push((*k, if matches!(c,reference::Color::B) {1} else {0}));
            reference_fingerprint(l,out); reference_fingerprint(r,out);
        }
    }
}
fn validate(t: &Tree, low: i64, high: i64, keys: &mut Vec<i64>) -> (usize,i64) {
    match t {
        Tree::E => (1,0),
        Tree::T(c,l,k,r) => {
            assert!(low < *k && *k < high);
            if matches!(c,Color::R) {
                for child in [l,r] {
                    if let Tree::T(cc,..) = child.as_ref() { assert!(matches!(cc,Color::B)); }
                }
            }
            let (lh,ld) = validate(l,low,*k,keys);
            keys.push(*k);
            let (rh,rd) = validate(r,*k,high,keys);
            assert_eq!(lh,rh);
            (lh+usize::from(matches!(c,Color::B)),1+ld.max(rd))
        }
    }
}
fn check_pair(t: &Tree, reference: &reference::Tree, expected: &[i64]) -> i64 {
    if let Tree::T(c,..) = t { assert!(matches!(c,Color::B)); }
    let mut keys=Vec::new(); let (_,depth)=validate(t,i64::MIN,i64::MAX,&mut keys);
    assert_eq!(keys,expected);
    let mut actual=Vec::new(); let mut want=Vec::new();
    fingerprint(t,&mut actual); reference_fingerprint(reference,&mut want);
    assert_eq!(actual,want,"complete shape and colors differ");
    depth
}
fn test_order(name: &str, order: Vec<i64>) {
    let mut t=std::rc::Rc::new(Tree::E);
    let mut reference=std::rc::Rc::new(reference::Tree::E);
    for &k in &order {
        t=Test_RBTree_insert(k,t);
        reference=reference::Test_RBTree_insert(k,reference);
    }
    let mut expected=order; expected.sort_unstable(); expected.dedup();
    let depth=check_pair(&t,&reference,&expected);
    assert_eq!(Test_RBTree_depth(t),depth);
    println!("{name}\tkeys={}\tdepth={depth}\texact_shape=true",expected.len());
}
fn main() {
    for order in [[3,2,1],[3,1,2],[1,3,2],[1,2,3]] { test_order("rotation",order.to_vec()); }
    test_order("descending",(1..=100000).rev().collect());
    test_order("ascending",(1..=100000).collect());
    let mut shuffled: Vec<i64>=(1..=100000).collect(); let mut rng=0x6a09e667f3bcc909u64;
    for i in (1..shuffled.len()).rev() { rng^=rng<<13; rng^=rng>>7; rng^=rng<<17; shuffled.swap(i,rng as usize%(i+1)); }
    test_order("shuffled",shuffled);
    test_order("duplicates",(0..100000).map(|k|(k*73)%4096+1).collect());
    let mut t=std::rc::Rc::new(Tree::E);
    let mut reference=std::rc::Rc::new(reference::Tree::E);
    let mut versions=Vec::new(); let mut expected=Vec::new();
    for n in 0..200 {
        versions.push((t.clone(),reference.clone(),expected.clone()));
        let k=(n*73)%200+1;
        t=Test_RBTree_insert(k,t); reference=reference::Test_RBTree_insert(k,reference);
        expected.push(k); expected.sort_unstable();
        check_pair(&t,&reference,&expected);
        for (old,old_ref,keys) in &versions { check_pair(old,old_ref,keys); }
    }
    println!("persistent_200\tok");
    let mut t=std::rc::Rc::new(Tree::T(Color::B,std::rc::Rc::new(Tree::E),1,std::rc::Rc::new(Tree::E)));
    let weak=std::rc::Rc::downgrade(&t);
    assert!(std::rc::Rc::get_mut(&mut t).is_none());
    t=Test_RBTree_insert(2,t);
    let mut keys=Vec::new(); validate(&t,i64::MIN,i64::MAX,&mut keys); assert_eq!(keys,vec![1,2]);
    assert!(weak.upgrade().is_none());
    println!("weak\tok");
}

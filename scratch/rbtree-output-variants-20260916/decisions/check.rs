fn exact(a: &Tree, b: &Tree) {
    match (a,b) {
        (Tree::E,Tree::E) => (),
        (Tree::T(c,l,k,r),Tree::T(d,m,j,n)) => { assert_eq!(matches!(c,Color::R),matches!(d,Color::R)); assert_eq!(k,j); exact(l,m); exact(r,n); },
        _ => panic!("shape differs"),
    }
}
fn valid(t: &Tree, lo:i64, hi:i64) -> (usize,usize) {
    match t {
        Tree::E => (0,1),
        Tree::T(c,l,k,r) => {
            assert!(lo < *k && *k < hi);
            if matches!(c,Color::R) { for v in [l,r] { assert!(!matches!(v.as_ref(),Tree::T(Color::R,..))); } }
            let (ln,lh)=valid(l,lo,*k); let (rn,rh)=valid(r,*k,hi); assert_eq!(lh,rh);
            (ln+rn+1,lh+usize::from(matches!(c,Color::B)))
        }
    }
}
fn run(keys:&[i64], snapshots:bool) {
    let mut a=std::rc::Rc::new(Tree::E); let mut b=std::rc::Rc::new(Tree::E);
    let mut versions=Vec::new(); let mut set=std::collections::BTreeSet::new();
    for (i,&key) in keys.iter().enumerate() {
        if snapshots { versions.push((a.clone(),b.clone())); }
        a=Test_RBTree_insert(key,a); b=reference::Test_RBTree_insert(key,b); set.insert(key);
        if keys.len() < 300 || i%997==0 { exact(&a,&b); assert_eq!(valid(&a,i64::MIN,i64::MAX).0,set.len()); }
    }
    exact(&a,&b); assert_eq!(valid(&a,i64::MIN,i64::MAX).0,set.len());
    assert_eq!(Test_RBTree_depth(a),reference::Test_RBTree_depth(b));
    for (old,expect) in versions { exact(&old,&expect); valid(&old,i64::MIN,i64::MAX); }
}
fn main() {
    // Exhaust all shape/color combinations read by the original guard, including invalid RB trees.
    let leaves=[std::rc::Rc::new(Tree::E),std::rc::Rc::new(Tree::T(Color::R,std::rc::Rc::new(Tree::E),1,std::rc::Rc::new(Tree::E))),std::rc::Rc::new(Tree::T(Color::B,std::rc::Rc::new(Tree::E),1,std::rc::Rc::new(Tree::E)))];
    let mut subtrees=vec![std::rc::Rc::new(Tree::E)];
    for c in [Color::R,Color::B] {for l in &leaves {for r in &leaves {subtrees.push(std::rc::Rc::new(Tree::T(c,l.clone(),2,r.clone())));}}}
    for c in [Color::R,Color::B] {for l in &subtrees {for r in &subtrees {assert_eq!(Test_RBTree_balance__purust_child_rebuilds(&c,l,&3,r),reference::Test_RBTree_balance__purust_child_rebuilds(&c,l,&3,r));}}}
    for order in [[3,2,1],[3,1,2],[1,3,2],[1,2,3]] {run(&order,false);run(&order,true);}
    let asc=(1..=100000).collect::<Vec<i64>>(); run(&asc,false);
    let desc=asc.iter().copied().rev().collect::<Vec<_>>(); run(&desc,false);
    let mut shuffled=(1..=10000).collect::<Vec<i64>>(); let mut seed=123456789u64;
    for i in (1..shuffled.len()).rev() {seed=seed.wrapping_mul(6364136223846793005).wrapping_add(1);shuffled.swap(i,(seed as usize)%(i+1));}
    run(&shuffled,false);run(&shuffled[..200],true);
    let dup=shuffled.iter().take(1000).copied().cycle().take(3000).collect::<Vec<_>>();run(&dup,false);
    let mut a=Test_RBTree_buildTree(100,std::rc::Rc::new(Tree::E));
    let mut b=reference::Test_RBTree_buildTree(100,std::rc::Rc::new(Tree::E));
    let wa=std::rc::Rc::downgrade(&a);let wb=std::rc::Rc::downgrade(&b);
    a=Test_RBTree_insert(101,a);b=reference::Test_RBTree_insert(101,b);exact(&a,&b);
    assert_eq!(wa.upgrade().is_some(),wb.upgrade().is_some());
    if let (Tree::T(_,l,_,_),Tree::T(_,m,_,_))=(a.as_ref(),b.as_ref()) {
       let wa=std::rc::Rc::downgrade(l);let wb=std::rc::Rc::downgrade(m);
       a=Test_RBTree_insert(-1,a);b=reference::Test_RBTree_insert(-1,b);exact(&a,&b);
       match (wa.upgrade(),wb.upgrade()) {(Some(x),Some(y))=>exact(&x,&y),(None,None)=>(),_=>panic!("weak behavior differs")}
    }
    println!("PASS exhaustive 722 guard color/shape combinations; exact shapes/colors/keys/depth, all rotations, asc/desc 100k, shuffle10k, duplicates, 200 retained snapshots, root and child Weak");
}

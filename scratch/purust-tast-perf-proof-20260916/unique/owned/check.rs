#![allow(warnings)]
include!("kernel.rs");
mod generated { include!("generated.rs"); }
fn validate(link: &Link, low: i64, high: i64, keys: &mut Vec<i64>) -> usize {
    let Some(node) = link else { return 1; };
    assert!(low < node.key && node.key < high);
    if node.color == Color::R { assert!(!red(&node.left) && !red(&node.right)); }
    let lh = validate(&node.left, low, node.key, keys);
    keys.push(node.key);
    let rh = validate(&node.right, node.key, high, keys);
    assert_eq!(lh, rh);
    lh + usize::from(node.color == Color::B)
}
fn shape(link: &Link, output: &mut Vec<Option<(bool,i64)>>) {
    match link {
        None => output.push(None),
        Some(node) => { output.push(Some((node.color == Color::R, node.key))); shape(&node.left, output); shape(&node.right, output); }
    }
}
fn generated_shape(tree: &generated::Tree, output: &mut Vec<Option<(bool,i64)>>) {
    match tree {
        generated::Tree::E => output.push(None),
        generated::Tree::T(c,left,key,right) => { output.push(Some((matches!(c,generated::Color::R),*key))); generated_shape(left,output); generated_shape(right,output); }
    }
}
fn check(label: &str, order: Vec<i64>) {
    let mut expected = order.clone(); expected.sort_unstable(); expected.dedup();
    let mut tree = None;
    let mut reference = std::rc::Rc::new(generated::Tree::E);
    for key in order { insert(key, &mut tree); reference = generated::Test_RBTree_insert(key,reference); }
    assert!(!red(&tree));
    let mut actual = Vec::new(); validate(&tree, i64::MIN,i64::MAX,&mut actual); assert_eq!(actual,expected);
    let mut s = Vec::new(); let mut r = Vec::new(); shape(&tree,&mut s); generated_shape(&reference,&mut r); assert_eq!(s,r,"shape mismatch: {label}");
    assert_eq!(depth(&tree),generated::Test_RBTree_depth(reference));
    println!("{label}\tkeys={}\tdepth={}\texact_generated_shape=true",expected.len(),depth(&tree));
}
fn main() {
    for order in [[3,2,1],[3,1,2],[1,3,2],[1,2,3]] { check("rotation",order.to_vec()); }
    check("descending_100k",(1..=100000).rev().collect());
    check("ascending_100k",(1..=100000).collect());
    let mut shuffled:Vec<i64> = (1..=100000).collect(); let mut rng=0x71e9af148634u64;
    for i in (1..shuffled.len()).rev() { rng^=rng<<13; rng^=rng>>7; rng^=rng<<17; shuffled.swap(i,(rng as usize)%(i+1)); }
    check("shuffled_100k",shuffled);
    check("duplicates_100k",(0..100000).map(|n| (n*73)%4096+1).collect());
    #[cfg(rc_owner)]
    {
        let mut tree = None;
        let mut snapshots = Vec::new();
        for n in 0..200 {
            let mut previous = Vec::new(); shape(&tree,&mut previous); snapshots.push((tree.clone(), previous));
            insert((n*73)%200+1,&mut tree);
            for (old,wanted) in &snapshots { let mut actual=Vec::new(); shape(old,&mut actual); assert_eq!(&actual,wanted); }
        }
        println!("persistent_snapshots_200\tok");
    }
    assert_eq!(run(100000),22);
}

#![allow(warnings)]
include!("kernel.rs");
mod generated { include!("generated.rs"); }
fn validate(tree: &Arena, root: Index, low: i64, high: i64, keys: &mut Vec<i64>) -> usize {
    if root == 0 { return 1; }
    let node = tree.nodes[root as usize];
    assert!(low < node.key && node.key < high);
    if node.red { assert!(!tree.red(node.left) && !tree.red(node.right)); }
    let lh = validate(tree,node.left,low,node.key,keys);
    keys.push(node.key);
    let rh = validate(tree,node.right,node.key,high,keys);
    assert_eq!(lh,rh);
    lh + usize::from(!node.red)
}
fn shape(tree: &Arena, root: Index, output: &mut Vec<Option<(bool,i64)>>) {
    if root == 0 { output.push(None); return; }
    let node = tree.nodes[root as usize];
    output.push(Some((node.red,node.key)));
    shape(tree,node.left,output); shape(tree,node.right,output);
}
fn generated_shape(tree: &generated::Tree, output: &mut Vec<Option<(bool,i64)>>) {
    match tree {
        generated::Tree::E => output.push(None),
        generated::Tree::T(c,l,k,r) => {
            output.push(Some((matches!(c,generated::Color::R),*k)));
            generated_shape(l,output); generated_shape(r,output);
        }
    }
}
fn check(label: &str, order: Vec<i64>) {
    let mut expected=order.clone(); expected.sort_unstable(); expected.dedup();
    let mut tree=Arena::new(0);
    let mut reference=std::rc::Rc::new(generated::Tree::E);
    for key in order { tree.insert(key); reference=generated::Test_RBTree_insert(key,reference); }
    assert!(!tree.red(tree.root));
    let mut actual=Vec::new(); validate(&tree,tree.root,i64::MIN,i64::MAX,&mut actual); assert_eq!(actual,expected);
    let mut a=Vec::new(); let mut b=Vec::new();
    shape(&tree,tree.root,&mut a); generated_shape(&reference,&mut b); assert_eq!(a,b,"shape mismatch: {label}");
    assert_eq!(tree.depth(tree.root),generated::Test_RBTree_depth(reference));
    assert_eq!(tree.nodes.len(),expected.len()+1);
    println!("{label}\tkeys={}\tdepth={}\texact_generated_shape=true",expected.len(),tree.depth(tree.root));
}
fn main() {
    for order in [[3,2,1],[3,1,2],[1,3,2],[1,2,3]] { check("rotation",order.to_vec()); }
    check("empty",vec![]);
    check("descending_100k",(1..=100000).rev().collect());
    check("ascending_100k",(1..=100000).collect());
    let mut shuffled: Vec<i64>=(1..=100000).collect(); let mut rng=0x71e9af148634u64;
    for i in (1..shuffled.len()).rev() { rng^=rng<<13; rng^=rng>>7; rng^=rng<<17; shuffled.swap(i,(rng as usize)%(i+1)); }
    check("shuffled_100k",shuffled);
    check("duplicates_100k",(0..100000).map(|n| (n*73)%4096+1).collect());
    assert_eq!(run(100000),22); assert_eq!(run(0),0);
    println!("node_bytes={}\tindex_bytes={}",std::mem::size_of::<Node>(),std::mem::size_of::<Index>());
}

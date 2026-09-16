#![allow(warnings)]
#![recursion_limit="512"]
include!("/Users/0x1/Documents/htdocs/altbak.pub/scratch/purust-tast-perf-proof-20260916/borrow/kernel-borrow-ins.rs");

#[inline(never)]
fn run(n: i64) -> i64 {
    let tree = Test_RBTree_buildTree(n, std::rc::Rc::new(Tree::E));
    Test_RBTree_depth(std::hint::black_box(tree))
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    let n: i64 = args[1].parse().unwrap();
    let count: usize = args[2].parse().unwrap();
    for sample in 0..count+3 {
        let begin = std::time::Instant::now();
        let value = run(std::hint::black_box(n));
        std::hint::black_box(value);
        let elapsed = begin.elapsed().as_nanos();
        assert_eq!(value, 22);
        if sample >= 3 { println!("{}", elapsed); }
    }
}

#[global_allocator]
static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;

fn timing_list(n: usize, seed: i64) -> Rc<List> {
    let mut out = Rc::new(List::Empty);
    for i in (0..n).rev() { out = Rc::new(List::Node(seed + i as i64, out)); }
    out
}
fn timing_tree(depth: usize, seed: i64) -> Rc<Tree> {
    if depth == 0 { return Rc::new(Tree::Empty); }
    Rc::new(Tree::Node(timing_tree(depth-1, seed+1), seed,
        timing_tree(depth-1, seed+(1_i64 << (depth-1)))))
}
fn list_metric(tree: &List) -> (i64, i64) {
    match tree {
        List::Empty => (0, 0),
        List::Node(key, next) => { let (sum, count) = list_metric(next); (key+sum, count+1) }
    }
}
fn tree_metric(tree: &Tree) -> (i64, i64) {
    match tree {
        Tree::Empty => (0, 0),
        Tree::Node(left, key, right) => {
            let (ls, ln) = tree_metric(left); let (rs, rn) = tree_metric(right);
            (ls+key+rs, ln+1+rn)
        }
    }
}
const CASES: [&str;5] = ["list_map", "list_filter", "tree_map", "tree_choose", "tree_update"];
fn repetitions(case: &str) -> usize {
    match case { "tree_choose" | "tree_update" => 32768, _ => 2048 }
}
fn workload(case: &str, mode: usize, seed: i64, loops: usize) -> (i64, i64) {
    let mut total = (0, 0);
    for iteration in 0..std::hint::black_box(loops) {
        let start = std::hint::black_box((seed + iteration as i64) & 1023);
        let metric;
        if case.starts_with("list_") {
            let input = timing_list(std::hint::black_box(256), start);
            let retained = match mode { 1 => Some(input.clone()), 2 => Some(first_List(&input)), _ => None };
            let transform: fn(Rc<List>)->Rc<List> = if case=="list_map" { list_map } else { list_filter };
            let result = transform(std::hint::black_box(input));
            metric = std::hint::black_box(list_metric(result.as_ref()));
            drop(result); drop(retained);
        } else {
            let depth = if case=="tree_map" { 7 } else { 3 };
            let input = timing_tree(std::hint::black_box(depth), start);
            let retained = match mode { 1 => Some(input.clone()), 2 => Some(first_Tree(&input)), _ => None };
            let transform: fn(Rc<Tree>)->Rc<Tree> = match case { "tree_map" => tree_map, "tree_choose" => tree_choose, _ => tree_update };
            let result = transform(std::hint::black_box(input));
            metric = std::hint::black_box(tree_metric(result.as_ref()));
            drop(result); drop(retained);
        }
        total.0 += metric.0; total.1 += metric.1;
    }
    std::hint::black_box(total)
}
fn expected(case: &str, seed: i64, loops: usize) -> (i64, i64) {
    let mut total = (0, 0);
    for iteration in 0..loops {
        let start = (seed + iteration as i64) & 1023;
        let (sum, count) = match case {
            "list_map" => (256*(start+1) + 256*255/2, 256),
            "list_filter" => { let first = start+(start&1); (128*first + 128*127, 128) },
            "tree_map" => (127*(start+1) + 127*126/2, 127),
            "tree_update" => (7*start + 7*6/2 + 1, 7),
            "tree_choose" => { let first = start + if start%2==0 {1} else {4}; (3*first + 3, 3) },
            _ => panic!("invalid case"),
        };
        total.0 += sum; total.1 += count;
    }
    total
}
fn main() {
    let args: Vec<_> = std::env::args().collect();
    match args.get(1).map(String::as_str) {
        Some("smoke") => {
            for case in CASES { for mode in 0..3 {
                for seed in [17, 18] { assert_eq!(workload(case, mode, seed, 2), expected(case, seed, 2)); }
            }}
            println!("PASS smoke: 5 workloads x 3 ownership modes x 2 dynamic seeds; no clocks");
        }
        Some("measure") => {
            use std::io::Write;
            let samples:usize = args[2].parse().unwrap();
            let seed:i64 = args[3].parse().unwrap();
            let divisor:usize = args[4].parse().unwrap();
            assert!(samples>0 && divisor>0);
            for case in CASES { for mode in 0..3 {
                let loops = (repetitions(case)/divisor).max(1);
                let reference = expected(case, seed, loops);
                assert_eq!(workload(case, mode, seed, loops), reference);
                for sample in 0..samples {
                    let start = std::time::Instant::now();
                    let result = workload(case, mode, seed, loops);
                    let ns = start.elapsed().as_nanos();
                    assert_eq!(result, reference);
                    println!("TIME {case} {mode} {sample} {loops} {ns}");
                    std::io::stdout().flush().unwrap();
                }
            }}
        }
        _ => panic!("Pass smoke or measure explicitly"),
    }
}

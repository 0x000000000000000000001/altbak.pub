#![allow(warnings)]
#![recursion_limit="512"]
include!("/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rbtree-output-variants-20260916/baseline.rs");
#[path="/Users/0x1/Documents/htdocs/altbak.pub-purust/scratch/rbtree-output-variants-20260916/region/region_alloc.rs"] mod region_alloc;
#[inline(never)] fn run(n:i64)->i64 { region_alloc::run(n) }

fn main() {
 let args: Vec<String>=std::env::args().collect();
 let n: i64=args[1].parse().unwrap();
 let samples: usize=args[2].parse().unwrap();
 let expected: i64=args[3].parse().unwrap();
 for sample in 0..samples+3 {
  let start=std::time::Instant::now();
  let result=run(std::hint::black_box(n));
  std::hint::black_box(result);
  let ns=start.elapsed().as_nanos();
  assert_eq!(result,expected);
  if sample>=3 { println!("{}",ns); }
 }
}

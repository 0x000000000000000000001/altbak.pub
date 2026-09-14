#![allow(warnings)]
use std::hint::black_box;
use std::rc::Rc;
use Purs_Test_ListOps::{List, Test_ListOps_filterEvens};

#[global_allocator]
static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;

fn input(n: usize, ratio: usize) -> Rc<List> {
    let mut root=Rc::new(List::Nil);
    for i in (0..n).rev() {
        let value=match ratio {0=>(2*i+1) as i64,100=>(2*i+2) as i64,_=>(i+1) as i64};
        root=Rc::new(List::Cons(purust_core::mk_int(value),root));
    }
    root
}
fn tail_at(mut root: &Rc<List>, index: usize) -> &Rc<List> {
    for _ in 0..index { if let List::Cons(_,tail)=root.as_ref(){root=tail;} }
    root
}
fn sum(root: &Rc<List>) -> i64 {
    let mut result=0; let mut cursor=root;
    while let List::Cons(value,tail)=cursor.as_ref() {
        result+=value.unwrap_int();cursor=tail;
    }
    result
}
fn values(root: &Rc<List>) -> Vec<i64> {
    let mut result=Vec::new();let mut cursor=root;
    while let List::Cons(value,tail)=cursor.as_ref(){result.push(value.unwrap_int());cursor=tail;}
    result
}
fn one(n: usize, ratio: usize, scenario: &str, validate: bool) -> i64 {
    // The complete input, owners and Weak are created inside this operation.
    let root=black_box(input(black_box(n),black_box(ratio)));
    let retained_root=(scenario=="root-retained").then(||root.clone());
    let retained_tail=(scenario=="interior-retained").then(||tail_at(&root,n/3).clone());
    // For ratio 50 this observes an even, retained element, so Weak really
    // excludes a possible reuse instead of observing only a discarded head.
    let weak_index=if n>1 {((n/3)|1).min(n-1)} else {0};
    let weak=(scenario=="weak").then(||Rc::downgrade(tail_at(&root,weak_index)));
    let original=validate.then(||values(&root));
    let output=black_box(Test_ListOps_filterEvens(root));
    let result=black_box(sum(black_box(&output)));
    if let Some(original)=original {
        let expected: Vec<_>=original.iter().rev().copied().filter(|v|v%2==0).collect();
        assert_eq!(values(&output),expected);
        if let Some(ref old)=retained_root{assert_eq!(values(old),original);}
        if let Some(ref old)=retained_tail{assert_eq!(values(old),original[n/3..]);}
        if let Some(ref weak)=weak{assert!(weak.upgrade().is_none());}
    }
    // No surviving List owner leaves the timed operation. Weak deallocation
    // and destruction of both output and retained versions are included.
    drop(output);drop(retained_root);drop(retained_tail);drop(weak);
    result
}
fn batch(n: usize, scenario: &str, iterations: usize) -> i64 {
    let mut checksum=0;
    for _ in 0..iterations {checksum+=black_box(one(n,50,scenario,false));}
    black_box(checksum)
}
fn run(args: Vec<String>) {
    let variant=&args[1];let mode=&args[2];
    if mode=="smoke" {
        let mut cases=0;
        for n in [0,1,2,31,900,9000] {for ratio in [0,50,100] {
            for scenario in ["unique","root-retained","interior-retained","weak"] {
                let expected=match ratio {0=>0,100=>(n*(n+1)) as i64,_=>((n/2)*(n/2+1)) as i64};
                assert_eq!(one(n,ratio,scenario,true),expected);cases+=1;
            }
        }}
        println!("{{\"variant\":\"{variant}\",\"smoke_passed\":true,\"cases\":{cases},\"clock_read\":false}}");
        return;
    }
    assert_eq!(mode,"time");
    assert_eq!(args[3],"--coordinated");
    let n:usize=args[4].parse().unwrap();let scenario=&args[5];
    let iterations:usize=args[6].parse().unwrap();
    assert!([900,9000].contains(&n));
    assert!(["unique","root-retained","interior-retained","weak"].contains(&scenario.as_str()));
    assert!(iterations>0);
    let expected=((n/2)*(n/2+1)*iterations) as i64;
    // One full batch is warmed up before the first read of Instant.
    assert_eq!(batch(n,scenario,iterations),expected);
    let mut samples=Vec::with_capacity(7);
    for sample in 0..7 {
        let start=std::time::Instant::now();
        let checksum=batch(n,scenario,iterations);
        let nanos=start.elapsed().as_nanos();
        assert_eq!(checksum,expected);
        samples.push(nanos);
    }
    println!("{{\"variant\":\"{}\",\"n\":{},\"ratio\":50,\"scenario\":\"{}\",\"iterations\":{},\"warmup_batches\":1,\"samples_ns\":{:?}}}",variant,n,scenario,iterations,samples);
}
fn main() {
    let args=std::env::args().collect();
    std::thread::Builder::new().stack_size(64*1024*1024).spawn(move||run(args)).unwrap().join().unwrap();
}

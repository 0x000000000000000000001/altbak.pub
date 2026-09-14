#![allow(warnings)]
use std::alloc::{GlobalAlloc, Layout, System};
use std::rc::Rc;
use std::sync::atomic::{AtomicBool, AtomicU64, AtomicUsize, Ordering::Relaxed};
use Purs_Test_ListOps::{List, Test_ListOps_filterEvens, Test_ListOps_sumEvens};
#[cfg(list_counted)]
use Purs_Test_ListOps::{probe_filter_counts, probe_list_new};

// Register the allocator's own raw pointer only while Rc::new(List) runs.
// The fixed address table needs no allocation and creates no Weak observers,
// which would otherwise disable the unique reuse path being measured.
const CAPACITY: usize = 65536;
static ADDRESSES: [AtomicUsize; CAPACITY] = [const { AtomicUsize::new(0) }; CAPACITY];
static NEXT_LIST: AtomicBool = AtomicBool::new(false);
static ALLOCS: AtomicU64 = AtomicU64::new(0);
static FREES: AtomicU64 = AtomicU64::new(0);
static LIST_ALLOCS: AtomicU64 = AtomicU64::new(0);
static LIST_FREES: AtomicU64 = AtomicU64::new(0);
fn register(address: usize) {
    let start = (address >> 4) % CAPACITY;
    for offset in 0..CAPACITY {
        let slot = &ADDRESSES[(start+offset)%CAPACITY];
        if slot.load(Relaxed) <= 1 { slot.store(address,Relaxed); return; }
    }
    panic!("list allocation registry full");
}
fn unregister(address: usize) -> bool {
    let start = (address >> 4) % CAPACITY;
    for offset in 0..CAPACITY {
        let slot = &ADDRESSES[(start+offset)%CAPACITY];
        let found = slot.load(Relaxed);
        if found == 0 { return false; }
        if found == address { slot.store(1,Relaxed); return true; }
    }
    false
}
struct Counter;
unsafe impl GlobalAlloc for Counter {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        let mark = NEXT_LIST.swap(false,Relaxed);
        let pointer = System.alloc(layout);
        ALLOCS.fetch_add(1,Relaxed);
        if mark { register(pointer as usize); LIST_ALLOCS.fetch_add(1,Relaxed); }
        pointer
    }
    unsafe fn dealloc(&self, pointer: *mut u8, layout: Layout) {
        if unregister(pointer as usize) { LIST_FREES.fetch_add(1,Relaxed); }
        FREES.fetch_add(1,Relaxed);
        System.dealloc(pointer,layout);
    }
}
#[global_allocator]
static GLOBAL: Counter = Counter;
#[no_mangle]
pub extern "C" fn probe_mark_list_allocation() {
    assert!(!NEXT_LIST.swap(true,Relaxed));
}
#[cfg(not(list_counted))]
fn probe_list_new(payload: List) -> Rc<List> { Rc::new(payload) }
#[cfg(not(list_counted))]
fn probe_filter_counts() -> [u64;7] { [0;7] }

fn heap() -> [u64;4] { [ALLOCS.load(Relaxed),FREES.load(Relaxed),LIST_ALLOCS.load(Relaxed),LIST_FREES.load(Relaxed)] }
fn delta(after: [u64;4], before: [u64;4]) -> [u64;4] { std::array::from_fn(|i| after[i]-before[i]) }
fn values(xs: &Rc<List>) -> Vec<i64> {
    let mut result=Vec::new(); let mut cursor=xs;
    while let List::Cons(value,tail)=cursor.as_ref() { result.push(value.unwrap_int());cursor=tail; }
    result
}
fn input(n: usize, ratio: usize) -> Rc<List> {
    let mut result=probe_list_new(List::Nil);
    for i in (0..n).rev() {
        let value=match ratio {0=>(2*i+1) as i64,100=>(2*i+2) as i64,_=>(i+1) as i64};
        result=probe_list_new(List::Cons(purust_core::mk_int(value),result));
    }
    result
}
fn tail_at(mut root: &Rc<List>, index: usize) -> &Rc<List> {
    for _ in 0..index { if let List::Cons(_,tail)=root.as_ref(){root=tail;} }
    root
}
fn one_case(variant: &str, n: usize, ratio: usize, mask: usize) -> ([u64;7],[u64;4],usize) {
    let root=input(n,ratio);
    let pivot=n/3;
    let tail=tail_at(&root,pivot);
    let old_tail=(mask&2!=0).then(||tail.clone());
    let weak_tail=(mask&8!=0).then(||Rc::downgrade(tail));
    let original=values(&root);
    let tail_original=values(tail);
    let old=(mask&1!=0).then(||root.clone());
    let weak=(mask&4!=0).then(||Rc::downgrade(&root));
    let expected:Vec<_>=original.iter().copied().rev().filter(|v|v%2==0).collect();
    let before=heap();probe_filter_counts();
    let result=Test_ListOps_filterEvens(root);
    let call_heap=delta(heap(),before);
    let paths=probe_filter_counts();
    assert_eq!(values(&result),expected,"{variant} n={n} ratio={ratio} mask={mask}");
    if let Some(ref old)=old{assert_eq!(values(old),original);}
    if let Some(ref old_tail)=old_tail{assert_eq!(values(old_tail),tail_original);}
    if let Some(ref weak)=weak {
        assert_eq!(weak.upgrade().is_some(),mask&1!=0 || (pivot==0 && mask&2!=0));
        if let Some(old)=weak.upgrade(){assert_eq!(values(&old),original);}
    }
    if let Some(ref weak)=weak_tail {
        assert_eq!(weak.upgrade().is_some(),mask&3!=0);
        if let Some(old)=weak.upgrade(){assert_eq!(values(&old),tail_original);}
    }
    #[cfg(list_counted)] {
        assert_eq!(paths[0]+paths[1],n as u64);
        assert_eq!(paths[3],expected.len() as u64);
        if variant=="baseline" { assert_eq!(paths[0],0,"adapter owner still retains the complete source list"); }
        if variant=="reuse" {
            let eligible_end=if mask&1!=0 {0} else if mask&2!=0 {pivot} else {n};
            let reused=original.iter().enumerate().filter(|(i,v)|
                *i<eligible_end && **v%2==0 && !(*i==0 && mask&4!=0) && !(*i==pivot && mask&8!=0)).count();
            assert_eq!(paths[4],reused as u64,"reuse only truly exclusive retained cells");
            assert_eq!(call_heap[2],(expected.len()-reused+1) as u64,"one Nil plus newly constructed kept cells");
        } else {
            assert_eq!(paths[4],0);
            assert_eq!(call_heap[2],(expected.len()+1) as u64);
        }
        if variant!="baseline" && mask==0 {assert_eq!(paths[0],n as u64,"precise adapter lifetime exposes unique nodes");}
    }
    drop(result);drop(old);drop(old_tail);
    if let Some(ref weak)=weak{assert!(weak.upgrade().is_none());}
    if let Some(ref weak)=weak_tail{assert!(weak.upgrade().is_none());}
    drop(weak);drop(weak_tail);
    (paths,call_heap,expected.len())
}

fn run_checks(variant: String) {
    // Initialize printing and generated thunk paths before accounting cases.
    println!("ListOps {variant} validation; no timing");
    drop(Test_ListOps_filterEvens(input(2,50)));
    probe_filter_counts();
    let mut cases=0;
    for n in [0,1,2,31,900,9000] {
        for ratio in [0,50,100] {
            for mask in 0..16 {
                let before=heap();
                let (paths,call_heap,kept)=one_case(&variant,n,ratio,mask);
                let all_heap=delta(heap(),before);
                #[cfg(list_counted)] assert_eq!(all_heap[2],all_heap[3],"Every List cell released, including retained/Weak cases");
                if n>=900 {
                    println!(concat!("{{\"kind\":\"filter\",\"variant\":\"{}\",\"n\":{},\"ratio\":{},\"mask\":{},\"kept\":{},",
                        "\"strong_unique\":{},\"strong_shared\":{},\"weak_visits\":{},\"reused\":{},\"consumed_unique\":{},\"consumed_shared\":{},",
                        "\"call_allocations\":{},\"call_frees\":{},\"call_list_allocations\":{},\"call_list_frees\":{},",
                        "\"all_allocations\":{},\"all_frees\":{},\"all_list_allocations\":{},\"all_list_frees\":{}}}"),
                        variant,n,ratio,mask,kept,paths[0],paths[1],paths[2],paths[4],paths[5],paths[6],
                        call_heap[0],call_heap[1],call_heap[2],call_heap[3],all_heap[0],all_heap[1],all_heap[2],all_heap[3]);
                }
                cases+=1;
            }
        }
    }
    for n in [0,1,2,31,900,9000] {
        let before=heap();probe_filter_counts();
        assert_eq!(Test_ListOps_sumEvens(n),(1..=n).filter(|v|v%2==0).sum::<i64>());
        let counts=probe_filter_counts();let h=delta(heap(),before);
        #[cfg(list_counted)] {
            assert_eq!(h[2],h[3]);
            if variant=="baseline"{assert_eq!(counts[0],0);}else{assert_eq!(counts[0],n as u64);}
            if variant=="reuse"{assert_eq!(counts[4],(n/2) as u64);}else{assert_eq!(counts[4],0);}
        }
        println!("{{\"kind\":\"sum\",\"variant\":\"{}\",\"n\":{},\"strong_unique\":{},\"strong_shared\":{},\"reused\":{},\"allocations\":{},\"frees\":{},\"list_allocations\":{},\"list_frees\":{}}}",
            variant,n,counts[0],counts[1],counts[4],h[0],h[1],h[2],h[3]);
    }
    println!("{{\"checks_passed\":true,\"sharing_cases\":{},\"sum_cases\":6,\"no_timings\":true}}",cases);
}
fn main() {
    let variant=std::env::args().nth(1).unwrap();
    // Ordinary recursive destruction of a retained 9000-element baseline list
    // gets an explicit stack budget; production generation stays untouched.
    std::thread::Builder::new().stack_size(64*1024*1024).spawn(move||run_checks(variant)).unwrap().join().unwrap();
}

use std::rc::Rc;
use Purs_Test_ListOps::{List, Test_ListOps_filterEvens, Test_ListOps_sumEvens, probe_filter_counts};

fn values(xs: &Rc<List>) -> Vec<i64> {
    let mut result = Vec::new();
    let mut cursor = xs;
    while let List::Cons(value, tail) = cursor.as_ref() {
        result.push(value.unwrap_int());
        cursor = tail;
    }
    result
}

fn input(n: usize) -> Rc<List> {
    let mut result = Rc::new(List::Nil);
    for i in (1..=n).rev() {
        result = Rc::new(List::Cons(purust_core::mk_int(i as i64), result));
    }
    result
}

fn main() {
    let mut cases = 0;
    for n in [0, 1, 2, 8, 31, 900] {
        for mask in 0..16 {
            let root = input(n);
            let tail = match root.as_ref() { List::Cons(_, tail) => tail, List::Nil => &root };
            let old_tail = (mask & 2 != 0).then(|| tail.clone());
            let weak_tail = (mask & 8 != 0).then(|| Rc::downgrade(tail));
            let old = (mask & 1 != 0).then(|| root.clone());
            let weak = (mask & 4 != 0).then(|| Rc::downgrade(&root));
            let original = values(&root);
            let tail_original = values(tail);
            let result = Test_ListOps_filterEvens(root);
            let expected: Vec<_> = original.iter().copied().rev().filter(|v| v % 2 == 0).collect();
            assert_eq!(values(&result), expected);
            if let Some(ref old) = old { assert_eq!(values(old), original); }
            if let Some(ref old) = old_tail { assert_eq!(values(old), tail_original); }
            drop(result); drop(old); drop(old_tail);
            if let Some(ref weak) = weak { assert!(weak.upgrade().is_none()); }
            if let Some(ref weak) = weak_tail { assert!(weak.upgrade().is_none()); }
            cases += 1;
        }
    }
    let filter = probe_filter_counts();
    for n in [0, 1, 2, 10, 31, 900] {
        assert_eq!(Test_ListOps_sumEvens(n), (1..=n).filter(|v| v % 2 == 0).sum::<i64>());
    }
    let sums = probe_filter_counts();
    println!("{{\"sharing_cases\":{cases},\"sum_cases\":6,\"filter_unique\":{},\"filter_shared\":{},\"sum_unique\":{},\"sum_shared\":{}}}", filter[0], filter[1], sums[0], sums[1]);
}

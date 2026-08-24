
enum List {
    Nil,
    Cons(i64, Box<List>),
}

fn range_list(start: i64, end: i64) -> Box<List> {
    let mut acc = Box::new(List::Nil);
    let mut curr = end;
    while curr >= start {
        acc = Box::new(List::Cons(curr, acc));
        curr -= 1;
    }
    acc
}

fn filter_evens(mut lst: &List) -> Box<List> {
    let mut evens = Vec::new();
    while let List::Cons(x, xs) = lst {
        if x % 2 == 0 {
            evens.push(*x);
        }
        lst = xs;
    }
    let mut acc = Box::new(List::Nil);
    for x in evens.iter().rev() {
        acc = Box::new(List::Cons(*x, acc));
    }
    acc
}

fn foldl<F>(f: F, mut acc: i64, mut lst: &List) -> i64
where
    F: Fn(i64, i64) -> i64,
{
    while let List::Cons(x, xs) = lst {
        acc = f(acc, *x);
        lst = xs;
    }
    acc
}

pub fn Test_ListOpsFFI_runListOpsFFI(mut limit: i64) -> i64 {
    let lst = range_list(1, limit);
    let filtered = filter_evens(&lst);
    foldl(|a, b| a + b, 0, &filtered)
}
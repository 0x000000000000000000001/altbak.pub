enum List<A> {
    Nil,
    Cons(A, Box<List<A>>),
}

fn range_list(start: i64, end: i64) -> Box<List<i64>> {
    let mut acc = Box::new(List::Nil);
    let mut curr = end;
    while curr >= start {
        acc = Box::new(List::Cons(curr, acc));
        curr -= 1;
    }
    acc
}

fn filter_evens(mut lst: &List<i64>) -> Box<List<i64>> {
    let mut acc = Box::new(List::Nil);
    while let List::Cons(x, xs) = lst {
        if x % 2 == 0 {
            acc = Box::new(List::Cons(*x, acc));
        }
        lst = xs;
    }
    acc
}

fn foldl<A, B, F>(f: F, mut acc: B, mut lst: &List<A>) -> B
where
    F: Fn(B, &A) -> B,
{
    while let List::Cons(x, xs) = lst {
        acc = f(acc, x);
        lst = xs;
    }
    acc
}

pub fn Test_ListOpsFFI_runListOpsFFI(limit: i64) -> i64 {
    let lst = range_list(1, limit);
    let filtered = filter_evens(&lst);
    foldl(|a, b| a + *b, 0, &filtered)
}

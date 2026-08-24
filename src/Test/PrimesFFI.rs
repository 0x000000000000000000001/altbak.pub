// Primes sieve
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

fn filter_primes(lst: Box<List>) -> Box<List> {
    match *lst {
        List::Nil => Box::new(List::Nil),
        List::Cons(p, xs) => {
            let mut filtered_xs = filter_divisible(p, xs);
            Box::new(List::Cons(p, filter_primes(filtered_xs)))
        }
    }
}

fn filter_divisible(p: i64, mut lst: Box<List>) -> Box<List> {
    let mut valid = Vec::new();
    while let List::Cons(x, xs) = *lst {
        if x % p != 0 {
            valid.push(x);
        }
        lst = xs;
    }
    let mut acc = Box::new(List::Nil);
    for x in valid.iter().rev() {
        acc = Box::new(List::Cons(*x, acc));
    }
    acc
}

fn list_length(mut lst: &List) -> i64 {
    let mut count = 0;
    while let List::Cons(_, xs) = lst {
        count += 1;
        lst = xs;
    }
    count
}

pub fn Test_PrimesFFI_runPrimesFFI(mut limit: i64) -> i64 {
    let lst = range_list(2, limit);
    let primes = filter_primes(lst);
    list_length(&primes)
}
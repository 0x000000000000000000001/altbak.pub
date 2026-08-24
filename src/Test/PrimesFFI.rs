enum List {
    Nil,
    Cons(i64, Box<List>),
}

fn range_list(start: i64, end: i64) -> List {
    fn go_func(curr: i64, acc: List, start: i64) -> List {
        if curr < start {
            acc
        } else {
            go_func(curr - 1, List::Cons(curr, Box::new(acc)), start)
        }
    }
    go_func(end, List::Nil, start)
}

fn reverse(lst: List) -> List {
    fn go_func(list: List, acc: List) -> List {
        match list {
            List::Nil => acc,
            List::Cons(x, xs) => go_func(*xs, List::Cons(x, Box::new(acc)))
        }
    }
    go_func(lst, List::Nil)
}

fn filter<P>(p: &P, lst: List) -> List 
where P: Fn(i64) -> bool
{
    fn go_func<P>(p: &P, list: List, acc: List) -> List 
    where P: Fn(i64) -> bool 
    {
        match list {
            List::Nil => reverse(acc),
            List::Cons(x, xs) => {
                if p(x) {
                    go_func(p, *xs, List::Cons(x, Box::new(acc)))
                } else {
                    go_func(p, *xs, acc)
                }
            }
        }
    }
    go_func(p, lst, List::Nil)
}

fn sieve(lst: List) -> List {
    match lst {
        List::Nil => List::Nil,
        List::Cons(p, xs) => {
            let filtered = filter(&|x| x % p != 0, *xs);
            List::Cons(p, Box::new(sieve(filtered)))
        }
    }
}

fn sum_list(lst: List) -> i64 {
    fn go_func(list: List, acc: i64) -> i64 {
        match list {
            List::Nil => acc,
            List::Cons(x, xs) => go_func(*xs, acc + x)
        }
    }
    go_func(lst, 0)
}

pub fn Test_PrimesFFI_runPrimesFFI(limit: i64) -> i64 {
    sum_list(sieve(range_list(2, limit)))
}

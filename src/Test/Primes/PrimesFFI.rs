enum List<A> {
    Nil,
    Cons(A, Box<List<A>>),
}

fn range_list(start: i64, end: i64) -> List<i64> {
    fn go_func(curr: i64, acc: List<i64>, start: i64) -> List<i64> {
        if curr < start {
            acc
        } else {
            go_func(curr - 1, List::Cons(curr, Box::new(acc)), start)
        }
    }
    go_func(end, List::Nil, start)
}

fn reverse<A>(lst: List<A>) -> List<A> {
    fn go_func<A>(list: List<A>, acc: List<A>) -> List<A> {
        match list {
            List::Nil => acc,
            List::Cons(x, xs) => go_func(*xs, List::Cons(x, Box::new(acc)))
        }
    }
    go_func(lst, List::Nil)
}

fn filter<A, P>(p: &P, lst: List<A>) -> List<A>
where P: Fn(&A) -> bool
{
    fn go_func<A, P>(p: &P, list: List<A>, acc: List<A>) -> List<A>
    where P: Fn(&A) -> bool
    {
        match list {
            List::Nil => reverse(acc),
            List::Cons(x, xs) => {
                if p(&x) {
                    go_func(p, *xs, List::Cons(x, Box::new(acc)))
                } else {
                    go_func(p, *xs, acc)
                }
            }
        }
    }
    go_func(p, lst, List::Nil)
}

fn sieve(lst: List<i64>) -> List<i64> {
    match lst {
        List::Nil => List::Nil,
        List::Cons(p, xs) => {
            let filtered = filter(&|x: &i64| *x % p != 0, *xs);
            List::Cons(p, Box::new(sieve(filtered)))
        }
    }
}

fn sum_list(lst: List<i64>) -> i64 {
    fn go_func(list: List<i64>, acc: i64) -> i64 {
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

use std::rc::Rc;

#[derive(Clone)]
pub enum List {
    Nil,
    Cons(i64, Rc<List>),
}

fn sieve(v: Rc<List>) -> Rc<List> {
    match &*v {
        List::Nil => Rc::new(List::Nil),
        List::Cons(p, xs) => {
            let filtered = filter(xs.clone(), *p);
            Rc::new(List::Cons(*p, sieve(filtered)))
        }
    }
}

fn filter(mut v: Rc<List>, p: i64) -> Rc<List> {
    let mut acc = Rc::new(List::Nil);
    loop {
        match &*v {
            List::Nil => {
                return reverse(acc);
            }
            List::Cons(x, xs) => {
                if *x % p != 0 {
                    acc = Rc::new(List::Cons(*x, acc));
                }
                v = xs.clone();
            }
        }
    }
}

fn reverse(mut v: Rc<List>) -> Rc<List> {
    let mut acc = Rc::new(List::Nil);
    loop {
        match &*v {
            List::Nil => {
                return acc;
            }
            List::Cons(x, xs) => {
                acc = Rc::new(List::Cons(*x, acc));
                v = xs.clone();
            }
        }
    }
}

fn range(start: i64, end: i64) -> Rc<List> {
    let mut curr = end;
    let mut acc = Rc::new(List::Nil);
    loop {
        if curr < start {
            return acc;
        }
        acc = Rc::new(List::Cons(curr, acc));
        curr -= 1;
    }
}

fn main() {
    let l = range(2, 500);
    let s = sieve(l);
    let mut curr = s;
    let mut sum = 0;
    loop {
        match &*curr {
            List::Nil => break,
            List::Cons(x, xs) => {
                sum += *x;
                curr = xs.clone();
            }
        }
    }
    println!("{}", sum);
}

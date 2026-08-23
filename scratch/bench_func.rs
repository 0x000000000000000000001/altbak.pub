use std::rc::Rc;
use std::time::Instant;
use std::ops::Deref;
use std::hint::black_box;

// 1. Approche actuelle (Pénalisée par le virtual dispatch)
type DynFunc = Rc<dyn Fn(i32, i32) -> i32>;

// 2. Approche idéale mais rigide (Pointeurs natifs)
type StatFunc = fn(i32, i32) -> i32;

// 3. Le Smart Wrapper (Inspiré de Fable)
enum Func<A, B, R> {
    Static(fn(A, B) -> R),
    Shared(Rc<dyn Fn(A, B) -> R>),
}

impl<A: 'static, B: 'static, R: 'static> Deref for Func<A, B, R> {
    type Target = dyn Fn(A, B) -> R;

    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func::Static(f) => f,
            Func::Shared(rc) => rc.as_ref(),
        }
    }
}

fn add(a: i32, b: i32) -> i32 {
    a + b
}

fn main() {
    let iterations = 100_000_000;
    println!("Démarrage du Benchmark ({} iterations)...\n", iterations);

    // --- Benchmark 1: Rc<dyn Fn> ---
    let start = Instant::now();
    let mut sum1 = 0;
    for i in 0..iterations {
        let dyn_func: DynFunc = Rc::new(add);
        sum1 = black_box(dyn_func(black_box(sum1), 1));
    }
    println!("1. Rc<dyn Fn> (Virtuel + Heap)  : {:?} (sum={})", start.elapsed(), sum1);

    // --- Benchmark 2: Static fn ---
    let start = Instant::now();
    let mut sum2 = 0;
    for i in 0..iterations {
        let stat_func: StatFunc = add;
        sum2 = black_box(stat_func(black_box(sum2), 1));
    }
    println!("2. Static fn (Natif)            : {:?} (sum={})", start.elapsed(), sum2);

    // --- Benchmark 3: Smart Wrapper (Mode Static) ---
    let start = Instant::now();
    let mut sum3 = 0;
    for i in 0..iterations {
        let smart_func: Func<i32, i32, i32> = Func::Static(add);
        sum3 = black_box(smart_func(black_box(sum3), 1));
    }
    println!("3. Smart Wrapper (Stack Seulement): {:?} (sum={})", start.elapsed(), sum3);

}

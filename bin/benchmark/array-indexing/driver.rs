// Timing/input harness only: both indexed-read loops come from ArrayIndexing.purs.
use std::alloc::{GlobalAlloc, Layout};
use std::sync::atomic::{AtomicUsize, Ordering};
use std::time::Instant;

use purust_core::{mk_array, UnknownType, Value};
use Purs_Test_ArrayIndexing::{Test_ArrayIndexing_boxedReads, Test_ArrayIndexing_nativeReads};

// Count requested bytes so the collector sees the same allocation signal as
// the Go harness; the counter is read only between timed regions.
struct Counting;

static ALLOCATED: AtomicUsize = AtomicUsize::new(0);

unsafe impl GlobalAlloc for Counting {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        ALLOCATED.fetch_add(layout.size(), Ordering::Relaxed);
        unsafe { mimalloc::MiMalloc.alloc(layout) }
    }
    unsafe fn dealloc(&self, ptr: *mut u8, layout: Layout) {
        unsafe { mimalloc::MiMalloc.dealloc(ptr, layout) }
    }
}

#[global_allocator]
static GLOBAL: Counting = Counting;

fn main() {
    let mut accesses: i64 = 1 << 23;
    let mut batches: usize = 10;
    let mut seed: i64 = 5;
    let args: Vec<String> = std::env::args().collect();
    let mut i = 1;
    while i + 1 < args.len() {
        match args[i].as_str() {
            "-accesses" => accesses = args[i + 1].parse().expect("accesses"),
            "-batches" => batches = args[i + 1].parse().expect("batches"),
            "-seed" => seed = args[i + 1].parse().expect("seed"),
            _ => {}
        }
        i += 2;
    }
    assert!(
        (16384..=1 << 23).contains(&accesses) && batches >= 3 && (0..=1000).contains(&seed),
        "invalid accesses, batches or seed"
    );

    for size in [16i64, 1024, 16384] {
        let values: Vec<i64> = (0..size).map(|i| (i * 17 + seed * 31) % 251 + 1).collect();
        let elements: Vec<UnknownType> = values.iter().map(|v| Value::Int(*v)).collect();
        let native = mk_array(elements);
        let boxed = native.clone();
        let start = seed % size;
        let oracle = |count: i64| -> i64 {
            let total: i64 = values.iter().sum();
            let mut result = total * (count / size);
            for k in 0..(count % size) {
                result += values[((start + k) % size) as usize];
            }
            result
        };
        for representation in ["native", "boxed"] {
            let call = |count: i64| -> i64 {
                if representation == "native" {
                    Test_ArrayIndexing_nativeReads(native.clone(), size, count, start)
                } else {
                    Test_ArrayIndexing_boxedReads(boxed.clone(), size, count, start)
                }
            };
            // A bounded preflight catches a whole-array copy before a long
            // batch could allocate terabytes on the largest input.
            let probe_expected = oracle(1024);
            let probe_before = ALLOCATED.load(Ordering::Relaxed);
            let probe = call(1024);
            let probe_bytes = (ALLOCATED.load(Ordering::Relaxed) - probe_before) as f64 / 1024.0;
            assert_eq!(probe, probe_expected, "preflight checksum mismatch");
            if probe_bytes > 256.0 {
                panic!(
                    "indexing regression: {} size={} bytes/access={:.3}",
                    representation, size, probe_bytes
                );
            }
            let expected = oracle(accesses);
            for _ in 0..3 {
                assert_eq!(call(accesses), expected, "warm-up checksum mismatch");
            }
            let mut nanoseconds: Vec<f64> = Vec::with_capacity(batches);
            let mut allocated: Vec<f64> = Vec::with_capacity(batches);
            for _ in 0..batches {
                let before = ALLOCATED.load(Ordering::Relaxed);
                let begin = Instant::now();
                let sink = call(accesses);
                let elapsed = begin.elapsed().as_nanos() as f64;
                let bytes = (ALLOCATED.load(Ordering::Relaxed) - before) as f64;
                assert_eq!(sink, expected, "invalid measured result");
                nanoseconds.push(elapsed / accesses as f64);
                allocated.push(bytes / accesses as f64);
            }
            let ns = nanoseconds
                .iter()
                .map(|v| format!("{v:.6}"))
                .collect::<Vec<_>>()
                .join(",");
            let bytes = allocated
                .iter()
                .map(|v| format!("{v:.6}"))
                .collect::<Vec<_>>()
                .join(",");
            println!(
                "{{\"runtime\":\"rust\",\"representation\":\"{representation}\",\"size\":{size},\"accesses\":{accesses},\"seed\":{seed},\"checksum\":{expected},\"warmups\":3,\"ns_per_access\":[{ns}],\"bytes_per_access\":[{bytes}],\"probe_bytes_per_access\":{probe_bytes:.6}}}"
            );
        }
    }
}

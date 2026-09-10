"""Count clones and heap operations on copies of actual generated Records Rust.

Run only after the uninstrumented measurements are finished. Each variant
recompiles its own purust_core against a private instrumented PerceusPtr.
"""
from pathlib import Path
import hashlib
import json
import subprocess


HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE / "build"
VARIANTS = {
    "before": BUILD / "before-output",
    "after": ROOT / "run/bak/rust/output/purust_output",
}
RUNTIME = ROOT.parent / "purust/purust/tests/runtime/perceus_ptr/src/lib.rs"

COUNTER = """
use std::sync::atomic::{AtomicUsize, Ordering};
static RECORD_CLONE_CALLS: AtomicUsize = AtomicUsize::new(0);
pub fn record_clone_calls() -> usize { RECORD_CLONE_CALLS.load(Ordering::Relaxed) }
pub fn reset_record_clone_calls() { RECORD_CLONE_CALLS.store(0, Ordering::Relaxed); }
"""

MAIN = r"""
use std::alloc::{GlobalAlloc, Layout, System};
use std::sync::atomic::{AtomicUsize, Ordering};
static RECORD_ALLOCS: AtomicUsize = AtomicUsize::new(0);
static RECORD_FREES: AtomicUsize = AtomicUsize::new(0);
struct RecordAllocator;
unsafe impl GlobalAlloc for RecordAllocator {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        RECORD_ALLOCS.fetch_add(1, Ordering::Relaxed);
        System.alloc(layout)
    }
    unsafe fn dealloc(&self, ptr: *mut u8, layout: Layout) {
        RECORD_FREES.fetch_add(1, Ordering::Relaxed);
        System.dealloc(ptr, layout)
    }
}
#[global_allocator]
static RECORD_ALLOCATOR: RecordAllocator = RecordAllocator;

fn values(r: &UnknownType) -> [i64; 4] {
    [r.get_a().unwrap_int(), r.get_b().get_c().unwrap_int(),
     r.get_b().get_d().get_e().unwrap_int(), r.get_b().get_d().get_f().unwrap_int()]
}
fn main() {
    for n in [0_i64, 1, 2, 10, 10000] {
        RECORD_ALLOCS.store(0, Ordering::Relaxed);
        RECORD_FREES.store(0, Ordering::Relaxed);
        perceus_ptr::reset_record_clone_calls();
        let initial = Test_Records_initial();
        let initial_allocations = RECORD_ALLOCS.load(Ordering::Relaxed);
        let initial_clones = perceus_ptr::record_clone_calls();
        let result = Test_Records_updateRec(std::hint::black_box(n), initial);
        // Capture update costs before owning getters verify the four fields.
        let after_update_clones = perceus_ptr::record_clone_calls();
        let after_update_allocations = RECORD_ALLOCS.load(Ordering::Relaxed);
        let found = values(&result);
        let after_values_clones = perceus_ptr::record_clone_calls();
        let after_values_allocations = RECORD_ALLOCS.load(Ordering::Relaxed);
        assert_eq!(found, [n, 2 * n, 3 * n, (1..=n).map(|i| i % 5).sum::<i64>()]);
        drop(result);
        let after_drop_clones = perceus_ptr::record_clone_calls();
        let total_allocations = RECORD_ALLOCS.load(Ordering::Relaxed);
        let total_frees = RECORD_FREES.load(Ordering::Relaxed);
        assert_eq!(total_allocations, total_frees, "kernel heap allocations must be released");
        println!(concat!("{{\"iterations\":{},\"initial_clones\":{},\"update_clones\":{},",
            "\"verification_clones\":{},\"drop_clones\":{},\"initial_allocations\":{},",
            "\"update_allocations\":{},\"verification_allocations\":{},",
            "\"total_allocations\":{},\"total_frees\":{},\"fields\":{:?}}}"),
            n, initial_clones, after_update_clones - initial_clones,
            after_values_clones - after_update_clones, after_drop_clones - after_values_clones,
            initial_allocations, after_update_allocations - initial_allocations,
            after_values_allocations - after_update_allocations, total_allocations, total_frees, found);
    }
}
"""


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def compile_rust(arguments):
    command = ["rustc", "--edition=2021", "-C", "opt-level=1", *map(str, arguments)]
    result = subprocess.run(command, text=True, capture_output=True)
    if result.returncode:
        raise RuntimeError(result.stdout + result.stderr)


def count_variant(name, source_directory):
    directory = BUILD / f"count-{name}"
    deps = directory / "deps"
    deps.mkdir(parents=True, exist_ok=True)
    # Only unchanged third-party dependencies come from the completed after
    # build. The before Cargo build may still be populating its own target.
    source_deps = VARIANTS["after"] / "target/release/deps"
    original = RUNTIME.read_text()
    marker = "    fn clone(&self) -> Self {\n"
    assert original.count(marker) == 1
    instrumented = original.replace(marker, marker + "        RECORD_CLONE_CALLS.fetch_add(1, Ordering::Relaxed);\n")
    runtime_copy = directory / "perceus_ptr.rs"
    runtime_copy.write_text(instrumented + COUNTER)
    runtime_lib = deps / "libperceus_ptr.rlib"
    metadata = f"record_borrow_count_{name}"
    compile_rust(["--crate-type=rlib", "--crate-name=perceus_ptr", "-C", f"metadata={metadata}",
                  runtime_copy, "-o", runtime_lib])

    # Crucially, the before kernel uses its before runtime, without injecting
    # the after variant's borrowed getters or another variant's compiled core.
    core_source = source_directory / "purust_core/src/lib.rs"
    core_copy = directory / "purust_core.rs"
    core_copy.write_bytes(core_source.read_bytes())
    regex_libs = list(source_deps.glob("libfancy_regex-*.rlib"))
    assert len(regex_libs) == 1, (name, regex_libs)
    core_lib = deps / "libpurust_core.rlib"
    compile_rust(["--crate-type=rlib", "--crate-name=purust_core", "-C", f"metadata={metadata}",
                  core_copy, "-o", core_lib, "-L", f"dependency={source_deps}",
                  "--extern", f"perceus_ptr={runtime_lib}", "--extern", f"fancy_regex={regex_libs[0]}"])

    record_source = source_directory / "Purs_Test_Records/src/lib.rs"
    source = record_source.read_text()
    kernel = source[source.index("pub fn Test_Records_updateRec("):source.index("pub fn Test_Records_describe(")]
    assert "pub fn Test_Records_initial()" in kernel
    kernel_copy = directory / "kernel.rs"
    kernel_copy.write_text("#![allow(warnings)]\npub use purust_core::*;\n" + kernel + MAIN)
    binary = directory / "kernel"
    compile_rust([kernel_copy, "-o", binary, "-L", f"dependency={deps}", "-L", f"dependency={source_deps}",
                  "--extern", f"perceus_ptr={runtime_lib}", "--extern", f"purust_core={core_lib}"])
    output = subprocess.check_output([str(binary)], text=True)
    (directory / "counts.log").write_text(output)
    rows = [json.loads(line) for line in output.splitlines()]
    assert [row["iterations"] for row in rows] == [0, 1, 2, 10, 10000]
    print(name, json.dumps(rows[-1]), flush=True)
    return {
        "source_sha256": {str(path): sha(path) for path in [RUNTIME, core_source, record_source]},
        "instrumented_core_sha256": sha(core_copy),
        "third_party_dependency_directory": str(source_deps),
        "results": rows,
    }


def main():
    results = {name: count_variant(name, directory) for name, directory in VARIANTS.items()}
    report = {
        "method": "Each exact generated Records kernel and its own copied purust_core are recompiled with O1 "
                  "and separate dependencies. A private PerceusPtr increments an atomic on clone entry. "
                  "The global allocator counts heap allocation/free calls using System. Counters reset before "
                  "initial construction; update snapshots precede the four-field verification and destruction.",
        "limitations": "Instrumentation may prevent optimizer elimination, so clone counts describe these "
                       "instrumented binaries rather than refcount instructions surviving in timed binaries. "
                       "No timing is performed. The counting allocator is System, not the timed runner's mimalloc. "
                       "Immediate Value::Int copies and other pointer types are not counted as PerceusPtr clones. "
                       "Heap counts exclude printing and include initial construction, verification and final destruction.",
        "rustc": subprocess.check_output(["rustc", "--version"], text=True).strip(),
        "variants": results,
    }
    (HERE / "counts.json").write_text(json.dumps(report, indent=2) + "\n")


if __name__ == "__main__":
    main()

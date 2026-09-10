"""Count PerceusPtr::clone calls; instrumentation is isolated from timed binaries."""
from pathlib import Path
import hashlib
import json
import subprocess


HERE = Path(__file__).resolve().parent
BUILD = HERE / "build"
WORK = BUILD / "workspace"
DEPS = WORK / "target/release/deps"
COUNT_DEPS = BUILD / "count-clone-deps"
RUNTIME = Path("/Users/0x1/Documents/htdocs/purust/purust/tests/runtime/perceus_ptr/src/lib.rs")
NAMES = ["before", "root_borrow", "path_borrow"]

COUNTER = """
use std::sync::atomic::{AtomicUsize, Ordering};
static RECORD_CLONE_CALLS: AtomicUsize = AtomicUsize::new(0);
pub fn record_clone_calls() -> usize { RECORD_CLONE_CALLS.load(Ordering::Relaxed) }
pub fn reset_record_clone_calls() { RECORD_CLONE_CALLS.store(0, Ordering::Relaxed); }
"""

MAIN = r"""
fn values(r: &UnknownType) -> [i64; 4] {
    [r.get_a().unwrap_int(), r.get_b().get_c().unwrap_int(),
     r.get_b().get_d().get_e().unwrap_int(), r.get_b().get_d().get_f().unwrap_int()]
}
fn main() {
    for n in [0_i64, 1, 2, 10, 10000] {
        let initial = Test_Records_initial();
        perceus_ptr::reset_record_clone_calls();
        let result = Test_Records_updateRec(std::hint::black_box(n), initial);
        // Snapshot precedes verification, which itself uses owning getters.
        let update_clones = perceus_ptr::record_clone_calls();
        let found = values(&result);
        let after_values = perceus_ptr::record_clone_calls();
        assert_eq!(found, [n, 2 * n, 3 * n, (1..=n).map(|i| i % 5).sum::<i64>()]);
        drop(result);
        let after_drop = perceus_ptr::record_clone_calls();
        println!("{{\"iterations\":{n},\"update_clones\":{update_clones},\"verification_clones\":{},\"drop_clones\":{},\"fields\":{:?}}}",
                 after_values - update_clones, after_drop - after_values, found);
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


def main():
    COUNT_DEPS.mkdir(exist_ok=True)
    original = RUNTIME.read_text()
    marker = "    fn clone(&self) -> Self {\n"
    assert original.count(marker) == 1
    instrumented = original.replace(marker, marker + "        RECORD_CLONE_CALLS.fetch_add(1, Ordering::Relaxed);\n")
    runtime_copy = BUILD / "count-clone-runtime.rs"
    runtime_copy.write_text(instrumented + COUNTER)
    runtime_lib = COUNT_DEPS / "libperceus_ptr.rlib"
    compile_rust(["--crate-type=rlib", "--crate-name=perceus_ptr", "-C", "metadata=record_clone_counter",
                  runtime_copy, "-o", runtime_lib])

    core_source = WORK / "purust_core/src/lib.rs"
    core_copy = BUILD / "count-clone-core.rs"
    core_copy.write_bytes(core_source.read_bytes())
    regex_libs = list(DEPS.glob("libfancy_regex-*.rlib"))
    assert len(regex_libs) == 1, regex_libs
    core_lib = COUNT_DEPS / "libpurust_core.rlib"
    compile_rust(["--crate-type=rlib", "--crate-name=purust_core", "-C", "metadata=record_clone_counter",
                  core_copy, "-o", core_lib, "-L", f"dependency={DEPS}",
                  "--extern", f"perceus_ptr={runtime_lib}", "--extern", f"fancy_regex={regex_libs[0]}"])

    results = {}
    for name in NAMES:
        source = (BUILD / f"Records-{name}.rs").read_text()
        kernel = source[source.index("pub fn Test_Records_updateRec("):source.index("pub fn Test_Records_describe(")]
        path = BUILD / f"count-clones-{name}.rs"
        path.write_text("#![allow(warnings)]\npub use purust_core::*;\n" + kernel + MAIN)
        binary = path.with_suffix("")
        compile_rust([path, "-o", binary, "-L", f"dependency={COUNT_DEPS}", "-L", f"dependency={DEPS}",
                      "--extern", f"perceus_ptr={runtime_lib}", "--extern", f"purust_core={core_lib}"])
        output = subprocess.check_output([str(binary)], text=True)
        results[name] = [json.loads(line) for line in output.splitlines()]
        print(name, json.dumps(results[name][-1]), flush=True)

    report = {
        "method": "Isolated copy of PerceusPtr increments AtomicUsize once on entry to clone. "
                  "The copied borrowed-getter purust_core is linked against that runtime. "
                  "Fresh initial record before reset; snapshot immediately after updateRec and before four-field verification.",
        "limitations": "Counts dynamic clone invocations in instrumented O1 binaries, not time or the number of "
                       "reference-counter instructions surviving in uninstrumented machine code. Atomic side effects "
                       "can prevent optimizer elimination. No timing is taken. Only PerceusPtr clones are counted, "
                       "not Value::Int copies, other pointer types, allocation or destruction operations.",
        "rustc": subprocess.check_output(["rustc", "--version"], text=True).strip(),
        "source_sha256": {str(path): sha(path) for path in [RUNTIME, core_source, *[BUILD / f"Records-{name}.rs" for name in NAMES]]},
        "results": results,
    }
    (HERE / "counts.json").write_text(json.dumps(report, indent=2) + "\n")


if __name__ == "__main__":
    main()

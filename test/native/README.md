# Native benchmark contracts

`bin/go/test-native` and `bin/rust/test-native` compile the exact 28 native kernels of each language in temporary directories. They check 218 results per language against `oracle.py` (small inputs plus all 14 nominal inputs). They execute no benchmark harness and record no performance measurements. Rust uses O1 with overflow checks for these contract checks; Go explicitly disables PGO. Clock tests check monotonic fractional microseconds and the opaque input contract.

`python3 test/native/runner/check.py` tests the runner orchestration with simulated tool effects. It does not build or run the suite.

The FFI State argument is the chain depth: 20 fresh chains of that depth. The Lazy argument is the number of repetitions: a fresh depth-1000 chain per repetition. Array.range includes both endpoints and descends when its second endpoint is below 1. Other recursive inputs are tested only in their terminating benchmark domain; the tests do not promise arbitrary Int overflow behavior outside this domain.

Rust also checks that list filtering reverses the selected cells, Church numerals apply their supplied numeral only once the final argument arrives, and red-black insertion preserves old snapshots, balancing invariants and shallow Rc sharing (including an existing Weak observer). Native FP is handwritten code with explicit functional structure, not an instruction-for-instruction copy of backend output. Cheatcode remains free to use deliberate algorithmic simplifications while preserving the benchmark result and argument.

# Isolated Go/Rust runners

`bin/go/run --build-only`, `bin/go/run --ffi --build-only`, and `bin/go/run --fficc --build-only` build three independent workspaces under `run/bak/go/modes/`. Rust has the same options under `run/bak/rust/modes/`. Add `--build-dir PATH` to select another isolated workspace. Source copies, dependencies, generated output, logs and executable stay there; the repository's AppX, output symlink, Spago files and neighbouring compilers are untouched.

Replace `--build-only` with `--run-only` to execute the corresponding saved binary. The manifest verifies mode, input source hashes, backend bundle, profile and executable SHA before execution; all output then passes the shared `bin/benchmark/validate.py` validator. A failed build invalidates its old manifest. `--clean` affects only that marked workspace. `--test MODULE [--expected VALUE]` generates an AppX inside its workspace. `--x` copies the extended `srx` suite there; the shared validator checks its output structure, not a core value oracle.

Rust release builds explicitly use `CARGO_PROFILE_RELEASE_OPT_LEVEL=3 CARGO_PROFILE_RELEASE_DEBUG=false` for all three columns. The generated Cargo manifest still declares O1 and debug=true; those were the former runner defaults. Setting these environment variables explicitly can reproduce that profile. The effective profile and toolchain are saved in `manifest.json`; O1 historical measurements are not interchangeable with new O3 results. No neighbouring compiler source is changed.

Go uses GOGC=800 by default for ordinary execution, PGO training and PGO execution alike; an explicit GOGC overrides all three. Ordinary builds pass `-pgo=off`. `--pgo` validates a training run before rebuilding with its CPU profile. `--build-only --pgo` requires `--pgo-profile FILE`, so build-only cannot silently execute a benchmark. PGO modes have separate default workspaces. Profiles and build commands are recorded; final performance comparisons still require controlled, separately coordinated runs.

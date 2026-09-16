"""Extract the real record update into an opaque call-boundary experiment.

This only generates sources. The separate harness controls compilation and
measurement. No production compiler or generated output is modified.
"""

from pathlib import Path
import hashlib
import json
import re


def sha256(data):
    return hashlib.sha256(data).hexdigest()


base = Path(__file__).resolve().parent
source = base.parents[1] / "output/purust_output/Purs_Test_Records/src/lib.rs"
source_bytes = source.read_bytes()
text = source_bytes.decode()
start = text.index("pub fn Test_Records_updateRec(")
end = text.index("pub fn Test_Records_initial(", start)
original = text[start:end]
(base / "kernel-original.rs").write_text(original)

assignment = original.index("let _tco_temp_1 =")
expression_start = original.index("{", assignment)
depth = 0
for index in range(expression_start, len(original)):
    char = original[index]
    if char == "{":
        depth += 1
    elif char == "}":
        depth -= 1
        if depth == 0:
            expression_end = index + 1
            break
else:
    raise ValueError("Unclosed generated record update expression")

expression = original[expression_start:expression_end]
assert original[expression_end] == ";"
assert expression.count("let _record_update_0 =") == 1
assert expression.count("let _record_child_update_0 =") == 1
assert expression.count("let _record_child_1_update_0 =") == 1
assert expression.count("let _record_child_1_update_1 =") == 1
assert expression.count("checked_rem_euclid") == 1
assert expression.count(".set_") == 8
renamed = re.sub(r"\bpurs_local_0\b", "n", expression)
renamed = re.sub(r"\bpurs_local_1\b", "r", renamed)
assert "purs_local_" not in renamed

callee = '''#![allow(warnings)]
use purust_core::*;

// The baseline below is extracted byte-for-byte from generated output.
include!("kernel-original.rs");

#[cfg(counted)]
use std::sync::atomic::{AtomicUsize, Ordering};
#[cfg(counted)]
static RECORD_CALLS: AtomicUsize = AtomicUsize::new(0);
#[cfg(counted)]
static SCALAR_CALLS: AtomicUsize = AtomicUsize::new(0);

#[cfg(counted)]
pub fn call_counts() -> [usize; 2] {
    [RECORD_CALLS.load(Ordering::Relaxed), SCALAR_CALLS.load(Ordering::Relaxed)]
}

#[cfg(counted)]
pub fn reset_counts() {
    RECORD_CALLS.store(0, Ordering::Relaxed);
    SCALAR_CALLS.store(0, Ordering::Relaxed);
}

// Exact generated update expression; only parameter identifiers are renamed.
#[inline(never)]
pub fn record_step(n: i64, r: Value) -> Value {
    #[cfg(counted)]
    RECORD_CALLS.fetch_add(1, Ordering::Relaxed);
''' + renamed + '''
}

// Four scalar fields follow the same update arithmetic as record_step.
// This function is compiled in the same separate crate as record_step.
#[inline(never)]
pub fn scalar_step(n: i64, fields: [i64; 4]) -> [i64; 4] {
    #[cfg(counted)]
    SCALAR_CALLS.fetch_add(1, Ordering::Relaxed);
    [
        fields[0] + 1,
        fields[1] + 2,
        fields[2] + 3,
        fields[3] + n.checked_rem_euclid(5).unwrap_or(0_i64),
    ]
}
'''
(base / "callee.rs").write_text(callee)
manifest = {
    "source": str(source),
    "source_sha256": sha256(source_bytes),
    "extracted_function": "Test_Records_updateRec",
    "kernel_original_sha256": sha256(original.encode()),
    "source_update_expression_sha256": sha256(expression.encode()),
    "record_step_expression_sha256": sha256(renamed.encode()),
    "callee_sha256": sha256(callee.encode()),
    "record_step_transformation": {
        "purs_local_0": "n",
        "purs_local_1": "r",
        "other_expression_changes": [],
    },
    "boundary": "Both step functions inline(never), in a separate Rust crate; build must disable LTO",
    "counters": "Atomic call counters exist only under --cfg counted",
    "scalar_arithmetic": ["a + 1", "c + 2", "e + 3", "f + n.checked_rem_euclid(5).unwrap_or(0_i64)"],
}
(base / "callee-manifest.json").write_text(json.dumps(manifest, indent=2) + "\n")
print(json.dumps(manifest, indent=2))

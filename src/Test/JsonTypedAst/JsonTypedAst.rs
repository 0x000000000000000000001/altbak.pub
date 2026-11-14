// JSON to Typed AST Rust driver, mirroring Test/JsonDecoding.rs and the Go and
// JavaScript drivers in the same suite:
//   * the corpus is a JSON array of { name, contents, benchmark? } entries read
//     from DIAG_CORPUS;
//   * `parse` measures Data.Argonaut.Parser.jsonParser, `decode` measures
//     decodeModule over the JSON value extracted once before timing, and
//     `combined` measures parseModule (parse + decode + usage validation);
//   * DIAG_PHASES selects the measured phases and their order exactly like the
//     Go driver; defaults are parse, decode, combined;
//   * the historical statistic stays two warm-ups plus five samples per phase,
//     with the process minimum reported as time_us. DIAG_WARMUPS and
//     DIAG_SAMPLES exist for focused profiling runs only;
//   * fingerprints are computed and validated outside the timed intervals;
//   * result destruction is timed separately after validation so the drop cost
//     of the decoded values is visible without contaminating decode/combined.
//
// A separate build with `--cfg diag_alloc` swaps mimalloc for the counting
// allocator below. DIAG_ALLOC=1 on that binary reports cumulative allocation
// and free totals per phase and per timed drop; the timed build never pays for
// the counters.
use std::io::Write;
use std::os::unix::io::FromRawFd;
use std::time::Instant;

fn driver_json_string(value: &str) -> String {
    Purs_Data_Argonaut_Core::purust_json_quote(value)
}

fn driver_json_number(value: f64) -> String {
    format!("{}", value)
}

fn driver_field(value: &crate::UnknownType, key: &str) -> crate::UnknownType {
    Purs_Data_Argonaut_Core::purust_json_object_get(value, key)
        .unwrap_or_else(|| panic!("Data.JsonTypedAst: corpus entry has no {}", key))
}

fn driver_usize_env(name: &str, default: usize) -> usize {
    std::env::var(name)
        .ok()
        .and_then(|value| value.parse::<usize>().ok())
        .unwrap_or(default)
}

/// Named, non-inlined profiling kernel: no fingerprint, no hashing, no
/// validation inside. The returned vector is dropped by the caller, so
/// destruction frames stay distinguishable from the decode frames.
#[inline(never)]
fn profile_batch(
    phase: &str,
    contents: &[String],
    parsed: &[crate::UnknownType],
    indices: &[usize],
    parse: &purust_core::Func1<String, crate::UnknownType>,
    decode: &purust_core::Func1<crate::UnknownType, crate::UnknownType>,
    decode_text: &purust_core::Func1<String, crate::UnknownType>,
) -> Vec<crate::UnknownType> {
    let mut results = Vec::with_capacity(indices.len());
    for index in indices {
        let result = match phase {
            "parse" => (*parse)(contents[*index].clone()),
            "decode" => (*decode)(parsed[*index].clone()),
            _ => (*decode_text)(contents[*index].clone()),
        };
        results.push(result);
    }
    results
}

/// Signals the profiling helper through an inherited pipe descriptor, so the
/// helper can start `sample` on an event instead of polling for readiness.
fn signal_ready() {
    if let Ok(fd) = std::env::var("DIAG_READY_FD") {
        if let Ok(fd) = fd.parse::<i32>() {
            let mut file = unsafe { std::fs::File::from_raw_fd(fd) };
            let _ = file.write_all(b"1");
            let _ = file.flush();
        }
    }
}

fn validate_result_one(
    phase: &str,
    index: usize,
    result: &crate::UnknownType,
    expected_json: &[String],
    expected_ast: &[String],
    encode: &purust_core::Func1<crate::UnknownType, String>,
    fingerprint: &purust_core::Func1<crate::UnknownType, String>,
) {
    let (raw, expected) = if phase == "parse" {
        ((*encode)(result.clone()), &expected_json[index])
    } else {
        ((*fingerprint)(result.clone()), &expected_ast[index])
    };
    let hash = Purs_Data_Argonaut_Core::purust_canonical_hash(&raw);
    assert_eq!(hash, *expected, "Data.JsonTypedAst: unstable {} output", phase);
}

/// Global allocator used only by the `--cfg diag_alloc` diagnostic build.
/// `realloc` is accounted as one free of the old block plus one allocation of
/// the new size, so `allocated_bytes - deallocated_bytes` equals the live byte
/// flow rather than the final resident size.
pub mod diag_alloc {
    use std::alloc::{GlobalAlloc, Layout};
    use std::sync::atomic::{AtomicU64, Ordering};

    pub static ALLOCATIONS: AtomicU64 = AtomicU64::new(0);
    pub static ALLOCATED_BYTES: AtomicU64 = AtomicU64::new(0);
    pub static DEALLOCATIONS: AtomicU64 = AtomicU64::new(0);
    pub static DEALLOCATED_BYTES: AtomicU64 = AtomicU64::new(0);

    pub struct CountingAlloc<T>(pub T);

    #[derive(Clone, Copy, Default)]
    pub struct Delta {
        pub allocations: u64,
        pub allocated_bytes: u64,
        pub deallocations: u64,
        pub deallocated_bytes: u64,
    }

    pub fn reset() {
        ALLOCATIONS.store(0, Ordering::Relaxed);
        ALLOCATED_BYTES.store(0, Ordering::Relaxed);
        DEALLOCATIONS.store(0, Ordering::Relaxed);
        DEALLOCATED_BYTES.store(0, Ordering::Relaxed);
    }

    pub fn delta() -> Delta {
        Delta {
            allocations: ALLOCATIONS.load(Ordering::Relaxed),
            allocated_bytes: ALLOCATED_BYTES.load(Ordering::Relaxed),
            deallocations: DEALLOCATIONS.load(Ordering::Relaxed),
            deallocated_bytes: DEALLOCATED_BYTES.load(Ordering::Relaxed),
        }
    }

    unsafe impl<T: GlobalAlloc> GlobalAlloc for CountingAlloc<T> {
        unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
            let pointer = self.0.alloc(layout);
            if !pointer.is_null() {
                ALLOCATIONS.fetch_add(1, Ordering::Relaxed);
                ALLOCATED_BYTES.fetch_add(layout.size() as u64, Ordering::Relaxed);
            }
            pointer
        }

        unsafe fn dealloc(&self, pointer: *mut u8, layout: Layout) {
            DEALLOCATIONS.fetch_add(1, Ordering::Relaxed);
            DEALLOCATED_BYTES.fetch_add(layout.size() as u64, Ordering::Relaxed);
            self.0.dealloc(pointer, layout);
        }

        unsafe fn alloc_zeroed(&self, layout: Layout) -> *mut u8 {
            let pointer = self.0.alloc_zeroed(layout);
            if !pointer.is_null() {
                ALLOCATIONS.fetch_add(1, Ordering::Relaxed);
                ALLOCATED_BYTES.fetch_add(layout.size() as u64, Ordering::Relaxed);
            }
            pointer
        }

        unsafe fn realloc(&self, pointer: *mut u8, layout: Layout, new_size: usize) -> *mut u8 {
            let new_pointer = self.0.realloc(pointer, layout, new_size);
            if !new_pointer.is_null() {
                DEALLOCATIONS.fetch_add(1, Ordering::Relaxed);
                DEALLOCATED_BYTES.fetch_add(layout.size() as u64, Ordering::Relaxed);
                ALLOCATIONS.fetch_add(1, Ordering::Relaxed);
                ALLOCATED_BYTES.fetch_add(new_size as u64, Ordering::Relaxed);
            }
            new_pointer
        }
    }
}

pub fn Test_JsonTypedAst_drive(
    parse: purust_core::Func1<String, crate::UnknownType>,
    decode: purust_core::Func1<crate::UnknownType, crate::UnknownType>,
    decode_text: purust_core::Func1<String, crate::UnknownType>,
    encode: purust_core::Func1<crate::UnknownType, String>,
    fingerprint: purust_core::Func1<crate::UnknownType, String>,
) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        let path = std::env::var("DIAG_CORPUS").expect("Data.JsonTypedAst: DIAG_CORPUS must be set");
        let raw = std::fs::read_to_string(&path)
            .unwrap_or_else(|error| panic!("Data.JsonTypedAst: cannot read {}: {}", path, error));
        let text = purust_core::purust_string_from_utf8(&raw);
        let corpus = Purs_Data_Argonaut_Core::purust_json_parse_text(&text)
            .unwrap_or_else(|error| panic!("Data.JsonTypedAst: corpus is not JSON: {}", error));
        let entries = Purs_Data_Argonaut_Core::purust_json_array_items(&corpus)
            .expect("Data.JsonTypedAst: corpus must be an array");
        assert!(!entries.is_empty(), "Data.JsonTypedAst: empty corpus");
        let mut names = Vec::with_capacity(entries.len());
        let mut contents = Vec::with_capacity(entries.len());
        let mut indices: Vec<usize> = Vec::with_capacity(entries.len());
        for (index, entry) in entries.iter().enumerate() {
            names.push(driver_field(entry, "name").unwrap_string());
            contents.push(driver_field(entry, "contents").unwrap_string());
            // Entries without an explicit `benchmark` flag are timed, matching
            // the frozen twelve-module corpus format.
            let timed = match Purs_Data_Argonaut_Core::purust_json_object_get(entry, "benchmark") {
                None => true,
                Some(value) => matches!(value.resolve(), crate::Value::Bool(true)),
            };
            if timed {
                indices.push(index);
            }
        }
        let modules = contents.len();
        assert!(!indices.is_empty(), "Data.JsonTypedAst: no timed cases");
        let parsed: Vec<crate::UnknownType> = contents
            .iter()
            .map(|content| parse(content.clone()))
            .collect();
        let expected_json: Vec<String> = contents
            .iter()
            .map(|content| Purs_Data_Argonaut_Core::purust_canonical_hash(content))
            .collect();
        let expected_ast: Vec<String> = parsed
            .iter()
            .map(|value| {
                let raw = fingerprint(decode(value.clone()));
                Purs_Data_Argonaut_Core::purust_canonical_hash(&raw)
            })
            .collect();
        if let Ok(profile_phase) = std::env::var("DIAG_PROFILE_PHASE") {
            // Dedicated profiling mode. Validation runs once before the
            // readiness signal and once after a natural loop completion; the
            // helper starts sampling only after readiness and terminates the
            // process, so expected fingerprints, corpus loading and report
            // construction never share a sampled window with the kernel.
            assert!(
                ["parse", "decode", "combined"].contains(&profile_phase.as_str()),
                "Data.JsonTypedAst: unknown profile phase {}",
                profile_phase
            );
            for index in &indices {
                let result = match profile_phase.as_str() {
                    "parse" => parse(contents[*index].clone()),
                    "decode" => decode(parsed[*index].clone()),
                    _ => decode_text(contents[*index].clone()),
                };
                validate_result_one(
                    &profile_phase, *index, &result, &expected_json, &expected_ast, &encode, &fingerprint,
                );
            }
            signal_ready();
            let passes = driver_usize_env("DIAG_PROFILE_PASSES", 10_000).max(1);
            let mut latest: Vec<crate::UnknownType> = Vec::new();
            for _ in 0..passes {
                latest = profile_batch(&profile_phase, &contents, &parsed, &indices, &parse, &decode, &decode_text);
                std::hint::black_box(&latest);
            }
            for (slot, result) in latest.iter().enumerate() {
                validate_result_one(
                    &profile_phase, indices[slot], result, &expected_json, &expected_ast, &encode, &fingerprint,
                );
            }
            println!(
                "{{\"profile_mode\":true,\"phase\":\"{}\",\"passes\":{}}}",
                profile_phase, passes
            );
            return crate::Value::Unit;
        }
        let warmups = driver_usize_env("DIAG_WARMUPS", 2);
        let samples = driver_usize_env("DIAG_SAMPLES", 5).max(1);
        let alloc_enabled = cfg!(diag_alloc)
            && std::env::var("DIAG_ALLOC").map(|value| value != "0").unwrap_or(false);
        let mut report = String::from("{\"backend\":\"rust\",\"allocator\":");
        report.push_str(if cfg!(diag_alloc) {
            "\"mimalloc+counting\""
        } else {
            "\"mimalloc\""
        });
        report.push_str(",\"alloc_diag\":");
        report.push_str(if alloc_enabled { "true" } else { "false" });
        report.push_str(",\"modules\":");
        report.push_str(&modules.to_string());
        report.push_str(",\"timed_cases\":");
        report.push_str(&indices.len().to_string());
        report.push_str(",\"warmups_per_phase\":");
        report.push_str(&warmups.to_string());
        report.push_str(",\"samples_per_phase\":");
        report.push_str(&samples.to_string());
        report.push_str(",\"names\":[");
        for (index, name) in names.iter().enumerate() {
            if index > 0 {
                report.push(',');
            }
            report.push_str(&driver_json_string(name));
        }
        report.push_str("],\"fingerprints\":[");
        for (index, hash) in expected_ast.iter().enumerate() {
            if index > 0 {
                report.push(',');
            }
            report.push_str(&driver_json_string(hash));
        }
        report.push_str("],\"json_fingerprints\":[");
        for (index, hash) in expected_json.iter().enumerate() {
            if index > 0 {
                report.push(',');
            }
            report.push_str(&driver_json_string(hash));
        }
        report.push_str("],\"phases\":{");
        // The Go and C drivers select the measured phases and their order
        // through DIAG_PHASES; keep the same protocol for paired campaigns.
        let requested = std::env::var("DIAG_PHASES").unwrap_or_else(|_| "parse,decode,combined".into());
        let mut phases: Vec<String> = Vec::new();
        for name in requested.split(',').map(|name| name.trim().to_owned()) {
            assert!(
                ["parse", "decode", "combined"].contains(&name.as_str()),
                "Data.JsonTypedAst: unknown phase {}",
                name
            );
            assert!(!phases.contains(&name), "Data.JsonTypedAst: duplicate phase {}", name);
            phases.push(name);
        }
        assert!(!phases.is_empty(), "Data.JsonTypedAst: no phases requested");
        let mut phase_order = String::from("[");
        for (index, phase) in phases.iter().enumerate() {
            if index > 0 {
                phase_order.push(',');
            }
            phase_order.push_str(&driver_json_string(phase));
        }
        phase_order.push(']');
        for (phase_index, phase) in phases.iter().enumerate() {
            if phase_index > 0 {
                report.push(',');
            }
            report.push('"');
            report.push_str(phase);
            report.push_str("\":{\"samples\":[");
            let mut best = f64::INFINITY;
            for pass in 0..(warmups + samples) {
                if alloc_enabled {
                    diag_alloc::reset();
                }
                let start = Instant::now();
                let mut results: Vec<crate::UnknownType> = Vec::with_capacity(indices.len());
                for index in &indices {
                    let result = match phase.as_str() {
                        "parse" => parse(contents[*index].clone()),
                        "decode" => decode(parsed[*index].clone()),
                        _ => decode_text(contents[*index].clone()),
                    };
                    results.push(result);
                }
                let elapsed = start.elapsed().as_nanos() as f64 / 1000.0;
                let phase_alloc = if alloc_enabled { Some(diag_alloc::delta()) } else { None };
                // Fingerprints and validation stay outside the timed interval.
                for (slot, result) in results.iter().enumerate() {
                    let index = indices[slot];
                    let (raw, expected) = if phase.as_str() == "parse" {
                        (encode(result.clone()), &expected_json[index])
                    } else {
                        (fingerprint(result.clone()), &expected_ast[index])
                    };
                    let hash = Purs_Data_Argonaut_Core::purust_canonical_hash(&raw);
                    assert_eq!(hash, *expected, "Data.JsonTypedAst: unstable {} output", phase);
                }
                // Destruction is timed separately, after validation.
                if alloc_enabled {
                    diag_alloc::reset();
                }
                let drop_start = Instant::now();
                drop(results);
                let drop_us = drop_start.elapsed().as_nanos() as f64 / 1000.0;
                let drop_alloc = if alloc_enabled { Some(diag_alloc::delta()) } else { None };
                if pass >= warmups {
                    if elapsed < best {
                        best = elapsed;
                    }
                    if pass > warmups {
                        report.push(',');
                    }
                    report.push_str("{\"time_us\":");
                    report.push_str(&driver_json_number(elapsed));
                    report.push_str(",\"drop_us\":");
                    report.push_str(&driver_json_number(drop_us));
                    if let Some(alloc) = phase_alloc {
                        report.push_str(",\"allocations\":");
                        report.push_str(&alloc.allocations.to_string());
                        report.push_str(",\"allocated_bytes\":");
                        report.push_str(&alloc.allocated_bytes.to_string());
                        report.push_str(",\"deallocations\":");
                        report.push_str(&alloc.deallocations.to_string());
                        report.push_str(",\"deallocated_bytes\":");
                        report.push_str(&alloc.deallocated_bytes.to_string());
                    }
                    if let Some(alloc) = drop_alloc {
                        report.push_str(",\"drop_allocations\":");
                        report.push_str(&alloc.allocations.to_string());
                        report.push_str(",\"drop_allocated_bytes\":");
                        report.push_str(&alloc.allocated_bytes.to_string());
                        report.push_str(",\"drop_deallocations\":");
                        report.push_str(&alloc.deallocations.to_string());
                        report.push_str(",\"drop_deallocated_bytes\":");
                        report.push_str(&alloc.deallocated_bytes.to_string());
                    }
                    report.push('}');
                }
            }
            report.push_str("],\"time_us\":");
            report.push_str(&driver_json_number(best));
            report.push('}');
        }
        report.push_str("},\"phase_order\":");
        report.push_str(&phase_order);
        if let Ok(version) = std::env::var("DIAG_RUSTC") {
            report.push_str(",\"rustc\":");
            report.push_str(&driver_json_string(&version));
        }
        if let Ok(profile) = std::env::var("DIAG_PROFILE") {
            report.push_str(",\"profile\":");
            report.push_str(&profile);
        }
        report.push_str("}");
        println!("{}", purust_core::purust_string_to_utf8_lossy(&report));
        crate::Value::Unit
    })))
}

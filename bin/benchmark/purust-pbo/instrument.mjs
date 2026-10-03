// Allocation/CPU attribution is a separate, untimed diagnostic binary.
import assert from 'node:assert/strict';
import { copyFileSync, existsSync, mkdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { dirname, join, relative, resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import { environment, hash, manifest, run, writeJson } from '../gopurs-aff/common.mjs';

const [rustArg, outArg, corpusArg] = process.argv.slice(2);
assert(rustArg && outArg && corpusArg, 'instrument.mjs GENERATED_RUST NEW_ARCHIVE COMPARE_AFF_RUNS');
const rust = resolve(rustArg), out = resolve(outArg), corpus = resolve(corpusArg);
assert(!existsSync(out)); mkdirSync(out, { recursive: true });
const files = { main: join(rust, 'src/main.rs'), metrics: join(rust, 'Purs_Purust_Metrics/src/lib.rs'),
  convert: join(rust, 'Purs_PureScript_Backend_Optimizer_Convert/src/lib.rs'),
  memo: join(rust, 'Purs_PureScript_Backend_Optimizer_BoundedMemo/src/lib.rs'),
  binary: join(rust, 'target/release/purust_output') };
const original = Object.fromEntries(Object.entries(files).map(([name, path]) => [name, readFileSync(path)]));
for (const [name, bytes] of Object.entries(original)) writeFileSync(join(out, 'original-' + name), bytes);
const result = { status: 'pending', binary_sha256: hash(original.binary),
  protocol: 'instrumented durations excluded; per-attempt thread CPU and allocation requests, phase-wide counters',
  original: Object.fromEntries(Object.entries(original).map(([name, bytes]) => [name, hash(bytes)])) };
const save = () => writeJson(join(out, 'results.json'), result);
save();
const counterPath = resolve(dirname(fileURLToPath(import.meta.url)), '../../../var/benchmark/purust-compiler-20261002/allocation-counter-parallel.rs');
const counter = readFileSync(counterPath, 'utf8') + `
#[no_mangle]
pub extern "C" fn purust_thread_alloc_counter(which: usize) -> u64 {
    SLOT.with(|slot| { let index = slot.get(); if index == usize::MAX { 0 } else {
        if which == 0 { COUNTERS[index].requests.load(Ordering::Relaxed) }
        else { COUNTERS[index].bytes.load(Ordering::Relaxed) }
    } })
}
`;
function replace(name, before, after) {
  const text = readFileSync(files[name], 'utf8'); assert.equal(text.split(before).length, 2, name);
  writeFileSync(files[name], text.replace(before, after));
}
try {
  replace('main', '#[global_allocator]\nstatic GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;', counter);
  replace('metrics', 'Value::Number(START.get_or_init', `
    extern "C" { fn purust_alloc_counter(which: usize) -> u64; }
    static CALLS: std::sync::atomic::AtomicUsize = std::sync::atomic::AtomicUsize::new(0);
    let call = CALLS.fetch_add(1, std::sync::atomic::Ordering::Relaxed);
    eprintln!("[alloc] {} {} {}", call, unsafe { purust_alloc_counter(0) }, unsafe { purust_alloc_counter(1) });
    Value::Number(START.get_or_init`);
  replace('memo', 'struct PurustMemo {', 'struct PurustMemo { probes: u64, hits: u64,');
  replace('memo', 'if let Some((_, _, result)) = cache.lock().unwrap().entries.get(key) { return result.clone(); }', `
    let mut state = cache.lock().unwrap(); state.probes += 1;
    if let Some((_, _, result)) = state.entries.get(key) {
        let result = result.clone(); state.hits += 1; return result;
    }`);
  writeFileSync(files.memo, readFileSync(files.memo, 'utf8') + `
impl Drop for PurustMemo {
    fn drop(&mut self) { eprintln!("[memo] {} {} {}", self.probes, self.hits, self.entries.len()); }
}
`);
  const source = readFileSync(files.convert, 'utf8');
  const headers = [...source.matchAll(/^pub fn PureScript_Backend_Optimizer_Convert_toBackendModuleWithLookup\([^\n]+\{\n/gm)];
  assert.equal(headers.length, 1);
  const header = headers[0][0];
  assert(header.includes('mut purs_local_0:') && header.includes('mut purs_local_1:'));
  replace('convert', header, header + `
    let _diag = PurustAttemptDiag::new(purs_local_1.get_name().unwrap_string());
    let _pending = _diag.pending.clone();
    let _lookup = purs_local_0;
    purs_local_0 = Func2::Shared(std::sync::Arc::new(move |module: String, ident: String| {
        let result = _lookup(module.clone(), ident);
        if matches!(result.as_ref(), ExternLookup::ExternPending) { _pending.lock().unwrap().insert(module); }
        result
    }));
`);
  writeFileSync(files.convert, readFileSync(files.convert, 'utf8') + `
#[repr(C)] struct PurustTimespec { sec: i64, nsec: i64 }
extern "C" {
    fn clock_gettime(clock: u32, out: *mut PurustTimespec) -> i32;
    fn purust_thread_alloc_counter(which: usize) -> u64;
}
fn purust_thread_cpu() -> u64 {
    let mut time = PurustTimespec { sec: 0, nsec: 0 };
    assert_eq!(unsafe { clock_gettime(16, &mut time) }, 0); // Darwin CLOCK_THREAD_CPUTIME_ID
    (time.sec as u64) * 1_000_000_000 + time.nsec as u64
}
struct PurustAttemptDiag {
    name: String, cpu: u64, requests: u64, bytes: u64, wall: std::time::Instant,
    pending: std::sync::Arc<std::sync::Mutex<std::collections::BTreeSet<String>>>,
}
impl PurustAttemptDiag {
    fn new(name: String) -> Self { Self { name, cpu: purust_thread_cpu(),
        requests: unsafe { purust_thread_alloc_counter(0) }, bytes: unsafe { purust_thread_alloc_counter(1) },
        wall: std::time::Instant::now(), pending: Default::default() } }
}
impl Drop for PurustAttemptDiag {
    fn drop(&mut self) {
        let cpu = purust_thread_cpu() - self.cpu;
        let requests = unsafe { purust_thread_alloc_counter(0) } - self.requests;
        let bytes = unsafe { purust_thread_alloc_counter(1) } - self.bytes;
        let pending = self.pending.lock().unwrap();
        eprintln!("[attempt] {}\\t{}\\t{}\\t{}\\t{}\\t{}\\t{}", self.name, !pending.is_empty(),
            cpu, self.wall.elapsed().as_nanos(), requests, bytes, pending.iter().cloned().collect::<Vec<_>>().join(","));
    }
}
`);
  for (const name of ['main', 'metrics', 'convert', 'memo']) copyFileSync(files[name], join(out, 'instrumented-' + name + '.rs'));
  result.build = run(out, 'cargo', 'cargo', ['build', '--offline', '--release', '--config', 'profile.release.lto=false',
    '--config', 'profile.release.opt-level=3', '--manifest-path', join(rust, 'Cargo.toml')], rust,
    { ...environment(), CARGO_BUILD_JOBS: '8', CARGO_INCREMENTAL: '0' }, 1200000);
  const binary = join(out, 'compiler-counted'); copyFileSync(files.binary, binary);
  result.instrumented_sha256 = hash(readFileSync(binary)); save();
  const cwd = join(corpus, 'inputs/gopurs-aff');
  for (const name of ['.purmeta', '.cache']) rmSync(join(cwd, name), { recursive: true, force: true });
  result.run = run(out, 'diagnostic', '/usr/bin/time', ['-l', binary, '--main', 'Test.Main', '--threaded',
    '--source', 'output', '--out', join(out, 'generated'), '--ffi-dir', relative(cwd, join(corpus, 'rust-ffi'))], cwd);
  const log = readFileSync(result.run.stderr, 'utf8');
  result.memo_caches = [...log.matchAll(/^\[memo\] (\d+) (\d+) (\d+)$/gm)]
    .map(([, probes, hits, size]) => ({ probes: Number(probes), hits: Number(hits), size: Number(size) }));
  result.memo_totals = Object.fromEntries(['probes', 'hits'].map(key => [key,
    result.memo_caches.reduce((sum, cache) => sum + cache[key], 0)]));
  result.attempts = [...log.matchAll(/^\[attempt\] (.+)\t(true|false)\t(\d+)\t(\d+)\t(\d+)\t(\d+)\t(.*)$/gm)]
    .map(([, module, deferred, cpu, wall, requests, bytes, pending]) => ({ module, deferred: deferred === 'true',
      cpu_ns: Number(cpu), wall_ns: Number(wall), requests: Number(requests), bytes: Number(bytes), pending: pending ? pending.split(',') : [] }));
  assert(result.attempts.length >= 238);
  result.attempt_totals = Object.fromEntries([false, true].map(deferred => {
    const attempts = result.attempts.filter(attempt => attempt.deferred === deferred);
    return [deferred ? 'rejected' : 'accepted', { count: attempts.length,
      ...Object.fromEntries(['cpu_ns', 'wall_ns', 'requests', 'bytes'].map(key => [key, attempts.reduce((sum, a) => sum + a[key], 0)])) }];
  }));
  result.snapshots = [...log.matchAll(/^\[alloc\] (\d+) (\d+) (\d+)$/gm)]
    .map(([, call, requests, bytes]) => ({ call: Number(call), requests: Number(requests), bytes: Number(bytes) }));
  assert.equal(result.snapshots.length, 10);
  const delta = (a, b) => Object.fromEntries(['requests', 'bytes'].map(key => [key, result.snapshots[b][key] - result.snapshots[a][key]]));
  result.phases = { total: delta(0, 9), load: delta(1, 2), prepare: delta(3, 4), optimize_generate: delta(5, 6), finalize: delta(7, 8) };
  const prior = JSON.parse(readFileSync(corpus.replace(/-runs$/, '.json'), 'utf8'));
  assert.deepEqual(manifest(join(out, 'generated'), path => /\.(rs|toml)$/.test(path)), prior.provenance.reference_files);
  result.identical_files = prior.provenance.reference_files.length;
  result.status = 'passed'; save();
  console.log(JSON.stringify({ attempts: result.attempt_totals, phases: result.phases }, null, 2));
} catch (error) { result.status = 'failed'; result.error = error.stack; save(); throw error; }
finally {
  for (const [name, bytes] of Object.entries(original)) writeFileSync(files[name], bytes);
  for (const [name, bytes] of Object.entries(original)) assert.equal(hash(readFileSync(files[name])), hash(bytes));
  result.restored = true; save();
}

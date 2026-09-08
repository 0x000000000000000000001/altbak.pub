// Run after npm run build. Measure Unit construction separately from its passage
// through generated functions; use the real Data.Unit FFI and generated runtime.
import assert from 'node:assert/strict';
import { mkdtempSync, readFileSync, writeFileSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { spawnSync } from 'node:child_process';
import { fileURLToPath } from 'node:url';
import { codegenModule, codegenPrelude } from '../../output/Purust.CodeGen/index.js';
import { empty as emptyMap } from '../../output/Data.Map/index.js';
import { empty as emptySet } from '../../output/Data.Set/index.js';
import { Just } from '../../output/Data.Maybe/index.js';
import { Tuple } from '../../output/Data.Tuple/index.js';
import { Any, Func, Int, Unit } from '../../output/PureScript.Backend.Optimizer.CoreFn/index.js';
import { Abs, App, Local, Typed } from '../../output/PureScript.Backend.Optimizer.Syntax/index.js';

const local = (name, level) => new Local(new Just(name), level);
const param = (name, level) => new Tuple(new Just(name), level);
const bindings = [new Tuple('identity', new Typed(new Func([Unit.value], Unit.value),
  new Abs([param('unit', 0)], local('unit', 0))))];
for (const [name, resultType] of [['integer', Int.value], ['erased', Any.value]]) {
  const continuationType = new Func([Unit.value], resultType);
  bindings.push(new Tuple(name, new Typed(new Func([continuationType, Unit.value], resultType),
    new Abs([param('continuation', 0), param('unit', 1)],
      new Typed(resultType, new App(local('continuation', 0), [local('unit', 1)]))))));
}
const generated = codegenModule(emptyMap)(emptyMap)(
  { name: 'UnitValues', dataDecls: [], classDecls: [] },
)({ name: 'UnitValues', bindings: [{ recursive: false, bindings }] });
const unitFFI = readFileSync(new URL('../../../purust-prelude/src/Data/Unit.rs', import.meta.url), 'utf8');
const runtime = fileURLToPath(new URL('../runtime/perceus_ptr/src/lib.rs', import.meta.url));
const rust = `${codegenPrelude(emptySet)}
extern crate self as purust_core;
#[path = ${JSON.stringify(runtime)}]
mod perceus_ptr;
${unitFFI}
${generated}

use std::alloc::{GlobalAlloc, Layout, System};
use std::sync::atomic::{AtomicUsize, Ordering};
static ALLOCATIONS: AtomicUsize = AtomicUsize::new(0);
static BYTES: AtomicUsize = AtomicUsize::new(0);
static DEALLOCATIONS: AtomicUsize = AtomicUsize::new(0);
static CALLS: AtomicUsize = AtomicUsize::new(0);
struct CountingAllocator;
unsafe impl GlobalAlloc for CountingAllocator {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        ALLOCATIONS.fetch_add(1, Ordering::Relaxed);
        BYTES.fetch_add(layout.size(), Ordering::Relaxed);
        System.alloc(layout)
    }
    unsafe fn dealloc(&self, pointer: *mut u8, layout: Layout) {
        DEALLOCATIONS.fetch_add(1, Ordering::Relaxed);
        System.dealloc(pointer, layout)
    }
}
#[global_allocator] static ALLOCATOR: CountingAllocator = CountingAllocator;
fn counts() -> (usize, usize, usize) {
    (ALLOCATIONS.load(Ordering::Relaxed), BYTES.load(Ordering::Relaxed),
     DEALLOCATIONS.load(Ordering::Relaxed))
}
fn delta(before: (usize, usize, usize)) -> (usize, usize, usize) {
    let after = counts();
    (after.0 - before.0, after.1 - before.1, after.2 - before.2)
}
fn main() {
    // Different results expose accidental replacement of a continuation result
    // by Unit, while the counter exposes missing or premature calls.
    for expected in [-7, 0, 42] {
        let before = CALLS.load(Ordering::Relaxed);
        let continuation = Func1::Shared(std::rc::Rc::new(move |_| {
            CALLS.fetch_add(1, Ordering::Relaxed);
            expected
        }));
        let unit = UnitValues_identity(Data_Unit_unit());
        assert_eq!(CALLS.load(Ordering::Relaxed), before);
        assert_eq!(UnitValues_integer(continuation.clone(), unit.clone()), expected);
        assert_eq!(UnitValues_integer(continuation, unit), expected);
        assert_eq!(CALLS.load(Ordering::Relaxed), before + 2);
    }
    assert_eq!(UnitValues_erased(Func1::Static(|_| mk_int(43)), Data_Unit_unit()).unwrap_int(), 43);
    assert!(UnitValues_erased(Func1::Static(|_| mk_bool(true)), Data_Unit_unit()).unwrap_bool());
    let before = CALLS.load(Ordering::Relaxed);
    let returned = UnitValues_erased(Func1::Static(|_| {
        Value::Func1(Func1::Static(|_| {
            CALLS.fetch_add(1, Ordering::Relaxed);
            mk_int(99)
        }))
    }), Data_Unit_unit());
    assert_eq!(CALLS.load(Ordering::Relaxed), before, "a returned function is a value");
    assert_eq!(returned.unwrap_func1()(mk_int(0)).unwrap_int(), 99);
    assert_eq!(CALLS.load(Ordering::Relaxed), before + 1);

    let before = counts();
    for _ in 0..1000 {
        drop(std::hint::black_box(Data_Unit_unit()));
    }
    let construction = delta(before);

    let unit = Data_Unit_unit();
    let continuation = Func1::Static(|_| 42);
    let before = counts();
    for _ in 0..1000 {
        assert_eq!(UnitValues_integer(continuation.clone(),
            UnitValues_identity(std::hint::black_box(unit.clone()))), 42);
    }
    let passage = delta(before);
    assert_eq!(passage, (0, 0, 0), "passing an existing Unit must not allocate");
    assert_eq!(construction.0, construction.2, "constructed Units must be released");
    println!("Unit construction x1000: allocations={}, requested_bytes={}, deallocations={}",
             construction.0, construction.1, construction.2);
    println!("Unit passage x1000: allocations={}, requested_bytes={}, deallocations={}",
             passage.0, passage.1, passage.2);
    println!("Minimal runtime: Value={} bytes, Record_a={} bytes",
             std::mem::size_of::<Value>(), std::mem::size_of::<Record_a>());
}
`;
const directory = mkdtempSync(join(tmpdir(), 'purust-unit-values-'));
try {
  const source = join(directory, 'unit-values.rs');
  const binary = join(directory, 'unit-values');
  writeFileSync(source, rust);
  const build = spawnSync('rustc', ['--edition=2021', '-C', 'opt-level=1', source, '-o', binary], { encoding: 'utf8' });
  assert.equal(build.status, 0, `rustc: ${build.error ?? ''}\n${build.stdout}\n${build.stderr}`);
  const run = spawnSync(binary, [], { encoding: 'utf8' });
  assert.equal(run.status, 0, `${run.error ?? ''}\n${run.stdout}\n${run.stderr}`);
  process.stdout.write(run.stdout);
} finally {
  rmSync(directory, { recursive: true, force: true });
}

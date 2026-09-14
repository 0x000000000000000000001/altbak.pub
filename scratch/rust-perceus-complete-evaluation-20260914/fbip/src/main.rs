//! Synthetic native ADT study. All stages use the same Rc layout and update
//! plan; only copying, payload reuse, retained uniqueness and field mutation
//! differ. This is a compiler-representation experiment, not emitted Purust.
use std::hint::black_box;
use std::rc::{Rc, Weak};
use std::time::Instant;

#[global_allocator]
static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;

#[derive(Clone, Copy, Default)]
struct Metrics {
    allocations: u64, cell_drops: u64, dups: u64, drops: u64,
    unique_checks: u64, unique_success: u64, shared_rejections: u64, weak_rejections: u64,
    payload_takes: u64, payload_rebuilds: u64, specialized_actions: u64,
    weak_observers: u64, weak_drops: u64,
}
#[cfg(feature = "counts")]
thread_local! { static METRICS: std::cell::Cell<Metrics> = std::cell::Cell::new(Metrics::default()); }
macro_rules! count {
    ($field:ident) => {
        #[cfg(feature = "counts")]
        METRICS.with(|slot| { let mut m = slot.get(); m.$field += 1; slot.set(m); });
    };
}
fn metrics() -> Metrics {
    #[cfg(feature = "counts")]
    { return METRICS.with(|slot| slot.get()); }
    #[cfg(not(feature = "counts"))]
    Metrics::default()
}
fn reset_metrics() {
    #[cfg(feature = "counts")]
    METRICS.with(|slot| slot.set(Metrics::default()));
}

struct Ptr(Rc<CellData>);
struct CellData { tree: Tree }
enum Tree { Empty, Node { a: i64, b: i64, left: Ptr, right: Ptr } }
impl Ptr {
    fn new(tree: Tree) -> Self { count!(allocations); Self(Rc::new(CellData { tree })) }
    fn tree(&self) -> &Tree { &self.0.tree }
    fn node(a: i64, b: i64, left: Ptr, right: Ptr) -> Self {
        Self::new(Tree::Node { a, b, left, right })
    }
    fn address(&self) -> usize { Rc::as_ptr(&self.0) as usize }
}
impl Clone for Ptr {
    fn clone(&self) -> Self { count!(dups); Self(self.0.clone()) }
}
impl Drop for Ptr { fn drop(&mut self) { count!(drops); } }
impl Drop for CellData { fn drop(&mut self) { count!(cell_drops); } }
struct Observer(Weak<CellData>);
impl Observer {
    fn new(ptr: &Ptr) -> Self { count!(weak_observers); Self(Rc::downgrade(&ptr.0)) }
}
impl Drop for Observer { fn drop(&mut self) { count!(weak_drops); } }

fn unique(root: &mut Ptr) -> Option<&mut CellData> {
    count!(unique_checks);
    #[cfg(feature = "counts")]
    {
        if Rc::strong_count(&root.0) != 1 { count!(shared_rejections); }
        else if Rc::weak_count(&root.0) != 0 { count!(weak_rejections); }
    }
    let result = Rc::get_mut(&mut root.0);
    if result.is_some() { count!(unique_success); }
    result
}

#[derive(Clone, Copy)]
enum Action { Scalar, Pair, Swap }
#[derive(Clone, Copy)]
struct Update { bits: u64, remaining: u32, action: Action, delta: i64 }
impl Update {
    fn next(self) -> Self { Self { bits: self.bits >> 1, remaining: self.remaining - 1, ..self } }
    fn apply(self, a: &mut i64, b: &mut i64, left: &mut Ptr, right: &mut Ptr) {
        match self.action {
            Action::Scalar => *b += self.delta,
            Action::Pair => { *a += self.delta; *b -= self.delta; }
            Action::Swap => std::mem::swap(left, right),
        }
    }
}
type Worker = fn(&mut Ptr, Update);

fn rebuild(mut tree: Tree, update: Update, worker: Worker) -> Tree {
    let Tree::Node { a, b, left, right } = &mut tree else { panic!("plan reached Empty") };
    if update.remaining == 0 { update.apply(a, b, left, right); }
    else if update.bits & 1 == 0 { worker(left, update.next()); }
    else { worker(right, update.next()); }
    tree
}

fn copy_update(root: &mut Ptr, update: Update, worker: Worker) {
    let Tree::Node { a, b, left, right } = root.tree() else { panic!("plan reached Empty") };
    let copy = Tree::Node { a: *a, b: *b, left: left.clone(), right: right.clone() };
    // The old parent stays alive while the updated child is computed, exactly
    // as on a persistent reconstruction path. All stages share this fallback.
    let result = rebuild(copy, update, worker);
    *root = Ptr::new(result);
    count!(payload_rebuilds);
}
fn persistent(root: &mut Ptr, update: Update) { copy_update(root, update, persistent); }

fn reuse_rechecked(root: &mut Ptr, update: Update) {
    let detached = unique(root).map(|cell| {
        count!(payload_takes);
        std::mem::replace(&mut cell.tree, Tree::Empty)
    });
    if let Some(tree) = detached {
        let result = rebuild(tree, update, reuse_rechecked);
        unique(root).expect("detached root cannot acquire owners or weak observers").tree = result;
        count!(payload_rebuilds);
    } else { copy_update(root, update, reuse_rechecked); }
}
fn reuse_retained(root: &mut Ptr, update: Update) {
    if let Some(cell) = unique(root) {
        count!(payload_takes);
        let tree = std::mem::replace(&mut cell.tree, Tree::Empty);
        cell.tree = rebuild(tree, update, reuse_retained);
        count!(payload_rebuilds);
    } else { copy_update(root, update, reuse_retained); }
}
fn fields_retained(root: &mut Ptr, update: Update) {
    if let Some(cell) = unique(root) {
        let Tree::Node { a, b, left, right } = &mut cell.tree else { panic!("plan reached Empty") };
        if update.remaining == 0 {
            update.apply(a, b, left, right);
            count!(specialized_actions);
        } else if update.bits & 1 == 0 { fields_retained(left, update.next()); }
        else { fields_retained(right, update.next()); }
    } else { copy_update(root, update, fields_retained); }
}

#[derive(Clone, PartialEq, Eq, Debug)]
enum Model { Empty, Node { a: i64, b: i64, left: Box<Model>, right: Box<Model> } }
fn model(depth: u32, seed: i64) -> Model {
    if depth == 0 { Model::Empty } else { Model::Node { a: seed, b: seed * 3,
        left: Box::new(model(depth - 1, seed * 2)), right: Box::new(model(depth - 1, seed * 2 + 1)) } }
}
fn native(depth: u32, seed: i64) -> Ptr {
    if depth == 0 { Ptr::new(Tree::Empty) } else {
        Ptr::node(seed, seed * 3, native(depth - 1, seed * 2), native(depth - 1, seed * 2 + 1))
    }
}
fn observe(root: &Ptr) -> Model {
    match root.tree() {
        Tree::Empty => Model::Empty,
        Tree::Node { a, b, left, right } => Model::Node { a: *a, b: *b,
            left: Box::new(observe(left)), right: Box::new(observe(right)) },
    }
}
fn oracle_update(root: &mut Model, update: Update) {
    let Model::Node { a, b, left, right } = root else { panic!("invalid oracle plan") };
    if update.remaining == 0 {
        match update.action {
            Action::Scalar => *b += update.delta,
            Action::Pair => { *a += update.delta; *b -= update.delta; }
            Action::Swap => std::mem::swap(left, right),
        }
    } else if update.bits & 1 == 0 { oracle_update(left, update.next()); }
    else { oracle_update(right, update.next()); }
}
fn mix(a: i64, b: i64, left: u64, right: u64) -> u64 {
    (a as u64).wrapping_mul(31).wrapping_add(b as u64).wrapping_mul(17)
        ^ left.rotate_left(7) ^ right.rotate_right(11)
}
fn checksum(root: &Ptr) -> u64 {
    match root.tree() {
        Tree::Empty => 0x517cc1b727220a95,
        Tree::Node { a, b, left, right } => mix(*a, *b, checksum(left), checksum(right)),
    }
}
fn model_checksum(root: &Model) -> u64 {
    match root {
        Model::Empty => 0x517cc1b727220a95,
        Model::Node { a, b, left, right } => mix(*a, *b, model_checksum(left), model_checksum(right)),
    }
}
fn plans(depth: u32, n: usize) -> Vec<Update> {
    (0..n).map(|i| Update { bits: (i as u64).wrapping_mul(0x9e3779b97f4a7c15),
        remaining: match i % 3 { 0 => depth - 1, 1 => depth / 2, _ => depth / 3 },
        action: match i % 3 { 0 => Action::Scalar, 1 => Action::Pair, _ => Action::Swap },
        delta: (i % 7 + 1) as i64 }).collect()
}

const STAGES: [&str; 4] = ["persistent", "reuse_rechecked", "reuse_retained", "fields_retained"];
const SCENARIOS: [&str; 6] = ["unique", "old-root", "mixed", "diamond", "weak-root", "weak-mixed"];
fn worker(stage: &str) -> Worker {
    match stage { "persistent" => persistent, "reuse_rechecked" => reuse_rechecked,
        "reuse_retained" => reuse_retained, "fields_retained" => fields_retained,
        _ => panic!("unknown stage") }
}
struct Expected { final_root: Model, held: Vec<Model> }
fn expected(scenario: &str, depth: u32, updates: &[Update]) -> Expected {
    let mut root = model(depth, 1);
    let mut held = Vec::new();
    let Model::Node { left, right, .. } = &mut root else { unreachable!() };
    if scenario == "diamond" { *right = left.clone(); }
    if scenario == "mixed" {
        held.push((**left).clone());
        let Model::Node { left: inner, .. } = right.as_ref() else { unreachable!() };
        held.push((**inner).clone());
    }
    for (i, update) in updates.iter().enumerate() {
        if scenario == "old-root" && i % (updates.len() / 8).max(1) == 0 { held.push(root.clone()); }
        oracle_update(&mut root, *update);
    }
    Expected { final_root: root, held }
}

struct Outcome { hash: u64, retained: usize, weak: usize, root_same: bool, metrics: Metrics }
fn execute(stage: &str, scenario: &str, depth: u32, updates: &[Update], oracle: Option<&Expected>) -> Outcome {
    reset_metrics();
    let mut root = native(depth, 1);
    let mut held = Vec::new();
    let mut observers = Vec::new();
    if scenario == "diamond" {
        let Tree::Node { left, right, .. } = &mut Rc::get_mut(&mut root.0).unwrap().tree else { unreachable!() };
        *right = left.clone();
    }
    if scenario == "weak-root" { observers.push(Observer::new(&root)); }
    let Tree::Node { left, right, .. } = root.tree() else { unreachable!() };
    if scenario == "mixed" {
        held.push(left.clone());
        let Tree::Node { left: inner, .. } = right.tree() else { unreachable!() };
        held.push(inner.clone());
    }
    if scenario == "weak-mixed" {
        observers.push(Observer::new(left));
        let Tree::Node { left: inner, right: inner_right, .. } = right.tree() else { unreachable!() };
        observers.push(Observer::new(inner)); observers.push(Observer::new(inner_right));
    }
    let root_address = root.address();
    let update = worker(stage);
    for (i, operation) in updates.iter().enumerate() {
        if scenario == "old-root" && i % (updates.len() / 8).max(1) == 0 { held.push(root.clone()); }
        update(&mut root, *operation);
    }
    let hash = black_box(checksum(black_box(&root)));
    if let Some(oracle) = oracle {
        assert_eq!(observe(&root), oracle.final_root, "{stage}/{scenario}: final tree");
        assert_eq!(held.len(), oracle.held.len());
        for (index, (actual, expected)) in held.iter().zip(&oracle.held).enumerate() {
            assert_eq!(observe(actual), *expected, "{stage}/{scenario}: retained handle {index}");
        }
    }
    let retained = held.len(); let weak = observers.len(); let root_same = root.address() == root_address;
    if stage != "persistent" && matches!(scenario, "unique" | "mixed" | "diamond" | "weak-mixed") {
        assert!(root_same, "eligible unique root must keep its cell");
    }
    drop(root); drop(held);
    for observer in &observers { assert_eq!(observer.0.strong_count(), 0, "all strong owners must be released"); }
    drop(observers);
    let result = metrics();
    #[cfg(feature = "counts")]
    {
        assert_eq!(result.allocations, result.cell_drops, "every allocated payload dropped");
        assert_eq!(result.allocations + result.dups, result.drops, "every pointer owner dropped");
        assert_eq!(result.weak_observers, result.weak_drops, "every weak observer dropped");
    }
    Outcome { hash, retained, weak, root_same, metrics: result }
}
fn print_outcome(stage: &str, scenario: &str, depth: u32, n: usize, result: Outcome) {
    let m = result.metrics;
    println!(concat!("{{\"stage\":\"{}\",\"scenario\":\"{}\",\"depth\":{},\"updates\":{},",
        "\"checksum\":{},\"retained_handles\":{},\"weak_observers\":{},\"root_same\":{},",
        "\"allocations\":{},\"cell_drops\":{},\"dups\":{},\"drops\":{},\"unique_checks\":{},",
        "\"unique_success\":{},\"shared_rejections\":{},\"weak_rejections\":{},",
        "\"payload_takes\":{},\"payload_rebuilds\":{},\"specialized_actions\":{},\"weak_created\":{},\"weak_drops\":{}}}"),
        stage, scenario, depth, n, result.hash, result.retained, result.weak, result.root_same,
        m.allocations, m.cell_drops, m.dups, m.drops, m.unique_checks, m.unique_success,
        m.shared_rejections, m.weak_rejections, m.payload_takes, m.payload_rebuilds, m.specialized_actions, m.weak_observers, m.weak_drops);
}

fn main() {
    let args: Vec<String> = std::env::args().collect();
    match args.get(1).map(String::as_str) {
        Some("validate") => {
            for depth in [7, 10] { for n in [128, 2048] {
                let updates = plans(depth, n);
                for scenario in SCENARIOS {
                    let oracle = expected(scenario, depth, &updates);
                    let hash = model_checksum(&oracle.final_root);
                    for stage in STAGES {
                        let result = execute(stage, scenario, depth, &updates, Some(&oracle));
                        assert_eq!(result.hash, hash);
                        print_outcome(stage, scenario, depth, n, result);
                    }
                }
            } }
        }
        Some("smoke") | Some("time") => {
            let timing = args[1] == "time";
            assert!(!timing || args.iter().any(|s| s == "--coordinated"), "timing requires coordinator signal");
            assert!(!timing || !cfg!(feature = "counts"), "no instrumented timing");
            let stage = &args[2]; let scenario = &args[3];
            assert!(SCENARIOS.contains(&scenario.as_str()));
            let depth: u32 = args[4].parse().unwrap(); let n: usize = args[5].parse().unwrap();
            assert!((3..=14).contains(&depth) && n > 0);
            let updates = plans(depth, n);
            let oracle = expected(scenario, depth, &updates); let hash = model_checksum(&oracle.final_root);
            assert_eq!(execute(stage, scenario, depth, &updates, Some(&oracle)).hash, hash);
            if timing {
                for _ in 0..7 {
                    let start = Instant::now();
                    let result = execute(stage, scenario, depth, black_box(&updates), None);
                    let elapsed = start.elapsed().as_nanos();
                    assert_eq!(result.hash, hash);
                    println!("{{\"elapsed_ns\":{elapsed},\"checksum\":{hash}}}");
                }
            } else { println!("{stage}/{scenario}: {depth} levels, {n} updates, oracle/lifetimes passed"); }
        }
        _ => panic!("use validate, smoke STAGE SCENARIO DEPTH UPDATES, or coordinated time"),
    }
}

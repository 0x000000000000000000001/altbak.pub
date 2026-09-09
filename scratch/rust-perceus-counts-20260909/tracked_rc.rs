// Count logical Rc operations; never use this instrumentation for timings.
use std::cell::{Cell, RefCell};
use std::collections::BTreeMap;
use std::ops::Deref;
use std::panic::Location;
use std::rc::Rc as NativeRc;

thread_local! {
    static PHASE: Cell<&'static str> = const { Cell::new("checks") };
    static EVENTS: RefCell<BTreeMap<(&'static str, &'static str, &'static str, u32), u64>> = RefCell::new(BTreeMap::new());
}
pub trait Kind { fn kind(&self) -> &'static str; }
fn record(op: &'static str, kind: &'static str, line: u32) {
    PHASE.with(|phase| EVENTS.with(|events| {
        *events.borrow_mut().entry((phase.get(), op, kind, line)).or_default() += 1;
    }));
}
pub fn phase(name: &'static str) { PHASE.with(|phase| phase.set(name)); }
pub fn print_events() {
    EVENTS.with(|events| {
        for ((phase, op, kind, line), count) in events.borrow().iter() {
            println!("{phase}\t{op}\t{kind}\t{line}\t{count}");
        }
    });
}

// Option<Rc<T>> retains Rc's one-word layout and lets unwrap_or_clone consume
// the native pointer without counting the wrapper's later empty Drop twice.
pub struct Rc<T: Kind>(Option<NativeRc<T>>);
impl<T: Kind> Rc<T> {
    #[track_caller]
    pub fn new(value: T) -> Self {
        record("new", value.kind(), Location::caller().line());
        Self(Some(NativeRc::new(value)))
    }
    pub fn as_ref(&self) -> &T { self.0.as_ref().unwrap().as_ref() }
    #[track_caller]
    pub fn get_mut(this: &mut Self) -> Option<&mut T> {
        let pointer = this.0.as_mut().unwrap();
        let op = if NativeRc::strong_count(pointer) != 1 { "get_mut_shared" }
            else if NativeRc::weak_count(pointer) != 0 { "get_mut_weak" }
            else { "get_mut_unique" };
        record(op, pointer.kind(), Location::caller().line());
        NativeRc::get_mut(pointer)
    }
    pub fn downgrade(this: &Self) -> std::rc::Weak<T> {
        NativeRc::downgrade(this.0.as_ref().unwrap())
    }
}
impl<T: Kind + Clone> Rc<T> {
    #[track_caller]
    pub fn unwrap_or_clone(mut this: Self) -> T {
        let pointer = this.0.take().unwrap();
        let op = if NativeRc::strong_count(&pointer) == 1 { "unwrap_unique" }
            else { "unwrap_shared" };
        record(op, pointer.kind(), Location::caller().line());
        // On the shared path T::clone recursively invokes the instrumented
        // child-pointer clones. The consumed outer reference is counted above.
        NativeRc::unwrap_or_clone(pointer)
    }
}
impl<T: Kind> Clone for Rc<T> {
    #[track_caller]
    fn clone(&self) -> Self {
        record("clone", self.kind(), Location::caller().line());
        Self(Some(self.0.as_ref().unwrap().clone()))
    }
}
impl<T: Kind> Drop for Rc<T> {
    fn drop(&mut self) {
        if let Some(pointer) = &self.0 {
            let op = if NativeRc::strong_count(pointer) == 1 { "drop_last" } else { "drop_shared" };
            record(op, pointer.kind(), 0);
        }
        // Rust then drops the native pointer, recursively dropping its payload
        // only for the last owner. No counter borrow survives into this step.
    }
}
impl<T: Kind> Deref for Rc<T> {
    type Target = T;
    fn deref(&self) -> &T { self.as_ref() }
}

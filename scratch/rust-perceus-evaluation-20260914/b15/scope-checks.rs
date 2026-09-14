// These deliberately ineligible payloads/callbacks establish why the plain
// native Tree prototype must not become an unrestricted consuming rewrite.
use std::cell::RefCell;
use std::rc::{Rc, Weak};

#[derive(Clone)]
struct Payload { id: i32, log: Rc<RefCell<Vec<i32>>> }
impl Drop for Payload {
    fn drop(&mut self) { self.log.borrow_mut().push(self.id); }
}
#[derive(Clone)]
enum DTree { E, T(Payload, Rc<DTree>, Rc<DTree>) }

fn fixture(log: &Rc<RefCell<Vec<i32>>>) -> Rc<DTree> {
    let leaf = |id| Rc::new(DTree::T(Payload { id, log: log.clone() }, Rc::new(DTree::E), Rc::new(DTree::E)));
    Rc::new(DTree::T(Payload { id: 0, log: log.clone() }, leaf(1), leaf(2)))
}

fn observed(owner: Rc<DTree>, callback: &mut dyn FnMut()) -> usize {
    match owner.as_ref() {
        DTree::E => 0,
        DTree::T(_, left, right) => {
            callback();
            let left_depth = observed(left.clone(), callback);
            let right_depth = observed(right.clone(), callback);
            1 + left_depth.max(right_depth)
        }
    }
}

fn consuming(owner: Rc<DTree>, callback: &mut dyn FnMut()) -> usize {
    match Rc::unwrap_or_clone(owner) {
        DTree::E => 0,
        DTree::T(_, left, right) => {
            callback();
            let left_depth = consuming(left, callback);
            let right_depth = consuming(right, callback);
            1 + left_depth.max(right_depth)
        }
    }
}

fn main() {
    let mut logs = Vec::new();
    let mut live_root_at_callback = Vec::new();
    for implementation in [observed, consuming] {
        let log = Rc::new(RefCell::new(Vec::new()));
        let owner = fixture(&log);
        let weak: Weak<DTree> = Rc::downgrade(&owner);
        let mut seen = Vec::new();
        assert_eq!(implementation(owner, &mut || seen.push(weak.upgrade().is_some())), 2);
        assert!(weak.upgrade().is_none());
        logs.push(log.borrow().clone());
        live_root_at_callback.push(seen[0]);
    }
    assert_eq!(live_root_at_callback, [true, false]);
    assert_ne!(logs[0], logs[1], "Opaque payload destructors expose a changed order");
    println!("Scope counterexamples: original/consuming drop logs {logs:?}; live root at first callback {live_root_at_callback:?}");

    std::panic::set_hook(Box::new(|_| {}));
    let mut unwind_logs = Vec::new();
    for implementation in [observed, consuming] {
        let log = Rc::new(RefCell::new(Vec::new()));
        let owner = fixture(&log);
        let weak = Rc::downgrade(&owner);
        let failed = std::panic::catch_unwind(std::panic::AssertUnwindSafe(|| {
            implementation(owner, &mut || panic!("deliberate callback panic"))
        }));
        assert!(failed.is_err());
        assert!(weak.upgrade().is_none());
        let dropped = log.borrow().clone();
        let mut sorted = dropped.clone();
        sorted.sort();
        assert_eq!(sorted, [0, 1, 2], "Unwinding still releases every payload once");
        unwind_logs.push(dropped);
    }
    println!("Panic scope: all payloads released exactly once; original/consuming unwind drop logs {unwind_logs:?}");
}

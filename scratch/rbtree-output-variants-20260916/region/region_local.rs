// Experimental variant of region_alloc.rs under a CLOSED SINGLE-THREAD contract.
// No thread, signal callback, or reentrant allocator access may touch this state.
// All allocations made while active must be dropped before Region::drop.
// This is not a generally thread-safe global allocator library.
use std::alloc::{GlobalAlloc, Layout};
use std::cell::UnsafeCell;
#[cfg(upstream_mimalloc)]
use mimalloc::MiMalloc as Upstream;
#[cfg(not(upstream_mimalloc))]
use std::alloc::System as Upstream;
static UPSTREAM: Upstream = Upstream;

struct State { base: *mut u8, capacity: usize, offset: usize, fallbacks: usize }
struct LocalState(UnsafeCell<State>);
// SAFETY CONTRACT: this experimental executable is single-threaded, including
// allocator calls. The compiler has not proved that property; harness supplies it.
unsafe impl Sync for LocalState {}
static STATE: LocalState = LocalState(UnsafeCell::new(State {
    base: std::ptr::null_mut(), capacity: 0, offset: 0, fallbacks: 0,
}));

struct RegionAllocator;
#[global_allocator]
static ALLOCATOR: RegionAllocator = RegionAllocator;

unsafe impl GlobalAlloc for RegionAllocator {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        let state = &mut *STATE.0.get();
        if !state.base.is_null() {
            if let Some(address) = (state.base as usize).checked_add(state.offset)
                .and_then(|p| p.checked_add(layout.align() - 1)) {
                let aligned = (address & !(layout.align() - 1)) - state.base as usize;
                if let Some(end) = aligned.checked_add(layout.size().max(1)) {
                    if end <= state.capacity {
                        state.offset = end;
                        return state.base.add(aligned);
                    }
                }
            }
            state.fallbacks += 1;
        }
        UPSTREAM.alloc(layout)
    }
    unsafe fn dealloc(&self, pointer: *mut u8, layout: Layout) {
        let state = &*STATE.0.get();
        let base = state.base as usize;
        let p = pointer as usize;
        if base != 0 && p >= base && p - base < state.capacity { return; }
        UPSTREAM.dealloc(pointer, layout);
    }
}

pub struct Region { pointer: *mut u8, layout: Layout }
impl Region {
    pub unsafe fn new(capacity: usize) -> Self {
        assert!((*STATE.0.get()).base.is_null());
        let layout = Layout::from_size_align(capacity.max(64), 64).unwrap();
        let pointer = UPSTREAM.alloc(layout);
        if pointer.is_null() { std::alloc::handle_alloc_error(layout); }
        *STATE.0.get() = State { base: pointer, capacity: layout.size(), offset: 0, fallbacks: 0 };
        Self { pointer, layout }
    }
    pub fn used(&self) -> usize { unsafe { (*STATE.0.get()).offset } }
    pub fn fallbacks(&self) -> usize { unsafe { (*STATE.0.get()).fallbacks } }
}
impl Drop for Region {
    fn drop(&mut self) {
        unsafe {
            (*STATE.0.get()).base = std::ptr::null_mut();
            UPSTREAM.dealloc(self.pointer, self.layout);
        }
    }
}

pub fn run(n: i64) -> i64 {
    let region = unsafe { Region::new((n.max(0) as usize + 1).checked_mul(64).unwrap()) };
    let tree = crate::Test_RBTree_buildTree(std::hint::black_box(n), std::rc::Rc::new(crate::Tree::E));
    let result = crate::Test_RBTree_depth(std::hint::black_box(tree));
    drop(region);
    result
}

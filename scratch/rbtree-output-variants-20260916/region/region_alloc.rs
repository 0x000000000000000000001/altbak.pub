// Scratch experiment: whole closed, single-threaded kernel executes in this region.
// All allocations made while active must be dropped before Region::drop.
// Rc layout, reference counting, and recursive destruction remain unchanged.
use std::alloc::{GlobalAlloc, Layout, System};
use std::sync::atomic::{AtomicPtr, AtomicUsize, Ordering};

struct RegionAllocator;
#[global_allocator]
static ALLOCATOR: RegionAllocator = RegionAllocator;
static BASE: AtomicPtr<u8> = AtomicPtr::new(std::ptr::null_mut());
static CAPACITY: AtomicUsize = AtomicUsize::new(0);
static OFFSET: AtomicUsize = AtomicUsize::new(0);
static FALLBACKS: AtomicUsize = AtomicUsize::new(0);

unsafe impl GlobalAlloc for RegionAllocator {
    unsafe fn alloc(&self, layout: Layout) -> *mut u8 {
        let base = BASE.load(Ordering::Acquire);
        if !base.is_null() {
            let capacity = CAPACITY.load(Ordering::Relaxed);
            let mut offset = OFFSET.load(Ordering::Relaxed);
            loop {
                let Some(address) = (base as usize).checked_add(offset)
                    .and_then(|p| p.checked_add(layout.align() - 1)) else { break };
                let aligned = (address & !(layout.align() - 1)) - base as usize;
                let Some(end) = aligned.checked_add(layout.size().max(1)) else { break };
                if end > capacity { break; }
                match OFFSET.compare_exchange_weak(offset, end, Ordering::Relaxed, Ordering::Relaxed) {
                    Ok(_) => return base.add(aligned),
                    Err(next) => offset = next,
                }
            }
            FALLBACKS.fetch_add(1, Ordering::Relaxed);
        }
        System.alloc(layout)
    }

    unsafe fn dealloc(&self, pointer: *mut u8, layout: Layout) {
        let base = BASE.load(Ordering::Acquire) as usize;
        let p = pointer as usize;
        if base != 0 && p >= base && p - base < CAPACITY.load(Ordering::Relaxed) { return; }
        System.dealloc(pointer, layout);
    }
}

pub struct Region { pointer: *mut u8, layout: Layout }
impl Region {
    // Caller proves no concurrent/escaping allocations in the dynamic region.
    pub unsafe fn new(capacity: usize) -> Self {
        assert!(BASE.load(Ordering::Acquire).is_null());
        let layout = Layout::from_size_align(capacity.max(64), 64).unwrap();
        let pointer = System.alloc(layout);
        if pointer.is_null() { std::alloc::handle_alloc_error(layout); }
        CAPACITY.store(layout.size(), Ordering::Relaxed);
        OFFSET.store(0, Ordering::Relaxed);
        FALLBACKS.store(0, Ordering::Relaxed);
        BASE.store(pointer, Ordering::Release);
        Self { pointer, layout }
    }
    pub fn used(&self) -> usize { OFFSET.load(Ordering::Relaxed) }
    pub fn fallbacks(&self) -> usize { FALLBACKS.load(Ordering::Relaxed) }
}
impl Drop for Region {
    fn drop(&mut self) {
        BASE.store(std::ptr::null_mut(), Ordering::Release);
        unsafe { System.dealloc(self.pointer, self.layout); }
    }
}

pub fn run(n: i64) -> i64 {
    // Capacity derives from known construction cardinality. Overflow remains correct:
    // excess allocations go through System. The complete block cost is in this call.
    let region = unsafe { Region::new((n.max(0) as usize + 1).checked_mul(64).unwrap()) };
    let tree = crate::Test_RBTree_buildTree(std::hint::black_box(n), std::rc::Rc::new(crate::Tree::E));
    let result = crate::Test_RBTree_depth(std::hint::black_box(tree));
    drop(region); // All node destructors and Rc drops ran before the block is freed.
    result
}

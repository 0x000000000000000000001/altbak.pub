// Counting uses a one-word wrapper around native Rc; `--cfg untracked` emits
// precisely the same IR with std::rc::Rc and no instrumentation.
use std::cell::Cell;
thread_local! {
    static EVENTS: Cell<[u64; 9]> = const { Cell::new([0; 9]) };
    static LIVE: Cell<u64> = const { Cell::new(0) };
    static PEAK: Cell<u64> = const { Cell::new(0) };
}
fn event(index: usize) {
    if !cfg!(untracked) {
        EVENTS.with(|s| { let mut v=s.get(); v[index]+=1; s.set(v); });
        if index==0 { LIVE.with(|s| { let next=s.get()+1;s.set(next);PEAK.with(|p|p.set(p.get().max(next))); }); }
        if index==2 || index==4 { LIVE.with(|s|s.set(s.get()-1)); }
    }
}
fn counts() -> [u64;9] { EVENTS.with(Cell::get) }
fn peak() -> u64 { PEAK.with(Cell::get) }
fn clear_counts() { EVENTS.with(|s|s.set([0;9]));LIVE.with(|s|s.set(0));PEAK.with(|s|s.set(0)); }
fn balance() {
    let c=counts();
    assert_eq!(c[0]+c[1], c[2]+c[3]+c[4], "owner balance");
    assert_eq!(c[0], c[2]+c[4], "allocated payloads all released");
    assert_eq!(LIVE.with(Cell::get), 0);
}
#[cfg(untracked)] use std::rc::Rc;
#[cfg(not(untracked))] use counted::Rc;
#[cfg(not(untracked))]
mod counted {
    use super::event;
    use std::ops::Deref;
    use std::rc::Rc as Native;
    pub struct Rc<T>(Option<Native<T>>);
    impl<T> Rc<T> {
        pub fn new(value:T)->Self { event(0); Self(Some(Native::new(value))) }
        pub fn as_ref(&self)->&T { self.0.as_ref().unwrap().as_ref() }
        pub fn get_mut(owner:&mut Self)->Option<&mut T> { Native::get_mut(owner.0.as_mut().unwrap()) }
        pub fn downgrade(owner:&Self)->std::rc::Weak<T> { Native::downgrade(owner.0.as_ref().unwrap()) }
        pub fn try_unwrap(mut owner:Self)->Result<T,Self> {
            match Native::try_unwrap(owner.0.take().unwrap()) {
                Ok(value)=>{event(4);Ok(value)},
                Err(ptr)=>{event(5);Err(Self(Some(ptr)))}
            }
        }
    }
    impl<T> Clone for Rc<T> { fn clone(&self)->Self { event(1);Self(Some(self.0.as_ref().unwrap().clone())) } }
    impl<T> Deref for Rc<T> { type Target=T; fn deref(&self)->&T { self.as_ref() } }
    impl<T> Drop for Rc<T> { fn drop(&mut self) {
        if let Some(ptr)=&self.0 { event(if Native::strong_count(ptr)==1 {2} else {3}); }
    }}
}

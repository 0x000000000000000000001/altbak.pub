use std::rc::Rc;
use std::alloc::{GlobalAlloc,Layout,System};
use std::sync::atomic::{AtomicUsize,Ordering};
static ALLOCS:AtomicUsize=AtomicUsize::new(0);
static FREES:AtomicUsize=AtomicUsize::new(0);
struct Counter;
unsafe impl GlobalAlloc for Counter {
 unsafe fn alloc(&self,l:Layout)->*mut u8 {ALLOCS.fetch_add(1,Ordering::Relaxed);System.alloc(l)}
 unsafe fn dealloc(&self,p:*mut u8,l:Layout){FREES.fetch_add(1,Ordering::Relaxed);System.dealloc(p,l)}
}
#[global_allocator] static GLOBAL:Counter=Counter;
#[derive(Clone)] enum Tree {E,N(Rc<Tree>,i64,Rc<Tree>)}
fn node(tree:&Tree)->(&Rc<Tree>,i64,&Rc<Tree>){if let Tree::N(l,k,r)=tree{(l,*k,r)}else{panic!()}}
fn leaf(k:i64)->Rc<Tree>{Rc::new(Tree::N(Rc::new(Tree::E),k,Rc::new(Tree::E)))}
fn rotate(mut left:Rc<Tree>,k:i64,right:Rc<Tree>,reuse:bool)->Rc<Tree>{
 if reuse && Rc::get_mut(&mut left).is_some(){
  let Tree::N(a,x,b)=std::mem::replace(Rc::get_mut(&mut left).unwrap(),Tree::E) else {panic!()};
  *Rc::get_mut(&mut left).unwrap()=Tree::N(b,k,right);
  Rc::new(Tree::N(a,x,left))
 }else{
  let (a,x,b)=node(&left);
  Rc::new(Tree::N(a.clone(),x,Rc::new(Tree::N(b.clone(),k,right))))
 }
}
fn main(){
 for (reuse,shared,expected) in [(false,false,2),(true,false,1),(true,true,2)]{
  let starts=(ALLOCS.load(Ordering::Relaxed),FREES.load(Ordering::Relaxed));
  let left=Rc::new(Tree::N(leaf(1),2,leaf(3)));
  let old=shared.then(||left.clone());
  let address=Rc::as_ptr(&left);
  let right=leaf(5);
  let before=ALLOCS.load(Ordering::Relaxed);
  let changed=rotate(left,4,right,reuse);
  let count=ALLOCS.load(Ordering::Relaxed)-before;
  assert_eq!(count,expected);
  let (a,x,b)=node(&changed);assert_eq!((node(a).1,x,node(b).1,node(node(b).0).1,node(node(b).2).1),(1,2,4,3,5));
  if reuse && !shared {assert_eq!(Rc::as_ptr(b),address);}
  if let Some(old)=old{assert_eq!((node(node(&old).0).1,node(&old).1,node(node(&old).2).1),(1,2,3));}
  drop(changed);
  assert_eq!(ALLOCS.load(Ordering::Relaxed)-starts.0,FREES.load(Ordering::Relaxed)-starts.1);
  println!("reuse={reuse}, shared={shared}: {count} allocations; values, retained version, cell address and frees checked");
 }
}

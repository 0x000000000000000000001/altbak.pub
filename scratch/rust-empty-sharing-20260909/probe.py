from pathlib import Path
import re,subprocess,json
HERE=Path(__file__).resolve().parent
ROOT=HERE.parents[1]
GEN=ROOT/'run/bak/rust/output/purust_output'
source=(HERE/'RBTree-before.rs').read_text()
starts=list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(',source,re.M))
functions={m[1]:source[m.start():starts[i+1].start() if i+1<len(starts) else len(source)] for i,m in enumerate(starts)}
names=['max','makeBlack','depth','balance','ins','insert','buildTree']
names += [name.removeprefix('Test_RBTree_') for name in functions if '__purust_rebuild_' in name or name.endswith('__purust_reuse')]
kernel=source[source.index('#[derive(Clone'):starts[0].start()]+'\n'.join(functions['Test_RBTree_'+n] for n in names)
empty='std::rc::Rc::new(crate::Tree::E)'
lines=[line for line in functions['Test_RBTree_ins'].splitlines() if line.count(empty)==2]
assert len(lines)==1
line=lines[0]
expr=line.strip(); rewritten='{ let empty = '+empty+'; '+expr.replace(empty,'empty.clone()',1).replace(empty,'empty',1)+' }'
shared=kernel.replace(line,rewritten)
assert shared!=kernel
harness=r'''
use std::alloc::{GlobalAlloc,Layout};
use std::sync::atomic::{AtomicUsize,Ordering};
static A:AtomicUsize=AtomicUsize::new(0);
static F:AtomicUsize=AtomicUsize::new(0);
static EMPTY:AtomicUsize=AtomicUsize::new(0);
static NODE:AtomicUsize=AtomicUsize::new(0);
struct Counting;
unsafe impl GlobalAlloc for Counting {
 unsafe fn alloc(&self,l:Layout)->*mut u8 {A.fetch_add(1,Ordering::Relaxed);mimalloc::MiMalloc.alloc(l)}
 unsafe fn dealloc(&self,p:*mut u8,l:Layout){F.fetch_add(1,Ordering::Relaxed);mimalloc::MiMalloc.dealloc(p,l)}
}
#[global_allocator] static GLOBAL:Counting=Counting;
fn counted(value:Tree)->std::rc::Rc<Tree>{
 match &value {Tree::E=>{EMPTY.fetch_add(1,Ordering::Relaxed);},Tree::T(..)=>{NODE.fetch_add(1,Ordering::Relaxed);}}
 std::rc::Rc::new(value)
}
fn main(){
 let a=A.load(Ordering::Relaxed);let f=F.load(Ordering::Relaxed);
 let seed=counted(Tree::E);
 EMPTY.store(0,Ordering::Relaxed);NODE.store(0,Ordering::Relaxed);
 let tree=Test_RBTree_ins(7,seed);
 let (single_e,single_n)=(EMPTY.load(Ordering::Relaxed),NODE.load(Ordering::Relaxed));
 let Tree::T(_,l,k,r)=tree.as_ref() else{panic!()};assert_eq!(*k,7);assert!(matches!(l.as_ref(),Tree::E)&&matches!(r.as_ref(),Tree::E));
 drop(tree);
 assert_eq!(A.load(Ordering::Relaxed)-a,F.load(Ordering::Relaxed)-f);
 EMPTY.store(0,Ordering::Relaxed);NODE.store(0,Ordering::Relaxed);
 let a=A.load(Ordering::Relaxed);let f=F.load(Ordering::Relaxed);
 let tree=Test_RBTree_buildTree(100000,counted(Tree::E));
 assert_eq!(Test_RBTree_depth(tree),22);
 let (e,n)=(EMPTY.load(Ordering::Relaxed),NODE.load(Ordering::Relaxed));
 assert_eq!(A.load(Ordering::Relaxed)-a,F.load(Ordering::Relaxed)-f);
 println!("{single_e} {single_n} {e} {n}");
}
'''
deps=GEN/'target/release/deps';mimalloc=list(deps.glob('libmimalloc-*.rlib'));assert len(mimalloc)==1
results={}
for side,code in [('before',kernel),('shared',shared)]:
 build=HERE/'build';build.mkdir(exist_ok=True)
 path=build/f'probe-{side}.rs';binary=build/f'probe-{side}'
 path.write_text('#![allow(warnings)]\n'+code.replace('std::rc::Rc::new(', 'counted(')+harness)
 subprocess.run(['rustc','--edition=2021','-C','opt-level=1','--extern',f'mimalloc={mimalloc[0]}','-L',f'dependency={deps}',str(path),'-o',str(binary)],check=True)
 values=list(map(int,subprocess.check_output([str(binary)],text=True).split()))
 results[side]=dict(zip(['single_empty','single_node','full_empty','full_node'],values))
 print(side,results[side],flush=True)
(HERE/'probe.json').write_text(json.dumps(results,indent=2)+'\n')

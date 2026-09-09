from pathlib import Path
import json,subprocess,statistics,hashlib
HERE=Path(__file__).resolve().parent
ROOT=HERE.parents[1]
GEN=ROOT/'run/bak/rust/output/purust_output'
DEPS=GEN/'target/release/deps'
source=(HERE/'LazyEvaluation-before.rs').read_text()
kernel=source[source.index('pub fn Test_LazyEvaluation_buildThunks('):source.index('pub fn Test_LazyEvaluation_act(')]
needle='(purs_local_1.clone())('
assert kernel.count(needle)==1
variants={'before':kernel,'borrowed':kernel.replace(needle,'(purs_local_1)(')}
extern=[]
for name in ['purust_core','Purs_Data_Unit','mimalloc']:
 paths=list(DEPS.glob('lib'+name+'-*.rlib'));assert len(paths)==1
 extern+=['--extern',f'{name}={paths[0]}']
prefix='#![allow(warnings)]\nuse purust_core::*;\nuse Purs_Data_Unit::Data_Unit_unit;\n'
checks='''
fn checks() {
 for n in [0,1,2,17,1000] {
  let calls=std::rc::Rc::new(std::cell::Cell::new(0)); let observed=calls.clone();
  let f=purust_core::Func1::Shared(std::rc::Rc::new(move |_| {observed.set(observed.get()+1);-7}));
  for i in 1..=3 {assert_eq!(Test_LazyEvaluation_buildThunks(n,f.clone(),()),n-7);assert_eq!(calls.get(),i);}
 }
 for n in [0,1,2,10] {assert_eq!(Test_LazyEvaluation_runManyTimes(n,17),1000*n+17);}
}
'''
old=(HERE/'count.py').read_text()
allocator=old.split("allocator='''",1)[1].split("'''",1)[0]
count_main='''
fn main(){
 checks();
 let a=ALLOCS.load(Ordering::Relaxed);let f=FREES.load(Ordering::Relaxed);let b=BYTES.load(Ordering::Relaxed);
 let result=Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0);
 let a=ALLOCS.load(Ordering::Relaxed)-a;let f=FREES.load(Ordering::Relaxed)-f;let b=BYTES.load(Ordering::Relaxed)-b;
 assert_eq!(result,1000000);assert_eq!(a,f);println!("{a} {f} {b}");
}
'''
time_main='''
#[global_allocator] static GLOBAL:mimalloc::MiMalloc=mimalloc::MiMalloc;
fn main(){
 checks();
 for i in 0..16 {
  let start=std::time::Instant::now();
  assert_eq!(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0),1000000);
  let ns=start.elapsed().as_nanos();if i>0 {println!("{ns}");}
 }
}
'''
allocs={}
BUILD=HERE/'build';BUILD.mkdir(exist_ok=True)
for side,k in variants.items():
 (BUILD/f'{side}-kernel.rs').write_text(k)
 for mode,tail in [('count',allocator+checks+count_main),('time',checks+time_main)]:
  path=BUILD/f'{side}-{mode}.rs';binary=BUILD/f'{side}-{mode}';path.write_text(prefix+k+tail)
  subprocess.run(['rustc','--edition=2021','-C','opt-level=1',*extern,'-L',f'dependency={DEPS}',str(path),'-o',str(binary)],check=True,capture_output=True)
  if mode=='count':allocs[side]=list(map(int,subprocess.check_output([str(binary)],text=True).split()))
runs={s:[] for s in variants}
for i in range(3):
 for side in (['before','borrowed'] if i%2==0 else ['borrowed','before']):
  nums=list(map(int,subprocess.check_output([str(BUILD/f'{side}-time')],text=True).split()));assert len(nums)==15
  runs[side].append(nums);print(i+1,side,statistics.median(nums)/1e6,'ms',flush=True)
result={'source_sha256':hashlib.sha256(source.encode()).hexdigest(),'method':'Current generated worktree kernel, exact one-expression change; O1/mimalloc; 3 alternating process pairs; 15 measurements after warmup per process; repeated calls, shared inputs and results checked; allocations counted separately.','median_ms':{s:statistics.median(sum(r,[]))/1e6 for s,r in runs.items()},'runs_ns':runs,'counts_order':['allocations','deallocations','bytes_requested'],'allocations':allocs}
(HERE/'prototype-results.json').write_text(json.dumps(result,indent=2)+'\n');print(result['median_ms'],allocs,flush=True)

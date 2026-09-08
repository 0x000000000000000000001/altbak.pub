from pathlib import Path
import hashlib, json, os, re, statistics, subprocess

HERE=Path(__file__).resolve().parent
GEN=HERE.parents[1]/'output/purust_output'
source=(GEN/'Purs_Test_LazyEvaluation/src/lib.rs').read_text()
start=source.index('pub fn Test_LazyEvaluation_buildThunks(')
split=source.index('pub fn Test_LazyEvaluation_runManyTimes(')
end=source.index('pub fn Test_LazyEvaluation_act(')
original=source[start:end]
replacement='''pub fn Test_LazyEvaluation_buildThunks(mut n: i64, mut acc: purust_core::Func1<crate::UnknownType, i64>, unit: crate::UnknownType) -> i64 {
    while n != 0 {
        let previous = acc;
        acc = purust_core::Func1::Shared(std::rc::Rc::new(move |_| previous(Data_Unit_unit()) + 1));
        n -= 1;
    }
    acc(unit)
}
'''+source[split:end]

harness='''
fn main() {
    eprintln!("layout: Value={} Record_a={}", std::mem::size_of::<Value>(), std::mem::size_of::<Record_a>());
    for n in [0,1,10,1000] {
        let calls = std::rc::Rc::new(std::cell::Cell::new(0));
        let counter = calls.clone();
        let f = purust_core::Func1::Shared(std::rc::Rc::new(move |_| { counter.set(counter.get()+1); 7 }));
        assert_eq!(Test_LazyEvaluation_buildThunks(n, f, Data_Unit_unit()), n+7);
        assert_eq!(calls.get(), 1);
    }
    assert_eq!(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0),1000000);
    for _ in 0..3 {
        let start = std::time::Instant::now();
        let value = std::hint::black_box(Test_LazyEvaluation_runManyTimes(std::hint::black_box(1000),0));
        let us = start.elapsed().as_secs_f64()*1e6;
        assert_eq!(value,1000000);
        println!("{:.3}",us);
    }
}
'''
deps=GEN/'target/release/deps'
externs=[]
for name in ['mimalloc','purust_core','Purs_Data_Unit']:
    files=list(deps.glob('lib'+name+'-*.rlib'))
    assert len(files)==1,files
    externs+=['--extern',f'{name}={files[0]}']
prefix='''#![allow(warnings)]
use purust_core::*;
use Purs_Data_Unit::Data_Unit_unit;
#[global_allocator] static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;
'''
cached_unit='''
thread_local! { static AUDIT_UNIT: UnknownType = Purs_Data_Unit::Data_Unit_unit(); }
fn Data_Unit_unit() -> UnknownType { AUDIT_UNIT.with(Clone::clone) }
'''
variants={'lazy-generated':original,'lazy-one-thunk':replacement,
          'lazy-cached-unit':cached_unit+original}
selected=os.environ.get('RUST_AUDIT_LAZY_VARIANTS')
if selected:
    variants={name:variants[name] for name in selected.split(',')}
for name,body in variants.items():
    path=HERE/(name+'.rs')
    variant_prefix=prefix.replace('use Purs_Data_Unit::Data_Unit_unit;','') if name=='lazy-cached-unit' else prefix
    path.write_text(variant_prefix+body+harness)
    subprocess.run(['rustc','--edition=2021','-C','opt-level=1','-C','debuginfo=1',*externs,'-L',f'dependency={deps}',str(path),'-o',str(HERE/name)],check=True)
samples={name:[] for name in variants}
for trial in range(3):
    order=list(variants) if trial%2==0 else list(reversed(variants))
    for name in order:
        print('MEASURE',trial,name,flush=True)
        raw=subprocess.check_output([str(HERE/name)],text=True)
        samples[name]+=[float(x) for x in raw.splitlines()]
result={'generated_sha256':hashlib.sha256(source.encode()).hexdigest(),
    'method':'opt-level=1, existing compiled runtime, mimalloc; 3 interleaved processes, 1 warm-up + 3 timed runs each; 1000 chains of depth1000; all values/assertions checked',
    'results':{name:{'n':len(xs),'min_us':min(xs),'median_us':statistics.median(xs),'max_us':max(xs),'samples_us':xs} for name,xs in samples.items()}}
(HERE/os.environ.get('RUST_AUDIT_LAZY_RESULTS','results-lazy.json')).write_text(json.dumps(result,indent=2)+'\n')
print(json.dumps(result,indent=2))

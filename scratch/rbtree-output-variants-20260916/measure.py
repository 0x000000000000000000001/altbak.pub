#!/usr/bin/env python3
"""Matched kernel timings; all builds finish before any samples are taken."""
from pathlib import Path
import argparse
import hashlib
import json
import platform
import random
import statistics
import subprocess
import time

ROOT=Path(__file__).resolve().parent
FLAGS=['--edition=2021','-C','opt-level=3','-C','codegen-units=1']
VARIANTS={
 'generated': ('baseline.rs','rc',[]),
 'guard_simple': ('decisions/kernel-simple-guard.rs','rc',[]),
 'guard_fused': ('decisions/kernel-fused-guard.rs','rc',[]),
 'region': ('baseline.rs','region',[]),
 'region_fused': ('decisions/kernel-fused-guard.rs','region',[]),
 'arena_u32': ('arena/kernel.rs','own',[]),
 'arena_grow': ('arena/kernel.rs','own',['--cfg','arena_grow']),
 'arena_usize': ('arena/kernel.rs','own',['--cfg','arena_usize']),
 'guard_unique': ('decisions/kernel-fused-unique.rs','rc',[]),
 'region_local': ('baseline.rs','region_local',[]),
 'region_local_fused': ('decisions/kernel-fused-guard.rs','region_local',[]),
 'region_local_unique': ('decisions/kernel-fused-unique.rs','region_local',[]),
 'arena_unchecked': ('arena/kernel-unchecked.rs','own',[]),
 'arena_unchecked_inline': ('arena/kernel-unchecked-inline.rs','own',[]),
 'arena_directional': ('arena/kernel-directional.rs','own',[]),
 'arena_unchecked_directional_inline': ('arena/kernel-unchecked-directional-inline.rs','own',[]),
}
RC_RUN='''
#[inline(never)]
fn run(n: i64) -> i64 {
 let tree=Test_RBTree_buildTree(std::hint::black_box(n),std::rc::Rc::new(Tree::E));
 Test_RBTree_depth(std::hint::black_box(tree))
}
'''
MAIN='''
fn main() {
 let args: Vec<String>=std::env::args().collect();
 let n: i64=args[1].parse().unwrap();
 let samples: usize=args[2].parse().unwrap();
 let expected: i64=args[3].parse().unwrap();
 for sample in 0..samples+3 {
  let start=std::time::Instant::now();
  let result=run(std::hint::black_box(n));
  std::hint::black_box(result);
  let ns=start.elapsed().as_nanos();
  assert_eq!(result,expected);
  if sample>=3 { println!("{}",ns); }
 }
}
'''
def sha(path): return hashlib.sha256(path.read_bytes()).hexdigest()
def main():
 p=argparse.ArgumentParser()
 p.add_argument('--compile-only',action='store_true')
 p.add_argument('--skip-compile',action='store_true')
 p.add_argument('--variants',default=','.join(VARIANTS))
 p.add_argument('--rounds',type=int,default=21)
 p.add_argument('--samples',type=int,default=10)
 p.add_argument('--n',type=int,default=100000)
 p.add_argument('--expected',type=int,default=22)
 p.add_argument('--out',default='timing')
 p.add_argument('--allocator',choices=['system','mimalloc'],default='system')
 a=p.parse_args()
 names=a.variants.split(',')
 assert 'generated' in names
 out=ROOT/a.out; out.mkdir(exist_ok=True)
 link_flags=[]; allocator_info={'name':a.allocator}
 if a.allocator=='mimalloc':
  deps=ROOT.parents[1]/'run/bak/rust/modes/pure/target/release/deps'
  libraries=list(deps.glob('libmimalloc-*.rlib')); assert len(libraries)==1
  library=libraries[0]
  link_flags=['--extern','mimalloc='+str(library),'-L','dependency='+str(deps)]
  allocator_info.update({'library':str(library),'library_sha256':sha(library)})
 manifest={'rustc':subprocess.check_output(['rustc','-Vv'],text=True),
  'platform':platform.platform(),'flags':FLAGS,'n':a.n,'rounds':a.rounds,'samples':a.samples,
  'warmups':3,'seed':16092026,'scope':'allocation + descending build + depth + full destruction/release',
  'source_manifest':json.loads((ROOT/'source-manifest.json').read_text()),'allocator':allocator_info,'variants':{}}
 for name in names:
  kernel,kind,cfg=VARIANTS[name]; kernel=ROOT/kernel
  source=out/(name+'.rs'); binary=out/name
  text='#![allow(warnings)]\n#![recursion_limit="512"]\ninclude!('+json.dumps(str(kernel))+');\n'
  extras={}
  if a.allocator=='mimalloc' and not kind.startswith('region'):
   text+='#[global_allocator]\nstatic GLOBAL: mimalloc::MiMalloc=mimalloc::MiMalloc;\n'
  if kind=='rc': text+=RC_RUN
  elif kind.startswith('region'):
   alloc=ROOT/('region/region_local.rs' if kind=='region_local' else 'region/region_alloc.rs')
   if a.allocator=='mimalloc':
    if kind=='region_local': cfg=[*cfg,'--cfg','upstream_mimalloc']
    else:
     original=alloc.read_text()
     replacement=original.replace('use std::alloc::{GlobalAlloc, Layout, System};',
      'use std::alloc::{GlobalAlloc, Layout};\nuse mimalloc::MiMalloc as System;')
     assert original!=replacement
     alloc=out/'region_mimalloc.rs'
     alloc.write_text(replacement)
   text+='#[path='+json.dumps(str(alloc))+'] mod region_alloc;\n'
   text+='#[inline(never)] fn run(n:i64)->i64 { region_alloc::run(n) }\n'
   extras['allocator_sha256']=sha(alloc)
  text+=MAIN
  if not a.skip_compile:
   source.write_text(text)
   result=subprocess.run(['rustc',*FLAGS,*cfg,*link_flags,str(source),'-o',str(binary)],capture_output=True,text=True,timeout=60)
   (out/(name+'.compile.log')).write_text(result.stdout+result.stderr)
   result.check_returncode()
  else: assert source.read_text()==text, 'harness changed; recompile'
  manifest['variants'][name]={'kernel':str(kernel),'kernel_sha256':sha(kernel),
   'harness_sha256':sha(source),'binary_sha256':sha(binary),'cfg':cfg,**extras}
 (out/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')
 if a.compile_only:
  print('Compiled '+', '.join(names)); return
 rng=random.Random(16092026); rounds=[]
 for idx in range(a.rounds):
  order=names.copy(); rng.shuffle(order); samples={}
  for name in order:
   result=subprocess.run([str(out/name),str(a.n),str(a.samples),str(a.expected)],capture_output=True,text=True,check=True,timeout=30)
   us=[int(s)/1000 for s in result.stdout.splitlines()]; assert len(us)==a.samples
   samples[name]={'us':us,'min_us':min(us),'median_us':statistics.median(us)}
  rounds.append({'index':idx,'order':order,'samples':samples})
  print('round',idx+1,{k:round(v['min_us'],2) for k,v in samples.items()},flush=True)
 summary={}
 for name in names:
  mins=[r['samples'][name]['min_us'] for r in rounds]
  gains=[100*(1-r['samples'][name]['min_us']/r['samples']['generated']['min_us']) for r in rounds]
  summary[name]={'median_best10_us':statistics.median(mins),
   'median_median10_us':statistics.median(r['samples'][name]['median_us'] for r in rounds),
   'median_paired_gain_pct':statistics.median(gains),'paired_gain_range_pct':[min(gains),max(gains)],
   'wins':sum(g>0 for g in gains)}
 result={'manifest':manifest,'rounds':rounds,'summary':summary,'finished_unix':time.time()}
 (out/'results.json').write_text(json.dumps(result,indent=2)+'\n')
 print(json.dumps(summary,indent=2))
if __name__=='__main__': main()

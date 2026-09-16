from pathlib import Path
import json,subprocess,hashlib
here=Path(__file__).resolve().parent
old=Path('/Users/0x1/Documents/htdocs/altbak.pub/scratch/purust-tast-perf-proof-20260916/unique')
fused=(here/'kernel-fused-guard.rs').read_text()
previous=(old/'kernel_unique_sites.rs').read_text()
helper=previous[previous.index('// EXPERIMENTAL specialization only.'):]
assert helper.count('fn assumed_unique_mut(')==1
count=fused.count('std::rc::Rc::get_mut(')
variant=fused.replace('std::rc::Rc::get_mut(', 'assumed_unique_mut(')+'\n'+helper
(here/'kernel-fused-unique.rs').write_text(variant)
check=(old/'check_sites.rs').read_text().replace('kernel_unique_sites.rs','kernel-fused-unique.rs')
(here/'check-fused-unique.rs').write_text(check)
subprocess.run(['rustc','--edition=2021','-C','opt-level=2','--cfg','verify_unique',str(here/'check-fused-unique.rs'),'-o',str(here/'check-fused-unique')],check=True)
result=subprocess.check_output([str(here/'check-fused-unique')],text=True,timeout=60)
(here/'unique-check.txt').write_text(result)
(here/'unique-manifest.json').write_text(json.dumps({'fused_source_sha256':hashlib.sha256(fused.encode()).hexdigest(),'replaced_get_mut_sites':count,'helper_source':str(old/'kernel_unique_sites.rs'),'checks':result,'contract':'Experimental specialization requiring exclusive provenance at every mutable site, including temporarily E payloads; no Weak aliases. Not a general persistent API. cfg verify_unique checks strong=1/weak=0 on every mutable call. Timing compiles without cfg.','timings':'parent only'},indent=2)+'\n')
print(result)

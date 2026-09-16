from pathlib import Path
import shutil, hashlib, difflib, json, os, subprocess, time
base=Path(__file__).resolve().parent
source=base.parent/'compact-guards-integration-20260916/after'
dest=base/'full-scalar'
shutil.copytree(source,dest,ignore=shutil.ignore_patterns('target'),dirs_exist_ok=True)
relative=Path('output/purust_output/Purs_Test_Records/src/lib.rs')
p=dest/relative
original=p.read_text()
start=original.index('pub fn Test_Records_updateRec')
end=original.index('pub fn Test_Records_initial',start)
expected=(base/'kernel-original.rs').read_text()
assert original[start:end]==expected, 'worker drift against checked probe'
variant=(base/'kernel-scalar.rs').read_text().replace('scalar_replacement','Test_Records_updateRec')
changed=original[:start]+variant+'\n'+original[end:]
p.write_text(changed)
(base/'full-scalar.patch').write_text(''.join(difflib.unified_diff(original.splitlines(True),changed.splitlines(True),fromfile=str(source/relative),tofile=str(p))))
modified=[]
for q in sorted((source/'output/purust_output').rglob('*')):
 if q.is_file() and 'target' not in q.relative_to(source).parts:
  other=dest/q.relative_to(source)
  if hashlib.sha256(q.read_bytes()).digest()!=hashlib.sha256(other.read_bytes()).digest():modified.append(str(q.relative_to(source)))
assert modified==[str(relative)],modified
cmd=['cargo','build','--release','--offline','--manifest-path',str(dest/'output/purust_output/Cargo.toml'),'--target-dir',str(dest/'target')]
env=os.environ.copy();env.update(CARGO_PROFILE_RELEASE_DEBUG='false',CARGO_PROFILE_RELEASE_OPT_LEVEL='3')
manifest={'source':str(source),'target':str(dest),'modified_files':modified,'source_records_sha256':hashlib.sha256(original.encode()).hexdigest(),'variant_records_sha256':hashlib.sha256(changed.encode()).hexdigest(),'command':cmd,'profile':{'CARGO_PROFILE_RELEASE_DEBUG':'false','CARGO_PROFILE_RELEASE_OPT_LEVEL':'3'}}
(base/'full-scalar-manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')
start=time.monotonic()
with (base/'full-scalar-build.log').open('w') as log:
 result=subprocess.run(cmd,env=env,stdout=log,stderr=subprocess.STDOUT,timeout=180)
print('exit',result.returncode,'elapsed_s',time.monotonic()-start)
print('binary',dest/'target/release/purust_output')
raise SystemExit(result.returncode)

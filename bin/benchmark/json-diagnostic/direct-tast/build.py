#!/usr/bin/env python3
"""Preserve an indexed TAST prototype beside a pinned ordinary diagnostic.

Only the combined operation changes. Parse/decode controls keep the public DOM
path, and every timed output retains the original fingerprint validation.
"""
import argparse, hashlib, json, os, shutil, subprocess
from pathlib import Path

def sha(path): return hashlib.sha256(path.read_bytes()).hexdigest()
def replace(path, before, after):
    text=path.read_text()
    if text.count(before)!=1: raise ValueError(f'Expected one replacement in {path}: {before}')
    path.write_text(text.replace(before,after))

parser=argparse.ArgumentParser(description=__doc__)
parser.add_argument('--baseline',type=Path,required=True)
parser.add_argument('--output',type=Path,required=True)
args=parser.parse_args();base=args.baseline.resolve();out=args.output.resolve()
root=Path(__file__).resolve().parent
shutil.copytree(base,out)
ffi=out/'output/purescript/PureScript_Backend_Optimizer_CoreFn_Json_ffi.go'
replace(ffi,'switch v := input.(type) {','switch v := input.(type) {\n\tcase directCursor: return v.native()')
replace(ffi,'func ntObjectOf(raw any) (ntObject, bool) { return gopurs_runtime.ReadJSONObject(raw) }',
    '''func ntObjectOf(raw any) (ntObject, bool) {
    if cursor, ok := raw.(directCursor); ok && cursor.kind() != '{' { return ntObject{}, false }
    return gopurs_runtime.ReadJSONObject(raw)
}''')
replace(ffi,'func cndIsNull(raw any) bool {','func cndIsNull(raw any) bool {\n\tif cursor, ok := raw.(directCursor); ok { return cursor.kind() == \'n\' }')
replace(ffi,'jsonValue := gopurs_runtime.Box(ntNative(raw))','jsonValue := gopurs_runtime.Box(directMaterialize(raw))')
driver=out/'output/purescript/Test_JsonTypedAst_ffi.go'
replace(driver,'case "combined":results[i]=gopurs_runtime.Apply(decode,gopurs_runtime.Apply(parse,gopurs_runtime.Str(f.Contents)))',
    'case "combined":results[i]=directTAST(parse,decode,f.Contents)')
shutil.copy2(root/'index.go',out/'output/purescript/direct_index.go')
env={k:v for k,v in os.environ.items() if not k.startswith(('GOPURS_','NEUTRAL_','DIAG_')) and k not in ['PPROF','GODEBUG','GOMEMLIMIT','GOFLAGS','GOEXPERIMENT','NODE_OPTIONS']}
env.update(GOMAXPROCS='14',GOGC='100',GOWORK='off')
command=['go','build','-pgo=off','-o',str(out/'benchmark'),'./main']
subprocess.run(command,cwd=out/'output',env=env,check=True)
manifest=json.loads((base/'manifest.json').read_text())
manifest.update(binary_sha256=sha(out/'benchmark'),prototype={
    'baseline_manifest_sha256':sha(base/'manifest.json'),
    'build_command':command,
    'sources':{str(p):sha(p) for p in [root/'index.go',Path(__file__).resolve()]},
    'scope':'combined only: validated structural index, lazy decoder inputs, owned final strings, complete existing TAST/type-table/usage decoder'})
(out/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')

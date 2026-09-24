#!/usr/bin/env python3
"""Rebuild a preserved diagnostic with changed native decoder internals only.

Public FFI signatures and generated wrappers must match the preserved build.
Source copies, emitted Go, commands and binaries remain in the new workspace.
"""
import argparse, hashlib, json, os, re, shutil, subprocess
from pathlib import Path

ROOT = Path(__file__).resolve().parents[4]
PBO = ROOT.parent/'purescript-backend-optimizer-gopurs'
sha = lambda p: hashlib.sha256(p.read_bytes()).hexdigest()

def body(text):
    return text.split('// --- Auto-generated FFI wrappers ---')[0]

def replace_ffi(source, original, prefix):
    text = source.read_text()
    declarations = re.findall(r'^func ([A-Z]\w*)\([^\n]+',text,re.M)
    previous = body(original)
    for name in declarations:
        signature = re.search(r'^func '+name+r'\([^\n]+',text,re.M)[0]
        signature = signature.replace('func '+name+'(', 'func '+prefix+'_'+name+'(')
        assert signature in previous, ('FFI signature changed',name)
        text = re.sub(r'\b'+name+r'\b',prefix+'_'+name,text)
    text = re.sub(r'^package \w+\n','package purescript\n',text, count=1,flags=re.M)
    marker = '// --- Auto-generated FFI wrappers ---'
    assert original.count(marker) == 1
    return text+'\n'+marker+original.split(marker,1)[1]

def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--baseline',type=Path,required=True)
    parser.add_argument('--output',type=Path,required=True)
    args = parser.parse_args()
    base, out = args.baseline.resolve(),args.output.resolve()
    manifest = json.loads((base/'manifest.json').read_text())
    assert sha(base/'benchmark') == manifest['binary_sha256']
    out.mkdir(parents=True)
    shutil.copytree(base/'output',out/'output',ignore=shutil.ignore_patterns('*_test.go'))
    changed = {}
    for relative,prefix in [('Json.go','PureScript_Backend_Optimizer_CoreFn_Json'),('Json/Text.go','PureScript_Backend_Optimizer_CoreFn_Json_Text')]:
        source = PBO/'src/PureScript/Backend/Optimizer/CoreFn'/relative
        snapshot = out/'sources'/relative
        snapshot.parent.mkdir(parents=True,exist_ok=True)
        shutil.copy2(source,snapshot)
        target = out/'output/purescript'/(prefix+'_ffi.go')
        target.write_text(replace_ffi(snapshot,target.read_text(),prefix))
        changed[str(source)] = sha(snapshot)
    env = {k:v for k,v in os.environ.items() if not k.startswith(('GOPURS_','NEUTRAL_','DIAG_')) and k not in ['PPROF','GODEBUG','GOMEMLIMIT','GOFLAGS','GOEXPERIMENT','NODE_OPTIONS']}
    env.update(GOWORK='off',GOMAXPROCS='14',GOGC='100')
    command = ['go','build','-pgo=off','-o',str(out/'benchmark'),'./main']
    subprocess.run(command,cwd=out/'output',env=env,check=True)
    manifest.update(binary_sha256=sha(out/'benchmark'),prototype={
        'scope':'canonical native type-table internals; existing public FFI and diagnostic driver',
        'baseline_manifest_sha256':sha(base/'manifest.json'),'sources':changed,
        'build_command':command,'builder_sha256':sha(Path(__file__).resolve()),
        'generated_go':{str(p.relative_to(out)):sha(p) for p in sorted((out/'output').rglob('*.go'))}})
    (out/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')
    print(out)

if __name__ == '__main__':
    main()

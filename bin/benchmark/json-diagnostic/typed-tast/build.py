#!/usr/bin/env python3
"""Build a typed-input specialization of the existing native TAST decoder.

Schema control flow, constructors, type resolution and usage validation are
preserved mechanically. Only JSON access is specialized; no any cursor boxes
or temporary JSON arrays remain. Generated source is retained for auditing.
"""
import argparse, hashlib, json, os, re, shutil, subprocess
from pathlib import Path

def sha(path): return hashlib.sha256(path.read_bytes()).hexdigest()

def function(text,name,replacement):
    pattern=rf'^func {name}\([^\n]*\{{[^\n]*\}}\n|^func {name}\([^\n]*\{{\n.*?^\}}\n'
    result,n=re.subn(pattern,lambda _:replacement+'\n',text,flags=re.M|re.S)
    if n!=1: raise ValueError((name,n))
    return result

def specialize(text):
    text=text.split('// --- Auto-generated FFI wrappers ---')[0]
    text=re.sub(r'\bPureScript_Backend_Optimizer_CoreFn_Json_(Decode\w+Impl|IsNonNegativeInteger)\b',r'\1',text)
    for name in ['DecodeTypeTableImpl','IsNonNegativeInteger','DecodeArrayImpl','DecodeAnnWithUsageImpl']:
        text=function(text,name,'')
    bodies={
      'ntNative':'func ntNative(input any) any { return input }',
      'ntIsNull':"func ntIsNull(input any) (bool,bool) { return input.kind() == 'n', true }",
      'ntObjectOf':"func ntObjectOf(raw any) (ntObject,bool) { return raw, raw.kind() == '{' }",
      'ndNullable':"func ndNullable(raw any) bool { return raw.kind() == 'n' }",
      'ndNumber':'func ndNumber(raw any) (float64,*ndFailure) { if n,ok:=raw.number();ok{return n,nil};return 0,ndPublic("Number") }',
      'ndArray':'func ndArray(raw any) (tcArray,*ndFailure) { if a,ok:=raw.array();ok{return a,nil};return tcArray{},ndPublic("Array") }',
      'cndNumber':'func cndNumber(raw any) float64 { if n,ok:=raw.number();ok{return n};cndFail("Number");return 0 }',
      'cndElements':'func cndElements(raw any) tcArray { if a,ok:=raw.array();ok{return a};cndFail("Array");return tcArray{} }',
      'cndIsNull':"func cndIsNull(raw any) bool { return raw.kind() == 'n' }",
    }
    for name,body in bodies.items():text=function(text,name,body)
    text=text.replace('type ntObject = gopurs_runtime.JSONObjectView','type ntObject = tcCursor')
    text=re.sub(r'ntNative\(([\w.\[\]]+)\)',r'\1',text)
    for typ,method in [('string','text'),('float64','number'),('bool','boolean'),('[]any','array')]:
        text=re.sub(r'(\w+)\.\('+re.escape(typ)+r'\)',rf'\1.{method}()',text)
    text=text.replace('for i := range entries {\n\t\ttable.refs[i] = table.decodeRef(entries[i])','for i, entry := range entries.each {\n\t\ttable.refs[i] = table.decodeRef(entry)')
    text=re.sub(r'len\((arr|entries|elements)\)',r'\1.count',text)
    text=re.sub(r'range (arr|elements)\b',r'range \1.each',text)
    text=text.replace('range cndElements(raw) {','range cndElements(raw).each {')
    text=re.sub(r'for (\w+), (\w+) := range ([^\n]+)\.each \{',r'for cursorLoop := \3.iter(); cursorLoop.more(); {\n\1, \2 := cursorLoop.next()',text)
    text=re.sub(r'elements\[(\d+)\]',r'elements.at(\1)',text)
    text=re.sub(r'cndString\((cndFieldOf\([^\n]+?\))\)',r'\1.StrVal()',text)
    text=text.replace('jsonValue := gopurs_runtime.Box(raw)','jsonValue := gopurs_runtime.Box(directMaterialize(directCursor(raw)))')
    text=text.replace('json gopurs_runtime.Value) (result gopurs_runtime.Value)','json tcCursor) (result gopurs_runtime.Value)')
    text=re.sub(r'\bany\b','tcCursor',text)
    # Prefix native helper identities so this experiment coexists with the DOM
    # controls in the same binary and cannot accidentally replace their path.
    text=re.sub(r'\b((?:nt|nd|cnd)[A-Z]\w*|decodeTypeTableNative)\b',lambda m:'tc_'+m[0],text)
    text=text.replace('func DecodeModuleImpl(','func tcDecodeModule(')
    # These strings only select a branch; their bytes cannot enter the output.
    # Assert the exact sites, leaving every stored name/literal/path owned.
    replacements={
      'if s, ok := raw.text(); ok {':'if s, ok := raw.borrowedText(); ok {',
      'typ, typOK := typRaw.text()':'typ, typOK := typRaw.borrowedText()',
    }
    for before,after in replacements.items():
        assert text.count(before)==1,before
        text=text.replace(before,after)
    for label in ['type','binderType','literalType','bindType']:
        before=f'kind := tc_cndFieldOf(obj, "{label}", tc_cndStringValue).StrVal()'
        assert text.count(before)==1,before
        text=text.replace(before,before.replace('tc_cndStringValue','tcBorrowedTag'))
    return text

def main():
    parser=argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--baseline',type=Path,required=True)
    parser.add_argument('--output',type=Path,required=True)
    parser.add_argument('--resume',action='store_true')
    parser.add_argument('--dom-control',action='store_true',help='build the ordinary DOM control with the same phase-order-aware driver')
    args=parser.parse_args();base=args.baseline.resolve();out=args.output.resolve()
    root=Path(__file__).resolve().parent
    if (out/'manifest.json').exists():
        raise ValueError('Preserve completed builds; choose a new output workspace')
    if not args.resume: shutil.copytree(base,out)
    output=out/'output/purescript'
    source=base/'output/purescript/PureScript_Backend_Optimizer_CoreFn_Json_ffi.go'
    if not args.dom_control:
        (output/'typed_tast_decoder.go').write_text(specialize(source.read_text()))
        shutil.copy2(root/'cursor.go',output/'typed_tast_cursor.go')
        shutil.copy2(root/'index.go',output/'typed_tast_index.go')
    driver=output/'Test_JsonTypedAst_ffi.go'
    text=(base/'output/purescript/Test_JsonTypedAst_ffi.go').read_text();old='case "combined":results[i]=gopurs_runtime.Apply(decode,gopurs_runtime.Apply(parse,gopurs_runtime.Str(f.Contents)))'
    assert text.count(old)==1
    if not args.dom_control:
        text=text.replace(old,'case "combined":results[i]=typedTAST(parse,decode,f.Contents)')
    text=text.replace('"encoding/json"','"encoding/json"\n    "strings"')
    old_order='for _,phase:=range []string{"parse","decode","combined"} {'
    assert text.count(old_order)==1
    text=text.replace(old_order,'''phaseOrder := []string{"parse","decode","combined"}
    if requested := os.Getenv("DIAG_PHASES"); requested != "" { phaseOrder = strings.Split(requested, ",") }
    seen := map[string]bool{}
    for _,phase := range phaseOrder {
        if seen[phase] || (phase != "parse" && phase != "decode" && phase != "combined") { panic("invalid DIAG_PHASES") }
        seen[phase] = true
    }
    for _,phase:=range phaseOrder {''')
    text=text.replace('"backend":"go"','"phase_order":phaseOrder,"backend":"go"')
    driver.write_text(text)
    env={k:v for k,v in os.environ.items() if not k.startswith(('GOPURS_','NEUTRAL_','DIAG_')) and k not in ['PPROF','GODEBUG','GOMEMLIMIT','GOFLAGS','GOEXPERIMENT','NODE_OPTIONS']}
    env.update(GOMAXPROCS='14',GOGC='100',GOWORK='off')
    command=['go','build','-pgo=off','-o',str(out/'benchmark'),'./main']
    subprocess.run(command,cwd=out/'output',env=env,check=True)
    manifest=json.loads((base/'manifest.json').read_text())
    manifest.update(binary_sha256=sha(out/'benchmark'),driver_capabilities=['phase-order'],prototype={
      'baseline_manifest_sha256':sha(base/'manifest.json'),'build_command':command,
      'sources':{str(p):sha(p) for p in [root/'cursor.go',root/'index.go',Path(__file__).resolve()]},
      'scope':('ordinary compact DOM control' if args.dom_control else 'combined only: typed cursors and array views directly feed final native constructors; full validation, eager type-table resolution and usage checks')})
    (out/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')

if __name__=='__main__':main()

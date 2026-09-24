#!/usr/bin/env python3
"""Differentially validate typed cursors and complete final TAST construction."""
import argparse,os,shutil,subprocess
from pathlib import Path

parser=argparse.ArgumentParser(description=__doc__)
parser.add_argument('--workspace',type=Path,required=True)
parser.add_argument('--tests',type=Path,required=True)
parser.add_argument('--corpus',type=Path,required=True)
parser.add_argument('--log',type=Path,required=True)
parser.add_argument('--integrated',action='store_true',help='exercise the public Json.Text.parseModule API and its parser/error fallback')
args=parser.parse_args();out=args.workspace.resolve()/'output'
root=Path(__file__).resolve().parent
for name in ['native_decode_test.go','native_decode_errors_test.go','type-table-native_test.go']:
    text=(args.tests/name).read_text()
    if name!='type-table-native_test.go':
        old='p.PureScript_Backend_Optimizer_CoreFn_Json_DecodeModuleImpl(rt.Value{}, p.Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), jsonValue)'
        assert old in text
        text=text.replace(old,'p.DirectDecodeModuleText('+('file.Contents' if name=='native_decode_test.go' else 'string(text)')+')')
        if args.integrated:
            text=text.replace('rt.CoerceToStruct[p.Constructor_Data_Show_Show[rt.Value]](p.Get_Data_Argonaut_Decode_Error_showJsonDecodeError()).V0','p.Get_Data_Argonaut_Decode_Error_printJsonDecodeError()')
            text=text.replace('return rt.Apply(show, err).StrVal()','if err.Type == rt.TypeString { return err.StrVal() }; return rt.Apply(show, err).StrVal()')
            text=text.replace('rt.Apply(show, left).StrVal()', 'decodeError(show, left)')
    else:
        text=text.replace('result := rt.Apply(p.Get_Control_Monad_ST_Internal_run(), p.Call_PureScript_Backend_Optimizer_CoreFn_TypeTable_decodeTypeTableST(input))','result := p.TypedDecodeTypeTableText(text)')
    (out/'purescript'/name).write_text(text)
text=(root.parent/'direct-tast/index_test.go').read_text()
text=text.replace('return PureScript_Backend_Optimizer_CoreFn_Json_DecodeModuleImpl(rt.Value{}, Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), rt.Any(cursor))','return tcDecodeModule(rt.Value{}, Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), tcCursor(cursor))')
if args.integrated:
    text=text.replace('return tcDecodeModule(rt.Value{}, Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), tcCursor(cursor))','_ = cursor; return rt.Apply(Get_PureScript_Backend_Optimizer_CoreFn_Json_Text_parseModule(), rt.Str(text))')
text+='''
func TypedDecodeTypeTableText(text string) rt.Value {
 cursor,ok:=directIndex(text);if !ok {panic("invalid type table")}
 return tc_decodeTypeTableNative(tcCursor(cursor))
}
'''
(out/'purescript/direct_index_test.go').write_text(text)
text=(root/'cursor_test.go').read_text()
if args.integrated:
    text=text.replace('tcDecodeModule(rt.Value{}, Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), tcCursor(cursor))','rt.Apply(Get_PureScript_Backend_Optimizer_CoreFn_Json_Text_parseModule(), rt.Str(text))')
    text=text.replace('cursor, ok := directIndex(text)\n\t\tif !ok {\n\t\t\tt.Fatal(file.Name)\n\t\t}', '_, ok := directIndex(text)\n\t\tif !ok {\n\t\t\tt.Fatal(file.Name)\n\t\t}')
    shutil.copy2(root.parents[4]/'purescript-backend-optimizer-gopurs/test/json-text-native_test.go',out/'purescript/json_text_native_test.go')
(out/'purescript/typed_cursor_test.go').write_text(text)
env={**os.environ,'GOWORK':'off','GOMAXPROCS':'1','GOGC':'100','DIAG_CORPUS':str(args.corpus.resolve())}
with args.log.open('w') as log:
    subprocess.run(['go','test','-count=1','-v','./purescript'],cwd=out,env=env,stdout=log,stderr=subprocess.STDOUT,check=True)
    subprocess.run(['go','test','-race','-count=1','-run','TestDirect|TestTyped|TestPublicModule','-v','./purescript'],cwd=out,env=env,stdout=log,stderr=subprocess.STDOUT,check=True)

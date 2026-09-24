#!/usr/bin/env python3
"""Run a preserved TAST differential suite against the indexed candidate."""
import argparse,os,shutil,subprocess
from pathlib import Path
parser=argparse.ArgumentParser(description=__doc__)
parser.add_argument('--workspace',type=Path,required=True)
parser.add_argument('--tests',type=Path,required=True)
parser.add_argument('--corpus',type=Path,required=True)
parser.add_argument('--log',type=Path,required=True)
args=parser.parse_args();out=args.workspace.resolve()/'output'
for name in ['native_decode_test.go','native_decode_errors_test.go','type-table-native_test.go']:
    text=(args.tests/name).read_text()
    if name!='type-table-native_test.go':
        text=text.replace('p.PureScript_Backend_Optimizer_CoreFn_Json_DecodeModuleImpl(rt.Value{}, p.Get_PureScript_Backend_Optimizer_CoreFn_Usage_validateSourceUsageModule(), jsonValue)',
            'p.DirectDecodeModuleText('+('file.Contents' if name=='native_decode_test.go' else 'string(text)')+')')
    (out/'purescript'/name).write_text(text)
shutil.copy2(Path(__file__).with_name('index_test.go'),out/'purescript/direct_index_test.go')
env={**os.environ,'GOWORK':'off','GOMAXPROCS':'1','GOGC':'100','DIAG_CORPUS':str(args.corpus.resolve())}
with args.log.open('w') as log:
    subprocess.run(['go','test','-count=1','-v','./purescript'],cwd=out,env=env,stdout=log,stderr=subprocess.STDOUT,check=True)
    subprocess.run(['go','test','-race','-count=1','-run','TestDirect','-v','./purescript'],cwd=out,env=env,stdout=log,stderr=subprocess.STDOUT,check=True)

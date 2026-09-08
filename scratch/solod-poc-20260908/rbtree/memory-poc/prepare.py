#!/usr/bin/env python3
"""Keep generated algorithms exact; redirect only Node construction sites."""
import difflib
import hashlib
import json
from pathlib import Path
import re

ROOT=Path(__file__).resolve().parent
SOURCE=ROOT.parents[3]/'run/bak/go/output/purescript/Test_RBTree.go'
source=SOURCE.read_text()
start=source.index('type Constructor_Test_RBTree_T struct {')
raw=source[start:]
assert 'gopurs_runtime' not in raw
adapted=raw
sites=0
while '&Constructor_Test_RBTree_T{' in adapted:
    adapted,n=re.subn(r'&Constructor_Test_RBTree_T\{([^{}]*)\}',r'AllocNode(\1)',adapted)
    assert n>0
    sites+=n
assert sites==86
stub='''package kernels
func ConfigurePool(capMiB int) {}
func ResetPool() {}
func AllocationCounts() (int64,int64,int64) { return -1,0,-1 }
'''
for name,body,allocator in [('original',raw,stub),('pool',adapted,(ROOT/'pool.go.txt').read_text())]:
    out=ROOT/name
    (out/'kernels').mkdir(parents=True,exist_ok=True)
    (out/'cmd').mkdir(exist_ok=True)
    (out/'go.mod').write_text('module rbmemory\n\ngo 1.22\n')
    (out/'kernels/kernels.go').write_text('package kernels\n\n'+body)
    (out/'kernels/allocator.go').write_text(allocator)
    (out/'kernels/validate.go').write_text((ROOT.parent/'validate.go').read_text())
    (out/'cmd/main.go').write_text((ROOT/'driver.go.txt').read_text())
(ROOT/'allocation.diff').write_text(''.join(difflib.unified_diff(raw.splitlines(True),adapted.splitlines(True),fromfile='generated-go',tofile='pool-constructors-go')))
(ROOT/'sources.json').write_text(json.dumps({'source':str(SOURCE),'source_sha256':hashlib.sha256(SOURCE.read_bytes()).hexdigest(),'extracted_sha256':hashlib.sha256(raw.encode()).hexdigest(),'constructor_sites':sites},indent=2)+'\n')
print(f'Extracted RBTree unchanged; {sites} constructor sites redirected in pool variant.')

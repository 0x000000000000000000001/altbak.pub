#!/usr/bin/env python3
"""Stage the same Go checks/probe against one generated package; do not compile."""
import argparse
import hashlib
import json
from pathlib import Path
import re
import shutil

parser = argparse.ArgumentParser(description=__doc__)
parser.add_argument('--generated', type=Path, required=True)
parser.add_argument('--target', type=Path, required=True)
args = parser.parse_args()
generated = args.generated.resolve(strict=True)
if not (generated / 'purescript/Test_RBTree.go').is_file():
    raise SystemExit('Expected generated output with Test_RBTree.go')
target = args.target.resolve()
target.mkdir()  # An existing directory is an error; never overwrite a run.
here = Path(__file__).resolve().parent
for source in (here / 'harness').glob('*.go'):
    shutil.copy2(source, target / source.name)
tree_source = (generated / 'purescript/Test_RBTree.go').read_text()
def entrypoint(name):
    workers = re.findall(r'^func (Call_Test_RBTree___gopurs_owned_' + name + r'_\d+)\(', tree_source, re.MULTILINE)
    if len(workers) > 1:
        raise SystemExit('Ambiguous owned entrypoint: ' + name)
    return workers[0] if workers else 'Call_Test_RBTree_' + name
build_entry, insert_entry = entrypoint('buildTree'), entrypoint('insert')
if '___gopurs_owned_' in build_entry:
    act = re.search(r'(?ms)^func Get_Test_RBTree_act\(.*?^}', tree_source)
    if act is None or build_entry + '(' not in act.group(0):
        raise SystemExit('Owned build wrapper is not the entry used by the benchmark act')
(target / 'entrypoints.go').write_text('package main\n\nimport ps "gopurs/output/purescript"\n\n'
    '// The build root is always empty; the same entry is selected by benchmark act.\n'
    f'func buildFresh(n int64) *ps.Constructor_Test_RBTree_T {{ return ps.{build_entry}(n, nil) }}\n\n'
    '// The caller must supply exclusive nodes and retain no root/subtree snapshots.\n'
    f'func insertFresh(key int64, root *ps.Constructor_Test_RBTree_T) *ps.Constructor_Test_RBTree_T {{ return ps.{insert_entry}(key, root) }}\n')
(target / 'entrypoints.json').write_text(json.dumps({'buildFresh': build_entry, 'insertFresh': insert_entry,
    'precondition': 'buildFresh starts from nil; insertFresh accepts only exclusively constructed, unretained nodes'}, indent=2) + '\n')
(target / 'go.mod').write_text('module adt-reuse-validation\n\ngo 1.22\n\n'
    'require gopurs/output v0.0.0\n\n'
    f'replace gopurs/output => {json.dumps(str(generated))}\n')
hashes = {str(path.relative_to(generated)): hashlib.sha256(path.read_bytes()).hexdigest()
          for path in sorted(generated.rglob('*.go'))}
(target / 'generated-sha256.json').write_text(json.dumps(hashes, indent=2) + '\n')
print(target)

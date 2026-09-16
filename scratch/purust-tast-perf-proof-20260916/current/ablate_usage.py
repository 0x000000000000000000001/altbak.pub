#!/usr/bin/env python3
"""Read saved TAST, ablate only usageCount/escapes, run existing backend isolated."""
import collections
import copy
import hashlib
import json
import pathlib
import re
import subprocess
import sys
import time

ROOT = pathlib.Path(__file__).resolve().parent
SAVED = ROOT.parents[1] / 'purust-fbip-20260916' / 'benchmark'
BACKEND = pathlib.Path('/Users/0x1/Documents/htdocs/purust/purust/bin/purust')

def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()

def strip(value, removed):
    if isinstance(value, dict):
        for key in ('usageCount', 'escapes'):
            if key in value:
                removed[key] += 1
                del value[key]
        for child in value.values(): strip(child, removed)
    elif isinstance(value, list):
        for child in value: strip(child, removed)

def occurrences(value, path='$', top='<module>'):
    if isinstance(value, dict):
        if 'identifier' in value and re.fullmatch(r'\$\.decls\[\d+\](?:\.binds\[\d+\])?', path):
            top = value.get('identifier', top)
        ann = value.get('annotation', {})
        if 'usageCount' in ann or 'escapes' in ann:
            yield dict(function=top, node=value.get('type', value.get('binderType', value.get('bindType', 'RecBinding'))),
                       name=value.get('argument', value.get('identifier', value.get('name'))),
                       usageCount=ann.get('usageCount'), escapes=ann.get('escapes'),
                       span=ann.get('sourceSpan'), path=path)
        for key, child in value.items():
            yield from occurrences(child, path + '.' + key, top)
    elif isinstance(value, list):
        for i, child in enumerate(value):
            yield from occurrences(child, path + f'[{i}]', top)

def main():
    total_removed = collections.Counter()
    source_hashes = {}
    for mode in ('with-usage', 'without-usage'):
        (ROOT / mode / 'input').mkdir(parents=True, exist_ok=True)
    records = []
    files = sorted((SAVED / 'output').glob('*/corefn.json'))
    for path in files:
        original = json.loads(path.read_text())
        source_hashes[str(path.relative_to(SAVED))] = sha(path)
        if path.parent.name == 'Test.RBTree':
            records = list(occurrences(original))
        # Both sides get exactly the same absolute paths so the backend can read
        # existing FFI/source files from its isolated working directory.
        original['modulePath'] = str((SAVED / original['modulePath']).resolve())
        ablated = copy.deepcopy(original)
        strip(ablated, total_removed)
        for mode, value in [('with-usage', original), ('without-usage', ablated)]:
            target = ROOT / mode / 'input' / path.parent.name / 'corefn.json'
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_text(json.dumps(value, separators=(',', ':')))
    stats = dict(total=len(records), by_kind=dict(collections.Counter(r['node'] for r in records)),
                 by_usage_escape=dict(collections.Counter(f"{r['usageCount']}/{r['escapes']}" for r in records)),
                 by_function=dict(collections.Counter(r['function'] for r in records)),
                 by_function_kind_usage_escape=dict(collections.Counter(
                     f"{r['function']}/{r['node']}/{r['usageCount']}/{r['escapes']}" for r in records)))
    (ROOT / 'metadata.json').write_text(json.dumps(dict(summary=stats, occurrences=records), indent=2)+'\n')
    if '--metadata-only' in sys.argv:
        path = ROOT / 'ablation-results.json'
        result = json.loads(path.read_text())
        result['metadata'] = stats
        path.write_text(json.dumps(result, indent=2)+'\n')
        print(json.dumps(stats, indent=2))
        return
    results = dict(input_modules=len(files), removed_fields=dict(total_removed), metadata=stats,
                   source_hashes=source_hashes, backend_sha256=sha(BACKEND.with_suffix('.js')), runs=[])
    for mode in ('with-usage', 'without-usage'):
        directory = ROOT / mode
        command = [str(BACKEND), '--main', 'AppX', '--source', 'input', '--out', 'rust',
                   '--ffi-dir', str(SAVED.parent)]
        start = time.monotonic()
        with (directory / 'backend.log').open('w') as log:
            process = subprocess.run(command, cwd=directory, stdout=log, stderr=subprocess.STDOUT)
        results['runs'].append(dict(mode=mode, command=command, cwd=str(directory),
                                    returncode=process.returncode, elapsed_s=time.monotonic()-start))
        if process.returncode != 0:
            (ROOT / 'ablation-results.json').write_text(json.dumps(results, indent=2)+'\n')
            raise RuntimeError(f'{mode} failed; inspect backend.log')
    a, b = [ROOT / mode / 'rust' for mode in ('with-usage', 'without-usage')]
    fa = {str(p.relative_to(a)): sha(p) for p in a.rglob('*') if p.is_file()}
    fb = {str(p.relative_to(b)): sha(p) for p in b.rglob('*') if p.is_file()}
    results['generated_files'] = len(fa)
    results['generated_rust_files'] = sum(p.endswith('.rs') for p in fa)
    results['differing_files'] = [p for p in sorted(fa.keys() | fb.keys()) if fa.get(p) != fb.get(p)]
    results['rust_hashes'] = fa
    results['rbtree_sha256'] = fa['Purs_Test_RBTree/src/lib.rs']
    saved_rb = SAVED / 'output/purust_output/Purs_Test_RBTree/src/lib.rs'
    results['matches_saved_rbtree'] = sha(saved_rb) == results['rbtree_sha256']
    (ROOT / 'ablation-results.json').write_text(json.dumps(results, indent=2)+'\n')
    print(json.dumps({k:v for k,v in results.items() if k not in ('source_hashes','rust_hashes')}, indent=2))

if __name__ == '__main__': main()

"""Freeze the current program and replace only the internal ins/balance path."""
from pathlib import Path
import difflib
import hashlib
import json
import shutil
import subprocess

AUDIT = Path(__file__).resolve().parent
PROJECT = AUDIT.parents[1]
SOURCE = PROJECT / 'run/bak/sharp/output/Main'
PRIOR = PROJECT / 'scratch/sharpurs-int-arithmetic-native-20260909'
PROBE = PROJECT / 'scratch/sharpurs-rbt-typed-probe-20260909/probe.fsx'
BACKEND = PROJECT.parent / 'sharpurs/sharpurs'


def require(condition, message):
    if not condition:
        raise ValueError(message)


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


require(not any((AUDIT / name).exists() for name in ['before', 'after', 'inputs.json']),
        'Preserve existing snapshots and manifest')
names = sorted(p.name for p in SOURCE.iterdir()
               if p.is_file() and p.suffix in {'.fs', '.cs', '.fsproj', '.csproj', '.props'})
before = {name: sha(SOURCE / name) for name in names}
require(len(names) == 355, 'Expected 355 build inputs')
require(before == json.loads((PRIOR / 'inputs.json').read_text())['afterSha256'],
        'Expected normal program after native Int arithmetic integration')
normal_runtime = SOURCE / 'bin/Release/net8.0'
normal_runtime_hashes = {p.name: sha(p) for p in sorted(normal_runtime.iterdir()) if p.is_file()}

original = (SOURCE / 'Test.RBTree.fs').read_text()
probe = PROBE.read_text()
kernel_start = probe.index('// Same four Okasaki rotations')
kernel_end = probe.index('let Test_RBTree_ins_tco ', kernel_start)
kernels = probe[kernel_start:kernel_end]
for old, new in [
    ('typedBalanceApply', 'Test_RBTree_balance_adt_native_apply'),
    ('typedBalance', 'Test_RBTree_balance_adt_native'),
    ('typedIns', 'Test_RBTree_ins_adt_native'),
]:
    kernels = kernels.replace(old, new)
require('sharpurs_apply' not in kernels and 'obj' not in kernels,
        'Native bodies must use only their concrete types')
original_ins = [line for line in original.splitlines() if line.startswith('let rec Test_RBTree_ins_tco ')]
require(len(original_ins) == 1, 'Expected exactly one recursive ins definition')
bridge = ('let rec Test_RBTree_ins_tco (v: obj) (v1: obj) : obj = '
          'box (Test_RBTree_ins_adt_native (unbox<int> v) (unbox<Test_RBTree_Tree> v1))')
patched = original.replace(original_ins[0], kernels + bridge, 1)
require(patched.replace(kernels + bridge, original_ins[0], 1) == original,
        'Only insert native kernels and replace ins entry body')
# Public balance, its raw and guarded direct entry points, public ins wrapper,
# insert/buildTree, constructor layout, depth and makeBlack remain unchanged.
for prefix in ['let Test_RBTree_balance_direct ', 'let Test_RBTree_balance_direct_apply ',
               'let Test_RBTree_balance =', 'and Test_RBTree_ins =',
               'let Test_RBTree_insert_direct ', 'let rec Test_RBTree_buildTree_tco ']:
    lines = [line for line in original.splitlines() if line.startswith(prefix)]
    require(len(lines) == 1 and lines[0] in patched, 'Preserve ' + prefix)

for variant in ['before', 'after']:
    destination = AUDIT / variant
    destination.mkdir()
    for name in names:
        shutil.copy2(SOURCE / name, destination / name)
(AUDIT / 'after/Test.RBTree.fs').write_text(patched)
(AUDIT / 'native-kernels.fs.txt').write_text(kernels)
(AUDIT / 'Test.RBTree.fs.generated').write_text(patched)
(AUDIT / 'Test.RBTree.fs.diff').write_text(''.join(difflib.unified_diff(
    original.splitlines(True), patched.splitlines(True),
    fromfile='before/Test.RBTree.fs', tofile='after/Test.RBTree.fs')))
after = {name: sha(AUDIT / 'after' / name) for name in names}
changed = [name for name in names if before[name] != after[name]]
require(changed == ['Test.RBTree.fs'], 'Exactly one generated file changes')
require(before == {name: sha(SOURCE / name) for name in names}, 'Normal sources changed during capture')
readme = PROJECT.parent / 'altbak.pub/README.md'
shutil.copy2(readme, AUDIT / 'README.baseline.md')
manifest = {
    'beforeSha256': before, 'afterSha256': after,
    'sourceGeneratedDirectory': str(SOURCE), 'sourceGeneratedSha256': before,
    'sourceRuntimeSha256': normal_runtime_hashes,
    'changedFiles': changed, 'priorManifestSha256': sha(PRIOR / 'inputs.json'),
    'priorAfterEqualsBefore': True, 'prototypeSha256': sha(PROBE),
    'officialReadmeSha256': sha(AUDIT / 'README.baseline.md'),
    'backendCommit': subprocess.check_output(['git', 'rev-parse', 'HEAD'], cwd=BACKEND, text=True).strip(),
    'backendBundleSha256': sha(BACKEND / 'bin/sharpurs.js'),
    'backendSourcesSha256': {str(p.relative_to(BACKEND)): sha(p)
                            for p in sorted((BACKEND / 'src').rglob('*.purs'))},
    'transformation': 'Native ins/balance internal kernels; ins_tco bridges obj ABI; every other generated definition remains unchanged',
}
(AUDIT / 'inputs.json').write_text(json.dumps(manifest, indent=2) + '\n')
print(f'Captured {len(names)} sources per variant; only Test.RBTree.fs differs; normal generation unchanged')

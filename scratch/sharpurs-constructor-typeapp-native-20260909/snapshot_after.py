"""Capture the normally generated integration; never patches or regenerates F#."""
from pathlib import Path
import difflib
import hashlib
import json
import shutil
import subprocess

from common import AUDIT, require, sha

manifest_path = AUDIT / 'inputs.json'
manifest = json.loads(manifest_path.read_text())
source = Path(manifest['sourceGeneratedDirectory'])
backend = AUDIT.parents[2] / 'sharpurs/sharpurs'
names = sorted(p.name for p in source.iterdir()
               if p.is_file() and p.suffix in {'.fs', '.cs', '.fsproj', '.csproj', '.props'})
require(set(names) == set(manifest['beforeSha256']), 'Generated source file set changed')
require(not (AUDIT / 'after').exists(), 'Preserve the existing after snapshot')
require('afterSha256' not in manifest, 'Preserve the existing integration manifest')
after = {name: sha(source / name) for name in names}
changed = [name for name in names if manifest['beforeSha256'][name] != after[name]]
require(changed, 'No generated change to measure')
(AUDIT / 'after').mkdir()
for name in names:
    shutil.copy2(source / name, AUDIT / 'after' / name)
for name in changed:
    original = (AUDIT / 'before' / name).read_text()
    generated = (source / name).read_text()
    (AUDIT / (name + '.diff')).write_text(''.join(difflib.unified_diff(
        original.splitlines(True), generated.splitlines(True),
        fromfile='before/' + name, tofile='after/' + name)))
    (AUDIT / (name + '.generated')).write_text(generated)
manifest.update({'afterSha256': after, 'sourceGeneratedSha256': after, 'changedFiles': changed,
    'backendAfterCommit': subprocess.check_output(['git', 'rev-parse', 'HEAD'], cwd=backend, text=True).strip(),
    'backendAfterBundleSha256': sha(backend / 'bin/sharpurs.js'),
    'backendAfterSourcesSha256': {str(p.relative_to(backend)): sha(p)
        for p in sorted((backend / 'src').rglob('*.purs'))}})
require(all(sha(source / name) == expected for name, expected in after.items()), 'Normal generation changed during capture')
shutil.copytree(backend / 'bin', AUDIT / 'backend-after-bin', symlinks=True)
require(sha(AUDIT / 'backend-after-bin/sharpurs.js') == manifest['backendAfterBundleSha256'], 'Bundle changed during capture')
manifest_path.write_text(json.dumps(manifest, indent=2) + '\n')
print(f'Captured {len(names)} integration inputs; {len(changed)} normally generated files changed:')
print('\n'.join(changed))

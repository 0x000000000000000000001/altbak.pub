"""Freeze the normal program before the compiler integration; refuses to overwrite."""
from pathlib import Path
import json
import shutil
import subprocess
from common import AUDIT, require, sha

project = AUDIT.parents[1]
source = project / 'run/bak/sharp/output/Main'
prior = project / 'scratch/sharpurs-rbt-typed-native-20260909'
backend = project.parent / 'sharpurs/sharpurs'
require(not (AUDIT / 'before').exists() and not (AUDIT / 'inputs.json').exists(),
        'Preserve existing baseline snapshot and manifest')
names = sorted(p.name for p in source.iterdir()
               if p.is_file() and p.suffix in {'.fs', '.cs', '.fsproj', '.csproj', '.props'})
hashes = {name: sha(source / name) for name in names}
previous = json.loads((prior / 'inputs.json').read_text())
require(hashes == previous['beforeSha256'], 'Expected normal program matching prior experiment before')
(AUDIT / 'before').mkdir()
for name in names:
    shutil.copy2(source / name, AUDIT / 'before' / name)
shutil.copy2(project.parent / 'altbak.pub/README.md', AUDIT / 'README.baseline.md')
manifest = {'beforeSha256': hashes, 'sourceGeneratedDirectory': str(source),
    'baselineMatchesPriorExperimentBefore': True,
    'baselinePreviousManifestSha256': sha(prior / 'inputs.json'),
    'backendBeforeCommit': subprocess.check_output(['git', 'rev-parse', 'HEAD'], cwd=backend, text=True).strip(),
    'backendBeforeBundleSha256': sha(backend / 'bin/sharpurs.js'),
    'backendSourceSnapshotNote': 'Working sources captured during concurrent integration development; generated before is proven by the previous experiment manifest. Bundle and generated inputs are frozen separately.',
    'backendBeforeSourcesSha256': {str(p.relative_to(backend)): sha(p)
        for p in sorted((backend / 'src').rglob('*.purs'))},
    'officialReadmeSha256': sha(AUDIT / 'README.baseline.md')}
require(all(sha(source / name) == expected for name, expected in hashes.items()),
        'Normal generation changed during baseline capture')
shutil.copy2(backend / 'bin/sharpurs.js', AUDIT / 'backend-before.bundle.js')
require(sha(AUDIT / 'backend-before.bundle.js') == manifest['backendBeforeBundleSha256'], 'Bundle changed during capture')
(AUDIT / 'inputs.json').write_text(json.dumps(manifest, indent=2) + '\n')
print(f'Captured before: {len(names)} inputs, identical to the prior experiment before')

"""Record normal Release build provenance; the normal inputs must match measured after."""
from datetime import datetime, timezone
import json
from pathlib import Path
from common import AUDIT, require, sha

manifest = json.loads((AUDIT / 'inputs.json').read_text())
source = Path(manifest['sourceGeneratedDirectory'])
inputs = manifest['afterSha256']
require(inputs == manifest['sourceGeneratedSha256'], 'Normal generation must match after')
require(all(sha(source / name) == expected for name, expected in inputs.items()), 'Normal build sources differ from after')
log = AUDIT / 'build-normal.log'
content = log.read_text()
require('Build succeeded.' in content and '0 Error(s)' in content, 'No successful normal build in log')
runtime = source / 'bin/Release/net8.0'
files = sorted(p for p in runtime.iterdir() if p.is_file())
require((runtime / 'Program.dll') in files, 'Normal Program.dll absent')
record = {'capturedUtc': datetime.now(timezone.utc).isoformat(), 'variant': 'normal',
    'buildLogSha256': sha(log), 'sourceSha256': inputs,
    'runtimeSha256': {p.name: sha(p) for p in files}}
(AUDIT / 'build-normal.json').write_text(json.dumps(record, indent=2) + '\n')
print(f'Recorded normal: {len(inputs)} inputs matching measured after and {len(files)} runtime files')

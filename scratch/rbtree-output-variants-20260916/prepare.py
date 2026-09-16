#!/usr/bin/env python3
"""Freeze the requested compiled output without rebuilding or editing it."""
from pathlib import Path
import hashlib
import json
import re

ROOT = Path(__file__).resolve().parent
SOURCE = ROOT.parents[1] / 'output/purust_output/Purs_Test_RBTree/src/lib.rs'
source = SOURCE.read_text()
starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
ranges = [(source.index('#[derive(Clone'), starts[0].start())]
excluded = {'R', 'B', 'E', 'T', 'describe', 'act'}
for i, match in enumerate(starts):
    end = starts[i+1].start() if i+1 < len(starts) else len(source)
    if match[1].removeprefix('Test_RBTree_') not in excluded:
        ranges.append((match.start(), end))
lines = [''] * len(source.splitlines())
for start, end in ranges:
    line = source[:start].count('\n')
    part = source[start:end].splitlines()
    lines[line:line+len(part)] = part
kernel = '\n'.join(lines) + '\n'
(ROOT / 'baseline.rs').write_text(kernel)
(ROOT / 'reference.rs').write_text(kernel.replace('crate::', 'crate::reference::'))
(ROOT / 'original-output.rs').write_text(source)
(ROOT / 'source-manifest.json').write_text(json.dumps({
    'source': str(SOURCE), 'resolved_source': str(SOURCE.resolve()),
    'source_sha256': hashlib.sha256(source.encode()).hexdigest(),
    'baseline_sha256': hashlib.sha256(kernel.encode()).hexdigest(),
    'extraction': 'Enums and all algorithm functions unchanged; unused module constructor and Effect adapters omitted.',
}, indent=2) + '\n')
print((ROOT / 'source-manifest.json').read_text())

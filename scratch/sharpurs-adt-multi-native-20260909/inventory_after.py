"""Describe actual normally generated changes; never modifies compiler output."""
import json
import re
from common import AUDIT, require, sha

manifest = json.loads((AUDIT / 'inputs.json').read_text())
declaration = re.compile(r'^(?:let(?: rec)?|and) ([A-Za-z0-9_]+)', re.MULTILINE)
changed = manifest['changedFiles']
records = []
for name in changed:
    before = (AUDIT / 'before' / name).read_text()
    after = (AUDIT / 'after' / name).read_text()
    old_names, new_names = set(declaration.findall(before)), set(declaration.findall(after))
    added = sorted(new_names - old_names)
    records.append({'file': name, 'beforeSha256': sha(AUDIT / 'before' / name),
                    'afterSha256': sha(AUDIT / 'after' / name),
                    'newNativeFunctionDeclarations': [n for n in added if n.endswith('_adt_native')],
                    'newNativeApplyDeclarations': [n for n in added if n.endswith('_adt_native_apply')],
                    'otherNewTopLevelDeclarations': [n for n in added if not n.endswith(('_adt_native', '_adt_native_apply'))],
                    'removedTopLevelDeclarations': sorted(old_names - new_names)})
require(all(record['file'].endswith('.fs') for record in records), 'Review changed non-F# build inputs')
result = {'source': 'Normal compiler generation; declaration inventory and source diffs, no patched F#.',
          'inputFiles': len(manifest['afterSha256']), 'changedFiles': len(records),
          'unchangedFiles': len(manifest['afterSha256']) - len(records),
          'newNativeFunctionDeclarations': sum(len(r['newNativeFunctionDeclarations']) for r in records),
          'files': records}
(AUDIT / 'generation-inventory.json').write_text(json.dumps(result, indent=2) + '\n')
print(json.dumps(result, indent=2))

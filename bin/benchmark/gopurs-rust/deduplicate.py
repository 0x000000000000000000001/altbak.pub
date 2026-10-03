"""APFS-copy-on-write deduplication of completed benchmark artifacts only."""
import ctypes
import hashlib
import json
import os
from pathlib import Path
import sys

out, *directories = map(Path, sys.argv[1:])
assert directories and not out.exists()
clonefile = ctypes.CDLL(None, use_errno=True).clonefile
clonefile.argtypes = [ctypes.c_char_p, ctypes.c_char_p, ctypes.c_int]
clonefile.restype = ctypes.c_int
groups = {}
records = []
for directory in directories:
    for path in sorted(directory.rglob('*')):
        if path.is_symlink() or not path.is_file() or 'target' in path.parts:
            continue
        size = path.stat().st_size
        if size < 4096:
            continue
        digest = hashlib.sha256(path.read_bytes()).hexdigest()
        key = (size, digest)
        if key not in groups:
            groups[key] = path
            continue
        source = groups[key]
        temporary = path.with_name(path.name + '.reflink-' + str(os.getpid()))
        assert not temporary.exists()
        if clonefile(os.fsencode(source), os.fsencode(temporary), 0) != 0:
            raise OSError(ctypes.get_errno(), f'clonefile {source} -> {temporary}')
        assert hashlib.sha256(temporary.read_bytes()).hexdigest() == digest
        os.chmod(temporary, path.stat().st_mode & 0o777)
        os.replace(temporary, path)
        records.append({'path': str(path), 'source': str(source), 'bytes': size, 'sha256': digest})
result = {'status': 'passed', 'method': 'independent APFS copy-on-write files; no hard links',
          'files': len(records), 'logical_duplicate_bytes': sum(r['bytes'] for r in records),
          'records': records}
out.write_text(json.dumps(result, indent=2) + '\n')
print(json.dumps({k: v for k, v in result.items() if k != 'records'}, indent=2))

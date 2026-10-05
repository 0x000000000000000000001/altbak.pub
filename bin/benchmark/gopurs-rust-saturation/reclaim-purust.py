"""Reclaim regenerable Purust caches, retaining every executable and diagnostic."""
import hashlib
import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
from datetime import datetime, timezone

htdocs = Path(__file__).resolve().parents[4]
archive = Path(sys.argv[1]).resolve()
archive.mkdir(parents=True, exist_ok=True)
report = archive / 'purust-reclamation.json'
assert not report.exists(), 'Keep previous reclamation evidence'
purust = htdocs / 'purust'
old = htdocs / 'altbak.pub/var/benchmark/purust-packages-20261004'
builds = [
    old / 'finish2-bootstrap/purust-native-build-8ffnVb',
    old / 'primed-bootstrap2/purust-native-build-Sv70Rv',
    old / 'unicode-bootstrap/purust-native-build-XEkFUG',
]
targets = [build / stage / 'target' for build in builds for stage in ['rust', 'rust-stage2']]

def sha(path):
    digest = hashlib.sha256()
    with path.open('rb') as stream:
        for block in iter(lambda: stream.read(1024 * 1024), b''):
            digest.update(block)
    return digest.hexdigest()

def files(root):
    for directory, dirs, names in os.walk(root, followlinks=False):
        dirs[:] = sorted(d for d in dirs if not (Path(directory) / d).is_symlink())
        for name in sorted(names):
            path = Path(directory) / name
            if not path.is_symlink():
                yield path

def record(path):
    return {'path': str(path), 'bytes': path.stat().st_size, 'sha256': sha(path)}

repos = sorted(p for p in purust.iterdir() if (p / '.git').is_dir())
statuses = {str(p): subprocess.check_output(['git', '-C', str(p), 'status', '--porcelain'], text=True) for p in repos}
protected = []
for repo in repos:
    tracked = subprocess.check_output(['git', '-C', str(repo), 'ls-files', '-z']).decode().split('\0')
    protected.extend(record(repo / name) for name in tracked if name and (repo / name).is_file() and not (repo / name).is_symlink())
protected.extend(record(p) for p in (purust / 'purust/bin').iterdir() if p.is_file() and not p.is_symlink())
removed = []
for target in targets:
    assert (target / '.rustc_info.json').exists()
    assert (target / 'release/purust_output').is_file()
    for path in files(target):
        # Delete compiled intermediate libraries/objects only. Even Cargo's
        # build-script executables stay at their original diagnostic paths.
        if path.suffix in ['.rlib', '.rmeta', '.o', '.bc'] and not os.access(path, os.X_OK):
            removed.append(record(path))
        elif os.access(path, os.X_OK):
            protected.append(record(path))
caches = []
for repo in repos:
    cache = repo / '.purmeta'
    if not cache.is_dir() or cache.is_symlink():
        continue
    if subprocess.check_output(['git', '-C', str(repo), 'ls-files', '--', '.purmeta']):
        continue
    cache_files = list(files(cache))
    assert all(p.suffix == '.purmeta' or p.name == '.DS_Store' for p in cache_files), str(cache)
    caches.append(cache)
    removed.extend(record(p) for p in cache_files)
state = {'status': 'prepared', 'started_at': datetime.now(timezone.utc).isoformat(),
         'available_before': shutil.disk_usage(htdocs).free, 'removed': removed,
         'protected': protected, 'git_status_before': statuses,
         'logical_bytes_removed': sum(p['bytes'] for p in removed)}
report.write_text(json.dumps(state, indent=2) + '\n')
for item in removed:
    Path(item['path']).unlink()
for cache in caches:
    for directory, dirs, names in os.walk(cache, topdown=False):
        if not dirs and not names:
            Path(directory).rmdir()
for item in protected:
    assert sha(Path(item['path'])) == item['sha256'], item['path']
for repo in repos:
    assert subprocess.check_output(['git', '-C', str(repo), 'status', '--porcelain'], text=True) == statuses[str(repo)]
state.update(status='passed', available_after=shutil.disk_usage(htdocs).free,
             finished_at=datetime.now(timezone.utc).isoformat())
report.write_text(json.dumps(state, indent=2) + '\n')
print(json.dumps({k: v for k, v in state.items() if k not in ['removed', 'protected', 'git_status_before']}, indent=2))

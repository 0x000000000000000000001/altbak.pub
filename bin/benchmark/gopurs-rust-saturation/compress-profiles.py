"""Losslessly archive completed raw profiles, checking bytes before reclamation."""
import gzip
import hashlib
import json
from pathlib import Path
import shutil
import sys
from datetime import datetime, timezone

archive = Path(sys.argv[1]).resolve()
label = sys.argv[2]
out = archive / (label + '.json')
assert not out.exists()
result = {'status': 'running', 'started_at': datetime.now(timezone.utc).isoformat(),
          'tool_sha256': hashlib.sha256(Path(__file__).read_bytes()).hexdigest(),
          'available_before': shutil.disk_usage(archive).free, 'files': []}

def save():
    out.write_text(json.dumps(result, indent=2) + '\n')

def digest(stream):
    sha, size = hashlib.sha256(), 0
    while True:
        block = stream.read(1024 * 1024)
        if not block:
            return sha.hexdigest(), size
        sha.update(block)
        size += len(block)

save()
try:
    for source in sorted((archive / 'profiles').glob('*/sample.txt')):
        summary = source.parent / 'summary.json'
        assert json.loads(summary.read_text())['status'] == 'passed'
        destination = source.with_suffix(source.suffix + '.gz')
        assert not destination.exists()
        with source.open('rb') as raw:
            original_hash, original_size = digest(raw)
        with source.open('rb') as raw, gzip.open(destination, 'xb', compresslevel=1) as compressed:
            shutil.copyfileobj(raw, compressed, 1024 * 1024)
        with gzip.open(destination, 'rb') as restored:
            assert digest(restored) == (original_hash, original_size)
        with destination.open('rb') as compressed:
            compressed_hash, compressed_size = digest(compressed)
        entry = {'original': str(source), 'original_sha256': original_hash, 'original_bytes': original_size,
                 'compressed': str(destination), 'compressed_sha256': compressed_hash,
                 'compressed_bytes': compressed_size, 'restored_bytes_verified': True,
                 'summary_sha256': hashlib.sha256(summary.read_bytes()).hexdigest()}
        result['files'].append(entry)
        save()
        source.unlink()
        entry['original_removed'] = True
        save()
    result['logical_bytes_reclaimed'] = sum(item['original_bytes'] - item['compressed_bytes'] for item in result['files'])
    result['available_after'] = shutil.disk_usage(archive).free
    result['status'] = 'passed'
    result['finished_at'] = datetime.now(timezone.utc).isoformat()
    save()
    print(json.dumps({key: value for key, value in result.items() if key != 'files'}, indent=2))
except Exception as error:
    result['status'] = 'failed'
    result['failure'] = repr(error)
    save()
    raise

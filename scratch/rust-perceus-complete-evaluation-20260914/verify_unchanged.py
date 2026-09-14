"""Check product sources and generated output against the start-of-study snapshot."""
from pathlib import Path
import hashlib
import json
import subprocess
import full_runner

HERE = Path(__file__).resolve().parent


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def check():
    before = json.loads((HERE / 'purust-before.json').read_text())
    root = Path(before['repository'])
    current = {name: sha(root / name) for name in before['sha256']}
    assert current == before['sha256'], 'Purust source changed during study'
    head = subprocess.check_output(['git', 'rev-parse', 'HEAD'], cwd=root, text=True).strip()
    status = subprocess.check_output(['git', 'status', '--porcelain'], cwd=root, text=True).strip()
    assert head == before['head'] and status == before['status']
    frozen = full_runner.verify_frozen()
    assert full_runner.inputs(full_runner.LIVE) == frozen['inputs'], 'Live generated output changed'
    result = {'purust_files_verified': len(current), 'purust_head': head,
              'purust_status': status, 'live_generated_files_verified': len(frozen['inputs']),
              'purust_unchanged': True, 'live_generated_unchanged': True,
              'snapshot_sha256': sha(HERE / 'full-runner-snapshot.json'),
              'before_sha256': sha(HERE / 'purust-before.json')}
    (HERE / 'no-product-changes.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps(result, indent=2))


if __name__ == '__main__':
    check()

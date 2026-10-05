"""Untimed all-thread sampling of one frozen, exact-output-checked case."""
import argparse
import collections
from datetime import datetime, timezone
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import subprocess
import time

parser = argparse.ArgumentParser()
parser.add_argument('archive', type=Path)
parser.add_argument('label')
parser.add_argument('case')
parser.add_argument('binary', type=Path)
parser.add_argument('--phase', choices=['all', 'prepare', 'optimize'], default='all')
args = parser.parse_args()
archive, binary = args.archive.resolve(), args.binary.resolve()
out = archive / 'profiles' / args.label
out.mkdir(parents=True, exist_ok=False)
shutil.copy2(__file__, out / 'profile.py')
item = json.loads((archive / 'cases' / args.case / 'definition.json').read_text())
cwd = Path(item['input'])

def sha(path):
    digest = hashlib.sha256()
    with path.open('rb') as stream:
        for block in iter(lambda: stream.read(1024 * 1024), b''):
            digest.update(block)
    return digest.hexdigest()

def clean():
    protected = {p['path'] for p in item['input_manifest']}
    for p in sorted((cwd / 'output').rglob('*')):
        if p.is_file() and (p.suffix == '.go' or p.name in ['go.mod', 'go.sum']):
            assert str(p.relative_to(cwd)) not in protected
            p.unlink()
    for name in ['.purmeta', '.cache']:
        if (cwd / name).exists():
            shutil.rmtree(cwd / name)

def verify_inputs():
    for p in item['input_manifest']:
        assert sha(cwd / p['path']) == p['sha256'], p['path']
    for p in item['sibling_inputs']:
        assert sha(cwd.parent / p['path']) == p['sha256'], p['path']

clean()
verify_inputs()
env = {k: v for k, v in os.environ.items()
       if not re.match(r'^(PURUST_|GOPURS_|NODE_|RUST|CARGO|GOGC$|GOMEMLIMIT$|GOMAXPROCS$|GODEBUG$)', k)}
env.update({key: '8' for key in ['GOPURS_JOBS', 'GOPURS_PREPARE_JOBS', 'GOPURS_PBO_JOBS', 'GOPURS_EMIT_JOBS']})
env.update(GOPURS_PIPELINE='1', GOWORK='off')
command = [str(binary), *item['invocation']]
started = time.monotonic()
timeline = []
with (out / 'stdout').open('w') as stdout, (out / 'stderr').open('w') as stderr:
    child = subprocess.Popen(command, cwd=cwd, env=env, stdout=stdout, stderr=subprocess.PIPE, text=True)
    if args.phase != 'all':
        marker = '[gopurs] ' + ('load TAST + sort:' if args.phase == 'prepare' else 'prepare + monomorphize:')
        for line in child.stderr:
            stderr.write(line)
            timeline.append({'offset_ms': (time.monotonic() - started) * 1000, 'line': line.rstrip()})
            if line.startswith(marker):
                break
    sample = subprocess.Popen(['sample', str(child.pid), '60', '1', '-file', str(out / 'sample.txt')],
                              stdout=subprocess.PIPE, stderr=subprocess.STDOUT, text=True)
    for line in child.stderr:
        stderr.write(line)
        timeline.append({'offset_ms': (time.monotonic() - started) * 1000, 'line': line.rstrip()})
    status = child.wait()
    sample_log, _ = sample.communicate()
    (out / 'sample.log').write_text(sample_log)
assert status == 0, f'compiler failed: {status}'
actual = []
for path in sorted((cwd / 'output').rglob('*')):
    if path.is_file() and (path.suffix == '.go' or path.name == 'go.mod'):
        actual.append({'path': str(path.relative_to(cwd / 'output')), 'bytes': path.stat().st_size, 'sha256': sha(path)})
assert actual == item['oracle']['files'], 'Profiled generation differs from the frozen oracle'
for p in actual:
    assert (cwd / 'output' / p['path']).read_bytes() == (Path(item['oracle']['directory']) / p['path']).read_bytes()
(out / 'generated-manifest.json').write_text(json.dumps(actual, indent=2) + '\n')
clean()
verify_inputs()

stack = []
leaves, families, allocation_callers, functions = (collections.Counter() for _ in range(4))
patterns = {
    'decoding': r'Optimizer_CoreFn_Json|Optimizer_CoreFn_TypeTable|Purs_Data_Argonaut',
    'Map/Set': r'Purs_Data_Map_Internal::|Purs_Data_Set::',
    'PBO': r'Purs_PureScript_Backend_Optimizer_',
    'monomorphization': r'Optimizer_Monomorphize::',
    'type substitution': r'Optimizer_TypeSubstitution::|substituteNeutralTypes|instantiateNeutralType',
    'gopurs': r'Purs_Gopurs_',
    'allocation': r'mi_malloc|mi_free|malloc|__rust_alloc|__rust_dealloc',
    'String clone': r'5alloc6string.*5clone|String.*clone',
    'Value clone': r'purust_core..Value.*Clone.*clone',
    'reference counts': r'Arc.*clone|Arc.*drop|arc.*drop_slow|perceus',
}

def finish():
    _, count, name, children = stack.pop()
    samples = count - children
    path = '\n'.join(frame[2] for frame in stack) + '\n' + name
    if samples <= 0 or 'Purs_' not in path:
        return
    if re.search(r'__ulock_wait|__psynch_cvwait|kevent|mach_msg.*trap|thread_switch|__psynch_mutexwait', name):
        return
    leaves[name] += samples
    for label, pattern in patterns.items():
        if re.search(pattern, path):
            families[label] += samples
    for frame in reversed([*stack, [0, 0, name, 0]]):
        if 'Purs_' in frame[2]:
            functions[frame[2]] += samples
            break
    if re.search(patterns['allocation'], name):
        for frame in reversed(stack):
            if 'Purs_' in frame[2]:
                allocation_callers[frame[2]] += samples
                break

profile = (out / 'sample.txt').open()
for line in profile:
    if line.startswith('Total number in stack'):
        break
    match = re.match(r'^([ +!:|]*)(\d+) (.+)$', line)
    if not match:
        continue
    prefix, count, name = match.groups()
    depth, count, name = len(prefix), int(count), name.split('  (in ')[0]
    while stack and stack[-1][0] >= depth:
        finish()
    if stack:
        stack[-1][3] += count
    stack.append([depth, count, name, 0])
while stack:
    finish()
profile.close()
result = {'status': 'passed', 'checked_at': datetime.now(timezone.utc).isoformat(),
          'command': command, 'cwd': str(cwd), 'phase': args.phase, 'binary_sha256': sha(binary),
          'capture_window': 'From the requested start marker until process exit; later phases are included',
          'tool_sha256': sha(out / 'profile.py'),
          'sample': {'path': str(out / 'sample.txt'), 'sha256': sha(out / 'sample.txt'),
                     'bytes': (out / 'sample.txt').stat().st_size},
          'identical_go_files': len(actual), 'timeline': timeline,
          'metric': 'All-thread wall-clock samples excluding known blocking leaves; inclusive families overlap and are not CPU percentages',
          'samples': sum(leaves.values()), 'inclusive_families': dict(families),
          'top_exclusive': leaves.most_common(50), 'nearest_purescript_frames': functions.most_common(60),
          'allocation_callers': allocation_callers.most_common(40)}
(out / 'summary.json').write_text(json.dumps(result, indent=2) + '\n')
print(json.dumps({k: v for k, v in result.items() if k != 'timeline'}, indent=2))

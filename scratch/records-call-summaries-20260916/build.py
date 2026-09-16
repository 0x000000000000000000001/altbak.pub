#!/usr/bin/env python3
from pathlib import Path
import hashlib
import json
import subprocess

BASE = Path(__file__).resolve().parent
ROOT = BASE.parents[1]
DEPS = ROOT / 'output/purust_output/target/release/deps'
def one(pattern):
    found = sorted(DEPS.glob(pattern))
    assert len(found) == 1, found
    return found[0]

libraries = {name: one(f'lib{name}-*.rlib') for name in ['purust_core', 'perceus_ptr', 'mimalloc']}
native = sorted((DEPS.parent/'build').glob('libmimalloc-sys-*/out/libmimalloc.a'))
assert len(native) == 1, native
common = ['rustc', '--edition=2021', '-C', 'opt-level=3', '-C', 'lto=off', '-C', 'embed-bitcode=no',
          '-L', f'dependency={DEPS}']
for name, lib in libraries.items():
    common += ['--extern', f'{name}={lib}']
common += ['-L', f'native={native[0].parent}']
commands = []
for counted in [False, True]:
    suffix = '-count' if counted else ''
    cfg = ['--cfg', 'counted'] if counted else []
    lib = BASE/f'libcall_fields{suffix}.rlib'
    commands.append(common + cfg + ['--crate-name', 'call_fields', '--crate-type=rlib', str(BASE/'callee.rs'), '-o', str(lib)])
    commands.append(common + cfg + ['--extern', f'call_fields={lib}', str(BASE/'probe.rs'), '-o', str(BASE/f'probe{suffix}')])
for cmd in commands:
    result = subprocess.run(cmd, capture_output=True, text=True, timeout=60)
    if result.returncode:
        print(result.stdout, result.stderr)
        raise SystemExit(result.returncode)
validated = subprocess.check_output([str(BASE/'probe'), 'validate'], text=True, timeout=60)
(BASE/'validation.txt').write_text(validated)
counts = []
for stride in [0,100,1]:
    for name in ['record', 'materialized', 'scalar']:
        raw = subprocess.check_output([str(BASE/'probe-count'), 'count', name, str(stride)], text=True, timeout=60)
        counts.append(json.loads(raw))
(BASE/'counts.json').write_text(json.dumps(counts, indent=2)+'\n')
manifest = {'commands': commands, 'source': str(ROOT/'output/purust_output/Purs_Test_Records/src/lib.rs'),
            'sha256': {str(p): hashlib.sha256(p.read_bytes()).hexdigest() for p in [BASE/'callee.rs', BASE/'probe.rs', BASE/'probe', *libraries.values()]},
            'profile': 'opt-level=3, lto=off, embed-bitcode=no, native MiMalloc',
            'timing_instrumented': False}
(BASE/'build-manifest.json').write_text(json.dumps(manifest, indent=2)+'\n')
print(validated, json.dumps(counts, indent=2))

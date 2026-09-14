from pathlib import Path
import json
import re
import shutil
import subprocess
from list_probe import HERE, OUT, FROZEN, prepare

prepare()
build = HERE / 'build/list-validation'
crate = build / 'list-crate'
(crate / 'src').mkdir(parents=True, exist_ok=True)
(build / 'src').mkdir(exist_ok=True)
original_dir = FROZEN / 'Purs_Test_ListOps'
manifest = (original_dir / 'Cargo.toml').read_text()
manifest = re.sub(r'path\s*=\s*"([^"]+)"', lambda m: 'path = ' + json.dumps(str((original_dir / m[1]).resolve())), manifest)
(crate / 'Cargo.toml').write_text(manifest)
(build / 'Cargo.toml').write_text('''[workspace]
[package]
name = "list_perceus_validation"
version = "0.1.0"
edition = "2021"
[profile.release]
opt-level = 1
debug = true
[dependencies]
''' + 'Purs_Test_ListOps = { path = "list-crate" }\npurust_core = { path = ' + json.dumps(str(FROZEN / 'purust_core')) + ' }\n')
shutil.copy2(OUT / 'checks.rs', build / 'src/main.rs')
shutil.copy2(FROZEN / 'Cargo.lock', build / 'Cargo.lock')
results = {}
for name in ['before', 'consumed']:
    shutil.copy2(OUT / f'ListOps-{name}-counted.rs', crate / 'src/lib.rs')
    with (OUT / f'validation-{name}.log').open('w') as log:
        subprocess.run(['cargo', 'clean', '--release', '-p', 'Purs_Test_ListOps', '-p', 'list_perceus_validation'], cwd=build, stdout=log, stderr=subprocess.STDOUT, check=True)
        subprocess.run(['cargo', 'run', '--release', '--offline', '--quiet'], cwd=build, stdout=log, stderr=subprocess.STDOUT, check=True)
    lines = (OUT / f'validation-{name}.log').read_text().splitlines()
    result = json.loads(next(line for line in reversed(lines) if line.startswith('{')))
    results[name] = result
    print(name, result, flush=True)
(OUT / 'validation.json').write_text(json.dumps(results, indent=2) + '\n')

#!/usr/bin/env python3
"""Build hand-written F# with stock Fable, then measure the generated Rust.

This reference translates the PureScript FP kernels to typed F#. It does not use
PureScript, sharpurs, the compatibility runtime, or the local Fable emitter patch.
"""
from __future__ import annotations

import argparse
from datetime import datetime, timezone
import math
import os
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import xml.etree.ElementTree as ET

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parent.parent
ASSETS = ROOT / 'tmp/fable_rust'
sys.path.insert(0, str(ASSETS))
from build_benchmark import Build, NUGET_CONFIG, sha256, tree_hashes


def build_native(args, directory):
    root = directory / 'build'
    root.mkdir()
    source = args.source.resolve()
    package = args.fable_package.resolve()
    dotnet = args.dotnet.resolve()
    package_tools = package / 'tools/net10.0/any'
    harness_source = ASSETS / 'bench_fable.rs'
    environment = os.environ.copy()
    environment.update({
        'DOTNET_CLI_HOME': str(root / 'dotnet-home'),
        'DOTNET_ROOT': str(dotnet.parent),
        'DOTNET_SKIP_FIRST_TIME_EXPERIENCE': '1',
        'DOTNET_CLI_TELEMETRY_OPTOUT': '1',
        'DOTNET_NOLOGO': '1',
        'DOTNET_GENERATE_ASPNET_CERTIFICATE': 'false',
        'CARGO_TARGET_DIR': str(root / 'target'),
        'PATH': str(dotnet.parent) + os.pathsep + environment.get('PATH', ''),
    })
    manifest = {
        'status': 'building', 'started_utc': datetime.now(timezone.utc).isoformat(),
        'pipeline': 'hand-written F# -> stock Fable -> Rust',
        'fable_patched': False, 'sharpurs_used': False,
        'source': str(source), 'source_sha256': sha256(source),
        'orchestrator_sha256': sha256(Path(__file__).resolve()),
        'build_helper_sha256': sha256(ASSETS / 'build_benchmark.py'),
        'harness_sha256': sha256(harness_source),
        'fable_package': str(package),
        'package_tools_sha256': tree_hashes(package_tools),
        'package_runtime_sha256': tree_hashes(package / 'fable-library-rust'),
        'commands': [], 'benchmark_executed': False,
        'reference_source_sha256': tree_hashes(ROOT / 'src/Test', {'.purs'}),
    }
    build = Build(root, environment, manifest)
    for name, command in (
        ('dotnet', [dotnet, '--version']),
        ('rustc', ['rustc', '--version']),
        ('cargo', ['cargo', '--version']),
    ):
        build.run('00-version-' + name, command, root, 30)
    fsharp = root / 'fsharp'
    fsharp.mkdir()
    copied = fsharp / 'NativeBench.fs'
    shutil.copy2(source, copied)
    if sha256(copied) != manifest['source_sha256']:
        raise RuntimeError('F# source changed while preparing build')
    (fsharp / 'NuGet.Config').write_text(NUGET_CONFIG)
    project = ET.Element('Project', Sdk='Microsoft.NET.Sdk')
    properties = ET.SubElement(project, 'PropertyGroup')
    for key, value in {
        'TargetFramework': 'net10.0',
        'EnableDefaultCompileItems': 'false',
        'DisableImplicitFSharpCoreReference': 'true',
    }.items():
        ET.SubElement(properties, key).text = value
    items = ET.SubElement(project, 'ItemGroup')
    ET.SubElement(items, 'Compile', Include='NativeBench.fs')
    reference = ET.SubElement(items, 'Reference', Include='FSharp.Core')
    ET.SubElement(reference, 'HintPath').text = str(package_tools / 'FSharp.Core.dll')
    ET.indent(project)
    fsproj = fsharp / 'NativeBench.fsproj'
    ET.ElementTree(project).write(fsproj, encoding='unicode')
    rust = root / 'rust'
    build.run('01-fable-rust', [dotnet, package_tools / 'fable.dll', fsproj,
              '--noCache', '--lang', 'Rust', '--outDir', rust], root, 180)
    harness = harness_source.read_text().replace('BenchFable', 'NativeBench')
    generated = (rust / 'NativeBench.rs').read_text()
    allocator = '#[global_allocator]\nstatic FABLE_BENCH_ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;\n'
    (rust / 'main.rs').write_text(generated.rstrip() + '\n\n' + allocator + '\n' + harness)
    (root / 'Cargo.toml').write_text('''[package]
name = "altbak-fable-native"
version = "0.1.0"
edition = "2024"

[[bin]]
name = "altbak-fable-native"
path = "rust/main.rs"

[dependencies]
fable_library_rust = { path = "rust/fable_modules/fable-library-rust", default-features = false }
mimalloc = "0.1.32"

[profile.release]
opt-level = 3
debug = false
lto = "thin"
''')
    build.run('02-cargo-release', ['cargo', 'build', '--offline', '--release',
              '--manifest-path', root / 'Cargo.toml'], root, 600)
    binary = root / 'target/release/altbak-fable-native'
    manifest.update({
        'status': 'built', 'executable': str(binary),
        'executable_sha256': sha256(binary),
        'cargo_lock_sha256': sha256(root / 'Cargo.lock'),
        'generated_rust_sha256': tree_hashes(rust, {'.rs', '.toml'}),
    })
    build.save()
    try:
        build.run('03-check-results', [binary, '--check-only'], root, args.timeout)
    except RuntimeError as error:
        manifest.update({'status': 'failed', 'error': str(error)})
        build.save()
        raise
    manifest['correctness_validated_cases'] = 14
    build.save()
    return binary, root / 'manifest.json'


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--update-readme', action='store_true')
    parser.add_argument('--build-only', action='store_true')
    parser.add_argument('--dotnet', type=Path,
                        default=os.environ.get('FABLE_DOTNET') or shutil.which('dotnet'))
    parser.add_argument('--fable-package', type=Path,
                        default=Path.home() / '.nuget/packages/fable/5.17.2')
    parser.add_argument('--source', type=Path, default=ASSETS / 'NativeBench.fs')
    parser.add_argument('--timeout', type=float, default=180,
                        help='Seconds allowed per measurement process')
    args = parser.parse_args()
    if args.build_only and args.update_readme:
        parser.error('--build-only cannot update measurements')
    if not math.isfinite(args.timeout) or args.timeout <= 0:
        parser.error('--timeout must be a positive finite number')
    if args.dotnet is None:
        parser.error('Supply --dotnet PATH to a .NET 10 executable, or set FABLE_DOTNET')
    for path in (args.dotnet, args.source,
                 args.fable_package / 'tools/net10.0/any/fable.dll',
                 args.fable_package / 'tools/net10.0/any/FSharp.Core.dll'):
        if not path.is_file():
            parser.error(f'Required file not found: {path}')
    for name in ('rustc', 'cargo'):
        if shutil.which(name) is None:
            parser.error(f'Executable not found on PATH: {name}')
    parent = ROOT / 'var/benchmark/fable-native-rust'
    parent.mkdir(parents=True, exist_ok=True)
    stamp = datetime.now(timezone.utc).strftime('%Y%m%dT%H%M%SZ-')
    directory = Path(tempfile.mkdtemp(prefix=stamp, dir=parent))
    print(f'Builds, validated results and provenance: {directory}', flush=True)
    binary, manifest = build_native(args, directory)
    if args.build_only:
        print('Build manifest: ' + str(manifest), flush=True)
        return
    command = [sys.executable, str(ASSETS / 'collect_benchmark.py'),
               '--variant', 'native-fsharp', '--binary', str(binary),
               '--output-dir', str(directory / 'measurements'),
               '--project-root', str(ROOT), '--build-manifest', str(manifest),
               '--allocator', 'mimalloc', '--timeout', str(args.timeout)]
    if args.update_readme:
        command += ['--readme', str(ROOT / 'README.md')]
    subprocess.run(command, cwd=ROOT, check=True)


if __name__ == '__main__':
    main()

#!/usr/bin/env python3
"""Build and measure sharpurs' numeric kernels through a locally patched Fable.

Requires .NET 10, the Fable 5.17.2 NuGet package, a matching Fable source checkout,
and cached Cargo dependencies. The generated sharpurs Main directory is an
explicit input; this command preserves its source files and records their hashes.
"""
import argparse
from datetime import datetime, timezone
import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parent.parent
ASSETS = ROOT / 'tmp/fable_rust'


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--update-readme', action='store_true')
    parser.add_argument('--build-only', action='store_true')
    parser.add_argument('--dotnet', type=Path, default=os.environ.get('FABLE_DOTNET') or shutil.which('dotnet'),
                        help='Path to a .NET 10 dotnet executable (or set FABLE_DOTNET)')
    parser.add_argument('--fable-package', type=Path, default=Path.home()/'.nuget/packages/fable/5.17.2')
    parser.add_argument('--fable-source', type=Path, default=ROOT.parent/'Fable')
    parser.add_argument('--source', type=Path, default=ROOT/'run/bak/sharp/output/Main')
    parser.add_argument('--timeout', type=float, default=180, help='Seconds allowed per measurement process')
    args = parser.parse_args()
    if args.build_only and args.update_readme:
        parser.error('--build-only cannot update measurements')
    if args.dotnet is None:
        parser.error('Supply --dotnet PATH to a .NET 10 executable, or set FABLE_DOTNET')
    parent = ROOT/'var/benchmark/fable-rust'
    parent.mkdir(parents=True, exist_ok=True)
    stamp = datetime.now(timezone.utc).strftime('%Y%m%dT%H%M%SZ-')
    directory = Path(tempfile.mkdtemp(prefix=stamp, dir=parent))
    print(f'Builds, validated results and provenance: {directory}', flush=True)
    subprocess.run([
        sys.executable, str(ASSETS/'build_benchmark.py'),
        '--source', str(args.source), '--build-dir', str(directory/'build'),
        '--dotnet', str(args.dotnet), '--fable-package', str(args.fable_package),
        '--fable-source', str(args.fable_source),
    ], cwd=ROOT, check=True)
    manifest = directory/'build/manifest.json'
    binary = json.loads(manifest.read_text())['executable']
    # Verify values before spending time on the calibrated benchmark protocol.
    with (directory/'check.log').open('w') as log:
        subprocess.run([binary, '--check-only'], cwd=ROOT, stdout=log,
                       stderr=subprocess.STDOUT, timeout=args.timeout, check=True)
    if args.build_only:
        print('14 kernels verified; build manifest: ' + str(manifest), flush=True)
        return
    print('14 kernels executed successfully; starting measurements.', flush=True)
    command = [
        sys.executable, str(ASSETS/'collect_benchmark.py'), '--binary', binary,
        '--output-dir', str(directory/'measurements'), '--project-root', str(ROOT),
        '--build-manifest', str(manifest), '--allocator', 'mimalloc',
        '--timeout', str(args.timeout),
    ]
    if args.update_readme:
        command += ['--readme', str(ROOT/'README.md')]
    subprocess.run(command, cwd=ROOT, check=True)
    print(f'Results: {directory / "measurements/results.json"}', flush=True)


if __name__ == '__main__':
    main()

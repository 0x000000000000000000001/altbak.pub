#!/usr/bin/env python3
"""Build one Rust benchmark mode and measure three validated processes.

The build uses the isolated workspace machinery in bin/native/driver.py. The
measurement reuses bin/benchmark/measure.py: a frozen plan, three independent
processes, validated outputs and the median per row. The printed total is the
sum of the fourteen cell medians, exactly like the published tables.
`--update-readme` writes that same result into the Rust table column, keeping
the existing table format and recomputing the /C ratio from the C (reference)
total.

The generated Purust Cargo.toml enables `lto = "thin"`; this runner only
overrides `opt-level` and `debug`, so every Rust mode measured here is built
with thin LTO. Modes `ffi`/`fficc` update the Rust table; mode `x` updates the
Extended Results table. Column 0 (`pure`) is published by
tmp/run_purust_benchmark.py.
"""
import argparse
import hashlib
import json
import math
import os
import re
import subprocess
import sys
from datetime import datetime, timezone
from pathlib import Path

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[2]
sys.path.insert(0, str(ROOT / 'bin/benchmark'))
from publish import EXTENDED_ROWS, ROWS  # noqa: E402
from validate import expected_cases  # noqa: E402

MODES = {'ffi': ['--ffi'], 'fficc': ['--fficc'], 'x': ['--x']}
# Keys order from publish.py GROUPS: Rust = [purust, sharpurs, fable-native, ffi, fficc],
# Extended Results = [js-x, es-x, go-x, rust-x].
COLUMNS = {'ffi': 3, 'fficc': 4, 'x': 3}
SECTIONS = {'ffi': 'Rust', 'fficc': 'Rust', 'x': 'Extended Results'}
README = ROOT / 'README.md'

CELL = re.compile(r'~\s*[0-9.]+ μs')
TOTAL = re.compile(r'~\s*[0-9.]+ ms')
RATIO = re.compile(r'/C = [0-9.]+x')


def digest(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def stamp():
    return datetime.now(timezone.utc).strftime('%Y%m%dT%H%M%SZ')


def section(text, title):
    match = re.search(r'#### ' + re.escape(title) + r'\n(.*?)(?=\n#### |\n### |\n## |\n> |\Z)', text, re.S)
    if not match:
        raise ValueError('Missing README section: ' + title)
    return match


def toolchain():
    backend = ROOT.parent / 'purust/purust'
    def git(*arguments):
        result = subprocess.run(['git', '-C', str(backend), *arguments], capture_output=True, text=True)
        return result.stdout.strip() if result.returncode == 0 else ''
    def bundle(name):
        path = backend / 'bin' / name
        return digest(path) if path.is_file() else 'missing'
    return {
        'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
        'cargo': subprocess.check_output(['cargo', '--version'], text=True).strip(),
        'purust_bundle_sha256': bundle('purust.js'),
        'purust_native_sha256': bundle('purust-native'),
        'purust_head': git('rev-parse', 'HEAD'),
        'purust_dirty': bool(git('status', '--porcelain')),
    }


def build(mode, build_dir, clean):
    # Build through bin/rust/run so the driver fingerprints its own __file__
    # exactly like the --run-only plan command below (same path spelling).
    command = [str(ROOT / 'bin/rust/run'), *MODES[mode],
               '--build-only', '--build-dir', str(build_dir)]
    if clean:
        command.append('--clean')
    print('Building Rust ' + mode + ' in ' + str(build_dir), flush=True)
    subprocess.run(command, check=True, env={**os.environ,
                   'CARGO_PROFILE_RELEASE_OPT_LEVEL': '3',
                   'CARGO_PROFILE_RELEASE_DEBUG': 'false'})


def plan(mode, build_dir):
    manifest = build_dir / 'manifest.json'
    binary = build_dir / 'benchmark'
    return {
        'key': 'rust-' + mode,
        'mode': mode,
        'command': ['env', 'CARGO_PROFILE_RELEASE_DEBUG=false', 'CARGO_PROFILE_RELEASE_OPT_LEVEL=3',
                    str(ROOT / 'bin/rust/run'), '--run-only', *MODES[mode],
                    '--build-dir', str(build_dir)],
        'cwd': str(ROOT),
        'artifact': str(build_dir),
        'manifest': str(manifest),
        'manifest_sha256': digest(manifest),
        'payload_sha256': {str(binary): digest(binary)},
    }


def measure(mode, build_dir, campaign_dir):
    if campaign_dir.exists():
        raise ValueError('Campaign directory already exists: ' + str(campaign_dir))
    plan_path = campaign_dir.with_suffix('.plan.json')
    plan_path.write_text(json.dumps([plan(mode, build_dir)], indent=2) + '\n')
    print('Plan: ' + str(plan_path) + '\nCampaign: ' + str(campaign_dir), flush=True)
    subprocess.run([sys.executable, '-B', str(ROOT / 'bin/benchmark/measure.py'),
                    '--plan', str(plan_path), '--output', str(campaign_dir)], check=True)
    return json.loads((campaign_dir / 'results.json').read_text())['rust-' + mode]


def verify(result, mode):
    expected = len(expected_cases(mode))
    if len(result['times_us']) != expected or not result['values_validated']:
        raise ValueError('Campaign did not validate every benchmark result')
    total = sum(result['times_us']) / 1000.0
    if not math.isclose(total, result['total_ms'], rel_tol=1e-12, abs_tol=1e-12):
        raise ValueError('Campaign total is not the sum of the median cells')
    if len(result.get('process_runs', [])) != 3:
        raise ValueError('A campaign needs three recorded processes')
    return total


def published(mode):
    """Return the README total for one Rust column, or None when unreadable."""
    try:
        rust = section(README.read_text(), SECTIONS[mode])
        total = None
        for line in rust.group(1).splitlines():
            if line.split('|', 1)[0].strip() == '**Total Execution Time**':
                cells = line.split('|')
                match = TOTAL.search(cells[COLUMNS[mode] + 1])
                total = float(match.group(0).split('~')[1].split('ms')[0]) if match else None
        return total
    except (OSError, ValueError):
        return None


def update_readme(mode, result):
    text = README.read_text()
    original = text
    rows = EXTENDED_ROWS if mode == 'x' else ROWS
    rust = section(text, SECTIONS[mode])
    lines = rust.group(1).splitlines(keepends=True)
    column = COLUMNS[mode]
    seen = 0
    for index, line in enumerate(lines):
        newline = '\n' if line.endswith('\n') else ''
        body = line[:-1] if newline else line
        label = body.split('|', 1)[0].strip()
        if label in rows:
            if seen >= len(result['times_us']):
                raise ValueError('More README rows than measured cells')
            value = result['times_us'][seen]
            seen += 1
            cells = body.split('|')
            if len(cells) <= column + 1:
                raise ValueError('Missing timing cell for ' + label)
            cells[column + 1], count = CELL.subn(f'~ {value:.2f} μs', cells[column + 1], count=1)
            if count != 1:
                raise ValueError('Missing timing cell for ' + label)
            lines[index] = '|'.join(cells) + newline
        elif label == '**Total Execution Time**':
            cells = body.split('|')
            if len(cells) <= column + 1:
                raise ValueError('Missing total cell for the selected column')
            updated, count = TOTAL.subn(f'~ {result["total_ms"]:.2f} ms', cells[column + 1], count=1)
            if count != 1:
                raise ValueError('Missing total cell for the selected column')
            if mode != 'x':
                reference = float(re.search(r'~\s*([0-9.]+) ms', section(text, 'C (reference)').group(1)).group(1))
                updated, count = RATIO.subn(f'/C = {result["total_ms"] / reference:.1f}x', updated, count=1)
                if count != 1:
                    raise ValueError('Missing /C annotation for the selected column')
            cells[column + 1] = updated
            lines[index] = '|'.join(cells) + newline
    if seen != len(rows):
        raise ValueError('Expected ' + str(len(rows)) + ' timing rows, found ' + str(seen))
    text = text[:rust.start(1)] + ''.join(lines) + text[rust.end(1):]
    if README.read_text() != original:
        raise RuntimeError('README changed while preparing the column')
    README.write_text(text)
    print('Updated README ' + SECTIONS[mode] + ' column (mode ' + mode + '); total '
          + f'{result["total_ms"]:.6f}' + ' ms')


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--mode', choices=list(MODES), required=True)
    parser.add_argument('--build-dir', type=Path, help='isolated build workspace')
    parser.add_argument('--campaign-dir', type=Path, help='new campaign output directory')
    parser.add_argument('--clean', action='store_true', help='rebuild the compiler and workspace from scratch')
    parser.add_argument('--build-only', action='store_true')
    parser.add_argument('--update-readme', action='store_true',
                        help='write the measured column, total and /C into README.md')
    parser.add_argument('--publish-only', action='store_true',
                        help='update README from an existing --campaign-dir without building or measuring')
    args = parser.parse_args()
    mode = args.mode
    campaign_dir = (args.campaign_dir or ROOT / 'var/benchmark' / ('rust-' + mode + '-' + stamp())).resolve()
    if args.publish_only:
        if not args.update_readme:
            parser.error('--publish-only requires --update-readme')
        if args.campaign_dir is None:
            parser.error('--publish-only requires --campaign-dir')
        result = json.loads((campaign_dir / 'results.json').read_text())['rust-' + mode]
        verify(result, mode)
        update_readme(mode, result)
        return
    build_dir = (args.build_dir or ROOT / 'run/bak/rust/modes' / ('campaign-' + mode)).resolve()
    if str(build_dir) in {str(ROOT), *map(str, ROOT.parents)}:
        raise ValueError('Unsafe build directory: ' + str(build_dir))
    if not args.build_only and args.update_readme and args.campaign_dir is None:
        raise ValueError('--update-readme needs a campaign directory or the default (recorded) one')
    build(mode, build_dir, args.clean)
    if args.build_only:
        print('Built Rust ' + mode + '; manifest ' + str(build_dir / 'manifest.json'))
        return
    result = measure(mode, build_dir, campaign_dir)
    (campaign_dir / 'toolchain.json').write_text(json.dumps(toolchain(), indent=2) + '\n')
    total = verify(result, mode)
    print('', flush=True)
    for label, value in zip(result['labels'], result['times_us']):
        print(f'  {label:60} ~ {value:14.3f} μs')
    print(f'  {"Total Execution Time":60} ~ {total:14.3f} ms')
    reference = published(mode)
    if reference:
        print(f'  Published README rust-{mode} total: {reference:.2f} ms; '
              f'this campaign: {result["total_ms"]:.2f} ms '
              f'({100.0 * (result["total_ms"] - reference) / reference:+.1f}%)')
    print('Campaign: ' + str(campaign_dir / 'results.json'), flush=True)
    if args.update_readme:
        update_readme(mode, result)


if __name__ == '__main__':
    try:
        main()
    except (OSError, ValueError, RuntimeError, subprocess.SubprocessError) as error:
        print(f'ERROR: {error}', file=sys.stderr)
        sys.exit(1)

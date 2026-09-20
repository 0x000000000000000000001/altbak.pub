#!/usr/bin/env python3
"""Measure an already-built Fable Rust binary in three validated processes.

No compilation occurs. The binary must implement altbak's numeric-kernel
protocol: 14 outputs, calibrated power-of-two batches, and best-of-ten timings.
"""
from __future__ import annotations
import argparse
from datetime import datetime, timezone
import hashlib
import importlib.util
import json
import math
import os
from pathlib import Path
import platform
import re
import statistics
import subprocess
import sys
import tempfile
import time

sys.dont_write_bytecode = True
TABLE_ROWS = [
    'AST Evaluation', 'Fibonacci', 'List Processing', 'Tail Call Optimization',
    'Deep Record Updates', 'Ackermann', 'Church Numerals', 'Prime Sieve',
    'Red-Black Tree', 'Polymorphism', 'State Monad', 'Lazy Evaluation',
    'Array Processing', 'RowToList',
]
PROCESS_COUNT = 3
VARIANTS = {
    'sharpurs': {'result_key': 'Fable', 'label': 'Fable Rust (sharpurs)'},
    'native-fsharp': {'result_key': 'Fable (native F#)', 'label': 'Fable Rust (native F#)'},
}


def utc_now() -> str:
    return datetime.now(timezone.utc).isoformat()


def sha256(path: Path) -> str:
    digest = hashlib.sha256()
    with path.open('rb') as stream:
        for chunk in iter(lambda: stream.read(1024 * 1024), b''):
            digest.update(chunk)
    return digest.hexdigest()


def discover_project_root() -> Path:
    # Works when installed in altbak.pub/tmp/fable_rust and from scratch tools.
    candidates = [*Path(__file__).resolve().parents, Path.cwd(), Path.cwd()/'altbak.pub']
    for candidate in candidates:
        if (candidate/'bin/benchmark/validate.py').is_file() and (candidate/'tmp/run_purust_benchmark.py').is_file():
            return candidate.resolve()
    raise ValueError('Cannot locate altbak.pub; provide --project-root PATH')


def write_json(path: Path, value) -> None:
    temporary = path.with_suffix(path.suffix + '.tmp')
    temporary.write_text(json.dumps(value, indent=2, allow_nan=False) + '\n')
    temporary.replace(path)


def snapshot(path: Path) -> dict:
    before = path.stat()
    digest = sha256(path)
    after = path.stat()
    fields = lambda s: (s.st_dev, s.st_ino, s.st_size, s.st_mtime_ns)
    if fields(before) != fields(after):
        raise RuntimeError(f'File changed while fingerprinting: {path}')
    return {'sha256': digest, 'identity': fields(after)}


def ensure_unchanged(path: Path, original: dict) -> None:
    if snapshot(path) != original:
        raise RuntimeError(f'File changed during measurement: {path}; retained results, no README update')


def checked_result(output: str, validate_output) -> dict:
    # Same assertions as tmp/run_purust_benchmark.py, plus explicit finiteness.
    result = validate_output(output)
    counts = [int(n) for n in re.findall(r'^Batch iterations: (\d+)$', output, re.M)]
    if len(counts) != 14 or any(n < 1 or n > 16777216 or n & (n-1) for n in counts):
        raise ValueError('Invalid batch sizes; expected 14 positive powers of two <= 16777216')
    if any(not math.isfinite(t) or t <= 0 for t in result['times_us']):
        raise ValueError('Nonpositive/nonfinite duration')
    if result.get('values_validated') is not True or len(result['values']) != 14:
        raise ValueError('The full 14-case result oracle was not validated')
    result['batch_iterations'] = counts
    return result


def table_coordinates(original: str, variant: str = 'sharpurs') -> tuple[list[str], int, int, int]:
    if variant not in VARIANTS:
        raise ValueError(f'Unknown Fable variant: {variant!r}')
    lines = original.splitlines(keepends=True)
    matches = []
    for line_number, line in enumerate(lines):
        if '|' not in line:
            continue
        for column, cell in enumerate(line.split('|')):
            if variant == 'sharpurs':
                selected = ('Fable' in cell
                            and ('sharpurs' in cell or 'Compiled Rust (Fable' in cell))
            else:
                selected = 'Hand-written F#' in cell and 'Fable' in cell
            if selected:
                matches.append((line_number, column))
    if len(matches) != 1:
        raise ValueError(f'Expected exactly one Fable header cell for {variant}')
    header, column = matches[0]
    header_cells = lines[header].rstrip('\r\n').split('|')
    label_column = 1 if not header_cells[0].strip() else 0
    if header + 16 >= len(lines):
        raise ValueError('Fable table is incomplete')
    separators = lines[header+1].rstrip('\r\n').split('|')
    if (len(separators) != len(header_cells)
            or any(not re.fullmatch(r'\s*:?-+:?\s*', separators[i])
                   for i, cell in enumerate(header_cells) if cell.strip())):
        raise ValueError('Unexpected Fable table separator')
    expected = TABLE_ROWS + ['**Total Execution Time**']
    for offset, label in enumerate(expected, start=2):
        cells = lines[header+offset].rstrip('\r\n').split('|')
        if len(cells) != len(header_cells) or cells[label_column].strip() != label:
            raise ValueError(f'Unexpected Fable table row/order: expected {label!r}')
    return lines, header, column, label_column


def render_readme(original: str, result: dict, variant: str = 'sharpurs') -> str:
    lines, header, column, _ = table_coordinates(original, variant)
    if len(result['times_us']) != len(TABLE_ROWS):
        raise ValueError('Expected exactly 14 benchmark timings')
    edits = {header+i+2: f"~ {value:.3f} μs" for i, value in enumerate(result['times_us'])}
    edits[header+16] = f"~ {result['sum_displayed_lines_ms']:.2f} ms"
    if len(edits) != 15:
        raise ValueError('Expected exactly 15 target cells')
    original_lines = list(lines)
    for number, value in edits.items():
        line = lines[number]
        newline = '\r\n' if line.endswith('\r\n') else '\n' if line.endswith('\n') else ''
        cells = line[:-len(newline)].split('|') if newline else line.split('|')
        old = cells[column]
        leading = re.match(r'^\s*', old).group()
        trailing = re.search(r'\s*$', old).group()
        cells[column] = leading + value + trailing
        lines[number] = '|'.join(cells) + newline
        old_cells = original_lines[number].split('|')
        new_cells = lines[number].split('|')
        if any(a != b for i, (a, b) in enumerate(zip(old_cells, new_cells)) if i != column):
            raise AssertionError('A non-Fable cell changed')
    if any(a != b for i, (a, b) in enumerate(zip(original_lines, lines)) if i not in edits):
        raise AssertionError('Text outside the 15 Fable cells changed')
    return ''.join(lines)


def aggregate(samples: list[dict]) -> dict:
    if len(samples) != PROCESS_COUNT:
        raise ValueError('Three validated process runs are required')
    result = dict(samples[0])
    result.pop('batch_iterations', None)
    result['times_us'] = [statistics.median(s['times_us'][i] for s in samples) for i in range(14)]
    # README follows the existing Purust convention: sum the per-row medians.
    result['total_ms'] = result['sum_displayed_lines_ms'] = sum(result['times_us']) / 1000.0
    result['median_process_total_ms'] = statistics.median(s['total_ms'] for s in samples)
    result['process_runs'] = samples
    result['aggregation'] = 'median per benchmark across three processes; displayed total sums median rows'
    return result


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--binary', required=True, type=Path)
    parser.add_argument('--variant', choices=VARIANTS, default='sharpurs',
                        help='Benchmark source route and README target column (default: sharpurs)')
    parser.add_argument('--output-dir', required=True, type=Path,
                        help='New empty results directory; never overwrites previous runs')
    parser.add_argument('--project-root', type=Path, help='Default: discover from this script or the working directory')
    parser.add_argument('--readme', type=Path, help='Update only the 14 Fable cells and their total after successful measurement')
    parser.add_argument('--timeout', type=float, default=180.0, help='Timeout in seconds for each complete process (default 180)')
    parser.add_argument('--allocator', default='mimalloc',
                        help='Provenance description supplied by builder (default: mimalloc)')
    parser.add_argument('--build-manifest', type=Path, help='Optional JSON build provenance to preserve and fingerprint')
    args = parser.parse_args()
    variant = VARIANTS[args.variant]
    label = variant['label']
    if not math.isfinite(args.timeout) or args.timeout <= 0:
        parser.error('--timeout must be a positive finite number')
    binary, directory = args.binary.resolve(), args.output_dir.resolve()
    try:
        root = args.project_root.resolve() if args.project_root else discover_project_root()
    except ValueError as exc:
        parser.error(str(exc))
    if not binary.is_file() or not os.access(binary, os.X_OK):
        parser.error('--binary must be an existing executable')
    if directory.exists() and any(directory.iterdir()):
        parser.error('--output-dir must be new or empty')
    readme = args.readme.resolve() if args.readme else None
    readme_bytes = readme.read_bytes() if readme else None
    readme_state = snapshot(readme) if readme else None
    if readme:
        if hashlib.sha256(readme_bytes).hexdigest() != readme_state['sha256']:
            raise RuntimeError('README changed while reading its initial snapshot')
        table_coordinates(readme_bytes.decode('utf-8'), args.variant)  # Fail before measuring malformed tables.
    validator_path = root / 'bin/benchmark/validate.py'
    spec = importlib.util.spec_from_file_location('altbak_fable_validator', validator_path)
    validator = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(validator)
    oracle_files = [root/'src/Test'/f'{name}.purs' for name in validator.CASES]
    watched = {p: snapshot(p) for p in [binary, validator_path, Path(__file__).resolve(), *oracle_files]}
    build_manifest = args.build_manifest.resolve() if args.build_manifest else None
    build_provenance = None
    if build_manifest:
        watched[build_manifest] = snapshot(build_manifest)
        build_provenance = json.loads(build_manifest.read_text())
        ensure_unchanged(build_manifest, watched[build_manifest])
    if readme:
        watched[readme] = readme_state
    directory.mkdir(parents=True, exist_ok=True)
    manifest = {
        'started_at_utc': utc_now(), 'status': 'measuring',
        'variant': args.variant, 'result_key': variant['result_key'],
        'binary': str(binary), 'binary_sha256': watched[binary]['sha256'],
        'command': [str(binary)], 'working_directory': str(directory),
        'platform': platform.platform(), 'machine': platform.machine(),
        'allocator': {'description': args.allocator, 'evidence': 'Builder-supplied CLI description; not inferred from binary'},
        'build_manifest': str(build_manifest) if build_manifest else None,
        'build_provenance': build_provenance,
        'protocol': {
            'process_runs': PROCESS_COUNT, 'timeout_seconds_per_process': args.timeout,
            'global_warmups': 3, 'local_warmups': 3, 'samples_per_process': 10,
            'calibration_target_ms': 10, 'maximum_batch_iterations': 16777216,
            'statistic': 'median across three processes of per-process minimum batch time divided by invocation count',
            'displayed_total': 'sum of per-row medians',
            'additional_total': 'median of the three printed process totals',
            'numeric_kernels': True,
            'implementation_note': 'Warm-up counts and clock behavior are declared protocol requirements; printed results and batch sizes are validated.',
        },
        'fingerprints': {str(p): s['sha256'] for p, s in watched.items()},
        'readme_requested': str(readme) if readme else None,
        'readme_updated': False,
        'processes': [],
    }
    write_json(directory/'provenance.json', manifest)
    samples = []
    try:
        for repetition in range(PROCESS_COUNT):
            for path, state in watched.items():
                ensure_unchanged(path, state)
            log = directory/f'fable-results-{repetition+1}.log'
            print(f'Measuring {label}, process {repetition+1}/{PROCESS_COUNT}; log: {log}', flush=True)
            started = utc_now()
            start = time.monotonic()
            process_info = {'number': repetition+1, 'started_at_utc': started, 'log': log.name}
            manifest['processes'].append(process_info)
            write_json(directory/'provenance.json', manifest)
            with log.open('w', encoding='utf-8') as output:
                try:
                    completed = subprocess.run([str(binary)], cwd=directory, stdout=output,
                                               stderr=subprocess.STDOUT, timeout=args.timeout)
                except subprocess.TimeoutExpired:
                    process_info.update({'timeout': True, 'elapsed_wall_seconds': time.monotonic()-start})
                    raise RuntimeError(f'Process {repetition+1} exceeded {args.timeout:g}s; log retained: {log}')
            process_info.update({'returncode': completed.returncode, 'elapsed_wall_seconds': time.monotonic()-start})
            if completed.returncode:
                raise RuntimeError(f'Process {repetition+1} failed with exit {completed.returncode}; see {log}')
            result = checked_result(log.read_text(), validator.validate_output)
            process_info['validated_cases'] = 14
            process_info['log_sha256'] = sha256(log)
            samples.append(result)
            write_json(directory/'validated-processes.json', samples)
            for path, state in watched.items():
                ensure_unchanged(path, state)
            write_json(directory/'provenance.json', manifest)
            print(f'{label}: 14/14 expected outputs verified; process total {result["total_ms"]:.6f} ms', flush=True)
        result = aggregate(samples)
        write_json(directory/'results.json', {variant['result_key']: result})
        for path, state in watched.items():
            ensure_unchanged(path, state)
        if readme:
            updated = render_readme(readme_bytes.decode('utf-8'), result, args.variant).encode('utf-8')
            ensure_unchanged(readme, readme_state)
            # Preserve permissions and use an atomic rename to avoid partial README writes.
            fd, temp_name = tempfile.mkstemp(prefix='.'+readme.name+'.fable-', dir=readme.parent)
            temp = Path(temp_name)
            try:
                with os.fdopen(fd, 'wb') as output:
                    output.write(updated)
                    output.flush()
                    os.fsync(output.fileno())
                temp.chmod(readme.stat().st_mode)
                ensure_unchanged(readme, readme_state)
                temp.replace(readme)
            finally:
                if temp.exists():
                    temp.unlink()
            manifest['readme_updated'] = True
            manifest['readme_sha256_after'] = sha256(readme)
        manifest.update({'status': 'complete', 'completed_at_utc': utc_now()})
        write_json(directory/'provenance.json', manifest)
        print(f'{label}: sum of median rows {result["total_ms"]:.6f} ms; median process total {result["median_process_total_ms"]:.6f} ms', flush=True)
        print(f'Results: {directory/"results.json"}', flush=True)
        if readme:
            print(f'README: only the 15 {args.variant} Fable target cells updated.', flush=True)
        return 0
    except (OSError, ValueError, RuntimeError, KeyboardInterrupt) as exc:
        manifest.update({'status': 'failed', 'failed_at_utc': utc_now(), 'error': str(exc)})
        write_json(directory/'provenance.json', manifest)
        print(f'{label} collection failed: {exc}', file=sys.stderr)
        return 1


if __name__ == '__main__':
    sys.exit(main())

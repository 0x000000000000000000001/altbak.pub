#!/usr/bin/env python3
"""Separate whole-process CPU/RSS diagnostic with the exact Aff output oracle."""
import argparse
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import statistics
import subprocess


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def files(directory):
    return sorted(({'path': str(p.relative_to(directory)), 'bytes': p.stat().st_size, 'sha256': sha(p)}
                   for p in directory.rglob('*') if p.is_file() and p.suffix in ['.rs', '.toml']), key=lambda p: p['path'])


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('comparison', type=Path, help='completed compare-aff JSON')
    parser.add_argument('out', type=Path)
    args = parser.parse_args()
    comparison = json.loads(args.comparison.read_text())
    assert comparison['validation']['status'] == 'passed'
    out = args.out.resolve()
    out.mkdir(parents=True, exist_ok=False)
    (out / 'generated').mkdir()
    source_run = next(r for r in comparison['runs'] if not r['warmup'])
    cwd = Path(source_run['cwd'])
    # Use the copied immutable corpus, never the original oracle directory.
    ffi = cwd.parent.parent / 'rust-ffi'
    oracle = comparison['provenance']['reference_files']
    def verify_inputs():
        for key, folder in [('frozen_inputs', cwd.parent), ('frozen_ffi', ffi)]:
            for file in comparison['provenance'][key]:
                assert sha(folder / file['path']) == file['sha256']
    verify_inputs()
    env = {k: v for k, v in os.environ.items() if not re.match(
        r'^(PURUST_|GOPURS_|NODE_|RUST|CARGO|GOGC$|GOMEMLIMIT$|GOMAXPROCS$|GODEBUG$)', k)}
    variants = comparison['variants']
    assert len(variants) == 2
    report = {'metric': 'whole process including startup, teardown and descendants; separate from backend total',
              'source_comparison': str(args.comparison.resolve()), 'source_sha256': sha(args.comparison),
              'protocol': 'one warmup per variant then three alternating pairs; /usr/bin/time -l on macOS; peak RSS bytes',
              'variants': variants, 'runs': [], 'status': 'pending'}
    save = lambda: (out / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
    save()
    try:
        for round in range(4):
            order = variants if round % 2 else variants[::-1]
            for variant in order:
                binary = Path(variant['binary'])
                assert sha(binary) == variant['sha256']
                name = f"{round}-{variant['label']}"
                for cache in ['.cache', '.purmeta']:
                    shutil.rmtree(cwd / cache, ignore_errors=True)
                generated = out / 'generated' / name
                command = ['/usr/bin/time', '-l', str(binary), '--main', 'Test.Main', '--threaded', '--source', 'output',
                           '--out', str(generated), '--ffi-dir', os.path.relpath(ffi, cwd)]
                p = subprocess.run(command, cwd=cwd, env=dict(env, **variant.get('env', {})), capture_output=True, text=True, timeout=180)
                (out / (name + '.stdout')).write_text(p.stdout)
                (out / (name + '.stderr')).write_text(p.stderr)
                run = {'label': variant['label'], 'round': round, 'warmup': round == 0, 'command': command,
                       'exit_code': p.returncode, 'stdout': name + '.stdout', 'stderr': name + '.stderr'}
                report['runs'].append(run)
                save()
                assert p.returncode == 0, p.stderr
                timing = re.search(r'(\d+\.\d+) real\s+(\d+\.\d+) user\s+(\d+\.\d+) sys', p.stderr)
                rss = re.search(r'(\d+)\s+maximum resident set size', p.stderr)
                assert timing and rss
                run.update(wall_seconds=float(timing[1]), user_seconds=float(timing[2]), system_seconds=float(timing[3]),
                           cpu_seconds=float(timing[2]) + float(timing[3]), max_rss_bytes=int(rss[1]))
                assert files(generated) == oracle, name
                run['identical_files'] = len(oracle)
                save()
                print(name, run['cpu_seconds'], run['max_rss_bytes'], flush=True)
        report['summary'] = {v['label']: {key: statistics.median(r[key] for r in report['runs']
            if r['label'] == v['label'] and not r['warmup'])
            for key in ['wall_seconds', 'user_seconds', 'system_seconds', 'cpu_seconds', 'max_rss_bytes']} for v in variants}
        verify_inputs()
        for variant in variants:
            assert sha(Path(variant['binary'])) == variant['sha256']
        report['status'] = 'passed'
        save()
    except BaseException as error:
        report['status'], report['failure'] = 'failed', repr(error)
        save()
        raise


if __name__ == '__main__':
    main()

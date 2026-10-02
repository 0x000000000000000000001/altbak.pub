#!/usr/bin/env python3
"""Build a Rust-only TAST candidate, or compare it with the frozen phase-A binary.

Builds and measurements are separate commands. Measurement uses archived
executables and oracles, so editing live PBO sources cannot silently change
either side of the before/after comparison.
"""
import argparse
import importlib.util
import json
import os
from pathlib import Path
import shutil
import statistics
import subprocess
import time

HERE = Path(__file__).resolve().parent
spec = importlib.util.spec_from_file_location('json_diagnostic', HERE.parent / 'json-diagnostic.py')
diag = importlib.util.module_from_spec(spec)
spec.loader.exec_module(diag)


def save(path, value):
    path.write_text(json.dumps(value, indent=2) + '\n')


def copy(source, destination):
    destination.parent.mkdir(parents=True, exist_ok=True)
    shutil.copy2(source, destination)
    return {'path': str(destination), 'sha256': diag.sha(destination)}


def build(args):
    out = args.out.resolve()
    out.mkdir(parents=True, exist_ok=False)
    (out / 'logs').mkdir()
    before = diag.fingerprint('JsonTypedAst')
    report = {'status': 'building', 'sources': before, 'profile': diag.RUST_TYPED_AST_PROFILE}
    save(out / 'build.json', report)
    try:
        for name, sha256 in before.items():
            source = Path(name)
            frozen = copy(source, out / 'sources' / source.relative_to(diag.ROOT.parent))
            assert frozen['sha256'] == sha256, name
        copy(Path(__file__).resolve(), out / 'harness/diagnostic.py')
        env = diag.environment()
        env.update(CARGO_BUILD_JOBS='8', CARGO_INCREMENTAL='0', GHCRTS='-N2')
        report['rust'] = diag.build_rust(out, env, 'JsonTypedAst')
        assert before == diag.fingerprint('JsonTypedAst'), 'Live sources changed during build'
        report['timed'] = copy(Path(report['rust']['binary']), out / 'rust-timed')
        report['alloc'] = copy(Path(report['rust']['alloc_binary']), out / 'rust-alloc')
        report['status'] = 'passed'
        save(out / 'build.json', report)
        print(json.dumps({key: report[key] for key in ['status', 'timed', 'alloc', 'profile']}, indent=2))
    except BaseException as error:
        report['failure'] = str(error)
        save(out / 'build.json', report)
        raise


def measure(args):
    baseline, candidate, out = args.baseline.resolve(), args.candidate.resolve(), args.out.resolve()
    baseline_manifest = json.loads((baseline / 'artifacts/provenance.json').read_text())
    candidate_manifest = json.loads((candidate / 'build.json').read_text())
    assert candidate_manifest['status'] == 'passed'
    original = next(item for item in baseline_manifest['frozen'] if item['copied'] == 'artifacts/rust-timed')
    before = baseline / original['copied']
    after = Path(candidate_manifest['timed']['path'])
    assert diag.sha(before) == original['sha256']
    assert diag.sha(after) == candidate_manifest['timed']['sha256']
    original_build = json.loads((baseline / 'artifacts/workspace-manifest.json').read_text())
    assert original_build['rust']['profile'] == candidate_manifest['profile'], 'Build profiles differ'
    assert original_build['rust']['rustc'] == candidate_manifest['rust']['rustc'], 'Rust compiler versions differ'
    original_driver = baseline / 'artifacts/sources/JsonTypedAst.rs'
    candidate_driver = candidate / 'sources/altbak.pub/src/Test/JsonTypedAst.rs'
    assert original_driver.read_bytes() == candidate_driver.read_bytes(), 'Benchmark driver changed'
    if args.corpus == 'gopurs238':
        corpus = (baseline / 'campaign-238/corpus.json').read_bytes()
        oracle = json.loads((baseline / 'oracle-gopurs238.json').read_text())
        assert diag.corpus_sha(corpus) == oracle['corpus_sha256']
        original_results = json.loads((baseline / 'campaign-238/results.json').read_text())
        diag.validate_result(original_results['results']['js'][0], oracle, 'JsonTypedAst')
    else:
        corpus, _ = diag.corpus_data('JsonTypedAst', 'fixture12')
        original_results = json.loads((baseline / 'campaign-12/results.json').read_text())
        assert diag.corpus_sha(corpus) == original_results['corpus']['corpus_sha256']
        oracle_path = diag.FIXTURES / 'expected.json'
        assert diag.sha(oracle_path) == original_build['sources'][str(oracle_path)], 'Historical oracle changed'
        oracle = json.loads(oracle_path.read_text())
    out.mkdir(parents=True, exist_ok=False)
    binaries = {label: copy(source, out / 'binaries' / label) for label, source in [('before', before), ('after', after)]}
    (out / 'corpus.json').write_bytes(corpus)
    save(out / 'oracle.json', oracle)
    copy(Path(__file__).resolve(), out / 'harness/diagnostic.py')
    corpus_hash, oracle_hash = diag.sha(out / 'corpus.json'), diag.sha(out / 'oracle.json')
    report = {'corpus': args.corpus, 'corpus_sha256': corpus_hash, 'modules': oracle['modules'],
              'oracle_sha256': oracle_hash, 'baseline': str(baseline), 'candidate': str(candidate),
              'binaries': binaries, 'profile': candidate_manifest['profile'],
              'protocol': {'processes_per_variant': 3, 'warmups_per_phase': 2, 'samples_per_phase': 5,
                           'statistic': 'median of three process minima',
                           'order': 'before/after, after/before, before/after; sequential',
                           'fingerprinting_timed': False, 'result_destruction_timed_separately': True},
              'runs': [], 'validation': 'pending'}
    env = diag.environment()
    env.update(DIAG_CORPUS=str(out / 'corpus.json'), DIAG_PROFILE=json.dumps(report['profile']))
    result_path = out / 'results.json'
    save(result_path, report)
    try:
        for process_index in range(3):
            order = ['before', 'after'] if process_index % 2 == 0 else ['after', 'before']
            for label in order:
                name = f'{process_index + 1}-{label}'
                started, load = time.monotonic(), os.getloadavg()
                try:
                    completed = subprocess.run([binaries[label]['path']], cwd=out, env=env,
                                               capture_output=True, text=True, timeout=240)
                except subprocess.TimeoutExpired as error:
                    for suffix, data in [('stdout', error.stdout), ('stderr', error.stderr)]:
                        (out / (name + '.' + suffix)).write_bytes(data.encode() if isinstance(data, str) else (data or b''))
                    report['runs'].append({'variant': label, 'process': process_index + 1, 'status': 'timeout',
                                           'wall_seconds': time.monotonic() - started,
                                           'stdout': name + '.stdout', 'stderr': name + '.stderr'})
                    raise
                (out / (name + '.stdout')).write_text(completed.stdout)
                (out / (name + '.stderr')).write_text(completed.stderr)
                record = {'variant': label, 'process': process_index + 1, 'status': completed.returncode,
                          'wall_seconds': time.monotonic() - started, 'load_average_before': load,
                          'stdout': name + '.stdout', 'stderr': name + '.stderr'}
                report['runs'].append(record)
                save(result_path, report)
                assert completed.returncode == 0, completed.stderr[-5000:]
                value = json.loads(completed.stdout)
                diag.validate_result(value, oracle, 'JsonTypedAst')
                assert value['warmups_per_phase'] == 2 and value['samples_per_phase'] == 5
                assert not value['alloc_diag']
                record['result'] = value
                save(result_path, report)
                print(name, {phase: round(data['time_us'] / 1000, 3) for phase, data in value['phases'].items()}, flush=True)
        summary = {}
        for label in binaries:
            values = [run['result'] for run in report['runs'] if run['variant'] == label]
            summary[label] = {phase: {
                'process_minima_us': [value['phases'][phase]['time_us'] for value in values],
                'median_process_minimum_us': statistics.median(value['phases'][phase]['time_us'] for value in values),
                'median_all_samples_us': statistics.median(sample['time_us'] for value in values for sample in value['phases'][phase]['samples']),
                'median_drop_us': statistics.median(sample['drop_us'] for value in values for sample in value['phases'][phase]['samples']),
            } for phase in ['parse', 'decode', 'combined']}
        report['summary'] = summary
        report['after_over_before'] = {phase: summary['after'][phase]['median_process_minimum_us'] /
                                      summary['before'][phase]['median_process_minimum_us'] for phase in summary['before']}
        for artifact in binaries.values():
            assert diag.sha(Path(artifact['path'])) == artifact['sha256']
        assert diag.sha(out / 'corpus.json') == corpus_hash
        assert diag.sha(out / 'oracle.json') == oracle_hash
        report['validation'] = 'passed: every JSON and TAST fingerprint matches the frozen oracle'
        save(result_path, report)
        print(json.dumps({'summary': summary, 'after_over_before': report['after_over_before']}, indent=2))
    except BaseException as error:
        report['failure'] = str(error)
        save(result_path, report)
        raise


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    commands = parser.add_subparsers(dest='action', required=True)
    build_parser = commands.add_parser('build')
    build_parser.add_argument('--out', type=Path, required=True)
    measure_parser = commands.add_parser('measure')
    measure_parser.add_argument('--baseline', type=Path, required=True)
    measure_parser.add_argument('--candidate', type=Path, required=True)
    measure_parser.add_argument('--out', type=Path, required=True)
    measure_parser.add_argument('--corpus', choices=['fixture12', 'gopurs238'], required=True)
    args = parser.parse_args()
    (build if args.action == 'build' else measure)(args)


if __name__ == '__main__':
    main()

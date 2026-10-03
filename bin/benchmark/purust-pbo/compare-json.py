#!/usr/bin/env python3
"""Shared Go/Rust JSON-to-TAST campaign on frozen 12/238-module inputs."""
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
spec = importlib.util.spec_from_file_location('diag', HERE.parent / 'json-diagnostic.py')
diag = importlib.util.module_from_spec(spec)
spec.loader.exec_module(diag)


def save(path, value):
    path.write_text(json.dumps(value, indent=2) + '\n')


def main():
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument('--reference', type=Path, required=True)
    p.add_argument('--before', type=Path, required=True, help='Rust-only diagnostic build directory')
    p.add_argument('--after', type=Path, required=True, help='Rust-only diagnostic build directory')
    p.add_argument('--out', type=Path, required=True)
    p.add_argument('--corpus', choices=['fixture12', 'gopurs238'], required=True)
    a = p.parse_args()
    reference, out = a.reference.resolve(), a.out.resolve()
    out.mkdir(parents=True, exist_ok=False)
    frozen = json.loads((reference / 'artifacts/provenance.json').read_text())
    go = next(x for x in frozen['frozen'] if x['copied'] == 'artifacts/go-benchmark')
    assert diag.sha(reference / go['copied']) == go['sha256']
    builds = {label: json.loads((path.resolve() / 'build.json').read_text())
              for label, path in [('before', a.before), ('after', a.after)]}
    assert all(b['status'] == 'passed' for b in builds.values())
    assert builds['before']['profile'] == builds['after']['profile']
    assert builds['before']['rust']['rustc'] == builds['after']['rust']['rustc']
    for path in [a.before, a.after]:
        for suffix in ['purs', 'rs']:
            assert (reference / f'artifacts/sources/JsonTypedAst.{suffix}').read_bytes() == (
                path / f'sources/altbak.pub/src/Test/JsonTypedAst.{suffix}').read_bytes()
    sources = {'go': reference / go['copied'], **{k: Path(b['timed']['path']) for k, b in builds.items()}}
    binaries = {}
    (out / 'binaries').mkdir()
    for label, source in sources.items():
        expected = go['sha256'] if label == 'go' else builds[label]['timed']['sha256']
        assert diag.sha(source) == expected
        dest = out / 'binaries' / label
        shutil.copy2(source, dest)
        binaries[label] = {'path': str(dest), 'sha256': expected}
    if a.corpus == 'gopurs238':
        corpus = (reference / 'campaign-238/corpus.json').read_bytes()
        oracle = json.loads((reference / 'oracle-gopurs238.json').read_text())
        assert diag.corpus_sha(corpus) == oracle['corpus_sha256']
    else:
        corpus, _ = diag.corpus_data('JsonTypedAst', 'fixture12')
        prior = json.loads((reference / 'campaign-12/results.json').read_text())
        assert diag.corpus_sha(corpus) == prior['corpus']['corpus_sha256']
        oracle = json.loads((diag.FIXTURES / 'expected.json').read_text())
        diag.validate_result(prior['results']['go'][0], oracle, 'JsonTypedAst')
    (out / 'corpus.json').write_bytes(corpus)
    save(out / 'oracle.json', oracle)
    shutil.copy2(__file__, out / 'harness.py')
    result = {'status': 'pending', 'corpus': a.corpus, 'modules': oracle['modules'],
              'corpus_sha256': diag.sha(out / 'corpus.json'), 'oracle_sha256': diag.sha(out / 'oracle.json'),
              'binaries': binaries, 'builds': {k: str(v.resolve()) for k, v in [('before', a.before), ('after', a.after)]},
              'rust_profile': builds['after']['profile'], 'runs': [],
              'protocol': {'processes_per_variant': 3, 'warmups_per_phase': 2, 'samples_per_phase': 5,
                           'statistic': 'median of three process minima', 'order': 'rotating runtime and phase orders',
                           'GOMAXPROCS': 1, 'GOGC': 100, 'file_io_timed': False, 'fingerprinting_timed': False,
                           'destruction': 'Rust final-result drop recorded separately; Go GC uses its normal policy'}}
    report = out / 'results.json'
    save(report, result)
    env = diag.environment()
    env.update(DIAG_CORPUS=str(out / 'corpus.json'), DIAG_PROFILE=json.dumps(result['rust_profile']))
    phases = ['parse', 'decode', 'combined']
    try:
        labels = list(binaries)
        for index in range(3):
            order = labels[index:] + labels[:index]
            phase_order = phases[index:] + phases[:index]
            for label in order:
                name = f'{index + 1}-{label}'
                started = time.monotonic()
                completed = subprocess.run([binaries[label]['path']], cwd=out,
                    env=dict(env, DIAG_PHASES=','.join(phase_order)), capture_output=True, text=True, timeout=300)
                (out / (name + '.stdout')).write_text(completed.stdout)
                (out / (name + '.stderr')).write_text(completed.stderr)
                record = {'variant': label, 'process': index + 1, 'exit_code': completed.returncode,
                          'wall_seconds': time.monotonic() - started, 'stdout': name + '.stdout', 'stderr': name + '.stderr'}
                result['runs'].append(record)
                save(report, result)
                assert completed.returncode == 0, completed.stderr[-5000:]
                value = json.loads(completed.stdout)
                diag.validate_result(value, oracle, 'JsonTypedAst')
                assert value['phase_order'] == phase_order
                if label != 'go':
                    assert value['warmups_per_phase'] == 2 and value['samples_per_phase'] == 5 and not value['alloc_diag']
                record['result'] = value
                save(report, result)
                print(name, {k: round(v['time_us'] / 1000, 3) for k, v in value['phases'].items()}, flush=True)
        result['summary'] = {}
        for label in labels:
            runs = [r['result'] for r in result['runs'] if r['variant'] == label]
            result['summary'][label] = {phase: {
                'process_minima_us': [r['phases'][phase]['time_us'] for r in runs],
                'median_process_minimum_us': statistics.median(r['phases'][phase]['time_us'] for r in runs),
                'median_all_samples_us': statistics.median(s['time_us'] for r in runs for s in r['phases'][phase]['samples']),
                'median_drop_us': statistics.median(s['drop_us'] for r in runs for s in r['phases'][phase]['samples']) if label != 'go' else None,
            } for phase in phases}
        for binary in binaries.values():
            assert diag.sha(Path(binary['path'])) == binary['sha256']
        assert diag.sha(out / 'corpus.json') == result['corpus_sha256']
        assert diag.sha(out / 'oracle.json') == result['oracle_sha256']
        result['status'] = 'passed'
        save(report, result)
        print(json.dumps(result['summary'], indent=2))
    except BaseException as error:
        result['status'], result['failure'] = 'failed', repr(error)
        save(report, result)
        raise


if __name__ == '__main__':
    main()

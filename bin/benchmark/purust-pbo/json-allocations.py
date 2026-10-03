#!/usr/bin/env python3
"""Allocation-only Rust diagnostic using a passed common JSON campaign's corpus."""
import argparse
import importlib.util
import json
from pathlib import Path
import shutil
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
spec = importlib.util.spec_from_file_location('diag', HERE.parent / 'json-diagnostic.py')
diag = importlib.util.module_from_spec(spec)
spec.loader.exec_module(diag)


def main():
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument('campaign', type=Path)
    p.add_argument('out', type=Path)
    a = p.parse_args()
    campaign = a.campaign.resolve()
    result = json.loads((campaign / 'results.json').read_text())
    assert result['status'] == 'passed'
    assert diag.sha(campaign / 'corpus.json') == result['corpus_sha256']
    assert diag.sha(campaign / 'oracle.json') == result['oracle_sha256']
    oracle = json.loads((campaign / 'oracle.json').read_text())
    out = a.out.resolve()
    out.mkdir(parents=True, exist_ok=False)
    for name in ['corpus.json', 'oracle.json']:
        shutil.copy2(campaign / name, out / name)
    shutil.copy2(__file__, out / 'harness.py')
    env = dict(diag.environment(), DIAG_CORPUS=str(out / 'corpus.json'), DIAG_ALLOC='1',
               DIAG_PHASES='parse,decode,combined', DIAG_PROFILE=json.dumps(result['rust_profile']))
    report = {'status': 'pending', 'campaign': str(campaign), 'corpus_sha256': result['corpus_sha256'],
              'protocol': 'Separate instrumented binaries; cumulative allocation requests/bytes, not RSS. Timings excluded from performance claims.',
              'runs': {}}
    save = lambda: (out / 'results.json').write_text(json.dumps(report, indent=2) + '\n')
    save()
    try:
        for label in ['before', 'after']:
            build = json.loads((Path(result['builds'][label]) / 'build.json').read_text())
            assert build['status'] == 'passed' and build['profile'] == result['rust_profile']
            binary = Path(build['alloc']['path'])
            assert diag.sha(binary) == build['alloc']['sha256']
            dest = out / label
            shutil.copy2(binary, dest)
            run = subprocess.run([str(dest)], cwd=out, env=env, capture_output=True, text=True, timeout=360)
            (out / (label + '.stdout')).write_text(run.stdout)
            (out / (label + '.stderr')).write_text(run.stderr)
            record = {'binary': {'path': str(dest), 'sha256': diag.sha(dest)}, 'exit_code': run.returncode}
            report['runs'][label] = record
            save()
            assert run.returncode == 0, run.stderr
            raw = json.loads(run.stdout)
            assert raw['alloc_diag']
            diag.validate_result(raw, oracle, 'JsonTypedAst')
            record['raw'] = raw
            record['summary'] = {phase: {key: statistics.median(s[key] for s in data['samples'])
                                        for key in data['samples'][0] if key not in ['time_us', 'drop_us']}
                                 for phase, data in raw['phases'].items()}
            save()
            print(label, json.dumps(record['summary']), flush=True)
        report['status'] = 'passed'
        save()
    except BaseException as error:
        report['status'], report['failure'] = 'failed', repr(error)
        save()
        raise


if __name__ == '__main__':
    main()

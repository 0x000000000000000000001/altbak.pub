#!/usr/bin/env python3
"""Verify preserved reference builds, semantic checks and paired measurements."""
import argparse
import importlib.util
import json
from pathlib import Path
import statistics

HERE = Path(__file__).resolve().parent
spec = importlib.util.spec_from_file_location('reference_run', HERE / 'run.py')
RUN = importlib.util.module_from_spec(spec)
spec.loader.exec_module(RUN)


def verify_files(directory, files):
    for path, expected in files.items():
        assert RUN.sha(directory / path) == expected, directory / path


def verify_campaign(cycle, campaign, builds, lifetime_phases):
    sampled = lifecycle = 0
    for suite in RUN.SUITES:
        work = cycle / builds / suite
        manifest = RUN.check(work)
        oracle = RUN.load(work / 'expected.json')
        timed = oracle.get('timed_cases', oracle['modules'])
        out = cycle / campaign / suite
        result = RUN.load(out / 'results.json')
        assert result['build_manifest_sha256'] == RUN.sha(work / 'manifest.json')
        verify_files(out, result['files'])
        expected = oracle['fingerprints'] if suite == 'JsonTypedAst' else [
            oracle['fingerprints'][i] for i, case in enumerate(RUN.load(work / 'corpus.json')) if case['benchmark']]
        for backend, reports in result['runs'].items():
            assert len(reports) == 6
            for index, report in enumerate(reports):
                assert RUN.load(out / f'{index:02}-{backend}.json') == report
                if backend == 'go':
                    RUN.DIAG.validate_result(report, oracle, suite)
                    assert report['gomaxprocs'] == 1
                    assert report['phase_order'] == result['protocol']['phase_orders'][index]
                else:
                    RUN.DIAG.validate_result_c(report, oracle, suite, manifest['corpus'])
                    assert list(report['phases']) == result['protocol']['phase_orders'][index]
                for phase in RUN.PHASES:
                    observation = report['phases'][phase]
                    assert len(observation['samples']) == 5
                    assert observation['time_us'] == min(s['time_us'] for s in observation['samples'])
                    for sample in observation['samples']:
                        assert sample['time_us'] > 0
                        if backend == 'c':
                            assert sample['release_us'] >= 0
                            assert abs(sample['lifecycle_us'] - sample['time_us'] - sample['release_us']) < 0.00001
                    sampled += timed * len(observation['samples'])
            for phase in RUN.PHASES:
                assert result['medians_us'][backend][phase] == statistics.median(r['phases'][phase]['time_us'] for r in reports)
        for backend, reports in result['lifecycle_runs'].items():
            assert len(reports) == 6
            for index, report in enumerate(reports):
                assert RUN.load(out / f'lifecycle-{index:02}-{backend}.json') == report
                assert report['first_fingerprints'] == report['second_fingerprints'] == expected
                lifecycle += 2 * timed
            totals = [sum(r[p]['time_us'] for p in lifetime_phases) for r in reports]
            assert result['lifecycle_subtotals_us'][backend] == totals
            assert result['lifecycle_subtotal_medians_us'][backend] == statistics.median(totals)
    return {'sampled_outputs': sampled, 'lifecycle_outputs': lifecycle}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--cycle', type=Path, required=True)
    args = parser.parse_args()
    cycle = args.cycle.resolve()
    before = RUN.load(cycle / 'before.json')
    verify_files(cycle, before['sources'])
    for build in before['builds'].values():
        verify_files(cycle, build['files'])
    campaigns = {
        'initial-paired': verify_campaign(cycle, 'initial-paired', 'initial-owned',
            ['construction', 'first_combined', 'second_combined', 'release_collection']),
        'paired': verify_campaign(cycle, 'paired', 'owned', RUN.LIFETIME),
    }
    validations = {}
    for suite, driver in RUN.SUITES.items():
        work = cycle / 'owned' / suite
        manifest = RUN.check(work)
        assert RUN.sha(work / 'benchmark-go') == RUN.sha(cycle / 'previous' / suite / 'benchmark')
        for relative, expected in manifest['sources'].items():
            live = HERE.parent / Path(relative).relative_to('source')
            assert RUN.sha(live) == expected, live
        for prefix, build_prefix in [('initial-validation', 'initial-owned'), ('validation', 'owned')]:
            out = cycle / prefix / suite
            verified = RUN.load(out / 'verification.json')
            assert verified['status'] == 'passed'
            assert verified['build_manifest_sha256'] == RUN.sha(cycle / build_prefix / suite / 'manifest.json')
            verify_files(out, verified['files'])
            assert RUN.load(out / 'expected.json') == RUN.load(out / 'validate-c.json') == RUN.load(out / 'validate-sanitized.json')
            assert RUN.load(out / 'validate-c-differences.json') == []
            validations[prefix + '/' + suite] = {key: verified[key] for key in ['cases', 'successes', 'rejections']}
    result = {'status': 'passed', 'campaigns': campaigns, 'validations': validations,
              'sampled_outputs': sum(c['sampled_outputs'] for c in campaigns.values()),
              'lifecycle_outputs': sum(c['lifecycle_outputs'] for c in campaigns.values()),
              'sources': {str(p.relative_to(RUN.ROOT)): RUN.sha(p) for p in [
                  *[HERE.parent / name for name in RUN.SUITES.values()],
                  *sorted(p for p in HERE.iterdir() if p.is_file()),
                  RUN.ROOT / 'README.md', RUN.ROOT / 'bin/benchmark/json-diagnostic.py',
                  RUN.ROOT / 'docs/benchmark-results/2026-09-25-json-owned-references.md']}}
    (cycle / 'final-verification.json').write_text(json.dumps(result, indent=2) + '\n')
    print(json.dumps({key: result[key] for key in ['status', 'sampled_outputs', 'lifecycle_outputs']}))


if __name__ == '__main__':
    main()

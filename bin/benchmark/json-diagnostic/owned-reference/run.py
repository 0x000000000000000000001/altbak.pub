#!/usr/bin/env python3
"""Build, validate and pair ordinary-owned C++ references with preserved Go builds."""
import argparse
import hashlib
import importlib.util
import itertools
import json
from pathlib import Path
import shutil
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[3]
SUITES = {'JsonDecoding': 'decoding.cc', 'JsonTypedAst': 'typed-ast.cc'}
PHASES = ['parse', 'decode', 'combined']
LIFETIME = ['construction', 'first_combined', 'second_combined', 'retained_collection', 'release_collection']


def load(path):
    return json.loads(path.read_text())


def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()


def module(name, path):
    spec = importlib.util.spec_from_file_location(name, path)
    result = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(result)
    return result


DIAG = module('diagnostic', ROOT / 'bin/benchmark/json-diagnostic.py')
CASES = module('reference_cases', HERE / 'cases.py')


def run(command, log, env, cwd=None):
    with log.open('w') as stream:
        subprocess.run(list(map(str, command)), cwd=cwd, env=env,
                       stdout=stream, stderr=subprocess.STDOUT, check=True)


def build(cycle, suite):
    work = cycle / 'owned' / suite
    work.mkdir(parents=True)
    before = load(cycle / 'before.json')['builds'][suite]
    original = Path(before['workspace'])
    source = work / 'source'
    source.mkdir()
    for name in SUITES.values():
        shutil.copy2(HERE.parent / name, source / name)
    shutil.copytree(HERE, source / 'owned-reference', ignore=shutil.ignore_patterns('__pycache__'))
    env = DIAG.environment()
    prefix = DIAG.SIMDJSON_PREFIX
    flags = ['clang++', '-O3', '-std=c++17', '-Wno-deprecated-declarations',
             '-I' + str(prefix / 'include'), '-L' + str(prefix / 'lib'), '-lsimdjson']
    define = ['-DREFERENCE_TAST'] if suite == 'JsonTypedAst' else []
    commands = {}
    for binary, input_file, extra in [
        ('benchmark-c', source / SUITES[suite], []),
        ('lifecycle-c', source / 'owned-reference/lifecycle.cc', define),
        ('validate-c', source / 'owned-reference/validate.cc', define),
        ('validate-sanitized', source / 'owned-reference/validate.cc',
         define + ['-O1', '-g', '-fsanitize=address,undefined', '-fno-omit-frame-pointer']),
    ]:
        command = flags + extra + ['-o', str(work / binary), str(input_file)]
        run(command, work / (binary + '-build.log'), env)
        commands[binary] = command
    previous = cycle / 'previous' / suite
    assert sha(previous / 'benchmark') == load(previous / 'manifest.json')['binary_sha256']
    shutil.copy2(previous / 'benchmark', work / 'benchmark-go')
    shutil.copytree(original / 'output', work / 'output', ignore=shutil.ignore_patterns('*_test.go'))
    template = HERE / 'lifecycle.go'
    main = template.read_text().replace('// BUILD_DECODER',
        'decode := p.Get_Test_JsonDecoding_decodeText(); combined = func(text string) rt.Value { return rt.Apply(decode, rt.Str(text)) }')
    if suite == 'JsonTypedAst':
        main = main.replace('Test_JsonDecoding', 'Test_JsonTypedAst').replace('if file.Benchmark {', 'if true {')
    (work / 'output/main/main.go').write_text(main)
    command = ['go', 'build', '-pgo=off', '-o', str(work / 'lifecycle-go'), './main']
    run(command, work / 'lifecycle-go-build.log', dict(env, GOMAXPROCS='14'), work / 'output')
    commands['lifecycle-go'] = command
    corpus, info = DIAG.corpus_data(suite)
    (work / 'corpus.json').write_bytes(corpus)
    shutil.copy2(DIAG.fixtures(suite) / 'expected.json', work / 'expected.json')
    generated = {str(p.relative_to(work / 'output')): sha(p) for p in (work / 'output').rglob('*.go') if not p.name.endswith('_test.go')}
    original_go = {str(p.relative_to(original / 'output')): sha(p) for p in (original / 'output').rglob('*.go') if not p.name.endswith('_test.go')}
    assert {p: h for p, h in generated.items() if p != 'main/main.go'} == {p: h for p, h in original_go.items() if p != 'main/main.go'}
    manifest = {'suite': suite, 'go_workspace': str(original), 'go_manifest_sha256': sha(original / 'manifest.json'),
                'sources': {str(p.relative_to(work)): sha(p) for p in source.rglob('*') if p.is_file()},
                'commands': commands, 'binaries': {name: sha(work / name) for name in [*commands, 'benchmark-go']},
                'compiler': subprocess.check_output(['clang++', '--version'], text=True).strip(),
                'simdjson': {'path': str((prefix / 'lib/libsimdjson.dylib').resolve()), 'sha256': sha(prefix / 'lib/libsimdjson.dylib')},
                'generated_go': generated, 'corpus': info, 'corpus_sha256': sha(work / 'corpus.json'),
                'oracle_sha256': sha(work / 'expected.json'), 'go_lifecycle_template_sha256': sha(template)}
    (work / 'manifest.json').write_text(json.dumps(manifest, indent=2) + '\n')
    print(suite, 'built', flush=True)


def check(work):
    manifest = load(work / 'manifest.json')
    for name, expected in manifest['binaries'].items():
        assert sha(work / name) == expected, name
    for path, expected in manifest['sources'].items():
        assert sha(work / path) == expected, path
    for path, expected in manifest['generated_go'].items():
        assert sha(work / 'output' / path) == expected, path
    assert sha(Path(manifest['simdjson']['path'])) == manifest['simdjson']['sha256']
    assert sha(work / 'corpus.json') == manifest['corpus_sha256']
    assert sha(work / 'expected.json') == manifest['oracle_sha256']
    return manifest


def validate(cycle, suite):
    work = cycle / 'owned' / suite
    manifest = check(work)
    out = cycle / 'validation' / suite
    out.mkdir(parents=True)
    frozen = load(work / 'corpus.json')
    cases = CASES.typed(frozen) if suite == 'JsonTypedAst' else CASES.application(frozen)
    (out / 'corpus.json').write_text(json.dumps(cases) + '\n')
    env = dict(DIAG.environment(), DIAG_CORPUS=str(out / 'corpus.json'))
    run(['node', HERE / 'oracle.mjs', suite, Path(manifest['go_workspace']) / 'js/output',
         out / 'corpus.json', out / 'expected.json'], out / 'oracle.log', env)
    expected = load(out / 'expected.json')
    for name in ['validate-c', 'validate-sanitized']:
        with (out / (name + '.json')).open('w') as stdout, (out / (name + '.log')).open('w') as stderr:
            subprocess.run([work / name], env=env, stdout=stdout, stderr=stderr, check=True)
        actual = load(out / (name + '.json'))
        differences = [dict(expected=want, actual=got, input=cases[i])
                       for i, (want, got) in enumerate(zip(expected, actual)) if want != got]
        (out / (name + '-differences.json')).write_text(json.dumps(differences, indent=2) + '\n')
        assert len(actual) == len(expected) and not differences, (suite, name, len(differences))
    result = {'cases': len(cases), 'successes': sum(r['accepted'] for r in expected),
              'rejections': sum(not r['accepted'] for r in expected), 'status': 'passed',
              'build_manifest_sha256': sha(work / 'manifest.json'),
              'files': {p.name: sha(p) for p in out.iterdir() if p.is_file()}}
    (out / 'verification.json').write_text(json.dumps(result, indent=2) + '\n')
    print(suite, result['cases'], 'differential/ownership cases passed, including sanitizers', flush=True)


def measure(cycle, suite, campaign):
    work = cycle / 'owned' / suite
    manifest = check(work)
    assert load(cycle / 'validation' / suite / 'verification.json')['status'] == 'passed'
    out = cycle / campaign / suite
    out.mkdir(parents=True)
    env = dict(DIAG.environment(), DIAG_CORPUS=str(work / 'corpus.json'))
    oracle = load(work / 'expected.json')
    runs = {'go': [], 'c': []}
    lifetime = {'go': [], 'c': []}
    timed = oracle.get('timed_cases', oracle['modules'])
    expected = oracle['fingerprints'] if suite == 'JsonTypedAst' else [oracle['fingerprints'][i] for i, case in enumerate(load(work / 'corpus.json')) if case['benchmark']]
    for index, order in enumerate(itertools.permutations(PHASES)):
        for backend in (['go', 'c'] if index % 2 == 0 else ['c', 'go']):
            log = out / f'{index:02}-{backend}.json'
            run([work / ('benchmark-' + backend)], log, dict(env, DIAG_PHASES=','.join(order)))
            report = load(log)
            if backend == 'go': DIAG.validate_result(report, oracle, suite)
            else: DIAG.validate_result_c(report, oracle, suite, manifest['corpus'])
            assert set(report['phases']) == set(PHASES)
            if backend == 'go': assert report['gomaxprocs'] == 1
            runs[backend].append(report)
            print(suite, campaign, index, backend, round(report['phases']['combined']['time_us'] / 1000, 3), flush=True)
    for index in range(6):
        for backend in (['go', 'c'] if index % 2 == 0 else ['c', 'go']):
            log = out / f'lifecycle-{index:02}-{backend}.json'
            run([work / ('lifecycle-' + backend)], log, env)
            report = load(log)
            assert report['first_fingerprints'] == report['second_fingerprints'] == expected
            lifetime[backend].append(report)
    check(work)
    subtotals = {backend: [sum(run[phase]['time_us'] for phase in LIFETIME) for run in reports]
                 for backend, reports in lifetime.items()}
    result = {'suite': suite, 'build_manifest_sha256': sha(work / 'manifest.json'),
              'protocol': {'processes_per_backend': 6, 'warmups': 2, 'samples': 5,
                           'GOMAXPROCS': 1, 'GOGC': 100, 'phase_orders': list(itertools.permutations(PHASES)),
                           'cell': 'median of six process minima; full retained outputs',
                           'lifecycle': 'fresh process getter construction, two retained complete decodes, deferred decoder garbage collection, final release; fingerprints and their separate scratch-clearing GC excluded'},
              'runs': runs, 'lifecycle_runs': lifetime,
              'medians_us': {backend: {phase: statistics.median(run['phases'][phase]['time_us'] for run in reports)
                                      for phase in PHASES} for backend, reports in runs.items()},
              'lifecycle_phase_medians_us': {backend: {phase: statistics.median(run[phase]['time_us'] for run in reports)
                                                      for phase in LIFETIME} for backend, reports in lifetime.items()},
              'lifecycle_subtotals_us': subtotals,
              'lifecycle_subtotal_medians_us': {backend: statistics.median(values) for backend, values in subtotals.items()},
              'retained_sampled_outputs': 6 * 2 * 3 * 5 * timed,
              'lifecycle_outputs': 6 * 2 * 2 * timed,
              'files': {p.name: sha(p) for p in out.glob('*.json')}}
    result['combined_go_over_c'] = result['medians_us']['go']['combined'] / result['medians_us']['c']['combined']
    result['lifecycle_go_over_c'] = result['lifecycle_subtotal_medians_us']['go'] / result['lifecycle_subtotal_medians_us']['c']
    (out / 'results.json').write_text(json.dumps(result, indent=2) + '\n')
    print(suite, campaign, 'combined ratio', result['combined_go_over_c'], 'lifecycle ratio', result['lifecycle_go_over_c'], flush=True)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('action', choices=['build', 'validate', 'measure'])
    parser.add_argument('--cycle', type=Path, required=True)
    parser.add_argument('--suite', choices=SUITES)
    parser.add_argument('--campaign', default='paired')
    args = parser.parse_args()
    for suite in [args.suite] if args.suite else SUITES:
        if args.action == 'measure': measure(args.cycle.resolve(), suite, args.campaign)
        else: {'build': build, 'validate': validate}[args.action](args.cycle.resolve(), suite)


if __name__ == '__main__':
    main()

#!/usr/bin/env python3
"""Build or measure JSON decoding diagnostics in Go and JavaScript."""
import argparse
import gzip
import hashlib
import json
import os
from pathlib import Path
import shutil
import statistics
import subprocess
import sys

ROOT = Path(__file__).resolve().parents[2]
SOURCES = ROOT / 'src/Test'
FIXTURES = ROOT / 'test/fixtures/json-typed-ast'
SOURCE_FILES = [SOURCES / ('JsonTypedAst.' + ext) for ext in ['purs', 'go', 'js']]
COMPILER = ROOT.parent / 'gopurs/gopurs'
PBO = ROOT.parent / 'purescript-backend-optimizer-gopurs'
SUITES = ['JsonTypedAst', 'JsonDecoding']
# Native C/C++ references exist for the JSON diagnostics only; they are built
# next to the Go and JS artifacts and validated against the same frozen oracle.
C_DRIVERS = {'JsonDecoding': ROOT / 'bin/benchmark/json-diagnostic/decoding.cc',
             'JsonTypedAst': ROOT / 'bin/benchmark/json-diagnostic/typed-ast.cc'}
# The typed-AST reference decodes every module with the same type-table and
# annotation rules as the PureScript decoder, but does not reproduce the pure
# usage-validation pass (it returns Unit and cannot change the fingerprint).
C_SCOPES = {
    'JsonDecoding': ('simdjson parsing, hand-written arena decoding; successful fingerprints '
                     'validated exactly, malformed inputs only required to fail'),
    'JsonTypedAst': ('simdjson parsing, hand-written arena decoding of the full typed module; '
                     'all twelve fingerprints validated exactly; the pure usage-validation '
                     'pass is not reproduced'),
}
SIMDJSON_PREFIX = Path(os.environ.get('SIMDJSON_PREFIX', '/opt/homebrew/opt/simdjson'))
# The C reference validates successful decodes bit for bit. Malformed inputs
# only have to be rejected: the exact Argonaut messages are a PureScript
# library contract and the fixture excludes error cases from timing anyway.
C_EXTRA_SUCCESSES = {'optional-fields-missing', 'optional-fields-null'}
C_SEQUENCE = ['go', 'c', 'js', 'js', 'c', 'go', 'go', 'c', 'js']

def typed_purs():
    candidates = list((ROOT.parent / 'purescript/.stack-work/dist').glob('*/**/build/purs/purs'))
    if not candidates:
        raise ValueError('Build the local TAST fork before running this diagnostic')
    return max(candidates, key=lambda path: path.stat().st_mtime).resolve()

def source_files(suite):
    if suite == 'JsonTypedAst': return SOURCE_FILES
    return [SOURCES / (suite + '.' + ext) for ext in ['purs', 'go', 'js']]

def fixtures(suite):
    return FIXTURES if suite == 'JsonTypedAst' else ROOT / 'test/fixtures/json-decoding'

def dependencies(suite):
    shared = ['argonaut-core','argonaut-codecs','arrays','effect','either','integers','maybe','ordered-collections','partial','prelude','strings','tuples']
    return sorted(shared + (['backend-optimizer'] if suite == 'JsonTypedAst' else []))

def check_suite(manifest, suite):
    # Existing JsonTypedAst manifests predate the suite selector.
    if manifest.get('suite', 'JsonTypedAst') != suite:
        raise ValueError('Build belongs to a different JSON diagnostic suite')

def corpus_data(suite):
    directory = fixtures(suite)
    if suite == 'JsonTypedAst':
        return (gzip.decompress((directory/'tast-corpus.json.gz').read_bytes()),
                json.loads((directory/'tast-corpus.json').read_text()))
    raw = (directory/'corpus.json').read_bytes()
    entries = json.loads(raw)
    return raw, {'modules': len(entries), 'files': [
        {'name': item['name'], 'benchmark': item['benchmark'],
         'bytes': len(item['contents'].encode('utf-8'))} for item in entries]}

def build_c(work, env, suite):
    driver = C_DRIVERS.get(suite)
    if driver is None:
        return None
    binary = work / 'benchmark-c'
    compiler = env.get('CXX', 'clang++')
    command = [compiler, '-O3', '-std=c++17', f'-I{SIMDJSON_PREFIX}/include',
               f'-L{SIMDJSON_PREFIX}/lib', '-lsimdjson', '-o', binary, driver]
    call(command, work, work / 'logs/clang.log', env)
    version = subprocess.check_output([compiler, '--version'], env=env, text=True).strip().splitlines()[0]
    return {'binary_sha256': sha(binary), 'sources': {str(driver): sha(driver)},
            'command': [str(part) for part in command], 'compiler': version,
            'parser': f'simdjson at {SIMDJSON_PREFIX}', 'suite': suite}


def check_c(work, manifest, env, suite):
    if suite not in C_DRIVERS:
        raise ValueError(f'No native C reference for suite {suite}')
    section = manifest.get('c')
    if not section or section['suite'] != suite:
        raise ValueError('Build has no native C reference for this suite')
    for path, expected in section['sources'].items():
        if sha(Path(path)) != expected:
            raise ValueError(f'Stale C reference source: {path}')
    if sha(work / 'benchmark-c') != section['binary_sha256']:
        raise ValueError('Stale C reference binary')
    return section


def validate_result_c(result, oracle, suite, corpus_info):
    if result['modules'] != oracle['modules']:
        raise ValueError('Wrong C reference module count')
    if suite == 'JsonTypedAst':
        if result['fingerprints'] != oracle['fingerprints']:
            raise ValueError('C reference decoded a different typed AST')
        if result['json_fingerprints'] != oracle['json_fingerprints']:
            raise ValueError('C reference parsed different JSON')
        return
    if result['names'] != oracle['names'] or result['timed_cases'] != oracle['timed_cases']:
        raise ValueError('Wrong JSON decoding corpus order or timed case count')
    if result['json_fingerprints'] != oracle['json_fingerprints']:
        raise ValueError('C reference parsed different JSON')
    expected_success = {item['name'] for item in corpus_info['files'] if item['benchmark']} | C_EXTRA_SUCCESSES
    for name, got, want in zip(result['names'], result['fingerprints'], oracle['fingerprints']):
        if name in expected_success:
            if got != want:
                raise ValueError(f'C reference decoded differently: {name}')
        elif got is not None:
            raise ValueError(f'C reference accepted a malformed input: {name}')


def validate_result(result, oracle, suite):
    if result['modules'] != oracle['modules']: raise ValueError('Wrong corpus module count')
    if suite == 'JsonDecoding' and any(result[key] != oracle[key] for key in ['names', 'timed_cases']):
        raise ValueError('Wrong JSON decoding corpus order or timed case count')
    if any(result[key] != oracle[key] for key in ['fingerprints', 'json_fingerprints']):
        raise ValueError('Different decoded values or JSON across backends/runs')

def sha(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()

def fingerprint(suite='JsonTypedAst'):
    paths = list(source_files(suite))
    paths += list(fixtures(suite).glob('*')) + [Path(__file__), COMPILER / 'bin/gopurs-native', typed_purs()]
    # Native library FFI is read during generation, independently of the
    # compiler binary. A library change must invalidate a previous build.
    for source_root in ([PBO / 'src'] if suite == 'JsonTypedAst' else []) + [p / 'src' for p in COMPILER.parent.glob('gopurs-*') if (p / 'spago.yaml').is_file()]:
        paths += [p for p in source_root.rglob('*') if p.suffix in {'.purs', '.go', '.js'}]
    return {str(p): sha(p) for p in sorted(paths)}

def environment():
    env = {k:v for k,v in os.environ.items() if not k.startswith(('GOPURS_', 'NEUTRAL_')) and k not in ['PPROF','GODEBUG','GOMEMLIMIT','GOFLAGS','GOEXPERIMENT','NODE_OPTIONS']}
    env.update(GOMAXPROCS='1', GOGC='100', GOWORK='off')
    # The standard purs can have the same version number without exporting
    # TAST metadata. Select and fingerprint the actual local fork explicitly.
    env['PATH'] = str(typed_purs().parent) + os.pathsep + str(COMPILER/'node_modules/.bin') + os.pathsep + env['PATH']
    return env

def call(command, cwd, log, env):
    print(' '.join(map(str,command)), flush=True)
    with log.open('w') as out:
        result = subprocess.run(list(map(str,command)),cwd=cwd,env=env,stdout=out,stderr=subprocess.STDOUT)
    if result.returncode:
        print(log.read_text()[-10000:],file=sys.stderr)
        raise RuntimeError(f'Failed: {log}')

def copy_sources(destination, native=True, suite='JsonTypedAst'):
    for source in source_files(suite):
        if not native and source.suffix == '.go': continue
        target = destination / 'Test' / source.relative_to(SOURCES)
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(source, target)

def build_js(work, env, suite='JsonTypedAst'):
    js = work / 'js'
    (js / 'src').mkdir(parents=True, exist_ok=True)
    copy_sources(js / 'src', native=False, suite=suite)
    config = (work / 'spago.yaml').read_text().split('workspace:')[0] + '    - console\n'
    config += 'workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n'
    packages = [('st', COMPILER.parent/'gopurs-st'), ('unsafe-coerce', COMPILER.parent/'gopurs-unsafe-coerce')]
    if suite == 'JsonTypedAst': packages.insert(0, ('backend-optimizer', PBO))
    for name, path in packages:
        config += f'    {name}:\n      path: {json.dumps(str(path))}\n'
    (js / 'spago.yaml').write_text(config)
    call(['spago', 'build'], js, work/'logs/js-purs.log', env)
    (js/'entry.mjs').write_text(f"import {{ main }} from './output/Test.{suite}/index.js'; main();\n")
    call([COMPILER/'node_modules/.bin/esbuild', js/'entry.mjs', '--bundle', '--platform=node', '--format=esm', '--outfile='+str(work/'benchmark.mjs')], js, work/'logs/bundle.log', env)

def main():
    p=argparse.ArgumentParser(description=__doc__)
    p.add_argument('action',choices=['build','build-js','measure'])
    p.add_argument('--suite',choices=SUITES,default='JsonTypedAst')
    p.add_argument('--workspace',type=Path)
    p.add_argument('--output',type=Path)
    p.add_argument('--runtime', choices=['go', 'js'], help='measure only the selected runtime; builds retain both artifacts')
    p.add_argument('--resume-build',action='store_true',help='resume a failed build in this workspace')
    args=p.parse_args();suite=args.suite;env=environment()
    work=(args.workspace or ROOT/'var/benchmark'/('json-diagnostic' if suite == 'JsonTypedAst' else 'json-decoding')).resolve()
    if args.action in ['build', 'build-js']:
        env['GOMAXPROCS'] = '14'
    if args.action=='build-js':
        manifest=json.loads((work/'manifest.json').read_text())
        check_suite(manifest, suite)
        before=fingerprint(suite)
        allowed={str(Path(__file__)),str(SOURCES/(suite + '.js'))}
        if {k:v for k,v in before.items() if k not in allowed}!={k:v for k,v in manifest['sources'].items() if k not in allowed} or sha(work/'benchmark')!=manifest['binary_sha256']:
            raise ValueError('Native build no longer matches sources')
        build_js(work,env,suite)
        if before!=fingerprint(suite): raise ValueError('Sources changed during JS build')
        manifest.update(sources=before,js_sha256=sha(work/'benchmark.mjs'))
        (work/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')
        return
    if args.action=='build':
        if work.exists() and (not args.resume_build or (work/'manifest.json').exists()): raise ValueError('Use a new workspace, or --resume-build for a failed build')
        (work/'src').mkdir(parents=True,exist_ok=True);(work/'logs').mkdir(exist_ok=True)
        copy_sources(work / 'src', suite=suite)
        packages={'backend-optimizer':PBO} if suite == 'JsonTypedAst' else {}
        packages.update({p.name.removeprefix('gopurs-'):p for p in COMPILER.parent.glob('gopurs-*') if (p/'spago.yaml').is_file()})
        config='package:\n  name: json-diagnostic\n  dependencies:\n'
        for dep in dependencies(suite):
            config+='    - '+dep+'\n'
        config+='workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n'
        for name,path in sorted(packages.items()): config+=f'    {name}:\n      path: {json.dumps(str(path))}\n'
        (work/'spago.yaml').write_text(config)
        before=fingerprint(suite)
        call(['spago','build'],work,work/'logs/purs.log',env)
        for f in (work/'output').glob('*/corefn.json'):
            if 'typeTable' not in json.loads(f.read_text()): raise ValueError('Use the TAST fork of purs')
        call([COMPILER/'bin/gopurs-native','--main','Test.' + suite],work,work/'logs/backend.log',env)
        build_env=dict(env,GOMAXPROCS='14')
        call(['go','mod','tidy'],work/'output',work/'logs/modules.log',build_env)
        call(['go','build','-pgo=off','-o',work/'benchmark','./main'],work/'output',work/'logs/go.log',build_env)
        build_js(work,env,suite)
        c_build=build_c(work,env,suite)
        if before!=fingerprint(suite): raise ValueError('Source changed during build')
        manifest={'suite':suite,'sources':before,'binary_sha256':sha(work/'benchmark'),'js_sha256':sha(work/'benchmark.mjs'),
                  'c':c_build,
                  'go':subprocess.check_output(['go','version'],env=env,text=True).strip(),
                  'purs':subprocess.check_output(['purs','--version'],env=env,text=True).strip(),
                  'generated_tast':{str(f.relative_to(work)):sha(f) for f in sorted((work/'output').glob('*/corefn.json'))}}
        (work/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')
        print('Build complete; measurements are a separate command.',flush=True)
        return
    manifest=json.loads((work/'manifest.json').read_text())
    check_suite(manifest, suite)
    if manifest['sources']!=fingerprint(suite) or manifest['binary_sha256']!=sha(work/'benchmark') or manifest['js_sha256']!=sha(work/'benchmark.mjs'): raise ValueError('Stale or modified build')
    if not args.output: p.error('measure requires --output NEW_DIRECTORY')
    out=args.output.resolve();out.mkdir(parents=True,exist_ok=False)
    corpus, corpus_info=corpus_data(suite)
    (out/'corpus.json').write_bytes(corpus);env['DIAG_CORPUS']=str(out/'corpus.json')
    default_runtimes = ['go', 'js'] + (['c'] if suite in C_DRIVERS else [])
    runtimes = [args.runtime] if args.runtime else default_runtimes
    results = {runtime: [] for runtime in runtimes}
    c_reference = check_c(work, manifest, env, suite) if 'c' in runtimes else None
    oracle = json.loads((fixtures(suite)/'expected.json').read_text())
    for index,backend in enumerate(C_SEQUENCE):
        if backend not in runtimes: continue
        log=out/f'{index}-{backend}.log'
        if backend == 'c':
            call([work/'benchmark-c'], work, log, env)
        else:
            call([work/'benchmark'] if backend=='go' else ['node',work/'benchmark.mjs'],work,log,env)
        result=json.loads(log.read_text())
        if backend == 'c':
            validate_result_c(result, oracle, suite, corpus_info)
        else:
            validate_result(result, oracle, suite)
        results[backend].append(result)
        print(backend,{phase:round(data['time_us']/1000,3) for phase,data in result['phases'].items()},flush=True)
    if manifest['sources']!=fingerprint(suite): raise ValueError('Source changed during measurements')
    medians={b:{phase:statistics.median(run['phases'][phase]['time_us'] for run in runs) for phase in ['parse','decode','combined']} for b,runs in results.items()}
    allocated={phase:statistics.median(sample['allocated_bytes'] for run in results.get('go', []) for sample in run['phases'][phase]['samples']) for phase in ['parse','decode','combined']} if 'go' in results else {}
    report={'suite':suite,'protocol':{'processes_per_backend':3,'warmups_per_phase':2,'samples_per_phase':5,'cell':'median of process minimum times','GOMAXPROCS':1,'GOGC':100,'file_io_timed':False,'fingerprinting_timed':False},'corpus':corpus_info,'medians_us':medians,'go_allocated_bytes_per_corpus':allocated,'results':results,'build':manifest}
    if c_reference is not None:
        report['c_reference'] = {key: c_reference[key] for key in ['compiler', 'parser', 'command', 'binary_sha256']}
        report['c_reference']['scope'] = C_SCOPES[suite]
    (out/'results.json').write_text(json.dumps(report,indent=2)+'\n')
    (out/'corpus.json').unlink()  # Recreated from the versioned fixture each run.
    print(json.dumps(medians,indent=2))

if __name__=='__main__':main()

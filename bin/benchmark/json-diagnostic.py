#!/usr/bin/env python3
"""Build or measure JSON decoding diagnostics in Go, JavaScript, C and Rust.

The JsonTypedAst suite also supports the frozen gopurs238 typed-AST corpus
(238 corefn.json modules) selected with --corpus gopurs238. A non-fixture
corpus needs a frozen JS PureScript oracle (--oracle PATH), created
automatically on first use; its timings are ignored and each runtime is
validated against the JS fingerprints before timings are reported.
"""
import argparse
import gzip
import hashlib
import json
import os
from pathlib import Path
import re
import shutil
import statistics
import subprocess
import sys

ROOT = Path(__file__).resolve().parents[2]
SOURCES = ROOT / 'src/Test'
FIXTURES = ROOT / 'test/fixtures/json-typed-ast'
SOURCE_FILES = [SOURCES / ('JsonTypedAst.' + ext) for ext in ['purs', 'go', 'js', 'rs']]
COMPILER = ROOT.parent / 'gopurs/gopurs'
PBO = ROOT.parent / 'purescript-backend-optimizer-gopurs'
PBO_PURUST = ROOT.parent / 'purescript-backend-optimizer-purust'
RUST_BACKEND = ROOT.parent / 'purust/purust/bin/purust'
SUITES = ['JsonTypedAst', 'JsonDecoding']
# The frozen 238-module typed-AST corpus produced by the gopurs-aff benchmark;
# prepared.json records the per-file SHA-256 and module order of the original
# `output` tree, and its tare is verified before any measurement.
CORPUS_GOPURS238 = ROOT / 'var/benchmark/gopurs-purust-aff-20261002/inputs/gopurs-aff/output'
CORPUS_GOPURS238_MANIFEST = ROOT / 'var/benchmark/gopurs-purust-aff-20261002/prepared.json'
# Diagnostic Rust profile for JsonTypedAst, aligned with the native compiler
# configuration: `purust --threaded` (Arc), Cargo opt-level=3, LTO off, mimalloc.
# This is deliberately distinct from the generic README `rustc -O3 + thin LTO`
# row and is what the archived binaries and results must be read against.
RUST_TYPED_AST_PROFILE = {'opt_level': '3', 'debug': 'false', 'lto': 'false',
                          'codegen_units': '16', 'threaded': True,
                          'allocator': 'mimalloc',
                          'alloc_diagnostic': '--cfg diag_alloc, DIAG_ALLOC=1'}
# Native C/C++ references exist for the JSON diagnostics only; they are built
# next to the Go and JS artifacts and validated against the same frozen oracle.
C_DRIVERS = {'JsonDecoding': ROOT / 'bin/benchmark/json-diagnostic/decoding.cc',
             'JsonTypedAst': ROOT / 'bin/benchmark/json-diagnostic/typed-ast.cc'}
# Native references use ordinary owned containers, fresh parsers and separately
# reported result destruction. The typed reference resolves the complete table
# and validates lexical source usage before publishing a module.
C_SCOPES = {
    'JsonDecoding': ('fresh simdjson parsing, ordinary owned C++ strings/vectors/optionals; '
                      'successful fingerprints validated exactly, malformed inputs only '
                      'required to fail; final result destruction separately reported'),
    'JsonTypedAst': ('fresh simdjson parsing, ordinary owned C++ strings/vectors/tree nodes '
                      'and shared type nodes; eager whole-table resolution and lexical '
                      'source-usage validation; all twelve fingerprints validated exactly; '
                      'final result destruction separately reported'),
}
SIMDJSON_PREFIX = Path(os.environ.get('SIMDJSON_PREFIX', '/opt/homebrew/opt/simdjson'))
# The C reference validates successful decodes bit for bit. Malformed inputs
# only have to be rejected: the exact Argonaut messages are a PureScript
# library contract and the fixture excludes error cases from timing anyway.
C_EXTRA_SUCCESSES = {'optional-fields-missing', 'optional-fields-null'}
C_SEQUENCE = ['go', 'rust', 'c', 'js', 'js', 'c', 'rust', 'go', 'go', 'c', 'js', 'rust']

def typed_purs():
    candidates = list((ROOT.parent / 'purescript/.stack-work/dist').glob('*/**/build/purs/purs'))
    if not candidates:
        raise ValueError('Build the local TAST fork before running this diagnostic')
    return max(candidates, key=lambda path: path.stat().st_mtime).resolve()

def source_files(suite):
    if suite == 'JsonTypedAst': return SOURCE_FILES
    return [SOURCES / (suite + '.' + ext) for ext in ['purs', 'go', 'js', 'rs']]

def fixtures(suite):
    return FIXTURES if suite == 'JsonTypedAst' else ROOT / 'test/fixtures/json-decoding'

def dependencies(suite):
    shared = ['argonaut-core','argonaut-codecs','arrays','effect','either','integers','maybe','ordered-collections','partial','prelude','strings','tuples']
    return sorted(shared + (['backend-optimizer'] if suite == 'JsonTypedAst' else []))

def check_suite(manifest, suite):
    # Existing JsonTypedAst manifests predate the suite selector.
    if manifest.get('suite', 'JsonTypedAst') != suite:
        raise ValueError('Build belongs to a different JSON diagnostic suite')

def corpus_sha(data):
    return hashlib.sha256(data).hexdigest()

def gopurs238_corpus(source=None, manifest=None):
    source = Path(source) if source else CORPUS_GOPURS238
    manifest = Path(manifest) if manifest else CORPUS_GOPURS238_MANIFEST
    if not source.is_dir():
        raise ValueError(f'gopurs238 corpus directory not found: {source}')
    expected = None
    if manifest.is_file():
        info = json.loads(manifest.read_text()).get('tast')
        if info and info.get('files'):
            expected = info
    entries, seen = [], {}
    for directory in sorted(path for path in source.iterdir() if (path/'corefn.json').is_file()):
        raw = (directory/'corefn.json').read_bytes()
        value = json.loads(raw)
        name = '.'.join(value['moduleName']) if isinstance(value.get('moduleName'), list) else directory.name
        seen[name] = {'sha256': corpus_sha(raw), 'bytes': len(raw)}
        entries.append({'name': name, 'contents': raw.decode('utf-8')})
    corpus = json.dumps(entries).encode('utf-8')
    import_info = {'kind': 'gopurs238', 'source': str(source), 'modules': len(entries),
                   'bytes': sum(item['bytes'] for item in seen.values()),
                   'sha256': corpus_sha(corpus)}
    if expected is not None:
        if [item['module'] for item in expected['files']] != [item['name'] for item in entries]:
            raise ValueError('gopurs238 corpus module list differs from prepared.json')
        for item in expected['files']:
            got = seen.get(item['module'])
            if not got or got['sha256'] != item['sha256'] or got['bytes'] != item['bytes']:
                raise ValueError(f"gopurs238 corpus file changed: {item['module']}")
        aggregate = hashlib.sha256(''.join(
            f"{item['sha256']}  {item['module']}/corefn.json\n" for item in expected['files']).encode()).hexdigest()
        if aggregate != expected['sha256']:
            raise ValueError('gopurs238 prepared.json manifest is inconsistent')
        import_info.update({'manifest': str(manifest), 'manifest_sha256': expected['sha256'],
                            'manifest_modules': expected['modules'], 'manifest_bytes': expected['bytes'],
                            'manifest_types': expected['types']})
    return corpus, import_info

def corpus_data(suite, kind='fixture12', source=None, manifest=None):
    directory = fixtures(suite)
    if suite == 'JsonTypedAst':
        if kind == 'fixture12':
            data = gzip.decompress((directory/'tast-corpus.json.gz').read_bytes())
            info = json.loads((directory/'tast-corpus.json').read_text())
            info.update({'kind': 'fixture12', 'corpus_bytes': len(data), 'corpus_sha256': corpus_sha(data)})
            return data, info
        return gopurs238_corpus(source, manifest)
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
    # Resolve the runner path so a build recorded from any working directory
    # compares equal on later invocations.
    paths += list(fixtures(suite).glob('*')) + [Path(__file__).resolve(), COMPILER / 'bin/gopurs-native', typed_purs()]
    # Native library FFI is read during generation, independently of the
    # compiler binary. A library change must invalidate a previous build.
    for source_root in ([PBO / 'src'] if suite == 'JsonTypedAst' else []) + [p / 'src' for p in COMPILER.parent.glob('gopurs-*') if (p / 'spago.yaml').is_file()]:
        paths += [p for p in source_root.rglob('*') if p.suffix in {'.purs', '.go', '.js'}]
    # The Rust build of JsonTypedAst links the purust fork's real FFI; any
    # change to its sources or cargo sidecars must invalidate a previous
    # measurement artifact.
    if suite == 'JsonTypedAst':
        paths += [p for p in (PBO_PURUST / 'src').rglob('*')
                  if p.suffix in {'.purs', '.rs'} or p.name.endswith('.rs.cargo.json')]
    # The Rust runtime compiles the same sources against the purust ports. The
    # launcher forwards to purust-native, and cargo sidecars carry the FFI
    # crate metadata consumed while those ports are built.
    paths += [RUST_BACKEND, RUST_BACKEND.with_suffix('.js'), RUST_BACKEND.parent / 'purust-native']
    for source_root in [p / 'src' for p in (ROOT.parent / 'purust').glob('purust-*') if (p / 'spago.yaml').is_file()]:
        paths += [p for p in source_root.rglob('*')
                  if p.suffix in {'.purs', '.rs'} or p.name.endswith('.rs.cargo.json')]
    return {str(p): sha(p) for p in sorted(set(paths)) if p.is_file()}

def environment():
    # Inherited overrides for the compilers/runtimes are removed so the
    # declared build profile is the one actually used, not an ambient override.
    env = {k:v for k,v in os.environ.items() if not k.startswith(('GOPURS_', 'NEUTRAL_', 'DIAG_', 'PURUST_', 'RUSTFLAGS', 'CARGO_PROFILE_')) and k not in ['PPROF','GODEBUG','GOMEMLIMIT','GOFLAGS','GOEXPERIMENT','NODE_OPTIONS','CARGO_TARGET_DIR','RUSTC','RUSTC_WRAPPER','RUSTC_WORKSPACE_WRAPPER']}
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
        # Rust FFI is consumed by the purust build only.
        if source.suffix == '.rs': continue
        if not native and source.suffix == '.go': continue
        target = destination / 'Test' / source.relative_to(SOURCES)
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copy2(source, target)

def rust_packages():
    packages = {}
    for directory in (ROOT.parent / 'purust').glob('purust-*'):
        config = directory / 'spago.yaml'
        if not config.is_file(): continue
        match = re.search(r'name:\s*"?([A-Za-z0-9_.-]+)"?\s*$', config.read_text(), re.MULTILINE)
        if match: packages[match.group(1)] = directory
    return packages


def patch_typed_ast_allocator(project):
    # The allocation diagnostic reuses the generated project but swaps the
    # global allocator for the counting wrapper behind --cfg diag_alloc. The
    # timed build keeps mimalloc untouched. Idempotent on purpose.
    main = project / 'src/main.rs'
    text = main.read_text()
    if 'CountingAlloc' in text:
        return
    marker = '#[global_allocator]\nstatic GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;'
    if marker not in text:
        raise ValueError('purust main.rs no longer contains the expected allocator marker')
    replacement = ('#[cfg(not(diag_alloc))]\n'
                   '#[global_allocator]\n'
                   'static GLOBAL: mimalloc::MiMalloc = mimalloc::MiMalloc;\n'
                   '#[cfg(diag_alloc)]\n'
                   '#[global_allocator]\n'
                   'static GLOBAL: Purs_Test_JsonTypedAst::diag_alloc::CountingAlloc<mimalloc::MiMalloc> = '
                   'Purs_Test_JsonTypedAst::diag_alloc::CountingAlloc(mimalloc::MiMalloc);')
    main.write_text(text.replace(marker, replacement))


def build_rust(work, env, suite):
    rust = work / 'rust'
    (rust / 'src/Test').mkdir(parents=True, exist_ok=True)
    for source in source_files(suite):
        # purs still requires the JS FFI module to typecheck; purust resolves
        # the Rust implementation next to the source.
        if source.suffix not in ['.purs', '.rs', '.js']: continue
        shutil.copy2(source, rust / 'src/Test' / source.name)
    config = 'package:\n  name: json-diagnostic-rust\n  dependencies:\n'
    for dep in dependencies(suite): config += '    - ' + dep + '\n'
    config += 'workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n'
    packages = rust_packages()
    if suite == 'JsonTypedAst':
        # The Rust runtime must link the purust fork's real FFI, not the
        # gopurs fork used by the native Go reference.
        packages['backend-optimizer'] = PBO_PURUST
    for name, path in sorted(packages.items()):
        config += f'    {name}:\n      path: {json.dumps(str(path))}\n'
    (rust / 'spago.yaml').write_text(config)
    call(['spago', 'build'], rust, work/'logs/rust-purs.log', env)
    project = rust / 'rust-project'
    command = [RUST_BACKEND, '--main', 'Test.' + suite, '--source', 'output', '--out', str(project)]
    if suite == 'JsonTypedAst':
        # Align the diagnostic with the native compiler configuration (Arc).
        command.append('--threaded')
    call(command, rust, work/'logs/rust-purust.log', env)
    rust_env = dict(env, CARGO_PROFILE_RELEASE_OPT_LEVEL='3', CARGO_PROFILE_RELEASE_DEBUG='false')
    section = {}
    if suite == 'JsonTypedAst':
        rust_env.update(CARGO_PROFILE_RELEASE_LTO='false', CARGO_PROFILE_RELEASE_CODEGEN_UNITS='16')
        section['profile'] = RUST_TYPED_AST_PROFILE
        section['cargo_env'] = {'CARGO_PROFILE_RELEASE_OPT_LEVEL': '3',
                                'CARGO_PROFILE_RELEASE_DEBUG': 'false',
                                'CARGO_PROFILE_RELEASE_LTO': 'false',
                                'CARGO_PROFILE_RELEASE_CODEGEN_UNITS': '16'}
        patch_typed_ast_allocator(project)
    call(['cargo', 'build', '--release', '--manifest-path', str(project / 'Cargo.toml'),
          '--bin', 'purust_output'], rust, work/'logs/rust-cargo.log', rust_env)
    binary = project / 'target/release/purust_output'
    if suite == 'JsonTypedAst':
        alloc_target = project / 'target-alloc'
        alloc_env = dict(rust_env, RUSTFLAGS='--cfg diag_alloc', CARGO_TARGET_DIR=str(alloc_target))
        call(['cargo', 'build', '--release', '--manifest-path', str(project / 'Cargo.toml'),
              '--bin', 'purust_output'], rust, work/'logs/rust-cargo-alloc.log', alloc_env)
        frozen_alloc = project / 'target/release/purust_output-alloc'
        shutil.copy2(alloc_target / 'release/purust_output', frozen_alloc)
        section.update({'alloc_binary': str(frozen_alloc), 'alloc_binary_sha256': sha(frozen_alloc),
                        'alloc_rustflags': '--cfg diag_alloc', 'alloc_target': str(alloc_target),
                        'alloc_cargo_env': {'CARGO_PROFILE_RELEASE_OPT_LEVEL': '3',
                                            'CARGO_PROFILE_RELEASE_DEBUG': 'false',
                                            'CARGO_PROFILE_RELEASE_LTO': 'false',
                                            'CARGO_PROFILE_RELEASE_CODEGEN_UNITS': '16',
                                            'RUSTFLAGS': '--cfg diag_alloc'}})
    sources = {str(path): sha(path) for path in sorted((rust/'src/Test').glob('*'))}
    sources[str(rust/'spago.yaml')] = sha(rust/'spago.yaml')
    result = {'suite': suite, 'binary_sha256': sha(binary), 'binary': str(binary), 'sources': sources,
              'rustc': subprocess.check_output(['rustc', '--version'], env=env, text=True).strip()}
    result.update(section)
    return result


def check_rust(work, manifest, suite):
    section = manifest.get('rust')
    if not section or section['suite'] != suite:
        raise ValueError('Build has no Rust runtime for this suite')
    for path, expected in section['sources'].items():
        if sha(Path(path)) != expected:
            raise ValueError(f'Stale Rust source: {path}')
    if sha(work / 'rust/rust-project/target/release/purust_output') != section['binary_sha256']:
        raise ValueError('Stale Rust binary')
    if section.get('alloc_binary_sha256') and sha(Path(section['alloc_binary'])) != section['alloc_binary_sha256']:
        raise ValueError('Stale Rust allocation diagnostic binary')
    return section


def build_js(work, env, suite='JsonTypedAst'):
    js = work / 'js'
    (js / 'src').mkdir(parents=True, exist_ok=True)
    copy_sources(js / 'src', native=False, suite=suite)
    config = (work / 'spago.yaml').read_text().split('workspace:')[0] + '    - console\n'
    config += 'workspace:\n  packageSet:\n    registry: 77.10.1\n  extraPackages:\n'
    # Exercise the same modified codec interface in both backends. Registry
    # codecs would silently omit the local Record.js implementation.
    packages = [('argonaut-codecs', COMPILER.parent/'gopurs-argonaut-codecs'),
                ('st', COMPILER.parent/'gopurs-st'), ('unsafe-coerce', COMPILER.parent/'gopurs-unsafe-coerce')]
    if suite == 'JsonTypedAst': packages.insert(0, ('backend-optimizer', PBO))
    for name, path in packages:
        config += f'    {name}:\n      path: {json.dumps(str(path))}\n'
    (js / 'spago.yaml').write_text(config)
    call(['spago', 'build'], js, work/'logs/js-purs.log', env)
    (js/'entry.mjs').write_text(f"import {{ main }} from './output/Test.{suite}/index.js'; main();\n")
    call([COMPILER/'node_modules/.bin/esbuild', js/'entry.mjs', '--bundle', '--platform=node', '--format=esm', '--outfile='+str(work/'benchmark.mjs')], js, work/'logs/bundle.log', env)

def run_capture(command, cwd, env):
    result = subprocess.run([str(part) for part in command], cwd=cwd, env=env,
                            capture_output=True, text=True)
    if result.returncode:
        raise RuntimeError(f'Oracle run failed: {result.stderr[-4000:]}')
    return result.stdout


def load_or_create_oracle(path, work, env, corpus, corpus_info):
    corpus_hash = corpus_info.get('sha256') or corpus_info.get('corpus_sha256')
    if path is not None and path.is_file():
        oracle = json.loads(path.read_text())
        if oracle.get('corpus_sha256') != corpus_hash:
            raise ValueError('Frozen oracle belongs to a different corpus')
        if oracle.get('modules') != corpus_info['modules']:
            raise ValueError('Frozen oracle has a different module count')
        return oracle
    if path is None:
        raise ValueError('--oracle PATH is required for non-fixture corpora')
    # Authority: the unmodified JS PureScript program. Its timings are ignored;
    # only fingerprints are frozen, and every runtime is validated against
    # them before campaign timings are reported.
    raw = run_capture(['node', work/'benchmark.mjs'], work, env)
    result = json.loads(raw)
    oracle = {'authority': 'js-purescript', 'corpus': corpus_info['kind'],
              'corpus_sha256': corpus_hash, 'modules': result['modules'],
              'fingerprints': result['fingerprints'],
              'json_fingerprints': result['json_fingerprints'],
              'authority_program_sha256': sha(work/'benchmark.mjs'),
              'timings_ignored': True}
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(oracle, indent=2)+'\n')
    print(f'Wrote JS authority oracle: {path}', flush=True)
    return oracle


def main():
    p=argparse.ArgumentParser(description=__doc__)
    p.add_argument('action',choices=['build','build-js','measure'])
    p.add_argument('--suite',choices=SUITES,default='JsonTypedAst')
    p.add_argument('--workspace',type=Path)
    p.add_argument('--output',type=Path)
    p.add_argument('--runtime', choices=['go', 'js', 'c', 'rust'], help='measure only the selected runtime; builds retain both artifacts')
    p.add_argument('--runtimes', help='comma-separated runtimes measured in the shared protocol order, e.g. go,js,rust')
    p.add_argument('--corpus', choices=['fixture12', 'gopurs238'], default='fixture12', help='typed-AST corpus; gopurs238 is the frozen 238-module corpus')
    p.add_argument('--corpus-source', type=Path, help='override the gopurs238 output directory')
    p.add_argument('--corpus-manifest', type=Path, help='override the prepared.json verifying gopurs238')
    p.add_argument('--oracle', type=Path, help='frozen JS authority for non-fixture corpora; created on first use')
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
        rust_build=build_rust(work,env,suite)
        if before!=fingerprint(suite): raise ValueError('Source changed during build')
        manifest={'suite':suite,'sources':before,'binary_sha256':sha(work/'benchmark'),'js_sha256':sha(work/'benchmark.mjs'),
                  'c':c_build, 'rust':rust_build,
                  'go':subprocess.check_output(['go','version'],env=env,text=True).strip(),
                  'purs':subprocess.check_output(['purs','--version'],env=env,text=True).strip(),
                  'generated_tast':{str(f.relative_to(work)):sha(f) for f in sorted((work/'output').glob('*/corefn.json'))}}
        if suite == 'JsonTypedAst':
            manifest['driver_capabilities'] = ['phase-order']
        (work/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')
        print('Build complete; measurements are a separate command.',flush=True)
        return
    manifest=json.loads((work/'manifest.json').read_text())
    check_suite(manifest, suite)
    if manifest['sources']!=fingerprint(suite) or manifest['binary_sha256']!=sha(work/'benchmark') or manifest['js_sha256']!=sha(work/'benchmark.mjs'): raise ValueError('Stale or modified build')
    if not args.output: p.error('measure requires --output NEW_DIRECTORY')
    if suite == 'JsonDecoding' and args.corpus != 'fixture12': p.error('alternate corpora are only available for JsonTypedAst')
    out=args.output.resolve();out.mkdir(parents=True,exist_ok=False)
    corpus, corpus_info=corpus_data(suite, args.corpus, args.corpus_source, args.corpus_manifest)
    corpus_path = out/'corpus.json'
    corpus_path.write_bytes(corpus);env['DIAG_CORPUS']=str(corpus_path)
    env['DIAG_RUSTC'] = manifest.get('rust',{}).get('rustc','')
    if suite == 'JsonTypedAst' and manifest.get('rust',{}).get('profile'):
        env['DIAG_PROFILE'] = json.dumps({**manifest['rust']['profile'], 'rustc': manifest['rust']['rustc']})
    default_runtimes = ['go', 'js'] + (['c'] if suite in C_DRIVERS else []) + (['rust'] if manifest.get('rust') else [])
    if args.runtime: runtimes=[args.runtime]
    elif args.runtimes: runtimes=[name.strip() for name in args.runtimes.split(',') if name.strip()]
    else: runtimes=default_runtimes
    unknown=[name for name in runtimes if name not in default_runtimes]
    if unknown: p.error('runtimes not available in this build: ' + ','.join(unknown))
    results = {runtime: [] for runtime in runtimes}
    c_reference = check_c(work, manifest, env, suite) if 'c' in runtimes else None
    if 'rust' in runtimes: check_rust(work, manifest, suite)
    if suite == 'JsonTypedAst' and args.corpus != 'fixture12':
        oracle = load_or_create_oracle(args.oracle, work, env, corpus, corpus_info)
    else:
        oracle = json.loads((fixtures(suite)/'expected.json').read_text())
    for index,backend in enumerate(C_SEQUENCE):
        if backend not in runtimes: continue
        log=out/f'{index}-{backend}.log'
        if backend == 'c':
            call([work/'benchmark-c'], work, log, env)
        elif backend == 'rust':
            call([work/'rust/rust-project/target/release/purust_output'], work, log, env)
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
    protocol={'processes_per_backend':3,'warmups_per_phase':2,'samples_per_phase':5,
              'cell':'median of process minimum times','runtimes':runtimes,
              'GOMAXPROCS':1,'GOGC':100,'file_io_timed':False,'fingerprinting_timed':False,
              'destruction_timed_separately':[name for name in runtimes if name in ('rust', 'c')]}
    report={'suite':suite,'corpus':{**corpus_info,'kind':args.corpus},'protocol':protocol,
            'medians_us':medians,'go_allocated_bytes_per_corpus':allocated,
            'results':results,'build':manifest}
    if suite == 'JsonTypedAst' and manifest.get('rust',{}).get('profile'):
        report['rust_profile'] = manifest['rust']['profile']
    if suite == 'JsonTypedAst' and args.corpus != 'fixture12':
        report['oracle'] = {key:value for key,value in oracle.items()
                            if key not in ['fingerprints','json_fingerprints']}
    if c_reference is not None:
        report['c_reference'] = {key: c_reference[key] for key in ['compiler', 'parser', 'command', 'binary_sha256']}
        report['c_reference']['scope'] = C_SCOPES[suite]
    (out/'results.json').write_text(json.dumps(report,indent=2)+'\n')
    if args.corpus == 'fixture12':
        corpus_path.unlink()  # Recreated from the versioned fixture each run.
    print(json.dumps(medians,indent=2))

if __name__=='__main__':main()

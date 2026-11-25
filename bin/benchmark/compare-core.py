#!/usr/bin/env python3
"""Measure frozen Go, official JS and a declared C core reference together.

Build all artifacts first. The C binary requires a build sidecar (by default
<binary>.manifest.json) with binary_sha256, source_sha256, toolchain, commands
and profile. This runner never compiles or republishes the historical catalogue.
"""
import argparse
from datetime import datetime, timezone
import hashlib
import json
import math
import os
from pathlib import Path
import platform
import shutil
import subprocess
import sys

sys.dont_write_bytecode = True
from measure import aggregate, check
from publish import ROWS, format_number, verify
from validate import CASES

ROOT = Path(__file__).resolve().parents[2]
C_REFERENCES = {
    'functional': {'backend': 'c-functional', 'label': 'C fonctionnel',
                   'source': 'tmp/bench_functional.c',
                   'contract': 'tmp/fp_reference_contract.md'},
    'imperative': {'backend': 'c-imperative', 'label': 'C impératif',
                   'source': 'tmp/bench_native.c',
                   'contract': 'tmp/imperative_reference_contract.md'},
}
BACKENDS = ('go', 'js', 'c-functional')
CORE_SOURCES = ['src/App.purs', 'src/Bench.purs'] + [
    f'src/Test/{module}/{module}.purs' for module in CASES]
PROTOCOL = {
    'processes_per_backend': 3, 'timed_batches_per_case': 10,
    'global_warmups': 3, 'local_warmups': 3,
    'calibration_target_ms': 10, 'maximum_batch_iterations': 16777216,
    'aggregation': 'median of the three process minima per case',
    'order': [list(BACKENDS[i:] + BACKENDS[:i]) for i in range(3)],
    'scope': '14 Stable core cases; no library or excluded diagnostics',
}


def c_reference(kind):
    if kind not in C_REFERENCES:
        raise ValueError(f'Unknown C reference kind: {kind}')
    return C_REFERENCES[kind]


def backends_for(kind):
    return ('go', 'js', c_reference(kind)['backend'])


def protocol_for(kind):
    backends = backends_for(kind)
    return {**PROTOCOL,
            'order': [list(backends[i:] + backends[:i]) for i in range(3)]}


def sha(path):
    return hashlib.sha256(Path(path).read_bytes()).hexdigest()


def timestamp():
    return datetime.now(timezone.utc).isoformat()


def load_average():
    return list(os.getloadavg()) if hasattr(os, 'getloadavg') else None


def json_file(path):
    return json.loads(Path(path).read_text())


def write_json(path, value):
    Path(path).write_text(json.dumps(value, ensure_ascii=False, indent=2,
                                    allow_nan=False) + '\n')


def verify_hashes(hashes, base=None):
    """Fail closed on changed/missing files, including build manifests."""
    for name, expected in hashes.items():
        path = Path(name) if base is None else Path(base) / name
        if not path.is_file() or sha(path) != expected:
            raise ValueError(f'Frozen input changed or is missing: {path}')


def verify_sources(directory, stored, names):
    for name in names:
        if name not in stored:
            raise ValueError(f'Build manifest omits source: {name}')
        verify_hashes({name: stored[name]}, ROOT)
        verify_hashes({name: stored[name]}, directory)


def inspect_go(directory):
    directory = Path(directory).resolve()
    path = directory / 'manifest.json'
    manifest = json_file(path)
    expected = {'language': 'go', 'mode': 'pure', 'test': None,
                'entry': 'App', 'include_library': False}
    if any(manifest.get(key) != value for key, value in expected.items()):
        raise ValueError('Expected a frozen Go App build containing exactly 14 core cases')
    profile = manifest['profile']
    if (profile.get('backend_profile') != 'release' or profile.get('pgo') is not False
            or str(profile.get('GOGC')) != '800'):
        raise ValueError('Go must be Release, GOGC=800, without PGO')
    names = CORE_SOURCES + ['src/Bench.go']
    verify_sources(directory, manifest['inputs'], names)
    binary = directory / 'benchmark'
    verify_hashes({str(binary): manifest['binary_sha256']})
    return {'manifest': manifest, 'manifest_path': str(path),
            'manifest_sha256': sha(path), 'command': [str(binary)],
            'cwd': str(directory), 'profile': profile,
            'toolchain': manifest['toolchain'],
            'payload_sha256': {str(binary): manifest['binary_sha256']},
            'source_sha256': {name: manifest['inputs'][name] for name in names}}


def js_payload(directory):
    output = Path(directory) / 'output'
    return {str(path.resolve()): sha(path) for path in sorted(output.rglob('*'))
            if path.is_file() and path.suffix in {'.js', '.json'}}


def inspect_js(directory):
    directory = Path(directory).resolve()
    path = directory / 'manifest.json'
    manifest = json_file(path)
    expected = {'backend': 'js', 'mode': 'pure', 'test': None,
                'entry': 'App', 'include_library': False}
    if any(manifest.get(key) != value for key, value in expected.items()):
        raise ValueError('Expected an official JavaScript App build with exactly 14 core cases')
    names = CORE_SOURCES + ['src/Bench.js']
    verify_sources(directory, manifest['source_sha256'], names)
    run = manifest['run']
    expected_run = ['node', '--input-type=module', '-e',
                    "import { main } from './output/App/index.js'; main();"]
    if run != expected_run:
        raise ValueError('Unexpected JavaScript runtime command; build the official App entry')
    node = shutil.which('node')
    if not node:
        raise ValueError('Node.js runtime is unavailable')
    node = str(Path(node).resolve())
    payload = js_payload(directory)
    if str(directory / 'output/App/index.js') not in payload:
        raise ValueError('JavaScript entry payload is missing')
    payload[node] = sha(node)
    compilers = [command[0] for command in manifest['commands']
                 if len(command) > 1 and command[1] == 'compile']
    if len(compilers) != 1:
        raise ValueError('Official PureScript compiler command is missing or ambiguous')
    compiler = str(Path(compilers[0]).resolve())
    compiler_version = subprocess.check_output([compiler, '--version'], text=True).strip()
    node_version = subprocess.check_output([node, '--version'], text=True).strip()
    return {'manifest': manifest, 'manifest_path': str(path),
            'manifest_sha256': sha(path), 'command': [node, *run[1:]],
            'cwd': str(directory), 'profile': {'compiler': 'official PureScript',
                                             'NODE_OPTIONS': os.environ.get('NODE_OPTIONS', '')},
            'toolchain': {'node': node_version, 'purescript': compiler_version,
                          'purescript_sha256': sha(compiler)},
            'payload_sha256': payload,
            'source_sha256': {name: manifest['source_sha256'][name] for name in names}}


def inspect_c(binary, manifest_path, source):
    binary, path, source = map(lambda p: Path(p).resolve(), (binary, manifest_path, source))
    manifest = json_file(path)
    source_name = str(source.relative_to(ROOT))
    if not manifest.get('commands') or not manifest.get('toolchain') or not manifest.get('profile'):
        raise ValueError('C build manifest must record commands, toolchain and profile')
    if source_name not in manifest.get('source_sha256', {}):
        raise ValueError(f'C build manifest omits {source_name}')
    verify_hashes(manifest['source_sha256'], ROOT)
    verify_hashes({str(binary): manifest['binary_sha256']})
    return {'manifest': manifest, 'manifest_path': str(path),
            'manifest_sha256': sha(path), 'command': [str(binary)],
            'cwd': str(binary.parent), 'profile': manifest['profile'],
            'toolchain': manifest['toolchain'],
            'payload_sha256': {str(binary): manifest['binary_sha256']},
            'source_sha256': manifest['source_sha256']}


def source_snapshot(c_source, c_kind='functional'):
    names = CORE_SOURCES + ['src/Bench.go', 'src/Bench.js',
                           str(Path(c_source).resolve().relative_to(ROOT)),
                           c_reference(c_kind)['contract'],
                           'bin/benchmark/compare-core.py', 'bin/benchmark/measure.py',
                           'bin/benchmark/validate.py', 'bin/benchmark/publish.py']
    catalogue = ROOT / 'bin/benchmark/ratio-measurements.json'
    if catalogue.is_file():
        names.append(str(catalogue.relative_to(ROOT)))
    return {name: sha(ROOT / name) for name in names}


def verify_frozen(builds, sources):
    verify_hashes(sources, ROOT)
    for key, build in builds.items():
        verify_hashes({build['manifest_path']: build['manifest_sha256']})
        verify_hashes(build['payload_sha256'])
        if key == 'js':
            current = js_payload(build['cwd'])
            expected = {name: value for name, value in build['payload_sha256'].items()
                        if Path(name).is_relative_to(Path(build['cwd']).resolve() / 'output')}
            if current != expected:
                raise ValueError('JavaScript payload inventory changed during the campaign')


def check_core(output):
    checked = check(output, 'pure')
    if 'library' in checked or len(checked['times_us']) != 14:
        raise ValueError('This comparison requires exactly 14 core cases')
    return checked


def rankings(results, c_kind='functional'):
    c_backend = c_reference(c_kind)['backend']
    if set(results) != set(backends_for(c_kind)):
        raise ValueError('All three complete backend columns are required')
    for key, result in results.items():
        verify(key, result)
    rows = []
    for index, (name, module) in enumerate(zip(ROWS, CASES)):
        times = {key: result['times_us'][index] for key, result in results.items()}
        samples = {key: [sample['times_us'][index] for sample in result['process_runs']]
                   for key, result in results.items()}
        variability = {}
        for reference, key in [(c_backend, 'go_c'), ('js', 'go_js')]:
            ratios = [go / other for go, other in zip(samples['go'], samples[reference])]
            variability[key] = {'ratios_by_cycle': ratios, 'minimum': min(ratios),
                                'maximum': max(ratios),
                                'interpretation': 'Observed ratios within each rotating cycle; '
                                                  'not a confidence interval.'}
        rows.append({'case': name, 'module': f'Test.{module}',
                     'label': results['go']['labels'][index],
                     'expected': results['go']['values'][index], 'times_us': times,
                     'process_samples_us': samples, 'ratio_variability': variability,
                     'go_over_c': times['go'] / times[c_backend],
                     'go_over_js': times['go'] / times['js']})
    return {'go_c': sorted(rows, key=lambda row: (-row['go_over_c'], row['case'])),
            'go_js': sorted(rows, key=lambda row: (-row['go_over_js'], row['case']))}


def historical_comparison(results):
    path = ROOT / 'bin/benchmark/ratio-measurements.json'
    if not path.is_file():
        return None
    catalogue = json_file(path)
    columns = catalogue.get('columns', {})
    rows = []
    for index, name in enumerate(ROWS):
        old = {key: columns.get(key, {}).get(name) for key in ['go-pure', 'js-pure', 'c']}
        values = {}
        for old_key, fresh_key in [('go-pure', 'go'), ('js-pure', 'js')]:
            item = old[old_key]
            if item and math.isfinite(item['time_us']) and item['time_us'] > 0:
                values[fresh_key] = {'archived': item,
                                    'fresh_over_archived': results[fresh_key]['times_us'][index] / item['time_us']}
        rows.append({'case': name, 'program_timings': values,
                     'c_archived': old['c']})
    return {'catalogue': str(path.relative_to(ROOT)), 'sha256': sha(path),
            'note': 'Historical observations only, not paired measurements. '
                    'Archived C timings retain their recorded source identity; '
                    'they may describe a different C implementation.',
            'cases': rows}


def render_report(report):
    kind = report.get('c_reference', {}).get('kind', 'functional')
    reference = c_reference(kind)
    if kind == 'functional':
        scope = ('La référence C traduit la structure fonctionnelle des cas. Les backends restent libres '
                 'd’optimiser : une forme close légitime est un résultat de compilation, pas une erreur. ')
        memory = ('Le C alloue via `malloc` avec un registre et libère les allocations dans chaque action '
                  'chronométrée ; aucune arène préallouée hors mesure n’est utilisée. ')
        verification = ('La correction des oracles ne prouve pas à elle seule '
                        'l’équivalence structurelle : le contrat et le source C doivent être lus avec ces résultats.')
    else:
        scope = ('La référence C est impérative et optimisée. Ses algorithmes et structures mutables '
                 'peuvent différer de ceux des sources PureScript, avec les mêmes entrées et résultats. '
                 'Les entrées restent opaques à l’optimisation ; les simplifications algébriques '
                 'qui calculent encore le résultat depuis l’entrée d’exécution restent permises. ')
        memory = ('Les allocations et libérations nécessaires au C sont incluses dans chaque action '
                  'chronométrée ; aucune arène préallouée hors mesure n’est utilisée. ')
        verification = ('Les oracles ne prouvent pas à eux seuls l’opacité des entrées ni '
                        'l’équivalence pour d’autres paramètres : le contrat, les vérifications '
                        'complémentaires et le source C doivent être lus avec ces résultats.')
    lines = [f"# Comparaison des 14 cas : Go, JavaScript officiel et {reference['label']}", '',
             f"Campagne terminée : {report['completed_at_utc']} (UTC).", '',
             'Les ratios sont les temps Go divisés par la référence : au-dessus de 1×, Go est plus lent. '
             'Le tri utilise les valeurs archivées avant arrondi. Les deux classements couvrent 14/14 cas.', '',
             '## Protocole', '',
             'Trois processus indépendants par backend, ordre tournant Go/JS/C, JS/C/Go, C/Go/JS. '
             'Trois échauffements globaux et locaux, calibration vers 10 ms, dix batches ; '
             'chaque durée publiée est la médiane des trois minima par processus. '
             'Tous les résultats et tailles de batches ont passé le validateur commun. '
             'Les sources et artefacts ont été figés et vérifiés avant et après chaque exécution.', '',
             'Go : Release, GOGC=800, sans PGO, GOMAXPROCS par défaut. '
             'JavaScript : compilateur PureScript officiel, runtime Node.js. '
             f"C : référence {('fonctionnelle' if kind == 'functional' else 'impérative')} "
             'compilée séparément avant la campagne. '
             'Les versions, commandes, profils, hashes, labels/paramètres, oracles et sorties brutes '
             'sont archivés dans `results.json` et `manifest.json`. Les sources mesurées et le contrat '
             'C sont copiés sous `sources/`.', '',
             'La machine n’est pas isolée des autres activités. La charge système au début et à la fin '
             'de chaque processus est archivée. Les intervalles min–max ci-dessous sont les ratios '
             'observés dans les trois cycles tournants, pas des intervalles de confiance. '
             'Le classement porte sur les ratios des médianes, qui peuvent différer de la médiane '
             'des ratios par cycle.', '',
             '## Portée de la comparaison', '',
             scope +
             'Ces ratios comparent les programmes complets optimisés ; ils ne sont pas des coûts '
             'par opération source ni la preuve d’un nombre identique d’opérations machine.', '',
             memory +
             'Go et JavaScript utilisent leur GC. '
             'Les représentations et coûts de libération restent donc différents. '
             'Les trois harnais consomment les résultats via des accès atomiques ; les instructions '
             'natives et les adaptations FFI Go/JS restent différentes. '
             'Les adaptations d’appel et harnais sont incluses ; les très petits temps peuvent '
             'être dominés par ces coûts. ' + verification, '']
    for title, key, backend, ratio in [(f"Go / {reference['label']}", 'go_c', reference['backend'], 'go_over_c'),
                                      ('Go / JavaScript officiel', 'go_js', 'js', 'go_over_js')]:
        lines += [f'## {title} — ordre décroissant', '',
                  '| Cas | Go (µs) | Référence (µs) | Ratio | Min–max des ratios par cycle |',
                  '| --- | ---: | ---: | ---: | ---: |']
        for row in report['rankings'][key]:
            spread = row['ratio_variability'][key]
            lines.append(f"| {row['case']} | {format_number(row['times_us']['go'])} | "
                         f"{format_number(row['times_us'][backend])} | {format_number(row[ratio])}× | "
                         f"{format_number(spread['minimum'])}–{format_number(spread['maximum'])}× |")
        lines += ['']
    history = report.get('historical_comparison')
    if history:
        lines += ['## Baselines historiques — comparaison séparée', '',
                  'Le catalogue reste inchangé par cette collecte. La référence C archivée garde '
                  'son identité et sa provenance ; elle peut différer de celle de cette campagne. Les variations suivantes '
                  'confrontent les mesures Go/JS fraîches aux baselines archivées du README ; '
                  'elles ne sont pas une expérience avant/après appariée et peuvent inclure des '
                  'changements de harnais, versions ou environnement.', '',
                  '| Cas | Go frais / Go archivé | JS frais / JS archivé | C archivé (µs) |',
                  '| --- | ---: | ---: | ---: |']
        for row in history['cases']:
            values = row['program_timings']
            cells = [format_number(values[key]['fresh_over_archived']) + '×' if key in values else '—'
                     for key in ['go', 'js']]
            old_c = row.get('c_archived', row.get('imperative_c_archived'))
            cells.append(format_number(old_c['time_us']) if old_c else '—')
            lines.append('| ' + row['case'] + ' | ' + ' | '.join(cells) + ' |')
        lines += ['']
    return '\n'.join(lines)


def run_campaign(builds, output, sources, timeout, c_kind='functional'):
    backends = backends_for(c_kind)
    protocol = protocol_for(c_kind)
    if set(builds) != set(backends):
        raise ValueError('Build identities do not match the declared C reference kind')
    output = Path(output).resolve()
    output.mkdir(parents=True, exist_ok=False)
    verify_hashes(sources, ROOT)
    for name in sources:
        destination = output / 'sources' / name
        destination.parent.mkdir(parents=True, exist_ok=True)
        shutil.copyfile(ROOT / name, destination)
    verify_hashes(sources, output / 'sources')
    env = dict(os.environ, GOGC='800')
    env.pop('GOMAXPROCS', None)
    env.pop('PPROF', None)
    manifest = {'started_at_utc': timestamp(), 'host': platform.platform(),
                'machine': platform.machine(), 'protocol': protocol,
                'c_reference': {'kind': c_kind, 'backend': c_reference(c_kind)['backend'],
                                'contract': c_reference(c_kind)['contract']},
                'load_average_start': load_average(), 'host_isolated': False,
                'environment': {name: env.get(name, 'unset') for name in
                                ['GOGC', 'GOMAXPROCS', 'GOMEMLIMIT', 'GODEBUG', 'NODE_OPTIONS']},
                'sources_sha256': sources, 'builds': builds, 'runs': []}
    write_json(output / 'manifest.json', manifest)
    samples = {key: [] for key in backends}
    for cycle, order in enumerate(protocol['order'], 1):
        for key in order:
            verify_frozen(builds, sources)
            build = builds[key]
            stem = f'{cycle}-{key}'
            started = timestamp()
            load_start = load_average()
            print(f'Measuring {key}: process {cycle}/3', flush=True)
            with (output / f'{stem}.stdout').open('w') as out, (output / f'{stem}.stderr').open('w') as err:
                process = subprocess.run(build['command'], cwd=build['cwd'], env=env,
                                         stdout=out, stderr=err, timeout=timeout)
            process.check_returncode()
            load_end = load_average()
            verify_frozen(builds, sources)
            raw = (output / f'{stem}.stdout').read_text()
            checked = check_core(raw)
            write_json(output / f'{stem}.json', checked)
            samples[key].append(checked)
            manifest['runs'].append({'backend': key, 'process': cycle,
                                     'started_at_utc': started, 'completed_at_utc': timestamp(),
                                     'load_average_start': load_start, 'load_average_end': load_end,
                                     'command': build['command'], 'cwd': build['cwd'],
                                     'files_sha256': {f'{stem}.{suffix}': sha(output / f'{stem}.{suffix}')
                                                      for suffix in ['stdout', 'stderr', 'json']}})
            write_json(output / 'manifest.json', manifest)
            print(f'{stem}: 14 oracles and calibrated batches verified', flush=True)
    results = {key: aggregate(values) for key, values in samples.items()}
    report = {'started_at_utc': manifest['started_at_utc'], 'completed_at_utc': timestamp(),
              'scope': {'core_cases': 14, 'backends': list(backends)},
              'c_reference': manifest['c_reference'],
              'protocol': protocol, 'environment': manifest['environment'],
              'host': manifest['host'], 'machine': manifest['machine'], 'host_isolated': False,
              'load_average_start': manifest['load_average_start'], 'load_average_end': load_average(),
              'runs': manifest['runs'],
              'sources_sha256': sources, 'builds': builds, 'results': results,
              'rankings': rankings(results, c_kind), 'historical_comparison': historical_comparison(results)}
    verify_frozen(builds, sources)
    write_json(output / 'results.json', report)
    (output / 'report.md').write_text(render_report(report))
    manifest['completed_at_utc'] = report['completed_at_utc']
    manifest['load_average_end'] = report['load_average_end']
    manifest['results_sha256'] = sha(output / 'results.json')
    manifest['report_sha256'] = sha(output / 'report.md')
    write_json(output / 'manifest.json', manifest)
    print(f'Complete comparison: {output / "report.md"}', flush=True)
    return report


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--go-dir', type=Path, required=True)
    parser.add_argument('--js-dir', type=Path, required=True)
    parser.add_argument('--c-binary', type=Path, required=True)
    parser.add_argument('--c-manifest', type=Path)
    parser.add_argument('--c-kind', choices=tuple(C_REFERENCES), default='functional',
                        help='C implementation identity (default: functional)')
    parser.add_argument('--c-source', type=Path,
                        help='defaults to bench_functional.c or bench_native.c according to --c-kind')
    parser.add_argument('--output', type=Path, required=True)
    parser.add_argument('--timeout', type=float, default=600)
    args = parser.parse_args()
    c_source = args.c_source or ROOT / c_reference(args.c_kind)['source']
    c_backend = c_reference(args.c_kind)['backend']
    c_manifest = args.c_manifest or Path(str(args.c_binary) + '.manifest.json')
    builds = {'go': inspect_go(args.go_dir), 'js': inspect_js(args.js_dir),
              c_backend: inspect_c(args.c_binary, c_manifest, c_source)}
    for name in CORE_SOURCES:
        if builds['go']['source_sha256'][name] != builds['js']['source_sha256'][name]:
            raise ValueError(f'Go and JS do not share the same PureScript source: {name}')
    sources = source_snapshot(c_source, args.c_kind)
    verify_frozen(builds, sources)
    run_campaign(builds, args.output, sources, args.timeout, args.c_kind)


if __name__ == '__main__':
    try:
        main()
    except (ValueError, OSError, KeyError, subprocess.SubprocessError) as error:
        print(f'ERROR: {error}', file=sys.stderr)
        sys.exit(1)

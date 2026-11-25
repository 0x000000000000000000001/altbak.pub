#!/usr/bin/env python3
"""Check core comparison provenance, complete oracles and unrounded rankings."""
import copy
import json
from pathlib import Path
import runpy
import subprocess
import sys
import tempfile
import unittest
from unittest.mock import patch

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / 'bin/benchmark'))
C = runpy.run_path(str(ROOT / 'bin/benchmark/compare-core.py'))
from measure import aggregate
from validate import expected_cases


def output_fixture(times=None):
    cases = expected_cases('pure')
    times = [1.0] * 14 if times is None else times
    rows = []
    for case, elapsed in zip(cases, times):
        rows.append(f"(Test)\n{case['label']}\n\n(Output & Warm-up)\n{case['value']}\n\n"
                    f"(Execution time - best of 10)\n\n{elapsed:.9f} μs\n\nBatch iterations: 1024\n")
    return '\n'.join(rows) + f'\nTotal exec time: {sum(times) / 1000:.9f} ms\n'


def column(times=None):
    checked = C['check_core'](output_fixture(times))
    return aggregate([copy.deepcopy(checked) for _ in range(3)])


class CoreComparison(unittest.TestCase):
    def test_fourteen_oracles_and_calibrated_batches_are_required(self):
        valid = output_fixture()
        self.assertEqual(len(C['check_core'](valid)['times_us']), 14)
        invalid = [valid.replace('(Output & Warm-up)\n7\n', '(Output & Warm-up)\n8\n', 1),
                   valid.replace('Batch iterations: 1024', 'Batch iterations: 1023', 1),
                   valid.replace('1.000000000 μs', '0.000000000 μs', 1),
                   output_fixture([1.0] * 13)]
        for text in invalid:
            with self.subTest(text=text[:50]), self.assertRaises(ValueError):
                C['check_core'](text)

    def test_rankings_use_raw_ratios_before_rounding(self):
        go = [1.0] * 14
        go[0], go[1] = 1.00001, 1.00002
        result = C['rankings']({'go': column(go), 'js': column([2.0] * 14),
                                'c-functional': column()})
        for key in ['go_c', 'go_js']:
            self.assertEqual(len(result[key]), 14)
            self.assertEqual([row['case'] for row in result[key][:2]], ['Fibonacci', 'AST Evaluation'])
        self.assertEqual(result['go_c'][0]['go_over_c'], 1.00002)
        self.assertEqual(result['go_js'][0]['go_over_js'], 0.50001)
        self.assertEqual(result['go_js'][0]['expected'], '55')

    def test_ratios_use_medians_and_preserve_variability_separately(self):
        values = {}
        for key, samples in {'go': [1, 6, 10], 'js': [1, 10, 3],
                             'c-functional': [2, 4, 5]}.items():
            values[key] = aggregate([C['check_core'](output_fixture([float(t)] * 14)) for t in samples])
        row = C['rankings'](values)['go_js'][0]
        self.assertEqual(row['go_over_js'], 2.0)  # median(1,6,10) / median(1,10,3)
        spread = row['ratio_variability']['go_js']
        self.assertEqual(spread['ratios_by_cycle'], [1.0, 0.6, 10 / 3])
        self.assertEqual(spread['minimum'], 0.6)
        self.assertEqual(spread['maximum'], 10 / 3)

    def test_imperative_rankings_require_and_retain_the_declared_identity(self):
        results = {'go': column([6.0] * 14), 'js': column([3.0] * 14),
                   'c-imperative': column([2.0] * 14)}
        rows = C['rankings'](results, 'imperative')
        row = rows['go_c'][0]
        self.assertEqual(row['go_over_c'], 3)
        self.assertEqual(row['go_over_js'], 2)
        self.assertEqual(row['times_us']['c-imperative'], 2)
        self.assertNotIn('c-functional', row['times_us'])
        with self.assertRaises(ValueError):
            C['rankings'](results)
        with self.assertRaises(ValueError):
            C['rankings']({**results, 'c-functional': column()}, 'imperative')

    def test_snapshot_freezes_the_selected_contract_and_shared_opaque_boundaries(self):
        globals_ = C['source_snapshot'].__globals__
        with patch.dict(globals_, sha=lambda path: str(path)):
            for kind, source, contract in [
                    ('functional', 'tmp/bench_functional.c', 'tmp/fp_reference_contract.md'),
                    ('imperative', 'tmp/bench_native.c', 'tmp/imperative_reference_contract.md')]:
                with self.subTest(kind=kind):
                    snapshot = C['source_snapshot'](ROOT / source, kind)
                    self.assertIn(source, snapshot)
                    self.assertIn(contract, snapshot)
                    self.assertIn('src/Bench.purs', snapshot)
                    self.assertIn('src/Bench.go', snapshot)
                    self.assertIn('src/Bench.js', snapshot)
                    other = ('tmp/fp_reference_contract.md' if kind == 'imperative'
                             else 'tmp/imperative_reference_contract.md')
                    self.assertNotIn(other, snapshot)

    def test_historical_c_keeps_provenance_without_assuming_an_implementation(self):
        stored = {'time_us': 7.0, 'source': 'archived-functional.json'}
        catalogue = {'columns': {'c': {name: stored for name in C['ROWS']}}}
        globals_ = C['historical_comparison'].__globals__
        with patch.dict(globals_, json_file=lambda _: catalogue, sha=lambda _: 'digest'):
            history = C['historical_comparison']({'go': column(), 'js': column()})
        self.assertEqual(history['cases'][0]['c_archived'], stored)
        self.assertNotIn('imperative_c_archived', history['cases'][0])
        self.assertIn('recorded source identity', history['note'])

    def test_partial_columns_or_tampered_aggregate_are_rejected(self):
        results = {key: column() for key in C['BACKENDS']}
        mutations = [lambda r: r.pop('js'),
                     lambda r: r['go']['times_us'].pop(),
                     lambda r: r['go']['values'].__setitem__(0, 'wrong'),
                     lambda r: r['go']['times_us'].__setitem__(0, 9.0),
                     lambda r: r['go']['process_runs'].pop()]
        for mutation in mutations:
            changed = copy.deepcopy(results)
            mutation(changed)
            with self.assertRaises(ValueError):
                C['rankings'](changed)

    def test_hashes_reject_changed_and_missing_files(self):
        with tempfile.TemporaryDirectory() as temporary:
            path = Path(temporary) / 'artifact'
            path.write_bytes(b'frozen')
            hashes = {str(path): C['sha'](path)}
            C['verify_hashes'](hashes)
            path.write_bytes(b'changed')
            with self.assertRaisesRegex(ValueError, 'Frozen input changed'):
                C['verify_hashes'](hashes)
            path.unlink()
            with self.assertRaisesRegex(ValueError, 'missing'):
                C['verify_hashes'](hashes)

    def test_js_payload_new_file_invalidates_frozen_inventory(self):
        with tempfile.TemporaryDirectory() as temporary:
            directory = Path(temporary)
            output = directory / 'output/App'
            output.mkdir(parents=True)
            (output / 'index.js').write_text('export const main = () => {};')
            manifest = directory / 'manifest.json'
            manifest.write_text('{}')
            builds = {'js': {'manifest_path': str(manifest), 'manifest_sha256': C['sha'](manifest),
                             'cwd': str(directory), 'payload_sha256': C['js_payload'](directory)}}
            C['verify_frozen'](builds, {})
            (output / 'foreign.js').write_text('export const extra = 1;')
            with self.assertRaisesRegex(ValueError, 'inventory changed'):
                C['verify_frozen'](builds, {})

    def test_c_binary_must_match_its_build_manifest(self):
        with tempfile.TemporaryDirectory(dir=ROOT / 'tmp') as temporary:
            directory = Path(temporary)
            source, binary = directory / 'reference.c', directory / 'reference'
            source.write_text('int main(void) { return 0; }')
            binary.write_bytes(b'compiled-placeholder')
            manifest = directory / 'reference.manifest.json'
            manifest.write_text(json.dumps({'binary_sha256': C['sha'](binary),
                'source_sha256': {str(source.relative_to(ROOT)): C['sha'](source)},
                'commands': [['clang', '-O3', str(source)]],
                'toolchain': 'test compiler', 'profile': {'optimization': '-O3'}}))
            C['inspect_c'](binary, manifest, source)
            source.write_text('int main(void) { return 1; }')
            with self.assertRaisesRegex(ValueError, 'Frozen input changed'):
                C['inspect_c'](binary, manifest, source)

    def test_three_rotating_processes_archive_reports_without_publication(self):
        with tempfile.TemporaryDirectory() as temporary:
            output = Path(temporary) / 'campaign'
            builds = {key: {'command': [key], 'cwd': temporary} for key in C['BACKENDS']}
            seen = []

            def fake_run(command, **kwargs):
                seen.append(command[0])
                kwargs['stdout'].write(output_fixture())
                return subprocess.CompletedProcess(command, 0)

            # Simulate only execution; freeze checks have independent mutation tests.
            globals_ = C['run_campaign'].__globals__
            with patch.dict(globals_, verify_frozen=lambda *_: None,
                            historical_comparison=lambda _: None), \
                    patch.object(C['platform'], 'platform', return_value='test host'), \
                    patch.object(subprocess, 'run', side_effect=fake_run):
                report = C['run_campaign'](builds, output, {}, 5)
            self.assertEqual(seen, ['go', 'js', 'c-functional', 'js', 'c-functional', 'go',
                                    'c-functional', 'go', 'js'])
            self.assertEqual(len(report['runs']), 9)
            self.assertFalse(report['host_isolated'])
            self.assertTrue(all(len(result['process_runs']) == 3 for result in report['results'].values()))
            manifest = json.loads((output / 'manifest.json').read_text())
            self.assertEqual(manifest['results_sha256'], C['sha'](output / 'results.json'))
            self.assertEqual(manifest['report_sha256'], C['sha'](output / 'report.md'))
            self.assertEqual(len(list(output.glob('*.stdout'))), 9)
            self.assertEqual(len(list(output.glob('*.stderr'))), 9)
            self.assertIn('14/14', (output / 'report.md').read_text())
            with self.assertRaises(FileExistsError):
                C['run_campaign'](builds, output, {}, 5)

    def test_imperative_campaign_archives_its_own_identity_and_protocol(self):
        with tempfile.TemporaryDirectory() as temporary:
            output = Path(temporary) / 'imperative'
            builds = {key: {'command': [key], 'cwd': temporary}
                      for key in ('go', 'js', 'c-imperative')}
            seen = []

            def fake_run(command, **kwargs):
                seen.append(command[0])
                kwargs['stdout'].write(output_fixture())
                return subprocess.CompletedProcess(command, 0)

            globals_ = C['run_campaign'].__globals__
            with patch.dict(globals_, verify_frozen=lambda *_: None,
                            historical_comparison=lambda _: None), \
                    patch.object(C['platform'], 'platform', return_value='test host'), \
                    patch.object(subprocess, 'run', side_effect=fake_run):
                report = C['run_campaign'](builds, output, {}, 5, 'imperative')
            self.assertEqual(seen, ['go', 'js', 'c-imperative', 'js', 'c-imperative', 'go',
                                    'c-imperative', 'go', 'js'])
            self.assertEqual(report['c_reference'], {
                'kind': 'imperative', 'backend': 'c-imperative',
                'contract': 'tmp/imperative_reference_contract.md'})
            manifest = json.loads((output / 'manifest.json').read_text())
            self.assertEqual(manifest['c_reference'], report['c_reference'])
            self.assertEqual(report['scope']['backends'], ['go', 'js', 'c-imperative'])
            self.assertNotIn('c-functional', report['results'])
            text = (output / 'report.md').read_text()
            self.assertIn('Go / C impératif', text)
            self.assertIn('entrées restent opaques', text)
            self.assertNotIn('C fonctionnel', text)
            self.assertNotIn('malloc` avec un registre', text)
            with self.assertRaisesRegex(ValueError, 'Build identities'):
                C['run_campaign'](builds, output / 'wrong-kind', {}, 5)


if __name__ == '__main__':
    unittest.main()

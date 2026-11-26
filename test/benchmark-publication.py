#!/usr/bin/env python3
"""Keep invalid or incompatible measurements out of the published tables."""
import copy
from contextlib import redirect_stdout
import hashlib
import io
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
PUBLISH = runpy.run_path(str(ROOT / 'bin/benchmark/publish.py'))


def fixture(key):
    mode = 'x' if key.endswith('-x') else 'fficc' if key.endswith('-fficc') else 'ffi' if key.endswith('-ffi') else 'pure'
    cases = PUBLISH['expected_cases'](mode)
    times = [float(i + 1) for i in range(len(cases))]
    sample = {'mode': mode, 'values_validated': True, 'values': [c['value'] for c in cases],
              'labels': [c['label'] for c in cases], 'times_us': times, 'total_ms': sum(times) / 1000}
    return dict(sample, process_runs=[copy.deepcopy(sample) for _ in range(3)])


def status_fixture():
    sections = []
    for title in ['JavaScript', 'Go', 'C (reference)']:
        keys = PUBLISH['GROUPS'][title]
        owned = [title == 'Go' and key != 'psgo-pure' for key in keys]
        headers = ['[gopurs](compiler)' if controlled else key
                   for key, controlled in zip(keys, owned)]
        row = lambda label, values: '| ' + label + ' | ' + ' | '.join(values) + ' |'
        lines = ['#### ' + title, '', 'Benchmark | ' + ' | '.join(headers) + ' |',
                 row('---', ['---'] * len(keys))]
        if any(owned):
            lines.extend([row('', ['**Stable**<br>0.00 ms' if x else '' for x in owned]),
                          row('', ['**WIP**' if x else '' for x in owned])])
        for index, label in enumerate(['small', 'large', 'future']):
            if any(owned) and label == 'future':
                lines.extend([row('**Total WIP**', ['—' if x else '' for x in owned]),
                              row('', ['**Backlog**' if x else '' for x in owned])])
            values = ([2, 8, 16] if title == 'JavaScript' else
                      [1, None, 4] if title == 'C (reference)' else [4, 800, None])
            cells = []
            for key in keys:
                value = [12, 900, 32][index] if key == 'psgo-pure' else values[index]
                cells.append('—' if value is None else PUBLISH['format_time'](value))
            lines.append(row('[' + label + '](fixture)', cells))
            if any(value != '—' for value in cells):
                lines.extend([row('↳ /C', ['—'] * len(keys)), row('↳ /JS', ['—'] * len(keys))])
        lines.append(row('**Total**', ['' if x else '—' for x in owned]))
        sections.append('\n'.join(lines))
    return '\n\n'.join(sections) + '\n'


class PublicationTests(unittest.TestCase):
    def test_all_columns_preserve_user_headers(self):
        results = {key: fixture(key) for keys in PUBLISH['GROUPS'].values() for key in keys}
        self.assertEqual(len(results), 40)
        for key, result in results.items():
            PUBLISH['verify'](key, result)
        original = (ROOT / 'README.md').read_text()
        rendered = PUBLISH['render'](original, results)
        headers = lambda text: [line for line in text.splitlines() if 'Benchmark ' in line or line.startswith('Benchmark ')]
        self.assertEqual(headers(original), headers(rendered))
        _, extended = PUBLISH['section_lines'](rendered, 'Extended Results')
        self.assertFalse(any('(/JS' in line or '(/C' in line for line in extended))
        ratios = [line for line in extended if line.startswith('↳ /JS |')]
        self.assertEqual(ratios, ['↳ /JS | — | 1.00x | 1.00x | 1.00x |'] * 5 + [
            '↳ /JS | — | <strong>1.00x</strong> | <strong>1.00x</strong> | <strong>1.00x</strong> |'])
        self.assertEqual(extended.count('↳ /C | — | — | — | — |'), 6)

    def test_section_provenance_notes_survive_publication(self):
        original = (ROOT / 'README.md').read_text()
        note = '> Timing precision and campaign provenance are retained.\n\n'
        original = original.replace('#### Go\n\n', '#### Go\n\n' + note, 1)
        results = {key: fixture(key) for keys in PUBLISH['GROUPS'].values() for key in keys}
        rendered = PUBLISH['render'](original, results)
        self.assertEqual(rendered.count(note), 1)
        _, lines = PUBLISH['section_lines'](rendered, 'Go')
        self.assertGreater(PUBLISH['find_benchmark_row'](lines, 'AST Evaluation'), 0)

    def test_application_tables_survive_core_and_wip_publication(self):
        original = (ROOT / 'README.md').read_text()
        start = original.index('### Application benchmark results\n')
        end = original.index('### Compilation times...', start)
        application = original[start:end]
        results = {key: fixture(key) for keys in PUBLISH['GROUPS'].values() for key in keys}
        self.assertIn(application, PUBLISH['render'](original, results))

        cases = PUBLISH['expected_cases']('wip')
        library = {'labels': [case['label'] for case in cases],
                   'times_us': [float(i + 1) for i in range(len(cases))]}
        wip = {key: {'library': copy.deepcopy(library)}
               for keys in PUBLISH['WIP_GROUPS'].values() for key in keys}
        self.assertIn(application, PUBLISH['render_wip'](original, wip))

    def test_wrong_labels_outputs_and_statistics_are_rejected(self):
        original = fixture('js-pure')
        for mutation in [lambda r: r['times_us'].__setitem__(0, 9),
                         lambda r: r['process_runs'][0]['values'].__setitem__(0, 'wrong'),
                         lambda r: r['process_runs'][0]['labels'].__setitem__(0, 'another test'),
                         lambda r: r.__setitem__('total_ms', 99)]:
            result = copy.deepcopy(original)
            mutation(result)
            with self.assertRaises(ValueError):
                PUBLISH['verify']('js-pure', result)

    def test_different_source_campaigns_cannot_be_combined(self):
        with tempfile.TemporaryDirectory(prefix='altbak-publication-') as temporary:
            root = Path(temporary)
            paths = []
            for index, key in enumerate(['js-pure', 'js-ffi']):
                folder = root / str(index)
                folder.mkdir()
                data = json.dumps({key: fixture(key)})
                (folder / 'results.json').write_text(data)
                (folder / 'manifest.json').write_text(json.dumps({
                    'completed_at_utc': 'test-only', 'sources': {'source': str(index)},
                    'protocol': {'processes': 3},
                    'results_sha256': hashlib.sha256(data.encode()).hexdigest()}))
                paths += ['--campaign', str(folder)]
            readme = root / 'README.md'
            readme.write_text('must stay unchanged')
            result = subprocess.run([sys.executable, '-B', str(ROOT / 'bin/benchmark/publish.py'),
                                     *paths, '--readme', str(readme)], capture_output=True, text=True)
            self.assertNotEqual(result.returncode, 0)
            self.assertIn('different sources or timing protocols', result.stderr)
            self.assertEqual(readme.read_text(), 'must stay unchanged')

    def test_status_totals_use_matching_cases_and_keep_external_columns_unclassified(self):
        original = status_fixture()
        measurements = {'go-pure': {'small': 4, 'large': 800},
                        'js-pure': {'small': 2, 'large': 8, 'future': 16},
                        'c': {'small': 1, 'future': 4}}
        rendered = PUBLISH['refresh_backend_status_totals'](original, measurements)
        self.assertEqual(PUBLISH['refresh_backend_status_totals'](rendered, measurements), rendered)
        _, lines = PUBLISH['section_lines'](rendered, 'Go')
        total = PUBLISH['find_table_row'](lines, '**Total WIP**')
        metric = lambda name: PUBLISH['table_cells'](lines[PUBLISH['metric_row'](lines, total, name)])
        # The unreferenced 800 us case must never enter the C ratio numerator.
        for name in ['↳ /C', '↳ μ/C', '↳ Q₂/C', '↳ G/C']:
            self.assertEqual(metric(name)[1], '<strong>4.00x</strong> (n=1/2)')
            self.assertEqual(metric(name)[3], '')  # psgo is not assigned a WIP total.
        self.assertEqual(metric('↳ /JS')[1], '<strong>80.40x</strong>')
        self.assertEqual(metric('↳ μ/JS')[1], '<strong>51.00x</strong>')
        marker = next(PUBLISH['table_cells'](line) for line in lines if '**Stable**<br>' in line)
        self.assertEqual(marker[1], '**Stable**<br>0.00 ms')
        self.assertEqual(marker[3], '')
        for title in ['JavaScript', 'C (reference)']:
            _, external = PUBLISH['section_lines'](rendered, title)
            self.assertFalse(any(word in '\n'.join(external) for word in ['Stable', 'WIP', 'Backlog']))
        _, reference = PUBLISH['section_lines'](rendered, 'C (reference)')
        self.assertIn('<strong>~ 0.0050 ms</strong> (n=2/3)', '\n'.join(reference))

    def test_new_measurement_does_not_silently_publish_backlog(self):
        rendered = PUBLISH['render_backend_status'](status_fixture(), {
            'go-pure': {'future': 999}, 'js-pure': {'future': 16}})
        _, lines = PUBLISH['section_lines'](rendered, 'Go')
        index = PUBLISH['find_benchmark_row'](lines, 'future')
        cells = PUBLISH['table_cells'](lines[index])
        self.assertEqual(cells[1], '—')
        self.assertEqual(cells[3], '~ 32.00 μs')
        for metric in ['↳ /C', '↳ /JS']:
            self.assertEqual(PUBLISH['table_cells'](lines[PUBLISH['metric_row'](lines, index, metric)])[1], '—')
        self.assertNotIn('999.00', rendered)

    def test_statuses_are_rejected_on_an_external_backend(self):
        original = status_fixture().replace('|  | **WIP** | **WIP** |  |',
                                             '|  | **WIP** | **WIP** | **WIP** |')
        with self.assertRaisesRegex(ValueError, 'external backend cannot have a status'):
            PUBLISH['refresh_backend_status_totals'](original)

    def test_cli_does_not_restore_deleted_ratio_catalogue(self):
        results = {key: fixture(key) for keys in PUBLISH['GROUPS'].values() for key in keys}
        with tempfile.TemporaryDirectory(prefix='altbak-publication-catalogue-') as temporary:
            root = Path(temporary)
            readme = root / 'README.md'
            readme.write_text((ROOT / 'README.md').read_text())
            campaign = root / 'campaign'
            campaign.mkdir()
            data = json.dumps(results)
            (campaign / 'results.json').write_text(data)
            (campaign / 'manifest.json').write_text(json.dumps({
                'completed_at_utc': 'test-only', 'sources': {}, 'protocol': {},
                'results_sha256': hashlib.sha256(data.encode()).hexdigest()}))
            main = PUBLISH['main']
            with patch.dict(main.__globals__, {'ROOT': root, 'fingerprints': lambda: {}}), \
                    patch.object(sys, 'argv', ['publish.py', '--campaign', str(campaign), '--readme', str(readme)]), \
                    redirect_stdout(io.StringIO()):
                main()
            self.assertFalse((root / 'bin/benchmark/ratio-measurements.json').exists())
            self.assertIn('**Backlog**', readme.read_text())


if __name__ == '__main__':
    unittest.main()

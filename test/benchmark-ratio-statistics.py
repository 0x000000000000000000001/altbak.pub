#!/usr/bin/env python3
"""Exercise ratio arithmetic and publication using synthetic times only."""
import math
from pathlib import Path
import runpy
import sys
import unittest

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / 'bin/benchmark'))
PUBLISH = runpy.run_path(str(ROOT / 'bin/benchmark/publish.py'))


def measured(value, approximate=False):
    return ('measured', value, approximate)


def table_fixture():
    tables = []
    for title, time in [('JavaScript', 4), ('Go', 2), ('C (reference)', 1)]:
        keys = PUBLISH['GROUPS'][title]
        def row(label, value):
            return '| ' + label + ' | ' + ' | '.join([value] * len(keys)) + ' |'
        lines = ['#### ' + title, '', '| Benchmark | ' + ' | '.join(keys) + ' |',
                 row('---', '---')]
        for block, size in [('Stable', 14), ('Incl WIP', 23)]:
            lines.append(row('**' + block + '**', ''))
            for index in range(size):
                value = '~ ' + str(time) + '.00 μs'
                if title == 'C (reference)':
                    if block == 'Incl WIP':
                        value = '—'
                    elif index >= 9:
                        value = '< 0.01 μs'
                lines.append(row('[' + block + str(index) + '](fixture)', value))
                # Deliberately unrelated ratios: statistics must use the times.
                lines.extend([row('↳ /C', '999.00x'), row('↳ /JS', '< 0.01x')])
            lines.extend([row('**Total ' + block + '**', '<strong>~ 1.00 ms</strong>'),
                          row('↳ /C', '<strong>99.00x</strong>'),
                          row('↳ /JS', '<strong>99.00x</strong>')])
        lines.extend([row('**Total Stable + Incl WIP**', '<strong>~ 2.00 ms</strong>'),
                      row('↳ /C', '—'), row('↳ /JS', '<strong>99.00x</strong>'),
                      row('**Excl WIP**', '')])
        tables.append('\n'.join(lines) + '\n')
    return '\n'.join(tables)


class RatioStatistics(unittest.TestCase):
    def test_small_measurements_keep_two_significant_fractional_digits(self):
        cases = [(12.345, '12.35'), (0, '0.00'), (0.1, '0.10'),
                 (0.01234, '0.012'), (0.009492, '0.0095'),
                 (0.000045, '0.000045'), (0.0000001, '0.00000010'),
                 (0.00999, '0.010'), (0.0000999, '0.00010'),
                 (-0.000045, '-0.000045')]
        for value, expected in cases:
            with self.subTest(value=value):
                self.assertEqual(PUBLISH['format_number'](value), expected)
        self.assertEqual(PUBLISH['format_time'](0.000045), '~ 0.000045 μs')
        self.assertGreater(float(PUBLISH['format_number'](math.ulp(0.0))), 0)

    def test_mean_median_and_geometric_mean_of_ratios(self):
        result = PUBLISH['summarize_ratios'](
            [measured(value) for value in [1, 2, 8]], [measured(1)] * 3)
        self.assertAlmostEqual(result['μ'], 11 / 3)
        self.assertEqual(result['Q₂'], 2)
        self.assertAlmostEqual(result['G'], 16 ** (1 / 3))
        self.assertEqual(result['count'], 3)
        self.assertFalse(result['approximate'])

    def test_missing_and_bounded_times_have_distinct_coverage(self):
        values = [measured(2), measured(4), measured(6)]
        reference = [measured(1), ('bounded', None, False), measured(2, True)]
        result = PUBLISH['summarize_ratios'](values, reference)
        self.assertEqual((result['count'], result['population']), (2, 3))
        self.assertEqual(result['Q₂'], 2.5)
        self.assertEqual(PUBLISH['statistics_cell'](result, 'Q₂'),
                         '<strong>≈ 2.50x</strong> (n=2/3)')
        reference[1] = ('missing', None, False)
        self.assertIsNone(PUBLISH['summarize_ratios'](values, reference))

    def test_tiny_positive_ratios_do_not_use_rounded_ratio_cells(self):
        result = PUBLISH['summarize_ratios'](
            [measured(1e-8), measured(1e-6)], [measured(1), measured(1)])
        self.assertAlmostEqual(result['G'] / 1e-7, 1)
        self.assertTrue(math.isfinite(result['G']))
        self.assertEqual(PUBLISH['statistics_cell'](result, 'G'),
                         '<strong>0.00000010x</strong>')

    def test_exact_small_measurements_restore_full_coverage_without_using_bounds(self):
        # Old bounds alone stay excluded. A newly published positive measurement
        # gets an explicit numeric display and participates in all 14 ratios.
        read = PUBLISH['timing_measurement']
        values = [measured(1)] * 14
        historical = [read('~ 1.00 μs', 'c', str(i), {}, {}) for i in range(9)]
        historical += [read('< 0.01 μs', 'c', str(i), {}, {}) for i in range(9, 14)]
        self.assertEqual(PUBLISH['summarize_ratios'](values, historical)['count'], 9)
        tiny = [0.000045, 0.000031, 0.000072, 0.000081, 0.000029]
        raw = dict(zip(map(str, range(9, 14)), tiny))
        catalogue = PUBLISH['updated_ratio_catalogue'](
            {}, {'c': raw}, {'c': {'path': 'measured.json', 'sha256': 'hash'}})
        measured_refs = historical[:9] + [read(PUBLISH['format_time'](value), 'c',
                                               str(i), {}, catalogue)
                                           for i, value in zip(range(9, 14), tiny)]
        result = PUBLISH['summarize_ratios'](values, measured_refs)
        self.assertEqual((result['count'], result['population']), (14, 14))
        self.assertNotIn('(n=', PUBLISH['statistics_cell'](result, 'μ'))
        self.assertAlmostEqual(result['μ'], (9 + sum(1 / value for value in tiny)) / 14)

    def test_raw_times_require_matching_display_and_incoming_times_take_priority(self):
        catalogue = {'columns': {'go-pure': {'case': {
            'display': '~ 2.00 μs', 'time_us': 1.996}}}}
        read = lambda cell, incoming={}: PUBLISH['timing_measurement'](
            cell, 'go-pure', 'case', incoming, catalogue)
        self.assertEqual(read('~ 2.00 μs'), measured(1.996))
        self.assertEqual(read('~ 2.01 μs'), measured(2.01, True))
        self.assertEqual(read('~ 2.00 μs', {'go-pure': {'case': 2.004}}), measured(2.004))
        self.assertEqual(read('< 0.01 μs', {'go-pure': {'case': 0.001}})[0], 'bounded')
        self.assertEqual(read('—', {'go-pure': {'case': 2.004}})[0], 'missing')

    def test_statistics_only_refresh_is_idempotent_and_preserves_existing_values(self):
        original = table_fixture()
        refresh = PUBLISH['refresh_ratio_statistics']
        rendered = refresh(original, catalogue={})
        self.assertEqual(refresh(rendered, catalogue={}), rendered)
        def without_statistics(text):
            return [line for line in text.splitlines()
                    if not any(line.startswith('| ↳ ' + name + '/') for name in ['μ', 'Q₂', 'G'])]
        self.assertEqual(without_statistics(original), without_statistics(rendered))
        _, go = PUBLISH['section_lines'](rendered, 'Go')
        def value(total, statistic):
            start = PUBLISH['find_table_row'](go, total)
            index = PUBLISH['metric_row'](go, start, statistic)
            return PUBLISH['table_cells'](go[index])[1]
        self.assertEqual(value('**Total Stable**', '↳ μ/C'),
                         '<strong>≈ 2.00x</strong> (n=9/14)')
        self.assertEqual(value('**Total Stable + Incl WIP**', '↳ G/C'), '—')
        self.assertEqual(value('**Total Stable + Incl WIP**', '↳ Q₂/JS'),
                         '<strong>≈ 0.50x</strong>')
        totals = PUBLISH['refresh_new_totals'](rendered, catalogue={})
        self.assertEqual(PUBLISH['refresh_new_totals'](totals, catalogue={}), totals)
        self.assertEqual(sum(line.startswith('| ↳ Q₂/') for line in totals.splitlines()), 18)

    def test_catalogue_updates_keep_other_measurements_and_remove_resolved_gaps(self):
        catalogue = {'sources': {'old.json': 'old-hash'},
                     'columns': {'go-pure': {'untouched': {'time_us': 3}}},
                     'missing_columns': [{'key': 'go-pure', 'block': 'stable',
                                          'missing': ['fresh', 'still missing']}]}
        updated = PUBLISH['updated_ratio_catalogue'](
            catalogue, {'go-pure': {'fresh': 2.004}},
            {'go-pure': {'path': 'new.json', 'sha256': 'new-hash'}})
        self.assertEqual(updated['columns']['go-pure']['untouched'], {'time_us': 3})
        self.assertEqual(updated['columns']['go-pure']['fresh'],
                         {'time_us': 2.004, 'display': '~ 2.00 μs', 'source': 'new.json'})
        self.assertEqual(updated['missing_columns'][0]['missing'], ['still missing'])
        self.assertNotIn('fresh', catalogue['columns']['go-pure'])
        self.assertEqual(catalogue['missing_columns'][0]['missing'], ['fresh', 'still missing'])

    def test_extended_statistics_use_the_extended_js_reference_and_separate_ratio_rows(self):
        lines = ['#### Extended Results', '', 'Benchmark | JS | ES | Go | Rust |',
                 '--- | --- | --- | --- | --- |']
        for label, ratio in zip(PUBLISH['EXTENDED_ROWS'], [1, 2, 4, 8, 16]):
            lines.append('[' + label + '](fixture) | ' + ' | '.join(
                PUBLISH['format_time'](value)
                for value in [1, ratio, ratio * 2, ratio * 3]) + ' |')
            # Deliberately unrelated ratios: statistics must use the times.
            lines.extend(['↳ /C | — | — | — | — |',
                          '↳ /JS | — | 999.00x | 999.00x | 999.00x |'])
        lines.append('**Total Execution Time** | ' + ' | '.join(['~ 1.00 ms'] * 4) + ' |')
        lines.extend(['↳ /C | — | — | — | — |',
                      '↳ /JS | — | <strong>99.00x</strong> | <strong>99.00x</strong> | <strong>99.00x</strong> |'])
        original = '\n'.join(lines) + '\n'
        refresh = PUBLISH['refresh_ratio_statistics']
        rendered = refresh(original, catalogue={})
        self.assertEqual(refresh(rendered, catalogue={}), rendered)
        without_statistics = [line for line in rendered.splitlines()
                              if not any(line.startswith('↳ ' + name + '/') for name in ['μ', 'Q₂', 'G'])]
        self.assertEqual(without_statistics, original.splitlines())
        total = rendered.splitlines().index(lines[-3])
        self.assertEqual([line.split('|', 1)[0].strip() for line in rendered.splitlines()[total + 1:]],
                         ['↳ /C', '↳ μ/C', '↳ Q₂/C', '↳ G/C',
                          '↳ /JS', '↳ μ/JS', '↳ Q₂/JS', '↳ G/JS'])
        self.assertIn('↳ μ/JS | — | <strong>≈ 6.20x</strong>', rendered)
        self.assertIn('↳ Q₂/JS | — | <strong>≈ 4.00x</strong>', rendered)
        self.assertIn('↳ G/JS | — | <strong>≈ 4.00x</strong>', rendered)
        self.assertIn('↳ G/C | — | — | — | — |', rendered)
        incoming = PUBLISH['campaign_times']({'go-x': {'times_us': [1, 2, 4, 8, 16]}})
        self.assertEqual(incoming['go-x']['Parallelism'], 16)


if __name__ == '__main__':
    unittest.main()

#!/usr/bin/env python3
"""Precision migration preserves provenance and never treats a bound as exact."""
import copy
from pathlib import Path
import runpy
import sys
import unittest

sys.dont_write_bytecode = True
ROOT = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(ROOT / 'bin/benchmark'))
import precision
import publish

fixture = runpy.run_path(str(ROOT / 'test/benchmark-ratio-statistics.py'))['table_fixture']


def case_cell(text, title, label, metric=None):
    _, lines = publish.section_lines(text, title)
    index = publish.find_benchmark_row(lines, label)
    if metric:
        index = publish.metric_row(lines, index, metric)
    return publish.table_cells(lines[index])[1]


class PrecisionPublication(unittest.TestCase):
    def test_all_table_displays_normalize_without_promoting_bounds(self):
        text = ('Note: 0.08 μs\n'
                '| [link](https://example.org/0x1) | ~ 16023 ms | 0.08x | < 0.01 μs | — |\n'
                'extended | <strong>0.03x</strong> | ~ 10.43 s |\n')
        expected = ('Note: 0.08 μs\n'
                    '| [link](https://example.org/0x1) | ~ 16023.00 ms | 0.080x | < 0.01 μs | — |\n'
                    'extended | <strong>0.030x</strong> | ~ 10.43 s |\n')
        self.assertEqual(publish.format_table_numbers(text), expected)
        self.assertEqual(publish.format_table_numbers(expected), expected)

    def test_precision_switches_at_one(self):
        cases = [(0.000045, '0.000045'), (0.009492, '0.0095'),
                 (0.1, '0.10'), (0.9999, '1.00'), (1, '1.00'),
                 (1.0045, '1.00'), (12.000045, '12.00'),
                 (1234.005678, '1234.01'), (12.00999, '12.01'),
                 (12.9999, '13.00'), (-12.000045, '-12.00'),
                 (-0.000045, '-0.000045'), (0, '0.00')]
        for value, expected in cases:
            with self.subTest(value=value):
                self.assertEqual(publish.format_number(value), expected)

    def test_fresh_small_c_measurements_restore_fourteen_pairs_and_totals(self):
        values = {f'Stable{i}': 1 if i < 9 else 0.000045 for i in range(14)}
        incoming = {'c': values}
        origins = {'c': {'path': 'validated-c.json', 'sha256': 'verified-hash'}}
        rendered, catalogue = precision.refresh(fixture(), {}, incoming, origins)
        self.assertEqual(case_cell(rendered, 'C (reference)', 'Stable9'), '~ 0.000045 μs')
        self.assertEqual(case_cell(rendered, 'Go', 'Stable9', '↳ /C'), '44444.44x')
        self.assertNotIn('(n=9/14)', rendered)
        _, go = publish.section_lines(rendered, 'Go')
        total = publish.find_table_row(go, '**Total Stable**')
        summary = publish.table_cells(go[publish.metric_row(go, total, '↳ μ/C')])[1]
        self.assertNotIn('(n=', summary)
        self.assertEqual(catalogue['sources']['validated-c.json'], 'verified-hash')
        self.assertEqual(catalogue['columns']['c']['Stable9']['time_us'], 0.000045)
        self.assertEqual(precision.refresh(rendered, catalogue), (rendered, catalogue))

    def test_matching_archive_restores_a_formatting_bound_without_changing_time(self):
        catalogue = {'sources': {'old.json': 'hash'}, 'columns': {'c': {'Stable9': {
            'time_us': 0.000045123, 'display': '< 0.01 μs', 'source': 'old.json'}}}}
        original = copy.deepcopy(catalogue)
        rendered, updated = precision.refresh(fixture(), catalogue)
        self.assertEqual(case_cell(rendered, 'C (reference)', 'Stable9'), '~ 0.000045 μs')
        self.assertEqual(case_cell(rendered, 'C (reference)', 'Stable10'), '< 0.01 μs')
        self.assertIn('(n=10/14)', rendered)
        self.assertEqual(updated['columns']['c']['Stable9']['time_us'], 0.000045123)
        self.assertEqual(updated['columns']['c']['Stable9']['source'], 'old.json')
        self.assertEqual(catalogue, original)

    def test_unmatched_or_missing_raw_values_do_not_promote_bounds(self):
        catalogue = {'columns': {'c': {'Stable9': {
            'time_us': 0.000045, 'display': '~ 0.001 μs', 'source': 'old.json'}}}}
        rendered, _ = precision.refresh(fixture(), catalogue)
        self.assertEqual(case_cell(rendered, 'C (reference)', 'Stable9'), '< 0.01 μs')
        self.assertEqual(case_cell(rendered, 'Go', 'Stable9', '↳ /C'), '—')
        self.assertIn('(n=9/14)', rendered)

    def test_tiny_ratio_uses_raw_times_and_explicit_digits(self):
        raw = 0.000045123
        catalogue = {'columns': {'go-pure': {'Stable0': {
            'time_us': raw, 'display': '~ 2.00 μs', 'source': 'old.json'}}}}
        rendered, _ = precision.refresh(fixture(), catalogue)
        self.assertEqual(case_cell(rendered, 'Go', 'Stable0'), '~ 0.000045 μs')
        self.assertEqual(case_cell(rendered, 'Go', 'Stable0', '↳ /JS'), '0.000011x')

    def test_invalid_new_times_are_rejected(self):
        for value in [0, -1, float('nan'), float('inf')]:
            with self.subTest(value=value), self.assertRaises(ValueError):
                precision.refresh(fixture(), {}, {'c': {'Stable9': value}},
                                  {'c': {'path': 'invalid.json', 'sha256': 'hash'}})


if __name__ == '__main__':
    unittest.main()

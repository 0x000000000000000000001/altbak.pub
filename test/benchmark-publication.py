#!/usr/bin/env python3
"""Keep invalid or incompatible measurements out of the published tables."""
import copy
import hashlib
import json
from pathlib import Path
import runpy
import subprocess
import sys
import tempfile
import unittest

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


class PublicationTests(unittest.TestCase):
    def test_all_columns_preserve_user_headers(self):
        results = {key: fixture(key) for keys in PUBLISH['GROUPS'].values() for key in keys}
        self.assertEqual(len(results), 39)
        for key, result in results.items():
            PUBLISH['verify'](key, result)
        original = (ROOT / 'README.md').read_text()
        rendered = PUBLISH['render'](original, results)
        headers = lambda text: [line for line in text.splitlines() if 'Benchmark ' in line or line.startswith('Benchmark ')]
        self.assertEqual(headers(original), headers(rendered))

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


if __name__ == '__main__':
    unittest.main()

#!/usr/bin/env python3
"""Regression checks for failures hidden by the former multi-backend runner."""
import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
BACKENDS = ['js', 'es', 'php', 'go', 'rust', 'sharp', 'java', 'wasm', 'scm', 'erl', 'cpp', 'psgo']


class Orchestrator(unittest.TestCase):
    def setUp(self):
        self.scratch = tempfile.TemporaryDirectory(prefix='altbak-orchestrator-')
        self.root = Path(self.scratch.name)
        (self.root / 'bin').mkdir()
        shutil.copyfile(ROOT / 'bin/run', self.root / 'bin/run')
        for name in BACKENDS:
            target = self.root / 'bin' / name / 'run'
            target.parent.mkdir()
            target.write_text('#!/usr/bin/env python3\nimport json, os, pathlib, sys\n'
                'with pathlib.Path("calls.jsonl").open("a") as log:\n'
                f' log.write(json.dumps([{name!r}, sys.argv[1:]]) + "\\n")\n'
                f'sys.exit(1 if os.environ.get("FAIL_BACKEND") == {name!r} else 0)\n')
            target.chmod(0o755)

    def tearDown(self):
        self.scratch.cleanup()

    def invoke(self, args, failure=''):
        env = dict(os.environ, FAIL_BACKEND=failure)
        result = subprocess.run(['bash', 'bin/run', *args], cwd=self.root, env=env,
                                capture_output=True, text=True)
        path = self.root / 'calls.jsonl'
        calls = [json.loads(line) for line in path.read_text().splitlines()] if path.exists() else []
        return result, calls

    def test_failed_backend_remains_failure_after_later_success(self):
        result, calls = self.invoke(['--ffi'], 'php')
        self.assertEqual(result.returncode, 1)
        self.assertEqual([name for name, _ in calls], [x for x in BACKENDS if x not in ['wasm', 'psgo']])
        self.assertTrue(all(args == ['--ffi'] for _, args in calls))
        self.assertIn('Benchmark failures: php', result.stderr)

    def test_all_core_backends_are_included(self):
        result, calls = self.invoke([])
        self.assertEqual(result.returncode, 0)
        self.assertEqual([name for name, _ in calls], BACKENDS)

    def test_invalid_options_do_not_launch_a_backend(self):
        for args in [['--typo'], ['--test'], ['--ffi', '--fficc']]:
            with self.subTest(args=args):
                result, calls = self.invoke(args)
                self.assertNotEqual(result.returncode, 0)
                self.assertEqual(calls, [])


if __name__ == '__main__':
    unittest.main()

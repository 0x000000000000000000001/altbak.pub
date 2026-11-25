#!/usr/bin/env python3
"""Compile the imperative reference and vary inputs after compilation."""
import os
from pathlib import Path
import random
import shutil
import subprocess
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[1]
SOURCE = ROOT / 'tmp/bench_native.c'


class ImperativeReference(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.compiler = shutil.which('clang')
        if not cls.compiler:
            raise unittest.SkipTest('clang is required for the native reference')
        cls.workspace = tempfile.TemporaryDirectory(prefix='altbak-imperative-reference-')
        cls.binary = Path(cls.workspace.name) / 'reference'
        subprocess.run([cls.compiler, '-std=c11', '-O3', '-Wall', '-Wextra', '-Werror',
                        str(SOURCE), '-o', str(cls.binary)], check=True, capture_output=True, text=True)

    @classmethod
    def tearDownClass(cls):
        cls.workspace.cleanup()

    def run_case(self, name, value):
        return int(subprocess.check_output([str(self.binary), '--case', name, str(value)], text=True))

    def test_all_benchmark_oracles_variants_and_tree_invariants(self):
        result = subprocess.check_output([str(self.binary), '--verify'], text=True)
        self.assertIn('14 core oracles, 56 input variants, 257 tree sizes', result)
        self.assertIn('rotations, duplicates and Int32 arithmetic', result)

    def test_inputs_are_selected_after_compilation(self):
        # These values enter through argv, after the optimized executable exists.
        # Arithmetic simplification is valid; substituting the benchmark's fixed
        # expected result is not.
        values = [0, 1, 2, 17] + random.Random(8217).sample(range(3, 300), 5)
        for value in values:
            for case, expected in [
                ('Polymorphism', value), ('StateMonad', value * 60),
                ('LazyEvaluation', value * 1000),
                ('ListOps', sum(i for i in range(1, value + 1) if i % 2 == 0)),
                ('TCO', sum(i % 3 for i in range(1, value + 1))),
                ('Records', sum(i % 5 for i in range(1, value + 1))),
            ]:
                with self.subTest(case=case, value=value):
                    self.assertEqual(self.run_case(case, value), expected)
        for value in range(-11, 12):
            step = 1 if value >= 1 else -1
            expected = sum(i for i in range(1, value + step, step) if i % 2 == 0)
            self.assertEqual(self.run_case('ArrayOps', value), expected)
        for value in range(7):
            self.assertEqual(self.run_case('Church', value), value ** 5)
        # The source deliberately ignores its opaque number and counts a fixed
        # record's fields. That input-independent result is part of its contract.
        for value in [-2147483648, -1, 0, 17, 10000, 2147483647]:
            self.assertEqual(self.run_case('RowToList', value), 5)

    def test_invalid_inputs_are_rejected_instead_of_returning_a_guess(self):
        for case, value in [('AstTree', -1), ('AstTree', 13), ('Ackermann', 4),
                            ('Fib', 'abc'), ('RowToList', 2147483648), ('unknown', 1)]:
            result = subprocess.run([str(self.binary), '--case', case, str(value)],
                                    text=True, capture_output=True)
            self.assertEqual(result.returncode, 2, (case, value, result))

    def test_sanitized_oracles_and_mutable_tree(self):
        binary = Path(self.workspace.name) / 'reference-sanitized'
        subprocess.run([self.compiler, '-std=c11', '-O1', '-g', '-Wall', '-Wextra', '-Werror',
                        '-fsanitize=address,undefined', '-fno-omit-frame-pointer',
                        str(SOURCE), '-o', str(binary)], check=True, capture_output=True, text=True)
        result = subprocess.run([str(binary), '--verify'], text=True, capture_output=True,
                                env={**os.environ, 'ASAN_OPTIONS': 'halt_on_error=1',
                                     'UBSAN_OPTIONS': 'halt_on_error=1'})
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        self.assertEqual(result.stderr, '')


if __name__ == '__main__':
    unittest.main()

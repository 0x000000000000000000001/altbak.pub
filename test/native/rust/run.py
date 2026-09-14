#!/usr/bin/env python3
"""Compile exact native Rust sources in a temporary directory, without timing."""
import os
from pathlib import Path
import subprocess
import sys
import tempfile

sys.dont_write_bytecode = True
sys.path.insert(0, str(Path(__file__).resolve().parents[1]))
from oracle import ROOT, NOMINAL, cases

with tempfile.TemporaryDirectory(prefix='altbak-native-rust-') as temporary:
    folder = Path(temporary)
    lines = ['#![allow(dead_code, unused_mut, unused_variables, non_snake_case)]']
    for name in NOMINAL:
        for suffix in ['FFI', 'FFICheatcode']:
            module = name + suffix
            lines.append(f'#[path = "{ROOT}/src/Test/{module}.rs"] mod {module};')
    lines += ['fn main() {', 'let mut failures = 0;']
    count = 0
    for name, argument, expected in cases():
        for suffix in ['FFI', 'FFICheatcode']:
            module = name + suffix
            lines += [f'let got = {module}::Test_{module}_run{module}(std::hint::black_box({argument}));',
                      f'if got != {expected} {{ failures += 1; eprintln!("FAIL {module}({argument}): expected {expected}, got {{}}", got); }}']
            count += 1
    lines += [f'assert_eq!(failures, 0, "native Rust contracts ({count} cases)");',
              f'println!("PASS native Rust: {count} values; no timing");', '}']
    source = folder / 'contracts.rs'
    source.write_text('\n'.join(lines) + '\n')
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', '-C', 'overflow-checks=yes',
                    str(source), '-o', str(folder / 'contracts')], check=True)
    subprocess.run([str(folder / 'contracts')], check=True)

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
    structural_modules = []
    for name in NOMINAL:
        for suffix in ['FFI', 'FFICheatcode']:
            module = name + suffix
            original = (ROOT / 'src' / 'Test' / (module + '.rs')).read_text()
            extra = Path(__file__).with_name(module + '.checks.rs')
            if extra.exists():
                original += '\n' + extra.read_text()
                structural_modules.append(module)
            copied = folder / (module + '.rs')
            copied.write_text(original)
            lines.append(f'#[path = "{copied}"] mod {module};')
    lines += ['fn main() {', 'let mut failures = 0;']
    count = 0
    for name, argument, expected in cases():
        for suffix in ['FFI', 'FFICheatcode']:
            module = name + suffix
            lines += [f'let got = {module}::Test_{module}_run{module}(std::hint::black_box({argument}));',
                      f'if got != {expected} {{ failures += 1; eprintln!("FAIL {module}({argument}): expected {expected}, got {{}}", got); }}']
            count += 1
    for module in structural_modules:
        lines.append(f'{module}::check_structure();')
    lines += [f'assert_eq!(failures, 0, "native Rust contracts ({count} cases)");',
              f'println!("PASS native Rust: {count} values; {len(structural_modules)} generic/structural modules; no benchmark timing");', '}']
    source = folder / 'contracts.rs'
    source.write_text('\n'.join(lines) + '\n')
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', '-C', 'overflow-checks=yes',
                    str(source), '-o', str(folder / 'contracts')], check=True)
    subprocess.run([str(folder / 'contracts')], check=True)

    clock = folder / 'clock.rs'
    clock.write_text(Path(__file__).with_name('clock.rs').read_text().replace('BENCH_PATH', str(ROOT / 'src/Bench.rs')))
    subprocess.run(['rustc', '--edition=2021', str(clock), '-o', str(folder / 'clock')], check=True)
    subprocess.run([str(folder / 'clock')], check=True)

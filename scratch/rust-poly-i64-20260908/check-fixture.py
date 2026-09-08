"""Execute both instantiations and the generic fallback emitted by trace.mjs."""
from pathlib import Path
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
GENERATED = HERE.parents[1] / 'run/bak/rust/output/purust_output'
DEPS = GENERATED / 'target/release/deps'
BUILD = HERE / 'build'
BUILD.mkdir(exist_ok=True)

def dependency(name):
    paths = list(DEPS.glob(f'lib{name}-*.rlib'))
    assert len(paths) == 1, (name, paths)
    return paths[0]

def run(args):
    result = subprocess.run(list(map(str, args)), capture_output=True, text=True)
    assert result.returncode == 0, result.stdout + result.stderr
    return result.stdout

manifest = (HERE / 'rust/Purs_PolyI64/Cargo.toml').read_text().split('[dependencies]\n')[1]
names = [n.replace('-', '_') for n in re.findall(r'^([\w-]+) =', manifest, re.M)]
externs = [arg for name in names for arg in ['--extern', f'{name}={dependency(name)}']]
common = ['rustc', '--edition=2021', '-C', 'opt-level=1', '-L', f'dependency={DEPS}']
run(common + ['--crate-type=rlib', '--crate-name', 'fixture', HERE / 'PolyI64-generated.rs',
              '-o', BUILD / 'libfixture.rlib'] + externs)
source = '''#![allow(warnings)]
use purust_core::*;
use fixture::*;
fn main() {
    for n in [0_i64, 1, 2, 7, 1000] {
        for initial in [-42_i64, 0, 17] {
            assert_eq!(PolyI64_intLoop(n, initial), initial + n);
            assert_eq!(PolyI64_polyLoop(PolyI64_intMonoidish(), n, mk_int(initial)).unwrap_int(), initial + n);
            let number = initial as f64 + 0.25;
            let expected = number + 0.5 * n as f64;
            assert_eq!(PolyI64_numberLoop(n, number), expected);
            assert_eq!(PolyI64_polyLoop(PolyI64_numberMonoidish(), n, mk_number(number)).unwrap_number(), expected);
        }
    }
    println!("60 assertions: Int, Number, both generic dictionary calls, zero iterations and varied initial values.");
}
'''
(BUILD / 'fixture-checks.rs').write_text(source)
run(common + [BUILD / 'fixture-checks.rs', '-o', BUILD / 'fixture-checks',
    '--extern', f'fixture={BUILD}/libfixture.rlib', '--extern', f'purust_core={dependency("purust_core")}'])
print(run([BUILD / 'fixture-checks']).strip())

# Keep only the accumulator implementation's machine code, excluding debug data.
for side in ['before', 'native']:
    assembly = (BUILD / f'poly_{side}.s').read_text()
    start = re.search(r'^__ZN[^\n]*Test_Polymorphism_act[^\n]*purs_local_2_rec_0_impl[^\n]*:$', assembly, re.M)
    assert start, side
    end = assembly.index('\t.cfi_endproc', start.end())
    excerpt = assembly[start.start():end] + '\t.cfi_endproc\n'
    (HERE / f'kernel-{side}.s').write_text(excerpt)
    if side == 'native':
        assert '\tadd\tx0, x19, x20' in excerpt
        assert not re.search(r'^\s*(b|b\.\w+|cbnz|cbz|tbnz|tbz)\s', excerpt, re.M)
print('Native accumulator assembly: one add, no loop branch (arm64).')

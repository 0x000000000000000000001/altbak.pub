"""Inspect the actual emitted kernel and archive its assembly, without edits."""
from pathlib import Path
import difflib
import hashlib
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
GENERATED = ROOT / 'run/bak/rust/output/purust_output'
DEPS = GENERATED / 'target/release/deps'
BUILD = HERE / 'build'
BUILD.mkdir(exist_ok=True)
source = GENERATED / 'Purs_Test_Polymorphism/src/lib.rs'
before = (HERE / 'Polymorphism-before.rs').read_text()
after = source.read_text()
(HERE / 'Polymorphism-after.rs').write_text(after)
(HERE / 'Polymorphism.diff').write_text(''.join(difflib.unified_diff(
    before.splitlines(True), after.splitlines(True), fromfile='before.rs', tofile='after.rs')))
act = after[after.index('pub fn Test_Polymorphism_act('):]
assert re.search(r'fn \w+_impl\([^\n]*mut \w+: i64\) -> i64', act)
assert not re.search(r'let _tco_temp_\d+ =[^\n]*(?:mk_int|unwrap_int)', act)
generic = after[after.index('pub fn Test_Polymorphism_polyLoop('):after.index('pub fn Test_Polymorphism_intMonoidish(')]
assert re.search(r'fn \w+_impl\([^\n]*mut \w+: crate::UnknownType\) -> crate::UnknownType', generic)
manifest = (GENERATED / 'Purs_Test_Polymorphism/Cargo.toml').read_text().split('[dependencies]\n')[1]
names = [name.replace('-', '_') for name in re.findall(r'^([\w-]+) =', manifest, re.M)]
externs = []
for name in names:
    paths = list(DEPS.glob(f'lib{name}-*.rlib'))
    assert len(paths) == 1, (name, paths)
    externs += ['--extern', f'{name}={paths[0]}']
assembly = BUILD / 'Polymorphism.s'
subprocess.run(['rustc', '--edition=2021', '--crate-type=rlib', '--crate-name', 'poly_verified',
    '-C', 'opt-level=1', str(source), '--emit', f'asm={assembly}',
    '-L', f'dependency={DEPS}'] + externs, check=True)
text = assembly.read_text()
start = re.search(r'^__ZN[^\n]*Test_Polymorphism_act[^\n]*purs_local_2_rec_0_impl[^\n]*:$', text, re.M)
assert start
end = text.index('\t.cfi_endproc', start.end())
kernel = text[start.start():end] + '\t.cfi_endproc\n'
assert re.search(r'^\s*add\s+x0,', kernel, re.M)
assert not re.search(r'^\s*(b|b\.\w+|cbnz|cbz|tbnz|tbz)\s', kernel, re.M)
(HERE / 'kernel-after.s').write_text(kernel)
roots = {'pbo_worktree': ROOT / '../purescript-backend-optimizer-purust',
         'purust': ROOT / '../purust/purust', 'altbak_worktree': ROOT}
tracked = [source, roots['purust'] / 'bin/purust.js',
    roots['pbo_worktree'] / 'src/PureScript/Backend/Optimizer/Convert.purs',
    roots['pbo_worktree'] / 'src/PureScript/Backend/Optimizer/Semantics.purs',
    roots['pbo_worktree'] / 'src/PureScript/Backend/Optimizer/CoreFn/TypeInstantiation.purs']
test_counts = {}
for name, expected in [('codegen', 14), ('tast', 3)]:
    log = (HERE / f'{name}-tests.log').read_text()
    passed = re.search(r'ℹ pass (\d+)', log)
    assert passed and int(passed[1]) == expected
    assert re.search(r'ℹ fail 0\b', log)
    test_counts[name] = int(passed[1])
metadata = {
    'base_heads': {name: subprocess.check_output(['git', 'rev-parse', 'HEAD'], cwd=path, text=True).strip()
                   for name, path in roots.items()},
    'sha256': {str(path.resolve()): hashlib.sha256(path.read_bytes()).hexdigest() for path in tracked},
    'native_accumulator': True, 'generic_accumulator_preserved': True,
    'arm64_kernel_add_without_loop': True, 'codegen_tests_passed': test_counts['codegen'], 'tast_tests_passed': test_counts['tast'],
    'runner_outputs_checked_per_process': 14,
}
(HERE / 'validation.json').write_text(json.dumps(metadata, indent=2) + '\n')
print('Generated i64 accumulator, generic fallback, no loop conversions; actual arm64 kernel reduces to addition.')

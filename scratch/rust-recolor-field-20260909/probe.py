"""Measure changing only the color of an already-unique RBTree node."""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import re
import runpy
import subprocess
import sys

HERE = Path(__file__).resolve().parent
BUILD = HERE/'build'
UNIQUE = HERE.parent/'rust-uniqueness-proof-20260909'
spec = importlib.util.spec_from_file_location('uniqueness_probe', UNIQUE/'probe.py')
common = importlib.util.module_from_spec(spec)
spec.loader.exec_module(common)
# Reuse the existing measurement/counting harness without overwriting its evidence.
common.HERE, common.BUILD = HERE, BUILD
PREVIOUS, GENERATED = common.PREVIOUS, common.SOURCE
# Integration regenerates the live output. Keep using the saved baseline for
# the independent prototype; reconstruct it from the recorded revision if needed.
SOURCE = BUILD/'RBTree-before.rs' if (BUILD/'RBTree-before.rs').exists() else GENERATED
common.SOURCE = SOURCE
INTEGRATED = False

def recolor(source):
    pattern = re.compile(
        r'\{ let mut (?P<cell>\w+) = (?P=cell); let _taken = .*?; match _taken \{ '
        r'.*? => \{ let _rebuilt = crate::Tree::T\(crate::Color::B,.*?; '
        r'\*std::rc::Rc::get_mut\(&mut (?P=cell)\)\.unwrap\(\) = _rebuilt; (?P=cell) \}, '
        r'std::option::Option::None => \{(?P<fallback>.*?)\}, _ => unreachable!\(\) \} \}')
    def replace(match):
        cell = match['cell']
        return (f'{{ let mut {cell} = {cell}; '
            f'if let Some(crate::Tree::T(_color, ..)) = std::rc::Rc::get_mut(&mut {cell}) '
            f'{{ *_color = crate::Color::B; {cell} }} else {{'+match['fallback']+'} }')
    result, n = pattern.subn(replace, source)
    assert n == 2, n
    return result, n

def variants():
    original = common.previous.kernel(SOURCE.read_text())
    if INTEGRATED:
        return {'before': original, 'integrated': common.previous.kernel(GENERATED.read_text())}, {'integrated': True}
    prototype, n = recolor(original)
    return {'before': original, 'recolor': prototype}, {'recolor_sites': n}

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['time', 'count', 'validate'])
    parser.add_argument('--assembly', action='store_true')
    parser.add_argument('--pairs', type=int, default=5)
    parser.add_argument('--output', default='timings.json')
    parser.add_argument('--reuse-binaries', action='store_true')
    parser.add_argument('--integrated', action='store_true')
    args = parser.parse_args()
    INTEGRATED = args.integrated
    BUILD.mkdir(exist_ok=True)
    versions, sites = variants()
    print(sites, flush=True)
    if args.mode == 'count':
        if INTEGRATED: common.HERE = BUILD
        common.count(versions)
        if INTEGRATED: (BUILD/'counts.json').replace(HERE/'counts-integrated.json')
    elif args.mode == 'validate':
        # validate.py is a harness parameterized by its imported `probe` module.
        # Route it to this experiment's variants and output directory.
        destination = HERE
        if INTEGRATED: HERE = BUILD
        sys.modules['probe'] = sys.modules[__name__]
        runpy.run_path(str(UNIQUE/'validate.py'), run_name='__main__')
        if INTEGRATED: (BUILD/'validation.json').replace(destination/'validation-integrated.json')
    else:
        if not args.reuse_binaries: common.compile_native(versions, args.assembly)
        else:
            for name, code in versions.items():
                assert (BUILD/f'time-{name}.rs').read_text() == '#![allow(warnings)]\n'+code+common.HARNESS
        common.measure(versions, args.output, args.pairs)
        result = json.loads((HERE/args.output).read_text())
        result['rustc'] = subprocess.check_output(['rustc', '--version'], text=True).strip()
        result['binary_sha256'] = {name: hashlib.sha256((BUILD/f'time-{name}').read_bytes()).hexdigest() for name in versions}
        (HERE/args.output).write_text(json.dumps(result, indent=2)+'\n')

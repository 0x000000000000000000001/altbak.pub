#!/usr/bin/env python3
"""Cumulative, scratch-only changes to the two real generated ListOps filters.

No timing command is provided. Full modules are inputs to the coordinator's
separate complete runner. All dependencies come from the prior frozen tree.
"""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
PREVIOUS = HERE.parents[1] / 'rust-perceus-evaluation-20260914'
FROZEN = PREVIOUS / 'build/frozen'
SOURCE = FROZEN / 'Purs_Test_ListOps/src/lib.rs'
DEPS = PREVIOUS / 'build/list-validation/target/release/deps'
BUILD = HERE / 'build'
VARIANTS = ['baseline', 'lifetimes', 'consuming', 'reuse']
spec = importlib.util.spec_from_file_location('previous_list_probe', PREVIOUS / 'list_probe.py')
previous = importlib.util.module_from_spec(spec)
spec.loader.exec_module(previous)

HEADER = re.compile(r'    } else if /\* OpIsTag Debug: Test_ListOps_Cons -> Func \*/ matches!\(\(/\*[^\n]*?\*/(purs_local_[23])\)\.as_ref\(\), crate::List::Cons\(\.\.\)\) \{\n')

def filter_blocks(source):
    result = []
    for match in HEADER.finditer(source):
        following = source[match.end():]
        if following.startswith(('        if /* Typed bool', '        let crate::List::Cons(probe_head', '        let probe_keep')):
            end = source.index('    };\n        }\n    }', match.end())
            result.append((match, end))
    assert len(result) == 2, len(result)
    return result

def precise_lifetimes(source):
    before = 'purust_core::Value::Class(std::rc::Rc::new(_f((_a0).unwrap_class::<std::rc::Rc<crate::List>>().clone(), (_a1).unwrap_class::<std::rc::Rc<crate::List>>().clone())))'
    after = '''{
        // Scratch: closed List Int call; acquire both owned arguments first.
        let probe_arg0 = (_a0).unwrap_class::<std::rc::Rc<crate::List>>().clone();
        let probe_arg1 = (_a1).unwrap_class::<std::rc::Rc<crate::List>>().clone();
        drop(_a1); drop(_a0);
        purust_core::Value::Class(std::rc::Rc::new(_f(probe_arg0, probe_arg1)))
    }'''
    assert source.count(before) == 2
    return source.replace(before, after)

def reuse_cells(source):
    result = source
    for match, end in reversed(filter_blocks(source)):
        owner = match[1]
        accumulator = 'purs_local_' + str(int(owner.rsplit('_', 1)[1]) + 1)
        block = source[match.start():end]
        start = block.index('        if /* Typed bool')
        finish = block.index(') {\n        {', start) + 1
        condition = block[start + len('        if '):finish]
        assert condition.count('probe_head.clone()') == 1
        borrowed_head = '{ match (' + owner + ').as_ref() { crate::List::Cons(ref f, ..) => f.clone(), _ => unreachable!() } }'
        condition = condition.replace('probe_head.clone()', borrowed_head)
        prefix = f'''        let probe_keep = {condition};
        if probe_keep {{
            if let std::option::Option::Some(probe_slot) = std::rc::Rc::get_mut(&mut {owner}) {{
                let crate::List::Cons(_, probe_tail_slot) = probe_slot else {{ unreachable!() }};
                let probe_next = std::mem::replace(probe_tail_slot, {accumulator}.clone());
                let probe_reused_cell = {owner};
                {owner} = probe_next;
                {accumulator} = probe_reused_cell;
                // PROBE_REUSED_CELL
                continue;
            }}
        }}
'''
        block = block[:start] + '        if probe_keep' + block[finish:]
        block = block[:match.end() - match.start()] + prefix + block[match.end() - match.start():]
        result = result[:match.start()] + block + result[end:]
    return result

COUNTERS = r'''
std::thread_local! { static PROBE_FILTER: std::cell::Cell<[u64;7]> = const { std::cell::Cell::new([0;7]) }; }
fn probe_hit(index: usize) { PROBE_FILTER.with(|v| { let mut a=v.get(); a[index]+=1; v.set(a); }); }
fn probe_filter_visit(node: &std::rc::Rc<List>) {
    probe_hit(if std::rc::Rc::strong_count(node) == 1 {0} else {1});
    if std::rc::Rc::weak_count(node) > 0 { probe_hit(2); }
    if let List::Cons(head, _) = node.as_ref() { if head.unwrap_int().rem_euclid(2) == 0 { probe_hit(3); } }
}
fn probe_list_consume(node: std::rc::Rc<List>) -> List {
    probe_hit(if std::rc::Rc::strong_count(&node) == 1 {5} else {6});
    std::rc::Rc::unwrap_or_clone(node)
}
pub fn probe_filter_counts() -> [u64;7] { PROBE_FILTER.with(|v| v.replace([0;7])) }
extern "C" { fn probe_mark_list_allocation(); }
pub fn probe_list_new(payload: List) -> std::rc::Rc<List> {
    unsafe { probe_mark_list_allocation(); }
    std::rc::Rc::new(payload)
}
'''

def instrument(source):
    result = source
    for match, end in reversed(filter_blocks(source)):
        owner = match[1]
        block = source[match.start():end]
        prefix = f'        probe_filter_visit(&{owner});\n'
        block = block[:match.end()-match.start()] + prefix + block[match.end()-match.start():]
        block = block.replace(f'std::rc::Rc::unwrap_or_clone({owner})', f'probe_list_consume({owner})')
        block = block.replace('// PROBE_REUSED_CELL', 'probe_hit(4); // PROBE_REUSED_CELL')
        result = result[:match.start()] + block + result[end:]
    result = result.replace('std::rc::Rc::new(crate::List::', 'probe_list_new(crate::List::')
    return result + COUNTERS

def digest(path):
    return hashlib.sha256(path.read_bytes()).hexdigest()

def prepare():
    HERE.mkdir(parents=True, exist_ok=True)
    original = SOURCE.read_text()
    lifetime = precise_lifetimes(original)
    consuming = previous.transform(lifetime, True)
    variants = {'baseline': original, 'lifetimes': lifetime,
                'consuming': consuming, 'reuse': reuse_cells(consuming)}
    for name, source in variants.items():
        (HERE / f'ListOps-{name}.rs').write_text(source)
        (HERE / f'ListOps-{name}-counted.rs').write_text(instrument(source))
    # Stage inversion establishes that no extra edit entered the lifetime step.
    assert (HERE / 'ListOps-baseline.rs').read_bytes() == SOURCE.read_bytes()
    metadata = {
        'source': str(SOURCE), 'source_sha256': digest(SOURCE),
        'variants': {name: {'plain': digest(HERE/f'ListOps-{name}.rs'),
            'counted': digest(HERE/f'ListOps-{name}-counted.rs')} for name in variants},
        'steps': ['Byte-identical baseline', 'Release two Value::Class argument wrappers after acquiring both Rc<List> arguments, before the closed filter call',
            'Add the prior consumed-pattern substitution in both filter loops',
            'Reuse each retained Cons cell only when Rc::get_mut succeeds; preserve its head and replace only its tail'],
        'unchanged': 'Range, foldl, thunk/closure generation, boxing representation, accumulator clones and every other module remain unchanged in all plain variants.',
        'scope': 'Manual scratch prototype for the actual closed List Int filter; not a compiler rule for arbitrary Value/Class payloads or callbacks.',
        'profile': 'rustc O1 debug=yes; System allocator for checks; full runner uses its own mimalloc configuration',
        'deps': str(DEPS), 'no_timings': True,
    }
    (HERE / 'metadata.json').write_text(json.dumps(metadata, indent=2)+'\n')

def run(command, log):
    process = subprocess.run(list(map(str,command)), capture_output=True, text=True)
    log.write_text(process.stdout+'\n'+process.stderr)
    if process.returncode: raise RuntimeError(f'{command}: {process.returncode}\n{process.stdout}\n{process.stderr}')
    return process.stdout

def build():
    prepare()
    BUILD.mkdir(exist_ok=True)
    manifest = (FROZEN/'Purs_Test_ListOps/Cargo.toml').read_text().split('[dependencies]')[1]
    names = [m.replace('-','_') for m in re.findall(r'^([\w-]+)\s*=',manifest,re.M)]
    externs = []
    for name in names:
        candidates = list(DEPS.glob(f'lib{name}-*.rlib'))
        assert len(candidates)==1,(name,candidates)
        externs += ['--extern', f'{name}={candidates[0]}']
    flags = ['--edition=2021','-C','opt-level=1','-C','debuginfo=2','-L',f'dependency={DEPS}']
    binaries = {}
    for name in VARIANTS:
        binaries[name] = {}
        for counted in [True,False]:
            suffix = '-counted' if counted else ''
            source = HERE/f'ListOps-{name}{suffix}.rs'
            lib = BUILD/f'lib{name}{suffix}.rlib'
            run(['rustc',*flags,'--crate-type=rlib','--crate-name=Purs_Test_ListOps',
                '-C',f'metadata=list_{name}_{counted}',source,'-o',lib,*externs], BUILD/f'compile-{name}{suffix}.log')
            binary = BUILD/f'checks-{name}{suffix}'
            cfg = ['--cfg','list_counted'] if counted else []
            run(['rustc',*flags,*cfg,HERE/'checks.rs','-o',binary,
                '--extern',f'Purs_Test_ListOps={lib}',*externs], BUILD/f'compile-checks-{name}{suffix}.log')
            binaries[name]['counted' if counted else 'plain'] = {'path':str(binary),'sha256':digest(binary)}
        print('compiled',name,flush=True)
    metadata = json.loads((HERE/'metadata.json').read_text())
    metadata['binaries'] = binaries
    metadata['harness_sha256'] = digest(HERE/'checks.rs')
    metadata['script_sha256'] = digest(HERE/'probe.py')
    (HERE/'metadata.json').write_text(json.dumps(metadata,indent=2)+'\n')

def check():
    results = {}
    for name in VARIANTS:
        print('validating',name,flush=True)
        rows = []
        for counted in [True,False]:
            suffix = '-counted' if counted else ''
            output = run([BUILD/f'checks-{name}{suffix}',name], HERE/f'validation-{name}{suffix}.log')
            parsed = [json.loads(line) for line in output.splitlines() if line.startswith('{')]
            if counted: rows = parsed
            else: assert parsed[-1]['checks_passed']
        results[name] = rows
    (HERE/'validation.json').write_text(json.dumps(results,indent=2)+'\n')
    print('All four cumulative variants passed, counted and uninstrumented. No timing.',flush=True)

if __name__ == '__main__':
    parser=argparse.ArgumentParser();parser.add_argument('command',choices=['prepare','build','check'])
    args=parser.parse_args()
    globals()[args.command]()

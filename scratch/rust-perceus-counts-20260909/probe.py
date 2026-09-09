"""Instrument the generated RBTree kernel without modifying the runner output."""
from pathlib import Path
import hashlib
import json
import re
import subprocess
from collections import defaultdict

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
BUILD = HERE / 'build'
SOURCE = ROOT / 'run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs'

def kernel(source):
    starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
    result = [''] * len(source.splitlines())
    enum_start = source.index('#[derive(Clone')
    ranges = [(enum_start, starts[0].start())]
    names = {'max', 'makeBlack', 'depth', 'balance', 'ins', 'insert', 'buildTree'}
    for i, match in enumerate(starts):
        name = match[1].removeprefix('Test_RBTree_')
        if name in names or '__purust_rebuild_' in name or name.endswith('__purust_reuse'):
            ranges.append((match.start(), starts[i+1].start() if i+1 < len(starts) else len(source)))
    for start, end in ranges:
        line = source[:start].count('\n')
        part = source[start:end].splitlines()
        result[line:line+len(part)] = part
    return '\n'.join(result) + '\n'

def borrow_projection_receivers(source, constructor='crate::Tree::T'):
    # Prototype only: eliminate ownership of a projected child used immediately
    # as the receiver of as_ref(). The retained local parent owns that child.
    pattern = re.compile(r'\(\{ if let ' + re.escape(constructor) +
        r'\([^()]*\) = \([A-Za-z_][A-Za-z_0-9]*\)\.as_ref\(\) '
        r'\{ f\.clone\(\) \} else \{ unreachable!\(\) \} \}\)\.as_ref\(\)')
    return pattern.subn(lambda m: m[0].replace('f.clone()', 'f'), source)

def count(side, source):
    counted = source.replace('std::rc::Rc', 'tracked::Rc')
    (BUILD/f'kernel-{side}-counted.rs').write_text(counted)
    (BUILD/f'count-{side}.rs').write_text('#![allow(warnings)]\n#[path="../tracked_rc.rs"] mod tracked;\n'
        f'include!("kernel-{side}-counted.rs");\ninclude!("../count_harness.rs");\n')
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(BUILD/f'count-{side}.rs'),
                    '-o', str(BUILD/f'count-{side}')], check=True)
    output = subprocess.check_output([str(BUILD/f'count-{side}')], text=True)
    (BUILD/f'events-{side}.tsv').write_text(output)
    return output

def summarize(output):
    events = []
    phases = defaultdict(lambda: defaultdict(int))
    sites = defaultdict(int)
    for line in output.splitlines():
        phase, op, kind, site, count = line.split('\t')
        count, site = int(count), int(site)
        events.append(dict(phase=phase, operation=op, kind=kind, line=site, count=count))
        phases[phase][op] += count
        if op == 'clone' and phase.startswith('unique_'): sites[(phase, site, kind)] += count
    totals = defaultdict(int)
    for counters in phases.values():
        for op, count in counters.items(): totals[op] += count
    assert totals['new'] + totals['clone'] == sum(totals[op] for op in
        ['drop_last', 'drop_shared', 'unwrap_unique', 'unwrap_shared']), totals
    assert totals['new'] == totals['drop_last'] + totals['unwrap_unique'], totals
    return {'phases': dict(phases), 'events': events}, sites

def main():
    BUILD.mkdir(exist_ok=True)
    source = SOURCE.read_text()
    original = kernel(source)
    (BUILD/'kernel-native.rs').write_text(original)
    prototype, replacements = borrow_projection_receivers(original)
    assert replacements > 0
    (BUILD/'kernel-prototype.rs').write_text(prototype)
    result = {'source': str(SOURCE), 'source_sha256': hashlib.sha256(source.encode()).hexdigest(),
        'method': 'Observed logical Rc API operations through a one-word wrapper around std::rc::Rc; not machine instructions or timings. Included kernel preserves original source line numbers.',
        'rustc': subprocess.check_output(['rustc','--version'],text=True).strip(),
        'prototype_replacements': replacements, 'variants': {},
        'checks': 'Four rotations; ordering/red-black invariants; all exact keys of 200 retained versions; weak-reference exclusion; global handle and payload lifetime balances.'}
    for side, code in [('before', original), ('prototype', prototype)]:
        data, sites = summarize(count(side, code))
        result['variants'][side] = data
        print(side, json.dumps(data['phases'],indent=2),flush=True)
        print('Top unique-path clone sites:',flush=True)
        for (phase, line, kind), n in sorted(sites.items(),key=lambda x:-x[1])[:8]:
            print(phase, line, kind, n,flush=True)
    for phase in result['variants']['before']['phases']:
        a = result['variants']['before']['phases'][phase]
        b = result['variants']['prototype']['phases'][phase]
        for op in ['new', 'drop_last', 'unwrap_unique', 'unwrap_shared', 'get_mut_unique', 'get_mut_shared', 'get_mut_weak']:
            assert a.get(op,0) == b.get(op,0), (phase,op,a,b)
        assert a.get('clone',0)-b.get('clone',0) == a.get('drop_shared',0)-b.get('drop_shared',0), phase
    (HERE/'counts.json').write_text(json.dumps(result,indent=2)+'\n')

if __name__ == '__main__': main()

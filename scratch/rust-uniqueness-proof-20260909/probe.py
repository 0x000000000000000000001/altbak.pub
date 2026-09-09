"""Isolated, safe-Rust uniqueness-proof experiments on the generated RBTree."""
from pathlib import Path
import argparse
import hashlib
import importlib.util
import json
import re
import statistics
import subprocess

HERE = Path(__file__).resolve().parent
PREVIOUS = HERE.parent / 'rust-perceus-counts-20260909'
spec = importlib.util.spec_from_file_location('previous_probe', PREVIOUS/'probe.py')
previous = importlib.util.module_from_spec(spec)
spec.loader.exec_module(previous)
BUILD = HERE/'build'
SOURCE = previous.SOURCE
DEPS = previous.ROOT/'run/bak/rust/output/purust_output/target/release/deps'
HARNESS = '''
#[global_allocator] static ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;
fn main() {
    for i in 0..16 {
        let start = std::time::Instant::now();
        let tree = Test_RBTree_buildTree(std::hint::black_box(100000), std::rc::Rc::new(Tree::E));
        assert_eq!(Test_RBTree_depth(tree), 22);
        let ns = start.elapsed().as_nanos();
        if i > 0 { println!("{}", ns); }
    }
}
'''

def local_proof(source):
    """Keep the original exclusive borrow for a direct take/reconstruction."""
    changed = 0
    result = []
    for line in source.splitlines(keepends=True):
        match = re.search(r'\*std::rc::Rc::get_mut\(&mut (\w+)\)\.unwrap\(\) = _rebuilt;', line)
        if match:
            cell = match[1]
            old = f'let _taken = std::rc::Rc::get_mut(&mut {cell}).and_then(|node| node.__purust_take());'
            assert line.count(old) == 1
            line = line.replace(old,
                f'let mut _unique_slot = std::rc::Rc::get_mut(&mut {cell}); '
                'let _taken = _unique_slot.as_deref_mut().and_then(|node| node.__purust_take());')
            line = line.replace(match[0], '*_unique_slot.unwrap() = _rebuilt;')
            changed += 1
        result.append(line)
    assert changed == 3, changed
    return ''.join(result), changed

def variants():
    original = previous.kernel(SOURCE.read_text())
    local, n = local_proof(original)
    worker, workers = worker_proof(original)
    return {'before': original, 'local': local, 'worker': worker}, {'local_sites': n, **workers}

def mask_comments(source):
    return re.sub(r'/\*.*?\*/|//[^\n]*', lambda m: ' ' * len(m[0]), source, flags=re.S)

def close_paren(source, start):
    masked = mask_comments(source)
    assert masked[start] == '('
    depth = 1
    for i in range(start+1, len(masked)):
        if masked[i] == '(': depth += 1
        elif masked[i] == ')': depth -= 1
        if depth == 0: return i
    raise AssertionError('Unclosed call')

def worker_proof(source):
    """The caller keeps its Rc and lends the extracted payload slot to a private worker.

    The worker fills that slot and returns (). Its children remain owned Rc values.
    Rust checks that the owning Rc cannot be moved/aliased while this borrow lives.
    No unsafe and no change to the public worker or shared/weak fallback.
    """
    worker = 'Test_RBTree_balance__purust_reuse'
    helper = 'Test_RBTree___purust_rebuild_T'
    fill = 'Test_RBTree___purust_fill_T'
    start = source.index('fn '+worker+'(')
    front, body = source[:start], source[start:]
    signature = 'mut __purust_cell: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree>'
    assert body.count(signature) == 1
    body = body.replace(signature, '__purust_cell: &mut crate::Tree) -> ()', 1)
    edits = []
    for match in re.finditer(re.escape(helper)+r'\(', body):
        end = close_paren(body, match.end()-1)
        # Only the outer return builder owns the parent's slot.
        if re.search(r',\s*__purust_cell\s*$', mask_comments(body[match.end():end])):
            edits.append((match.start(), match.end()-1, fill))
    assert len(edits) > 0
    for begin, end, replacement in reversed(edits):
        body = body[:begin]+replacement+body[end:]
    n_calls = 0
    lines = []
    for line in front.splitlines(keepends=True):
        match = re.search(re.escape(worker)+r'\(', line)
        if match:
            cell_match = re.search(r'let _taken = std::rc::Rc::get_mut\(&mut (\w+)\)\.and_then\(\|node\| node\.__purust_take\(\)\);', line)
            assert cell_match
            cell = cell_match[1]
            end = close_paren(line, match.end()-1)
            call = line[match.start():end+1]
            assert call.endswith(', '+cell+')')
            call = call[:-len(cell)-1]+'_unique_slot.unwrap())'
            line = line[:match.start()]+'{ '+call+'; '+cell+' }'+line[end+1:]
            line = line.replace(cell_match[0],
                f'let mut _unique_slot = std::rc::Rc::get_mut(&mut {cell}); '
                'let _taken = _unique_slot.as_deref_mut().and_then(|node| node.__purust_take());')
            n_calls += 1
        lines.append(line)
    assert n_calls == 2, n_calls
    definition = f'\nfn {fill}(a0: crate::Color, a1: std::rc::Rc<crate::Tree>, a2: i64, a3: std::rc::Rc<crate::Tree>, slot: &mut crate::Tree) {{\n    *slot = crate::Tree::T(a0, a1, a2, a3);\n}}\n'
    return ''.join(lines)+body+definition, {'private_worker_calls': n_calls, 'private_worker_return_sites': len(edits)}

def compile_native(versions, assembly=False):
    mimalloc, = DEPS.glob('libmimalloc-*.rlib')
    common = ['rustc', '--edition=2021', '-C', 'opt-level=1',
              '--extern', f'mimalloc={mimalloc}', '-L', f'dependency={DEPS}']
    for name, code in versions.items():
        path = BUILD/f'time-{name}.rs'
        path.write_text('#![allow(warnings)]\n'+code+HARNESS)
        subprocess.run(common + [str(path), '-o', str(BUILD/f'time-{name}')], check=True)
        if assembly:
            subprocess.run(common + ['-C', 'debuginfo=line-tables-only', str(path), '--emit',
                f'asm={BUILD}/{name}.s,llvm-ir={BUILD}/{name}.ll'], check=True)

def measure(versions, output, pairs):
    runs = {name: [] for name in versions}
    for pair in range(pairs):
        names = list(versions) if pair % 2 == 0 else list(reversed(versions))
        for name in names:
            values = list(map(int, subprocess.check_output([str(BUILD/f'time-{name}')], text=True).split()))
            assert len(values) == 15
            runs[name].append(values)
            print(pair+1, name, statistics.median(values)/1e6, 'ms', flush=True)
    result = {'method': f'{pairs} alternating process pairs, 15 samples after one warmup; '
              'O1/mimalloc/native Rc; build+depth+destruction. No concurrent instrumentation or compilation.',
              'source_sha256': hashlib.sha256(SOURCE.read_bytes()).hexdigest(),
              'runs_ns': runs, 'median_ms': {name: statistics.median(sum(v, []))/1e6 for name, v in runs.items()}}
    (HERE/output).write_text(json.dumps(result, indent=2)+'\n')
    print(result['median_ms'], flush=True)

def count(versions):
    results = {}
    for name, code in versions.items():
        (BUILD/f'kernel-{name}.rs').write_text(code.replace('std::rc::Rc', 'tracked::Rc'))
        path = BUILD/f'count-{name}.rs'
        path.write_text('#![allow(warnings)]\n'
            f'#[path="{PREVIOUS}/tracked_rc.rs"] mod tracked;\n'
            f'include!("kernel-{name}.rs");\ninclude!("{PREVIOUS}/count_harness.rs");\n')
        subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(BUILD/f'count-{name}')], check=True)
        output = subprocess.check_output([str(BUILD/f'count-{name}')], text=True)
        data, _ = previous.summarize(output)
        results[name] = data
        print(name, json.dumps(data['phases']), flush=True)
    baseline = results['before']['phases']
    for name, data in results.items():
        for phase, before in baseline.items():
            for op in ['new', 'clone', 'drop_last', 'drop_shared', 'unwrap_unique', 'unwrap_shared', 'get_mut_shared', 'get_mut_weak']:
                assert before.get(op, 0) == data['phases'][phase].get(op, 0), (name, phase, op)
    (HERE/'counts.json').write_text(json.dumps(results, indent=2)+'\n')

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('mode', choices=['time', 'count'])
    parser.add_argument('--assembly', action='store_true')
    parser.add_argument('--output', default='timings-local.json')
    parser.add_argument('--variants', nargs='+')
    parser.add_argument('--pairs', type=int, default=3)
    parser.add_argument('--reuse-binaries', action='store_true')
    args = parser.parse_args()
    BUILD.mkdir(exist_ok=True)
    versions, sites = variants()
    selected = args.variants or (['before', 'local'] if args.mode == 'time' else list(versions))
    versions = {name: versions[name] for name in selected}
    print(sites, flush=True)
    if args.mode == 'count': count(versions)
    else:
        if not args.reuse_binaries: compile_native(versions, args.assembly)
        measure(versions, args.output, args.pairs)

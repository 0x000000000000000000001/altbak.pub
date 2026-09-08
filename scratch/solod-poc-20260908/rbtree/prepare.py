#!/usr/bin/env python3
"""Extract actual generated RBTree; mechanically adapt syntax/allocation only."""
import difflib
import hashlib
import json
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parent
SOURCE = ROOT.parents[2] / 'run/bak/go/output/purescript/Test_RBTree.go'
text = SOURCE.read_text()
start = text.index('type Constructor_Test_RBTree_T struct {')
raw = text[start:]
assert 'gopurs_runtime' not in raw and 'unsafe.' not in raw
syntax = raw
for label in ('depth', 'ins', 'buildTree'):
    assert f'if false {{ continue {label} }}' in syntax
    syntax = syntax.replace(f'{label}:\n', '', 1)
    syntax = syntax.replace(f'if false {{ continue {label} }}\n', '')
    syntax = syntax.replace(f'continue {label}', 'continue')
syntax, dead_count = re.subn(r'^__t\d+ = func\(\) \*Constructor_Test_RBTree_T \{ panic\("unreachable"\) \}\(\)\n', '', syntax, flags=re.M)
syntax, panic_count = re.subn(r'^__t\d+ = func\(\) (?:\*Constructor_Test_RBTree_T|int64) \{ panic\("Failed pattern match"\) \}\(\)$', 'panic("Failed pattern match")', syntax, flags=re.M)
assert dead_count == 1 and panic_count == 4
assert 'func()' not in syntax
arena = syntax
allocations = 0
while '&Constructor_Test_RBTree_T{' in arena:
    arena, count = re.subn(r'&Constructor_Test_RBTree_T\{([^{}]*)\}', r'AllocNode(\1)', arena)
    assert count > 0
    allocations += count

helper = '''
// The harness owns this fixed backing allocation for the process lifetime.
// Reset only after the previous tree and its shared intermediate nodes are dead.
var nodePool []Constructor_Test_RBTree_T
var nodesUsed int
func BeginArena(pool []Constructor_Test_RBTree_T) {
    nodePool = pool
    nodesUsed = 0
}
func ResetArena() { nodesUsed = 0 }
func NodesUsed() int { return nodesUsed }
func AllocNode(rc uint32, color uint32, left *Constructor_Test_RBTree_T, key int64, right *Constructor_Test_RBTree_T) *Constructor_Test_RBTree_T {
    if nodesUsed >= len(nodePool) { panic("node arena exhausted") }
    p := &nodePool[nodesUsed]
    nodesUsed++
    *p = Constructor_Test_RBTree_T{rc, color, left, key, right}
    return p
}
'''
for name, body in [('original', raw), ('syntax', syntax), ('arena', arena + helper)]:
    directory = ROOT / name
    (directory / 'kernels').mkdir(parents=True, exist_ok=True)
    (directory / 'cmd').mkdir(exist_ok=True)
    (directory / 'go.mod').write_text('module rbpoc\n\ngo 1.22\n')
    (directory / 'kernels/kernels.go').write_text('package kernels\n\n' + body)
    if (ROOT / 'validate.go').exists():
        (directory / 'kernels/validate.go').write_text((ROOT / 'validate.go').read_text())
    driver = (ROOT / 'driver.go.txt').read_text()
    setup = '''
    bits := int64(0)
    for k := n; k > 0; k >>= 1 { bits++ }
    capacity := (n+1)*(6*bits+4)
    pool := make([]kernels.Constructor_Test_RBTree_T, int(capacity))
    kernels.BeginArena(pool)
'''
    driver = driver.replace('// ARENA_SETUP', setup if name == 'arena' else '')
    driver = driver.replace('// ARENA_RESET', 'kernels.ResetArena()' if name == 'arena' else '')
    driver = driver.replace('ALLOCATIONS', 'kernels.NodesUsed()' if name == 'arena' else '-1')
    (directory / 'cmd/main.go').write_text(driver)

(ROOT / 'syntax.diff').write_text(''.join(difflib.unified_diff(raw.splitlines(True), syntax.splitlines(True), fromfile='generated-go', tofile='syntax-adapted-go')))
(ROOT / 'allocation.diff').write_text(''.join(difflib.unified_diff(syntax.splitlines(True), (arena+helper).splitlines(True), fromfile='syntax-adapted-go', tofile='arena-go-and-solod')))
(ROOT / 'sources.json').write_text(json.dumps({
    'source': str(SOURCE), 'line': text[:start].count('\n')+1,
    'source_sha256': hashlib.sha256(SOURCE.read_bytes()).hexdigest(),
    'extracted_sha256': hashlib.sha256(raw.encode()).hexdigest(),
    'constructor_sites_replaced': allocations,
    'failed_match_iife_replaced': panic_count,
    'unreachable_iife_removed': dead_count,
}, indent=2)+'\n')
print(f'Extracted actual RBTree; {allocations} constructor sites redirected to checked arena.')

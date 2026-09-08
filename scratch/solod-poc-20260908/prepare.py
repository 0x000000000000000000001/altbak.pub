#!/usr/bin/env python3
"""Extract generated kernels and make narrowly scoped Solod syntax adaptations."""
import difflib
import hashlib
import json
from pathlib import Path
import re

ROOT = Path(__file__).resolve().parent
SOURCE = ROOT.parent.parent / 'run/bak/go/output/purescript'
TARGETS = [('Fib', 'fib'), ('Ackermann', 'ackermann'), ('TCO', 'deepTailRec')]
manifest = []
original = []
adapted = []
for module, name in TARGETS:
    path = SOURCE / f'Test_{module}.go'
    text = path.read_text()
    symbol = f'Call_Test_{module}_{name}'
    start = text.index(f'func {symbol}(')
    # These three numeric functions have balanced braces inside their only
    # string literal ("unreachable") and contain no brace-bearing comments.
    opening = text.index('{', start)
    depth = 1
    end = opening + 1
    while depth:
        depth += (text[end] == '{') - (text[end] == '}')
        end += 1
    raw = text[start:end] + '\n'
    assert raw.count('for {') == 1, 'single loop required for label removal'
    assert f'if false {{ continue {name} }}' in raw
    clean = raw.replace(f'{name}:\n', '', 1)
    clean = clean.replace(f'if false {{ continue {name} }}\n', '')
    clean = clean.replace(f'continue {name}', 'continue')
    clean = re.sub(r'^__t\d+ = func\(\) int64 \{ panic\("unreachable"\) \}\(\)\n', '', clean, flags=re.M)
    assert 'func()' not in clean
    original.append(raw)
    adapted.append(clean)
    manifest.append({'source': str(path), 'line': text[:start].count('\n') + 1,
                     'symbol': symbol, 'source_sha256': hashlib.sha256(path.read_bytes()).hexdigest(),
                     'extracted_sha256': hashlib.sha256(raw.encode()).hexdigest()})

driver = '''package main
import (
    "fmt"
    "os"
    "strconv"
    "poc/kernels"
)
func main() {
    if len(os.Args) != 4 { panic("usage: runner kind iterations n") }
    count, err := strconv.ParseInt(os.Args[2],10,64); if err != nil { panic(err) }
    n, err := strconv.ParseInt(os.Args[3],10,64); if err != nil { panic(err) }
    var sum int64
    switch os.Args[1] {
    case "fib":
        for i := int64(0); i < count; i++ { sum += kernels.Call_Test_Fib_fib(n+(i&1)) }
    case "ackermann":
        for i := int64(0); i < count; i++ { sum += kernels.Call_Test_Ackermann_ackermann(3,n+(i&1)) }
    case "tco":
        for i := int64(0); i < count; i++ { sum += kernels.Call_Test_TCO_deepTailRec(n+(i&1),0) }
    default: panic("unknown kind")
    }
    fmt.Println(sum)
}
'''
for name, bodies in [('original', original), ('adapted', adapted)]:
    directory = ROOT / name
    (directory / 'kernels').mkdir(parents=True, exist_ok=True)
    (directory / 'cmd').mkdir(exist_ok=True)
    (directory / 'go.mod').write_text('module poc\n\ngo 1.22\n')
    (directory / 'kernels/kernels.go').write_text('package kernels\n\n' + '\n'.join(bodies))
    (directory / 'cmd/main.go').write_text(driver)
(ROOT / 'adaptations.diff').write_text(''.join(difflib.unified_diff(
    ('\n'.join(original)).splitlines(True), ('\n'.join(adapted)).splitlines(True),
    fromfile='extracted-generated-go', tofile='solod-compatible-go')))
(ROOT / 'sources.json').write_text(json.dumps(manifest, indent=2) + '\n')
print('Extracted 3 generated functions; wrote original, adapted, adaptations.diff and sources.json')

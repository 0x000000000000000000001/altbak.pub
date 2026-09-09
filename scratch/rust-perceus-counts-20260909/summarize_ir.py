"""Summarize optional LLVM/assembly diagnostics produced by EMIT_IR=1 time.py."""
import json
import re
from pathlib import Path

HERE = Path(__file__).resolve().parent
variants = {}
excerpt = []
for side in ['before','prototype']:
    text = (HERE/f'build/{side}.ll').read_text()
    variants[side] = {}
    for match in re.finditer(r'^define[^\n]*?@([^ (]+)[^\n]*\{', text, re.M):
        if 'RBTree' not in match[1]: continue
        body = text[match.end():text.index('\n}',match.end())]
        name = re.search(r'Test_RBTree_\w+?(?=17h)',match[1])[0]
        variants[side][name] = {
            'ir_lines':body.count('\n'),
            'static_loads':len(re.findall(r' = load ',body)),
            'static_stores':len(re.findall(r'^  store ',body,re.M)),
            'static_branches':len(re.findall(r'^  br ',body,re.M))}
    lines = (HERE/f'build/{side}.s').read_text().splitlines()
    # The standalone source has one additional #![allow(warnings)] line.
    hit = next(i for i,line in enumerate(lines) if re.search(r'\.loc\s+\d+\s+1742\s',line))
    excerpt.extend([side, *lines[max(0,hit-14):hit+24], ''])
result = {'method':'Static optimized LLVM IR instruction-site counts; not dynamic execution counts or a speed estimate.',
    'variants':variants}
(HERE/'assembly-summary.json').write_text(json.dumps(result,indent=2)+'\n')
(HERE/'assembly-excerpt.txt').write_text('\n'.join(excerpt)+'\n')

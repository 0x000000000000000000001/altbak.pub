"""Archive compact reproducibility metadata and optimized-machine-code evidence."""
import hashlib
import json
import re
import subprocess
import probe

summary = {}
for name in ['before', 'local', 'worker']:
    text = (probe.BUILD/f'{name}.ll').read_text()
    summary[name] = {}
    for match in re.finditer(r'^define[^\n]*?@([^ (]+)[^\n]*\{', text, re.M):
        name_match = re.search(r'(?:Test_RBTree_\w+?|4main)(?=17h)', match[1])
        if not name_match: continue
        body = text[match.end():text.index('\n}', match.end())]
        summary[name][name_match[0]] = {
            'ir_lines': body.count('\n'),
            'static_loads': len(re.findall(r' = load ', body)),
            'static_stores': len(re.findall(r'^  store ', body, re.M)),
            'static_branches': len(re.findall(r'^  br ', body, re.M))}
(probe.HERE/'assembly-summary.json').write_text(json.dumps({
    'method': 'Static optimized LLVM sites, not executed instructions or a performance prediction.',
    'variants': summary}, indent=2)+'\n')
lines = (probe.BUILD/'before.s').read_text().splitlines()
start = next(i for i, line in enumerate(lines) if re.match(r'^__ZN.*Test_RBTree___purust_rebuild_T.*:$', line))
snippet = ['Original rebuild helper: native arm64 uniqueness check survives O1.',
           'The loads at Rc-box offsets 8 and 0 check weak and strong counts against 1.',
           *lines[start:start+58]]
lines = (probe.BUILD/'worker.s').read_text().splitlines()
start = next(i for i, line in enumerate(lines) if re.match(r'^__ZN.*Test_RBTree___purust_fill_T.*:$', line))
snippet.extend(['', 'Borrowed-slot fill helper: receives the payload pointer; no Rc uniqueness check.',
                *lines[start:start+75]])
(probe.HERE/'assembly-excerpt.txt').write_text('\n'.join(snippet)+'\n')
paths = {'altbak_worktree': probe.previous.ROOT, 'purust': probe.previous.ROOT.parent/'purust/purust',
         'pbo_worktree': probe.previous.ROOT.parent/'purescript-backend-optimizer-purust'}
metadata = {
    'revisions': {name: subprocess.check_output(['git', '-C', str(path), 'rev-parse', 'HEAD'], text=True).strip()
                  for name, path in paths.items()},
    'source': str(probe.SOURCE), 'source_sha256': hashlib.sha256(probe.SOURCE.read_bytes()).hexdigest(),
    'rustc': subprocess.check_output(['rustc', '--version'], text=True).strip(),
    'README_reference': {'path': str(probe.previous.ROOT.parent/'altbak.pub/README.md'),
                         'rbtree_ms': 18.985, 'total_ms': 21.82, 'native_rbtree_ms': 36.07, 'native_total_ms': 36.13},
    'reused_harnesses': {name: hashlib.sha256((probe.PREVIOUS/name).read_bytes()).hexdigest()
                         for name in ['probe.py', 'tracked_rc.rs', 'count_harness.rs']},
    'binary_sha256': {name: hashlib.sha256((probe.BUILD/f'time-{name}').read_bytes()).hexdigest()
                      for name in ['before', 'local', 'worker']}}
(probe.HERE/'metadata.json').write_text(json.dumps(metadata, indent=2)+'\n')
print(json.dumps(summary, indent=2))

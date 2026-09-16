from pathlib import Path
import hashlib, json, re, subprocess
HERE = Path(__file__).resolve().parent
SOURCE = Path('/Users/0x1/Documents/htdocs/altbak.pub-purust/output/purust_output/Purs_Test_RBTree/src/lib.rs')
source = SOURCE.read_text()
starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
excluded = {'R', 'B', 'E', 'T', 'describe', 'act'}
ranges = [(source.index('#[derive(Clone'), starts[0].start())]
for i, match in enumerate(starts):
    if match[1].removeprefix('Test_RBTree_') not in excluded:
        ranges.append((match.start(), starts[i + 1].start() if i + 1 < len(starts) else len(source)))
lines = [''] * len(source.splitlines())
for start, end in ranges:
    at = source[:start].count('\n')
    part = source[start:end].splitlines()
    lines[at:at+len(part)] = part
kernel = '\n'.join(lines) + '\n'
(HERE / 'kernel.rs').write_text(kernel)
(HERE / 'manifest.json').write_text(json.dumps({
    'source': str(SOURCE), 'source_sha256': hashlib.sha256(source.encode()).hexdigest(),
    'kernel_sha256': hashlib.sha256(kernel.encode()).hexdigest(),
    'kernel_transform': 'Only extraction: retained enums and kernel function bodies are unchanged.',
    'variant': 'Global allocator fresh block per invocation; Rc and drop semantics unchanged.',
}, indent=2)+'\n')
subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=3', str(HERE / 'validate.rs'), '-o', str(HERE / 'validate')], check=True)
result = subprocess.check_output([str(HERE/'validate')], text=True, timeout=60)
(HERE/'validation.txt').write_text(result)
print(result, end='')

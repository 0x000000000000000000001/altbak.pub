"""Check invariants against the unchanged functions extracted from generated Rust."""
from pathlib import Path
import ast
import difflib
import json
import re
import subprocess

HERE = Path(__file__).resolve().parent
ROOT = HERE.parents[1]
after = (ROOT/'output/purust_output/Purs_Test_RBTree/src/lib.rs').read_text()
before = (HERE/'RBTree-before.rs').read_text()
(HERE/'RBTree-after.rs').write_text(after)
(HERE/'RBTree.diff').write_text(''.join(difflib.unified_diff(
    before.splitlines(True), after.splitlines(True), fromfile='RBTree-before.rs', tofile='RBTree-after.rs')))
# In this module the entire change must be the removal of native-local parent
# clones immediately before as_ref(). The child clones must stay intact.
borrowed = re.sub(r'\((purs_local_\w+)\.clone\(\)\)\.as_ref\(\)', r'(\1).as_ref()', before)
assert borrowed == after
emission = {
    'clone_sites_before': before.count('.clone()'),
    'clone_sites_after': after.count('.clone()'),
    'field_clone_sites_before': before.count('{ f.clone() }'),
    'field_clone_sites_after': after.count('{ f.clone() }'),
    'only_local_parent_clones_removed': True,
}
(HERE/'emission.json').write_text(json.dumps(emission, indent=2)+'\n')
print(emission, flush=True)

# Reuse the audit's ordering, black-height, red-child, duplicate and persistence
# checks without executing its prototype transforms or timing loops.
audit = ast.parse((HERE.parent/'rust-audit-20260908/measure.py').read_text())
harness = next(ast.literal_eval(n.value) for n in audit.body
               if isinstance(n, ast.Assign) and any(isinstance(t, ast.Name) and t.id == 'harness' for t in n.targets))
checks = harness[harness.index('fn validate('):harness.index('fn run(')]
starts = list(re.finditer(r'^pub fn (Test_RBTree_\w+)\(', after, re.M))
functions = {m[1]: after[m.start():starts[i+1].start() if i+1 < len(starts) else len(after)]
             for i, m in enumerate(starts)}
enums = after[after.index('#[derive(Clone)]'):starts[0].start()]
body = enums + '\n'.join(functions['Test_RBTree_'+name]
                         for name in ['max', 'makeBlack', 'depth', 'balance', 'ins', 'insert', 'buildTree'])
source = HERE/'rbtree-invariants.rs'
binary = HERE/'rbtree-invariants'
source.write_text('#![allow(warnings)]\n'+body+checks+'\nfn main() { correctness(); }\n')
subprocess.run(['rustc', '--edition=2021', str(source), '-o', str(binary)], check=True)
subprocess.run([str(binary)], check=True)
print('Generated RBTree: ordering, black heights, red children, duplicates and old version preserved.')

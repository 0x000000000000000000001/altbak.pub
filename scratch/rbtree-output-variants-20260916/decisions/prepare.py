from pathlib import Path
import re, hashlib, json, subprocess

HERE = Path(__file__).resolve().parent
SOURCE = Path('/Users/0x1/Documents/htdocs/altbak.pub-purust/output/purust_output/Purs_Test_RBTree/src/lib.rs')
source = SOURCE.read_text()
starts = list(re.finditer(r'^(?:pub )?fn (Test_RBTree_\w+)\(', source, re.M))
excluded = {'R', 'B', 'E', 'T', 'describe', 'act'}
ranges = [(source.index('#[derive(Clone'), starts[0].start())]
for index, match in enumerate(starts):
    end = starts[index + 1].start() if index + 1 < len(starts) else len(source)
    if match[1].removeprefix('Test_RBTree_') not in excluded:
        ranges.append((match.start(), end))
lines = [''] * len(source.splitlines())
for start, end in ranges:
    line = source[:start].count('\n')
    part = source[start:end].splitlines()
    lines[line:line + len(part)] = part
baseline = '\n'.join(lines) + '\n'
(HERE/'kernel-baseline.rs').write_text(baseline)
classifier = '''#[inline]
fn classify_rotation(color: &Color, left: &std::rc::Rc<Tree>, _key: &i64, right: &std::rc::Rc<Tree>) -> u8 {
    if !matches!(color, Color::B) { return 0; }
    if let Tree::T(Color::R, ll, _, lr) = left.as_ref() {
        if matches!(ll.as_ref(), Tree::T(Color::R, ..)) { return 1; }
        if matches!(lr.as_ref(), Tree::T(Color::R, ..)) { return 2; }
    }
    if let Tree::T(Color::R, rl, _, rr) = right.as_ref() {
        if matches!(rl.as_ref(), Tree::T(Color::R, ..)) { return 3; }
        if matches!(rr.as_ref(), Tree::T(Color::R, ..)) { return 4; }
    }
    0
}
'''
guard = '''fn Test_RBTree_balance__purust_child_rebuilds(color: &Color, left: &std::rc::Rc<Tree>, key: &i64, right: &std::rc::Rc<Tree>) -> bool {
    classify_rotation(color, left, key, right) == 0
}
'''
simplified, count = re.subn(r'^fn Test_RBTree_balance__purust_child_rebuilds\([^\n]*\n', guard, baseline, flags=re.M)
assert count == 1
simplified += classifier
(HERE/'kernel-simple-guard.rs').write_text(simplified)
call = 'if Test_RBTree_balance__purust_child_rebuilds(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3)'
replace = 'let rotation = classify_rotation(_purust_post_field_0, _purust_post_field_1, _purust_post_field_2, _purust_post_field_3); if rotation == 0'
assert simplified.count(call) == 2
fused = simplified.replace(call, replace)
call = 'else if Test_RBTree_balance__purust_permute_fields(_purust_post_slot)'
assert fused.count(call) == 2
fused = fused.replace(call, 'else if rotation == 1 && Test_RBTree_balance__purust_permute_fields(_purust_post_slot)')
guardline = 'if !Test_RBTree_balance__purust_permute_fields__matches(_purust_permute_0_0, _purust_permute_0_1, _purust_permute_0_2, _purust_permute_0_3) { return false; }'
assert fused.count(guardline) == 1
fused = fused.replace(guardline, '// LL shape already established by classify_rotation before calling this private worker.')
(HERE/'kernel-fused-guard.rs').write_text(fused)
# The original implementation uses the same public enum, enabling exact structural comparison.
reference = baseline[baseline.index('fn Test_RBTree___purust_rebuild_E'):]
reference = reference.replace('fn Test_RBTree_balance__purust_child_rebuilds(', 'pub fn Test_RBTree_balance__purust_child_rebuilds(')
(HERE/'reference.rs').write_text(reference)
results={}
for name in ['simple-guard','fused-guard']:
    main = '#![allow(warnings)]\ninclude!("kernel-'+name+'.rs");\nmod reference { include!("reference.rs"); }\ninclude!("check.rs");\n'
    (HERE/f'check-{name}.rs').write_text(main)
    subprocess.run(['rustc','--edition=2021','-C','opt-level=2',str(HERE/f'check-{name}.rs'),'-o',str(HERE/f'check-{name}')],check=True)
    results[name]=subprocess.check_output([str(HERE/f'check-{name}')],text=True,timeout=90).strip()
manifest = {'source':str(SOURCE),'sha256':hashlib.sha256(source.encode()).hexdigest(),'changes':{'simple-guard':'Replace only rebuild guard by match classifier; old LL permutation guard unchanged.','fused-guard':'Classify once at two post-child call sites; branch by classification and omit redundant LL predicate inside private permutation worker. All get_mut checks and generic fallbacks preserved.'},'checks':results,'timings':'not run; parent orchestrates sequentially'}
(HERE/'manifest.json').write_text(json.dumps(manifest,indent=2)+'\n')
print(json.dumps(manifest,indent=2))

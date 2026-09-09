"""Run the invariants again with native Rc, including mixed sharing and weak handles."""
import hashlib
import json
import re
import subprocess
import probe

versions, _ = probe.variants()
base = (probe.PREVIOUS/'count_harness.rs').read_text()
base = base[base.index('fn validate('):]
base = base.replace('tracked::Rc', 'std::rc::Rc')
base = re.sub(r'    tracked::(?:phase\("[^"]*"\)|print_events\(\));\n', '', base)
base = base.replace('fn main() {', 'fn original_checks() {')
extra = r'''
fn main() {
    original_checks();
    // Unique, strongly shared, weak-only and shared+weak recoloration.
    for shared in [false, true] {
        for weakly_shared in [false, true] {
            let left = std::rc::Rc::new(Tree::E);
            let right = std::rc::Rc::new(Tree::E);
            let left_weak = std::rc::Rc::downgrade(&left);
            let right_weak = std::rc::Rc::downgrade(&right);
            let tree = std::rc::Rc::new(Tree::T(Color::R, left, 1, right));
            let old = shared.then(|| tree.clone());
            let weak = weakly_shared.then(|| std::rc::Rc::downgrade(&tree));
            let address = std::rc::Rc::as_ptr(&tree);
            let tree = Test_RBTree_makeBlack(tree);
            check(&tree, 1);
            if !shared && !weakly_shared { assert_eq!(address, std::rc::Rc::as_ptr(&tree)); }
            if let Some(ref old) = old { assert!(matches!(old.as_ref(), Tree::T(Color::R, ..))); }
            if let Some(ref weak) = weak { assert_eq!(weak.upgrade().is_some(), shared); }
            assert!(left_weak.upgrade().is_some() && right_weak.upgrade().is_some());
            drop(tree); drop(old);
            assert!(left_weak.upgrade().is_none() && right_weak.upgrade().is_none());
        }
    }
    let mut tree = std::rc::Rc::new(Tree::E);
    let mut versions = Vec::new();
    let mut expected = Vec::new();
    let mut state = 1234567u64;
    // Most roots stay unique, but retained ancestors leave shared subtrees.
    for n in 0..512 {
        if n % 17 == 0 { versions.push((tree.clone(), expected.clone())); }
        let weak = (n % 23 == 0).then(|| std::rc::Rc::downgrade(&tree));
        state = state.wrapping_mul(6364136223846793005).wrapping_add(1);
        let key = (state >> 32) as i64 % 400;
        tree = Test_RBTree_insert(key, tree);
        if !expected.contains(&key) { expected.push(key); expected.sort(); }
        check(&tree, expected.len()); assert_eq!(keys(&tree), expected);
        for (old, expected) in &versions { check(old, expected.len()); assert_eq!(&keys(old), expected); }
        drop(weak);
    }
    let weak_roots: Vec<_> = versions.iter().map(|(old, _)| std::rc::Rc::downgrade(old)).collect();
    drop(tree); drop(versions);
    assert!(weak_roots.iter().all(|weak| weak.upgrade().is_none()));
    println!("native Rc checks passed");
}
'''
result = {'method': 'Native std::rc::Rc, no counter wrapper. Original four rotations, 100k nodes, '
          '200 retained versions; four strong/weak recoloration cases; 512 mixed-sharing insertions '
          'including duplicates; preserved keys/invariants/lifetimes.', 'variants': {}}
for name, code in versions.items():
    path = probe.BUILD/f'validate-{name}.rs'
    path.write_text('#![allow(warnings)]\n'+code+base+extra)
    binary = probe.BUILD/f'validate-{name}'
    subprocess.run(['rustc', '--edition=2021', '-C', 'opt-level=1', str(path), '-o', str(binary)], check=True)
    output = subprocess.check_output([str(binary)], text=True).strip()
    result['variants'][name] = {'result': output, 'source_sha256': hashlib.sha256(path.read_bytes()).hexdigest()}
    print(name, output, flush=True)
(probe.HERE/'validation.json').write_text(json.dumps(result, indent=2)+'\n')

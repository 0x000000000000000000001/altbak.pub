from pathlib import Path
import hashlib
import json
import re

HERE = Path(__file__).resolve().parent
SOURCE = Path('/Users/0x1/Documents/htdocs/altbak.pub/scratch/purust-fbip-20260916/benchmark/output/purust_output/Purs_Test_RBTree/src/lib.rs')
source = SOURCE.read_text()
# Remove external-library entry points only. Keep every benchmark kernel body.
kernel = source[source.index('#[derive(Clone, Copy)]'):]
removed = ['Test_RBTree_T', 'Test_RBTree_describe', 'Test_RBTree_act']
for name in removed:
    pattern = r'^pub fn ' + re.escape(name) + r'\([^\n]*\n.*?(?=^(?:pub )?fn |\Z)'
    kernel, count = re.subn(pattern, '', kernel, flags=re.M | re.S)
    assert count == 1, (name, count)
assert 'purust_core::' not in re.sub(r'/\*.*?\*/', '', kernel, flags=re.S)
(HERE / 'kernel_baseline.rs').write_text(kernel)
count = kernel.count('std::rc::Rc::get_mut(')
unique = kernel.replace('std::rc::Rc::get_mut(', 'assumed_unique_mut(')
helper = r'''
// EXPERIMENTAL specialization only. The general persistent API must keep
// Rc::get_mut and its fallback. Preconditions for the optimized specialization:
// every nonempty cell reachable by the mutable sites has exactly one strong
// owner; no Weak reference; no other outstanding reference into the payload.
// Empty leaves may be shared, so those still use the standard dynamic check.
#[inline(always)]
fn assumed_unique_mut(cell: &mut std::rc::Rc<Tree>) -> Option<&mut Tree> {
    #[cfg(verify_unique)]
    UNIQUE_CALLS.fetch_add(1, std::sync::atomic::Ordering::Relaxed);
    if matches!(cell.as_ref(), Tree::E) {
        #[cfg(verify_unique)]
        EMPTY_CHECKS.fetch_add(1, std::sync::atomic::Ordering::Relaxed);
        return std::rc::Rc::get_mut(cell);
    }
    #[cfg(verify_unique)]
    {
        assert_eq!(std::rc::Rc::strong_count(cell), 1, "nonempty alias at specialized mutation");
        assert_eq!(std::rc::Rc::weak_count(cell), 0, "weak alias at specialized mutation");
    }
    // SAFETY: the experimental worker's explicit preconditions above replace
    // the same strong/weak uniqueness conditions tested by Rc::get_mut.
    Some(unsafe { &mut *(std::rc::Rc::as_ptr(cell) as *mut Tree) })
}
#[cfg(verify_unique)]
static UNIQUE_CALLS: std::sync::atomic::AtomicU64 = std::sync::atomic::AtomicU64::new(0);
#[cfg(verify_unique)]
static EMPTY_CHECKS: std::sync::atomic::AtomicU64 = std::sync::atomic::AtomicU64::new(0);
'''
(HERE / 'kernel_unique.rs').write_text(unique + helper)
empty_guard = '''    if matches!(cell.as_ref(), Tree::E) {
        #[cfg(verify_unique)]
        EMPTY_CHECKS.fetch_add(1, std::sync::atomic::Ordering::Relaxed);
        return std::rc::Rc::get_mut(cell);
    }
'''
assert helper.count(empty_guard) == 1
site_helper = helper.replace(empty_guard, '').replace(
    '// Empty leaves may be shared, so those still use the standard dynamic check.',
    '// Stronger site-level contract: every selected mutable cell is unique, even\n'
    '// after its payload has temporarily been replaced by E. Shared leaf E must\n'
    '// never reach these mutable sites. This requires tracking cell provenance.'
).replace('nonempty alias at specialized mutation', 'alias at specialized mutation site')
(HERE / 'kernel_unique_sites.rs').write_text(unique + site_helper)
manifest = dict(source=str(SOURCE), source_sha256=hashlib.sha256(source.encode()).hexdigest(),
                removed_non_kernel_functions=removed, replaced_get_mut_sites=count,
                baseline_sha256=hashlib.sha256(kernel.encode()).hexdigest(),
                unique_sha256=hashlib.sha256((unique+helper).encode()).hexdigest())
(HERE / 'manifest.json').write_text(json.dumps(manifest, indent=2)+'\n')
print(json.dumps(manifest, indent=2))

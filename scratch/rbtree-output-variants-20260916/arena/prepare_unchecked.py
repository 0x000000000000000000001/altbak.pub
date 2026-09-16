#!/usr/bin/env python3
from pathlib import Path
HERE = Path(__file__).resolve().parent
source = (HERE/'kernel.rs').read_text()
wrapper = '''
// Scratch-only unchecked indexing. All indices originate from the sentinel,
// newly appended nodes, or links between those nodes. Nodes are never removed.
// Compile with --cfg arena_verify to assert the invariant at EVERY access.
struct NodeStore(Vec<Node>);
impl std::ops::Deref for NodeStore {
    type Target = Vec<Node>;
    fn deref(&self) -> &Self::Target { &self.0 }
}
impl std::ops::DerefMut for NodeStore {
    fn deref_mut(&mut self) -> &mut Self::Target { &mut self.0 }
}
impl std::ops::Index<usize> for NodeStore {
    type Output = Node;
    #[inline(always)]
    fn index(&self, index: usize) -> &Node {
        #[cfg(arena_verify)]
        assert!(index < self.0.len(), "arena read index out of bounds");
        // SAFETY: internal links/sentinel only; append-only store preserves validity.
        unsafe { self.0.get_unchecked(index) }
    }
}
impl std::ops::IndexMut<usize> for NodeStore {
    #[inline(always)]
    fn index_mut(&mut self, index: usize) -> &mut Node {
        #[cfg(arena_verify)]
        assert!(index < self.0.len(), "arena write index out of bounds");
        // SAFETY: same range invariant; exclusive &mut access supplies alias safety.
        unsafe { self.0.get_unchecked_mut(index) }
    }
}
'''
source = source.replace('// Isolated, safe arena lowering of the generated Okasaki algorithm.', '// Scratch-only bounds-elision variant; algorithm otherwise identical to kernel.rs.')
source = source.replace('struct Arena { nodes: Vec<Node>, root: Index }', wrapper + '\nstruct Arena { nodes: NodeStore, root: Index }')
assert source.count('Self { nodes, root: 0 }') == 1
source = source.replace('Self { nodes, root: 0 }', 'Self { nodes: NodeStore(nodes), root: 0 }')
(HERE/'kernel-unchecked.rs').write_text(source)
(HERE/'check-unchecked.rs').write_text((HERE/'check.rs').read_text().replace('include!("kernel.rs")','include!("kernel-unchecked.rs")'))

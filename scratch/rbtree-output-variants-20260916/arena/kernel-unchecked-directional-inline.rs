// Scratch-only bounds-elision variant; algorithm otherwise identical to kernel.rs.
// A closed exclusive region replaces Rc ownership. Index zero is E.
#[cfg(not(arena_usize))]
type Index = u32;
#[cfg(arena_usize)]
type Index = usize;
#[derive(Clone, Copy)]
struct Node { key: i64, left: Index, right: Index, red: bool }

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

struct Arena { nodes: NodeStore, root: Index }
impl Arena {
    fn new(capacity: usize) -> Self {
        let mut nodes = Vec::with_capacity(capacity);
        nodes.push(Node { key: 0, left: 0, right: 0, red: false });
        Self { nodes: NodeStore(nodes), root: 0 }
    }
    #[inline(always)]
    fn red(&self, t: Index) -> bool { self.nodes[t as usize].red }
    #[inline(always)]
    fn rotate_right(&mut self, old: Index) -> Index {
        let promoted = self.nodes[old as usize].left;
        self.nodes[old as usize].left = self.nodes[promoted as usize].right;
        self.nodes[promoted as usize].right = old;
        promoted
    }
    #[inline(always)]
    fn rotate_left(&mut self, old: Index) -> Index {
        let promoted = self.nodes[old as usize].right;
        self.nodes[old as usize].right = self.nodes[promoted as usize].left;
        self.nodes[promoted as usize].left = old;
        promoted
    }
    #[inline(always)]
    fn balance<const LEFT: bool>(&mut self, root: Index) -> Index {
        let node = self.nodes[root as usize];
        if node.red { return root; }
        // A valid input subtree has no adjacent red nodes. Only the child
        // changed by ins can acquire a red-red edge; the other side is unchanged.
        let direction = if LEFT {
            if self.red(node.left) {
                let left = self.nodes[node.left as usize];
                if self.red(left.left) { 1 } else if self.red(left.right) { 2 } else { 0 }
            } else { 0 }
        } else {
            if self.red(node.right) {
                let right = self.nodes[node.right as usize];
                if self.red(right.left) { 3 } else if self.red(right.right) { 4 } else { 0 }
            } else { 0 }
        };
        let root = match direction {
            1 => self.rotate_right(root),
            2 => {
                let left = self.rotate_left(node.left);
                self.nodes[root as usize].left = left;
                self.rotate_right(root)
            },
            3 => {
                let right = self.rotate_right(node.right);
                self.nodes[root as usize].right = right;
                self.rotate_left(root)
            },
            4 => self.rotate_left(root),
            _ => return root,
        };
        let node = self.nodes[root as usize];
        self.nodes[root as usize].red = true;
        self.nodes[node.left as usize].red = false;
        self.nodes[node.right as usize].red = false;
        root
    }
    fn ins(&mut self, key: i64, root: Index) -> Index {
        if root == 0 {
            let next = Index::try_from(self.nodes.len()).expect("arena index overflow");
            self.nodes.push(Node { key, left: 0, right: 0, red: true });
            return next;
        }
        let node = self.nodes[root as usize];
        if key < node.key {
            let left = self.ins(key, node.left);
            self.nodes[root as usize].left = left;
            self.balance::<true>(root)
        } else if key > node.key {
            let right = self.ins(key, node.right);
            self.nodes[root as usize].right = right;
            self.balance::<false>(root)
        } else { root }
    }
    fn insert(&mut self, key: i64) {
        let root = self.ins(key, self.root);
        self.nodes[root as usize].red = false;
        self.root = root;
    }
    fn depth(&self, root: Index) -> i64 {
        if root == 0 { return 0; }
        let node = &self.nodes[root as usize];
        1 + self.depth(node.left).max(self.depth(node.right))
    }
}
#[inline(never)]
pub fn run(n: i64) -> i64 {
    #[cfg(not(arena_grow))]
    let capacity = usize::try_from(n.max(0)).unwrap().checked_add(1).unwrap();
    #[cfg(arena_grow)]
    let capacity = 0;
    let mut tree = Arena::new(capacity);
    for key in (1..=n).rev() { tree.insert(key); }
    let tree = std::hint::black_box(tree);
    let answer = tree.depth(tree.root);
    drop(tree);
    answer
}

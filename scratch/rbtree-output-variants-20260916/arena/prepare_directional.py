#!/usr/bin/env python3
"""Specialize balancing for the only child modified by recursive insertion."""
from pathlib import Path
HERE=Path(__file__).resolve().parent
old='''        let direction = if self.red(node.left) {
            let left = self.nodes[node.left as usize];
            if self.red(left.left) { 1 } else if self.red(left.right) { 2 } else { 0 }
        } else { 0 };
        let direction = if direction == 0 && self.red(node.right) {
            let right = self.nodes[node.right as usize];
            if self.red(right.left) { 3 } else if self.red(right.right) { 4 } else { 0 }
        } else { direction };'''
new='''        // A valid input subtree has no adjacent red nodes. Only the child
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
        };'''
for original, output in [('kernel.rs','kernel-directional.rs'),('kernel-unchecked-inline.rs','kernel-unchecked-directional-inline.rs')]:
    source=(HERE/original).read_text()
    assert source.count(old)==1
    source=source.replace(old,new)
    source=source.replace('fn balance(&mut self, root: Index)', 'fn balance<const LEFT: bool>(&mut self, root: Index)')
    before='''            self.nodes[root as usize].left = left;
        } else if key > node.key {
            let right = self.ins(key, node.right);
            self.nodes[root as usize].right = right;
        } else { return root; }
        self.balance(root)'''
    after='''            self.nodes[root as usize].left = left;
            self.balance::<true>(root)
        } else if key > node.key {
            let right = self.ins(key, node.right);
            self.nodes[root as usize].right = right;
            self.balance::<false>(root)
        } else { root }'''
    assert source.count(before)==1
    source=source.replace(before,after)
    (HERE/output).write_text(source)
    (HERE/('check-'+output.removeprefix('kernel-'))).write_text((HERE/'check.rs').read_text().replace('include!("kernel.rs")',f'include!("{output}")'))

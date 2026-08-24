#[derive(Clone, Copy)]
struct Node {
    color: u8,
    left: u32,
    value: i64,
    right: u32,
}

fn alloc(arena: &mut Vec<Node>, color: u8, left: u32, value: i64, right: u32) -> u32 {
    let idx = arena.len() as u32;
    arena.push(Node { color, left, value, right });
    idx
}

fn balance(arena: &mut Vec<Node>, c: u8, a: u32, x: i64, b: u32) -> u32 {
    if c == 1 { // black
        if a != 0 {
            let a_node = arena[a as usize];
            if a_node.color == 0 { // red
                if a_node.left != 0 {
                    let al = arena[a_node.left as usize];
                    if al.color == 0 {
                        let left = alloc(arena, 1, al.left, al.value, al.right);
                        let right = alloc(arena, 1, a_node.right, x, b);
                        return alloc(arena, 0, left, a_node.value, right);
                    }
                }
                if a_node.right != 0 {
                    let ar = arena[a_node.right as usize];
                    if ar.color == 0 {
                        let left = alloc(arena, 1, a_node.left, a_node.value, ar.left);
                        let right = alloc(arena, 1, ar.right, x, b);
                        return alloc(arena, 0, left, ar.value, right);
                    }
                }
            }
        }
        if b != 0 {
            let b_node = arena[b as usize];
            if b_node.color == 0 {
                if b_node.left != 0 {
                    let bl = arena[b_node.left as usize];
                    if bl.color == 0 {
                        let left = alloc(arena, 1, a, x, bl.left);
                        let right = alloc(arena, 1, bl.right, b_node.value, b_node.right);
                        return alloc(arena, 0, left, bl.value, right);
                    }
                }
                if b_node.right != 0 {
                    let br = arena[b_node.right as usize];
                    if br.color == 0 {
                        let left = alloc(arena, 1, a, x, b_node.left);
                        let right = alloc(arena, 1, br.left, br.value, br.right);
                        return alloc(arena, 0, left, b_node.value, right);
                    }
                }
            }
        }
    }
    alloc(arena, c, a, x, b)
}

fn ins(arena: &mut Vec<Node>, x: i64, t: u32) -> u32 {
    if t == 0 {
        return alloc(arena, 0, 0, x, 0);
    }
    let node = arena[t as usize];
    if x < node.value {
        let left = ins(arena, x, node.left);
        balance(arena, node.color, left, node.value, node.right)
    } else if x > node.value {
        let right = ins(arena, x, node.right);
        balance(arena, node.color, node.left, node.value, right)
    } else {
        t
    }
}

fn depth(arena: &Vec<Node>, t: u32) -> i64 {
    if t == 0 {
        return 0;
    }
    let node = &arena[t as usize];
    let l = depth(arena, node.left);
    let r = depth(arena, node.right);
    1 + std::cmp::max(l, r)
}

pub fn Test_RBTreeFFICheatcode_runRBTreeFFICheatcode(limit: i64) -> i64 {
    let mut arena = Vec::with_capacity(5_000_000);
    arena.push(Node { color: 1, left: 0, value: 0, right: 0 }); // index 0 is null

    let mut acc = 0;
    let mut i = limit;
    while i > 0 {
        let res = ins(&mut arena, i, acc);
        arena[res as usize].color = 1; // root is always black
        acc = res;
        i -= 1;
    }
    depth(&arena, acc)
}

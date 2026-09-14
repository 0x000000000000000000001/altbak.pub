// Control: same eligibility and final cell placement as rotation.rs, but three
// complete payload extractions/reconstructions instead of field permutations.
fn probe_rotate_ll(slot: &mut Tree) -> bool {
    let Tree::T(color, left, key, right) = slot else { return false; };
    if !matches!(color, Color::B) { return false; }
    {
        let Some(parent) = std::rc::Rc::get_mut(left) else { return false; };
        let Tree::T(parent_color, child, _, _) = parent else { return false; };
        if !matches!(parent_color, Color::R) { return false; }
        {
            let Some(grandchild) = std::rc::Rc::get_mut(child) else { return false; };
            let Tree::T(grandchild_color, _, _, _) = grandchild else { return false; };
            if !matches!(grandchild_color, Color::R) { return false; }
            let Tree::T(_, a, x, b) = grandchild.__purust_take().unwrap() else { unreachable!() };
            *grandchild = Tree::T(Color::B, a, x, b); // PROBE_REBUILD
        }
        let Tree::T(_, child, y, c) = parent.__purust_take().unwrap() else { unreachable!() };
        let d = std::mem::replace(right, child);
        *parent = Tree::T(Color::B, c, *key, d); // PROBE_REBUILD
        *key = y;
    }
    let Tree::T(_, parent, y, child) = slot.__purust_take().unwrap() else { unreachable!() };
    *slot = Tree::T(Color::R, child, y, parent); // PROBE_REBUILD
    // PROBE_SUCCESS
    true
}

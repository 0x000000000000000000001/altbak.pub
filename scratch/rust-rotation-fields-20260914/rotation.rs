// Isolated prototype, deliberately tied to the benchmark's LL constructor shape.
// A compiler rule would have to derive the transformation from the TAST.
fn probe_rotate_ll(slot: &mut Tree) -> bool {
    let Tree::T(color, left, key, right) = slot else { return false; };
    if !matches!(color, Color::B) { return false; }
    {
        let Some(parent) = std::rc::Rc::get_mut(left) else { return false; };
        let Tree::T(parent_color, child, parent_key, parent_right) = parent else { return false; };
        if !matches!(parent_color, Color::R) { return false; }
        {
            let Some(grandchild) = std::rc::Rc::get_mut(child) else { return false; };
            let Tree::T(grandchild_color, _, _, _) = grandchild else { return false; };
            if !matches!(grandchild_color, Color::R) { return false; }
            // All fallible guards precede the first mutation.
            *grandchild_color = Color::B;
        }
        *parent_color = Color::B;
        std::mem::swap(parent_key, key);
        std::mem::swap(child, parent_right);
        std::mem::swap(parent_right, right);
    }
    std::mem::swap(left, right);
    *color = Color::R;
    // PROBE_SUCCESS
    true
}

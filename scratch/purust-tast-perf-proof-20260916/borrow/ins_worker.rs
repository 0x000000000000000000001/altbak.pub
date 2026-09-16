// Experiment: retain the generated balance guards, field permutation and generic
// fallback. Only pass recursive child storage as a mutable slot. This needs a
// consumed-argument/result contract, not just a read-only borrow contract.
pub fn Test_RBTree_ins(mut key: i64, mut tree: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    Test_RBTree_ins__borrow(key, &mut tree);
    tree
}

fn Test_RBTree_ins__borrow(key: i64, tree: &mut std::rc::Rc<crate::Tree>) {
    if matches!(tree.as_ref(), crate::Tree::E) {
        // The empty singleton remains shared. This borrowed ABI costs one extra
        // temporary clone here compared with the generated consuming ABI.
        *tree = std::rc::Rc::new(crate::Tree::T(crate::Color::R, tree.clone(), key, tree.clone()));
        return;
    }

    let fields = if let Some(slot) = std::rc::Rc::get_mut(tree) {
        let crate::Tree::T(color, left, node_key, right) = slot else { unreachable!() };
        if key < *node_key {
            Test_RBTree_ins__borrow(key, left);
        } else if key > *node_key {
            Test_RBTree_ins__borrow(key, right);
        } else {
            return;
        }
        if !matches!(color, crate::Color::B)
            || Test_RBTree_balance__purust_child_rebuilds(color, left, node_key, right) {
            return;
        }
        if Test_RBTree_balance__purust_permute_fields(slot) {
            return;
        }
        let Some(crate::Tree::T(color, left, node_key, right)) = slot.__purust_take()
            else { unreachable!() };
        Some((color, left, node_key, right))
    } else {
        None
    };

    if let Some((color, left, node_key, right)) = fields {
        // Uncommon rotations use the exact generated reuse worker, with a safe
        // temporary reference in the caller slot while transferring its cell.
        let cell = std::mem::replace(tree, left.clone());
        *tree = Test_RBTree_balance__purust_reuse(color, left, node_key, right, cell);
    } else {
        // Shared or weakly referenced cells keep persistent semantics through
        // the exact original emitted function and all its generated fallbacks.
        *tree = Test_RBTree_ins__original(key, tree.clone());
    }
}

// Same four Okasaki cases and persistent Rc sharing, with borrowed projections.
fn rb_node(color: std::rc::Rc<crate::Color>, a: std::rc::Rc<crate::Tree>, x: i64, b: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    std::rc::Rc::new(crate::Tree::T(color, a, x, b))
}
pub fn Test_RBTree_balance(color: std::rc::Rc<crate::Color>, left: std::rc::Rc<crate::Tree>, key: i64, right: std::rc::Rc<crate::Tree>) -> std::rc::Rc<crate::Tree> {
    if matches!(color.as_ref(), Color::B) {
        if let Tree::T(lc, ll, lk, lr) = left.as_ref() {
            if matches!(lc.as_ref(), Color::R) {
                if let Tree::T(llc, a, x, b) = ll.as_ref() {
                    if matches!(llc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), a.clone(), *x, b.clone()), *lk,
                            rb_node(std::rc::Rc::new(crate::Color::B), lr.clone(), key, right));
                    }
                }
                if let Tree::T(lrc, b, y, c) = lr.as_ref() {
                    if matches!(lrc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), ll.clone(), *lk, b.clone()), *y,
                            rb_node(std::rc::Rc::new(crate::Color::B), c.clone(), key, right));
                    }
                }
            }
        }
        if let Tree::T(rc, rl, rk, rr) = right.as_ref() {
            if matches!(rc.as_ref(), Color::R) {
                if let Tree::T(rlc, b, y, c) = rl.as_ref() {
                    if matches!(rlc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), left, key, b.clone()), *y,
                            rb_node(std::rc::Rc::new(crate::Color::B), c.clone(), *rk, rr.clone()));
                    }
                }
                if let Tree::T(rrc, c, z, d) = rr.as_ref() {
                    if matches!(rrc.as_ref(), Color::R) {
                        return rb_node(std::rc::Rc::new(crate::Color::R),
                            rb_node(std::rc::Rc::new(crate::Color::B), left, key, rl.clone()), *rk,
                            rb_node(std::rc::Rc::new(crate::Color::B), c.clone(), *z, d.clone()));
                    }
                }
            }
        }
    }
    rb_node(color, left, key, right)
}

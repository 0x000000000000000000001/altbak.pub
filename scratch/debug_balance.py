import re

with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "r") as f:
    content = f.read()

target = """    fn make_t(c: crate::UnknownType, l: crate::UnknownType, x: i64, r: crate::UnknownType, reuse: Option<crate::UnknownType>) -> crate::UnknownType {
        if let Some(crate::Value::Record(mut ptr)) = reuse {
            let mut_ptr = perceus_ptr::PerceusPtr::make_mut(&mut ptr);
            if let Some(rc) = mut_ptr.vals.as_mut() {
                let vec = std::rc::Rc::make_mut(rc);
                vec[0] = c;
                vec[1] = l;
                vec[2] = crate::mk_int(x);
                vec[3] = r;
            }
            crate::Value::Record(ptr)
        } else {"""

replacement = """    fn make_t(c: crate::UnknownType, l: crate::UnknownType, x: i64, r: crate::UnknownType, reuse: Option<crate::UnknownType>) -> crate::UnknownType {
        if let Some(crate::Value::Record(mut ptr)) = reuse {
            if ptr.is_unique() {
                // println!("Reusing ptr!");
            }
            let mut_ptr = perceus_ptr::PerceusPtr::make_mut(&mut ptr);
            if let Some(rc) = mut_ptr.vals.as_mut() {
                if std::rc::Rc::strong_count(rc) == 1 {
                    // println!("Reusing vec!");
                } else {
                    // println!("Cloning vec!");
                }
                let vec = std::rc::Rc::make_mut(rc);
                vec[0] = c;
                vec[1] = l;
                vec[2] = crate::mk_int(x);
                vec[3] = r;
            }
            crate::Value::Record(ptr)
        } else {"""

if target in content:
    content = content.replace(target, replacement)
    with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "w") as f:
        f.write(content)
    print("Patched debug!")
else:
    print("Target not found!")

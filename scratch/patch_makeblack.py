import re

with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "r") as f:
    content = f.read()

target = """{
        let mut _reuse_ptr = match v { crate::Value::Record(ptr) => ptr, _ => unreachable!() };
        {
            let _mut = perceus_ptr::PerceusPtr::make_mut(&mut _reuse_ptr);
            // wait, Rc::get_mut only works if Rc is unique, but make_mut guarantees PerceusPtr is unique, but what about the Rc inside it?
            // The Rc inside was allocated uniquely, and FBIP clone clones the PerceusPtr, not the Rc inside.
            // Wait, if it was unique, Rc::get_mut works. If it was cloned, PerceusPtr clones, but we create a new Rc when we mutate?
            // Wait, PerceusPtr cloning clones the PerceusPtr, which holds the Record_a. Record_a holds Rc. So Rc is shared.
            // If make_mut clones the Record_a, Rc's refcount increases!
            // So Rc is NOT unique! We can't use Rc::get_mut!
            // Let's just allocate a new Vec!
            let mut new_vals = (*_mut.vals.as_ref().unwrap()).clone(); // clones the Rc
            let mut vec = std::rc::Rc::make_mut(&mut _mut.vals.as_mut().unwrap()); // This is safer!
            vec[0] = crate::Value::Record(perceus_ptr::PerceusPtr::new(Record_a { tag: "B", vals: None, ..Default::default() }));
        }
        crate::Value::Record(_reuse_ptr)
}"""

replacement = """{
        let mut _reuse_ptr = match v { crate::Value::Record(ptr) => ptr, _ => unreachable!() };
        {
            let _mut = perceus_ptr::PerceusPtr::make_mut(&mut _reuse_ptr);
            if let Some(rc) = _mut.vals.as_mut() {
                let vec = std::rc::Rc::make_mut(rc);
                vec[0] = crate::Value::Record(perceus_ptr::PerceusPtr::new(Record_a { tag: "B", vals: None, ..Default::default() }));
            }
        }
        crate::Value::Record(_reuse_ptr)
}"""

if target in content:
    content = content.replace(target, replacement)
    with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "w") as f:
        f.write(content)
    print("Patched makeBlack successfully!")
else:
    print("Target not found in makeBlack!")

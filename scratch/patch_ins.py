import re

with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "r") as f:
    content = f.read()

target = """    } else {
        let mut _reuse = v1;
        {
            let _mut = perceus_ptr::PerceusPtr::make_mut(_reuse.as_record_mut());
            // It's already a T with the same values, no need to update!
        }
        crate::Value::Record(_reuse)
    }"""

replacement = """    } else {
        let mut _reuse_ptr = match v1 { crate::Value::Record(ptr) => ptr, _ => unreachable!() };
        {
            let _mut = perceus_ptr::PerceusPtr::make_mut(&mut _reuse_ptr);
        }
        crate::Value::Record(_reuse_ptr)
    }"""

if target in content:
    content = content.replace(target, replacement)
    with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "w") as f:
        f.write(content)
    print("Patched ins successfully!")
else:
    print("Target not found!")

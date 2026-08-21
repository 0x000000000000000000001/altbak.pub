import re

with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "r") as f:
    content = f.read()

# We will replace the entire Test_RBTree_balance function.
target_start = "pub fn Test_RBTree_balance(mut v: crate::UnknownType, mut v1: crate::UnknownType, mut v2: i64, mut v3: crate::UnknownType) -> crate::UnknownType {"
target_end = "pub fn Test_RBTree_ins(mut v: i64, mut v1: crate::UnknownType) -> crate::UnknownType {"

start_idx = content.find(target_start)
end_idx = content.find(target_end)

new_balance = """pub fn Test_RBTree_balance(mut v: crate::UnknownType, mut v1: crate::UnknownType, mut v2: i64, mut v3: crate::UnknownType) -> crate::UnknownType {
    // Helper to get ptr and vals
    fn get_vals(t: &crate::UnknownType) -> Option<&Vec<crate::UnknownType>> {
        match t {
            crate::Value::Record(ptr) => {
                let rec = &**ptr;
                if rec.tag == "T" {
                    rec.vals.as_ref().map(|rc| &**rc)
                } else { None }
            },
            _ => None
        }
    }
    
    fn make_t(c: crate::UnknownType, l: crate::UnknownType, x: i64, r: crate::UnknownType, reuse: Option<crate::UnknownType>) -> crate::UnknownType {
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
        } else {
            crate::Value::Record(perceus_ptr::PerceusPtr::new(Record_a { 
                tag: "T", 
                vals: Some(std::rc::Rc::new(vec![c, l, crate::mk_int(x), r])), 
                ..Default::default() 
            }))
        }
    }
    
    fn make_r() -> crate::UnknownType {
        crate::Value::Record(perceus_ptr::PerceusPtr::new(Record_a { tag: "R", vals: None, ..Default::default() }))
    }
    fn make_b() -> crate::UnknownType {
        crate::Value::Record(perceus_ptr::PerceusPtr::new(Record_a { tag: "B", vals: None, ..Default::default() }))
    }

    let is_b = match &v { crate::Value::Record(ptr) => ptr.tag == "B", _ => false };
    
    if is_b {
        // match (T R (T R a x b) y c) z d
        if let Some(v1_vals) = get_vals(&v1) {
            if v1_vals[0].unwrap_record().tag == "R" {
                if let Some(v1_1_vals) = get_vals(&v1_vals[1]) {
                    if v1_1_vals[0].unwrap_record().tag == "R" {
                        // T R (T B a x b) y (T B c z d)
                        let a = v1_1_vals[1].clone();
                        let x = v1_1_vals[2].clone().unwrap_int();
                        let b = v1_1_vals[3].clone();
                        let y = v1_vals[2].clone().unwrap_int();
                        let c = v1_vals[3].clone();
                        let z = v2;
                        let d = v3;
                        let left = make_t(make_b(), a, x, b, Some(v1_vals[1].clone()));
                        let right = make_t(make_b(), c, z, d, Some(v1.clone()));
                        return make_t(make_r(), left, y, right, None);
                    }
                }
                // match (T R a x (T R b y c)) z d
                if let Some(v1_3_vals) = get_vals(&v1_vals[3]) {
                    if v1_3_vals[0].unwrap_record().tag == "R" {
                        // T R (T B a x b) y (T B c z d)
                        let a = v1_vals[1].clone();
                        let x = v1_vals[2].clone().unwrap_int();
                        let b = v1_3_vals[1].clone();
                        let y = v1_3_vals[2].clone().unwrap_int();
                        let c = v1_3_vals[3].clone();
                        let z = v2;
                        let d = v3;
                        let left = make_t(make_b(), a, x, b, Some(v1_vals[3].clone()));
                        let right = make_t(make_b(), c, z, d, Some(v1.clone()));
                        return make_t(make_r(), left, y, right, None);
                    }
                }
            }
        }
        
        // match a x (T R (T R b y c) z d)
        if let Some(v3_vals) = get_vals(&v3) {
            if v3_vals[0].unwrap_record().tag == "R" {
                if let Some(v3_1_vals) = get_vals(&v3_vals[1]) {
                    if v3_1_vals[0].unwrap_record().tag == "R" {
                        let a = v1;
                        let x = v2;
                        let b = v3_1_vals[1].clone();
                        let y = v3_1_vals[2].clone().unwrap_int();
                        let c = v3_1_vals[3].clone();
                        let z = v3_vals[2].clone().unwrap_int();
                        let d = v3_vals[3].clone();
                        let left = make_t(make_b(), a, x, b, Some(v3_vals[1].clone()));
                        let right = make_t(make_b(), c, z, d, Some(v3.clone()));
                        return make_t(make_r(), left, y, right, None);
                    }
                }
                
                // match a x (T R b y (T R c z d))
                if let Some(v3_3_vals) = get_vals(&v3_vals[3]) {
                    if v3_3_vals[0].unwrap_record().tag == "R" {
                        let a = v1;
                        let x = v2;
                        let b = v3_vals[1].clone();
                        let y = v3_vals[2].clone().unwrap_int();
                        let c = v3_3_vals[1].clone();
                        let z = v3_3_vals[2].clone().unwrap_int();
                        let d = v3_3_vals[3].clone();
                        let left = make_t(make_b(), a, x, b, Some(v3_vals[3].clone()));
                        let right = make_t(make_b(), c, z, d, Some(v3.clone()));
                        return make_t(make_r(), left, y, right, None);
                    }
                }
            }
        }
    }
    
    // Fallback
    make_t(v, v1, v2, v3, None)
}

"""

new_content = content[:start_idx] + new_balance + content[end_idx:]

with open("run/bak/rust/output/purust_output/Purs_Test_RBTree/src/lib.rs", "w") as f:
    f.write(new_content)

print("Replaced balance!")

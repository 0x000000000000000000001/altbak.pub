pub fn Data_Ord_ordIntImpl(mut lt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut eq: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut gt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut a: i64, mut b: i64) -> std::rc::Rc<Purs_Data_Ordering::Ordering> {
    match a.cmp(&b) {
        std::cmp::Ordering::Less => lt,
        std::cmp::Ordering::Equal => eq,
        std::cmp::Ordering::Greater => gt,
    }
}

pub fn Data_Ord_ordNumberImpl(mut lt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut eq: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut gt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut a: f64, mut b: f64) -> std::rc::Rc<Purs_Data_Ordering::Ordering> {
    if a < b { lt }
    else if a > b { gt }
    else { eq }
}

pub fn Data_Ord_ordCharImpl(mut lt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut eq: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut gt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut a: char, mut b: char) -> std::rc::Rc<Purs_Data_Ordering::Ordering> {
    match a.cmp(&b) {
        std::cmp::Ordering::Less => lt,
        std::cmp::Ordering::Equal => eq,
        std::cmp::Ordering::Greater => gt,
    }
}

pub fn Data_Ord_ordStringImpl(mut lt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut eq: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut gt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut a: String, mut b: String) -> std::rc::Rc<Purs_Data_Ordering::Ordering> {
    match a.cmp(&b) {
        std::cmp::Ordering::Less => lt,
        std::cmp::Ordering::Equal => eq,
        std::cmp::Ordering::Greater => gt,
    }
}

pub fn Data_Ord_ordBooleanImpl(mut lt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut eq: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut gt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut a: bool, mut b: bool) -> std::rc::Rc<Purs_Data_Ordering::Ordering> {
    match a.cmp(&b) {
        std::cmp::Ordering::Less => lt,
        std::cmp::Ordering::Equal => eq,
        std::cmp::Ordering::Greater => gt,
    }
}

pub fn Data_Ord_ordArrayImpl(mut f: purust_core::Func2<crate::UnknownType, crate::UnknownType, i64>, mut a0: crate::UnknownType, mut a1: crate::UnknownType) -> i64 {
    let arr1 = a0.unwrap_array();
    let arr2 = a1.unwrap_array();
    let len = std::cmp::min(arr1.len(), arr2.len());
    for i in 0..len {
        let cmp = f(arr1[i].clone(), arr2[i].clone());
        if cmp != 0 {
            return cmp;
        }
    }
    match arr1.len().cmp(&arr2.len()) {
        std::cmp::Ordering::Less => -1,
        std::cmp::Ordering::Equal => 0,
        std::cmp::Ordering::Greater => 1,
    }
}

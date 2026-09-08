// FFI implementation for Data.String.Common in Rust

pub fn Data_String_Common__localeCompare(mut lt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut eq: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut gt: std::rc::Rc<Purs_Data_Ordering::Ordering>, mut a: String, mut b: String) -> std::rc::Rc<Purs_Data_Ordering::Ordering> {
    match a.cmp(&b) {
        std::cmp::Ordering::Less => lt,
        std::cmp::Ordering::Equal => eq,
        std::cmp::Ordering::Greater => gt,
    }
}

pub fn Data_String_Common_replace(mut pat: String, mut rep: String, mut s: String) -> String {
    s.replacen(&pat, &rep, 1)
}

pub fn Data_String_Common_replaceAll(mut pat: String, mut rep: String, mut s: String) -> String {
    s.replace(&pat, &rep)
}

pub fn Data_String_Common_split(mut pat: String, mut s: String) -> std::rc::Rc<Vec<crate::UnknownType>> {
    let parts: Vec<crate::UnknownType> = if pat.is_empty() {
        if s.is_empty() {
            vec![]
        } else {
            s.chars().map(|c| crate::Value::String(c.to_string())).collect()
        }
    } else {
        s.split(&pat).map(|part| crate::Value::String(part.to_string())).collect()
    };
    std::rc::Rc::new(parts)
}

pub fn Data_String_Common_toLower(mut s: String) -> String {
    s.to_lowercase()
}

pub fn Data_String_Common_toUpper(mut s: String) -> String {
    s.to_uppercase()
}

pub fn Data_String_Common_trim(mut s: String) -> String {
    s.trim().to_string()
}

pub fn Data_String_Common_joinWith(mut sep: String, mut arr: std::rc::Rc<Vec<crate::UnknownType>>) -> String {
    let strings: Vec<String> = arr.iter().map(|item| item.unwrap_string()).collect();
    strings.join(&sep)
}

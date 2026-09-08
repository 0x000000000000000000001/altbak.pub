pub fn Effect_Console_log(mut a0: crate::UnknownType) -> crate::UnknownType {
    println!("{}", a0.unwrap_string());
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_error(mut a0: crate::UnknownType) -> crate::UnknownType {
    eprintln!("{}", a0.unwrap_string());
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_warn(mut a0: crate::UnknownType) -> crate::UnknownType {
    eprintln!("WARN: {}", a0.unwrap_string());
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_info(mut a0: crate::UnknownType) -> crate::UnknownType {
    println!("INFO: {}", a0.unwrap_string());
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_debug(mut a0: crate::UnknownType) -> crate::UnknownType {
    println!("DEBUG: {}", a0.unwrap_string());
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_clear() -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_group(mut a0: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_groupCollapsed(mut a0: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_groupEnd() -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_time(mut a0: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_timeEnd(mut a0: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

pub fn Effect_Console_timeLog(mut a0: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

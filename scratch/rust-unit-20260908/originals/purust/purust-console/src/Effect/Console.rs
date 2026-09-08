use std::rc::Rc;
use crate::{UnknownType, Record_a};

pub fn Effect_Console_log(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            println!("{}", a0);
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_error(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            eprintln!("{}", a0);
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_warn(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            eprintln!("WARN: {}", a0);
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_info(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            println!("INFO: {}", a0);
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_debug(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            println!("DEBUG: {}", a0);
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_clear() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_group(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_groupCollapsed(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_groupEnd() -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_time(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_timeEnd(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

pub fn Effect_Console_timeLog(mut a0: String) -> UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(Rc::new(move |mut _u: UnknownType| -> UnknownType {
            crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
        })))
}

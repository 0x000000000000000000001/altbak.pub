fn purust_effect_unit() -> crate::UnknownType {
    crate::Value::Record_a(perceus_ptr::PerceusPtr::new(crate::Record_a { ..Default::default() }))
}

fn purust_effect_run(action: &crate::UnknownType) -> crate::UnknownType {
    action.unwrap_func1()(purust_effect_unit())
}

pub fn Effect_pureE(value: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| value.clone())))
}

pub fn Effect_bindE(action: crate::UnknownType, next: purust_core::Func1<crate::UnknownType, crate::UnknownType>) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        let value = purust_effect_run(&action);
        purust_effect_run(&next(value))
    })))
}

pub fn Effect_untilE(condition: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        while !purust_effect_run(&condition).unwrap_bool() {}
        purust_effect_unit()
    })))
}

pub fn Effect_whileE(condition: crate::UnknownType, body: crate::UnknownType) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        while purust_effect_run(&condition).unwrap_bool() {
            purust_effect_run(&body);
        }
        purust_effect_unit()
    })))
}

pub fn Effect_forE(lo: i64, hi: i64, body: purust_core::Func1<i64, crate::UnknownType>) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        for index in lo..hi {
            purust_effect_run(&body(index));
        }
        purust_effect_unit()
    })))
}

pub fn Effect_foreachE(values: crate::UnknownType, body: purust_core::Func1<crate::UnknownType, crate::UnknownType>) -> crate::UnknownType {
    crate::Value::Func1(purust_core::Func1::Shared(std::rc::Rc::new(move |_| {
        for value in values.unwrap_array().iter() {
            purust_effect_run(&body(value.clone()));
        }
        purust_effect_unit()
    })))
}

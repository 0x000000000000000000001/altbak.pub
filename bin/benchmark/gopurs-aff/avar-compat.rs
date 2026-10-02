// Adapt constructor arguments from gopurs-avar to the Purust runtime's utility
// record. Queueing, cancellation and synchronized callback delivery are shared.
fn gopurs_avar_field(util: &Value, name: &str) -> Value {
    let Value::Record_a(record) = util.resolve() else { panic!("Expected AVar utility record") };
    record.get_field(name).expect("AVar utility field").clone()
}

fn gopurs_avar_util(overrides: Vec<(&str, Value)>) -> Value {
    let mut record = purust_core::Record_a::default();
    record.set_field("left", Value::Func1(purust_core::Func1::Static(|value|
        Value::Class(Rc::new(Rc::new(Purs_Data_Either::Either::Left(value)))))));
    record.set_field("right", Value::Func1(purust_core::Func1::Static(|value|
        Value::Class(Rc::new(Rc::new(Purs_Data_Either::Either::Right(value)))))));
    for (name, value) in overrides { record.set_field(name, value); }
    Value::Record_a(perceus_ptr::PerceusPtr::new(record))
}

pub fn Effect_AVar__putVar() -> UnknownType {
    Value::Func5(purust_core::Func5::Static(|left, right, value, avar, callback| {
        gopurs_base_AVar__putVar().unwrap_func4()(gopurs_avar_util(vec![("left", left), ("right", right)]), value, avar, callback)
    }))
}

pub fn Effect_AVar__takeVar() -> UnknownType {
    Value::Func4(purust_core::Func4::Static(|left, right, avar, callback| {
        gopurs_base_AVar__takeVar().unwrap_func3()(gopurs_avar_util(vec![("left", left), ("right", right)]), avar, callback)
    }))
}

pub fn Effect_AVar__readVar() -> UnknownType {
    Value::Func4(purust_core::Func4::Static(|left, right, avar, callback| {
        gopurs_base_AVar__readVar().unwrap_func3()(gopurs_avar_util(vec![("left", left), ("right", right)]), avar, callback)
    }))
}

pub fn Effect_AVar__killVar() -> UnknownType {
    Value::Func2(purust_core::Func2::Static(|error, avar| {
        gopurs_base_AVar__killVar().unwrap_func3()(gopurs_avar_util(vec![]), error, avar)
    }))
}

pub fn Effect_AVar__tryPutVar() -> UnknownType {
    Value::Func2(purust_core::Func2::Static(|value, avar| {
        gopurs_base_AVar__tryPutVar().unwrap_func3()(gopurs_avar_util(vec![]), value, avar)
    }))
}

pub fn Effect_AVar__tryTakeVar() -> UnknownType {
    Value::Func5(purust_core::Func5::Static(|left, right, nothing, just, avar| {
        let util = gopurs_avar_util(vec![("left", left), ("right", right), ("nothing", nothing), ("just", just)]);
        gopurs_base_AVar__tryTakeVar().unwrap_func2()(util, avar)
    }))
}

pub fn Effect_AVar__tryReadVar() -> UnknownType {
    Value::Func3(purust_core::Func3::Static(|nothing, just, avar| {
        gopurs_base_AVar__tryReadVar().unwrap_func2()(gopurs_avar_util(vec![("nothing", nothing), ("just", just)]), avar)
    }))
}

pub fn Effect_AVar__status() -> UnknownType {
    Value::Func4(purust_core::Func4::Static(|killed, filled, empty, avar| {
        gopurs_base_AVar__status().unwrap_func2()(gopurs_avar_util(vec![("killed", killed), ("filled", filled), ("empty", empty)]), avar)
    }))
}

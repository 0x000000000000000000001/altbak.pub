use std::rc::Rc;

#[derive(Clone)]
pub enum Value {
    Int(i64),
    Func(Rc<dyn Fn(Value) -> Value>),
}

impl Value {
    pub fn unwrap_int(&self) -> i64 {
        if let Value::Int(v) = self { *v } else { panic!("Expected Int"); }
    }
    pub fn unwrap_func(&self) -> Rc<dyn Fn(Value) -> Value> {
        if let Value::Func(v) = self { v.clone() } else { panic!("Expected Func"); }
    }
}

pub fn mk_int(val: i64) -> Value { Value::Int(val) }

fn main() {
    // We have a typed closure
    let typed_closure: Rc<dyn Fn(i64) -> i64> = Rc::new(|x| x + 1);
    
    // We want to box it into Value
    let typed_closure_clone = typed_closure.clone();
    let boxed: Value = Value::Func(Rc::new(move |arg: Value| -> Value {
        mk_int(typed_closure_clone(arg.unwrap_int()))
    }));
    
    // We have a boxed closure, we want to unbox it into a typed closure
    let boxed_func = boxed.unwrap_func();
    let unboxed: Rc<dyn Fn(i64) -> i64> = Rc::new(move |arg: i64| -> i64 {
        boxed_func(mk_int(arg)).unwrap_int()
    });
    
    println!("{}", unboxed(42));
}

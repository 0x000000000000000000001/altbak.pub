#![allow(warnings)]

use perceus_ptr::PerceusPtr;

#[derive(Clone)]
pub enum Void {}

#[derive(Clone)]
pub enum Value {
    Unit,
    Null,
    Int(i64),
    Number(f64),
    Bool(bool),
    String(String),
    Char(char),
    Array(std::rc::Rc<Vec<UnknownType>>),
    Func1(Func1<UnknownType, UnknownType>),
    Func2(Func2<UnknownType, UnknownType, UnknownType>),
    Func3(Func3<UnknownType, UnknownType, UnknownType, UnknownType>),
    Func4(Func4<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func5(Func5<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func6(Func6<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func7(Func7<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func8(Func8<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func9(Func9<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func10(Func10<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func11(Func11<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Func12(Func12<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType>),
    Class(std::rc::Rc<dyn std::any::Any>),
    Thunk(perceus_ptr::PerceusPtr<Thunk>),
    Record_a(perceus_ptr::PerceusPtr<Record_a>),
    DynamicRecord(perceus_ptr::PerceusPtr<RecordFields>),
    Record_a_b_keep_z(perceus_ptr::PerceusPtr<Record_a_b_keep_z>),
    Record_c_d(perceus_ptr::PerceusPtr<Record_c_d>),
    Record_e_f(perceus_ptr::PerceusPtr<Record_e_f>),
}

impl Value {
    pub fn resolve(&self) -> &Self {
        let mut value = self;
        while let Value::Thunk(thunk) = value {
            value = thunk.value.get().expect("recursive value used before initialization");
        }
        value
    }
    pub fn unwrap_unit(&self) {
        if !matches!(self.resolve(), Value::Unit) { panic!("Expected Unit"); }
    }
    pub fn unwrap_int(&self) -> i64 {
        if let Value::Int(v) = self.resolve() { *v } else { panic!("Expected Int"); }
    }
    pub fn unwrap_number(&self) -> f64 {
        // Foreign numbers can originate from a native PureScript Int.
        match self.resolve() { Value::Number(v) => *v, Value::Int(v) => *v as f64, _ => panic!("Expected Number") }
    }
    pub fn unwrap_bool(&self) -> bool {
        if let Value::Bool(v) = self.resolve() { *v } else { panic!("Expected Bool"); }
    }
    pub fn unwrap_string(&self) -> String {
        if let Value::String(v) = self.resolve() { v.clone() } else { panic!("Expected String"); }
    }
    pub fn unwrap_char(&self) -> char {
        if let Value::Char(v) = self.resolve() { *v } else { panic!("Expected Char"); }
    }
    pub fn unwrap_array(&self) -> std::rc::Rc<Vec<UnknownType>> {
        if let Value::Array(v) = self.resolve() { v.clone() } else { panic!("Expected Array"); }
    }
    pub fn unwrap_func1(&self) -> Func1<UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func1(v) = value { v.clone() } else if let Value::Record_a(v) = value { v.call.clone().unwrap() } else if let Value::Func2(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func1(Func1::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType| -> UnknownType { f2(a0.clone(), a1) } }))) })) } else if let Value::Func3(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func2(Func2::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2) } }))) })) } else if let Value::Func4(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func3(Func3::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3) } }))) })) } else if let Value::Func5(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func4(Func4::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4) } }))) })) } else if let Value::Func6(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func5(Func5::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5) } }))) })) } else if let Value::Func7(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func6(Func6::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6) } }))) })) } else if let Value::Func8(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func7(Func7::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7) } }))) })) } else if let Value::Func9(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func8(Func8::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8) } }))) })) } else if let Value::Func10(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func9(Func9::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8, a9) } }))) })) } else if let Value::Func11(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func10(Func10::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8, a9, a10) } }))) })) } else if let Value::Func12(v) = value { let f = v.clone(); Func1::Shared(std::rc::Rc::new(move |a0: UnknownType| -> UnknownType { crate::Value::Func11(Func11::Shared(std::rc::Rc::new({ let f2 = f.clone(); move |mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType, mut a11: UnknownType| -> UnknownType { f2(a0.clone(), a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11) } }))) })) } else { panic!("Expected Func1"); }
    }
    pub fn unwrap_func2(&self) -> Func2<UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func2(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func2::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1) })) } else { panic!("Expected Func2 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func3(&self) -> Func3<UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func3(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func3::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2) })) } else { panic!("Expected Func3 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func4(&self) -> Func4<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func4(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func4::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3) })) } else { panic!("Expected Func4 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func5(&self) -> Func5<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func5(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func5::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4) })) } else { panic!("Expected Func5 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func6(&self) -> Func6<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func6(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func6::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5) })) } else { panic!("Expected Func6 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func7(&self) -> Func7<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func7(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func7::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6) })) } else { panic!("Expected Func7 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func8(&self) -> Func8<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func8(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func8::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7) })) } else { panic!("Expected Func8 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func9(&self) -> Func9<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func9(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func9::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8) })) } else { panic!("Expected Func9 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func10(&self) -> Func10<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func10(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func10::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8).unwrap_func1()(a9) })) } else { panic!("Expected Func10 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func11(&self) -> Func11<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func11(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func11::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8).unwrap_func1()(a9).unwrap_func1()(a10) })) } else { panic!("Expected Func11 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_func12(&self) -> Func12<UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType, UnknownType> {
        let value = self.resolve();
        if let Value::Func12(v) = value { v.clone() } else if let Value::Func1(v) = value { let f = v.clone(); Func12::Shared(std::rc::Rc::new(move |mut a0: UnknownType, mut a1: UnknownType, mut a2: UnknownType, mut a3: UnknownType, mut a4: UnknownType, mut a5: UnknownType, mut a6: UnknownType, mut a7: UnknownType, mut a8: UnknownType, mut a9: UnknownType, mut a10: UnknownType, mut a11: UnknownType| -> UnknownType { f(a0).unwrap_func1()(a1).unwrap_func1()(a2).unwrap_func1()(a3).unwrap_func1()(a4).unwrap_func1()(a5).unwrap_func1()(a6).unwrap_func1()(a7).unwrap_func1()(a8).unwrap_func1()(a9).unwrap_func1()(a10).unwrap_func1()(a11) })) } else { panic!("Expected Func12 or Func1 (curried) - got something else"); }
    }
    pub fn unwrap_class<T: 'static>(&self) -> &T {
        if let Value::Class(v) = self.resolve() { v.downcast_ref::<T>().unwrap() } else { panic!("Expected Class"); }
    }
    pub fn drop_explicit(self) {
    }
    pub fn __purust_ctor_tag(&self) -> &'static str {
        if let Value::Record_a(r) = self.resolve() { r.tag } else { panic!("Expected Record_a for tag"); }
    }
    pub fn get_a(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.a.clone().unwrap(),
            Value::Record_a(r) => r.a.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("a").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field a"),
        }
    }
    pub fn get_b(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.b.clone().unwrap(),
            Value::Record_a(r) => r.b.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("b").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field b"),
        }
    }
    pub fn get_c(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_c_d(r) => r.c.clone().unwrap(),
            Value::Record_a(r) => r.c.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("c").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field c"),
        }
    }
    pub fn get_d(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_c_d(r) => r.d.clone().unwrap(),
            Value::Record_a(r) => r.d.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("d").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field d"),
        }
    }
    pub fn get_e(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_e_f(r) => r.e.clone().unwrap(),
            Value::Record_a(r) => r.e.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("e").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field e"),
        }
    }
    pub fn get_f(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_e_f(r) => r.f.clone().unwrap(),
            Value::Record_a(r) => r.f.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("f").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field f"),
        }
    }
    pub fn get_keep(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.keep.clone().unwrap(),
            Value::Record_a(r) => r.keep.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("keep").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field keep"),
        }
    }
    pub fn get_z(&self) -> UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.z.clone().unwrap(),
            Value::Record_a(r) => r.z.clone().unwrap(),
            Value::DynamicRecord(r) => r.get("z").cloned().expect("Missing record field"),
            _ => panic!("Expected record with field z"),
        }
    }
    pub fn __purust_get_field(&self, name: &str) -> Option<Value> {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => match name {
                "a" => r.a.clone(),
                "b" => r.b.clone(),
                "keep" => r.keep.clone(),
                "z" => r.z.clone(),
                _ => None,
            },
            Value::Record_c_d(r) => match name {
                "c" => r.c.clone(),
                "d" => r.d.clone(),
                _ => None,
            },
            Value::Record_e_f(r) => match name {
                "e" => r.e.clone(),
                "f" => r.f.clone(),
                _ => None,
            },
            Value::Record_a(r) => match name {
                "a" => r.a.clone(),
                "b" => r.b.clone(),
                "c" => r.c.clone(),
                "d" => r.d.clone(),
                "e" => r.e.clone(),
                "f" => r.f.clone(),
                "keep" => r.keep.clone(),
                "z" => r.z.clone(),
                _ => None,
            },
            Value::DynamicRecord(r) => r.get(name).cloned(),
            _ => panic!("Expected record"),
        }
    }
    pub fn __purust_set_field(mut self, name: &str, value: Value) -> Value {
        if matches!(self, Value::Thunk(_)) { self = self.resolve().clone(); }
        match &mut self {
            Value::Record_a_b_keep_z(r) => match name {
                "a" => { perceus_ptr::PerceusPtr::make_mut(r).a = Some(value); return self; },
                "b" => { perceus_ptr::PerceusPtr::make_mut(r).b = Some(value); return self; },
                "keep" => { perceus_ptr::PerceusPtr::make_mut(r).keep = Some(value); return self; },
                "z" => { perceus_ptr::PerceusPtr::make_mut(r).z = Some(value); return self; },
                _ => {},
            },
            Value::Record_c_d(r) => match name {
                "c" => { perceus_ptr::PerceusPtr::make_mut(r).c = Some(value); return self; },
                "d" => { perceus_ptr::PerceusPtr::make_mut(r).d = Some(value); return self; },
                _ => {},
            },
            Value::Record_e_f(r) => match name {
                "e" => { perceus_ptr::PerceusPtr::make_mut(r).e = Some(value); return self; },
                "f" => { perceus_ptr::PerceusPtr::make_mut(r).f = Some(value); return self; },
                _ => {},
            },
            Value::Record_a(r) => match name {
                "a" => { perceus_ptr::PerceusPtr::make_mut(r).a = Some(value); return self; },
                "b" => { perceus_ptr::PerceusPtr::make_mut(r).b = Some(value); return self; },
                "c" => { perceus_ptr::PerceusPtr::make_mut(r).c = Some(value); return self; },
                "d" => { perceus_ptr::PerceusPtr::make_mut(r).d = Some(value); return self; },
                "e" => { perceus_ptr::PerceusPtr::make_mut(r).e = Some(value); return self; },
                "f" => { perceus_ptr::PerceusPtr::make_mut(r).f = Some(value); return self; },
                "keep" => { perceus_ptr::PerceusPtr::make_mut(r).keep = Some(value); return self; },
                "z" => { perceus_ptr::PerceusPtr::make_mut(r).z = Some(value); return self; },
                _ => {},
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert(name.to_owned(), value); return self; },
            _ => panic!("Expected record"),
        }
        let mut fields = RecordFields::new();
        match &self {
            Value::Record_a_b_keep_z(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.keep { fields.insert("keep".to_owned(), value.clone()); }
                if let Some(value) = &r.z { fields.insert("z".to_owned(), value.clone()); }
            },
            Value::Record_c_d(r) => {
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
            },
            Value::Record_e_f(r) => {
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
            },
            Value::Record_a(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
                if let Some(value) = &r.keep { fields.insert("keep".to_owned(), value.clone()); }
                if let Some(value) = &r.z { fields.insert("z".to_owned(), value.clone()); }
            },
            _ => unreachable!(),
        }
        fields.insert(name.to_owned(), value);
        Value::DynamicRecord(perceus_ptr::PerceusPtr::new(fields))
    }
    pub fn __purust_record_fields(&self) -> Option<RecordFields> {
        let mut fields = RecordFields::new();
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.keep { fields.insert("keep".to_owned(), value.clone()); }
                if let Some(value) = &r.z { fields.insert("z".to_owned(), value.clone()); }
            },
            Value::Record_c_d(r) => {
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
            },
            Value::Record_e_f(r) => {
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
            },
            Value::Record_a(r) => {
                if let Some(value) = &r.a { fields.insert("a".to_owned(), value.clone()); }
                if let Some(value) = &r.b { fields.insert("b".to_owned(), value.clone()); }
                if let Some(value) = &r.c { fields.insert("c".to_owned(), value.clone()); }
                if let Some(value) = &r.d { fields.insert("d".to_owned(), value.clone()); }
                if let Some(value) = &r.e { fields.insert("e".to_owned(), value.clone()); }
                if let Some(value) = &r.f { fields.insert("f".to_owned(), value.clone()); }
                if let Some(value) = &r.keep { fields.insert("keep".to_owned(), value.clone()); }
                if let Some(value) = &r.z { fields.insert("z".to_owned(), value.clone()); }
            },
            Value::DynamicRecord(r) => return Some((**r).clone()),
            _ => return None,
        }
        Some(fields)
    }
    pub fn __purust_borrow_a(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.a.as_ref().unwrap(),
            Value::Record_a(r) => r.a.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("a").expect("Missing record field"),
            _ => panic!("Expected record with field a"),
        }
    }
    pub fn __purust_borrow_b(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.b.as_ref().unwrap(),
            Value::Record_a(r) => r.b.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("b").expect("Missing record field"),
            _ => panic!("Expected record with field b"),
        }
    }
    pub fn __purust_borrow_c(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_c_d(r) => r.c.as_ref().unwrap(),
            Value::Record_a(r) => r.c.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("c").expect("Missing record field"),
            _ => panic!("Expected record with field c"),
        }
    }
    pub fn __purust_borrow_d(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_c_d(r) => r.d.as_ref().unwrap(),
            Value::Record_a(r) => r.d.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("d").expect("Missing record field"),
            _ => panic!("Expected record with field d"),
        }
    }
    pub fn __purust_borrow_e(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_e_f(r) => r.e.as_ref().unwrap(),
            Value::Record_a(r) => r.e.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("e").expect("Missing record field"),
            _ => panic!("Expected record with field e"),
        }
    }
    pub fn __purust_borrow_f(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_e_f(r) => r.f.as_ref().unwrap(),
            Value::Record_a(r) => r.f.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("f").expect("Missing record field"),
            _ => panic!("Expected record with field f"),
        }
    }
    pub fn __purust_borrow_keep(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.keep.as_ref().unwrap(),
            Value::Record_a(r) => r.keep.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("keep").expect("Missing record field"),
            _ => panic!("Expected record with field keep"),
        }
    }
    pub fn __purust_borrow_z(&self) -> &UnknownType {
        match self.resolve() {
            Value::Record_a_b_keep_z(r) => r.z.as_ref().unwrap(),
            Value::Record_a(r) => r.z.as_ref().unwrap(),
            Value::DynamicRecord(r) => r.get("z").expect("Missing record field"),
            _ => panic!("Expected record with field z"),
        }
    }
    pub fn set_a(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b_keep_z(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.a = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.a = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("a".to_owned(), val); },
            _ => panic!("Expected record with field a"),
        }
    }
    pub fn set_b(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b_keep_z(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.b = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.b = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("b".to_owned(), val); },
            _ => panic!("Expected record with field b"),
        }
    }
    pub fn set_c(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_c_d(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.c = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.c = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("c".to_owned(), val); },
            _ => panic!("Expected record with field c"),
        }
    }
    pub fn set_d(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_c_d(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.d = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.d = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("d".to_owned(), val); },
            _ => panic!("Expected record with field d"),
        }
    }
    pub fn set_e(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_e_f(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.e = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.e = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("e".to_owned(), val); },
            _ => panic!("Expected record with field e"),
        }
    }
    pub fn set_f(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_e_f(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.f = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.f = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("f".to_owned(), val); },
            _ => panic!("Expected record with field f"),
        }
    }
    pub fn set_keep(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b_keep_z(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.keep = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.keep = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("keep".to_owned(), val); },
            _ => panic!("Expected record with field keep"),
        }
    }
    pub fn set_z(&mut self, val: UnknownType) {
        if matches!(self, Value::Thunk(_)) { *self = self.resolve().clone(); }
        match self {
            Value::Record_a_b_keep_z(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.z = Some(val);
            },
            Value::Record_a(r) => {
                let mut mut_r = perceus_ptr::PerceusPtr::make_mut(r);
                mut_r.z = Some(val);
            },
            Value::DynamicRecord(r) => { perceus_ptr::PerceusPtr::make_mut(r).insert("z".to_owned(), val); },
            _ => panic!("Expected record with field z"),
        }
    }
}

pub type UnknownType = Value;


// Internal strings store one Rust scalar per UTF-16 code unit, in code-unit
// order. Scalars >= D800 are shifted past Rust's surrogate hole.
#[inline]
pub fn purust_char_from_code_unit(unit: u16) -> char {
    let value = unit as u32;
    char::from_u32(if value < 0xd800 { value } else { value + 0x800 }).unwrap()
}

#[inline]
pub fn purust_char_to_code_unit(value: char) -> u16 {
    let value = value as u32;
    assert!(value <= 0x107ff, "Expected an encoded UTF-16 code unit");
    (if value < 0xd800 { value } else { value - 0x800 }) as u16
}

pub fn purust_string_from_utf16(units: &[u16]) -> String {
    units.iter().copied().map(purust_char_from_code_unit).collect()
}

pub fn purust_string_to_utf16(value: &str) -> Vec<u16> {
    value.chars().map(purust_char_to_code_unit).collect()
}

pub fn purust_string_from_utf8(value: &str) -> String {
    if value.is_ascii() { return value.to_owned(); }
    value.encode_utf16().map(purust_char_from_code_unit).collect()
}

pub fn purust_string_to_utf8_lossy(value: &str) -> std::string::String {
    if value.is_ascii() { return value.to_owned(); }
    std::string::String::from_utf16_lossy(&purust_string_to_utf16(value))
}

pub mod module_values {
    use std::collections::HashMap;
    use std::panic::{catch_unwind, resume_unwind, AssertUnwindSafe};
    use std::sync::{Condvar, Mutex, OnceLock};
    use std::thread::{self, ThreadId};

    enum State { Empty, Running(ThreadId), Ready, Poisoned }
    struct Wait { owner: ThreadId, cell: usize }
    static WAITS: OnceLock<Mutex<HashMap<ThreadId, Wait>>> = OnceLock::new();

    fn waits() -> &'static Mutex<HashMap<ThreadId, Wait>> {
        WAITS.get_or_init(|| Mutex::new(HashMap::new()))
    }

    struct Waiting { thread: ThreadId }
    impl Waiting {
        // Called with this cell's state locked, so its owner cannot complete
        // between registering the edge and entering the condition-variable wait.
        fn enter(thread: ThreadId, owner: ThreadId, cell: usize) -> Option<Self> {
            let mut graph = waits().lock().unwrap();
            let mut next = owner;
            loop {
                if next == thread { return None; }
                match graph.get(&next) {
                    Some(wait) => next = wait.owner,
                    None => break,
                }
            }
            graph.insert(thread, Wait { owner, cell });
            Some(Self { thread })
        }
    }
    impl Drop for Waiting {
        fn drop(&mut self) { waits().lock().unwrap().remove(&self.thread); }
    }

    pub struct Cell<T> {
        value: OnceLock<T>,
        state: Mutex<State>,
        ready: Condvar,
    }
    impl<T> Cell<T> {
        pub const fn new() -> Self {
            Self { value: OnceLock::new(), state: Mutex::new(State::Empty), ready: Condvar::new() }
        }

        pub fn get_or_init(&self, name: &'static str, init: impl FnOnce() -> T) -> &T {
            if let Some(value) = self.value.get() { return value; }
            let current = thread::current().id();
            let key = self as *const Self as usize;
            let mut state = self.state.lock().unwrap();
            loop {
                match *state {
                    State::Ready => return self.value.get().unwrap(),
                    State::Poisoned => {
                        drop(state);
                        panic!("module value initialization previously failed: {}", name);
                    }
                    State::Running(owner) => {
                        let waiting = match Waiting::enter(current, owner, key) {
                            Some(waiting) => waiting,
                            None => {
                                drop(state);
                                panic!("cyclic module value initialization: {}", name);
                            }
                        };
                        state = self.ready.wait(state).unwrap();
                        drop(waiting);
                    }
                    State::Empty => { *state = State::Running(current); break; }
                }
            }
            drop(state);
            // No state/graph mutex is held while executing user code, cloning
            // its result, or unwinding its destructors. Failure is permanent:
            // callers never silently replay a partially executed initializer.
            let outcome = catch_unwind(AssertUnwindSafe(init));
            let failure = match outcome {
                Ok(value) => {
                    assert!(self.value.set(value).is_ok(), "module value initialized twice");
                    None
                }
                Err(error) => Some(error),
            };
            let mut state = self.state.lock().unwrap();
            *state = if failure.is_some() { State::Poisoned } else { State::Ready };
            // Remove completed dependencies before this owner can start another
            // initialization. Otherwise stale wait edges could report a cycle.
            waits().lock().unwrap().retain(|_, wait| wait.cell != key);
            self.ready.notify_all();
            drop(state);
            if let Some(error) = failure { resume_unwind(error); }
            self.value.get().unwrap()
        }
    }
}

// Own properties retain insertion order. Enumeration puts array indices first,
// as JS Object.keys does; replacing a value never moves its property.
#[derive(Clone, Default)]
pub struct RecordFields(Vec<(String, Value)>);

impl RecordFields {
    pub fn new() -> Self { Self::default() }
    pub fn get(&self, name: &str) -> Option<&Value> {
        self.0.iter().find(|(key, _)| key == name).map(|(_, value)| value)
    }
    pub fn insert(&mut self, name: String, value: Value) -> Option<Value> {
        if let Some((_, old)) = self.0.iter_mut().find(|(key, _)| key == &name) {
            return Some(std::mem::replace(old, value));
        }
        self.0.push((name, value));
        None
    }
    pub fn remove(&mut self, name: &str) -> Option<Value> {
        self.0.iter().position(|(key, _)| key == name).map(|i| self.0.remove(i).1)
    }
    pub fn entries(&self) -> Vec<(String, Value)> {
        let mut entries = self.0.clone();
        entries.sort_by_key(|(key, _)| {
            key.parse::<u32>().ok().filter(|n| *n != u32::MAX && n.to_string() == *key)
                .map(|n| (0, n)).unwrap_or((1, 0))
        });
        entries
    }
}

// Shared mutable own-property storage used by native object FFI. Keeping the
// carrier here lets Foreign readers inspect it without a library dependency cycle.
pub struct SharedRecord(std::sync::Mutex<RecordFields>);
impl SharedRecord {
    pub fn empty() -> Self { Self(std::sync::Mutex::new(RecordFields::new())) }
    pub fn from_entries(entries: Vec<(String, Value)>) -> Self {
        let mut fields = RecordFields::new();
        for (key, value) in entries { fields.insert(key, value); }
        Self(std::sync::Mutex::new(fields))
    }
    pub fn snapshot(&self) -> Self { Self(std::sync::Mutex::new(self.lock().clone())) }
    pub fn get(&self, key: &str) -> Option<Value> { self.lock().get(key).cloned() }
    pub fn entries(&self) -> Vec<(String, Value)> { self.lock().entries() }
    pub fn insert(&self, key: String, value: Value) -> Option<Value> { self.lock().insert(key, value) }
    pub fn remove(&self, key: &str) -> Option<Value> { self.lock().remove(key) }
    fn lock(&self) -> std::sync::MutexGuard<'_, RecordFields> {
        self.0.lock().unwrap_or_else(|poisoned| poisoned.into_inner())
    }
}

impl Value {
    // Immutable PureScript records can enter Foreign.Object through Foreign
    // readers. Keep existing object handles shared; materialize own fields
    // only when crossing from a native immutable record representation.
    pub fn __purust_foreign_object(&self) -> std::rc::Rc<SharedRecord> {
        if let Value::Class(native) = self.resolve() {
            return native.downcast_ref::<std::rc::Rc<SharedRecord>>()
                .expect("Expected a Foreign.Object handle").clone();
        }
        let fields = self.__purust_record_fields().expect("Expected an object or record");
        std::rc::Rc::new(SharedRecord::from_entries(fields.entries()))
    }
}
pub fn mk_unit(_val: ()) -> UnknownType { Value::Unit }
pub fn mk_int(val: i64) -> UnknownType { Value::Int(val) }
pub fn mk_bool(val: bool) -> UnknownType { Value::Bool(val) }
pub fn mk_number(val: f64) -> UnknownType { Value::Number(val) }
pub fn mk_string(val: &str) -> UnknownType { Value::String(val.to_string()) }
pub fn mk_char(val: char) -> UnknownType { Value::Char(val) }
pub fn mk_array(val: Vec<UnknownType>) -> UnknownType { Value::Array(std::rc::Rc::new(val)) }

#[derive(Clone, Default)]
pub struct Thunk {
    pub value: std::sync::OnceLock<Value>,
}

#[derive(Clone, Default)]
pub struct Record_a {
    pub tag: &'static str,
    pub vals: Option<std::rc::Rc<Vec<UnknownType>>>,
    pub call: Option<Func1<UnknownType, UnknownType>>,
    pub a: Option<UnknownType>,
    pub b: Option<UnknownType>,
    pub c: Option<UnknownType>,
    pub d: Option<UnknownType>,
    pub e: Option<UnknownType>,
    pub f: Option<UnknownType>,
    pub keep: Option<UnknownType>,
    pub z: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_a_b_keep_z {
    pub a: Option<UnknownType>,
    pub b: Option<UnknownType>,
    pub keep: Option<UnknownType>,
    pub z: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_c_d {
    pub c: Option<UnknownType>,
    pub d: Option<UnknownType>,
}

#[derive(Clone, Default)]
pub struct Record_e_f {
    pub e: Option<UnknownType>,
    pub f: Option<UnknownType>,
}



#[derive(Clone)]
pub enum Func1<T1, R> {
    Static(fn(T1) -> R),
    Shared(std::rc::Rc<dyn Fn(T1) -> R>),
}

impl<T1: 'static, R: 'static> std::ops::Deref for Func1<T1, R> {
    type Target = dyn Fn(T1) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func1::Static(f) => f,
            Func1::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func2<T1, T2, R> {
    Static(fn(T1, T2) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2) -> R>),
}

impl<T1: 'static, T2: 'static, R: 'static> std::ops::Deref for Func2<T1, T2, R> {
    type Target = dyn Fn(T1, T2) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func2::Static(f) => f,
            Func2::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func3<T1, T2, T3, R> {
    Static(fn(T1, T2, T3) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, R: 'static> std::ops::Deref for Func3<T1, T2, T3, R> {
    type Target = dyn Fn(T1, T2, T3) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func3::Static(f) => f,
            Func3::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func4<T1, T2, T3, T4, R> {
    Static(fn(T1, T2, T3, T4) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, R: 'static> std::ops::Deref for Func4<T1, T2, T3, T4, R> {
    type Target = dyn Fn(T1, T2, T3, T4) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func4::Static(f) => f,
            Func4::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func5<T1, T2, T3, T4, T5, R> {
    Static(fn(T1, T2, T3, T4, T5) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, R: 'static> std::ops::Deref for Func5<T1, T2, T3, T4, T5, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func5::Static(f) => f,
            Func5::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func6<T1, T2, T3, T4, T5, T6, R> {
    Static(fn(T1, T2, T3, T4, T5, T6) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, R: 'static> std::ops::Deref for Func6<T1, T2, T3, T4, T5, T6, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func6::Static(f) => f,
            Func6::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func7<T1, T2, T3, T4, T5, T6, T7, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, R: 'static> std::ops::Deref for Func7<T1, T2, T3, T4, T5, T6, T7, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func7::Static(f) => f,
            Func7::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func8<T1, T2, T3, T4, T5, T6, T7, T8, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, R: 'static> std::ops::Deref for Func8<T1, T2, T3, T4, T5, T6, T7, T8, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func8::Static(f) => f,
            Func8::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func9<T1, T2, T3, T4, T5, T6, T7, T8, T9, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, R: 'static> std::ops::Deref for Func9<T1, T2, T3, T4, T5, T6, T7, T8, T9, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func9::Static(f) => f,
            Func9::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func10<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, T10: 'static, R: 'static> std::ops::Deref for Func10<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func10::Static(f) => f,
            Func10::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func11<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, T10: 'static, T11: 'static, R: 'static> std::ops::Deref for Func11<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func11::Static(f) => f,
            Func11::Shared(rc) => rc.as_ref(),
        }
    }
}

#[derive(Clone)]
pub enum Func12<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, R> {
    Static(fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) -> R),
    Shared(std::rc::Rc<dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) -> R>),
}

impl<T1: 'static, T2: 'static, T3: 'static, T4: 'static, T5: 'static, T6: 'static, T7: 'static, T8: 'static, T9: 'static, T10: 'static, T11: 'static, T12: 'static, R: 'static> std::ops::Deref for Func12<T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12, R> {
    type Target = dyn Fn(T1, T2, T3, T4, T5, T6, T7, T8, T9, T10, T11, T12) -> R;
    #[inline(always)]
    fn deref(&self) -> &Self::Target {
        match self {
            Func12::Static(f) => f,
            Func12::Shared(rc) => rc.as_ref(),
        }
    }
}


extern crate self as purust_core;
#[path = "/Users/0x1/Documents/htdocs/purust/purust/tests/runtime/perceus_ptr/src/lib.rs"] mod perceus_ptr;
// Code generated by purust for module ScalarRecords



#[inline(never)]
fn ScalarRecords_nextFields__purust_record_field_0_0(mut purs_local_62: i64) -> i64 {
    // AST: Typed(Abs(..., Typed(Typed(PrimOp(...)))))
/* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_62 + /* Typed i64 <- i64 : Lit */1)
}

#[inline(never)]
fn ScalarRecords_nextFields__purust_record_field_1_0(mut purs_local_62: i64) -> i64 {
    // AST: Typed(Abs(..., Typed(Typed(PrimOp(...)))))
/* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_62 + /* Typed i64 <- i64 : Lit */2)
}

#[inline(never)]
fn ScalarRecords_nextFields__purust_record_field_2_0(mut purs_local_62: i64) -> i64 {
    // AST: Typed(Abs(..., Typed(Typed(PrimOp(...)))))
/* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_62 + /* Typed i64 <- i64 : Lit */3)
}

#[inline(never)]
fn ScalarRecords_nextFields__purust_record_field_3_0(mut purs_local_62: i64, mut purs_local_63: i64) -> i64 {
    // AST: Typed(Abs(..., Typed(Typed(PrimOp(...)))))
/* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_63 + /* Typed i64 <- i64 : PrimOp(...) */{ let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */purs_local_62; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) })
}

fn ScalarRecords_nextWithLet__purust_record_field_0_0(mut purs_local_63: i64, mut purs_local_64: i64) -> i64 {
    // AST: Typed(Abs(..., Let(...)))
{
    let mut purs_local_62 = /* Typed i64 <- i64 : Local(...) */purs_local_63;
    drop(purs_local_62);
    /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_64 + /* Typed i64 <- i64 : Lit */1)
}
}

fn ScalarRecords_nextWithLet__purust_record_field_1_0(mut purs_local_63: i64) -> i64 {
    // AST: Typed(Abs(..., Let(...)))
{
    let mut purs_local_62 = /* Typed i64 <- i64 : Local(...) */purs_local_63;
    /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_62 + /* Typed i64 <- i64 : Lit */2)
}
}

fn ScalarRecords_nextWithLet__purust_record_field_2_0(mut purs_local_63: i64, mut purs_local_64: i64) -> i64 {
    // AST: Typed(Abs(..., Let(...)))
{
    let mut purs_local_62 = /* Typed i64 <- i64 : Local(...) */purs_local_63;
    drop(purs_local_62);
    /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_64 + /* Typed i64 <- i64 : Lit */3)
}
}

fn ScalarRecords_nextWithLet__purust_record_field_3_0(mut purs_local_63: i64, mut purs_local_64: i64, mut purs_local_65: i64) -> i64 {
    // AST: Typed(Abs(..., Let(...)))
{
    let mut purs_local_62 = /* Typed i64 <- i64 : Local(...) */purs_local_64;
    drop(purs_local_62);
    /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_65 + /* Typed i64 <- i64 : PrimOp(...) */{ let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */purs_local_63; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) })
}
}

fn ScalarRecords_rotateFields__purust_record_loop_0(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType, mut purs_local_2: i64, mut purs_local_3: i64, mut purs_local_4: i64, mut purs_local_5: i64, mut purs_local_6: i64, mut purs_local_7: bool) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Branch(...))))
    loop {
        break /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == /* Typed i64 <- i64 : Lit */0) {
        /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : Local(...) */purs_local_7 {
        /* Typed crate::UnknownType <- crate::UnknownType : Update */{
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_2);
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_3);
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_4);
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_5);
    let mut _base = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
    _base.set_a(_record_update_0);
    let mut _record_child = _base.get_b();
    _base.set_b(purust_core::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    let mut _record_child_1 = _record_child.get_d();
    _record_child.set_d(purust_core::Value::Unit);
    _record_child_1.set_e(_record_child_1_update_0);
    _record_child_1.set_f(_record_child_1_update_1);
    _record_child.set_d(_record_child_1);
    _base.set_b(_record_child);
    _base
}
    } else {
        /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1
    }
    } else {
        {
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
        let _tco_temp_2 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_3 + /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_3 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_2 + /* Typed i64 <- i64 : Lit */2);
        let _tco_temp_4 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_4 + /* Typed i64 <- i64 : Lit */3);
        let _tco_temp_5 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_5 + /* Typed i64 <- i64 : PrimOp(...) */{ let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */purs_local_0; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) });
        let _tco_temp_6 = /* Typed i64 <- i64 : Local(...) */purs_local_6;
        let _tco_temp_7 = /* Typed bool <- bool : Lit */true;
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        purs_local_2 = _tco_temp_2;
        purs_local_3 = _tco_temp_3;
        purs_local_4 = _tco_temp_4;
        purs_local_5 = _tco_temp_5;
        purs_local_6 = _tco_temp_6;
        purs_local_7 = _tco_temp_7;
        continue;
    }
    };
    }
}

pub fn ScalarRecords_rotateFields(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Let(...))))
{
    let mut purs_local_2 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_a()).unwrap_int();
    {
    let mut purs_local_3 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_c()).unwrap_int();
    {
    let mut purs_local_4 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int();
    {
    let mut purs_local_5 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int();
    {
    let mut purs_local_6 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_keep()).unwrap_int();
    /* Typed crate::UnknownType <- crate::UnknownType : App(Typed(Var(...))) */ScalarRecords_rotateFields__purust_record_loop_0(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1, /* Typed i64 <- i64 : Local(...) */purs_local_2, /* Typed i64 <- i64 : Local(...) */purs_local_3, /* Typed i64 <- i64 : Local(...) */purs_local_4, /* Typed i64 <- i64 : Local(...) */purs_local_5, /* Typed i64 <- i64 : Local(...) */purs_local_6, /* Typed bool <- bool : Lit */false)
}
}
}
}
}
}

fn ScalarRecords_withLet__purust_record_loop_0(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType, mut purs_local_41: i64, mut purs_local_42: i64, mut purs_local_43: i64, mut purs_local_44: i64, mut purs_local_45: i64, mut purs_local_46: bool) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Let(...))))
    loop {
        break {
    let mut purs_local_40 = /* Typed i64 <- i64 : Local(...) */purs_local_41.clone();
    /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == /* Typed i64 <- i64 : Lit */0) {
        /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : Local(...) */purs_local_46 {
        /* Typed crate::UnknownType <- crate::UnknownType : Update */{
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_41);
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_42);
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_43);
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_44);
    let mut _base = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
    _base.set_a(_record_update_0);
    let mut _record_child = _base.get_b();
    _base.set_b(purust_core::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    let mut _record_child_1 = _record_child.get_d();
    _record_child.set_d(purust_core::Value::Unit);
    _record_child_1.set_e(_record_child_1_update_0);
    _record_child_1.set_f(_record_child_1_update_1);
    _record_child.set_d(_record_child_1);
    _base.set_b(_record_child);
    _base
}
    } else {
        /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1
    }
    } else {
        {
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
        let _tco_temp_2 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_42 + /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_3 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_40 + /* Typed i64 <- i64 : Lit */2);
        let _tco_temp_4 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_43 + /* Typed i64 <- i64 : Lit */3);
        let _tco_temp_5 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_44 + /* Typed i64 <- i64 : PrimOp(...) */{ let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */purs_local_0; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) });
        let _tco_temp_6 = /* Typed i64 <- i64 : Local(...) */purs_local_45;
        let _tco_temp_7 = /* Typed bool <- bool : Lit */true;
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        purs_local_41 = _tco_temp_2;
        purs_local_42 = _tco_temp_3;
        purs_local_43 = _tco_temp_4;
        purs_local_44 = _tco_temp_5;
        purs_local_45 = _tco_temp_6;
        purs_local_46 = _tco_temp_7;
        continue;
    }
    }
};
    }
}

pub fn ScalarRecords_withLet(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Let(...))))
{
    let mut purs_local_41 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_a()).unwrap_int();
    {
    let mut purs_local_42 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_c()).unwrap_int();
    {
    let mut purs_local_43 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int();
    {
    let mut purs_local_44 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int();
    {
    let mut purs_local_45 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_keep()).unwrap_int();
    /* Typed crate::UnknownType <- crate::UnknownType : App(Typed(Var(...))) */ScalarRecords_withLet__purust_record_loop_0(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1, /* Typed i64 <- i64 : Local(...) */purs_local_41, /* Typed i64 <- i64 : Local(...) */purs_local_42, /* Typed i64 <- i64 : Local(...) */purs_local_43, /* Typed i64 <- i64 : Local(...) */purs_local_44, /* Typed i64 <- i64 : Local(...) */purs_local_45, /* Typed bool <- bool : Lit */false)
}
}
}
}
}
}

#[inline(never)]
pub fn ScalarRecords_nextFields(mut purs_local_60: i64, mut purs_local_61: crate::UnknownType) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Update)))
/* Typed crate::UnknownType <- crate::UnknownType : Update */{
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_b().__purust_borrow_c()).unwrap_int() + /* Typed i64 <- i64 : Lit */1));
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_a()).unwrap_int() + /* Typed i64 <- i64 : Lit */2));
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int() + /* Typed i64 <- i64 : Lit */3));
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int() + /* Typed i64 <- i64 : PrimOp(...) */{ let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */purs_local_60; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) }));
    let mut _base = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_61;
    _base.set_a(_record_update_0);
    let mut _record_child = _base.get_b();
    _base.set_b(purust_core::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    let mut _record_child_1 = _record_child.get_d();
    _record_child.set_d(purust_core::Value::Unit);
    _record_child_1.set_e(_record_child_1_update_0);
    _record_child_1.set_f(_record_child_1_update_1);
    _record_child.set_d(_record_child_1);
    _base.set_b(_record_child);
    _base
}
}

fn ScalarRecords_throughCall__purust_record_loop_0(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType, mut purs_local_2: i64, mut purs_local_3: i64, mut purs_local_4: i64, mut purs_local_5: i64, mut purs_local_6: i64, mut purs_local_7: bool) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Branch(...))))
    loop {
        break /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == /* Typed i64 <- i64 : Lit */0) {
        /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : Local(...) */purs_local_7 {
        /* Typed crate::UnknownType <- crate::UnknownType : Update */{
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_2);
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_3);
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_4);
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_5);
    let mut _base = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
    _base.set_a(_record_update_0);
    let mut _record_child = _base.get_b();
    _base.set_b(purust_core::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    let mut _record_child_1 = _record_child.get_d();
    _record_child.set_d(purust_core::Value::Unit);
    _record_child_1.set_e(_record_child_1_update_0);
    _record_child_1.set_f(_record_child_1_update_1);
    _record_child.set_d(_record_child_1);
    _base.set_b(_record_child);
    _base
}
    } else {
        /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1
    }
    } else {
        {
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
        let _tco_temp_2 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextFields__purust_record_field_0_0(/* Typed i64 <- i64 : Local(...) */purs_local_3);
        let _tco_temp_3 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextFields__purust_record_field_1_0(/* Typed i64 <- i64 : Local(...) */purs_local_2);
        let _tco_temp_4 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextFields__purust_record_field_2_0(/* Typed i64 <- i64 : Local(...) */purs_local_4);
        let _tco_temp_5 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextFields__purust_record_field_3_0(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed i64 <- i64 : Local(...) */purs_local_5);
        let _tco_temp_6 = /* Typed i64 <- i64 : Local(...) */purs_local_6;
        let _tco_temp_7 = /* Typed bool <- bool : Lit */true;
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        purs_local_2 = _tco_temp_2;
        purs_local_3 = _tco_temp_3;
        purs_local_4 = _tco_temp_4;
        purs_local_5 = _tco_temp_5;
        purs_local_6 = _tco_temp_6;
        purs_local_7 = _tco_temp_7;
        continue;
    }
    };
    }
}

pub fn ScalarRecords_throughCall(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Let(...))))
{
    let mut purs_local_2 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_a()).unwrap_int();
    {
    let mut purs_local_3 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_c()).unwrap_int();
    {
    let mut purs_local_4 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int();
    {
    let mut purs_local_5 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int();
    {
    let mut purs_local_6 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_keep()).unwrap_int();
    /* Typed crate::UnknownType <- crate::UnknownType : App(Typed(Var(...))) */ScalarRecords_throughCall__purust_record_loop_0(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1, /* Typed i64 <- i64 : Local(...) */purs_local_2, /* Typed i64 <- i64 : Local(...) */purs_local_3, /* Typed i64 <- i64 : Local(...) */purs_local_4, /* Typed i64 <- i64 : Local(...) */purs_local_5, /* Typed i64 <- i64 : Local(...) */purs_local_6, /* Typed bool <- bool : Lit */false)
}
}
}
}
}
}

pub fn ScalarRecords_nextWithLet(mut purs_local_60: i64, mut purs_local_61: crate::UnknownType) -> crate::UnknownType {
    // AST: Typed(Abs(..., Let(...)))
{
    let mut purs_local_62 = /* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_a()).unwrap_int();
    /* Typed crate::UnknownType <- crate::UnknownType : Update */{
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_b().__purust_borrow_c()).unwrap_int() + /* Typed i64 <- i64 : Lit */1));
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_62 + /* Typed i64 <- i64 : Lit */2));
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int() + /* Typed i64 <- i64 : Lit */3));
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : PrimOp(...) */(/* purust record: borrowed scalar */((&purs_local_61).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int() + /* Typed i64 <- i64 : PrimOp(...) */{ let _mod_l: i64 = /* Typed i64 <- i64 : Local(...) */purs_local_60; let _mod_r: i64 = /* Typed i64 <- i64 : Lit */5; _mod_l.checked_rem_euclid(_mod_r).unwrap_or(0_i64) }));
    let mut _base = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_61;
    _base.set_a(_record_update_0);
    let mut _record_child = _base.get_b();
    _base.set_b(purust_core::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    let mut _record_child_1 = _record_child.get_d();
    _record_child.set_d(purust_core::Value::Unit);
    _record_child_1.set_e(_record_child_1_update_0);
    _record_child_1.set_f(_record_child_1_update_1);
    _record_child.set_d(_record_child_1);
    _base.set_b(_record_child);
    _base
}
}
}

fn ScalarRecords_throughLetCall__purust_record_loop_0(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType, mut purs_local_2: i64, mut purs_local_3: i64, mut purs_local_4: i64, mut purs_local_5: i64, mut purs_local_6: i64, mut purs_local_7: bool) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Branch(...))))
    loop {
        break /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() == /* Typed i64 <- i64 : Lit */0) {
        /* Typed crate::UnknownType <- crate::UnknownType : Branch(...) */if /* Typed bool <- bool : Local(...) */purs_local_7 {
        /* Typed crate::UnknownType <- crate::UnknownType : Update */{
    let _record_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_2);
    let _record_child_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_3);
    let _record_child_1_update_0 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_4);
    let _record_child_1_update_1 = crate::mk_int(/* Typed i64 <- i64 : Local(...) */purs_local_5);
    let mut _base = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
    _base.set_a(_record_update_0);
    let mut _record_child = _base.get_b();
    _base.set_b(purust_core::Value::Unit);
    _record_child.set_c(_record_child_update_0);
    let mut _record_child_1 = _record_child.get_d();
    _record_child.set_d(purust_core::Value::Unit);
    _record_child_1.set_e(_record_child_1_update_0);
    _record_child_1.set_f(_record_child_1_update_1);
    _record_child.set_d(_record_child_1);
    _base.set_b(_record_child);
    _base
}
    } else {
        /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1
    }
    } else {
        {
        let _tco_temp_0 = /* Typed i64 <- i64 : PrimOp(...) */(/* Typed i64 <- i64 : Local(...) */purs_local_0.clone() - /* Typed i64 <- i64 : Lit */1);
        let _tco_temp_1 = /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1;
        let _tco_temp_2 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextWithLet__purust_record_field_0_0(/* Typed i64 <- i64 : Local(...) */purs_local_2.clone(), /* Typed i64 <- i64 : Local(...) */purs_local_3);
        let _tco_temp_3 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextWithLet__purust_record_field_1_0(/* Typed i64 <- i64 : Local(...) */purs_local_2.clone());
        let _tco_temp_4 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextWithLet__purust_record_field_2_0(/* Typed i64 <- i64 : Local(...) */purs_local_2.clone(), /* Typed i64 <- i64 : Local(...) */purs_local_4);
        let _tco_temp_5 = /* Typed i64 <- i64 : App(Typed(Var(...))) */ScalarRecords_nextWithLet__purust_record_field_3_0(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed i64 <- i64 : Local(...) */purs_local_2, /* Typed i64 <- i64 : Local(...) */purs_local_5);
        let _tco_temp_6 = /* Typed i64 <- i64 : Local(...) */purs_local_6;
        let _tco_temp_7 = /* Typed bool <- bool : Lit */true;
        purs_local_0 = _tco_temp_0;
        purs_local_1 = _tco_temp_1;
        purs_local_2 = _tco_temp_2;
        purs_local_3 = _tco_temp_3;
        purs_local_4 = _tco_temp_4;
        purs_local_5 = _tco_temp_5;
        purs_local_6 = _tco_temp_6;
        purs_local_7 = _tco_temp_7;
        continue;
    }
    };
    }
}

pub fn ScalarRecords_throughLetCall(mut purs_local_0: i64, mut purs_local_1: crate::UnknownType) -> crate::UnknownType {
    // AST: Typed(Abs(..., Typed(Let(...))))
{
    let mut purs_local_2 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_a()).unwrap_int();
    {
    let mut purs_local_3 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_c()).unwrap_int();
    {
    let mut purs_local_4 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_e()).unwrap_int();
    {
    let mut purs_local_5 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_b().__purust_borrow_d().__purust_borrow_f()).unwrap_int();
    {
    let mut purs_local_6 = /* purust record: borrowed scalar */((&purs_local_1).__purust_borrow_keep()).unwrap_int();
    /* Typed crate::UnknownType <- crate::UnknownType : App(Typed(Var(...))) */ScalarRecords_throughLetCall__purust_record_loop_0(/* Typed i64 <- i64 : Local(...) */purs_local_0, /* Typed crate::UnknownType <- crate::UnknownType : Local(...) */purs_local_1, /* Typed i64 <- i64 : Local(...) */purs_local_2, /* Typed i64 <- i64 : Local(...) */purs_local_3, /* Typed i64 <- i64 : Local(...) */purs_local_4, /* Typed i64 <- i64 : Local(...) */purs_local_5, /* Typed i64 <- i64 : Local(...) */purs_local_6, /* Typed bool <- bool : Lit */false)
}
}
}
}
}
}


fn seed(v: [i64;4]) -> Value {
    Value::Record_a_b_keep_z(perceus_ptr::PerceusPtr::new(Record_a_b_keep_z {
        a: Some(mk_int(v[0])), keep: Some(mk_int(99)), z: Some(mk_int(777)),
        b: Some(Value::Record_c_d(perceus_ptr::PerceusPtr::new(Record_c_d {
            c: Some(mk_int(v[1])), d: Some(Value::Record_e_f(perceus_ptr::PerceusPtr::new(Record_e_f {
                e: Some(mk_int(v[2])), f: Some(mk_int(v[3]))
            })))
        })))
    }))
}
fn values(r: &Value) -> [i64;4] {
    [r.__purust_borrow_a().unwrap_int(), r.__purust_borrow_b().__purust_borrow_c().unwrap_int(),
     r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_e().unwrap_int(),
     r.__purust_borrow_b().__purust_borrow_d().__purust_borrow_f().unwrap_int()]
}
fn expected(mut n:i64, mut v:[i64;4])->[i64;4] {
    while n != 0 { v=[v[1]+1,v[0]+2,v[2]+3,v[3]+n.checked_rem_euclid(5).unwrap_or(0)]; n-=1; } v
}

#[global_allocator] static BENCH_ALLOCATOR: mimalloc::MiMalloc = mimalloc::MiMalloc;
#[inline(never)]
fn reference_loop(mut n: i64, mut r: Value) -> Value {
    while n != 0 { r = ScalarRecords_nextFields(n, r); n -= 1; }
    r
}
fn main() {
    let args: Vec<String> = std::env::args().collect();
    let f: fn(i64, Value) -> Value = match args[1].as_str() {
        "reference" => reference_loop, "generated" => ScalarRecords_throughCall, _ => panic!()
    };
    let input = seed([0,0,0,0]);
    let oracle = expected(10000,[0,0,0,0]);
    for _ in 0..3 { drop(std::hint::black_box(f(std::hint::black_box(10000), std::hint::black_box(input.clone())))); }
    for _ in 0..10 {
        let start=std::time::Instant::now();
        let result=f(std::hint::black_box(10000),std::hint::black_box(input.clone()));
        let actual=std::hint::black_box(values(&result)); drop(result);
        let elapsed=start.elapsed().as_nanos();
        assert_eq!(actual,oracle); assert_eq!(values(&input),[0,0,0,0]);
        println!("{elapsed}");
    }
}

pub mod PureScript_Data_Function_Uncurried {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    pub mod Data_Function_Uncurried_FFI {
        use super::*;
        use fable_library_rust::Native_::Func0;
        use fable_library_rust::Native_::defaultOf;
        pub fn mkFn0(r#fn: &dyn Any) -> &dyn Any {
            &Func0::new({
                            let r#fn = r#fn.clone();
                            move || r#fn(defaultOf())
                        })
        }
        pub fn mkFn1(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn2(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn3(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn4(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn5(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn6(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn7(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn8(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn9(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn mkFn10(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn0(r#fn: &dyn Any) -> &dyn Any { r#fn() }
        pub fn runFn1(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn2(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn3(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn4(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn5(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn6(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn7(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn8(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn9(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
        pub fn runFn10(r#fn: &dyn Any) -> &dyn Any { r#fn.clone() }
    }
    pub fn Data_Function_Uncurried_mkFn0() -> &dyn Any {
        static Data_Function_Uncurried_mkFn0: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn0.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn0(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn10() -> &dyn Any {
        static Data_Function_Uncurried_mkFn10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn10.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn10(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn2() -> &dyn Any {
        static Data_Function_Uncurried_mkFn2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn2.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn2(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn3() -> &dyn Any {
        static Data_Function_Uncurried_mkFn3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn3.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn3(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn4() -> &dyn Any {
        static Data_Function_Uncurried_mkFn4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn4.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn4(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn5() -> &dyn Any {
        static Data_Function_Uncurried_mkFn5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn5.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn5(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn6() -> &dyn Any {
        static Data_Function_Uncurried_mkFn6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn6.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn6(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn7() -> &dyn Any {
        static Data_Function_Uncurried_mkFn7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn7.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn7(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn8() -> &dyn Any {
        static Data_Function_Uncurried_mkFn8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn8.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn8(r#fn)))
    }
    pub fn Data_Function_Uncurried_mkFn9() -> &dyn Any {
        static Data_Function_Uncurried_mkFn9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn9.get_or_init(||
                                                      &Func1::new(move |r#fn|
                                                                      PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::mkFn9(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn0() -> &dyn Any {
        static Data_Function_Uncurried_runFn0: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn0.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn0(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn10() -> &dyn Any {
        static Data_Function_Uncurried_runFn10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn10.get_or_init(||
                                                        &Func1::new(move
                                                                        |r#fn|
                                                                        PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn10(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn2() -> &dyn Any {
        static Data_Function_Uncurried_runFn2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn2.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn2(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn3() -> &dyn Any {
        static Data_Function_Uncurried_runFn3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn3.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn3(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn4() -> &dyn Any {
        static Data_Function_Uncurried_runFn4: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn4.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn4(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn5() -> &dyn Any {
        static Data_Function_Uncurried_runFn5: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn5.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn5(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn6() -> &dyn Any {
        static Data_Function_Uncurried_runFn6: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn6.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn6(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn7() -> &dyn Any {
        static Data_Function_Uncurried_runFn7: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn7.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn7(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn8() -> &dyn Any {
        static Data_Function_Uncurried_runFn8: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn8.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn8(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn9() -> &dyn Any {
        static Data_Function_Uncurried_runFn9: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn9.get_or_init(||
                                                       &Func1::new(move |r#fn|
                                                                       PureScript_Data_Function_Uncurried::Data_Function_Uncurried_FFI::runFn9(r#fn)))
    }
    pub fn Data_Function_Uncurried_runFn1() -> &dyn Any {
        static Data_Function_Uncurried_runFn1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_runFn1.get_or_init(||
                                                       &Func1::new(move |f|
                                                                       f.clone()))
    }
    pub fn Data_Function_Uncurried_mkFn1() -> &dyn Any {
        static Data_Function_Uncurried_mkFn1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Function_Uncurried_mkFn1.get_or_init(||
                                                      &Func1::new(move |f|
                                                                      f.clone()))
    }
}

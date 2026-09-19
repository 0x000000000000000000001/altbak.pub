pub mod PureScript_Partial {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Partial_FFI {
        use super::*;
        pub fn _crashWith<a: Clone + 'static, b: Clone + 'static>(msg: a)
         -> b {
            panic!("{}", Sharpurs_Prelude::unbox(&&msg),)
        }
    }
    pub fn Partial__crashWith() -> &dyn Any {
        static Partial__crashWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Partial__crashWith.get_or_init(||
                                           &Func1::new(move |arg0|
                                                           &PureScript_Partial::Partial_FFI::_crashWith(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Partial_crashWith() -> &dyn Any {
        static Partial_crashWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Partial_crashWith.get_or_init(||
                                          &Func1::new(move |usd__unused|
                                                          &PureScript_Partial::Partial__crashWith()))
    }
    pub fn Partial_crash() -> &dyn Any {
        static Partial_crash: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Partial_crash.get_or_init(||
                                      &Func1::new(move |usd__unused|
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial::Partial_crashWith(),
                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                       &&&string("Partial.crash: partial function"))))
    }
}

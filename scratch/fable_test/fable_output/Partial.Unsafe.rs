pub mod PureScript_Partial_Unsafe {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_c2816d8::PureScript_Partial;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Partial_Unsafe_FFI {
        use super::*;
        use fable_library_rust::Native_::defaultOf;
        pub fn _unsafePartial<a: Clone + 'static>(f: a) -> &dyn Any {
            Sharpurs_Prelude::sharpurs_apply(&&f, &&defaultOf())
        }
    }
    pub fn Partial_Unsafe__unsafePartial() -> &dyn Any {
        static Partial_Unsafe__unsafePartial: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Partial_Unsafe__unsafePartial.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Partial_Unsafe::Partial_Unsafe_FFI::_unsafePartial(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Partial_Unsafe_unsafePartial() -> &dyn Any {
        static Partial_Unsafe_unsafePartial: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Partial_Unsafe_unsafePartial.get_or_init(||
                                                     &PureScript_Partial_Unsafe::Partial_Unsafe__unsafePartial())
    }
    pub fn Partial_Unsafe_unsafeCrashWith() -> &dyn Any {
        static Partial_Unsafe_unsafeCrashWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Partial_Unsafe_unsafeCrashWith.get_or_init(||
                                                       &Func1::new(move |msg|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                        &&&Func1::new({
                                                                                                                          let msg
                                                                                                                              =
                                                                                                                              msg.clone();
                                                                                                                          move
                                                                                                                              |usd__unused|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial::Partial_crashWith(),
                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                               &&&msg)
                                                                                                                      }))))
    }
}

pub mod PureScript_Effect_Exception_Unsafe {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::MutCell;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_ceb943a5::PureScript_Effect_Exception;
    use crate::module_90c22cd8::PureScript_Effect_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Effect_Exception_Unsafe_unsafeThrowException() -> &dyn Any {
        static Effect_Exception_Unsafe_unsafeThrowException:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Exception_Unsafe_unsafeThrowException.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                         &&&PureScript_Effect_Unsafe::Effect_Unsafe_unsafePerformEffect()),
                                                                                                      &&&PureScript_Effect_Exception::Effect_Exception_throwException()))
    }
    pub fn Effect_Exception_Unsafe_unsafeThrow() -> &dyn Any {
        static Effect_Exception_Unsafe_unsafeThrow: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Effect_Exception_Unsafe_unsafeThrow.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                &&&PureScript_Effect_Exception_Unsafe::Effect_Exception_Unsafe_unsafeThrowException()),
                                                                                             &&&PureScript_Effect_Exception::Effect_Exception_error()))
    }
}

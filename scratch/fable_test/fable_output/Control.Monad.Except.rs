pub mod PureScript_Control_Monad_Except {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_8ea4b64e::PureScript_Control_Monad_Except_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Except_unwrap() -> &dyn Any {
        static Control_Monad_Except_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_unwrap.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_Except_withExcept() -> &dyn Any {
        static Control_Monad_Except_withExcept: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_withExcept.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_withExceptT(),
                                                                                         &&&PureScript_Data_Identity::Data_Identity_functorIdentity()))
    }
    pub fn Control_Monad_Except_runExcept() -> &dyn Any {
        static Control_Monad_Except_runExcept: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_runExcept.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                           &&&PureScript_Control_Monad_Except::Control_Monad_Except_unwrap()),
                                                                                        &&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_runExceptT()))
    }
    pub fn Control_Monad_Except_mapExcept() -> &dyn Any {
        static Control_Monad_Except_mapExcept: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Except_mapExcept.get_or_init(||
                                                       &Func1::new(move |f|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Except_Trans::Control_Monad_Except_Trans_mapExceptT(),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                              &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                 f),
                                                                                                                                                                              &&&PureScript_Control_Monad_Except::Control_Monad_Except_unwrap())))))
    }
}

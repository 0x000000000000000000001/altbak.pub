pub mod PureScript_Control_Comonad_Traced {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_7d1f6438::PureScript_Control_Comonad_Traced_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Traced_traced() -> &dyn Any {
        static Control_Comonad_Traced_traced: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_traced.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                          &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                       &&&PureScript_Control_Comonad_Traced_Trans::Control_Comonad_Traced_Trans_TracedT()))
    }
    pub fn Control_Comonad_Traced_runTraced() -> &dyn Any {
        static Control_Comonad_Traced_runTraced: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Traced_runTraced.get_or_init(||
                                                         &Func1::new(move |v|
                                                                         {
                                                                             let t =
                                                                                 Sharpurs_Prelude::unbox(v);
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                              &&&t)
                                                                         }))
    }
}

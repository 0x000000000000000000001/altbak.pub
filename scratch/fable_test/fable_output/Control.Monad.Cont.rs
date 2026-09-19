pub mod PureScript_Control_Monad_Cont {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_dbc0dd87::PureScript_Control_Monad_Cont_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Cont_unwrap() -> &dyn Any {
        static Control_Monad_Cont_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_unwrap.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_Cont_withCont() -> &dyn Any {
        static Control_Monad_Cont_withCont: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_withCont.get_or_init(||
                                                    &Func1::new(move |f|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_withContT(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                              &&&PureScript_Data_Identity::Data_Identity_Identity())),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                              f),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                              &&&PureScript_Control_Monad_Cont::Control_Monad_Cont_unwrap()))))))
    }
    pub fn Control_Monad_Cont_runCont() -> &dyn Any {
        static Control_Monad_Cont_runCont: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_runCont.get_or_init(||
                                                   &Func1::new(move |cc|
                                                                   &Func1::new({
                                                                                   let cc
                                                                                       =
                                                                                       cc.clone();
                                                                                   move
                                                                                       |k|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_runContT(),
                                                                                                                                                                                              &&&cc),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                 &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                                                                              k)))
                                                                               })))
    }
    pub fn Control_Monad_Cont_mapCont() -> &dyn Any {
        static Control_Monad_Cont_mapCont: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_mapCont.get_or_init(||
                                                   &Func1::new(move |f|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_mapContT(),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                          &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                             f),
                                                                                                                                                                          &&&PureScript_Control_Monad_Cont::Control_Monad_Cont_unwrap())))))
    }
    pub fn Control_Monad_Cont_cont() -> &dyn Any {
        static Control_Monad_Cont_cont: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_cont.get_or_init(||
                                                &Func1::new(move |f|
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Cont_Trans::Control_Monad_Cont_Trans_ContT(),
                                                                                                 &&&Func1::new({
                                                                                                                   let f
                                                                                                                       =
                                                                                                                       f.clone();
                                                                                                                   move
                                                                                                                       |c|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                 &&&PureScript_Control_Monad_Cont::Control_Monad_Cont_unwrap()),
                                                                                                                                                                                                                              c)))
                                                                                                               }))))
    }
}

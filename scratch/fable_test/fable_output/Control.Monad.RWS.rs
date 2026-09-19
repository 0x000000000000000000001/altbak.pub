pub mod PureScript_Control_Monad_RWS {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_9e2ffde7::PureScript_Control_Monad_RWS_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_RWS_pure() -> &dyn Any {
        static Control_Monad_RWS_pure: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_pure.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                &&&PureScript_Data_Identity::Data_Identity_applicativeIdentity()))
    }
    pub fn Control_Monad_RWS_unwrap() -> &dyn Any {
        static Control_Monad_RWS_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_unwrap.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_RWS_unwrap1() -> &dyn Any {
        static Control_Monad_RWS_unwrap1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_unwrap1.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_RWS_unwrap2() -> &dyn Any {
        static Control_Monad_RWS_unwrap2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_unwrap2.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_RWS_withRWS() -> &dyn Any {
        static Control_Monad_RWS_withRWS: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_withRWS.get_or_init(||
                                                  &PureScript_Control_Monad_RWS_Trans::Control_Monad_RWS_Trans_withRWST())
    }
    pub fn Control_Monad_RWS_rws() -> &dyn Any {
        static Control_Monad_RWS_rws: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_rws.get_or_init(||
                                              &Func1::new(move |f|
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_RWS_Trans::Control_Monad_RWS_Trans_RWST(),
                                                                                               &&&Func1::new({
                                                                                                                 let f
                                                                                                                     =
                                                                                                                     f.clone();
                                                                                                                 move
                                                                                                                     |r|
                                                                                                                     &Func1::new({
                                                                                                                                     let r
                                                                                                                                         =
                                                                                                                                         r.clone();
                                                                                                                                     move
                                                                                                                                         |s|
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                             &&&PureScript_Control_Monad_RWS::Control_Monad_RWS_pure()),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                &&&r),
                                                                                                                                                                                                             s))
                                                                                                                                 })
                                                                                                             }))))
    }
    pub fn Control_Monad_RWS_runRWS() -> &dyn Any {
        static Control_Monad_RWS_runRWS: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_runRWS.get_or_init(||
                                                 &Func1::new(move |m|
                                                                 &Func1::new({
                                                                                 let m
                                                                                     =
                                                                                     m.clone();
                                                                                 move
                                                                                     |r|
                                                                                     &Func1::new({
                                                                                                     let r
                                                                                                         =
                                                                                                         r.clone();
                                                                                                     move
                                                                                                         |s|
                                                                                                         &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&m),
                                                                                                                                                                                                        &&&r),
                                                                                                                                                                     s))
                                                                                                 })
                                                                             })))
    }
    pub fn Control_Monad_RWS_mapRWS() -> &dyn Any {
        static Control_Monad_RWS_mapRWS: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_mapRWS.get_or_init(||
                                                 &Func1::new(move |f|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_RWS_Trans::Control_Monad_RWS_Trans_mapRWST(),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                        &&&PureScript_Control_Monad_RWS::Control_Monad_RWS_unwrap()),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                           f),
                                                                                                                                                                        &&&PureScript_Data_Identity::Data_Identity_Identity())))))
    }
    pub fn Control_Monad_RWS_execRWS() -> &dyn Any {
        static Control_Monad_RWS_execRWS: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_execRWS.get_or_init(||
                                                  &Func1::new(move |m|
                                                                  &Func1::new({
                                                                                  let m
                                                                                      =
                                                                                      m.clone();
                                                                                  move
                                                                                      |r|
                                                                                      &Func1::new({
                                                                                                      let r
                                                                                                          =
                                                                                                          r.clone();
                                                                                                      move
                                                                                                          |s|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                              &&&PureScript_Control_Monad_RWS::Control_Monad_RWS_unwrap1()),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_RWS_Trans::Control_Monad_RWS_Trans_execRWST(),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Identity::Data_Identity_monadIdentity()),
                                                                                                                                                                                                                                                    &&&m),
                                                                                                                                                                                                                 &&&r),
                                                                                                                                                                              s))
                                                                                                  })
                                                                              })))
    }
    pub fn Control_Monad_RWS_evalRWS() -> &dyn Any {
        static Control_Monad_RWS_evalRWS: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_RWS_evalRWS.get_or_init(||
                                                  &Func1::new(move |m|
                                                                  &Func1::new({
                                                                                  let m
                                                                                      =
                                                                                      m.clone();
                                                                                  move
                                                                                      |r|
                                                                                      &Func1::new({
                                                                                                      let r
                                                                                                          =
                                                                                                          r.clone();
                                                                                                      move
                                                                                                          |s|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                              &&&PureScript_Control_Monad_RWS::Control_Monad_RWS_unwrap2()),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_RWS_Trans::Control_Monad_RWS_Trans_evalRWST(),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Identity::Data_Identity_monadIdentity()),
                                                                                                                                                                                                                                                    &&&m),
                                                                                                                                                                                                                 &&&r),
                                                                                                                                                                              s))
                                                                                                  })
                                                                              })))
    }
}

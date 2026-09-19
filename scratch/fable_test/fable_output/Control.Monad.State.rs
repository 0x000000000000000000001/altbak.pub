pub mod PureScript_Control_Monad_State {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_3bc71566::PureScript_Control_Monad_State_Trans;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_1becb483::PureScript_Data_Identity;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_State_unwrap() -> &dyn Any {
        static Control_Monad_State_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_unwrap.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Control_Monad_State_withState() -> &dyn Any {
        static Control_Monad_State_withState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_withState.get_or_init(||
                                                      &PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_withStateT())
    }
    pub fn Control_Monad_State_runState() -> &dyn Any {
        static Control_Monad_State_runState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_runState.get_or_init(||
                                                     &Func1::new(move |v|
                                                                     {
                                                                         let s =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                             &&&PureScript_Control_Monad_State::Control_Monad_State_unwrap()),
                                                                                                          &&&s)
                                                                     }))
    }
    pub fn Control_Monad_State_mapState() -> &dyn Any {
        static Control_Monad_State_mapState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_mapState.get_or_init(||
                                                     &Func1::new(move |f|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Trans::Control_Monad_State_Trans_mapStateT(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                            &&&PureScript_Data_Identity::Data_Identity_Identity()),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                               f),
                                                                                                                                                                            &&&PureScript_Control_Monad_State::Control_Monad_State_unwrap())))))
    }
    pub fn Control_Monad_State_execState() -> &dyn Any {
        static Control_Monad_State_execState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_execState.get_or_init(||
                                                      &Func1::new(move |v|
                                                                      &Func1::new({
                                                                                      let v
                                                                                          =
                                                                                          v.clone();
                                                                                      move
                                                                                          |s|
                                                                                          &match Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&v),
                                                                                                                                                            &&&Sharpurs_Prelude::unbox(s))).as_ref()
                                                                                               {
                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                          x)
                                                                                               =>
                                                                                               x.clone(),
                                                                                           }
                                                                                  })))
    }
    pub fn Control_Monad_State_evalState() -> &dyn Any {
        static Control_Monad_State_evalState: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_evalState.get_or_init(||
                                                      &Func1::new(move |v|
                                                                      &Func1::new({
                                                                                      let v
                                                                                          =
                                                                                          v.clone();
                                                                                      move
                                                                                          |s|
                                                                                          &match Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&v),
                                                                                                                                                            &&&Sharpurs_Prelude::unbox(s))).as_ref()
                                                                                               {
                                                                                               Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                          _)
                                                                                               =>
                                                                                               x.clone(),
                                                                                           }
                                                                                  })))
    }
}

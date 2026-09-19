pub mod PureScript_Control_Monad_State_Class {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_State_Class_MonadStateusd_Dict() -> &dyn Any {
        static Control_Monad_State_Class_MonadStateusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Class_MonadStateusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Control_Monad_State_Class_state() -> &dyn Any {
        static Control_Monad_State_Class_state: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Class_state.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("state"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_State_Class_put() -> &dyn Any {
        static Control_Monad_State_Class_put: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Class_put.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadState|
                                                                      &Func1::new({
                                                                                      let dictMonadState
                                                                                          =
                                                                                          dictMonadState.clone();
                                                                                      move
                                                                                          |s|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                              &&&dictMonadState),
                                                                                                                           &&&Func1::new({
                                                                                                                                             let s
                                                                                                                                                 =
                                                                                                                                                 s.clone();
                                                                                                                                             move
                                                                                                                                                 |v|
                                                                                                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                                                                                                                                         &s))
                                                                                                                                         }))
                                                                                  })))
    }
    pub fn Control_Monad_State_Class_modify_() -> &dyn Any {
        static Control_Monad_State_Class_modify_: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Class_modify_.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadState|
                                                                          &Func1::new({
                                                                                          let dictMonadState
                                                                                              =
                                                                                              dictMonadState.clone();
                                                                                          move
                                                                                              |f|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                                  &&&dictMonadState),
                                                                                                                               &&&Func1::new({
                                                                                                                                                 let f
                                                                                                                                                     =
                                                                                                                                                     f.clone();
                                                                                                                                                 move
                                                                                                                                                     |s|
                                                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&PureScript_Data_Unit::Data_Unit_unit(),
                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                              s)))
                                                                                                                                             }))
                                                                                      })))
    }
    pub fn Control_Monad_State_Class_modify() -> &dyn Any {
        static Control_Monad_State_Class_modify: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Class_modify.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictMonadState|
                                                                         &Func1::new({
                                                                                         let dictMonadState
                                                                                             =
                                                                                             dictMonadState.clone();
                                                                                         move
                                                                                             |f|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                                 &&&dictMonadState),
                                                                                                                              &&&Func1::new({
                                                                                                                                                let f
                                                                                                                                                    =
                                                                                                                                                    f.clone();
                                                                                                                                                move
                                                                                                                                                    |s|
                                                                                                                                                    {
                                                                                                                                                        let s_prime =
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                             s);
                                                                                                                                                        &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&s_prime,
                                                                                                                                                                                                                &s_prime))
                                                                                                                                                    }
                                                                                                                                            }))
                                                                                     })))
    }
    pub fn Control_Monad_State_Class_gets() -> &dyn Any {
        static Control_Monad_State_Class_gets: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Class_gets.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictMonadState|
                                                                       &Func1::new({
                                                                                       let dictMonadState
                                                                                           =
                                                                                           dictMonadState.clone();
                                                                                       move
                                                                                           |f|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                                               &&&dictMonadState),
                                                                                                                            &&&Func1::new({
                                                                                                                                              let f
                                                                                                                                                  =
                                                                                                                                                  f.clone();
                                                                                                                                              move
                                                                                                                                                  |s|
                                                                                                                                                  &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                           s),
                                                                                                                                                                                                          s.clone()))
                                                                                                                                          }))
                                                                                   })))
    }
    pub fn Control_Monad_State_Class_get() -> &dyn Any {
        static Control_Monad_State_Class_get: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_State_Class_get.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadState|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_State_Class::Control_Monad_State_Class_state(),
                                                                                                                                          dictMonadState),
                                                                                                       &&&Func1::new(move
                                                                                                                         |s|
                                                                                                                         &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(s.clone(),
                                                                                                                                                                                 s.clone()))))))
    }
}

pub mod PureScript_Control_Monad_Writer_Class {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Writer_Class_MonadTellusd_Dict() -> &dyn Any {
        static Control_Monad_Writer_Class_MonadTellusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Class_MonadTellusd_Dict.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Control_Monad_Writer_Class_MonadWriterusd_Dict() -> &dyn Any {
        static Control_Monad_Writer_Class_MonadWriterusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Class_MonadWriterusd_Dict.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |x|
                                                                                       x.clone()))
    }
    pub fn Control_Monad_Writer_Class_tell() -> &dyn Any {
        static Control_Monad_Writer_Class_tell: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Class_tell.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("tell"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Writer_Class_pass() -> &dyn Any {
        static Control_Monad_Writer_Class_pass: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Class_pass.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("pass"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Writer_Class_listen() -> &dyn Any {
        static Control_Monad_Writer_Class_listen: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Class_listen.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("listen"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Writer_Class_listens() -> &dyn Any {
        static Control_Monad_Writer_Class_listens: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Class_listens.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictMonadWriter|
                                                                           {
                                                                               let Monad1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad1"),
                                                                                                                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadTell1"),
                                                                                                                                                                                            Sharpurs_Prelude::unbox(dictMonadWriter)),
                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                               let Bind1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                           Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined());
                                                                               let pure_var =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               &Func1::new({
                                                                                               let Bind1
                                                                                                   =
                                                                                                   Bind1.clone();
                                                                                               let dictMonadWriter
                                                                                                   =
                                                                                                   dictMonadWriter.clone();
                                                                                               let pure_var
                                                                                                   =
                                                                                                   pure_var.clone();
                                                                                               move
                                                                                                   |f|
                                                                                                   &Func1::new({
                                                                                                                   let f
                                                                                                                       =
                                                                                                                       f.clone();
                                                                                                                   move
                                                                                                                       |m|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                              &&&Bind1),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_listen(),
                                                                                                                                                                                                                                                                 &&&dictMonadWriter),
                                                                                                                                                                                                                              m)),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue:
                                                                                                                                                                                      LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                  &&&pure_var),
                                                                                                                                                                                                               &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                 Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                             }))))
                                                                                                                                                                          }))
                                                                                                               })
                                                                                           })
                                                                           }))
    }
    pub fn Control_Monad_Writer_Class_censor() -> &dyn Any {
        static Control_Monad_Writer_Class_censor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Writer_Class_censor.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictMonadWriter|
                                                                          {
                                                                              let Monad1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad1"),
                                                                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("MonadTell1"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonadWriter)),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              let Bind1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                          Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined());
                                                                              let pure_var =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                                                             Sharpurs_Prelude::unbox(&&Monad1)),
                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                              &Func1::new({
                                                                                              let Bind1
                                                                                                  =
                                                                                                  Bind1.clone();
                                                                                              let dictMonadWriter
                                                                                                  =
                                                                                                  dictMonadWriter.clone();
                                                                                              let pure_var
                                                                                                  =
                                                                                                  pure_var.clone();
                                                                                              move
                                                                                                  |f|
                                                                                                  &Func1::new({
                                                                                                                  let f
                                                                                                                      =
                                                                                                                      f.clone();
                                                                                                                  move
                                                                                                                      |m|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Writer_Class::Control_Monad_Writer_Class_pass(),
                                                                                                                                                                                          &&&dictMonadWriter),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                                                &&&Bind1),
                                                                                                                                                                                                                             m),
                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                            |a|
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                &&&pure_var),
                                                                                                                                                                                                                                             &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(a.clone(),
                                                                                                                                                                                                                                                                                                       &f))))))
                                                                                                              })
                                                                                          })
                                                                          }))
    }
}

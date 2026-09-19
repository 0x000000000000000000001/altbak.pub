pub mod PureScript_Data_DateTime_Gen {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_41690983::PureScript_Data_Date_Gen;
    use crate::module_867835b4::PureScript_Data_DateTime::Data_DateTime_DateTime;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_b10b4242::PureScript_Data_Time_Gen;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_DateTime_Gen_genDateTime() -> &dyn Any {
        static Data_DateTime_Gen_genDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Gen_genDateTime.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictMonadGen|
                                                                      {
                                                                          let Bind1 =
                                                                              Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                      Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonadGen)),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                               &&&Sharpurs_Prelude::Prim_undefined());
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&Bind1)),
                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&Bind1)),
                                                                                                                                                                                                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                                      |usd__arg1|
                                                                                                                                                                                                                                      Func1::new({
                                                                                                                                                                                                                                                     let usd__arg1
                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                         usd__arg1.clone();
                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                         |usd__arg2|
                                                                                                                                                                                                                                                         &LrcPtr::new(Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                             usd__arg2.clone()))
                                                                                                                                                                                                                                                 }))),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Gen::Data_Date_Gen_genDate(),
                                                                                                                                                                                                                    dictMonadGen))),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Gen::Data_Time_Gen_genTime(),
                                                                                                                                              dictMonadGen))
                                                                      }))
    }
}

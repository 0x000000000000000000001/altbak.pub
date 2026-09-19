pub mod PureScript_Control_Alternative {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty as empty_1;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Control_Alternative_Alternativeusd_Dict() -> &dyn Any {
        static Control_Alternative_Alternativeusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Alternative_Alternativeusd_Dict.get_or_init(||
                                                                &Func1::new(move
                                                                                |x|
                                                                                x.clone()))
    }
    pub fn Control_Alternative_guard() -> &dyn Any {
        static Control_Alternative_guard: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Alternative_guard.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictAlternative|
                                                                  {
                                                                      let Applicative0 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                  Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                           &&&Sharpurs_Prelude::Prim_undefined());
                                                                      let empty =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                      &Func1::new({
                                                                                      let Applicative0
                                                                                          =
                                                                                          Applicative0.clone();
                                                                                      let empty
                                                                                          =
                                                                                          empty.clone();
                                                                                      move
                                                                                          |v|
                                                                                          {
                                                                                              let matchValue =
                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                              match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                               &matchValue)
                                                                                                  {
                                                                                                  0_i32
                                                                                                  =>
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                      &&&Applicative0),
                                                                                                                                   &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                  _
                                                                                                  =>
                                                                                                  match &Sharpurs_Prelude::_007cLitBool_007c__007c(false,
                                                                                                                                                   &matchValue)
                                                                                                      {
                                                                                                      0_i32
                                                                                                      =>
                                                                                                      &empty,
                                                                                                      _
                                                                                                      =>
                                                                                                      panic!("{}",
                                                                                                             LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Control.Alternative.fs"),
                                  Data1: 9_i32,
                                  Data2: 447_i32,}).get_Message(),),
                                                                                                  },
                                                                                              }
                                                                                          }
                                                                                  })
                                                                  }))
    }
    pub fn Control_Alternative_alternativeArray() -> &dyn Any {
        static Control_Alternative_alternativeArray: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Alternative_alternativeArray.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                              &&&add(string("Applicative0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Control_Applicative::Control_Applicative_applicativeArray()),
                                                                                                     add(string("Plus1"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused_1|
                                                                                                                          &PureScript_Control_Plus::Control_Plus_plusArray()),
                                                                                                         empty_1::<string,
                                                                                                                   &dyn Any>()))))
    }
}

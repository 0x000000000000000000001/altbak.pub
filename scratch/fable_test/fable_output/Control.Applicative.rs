pub mod PureScript_Control_Applicative {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::NativeArray_::new_array;
    use fable_library_rust::String_::string;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Control_Applicative_Applicativeusd_Dict() -> &dyn Any {
        static Control_Applicative_Applicativeusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Applicative_Applicativeusd_Dict.get_or_init(||
                                                                &Func1::new(move
                                                                                |x|
                                                                                x.clone()))
    }
    pub fn Control_Applicative_pure() -> &dyn Any {
        static Control_Applicative_pure: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Applicative_pure.get_or_init(||
                                                 &Func1::new(move |dict|
                                                                 find(string("pure"),
                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Applicative_unless() -> &dyn Any {
        static Control_Applicative_unless: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Applicative_unless.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictApplicative|
                                                                   &Func1::new({
                                                                                   let dictApplicative
                                                                                       =
                                                                                       dictApplicative.clone();
                                                                                   move
                                                                                       |v|
                                                                                       &Func1::new({
                                                                                                       let v
                                                                                                           =
                                                                                                           v.clone();
                                                                                                       move
                                                                                                           |v1|
                                                                                                           {
                                                                                                               let matchValue =
                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                               let matchValue_1 =
                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                               match &Sharpurs_Prelude::_007cLitBool_007c__007c(false,
                                                                                                                                                                &matchValue)
                                                                                                                   {
                                                                                                                   0_i32
                                                                                                                   =>
                                                                                                                   m.clone(),
                                                                                                                   _
                                                                                                                   =>
                                                                                                                   match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                    &matchValue)
                                                                                                                       {
                                                                                                                       0_i32
                                                                                                                       =>
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                           &&&dictApplicative),
                                                                                                                                                        &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                       _
                                                                                                                       =>
                                                                                                                       panic!("{}",
                                                                                                                              LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Control.Applicative.fs"),
                                  Data1: 11_i32,
                                  Data2: 122_i32,}).get_Message(),),
                                                                                                                   },
                                                                                                               }
                                                                                                           }
                                                                                                   })
                                                                               })))
    }
    pub fn Control_Applicative_when() -> &dyn Any {
        static Control_Applicative_when: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Applicative_when.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictApplicative|
                                                                 &Func1::new({
                                                                                 let dictApplicative
                                                                                     =
                                                                                     dictApplicative.clone();
                                                                                 move
                                                                                     |v|
                                                                                     &Func1::new({
                                                                                                     let v
                                                                                                         =
                                                                                                         v.clone();
                                                                                                     move
                                                                                                         |v1|
                                                                                                         {
                                                                                                             let matchValue =
                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                             let matchValue_1 =
                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                             match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                              &matchValue)
                                                                                                                 {
                                                                                                                 0_i32
                                                                                                                 =>
                                                                                                                 m.clone(),
                                                                                                                 _
                                                                                                                 =>
                                                                                                                 match &Sharpurs_Prelude::_007cLitBool_007c__007c(false,
                                                                                                                                                                  &matchValue)
                                                                                                                     {
                                                                                                                     0_i32
                                                                                                                     =>
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                         &&&dictApplicative),
                                                                                                                                                      &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                     _
                                                                                                                     =>
                                                                                                                     panic!("{}",
                                                                                                                            LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Control.Applicative.fs"),
                                  Data1: 13_i32,
                                  Data2: 120_i32,}).get_Message(),),
                                                                                                                 },
                                                                                                             }
                                                                                                         }
                                                                                                 })
                                                                             })))
    }
    pub fn Control_Applicative_liftA1() -> &dyn Any {
        static Control_Applicative_liftA1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Applicative_liftA1.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictApplicative|
                                                                   {
                                                                       let Apply0 =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                   Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                       &Func1::new({
                                                                                       let Apply0
                                                                                           =
                                                                                           Apply0.clone();
                                                                                       let dictApplicative
                                                                                           =
                                                                                           dictApplicative.clone();
                                                                                       move
                                                                                           |f|
                                                                                           &Func1::new({
                                                                                                           let f
                                                                                                               =
                                                                                                               f.clone();
                                                                                                           move
                                                                                                               |a|
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                      &&&Apply0),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                         &&&dictApplicative),
                                                                                                                                                                                                                      &&&f)),
                                                                                                                                                a)
                                                                                                       })
                                                                                   })
                                                                   }))
    }
    pub fn Control_Applicative_applicativeProxy() -> &dyn Any {
        static Control_Applicative_applicativeProxy: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Applicative_applicativeProxy.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                              &&&add(string("pure"),
                                                                                                     &&Func1::new(move
                                                                                                                      |v|
                                                                                                                      &LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                     add(string("Apply0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Control_Apply::Control_Apply_applyProxy()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
    pub fn Control_Applicative_applicativeFn() -> &dyn Any {
        static Control_Applicative_applicativeFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Applicative_applicativeFn.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                           &&&add(string("pure"),
                                                                                                  &&Func1::new(move
                                                                                                                   |x|
                                                                                                                   &Func1::new({
                                                                                                                                   let x
                                                                                                                                       =
                                                                                                                                       x.clone();
                                                                                                                                   move
                                                                                                                                       |v|
                                                                                                                                       &x
                                                                                                                               })),
                                                                                                  add(string("Apply0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Control_Apply::Control_Apply_applyFn()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))
    }
    pub fn Control_Applicative_applicativeArray() -> &dyn Any {
        static Control_Applicative_applicativeArray: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Applicative_applicativeArray.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                              &&&add(string("pure"),
                                                                                                     &&Func1::new(move
                                                                                                                      |x|
                                                                                                                      &new_array(&[x.clone()])),
                                                                                                     add(string("Apply0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Control_Apply::Control_Apply_applyArray()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))
    }
}

pub mod PureScript_Data_Show {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_38c1e8e1::PureScript_Data_Void;
    use crate::module_312cfc22::PureScript_Record_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub mod Data_Show_FFI {
        use super::*;
        use fable_library_rust::NativeArray_::count;
        use fable_library_rust::String_::append;
        use fable_library_rust::String_::ofChar;
        use fable_library_rust::String_::toString;
        pub fn showIntImpl(x: &dyn Any) -> &dyn Any {
            &toString(Sharpurs_Prelude::unbox(x))
        }
        pub fn showNumberImpl(x: &dyn Any) -> &dyn Any {
            &toString(Sharpurs_Prelude::unbox(x))
        }
        pub fn showStringImpl(x: &dyn Any) -> &dyn Any {
            &append(append(string("\""), Sharpurs_Prelude::unbox(x)),
                    string("\""))
        }
        pub fn showCharImpl<a: Clone + 'static>(c: a) -> string {
            ofChar(Sharpurs_Prelude::unbox(&&c))
        }
        pub fn showArrayImpl(f: &dyn Any, xs: &dyn Any) -> &dyn Any {
            let xs_ = Sharpurs_Prelude::unbox(xs);
            let res: MutCell<string> = MutCell::new(string("["));
            for i in 0_i32..=count(xs_.clone()) - 1_i32 {
                res.set(append(res.get(),
                               Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(f,
                                                                                         &xs_[i].clone()))));
                if i < count(xs_.clone()) - 1_i32 {
                    res.set(append(res.get(), string(",")));
                }
            }
            res.set(append(res.get(), string("]")));
            &res.get()
        }
    }
    pub fn Data_Show_showArrayImpl() -> &dyn Any {
        static Data_Show_showArrayImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showArrayImpl.get_or_init(||
                                                &Func1::new(move |arg0|
                                                                &Func1::new({
                                                                                let arg0
                                                                                    =
                                                                                    arg0.clone();
                                                                                move
                                                                                    |arg1|
                                                                                    &PureScript_Data_Show::Data_Show_FFI::showArrayImpl(&Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                        &Sharpurs_Prelude::unbox(arg1))
                                                                            })))
    }
    pub fn Data_Show_showCharImpl() -> &dyn Any {
        static Data_Show_showCharImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showCharImpl.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               &PureScript_Data_Show::Data_Show_FFI::showCharImpl(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_Show_showIntImpl() -> &dyn Any {
        static Data_Show_showIntImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showIntImpl.get_or_init(||
                                              &Func1::new(move |arg0|
                                                              &PureScript_Data_Show::Data_Show_FFI::showIntImpl(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_Show_showNumberImpl() -> &dyn Any {
        static Data_Show_showNumberImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showNumberImpl.get_or_init(||
                                                 &Func1::new(move |arg0|
                                                                 &PureScript_Data_Show::Data_Show_FFI::showNumberImpl(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_Show_showStringImpl() -> &dyn Any {
        static Data_Show_showStringImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showStringImpl.get_or_init(||
                                                 &Func1::new(move |arg0|
                                                                 &PureScript_Data_Show::Data_Show_FFI::showStringImpl(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_Show_ShowRecordFieldsusd_Dict() -> &dyn Any {
        static Data_Show_ShowRecordFieldsusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_ShowRecordFieldsusd_Dict.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Data_Show_Showusd_Dict() -> &dyn Any {
        static Data_Show_Showusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_Showusd_Dict.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Show_showVoid() -> &dyn Any {
        static Data_Show_showVoid: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showVoid.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                            &&&add(string("show"),
                                                                                   &&PureScript_Data_Void::Data_Void_absurd(),
                                                                                   empty::<string,
                                                                                           &dyn Any>())))
    }
    pub fn Data_Show_showUnit() -> &dyn Any {
        static Data_Show_showUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showUnit.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                            &&&add(string("show"),
                                                                                   &&Func1::new(move
                                                                                                    |v|
                                                                                                    &string("unit")),
                                                                                   empty::<string,
                                                                                           &dyn Any>())))
    }
    pub fn Data_Show_showString() -> &dyn Any {
        static Data_Show_showString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showString.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                              &&&add(string("show"),
                                                                                     &&PureScript_Data_Show::Data_Show_showStringImpl(),
                                                                                     empty::<string,
                                                                                             &dyn Any>())))
    }
    pub fn Data_Show_showRecordFieldsNil() -> &dyn Any {
        static Data_Show_showRecordFieldsNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showRecordFieldsNil.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_ShowRecordFieldsusd_Dict(),
                                                                                       &&&add(string("showRecordFields"),
                                                                                              &&Func1::new(move
                                                                                                               |v|
                                                                                                               &Func1::new(move
                                                                                                                               |v1|
                                                                                                                               &string(""))),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Show_showRecordFields() -> &dyn Any {
        static Data_Show_showRecordFields: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showRecordFields.get_or_init(||
                                                   &Func1::new(move |dict|
                                                                   find(string("showRecordFields"),
                                                                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Show_showRecord() -> &dyn Any {
        static Data_Show_showRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showRecord.get_or_init(||
                                             &Func1::new(move |usd__unused|
                                                             &Func1::new(move
                                                                             |usd__unused_1|
                                                                             &Func1::new(move
                                                                                             |dictShowRecordFields|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                              &&&add(string("show"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let dictShowRecordFields
                                                                                                                                                          =
                                                                                                                                                          dictShowRecordFields.clone();
                                                                                                                                                      move
                                                                                                                                                          |record|
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                              &&&string("{")),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_showRecordFields(),
                                                                                                                                                                                                                                                                                                                                                                          &&&dictShowRecordFields),
                                                                                                                                                                                                                                                                                                                                       &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                                                                                    record)),
                                                                                                                                                                                                                              &&&string("}")))
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>()))))))
    }
    pub fn Data_Show_showProxy() -> &dyn Any {
        static Data_Show_showProxy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showProxy.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                             &&&add(string("show"),
                                                                                    &&Func1::new(move
                                                                                                     |v|
                                                                                                     &string("Proxy")),
                                                                                    empty::<string,
                                                                                            &dyn Any>())))
    }
    pub fn Data_Show_showNumber() -> &dyn Any {
        static Data_Show_showNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showNumber.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                              &&&add(string("show"),
                                                                                     &&PureScript_Data_Show::Data_Show_showNumberImpl(),
                                                                                     empty::<string,
                                                                                             &dyn Any>())))
    }
    pub fn Data_Show_showInt() -> &dyn Any {
        static Data_Show_showInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showInt.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                           &&&add(string("show"),
                                                                                  &&PureScript_Data_Show::Data_Show_showIntImpl(),
                                                                                  empty::<string,
                                                                                          &dyn Any>())))
    }
    pub fn Data_Show_showChar() -> &dyn Any {
        static Data_Show_showChar: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showChar.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                            &&&add(string("show"),
                                                                                   &&PureScript_Data_Show::Data_Show_showCharImpl(),
                                                                                   empty::<string,
                                                                                           &dyn Any>())))
    }
    pub fn Data_Show_showBoolean() -> &dyn Any {
        static Data_Show_showBoolean: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showBoolean.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                               &&&add(string("show"),
                                                                                      &&Func1::new(move
                                                                                                       |v|
                                                                                                       {
                                                                                                           let matchValue =
                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                           match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                            &matchValue)
                                                                                                               {
                                                                                                               0_i32
                                                                                                               =>
                                                                                                               &string("true"),
                                                                                                               _
                                                                                                               =>
                                                                                                               match &Sharpurs_Prelude::_007cLitBool_007c__007c(false,
                                                                                                                                                                &matchValue)
                                                                                                                   {
                                                                                                                   0_i32
                                                                                                                   =>
                                                                                                                   &string("false"),
                                                                                                                   _
                                                                                                                   =>
                                                                                                                   panic!("{}",
                                                                                                                          LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Show.fs"),
                                  Data1: 54_i32,
                                  Data2: 145_i32,}).get_Message(),),
                                                                                                               },
                                                                                                           }
                                                                                                       }),
                                                                                      empty::<string,
                                                                                              &dyn Any>())))
    }
    pub fn Data_Show_show() -> &dyn Any {
        static Data_Show_show: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Show_show.get_or_init(||
                                       &Func1::new(move |dict|
                                                       find(string("show"),
                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_Show_showArray() -> &dyn Any {
        static Data_Show_showArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showArray.get_or_init(||
                                            &Func1::new(move |dictShow|
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                             &&&add(string("show"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_showArrayImpl(),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                         dictShow)),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Show_showRecordFieldsCons() -> &dyn Any {
        static Data_Show_showRecordFieldsCons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showRecordFieldsCons.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictIsSymbol|
                                                                       &Func1::new({
                                                                                       let dictIsSymbol
                                                                                           =
                                                                                           dictIsSymbol.clone();
                                                                                       move
                                                                                           |dictShowRecordFields|
                                                                                           &Func1::new({
                                                                                                           let dictShowRecordFields
                                                                                                               =
                                                                                                               dictShowRecordFields.clone();
                                                                                                           move
                                                                                                               |dictShow|
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_ShowRecordFieldsusd_Dict(),
                                                                                                                                                &&&add(string("showRecordFields"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let dictShow
                                                                                                                                                                            =
                                                                                                                                                                            dictShow.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |v|
                                                                                                                                                                            &Func1::new(move
                                                                                                                                                                                            |record|
                                                                                                                                                                                            {
                                                                                                                                                                                                let tail =
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_showRecordFields(),
                                                                                                                                                                                                                                                                                                           &&&dictShowRecordFields),
                                                                                                                                                                                                                                                                        &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor)),
                                                                                                                                                                                                                                     record);
                                                                                                                                                                                                let key =
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                                        &&&dictIsSymbol),
                                                                                                                                                                                                                                     &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                                                let focus =
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeGet(),
                                                                                                                                                                                                                                                                        &&&key),
                                                                                                                                                                                                                                     record);
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                    &&&string(" ")),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                       &&&key),
                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                          &&&string(": ")),
                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&dictShow),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&focus)),
                                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                                                                                                &&&string(",")),
                                                                                                                                                                                                                                                                                                                                                                             &&&tail)))))
                                                                                                                                                                                            })
                                                                                                                                                                    }),
                                                                                                                                                       empty::<string,
                                                                                                                                                               &dyn Any>()))
                                                                                                       })
                                                                                   })))
    }
    pub fn Data_Show_showRecordFieldsConsNil() -> &dyn Any {
        static Data_Show_showRecordFieldsConsNil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Show_showRecordFieldsConsNil.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictIsSymbol|
                                                                          &Func1::new({
                                                                                          let dictIsSymbol
                                                                                              =
                                                                                              dictIsSymbol.clone();
                                                                                          move
                                                                                              |dictShow|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_ShowRecordFieldsusd_Dict(),
                                                                                                                               &&&add(string("showRecordFields"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let dictShow
                                                                                                                                                           =
                                                                                                                                                           dictShow.clone();
                                                                                                                                                       move
                                                                                                                                                           |v|
                                                                                                                                                           &Func1::new(move
                                                                                                                                                                           |record|
                                                                                                                                                                           {
                                                                                                                                                                               let key =
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                       &&&dictIsSymbol),
                                                                                                                                                                                                                    &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor));
                                                                                                                                                                               let focus =
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Record_Unsafe::Record_Unsafe_unsafeGet(),
                                                                                                                                                                                                                                                       &&&key),
                                                                                                                                                                                                                    record);
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                   &&&string(" ")),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                      &&&key),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                         &&&string(": ")),
                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                                                  &&&dictShow),
                                                                                                                                                                                                                                                                                                                                                                                               &&&focus)),
                                                                                                                                                                                                                                                                                                                         &&&string(" ")))))
                                                                                                                                                                           })
                                                                                                                                                   }),
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>()))
                                                                                      })))
    }
}

pub mod PureScript_Data_Reflectable {
    use super::*;
    use fable_library_rust::Interfaces_::System::IComparable;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_48ec9431::PureScript_Type_Proxy::Type_Proxy_Proxy;
    pub mod Data_Reflectable_FFI {
        use super::*;
        pub fn unsafeCoerce<a: Clone + 'static>(x: a) -> a { x }
    }
    pub fn Data_Reflectable_unsafeCoerce() -> &dyn Any {
        static Data_Reflectable_unsafeCoerce: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_unsafeCoerce.get_or_init(||
                                                      &Func1::new(move |arg0|
                                                                      &PureScript_Data_Reflectable::Data_Reflectable_FFI::unsafeCoerce(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_Reflectable_Reifiableusd_Dict() -> &dyn Any {
        static Data_Reflectable_Reifiableusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_Reifiableusd_Dict.get_or_init(||
                                                           &Func1::new(move
                                                                           |x|
                                                                           x.clone()))
    }
    pub fn Data_Reflectable_Reflectableusd_Dict() -> &dyn Any {
        static Data_Reflectable_Reflectableusd_Dict: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Reflectable_Reflectableusd_Dict.get_or_init(||
                                                             &Func1::new(move
                                                                             |x|
                                                                             x.clone()))
    }
    pub fn Data_Reflectable_reifiableString() -> &dyn Any {
        static Data_Reflectable_reifiableString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_reifiableString.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Reflectable::Data_Reflectable_Reifiableusd_Dict(),
                                                                                          &&&empty::<LrcPtr<dyn IComparable>,
                                                                                                     &dyn Any>()))
    }
    pub fn Data_Reflectable_reifiableOrdering() -> &dyn Any {
        static Data_Reflectable_reifiableOrdering: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_reifiableOrdering.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Reflectable::Data_Reflectable_Reifiableusd_Dict(),
                                                                                            &&&empty::<LrcPtr<dyn IComparable>,
                                                                                                       &dyn Any>()))
    }
    pub fn Data_Reflectable_reifiableInt() -> &dyn Any {
        static Data_Reflectable_reifiableInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_reifiableInt.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Reflectable::Data_Reflectable_Reifiableusd_Dict(),
                                                                                       &&&empty::<LrcPtr<dyn IComparable>,
                                                                                                  &dyn Any>()))
    }
    pub fn Data_Reflectable_reifiableBoolean() -> &dyn Any {
        static Data_Reflectable_reifiableBoolean: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_reifiableBoolean.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Reflectable::Data_Reflectable_Reifiableusd_Dict(),
                                                                                           &&&empty::<LrcPtr<dyn IComparable>,
                                                                                                      &dyn Any>()))
    }
    pub fn Data_Reflectable_reifyType() -> &dyn Any {
        static Data_Reflectable_reifyType: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_reifyType.get_or_init(||
                                                   &Func1::new(move
                                                                   |usd__unused|
                                                                   &Func1::new(move
                                                                                   |s|
                                                                                   &Func1::new({
                                                                                                   let s
                                                                                                       =
                                                                                                       s.clone();
                                                                                                   move
                                                                                                       |f|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&&PureScript_Data_Reflectable::Data_Reflectable_unsafeCoerce(),
                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                let f
                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                    f.clone();
                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                    |dictReflectable|
                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                     dictReflectable)
                                                                                                                                                                                                                            })),
                                                                                                                                                                           &&&add(string("reflectType"),
                                                                                                                                                                                  &&Func1::new(move
                                                                                                                                                                                                   |v|
                                                                                                                                                                                                   &s),
                                                                                                                                                                                  empty::<string,
                                                                                                                                                                                          &dyn Any>())),
                                                                                                                                        &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))
                                                                                               }))))
    }
    pub fn Data_Reflectable_reflectType() -> &dyn Any {
        static Data_Reflectable_reflectType: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Reflectable_reflectType.get_or_init(||
                                                     &Func1::new(move |dict|
                                                                     find(string("reflectType"),
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

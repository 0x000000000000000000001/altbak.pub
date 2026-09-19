pub mod PureScript_Data_Symbol {
    use super::*;
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
    pub mod Data_Symbol_FFI {
        use super::*;
        pub fn unsafeCoerce<a: Clone + 'static>(x: a) -> a { x }
    }
    pub fn Data_Symbol_unsafeCoerce() -> &dyn Any {
        static Data_Symbol_unsafeCoerce: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Symbol_unsafeCoerce.get_or_init(||
                                                 &Func1::new(move |arg0|
                                                                 &PureScript_Data_Symbol::Data_Symbol_FFI::unsafeCoerce(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_Symbol_IsSymbolusd_Dict() -> &dyn Any {
        static Data_Symbol_IsSymbolusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Symbol_IsSymbolusd_Dict.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Symbol_reifySymbol() -> &dyn Any {
        static Data_Symbol_reifySymbol: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Symbol_reifySymbol.get_or_init(||
                                                &Func1::new(move |s|
                                                                &Func1::new({
                                                                                let s
                                                                                    =
                                                                                    s.clone();
                                                                                move
                                                                                    |f|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&&PureScript_Data_Symbol::Data_Symbol_unsafeCoerce(),
                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                             let f
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 f.clone();
                                                                                                                                                                                                             move
                                                                                                                                                                                                                 |dictIsSymbol|
                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                  dictIsSymbol)
                                                                                                                                                                                                         })),
                                                                                                                                                        &&&add(string("reflectSymbol"),
                                                                                                                                                               &&Func1::new(move
                                                                                                                                                                                |v|
                                                                                                                                                                                &s),
                                                                                                                                                               empty::<string,
                                                                                                                                                                       &dyn Any>())),
                                                                                                                     &&&LrcPtr::new(Type_Proxy_Proxy::Type_Proxy_Proxyusd_Ctor))
                                                                            })))
    }
    pub fn Data_Symbol_reflectSymbol() -> &dyn Any {
        static Data_Symbol_reflectSymbol: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Symbol_reflectSymbol.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("reflectSymbol"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

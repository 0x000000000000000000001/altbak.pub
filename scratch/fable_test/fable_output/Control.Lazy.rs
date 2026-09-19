pub mod PureScript_Control_Lazy {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Lazy_Lazyusd_Dict() -> &dyn Any {
        static Control_Lazy_Lazyusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Lazy_Lazyusd_Dict.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Control_Lazy_lazyUnit() -> &dyn Any {
        static Control_Lazy_lazyUnit: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Lazy_lazyUnit.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_Lazyusd_Dict(),
                                                                               &&&add(string("defer"),
                                                                                      &&Func1::new(move
                                                                                                       |v|
                                                                                                       &PureScript_Data_Unit::Data_Unit_unit()),
                                                                                      empty::<string,
                                                                                              &dyn Any>())))
    }
    pub fn Control_Lazy_lazyFn() -> &dyn Any {
        static Control_Lazy_lazyFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Lazy_lazyFn.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_Lazyusd_Dict(),
                                                                             &&&add(string("defer"),
                                                                                    &&Func1::new(move
                                                                                                     |f|
                                                                                                     &Func1::new({
                                                                                                                     let f
                                                                                                                         =
                                                                                                                         f.clone();
                                                                                                                     move
                                                                                                                         |x|
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                             &&&PureScript_Data_Unit::Data_Unit_unit()),
                                                                                                                                                          x)
                                                                                                                 })),
                                                                                    empty::<string,
                                                                                            &dyn Any>())))
    }
    pub fn Control_Lazy_defer() -> &dyn Any {
        static Control_Lazy_defer: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Lazy_defer.get_or_init(||
                                           &Func1::new(move |dict|
                                                           find(string("defer"),
                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Lazy_fix() -> &dyn Any {
        static Control_Lazy_fix: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Lazy_fix.get_or_init(||
                                         &Func1::new(move |dictLazy|
                                                         &Func1::new({
                                                                         let dictLazy
                                                                             =
                                                                             dictLazy.clone();
                                                                         move
                                                                             |f|
                                                                             {
                                                                                 let go_2 =
                                                                                     Func0::new({
                                                                                                    let f
                                                                                                        =
                                                                                                        f.clone();
                                                                                                    let go_1
                                                                                                        =
                                                                                                        go_1.clone();
                                                                                                    move
                                                                                                        ||
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                                                            &&&dictLazy),
                                                                                                                                         &&&Func1::new({
                                                                                                                                                           let go_1
                                                                                                                                                               =
                                                                                                                                                               go_1.clone();
                                                                                                                                                           move
                                                                                                                                                               |v|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                &&&go_1.Value)
                                                                                                                                                       }))
                                                                                                });
                                                                                 let go_1 =
                                                                                     Lazy(go_2);
                                                                                 let go =
                                                                                     go_1.Value;
                                                                                 &go
                                                                             }
                                                                     })))
    }
}

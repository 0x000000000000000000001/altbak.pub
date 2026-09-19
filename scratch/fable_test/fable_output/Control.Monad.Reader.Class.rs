pub mod PureScript_Control_Monad_Reader_Class {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Reader_Class_MonadAskusd_Dict() -> &dyn Any {
        static Control_Monad_Reader_Class_MonadAskusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Class_MonadAskusd_Dict.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |x|
                                                                                    x.clone()))
    }
    pub fn Control_Monad_Reader_Class_MonadReaderusd_Dict() -> &dyn Any {
        static Control_Monad_Reader_Class_MonadReaderusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Class_MonadReaderusd_Dict.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |x|
                                                                                       x.clone()))
    }
    pub fn Control_Monad_Reader_Class_monadAskFun() -> &dyn Any {
        static Control_Monad_Reader_Class_monadAskFun:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Class_monadAskFun.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadAskusd_Dict(),
                                                                                                &&&add(string("ask"),
                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                         &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                                       add(string("Monad0"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused|
                                                                                                                            &PureScript_Control_Monad::Control_Monad_monadFn()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
    pub fn Control_Monad_Reader_Class_monadReaderFun() -> &dyn Any {
        static Control_Monad_Reader_Class_monadReaderFun:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Class_monadReaderFun.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_MonadReaderusd_Dict(),
                                                                                                   &&&add(string("local"),
                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                          add(string("MonadAsk0"),
                                                                                                              &&Func1::new(move
                                                                                                                               |usd__unused|
                                                                                                                               &PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_monadAskFun()),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>()))))
    }
    pub fn Control_Monad_Reader_Class_local() -> &dyn Any {
        static Control_Monad_Reader_Class_local: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Class_local.get_or_init(||
                                                         &Func1::new(move
                                                                         |dict|
                                                                         find(string("local"),
                                                                              Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Reader_Class_ask() -> &dyn Any {
        static Control_Monad_Reader_Class_ask: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Class_ask.get_or_init(||
                                                       &Func1::new(move |dict|
                                                                       find(string("ask"),
                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Reader_Class_asks() -> &dyn Any {
        static Control_Monad_Reader_Class_asks: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Reader_Class_asks.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictMonadAsk|
                                                                        {
                                                                            let Functor0 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Monad0"),
                                                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictMonadAsk)),
                                                                                                                                                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined());
                                                                            let ask1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Reader_Class::Control_Monad_Reader_Class_ask(),
                                                                                                                 dictMonadAsk);
                                                                            &Func1::new({
                                                                                            let Functor0
                                                                                                =
                                                                                                Functor0.clone();
                                                                                            let ask1
                                                                                                =
                                                                                                ask1.clone();
                                                                                            move
                                                                                                |f|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                       &&&Functor0),
                                                                                                                                                                    f),
                                                                                                                                 &&&ask1)
                                                                                        })
                                                                        }))
    }
}

pub mod PureScript_Control_Monad_Gen_Class {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Gen_Class_MonadGenusd_Dict() -> &dyn Any {
        static Control_Monad_Gen_Class_MonadGenusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Class_MonadGenusd_Dict.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |x|
                                                                                 x.clone()))
    }
    pub fn Control_Monad_Gen_Class_sized() -> &dyn Any {
        static Control_Monad_Gen_Class_sized: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Class_sized.get_or_init(||
                                                      &Func1::new(move |dict|
                                                                      find(string("sized"),
                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Gen_Class_resize() -> &dyn Any {
        static Control_Monad_Gen_Class_resize: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Class_resize.get_or_init(||
                                                       &Func1::new(move |dict|
                                                                       find(string("resize"),
                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Gen_Class_chooseInt() -> &dyn Any {
        static Control_Monad_Gen_Class_chooseInt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Class_chooseInt.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("chooseInt"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Gen_Class_chooseFloat() -> &dyn Any {
        static Control_Monad_Gen_Class_chooseFloat: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Control_Monad_Gen_Class_chooseFloat.get_or_init(||
                                                            &Func1::new(move
                                                                            |dict|
                                                                            find(string("chooseFloat"),
                                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Monad_Gen_Class_chooseBool() -> &dyn Any {
        static Control_Monad_Gen_Class_chooseBool: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Gen_Class_chooseBool.get_or_init(||
                                                           &Func1::new(move
                                                                           |dict|
                                                                           find(string("chooseBool"),
                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

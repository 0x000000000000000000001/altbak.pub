pub mod PureScript_Control_Monad_Cont_Class {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Monad_Cont_Class_MonadContusd_Dict() -> &dyn Any {
        static Control_Monad_Cont_Class_MonadContusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Class_MonadContusd_Dict.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |x|
                                                                                   x.clone()))
    }
    pub fn Control_Monad_Cont_Class_callCC() -> &dyn Any {
        static Control_Monad_Cont_Class_callCC: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Monad_Cont_Class_callCC.get_or_init(||
                                                        &Func1::new(move
                                                                        |dict|
                                                                        find(string("callCC"),
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

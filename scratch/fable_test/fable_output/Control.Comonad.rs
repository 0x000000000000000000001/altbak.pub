pub mod PureScript_Control_Comonad {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Comonad_Comonadusd_Dict() -> &dyn Any {
        static Control_Comonad_Comonadusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_Comonadusd_Dict.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Control_Comonad_extract() -> &dyn Any {
        static Control_Comonad_extract: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Comonad_extract.get_or_init(||
                                                &Func1::new(move |dict|
                                                                find(string("extract"),
                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

pub mod PureScript_Control_Plus {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::NativeArray_::new_empty;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Plus_Plususd_Dict() -> &dyn Any {
        static Control_Plus_Plususd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Plus_Plususd_Dict.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  x.clone()))
    }
    pub fn Control_Plus_plusArray() -> &dyn Any {
        static Control_Plus_plusArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Plus_plusArray.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                &&&add(string("empty"),
                                                                                       &&new_empty::<&dyn Any>(),
                                                                                       add(string("Alt0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Control_Alt::Control_Alt_altArray()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Control_Plus_empty() -> &dyn Any {
        static Control_Plus_empty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Plus_empty.get_or_init(||
                                           &Func1::new(move |dict|
                                                           find(string("empty"),
                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

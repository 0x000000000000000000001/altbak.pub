pub mod PureScript_Control_Alt {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Alt_Altusd_Dict() -> &dyn Any {
        static Control_Alt_Altusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Alt_Altusd_Dict.get_or_init(||
                                                &Func1::new(move |x|
                                                                x.clone()))
    }
    pub fn Control_Alt_altArray() -> &dyn Any {
        static Control_Alt_altArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Alt_altArray.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                              &&&add(string("alt"),
                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                       &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                     add(string("Functor0"),
                                                                                         &&Func1::new(move
                                                                                                          |usd__unused|
                                                                                                          &PureScript_Data_Functor::Data_Functor_functorArray()),
                                                                                         empty::<string,
                                                                                                 &dyn Any>()))))
    }
    pub fn Control_Alt_alt() -> &dyn Any {
        static Control_Alt_alt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Alt_alt.get_or_init(||
                                        &Func1::new(move |dict|
                                                        find(string("alt"),
                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

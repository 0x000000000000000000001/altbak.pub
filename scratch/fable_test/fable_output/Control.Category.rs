pub mod PureScript_Control_Category {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Control_Category_Categoryusd_Dict() -> &dyn Any {
        static Control_Category_Categoryusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Category_Categoryusd_Dict.get_or_init(||
                                                          &Func1::new(move |x|
                                                                          x.clone()))
    }
    pub fn Control_Category_identity() -> &dyn Any {
        static Control_Category_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Category_identity.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("identity"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Control_Category_categoryFn() -> &dyn Any {
        static Control_Category_categoryFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Control_Category_categoryFn.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_Categoryusd_Dict(),
                                                                                     &&&add(string("identity"),
                                                                                            &&Func1::new(move
                                                                                                             |x|
                                                                                                             x.clone()),
                                                                                            add(string("Semigroupoid0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
}

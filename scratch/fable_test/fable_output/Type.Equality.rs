pub mod PureScript_Type_Equality {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Type_Equality_TypeEqualsusd_Dict() -> &dyn Any {
        static Type_Equality_TypeEqualsusd_Dict: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Equality_TypeEqualsusd_Dict.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Type_Equality_To() -> &dyn Any {
        static Type_Equality_To: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Equality_To.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Type_Equality_From() -> &dyn Any {
        static Type_Equality_From: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Equality_From.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Type_Equality_refl() -> &dyn Any {
        static Type_Equality_refl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Equality_refl.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Type_Equality::Type_Equality_TypeEqualsusd_Dict(),
                                                                            &&&add(string("proof"),
                                                                                   &&Func1::new(move
                                                                                                    |a|
                                                                                                    a.clone()),
                                                                                   add(string("Coercible0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &Sharpurs_Prelude::Prim_undefined()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Type_Equality_proof() -> &dyn Any {
        static Type_Equality_proof: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Equality_proof.get_or_init(||
                                            &Func1::new(move |dict|
                                                            find(string("proof"),
                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Type_Equality_to() -> &dyn Any {
        static Type_Equality_to: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Equality_to.get_or_init(||
                                         &Func1::new(move |dictTypeEquals|
                                                         &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Type_Equality::Type_Equality_proof(),
                                                                                                                                                        dictTypeEquals),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Type_Equality::Type_Equality_To(),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |a|
                                                                                                                                                                          a.clone()))))))
    }
    pub fn Type_Equality_from() -> &dyn Any {
        static Type_Equality_from: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Type_Equality_from.get_or_init(||
                                           &Func1::new(move |dictTypeEquals|
                                                           &Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Type_Equality::Type_Equality_proof(),
                                                                                                                                                          dictTypeEquals),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Type_Equality::Type_Equality_From(),
                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                            |a|
                                                                                                                                                                            a.clone()))))))
    }
}

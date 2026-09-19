pub mod PureScript_Data_String_NonEmpty_CaseInsensitive {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_d71935fd::PureScript_Data_String_NonEmpty_Internal;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString()
     -> &dyn Any {
        static
         Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString.get_or_init(||
                                                                                           &Func1::new(move
                                                                                                           |x|
                                                                                                           x.clone()))
    }
    pub fn Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString()
     -> &dyn Any {
        static
         Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString.get_or_init(||
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                                &&&add(string("show"),
                                                                                                                                       &&Func1::new(move
                                                                                                                                                        |v|
                                                                                                                                                        {
                                                                                                                                                            let s =
                                                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&&string("(CaseInsensitiveNonEmptyString ")),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_showNonEmptyString()),
                                                                                                                                                                                                                                                                                                      &&&s)),
                                                                                                                                                                                                                                &&&string(")")))
                                                                                                                                                        }),
                                                                                                                                       empty::<string,
                                                                                                                                               &dyn Any>())))
    }
    pub fn Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString()
     -> &dyn Any {
        static
         Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString.get_or_init(||
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                                                   &&&add(string("Coercible0"),
                                                                                                                                          &&Func1::new(move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                          empty::<string,
                                                                                                                                                  &dyn Any>())))
    }
    pub fn Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString()
     -> &dyn Any {
        static
         Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString.get_or_init(||
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                              &&&add(string("eq"),
                                                                                                                                     &&Func1::new(move
                                                                                                                                                      |v|
                                                                                                                                                      &Func1::new({
                                                                                                                                                                      let v
                                                                                                                                                                          =
                                                                                                                                                                          v.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |v1|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                              let matchValue_1 =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_eqNonEmptyString()),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_toLower(),
                                                                                                                                                                                                                                                                                     &&&matchValue)),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_toLower(),
                                                                                                                                                                                                                                                  &&&matchValue_1))
                                                                                                                                                                          }
                                                                                                                                                                  })),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
    }
    pub fn Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString()
     -> &dyn Any {
        static
         Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString.get_or_init(||
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                               &&&add(string("compare"),
                                                                                                                                      &&Func1::new(move
                                                                                                                                                       |v|
                                                                                                                                                       &Func1::new({
                                                                                                                                                                       let v
                                                                                                                                                                           =
                                                                                                                                                                           v.clone();
                                                                                                                                                                       move
                                                                                                                                                                           |v1|
                                                                                                                                                                           {
                                                                                                                                                                               let matchValue =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                               let matchValue_1 =
                                                                                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_ordNonEmptyString()),
                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_toLower(),
                                                                                                                                                                                                                                                                                      &&&matchValue)),
                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_toLower(),
                                                                                                                                                                                                                                                   &&&matchValue_1))
                                                                                                                                                                           }
                                                                                                                                                                   })),
                                                                                                                                      add(string("Eq0"),
                                                                                                                                          &&Func1::new(move
                                                                                                                                                           |usd__unused|
                                                                                                                                                           &PureScript_Data_String_NonEmpty_CaseInsensitive::Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString()),
                                                                                                                                          empty::<string,
                                                                                                                                                  &dyn Any>()))))
    }
}

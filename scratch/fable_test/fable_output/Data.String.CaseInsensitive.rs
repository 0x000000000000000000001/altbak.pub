pub mod PureScript_Data_String_CaseInsensitive {
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
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_String_CaseInsensitive_CaseInsensitiveString() -> &dyn Any {
        static Data_String_CaseInsensitive_CaseInsensitiveString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CaseInsensitive_CaseInsensitiveString.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |x|
                                                                                          x.clone()))
    }
    pub fn Data_String_CaseInsensitive_showCaseInsensitiveString()
     -> &dyn Any {
        static Data_String_CaseInsensitive_showCaseInsensitiveString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CaseInsensitive_showCaseInsensitiveString.get_or_init(||
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                               &&&add(string("show"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |v|
                                                                                                                                       {
                                                                                                                                           let s =
                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                               &&&string("(CaseInsensitiveString ")),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Show::Data_Show_showString()),
                                                                                                                                                                                                                                                                                     &&&s)),
                                                                                                                                                                                                               &&&string(")")))
                                                                                                                                       }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>())))
    }
    pub fn Data_String_CaseInsensitive_newtypeCaseInsensitiveString()
     -> &dyn Any {
        static Data_String_CaseInsensitive_newtypeCaseInsensitiveString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CaseInsensitive_newtypeCaseInsensitiveString.get_or_init(||
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                                  &&&add(string("Coercible0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &Sharpurs_Prelude::Prim_undefined()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>())))
    }
    pub fn Data_String_CaseInsensitive_eqCaseInsensitiveString() -> &dyn Any {
        static Data_String_CaseInsensitive_eqCaseInsensitiveString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CaseInsensitive_eqCaseInsensitiveString.get_or_init(||
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
                                                                                                                                                                                                                                                                    &&&PureScript_Data_Eq::Data_Eq_eqString()),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                                                             1_i32.get_Message(),),
                                                                                                                                                                                                                                                                    &&&matchValue)),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                          1_i32.get_Message(),),
                                                                                                                                                                                                                                 &&&matchValue_1))
                                                                                                                                                         }
                                                                                                                                                 })),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>())))
    }
    pub fn Data_String_CaseInsensitive_ordCaseInsensitiveString()
     -> &dyn Any {
        static Data_String_CaseInsensitive_ordCaseInsensitiveString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_CaseInsensitive_ordCaseInsensitiveString.get_or_init(||
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
                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ord::Data_Ord_ordString()),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                                                              1_i32.get_Message(),),
                                                                                                                                                                                                                                                                     &&&matchValue)),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                           1_i32.get_Message(),),
                                                                                                                                                                                                                                  &&&matchValue_1))
                                                                                                                                                          }
                                                                                                                                                  })),
                                                                                                                     add(string("Eq0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_String_CaseInsensitive::Data_String_CaseInsensitive_eqCaseInsensitiveString()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))
    }
}

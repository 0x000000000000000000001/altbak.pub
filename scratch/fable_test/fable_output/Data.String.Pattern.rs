pub mod PureScript_Data_String_Pattern {
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
    pub fn Data_String_Pattern_Replacement() -> &dyn Any {
        static Data_String_Pattern_Replacement: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_Replacement.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Data_String_Pattern_Pattern() -> &dyn Any {
        static Data_String_Pattern_Pattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_Pattern.get_or_init(||
                                                    &Func1::new(move |x|
                                                                    x.clone()))
    }
    pub fn Data_String_Pattern_showReplacement() -> &dyn Any {
        static Data_String_Pattern_showReplacement: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_Pattern_showReplacement.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                             &&&add(string("show"),
                                                                                                    &&Func1::new(move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let s =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                             &&&string("(Replacement ")),
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
    pub fn Data_String_Pattern_showPattern() -> &dyn Any {
        static Data_String_Pattern_showPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_showPattern.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                         &&&add(string("show"),
                                                                                                &&Func1::new(move
                                                                                                                 |v|
                                                                                                                 {
                                                                                                                     let s =
                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                         &&&string("(Pattern ")),
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
    pub fn Data_String_Pattern_newtypeReplacement() -> &dyn Any {
        static Data_String_Pattern_newtypeReplacement:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_newtypeReplacement.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                &&&add(string("Coercible0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &Sharpurs_Prelude::Prim_undefined()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Data_String_Pattern_newtypePattern() -> &dyn Any {
        static Data_String_Pattern_newtypePattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_newtypePattern.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                            &&&add(string("Coercible0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &Sharpurs_Prelude::Prim_undefined()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>())))
    }
    pub fn Data_String_Pattern_eqReplacement() -> &dyn Any {
        static Data_String_Pattern_eqReplacement: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_eqReplacement.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                           &&&add(string("eq"),
                                                                                                  &&Func1::new(move
                                                                                                                   |x|
                                                                                                                   &Func1::new({
                                                                                                                                   let x
                                                                                                                                       =
                                                                                                                                       x.clone();
                                                                                                                                   move
                                                                                                                                       |y|
                                                                                                                                       {
                                                                                                                                           let matchValue =
                                                                                                                                               Sharpurs_Prelude::unbox(&&x);
                                                                                                                                           let matchValue_1 =
                                                                                                                                               Sharpurs_Prelude::unbox(y);
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Eq::Data_Eq_eqString()),
                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                            &&&matchValue_1)
                                                                                                                                       }
                                                                                                                               })),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_String_Pattern_ordReplacement() -> &dyn Any {
        static Data_String_Pattern_ordReplacement: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_ordReplacement.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                            &&&add(string("compare"),
                                                                                                   &&Func1::new(move
                                                                                                                    |x|
                                                                                                                    &Func1::new({
                                                                                                                                    let x
                                                                                                                                        =
                                                                                                                                        x.clone();
                                                                                                                                    move
                                                                                                                                        |y|
                                                                                                                                        {
                                                                                                                                            let matchValue =
                                                                                                                                                Sharpurs_Prelude::unbox(&&x);
                                                                                                                                            let matchValue_1 =
                                                                                                                                                Sharpurs_Prelude::unbox(y);
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordString()),
                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                             &&&matchValue_1)
                                                                                                                                        }
                                                                                                                                })),
                                                                                                   add(string("Eq0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_String_Pattern::Data_String_Pattern_eqReplacement()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_String_Pattern_eqPattern() -> &dyn Any {
        static Data_String_Pattern_eqPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_eqPattern.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                       &&&add(string("eq"),
                                                                                              &&Func1::new(move
                                                                                                               |x|
                                                                                                               &Func1::new({
                                                                                                                               let x
                                                                                                                                   =
                                                                                                                                   x.clone();
                                                                                                                               move
                                                                                                                                   |y|
                                                                                                                                   {
                                                                                                                                       let matchValue =
                                                                                                                                           Sharpurs_Prelude::unbox(&&x);
                                                                                                                                       let matchValue_1 =
                                                                                                                                           Sharpurs_Prelude::unbox(y);
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                              &&&PureScript_Data_Eq::Data_Eq_eqString()),
                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                        &&&matchValue_1)
                                                                                                                                   }
                                                                                                                           })),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_String_Pattern_ordPattern() -> &dyn Any {
        static Data_String_Pattern_ordPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_Pattern_ordPattern.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                        &&&add(string("compare"),
                                                                                               &&Func1::new(move
                                                                                                                |x|
                                                                                                                &Func1::new({
                                                                                                                                let x
                                                                                                                                    =
                                                                                                                                    x.clone();
                                                                                                                                move
                                                                                                                                    |y|
                                                                                                                                    {
                                                                                                                                        let matchValue =
                                                                                                                                            Sharpurs_Prelude::unbox(&&x);
                                                                                                                                        let matchValue_1 =
                                                                                                                                            Sharpurs_Prelude::unbox(y);
                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                               &&&PureScript_Data_Ord::Data_Ord_ordString()),
                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                         &&&matchValue_1)
                                                                                                                                    }
                                                                                                                            })),
                                                                                               add(string("Eq0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_String_Pattern::Data_String_Pattern_eqPattern()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
}

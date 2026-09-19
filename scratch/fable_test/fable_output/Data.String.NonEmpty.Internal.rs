pub mod PureScript_Data_String_NonEmpty_Internal {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_eea316d6::PureScript_Data_String_CodeUnits;
    use crate::module_91241446::PureScript_Data_String_Pattern;
    use crate::module_20f337b3::PureScript_Data_Symbol;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_2a7662d2::PureScript_Unsafe_Coerce;
    pub fn Data_String_NonEmpty_Internal_fromJust() -> &dyn Any {
        static Data_String_NonEmpty_Internal_fromJust:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_fromJust.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_String_NonEmpty_Internal_NonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_NonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_NonEmptyString.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |x|
                                                                                     x.clone()))
    }
    pub fn Data_String_NonEmpty_Internal_NonEmptyReplacement() -> &dyn Any {
        static Data_String_NonEmpty_Internal_NonEmptyReplacement:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_NonEmptyReplacement.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |x|
                                                                                          x.clone()))
    }
    pub fn Data_String_NonEmpty_Internal_MakeNonEmptyusd_Dict() -> &dyn Any {
        static Data_String_NonEmpty_Internal_MakeNonEmptyusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_MakeNonEmptyusd_Dict.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |x|
                                                                                           x.clone()))
    }
    pub fn Data_String_NonEmpty_Internal_toUpper() -> &dyn Any {
        static Data_String_NonEmpty_Internal_toUpper:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_toUpper.get_or_init(||
                                                              &Func1::new(move
                                                                              |v|
                                                                              {
                                                                                  let s =
                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                               1_i32.get_Message(),),
                                                                                                                                                      &&&s))
                                                                              }))
    }
    pub fn Data_String_NonEmpty_Internal_toString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_toString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_toString.get_or_init(||
                                                               &Func1::new(move
                                                                               |v|
                                                                               &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_String_NonEmpty_Internal_toLower() -> &dyn Any {
        static Data_String_NonEmpty_Internal_toLower:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_toLower.get_or_init(||
                                                              &Func1::new(move
                                                                              |v|
                                                                              {
                                                                                  let s =
                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                               1_i32.get_Message(),),
                                                                                                                                                      &&&s))
                                                                              }))
    }
    pub fn Data_String_NonEmpty_Internal_showNonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_showNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_showNonEmptyString.get_or_init(||
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                          &&&add(string("show"),
                                                                                                                 &&Func1::new(move
                                                                                                                                  |v|
                                                                                                                                  {
                                                                                                                                      let s =
                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                          &&&string("(NonEmptyString.unsafeFromString ")),
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
    pub fn Data_String_NonEmpty_Internal_showNonEmptyReplacement()
     -> &dyn Any {
        static Data_String_NonEmpty_Internal_showNonEmptyReplacement:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_showNonEmptyReplacement.get_or_init(||
                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                               &&&add(string("show"),
                                                                                                                      &&Func1::new(move
                                                                                                                                       |v|
                                                                                                                                       {
                                                                                                                                           let s =
                                                                                                                                               Sharpurs_Prelude::unbox(v);
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                               &&&string("(NonEmptyReplacement ")),
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
    pub fn Data_String_NonEmpty_Internal_semigroupNonEmptyString()
     -> &dyn Any {
        static Data_String_NonEmpty_Internal_semigroupNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_semigroupNonEmptyString.get_or_init(||
                                                                              &PureScript_Data_Semigroup::Data_Semigroup_semigroupString())
    }
    pub fn Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement()
     -> &dyn Any {
        static Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement.get_or_init(||
                                                                                   &PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_semigroupNonEmptyString())
    }
    pub fn Data_String_NonEmpty_Internal_replaceAll() -> &dyn Any {
        static Data_String_NonEmpty_Internal_replaceAll:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_replaceAll.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |pat|
                                                                                 &Func1::new({
                                                                                                 let pat
                                                                                                     =
                                                                                                     pat.clone();
                                                                                                 move
                                                                                                     |v|
                                                                                                     &Func1::new({
                                                                                                                     let v
                                                                                                                         =
                                                                                                                         v.clone();
                                                                                                                     move
                                                                                                                         |v1|
                                                                                                                         {
                                                                                                                             let matchValue =
                                                                                                                                 Sharpurs_Prelude::unbox(&&pat);
                                                                                                                             let matchValue_1 =
                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                             let matchValue_2 =
                                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                                                                1_i32.get_Message(),),
                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Replacement(),
                                                                                                                                                                                                                                                                       &&&matchValue_1)),
                                                                                                                                                                                                 &&&matchValue_2))
                                                                                                                         }
                                                                                                                 })
                                                                                             })))
    }
    pub fn Data_String_NonEmpty_Internal_replace() -> &dyn Any {
        static Data_String_NonEmpty_Internal_replace:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_replace.get_or_init(||
                                                              &Func1::new(move
                                                                              |pat|
                                                                              &Func1::new({
                                                                                              let pat
                                                                                                  =
                                                                                                  pat.clone();
                                                                                              move
                                                                                                  |v|
                                                                                                  &Func1::new({
                                                                                                                  let v
                                                                                                                      =
                                                                                                                      v.clone();
                                                                                                                  move
                                                                                                                      |v1|
                                                                                                                      {
                                                                                                                          let matchValue =
                                                                                                                              Sharpurs_Prelude::unbox(&&pat);
                                                                                                                          let matchValue_1 =
                                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                                          let matchValue_2 =
                                                                                                                              Sharpurs_Prelude::unbox(v1);
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                                                             1_i32.get_Message(),),
                                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Pattern::Data_String_Pattern_Replacement(),
                                                                                                                                                                                                                                                                    &&&matchValue_1)),
                                                                                                                                                                                              &&&matchValue_2))
                                                                                                                      }
                                                                                                              })
                                                                                          })))
    }
    pub fn Data_String_NonEmpty_Internal_prependString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_prependString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_prependString.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |s1|
                                                                                    &Func1::new({
                                                                                                    let s1
                                                                                                        =
                                                                                                        s1.clone();
                                                                                                    move
                                                                                                        |v|
                                                                                                        {
                                                                                                            let matchValue =
                                                                                                                Sharpurs_Prelude::unbox(&&s1);
                                                                                                            let matchValue_1 =
                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                &&&matchValue_1))
                                                                                                        }
                                                                                                })))
    }
    pub fn Data_String_NonEmpty_Internal_ordNonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_ordNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_ordNonEmptyString.get_or_init(||
                                                                        &PureScript_Data_Ord::Data_Ord_ordString())
    }
    pub fn Data_String_NonEmpty_Internal_ordNonEmptyReplacement()
     -> &dyn Any {
        static Data_String_NonEmpty_Internal_ordNonEmptyReplacement:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_ordNonEmptyReplacement.get_or_init(||
                                                                             &PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_ordNonEmptyString())
    }
    pub fn Data_String_NonEmpty_Internal_nonEmptyNonEmpty() -> &dyn Any {
        static Data_String_NonEmpty_Internal_nonEmptyNonEmpty:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_nonEmptyNonEmpty.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |dictIsSymbol|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_MakeNonEmptyusd_Dict(),
                                                                                                                        &&&add(string("nes"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let dictIsSymbol
                                                                                                                                                    =
                                                                                                                                                    dictIsSymbol.clone();
                                                                                                                                                move
                                                                                                                                                    |p|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Symbol::Data_Symbol_reflectSymbol(),
                                                                                                                                                                                                                                                           &&&dictIsSymbol),
                                                                                                                                                                                                                        p))
                                                                                                                                            }),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_String_NonEmpty_Internal_nes() -> &dyn Any {
        static Data_String_NonEmpty_Internal_nes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_nes.get_or_init(||
                                                          &Func1::new(move
                                                                          |dict|
                                                                          find(string("nes"),
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_String_NonEmpty_Internal_makeNonEmptyBad() -> &dyn Any {
        static Data_String_NonEmpty_Internal_makeNonEmptyBad:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_makeNonEmptyBad.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |usd__unused|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_MakeNonEmptyusd_Dict(),
                                                                                                                       &&&add(string("nes"),
                                                                                                                              &&Func1::new(move
                                                                                                                                               |v|
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                                                                                &&&string(""))),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_String_NonEmpty_Internal_localeCompare() -> &dyn Any {
        static Data_String_NonEmpty_Internal_localeCompare:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_localeCompare.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |v|
                                                                                    &Func1::new({
                                                                                                    let v
                                                                                                        =
                                                                                                        v.clone();
                                                                                                    move
                                                                                                        |v1|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                     1_i32.get_Message(),),
                                                                                                                                                                            &&&Sharpurs_Prelude::unbox(&&v)),
                                                                                                                                         &&&Sharpurs_Prelude::unbox(v1))
                                                                                                })))
    }
    pub fn Data_String_NonEmpty_Internal_liftS() -> &dyn Any {
        static Data_String_NonEmpty_Internal_liftS: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_liftS.get_or_init(||
                                                            &Func1::new(move
                                                                            |f|
                                                                            &Func1::new({
                                                                                            let f
                                                                                                =
                                                                                                f.clone();
                                                                                            move
                                                                                                |v|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(&&f),
                                                                                                                                 &&&Sharpurs_Prelude::unbox(v))
                                                                                        })))
    }
    pub fn Data_String_NonEmpty_Internal_startsWith() -> &dyn Any {
        static Data_String_NonEmpty_Internal_startsWith:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_startsWith.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                     &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_liftS()),
                                                                                                  &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_startsWith()))
    }
    pub fn Data_String_NonEmpty_Internal_joinWith1() -> &dyn Any {
        static Data_String_NonEmpty_Internal_joinWith1:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_joinWith1.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictFoldable1|
                                                                                {
                                                                                    let Foldable0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                Sharpurs_Prelude::unbox(dictFoldable1)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    &Func1::new({
                                                                                                    let Foldable0
                                                                                                        =
                                                                                                        Foldable0.clone();
                                                                                                    move
                                                                                                        |v|
                                                                                                        {
                                                                                                            let splice =
                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString()),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_intercalate(),
                                                                                                                                                                                                                                                      &&&Foldable0),
                                                                                                                                                                                                                   &&&PureScript_Data_Monoid::Data_Monoid_monoidString()),
                                                                                                                                                                                &&&splice))
                                                                                                        }
                                                                                                })
                                                                                }))
    }
    pub fn Data_String_NonEmpty_Internal_joinWith() -> &dyn Any {
        static Data_String_NonEmpty_Internal_joinWith:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_joinWith.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictFoldable|
                                                                               &Func1::new({
                                                                                               let dictFoldable
                                                                                                   =
                                                                                                   dictFoldable.clone();
                                                                                               move
                                                                                                   |splice|
                                                                                                   {
                                                                                                       let coe =
                                                                                                           &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce();
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_intercalate(),
                                                                                                                                                                                                                                                                                    &&&dictFoldable),
                                                                                                                                                                                                                                                 &&&PureScript_Data_Monoid::Data_Monoid_monoidString()),
                                                                                                                                                                                                              splice)),
                                                                                                                                        &&&coe)
                                                                                                   }
                                                                                           })))
    }
    pub fn Data_String_NonEmpty_Internal_join1With() -> &dyn Any {
        static Data_String_NonEmpty_Internal_join1With:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_join1With.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictFoldable1|
                                                                                {
                                                                                    let Foldable0 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                Sharpurs_Prelude::unbox(dictFoldable1)),
                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined());
                                                                                    &Func1::new({
                                                                                                    let Foldable0
                                                                                                        =
                                                                                                        Foldable0.clone();
                                                                                                    move
                                                                                                        |splice|
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                            &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString()),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_joinWith(),
                                                                                                                                                                                                               &&&Foldable0),
                                                                                                                                                                            splice))
                                                                                                })
                                                                                }))
    }
    pub fn Data_String_NonEmpty_Internal_fromString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_fromString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_fromString.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |v|
                                                                                 {
                                                                                     let matchValue =
                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                     match &Sharpurs_Prelude::_007cLitString_007c__007c(string(""),
                                                                                                                                        &matchValue)
                                                                                         {
                                                                                         0_i32
                                                                                         =>
                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                         _
                                                                                         =>
                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                                                                                 &&&matchValue))),
                                                                                     }
                                                                                 }))
    }
    pub fn Data_String_NonEmpty_Internal_stripPrefix() -> &dyn Any {
        static Data_String_NonEmpty_Internal_stripPrefix:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_stripPrefix.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |pat|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_composeKleisliFlipped(),
                                                                                                                                                                                         &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                      &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_liftS(),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_stripPrefix(),
                                                                                                                                                                                         pat)))))
    }
    pub fn Data_String_NonEmpty_Internal_stripSuffix() -> &dyn Any {
        static Data_String_NonEmpty_Internal_stripSuffix:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_stripSuffix.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |pat|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_composeKleisliFlipped(),
                                                                                                                                                                                         &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                      &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_liftS(),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_stripSuffix(),
                                                                                                                                                                                         pat)))))
    }
    pub fn Data_String_NonEmpty_Internal_trim() -> &dyn Any {
        static Data_String_NonEmpty_Internal_trim: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_trim.get_or_init(||
                                                           &Func1::new(move
                                                                           |v|
                                                                           {
                                                                               let s =
                                                                                   Sharpurs_Prelude::unbox(v);
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                            1_i32.get_Message(),),
                                                                                                                                                   &&&s))
                                                                           }))
    }
    pub fn Data_String_NonEmpty_Internal_unsafeFromString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_unsafeFromString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_unsafeFromString.get_or_init(||
                                                                       &Func1::new(move
                                                                                       |usd__unused|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                           &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromJust()),
                                                                                                                        &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString())))
    }
    pub fn Data_String_NonEmpty_Internal_eqNonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_eqNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_eqNonEmptyString.get_or_init(||
                                                                       &PureScript_Data_Eq::Data_Eq_eqString())
    }
    pub fn Data_String_NonEmpty_Internal_eqNonEmptyReplacement() -> &dyn Any {
        static Data_String_NonEmpty_Internal_eqNonEmptyReplacement:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_eqNonEmptyReplacement.get_or_init(||
                                                                            &PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_eqNonEmptyString())
    }
    pub fn Data_String_NonEmpty_Internal_endsWith() -> &dyn Any {
        static Data_String_NonEmpty_Internal_endsWith:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_endsWith.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                   &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_liftS()),
                                                                                                &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_endsWith()))
    }
    pub fn Data_String_NonEmpty_Internal_contains() -> &dyn Any {
        static Data_String_NonEmpty_Internal_contains:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_contains.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                   &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_liftS()),
                                                                                                &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_contains()))
    }
    pub fn Data_String_NonEmpty_Internal_appendString() -> &dyn Any {
        static Data_String_NonEmpty_Internal_appendString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_Internal_appendString.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |v|
                                                                                   &Func1::new({
                                                                                                   let v
                                                                                                       =
                                                                                                       v.clone();
                                                                                                   move
                                                                                                       |s2|
                                                                                                       {
                                                                                                           let matchValue =
                                                                                                               Sharpurs_Prelude::unbox(&&v);
                                                                                                           let matchValue_1 =
                                                                                                               Sharpurs_Prelude::unbox(s2);
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString(),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                     &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                  &&&matchValue),
                                                                                                                                                                               &&&matchValue_1))
                                                                                                       }
                                                                                               })))
    }
}

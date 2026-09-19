pub mod PureScript_Data_String_NonEmpty_CodeUnits {
    use super::*;
    use fable_library_rust::Array_::equals;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::NativeArray_::count;
    use fable_library_rust::NativeArray_::new_empty;
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_34a61018::PureScript_Data_Array_NonEmpty;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_2563115d::PureScript_Data_Semigroup_Foldable;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_eea316d6::PureScript_Data_String_CodeUnits;
    use crate::module_d71935fd::PureScript_Data_String_NonEmpty_Internal;
    use crate::module_b0fd79e4::PureScript_Data_String_Unsafe;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_String_NonEmpty_CodeUnits_toNonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_toNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_toNonEmptyString.get_or_init(||
                                                                        &PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString())
    }
    pub fn Data_String_NonEmpty_CodeUnits_snoc() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_snoc: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_snoc.get_or_init(||
                                                            &Func1::new(move
                                                                            |c|
                                                                            &Func1::new({
                                                                                            let c
                                                                                                =
                                                                                                c.clone();
                                                                                            move
                                                                                                |s|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                       s),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_singleton(),
                                                                                                                                                                                                       &&&c)))
                                                                                        })))
    }
    pub fn Data_String_NonEmpty_CodeUnits_singleton() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_singleton:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_singleton.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                     &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString()),
                                                                                                  &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_singleton()))
    }
    pub fn Data_String_NonEmpty_CodeUnits_liftS() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_liftS: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_liftS.get_or_init(||
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
    pub fn Data_String_NonEmpty_CodeUnits_takeWhile() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_takeWhile:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_takeWhile.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |f|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                     &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString()),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS(),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_takeWhile(),
                                                                                                                                                                                        f)))))
    }
    pub fn Data_String_NonEmpty_CodeUnits_lastIndexOf_prime() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_lastIndexOf_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_lastIndexOf_prime.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |pat|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_lastIndexOf_prime(),
                                                                                                                                                             pat))))
    }
    pub fn Data_String_NonEmpty_CodeUnits_lastIndexOf() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_lastIndexOf:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_lastIndexOf.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                       &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS()),
                                                                                                    &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_lastIndexOf()))
    }
    pub fn Data_String_NonEmpty_CodeUnits_indexOf_prime() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_indexOf_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_indexOf_prime.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |pat|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                         &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS()),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_indexOf_prime(),
                                                                                                                                                         pat))))
    }
    pub fn Data_String_NonEmpty_CodeUnits_indexOf() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_indexOf:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_indexOf.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                   &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS()),
                                                                                                &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_indexOf()))
    }
    pub fn Data_String_NonEmpty_CodeUnits_fromNonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_fromNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_fromNonEmptyString.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |v|
                                                                                          &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_String_NonEmpty_CodeUnits_length() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_length:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_length.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                  &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length()),
                                                                                               &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString()))
    }
    pub fn Data_String_NonEmpty_CodeUnits_splitAt() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_splitAt:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_splitAt.get_or_init(||
                                                               &Func1::new(move
                                                                               |i|
                                                                               &Func1::new({
                                                                                               let i
                                                                                                   =
                                                                                                   i.clone();
                                                                                               move
                                                                                                   |nes|
                                                                                                   {
                                                                                                       let matchValue =
                                                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_splitAt(),
                                                                                                                                                                                                         &&&i),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString(),
                                                                                                                                                                                                         nes)));
                                                                                                       {
                                                                                                           let activePatternResult =
                                                                                                               Sharpurs_Prelude::_007cHasProp_007c__007c(string("before"),
                                                                                                                                                         &matchValue);
                                                                                                           if activePatternResult.is_some()
                                                                                                              {
                                                                                                               let activePatternResult_1 =
                                                                                                                   Sharpurs_Prelude::_007cHasProp_007c__007c(string("after"),
                                                                                                                                                             &matchValue);
                                                                                                               if activePatternResult_1.is_some()
                                                                                                                  {
                                                                                                                   let after =
                                                                                                                       getValue(activePatternResult_1);
                                                                                                                   let before =
                                                                                                                       getValue(activePatternResult);
                                                                                                                   &add(string("before"),
                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString(),
                                                                                                                                                          &&&before),
                                                                                                                        add(string("after"),
                                                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString(),
                                                                                                                                                              &&&after),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>()))
                                                                                                               } else {
                                                                                                                   panic!("{}",
                                                                                                                          LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.NonEmpty.CodeUnits.fs"),
                                  Data1: 29_i32,
                                  Data2: 320_i32,}).get_Message(),)
                                                                                                               }
                                                                                                           } else {
                                                                                                               panic!("{}",
                                                                                                                      LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.NonEmpty.CodeUnits.fs"),
                                  Data1: 29_i32,
                                  Data2: 320_i32,}).get_Message(),)
                                                                                                           }
                                                                                                       }
                                                                                                   }
                                                                                           })))
    }
    pub fn Data_String_NonEmpty_CodeUnits_take() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_take: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_take.get_or_init(||
                                                            &Func1::new(move
                                                                            |i|
                                                                            &Func1::new({
                                                                                            let i
                                                                                                =
                                                                                                i.clone();
                                                                                            move
                                                                                                |nes|
                                                                                                {
                                                                                                    let s =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString(),
                                                                                                                                         nes);
                                                                                                    let matchValue =
                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                        &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                     &&&i),
                                                                                                                                                                  &&&1_i32));
                                                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                     &matchValue)
                                                                                                        {
                                                                                                        0_i32
                                                                                                        =>
                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                        _
                                                                                                        =>
                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString(),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_take(),
                                                                                                                                                                                                                                                                      &&&i),
                                                                                                                                                                                                                                   &&&s)))),
                                                                                                    }
                                                                                                }
                                                                                        })))
    }
    pub fn Data_String_NonEmpty_CodeUnits_takeRight() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_takeRight:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_takeRight.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |i|
                                                                                 &Func1::new({
                                                                                                 let i
                                                                                                     =
                                                                                                     i.clone();
                                                                                                 move
                                                                                                     |nes|
                                                                                                     {
                                                                                                         let s =
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString(),
                                                                                                                                              nes);
                                                                                                         let matchValue =
                                                                                                             Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                          &&&i),
                                                                                                                                                                       &&&1_i32));
                                                                                                         match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                          &matchValue)
                                                                                                             {
                                                                                                             0_i32
                                                                                                             =>
                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                             _
                                                                                                             =>
                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString(),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_takeRight(),
                                                                                                                                                                                                                                                                           &&&i),
                                                                                                                                                                                                                                        &&&s)))),
                                                                                                         }
                                                                                                     }
                                                                                             })))
    }
    pub fn Data_String_NonEmpty_CodeUnits_toChar() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_toChar:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_toChar.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                  &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_toChar()),
                                                                                               &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString()))
    }
    pub fn Data_String_NonEmpty_CodeUnits_toCharArray() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_toCharArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_toCharArray.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                       &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_toCharArray()),
                                                                                                    &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString()))
    }
    pub fn Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray.get_or_init(||
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                    |usd__unused|
                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                     &&&Sharpurs_Prelude::Prim_undefined())))),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                  &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_fromArray()),
                                                                                                                                               &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toCharArray())))
    }
    pub fn Data_String_NonEmpty_CodeUnits_uncons() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_uncons:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_uncons.get_or_init(||
                                                              &Func1::new(move
                                                                              |nes|
                                                                              {
                                                                                  let s =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString(),
                                                                                                                       nes);
                                                                                  &add(string("head"),
                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_Unsafe::Data_String_Unsafe_charAt(),
                                                                                                                                                            &&&0_i32),
                                                                                                                         &&&s),
                                                                                       add(string("tail"),
                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString(),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_drop(),
                                                                                                                                                                                                   &&&1_i32),
                                                                                                                                                                &&&s)),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))
                                                                              }))
    }
    pub fn Data_String_NonEmpty_CodeUnits_fromFoldable1() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_fromFoldable1:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_fromFoldable1.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictFoldable1|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                            dictFoldable1),
                                                                                                                                                         &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_semigroupNonEmptyString()),
                                                                                                                      &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_singleton())))
    }
    pub fn Data_String_NonEmpty_CodeUnits_fromCharArray() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_fromCharArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_fromCharArray.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |v|
                                                                                     {
                                                                                         let matchValue =
                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                         if if !equals(matchValue.clone(),
                                                                                                       new_empty::<&dyn Any>())
                                                                                               {
                                                                                                count(matchValue.clone())
                                                                                                    ==
                                                                                                    0_i32
                                                                                            } else {
                                                                                                false
                                                                                            }
                                                                                            {
                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                         } else {
                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString(),
                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_fromCharArray(),
                                                                                                                                                                                                                        &&&matchValue))))
                                                                                         }
                                                                                     }))
    }
    pub fn Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray()
     -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray.get_or_init(||
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                      |usd__unused|
                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined())))),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                    &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromCharArray()),
                                                                                                                                                 &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())))
    }
    pub fn Data_String_NonEmpty_CodeUnits_dropWhile() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_dropWhile:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_dropWhile.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |f|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                     &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString()),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS(),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_dropWhile(),
                                                                                                                                                                                        f)))))
    }
    pub fn Data_String_NonEmpty_CodeUnits_dropRight() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_dropRight:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_dropRight.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |i|
                                                                                 &Func1::new({
                                                                                                 let i
                                                                                                     =
                                                                                                     i.clone();
                                                                                                 move
                                                                                                     |nes|
                                                                                                     {
                                                                                                         let s =
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString(),
                                                                                                                                              nes);
                                                                                                         let matchValue =
                                                                                                             Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                          &&&i),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length(),
                                                                                                                                                                                                          &&&s)));
                                                                                                         match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                          &matchValue)
                                                                                                             {
                                                                                                             0_i32
                                                                                                             =>
                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                             _
                                                                                                             =>
                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString(),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_dropRight(),
                                                                                                                                                                                                                                                                           &&&i),
                                                                                                                                                                                                                                        &&&s)))),
                                                                                                         }
                                                                                                     }
                                                                                             })))
    }
    pub fn Data_String_NonEmpty_CodeUnits_drop() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_drop: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_drop.get_or_init(||
                                                            &Func1::new(move
                                                                            |i|
                                                                            &Func1::new({
                                                                                            let i
                                                                                                =
                                                                                                i.clone();
                                                                                            move
                                                                                                |nes|
                                                                                                {
                                                                                                    let s =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_fromNonEmptyString(),
                                                                                                                                         nes);
                                                                                                    let matchValue =
                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                        &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                     &&&i),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_length(),
                                                                                                                                                                                                     &&&s)));
                                                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                     &matchValue)
                                                                                                        {
                                                                                                        0_i32
                                                                                                        =>
                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                        _
                                                                                                        =>
                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString(),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_drop(),
                                                                                                                                                                                                                                                                      &&&i),
                                                                                                                                                                                                                                   &&&s)))),
                                                                                                    }
                                                                                                }
                                                                                        })))
    }
    pub fn Data_String_NonEmpty_CodeUnits_countPrefix() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_countPrefix:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_countPrefix.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                       &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS()),
                                                                                                    &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_countPrefix()))
    }
    pub fn Data_String_NonEmpty_CodeUnits_cons() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_cons: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_cons.get_or_init(||
                                                            &Func1::new(move
                                                                            |c|
                                                                            &Func1::new({
                                                                                            let c
                                                                                                =
                                                                                                c.clone();
                                                                                            move
                                                                                                |s|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_toNonEmptyString(),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_singleton(),
                                                                                                                                                                                                                                          &&&c)),
                                                                                                                                                                    s))
                                                                                        })))
    }
    pub fn Data_String_NonEmpty_CodeUnits_charAt() -> &dyn Any {
        static Data_String_NonEmpty_CodeUnits_charAt:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodeUnits_charAt.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                  &&&PureScript_Data_String_NonEmpty_CodeUnits::Data_String_NonEmpty_CodeUnits_liftS()),
                                                                                               &&&PureScript_Data_String_CodeUnits::Data_String_CodeUnits_charAt()))
    }
}

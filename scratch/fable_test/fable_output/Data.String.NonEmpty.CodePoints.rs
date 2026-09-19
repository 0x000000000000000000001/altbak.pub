pub mod PureScript_Data_String_NonEmpty_CodePoints {
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
    use crate::module_d71935fd::PureScript_Data_String_NonEmpty_Internal;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_String_NonEmpty_CodePoints_toNonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_toNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_toNonEmptyString.get_or_init(||
                                                                         &PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_NonEmptyString())
    }
    pub fn Data_String_NonEmpty_CodePoints_snoc() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_snoc: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_snoc.get_or_init(||
                                                             &Func1::new(move
                                                                             |c|
                                                                             &Func1::new({
                                                                                             let c
                                                                                                 =
                                                                                                 c.clone();
                                                                                             move
                                                                                                 |s|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_toNonEmptyString(),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                        s),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                 1_i32.get_Message(),),
                                                                                                                                                                                                        &&&c)))
                                                                                         })))
    }
    pub fn Data_String_NonEmpty_CodePoints_singleton() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_singleton:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_singleton.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                      &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_toNonEmptyString()),
                                                                                                   &&panic!("{}",
                                                                                                            1_i32.get_Message(),)))
    }
    pub fn Data_String_NonEmpty_CodePoints_liftS() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_liftS:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_liftS.get_or_init(||
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
    pub fn Data_String_NonEmpty_CodePoints_takeWhile() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_takeWhile:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_takeWhile.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |f|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                      &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS(),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                  1_i32.get_Message(),),
                                                                                                                                                                                         f)))))
    }
    pub fn Data_String_NonEmpty_CodePoints_lastIndexOf_prime() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_lastIndexOf_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_lastIndexOf_prime.get_or_init(||
                                                                          &Func1::new(move
                                                                                          |pat|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                              &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS()),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                       1_i32.get_Message(),),
                                                                                                                                                              pat))))
    }
    pub fn Data_String_NonEmpty_CodePoints_lastIndexOf() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_lastIndexOf:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_lastIndexOf.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                        &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS()),
                                                                                                     &&panic!("{}",
                                                                                                              1_i32.get_Message(),)))
    }
    pub fn Data_String_NonEmpty_CodePoints_indexOf_prime() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_indexOf_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_indexOf_prime.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |pat|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                          &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                   1_i32.get_Message(),),
                                                                                                                                                          pat))))
    }
    pub fn Data_String_NonEmpty_CodePoints_indexOf() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_indexOf:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_indexOf.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                    &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS()),
                                                                                                 &&panic!("{}",
                                                                                                          1_i32.get_Message(),)))
    }
    pub fn Data_String_NonEmpty_CodePoints_fromNonEmptyString() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_fromNonEmptyString:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_fromNonEmptyString.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |v|
                                                                                           &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_String_NonEmpty_CodePoints_length() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_length:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_length.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                   &&panic!("{}",
                                                                                                                                            1_i32.get_Message(),)),
                                                                                                &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_fromNonEmptyString()))
    }
    pub fn Data_String_NonEmpty_CodePoints_splitAt() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_splitAt:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_splitAt.get_or_init(||
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
                                                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                   1_i32.get_Message(),),
                                                                                                                                                                                                          &&&i),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_fromNonEmptyString(),
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
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.NonEmpty.CodePoints.fs"),
                                  Data1: 29_i32,
                                  Data2: 323_i32,}).get_Message(),)
                                                                                                                }
                                                                                                            } else {
                                                                                                                panic!("{}",
                                                                                                                       LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.String.NonEmpty.CodePoints.fs"),
                                  Data1: 29_i32,
                                  Data2: 323_i32,}).get_Message(),)
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                            })))
    }
    pub fn Data_String_NonEmpty_CodePoints_take() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_take: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_take.get_or_init(||
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
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_fromNonEmptyString(),
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
                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_toNonEmptyString(),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                                                                1_i32.get_Message(),),
                                                                                                                                                                                                                                                                       &&&i),
                                                                                                                                                                                                                                    &&&s)))),
                                                                                                     }
                                                                                                 }
                                                                                         })))
    }
    pub fn Data_String_NonEmpty_CodePoints_toCodePointArray() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_toCodePointArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_toCodePointArray.get_or_init(||
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                             &&panic!("{}",
                                                                                                                                                      1_i32.get_Message(),)),
                                                                                                          &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_fromNonEmptyString()))
    }
    pub fn Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray()
     -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray.get_or_init(||
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
                                                                                                                                                     &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_toCodePointArray())))
    }
    pub fn Data_String_NonEmpty_CodePoints_uncons() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_uncons:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_uncons.get_or_init(||
                                                               &Func1::new(move
                                                                               |nes|
                                                                               {
                                                                                   let s =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_fromNonEmptyString(),
                                                                                                                        nes);
                                                                                   &add(string("head"),
                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                               |usd__unused|
                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                         1_i32.get_Message(),),
                                                                                                                                                                                                &&&0_i32),
                                                                                                                                                             &&&s)),
                                                                                        add(string("tail"),
                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                             1_i32.get_Message(),),
                                                                                                                                                                                                    &&&1_i32),
                                                                                                                                                                 &&&s)),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))
                                                                               }))
    }
    pub fn Data_String_NonEmpty_CodePoints_fromFoldable1() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_fromFoldable1:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_fromFoldable1.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictFoldable1|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                                                                                                                             dictFoldable1),
                                                                                                                                                          &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_semigroupNonEmptyString()),
                                                                                                                       &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_singleton())))
    }
    pub fn Data_String_NonEmpty_CodePoints_fromCodePointArray() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_fromCodePointArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_fromCodePointArray.get_or_init(||
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
                                                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_toNonEmptyString(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                       1_i32.get_Message(),),
                                                                                                                                                                                                                              &&&matchValue))))
                                                                                               }
                                                                                           }))
    }
    pub fn Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray()
     -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray.get_or_init(||
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                            |usd__unused|
                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined())))),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                          &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_fromCodePointArray()),
                                                                                                                                                       &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())))
    }
    pub fn Data_String_NonEmpty_CodePoints_dropWhile() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_dropWhile:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_dropWhile.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |f|
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                      &&&PureScript_Data_String_NonEmpty_Internal::Data_String_NonEmpty_Internal_fromString()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS(),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                  1_i32.get_Message(),),
                                                                                                                                                                                         f)))))
    }
    pub fn Data_String_NonEmpty_CodePoints_drop() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_drop: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_drop.get_or_init(||
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
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_fromNonEmptyString(),
                                                                                                                                          nes);
                                                                                                     let matchValue =
                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                         &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                      &&&i),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                               1_i32.get_Message(),),
                                                                                                                                                                                                      &&&s)));
                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                      &matchValue)
                                                                                                         {
                                                                                                         0_i32
                                                                                                         =>
                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                         _
                                                                                                         =>
                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_toNonEmptyString(),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                                                                1_i32.get_Message(),),
                                                                                                                                                                                                                                                                       &&&i),
                                                                                                                                                                                                                                    &&&s)))),
                                                                                                     }
                                                                                                 }
                                                                                         })))
    }
    pub fn Data_String_NonEmpty_CodePoints_countPrefix() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_countPrefix:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_countPrefix.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                        &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS()),
                                                                                                     &&panic!("{}",
                                                                                                              1_i32.get_Message(),)))
    }
    pub fn Data_String_NonEmpty_CodePoints_cons() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_cons: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_cons.get_or_init(||
                                                             &Func1::new(move
                                                                             |c|
                                                                             &Func1::new({
                                                                                             let c
                                                                                                 =
                                                                                                 c.clone();
                                                                                             move
                                                                                                 |s|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_toNonEmptyString(),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&panic!("{}",
                                                                                                                                                                                                                                                    1_i32.get_Message(),),
                                                                                                                                                                                                                                           &&&c)),
                                                                                                                                                                     s))
                                                                                         })))
    }
    pub fn Data_String_NonEmpty_CodePoints_codePointAt() -> &dyn Any {
        static Data_String_NonEmpty_CodePoints_codePointAt:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_String_NonEmpty_CodePoints_codePointAt.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                        &&&PureScript_Data_String_NonEmpty_CodePoints::Data_String_NonEmpty_CodePoints_liftS()),
                                                                                                     &&panic!("{}",
                                                                                                              1_i32.get_Message(),)))
    }
}

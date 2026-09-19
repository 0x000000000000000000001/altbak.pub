pub mod PureScript_Data_Array_NonEmpty {
    use super::*;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_ba715d91::PureScript_Data_Array_NonEmpty_Internal;
    use crate::module_2d8e16c::PureScript_Data_Array;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_2563115d::PureScript_Data_Semigroup_Foldable;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_b1754f14::PureScript_Data_Unfoldable1;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_285149e9::PureScript_Safe_Coerce;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use crate::module_2a7662d2::PureScript_Unsafe_Coerce;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_Array_NonEmpty_intercalate1() -> &dyn Any {
        static Data_Array_NonEmpty_intercalate1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_intercalate1.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_intercalate(),
                                                                                          &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))
    }
    pub fn Data_Array_NonEmpty_foldMap11() -> &dyn Any {
        static Data_Array_NonEmpty_foldMap11: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_foldMap11.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldMap1(),
                                                                                       &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))
    }
    pub fn Data_Array_NonEmpty_fold11() -> &dyn Any {
        static Data_Array_NonEmpty_fold11: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_fold11.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_fold1(),
                                                                                    &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))
    }
    pub fn Data_Array_NonEmpty_fromJust() -> &dyn Any {
        static Data_Array_NonEmpty_fromJust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_fromJust.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                      &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Array_NonEmpty_unsafeIndex1() -> &dyn Any {
        static Data_Array_NonEmpty_unsafeIndex1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_unsafeIndex1.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unsafeIndex(),
                                                                                          &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_Array_NonEmpty_unsafeFromArrayF() -> &dyn Any {
        static Data_Array_NonEmpty_unsafeFromArrayF: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Array_NonEmpty_unsafeFromArrayF.get_or_init(||
                                                             &PureScript_Unsafe_Coerce::Unsafe_Coerce_unsafeCoerce())
    }
    pub fn Data_Array_NonEmpty_unsafeFromArray() -> &dyn Any {
        static Data_Array_NonEmpty_unsafeFromArray: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Array_NonEmpty_unsafeFromArray.get_or_init(||
                                                            &PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_NonEmptyArray())
    }
    pub fn Data_Array_NonEmpty_transpose() -> &dyn Any {
        static Data_Array_NonEmpty_transpose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_transpose.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined())),
                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&&PureScript_Data_Array::Data_Array_transpose()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()))))
    }
    pub fn Data_Array_NonEmpty_toArray() -> &dyn Any {
        static Data_Array_NonEmpty_toArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_toArray.get_or_init(||
                                                    &Func1::new(move |v|
                                                                    &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_Array_NonEmpty_unionBy_prime() -> &dyn Any {
        static Data_Array_NonEmpty_unionBy_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_unionBy_prime.get_or_init(||
                                                          &Func1::new(move
                                                                          |eq|
                                                                          &Func1::new({
                                                                                          let eq
                                                                                              =
                                                                                              eq.clone();
                                                                                          move
                                                                                              |xs|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                  &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_unionBy(),
                                                                                                                                                                                                     &&&eq),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                     xs)))
                                                                                      })))
    }
    pub fn Data_Array_NonEmpty_union_prime() -> &dyn Any {
        static Data_Array_NonEmpty_union_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_union_prime.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictEq|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unionBy_prime(),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                            dictEq))))
    }
    pub fn Data_Array_NonEmpty_unionBy() -> &dyn Any {
        static Data_Array_NonEmpty_unionBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_unionBy.get_or_init(||
                                                    &Func1::new(move |eq|
                                                                    &Func1::new({
                                                                                    let eq
                                                                                        =
                                                                                        eq.clone();
                                                                                    move
                                                                                        |xs|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unionBy_prime(),
                                                                                                                                                                                                                                  &&&eq),
                                                                                                                                                                                               xs)),
                                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())
                                                                                })))
    }
    pub fn Data_Array_NonEmpty_union() -> &dyn Any {
        static Data_Array_NonEmpty_union: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_union.get_or_init(||
                                                  &Func1::new(move |dictEq|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unionBy(),
                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                      dictEq))))
    }
    pub fn Data_Array_NonEmpty_unzip() -> &dyn Any {
        static Data_Array_NonEmpty_unzip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_unzip.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                               &&&PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorTuple()),
                                                                                                                                                                                            &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray())),
                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                         &&&PureScript_Data_Array::Data_Array_unzip()),
                                                                                                                      &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())))
    }
    pub fn Data_Array_NonEmpty_updateAt() -> &dyn Any {
        static Data_Array_NonEmpty_updateAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_updateAt.get_or_init(||
                                                     &Func1::new(move |i|
                                                                     &Func1::new({
                                                                                     let i
                                                                                         =
                                                                                         i.clone();
                                                                                     move
                                                                                         |x|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArrayF()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_updateAt(),
                                                                                                                                                                                                                                                                      &&&i),
                                                                                                                                                                                                                                   x)),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray()))
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_zip() -> &dyn Any {
        static Data_Array_NonEmpty_zip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_zip.get_or_init(||
                                                &Func1::new(move |xs|
                                                                &Func1::new({
                                                                                let xs
                                                                                    =
                                                                                    xs.clone();
                                                                                move
                                                                                    |ys|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                        &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_zip(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                                              &&&xs)),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                           ys)))
                                                                            })))
    }
    pub fn Data_Array_NonEmpty_zipWith() -> &dyn Any {
        static Data_Array_NonEmpty_zipWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_zipWith.get_or_init(||
                                                    &Func1::new(move |f|
                                                                    &Func1::new({
                                                                                    let f
                                                                                        =
                                                                                        f.clone();
                                                                                    move
                                                                                        |xs|
                                                                                        &Func1::new({
                                                                                                        let xs
                                                                                                            =
                                                                                                            xs.clone();
                                                                                                        move
                                                                                                            |ys|
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_zipWith(),
                                                                                                                                                                                                                                                      &&&f),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                                                                      &&&xs)),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                                   ys)))
                                                                                                    })
                                                                                })))
    }
    pub fn Data_Array_NonEmpty_zipWithA() -> &dyn Any {
        static Data_Array_NonEmpty_zipWithA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_zipWithA.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictApplicative|
                                                                     &Func1::new({
                                                                                     let dictApplicative
                                                                                         =
                                                                                         dictApplicative.clone();
                                                                                     move
                                                                                         |f|
                                                                                         &Func1::new({
                                                                                                         let f
                                                                                                             =
                                                                                                             f.clone();
                                                                                                         move
                                                                                                             |xs|
                                                                                                             &Func1::new({
                                                                                                                             let xs
                                                                                                                                 =
                                                                                                                                 xs.clone();
                                                                                                                             move
                                                                                                                                 |ys|
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                     &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArrayF()),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_zipWithA(),
                                                                                                                                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                                                                                                                                           &&&f),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                                                                                           &&&xs)),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                                                        ys)))
                                                                                                                         })
                                                                                                     })
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_splitAt() -> &dyn Any {
        static Data_Array_NonEmpty_splitAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_splitAt.get_or_init(||
                                                    &Func1::new(move |i|
                                                                    &Func1::new({
                                                                                    let i
                                                                                        =
                                                                                        i.clone();
                                                                                    move
                                                                                        |xs|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_splitAt(),
                                                                                                                                                                                               &&&i)),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                            xs))
                                                                                })))
    }
    pub fn Data_Array_NonEmpty_some() -> &dyn Any {
        static Data_Array_NonEmpty_some: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_some.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictAlternative|
                                                                 {
                                                                     let some1 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_some(),
                                                                                                          dictAlternative);
                                                                     &Func1::new({
                                                                                     let some1
                                                                                         =
                                                                                         some1.clone();
                                                                                     move
                                                                                         |dictLazy|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArrayF()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&some1,
                                                                                                                                                             dictLazy))
                                                                                 })
                                                                 }))
    }
    pub fn Data_Array_NonEmpty_snoc_prime() -> &dyn Any {
        static Data_Array_NonEmpty_snoc_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_snoc_prime.get_or_init(||
                                                       &Func1::new(move |xs|
                                                                       &Func1::new({
                                                                                       let xs
                                                                                           =
                                                                                           xs.clone();
                                                                                       move
                                                                                           |x|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                               &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_snoc(),
                                                                                                                                                                                                  &&&xs),
                                                                                                                                                               x))
                                                                                   })))
    }
    pub fn Data_Array_NonEmpty_snoc() -> &dyn Any {
        static Data_Array_NonEmpty_snoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_snoc.get_or_init(||
                                                 &Func1::new(move |xs|
                                                                 &Func1::new({
                                                                                 let xs
                                                                                     =
                                                                                     xs.clone();
                                                                                 move
                                                                                     |x|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_snoc(),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                                               &&&xs)),
                                                                                                                                                         x))
                                                                             })))
    }
    pub fn Data_Array_NonEmpty_singleton() -> &dyn Any {
        static Data_Array_NonEmpty_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_singleton.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                       &&&PureScript_Data_Array::Data_Array_singleton()))
    }
    pub fn Data_Array_NonEmpty_replicate() -> &dyn Any {
        static Data_Array_NonEmpty_replicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_replicate.get_or_init(||
                                                      &Func1::new(move |i|
                                                                      &Func1::new({
                                                                                      let i
                                                                                          =
                                                                                          i.clone();
                                                                                      move
                                                                                          |x|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                              &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_replicate(),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_max(),
                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                       &&&1_i32),
                                                                                                                                                                                                                                    &&&i)),
                                                                                                                                                              x))
                                                                                  })))
    }
    pub fn Data_Array_NonEmpty_range() -> &dyn Any {
        static Data_Array_NonEmpty_range: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_range.get_or_init(||
                                                  &Func1::new(move |x|
                                                                  &Func1::new({
                                                                                  let x
                                                                                      =
                                                                                      x.clone();
                                                                                  move
                                                                                      |y|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_range(),
                                                                                                                                                                                             &&&x),
                                                                                                                                                          y))
                                                                              })))
    }
    pub fn Data_Array_NonEmpty_prependArray() -> &dyn Any {
        static Data_Array_NonEmpty_prependArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_prependArray.get_or_init(||
                                                         &Func1::new(move |xs|
                                                                         &Func1::new({
                                                                                         let xs
                                                                                             =
                                                                                             xs.clone();
                                                                                         move
                                                                                             |ys|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                 &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                       &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                                                                                                                                    &&&xs),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                    ys)))
                                                                                     })))
    }
    pub fn Data_Array_NonEmpty_modifyAt() -> &dyn Any {
        static Data_Array_NonEmpty_modifyAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_modifyAt.get_or_init(||
                                                     &Func1::new(move |i|
                                                                     &Func1::new({
                                                                                     let i
                                                                                         =
                                                                                         i.clone();
                                                                                     move
                                                                                         |f|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArrayF()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_modifyAt(),
                                                                                                                                                                                                                                                                      &&&i),
                                                                                                                                                                                                                                   f)),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray()))
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_intersectBy_prime() -> &dyn Any {
        static Data_Array_NonEmpty_intersectBy_prime:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_intersectBy_prime.get_or_init(||
                                                              &Func1::new(move
                                                                              |eq|
                                                                              &Func1::new({
                                                                                              let eq
                                                                                                  =
                                                                                                  eq.clone();
                                                                                              move
                                                                                                  |xs|
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_intersectBy(),
                                                                                                                                                                      &&&eq),
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                      xs))
                                                                                          })))
    }
    pub fn Data_Array_NonEmpty_intersectBy() -> &dyn Any {
        static Data_Array_NonEmpty_intersectBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_intersectBy.get_or_init(||
                                                        &Func1::new(move |eq|
                                                                        &Func1::new({
                                                                                        let eq
                                                                                            =
                                                                                            eq.clone();
                                                                                        move
                                                                                            |xs|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_intersectBy_prime(),
                                                                                                                                                                                                                                      &&&eq),
                                                                                                                                                                                                   xs)),
                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())
                                                                                    })))
    }
    pub fn Data_Array_NonEmpty_intersect_prime() -> &dyn Any {
        static Data_Array_NonEmpty_intersect_prime: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Array_NonEmpty_intersect_prime.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictEq|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_intersectBy_prime(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                dictEq))))
    }
    pub fn Data_Array_NonEmpty_intersect() -> &dyn Any {
        static Data_Array_NonEmpty_intersect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_intersect.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictEq|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_intersectBy(),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                          dictEq))))
    }
    pub fn Data_Array_NonEmpty_intercalate() -> &dyn Any {
        static Data_Array_NonEmpty_intercalate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_intercalate.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictSemigroup|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_intercalate1(),
                                                                                                         dictSemigroup)))
    }
    pub fn Data_Array_NonEmpty_insertAt() -> &dyn Any {
        static Data_Array_NonEmpty_insertAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_insertAt.get_or_init(||
                                                     &Func1::new(move |i|
                                                                     &Func1::new({
                                                                                     let i
                                                                                         =
                                                                                         i.clone();
                                                                                     move
                                                                                         |x|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArrayF()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_insertAt(),
                                                                                                                                                                                                                                                                      &&&i),
                                                                                                                                                                                                                                   x)),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray()))
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_fromFoldable1() -> &dyn Any {
        static Data_Array_NonEmpty_fromFoldable1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_fromFoldable1.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictFoldable1|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                              &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_fromFoldable(),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Foldable0"),
                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictFoldable1)),
                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined())))))
    }
    pub fn Data_Array_NonEmpty_fromArray() -> &dyn Any {
        static Data_Array_NonEmpty_fromArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_fromArray.get_or_init(||
                                                      &Func1::new(move |xs|
                                                                      {
                                                                          let matchValue =
                                                                              Sharpurs_Prelude::unbox(xs);
                                                                          if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                             &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_length(),
                                                                                                                                                                                                             &&&matchValue)),
                                                                                                                                       &&&0_i32))
                                                                             {
                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray(),
                                                                                                                                                                      &&&matchValue)))
                                                                          } else {
                                                                              if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                 {
                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                              } else {
                                                                                  panic!("{}",
                                                                                         LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Array.NonEmpty.fs"),
                                  Data1: 75_i32,
                                  Data2: 67_i32,}).get_Message(),)
                                                                              }
                                                                          }
                                                                      }))
    }
    pub fn Data_Array_NonEmpty_fromFoldable() -> &dyn Any {
        static Data_Array_NonEmpty_fromFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_fromFoldable.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictFoldable|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_fromArray()),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_fromFoldable(),
                                                                                                                                             dictFoldable))))
    }
    pub fn Data_Array_NonEmpty_transpose_prime() -> &dyn Any {
        static Data_Array_NonEmpty_transpose_prime: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Array_NonEmpty_transpose_prime.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_fromArray()),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                   &&&PureScript_Data_Array::Data_Array_transpose()),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Safe_Coerce::Safe_Coerce_coerce(),
                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()))))
    }
    pub fn Data_Array_NonEmpty_foldr1() -> &dyn Any {
        static Data_Array_NonEmpty_foldr1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_foldr1.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldr1(),
                                                                                    &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))
    }
    pub fn Data_Array_NonEmpty_foldl1() -> &dyn Any {
        static Data_Array_NonEmpty_foldl1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_foldl1.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Foldable::Data_Semigroup_Foldable_foldl1(),
                                                                                    &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))
    }
    pub fn Data_Array_NonEmpty_foldMap1() -> &dyn Any {
        static Data_Array_NonEmpty_foldMap1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_foldMap1.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictSemigroup|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_foldMap11(),
                                                                                                      dictSemigroup)))
    }
    pub fn Data_Array_NonEmpty_fold1() -> &dyn Any {
        static Data_Array_NonEmpty_fold1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_fold1.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictSemigroup|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_fold11(),
                                                                                                   dictSemigroup)))
    }
    pub fn Data_Array_NonEmpty_difference_prime() -> &dyn Any {
        static Data_Array_NonEmpty_difference_prime: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Array_NonEmpty_difference_prime.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictEq|
                                                                             {
                                                                                 let difference1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_difference(),
                                                                                                                      dictEq);
                                                                                 &Func1::new({
                                                                                                 let difference1
                                                                                                     =
                                                                                                     difference1.clone();
                                                                                                 move
                                                                                                     |xs|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                         &&&difference1),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                         xs))
                                                                                             })
                                                                             }))
    }
    pub fn Data_Array_NonEmpty_cons_prime() -> &dyn Any {
        static Data_Array_NonEmpty_cons_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_cons_prime.get_or_init(||
                                                       &Func1::new(move |x|
                                                                       &Func1::new({
                                                                                       let x
                                                                                           =
                                                                                           x.clone();
                                                                                       move
                                                                                           |xs|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                               &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_cons(),
                                                                                                                                                                                                  &&&x),
                                                                                                                                                               xs))
                                                                                   })))
    }
    pub fn Data_Array_NonEmpty_fromNonEmpty() -> &dyn Any {
        static Data_Array_NonEmpty_fromNonEmpty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_fromNonEmpty.get_or_init(||
                                                         &Func1::new(move |v|
                                                                         {
                                                                             let matchValue:
                                                                                     LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                 Sharpurs_Prelude::unbox(v);
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_cons_prime(),
                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                        Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                               _)
                                                                                                                                                        =>
                                                                                                                                                        x.clone(),
                                                                                                                                                    }),
                                                                                                              &&&match matchValue.as_ref()
                                                                                                                     {
                                                                                                                     Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                            x)
                                                                                                                     =>
                                                                                                                     x.clone(),
                                                                                                                 })
                                                                         }))
    }
    pub fn Data_Array_NonEmpty_concatMap() -> &dyn Any {
        static Data_Array_NonEmpty_concatMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_concatMap.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                          &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_bindNonEmptyArray())))
    }
    pub fn Data_Array_NonEmpty_concat() -> &dyn Any {
        static Data_Array_NonEmpty_concat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_concat.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                       &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                          &&&PureScript_Data_Array::Data_Array_concat()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray()),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                &&&PureScript_Data_Array_NonEmpty_Internal::Data_Array_NonEmpty_Internal_functorNonEmptyArray()),
                                                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())))))
    }
    pub fn Data_Array_NonEmpty_appendArray() -> &dyn Any {
        static Data_Array_NonEmpty_appendArray: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_appendArray.get_or_init(||
                                                        &Func1::new(move |xs|
                                                                        &Func1::new({
                                                                                        let xs
                                                                                            =
                                                                                            xs.clone();
                                                                                        move
                                                                                            |ys|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupArray()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray(),
                                                                                                                                                                                                                                      &&&xs)),
                                                                                                                                                                ys))
                                                                                    })))
    }
    pub fn Data_Array_NonEmpty_alterAt() -> &dyn Any {
        static Data_Array_NonEmpty_alterAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_alterAt.get_or_init(||
                                                    &Func1::new(move |i|
                                                                    &Func1::new({
                                                                                    let i
                                                                                        =
                                                                                        i.clone();
                                                                                    move
                                                                                        |f|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_alterAt(),
                                                                                                                                                                                                                                  &&&i),
                                                                                                                                                                                               f)),
                                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())
                                                                                })))
    }
    pub fn Data_Array_NonEmpty_adaptMaybe() -> &dyn Any {
        static Data_Array_NonEmpty_adaptMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_adaptMaybe.get_or_init(||
                                                       &Func1::new(move |f|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                           &&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial()),
                                                                                                        &&&Func1::new({
                                                                                                                          let f
                                                                                                                              =
                                                                                                                              f.clone();
                                                                                                                          move
                                                                                                                              |usd__unused|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                  &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_fromJust()),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                     &&&f),
                                                                                                                                                                                                  &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray()))
                                                                                                                      }))))
    }
    pub fn Data_Array_NonEmpty_head() -> &dyn Any {
        static Data_Array_NonEmpty_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_head.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptMaybe(),
                                                                                  &&&PureScript_Data_Array::Data_Array_head()))
    }
    pub fn Data_Array_NonEmpty_init() -> &dyn Any {
        static Data_Array_NonEmpty_init: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_init.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptMaybe(),
                                                                                  &&&PureScript_Data_Array::Data_Array_init()))
    }
    pub fn Data_Array_NonEmpty_last() -> &dyn Any {
        static Data_Array_NonEmpty_last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_last.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptMaybe(),
                                                                                  &&&PureScript_Data_Array::Data_Array_last()))
    }
    pub fn Data_Array_NonEmpty_tail() -> &dyn Any {
        static Data_Array_NonEmpty_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_tail.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptMaybe(),
                                                                                  &&&PureScript_Data_Array::Data_Array_tail()))
    }
    pub fn Data_Array_NonEmpty_uncons() -> &dyn Any {
        static Data_Array_NonEmpty_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_uncons.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptMaybe(),
                                                                                    &&&PureScript_Data_Array::Data_Array_uncons()))
    }
    pub fn Data_Array_NonEmpty_toNonEmpty() -> &dyn Any {
        static Data_Array_NonEmpty_toNonEmpty: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_toNonEmpty.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_composeFlipped(),
                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                           &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_uncons()),
                                                                                        &&&Func1::new(move
                                                                                                          |v|
                                                                                                          {
                                                                                                              let matchValue =
                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                              {
                                                                                                                  let activePatternResult =
                                                                                                                      Sharpurs_Prelude::_007cHasProp_007c__007c(string("head"),
                                                                                                                                                                &matchValue);
                                                                                                                  if activePatternResult.is_some()
                                                                                                                     {
                                                                                                                      let activePatternResult_1 =
                                                                                                                          Sharpurs_Prelude::_007cHasProp_007c__007c(string("tail"),
                                                                                                                                                                    &matchValue);
                                                                                                                      if activePatternResult_1.is_some()
                                                                                                                         {
                                                                                                                          let x =
                                                                                                                              getValue(activePatternResult);
                                                                                                                          let xs =
                                                                                                                              getValue(activePatternResult_1);
                                                                                                                          &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&x,
                                                                                                                                                                                              &xs))
                                                                                                                      } else {
                                                                                                                          panic!("{}",
                                                                                                                                 LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Array.NonEmpty.fs"),
                                  Data1: 115_i32,
                                  Data2: 280_i32,}).get_Message(),)
                                                                                                                      }
                                                                                                                  } else {
                                                                                                                      panic!("{}",
                                                                                                                             LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Array.NonEmpty.fs"),
                                  Data1: 115_i32,
                                  Data2: 280_i32,}).get_Message(),)
                                                                                                                  }
                                                                                                              }
                                                                                                          })))
    }
    pub fn Data_Array_NonEmpty_unsnoc() -> &dyn Any {
        static Data_Array_NonEmpty_unsnoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_unsnoc.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptMaybe(),
                                                                                    &&&PureScript_Data_Array::Data_Array_unsnoc()))
    }
    pub fn Data_Array_NonEmpty_adaptAny() -> &dyn Any {
        static Data_Array_NonEmpty_adaptAny: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_adaptAny.get_or_init(||
                                                     &Func1::new(move |f|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                         f),
                                                                                                      &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_toArray())))
    }
    pub fn Data_Array_NonEmpty_all() -> &dyn Any {
        static Data_Array_NonEmpty_all: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_all.get_or_init(||
                                                &Func1::new(move |p|
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                    &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_all(),
                                                                                                                                    p))))
    }
    pub fn Data_Array_NonEmpty_any() -> &dyn Any {
        static Data_Array_NonEmpty_any: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_any.get_or_init(||
                                                &Func1::new(move |p|
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                    &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_any(),
                                                                                                                                    p))))
    }
    pub fn Data_Array_NonEmpty_catMaybes() -> &dyn Any {
        static Data_Array_NonEmpty_catMaybes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_catMaybes.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny(),
                                                                                       &&&PureScript_Data_Array::Data_Array_catMaybes()))
    }
    pub fn Data_Array_NonEmpty_delete() -> &dyn Any {
        static Data_Array_NonEmpty_delete: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_delete.get_or_init(||
                                                   &Func1::new(move |dictEq|
                                                                   &Func1::new({
                                                                                   let dictEq
                                                                                       =
                                                                                       dictEq.clone();
                                                                                   move
                                                                                       |x|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                           &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_delete(),
                                                                                                                                                                                              &&&dictEq),
                                                                                                                                                           x))
                                                                               })))
    }
    pub fn Data_Array_NonEmpty_deleteAt() -> &dyn Any {
        static Data_Array_NonEmpty_deleteAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_deleteAt.get_or_init(||
                                                     &Func1::new(move |i|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_deleteAt(),
                                                                                                                                         i))))
    }
    pub fn Data_Array_NonEmpty_deleteBy() -> &dyn Any {
        static Data_Array_NonEmpty_deleteBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_deleteBy.get_or_init(||
                                                     &Func1::new(move |f|
                                                                     &Func1::new({
                                                                                     let f
                                                                                         =
                                                                                         f.clone();
                                                                                     move
                                                                                         |x|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_deleteBy(),
                                                                                                                                                                                                &&&f),
                                                                                                                                                             x))
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_difference() -> &dyn Any {
        static Data_Array_NonEmpty_difference: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_difference.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictEq|
                                                                       &Func1::new({
                                                                                       let dictEq
                                                                                           =
                                                                                           dictEq.clone();
                                                                                       move
                                                                                           |xs|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                               &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_difference_prime(),
                                                                                                                                                                                                  &&&dictEq),
                                                                                                                                                               xs))
                                                                                   })))
    }
    pub fn Data_Array_NonEmpty_drop() -> &dyn Any {
        static Data_Array_NonEmpty_drop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_drop.get_or_init(||
                                                 &Func1::new(move |i|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                     &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_drop(),
                                                                                                                                     i))))
    }
    pub fn Data_Array_NonEmpty_dropEnd() -> &dyn Any {
        static Data_Array_NonEmpty_dropEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_dropEnd.get_or_init(||
                                                    &Func1::new(move |i|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                        &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_dropEnd(),
                                                                                                                                        i))))
    }
    pub fn Data_Array_NonEmpty_dropWhile() -> &dyn Any {
        static Data_Array_NonEmpty_dropWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_dropWhile.get_or_init(||
                                                      &Func1::new(move |f|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_dropWhile(),
                                                                                                                                          f))))
    }
    pub fn Data_Array_NonEmpty_elem() -> &dyn Any {
        static Data_Array_NonEmpty_elem: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_elem.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 &Func1::new({
                                                                                 let dictEq
                                                                                     =
                                                                                     dictEq.clone();
                                                                                 move
                                                                                     |x|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_elem(),
                                                                                                                                                                                            &&&dictEq),
                                                                                                                                                         x))
                                                                             })))
    }
    pub fn Data_Array_NonEmpty_elemIndex() -> &dyn Any {
        static Data_Array_NonEmpty_elemIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_elemIndex.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictEq|
                                                                      &Func1::new({
                                                                                      let dictEq
                                                                                          =
                                                                                          dictEq.clone();
                                                                                      move
                                                                                          |x|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                              &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_elemIndex(),
                                                                                                                                                                                                 &&&dictEq),
                                                                                                                                                              x))
                                                                                  })))
    }
    pub fn Data_Array_NonEmpty_elemLastIndex() -> &dyn Any {
        static Data_Array_NonEmpty_elemLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_elemLastIndex.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictEq|
                                                                          &Func1::new({
                                                                                          let dictEq
                                                                                              =
                                                                                              dictEq.clone();
                                                                                          move
                                                                                              |x|
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                  &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_elemLastIndex(),
                                                                                                                                                                                                     &&&dictEq),
                                                                                                                                                                  x))
                                                                                      })))
    }
    pub fn Data_Array_NonEmpty_filter() -> &dyn Any {
        static Data_Array_NonEmpty_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_filter.get_or_init(||
                                                   &Func1::new(move |f|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                       &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_filter(),
                                                                                                                                       f))))
    }
    pub fn Data_Array_NonEmpty_filterA() -> &dyn Any {
        static Data_Array_NonEmpty_filterA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_filterA.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictApplicative|
                                                                    &Func1::new({
                                                                                    let dictApplicative
                                                                                        =
                                                                                        dictApplicative.clone();
                                                                                    move
                                                                                        |f|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                            &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_filterA(),
                                                                                                                                                                                               &&&dictApplicative),
                                                                                                                                                            f))
                                                                                })))
    }
    pub fn Data_Array_NonEmpty_find() -> &dyn Any {
        static Data_Array_NonEmpty_find: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_find.get_or_init(||
                                                 &Func1::new(move |p|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                     &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_find(),
                                                                                                                                     p))))
    }
    pub fn Data_Array_NonEmpty_findIndex() -> &dyn Any {
        static Data_Array_NonEmpty_findIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_findIndex.get_or_init(||
                                                      &Func1::new(move |p|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findIndex(),
                                                                                                                                          p))))
    }
    pub fn Data_Array_NonEmpty_findLastIndex() -> &dyn Any {
        static Data_Array_NonEmpty_findLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_findLastIndex.get_or_init(||
                                                          &Func1::new(move |x|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                              &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findLastIndex(),
                                                                                                                                              x))))
    }
    pub fn Data_Array_NonEmpty_findMap() -> &dyn Any {
        static Data_Array_NonEmpty_findMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_findMap.get_or_init(||
                                                    &Func1::new(move |p|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                        &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_findMap(),
                                                                                                                                        p))))
    }
    pub fn Data_Array_NonEmpty_foldM() -> &dyn Any {
        static Data_Array_NonEmpty_foldM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_foldM.get_or_init(||
                                                  &Func1::new(move |dictMonad|
                                                                  &Func1::new({
                                                                                  let dictMonad
                                                                                      =
                                                                                      dictMonad.clone();
                                                                                  move
                                                                                      |f|
                                                                                      &Func1::new({
                                                                                                      let f
                                                                                                          =
                                                                                                          f.clone();
                                                                                                      move
                                                                                                          |acc|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                              &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_foldM(),
                                                                                                                                                                                                                                                    &&&dictMonad),
                                                                                                                                                                                                                 &&&f),
                                                                                                                                                                              acc))
                                                                                                  })
                                                                              })))
    }
    pub fn Data_Array_NonEmpty_foldRecM() -> &dyn Any {
        static Data_Array_NonEmpty_foldRecM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_foldRecM.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonadRec|
                                                                     &Func1::new({
                                                                                     let dictMonadRec
                                                                                         =
                                                                                         dictMonadRec.clone();
                                                                                     move
                                                                                         |f|
                                                                                         &Func1::new({
                                                                                                         let f
                                                                                                             =
                                                                                                             f.clone();
                                                                                                         move
                                                                                                             |acc|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                 &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_foldRecM(),
                                                                                                                                                                                                                                                       &&&dictMonadRec),
                                                                                                                                                                                                                    &&&f),
                                                                                                                                                                                 acc))
                                                                                                     })
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_index() -> &dyn Any {
        static Data_Array_NonEmpty_index: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_index.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny(),
                                                                                   &&&PureScript_Data_Array::Data_Array_index()))
    }
    pub fn Data_Array_NonEmpty_length() -> &dyn Any {
        static Data_Array_NonEmpty_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_length.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny(),
                                                                                    &&&PureScript_Data_Array::Data_Array_length()))
    }
    pub fn Data_Array_NonEmpty_mapMaybe() -> &dyn Any {
        static Data_Array_NonEmpty_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_mapMaybe.get_or_init(||
                                                     &Func1::new(move |f|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_mapMaybe(),
                                                                                                                                         f))))
    }
    pub fn Data_Array_NonEmpty_notElem() -> &dyn Any {
        static Data_Array_NonEmpty_notElem: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_notElem.get_or_init(||
                                                    &Func1::new(move |dictEq|
                                                                    &Func1::new({
                                                                                    let dictEq
                                                                                        =
                                                                                        dictEq.clone();
                                                                                    move
                                                                                        |x|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                            &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_notElem(),
                                                                                                                                                                                               &&&dictEq),
                                                                                                                                                            x))
                                                                                })))
    }
    pub fn Data_Array_NonEmpty_partition() -> &dyn Any {
        static Data_Array_NonEmpty_partition: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_partition.get_or_init(||
                                                      &Func1::new(move |f|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_partition(),
                                                                                                                                          f))))
    }
    pub fn Data_Array_NonEmpty_slice() -> &dyn Any {
        static Data_Array_NonEmpty_slice: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_slice.get_or_init(||
                                                  &Func1::new(move |start|
                                                                  &Func1::new({
                                                                                  let start
                                                                                      =
                                                                                      start.clone();
                                                                                  move
                                                                                      |end_var|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_slice(),
                                                                                                                                                                                             &&&start),
                                                                                                                                                          end_var))
                                                                              })))
    }
    pub fn Data_Array_NonEmpty_span() -> &dyn Any {
        static Data_Array_NonEmpty_span: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_span.get_or_init(||
                                                 &Func1::new(move |f|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                     &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_span(),
                                                                                                                                     f))))
    }
    pub fn Data_Array_NonEmpty_take() -> &dyn Any {
        static Data_Array_NonEmpty_take: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_take.get_or_init(||
                                                 &Func1::new(move |i|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                     &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_take(),
                                                                                                                                     i))))
    }
    pub fn Data_Array_NonEmpty_takeEnd() -> &dyn Any {
        static Data_Array_NonEmpty_takeEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_takeEnd.get_or_init(||
                                                    &Func1::new(move |i|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                        &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_takeEnd(),
                                                                                                                                        i))))
    }
    pub fn Data_Array_NonEmpty_takeWhile() -> &dyn Any {
        static Data_Array_NonEmpty_takeWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_takeWhile.get_or_init(||
                                                      &Func1::new(move |f|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny()),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_takeWhile(),
                                                                                                                                          f))))
    }
    pub fn Data_Array_NonEmpty_toUnfoldable() -> &dyn Any {
        static Data_Array_NonEmpty_toUnfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_toUnfoldable.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictUnfoldable|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny(),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_toUnfoldable(),
                                                                                                                                             dictUnfoldable))))
    }
    pub fn Data_Array_NonEmpty_unsafeAdapt() -> &dyn Any {
        static Data_Array_NonEmpty_unsafeAdapt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_unsafeAdapt.get_or_init(||
                                                        &Func1::new(move |f|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                            &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeFromArray()),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny(),
                                                                                                                                            f))))
    }
    pub fn Data_Array_NonEmpty_cons() -> &dyn Any {
        static Data_Array_NonEmpty_cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_cons.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                     &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_cons(),
                                                                                                                                     x))))
    }
    pub fn Data_Array_NonEmpty_group() -> &dyn Any {
        static Data_Array_NonEmpty_group: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_group.get_or_init(||
                                                  &Func1::new(move |dictEq|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                      &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_group(),
                                                                                                                                      dictEq))))
    }
    pub fn Data_Array_NonEmpty_groupAllBy() -> &dyn Any {
        static Data_Array_NonEmpty_groupAllBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_groupAllBy.get_or_init(||
                                                       &Func1::new(move |op|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                           &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_groupAllBy(),
                                                                                                                                           op))))
    }
    pub fn Data_Array_NonEmpty_groupAll() -> &dyn Any {
        static Data_Array_NonEmpty_groupAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_groupAll.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictOrd|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_groupAllBy(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                         dictOrd))))
    }
    pub fn Data_Array_NonEmpty_groupBy() -> &dyn Any {
        static Data_Array_NonEmpty_groupBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_groupBy.get_or_init(||
                                                    &Func1::new(move |op|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                        &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_groupBy(),
                                                                                                                                        op))))
    }
    pub fn Data_Array_NonEmpty_insert() -> &dyn Any {
        static Data_Array_NonEmpty_insert: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_insert.get_or_init(||
                                                   &Func1::new(move |dictOrd|
                                                                   &Func1::new({
                                                                                   let dictOrd
                                                                                       =
                                                                                       dictOrd.clone();
                                                                                   move
                                                                                       |x|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                           &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_insert(),
                                                                                                                                                                                              &&&dictOrd),
                                                                                                                                                           x))
                                                                               })))
    }
    pub fn Data_Array_NonEmpty_insertBy() -> &dyn Any {
        static Data_Array_NonEmpty_insertBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_insertBy.get_or_init(||
                                                     &Func1::new(move |f|
                                                                     &Func1::new({
                                                                                     let f
                                                                                         =
                                                                                         f.clone();
                                                                                     move
                                                                                         |x|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_insertBy(),
                                                                                                                                                                                                &&&f),
                                                                                                                                                             x))
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_intersperse() -> &dyn Any {
        static Data_Array_NonEmpty_intersperse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_intersperse.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                            &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_intersperse(),
                                                                                                                                            x))))
    }
    pub fn Data_Array_NonEmpty_mapWithIndex() -> &dyn Any {
        static Data_Array_NonEmpty_mapWithIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_mapWithIndex.get_or_init(||
                                                         &Func1::new(move |f|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_mapWithIndex(),
                                                                                                                                             f))))
    }
    pub fn Data_Array_NonEmpty_modifyAtIndices() -> &dyn Any {
        static Data_Array_NonEmpty_modifyAtIndices: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Array_NonEmpty_modifyAtIndices.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFoldable|
                                                                            &Func1::new({
                                                                                            let dictFoldable
                                                                                                =
                                                                                                dictFoldable.clone();
                                                                                            move
                                                                                                |is|
                                                                                                &Func1::new({
                                                                                                                let is
                                                                                                                    =
                                                                                                                    is.clone();
                                                                                                                move
                                                                                                                    |f|
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                        &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_modifyAtIndices(),
                                                                                                                                                                                                                                                              &&&dictFoldable),
                                                                                                                                                                                                                           &&&is),
                                                                                                                                                                                        f))
                                                                                                            })
                                                                                        })))
    }
    pub fn Data_Array_NonEmpty_nub() -> &dyn Any {
        static Data_Array_NonEmpty_nub: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_nub.get_or_init(||
                                                &Func1::new(move |dictOrd|
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt(),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_nub(),
                                                                                                                                    dictOrd))))
    }
    pub fn Data_Array_NonEmpty_nubBy() -> &dyn Any {
        static Data_Array_NonEmpty_nubBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_nubBy.get_or_init(||
                                                  &Func1::new(move |f|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                      &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_nubBy(),
                                                                                                                                      f))))
    }
    pub fn Data_Array_NonEmpty_nubByEq() -> &dyn Any {
        static Data_Array_NonEmpty_nubByEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_nubByEq.get_or_init(||
                                                    &Func1::new(move |f|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                        &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_nubByEq(),
                                                                                                                                        f))))
    }
    pub fn Data_Array_NonEmpty_nubEq() -> &dyn Any {
        static Data_Array_NonEmpty_nubEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_nubEq.get_or_init(||
                                                  &Func1::new(move |dictEq|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt(),
                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_nubEq(),
                                                                                                                                      dictEq))))
    }
    pub fn Data_Array_NonEmpty_reverse() -> &dyn Any {
        static Data_Array_NonEmpty_reverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_reverse.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt(),
                                                                                     &&&PureScript_Data_Array::Data_Array_reverse()))
    }
    pub fn Data_Array_NonEmpty_scanl() -> &dyn Any {
        static Data_Array_NonEmpty_scanl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_scanl.get_or_init(||
                                                  &Func1::new(move |f|
                                                                  &Func1::new({
                                                                                  let f
                                                                                      =
                                                                                      f.clone();
                                                                                  move
                                                                                      |x|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_scanl(),
                                                                                                                                                                                             &&&f),
                                                                                                                                                          x))
                                                                              })))
    }
    pub fn Data_Array_NonEmpty_scanr() -> &dyn Any {
        static Data_Array_NonEmpty_scanr: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_scanr.get_or_init(||
                                                  &Func1::new(move |f|
                                                                  &Func1::new({
                                                                                  let f
                                                                                      =
                                                                                      f.clone();
                                                                                  move
                                                                                      |x|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                          &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_scanr(),
                                                                                                                                                                                             &&&f),
                                                                                                                                                          x))
                                                                              })))
    }
    pub fn Data_Array_NonEmpty_sort() -> &dyn Any {
        static Data_Array_NonEmpty_sort: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_sort.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt(),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sort(),
                                                                                                                                     dictOrd))))
    }
    pub fn Data_Array_NonEmpty_sortBy() -> &dyn Any {
        static Data_Array_NonEmpty_sortBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_sortBy.get_or_init(||
                                                   &Func1::new(move |f|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                       &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sortBy(),
                                                                                                                                       f))))
    }
    pub fn Data_Array_NonEmpty_sortWith() -> &dyn Any {
        static Data_Array_NonEmpty_sortWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_sortWith.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictOrd|
                                                                     &Func1::new({
                                                                                     let dictOrd
                                                                                         =
                                                                                         dictOrd.clone();
                                                                                     move
                                                                                         |f|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                             &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_sortWith(),
                                                                                                                                                                                                &&&dictOrd),
                                                                                                                                                             f))
                                                                                 })))
    }
    pub fn Data_Array_NonEmpty_updateAtIndices() -> &dyn Any {
        static Data_Array_NonEmpty_updateAtIndices: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Array_NonEmpty_updateAtIndices.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictFoldable|
                                                                            &Func1::new({
                                                                                            let dictFoldable
                                                                                                =
                                                                                                dictFoldable.clone();
                                                                                            move
                                                                                                |pairs|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                    &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeAdapt()),
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array::Data_Array_updateAtIndices(),
                                                                                                                                                                                                       &&&dictFoldable),
                                                                                                                                                                    pairs))
                                                                                        })))
    }
    pub fn Data_Array_NonEmpty_unsafeIndex() -> &dyn Any {
        static Data_Array_NonEmpty_unsafeIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_unsafeIndex.get_or_init(||
                                                        &Func1::new(move
                                                                        |usd__unused|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_adaptAny(),
                                                                                                         &&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeIndex1())))
    }
    pub fn Data_Array_NonEmpty_toUnfoldable1() -> &dyn Any {
        static Data_Array_NonEmpty_toUnfoldable1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Array_NonEmpty_toUnfoldable1.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictUnfoldable1|
                                                                          &Func1::new({
                                                                                          let dictUnfoldable1
                                                                                              =
                                                                                              dictUnfoldable1.clone();
                                                                                          move
                                                                                              |xs|
                                                                                              {
                                                                                                  let len =
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_length(),
                                                                                                                                       xs);
                                                                                                  let f =
                                                                                                      &Func1::new({
                                                                                                                      let len
                                                                                                                          =
                                                                                                                          len.clone();
                                                                                                                      let xs
                                                                                                                          =
                                                                                                                          xs.clone();
                                                                                                                      move
                                                                                                                          |i|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                                   |usd__arg1|
                                                                                                                                                                                                                                                   Func1::new({
                                                                                                                                                                                                                                                                  let usd__arg1
                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                      usd__arg1.clone();
                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                      |usd__arg2|
                                                                                                                                                                                                                                                                      &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                              usd__arg2.clone()))
                                                                                                                                                                                                                                                              })),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                                                                                                                                                          &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                            |usd__unused|
                                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Array_NonEmpty::Data_Array_NonEmpty_unsafeIndex(),
                                                                                                                                                                                                                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                                                                                                                                                                       &&&xs),
                                                                                                                                                                                                                                                                    i))),
                                                                                                                                                           &&{
                                                                                                                                                                 let matchValue =
                                                                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                                  i),
                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                     &&&len),
                                                                                                                                                                                                                                                                  &&&1_i32)));
                                                                                                                                                                 match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                  &matchValue)
                                                                                                                                                                     {
                                                                                                                                                                     0_i32
                                                                                                                                                                     =>
                                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                i),
                                                                                                                                                                                                                                                             &&&1_i32))),
                                                                                                                                                                     _
                                                                                                                                                                     =>
                                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                                                                 }
                                                                                                                                                             })
                                                                                                                  });
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable1::Data_Unfoldable1_unfoldr1(),
                                                                                                                                                                                                         &&&dictUnfoldable1),
                                                                                                                                                                      &&&f),
                                                                                                                                   &&&0_i32)
                                                                                              }
                                                                                      })))
    }
}

pub mod PureScript_Data_List_NonEmpty {
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
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_d662adf2::PureScript_Data_List_Types;
    use crate::module_d662adf2::PureScript_Data_List_Types::Data_List_Types_List;
    use crate::module_843b47b7::PureScript_Data_List;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_abab3d09::PureScript_Data_Semigroup_Traversable;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_e7bd458d::PureScript_Data_Tuple;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_List_NonEmpty_identity() -> &dyn Any {
        static Data_List_NonEmpty_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_identity.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                     &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_List_NonEmpty_zipWith() -> &dyn Any {
        static Data_List_NonEmpty_zipWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_zipWith.get_or_init(||
                                                   &Func1::new(move |f|
                                                                   &Func1::new({
                                                                                   let f
                                                                                       =
                                                                                       f.clone();
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
                                                                                                                   Sharpurs_Prelude::unbox(&&f);
                                                                                                               let matchValue_1:
                                                                                                                       LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                               let matchValue_2:
                                                                                                                       LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                   Sharpurs_Prelude::unbox(v1);
                                                                                                               let f1 =
                                                                                                                   matchValue;
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                 Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                                                       &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                          }),
                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_zipWith(),
                                                                                                                                                                                                                                                                                                                             &&&f1),
                                                                                                                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                 Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                                                       &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                          }))))
                                                                                                           }
                                                                                                   })
                                                                               })))
    }
    pub fn Data_List_NonEmpty_zipWithA() -> &dyn Any {
        static Data_List_NonEmpty_zipWithA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_zipWithA.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictApplicative|
                                                                    {
                                                                        let Apply0 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                    Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                             &&&Sharpurs_Prelude::Prim_undefined());
                                                                        &Func1::new({
                                                                                        let Apply0
                                                                                            =
                                                                                            Apply0.clone();
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
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Traversable::Data_Semigroup_Traversable_sequence1(),
                                                                                                                                                                                                                                           &&&PureScript_Data_List_Types::Data_List_Types_traversable1NonEmptyList()),
                                                                                                                                                                                                        &&&Apply0),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_zipWith(),
                                                                                                                                                                                                                                                                              &&&f),
                                                                                                                                                                                                                                           &&&xs),
                                                                                                                                                                                                        ys))
                                                                                                                            })
                                                                                                        })
                                                                                    })
                                                                    }))
    }
    pub fn Data_List_NonEmpty_zip() -> &dyn Any {
        static Data_List_NonEmpty_zip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_zip.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_zipWith(),
                                                                                &&&Func1::new(move
                                                                                                  |usd__arg1|
                                                                                                  Func1::new({
                                                                                                                 let usd__arg1
                                                                                                                     =
                                                                                                                     usd__arg1.clone();
                                                                                                                 move
                                                                                                                     |usd__arg2|
                                                                                                                     &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                             usd__arg2.clone()))
                                                                                                             }))))
    }
    pub fn Data_List_NonEmpty_wrappedOperation2() -> &dyn Any {
        static Data_List_NonEmpty_wrappedOperation2: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_NonEmpty_wrappedOperation2.get_or_init(||
                                                             &Func1::new(move
                                                                             |name|
                                                                             &Func1::new({
                                                                                             let name
                                                                                                 =
                                                                                                 name.clone();
                                                                                             move
                                                                                                 |f|
                                                                                                 &Func1::new({
                                                                                                                 let f
                                                                                                                     =
                                                                                                                     f.clone();
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
                                                                                                                                                 Sharpurs_Prelude::unbox(&&name);
                                                                                                                                             let matchValue_1 =
                                                                                                                                                 Sharpurs_Prelude::unbox(&&f);
                                                                                                                                             let matchValue_2:
                                                                                                                                                     LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                                             let matchValue_3:
                                                                                                                                                     LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                                                             let matchValue_5:
                                                                                                                                                     LrcPtr<Data_List_Types_List> =
                                                                                                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                               &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                      Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                  },
                                                                                                                                                                                                                                                                                                                 &match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                      Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                  }))),
                                                                                                                                                                                                            &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               },
                                                                                                                                                                                                                                                                              &match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               }))));
                                                                                                                                             match matchValue_5.as_ref()
                                                                                                                                                 {
                                                                                                                                                 Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                                                 =>
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafeCrashWith(),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                        &&&string("Impossible: empty list in NonEmptyList ")),
                                                                                                                                                                                                                     &&&matchValue)),
                                                                                                                                                 Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_5_1_0,
                                                                                                                                                                                                    matchValue_5_1_1)
                                                                                                                                                 =>
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                                                  &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue_5.as_ref()
                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                        &match matchValue_5.as_ref()
                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                             Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                         }))),
                                                                                                                                             }
                                                                                                                                         }
                                                                                                                                 })
                                                                                                             })
                                                                                         })))
    }
    pub fn Data_List_NonEmpty_wrappedOperation() -> &dyn Any {
        static Data_List_NonEmpty_wrappedOperation: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_NonEmpty_wrappedOperation.get_or_init(||
                                                            &Func1::new(move
                                                                            |name|
                                                                            &Func1::new({
                                                                                            let name
                                                                                                =
                                                                                                name.clone();
                                                                                            move
                                                                                                |f|
                                                                                                &Func1::new({
                                                                                                                let f
                                                                                                                    =
                                                                                                                    f.clone();
                                                                                                                move
                                                                                                                    |v|
                                                                                                                    {
                                                                                                                        let matchValue =
                                                                                                                            Sharpurs_Prelude::unbox(&&name);
                                                                                                                        let matchValue_1 =
                                                                                                                            Sharpurs_Prelude::unbox(&&f);
                                                                                                                        let matchValue_2:
                                                                                                                                LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                        let matchValue_4:
                                                                                                                                LrcPtr<Data_List_Types_List> =
                                                                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                       &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                         &match matchValue_2.as_ref()
                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                          }))));
                                                                                                                        match matchValue_4.as_ref()
                                                                                                                            {
                                                                                                                            Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                                            =>
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafeCrashWith(),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                   &&&string("Impossible: empty list in NonEmptyList ")),
                                                                                                                                                                                                &&&matchValue)),
                                                                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_4_1_0,
                                                                                                                                                                               matchValue_4_1_1)
                                                                                                                            =>
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                             &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue_4.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                    },
                                                                                                                                                                                                                                   &match matchValue_4.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                    }))),
                                                                                                                        }
                                                                                                                    }
                                                                                                            })
                                                                                        })))
    }
    pub fn Data_List_NonEmpty_updateAt() -> &dyn Any {
        static Data_List_NonEmpty_updateAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_updateAt.get_or_init(||
                                                    &Func1::new(move |i|
                                                                    &Func1::new({
                                                                                    let i
                                                                                        =
                                                                                        i.clone();
                                                                                    move
                                                                                        |a|
                                                                                        &Func1::new({
                                                                                                        let a
                                                                                                            =
                                                                                                            a.clone();
                                                                                                        move
                                                                                                            |v|
                                                                                                            {
                                                                                                                let matchValue =
                                                                                                                    Sharpurs_Prelude::unbox(&&i);
                                                                                                                let matchValue_1 =
                                                                                                                    Sharpurs_Prelude::unbox(&&a);
                                                                                                                let matchValue_2:
                                                                                                                        LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                             &&&0_i32))
                                                                                                                   {
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                                                                            &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&matchValue_1,
                                                                                                                                                                                                                                                                                  &match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                       Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                   })))))
                                                                                                                } else {
                                                                                                                    if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                       {
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                               &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList()),
                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                 let matchValue_2
                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                     matchValue_2.clone();
                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                     |v1|
                                                                                                                                                                                                                                                     &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                         v1.clone()))
                                                                                                                                                                                                                                             }))),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_updateAt(),
                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                                                                                     &&&1_i32)),
                                                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                                                            &&&match matchValue_2.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                               }))
                                                                                                                    } else {
                                                                                                                        panic!("{}",
                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.NonEmpty.fs"),
                                  Data1: 19_i32,
                                  Data2: 108_i32,}).get_Message(),)
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                    })
                                                                                })))
    }
    pub fn Data_List_NonEmpty_unzip() -> &dyn Any {
        static Data_List_NonEmpty_unzip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_unzip.get_or_init(||
                                                 &Func1::new(move |ts|
                                                                 &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                &&&PureScript_Data_List_Types::Data_List_Types_functorNonEmptyList()),
                                                                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_fst()),
                                                                                                                                                          ts),
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                &&&PureScript_Data_List_Types::Data_List_Types_functorNonEmptyList()),
                                                                                                                                                                                             &&&PureScript_Data_Tuple::Data_Tuple_snd()),
                                                                                                                                                          ts)))))
    }
    pub fn Data_List_NonEmpty_unsnoc() -> &dyn Any {
        static Data_List_NonEmpty_unsnoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_unsnoc.get_or_init(||
                                                  &Func1::new(move |v|
                                                                  {
                                                                      let matchValue:
                                                                              LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                          Sharpurs_Prelude::unbox(v);
                                                                      let x =
                                                                          match matchValue.as_ref()
                                                                              {
                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                     _)
                                                                              =>
                                                                              x.clone(),
                                                                          };
                                                                      let matchValue_1:
                                                                              LrcPtr<Data_Maybe_Maybe> =
                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_unsnoc(),
                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                            {
                                                                                                                                            Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                   x)
                                                                                                                                            =>
                                                                                                                                            x.clone(),
                                                                                                                                        }));
                                                                      match matchValue_1.as_ref()
                                                                          {
                                                                          Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                          => {
                                                                              let un =
                                                                                  matchValue_1_1_0.clone();
                                                                              &add(string("init"),
                                                                                   &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&x,
                                                                                                                                                    find(string("init"),
                                                                                                                                                         Sharpurs_Prelude::unbox(&&un)))),
                                                                                   add(string("last"),
                                                                                       &find(string("last"),
                                                                                             Sharpurs_Prelude::unbox(&&un)),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))
                                                                          }
                                                                          _ =>
                                                                          &add(string("init"),
                                                                               &&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor),
                                                                               add(string("last"),
                                                                                   &&x,
                                                                                   empty::<string,
                                                                                           &dyn Any>())),
                                                                      }
                                                                  }))
    }
    pub fn Data_List_NonEmpty_unionBy() -> &dyn Any {
        static Data_List_NonEmpty_unionBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_unionBy.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation2(),
                                                                                                                                                          &&&string("unionBy"))),
                                                                                    &&&PureScript_Data_List::Data_List_unionBy()))
    }
    pub fn Data_List_NonEmpty_union() -> &dyn Any {
        static Data_List_NonEmpty_union: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_union.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation2(),
                                                                                                                                     &&&string("union")),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_union(),
                                                                                                                                     dictEq))))
    }
    pub fn Data_List_NonEmpty_uncons() -> &dyn Any {
        static Data_List_NonEmpty_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_uncons.get_or_init(||
                                                  &Func1::new(move |v|
                                                                  {
                                                                      let matchValue:
                                                                              LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                          Sharpurs_Prelude::unbox(v);
                                                                      &add(string("head"),
                                                                           &&match matchValue.as_ref()
                                                                                 {
                                                                                 Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                        _)
                                                                                 =>
                                                                                 x.clone(),
                                                                             },
                                                                           add(string("tail"),
                                                                               &&match matchValue.as_ref()
                                                                                     {
                                                                                     Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                            x)
                                                                                     =>
                                                                                     x.clone(),
                                                                                 },
                                                                               empty::<string,
                                                                                       &dyn Any>()))
                                                                  }))
    }
    pub fn Data_List_NonEmpty_toList() -> &dyn Any {
        static Data_List_NonEmpty_toList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_toList.get_or_init(||
                                                  &Func1::new(move |v|
                                                                  {
                                                                      let matchValue:
                                                                              LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                          Sharpurs_Prelude::unbox(v);
                                                                      &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                           {
                                                                                                                                           Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                  _)
                                                                                                                                           =>
                                                                                                                                           x.clone(),
                                                                                                                                       },
                                                                                                                                      &match matchValue.as_ref()
                                                                                                                                           {
                                                                                                                                           Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                  x)
                                                                                                                                           =>
                                                                                                                                           x.clone(),
                                                                                                                                       }))
                                                                  }))
    }
    pub fn Data_List_NonEmpty_toUnfoldable() -> &dyn Any {
        static Data_List_NonEmpty_toUnfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_toUnfoldable.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictUnfoldable|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                                                                  dictUnfoldable),
                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                 |xs|
                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                       |rec_var|
                                                                                                                                                                                                                                                                                       &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(find(string("head"),
                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(rec_var)),
                                                                                                                                                                                                                                                                                                                                               find(string("tail"),
                                                                                                                                                                                                                                                                                                                                                    Sharpurs_Prelude::unbox(rec_var)))))),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_uncons(),
                                                                                                                                                                                                                                                                     xs))))),
                                                                                                         &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_toList())))
    }
    pub fn Data_List_NonEmpty_tail() -> &dyn Any {
        static Data_List_NonEmpty_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_tail.get_or_init(||
                                                &Func1::new(move |v|
                                                                &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                     {
                                                                     Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                            x)
                                                                     =>
                                                                     x.clone(),
                                                                 }))
    }
    pub fn Data_List_NonEmpty_sortBy() -> &dyn Any {
        static Data_List_NonEmpty_sortBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_sortBy.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                                         &&&string("sortBy"))),
                                                                                   &&&PureScript_Data_List::Data_List_sortBy()))
    }
    pub fn Data_List_NonEmpty_sort() -> &dyn Any {
        static Data_List_NonEmpty_sort: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_sort.get_or_init(||
                                                &Func1::new(move |dictOrd|
                                                                {
                                                                    let compare =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                         dictOrd);
                                                                    &Func1::new({
                                                                                    let compare
                                                                                        =
                                                                                        compare.clone();
                                                                                    move
                                                                                        |xs|
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_sortBy(),
                                                                                                                                                            &&&compare),
                                                                                                                         xs)
                                                                                })
                                                                }))
    }
    pub fn Data_List_NonEmpty_snoc() -> &dyn Any {
        static Data_List_NonEmpty_snoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_snoc.get_or_init(||
                                                &Func1::new(move |v|
                                                                &Func1::new({
                                                                                let v
                                                                                    =
                                                                                    v.clone();
                                                                                move
                                                                                    |y|
                                                                                    {
                                                                                        let matchValue:
                                                                                                LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                                        let matchValue_1 =
                                                                                            Sharpurs_Prelude::unbox(y);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                         &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                    Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                    =>
                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                },
                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_snoc(),
                                                                                                                                                                                                                                                                   &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                          Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                &&&matchValue_1))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_List_NonEmpty_singleton() -> &dyn Any {
        static Data_List_NonEmpty_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_singleton.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList()),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_NonEmpty::Data_NonEmpty_singleton(),
                                                                                                                         &&&PureScript_Data_List_Types::Data_List_Types_plusList())))
    }
    pub fn Data_List_NonEmpty_snoc_prime() -> &dyn Any {
        static Data_List_NonEmpty_snoc_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_snoc_prime.get_or_init(||
                                                      &Func1::new(move |v|
                                                                      &Func1::new({
                                                                                      let v
                                                                                          =
                                                                                          v.clone();
                                                                                      move
                                                                                          |v1|
                                                                                          {
                                                                                              let matchValue:
                                                                                                      LrcPtr<Data_List_Types_List> =
                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                              let matchValue_1 =
                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                              match matchValue.as_ref()
                                                                                                  {
                                                                                                  Data_List_Types_List::Data_List_Types_Nilusd_Ctor
                                                                                                  =>
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_singleton(),
                                                                                                                                   &&&matchValue_1),
                                                                                                  Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                     matchValue_1_1)
                                                                                                  =>
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                   &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                              {
                                                                                                                                                                                                              Data_List_Types_List::Data_List_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                              _
                                                                                                                                                                                                              =>
                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                          },
                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_snoc(),
                                                                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                    Data_List_Types_List::Data_List_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                                }),
                                                                                                                                                                                                                                          &&&matchValue_1)))),
                                                                                              }
                                                                                          }
                                                                                  })))
    }
    pub fn Data_List_NonEmpty_reverse() -> &dyn Any {
        static Data_List_NonEmpty_reverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_reverse.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                       &&&string("reverse")),
                                                                                    &&&PureScript_Data_List::Data_List_reverse()))
    }
    pub fn Data_List_NonEmpty_nubEq() -> &dyn Any {
        static Data_List_NonEmpty_nubEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_nubEq.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                     &&&string("nubEq")),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_nubEq(),
                                                                                                                                     dictEq))))
    }
    pub fn Data_List_NonEmpty_nubByEq() -> &dyn Any {
        static Data_List_NonEmpty_nubByEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_nubByEq.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                                          &&&string("nubByEq"))),
                                                                                    &&&PureScript_Data_List::Data_List_nubByEq()))
    }
    pub fn Data_List_NonEmpty_nubBy() -> &dyn Any {
        static Data_List_NonEmpty_nubBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_nubBy.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                                        &&&string("nubBy"))),
                                                                                  &&&PureScript_Data_List::Data_List_nubBy()))
    }
    pub fn Data_List_NonEmpty_nub() -> &dyn Any {
        static Data_List_NonEmpty_nub: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_nub.get_or_init(||
                                               &Func1::new(move |dictOrd|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                   &&&string("nub")),
                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_nub(),
                                                                                                                                   dictOrd))))
    }
    pub fn Data_List_NonEmpty_modifyAt() -> &dyn Any {
        static Data_List_NonEmpty_modifyAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_modifyAt.get_or_init(||
                                                    &Func1::new(move |i|
                                                                    &Func1::new({
                                                                                    let i
                                                                                        =
                                                                                        i.clone();
                                                                                    move
                                                                                        |f|
                                                                                        &Func1::new({
                                                                                                        let f
                                                                                                            =
                                                                                                            f.clone();
                                                                                                        move
                                                                                                            |v|
                                                                                                            {
                                                                                                                let matchValue =
                                                                                                                    Sharpurs_Prelude::unbox(&&i);
                                                                                                                let matchValue_1 =
                                                                                                                    Sharpurs_Prelude::unbox(&&f);
                                                                                                                let matchValue_2:
                                                                                                                        LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                             &&&0_i32))
                                                                                                                   {
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                                                                            &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                   &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                  &match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                       Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                   })))))
                                                                                                                } else {
                                                                                                                    if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                       {
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                               &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList()),
                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                 let matchValue_2
                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                     matchValue_2.clone();
                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                     |v1|
                                                                                                                                                                                                                                                     &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                         v1.clone()))
                                                                                                                                                                                                                                             }))),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_modifyAt(),
                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                                                                                     &&&1_i32)),
                                                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                                                            &&&match matchValue_2.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                               }))
                                                                                                                    } else {
                                                                                                                        panic!("{}",
                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.NonEmpty.fs"),
                                  Data1: 57_i32,
                                  Data2: 108_i32,}).get_Message(),)
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                    })
                                                                                })))
    }
    pub fn Data_List_NonEmpty_lift() -> &dyn Any {
        static Data_List_NonEmpty_lift: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_lift.get_or_init(||
                                                &Func1::new(move |f|
                                                                &Func1::new({
                                                                                let f
                                                                                    =
                                                                                    f.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&f);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                         &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                            },
                                                                                                                                                                                           &match matchValue_1.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                =>
                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                            })))
                                                                                    }
                                                                            })))
    }
    pub fn Data_List_NonEmpty_mapMaybe() -> &dyn Any {
        static Data_List_NonEmpty_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_mapMaybe.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                        &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                     &&&PureScript_Data_List::Data_List_mapMaybe()))
    }
    pub fn Data_List_NonEmpty_partition() -> &dyn Any {
        static Data_List_NonEmpty_partition: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_partition.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                      &&&PureScript_Data_List::Data_List_partition()))
    }
    pub fn Data_List_NonEmpty_span() -> &dyn Any {
        static Data_List_NonEmpty_span: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_span.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                    &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                 &&&PureScript_Data_List::Data_List_span()))
    }
    pub fn Data_List_NonEmpty_take() -> &dyn Any {
        static Data_List_NonEmpty_take: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_take.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                    &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                 &&&PureScript_Data_List::Data_List_take()))
    }
    pub fn Data_List_NonEmpty_takeWhile() -> &dyn Any {
        static Data_List_NonEmpty_takeWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_takeWhile.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                      &&&PureScript_Data_List::Data_List_takeWhile()))
    }
    pub fn Data_List_NonEmpty_length() -> &dyn Any {
        static Data_List_NonEmpty_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_length.get_or_init(||
                                                  &Func1::new(move |v|
                                                                  {
                                                                      let matchValue:
                                                                              LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                          Sharpurs_Prelude::unbox(v);
                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                             &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                          &&&1_i32),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_length(),
                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                 {
                                                                                                                                                 Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                        x)
                                                                                                                                                 =>
                                                                                                                                                 x.clone(),
                                                                                                                                             }))
                                                                  }))
    }
    pub fn Data_List_NonEmpty_last() -> &dyn Any {
        static Data_List_NonEmpty_last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_last.get_or_init(||
                                                &Func1::new(move |v|
                                                                {
                                                                    let matchValue:
                                                                            LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                        Sharpurs_Prelude::unbox(v);
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromMaybe(),
                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                               {
                                                                                                                                               Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                      _)
                                                                                                                                               =>
                                                                                                                                               x.clone(),
                                                                                                                                           }),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_last(),
                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                               {
                                                                                                                                               Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                      x)
                                                                                                                                               =>
                                                                                                                                               x.clone(),
                                                                                                                                           }))
                                                                }))
    }
    pub fn Data_List_NonEmpty_intersectBy() -> &dyn Any {
        static Data_List_NonEmpty_intersectBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_intersectBy.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation2(),
                                                                                                                                                              &&&string("intersectBy"))),
                                                                                        &&&PureScript_Data_List::Data_List_intersectBy()))
    }
    pub fn Data_List_NonEmpty_intersect() -> &dyn Any {
        static Data_List_NonEmpty_intersect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_intersect.get_or_init(||
                                                     &Func1::new(move |dictEq|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation2(),
                                                                                                                                         &&&string("intersect")),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_intersect(),
                                                                                                                                         dictEq))))
    }
    pub fn Data_List_NonEmpty_insertAt() -> &dyn Any {
        static Data_List_NonEmpty_insertAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_insertAt.get_or_init(||
                                                    &Func1::new(move |i|
                                                                    &Func1::new({
                                                                                    let i
                                                                                        =
                                                                                        i.clone();
                                                                                    move
                                                                                        |a|
                                                                                        &Func1::new({
                                                                                                        let a
                                                                                                            =
                                                                                                            a.clone();
                                                                                                        move
                                                                                                            |v|
                                                                                                            {
                                                                                                                let matchValue =
                                                                                                                    Sharpurs_Prelude::unbox(&&i);
                                                                                                                let matchValue_1 =
                                                                                                                    Sharpurs_Prelude::unbox(&&a);
                                                                                                                let matchValue_2:
                                                                                                                        LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                             &&&0_i32))
                                                                                                                   {
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                                                                            &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&matchValue_1,
                                                                                                                                                                                                                                                                                  &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                       Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                                   },
                                                                                                                                                                                                                                                                                                                                                  &match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                                       Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                                   })))))))
                                                                                                                } else {
                                                                                                                    if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                       {
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                               &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList()),
                                                                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                                                                 let matchValue_2
                                                                                                                                                                                                                                                     =
                                                                                                                                                                                                                                                     matchValue_2.clone();
                                                                                                                                                                                                                                                 move
                                                                                                                                                                                                                                                     |v1|
                                                                                                                                                                                                                                                     &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                          },
                                                                                                                                                                                                                                                                                                                         v1.clone()))
                                                                                                                                                                                                                                             }))),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_insertAt(),
                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                                                                                     &&&1_i32)),
                                                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                                                            &&&match matchValue_2.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                               }))
                                                                                                                    } else {
                                                                                                                        panic!("{}",
                                                                                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.NonEmpty.fs"),
                                  Data1: 79_i32,
                                  Data2: 108_i32,}).get_Message(),)
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                    })
                                                                                })))
    }
    pub fn Data_List_NonEmpty_init() -> &dyn Any {
        static Data_List_NonEmpty_init: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_init.get_or_init(||
                                                &Func1::new(move |v|
                                                                {
                                                                    let matchValue:
                                                                            LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                        Sharpurs_Prelude::unbox(v);
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                           &&&LrcPtr::new(Data_List_Types_List::Data_List_Types_Nilusd_Ctor)),
                                                                                                                                        &&&Func1::new(move
                                                                                                                                                          |v1|
                                                                                                                                                          &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                           },
                                                                                                                                                                                                                          v1.clone())))),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_init(),
                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                               {
                                                                                                                                               Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                      x)
                                                                                                                                               =>
                                                                                                                                               x.clone(),
                                                                                                                                           }))
                                                                }))
    }
    pub fn Data_List_NonEmpty_index() -> &dyn Any {
        static Data_List_NonEmpty_index: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_index.get_or_init(||
                                                 &Func1::new(move |v|
                                                                 &Func1::new({
                                                                                 let v
                                                                                     =
                                                                                     v.clone();
                                                                                 move
                                                                                     |i|
                                                                                     {
                                                                                         let matchValue:
                                                                                                 LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                         let matchValue_1 =
                                                                                             Sharpurs_Prelude::unbox(i);
                                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                            &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                         &&&matchValue_1),
                                                                                                                                                      &&&0_i32))
                                                                                            {
                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                         {
                                                                                                                                                         Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                _)
                                                                                                                                                         =>
                                                                                                                                                         x.clone(),
                                                                                                                                                     }))
                                                                                         } else {
                                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                {
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_index(),
                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                            Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                   x)
                                                                                                                                                                            =>
                                                                                                                                                                            x.clone(),
                                                                                                                                                                        }),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                        &&&matchValue_1),
                                                                                                                                                                     &&&1_i32))
                                                                                             } else {
                                                                                                 panic!("{}",
                                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.NonEmpty.fs"),
                                  Data1: 83_i32,
                                  Data2: 83_i32,}).get_Message(),)
                                                                                             }
                                                                                         }
                                                                                     }
                                                                             })))
    }
    pub fn Data_List_NonEmpty_head() -> &dyn Any {
        static Data_List_NonEmpty_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_head.get_or_init(||
                                                &Func1::new(move |v|
                                                                &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                     {
                                                                     Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                            _)
                                                                     =>
                                                                     x.clone(),
                                                                 }))
    }
    pub fn Data_List_NonEmpty_groupBy() -> &dyn Any {
        static Data_List_NonEmpty_groupBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_groupBy.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                                          &&&string("groupBy"))),
                                                                                    &&&PureScript_Data_List::Data_List_groupBy()))
    }
    pub fn Data_List_NonEmpty_groupAllBy() -> &dyn Any {
        static Data_List_NonEmpty_groupAllBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_groupAllBy.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                                             &&&string("groupAllBy"))),
                                                                                       &&&PureScript_Data_List::Data_List_groupAllBy()))
    }
    pub fn Data_List_NonEmpty_groupAll() -> &dyn Any {
        static Data_List_NonEmpty_groupAll: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_groupAll.get_or_init(||
                                                    &Func1::new(move |dictOrd|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                        &&&string("groupAll")),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_groupAll(),
                                                                                                                                        dictOrd))))
    }
    pub fn Data_List_NonEmpty_group() -> &dyn Any {
        static Data_List_NonEmpty_group: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_group.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_wrappedOperation(),
                                                                                                                                     &&&string("group")),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_group(),
                                                                                                                                     dictEq))))
    }
    pub fn Data_List_NonEmpty_fromList() -> &dyn Any {
        static Data_List_NonEmpty_fromList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_fromList.get_or_init(||
                                                    &Func1::new(move |v|
                                                                    {
                                                                        let matchValue:
                                                                                LrcPtr<Data_List_Types_List> =
                                                                            Sharpurs_Prelude::unbox(v);
                                                                        match matchValue.as_ref()
                                                                            {
                                                                            Data_List_Types_List::Data_List_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                               matchValue_1_1)
                                                                            =>
                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                                    &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                          matchValue_1_1))))),
                                                                            _
                                                                            =>
                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                        }
                                                                    }))
    }
    pub fn Data_List_NonEmpty_fromFoldable() -> &dyn Any {
        static Data_List_NonEmpty_fromFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_fromFoldable.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictFoldable|
                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                            &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_fromList()),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_fromFoldable(),
                                                                                                                                            dictFoldable))))
    }
    pub fn Data_List_NonEmpty_foldM() -> &dyn Any {
        static Data_List_NonEmpty_foldM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_foldM.get_or_init(||
                                                 &Func1::new(move |dictMonad|
                                                                 {
                                                                     let Bind1 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                                                                                 Sharpurs_Prelude::unbox(dictMonad)),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let Bind1
                                                                                         =
                                                                                         Bind1.clone();
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
                                                                                                             |b|
                                                                                                             &Func1::new({
                                                                                                                             let b
                                                                                                                                 =
                                                                                                                                 b.clone();
                                                                                                                             move
                                                                                                                                 |v|
                                                                                                                                 {
                                                                                                                                     let matchValue =
                                                                                                                                         Sharpurs_Prelude::unbox(&&f);
                                                                                                                                     let matchValue_1 =
                                                                                                                                         Sharpurs_Prelude::unbox(&&b);
                                                                                                                                     let matchValue_2:
                                                                                                                                             LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                                     let f1 =
                                                                                                                                         matchValue;
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                            &&&Bind1),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                                                                                                            &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                   Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                               })),
                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                        let f1
                                                                                                                                                                                            =
                                                                                                                                                                                            f1.clone();
                                                                                                                                                                                        let matchValue_2
                                                                                                                                                                                            =
                                                                                                                                                                                            matchValue_2.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |b_prime|
                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_foldM(),
                                                                                                                                                                                                                                                                                                                                      &&&dictMonad),
                                                                                                                                                                                                                                                                                                   &&&f1),
                                                                                                                                                                                                                                                                b_prime),
                                                                                                                                                                                                                             &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                    Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                })
                                                                                                                                                                                    }))
                                                                                                                                 }
                                                                                                                         })
                                                                                                     })
                                                                                 })
                                                                 }))
    }
    pub fn Data_List_NonEmpty_findLastIndex() -> &dyn Any {
        static Data_List_NonEmpty_findLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_findLastIndex.get_or_init(||
                                                         &Func1::new(move |f|
                                                                         &Func1::new({
                                                                                         let f
                                                                                             =
                                                                                             f.clone();
                                                                                         move
                                                                                             |v|
                                                                                             {
                                                                                                 let matchValue =
                                                                                                     Sharpurs_Prelude::unbox(&&f);
                                                                                                 let matchValue_1:
                                                                                                         LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                 let f1 =
                                                                                                     matchValue;
                                                                                                 let matchValue_3:
                                                                                                         LrcPtr<Data_Maybe_Maybe> =
                                                                                                     Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_findLastIndex(),
                                                                                                                                                                                                   &&&f1),
                                                                                                                                                                &&&match matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                       Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                              x)
                                                                                                                                                                       =>
                                                                                                                                                                       x.clone(),
                                                                                                                                                                   }));
                                                                                                 if let Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                        =
                                                                                                        matchValue_3.as_ref()
                                                                                                    {
                                                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&f1,
                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                         Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                _)
                                                                                                                                                                         =>
                                                                                                                                                                         x.clone(),
                                                                                                                                                                     }))
                                                                                                        {
                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&0_i32))
                                                                                                     } else {
                                                                                                         if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                            {
                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                         } else {
                                                                                                             panic!("{}",
                                                                                                                    string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                                                                                                         }
                                                                                                     }
                                                                                                 } else {
                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                       _
                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                       unreachable!(),
                                                                                                                                                                                                                                   }),
                                                                                                                                                                                             &&&1_i32)))
                                                                                                 }
                                                                                             }
                                                                                     })))
    }
    pub fn Data_List_NonEmpty_findIndex() -> &dyn Any {
        static Data_List_NonEmpty_findIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_findIndex.get_or_init(||
                                                     &Func1::new(move |f|
                                                                     &Func1::new({
                                                                                     let f
                                                                                         =
                                                                                         f.clone();
                                                                                     move
                                                                                         |v|
                                                                                         {
                                                                                             let matchValue =
                                                                                                 Sharpurs_Prelude::unbox(&&f);
                                                                                             let matchValue_1:
                                                                                                     LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                 Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                        _)
                                                                                                                                                                 =>
                                                                                                                                                                 x.clone(),
                                                                                                                                                             }))
                                                                                                {
                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&0_i32))
                                                                                             } else {
                                                                                                 if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                    {
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                            &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                         &&&Func1::new(move
                                                                                                                                                                                           |v1|
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                               v1),
                                                                                                                                                                                                                            &&&1_i32))),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_findIndex(),
                                                                                                                                                                                                            &&&matchValue),
                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                       x)
                                                                                                                                                                                =>
                                                                                                                                                                                x.clone(),
                                                                                                                                                                            }))
                                                                                                 } else {
                                                                                                     panic!("{}",
                                                                                                            LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.NonEmpty.fs"),
                                  Data1: 103_i32,
                                  Data2: 87_i32,}).get_Message(),)
                                                                                                 }
                                                                                             }
                                                                                         }
                                                                                 })))
    }
    pub fn Data_List_NonEmpty_filterM() -> &dyn Any {
        static Data_List_NonEmpty_filterM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_filterM.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictMonad|
                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                       &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_filterM(),
                                                                                                                                       dictMonad))))
    }
    pub fn Data_List_NonEmpty_filter() -> &dyn Any {
        static Data_List_NonEmpty_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_filter.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                      &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                   &&&PureScript_Data_List::Data_List_filter()))
    }
    pub fn Data_List_NonEmpty_elemLastIndex() -> &dyn Any {
        static Data_List_NonEmpty_elemLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_elemLastIndex.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictEq|
                                                                         &Func1::new({
                                                                                         let dictEq
                                                                                             =
                                                                                             dictEq.clone();
                                                                                         move
                                                                                             |x|
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_findLastIndex(),
                                                                                                                              &&&Func1::new({
                                                                                                                                                let x
                                                                                                                                                    =
                                                                                                                                                    x.clone();
                                                                                                                                                move
                                                                                                                                                    |v|
                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                           &&&dictEq),
                                                                                                                                                                                                                        v),
                                                                                                                                                                                     &&&x)
                                                                                                                                            }))
                                                                                     })))
    }
    pub fn Data_List_NonEmpty_elemIndex() -> &dyn Any {
        static Data_List_NonEmpty_elemIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_elemIndex.get_or_init(||
                                                     &Func1::new(move |dictEq|
                                                                     &Func1::new({
                                                                                     let dictEq
                                                                                         =
                                                                                         dictEq.clone();
                                                                                     move
                                                                                         |x|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_findIndex(),
                                                                                                                          &&&Func1::new({
                                                                                                                                            let x
                                                                                                                                                =
                                                                                                                                                x.clone();
                                                                                                                                            move
                                                                                                                                                |v|
                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                       &&&dictEq),
                                                                                                                                                                                                                    v),
                                                                                                                                                                                 &&&x)
                                                                                                                                        }))
                                                                                 })))
    }
    pub fn Data_List_NonEmpty_dropWhile() -> &dyn Any {
        static Data_List_NonEmpty_dropWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_dropWhile.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                      &&&PureScript_Data_List::Data_List_dropWhile()))
    }
    pub fn Data_List_NonEmpty_drop() -> &dyn Any {
        static Data_List_NonEmpty_drop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_drop.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                    &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift()),
                                                                                 &&&PureScript_Data_List::Data_List_drop()))
    }
    pub fn Data_List_NonEmpty_cons_prime() -> &dyn Any {
        static Data_List_NonEmpty_cons_prime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_cons_prime.get_or_init(||
                                                      &Func1::new(move |x|
                                                                      &Func1::new({
                                                                                      let x
                                                                                          =
                                                                                          x.clone();
                                                                                      move
                                                                                          |xs|
                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                           &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&x,
                                                                                                                                                                                                 xs.clone())))
                                                                                  })))
    }
    pub fn Data_List_NonEmpty_cons() -> &dyn Any {
        static Data_List_NonEmpty_cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_cons.get_or_init(||
                                                &Func1::new(move |y|
                                                                &Func1::new({
                                                                                let y
                                                                                    =
                                                                                    y.clone();
                                                                                move
                                                                                    |v|
                                                                                    {
                                                                                        let matchValue =
                                                                                            Sharpurs_Prelude::unbox(&&y);
                                                                                        let matchValue_1:
                                                                                                LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                         &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&matchValue,
                                                                                                                                                                                               &LrcPtr::new(Data_List_Types_List::Data_List_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                    Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                },
                                                                                                                                                                                                                                                               &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                    Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                })))))
                                                                                    }
                                                                            })))
    }
    pub fn Data_List_NonEmpty_concatMap() -> &dyn Any {
        static Data_List_NonEmpty_concatMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_concatMap.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                         &&&PureScript_Data_List_Types::Data_List_Types_bindNonEmptyList())))
    }
    pub fn Data_List_NonEmpty_concat() -> &dyn Any {
        static Data_List_NonEmpty_concat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_concat.get_or_init(||
                                                  &Func1::new(move |v|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                         &&&PureScript_Data_List_Types::Data_List_Types_bindNonEmptyList()),
                                                                                                                                      v),
                                                                                                   &&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_identity())))
    }
    pub fn Data_List_NonEmpty_catMaybes() -> &dyn Any {
        static Data_List_NonEmpty_catMaybes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_catMaybes.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_NonEmpty::Data_List_NonEmpty_lift(),
                                                                                      &&&PureScript_Data_List::Data_List_catMaybes()))
    }
    pub fn Data_List_NonEmpty_appendFoldable() -> &dyn Any {
        static Data_List_NonEmpty_appendFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_NonEmpty_appendFoldable.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictFoldable|
                                                                          &Func1::new({
                                                                                          let dictFoldable
                                                                                              =
                                                                                              dictFoldable.clone();
                                                                                          move
                                                                                              |v|
                                                                                              &Func1::new({
                                                                                                              let v
                                                                                                                  =
                                                                                                                  v.clone();
                                                                                                              move
                                                                                                                  |ys|
                                                                                                                  {
                                                                                                                      let matchValue:
                                                                                                                              LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                      let matchValue_1 =
                                                                                                                          Sharpurs_Prelude::unbox(ys);
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Types::Data_List_Types_NonEmptyList(),
                                                                                                                                                       &&&LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                  Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                              },
                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_List_Types::Data_List_Types_semigroupList()),
                                                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                        Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List::Data_List_fromFoldable(),
                                                                                                                                                                                                                                                                                                                                    &&&dictFoldable),
                                                                                                                                                                                                                                                                                                 &&&matchValue_1)))))
                                                                                                                  }
                                                                                                          })
                                                                                      })))
    }
}

pub mod PureScript_Data_List_Lazy {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::Native_::fix1;
    use fable_library_rust::Option_::getValue;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class;
    use crate::module_e6eed311::PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_Step;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_720e12db::PureScript_Data_Lazy;
    use crate::module_4df90a9e::PureScript_Data_List_Internal;
    use crate::module_44df2f32::PureScript_Data_List_Lazy_Types;
    use crate::module_44df2f32::PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_List_Lazy_unwrap() -> &dyn Any {
        static Data_List_Lazy_unwrap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_unwrap.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                               &&&Sharpurs_Prelude::Prim_undefined()))
    }
    pub fn Data_List_Lazy_one() -> &dyn Any {
        static Data_List_Lazy_one: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_one.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                            &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()))
    }
    pub fn Data_List_Lazy_identity() -> &dyn Any {
        static Data_List_Lazy_identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_identity.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                 &&&PureScript_Control_Category::Control_Category_categoryFn()))
    }
    pub fn Data_List_Lazy_Pattern() -> &dyn Any {
        static Data_List_Lazy_Pattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_Pattern.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_List_Lazy_zipWith_004016() -> &dyn Any {
        &Func1::new(move |f|
                        Func1::new({
                                       let f = f.clone();
                                       move |xs|
                                           Func1::new({
                                                          let xs = xs.clone();
                                                          move |ys|
                                                              PureScript_Data_List_Lazy::Data_List_Lazy_zipWith_tco(&f,
                                                                                                                    &xs,
                                                                                                                    ys)
                                                      })
                                   }))
    }
    pub fn Data_List_Lazy_zipWith_004016_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_zipWith_004016_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_zipWith_004016_002d1.get_or_init(||
                                                            Lazy(Data_List_Lazy_zipWith_004016.clone()))
    }
    pub fn Data_List_Lazy_zipWith_tco(f: &dyn Any, xs: &dyn Any, ys: &dyn Any)
     -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_zipWith_004016_002d1 =
                                Data_List_Lazy_zipWith_004016_002d1.clone();
                            let f = f.clone();
                            move |v|
                                &Func1::new({
                                                let Data_List_Lazy_zipWith_004016_002d1
                                                    =
                                                    Data_List_Lazy_zipWith_004016_002d1.clone();
                                                let v = v.clone();
                                                move |v1|
                                                    {
                                                        let matchValue:
                                                                LrcPtr<Data_List_Lazy_Types_Step> =
                                                            Sharpurs_Prelude::unbox(&&v);
                                                        let matchValue_1:
                                                                LrcPtr<Data_List_Lazy_Types_Step> =
                                                            Sharpurs_Prelude::unbox(v1);
                                                        if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                            matchValue_1_1)
                                                               =
                                                               matchValue.as_ref()
                                                           {
                                                            if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                matchValue_1_1_1)
                                                                   =
                                                                   matchValue_1.as_ref()
                                                               {
                                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                     Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                     _
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                 }),
                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                  Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                               _)
                                                                                                                                                                                  =>
                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                  _
                                                                                                                                                                                  =>
                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                              }),
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_zipWith_004016_002d1.Value,
                                                                                                                                                                                                                                                 &&&f),
                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                     Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                     _
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                 }),
                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                  Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                               x)
                                                                                                                                                                                  =>
                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                  _
                                                                                                                                                                                  =>
                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                              })))
                                                            } else {
                                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                                            }
                                                        } else {
                                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                                        }
                                                    }
                                            })
                        });
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_applyLazy()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                        &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                     &&&go),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                     xs))),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                               ys)))
    }
    pub fn Data_List_Lazy_zipWith() -> &dyn Any {
        static Data_List_Lazy_zipWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_zipWith.get_or_init(||
                                               Data_List_Lazy_zipWith_004016_002d1.Value)
    }
    pub fn Data_List_Lazy_zipWithA() -> &dyn Any {
        static Data_List_Lazy_zipWithA: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_zipWithA.get_or_init(||
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
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequence(),
                                                                                                                                                                                                                                   &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_traversableList()),
                                                                                                                                                                                                &&&dictApplicative),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_zipWith(),
                                                                                                                                                                                                                                                                      &&&f),
                                                                                                                                                                                                                                   &&&xs),
                                                                                                                                                                                                ys))
                                                                                                                    })
                                                                                                })
                                                                            })))
    }
    pub fn Data_List_Lazy_zip() -> &dyn Any {
        static Data_List_Lazy_zip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_zip.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_zipWith(),
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
    pub fn Data_List_Lazy_updateAt_004024() -> &dyn Any {
        &Func1::new(move |n|
                        Func1::new({
                                       let n = n.clone();
                                       move |x|
                                           Func1::new({
                                                          let x = x.clone();
                                                          move |xs|
                                                              PureScript_Data_List_Lazy::Data_List_Lazy_updateAt_tco(&n,
                                                                                                                     &x,
                                                                                                                     xs)
                                                      })
                                   }))
    }
    pub fn Data_List_Lazy_updateAt_004024_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_updateAt_004024_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_updateAt_004024_002d1.get_or_init(||
                                                             Lazy(Data_List_Lazy_updateAt_004024.clone()))
    }
    pub fn Data_List_Lazy_updateAt_tco(n: &dyn Any, x: &dyn Any, xs: &dyn Any)
     -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_updateAt_004024_002d1 =
                                Data_List_Lazy_updateAt_004024_002d1.clone();
                            let x = x.clone();
                            move |v|
                                &Func1::new({
                                                let Data_List_Lazy_updateAt_004024_002d1
                                                    =
                                                    Data_List_Lazy_updateAt_004024_002d1.clone();
                                                let v = v.clone();
                                                move |v1|
                                                    {
                                                        let matchValue =
                                                            Sharpurs_Prelude::unbox(&&v);
                                                        let matchValue_1:
                                                                LrcPtr<Data_List_Lazy_Types_Step> =
                                                            Sharpurs_Prelude::unbox(v1);
                                                        if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                            matchValue_1_1_1)
                                                               =
                                                               matchValue_1.as_ref()
                                                           {
                                                            if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                        &matchValue).is_some()
                                                               {
                                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&x,
                                                                                                                                          &match matchValue_1.as_ref()
                                                                                                                                               {
                                                                                                                                               Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                            x)
                                                                                                                                               =>
                                                                                                                                               x.clone(),
                                                                                                                                               _
                                                                                                                                               =>
                                                                                                                                               unreachable!(),
                                                                                                                                           }))
                                                            } else {
                                                                if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                    matchValue_1_1_1)
                                                                       =
                                                                       matchValue_1.as_ref()
                                                                   {
                                                                    &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                _)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                                   _
                                                                                                                                                   =>
                                                                                                                                                   unreachable!(),
                                                                                                                                               },
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_updateAt_004024_002d1.Value,
                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                                                                        &&&1_i32)),
                                                                                                                                                                                                                  &&&x),
                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                      =>
                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                      _
                                                                                                                                                                                      =>
                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                  })))
                                                                } else {
                                                                    panic!("{}",
                                                                           string("Match failure"),)
                                                                }
                                                            }
                                                        } else {
                                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                                        }
                                                    }
                                            })
                        });
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                  n)),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                               xs)))
    }
    pub fn Data_List_Lazy_updateAt() -> &dyn Any {
        static Data_List_Lazy_updateAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_updateAt.get_or_init(||
                                                Data_List_Lazy_updateAt_004024_002d1.Value)
    }
    pub fn Data_List_Lazy_unzip() -> &dyn Any {
        static Data_List_Lazy_unzip: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_unzip.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                 &&&Func1::new(move
                                                                                                                                   |v|
                                                                                                                                   {
                                                                                                                                       let matchValue:
                                                                                                                                               LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                       &Func1::new(move
                                                                                                                                                       |v1|
                                                                                                                                                       {
                                                                                                                                                           let matchValue_1:
                                                                                                                                                                   LrcPtr<Data_Tuple_Tuple> =
                                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                                           &LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                          }),
                                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                              Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                          }),
                                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                           Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                       })))
                                                                                                                                                       })
                                                                                                                                   })),
                                                                              &&&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                                        &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()))))
    }
    pub fn Data_List_Lazy_uncons() -> &dyn Any {
        static Data_List_Lazy_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_uncons.get_or_init(||
                                              &Func1::new(move |xs|
                                                              {
                                                                  let matchValue:
                                                                          LrcPtr<Data_List_Lazy_Types_Step> =
                                                                      Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                 xs));
                                                                  match matchValue.as_ref()
                                                                      {
                                                                      Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                   matchValue_1_1)
                                                                      =>
                                                                      &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&add(string("head"),
                                                                                                                                  &matchValue_1_0,
                                                                                                                                  add(string("tail"),
                                                                                                                                      &matchValue_1_1,
                                                                                                                                      empty::<string,
                                                                                                                                              &dyn Any>())))),
                                                                      _ =>
                                                                      &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                  }
                                                              }))
    }
    pub fn Data_List_Lazy_toUnfoldable() -> &dyn Any {
        static Data_List_Lazy_toUnfoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_toUnfoldable.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictUnfoldable|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
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
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                                                                                                                           xs))))))
    }
    pub fn Data_List_Lazy_takeWhile_004034() -> &dyn Any {
        &Func1::new(move |p|
                        PureScript_Data_List_Lazy::Data_List_Lazy_takeWhile_tco(p))
    }
    pub fn Data_List_Lazy_takeWhile_004034_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_takeWhile_004034_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_takeWhile_004034_002d1.get_or_init(||
                                                              Lazy(Data_List_Lazy_takeWhile_004034.clone()))
    }
    pub fn Data_List_Lazy_takeWhile_tco(p: &dyn Any) -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_takeWhile_004034_002d1 =
                                Data_List_Lazy_takeWhile_004034_002d1.clone();
                            let p = p.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                        matchValue_1_1)
                                           = matchValue.as_ref() {
                                        if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                     &&&match matchValue.as_ref()
                                                                                                            {
                                                                                                            Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                         _)
                                                                                                            =>
                                                                                                            x.clone(),
                                                                                                            _
                                                                                                            =>
                                                                                                            unreachable!(),
                                                                                                        }))
                                           {
                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                           {
                                                                                                                           Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                        _)
                                                                                                                           =>
                                                                                                                           x.clone(),
                                                                                                                           _
                                                                                                                           =>
                                                                                                                           unreachable!(),
                                                                                                                       },
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_takeWhile_004034_002d1.Value,
                                                                                                                                                                                          &&&p),
                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                              {
                                                                                                                                                              Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                           x)
                                                                                                                                                              =>
                                                                                                                                                              x.clone(),
                                                                                                                                                              _
                                                                                                                                                              =>
                                                                                                                                                              unreachable!(),
                                                                                                                                                          })))
                                        } else {
                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                        }
                                    } else {
                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                     &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                  &&&go)),
                                                                            &&&PureScript_Data_List_Lazy::Data_List_Lazy_unwrap()))
    }
    pub fn Data_List_Lazy_takeWhile() -> &dyn Any {
        static Data_List_Lazy_takeWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_takeWhile.get_or_init(||
                                                 Data_List_Lazy_takeWhile_004034_002d1.Value)
    }
    pub fn Data_List_Lazy_take_004038() -> &dyn Any {
        &Func1::new(move |n|
                        PureScript_Data_List_Lazy::Data_List_Lazy_take_tco(n))
    }
    pub fn Data_List_Lazy_take_004038_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_take_004038_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_take_004038_002d1.get_or_init(||
                                                         Lazy(Data_List_Lazy_take_004038.clone()))
    }
    pub fn Data_List_Lazy_take_tco(n: &dyn Any) -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_take_004038_002d1 =
                                Data_List_Lazy_take_004038_002d1.clone();
                            move |v|
                                &Func1::new({
                                                let Data_List_Lazy_take_004038_002d1
                                                    =
                                                    Data_List_Lazy_take_004038_002d1.clone();
                                                let v = v.clone();
                                                move |v1|
                                                    {
                                                        let matchValue =
                                                            Sharpurs_Prelude::unbox(&&v);
                                                        let matchValue_1:
                                                                LrcPtr<Data_List_Lazy_Types_Step> =
                                                            Sharpurs_Prelude::unbox(v1);
                                                        match matchValue_1.as_ref()
                                                            {
                                                            Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                         matchValue_1_1_1)
                                                            =>
                                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_take_004038_002d1.Value,
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                                                                             &&&1_i32)),
                                                                                                                                                                       &&matchValue_1_1_1))),
                                                            _ =>
                                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                                        }
                                                    }
                                            })
                        });
        let matchValue_3 =
            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                            &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                         n),
                                                                      &&&0_i32));
        match &Sharpurs_Prelude::_007cLitBool_007c__007c(true, &matchValue_3)
            {
            0_i32 =>
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                             &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()),
            _ =>
            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                      &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                         &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                                                         n))),
                                                                                &&&PureScript_Data_List_Lazy::Data_List_Lazy_unwrap())),
        }
    }
    pub fn Data_List_Lazy_take() -> &dyn Any {
        static Data_List_Lazy_take: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_take.get_or_init(||
                                            Data_List_Lazy_take_004038_002d1.Value)
    }
    pub fn Data_List_Lazy_tail() -> &dyn Any {
        static Data_List_Lazy_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_tail.get_or_init(||
                                            &Func1::new(move |xs|
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                   &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                &&&Func1::new(move
                                                                                                                                                  |v|
                                                                                                                                                  find(string("tail"),
                                                                                                                                                       Sharpurs_Prelude::unbox(v)))),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                                                                xs))))
    }
    pub fn Data_List_Lazy_stripPrefix() -> &dyn Any {
        static Data_List_Lazy_stripPrefix: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_stripPrefix.get_or_init(||
                                                   &Func1::new(move |dictEq|
                                                                   &Func1::new({
                                                                                   let dictEq
                                                                                       =
                                                                                       dictEq.clone();
                                                                                   move
                                                                                       |v|
                                                                                       &Func1::new({
                                                                                                       let v
                                                                                                           =
                                                                                                           v.clone();
                                                                                                       move
                                                                                                           |s|
                                                                                                           {
                                                                                                               let matchValue =
                                                                                                                   Sharpurs_Prelude::unbox(&&v);
                                                                                                               let matchValue_1 =
                                                                                                                   Sharpurs_Prelude::unbox(s);
                                                                                                               let go =
                                                                                                                   &Func1::new(move
                                                                                                                                   |prefix|
                                                                                                                                   &Func1::new({
                                                                                                                                                   let prefix
                                                                                                                                                       =
                                                                                                                                                       prefix.clone();
                                                                                                                                                   move
                                                                                                                                                       |input|
                                                                                                                                                       {
                                                                                                                                                           let matchValue_3:
                                                                                                                                                                   LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                          &&&prefix));
                                                                                                                                                           match matchValue_3.as_ref()
                                                                                                                                                               {
                                                                                                                                                               Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_3_1_0,
                                                                                                                                                                                                                            matchValue_3_1_1)
                                                                                                                                                               =>
                                                                                                                                                               {
                                                                                                                                                                   let matchValue_4:
                                                                                                                                                                           LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                                  input));
                                                                                                                                                                   if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_4_1_0,
                                                                                                                                                                                                                                       matchValue_4_1_1)
                                                                                                                                                                          =
                                                                                                                                                                          matchValue_4.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                       if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                          &&&dictEq),
                                                                                                                                                                                                                                                                       &&matchValue_3_1_0),
                                                                                                                                                                                                                                    &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                           Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                       }))
                                                                                                                                                                          {
                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                               &&&Func1::new(move
                                                                                                                                                                                                                                                                 |usd__arg1_1|
                                                                                                                                                                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1_1.clone())))),
                                                                                                                                                                                                            &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Loopusd_Ctor(&add(string("a"),
                                                                                                                                                                                                                                                                                                   &matchValue_3_1_1,
                                                                                                                                                                                                                                                                                                   add(string("b"),
                                                                                                                                                                                                                                                                                                       &&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                                                                         },
                                                                                                                                                                                                                                                                                                       empty::<string,
                                                                                                                                                                                                                                                                                                               &dyn Any>())))))
                                                                                                                                                                       } else {
                                                                                                                                                                           &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                                                                       }
                                                                                                                                                                   } else {
                                                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                                                                   }
                                                                                                                                                               }
                                                                                                                                                               _
                                                                                                                                                               =>
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                     |usd__arg1|
                                                                                                                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                &&&LrcPtr::new(Control_Monad_Rec_Class_Step::Control_Monad_Rec_Class_Doneusd_Ctor(input.clone()))),
                                                                                                                                                           }
                                                                                                                                                       }
                                                                                                                                               }));
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_tailRecM2(),
                                                                                                                                                                                                                                                         &&&PureScript_Control_Monad_Rec_Class::Control_Monad_Rec_Class_monadRecMaybe()),
                                                                                                                                                                                                                      &&&go),
                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                &&&matchValue_1)
                                                                                                           }
                                                                                                   })
                                                                               })))
    }
    pub fn Data_List_Lazy_span_004046() -> &dyn Any {
        &Func1::new(move |p|
                        Func1::new({
                                       let p = p.clone();
                                       move |xs|
                                           PureScript_Data_List_Lazy::Data_List_Lazy_span_tco(&p,
                                                                                              xs)
                                   }))
    }
    pub fn Data_List_Lazy_span_004046_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_span_004046_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_span_004046_002d1.get_or_init(||
                                                         Lazy(Data_List_Lazy_span_004046.clone()))
    }
    pub fn Data_List_Lazy_span_tco(p: &dyn Any, xs: &dyn Any) -> &dyn Any {
        let matchValue: LrcPtr<Data_Maybe_Maybe> =
            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                       xs));
        if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0) =
               matchValue.as_ref() {
            let activePatternResult =
                Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref() {
                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   });
            let activePatternResult_1 =
                Sharpurs_Prelude::_007cHasProp_007c__007c(string("head"),
                                                          &activePatternResult);
            if activePatternResult_1.is_some() {
                let activePatternResult_2 =
                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("tail"),
                                                              &activePatternResult);
                if activePatternResult_2.is_some() {
                    if {
                           let xs_prime =
                               getValue(activePatternResult_2.clone());
                           Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(p,
                                                                                     &&&getValue(activePatternResult_1.clone())))
                       } {
                        let x_1 = getValue(activePatternResult_1);
                        let xs_prime_1 = getValue(activePatternResult_2);
                        let matchValue_1 =
                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_span_004046_002d1.Value,
                                                                                                                          p),
                                                                                       &&&xs_prime_1));
                        {
                            let activePatternResult_3 =
                                Sharpurs_Prelude::_007cHasProp_007c__007c(string("init"),
                                                                          &matchValue_1);
                            if activePatternResult_3.is_some() {
                                let activePatternResult_4 =
                                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("rest"),
                                                                              &matchValue_1);
                                if activePatternResult_4.is_some() {
                                    let ys = getValue(activePatternResult_3);
                                    let zs = getValue(activePatternResult_4);
                                    &add(string("init"),
                                         &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                              &&&x_1),
                                                                           &&&ys),
                                         add(string("rest"), &&zs,
                                             empty::<string, &dyn Any>()))
                                } else {
                                    panic!("{}",
                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 45_i32,
                                  Data2: 458_i32,}).get_Message(),)
                                }
                            } else {
                                panic!("{}",
                                       LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 45_i32,
                                  Data2: 458_i32,}).get_Message(),)
                            }
                        }
                    } else {
                        &add(string("init"),
                             &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                             add(string("rest"), xs.clone(),
                                 empty::<string, &dyn Any>()))
                    }
                } else {
                    &add(string("init"),
                         &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                         add(string("rest"), xs.clone(),
                             empty::<string, &dyn Any>()))
                }
            } else {
                &add(string("init"),
                     &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                     add(string("rest"), xs.clone(),
                         empty::<string, &dyn Any>()))
            }
        } else {
            &add(string("init"),
                 &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                 add(string("rest"), xs.clone(), empty::<string, &dyn Any>()))
        }
    }
    pub fn Data_List_Lazy_span() -> &dyn Any {
        static Data_List_Lazy_span: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_span.get_or_init(||
                                            Data_List_Lazy_span_004046_002d1.Value)
    }
    pub fn Data_List_Lazy_snoc() -> &dyn Any {
        static Data_List_Lazy_snoc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_snoc.get_or_init(||
                                            &Func1::new(move |xs|
                                                            &Func1::new({
                                                                            let xs
                                                                                =
                                                                                xs.clone();
                                                                            move
                                                                                |x|
                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                          &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons()),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                          x),
                                                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())),
                                                                                                                 &&&xs)
                                                                        })))
    }
    pub fn Data_List_Lazy_singleton() -> &dyn Any {
        static Data_List_Lazy_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_singleton.get_or_init(||
                                                 &Func1::new(move |a|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                     a),
                                                                                                  &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())))
    }
    pub fn Data_List_Lazy_showPattern() -> &dyn Any {
        static Data_List_Lazy_showPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_showPattern.get_or_init(||
                                                   &Func1::new(move |dictShow|
                                                                   {
                                                                       let showList =
                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_showList(),
                                                                                                            dictShow);
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                        &&&add(string("show"),
                                                                                                               &&Func1::new({
                                                                                                                                let showList
                                                                                                                                    =
                                                                                                                                    showList.clone();
                                                                                                                                move
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
                                                                                                                                                                                                                                                                                                                     &&&showList),
                                                                                                                                                                                                                                                                                  &&&s)),
                                                                                                                                                                                                            &&&string(")")))
                                                                                                                                    }
                                                                                                                            }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))
                                                                   }))
    }
    pub fn Data_List_Lazy_scanlLazy_004056() -> &dyn Any {
        &Func1::new(move |f|
                        Func1::new({
                                       let f = f.clone();
                                       move |acc|
                                           Func1::new({
                                                          let acc =
                                                              acc.clone();
                                                          move |xs|
                                                              PureScript_Data_List_Lazy::Data_List_Lazy_scanlLazy_tco(&f,
                                                                                                                      &acc,
                                                                                                                      xs)
                                                      })
                                   }))
    }
    pub fn Data_List_Lazy_scanlLazy_004056_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_scanlLazy_004056_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_scanlLazy_004056_002d1.get_or_init(||
                                                              Lazy(Data_List_Lazy_scanlLazy_004056.clone()))
    }
    pub fn Data_List_Lazy_scanlLazy_tco(f: &dyn Any, acc: &dyn Any,
                                        xs: &dyn Any) -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_scanlLazy_004056_002d1 =
                                Data_List_Lazy_scanlLazy_004056_002d1.clone();
                            let acc = acc.clone();
                            let f = f.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    match matchValue.as_ref() {
                                        Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                     matchValue_1_1)
                                        => {
                                            let acc_prime =
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                    &&&acc),
                                                                                 &&matchValue_1_0);
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                     |usd__arg1|
                                                                                                                                                                     Func1::new({
                                                                                                                                                                                    let usd__arg1
                                                                                                                                                                                        =
                                                                                                                                                                                        usd__arg1.clone();
                                                                                                                                                                                    move
                                                                                                                                                                                        |usd__arg2|
                                                                                                                                                                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                  usd__arg2.clone()))
                                                                                                                                                                                })),
                                                                                                                                                   &&&acc_prime)),
                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_scanlLazy_004056_002d1.Value,
                                                                                                                                                                                      &&&f),
                                                                                                                                                   &&&acc_prime),
                                                                                                                &&matchValue_1_1))
                                        }
                                        _ =>
                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                               &&&go),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                               xs)))
    }
    pub fn Data_List_Lazy_scanlLazy() -> &dyn Any {
        static Data_List_Lazy_scanlLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_scanlLazy.get_or_init(||
                                                 Data_List_Lazy_scanlLazy_004056_002d1.Value)
    }
    pub fn Data_List_Lazy_reverse() -> &dyn Any {
        static Data_List_Lazy_reverse: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_reverse.get_or_init(||
                                               &Func1::new(move |xs|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                   &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_lazyList()),
                                                                                                &&&Func1::new({
                                                                                                                  let xs
                                                                                                                      =
                                                                                                                      xs.clone();
                                                                                                                  move
                                                                                                                      |v|
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons())),
                                                                                                                                                                                          &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()),
                                                                                                                                                       &&&xs)
                                                                                                              }))))
    }
    pub fn Data_List_Lazy_replicateM_004062() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Data_List_Lazy::Data_List_Lazy_replicateM_tco(dictMonad))
    }
    pub fn Data_List_Lazy_replicateM_004062_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_replicateM_004062_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_replicateM_004062_002d1.get_or_init(||
                                                               Lazy(Data_List_Lazy_replicateM_004062.clone()))
    }
    pub fn Data_List_Lazy_replicateM_tco(dictMonad: &dyn Any) -> &dyn Any {
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Data_List_Lazy_replicateM_004062_002d1 =
                            Data_List_Lazy_replicateM_004062_002d1.clone();
                        let dictMonad = dictMonad.clone();
                        move |n|
                            &Func1::new({
                                            let Data_List_Lazy_replicateM_004062_002d1
                                                =
                                                Data_List_Lazy_replicateM_004062_002d1.clone();
                                            let n = n.clone();
                                            move |m|
                                                {
                                                    let matchValue =
                                                        Sharpurs_Prelude::unbox(&&n);
                                                    let matchValue_1 =
                                                        Sharpurs_Prelude::unbox(m);
                                                    if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                       &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                    &&&matchValue),
                                                                                                                 &&&PureScript_Data_List_Lazy::Data_List_Lazy_one()))
                                                       {
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                            &&&Applicative0),
                                                                                         &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())
                                                    } else {
                                                        if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                           {
                                                            let m1_3 =
                                                                matchValue_1;
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                   &&&Bind1),
                                                                                                                                &&&m1_3),
                                                                                             &&&Func1::new({
                                                                                                               let Data_List_Lazy_replicateM_004062_002d1
                                                                                                                   =
                                                                                                                   Data_List_Lazy_replicateM_004062_002d1.clone();
                                                                                                               let m1_3
                                                                                                                   =
                                                                                                                   m1_3.clone();
                                                                                                               move
                                                                                                                   |a|
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                          &&&Bind1),
                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_replicateM_004062_002d1.Value,
                                                                                                                                                                                                                                                                                                &&&dictMonad),
                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                                                                                                                                &&&PureScript_Data_List_Lazy::Data_List_Lazy_one())),
                                                                                                                                                                                                                          &&&m1_3)),
                                                                                                                                                    &&&Func1::new({
                                                                                                                                                                      let a
                                                                                                                                                                          =
                                                                                                                                                                          a.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |as_var|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                              &&&Applicative0),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                 &&&a),
                                                                                                                                                                                                                                              as_var))
                                                                                                                                                                  }))
                                                                                                           }))
                                                        } else {
                                                            panic!("{}",
                                                                   LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 61_i32,
                                  Data2: 394_i32,}).get_Message(),)
                                                        }
                                                    }
                                                }
                                        })
                    })
    }
    pub fn Data_List_Lazy_replicateM() -> &dyn Any {
        static Data_List_Lazy_replicateM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_replicateM.get_or_init(||
                                                  Data_List_Lazy_replicateM_004062_002d1.Value)
    }
    pub fn Data_List_Lazy_repeat() -> &dyn Any {
        static Data_List_Lazy_repeat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_repeat.get_or_init(||
                                              &Func1::new(move |x|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_fix(),
                                                                                                                                  &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_lazyList()),
                                                                                               &&&Func1::new({
                                                                                                                 let x
                                                                                                                     =
                                                                                                                     x.clone();
                                                                                                                 move
                                                                                                                     |xs|
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                         &&&x),
                                                                                                                                                      xs)
                                                                                                             }))))
    }
    pub fn Data_List_Lazy_replicate() -> &dyn Any {
        static Data_List_Lazy_replicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_replicate.get_or_init(||
                                                 &Func1::new(move |i|
                                                                 &Func1::new({
                                                                                 let i
                                                                                     =
                                                                                     i.clone();
                                                                                 move
                                                                                     |xs|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_take(),
                                                                                                                                                         &&&i),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_repeat(),
                                                                                                                                                         xs))
                                                                             })))
    }
    pub fn Data_List_Lazy_range() -> &dyn Any {
        static Data_List_Lazy_range: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_range.get_or_init(||
                                             &Func1::new(move |start|
                                                             &Func1::new({
                                                                             let start
                                                                                 =
                                                                                 start.clone();
                                                                             move
                                                                                 |end_var|
                                                                                 {
                                                                                     let matchValue =
                                                                                         Sharpurs_Prelude::unbox(&&start);
                                                                                     let matchValue_1 =
                                                                                         Sharpurs_Prelude::unbox(end_var);
                                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                        &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                     &&&matchValue),
                                                                                                                                                  &&&matchValue_1))
                                                                                        {
                                                                                         let g =
                                                                                             &Func1::new({
                                                                                                             let matchValue_1
                                                                                                                 =
                                                                                                                 matchValue_1.clone();
                                                                                                             move
                                                                                                                 |x|
                                                                                                                 {
                                                                                                                     let matchValue_3 =
                                                                                                                         Sharpurs_Prelude::unbox(x);
                                                                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                        &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                     &&&matchValue_3),
                                                                                                                                                                                  &&&matchValue_1))
                                                                                                                        {
                                                                                                                         let x1_2 =
                                                                                                                             matchValue_3;
                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&x1_2,
                                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                            &&&x1_2),
                                                                                                                                                                                                                                                                         &&&1_i32)))))
                                                                                                                     } else {
                                                                                                                         if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                            {
                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                         } else {
                                                                                                                             panic!("{}",
                                                                                                                                    LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 69_i32,
                                  Data2: 389_i32,}).get_Message(),)
                                                                                                                         }
                                                                                                                     }
                                                                                                                 }
                                                                                                         });
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_unfoldableList()),
                                                                                                                                                             &&&g),
                                                                                                                          &&&matchValue)
                                                                                     } else {
                                                                                         if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                            {
                                                                                             let f =
                                                                                                 &Func1::new({
                                                                                                                 let matchValue_1
                                                                                                                     =
                                                                                                                     matchValue_1.clone();
                                                                                                                 move
                                                                                                                     |x_1|
                                                                                                                     {
                                                                                                                         let matchValue_4 =
                                                                                                                             Sharpurs_Prelude::unbox(x_1);
                                                                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                            &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                         &&&matchValue_4),
                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                            {
                                                                                                                             let x1_6 =
                                                                                                                                 matchValue_4;
                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&x1_6,
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                                                                &&&x1_6),
                                                                                                                                                                                                                                                                             &&&1_i32)))))
                                                                                                                         } else {
                                                                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                {
                                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                                                             } else {
                                                                                                                                 panic!("{}",
                                                                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 69_i32,
                                  Data2: 1232_i32,}).get_Message(),)
                                                                                                                             }
                                                                                                                         }
                                                                                                                     }
                                                                                                             });
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Unfoldable::Data_Unfoldable_unfoldr(),
                                                                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_unfoldableList()),
                                                                                                                                                                 &&&f),
                                                                                                                              &&&matchValue)
                                                                                         } else {
                                                                                             panic!("{}",
                                                                                                    LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 69_i32,
                                  Data2: 89_i32,}).get_Message(),)
                                                                                         }
                                                                                     }
                                                                                 }
                                                                         })))
    }
    pub fn Data_List_Lazy_partition() -> &dyn Any {
        static Data_List_Lazy_partition: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_partition.get_or_init(||
                                                 &Func1::new(move |f|
                                                                 {
                                                                     let go =
                                                                         &Func1::new({
                                                                                         let f
                                                                                             =
                                                                                             f.clone();
                                                                                         move
                                                                                             |x|
                                                                                             &Func1::new({
                                                                                                             let x
                                                                                                                 =
                                                                                                                 x.clone();
                                                                                                             move
                                                                                                                 |v|
                                                                                                                 {
                                                                                                                     let matchValue =
                                                                                                                         Sharpurs_Prelude::unbox(&&x);
                                                                                                                     let matchValue_1 =
                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                     {
                                                                                                                         let activePatternResult =
                                                                                                                             Sharpurs_Prelude::_007cHasProp_007c__007c(string("yes"),
                                                                                                                                                                       &matchValue_1);
                                                                                                                         if activePatternResult.is_some()
                                                                                                                            {
                                                                                                                             let activePatternResult_1 =
                                                                                                                                 Sharpurs_Prelude::_007cHasProp_007c__007c(string("no"),
                                                                                                                                                                           &matchValue_1);
                                                                                                                             if activePatternResult_1.is_some()
                                                                                                                                {
                                                                                                                                 let ns =
                                                                                                                                     getValue(activePatternResult_1);
                                                                                                                                 let ys =
                                                                                                                                     getValue(activePatternResult);
                                                                                                                                 let matchValue_3 =
                                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                               &&&matchValue));
                                                                                                                                 match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                  &matchValue_3)
                                                                                                                                     {
                                                                                                                                     0_i32
                                                                                                                                     =>
                                                                                                                                     &add(string("yes"),
                                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                            &&&ys),
                                                                                                                                          add(string("no"),
                                                                                                                                              &&ns,
                                                                                                                                              empty::<string,
                                                                                                                                                      &dyn Any>())),
                                                                                                                                     _
                                                                                                                                     =>
                                                                                                                                     &add(string("yes"),
                                                                                                                                          &&ys,
                                                                                                                                          add(string("no"),
                                                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                &&&ns),
                                                                                                                                              empty::<string,
                                                                                                                                                      &dyn Any>())),
                                                                                                                                 }
                                                                                                                             } else {
                                                                                                                                 panic!("{}",
                                                                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 71_i32,
                                  Data2: 115_i32,}).get_Message(),)
                                                                                                                             }
                                                                                                                         } else {
                                                                                                                             panic!("{}",
                                                                                                                                    LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 71_i32,
                                  Data2: 115_i32,}).get_Message(),)
                                                                                                                         }
                                                                                                                     }
                                                                                                                 }
                                                                                                         })
                                                                                     });
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                         &&&go),
                                                                                                      &&&add(string("yes"),
                                                                                                             &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                             add(string("no"),
                                                                                                                 &&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil(),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>())))
                                                                 }))
    }
    pub fn Data_List_Lazy_null() -> &dyn Any {
        static Data_List_Lazy_null: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_null.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                   &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                &&&PureScript_Data_Maybe::Data_Maybe_isNothing()),
                                                                             &&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons()))
    }
    pub fn Data_List_Lazy_nubBy() -> &dyn Any {
        static Data_List_Lazy_nubBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_nubBy.get_or_init(||
                                             &Func1::new(move |p|
                                                             {
                                                                 let goStep_2 =
                                                                     Func0::new({
                                                                                    let goStep_tco
                                                                                        =
                                                                                        goStep_tco.clone();
                                                                                    move
                                                                                        ||
                                                                                        &Func1::new({
                                                                                                        let goStep_tco
                                                                                                            =
                                                                                                            goStep_tco.clone();
                                                                                                        move
                                                                                                            |v|
                                                                                                            Func1::new({
                                                                                                                           let goStep_tco
                                                                                                                               =
                                                                                                                               goStep_tco.clone();
                                                                                                                           let v
                                                                                                                               =
                                                                                                                               v.clone();
                                                                                                                           move
                                                                                                                               |v1|
                                                                                                                               goStep_tco(v)(v1.clone())
                                                                                                                       })
                                                                                                    })
                                                                                });
                                                                 let goStep_1 =
                                                                     Lazy(goStep_2);
                                                                 let go_2 =
                                                                     Func0::new({
                                                                                    let go_tco
                                                                                        =
                                                                                        go_tco.clone();
                                                                                    move
                                                                                        ||
                                                                                        &Func1::new({
                                                                                                        let go_tco
                                                                                                            =
                                                                                                            go_tco.clone();
                                                                                                        move
                                                                                                            |s|
                                                                                                            Func1::new({
                                                                                                                           let go_tco
                                                                                                                               =
                                                                                                                               go_tco.clone();
                                                                                                                           let s
                                                                                                                               =
                                                                                                                               s.clone();
                                                                                                                           move
                                                                                                                               |v_1|
                                                                                                                               go_tco(s)(v_1.clone())
                                                                                                                       })
                                                                                                    })
                                                                                });
                                                                 let go_1 =
                                                                     Lazy(go_2);
                                                                 let goStep_tco =
                                                                     Func1::new({
                                                                                    let go_tco
                                                                                        =
                                                                                        go_tco.clone();
                                                                                    let p
                                                                                        =
                                                                                        p.clone();
                                                                                    move
                                                                                        |v_2|
                                                                                        Func1::new({
                                                                                                       let go_tco
                                                                                                           =
                                                                                                           go_tco.clone();
                                                                                                       let v_2
                                                                                                           =
                                                                                                           v_2.clone();
                                                                                                       move
                                                                                                           |v1_1|
                                                                                                           {
                                                                                                               let matchValue =
                                                                                                                   Sharpurs_Prelude::unbox(&&v_2);
                                                                                                               let matchValue_1:
                                                                                                                       LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                                                   Sharpurs_Prelude::unbox(v1_1);
                                                                                                               match matchValue_1.as_ref()
                                                                                                                   {
                                                                                                                   Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                matchValue_1_1_1)
                                                                                                                   =>
                                                                                                                   {
                                                                                                                       let as_var =
                                                                                                                           matchValue_1_1_1.clone();
                                                                                                                       let a =
                                                                                                                           matchValue_1_1_0.clone();
                                                                                                                       let matchValue_3 =
                                                                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Internal::Data_List_Internal_insertAndLookupBy(),
                                                                                                                                                                                                                                                            &&&p),
                                                                                                                                                                                                                         &&&a),
                                                                                                                                                                                      &&&matchValue));
                                                                                                                       {
                                                                                                                           let activePatternResult =
                                                                                                                               Sharpurs_Prelude::_007cHasProp_007c__007c(string("found"),
                                                                                                                                                                         &matchValue_3);
                                                                                                                           if activePatternResult.is_some()
                                                                                                                              {
                                                                                                                               let activePatternResult_1 =
                                                                                                                                   Sharpurs_Prelude::_007cHasProp_007c__007c(string("result"),
                                                                                                                                                                             &matchValue_3);
                                                                                                                               if activePatternResult_1.is_some()
                                                                                                                                  {
                                                                                                                                   let found =
                                                                                                                                       getValue(activePatternResult);
                                                                                                                                   let s_prime =
                                                                                                                                       getValue(activePatternResult_1);
                                                                                                                                   let matchValue_4 =
                                                                                                                                       Sharpurs_Prelude::unbox(&&found);
                                                                                                                                   match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                    &matchValue_4)
                                                                                                                                       {
                                                                                                                                       0_i32
                                                                                                                                       =>
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                        &&go_tco(&s_prime)(&as_var)),
                                                                                                                                       _
                                                                                                                                       =>
                                                                                                                                       &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&a,
                                                                                                                                                                                                                 go_tco(&s_prime)(&as_var))),
                                                                                                                                   }
                                                                                                                               } else {
                                                                                                                                   panic!("{}",
                                                                                                                                          LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 77_i32,
                                  Data2: 627_i32,}).get_Message(),)
                                                                                                                               }
                                                                                                                           } else {
                                                                                                                               panic!("{}",
                                                                                                                                      LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 77_i32,
                                  Data2: 627_i32,}).get_Message(),)
                                                                                                                           }
                                                                                                                       }
                                                                                                                   }
                                                                                                                   _
                                                                                                                   =>
                                                                                                                   &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                                                                                               }
                                                                                                           }
                                                                                                   })
                                                                                });
                                                                 let goStep =
                                                                     goStep_1.Value;
                                                                 let go_tco =
                                                                     Func1::new({
                                                                                    let goStep_1
                                                                                        =
                                                                                        goStep_1.clone();
                                                                                    move
                                                                                        |s_2|
                                                                                        Func1::new({
                                                                                                       let goStep_1
                                                                                                           =
                                                                                                           goStep_1.clone();
                                                                                                       let s_2
                                                                                                           =
                                                                                                           s_2.clone();
                                                                                                       move
                                                                                                           |v_3|
                                                                                                           {
                                                                                                               let matchValue_5 =
                                                                                                                   Sharpurs_Prelude::unbox(&&s_2);
                                                                                                               let matchValue_6 =
                                                                                                                   Sharpurs_Prelude::unbox(v_3);
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&goStep_1.Value,
                                                                                                                                                                                                                                                         &&&matchValue_5)),
                                                                                                                                                                                   &&&matchValue_6))
                                                                                                           }
                                                                                                   })
                                                                                });
                                                                 let go =
                                                                     go_1.Value;
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                  &&&PureScript_Data_List_Internal::Data_List_Internal_emptySet())
                                                             }))
    }
    pub fn Data_List_Lazy_nub() -> &dyn Any {
        static Data_List_Lazy_nub: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_nub.get_or_init(||
                                           &Func1::new(move |dictOrd|
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_nubBy(),
                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                               dictOrd))))
    }
    pub fn Data_List_Lazy_newtypePattern() -> &dyn Any {
        static Data_List_Lazy_newtypePattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_newtypePattern.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                       &&&add(string("Coercible0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &Sharpurs_Prelude::Prim_undefined()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_List_Lazy_mapMaybe_004096() -> &dyn Any {
        &Func1::new(move |f|
                        PureScript_Data_List_Lazy::Data_List_Lazy_mapMaybe_tco(f))
    }
    pub fn Data_List_Lazy_mapMaybe_004096_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_mapMaybe_004096_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_mapMaybe_004096_002d1.get_or_init(||
                                                             Lazy(Data_List_Lazy_mapMaybe_004096.clone()))
    }
    pub fn Data_List_Lazy_mapMaybe_tco(f: &dyn Any) -> &dyn Any {
        let go_2 =
            Func0::new({
                           let go_tco = go_tco.clone();
                           move ||
                               &Func1::new({
                                               let go_tco = go_tco.clone();
                                               move |v| go_tco(v.clone())
                                           })
                       });
        let go_1 = Lazy(go_2);
        let go_tco =
            Func1::new({
                           let Data_List_Lazy_mapMaybe_004096_002d1 =
                               Data_List_Lazy_mapMaybe_004096_002d1.clone();
                           let f = f.clone();
                           move |v_1|
                               {
                                   let v_1 = v_1.clone();
                                   '_go_tco:
                                       loop  {
                                           break '_go_tco
                                               ({
                                                    let matchValue:
                                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                                        Sharpurs_Prelude::unbox(&&v_1);
                                                    match matchValue.as_ref()
                                                        {
                                                        Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                     matchValue_1_1)
                                                        => {
                                                            let xs =
                                                                matchValue_1_1.clone();
                                                            let matchValue_1:
                                                                    LrcPtr<Data_Maybe_Maybe> =
                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                           &&matchValue_1_0));
                                                            match matchValue_1.as_ref()
                                                                {
                                                                Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                                                                =>
                                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_mapMaybe_004096_002d1.Value,
                                                                                                                                                                                                              &&&f),
                                                                                                                                                                           &&&xs))),
                                                                _ => {
                                                                    let v_1_temp =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                         &&&xs);
                                                                    v_1.set(v_1_temp);
                                                                    continue
                                                                        '_go_tco

                                                                }
                                                            }
                                                        }
                                                        _ =>
                                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                                    }
                                                }) ;
                                       }
                               }
                       });
        let go = go_1.Value;
        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                     &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                  &&&go)),
                                                                            &&&PureScript_Data_List_Lazy::Data_List_Lazy_unwrap()))
    }
    pub fn Data_List_Lazy_mapMaybe() -> &dyn Any {
        static Data_List_Lazy_mapMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_mapMaybe.get_or_init(||
                                                Data_List_Lazy_mapMaybe_004096_002d1.Value)
    }
    pub fn Data_List_Lazy_some_0040100() -> &dyn Any {
        &Func1::new(move |dictAlternative|
                        PureScript_Data_List_Lazy::Data_List_Lazy_some_tco(dictAlternative))
    }
    pub fn Data_List_Lazy_some_0040100_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_some_0040100_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_some_0040100_002d1.get_or_init(||
                                                          Lazy(Data_List_Lazy_some_0040100.clone()))
    }
    pub fn Data_List_Lazy_many_0040102() -> &dyn Any {
        &Func1::new(move |dictAlternative|
                        PureScript_Data_List_Lazy::Data_List_Lazy_many_tco(dictAlternative))
    }
    pub fn Data_List_Lazy_many_0040102_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_many_0040102_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_many_0040102_002d1.get_or_init(||
                                                          Lazy(Data_List_Lazy_many_0040102.clone()))
    }
    pub fn Data_List_Lazy_some_tco(dictAlternative: &dyn Any) -> &dyn Any {
        let Apply0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Functor0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Apply0 = Apply0.clone();
                        let Data_List_Lazy_many_0040102_002d1 =
                            Data_List_Lazy_many_0040102_002d1.clone();
                        let Functor0 = Functor0.clone();
                        let dictAlternative = dictAlternative.clone();
                        move |dictLazy|
                            &Func1::new({
                                            let Data_List_Lazy_many_0040102_002d1
                                                =
                                                Data_List_Lazy_many_0040102_002d1.clone();
                                            let dictLazy = dictLazy.clone();
                                            move |v|
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                       &&&Apply0),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                             &&&Functor0),
                                                                                                                                                                                          &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons()),
                                                                                                                                                       v)),
                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                                       &&&dictLazy),
                                                                                                                    &&&Func1::new({
                                                                                                                                      let Data_List_Lazy_many_0040102_002d1
                                                                                                                                          =
                                                                                                                                          Data_List_Lazy_many_0040102_002d1.clone();
                                                                                                                                      let v
                                                                                                                                          =
                                                                                                                                          v.clone();
                                                                                                                                      move
                                                                                                                                          |v1|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_many_0040102_002d1.Value,
                                                                                                                                                                                                                                                 &&&dictAlternative),
                                                                                                                                                                                                              &&&dictLazy),
                                                                                                                                                                           &&&v)
                                                                                                                                  })))
                                        })
                    })
    }
    pub fn Data_List_Lazy_some() -> &dyn Any {
        static Data_List_Lazy_some: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_some.get_or_init(||
                                            Data_List_Lazy_some_0040100_002d1.Value)
    }
    pub fn Data_List_Lazy_many_tco(dictAlternative: &dyn Any) -> &dyn Any {
        let Alt0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Plus1"),
                                                                                                                     Sharpurs_Prelude::unbox(dictAlternative)),
                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()))),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictAlternative)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Alt0 = Alt0.clone();
                        let Applicative0 = Applicative0.clone();
                        let Data_List_Lazy_some_0040100_002d1 =
                            Data_List_Lazy_some_0040100_002d1.clone();
                        let dictAlternative = dictAlternative.clone();
                        move |dictLazy|
                            &Func1::new({
                                            let Data_List_Lazy_some_0040100_002d1
                                                =
                                                Data_List_Lazy_some_0040100_002d1.clone();
                                            let dictLazy = dictLazy.clone();
                                            move |v|
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                       &&&Alt0),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_some_0040100_002d1.Value,
                                                                                                                                                                                                                             &&&dictAlternative),
                                                                                                                                                                                          &&&dictLazy),
                                                                                                                                                       v)),
                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                       &&&Applicative0),
                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()))
                                        })
                    })
    }
    pub fn Data_List_Lazy_many() -> &dyn Any {
        static Data_List_Lazy_many: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_many.get_or_init(||
                                            Data_List_Lazy_many_0040102_002d1.Value)
    }
    pub fn Data_List_Lazy_length() -> &dyn Any {
        static Data_List_Lazy_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_length.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                  &&&Func1::new(move
                                                                                                                                    |l|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let l
                                                                                                                                                        =
                                                                                                                                                        l.clone();
                                                                                                                                                    move
                                                                                                                                                        |v|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                            &&&l),
                                                                                                                                                                                         &&&1_i32)
                                                                                                                                                }))),
                                                                               &&&0_i32))
    }
    pub fn Data_List_Lazy_last() -> &dyn Any {
        static Data_List_Lazy_last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_last.get_or_init(||
                                            {
                                                let go_2 =
                                                    Func0::new({
                                                                   let go_tco
                                                                       =
                                                                       go_tco.clone();
                                                                   move ||
                                                                       &Func1::new({
                                                                                       let go_tco
                                                                                           =
                                                                                           go_tco.clone();
                                                                                       move
                                                                                           |v|
                                                                                           go_tco(v.clone())
                                                                                   })
                                                               });
                                                let go_1 = Lazy(go_2);
                                                fn go_tco(v_1: &dyn Any)
                                                 -> &dyn Any {
                                                    let v_1 = v_1.clone();
                                                    '_go_tco:
                                                        loop  {
                                                            break '_go_tco
                                                                ({
                                                                     let matchValue:
                                                                             LrcPtr<Data_List_Lazy_Types_Step> =
                                                                         Sharpurs_Prelude::unbox(&&v_1);
                                                                     if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                         matchValue_1_1)
                                                                            =
                                                                            matchValue.as_ref()
                                                                        {
                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_null(),
                                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                                             {
                                                                                                                                             Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                          x)
                                                                                                                                             =>
                                                                                                                                             x.clone(),
                                                                                                                                             _
                                                                                                                                             =>
                                                                                                                                             unreachable!(),
                                                                                                                                         }))
                                                                            {
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue.as_ref()
                                                                                                                                         {
                                                                                                                                         Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                      _)
                                                                                                                                         =>
                                                                                                                                         x.clone(),
                                                                                                                                         _
                                                                                                                                         =>
                                                                                                                                         unreachable!(),
                                                                                                                                     }))
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 let v_1_temp =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                             {
                                                                                                                             Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                          x)
                                                                                                                             =>
                                                                                                                             x.clone(),
                                                                                                                             _
                                                                                                                             =>
                                                                                                                             unreachable!(),
                                                                                                                         });
                                                                                 v_1.set(v_1_temp);
                                                                                 continue
                                                                                     '_go_tco

                                                                             } else {
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                             }
                                                                         }
                                                                     } else {
                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                     }
                                                                 }) ;
                                                        }
                                                }
                                                let go = go_1.Value;
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                    &&&go),
                                                                                 &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step())
                                            })
    }
    pub fn Data_List_Lazy_iterate() -> &dyn Any {
        static Data_List_Lazy_iterate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_iterate.get_or_init(||
                                               &Func1::new(move |f|
                                                               &Func1::new({
                                                                               let f
                                                                                   =
                                                                                   f.clone();
                                                                               move
                                                                                   |x|
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_fix(),
                                                                                                                                                       &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_lazyList()),
                                                                                                                    &&&Func1::new({
                                                                                                                                      let x
                                                                                                                                          =
                                                                                                                                          x.clone();
                                                                                                                                      move
                                                                                                                                          |xs|
                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                              &&&x),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_functorList()),
                                                                                                                                                                                                                                                 &&&f),
                                                                                                                                                                                                              xs))
                                                                                                                                  }))
                                                                           })))
    }
    pub fn Data_List_Lazy_insertAt_0040118() -> &dyn Any {
        &Func1::new(move |v|
                        Func1::new({
                                       let v = v.clone();
                                       move |v1|
                                           Func1::new({
                                                          let v1 = v1.clone();
                                                          move |v2|
                                                              PureScript_Data_List_Lazy::Data_List_Lazy_insertAt_tco(&v,
                                                                                                                     &v1,
                                                                                                                     v2)
                                                      })
                                   }))
    }
    pub fn Data_List_Lazy_insertAt_0040118_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_insertAt_0040118_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_insertAt_0040118_002d1.get_or_init(||
                                                              Lazy(Data_List_Lazy_insertAt_0040118.clone()))
    }
    pub fn Data_List_Lazy_insertAt_tco(v: &dyn Any, v1: &dyn Any,
                                       v2: &dyn Any) -> &dyn Any {
        let matchValue = Sharpurs_Prelude::unbox(v);
        let matchValue_1 = Sharpurs_Prelude::unbox(v1);
        let matchValue_2 = Sharpurs_Prelude::unbox(v2);
        match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32, &matchValue) {
            0_i32 =>
            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                x),
                                             xs),
            _ => {
                let x_1 = matchValue_1;
                let go =
                    &Func1::new({
                                    let Data_List_Lazy_insertAt_0040118_002d1
                                        =
                                        Data_List_Lazy_insertAt_0040118_002d1.clone();
                                    let x_1 = x_1.clone();
                                    move |v3|
                                        {
                                            let matchValue_4:
                                                    LrcPtr<Data_List_Lazy_Types_Step> =
                                                Sharpurs_Prelude::unbox(v3);
                                            match matchValue_4.as_ref() {
                                                Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_4_1_0,
                                                                                                             matchValue_4_1_1)
                                                =>
                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_4_1_0,
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_insertAt_0040118_002d1.Value,
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                                                    &&&1_i32)),
                                                                                                                                                                                              &&&x_1),
                                                                                                                                                           &&matchValue_4_1_1))),
                                                _ =>
                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&x_1,
                                                                                                                          &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())),
                                            }
                                        }
                                });
                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                          &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                       &&&go),
                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                       &&&matchValue_2)))
            }
        }
    }
    pub fn Data_List_Lazy_insertAt() -> &dyn Any {
        static Data_List_Lazy_insertAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_insertAt.get_or_init(||
                                                Data_List_Lazy_insertAt_0040118_002d1.Value)
    }
    pub fn Data_List_Lazy_init() -> &dyn Any {
        static Data_List_Lazy_init: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_init.get_or_init(||
                                            {
                                                let go_2 =
                                                    Func0::new({
                                                                   let go_tco
                                                                       =
                                                                       go_tco.clone();
                                                                   move ||
                                                                       &Func1::new({
                                                                                       let go_tco
                                                                                           =
                                                                                           go_tco.clone();
                                                                                       move
                                                                                           |v|
                                                                                           go_tco(v.clone())
                                                                                   })
                                                               });
                                                let go_1 = Lazy(go_2);
                                                fn go_tco(v_1: &dyn Any)
                                                 -> &dyn Any {
                                                    let matchValue:
                                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                                        Sharpurs_Prelude::unbox(v_1);
                                                    if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                        matchValue_1_1)
                                                           =
                                                           matchValue.as_ref()
                                                       {
                                                        if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_null(),
                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                            {
                                                                                                                            Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                         x)
                                                                                                                            =>
                                                                                                                            x.clone(),
                                                                                                                            _
                                                                                                                            =>
                                                                                                                            unreachable!(),
                                                                                                                        }))
                                                           {
                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()))
                                                        } else {
                                                            if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                               {
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                              {
                                                                                                                                                                              Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                           _)
                                                                                                                                                                              =>
                                                                                                                                                                              x.clone(),
                                                                                                                                                                              _
                                                                                                                                                                              =>
                                                                                                                                                                              unreachable!(),
                                                                                                                                                                          })),
                                                                                                 &&go_tco(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                  {
                                                                                                                                                  Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                               x)
                                                                                                                                                  =>
                                                                                                                                                  x.clone(),
                                                                                                                                                  _
                                                                                                                                                  =>
                                                                                                                                                  unreachable!(),
                                                                                                                                              })))
                                                            } else {
                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                            }
                                                        }
                                                    } else {
                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                    }
                                                }
                                                let go = go_1.Value;
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                    &&&go),
                                                                                 &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step())
                                            })
    }
    pub fn Data_List_Lazy_index() -> &dyn Any {
        static Data_List_Lazy_index: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_index.get_or_init(||
                                             &Func1::new(move |xs|
                                                             {
                                                                 let go_2 =
                                                                     Func0::new({
                                                                                    let go_tco
                                                                                        =
                                                                                        go_tco.clone();
                                                                                    move
                                                                                        ||
                                                                                        &Func1::new({
                                                                                                        let go_tco
                                                                                                            =
                                                                                                            go_tco.clone();
                                                                                                        move
                                                                                                            |v|
                                                                                                            Func1::new({
                                                                                                                           let go_tco
                                                                                                                               =
                                                                                                                               go_tco.clone();
                                                                                                                           let v
                                                                                                                               =
                                                                                                                               v.clone();
                                                                                                                           move
                                                                                                                               |v1|
                                                                                                                               go_tco(v)(v1.clone())
                                                                                                                       })
                                                                                                    })
                                                                                });
                                                                 let go_1 =
                                                                     Lazy(go_2);
                                                                 fn go_tco(v_1:
                                                                               _)
                                                                  ->
                                                                      Func1<&dyn Any,
                                                                            &dyn Any> {
                                                                     Func1::new({
                                                                                    let go_tco
                                                                                        =
                                                                                        go_tco.clone();
                                                                                    let v_1
                                                                                        =
                                                                                        v_1.clone();
                                                                                    move
                                                                                        |v1_1|
                                                                                        {
                                                                                            let matchValue:
                                                                                                    LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                                Sharpurs_Prelude::unbox(&&v_1);
                                                                                            let matchValue_1 =
                                                                                                Sharpurs_Prelude::unbox(v1_1);
                                                                                            if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                matchValue_1_1)
                                                                                                   =
                                                                                                   matchValue.as_ref()
                                                                                               {
                                                                                                if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                            &matchValue_1).is_some()
                                                                                                   {
                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                {
                                                                                                                                                                Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                             _)
                                                                                                                                                                =>
                                                                                                                                                                x.clone(),
                                                                                                                                                                _
                                                                                                                                                                =>
                                                                                                                                                                unreachable!(),
                                                                                                                                                            }))
                                                                                                } else {
                                                                                                    if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                        matchValue_1_1)
                                                                                                           =
                                                                                                           matchValue.as_ref()
                                                                                                       {
                                                                                                        go_tco(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                &&&match matchValue.as_ref()
                                                                                                                                                       {
                                                                                                                                                       Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                    x)
                                                                                                                                                       =>
                                                                                                                                                       x.clone(),
                                                                                                                                                       _
                                                                                                                                                       =>
                                                                                                                                                       unreachable!(),
                                                                                                                                                   }))(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                           &&&matchValue_1),
                                                                                                                                                                                        &&&1_i32))
                                                                                                    } else {
                                                                                                        panic!("{}",
                                                                                                               string("Match failure"),)
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                                            }
                                                                                        }
                                                                                })
                                                                 }
                                                                 let go =
                                                                     go_1.Value;
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                     xs))
                                                             }))
    }
    pub fn Data_List_Lazy_head() -> &dyn Any {
        static Data_List_Lazy_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_head.get_or_init(||
                                            &Func1::new(move |xs|
                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                   &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                &&&Func1::new(move
                                                                                                                                                  |v|
                                                                                                                                                  find(string("head"),
                                                                                                                                                       Sharpurs_Prelude::unbox(v)))),
                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                                                                xs))))
    }
    pub fn Data_List_Lazy_transpose_0040140() -> &dyn Any {
        &Func1::new(move |xs|
                        PureScript_Data_List_Lazy::Data_List_Lazy_transpose_tco(xs))
    }
    pub fn Data_List_Lazy_transpose_0040140_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_transpose_0040140_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_transpose_0040140_002d1.get_or_init(||
                                                               Lazy(Data_List_Lazy_transpose_0040140.clone()))
    }
    pub fn Data_List_Lazy_transpose_tco(xs: &dyn Any) -> &dyn Any {
        let matchValue: LrcPtr<Data_Maybe_Maybe> =
            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                       xs));
        if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0) =
               matchValue.as_ref() {
            let activePatternResult =
                Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref() {
                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                       => x.clone(),
                                                       _ => unreachable!(),
                                                   });
            let activePatternResult_1 =
                Sharpurs_Prelude::_007cHasProp_007c__007c(string("head"),
                                                          &activePatternResult);
            if activePatternResult_1.is_some() {
                let activePatternResult_2 =
                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("tail"),
                                                              &activePatternResult);
                if activePatternResult_2.is_some() {
                    let h = getValue(activePatternResult_1);
                    let xss = getValue(activePatternResult_2);
                    let matchValue_1: LrcPtr<Data_Maybe_Maybe> =
                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                   &&&h));
                    if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_1_0)
                           = matchValue_1.as_ref() {
                        let activePatternResult_3 =
                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue_1.as_ref()
                                                                   {
                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                   =>
                                                                   x.clone(),
                                                                   _ =>
                                                                   unreachable!(),
                                                               });
                        let activePatternResult_4 =
                            Sharpurs_Prelude::_007cHasProp_007c__007c(string("head"),
                                                                      &activePatternResult_3);
                        if activePatternResult_4.is_some() {
                            let activePatternResult_5 =
                                Sharpurs_Prelude::_007cHasProp_007c__007c(string("tail"),
                                                                          &activePatternResult_3);
                            if activePatternResult_5.is_some() {
                                let x = getValue(activePatternResult_4);
                                let xs_prime =
                                    getValue(activePatternResult_5);
                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                          &&&x),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_mapMaybe(),
                                                                                                                                                                                                             &&&PureScript_Data_List_Lazy::Data_List_Lazy_head()),
                                                                                                                                                                          &&&xss))),
                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_transpose_0040140_002d1.Value,
                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                          &&&xs_prime),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_mapMaybe(),
                                                                                                                                                                                                             &&&PureScript_Data_List_Lazy::Data_List_Lazy_tail()),
                                                                                                                                                                          &&&xss))))
                            } else {
                                panic!("{}",
                                       string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                            }
                        } else {
                            panic!("{}",
                                   string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                        }
                    } else {
                        Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_transpose_0040140_002d1.Value,
                                                         &&&xss)
                    }
                } else {
                    panic!("{}",
                           string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                }
            } else {
                panic!("{}",
                       string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
            }
        } else { xs.clone() }
    }
    pub fn Data_List_Lazy_transpose() -> &dyn Any {
        static Data_List_Lazy_transpose: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_transpose.get_or_init(||
                                                 Data_List_Lazy_transpose_0040140_002d1.Value)
    }
    pub fn Data_List_Lazy_groupBy_0040144() -> &dyn Any {
        &Func1::new(move |eq|
                        PureScript_Data_List_Lazy::Data_List_Lazy_groupBy_tco(eq))
    }
    pub fn Data_List_Lazy_groupBy_0040144_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_groupBy_0040144_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_groupBy_0040144_002d1.get_or_init(||
                                                             Lazy(Data_List_Lazy_groupBy_0040144.clone()))
    }
    pub fn Data_List_Lazy_groupBy_tco(eq: &dyn Any) -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_groupBy_0040144_002d1 =
                                Data_List_Lazy_groupBy_0040144_002d1.clone();
                            let eq = eq.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    match matchValue.as_ref() {
                                        Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                     matchValue_1_1)
                                        => {
                                            let x = matchValue_1_0.clone();
                                            let matchValue_1 =
                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_span(),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&eq,
                                                                                                                                                                                 &&&x)),
                                                                                                           &&matchValue_1_1));
                                            {
                                                let activePatternResult =
                                                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("init"),
                                                                                              &matchValue_1);
                                                if activePatternResult.is_some()
                                                   {
                                                    let activePatternResult_1 =
                                                        Sharpurs_Prelude::_007cHasProp_007c__007c(string("rest"),
                                                                                                  &matchValue_1);
                                                    if activePatternResult_1.is_some()
                                                       {
                                                        let ys =
                                                            getValue(activePatternResult);
                                                        let zs =
                                                            getValue(activePatternResult_1);
                                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                        let x
                                                                                                                                                                                                                            =
                                                                                                                                                                                                                            x.clone();
                                                                                                                                                                                                                        let ys
                                                                                                                                                                                                                            =
                                                                                                                                                                                                                            ys.clone();
                                                                                                                                                                                                                        move
                                                                                                                                                                                                                            |v2|
                                                                                                                                                                                                                            &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&x,
                                                                                                                                                                                                                                                                                                &ys))
                                                                                                                                                                                                                    }))),
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_groupBy_0040144_002d1.Value,
                                                                                                                                                                                                      &&&eq),
                                                                                                                                                                   &&&zs)))
                                                    } else {
                                                        panic!("{}",
                                                               LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 143_i32,
                                  Data2: 415_i32,}).get_Message(),)
                                                    }
                                                } else {
                                                    panic!("{}",
                                                           LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.List.Lazy.fs"),
                                  Data1: 143_i32,
                                  Data2: 415_i32,}).get_Message(),)
                                                }
                                            }
                                        }
                                        _ =>
                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                     &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                  &&&go)),
                                                                            &&&PureScript_Data_List_Lazy::Data_List_Lazy_unwrap()))
    }
    pub fn Data_List_Lazy_groupBy() -> &dyn Any {
        static Data_List_Lazy_groupBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_groupBy.get_or_init(||
                                               Data_List_Lazy_groupBy_0040144_002d1.Value)
    }
    pub fn Data_List_Lazy_group() -> &dyn Any {
        static Data_List_Lazy_group: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_group.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_groupBy(),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                 dictEq))))
    }
    pub fn Data_List_Lazy_fromStep() -> &dyn Any {
        static Data_List_Lazy_fromStep: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_fromStep.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                    &&&PureScript_Data_Lazy::Data_Lazy_applicativeLazy())))
    }
    pub fn Data_List_Lazy_insertBy_0040152() -> &dyn Any {
        &Func1::new(move |cmp|
                        Func1::new({
                                       let cmp = cmp.clone();
                                       move |x|
                                           Func1::new({
                                                          let x = x.clone();
                                                          move |xs|
                                                              PureScript_Data_List_Lazy::Data_List_Lazy_insertBy_tco(&cmp,
                                                                                                                     &x,
                                                                                                                     xs)
                                                      })
                                   }))
    }
    pub fn Data_List_Lazy_insertBy_0040152_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_insertBy_0040152_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_insertBy_0040152_002d1.get_or_init(||
                                                              Lazy(Data_List_Lazy_insertBy_0040152.clone()))
    }
    pub fn Data_List_Lazy_insertBy_tco(cmp: &dyn Any, x: &dyn Any,
                                       xs: &dyn Any) -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_insertBy_0040152_002d1 =
                                Data_List_Lazy_insertBy_0040152_002d1.clone();
                            let cmp = cmp.clone();
                            let x = x.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    match matchValue.as_ref() {
                                        Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                     matchValue_1_1)
                                        => {
                                            let y = matchValue_1_0.clone();
                                            if let Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                   =
                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&cmp,
                                                                                                                                                 &&&x),
                                                                                                              &&&y)).as_ref()
                                               {
                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&y,
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_insertBy_0040152_002d1.Value,
                                                                                                                                                                                                                                 &&&cmp),
                                                                                                                                                                                              &&&x),
                                                                                                                                                           &&matchValue_1_1)))
                                            } else {
                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&x,
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_fromStep(),
                                                                                                                                                           matchValue)))
                                            }
                                        }
                                        _ =>
                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&x,
                                                                                                                  &PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())),
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                               &&&go),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                               xs)))
    }
    pub fn Data_List_Lazy_insertBy() -> &dyn Any {
        static Data_List_Lazy_insertBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_insertBy.get_or_init(||
                                                Data_List_Lazy_insertBy_0040152_002d1.Value)
    }
    pub fn Data_List_Lazy_insert() -> &dyn Any {
        static Data_List_Lazy_insert: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_insert.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_insertBy(),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                  dictOrd))))
    }
    pub fn Data_List_Lazy_fromFoldable() -> &dyn Any {
        static Data_List_Lazy_fromFoldable: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_fromFoldable.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictFoldable|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                           dictFoldable),
                                                                                                                                        &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons()),
                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())))
    }
    pub fn Data_List_Lazy_foldrLazy() -> &dyn Any {
        static Data_List_Lazy_foldrLazy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_foldrLazy.get_or_init(||
                                                 &Func1::new(move |dictLazy|
                                                                 &Func1::new({
                                                                                 let dictLazy
                                                                                     =
                                                                                     dictLazy.clone();
                                                                                 move
                                                                                     |op|
                                                                                     &Func1::new({
                                                                                                     let op
                                                                                                         =
                                                                                                         op.clone();
                                                                                                     move
                                                                                                         |z|
                                                                                                         {
                                                                                                             let go_2 =
                                                                                                                 Func0::new({
                                                                                                                                let go_tco
                                                                                                                                    =
                                                                                                                                    go_tco.clone();
                                                                                                                                move
                                                                                                                                    ||
                                                                                                                                    &Func1::new({
                                                                                                                                                    let go_tco
                                                                                                                                                        =
                                                                                                                                                        go_tco.clone();
                                                                                                                                                    move
                                                                                                                                                        |xs|
                                                                                                                                                        go_tco(xs.clone())
                                                                                                                                                })
                                                                                                                            });
                                                                                                             let go_1 =
                                                                                                                 Lazy(go_2);
                                                                                                             let go_tco =
                                                                                                                 Func1::new({
                                                                                                                                let z
                                                                                                                                    =
                                                                                                                                    z.clone();
                                                                                                                                move
                                                                                                                                    |xs_1|
                                                                                                                                    fix1(&(move
                                                                                                                                               |go_tco,
                                                                                                                                                xs_1|
                                                                                                                                               {
                                                                                                                                                   let matchValue:
                                                                                                                                                           LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                                                  xs_1));
                                                                                                                                                   match matchValue.as_ref()
                                                                                                                                                       {
                                                                                                                                                       Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor
                                                                                                                                                       =>
                                                                                                                                                       &z,
                                                                                                                                                       Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                    matchValue_1_1)
                                                                                                                                                       =>
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                                                                                                                           &&&dictLazy),
                                                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                                                          let go_tco
                                                                                                                                                                                                              =
                                                                                                                                                                                                              go_tco.clone();
                                                                                                                                                                                                          move
                                                                                                                                                                                                              |v1|
                                                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&op,
                                                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                         Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                               &&go_tco(&match matchValue.as_ref()
                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                             Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                         }))
                                                                                                                                                                                                      })),
                                                                                                                                                   }
                                                                                                                                               }),
                                                                                                                                         xs_1.clone())
                                                                                                                            });
                                                                                                             let go =
                                                                                                                 go_1.Value;
                                                                                                             &go
                                                                                                         }
                                                                                                 })
                                                                             })))
    }
    pub fn Data_List_Lazy_foldM_0040168() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Data_List_Lazy::Data_List_Lazy_foldM_tco(dictMonad))
    }
    pub fn Data_List_Lazy_foldM_0040168_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_foldM_0040168_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_foldM_0040168_002d1.get_or_init(||
                                                           Lazy(Data_List_Lazy_foldM_0040168.clone()))
    }
    pub fn Data_List_Lazy_foldM_tco(dictMonad: &dyn Any) -> &dyn Any {
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Data_List_Lazy_foldM_0040168_002d1 =
                            Data_List_Lazy_foldM_0040168_002d1.clone();
                        let dictMonad = dictMonad.clone();
                        move |f|
                            &Func1::new({
                                            let Data_List_Lazy_foldM_0040168_002d1
                                                =
                                                Data_List_Lazy_foldM_0040168_002d1.clone();
                                            let f = f.clone();
                                            move |b|
                                                &Func1::new({
                                                                let Data_List_Lazy_foldM_0040168_002d1
                                                                    =
                                                                    Data_List_Lazy_foldM_0040168_002d1.clone();
                                                                let b =
                                                                    b.clone();
                                                                move |xs|
                                                                    {
                                                                        let matchValue:
                                                                                LrcPtr<Data_Maybe_Maybe> =
                                                                            Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                                                                       xs));
                                                                        if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                               =
                                                                               matchValue.as_ref()
                                                                           {
                                                                            let activePatternResult =
                                                                                Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                                       {
                                                                                                                       Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                                       =>
                                                                                                                       x.clone(),
                                                                                                                       _
                                                                                                                       =>
                                                                                                                       unreachable!(),
                                                                                                                   });
                                                                            let activePatternResult_1 =
                                                                                Sharpurs_Prelude::_007cHasProp_007c__007c(string("head"),
                                                                                                                          &activePatternResult);
                                                                            if activePatternResult_1.is_some()
                                                                               {
                                                                                let activePatternResult_2 =
                                                                                    Sharpurs_Prelude::_007cHasProp_007c__007c(string("tail"),
                                                                                                                              &activePatternResult);
                                                                                if activePatternResult_2.is_some()
                                                                                   {
                                                                                    let a =
                                                                                        getValue(activePatternResult_1);
                                                                                    let as_var =
                                                                                        getValue(activePatternResult_2);
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                           &&&Bind1),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                              &&&b),
                                                                                                                                                                                           &&&a)),
                                                                                                                     &&&Func1::new({
                                                                                                                                       let Data_List_Lazy_foldM_0040168_002d1
                                                                                                                                           =
                                                                                                                                           Data_List_Lazy_foldM_0040168_002d1.clone();
                                                                                                                                       let as_var
                                                                                                                                           =
                                                                                                                                           as_var.clone();
                                                                                                                                       move
                                                                                                                                           |b_prime|
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_foldM_0040168_002d1.Value,
                                                                                                                                                                                                                                                                                     &&&dictMonad),
                                                                                                                                                                                                                                                  &&&f),
                                                                                                                                                                                                               b_prime),
                                                                                                                                                                            &&&as_var)
                                                                                                                                   }))
                                                                                } else {
                                                                                    panic!("{}",
                                                                                           string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                                                                                }
                                                                            } else {
                                                                                panic!("{}",
                                                                                       string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                                                                            }
                                                                        } else {
                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                &&&Applicative0),
                                                                                                             &&&b)
                                                                        }
                                                                    }
                                                            })
                                        })
                    })
    }
    pub fn Data_List_Lazy_foldM() -> &dyn Any {
        static Data_List_Lazy_foldM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_foldM.get_or_init(||
                                             Data_List_Lazy_foldM_0040168_002d1.Value)
    }
    pub fn Data_List_Lazy_findIndex() -> &dyn Any {
        static Data_List_Lazy_findIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_findIndex.get_or_init(||
                                                 &Func1::new(move |r#fn|
                                                                 {
                                                                     let go_2 =
                                                                         Func0::new({
                                                                                        let go_tco
                                                                                            =
                                                                                            go_tco.clone();
                                                                                        move
                                                                                            ||
                                                                                            &Func1::new({
                                                                                                            let go_tco
                                                                                                                =
                                                                                                                go_tco.clone();
                                                                                                            move
                                                                                                                |n|
                                                                                                                Func1::new({
                                                                                                                               let go_tco
                                                                                                                                   =
                                                                                                                                   go_tco.clone();
                                                                                                                               let n
                                                                                                                                   =
                                                                                                                                   n.clone();
                                                                                                                               move
                                                                                                                                   |list|
                                                                                                                                   go_tco(n)(list.clone())
                                                                                                                           })
                                                                                                        })
                                                                                    });
                                                                     let go_1 =
                                                                         Lazy(go_2);
                                                                     let go_tco =
                                                                         Func1::new({
                                                                                        let r#fn
                                                                                            =
                                                                                            r#fn.clone();
                                                                                        move
                                                                                            |n_1|
                                                                                            fix1(&(move
                                                                                                       |go_tco,
                                                                                                        n_1|
                                                                                                       Func1::new({
                                                                                                                      let go_tco
                                                                                                                          =
                                                                                                                          go_tco.clone();
                                                                                                                      let n_1
                                                                                                                          =
                                                                                                                          n_1.clone();
                                                                                                                      move
                                                                                                                          |list_1|
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                                 &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                                                                                                                                                                 list_1)),
                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                             let go_tco
                                                                                                                                                                                 =
                                                                                                                                                                                 go_tco.clone();
                                                                                                                                                                             move
                                                                                                                                                                                 |o|
                                                                                                                                                                                 {
                                                                                                                                                                                     let matchValue =
                                                                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&r#fn,
                                                                                                                                                                                                                                                   &&find(string("head"),
                                                                                                                                                                                                                                                          Sharpurs_Prelude::unbox(o))));
                                                                                                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                      &matchValue)
                                                                                                                                                                                         {
                                                                                                                                                                                         0_i32
                                                                                                                                                                                         =>
                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                             &&&PureScript_Data_Maybe::Data_Maybe_applicativeMaybe()),
                                                                                                                                                                                                                          &&&n_1),
                                                                                                                                                                                         _
                                                                                                                                                                                         =>
                                                                                                                                                                                         go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                                                    &&&n_1),
                                                                                                                                                                                                                                 &&&1_i32))(find(string("tail"),
                                                                                                                                                                                                                                                 Sharpurs_Prelude::unbox(o))),
                                                                                                                                                                                     }
                                                                                                                                                                                 }
                                                                                                                                                                         }))
                                                                                                                  })),
                                                                                                 n_1.clone())
                                                                                    });
                                                                     let go =
                                                                         go_1.Value;
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                      &&&0_i32)
                                                                 }))
    }
    pub fn Data_List_Lazy_findLastIndex() -> &dyn Any {
        static Data_List_Lazy_findLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_findLastIndex.get_or_init(||
                                                     &Func1::new(move |r#fn|
                                                                     &Func1::new({
                                                                                     let r#fn
                                                                                         =
                                                                                         r#fn.clone();
                                                                                     move
                                                                                         |xs|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                               let xs
                                                                                                                                                                                   =
                                                                                                                                                                                   xs.clone();
                                                                                                                                                                               move
                                                                                                                                                                                   |v|
                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_length(),
                                                                                                                                                                                                                                                                                                                                                                &&&xs)),
                                                                                                                                                                                                                                                                                          &&&1_i32)),
                                                                                                                                                                                                                    v)
                                                                                                                                                                           })),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_findIndex(),
                                                                                                                                                                                                &&&r#fn),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_reverse(),
                                                                                                                                                                                                xs)))
                                                                                 })))
    }
    pub fn Data_List_Lazy_filterM_0040182() -> &dyn Any {
        &Func1::new(move |dictMonad|
                        PureScript_Data_List_Lazy::Data_List_Lazy_filterM_tco(dictMonad))
    }
    pub fn Data_List_Lazy_filterM_0040182_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_filterM_0040182_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_filterM_0040182_002d1.get_or_init(||
                                                             Lazy(Data_List_Lazy_filterM_0040182.clone()))
    }
    pub fn Data_List_Lazy_filterM_tco(dictMonad: &dyn Any) -> &dyn Any {
        let Applicative0 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Applicative0"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        let Bind1 =
            Sharpurs_Prelude::sharpurs_apply(&&find(string("Bind1"),
                                                    Sharpurs_Prelude::unbox(dictMonad)),
                                             &&&Sharpurs_Prelude::Prim_undefined());
        &Func1::new({
                        let Applicative0 = Applicative0.clone();
                        let Bind1 = Bind1.clone();
                        let Data_List_Lazy_filterM_0040182_002d1 =
                            Data_List_Lazy_filterM_0040182_002d1.clone();
                        let dictMonad = dictMonad.clone();
                        move |p|
                            &Func1::new({
                                            let Data_List_Lazy_filterM_0040182_002d1
                                                =
                                                Data_List_Lazy_filterM_0040182_002d1.clone();
                                            let p = p.clone();
                                            move |list|
                                                {
                                                    let matchValue:
                                                            LrcPtr<Data_Maybe_Maybe> =
                                                        Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                                                   list));
                                                    if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                           =
                                                           matchValue.as_ref()
                                                       {
                                                        let activePatternResult =
                                                            Sharpurs_Prelude::_007cUnbox_007c(&match matchValue.as_ref()
                                                                                                   {
                                                                                                   Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(x)
                                                                                                   =>
                                                                                                   x.clone(),
                                                                                                   _
                                                                                                   =>
                                                                                                   unreachable!(),
                                                                                               });
                                                        let activePatternResult_1 =
                                                            Sharpurs_Prelude::_007cHasProp_007c__007c(string("head"),
                                                                                                      &activePatternResult);
                                                        if activePatternResult_1.is_some()
                                                           {
                                                            let activePatternResult_2 =
                                                                Sharpurs_Prelude::_007cHasProp_007c__007c(string("tail"),
                                                                                                          &activePatternResult);
                                                            if activePatternResult_2.is_some()
                                                               {
                                                                let x =
                                                                    getValue(activePatternResult_1);
                                                                let xs =
                                                                    getValue(activePatternResult_2);
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                       &&&Bind1),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                       &&&x)),
                                                                                                 &&&Func1::new({
                                                                                                                   let Data_List_Lazy_filterM_0040182_002d1
                                                                                                                       =
                                                                                                                       Data_List_Lazy_filterM_0040182_002d1.clone();
                                                                                                                   let x
                                                                                                                       =
                                                                                                                       x.clone();
                                                                                                                   let xs
                                                                                                                       =
                                                                                                                       xs.clone();
                                                                                                                   move
                                                                                                                       |b|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                                              &&&Bind1),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_filterM_0040182_002d1.Value,
                                                                                                                                                                                                                                                                                                    &&&dictMonad),
                                                                                                                                                                                                                                                                 &&&p),
                                                                                                                                                                                                                              &&&xs)),
                                                                                                                                                        &&&Func1::new({
                                                                                                                                                                          let b
                                                                                                                                                                              =
                                                                                                                                                                              b.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |xs_prime|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                  &&&Applicative0),
                                                                                                                                                                                                               &&{
                                                                                                                                                                                                                     let matchValue_1 =
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(&&b);
                                                                                                                                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                                                      &matchValue_1)
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                         0_i32
                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                             &&&x),
                                                                                                                                                                                                                                                          xs_prime),
                                                                                                                                                                                                                         _
                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                         xs_prime.clone(),
                                                                                                                                                                                                                     }
                                                                                                                                                                                                                 })
                                                                                                                                                                      }))
                                                                                                               }))
                                                            } else {
                                                                panic!("{}",
                                                                       string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                                                            }
                                                        } else {
                                                            panic!("{}",
                                                                   string("Match failure: PureScript_Data_Maybe.Data_Maybe_Maybe"),)
                                                        }
                                                    } else {
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                            &&&Applicative0),
                                                                                         &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil())
                                                    }
                                                }
                                        })
                    })
    }
    pub fn Data_List_Lazy_filterM() -> &dyn Any {
        static Data_List_Lazy_filterM: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_filterM.get_or_init(||
                                               Data_List_Lazy_filterM_0040182_002d1.Value)
    }
    pub fn Data_List_Lazy_filter_0040192() -> &dyn Any {
        &Func1::new(move |p|
                        PureScript_Data_List_Lazy::Data_List_Lazy_filter_tco(p))
    }
    pub fn Data_List_Lazy_filter_0040192_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_filter_0040192_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_filter_0040192_002d1.get_or_init(||
                                                            Lazy(Data_List_Lazy_filter_0040192.clone()))
    }
    pub fn Data_List_Lazy_filter_tco(p: &dyn Any) -> &dyn Any {
        let go_2 =
            Func0::new({
                           let go_tco = go_tco.clone();
                           move ||
                               &Func1::new({
                                               let go_tco = go_tco.clone();
                                               move |v| go_tco(v.clone())
                                           })
                       });
        let go_1 = Lazy(go_2);
        let go_tco =
            Func1::new({
                           let Data_List_Lazy_filter_0040192_002d1 =
                               Data_List_Lazy_filter_0040192_002d1.clone();
                           let p = p.clone();
                           move |v_1|
                               {
                                   let v_1 = v_1.clone();
                                   '_go_tco:
                                       loop  {
                                           break '_go_tco
                                               ({
                                                    let matchValue:
                                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                                        Sharpurs_Prelude::unbox(&&v_1);
                                                    if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                        matchValue_1_1)
                                                           =
                                                           matchValue.as_ref()
                                                       {
                                                        if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                            {
                                                                                                                            Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                         _)
                                                                                                                            =>
                                                                                                                            x.clone(),
                                                                                                                            _
                                                                                                                            =>
                                                                                                                            unreachable!(),
                                                                                                                        }))
                                                           {
                                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                                           {
                                                                                                                                           Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                        _)
                                                                                                                                           =>
                                                                                                                                           x.clone(),
                                                                                                                                           _
                                                                                                                                           =>
                                                                                                                                           unreachable!(),
                                                                                                                                       },
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_filter_0040192_002d1.Value,
                                                                                                                                                                                                          &&&p),
                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                              {
                                                                                                                                                                              Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                           x)
                                                                                                                                                                              =>
                                                                                                                                                                              x.clone(),
                                                                                                                                                                              _
                                                                                                                                                                              =>
                                                                                                                                                                              unreachable!(),
                                                                                                                                                                          })))
                                                        } else {
                                                            if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                               {
                                                                let v_1_temp =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                     &&&match matchValue.as_ref()
                                                                                                            {
                                                                                                            Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                         x)
                                                                                                            =>
                                                                                                            x.clone(),
                                                                                                            _
                                                                                                            =>
                                                                                                            unreachable!(),
                                                                                                        });
                                                                v_1.set(v_1_temp);
                                                                continue
                                                                    '_go_tco
                                                            } else {
                                                                panic!("{}",
                                                                       string("Match failure: PureScript_Data_List_Lazy_Types.Data_List_Lazy_Types_Step"),)
                                                            }
                                                        }
                                                    } else {
                                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                                    }
                                                }) ;
                                       }
                               }
                       });
        let go = go_1.Value;
        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                     &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                  &&&go)),
                                                                            &&&PureScript_Data_List_Lazy::Data_List_Lazy_unwrap()))
    }
    pub fn Data_List_Lazy_filter() -> &dyn Any {
        static Data_List_Lazy_filter: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_filter.get_or_init(||
                                              Data_List_Lazy_filter_0040192_002d1.Value)
    }
    pub fn Data_List_Lazy_intersectBy() -> &dyn Any {
        static Data_List_Lazy_intersectBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_intersectBy.get_or_init(||
                                                   &Func1::new(move |eq|
                                                                   &Func1::new({
                                                                                   let eq
                                                                                       =
                                                                                       eq.clone();
                                                                                   move
                                                                                       |xs|
                                                                                       &Func1::new({
                                                                                                       let xs
                                                                                                           =
                                                                                                           xs.clone();
                                                                                                       move
                                                                                                           |ys|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_filter(),
                                                                                                                                                                               &&&Func1::new({
                                                                                                                                                                                                 let ys
                                                                                                                                                                                                     =
                                                                                                                                                                                                     ys.clone();
                                                                                                                                                                                                 move
                                                                                                                                                                                                     |x|
                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_any(),
                                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&eq,
                                                                                                                                                                                                                                                                                                            x)),
                                                                                                                                                                                                                                      &&&ys)
                                                                                                                                                                                             })),
                                                                                                                                            &&&xs)
                                                                                                   })
                                                                               })))
    }
    pub fn Data_List_Lazy_intersect() -> &dyn Any {
        static Data_List_Lazy_intersect: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_intersect.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_intersectBy(),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                     dictEq))))
    }
    pub fn Data_List_Lazy_nubByEq_0040200() -> &dyn Any {
        &Func1::new(move |eq|
                        PureScript_Data_List_Lazy::Data_List_Lazy_nubByEq_tco(eq))
    }
    pub fn Data_List_Lazy_nubByEq_0040200_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_nubByEq_0040200_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_nubByEq_0040200_002d1.get_or_init(||
                                                             Lazy(Data_List_Lazy_nubByEq_0040200.clone()))
    }
    pub fn Data_List_Lazy_nubByEq_tco(eq: &dyn Any) -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_nubByEq_0040200_002d1 =
                                Data_List_Lazy_nubByEq_0040200_002d1.clone();
                            let eq = eq.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    match matchValue.as_ref() {
                                        Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                     matchValue_1_1)
                                        => {
                                            let x = matchValue_1_0.clone();
                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&x,
                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_nubByEq_0040200_002d1.Value,
                                                                                                                                                                                          &&&eq),
                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_filter(),
                                                                                                                                                                                                                             &&&Func1::new({
                                                                                                                                                                                                                                               let x
                                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                                   x.clone();
                                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                                   |y|
                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&eq,
                                                                                                                                                                                                                                                                                                                                                          &&&x),
                                                                                                                                                                                                                                                                                                                       y))
                                                                                                                                                                                                                                           })),
                                                                                                                                                                                          &&matchValue_1_1))))
                                        }
                                        _ =>
                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                     &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                  &&&go)),
                                                                            &&&PureScript_Data_List_Lazy::Data_List_Lazy_unwrap()))
    }
    pub fn Data_List_Lazy_nubByEq() -> &dyn Any {
        static Data_List_Lazy_nubByEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_nubByEq.get_or_init(||
                                               Data_List_Lazy_nubByEq_0040200_002d1.Value)
    }
    pub fn Data_List_Lazy_nubEq() -> &dyn Any {
        static Data_List_Lazy_nubEq: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_nubEq.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_nubByEq(),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                 dictEq))))
    }
    pub fn Data_List_Lazy_eqPattern() -> &dyn Any {
        static Data_List_Lazy_eqPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_eqPattern.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 {
                                                                     let eqList =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_eqList(),
                                                                                                          dictEq);
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                      &&&add(string("eq"),
                                                                                                             &&Func1::new({
                                                                                                                              let eqList
                                                                                                                                  =
                                                                                                                                  eqList.clone();
                                                                                                                              move
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
                                                                                                                                                                                                                                                                 &&&eqList),
                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                           &&&matchValue_1)
                                                                                                                                                      }
                                                                                                                                              })
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))
                                                                 }))
    }
    pub fn Data_List_Lazy_ordPattern() -> &dyn Any {
        static Data_List_Lazy_ordPattern: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_ordPattern.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  {
                                                                      let ordList =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_ordList(),
                                                                                                           dictOrd);
                                                                      let eqPattern1 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_eqPattern(),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                     Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                       &&&add(string("compare"),
                                                                                                              &&Func1::new({
                                                                                                                               let ordList
                                                                                                                                   =
                                                                                                                                   ordList.clone();
                                                                                                                               move
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
                                                                                                                                                                                                                                                                  &&&ordList),
                                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                                            &&&matchValue_1)
                                                                                                                                                       }
                                                                                                                                               })
                                                                                                                           }),
                                                                                                              add(string("Eq0"),
                                                                                                                  &&Func1::new({
                                                                                                                                   let eqPattern1
                                                                                                                                       =
                                                                                                                                       eqPattern1.clone();
                                                                                                                                   move
                                                                                                                                       |usd__unused|
                                                                                                                                       &eqPattern1
                                                                                                                               }),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>())))
                                                                  }))
    }
    pub fn Data_List_Lazy_elemLastIndex() -> &dyn Any {
        static Data_List_Lazy_elemLastIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_elemLastIndex.get_or_init(||
                                                     &Func1::new(move |dictEq|
                                                                     &Func1::new({
                                                                                     let dictEq
                                                                                         =
                                                                                         dictEq.clone();
                                                                                     move
                                                                                         |x|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_findLastIndex(),
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
    pub fn Data_List_Lazy_elemIndex() -> &dyn Any {
        static Data_List_Lazy_elemIndex: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_elemIndex.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 &Func1::new({
                                                                                 let dictEq
                                                                                     =
                                                                                     dictEq.clone();
                                                                                 move
                                                                                     |x|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_findIndex(),
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
    pub fn Data_List_Lazy_dropWhile() -> &dyn Any {
        static Data_List_Lazy_dropWhile: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_dropWhile.get_or_init(||
                                                 &Func1::new(move |p|
                                                                 {
                                                                     let go_2 =
                                                                         Func0::new({
                                                                                        let go_tco
                                                                                            =
                                                                                            go_tco.clone();
                                                                                        move
                                                                                            ||
                                                                                            &Func1::new({
                                                                                                            let go_tco
                                                                                                                =
                                                                                                                go_tco.clone();
                                                                                                            move
                                                                                                                |v|
                                                                                                                go_tco(v.clone())
                                                                                                        })
                                                                                    });
                                                                     let go_1 =
                                                                         Lazy(go_2);
                                                                     let go_tco =
                                                                         Func1::new({
                                                                                        let p
                                                                                            =
                                                                                            p.clone();
                                                                                        move
                                                                                            |v_1|
                                                                                            {
                                                                                                let v_1 =
                                                                                                    v_1.clone();
                                                                                                '_go_tco:
                                                                                                    loop 
                                                                                                         {
                                                                                                        break
                                                                                                            '_go_tco
                                                                                                            ({
                                                                                                                 let matchValue:
                                                                                                                         LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                                                     Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                 if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                                                                     matchValue_1_1)
                                                                                                                        =
                                                                                                                        matchValue.as_ref()
                                                                                                                    {
                                                                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                         Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                         =>
                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                         _
                                                                                                                                                                                         =>
                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                     }))
                                                                                                                        {
                                                                                                                         let v_1_temp =
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                     Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                  x)
                                                                                                                                                                     =>
                                                                                                                                                                     x.clone(),
                                                                                                                                                                     _
                                                                                                                                                                     =>
                                                                                                                                                                     unreachable!(),
                                                                                                                                                                 });
                                                                                                                         v_1.set(v_1_temp);
                                                                                                                         continue
                                                                                                                             '_go_tco

                                                                                                                     } else {
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_fromStep(),
                                                                                                                                                          &&&matchValue)
                                                                                                                     }
                                                                                                                 } else {
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_fromStep(),
                                                                                                                                                      &&&matchValue)
                                                                                                                 }
                                                                                                             })
                                                                                                            ;
                                                                                                    }
                                                                                            }
                                                                                    });
                                                                     let go =
                                                                         go_1.Value;
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                         &&&go),
                                                                                                      &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step())
                                                                 }))
    }
    pub fn Data_List_Lazy_drop() -> &dyn Any {
        static Data_List_Lazy_drop: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_drop.get_or_init(||
                                            &Func1::new(move |n|
                                                            {
                                                                let go_2 =
                                                                    Func0::new({
                                                                                   let go_tco
                                                                                       =
                                                                                       go_tco.clone();
                                                                                   move
                                                                                       ||
                                                                                       &Func1::new({
                                                                                                       let go_tco
                                                                                                           =
                                                                                                           go_tco.clone();
                                                                                                       move
                                                                                                           |v|
                                                                                                           Func1::new({
                                                                                                                          let go_tco
                                                                                                                              =
                                                                                                                              go_tco.clone();
                                                                                                                          let v
                                                                                                                              =
                                                                                                                              v.clone();
                                                                                                                          move
                                                                                                                              |v1|
                                                                                                                              go_tco(v)(v1.clone())
                                                                                                                      })
                                                                                                   })
                                                                               });
                                                                let go_1 =
                                                                    Lazy(go_2);
                                                                fn go_tco(v_1:
                                                                              _)
                                                                 ->
                                                                     Func1<&dyn Any,
                                                                           &dyn Any> {
                                                                    Func1::new({
                                                                                   let go_tco
                                                                                       =
                                                                                       go_tco.clone();
                                                                                   let v_1
                                                                                       =
                                                                                       v_1.clone();
                                                                                   move
                                                                                       |v1_1|
                                                                                       {
                                                                                           let matchValue =
                                                                                               Sharpurs_Prelude::unbox(&&v_1);
                                                                                           let matchValue_1:
                                                                                                   LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                               Sharpurs_Prelude::unbox(v1_1);
                                                                                           match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                           &matchValue)
                                                                                               {
                                                                                               0_i32
                                                                                               =>
                                                                                               xs.clone(),
                                                                                               _
                                                                                               =>
                                                                                               match matchValue_1.as_ref()
                                                                                                   {
                                                                                                   Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                matchValue_1_1_1)
                                                                                                   =>
                                                                                                   go_tco(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                 &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                              &&&matchValue),
                                                                                                                                           &&&1_i32))(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                                                                       &&matchValue_1_1_1)),
                                                                                                   _
                                                                                                   =>
                                                                                                   &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor),
                                                                                               },
                                                                                           }
                                                                                       }
                                                                               })
                                                                }
                                                                let go =
                                                                    go_1.Value;
                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                       &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List()),
                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                                                                                                             n))),
                                                                                                                                    &&&PureScript_Data_List_Lazy::Data_List_Lazy_unwrap()))
                                                            }))
    }
    pub fn Data_List_Lazy_slice() -> &dyn Any {
        static Data_List_Lazy_slice: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_slice.get_or_init(||
                                             &Func1::new(move |start|
                                                             &Func1::new({
                                                                             let start
                                                                                 =
                                                                                 start.clone();
                                                                             move
                                                                                 |end_var|
                                                                                 &Func1::new({
                                                                                                 let end_var
                                                                                                     =
                                                                                                     end_var.clone();
                                                                                                 move
                                                                                                     |xs|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_take(),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                               &&&end_var),
                                                                                                                                                                                                            &&&start)),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_drop(),
                                                                                                                                                                                                            &&&start),
                                                                                                                                                                         xs))
                                                                                             })
                                                                         })))
    }
    pub fn Data_List_Lazy_deleteBy_0040232() -> &dyn Any {
        &Func1::new(move |eq|
                        Func1::new({
                                       let eq = eq.clone();
                                       move |x|
                                           Func1::new({
                                                          let x = x.clone();
                                                          move |xs|
                                                              PureScript_Data_List_Lazy::Data_List_Lazy_deleteBy_tco(&eq,
                                                                                                                     &x,
                                                                                                                     xs)
                                                      })
                                   }))
    }
    pub fn Data_List_Lazy_deleteBy_0040232_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_deleteBy_0040232_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_deleteBy_0040232_002d1.get_or_init(||
                                                              Lazy(Data_List_Lazy_deleteBy_0040232.clone()))
    }
    pub fn Data_List_Lazy_deleteBy_tco(eq: &dyn Any, x: &dyn Any,
                                       xs: &dyn Any) -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_deleteBy_0040232_002d1 =
                                Data_List_Lazy_deleteBy_0040232_002d1.clone();
                            let eq = eq.clone();
                            let x = x.clone();
                            move |v|
                                {
                                    let matchValue:
                                            LrcPtr<Data_List_Lazy_Types_Step> =
                                        Sharpurs_Prelude::unbox(v);
                                    if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                        matchValue_1_1)
                                           = matchValue.as_ref() {
                                        if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&eq,
                                                                                                                                        &&&x),
                                                                                                     &&&match matchValue.as_ref()
                                                                                                            {
                                                                                                            Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                         _)
                                                                                                            =>
                                                                                                            x.clone(),
                                                                                                            _
                                                                                                            =>
                                                                                                            unreachable!(),
                                                                                                        }))
                                           {
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                             &&&match matchValue.as_ref()
                                                                                    {
                                                                                    Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                 x)
                                                                                    =>
                                                                                    x.clone(),
                                                                                    _
                                                                                    =>
                                                                                    unreachable!(),
                                                                                })
                                        } else {
                                            if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                               {
                                                &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&match matchValue.as_ref()
                                                                                                                               {
                                                                                                                               Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                            _)
                                                                                                                               =>
                                                                                                                               x.clone(),
                                                                                                                               _
                                                                                                                               =>
                                                                                                                               unreachable!(),
                                                                                                                           },
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_deleteBy_0040232_002d1.Value,
                                                                                                                                                                                                                                 &&&eq),
                                                                                                                                                                                              &&&x),
                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                  {
                                                                                                                                                                  Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                               x)
                                                                                                                                                                  =>
                                                                                                                                                                  x.clone(),
                                                                                                                                                                  _
                                                                                                                                                                  =>
                                                                                                                                                                  unreachable!(),
                                                                                                                                                              })))
                                            } else {
                                                panic!("{}",
                                                       string("Match failure: PureScript_Data_List_Lazy_Types.Data_List_Lazy_Types_Step"),)
                                            }
                                        }
                                    } else {
                                        &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                    }
                                }
                        });
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                               &&&go),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                               xs)))
    }
    pub fn Data_List_Lazy_deleteBy() -> &dyn Any {
        static Data_List_Lazy_deleteBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_deleteBy.get_or_init(||
                                                Data_List_Lazy_deleteBy_0040232_002d1.Value)
    }
    pub fn Data_List_Lazy_unionBy() -> &dyn Any {
        static Data_List_Lazy_unionBy: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_unionBy.get_or_init(||
                                               &Func1::new(move |eq|
                                                               &Func1::new({
                                                                               let eq
                                                                                   =
                                                                                   eq.clone();
                                                                               move
                                                                                   |xs|
                                                                                   &Func1::new({
                                                                                                   let xs
                                                                                                       =
                                                                                                       xs.clone();
                                                                                                   move
                                                                                                       |ys|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                              &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                           &&&xs),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                                    &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_deleteBy(),
                                                                                                                                                                                                                                                                                                                       &&&eq))),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_nubByEq(),
                                                                                                                                                                                                                                                                                    &&&eq),
                                                                                                                                                                                                                                                 ys)),
                                                                                                                                                                           &&&xs))
                                                                                               })
                                                                           })))
    }
    pub fn Data_List_Lazy_union() -> &dyn Any {
        static Data_List_Lazy_union: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_union.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_unionBy(),
                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                 dictEq))))
    }
    pub fn Data_List_Lazy_deleteAt_0040240() -> &dyn Any {
        &Func1::new(move |n|
                        Func1::new({
                                       let n = n.clone();
                                       move |xs|
                                           PureScript_Data_List_Lazy::Data_List_Lazy_deleteAt_tco(&n,
                                                                                                  xs)
                                   }))
    }
    pub fn Data_List_Lazy_deleteAt_0040240_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_deleteAt_0040240_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_deleteAt_0040240_002d1.get_or_init(||
                                                              Lazy(Data_List_Lazy_deleteAt_0040240.clone()))
    }
    pub fn Data_List_Lazy_deleteAt_tco(n: &dyn Any, xs: &dyn Any)
     -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_deleteAt_0040240_002d1 =
                                Data_List_Lazy_deleteAt_0040240_002d1.clone();
                            move |v|
                                &Func1::new({
                                                let Data_List_Lazy_deleteAt_0040240_002d1
                                                    =
                                                    Data_List_Lazy_deleteAt_0040240_002d1.clone();
                                                let v = v.clone();
                                                move |v1|
                                                    {
                                                        let matchValue =
                                                            Sharpurs_Prelude::unbox(&&v);
                                                        let matchValue_1:
                                                                LrcPtr<Data_List_Lazy_Types_Step> =
                                                            Sharpurs_Prelude::unbox(v1);
                                                        if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                            matchValue_1_1_1)
                                                               =
                                                               matchValue_1.as_ref()
                                                           {
                                                            if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                        &matchValue).is_some()
                                                               {
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                        {
                                                                                                        Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                     x)
                                                                                                        =>
                                                                                                        x.clone(),
                                                                                                        _
                                                                                                        =>
                                                                                                        unreachable!(),
                                                                                                    })
                                                            } else {
                                                                if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                    matchValue_1_1_1)
                                                                       =
                                                                       matchValue_1.as_ref()
                                                                   {
                                                                    &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                _)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                                   _
                                                                                                                                                   =>
                                                                                                                                                   unreachable!(),
                                                                                                                                               },
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_deleteAt_0040240_002d1.Value,
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                        &&&matchValue),
                                                                                                                                                                                                                                                     &&&1_i32)),
                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                      =>
                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                      _
                                                                                                                                                                                      =>
                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                  })))
                                                                } else {
                                                                    panic!("{}",
                                                                           string("Match failure"),)
                                                                }
                                                            }
                                                        } else {
                                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                                        }
                                                    }
                                            })
                        });
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                  n)),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                               xs)))
    }
    pub fn Data_List_Lazy_deleteAt() -> &dyn Any {
        static Data_List_Lazy_deleteAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_deleteAt.get_or_init(||
                                                Data_List_Lazy_deleteAt_0040240_002d1.Value)
    }
    pub fn Data_List_Lazy_delete() -> &dyn Any {
        static Data_List_Lazy_delete: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_delete.get_or_init(||
                                              &Func1::new(move |dictEq|
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_deleteBy(),
                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                  dictEq))))
    }
    pub fn Data_List_Lazy_difference() -> &dyn Any {
        static Data_List_Lazy_difference: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_difference.get_or_init(||
                                                  &Func1::new(move |dictEq|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                      &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_foldableList()),
                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_delete(),
                                                                                                                                                                         dictEq)))))
    }
    pub fn Data_List_Lazy_cycle() -> &dyn Any {
        static Data_List_Lazy_cycle: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_cycle.get_or_init(||
                                             &Func1::new(move |xs|
                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_fix(),
                                                                                                                                 &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_lazyList()),
                                                                                              &&&Func1::new({
                                                                                                                let xs
                                                                                                                    =
                                                                                                                    xs.clone();
                                                                                                                move
                                                                                                                    |ys|
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                                        &&&xs),
                                                                                                                                                     ys)
                                                                                                            }))))
    }
    pub fn Data_List_Lazy_concatMap() -> &dyn Any {
        static Data_List_Lazy_concatMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_concatMap.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_bindList())))
    }
    pub fn Data_List_Lazy_concat() -> &dyn Any {
        static Data_List_Lazy_concat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_concat.get_or_init(||
                                              &Func1::new(move |v|
                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                     &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_bindList()),
                                                                                                                                  v),
                                                                                               &&&PureScript_Data_List_Lazy::Data_List_Lazy_identity())))
    }
    pub fn Data_List_Lazy_catMaybes() -> &dyn Any {
        static Data_List_Lazy_catMaybes: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_catMaybes.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_mapMaybe(),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                     &&&PureScript_Control_Category::Control_Category_categoryFn())))
    }
    pub fn Data_List_Lazy_alterAt_0040256() -> &dyn Any {
        &Func1::new(move |n|
                        Func1::new({
                                       let n = n.clone();
                                       move |f|
                                           Func1::new({
                                                          let f = f.clone();
                                                          move |xs|
                                                              PureScript_Data_List_Lazy::Data_List_Lazy_alterAt_tco(&n,
                                                                                                                    &f,
                                                                                                                    xs)
                                                      })
                                   }))
    }
    pub fn Data_List_Lazy_alterAt_0040256_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_List_Lazy_alterAt_0040256_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_List_Lazy_alterAt_0040256_002d1.get_or_init(||
                                                             Lazy(Data_List_Lazy_alterAt_0040256.clone()))
    }
    pub fn Data_List_Lazy_alterAt_tco(n: &dyn Any, f: &dyn Any, xs: &dyn Any)
     -> &dyn Any {
        let go =
            &Func1::new({
                            let Data_List_Lazy_alterAt_0040256_002d1 =
                                Data_List_Lazy_alterAt_0040256_002d1.clone();
                            let f = f.clone();
                            move |v|
                                &Func1::new({
                                                let Data_List_Lazy_alterAt_0040256_002d1
                                                    =
                                                    Data_List_Lazy_alterAt_0040256_002d1.clone();
                                                let v = v.clone();
                                                move |v1|
                                                    {
                                                        let matchValue =
                                                            Sharpurs_Prelude::unbox(&&v);
                                                        let matchValue_1:
                                                                LrcPtr<Data_List_Lazy_Types_Step> =
                                                            Sharpurs_Prelude::unbox(v1);
                                                        if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                            matchValue_1_1_1)
                                                               =
                                                               matchValue_1.as_ref()
                                                           {
                                                            if Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                        &matchValue).is_some()
                                                               {
                                                                let matchValue_3:
                                                                        LrcPtr<Data_Maybe_Maybe> =
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                      {
                                                                                                                                      Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                   _)
                                                                                                                                      =>
                                                                                                                                      x.clone(),
                                                                                                                                      _
                                                                                                                                      =>
                                                                                                                                      unreachable!(),
                                                                                                                                  }));
                                                                match matchValue_3.as_ref()
                                                                    {
                                                                    Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_3_1_0)
                                                                    =>
                                                                    &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_3_1_0,
                                                                                                                                              &match matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                x)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                                   _
                                                                                                                                                   =>
                                                                                                                                                   unreachable!(),
                                                                                                                                               })),
                                                                    _ =>
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                     &&&match matchValue_1.as_ref()
                                                                                                            {
                                                                                                            Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                         x)
                                                                                                            =>
                                                                                                            x.clone(),
                                                                                                            _
                                                                                                            =>
                                                                                                            unreachable!(),
                                                                                                        }),
                                                                }
                                                            } else {
                                                                if let Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_1_0,
                                                                                                                                    matchValue_1_1_1)
                                                                       =
                                                                       matchValue_1.as_ref()
                                                                   {
                                                                    &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                   Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(x,
                                                                                                                                                                                                                _)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                                   _
                                                                                                                                                   =>
                                                                                                                                                   unreachable!(),
                                                                                                                                               },
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_List_Lazy_alterAt_0040256_002d1.Value,
                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                                                                        &&&1_i32)),
                                                                                                                                                                                                                  &&&f),
                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                      Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(_,
                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                      =>
                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                      _
                                                                                                                                                                                      =>
                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                  })))
                                                                } else {
                                                                    panic!("{}",
                                                                           string("Match failure"),)
                                                                }
                                                            }
                                                        } else {
                                                            &LrcPtr::new(Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Nilusd_Ctor)
                                                        }
                                                    }
                                            })
                        });
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_List(),
                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Data_Lazy::Data_Lazy_functorLazy()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&go,
                                                                                                                                                  n)),
                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_unwrap(),
                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                               xs)))
    }
    pub fn Data_List_Lazy_alterAt() -> &dyn Any {
        static Data_List_Lazy_alterAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_alterAt.get_or_init(||
                                               Data_List_Lazy_alterAt_0040256_002d1.Value)
    }
    pub fn Data_List_Lazy_modifyAt() -> &dyn Any {
        static Data_List_Lazy_modifyAt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_modifyAt.get_or_init(||
                                                &Func1::new(move |n|
                                                                &Func1::new({
                                                                                let n
                                                                                    =
                                                                                    n.clone();
                                                                                move
                                                                                    |f|
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_alterAt(),
                                                                                                                                                        &&&n),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                             |usd__arg1|
                                                                                                                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                        f))
                                                                            })))
    }
}

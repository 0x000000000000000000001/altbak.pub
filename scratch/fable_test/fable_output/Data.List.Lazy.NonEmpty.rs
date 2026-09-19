pub mod PureScript_Data_List_Lazy_NonEmpty {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_720e12db::PureScript_Data_Lazy;
    use crate::module_44df2f32::PureScript_Data_List_Lazy_Types;
    use crate::module_44df2f32::PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_Step;
    use crate::module_529acc77::PureScript_Data_List_Lazy;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_d6a130cf::PureScript_Data_NonEmpty::Data_NonEmpty_NonEmpty;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_e7bd458d::PureScript_Data_Tuple::Data_Tuple_Tuple;
    use crate::module_98c530c5::PureScript_Data_Unfoldable;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_List_Lazy_NonEmpty_uncons() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_uncons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_uncons.get_or_init(||
                                                       &Func1::new(move |v|
                                                                       {
                                                                           let nel =
                                                                               Sharpurs_Prelude::unbox(v);
                                                                           let matchValue:
                                                                                   LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                          &&&nel));
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
    pub fn Data_List_Lazy_NonEmpty_toList() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_toList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_toList.get_or_init(||
                                                       &Func1::new(move |v|
                                                                       {
                                                                           let nel =
                                                                               Sharpurs_Prelude::unbox(v);
                                                                           let matchValue:
                                                                                   LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                          &&&nel));
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
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
    pub fn Data_List_Lazy_NonEmpty_toUnfoldable() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_toUnfoldable: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_toUnfoldable.get_or_init(||
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
                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_uncons(),
                                                                                                                                                                                                                                                                          xs))))),
                                                                                                              &&&PureScript_Data_List_Lazy_NonEmpty::Data_List_Lazy_NonEmpty_toList())))
    }
    pub fn Data_List_Lazy_NonEmpty_tail() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_tail: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_tail.get_or_init(||
                                                     &Func1::new(move |v|
                                                                     {
                                                                         let nel =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         &match Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                           &&&nel)).as_ref()
                                                                              {
                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                     x)
                                                                              =>
                                                                              x.clone(),
                                                                          }
                                                                     }))
    }
    pub fn Data_List_Lazy_NonEmpty_singleton() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_singleton: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_singleton.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_applicativeNonEmptyList()))
    }
    pub fn Data_List_Lazy_NonEmpty_repeat() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_repeat: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_repeat.get_or_init(||
                                                       &Func1::new(move |x|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                           &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                           &&&Func1::new({
                                                                                                                                                             let x
                                                                                                                                                                 =
                                                                                                                                                                 x.clone();
                                                                                                                                                             move
                                                                                                                                                                 |v|
                                                                                                                                                                 &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&x,
                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_repeat(),
                                                                                                                                                                                                                                                                      &&&x)))
                                                                                                                                                         })))))
    }
    pub fn Data_List_Lazy_NonEmpty_length() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_length: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_length.get_or_init(||
                                                       &Func1::new(move |v|
                                                                       {
                                                                           let nel =
                                                                               Sharpurs_Prelude::unbox(v);
                                                                           let matchValue:
                                                                                   LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                          &&&nel));
                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                  &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                               &&&1_i32),
                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_length(),
                                                                                                                                               &&&match matchValue.as_ref()
                                                                                                                                                      {
                                                                                                                                                      Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                             x)
                                                                                                                                                      =>
                                                                                                                                                      x.clone(),
                                                                                                                                                  }))
                                                                       }))
    }
    pub fn Data_List_Lazy_NonEmpty_last() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_last.get_or_init(||
                                                     &Func1::new(move |v|
                                                                     {
                                                                         let nel =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         let matchValue:
                                                                                 LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                        &&&nel));
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromMaybe(),
                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                    {
                                                                                                                                                    Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                           _)
                                                                                                                                                    =>
                                                                                                                                                    x.clone(),
                                                                                                                                                }),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_last(),
                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                    {
                                                                                                                                                    Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                           x)
                                                                                                                                                    =>
                                                                                                                                                    x.clone(),
                                                                                                                                                }))
                                                                     }))
    }
    pub fn Data_List_Lazy_NonEmpty_iterate() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_iterate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_iterate.get_or_init(||
                                                        &Func1::new(move |f|
                                                                        &Func1::new({
                                                                                        let f
                                                                                            =
                                                                                            f.clone();
                                                                                        move
                                                                                            |x|
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList()),
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                  let x
                                                                                                                                                                                      =
                                                                                                                                                                                      x.clone();
                                                                                                                                                                                  move
                                                                                                                                                                                      |v|
                                                                                                                                                                                      &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&x,
                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_iterate(),
                                                                                                                                                                                                                                                                                                                              &&&f),
                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                              &&&x))))
                                                                                                                                                                              })))
                                                                                    })))
    }
    pub fn Data_List_Lazy_NonEmpty_init() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_init: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_init.get_or_init(||
                                                     &Func1::new(move |v|
                                                                     {
                                                                         let nel =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         let matchValue:
                                                                                 LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                             Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                        &&&nel));
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_maybe(),
                                                                                                                                                                                &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_nil()),
                                                                                                                                             &&&Func1::new(move
                                                                                                                                                               |v2|
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                   &&&match matchValue.as_ref()
                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                          Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                v2))),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_init(),
                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                    {
                                                                                                                                                    Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                           x)
                                                                                                                                                    =>
                                                                                                                                                    x.clone(),
                                                                                                                                                }))
                                                                     }))
    }
    pub fn Data_List_Lazy_NonEmpty_head() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_head: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_head.get_or_init(||
                                                     &Func1::new(move |v|
                                                                     {
                                                                         let nel =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         &match Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                           &&&nel)).as_ref()
                                                                              {
                                                                              Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                     _)
                                                                              =>
                                                                              x.clone(),
                                                                          }
                                                                     }))
    }
    pub fn Data_List_Lazy_NonEmpty_fromList() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_fromList: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_fromList.get_or_init(||
                                                         &Func1::new(move |l|
                                                                         {
                                                                             let matchValue:
                                                                                     LrcPtr<Data_List_Lazy_Types_Step> =
                                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_step(),
                                                                                                                                            l));
                                                                             match matchValue.as_ref()
                                                                                 {
                                                                                 Data_List_Lazy_Types_Step::Data_List_Lazy_Types_Consusd_Ctor(matchValue_1_0,
                                                                                                                                              matchValue_1_1)
                                                                                 =>
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                                                                              |v1|
                                                                                                                                                                                                                              &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                                                                                  matchValue_1_1))))))),
                                                                                 _
                                                                                 =>
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                             }
                                                                         }))
    }
    pub fn Data_List_Lazy_NonEmpty_fromFoldable() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_fromFoldable: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_fromFoldable.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictFoldable|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                 &&&PureScript_Data_List_Lazy_NonEmpty::Data_List_Lazy_NonEmpty_fromList()),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_fromFoldable(),
                                                                                                                                                 dictFoldable))))
    }
    pub fn Data_List_Lazy_NonEmpty_cons() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_cons: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_cons.get_or_init(||
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
                                                                                             let matchValue_1 =
                                                                                                 Sharpurs_Prelude::unbox(v);
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                 &&&Func1::new({
                                                                                                                                                                                   let matchValue_1
                                                                                                                                                                                       =
                                                                                                                                                                                       matchValue_1.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |v1|
                                                                                                                                                                                       {
                                                                                                                                                                                           let matchValue_3:
                                                                                                                                                                                                   LrcPtr<Data_NonEmpty_NonEmpty> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_force(),
                                                                                                                                                                                                                                                          &&&matchValue_1));
                                                                                                                                                                                           &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(&matchValue,
                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_cons(),
                                                                                                                                                                                                                                                                                                                                   &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                                          Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                                &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                       Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                   })))
                                                                                                                                                                                       }
                                                                                                                                                                               })))
                                                                                         }
                                                                                 })))
    }
    pub fn Data_List_Lazy_NonEmpty_concatMap() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_concatMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_concatMap.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                              &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_bindNonEmptyList())))
    }
    pub fn Data_List_Lazy_NonEmpty_appendFoldable() -> &dyn Any {
        static Data_List_Lazy_NonEmpty_appendFoldable:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_List_Lazy_NonEmpty_appendFoldable.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictFoldable|
                                                                               &Func1::new({
                                                                                               let dictFoldable
                                                                                                   =
                                                                                                   dictFoldable.clone();
                                                                                               move
                                                                                                   |nel|
                                                                                                   &Func1::new({
                                                                                                                   let nel
                                                                                                                       =
                                                                                                                       nel.clone();
                                                                                                                   move
                                                                                                                       |ys|
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_NonEmptyList(),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Lazy::Data_Lazy_defer(),
                                                                                                                                                                                           &&&Func1::new({
                                                                                                                                                                                                             let ys
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 ys.clone();
                                                                                                                                                                                                             move
                                                                                                                                                                                                                 |v|
                                                                                                                                                                                                                 &LrcPtr::new(Data_NonEmpty_NonEmpty::Data_NonEmpty_NonEmptyusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_NonEmpty::Data_List_Lazy_NonEmpty_head(),
                                                                                                                                                                                                                                                                                                                      &&&nel),
                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_List_Lazy_Types::Data_List_Lazy_Types_semigroupList()),
                                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy_NonEmpty::Data_List_Lazy_NonEmpty_tail(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&nel)),
                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_List_Lazy::Data_List_Lazy_fromFoldable(),
                                                                                                                                                                                                                                                                                                                                                                                            &&&dictFoldable),
                                                                                                                                                                                                                                                                                                                                                         &&&ys))))
                                                                                                                                                                                                         })))
                                                                                                               })
                                                                                           })))
    }
}

pub mod PureScript_Data_Functor_Coproduct {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_173929b2::PureScript_Data_Either::Data_Either_Either;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Functor_Coproduct_Coproduct() -> &dyn Any {
        static Data_Functor_Coproduct_Coproduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_Coproduct.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Functor_Coproduct_showCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_showCoproduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Coproduct_showCoproduct.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictShow|
                                                                             &Func1::new({
                                                                                             let dictShow
                                                                                                 =
                                                                                                 dictShow.clone();
                                                                                             move
                                                                                                 |dictShow1|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                                  &&&add(string("show"),
                                                                                                                                         &&Func1::new({
                                                                                                                                                          let dictShow1
                                                                                                                                                              =
                                                                                                                                                              dictShow1.clone();
                                                                                                                                                          move
                                                                                                                                                              |v|
                                                                                                                                                              {
                                                                                                                                                                  let matchValue:
                                                                                                                                                                          LrcPtr<Data_Either_Either> =
                                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                                  match matchValue.as_ref()
                                                                                                                                                                      {
                                                                                                                                                                      Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_1_0)
                                                                                                                                                                      =>
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                          &&&string("(right ")),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictShow1),
                                                                                                                                                                                                                                                                                                                &&matchValue_1_0)),
                                                                                                                                                                                                                                          &&&string(")"))),
                                                                                                                                                                      Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_0_0)
                                                                                                                                                                      =>
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                          &&&string("(left ")),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                   &&&dictShow),
                                                                                                                                                                                                                                                                                                                &&matchValue_0_0)),
                                                                                                                                                                                                                                          &&&string(")"))),
                                                                                                                                                                  }
                                                                                                                                                              }
                                                                                                                                                      }),
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>()))
                                                                                         })))
    }
    pub fn Data_Functor_Coproduct_right() -> &dyn Any {
        static Data_Functor_Coproduct_right: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_right.get_or_init(||
                                                     &Func1::new(move |ga|
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct(),
                                                                                                      &&&LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(ga.clone())))))
    }
    pub fn Data_Functor_Coproduct_newtypeCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_newtypeCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_newtypeCoproduct.get_or_init(||
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                 &&&add(string("Coercible0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &Sharpurs_Prelude::Prim_undefined()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))
    }
    pub fn Data_Functor_Coproduct_left() -> &dyn Any {
        static Data_Functor_Coproduct_left: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_left.get_or_init(||
                                                    &Func1::new(move |fa|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct(),
                                                                                                     &&&LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(fa.clone())))))
    }
    pub fn Data_Functor_Coproduct_functorCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_functorCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_functorCoproduct.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictFunctor|
                                                                                &Func1::new({
                                                                                                let dictFunctor
                                                                                                    =
                                                                                                    dictFunctor.clone();
                                                                                                move
                                                                                                    |dictFunctor1|
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                                                     &&&add(string("map"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let dictFunctor1
                                                                                                                                                                 =
                                                                                                                                                                 dictFunctor1.clone();
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
                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                                         let matchValue_1 =
                                                                                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                                                                                         let f1 =
                                                                                                                                                                                             matchValue;
                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct(),
                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorEither()),
                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                         &&&dictFunctor),
                                                                                                                                                                                                                                                                                                                                                                      &&&f1)),
                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                      &&&dictFunctor1),
                                                                                                                                                                                                                                                                                                                                   &&&f1)),
                                                                                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                                                                                     }
                                                                                                                                                                             })
                                                                                                                                                         }),
                                                                                                                                            empty::<string,
                                                                                                                                                    &dyn Any>()))
                                                                                            })))
    }
    pub fn Data_Functor_Coproduct_eq1Coproduct() -> &dyn Any {
        static Data_Functor_Coproduct_eq1Coproduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Coproduct_eq1Coproduct.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictEq1|
                                                                            &Func1::new({
                                                                                            let dictEq1
                                                                                                =
                                                                                                dictEq1.clone();
                                                                                            move
                                                                                                |dictEq11|
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                                                 &&&add(string("eq1"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let dictEq11
                                                                                                                                                             =
                                                                                                                                                             dictEq11.clone();
                                                                                                                                                         move
                                                                                                                                                             |dictEq|
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
                                                                                                                                                                                                     |v1|
                                                                                                                                                                                                     {
                                                                                                                                                                                                         let matchValue =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                         let matchValue_1 =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                         let matchValue_3:
                                                                                                                                                                                                                 LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&matchValue);
                                                                                                                                                                                                         let matchValue_4:
                                                                                                                                                                                                                 LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                             Sharpurs_Prelude::unbox(&&matchValue_1);
                                                                                                                                                                                                         if let Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_3_1_0)
                                                                                                                                                                                                                =
                                                                                                                                                                                                                matchValue_3.as_ref()
                                                                                                                                                                                                            {
                                                                                                                                                                                                             if let Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    matchValue_4.as_ref()
                                                                                                                                                                                                                {
                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                                                           &&&dictEq11),
                                                                                                                                                                                                                                                                                                                        &&&dictEq),
                                                                                                                                                                                                                                                                                     &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                            Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                            _
                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                  &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                         Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                     })
                                                                                                                                                                                                             } else {
                                                                                                                                                                                                                 &false
                                                                                                                                                                                                             }
                                                                                                                                                                                                         } else {
                                                                                                                                                                                                             if let Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_4_0_0)
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    matchValue_4.as_ref()
                                                                                                                                                                                                                {
                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                                                                                                                                                                                                                                           &&&dictEq1),
                                                                                                                                                                                                                                                                                                                        &&&dictEq),
                                                                                                                                                                                                                                                                                     &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                            Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                            _
                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                  &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                         Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                     })
                                                                                                                                                                                                             } else {
                                                                                                                                                                                                                 &false
                                                                                                                                                                                                             }
                                                                                                                                                                                                         }
                                                                                                                                                                                                     }
                                                                                                                                                                                             })
                                                                                                                                                                         })
                                                                                                                                                     }),
                                                                                                                                        empty::<string,
                                                                                                                                                &dyn Any>()))
                                                                                        })))
    }
    pub fn Data_Functor_Coproduct_eqCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_eqCoproduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_eqCoproduct.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictEq1|
                                                                           {
                                                                               let eq1Coproduct1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_eq1Coproduct(),
                                                                                                                    dictEq1);
                                                                               &Func1::new({
                                                                                               let eq1Coproduct1
                                                                                                   =
                                                                                                   eq1Coproduct1.clone();
                                                                                               move
                                                                                                   |dictEq11|
                                                                                                   {
                                                                                                       let eq1 =
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq1(),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&eq1Coproduct1,
                                                                                                                                                                               dictEq11));
                                                                                                       &Func1::new({
                                                                                                                       let eq1
                                                                                                                           =
                                                                                                                           eq1.clone();
                                                                                                                       move
                                                                                                                           |dictEq|
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                                                            &&&add(string("eq"),
                                                                                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&eq1,
                                                                                                                                                                                                     dictEq),
                                                                                                                                                                   empty::<string,
                                                                                                                                                                           &dyn Any>()))
                                                                                                                   })
                                                                                                   }
                                                                                           })
                                                                           }))
    }
    pub fn Data_Functor_Coproduct_ord1Coproduct() -> &dyn Any {
        static Data_Functor_Coproduct_ord1Coproduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Coproduct_ord1Coproduct.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictOrd1|
                                                                             {
                                                                                 let eq1Coproduct1 =
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_eq1Coproduct(),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                 &Func1::new({
                                                                                                 let dictOrd1
                                                                                                     =
                                                                                                     dictOrd1.clone();
                                                                                                 let eq1Coproduct1
                                                                                                     =
                                                                                                     eq1Coproduct1.clone();
                                                                                                 move
                                                                                                     |dictOrd11|
                                                                                                     {
                                                                                                         let eq1Coproduct2 =
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&eq1Coproduct1,
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                                        Sharpurs_Prelude::unbox(dictOrd11)),
                                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                                                          &&&add(string("compare1"),
                                                                                                                                                 &&Func1::new({
                                                                                                                                                                  let dictOrd11
                                                                                                                                                                      =
                                                                                                                                                                      dictOrd11.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |dictOrd|
                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                      let dictOrd
                                                                                                                                                                                          =
                                                                                                                                                                                          dictOrd.clone();
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
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                                                  let matchValue_1 =
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                                                  let matchValue_3:
                                                                                                                                                                                                                          LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&matchValue);
                                                                                                                                                                                                                  let matchValue_4:
                                                                                                                                                                                                                          LrcPtr<Data_Either_Either> =
                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&matchValue_1);
                                                                                                                                                                                                                  if let Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_3_1_0)
                                                                                                                                                                                                                         =
                                                                                                                                                                                                                         matchValue_3.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                      if let Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_4_1_0)
                                                                                                                                                                                                                             =
                                                                                                                                                                                                                             matchValue_4.as_ref()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                                                                    &&&dictOrd11),
                                                                                                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                                                                                                              &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                     Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                           &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                  Data_Either_Either::Data_Either_Rightusd_Ctor(x)
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                      } else {
                                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                                      }
                                                                                                                                                                                                                  } else {
                                                                                                                                                                                                                      if let Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_4_0_0)
                                                                                                                                                                                                                             =
                                                                                                                                                                                                                             matchValue_4.as_ref()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                                                                                                                                                                                                                                                    &&&dictOrd1),
                                                                                                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                                                                                                              &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                     Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                           &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                  Data_Either_Either::Data_Either_Leftusd_Ctor(x)
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                                                              })
                                                                                                                                                                                                                      } else {
                                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                                      }
                                                                                                                                                                                                                  }
                                                                                                                                                                                                              }
                                                                                                                                                                                                      })
                                                                                                                                                                                  })
                                                                                                                                                              }),
                                                                                                                                                 add(string("Eq10"),
                                                                                                                                                     &&Func1::new({
                                                                                                                                                                      let eq1Coproduct2
                                                                                                                                                                          =
                                                                                                                                                                          eq1Coproduct2.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |usd__unused|
                                                                                                                                                                          &eq1Coproduct2
                                                                                                                                                                  }),
                                                                                                                                                     empty::<string,
                                                                                                                                                             &dyn Any>())))
                                                                                                     }
                                                                                             })
                                                                             }))
    }
    pub fn Data_Functor_Coproduct_ordCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_ordCoproduct: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Functor_Coproduct_ordCoproduct.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictOrd1|
                                                                            {
                                                                                let ord1Coproduct1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_ord1Coproduct(),
                                                                                                                     dictOrd1);
                                                                                let eqCoproduct1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_eqCoproduct(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                               Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                &Func1::new({
                                                                                                let eqCoproduct1
                                                                                                    =
                                                                                                    eqCoproduct1.clone();
                                                                                                let ord1Coproduct1
                                                                                                    =
                                                                                                    ord1Coproduct1.clone();
                                                                                                move
                                                                                                    |dictOrd11|
                                                                                                    {
                                                                                                        let compare1 =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare1(),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&ord1Coproduct1,
                                                                                                                                                                                dictOrd11));
                                                                                                        let eqCoproduct2 =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&eqCoproduct1,
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq10"),
                                                                                                                                                                                       Sharpurs_Prelude::unbox(dictOrd11)),
                                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                        &Func1::new({
                                                                                                                        let compare1
                                                                                                                            =
                                                                                                                            compare1.clone();
                                                                                                                        let eqCoproduct2
                                                                                                                            =
                                                                                                                            eqCoproduct2.clone();
                                                                                                                        move
                                                                                                                            |dictOrd|
                                                                                                                            {
                                                                                                                                let eqCoproduct3 =
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&eqCoproduct2,
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                                                               Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                                                                 &&&add(string("compare"),
                                                                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&compare1,
                                                                                                                                                                                                          dictOrd),
                                                                                                                                                                        add(string("Eq0"),
                                                                                                                                                                            &&Func1::new({
                                                                                                                                                                                             let eqCoproduct3
                                                                                                                                                                                                 =
                                                                                                                                                                                                 eqCoproduct3.clone();
                                                                                                                                                                                             move
                                                                                                                                                                                                 |usd__unused|
                                                                                                                                                                                                 &eqCoproduct3
                                                                                                                                                                                         }),
                                                                                                                                                                            empty::<string,
                                                                                                                                                                                    &dyn Any>())))
                                                                                                                            }
                                                                                                                    })
                                                                                                    }
                                                                                            })
                                                                            }))
    }
    pub fn Data_Functor_Coproduct_coproduct() -> &dyn Any {
        static Data_Functor_Coproduct_coproduct: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_coproduct.get_or_init(||
                                                         &Func1::new(move |v|
                                                                         &Func1::new({
                                                                                         let v
                                                                                             =
                                                                                             v.clone();
                                                                                         move
                                                                                             |v1|
                                                                                             &Func1::new({
                                                                                                             let v1
                                                                                                                 =
                                                                                                                 v1.clone();
                                                                                                             move
                                                                                                                 |v2|
                                                                                                                 {
                                                                                                                     let matchValue =
                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                     let matchValue_1 =
                                                                                                                         Sharpurs_Prelude::unbox(&&v1);
                                                                                                                     let matchValue_2:
                                                                                                                             LrcPtr<Data_Either_Either> =
                                                                                                                         Sharpurs_Prelude::unbox(v2);
                                                                                                                     match matchValue_2.as_ref()
                                                                                                                         {
                                                                                                                         Data_Either_Either::Data_Either_Rightusd_Ctor(matchValue_2_1_0)
                                                                                                                         =>
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                          &&matchValue_2_1_0),
                                                                                                                         Data_Either_Either::Data_Either_Leftusd_Ctor(matchValue_2_0_0)
                                                                                                                         =>
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                          &&matchValue_2_0_0),
                                                                                                                     }
                                                                                                                 }
                                                                                                         })
                                                                                     })))
    }
    pub fn Data_Functor_Coproduct_extendCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_extendCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_extendCoproduct.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictExtend|
                                                                               {
                                                                                   let functorCoproduct1 =
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_functorCoproduct(),
                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                  Sharpurs_Prelude::unbox(dictExtend)),
                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                   &Func1::new({
                                                                                                   let dictExtend
                                                                                                       =
                                                                                                       dictExtend.clone();
                                                                                                   let functorCoproduct1
                                                                                                       =
                                                                                                       functorCoproduct1.clone();
                                                                                                   move
                                                                                                       |dictExtend1|
                                                                                                       {
                                                                                                           let functorCoproduct2 =
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&functorCoproduct1,
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                                                                                          Sharpurs_Prelude::unbox(dictExtend1)),
                                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                                                                            &&&add(string("extend"),
                                                                                                                                                   &&Func1::new({
                                                                                                                                                                    let dictExtend1
                                                                                                                                                                        =
                                                                                                                                                                        dictExtend1.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |f|
                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                            &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                       |usd__arg1|
                                                                                                                                                                                                                                                                                                                                                                       &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                                                                                                                        &&&dictExtend),
                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                           f),
                                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                             |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                                                                                                                                             &LrcPtr::new(Data_Either_Either::Data_Either_Leftusd_Ctor(usd__arg1_1.clone())))))))),
                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                    |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_2.clone())))),
                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                                                                                     &&&dictExtend1),
                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                        f),
                                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct()),
                                                                                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                          |usd__arg1_3|
                                                                                                                                                                                                                                                                                                                                                                                                          &LrcPtr::new(Data_Either_Either::Data_Either_Rightusd_Ctor(usd__arg1_3.clone())))))))))
                                                                                                                                                                }),
                                                                                                                                                   add(string("Functor0"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let functorCoproduct2
                                                                                                                                                                            =
                                                                                                                                                                            functorCoproduct2.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |usd__unused|
                                                                                                                                                                            &functorCoproduct2
                                                                                                                                                                    }),
                                                                                                                                                       empty::<string,
                                                                                                                                                               &dyn Any>())))
                                                                                                       }
                                                                                               })
                                                                               }))
    }
    pub fn Data_Functor_Coproduct_comonadCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_comonadCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_comonadCoproduct.get_or_init(||
                                                                &Func1::new(move
                                                                                |dictComonad|
                                                                                {
                                                                                    let extract =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                         dictComonad);
                                                                                    let extendCoproduct1 =
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_extendCoproduct(),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                                   Sharpurs_Prelude::unbox(dictComonad)),
                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                    &Func1::new({
                                                                                                    let extendCoproduct1
                                                                                                        =
                                                                                                        extendCoproduct1.clone();
                                                                                                    let extract
                                                                                                        =
                                                                                                        extract.clone();
                                                                                                    move
                                                                                                        |dictComonad1|
                                                                                                        {
                                                                                                            let extendCoproduct2 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&extendCoproduct1,
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Extend0"),
                                                                                                                                                                                           Sharpurs_Prelude::unbox(dictComonad1)),
                                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                                                                                             &&&add(string("extract"),
                                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_coproduct(),
                                                                                                                                                                                                                         &&&extract),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_extract(),
                                                                                                                                                                                                                         dictComonad1)),
                                                                                                                                                    add(string("Extend0"),
                                                                                                                                                        &&Func1::new({
                                                                                                                                                                         let extendCoproduct2
                                                                                                                                                                             =
                                                                                                                                                                             extendCoproduct2.clone();
                                                                                                                                                                         move
                                                                                                                                                                             |usd__unused|
                                                                                                                                                                             &extendCoproduct2
                                                                                                                                                                     }),
                                                                                                                                                        empty::<string,
                                                                                                                                                                &dyn Any>())))
                                                                                                        }
                                                                                                })
                                                                                }))
    }
    pub fn Data_Functor_Coproduct_bihoistCoproduct() -> &dyn Any {
        static Data_Functor_Coproduct_bihoistCoproduct:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Functor_Coproduct_bihoistCoproduct.get_or_init(||
                                                                &Func1::new(move
                                                                                |natF|
                                                                                &Func1::new({
                                                                                                let natF
                                                                                                    =
                                                                                                    natF.clone();
                                                                                                move
                                                                                                    |natG|
                                                                                                    &Func1::new({
                                                                                                                    let natG
                                                                                                                        =
                                                                                                                        natG.clone();
                                                                                                                    move
                                                                                                                        |v|
                                                                                                                        {
                                                                                                                            let matchValue =
                                                                                                                                Sharpurs_Prelude::unbox(&&natF);
                                                                                                                            let matchValue_1 =
                                                                                                                                Sharpurs_Prelude::unbox(&&natG);
                                                                                                                            let matchValue_2 =
                                                                                                                                Sharpurs_Prelude::unbox(v);
                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Coproduct::Data_Functor_Coproduct_Coproduct(),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Bifunctor::Data_Bifunctor_bifunctorEither()),
                                                                                                                                                                                                                                                                      &&&matchValue),
                                                                                                                                                                                                                                   &&&matchValue_1),
                                                                                                                                                                                                &&&matchValue_2))
                                                                                                                        }
                                                                                                                })
                                                                                            })))
    }
}

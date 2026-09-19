pub mod PureScript_Data_Tuple {
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
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_209e0d2c::PureScript_Control_Lazy;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_13840e4f::PureScript_Data_BooleanAlgebra;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_cb047b05::PureScript_Data_CommutativeRing;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep;
    use crate::module_b37db3cd::PureScript_Data_Generic_Rep::Data_Generic_Rep_Product;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_9be97d93::PureScript_Data_Unit;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug,)]
    pub enum Data_Tuple_Tuple {
        Data_Tuple_Tupleusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for PureScript_Data_Tuple::Data_Tuple_Tuple {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Tuple_Tuple() -> &dyn Any {
        static Data_Tuple_Tuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_Tuple.get_or_init(||
                                         &Func1::new(move |usd__arg1|
                                                         Func1::new({
                                                                        let usd__arg1
                                                                            =
                                                                            usd__arg1.clone();
                                                                        move
                                                                            |usd__arg2|
                                                                            &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                           usd__arg2.clone()))
                                                                    })))
    }
    pub fn Data_Tuple_uncurry() -> &dyn Any {
        static Data_Tuple_uncurry: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_uncurry.get_or_init(||
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
                                                                                           LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                       &&&match matchValue_1.as_ref()
                                                                                                                                                              {
                                                                                                                                                              PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                _)
                                                                                                                                                              =>
                                                                                                                                                              x.clone(),
                                                                                                                                                          }),
                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                           {
                                                                                                                           PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                             x)
                                                                                                                           =>
                                                                                                                           x.clone(),
                                                                                                                       })
                                                                               }
                                                                       })))
    }
    pub fn Data_Tuple_swap() -> &dyn Any {
        static Data_Tuple_swap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_swap.get_or_init(||
                                        &Func1::new(move |v|
                                                        {
                                                            let matchValue:
                                                                    LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                Sharpurs_Prelude::unbox(v);
                                                            &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                {
                                                                                                                                                PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                  x)
                                                                                                                                                =>
                                                                                                                                                x.clone(),
                                                                                                                                            },
                                                                                                                                           &match matchValue.as_ref()
                                                                                                                                                {
                                                                                                                                                PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                  _)
                                                                                                                                                =>
                                                                                                                                                x.clone(),
                                                                                                                                            }))
                                                        }))
    }
    pub fn Data_Tuple_snd() -> &dyn Any {
        static Data_Tuple_snd: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Tuple_snd.get_or_init(||
                                       &Func1::new(move |v|
                                                       &match Sharpurs_Prelude::unbox(v).as_ref()
                                                            {
                                                            PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                              x)
                                                            => x.clone(),
                                                        }))
    }
    pub fn Data_Tuple_showTuple() -> &dyn Any {
        static Data_Tuple_showTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_showTuple.get_or_init(||
                                             &Func1::new(move |dictShow|
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
                                                                                                                                                          LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                      &&&string("(Tuple ")),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                               &&&dictShow),
                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                   PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                               })),
                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                            &&&string(" ")),
                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                     &&&dictShow1),
                                                                                                                                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                                                                                         PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                                                                                     })),
                                                                                                                                                                                                                                                                                            &&&string(")")))))
                                                                                                                                              }
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))
                                                                         })))
    }
    pub fn Data_Tuple_semiringTuple() -> &dyn Any {
        static Data_Tuple_semiringTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_semiringTuple.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictSemiring|
                                                                 &Func1::new({
                                                                                 let dictSemiring
                                                                                     =
                                                                                     dictSemiring.clone();
                                                                                 move
                                                                                     |dictSemiring1|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                                                      &&&add(string("add"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let dictSemiring1
                                                                                                                                                  =
                                                                                                                                                  dictSemiring1.clone();
                                                                                                                                              move
                                                                                                                                                  |v|
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let v
                                                                                                                                                                      =
                                                                                                                                                                      v.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |v1|
                                                                                                                                                                      {
                                                                                                                                                                          let matchValue:
                                                                                                                                                                                  LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                              Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                          let matchValue_1:
                                                                                                                                                                                  LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                              Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                          &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                &&&dictSemiring),
                                                                                                                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                    PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                }),
                                                                                                                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                 PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                &&&dictSemiring1),
                                                                                                                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                                                    PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                                                }),
                                                                                                                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                 PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                             })))
                                                                                                                                                                      }
                                                                                                                                                              })
                                                                                                                                          }),
                                                                                                                             add(string("one"),
                                                                                                                                 &&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                                                                                  &&&dictSemiring),
                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                                                                                  dictSemiring1))),
                                                                                                                                 add(string("mul"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let dictSemiring1
                                                                                                                                                          =
                                                                                                                                                          dictSemiring1.clone();
                                                                                                                                                      move
                                                                                                                                                          |v_1|
                                                                                                                                                          &Func1::new({
                                                                                                                                                                          let v_1
                                                                                                                                                                              =
                                                                                                                                                                              v_1.clone();
                                                                                                                                                                          move
                                                                                                                                                                              |v1_1|
                                                                                                                                                                              {
                                                                                                                                                                                  let matchValue_3:
                                                                                                                                                                                          LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                  let matchValue_4:
                                                                                                                                                                                          LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                      Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                  &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                                                                        &&&dictSemiring),
                                                                                                                                                                                                                                                                                                                                     &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                            PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                                                                  &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                         PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                                                                        &&&dictSemiring1),
                                                                                                                                                                                                                                                                                                                                     &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                            PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                                                                  &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                         PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                     })))
                                                                                                                                                                              }
                                                                                                                                                                      })
                                                                                                                                                  }),
                                                                                                                                     add(string("zero"),
                                                                                                                                         &&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                                                                                                          &&&dictSemiring),
                                                                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                                                                                                          dictSemiring1))),
                                                                                                                                         empty::<string,
                                                                                                                                                 &dyn Any>())))))
                                                                             })))
    }
    pub fn Data_Tuple_semigroupoidTuple() -> &dyn Any {
        static Data_Tuple_semigroupoidTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_semigroupoidTuple.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_Semigroupoidusd_Dict(),
                                                                                      &&&add(string("compose"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              &Func1::new({
                                                                                                                              let v
                                                                                                                                  =
                                                                                                                                  v.clone();
                                                                                                                              move
                                                                                                                                  |v1|
                                                                                                                                  {
                                                                                                                                      let matchValue:
                                                                                                                                              LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                      &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match Sharpurs_Prelude::unbox(v1).as_ref()
                                                                                                                                                                                                                          {
                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                      },
                                                                                                                                                                                                                     &match matchValue.as_ref()
                                                                                                                                                                                                                          {
                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                      }))
                                                                                                                                  }
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Tuple_semigroupTuple() -> &dyn Any {
        static Data_Tuple_semigroupTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_semigroupTuple.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictSemigroup|
                                                                  &Func1::new({
                                                                                  let dictSemigroup
                                                                                      =
                                                                                      dictSemigroup.clone();
                                                                                  move
                                                                                      |dictSemigroup1|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                       &&&add(string("append"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let dictSemigroup1
                                                                                                                                                   =
                                                                                                                                                   dictSemigroup1.clone();
                                                                                                                                               move
                                                                                                                                                   |v|
                                                                                                                                                   &Func1::new({
                                                                                                                                                                   let v
                                                                                                                                                                       =
                                                                                                                                                                       v.clone();
                                                                                                                                                                   move
                                                                                                                                                                       |v1|
                                                                                                                                                                       {
                                                                                                                                                                           let matchValue:
                                                                                                                                                                                   LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                               Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                           let matchValue_1:
                                                                                                                                                                                   LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                               Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                           &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                     PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                  PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                 &&&dictSemigroup1),
                                                                                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                     PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                  PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                              })))
                                                                                                                                                                       }
                                                                                                                                                               })
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>()))
                                                                              })))
    }
    pub fn Data_Tuple_ringTuple() -> &dyn Any {
        static Data_Tuple_ringTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_ringTuple.get_or_init(||
                                             &Func1::new(move |dictRing|
                                                             {
                                                                 let semiringTuple1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_semiringTuple(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                                Sharpurs_Prelude::unbox(dictRing)),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 &Func1::new({
                                                                                 let dictRing
                                                                                     =
                                                                                     dictRing.clone();
                                                                                 let semiringTuple1
                                                                                     =
                                                                                     semiringTuple1.clone();
                                                                                 move
                                                                                     |dictRing1|
                                                                                     {
                                                                                         let semiringTuple2 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&semiringTuple1,
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                                                        Sharpurs_Prelude::unbox(dictRing1)),
                                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                                                                          &&&add(string("sub"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let dictRing1
                                                                                                                                                      =
                                                                                                                                                      dictRing1.clone();
                                                                                                                                                  move
                                                                                                                                                      |v|
                                                                                                                                                      &Func1::new({
                                                                                                                                                                      let v
                                                                                                                                                                          =
                                                                                                                                                                          v.clone();
                                                                                                                                                                      move
                                                                                                                                                                          |v1|
                                                                                                                                                                          {
                                                                                                                                                                              let matchValue:
                                                                                                                                                                                      LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                              let matchValue_1:
                                                                                                                                                                                      LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                  Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                              &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                    &&&dictRing),
                                                                                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                        PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                     PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                                    &&&dictRing1),
                                                                                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                                                        PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                     PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                 })))
                                                                                                                                                                          }
                                                                                                                                                                  })
                                                                                                                                              }),
                                                                                                                                 add(string("Semiring0"),
                                                                                                                                     &&Func1::new({
                                                                                                                                                      let semiringTuple2
                                                                                                                                                          =
                                                                                                                                                          semiringTuple2.clone();
                                                                                                                                                      move
                                                                                                                                                          |usd__unused|
                                                                                                                                                          &semiringTuple2
                                                                                                                                                  }),
                                                                                                                                     empty::<string,
                                                                                                                                             &dyn Any>())))
                                                                                     }
                                                                             })
                                                             }))
    }
    pub fn Data_Tuple_monoidTuple() -> &dyn Any {
        static Data_Tuple_monoidTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_monoidTuple.get_or_init(||
                                               &Func1::new(move |dictMonoid|
                                                               {
                                                                   let semigroupTuple1 =
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_semigroupTuple(),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                  Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                   &Func1::new({
                                                                                   let dictMonoid
                                                                                       =
                                                                                       dictMonoid.clone();
                                                                                   let semigroupTuple1
                                                                                       =
                                                                                       semigroupTuple1.clone();
                                                                                   move
                                                                                       |dictMonoid1|
                                                                                       {
                                                                                           let semigroupTuple2 =
                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&semigroupTuple1,
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                                          Sharpurs_Prelude::unbox(dictMonoid1)),
                                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                            &&&add(string("mempty"),
                                                                                                                                   &&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                    &&&dictMonoid),
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                                                                    dictMonoid1))),
                                                                                                                                   add(string("Semigroup0"),
                                                                                                                                       &&Func1::new({
                                                                                                                                                        let semigroupTuple2
                                                                                                                                                            =
                                                                                                                                                            semigroupTuple2.clone();
                                                                                                                                                        move
                                                                                                                                                            |usd__unused|
                                                                                                                                                            &semigroupTuple2
                                                                                                                                                    }),
                                                                                                                                       empty::<string,
                                                                                                                                               &dyn Any>())))
                                                                                       }
                                                                               })
                                                               }))
    }
    pub fn Data_Tuple_heytingAlgebraTuple() -> &dyn Any {
        static Data_Tuple_heytingAlgebraTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_heytingAlgebraTuple.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictHeytingAlgebra|
                                                                       &Func1::new({
                                                                                       let dictHeytingAlgebra
                                                                                           =
                                                                                           dictHeytingAlgebra.clone();
                                                                                       move
                                                                                           |dictHeytingAlgebra1|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_HeytingAlgebrausd_Dict(),
                                                                                                                            &&&add(string("tt"),
                                                                                                                                   &&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                                                                                    &&&dictHeytingAlgebra),
                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                                                                                    dictHeytingAlgebra1))),
                                                                                                                                   add(string("ff"),
                                                                                                                                       &&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ff(),
                                                                                                                                                                                                                                                        &&&dictHeytingAlgebra),
                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ff(),
                                                                                                                                                                                                                                                        dictHeytingAlgebra1))),
                                                                                                                                       add(string("implies"),
                                                                                                                                           &&Func1::new({
                                                                                                                                                            let dictHeytingAlgebra1
                                                                                                                                                                =
                                                                                                                                                                dictHeytingAlgebra1.clone();
                                                                                                                                                            move
                                                                                                                                                                |v|
                                                                                                                                                                &Func1::new({
                                                                                                                                                                                let v
                                                                                                                                                                                    =
                                                                                                                                                                                    v.clone();
                                                                                                                                                                                move
                                                                                                                                                                                    |v1|
                                                                                                                                                                                    {
                                                                                                                                                                                        let matchValue:
                                                                                                                                                                                                LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                                                        let matchValue_1:
                                                                                                                                                                                                LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                            Sharpurs_Prelude::unbox(v1);
                                                                                                                                                                                        &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_implies(),
                                                                                                                                                                                                                                                                                                                                                                              &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                                                  PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                               PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_implies(),
                                                                                                                                                                                                                                                                                                                                                                              &&&dictHeytingAlgebra1),
                                                                                                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                                                                                  PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                               PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                                                                                           })))
                                                                                                                                                                                    }
                                                                                                                                                                            })
                                                                                                                                                        }),
                                                                                                                                           add(string("conj"),
                                                                                                                                               &&Func1::new({
                                                                                                                                                                let dictHeytingAlgebra1
                                                                                                                                                                    =
                                                                                                                                                                    dictHeytingAlgebra1.clone();
                                                                                                                                                                move
                                                                                                                                                                    |v_1|
                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                    let v_1
                                                                                                                                                                                        =
                                                                                                                                                                                        v_1.clone();
                                                                                                                                                                                    move
                                                                                                                                                                                        |v1_1|
                                                                                                                                                                                        {
                                                                                                                                                                                            let matchValue_3:
                                                                                                                                                                                                    LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                                            let matchValue_4:
                                                                                                                                                                                                    LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                                            &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                                                                                                  &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                               &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                      PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                                                                            &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                   PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                                                                                                  &&&dictHeytingAlgebra1),
                                                                                                                                                                                                                                                                                                                                               &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                                                                                      PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                                                                                            &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                   PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                               })))
                                                                                                                                                                                        }
                                                                                                                                                                                })
                                                                                                                                                            }),
                                                                                                                                               add(string("disj"),
                                                                                                                                                   &&Func1::new({
                                                                                                                                                                    let dictHeytingAlgebra1
                                                                                                                                                                        =
                                                                                                                                                                        dictHeytingAlgebra1.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |v_2|
                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                        let v_2
                                                                                                                                                                                            =
                                                                                                                                                                                            v_2.clone();
                                                                                                                                                                                        move
                                                                                                                                                                                            |v1_2|
                                                                                                                                                                                            {
                                                                                                                                                                                                let matchValue_6:
                                                                                                                                                                                                        LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(&&v_2);
                                                                                                                                                                                                let matchValue_7:
                                                                                                                                                                                                        LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(v1_2);
                                                                                                                                                                                                &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                                                                   &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                                                &&&match matchValue_7.as_ref()
                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                       PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                   }),
                                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                                                                                                                                                                                                      &&&dictHeytingAlgebra1),
                                                                                                                                                                                                                                                                                                                                                   &&&match matchValue_6.as_ref()
                                                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                                                                                                &&&match matchValue_7.as_ref()
                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                       PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                   })))
                                                                                                                                                                                            }
                                                                                                                                                                                    })
                                                                                                                                                                }),
                                                                                                                                                   add(string("not"),
                                                                                                                                                       &&Func1::new({
                                                                                                                                                                        let dictHeytingAlgebra1
                                                                                                                                                                            =
                                                                                                                                                                            dictHeytingAlgebra1.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |v_3|
                                                                                                                                                                            {
                                                                                                                                                                                let matchValue_9:
                                                                                                                                                                                        LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                    Sharpurs_Prelude::unbox(v_3);
                                                                                                                                                                                &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                                                                                   &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                                                &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                       PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                   }),
                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                                                                                                                                                   &&&dictHeytingAlgebra1),
                                                                                                                                                                                                                                                                                                &&&match matchValue_9.as_ref()
                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                       PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                   })))
                                                                                                                                                                            }
                                                                                                                                                                    }),
                                                                                                                                                       empty::<string,
                                                                                                                                                               &dyn Any>())))))))
                                                                                   })))
    }
    pub fn Data_Tuple_genericTuple() -> &dyn Any {
        static Data_Tuple_genericTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_genericTuple.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Genericusd_Dict(),
                                                                                 &&&add(string("to"),
                                                                                        &&Func1::new(move
                                                                                                         |x|
                                                                                                         {
                                                                                                             let matchValue:
                                                                                                                     LrcPtr<Data_Generic_Rep_Product> =
                                                                                                                 Sharpurs_Prelude::unbox(x);
                                                                                                             &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                 Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(x,
                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                 =>
                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                             },
                                                                                                                                                                                            &match matchValue.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                 Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(_,
                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                 =>
                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                             }))
                                                                                                         }),
                                                                                        add(string("from"),
                                                                                            &&Func1::new(move
                                                                                                             |x_1|
                                                                                                             {
                                                                                                                 let matchValue_1:
                                                                                                                         LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                     Sharpurs_Prelude::unbox(x_1);
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Constructor(),
                                                                                                                                                  &&&LrcPtr::new(Data_Generic_Rep_Product::Data_Generic_Rep_Productusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                                                             &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                    PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                }),
                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Generic_Rep::Data_Generic_Rep_Argument(),
                                                                                                                                                                                                                                                             &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                    PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                }))))
                                                                                                             }),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Tuple_functorTuple() -> &dyn Any {
        static Data_Tuple_functorTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_functorTuple.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                 &&&add(string("map"),
                                                                                        &&Func1::new(move
                                                                                                         |f|
                                                                                                         &Func1::new({
                                                                                                                         let f
                                                                                                                             =
                                                                                                                             f.clone();
                                                                                                                         move
                                                                                                                             |m|
                                                                                                                             {
                                                                                                                                 let matchValue:
                                                                                                                                         LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                     Sharpurs_Prelude::unbox(m);
                                                                                                                                 &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                     PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                 },
                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                        PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                    })))
                                                                                                                             }
                                                                                                                     })),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Data_Tuple_invariantTuple() -> &dyn Any {
        static Data_Tuple_invariantTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_invariantTuple.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                   &&&add(string("imap"),
                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                            &&&PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Tuple_fst() -> &dyn Any {
        static Data_Tuple_fst: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Tuple_fst.get_or_init(||
                                       &Func1::new(move |v|
                                                       &match Sharpurs_Prelude::unbox(v).as_ref()
                                                            {
                                                            PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                              _)
                                                            => x.clone(),
                                                        }))
    }
    pub fn Data_Tuple_lazyTuple() -> &dyn Any {
        static Data_Tuple_lazyTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_lazyTuple.get_or_init(||
                                             &Func1::new(move |dictLazy|
                                                             {
                                                                 let defer =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                      dictLazy);
                                                                 &Func1::new({
                                                                                 let defer
                                                                                     =
                                                                                     defer.clone();
                                                                                 move
                                                                                     |dictLazy1|
                                                                                     {
                                                                                         let defer1 =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_defer(),
                                                                                                                              dictLazy1);
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Lazy::Control_Lazy_Lazyusd_Dict(),
                                                                                                                          &&&add(string("defer"),
                                                                                                                                 &&Func1::new({
                                                                                                                                                  let defer1
                                                                                                                                                      =
                                                                                                                                                      defer1.clone();
                                                                                                                                                  move
                                                                                                                                                      |f|
                                                                                                                                                      &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                         &&&defer),
                                                                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                                                                        let f
                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                            f.clone();
                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                            |v|
                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_fst(),
                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                                                                                                                                                    })),
                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                                                                                                                                                         &&&defer1),
                                                                                                                                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                                                                                                                                        let f
                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                            f.clone();
                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                            |v_1|
                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_snd(),
                                                                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Unit::Data_Unit_unit()))
                                                                                                                                                                                                                                                                                    }))))
                                                                                                                                              }),
                                                                                                                                 empty::<string,
                                                                                                                                         &dyn Any>()))
                                                                                     }
                                                                             })
                                                             }))
    }
    pub fn Data_Tuple_extendTuple() -> &dyn Any {
        static Data_Tuple_extendTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_extendTuple.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                &&&add(string("extend"),
                                                                                       &&Func1::new(move
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
                                                                                                                                let matchValue_1:
                                                                                                                                        LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                                &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                    {
                                                                                                                                                                                                                    PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                },
                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                &&&matchValue_1)))
                                                                                                                            }
                                                                                                                    })),
                                                                                       add(string("Functor0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Tuple_eqTuple() -> &dyn Any {
        static Data_Tuple_eqTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_eqTuple.get_or_init(||
                                           &Func1::new(move |dictEq|
                                                           &Func1::new({
                                                                           let dictEq
                                                                               =
                                                                               dictEq.clone();
                                                                           move
                                                                               |dictEq1|
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                &&&add(string("eq"),
                                                                                                                       &&Func1::new({
                                                                                                                                        let dictEq1
                                                                                                                                            =
                                                                                                                                            dictEq1.clone();
                                                                                                                                        move
                                                                                                                                            |x|
                                                                                                                                            &Func1::new({
                                                                                                                                                            let x
                                                                                                                                                                =
                                                                                                                                                                x.clone();
                                                                                                                                                            move
                                                                                                                                                                |y|
                                                                                                                                                                {
                                                                                                                                                                    let matchValue:
                                                                                                                                                                            LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                        Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                    let matchValue_1:
                                                                                                                                                                            LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                        Sharpurs_Prelude::unbox(y);
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                           &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                 &&&dictEq),
                                                                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                     PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                  PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                              })),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                              &&&dictEq1),
                                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                  PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                               PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                           }))
                                                                                                                                                                }
                                                                                                                                                        })
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>()))
                                                                       })))
    }
    pub fn Data_Tuple_ordTuple() -> &dyn Any {
        static Data_Tuple_ordTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_ordTuple.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            {
                                                                let eqTuple1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_eqTuple(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                               Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                &Func1::new({
                                                                                let dictOrd
                                                                                    =
                                                                                    dictOrd.clone();
                                                                                let eqTuple1
                                                                                    =
                                                                                    eqTuple1.clone();
                                                                                move
                                                                                    |dictOrd1|
                                                                                    {
                                                                                        let eqTuple2 =
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&eqTuple1,
                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                       Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                         &&&add(string("compare"),
                                                                                                                                &&Func1::new({
                                                                                                                                                 let dictOrd1
                                                                                                                                                     =
                                                                                                                                                     dictOrd1.clone();
                                                                                                                                                 move
                                                                                                                                                     |x|
                                                                                                                                                     &Func1::new({
                                                                                                                                                                     let x
                                                                                                                                                                         =
                                                                                                                                                                         x.clone();
                                                                                                                                                                     move
                                                                                                                                                                         |y|
                                                                                                                                                                         {
                                                                                                                                                                             let matchValue:
                                                                                                                                                                                     LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                             let matchValue_1:
                                                                                                                                                                                     LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(y);
                                                                                                                                                                             let matchValue_3:
                                                                                                                                                                                     LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                 Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                  &&&dictOrd),
                                                                                                                                                                                                                                                                               &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                      PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                            &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                   PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                               }));
                                                                                                                                                                             match matchValue_3.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                 Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                 =>
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                                                                                                 Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                                                                 =>
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                                                                                                                 _
                                                                                                                                                                                 =>
                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                        &&&dictOrd1),
                                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                            PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                         PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                     }),
                                                                                                                                                                             }
                                                                                                                                                                         }
                                                                                                                                                                 })
                                                                                                                                             }),
                                                                                                                                add(string("Eq0"),
                                                                                                                                    &&Func1::new({
                                                                                                                                                     let eqTuple2
                                                                                                                                                         =
                                                                                                                                                         eqTuple2.clone();
                                                                                                                                                     move
                                                                                                                                                         |usd__unused|
                                                                                                                                                         &eqTuple2
                                                                                                                                                 }),
                                                                                                                                    empty::<string,
                                                                                                                                            &dyn Any>())))
                                                                                    }
                                                                            })
                                                            }))
    }
    pub fn Data_Tuple_eq1Tuple() -> &dyn Any {
        static Data_Tuple_eq1Tuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_eq1Tuple.get_or_init(||
                                            &Func1::new(move |dictEq|
                                                            {
                                                                let eqTuple1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_eqTuple(),
                                                                                                     dictEq);
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                 &&&add(string("eq1"),
                                                                                                        &&Func1::new({
                                                                                                                         let eqTuple1
                                                                                                                             =
                                                                                                                             eqTuple1.clone();
                                                                                                                         move
                                                                                                                             |dictEq1|
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&eqTuple1,
                                                                                                                                                                                                 dictEq1))
                                                                                                                     }),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))
                                                            }))
    }
    pub fn Data_Tuple_ord1Tuple() -> &dyn Any {
        static Data_Tuple_ord1Tuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_ord1Tuple.get_or_init(||
                                             &Func1::new(move |dictOrd|
                                                             {
                                                                 let ordTuple1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_ordTuple(),
                                                                                                      dictOrd);
                                                                 let eq1Tuple1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_eq1Tuple(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                  &&&add(string("compare1"),
                                                                                                         &&Func1::new({
                                                                                                                          let ordTuple1
                                                                                                                              =
                                                                                                                              ordTuple1.clone();
                                                                                                                          move
                                                                                                                              |dictOrd1|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&ordTuple1,
                                                                                                                                                                                                  dictOrd1))
                                                                                                                      }),
                                                                                                         add(string("Eq10"),
                                                                                                             &&Func1::new({
                                                                                                                              let eq1Tuple1
                                                                                                                                  =
                                                                                                                                  eq1Tuple1.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &eq1Tuple1
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Data_Tuple_curry() -> &dyn Any {
        static Data_Tuple_curry: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_curry.get_or_init(||
                                         &Func1::new(move |f|
                                                         &Func1::new({
                                                                         let f
                                                                             =
                                                                             f.clone();
                                                                         move
                                                                             |a|
                                                                             &Func1::new({
                                                                                             let a
                                                                                                 =
                                                                                                 a.clone();
                                                                                             move
                                                                                                 |b|
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                  &&&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(&a,
                                                                                                                                                                                                                   b.clone())))
                                                                                         })
                                                                     })))
    }
    pub fn Data_Tuple_comonadTuple() -> &dyn Any {
        static Data_Tuple_comonadTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_comonadTuple.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                                 &&&add(string("extract"),
                                                                                        &&PureScript_Data_Tuple::Data_Tuple_snd(),
                                                                                        add(string("Extend0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &PureScript_Data_Tuple::Data_Tuple_extendTuple()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Tuple_commutativeRingTuple() -> &dyn Any {
        static Data_Tuple_commutativeRingTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_commutativeRingTuple.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictCommutativeRing|
                                                                        {
                                                                            let ringTuple1 =
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_ringTuple(),
                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                           Sharpurs_Prelude::unbox(dictCommutativeRing)),
                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                            &Func1::new({
                                                                                            let ringTuple1
                                                                                                =
                                                                                                ringTuple1.clone();
                                                                                            move
                                                                                                |dictCommutativeRing1|
                                                                                                {
                                                                                                    let ringTuple2 =
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&ringTuple1,
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                                                   Sharpurs_Prelude::unbox(dictCommutativeRing1)),
                                                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                                                                     &&&add(string("Ring0"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let ringTuple2
                                                                                                                                                                 =
                                                                                                                                                                 ringTuple2.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused|
                                                                                                                                                                 &ringTuple2
                                                                                                                                                         }),
                                                                                                                                            empty::<string,
                                                                                                                                                    &dyn Any>()))
                                                                                                }
                                                                                        })
                                                                        }))
    }
    pub fn Data_Tuple_boundedTuple() -> &dyn Any {
        static Data_Tuple_boundedTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_boundedTuple.get_or_init(||
                                                &Func1::new(move |dictBounded|
                                                                {
                                                                    let ordTuple1 =
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_ordTuple(),
                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                   Sharpurs_Prelude::unbox(dictBounded)),
                                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined()));
                                                                    &Func1::new({
                                                                                    let dictBounded
                                                                                        =
                                                                                        dictBounded.clone();
                                                                                    let ordTuple1
                                                                                        =
                                                                                        ordTuple1.clone();
                                                                                    move
                                                                                        |dictBounded1|
                                                                                        {
                                                                                            let ordTuple2 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&ordTuple1,
                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                                           Sharpurs_Prelude::unbox(dictBounded1)),
                                                                                                                                                                    &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                                                             &&&add(string("top"),
                                                                                                                                    &&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                                                     &&&dictBounded),
                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                                                     dictBounded1))),
                                                                                                                                    add(string("bottom"),
                                                                                                                                        &&LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                                                         &&&dictBounded),
                                                                                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                                                         dictBounded1))),
                                                                                                                                        add(string("Ord0"),
                                                                                                                                            &&Func1::new({
                                                                                                                                                             let ordTuple2
                                                                                                                                                                 =
                                                                                                                                                                 ordTuple2.clone();
                                                                                                                                                             move
                                                                                                                                                                 |usd__unused|
                                                                                                                                                                 &ordTuple2
                                                                                                                                                         }),
                                                                                                                                            empty::<string,
                                                                                                                                                    &dyn Any>()))))
                                                                                        }
                                                                                })
                                                                }))
    }
    pub fn Data_Tuple_booleanAlgebraTuple() -> &dyn Any {
        static Data_Tuple_booleanAlgebraTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_booleanAlgebraTuple.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictBooleanAlgebra|
                                                                       {
                                                                           let heytingAlgebraTuple1 =
                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_heytingAlgebraTuple(),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebra0"),
                                                                                                                                                          Sharpurs_Prelude::unbox(dictBooleanAlgebra)),
                                                                                                                                                   &&&Sharpurs_Prelude::Prim_undefined()));
                                                                           &Func1::new({
                                                                                           let heytingAlgebraTuple1
                                                                                               =
                                                                                               heytingAlgebraTuple1.clone();
                                                                                           move
                                                                                               |dictBooleanAlgebra1|
                                                                                               {
                                                                                                   let heytingAlgebraTuple2 =
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&heytingAlgebraTuple1,
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("HeytingAlgebra0"),
                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictBooleanAlgebra1)),
                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_BooleanAlgebrausd_Dict(),
                                                                                                                                    &&&add(string("HeytingAlgebra0"),
                                                                                                                                           &&Func1::new({
                                                                                                                                                            let heytingAlgebraTuple2
                                                                                                                                                                =
                                                                                                                                                                heytingAlgebraTuple2.clone();
                                                                                                                                                            move
                                                                                                                                                                |usd__unused|
                                                                                                                                                                &heytingAlgebraTuple2
                                                                                                                                                        }),
                                                                                                                                           empty::<string,
                                                                                                                                                   &dyn Any>()))
                                                                                               }
                                                                                       })
                                                                       }))
    }
    pub fn Data_Tuple_applyTuple() -> &dyn Any {
        static Data_Tuple_applyTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_applyTuple.get_or_init(||
                                              &Func1::new(move |dictSemigroup|
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                               &&&add(string("apply"),
                                                                                                      &&Func1::new({
                                                                                                                       let dictSemigroup
                                                                                                                           =
                                                                                                                           dictSemigroup.clone();
                                                                                                                       move
                                                                                                                           |v|
                                                                                                                           &Func1::new({
                                                                                                                                           let v
                                                                                                                                               =
                                                                                                                                               v.clone();
                                                                                                                                           move
                                                                                                                                               |v1|
                                                                                                                                               {
                                                                                                                                                   let matchValue:
                                                                                                                                                           LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                   let matchValue_1:
                                                                                                                                                           LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                   &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                         &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                                                             PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                                                   &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                      }),
                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                      },
                                                                                                                                                                                                                                                                   &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                      })))
                                                                                                                                               }
                                                                                                                                       })
                                                                                                                   }),
                                                                                                      add(string("Functor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Tuple::Data_Tuple_functorTuple()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))))
    }
    pub fn Data_Tuple_bindTuple() -> &dyn Any {
        static Data_Tuple_bindTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_bindTuple.get_or_init(||
                                             &Func1::new(move |dictSemigroup|
                                                             {
                                                                 let applyTuple1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_applyTuple(),
                                                                                                      dictSemigroup);
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                  &&&add(string("bind"),
                                                                                                         &&Func1::new({
                                                                                                                          let dictSemigroup
                                                                                                                              =
                                                                                                                              dictSemigroup.clone();
                                                                                                                          move
                                                                                                                              |v|
                                                                                                                              &Func1::new({
                                                                                                                                              let v
                                                                                                                                                  =
                                                                                                                                                  v.clone();
                                                                                                                                              move
                                                                                                                                                  |f|
                                                                                                                                                  {
                                                                                                                                                      let matchValue:
                                                                                                                                                              LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                      let matchValue_3:
                                                                                                                                                              LrcPtr<PureScript_Data_Tuple::Data_Tuple_Tuple> =
                                                                                                                                                          Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(f),
                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                            {
                                                                                                                                                                                                                            PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                        }));
                                                                                                                                                      &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                            &&&dictSemigroup),
                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                  _)
                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                      &&&match matchValue_3.as_ref()
                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                             PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                         }),
                                                                                                                                                                                                                                     &match matchValue_3.as_ref()
                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                          PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(_,
                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                      }))
                                                                                                                                                  }
                                                                                                                                          })
                                                                                                                      }),
                                                                                                         add(string("Apply0"),
                                                                                                             &&Func1::new({
                                                                                                                              let applyTuple1
                                                                                                                                  =
                                                                                                                                  applyTuple1.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &applyTuple1
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Data_Tuple_applicativeTuple() -> &dyn Any {
        static Data_Tuple_applicativeTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_applicativeTuple.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictMonoid|
                                                                    {
                                                                        let applyTuple1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_applyTuple(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                         &&&add(string("pure"),
                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                    |usd__arg1|
                                                                                                                                                                    Func1::new({
                                                                                                                                                                                   let usd__arg1
                                                                                                                                                                                       =
                                                                                                                                                                                       usd__arg1.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |usd__arg2|
                                                                                                                                                                                       &LrcPtr::new(PureScript_Data_Tuple::Data_Tuple_Tuple::Data_Tuple_Tupleusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                      usd__arg2.clone()))
                                                                                                                                                                               })),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                     dictMonoid)),
                                                                                                                add(string("Apply0"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let applyTuple1
                                                                                                                                         =
                                                                                                                                         applyTuple1.clone();
                                                                                                                                     move
                                                                                                                                         |usd__unused|
                                                                                                                                         &applyTuple1
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>())))
                                                                    }))
    }
    pub fn Data_Tuple_monadTuple() -> &dyn Any {
        static Data_Tuple_monadTuple: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Tuple_monadTuple.get_or_init(||
                                              &Func1::new(move |dictMonoid|
                                                              {
                                                                  let applicativeTuple1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_applicativeTuple(),
                                                                                                       dictMonoid);
                                                                  let bindTuple1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Tuple::Data_Tuple_bindTuple(),
                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                 Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                   &&&add(string("Applicative0"),
                                                                                                          &&Func1::new({
                                                                                                                           let applicativeTuple1
                                                                                                                               =
                                                                                                                               applicativeTuple1.clone();
                                                                                                                           move
                                                                                                                               |usd__unused|
                                                                                                                               &applicativeTuple1
                                                                                                                       }),
                                                                                                          add(string("Bind1"),
                                                                                                              &&Func1::new({
                                                                                                                               let bindTuple1
                                                                                                                                   =
                                                                                                                                   bindTuple1.clone();
                                                                                                                               move
                                                                                                                                   |usd__unused_1|
                                                                                                                                   &bindTuple1
                                                                                                                           }),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
                                                              }))
    }
}

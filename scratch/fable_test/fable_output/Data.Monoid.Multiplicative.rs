pub mod PureScript_Data_Monoid_Multiplicative {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Monoid_Multiplicative_Multiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_Multiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_Multiplicative.get_or_init(||
                                                                  &Func1::new(move
                                                                                  |x|
                                                                                  x.clone()))
    }
    pub fn Data_Monoid_Multiplicative_showMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_showMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_showMultiplicative.get_or_init(||
                                                                      &Func1::new(move
                                                                                      |dictShow|
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                       &&&add(string("show"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let dictShow
                                                                                                                                                   =
                                                                                                                                                   dictShow.clone();
                                                                                                                                               move
                                                                                                                                                   |v|
                                                                                                                                                   {
                                                                                                                                                       let a =
                                                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                           &&&string("(Multiplicative ")),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                    &&&dictShow),
                                                                                                                                                                                                                                                                                                 &&&a)),
                                                                                                                                                                                                                           &&&string(")")))
                                                                                                                                                   }
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Monoid_Multiplicative_semigroupMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_semigroupMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_semigroupMultiplicative.get_or_init(||
                                                                           &Func1::new(move
                                                                                           |dictSemiring|
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                            &&&add(string("append"),
                                                                                                                                   &&Func1::new({
                                                                                                                                                    let dictSemiring
                                                                                                                                                        =
                                                                                                                                                        dictSemiring.clone();
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
                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative(),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                          &&&dictSemiring),
                                                                                                                                                                                                                                                                                       &&&matchValue),
                                                                                                                                                                                                                                                    &&&matchValue_1))
                                                                                                                                                                            }
                                                                                                                                                                    })
                                                                                                                                                }),
                                                                                                                                   empty::<string,
                                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Monoid_Multiplicative_ordMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_ordMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_ordMultiplicative.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |dictOrd|
                                                                                     dictOrd.clone()))
    }
    pub fn Data_Monoid_Multiplicative_monoidMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_monoidMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_monoidMultiplicative.get_or_init(||
                                                                        &Func1::new(move
                                                                                        |dictSemiring|
                                                                                        {
                                                                                            let semigroupMultiplicative1 =
                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_semigroupMultiplicative(),
                                                                                                                                 dictSemiring);
                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                             &&&add(string("mempty"),
                                                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative(),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_one(),
                                                                                                                                                                                                         dictSemiring)),
                                                                                                                                    add(string("Semigroup0"),
                                                                                                                                        &&Func1::new({
                                                                                                                                                         let semigroupMultiplicative1
                                                                                                                                                             =
                                                                                                                                                             semigroupMultiplicative1.clone();
                                                                                                                                                         move
                                                                                                                                                             |usd__unused|
                                                                                                                                                             &semigroupMultiplicative1
                                                                                                                                                     }),
                                                                                                                                        empty::<string,
                                                                                                                                                &dyn Any>())))
                                                                                        }))
    }
    pub fn Data_Monoid_Multiplicative_functorMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_functorMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_functorMultiplicative.get_or_init(||
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
                                                                                                                                                          let v =
                                                                                                                                                              Sharpurs_Prelude::unbox(m);
                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative(),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                              &&&v))
                                                                                                                                                      }
                                                                                                                                              })),
                                                                                                                 empty::<string,
                                                                                                                         &dyn Any>())))
    }
    pub fn Data_Monoid_Multiplicative_eqMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_eqMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_eqMultiplicative.get_or_init(||
                                                                    &Func1::new(move
                                                                                    |dictEq|
                                                                                    dictEq.clone()))
    }
    pub fn Data_Monoid_Multiplicative_eq1Multiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_eq1Multiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_eq1Multiplicative.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                      &&&add(string("eq1"),
                                                                                                             &&Func1::new(move
                                                                                                                              |dictEq|
                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_eqMultiplicative(),
                                                                                                                                                                                                  dictEq))),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Data_Monoid_Multiplicative_ord1Multiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_ord1Multiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_ord1Multiplicative.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                       &&&add(string("compare1"),
                                                                                                              &&Func1::new(move
                                                                                                                               |dictOrd|
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_ordMultiplicative(),
                                                                                                                                                                                                   dictOrd))),
                                                                                                              add(string("Eq10"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |usd__unused|
                                                                                                                                   &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_eq1Multiplicative()),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Monoid_Multiplicative_boundedMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_boundedMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_boundedMultiplicative.get_or_init(||
                                                                         &Func1::new(move
                                                                                         |dictBounded|
                                                                                         dictBounded.clone()))
    }
    pub fn Data_Monoid_Multiplicative_applyMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_applyMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_applyMultiplicative.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                                        &&&add(string("apply"),
                                                                                                               &&Func1::new(move
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
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative(),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                            &&&matchValue_1))
                                                                                                                                                    }
                                                                                                                                            })),
                                                                                                               add(string("Functor0"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused|
                                                                                                                                    &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_functorMultiplicative()),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Monoid_Multiplicative_bindMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_bindMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_bindMultiplicative.get_or_init(||
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                                       &&&add(string("bind"),
                                                                                                              &&Func1::new(move
                                                                                                                               |v|
                                                                                                                               &Func1::new({
                                                                                                                                               let v
                                                                                                                                                   =
                                                                                                                                                   v.clone();
                                                                                                                                               move
                                                                                                                                                   |f|
                                                                                                                                                   {
                                                                                                                                                       let matchValue =
                                                                                                                                                           Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(f),
                                                                                                                                                                                        &&&matchValue)
                                                                                                                                                   }
                                                                                                                                           })),
                                                                                                              add(string("Apply0"),
                                                                                                                  &&Func1::new(move
                                                                                                                                   |usd__unused|
                                                                                                                                   &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_applyMultiplicative()),
                                                                                                                  empty::<string,
                                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Monoid_Multiplicative_applicativeMultiplicative()
     -> &dyn Any {
        static Data_Monoid_Multiplicative_applicativeMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_applicativeMultiplicative.get_or_init(||
                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                              &&&add(string("pure"),
                                                                                                                     &&PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_Multiplicative(),
                                                                                                                     add(string("Apply0"),
                                                                                                                         &&Func1::new(move
                                                                                                                                          |usd__unused|
                                                                                                                                          &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_applyMultiplicative()),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Monoid_Multiplicative_monadMultiplicative() -> &dyn Any {
        static Data_Monoid_Multiplicative_monadMultiplicative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Multiplicative_monadMultiplicative.get_or_init(||
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                                        &&&add(string("Applicative0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_applicativeMultiplicative()),
                                                                                                               add(string("Bind1"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused_1|
                                                                                                                                    &PureScript_Data_Monoid_Multiplicative::Data_Monoid_Multiplicative_bindMultiplicative()),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))
    }
}

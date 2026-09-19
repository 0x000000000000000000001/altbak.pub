pub mod PureScript_Data_Monoid_Additive {
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
    pub fn Data_Monoid_Additive_Additive() -> &dyn Any {
        static Data_Monoid_Additive_Additive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_Additive.get_or_init(||
                                                      &Func1::new(move |x|
                                                                      x.clone()))
    }
    pub fn Data_Monoid_Additive_showAdditive() -> &dyn Any {
        static Data_Monoid_Additive_showAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_showAdditive.get_or_init(||
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
                                                                                                                                                                                                               &&&string("(Additive ")),
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
    pub fn Data_Monoid_Additive_semigroupAdditive() -> &dyn Any {
        static Data_Monoid_Additive_semigroupAdditive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_semigroupAdditive.get_or_init(||
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
                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive(),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                              &&&dictSemiring),
                                                                                                                                                                                                                                                                           &&&matchValue),
                                                                                                                                                                                                                                        &&&matchValue_1))
                                                                                                                                                                }
                                                                                                                                                        })
                                                                                                                                    }),
                                                                                                                       empty::<string,
                                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Monoid_Additive_ordAdditive() -> &dyn Any {
        static Data_Monoid_Additive_ordAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_ordAdditive.get_or_init(||
                                                         &Func1::new(move
                                                                         |dictOrd|
                                                                         dictOrd.clone()))
    }
    pub fn Data_Monoid_Additive_monoidAdditive() -> &dyn Any {
        static Data_Monoid_Additive_monoidAdditive: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Additive_monoidAdditive.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictSemiring|
                                                                            {
                                                                                let semigroupAdditive1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_semigroupAdditive(),
                                                                                                                     dictSemiring);
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                 &&&add(string("mempty"),
                                                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive(),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                                                                                             dictSemiring)),
                                                                                                                        add(string("Semigroup0"),
                                                                                                                            &&Func1::new({
                                                                                                                                             let semigroupAdditive1
                                                                                                                                                 =
                                                                                                                                                 semigroupAdditive1.clone();
                                                                                                                                             move
                                                                                                                                                 |usd__unused|
                                                                                                                                                 &semigroupAdditive1
                                                                                                                                         }),
                                                                                                                            empty::<string,
                                                                                                                                    &dyn Any>())))
                                                                            }))
    }
    pub fn Data_Monoid_Additive_functorAdditive() -> &dyn Any {
        static Data_Monoid_Additive_functorAdditive: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Additive_functorAdditive.get_or_init(||
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
                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive(),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                  &&&v))
                                                                                                                                          }
                                                                                                                                  })),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>())))
    }
    pub fn Data_Monoid_Additive_eqAdditive() -> &dyn Any {
        static Data_Monoid_Additive_eqAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_eqAdditive.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictEq|
                                                                        dictEq.clone()))
    }
    pub fn Data_Monoid_Additive_eq1Additive() -> &dyn Any {
        static Data_Monoid_Additive_eq1Additive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_eq1Additive.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                          &&&add(string("eq1"),
                                                                                                 &&Func1::new(move
                                                                                                                  |dictEq|
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_eqAdditive(),
                                                                                                                                                                                      dictEq))),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>())))
    }
    pub fn Data_Monoid_Additive_ord1Additive() -> &dyn Any {
        static Data_Monoid_Additive_ord1Additive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_ord1Additive.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                           &&&add(string("compare1"),
                                                                                                  &&Func1::new(move
                                                                                                                   |dictOrd|
                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_ordAdditive(),
                                                                                                                                                                                       dictOrd))),
                                                                                                  add(string("Eq10"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_eq1Additive()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Monoid_Additive_boundedAdditive() -> &dyn Any {
        static Data_Monoid_Additive_boundedAdditive: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Additive_boundedAdditive.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictBounded|
                                                                             dictBounded.clone()))
    }
    pub fn Data_Monoid_Additive_applyAdditive() -> &dyn Any {
        static Data_Monoid_Additive_applyAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_applyAdditive.get_or_init(||
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
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive(),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                &&&matchValue_1))
                                                                                                                                        }
                                                                                                                                })),
                                                                                                   add(string("Functor0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_functorAdditive()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Monoid_Additive_bindAdditive() -> &dyn Any {
        static Data_Monoid_Additive_bindAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_bindAdditive.get_or_init(||
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
                                                                                                                       &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_applyAdditive()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Monoid_Additive_applicativeAdditive() -> &dyn Any {
        static Data_Monoid_Additive_applicativeAdditive:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_applicativeAdditive.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                  &&&add(string("pure"),
                                                                                                         &&PureScript_Data_Monoid_Additive::Data_Monoid_Additive_Additive(),
                                                                                                         add(string("Apply0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_applyAdditive()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Monoid_Additive_monadAdditive() -> &dyn Any {
        static Data_Monoid_Additive_monadAdditive: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Additive_monadAdditive.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                            &&&add(string("Applicative0"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_applicativeAdditive()),
                                                                                                   add(string("Bind1"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused_1|
                                                                                                                        &PureScript_Data_Monoid_Additive::Data_Monoid_Additive_bindAdditive()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
}

pub mod PureScript_Data_Monoid_Conj {
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
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Monoid_Conj_Conj() -> &dyn Any {
        static Data_Monoid_Conj_Conj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_Conj.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Monoid_Conj_showConj() -> &dyn Any {
        static Data_Monoid_Conj_showConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_showConj.get_or_init(||
                                                  &Func1::new(move |dictShow|
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
                                                                                                                                                                                                       &&&string("(Conj ")),
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
    pub fn Data_Monoid_Conj_semiringConj() -> &dyn Any {
        static Data_Monoid_Conj_semiringConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_semiringConj.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictHeytingAlgebra|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                                       &&&add(string("zero"),
                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                   dictHeytingAlgebra)),
                                                                                                              add(string("one"),
                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_ff(),
                                                                                                                                                                                       dictHeytingAlgebra)),
                                                                                                                  add(string("add"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let dictHeytingAlgebra
                                                                                                                                           =
                                                                                                                                           dictHeytingAlgebra.clone();
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
                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                             &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                                       &&&matchValue_1))
                                                                                                                                                               }
                                                                                                                                                       })
                                                                                                                                   }),
                                                                                                                      add(string("mul"),
                                                                                                                          &&Func1::new({
                                                                                                                                           let dictHeytingAlgebra
                                                                                                                                               =
                                                                                                                                               dictHeytingAlgebra.clone();
                                                                                                                                           move
                                                                                                                                               |v_1|
                                                                                                                                               &Func1::new({
                                                                                                                                                               let v_1
                                                                                                                                                                   =
                                                                                                                                                                   v_1.clone();
                                                                                                                                                               move
                                                                                                                                                                   |v1_1|
                                                                                                                                                                   {
                                                                                                                                                                       let matchValue_3 =
                                                                                                                                                                           Sharpurs_Prelude::unbox(&&v_1);
                                                                                                                                                                       let matchValue_4 =
                                                                                                                                                                           Sharpurs_Prelude::unbox(v1_1);
                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                                                                                                                                 &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                              &&&matchValue_3),
                                                                                                                                                                                                                                           &&&matchValue_4))
                                                                                                                                                                   }
                                                                                                                                                           })
                                                                                                                                       }),
                                                                                                                          empty::<string,
                                                                                                                                  &dyn Any>())))))))
    }
    pub fn Data_Monoid_Conj_semigroupConj() -> &dyn Any {
        static Data_Monoid_Conj_semigroupConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_semigroupConj.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictHeytingAlgebra|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                        &&&add(string("append"),
                                                                                                               &&Func1::new({
                                                                                                                                let dictHeytingAlgebra
                                                                                                                                    =
                                                                                                                                    dictHeytingAlgebra.clone();
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
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                      &&&dictHeytingAlgebra),
                                                                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                                                                &&&matchValue_1))
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                            }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Monoid_Conj_ordConj() -> &dyn Any {
        static Data_Monoid_Conj_ordConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_ordConj.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 dictOrd.clone()))
    }
    pub fn Data_Monoid_Conj_monoidConj() -> &dyn Any {
        static Data_Monoid_Conj_monoidConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_monoidConj.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictHeytingAlgebra|
                                                                    {
                                                                        let semigroupConj1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_semigroupConj(),
                                                                                                             dictHeytingAlgebra);
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                         &&&add(string("mempty"),
                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_tt(),
                                                                                                                                                                                     dictHeytingAlgebra)),
                                                                                                                add(string("Semigroup0"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let semigroupConj1
                                                                                                                                         =
                                                                                                                                         semigroupConj1.clone();
                                                                                                                                     move
                                                                                                                                         |usd__unused|
                                                                                                                                         &semigroupConj1
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>())))
                                                                    }))
    }
    pub fn Data_Monoid_Conj_functorConj() -> &dyn Any {
        static Data_Monoid_Conj_functorConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_functorConj.get_or_init(||
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
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                          &&&v))
                                                                                                                                  }
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Monoid_Conj_eqConj() -> &dyn Any {
        static Data_Monoid_Conj_eqConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_eqConj.get_or_init(||
                                                &Func1::new(move |dictEq|
                                                                dictEq.clone()))
    }
    pub fn Data_Monoid_Conj_eq1Conj() -> &dyn Any {
        static Data_Monoid_Conj_eq1Conj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_eq1Conj.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                  &&&add(string("eq1"),
                                                                                         &&Func1::new(move
                                                                                                          |dictEq|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_eqConj(),
                                                                                                                                                                              dictEq))),
                                                                                         empty::<string,
                                                                                                 &dyn Any>())))
    }
    pub fn Data_Monoid_Conj_ord1Conj() -> &dyn Any {
        static Data_Monoid_Conj_ord1Conj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_ord1Conj.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                   &&&add(string("compare1"),
                                                                                          &&Func1::new(move
                                                                                                           |dictOrd|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_ordConj(),
                                                                                                                                                                               dictOrd))),
                                                                                          add(string("Eq10"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_eq1Conj()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Monoid_Conj_boundedConj() -> &dyn Any {
        static Data_Monoid_Conj_boundedConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_boundedConj.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBounded|
                                                                     dictBounded.clone()))
    }
    pub fn Data_Monoid_Conj_applyConj() -> &dyn Any {
        static Data_Monoid_Conj_applyConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_applyConj.get_or_init(||
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
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                        &&&matchValue_1))
                                                                                                                                }
                                                                                                                        })),
                                                                                           add(string("Functor0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_functorConj()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Monoid_Conj_bindConj() -> &dyn Any {
        static Data_Monoid_Conj_bindConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_bindConj.get_or_init(||
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
                                                                                                               &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_applyConj()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Monoid_Conj_applicativeConj() -> &dyn Any {
        static Data_Monoid_Conj_applicativeConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_applicativeConj.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                          &&&add(string("pure"),
                                                                                                 &&PureScript_Data_Monoid_Conj::Data_Monoid_Conj_Conj(),
                                                                                                 add(string("Apply0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_applyConj()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Monoid_Conj_monadConj() -> &dyn Any {
        static Data_Monoid_Conj_monadConj: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Conj_monadConj.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                    &&&add(string("Applicative0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_applicativeConj()),
                                                                                           add(string("Bind1"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused_1|
                                                                                                                &PureScript_Data_Monoid_Conj::Data_Monoid_Conj_bindConj()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
}

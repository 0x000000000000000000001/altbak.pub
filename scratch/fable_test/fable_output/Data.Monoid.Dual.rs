pub mod PureScript_Data_Monoid_Dual {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
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
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Monoid_Dual_Dual() -> &dyn Any {
        static Data_Monoid_Dual_Dual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_Dual.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Monoid_Dual_showDual() -> &dyn Any {
        static Data_Monoid_Dual_showDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_showDual.get_or_init(||
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
                                                                                                                                                                                                       &&&string("(Dual ")),
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
    pub fn Data_Monoid_Dual_semigroupDual() -> &dyn Any {
        static Data_Monoid_Dual_semigroupDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_semigroupDual.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictSemigroup|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                        &&&add(string("append"),
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
                                                                                                                                                            let matchValue =
                                                                                                                                                                Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                            let matchValue_1 =
                                                                                                                                                                Sharpurs_Prelude::unbox(v1);
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual(),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                      &&&dictSemigroup),
                                                                                                                                                                                                                                                                   &&&matchValue_1),
                                                                                                                                                                                                                                &&&matchValue))
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                            }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Monoid_Dual_ordDual() -> &dyn Any {
        static Data_Monoid_Dual_ordDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_ordDual.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 dictOrd.clone()))
    }
    pub fn Data_Monoid_Dual_monoidDual() -> &dyn Any {
        static Data_Monoid_Dual_monoidDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_monoidDual.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictMonoid|
                                                                    {
                                                                        let semigroupDual1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_semigroupDual(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                         &&&add(string("mempty"),
                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual(),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                     dictMonoid)),
                                                                                                                add(string("Semigroup0"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let semigroupDual1
                                                                                                                                         =
                                                                                                                                         semigroupDual1.clone();
                                                                                                                                     move
                                                                                                                                         |usd__unused|
                                                                                                                                         &semigroupDual1
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>())))
                                                                    }))
    }
    pub fn Data_Monoid_Dual_functorDual() -> &dyn Any {
        static Data_Monoid_Dual_functorDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_functorDual.get_or_init(||
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
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual(),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                          &&&v))
                                                                                                                                  }
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Monoid_Dual_eqDual() -> &dyn Any {
        static Data_Monoid_Dual_eqDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_eqDual.get_or_init(||
                                                &Func1::new(move |dictEq|
                                                                dictEq.clone()))
    }
    pub fn Data_Monoid_Dual_eq1Dual() -> &dyn Any {
        static Data_Monoid_Dual_eq1Dual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_eq1Dual.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                  &&&add(string("eq1"),
                                                                                         &&Func1::new(move
                                                                                                          |dictEq|
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_eqDual(),
                                                                                                                                                                              dictEq))),
                                                                                         empty::<string,
                                                                                                 &dyn Any>())))
    }
    pub fn Data_Monoid_Dual_ord1Dual() -> &dyn Any {
        static Data_Monoid_Dual_ord1Dual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_ord1Dual.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                   &&&add(string("compare1"),
                                                                                          &&Func1::new(move
                                                                                                           |dictOrd|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_ordDual(),
                                                                                                                                                                               dictOrd))),
                                                                                          add(string("Eq10"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_eq1Dual()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Monoid_Dual_boundedDual() -> &dyn Any {
        static Data_Monoid_Dual_boundedDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_boundedDual.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBounded|
                                                                     dictBounded.clone()))
    }
    pub fn Data_Monoid_Dual_applyDual() -> &dyn Any {
        static Data_Monoid_Dual_applyDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_applyDual.get_or_init(||
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
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual(),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                        &&&matchValue_1))
                                                                                                                                }
                                                                                                                        })),
                                                                                           add(string("Functor0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_functorDual()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Monoid_Dual_bindDual() -> &dyn Any {
        static Data_Monoid_Dual_bindDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_bindDual.get_or_init(||
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
                                                                                                               &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_applyDual()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Monoid_Dual_applicativeDual() -> &dyn Any {
        static Data_Monoid_Dual_applicativeDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_applicativeDual.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                          &&&add(string("pure"),
                                                                                                 &&PureScript_Data_Monoid_Dual::Data_Monoid_Dual_Dual(),
                                                                                                 add(string("Apply0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_applyDual()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Monoid_Dual_monadDual() -> &dyn Any {
        static Data_Monoid_Dual_monadDual: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Dual_monadDual.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                    &&&add(string("Applicative0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_applicativeDual()),
                                                                                           add(string("Bind1"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused_1|
                                                                                                                &PureScript_Data_Monoid_Dual::Data_Monoid_Dual_bindDual()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
}

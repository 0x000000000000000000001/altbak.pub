pub mod PureScript_Data_Identity {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_df3c4667::PureScript_Control_Comonad;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Identity_Identity() -> &dyn Any {
        static Data_Identity_Identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_Identity.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Identity_showIdentity() -> &dyn Any {
        static Data_Identity_showIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_showIdentity.get_or_init(||
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
                                                                                                                                    let x =
                                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                        &&&string("(Identity ")),
                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                 &&&dictShow),
                                                                                                                                                                                                                                                                              &&&x)),
                                                                                                                                                                                                        &&&string(")")))
                                                                                                                                }
                                                                                                                        }),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>()))))
    }
    pub fn Data_Identity_semiringIdentity() -> &dyn Any {
        static Data_Identity_semiringIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_semiringIdentity.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictSemiring|
                                                                       dictSemiring.clone()))
    }
    pub fn Data_Identity_semigroupIdentity() -> &dyn Any {
        static Data_Identity_semigroupIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_semigroupIdentity.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictSemigroup|
                                                                        dictSemigroup.clone()))
    }
    pub fn Data_Identity_ringIdentity() -> &dyn Any {
        static Data_Identity_ringIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_ringIdentity.get_or_init(||
                                                   &Func1::new(move |dictRing|
                                                                   dictRing.clone()))
    }
    pub fn Data_Identity_ordIdentity() -> &dyn Any {
        static Data_Identity_ordIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_ordIdentity.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  dictOrd.clone()))
    }
    pub fn Data_Identity_newtypeIdentity() -> &dyn Any {
        static Data_Identity_newtypeIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_newtypeIdentity.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                       &&&add(string("Coercible0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &Sharpurs_Prelude::Prim_undefined()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Identity_monoidIdentity() -> &dyn Any {
        static Data_Identity_monoidIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_monoidIdentity.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictMonoid|
                                                                     dictMonoid.clone()))
    }
    pub fn Data_Identity_lazyIdentity() -> &dyn Any {
        static Data_Identity_lazyIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_lazyIdentity.get_or_init(||
                                                   &Func1::new(move |dictLazy|
                                                                   dictLazy.clone()))
    }
    pub fn Data_Identity_heytingAlgebraIdentity() -> &dyn Any {
        static Data_Identity_heytingAlgebraIdentity: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Identity_heytingAlgebraIdentity.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictHeytingAlgebra|
                                                                             dictHeytingAlgebra.clone()))
    }
    pub fn Data_Identity_functorIdentity() -> &dyn Any {
        static Data_Identity_functorIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_functorIdentity.get_or_init(||
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
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                           &&&v))
                                                                                                                                   }
                                                                                                                           })),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Identity_invariantIdentity() -> &dyn Any {
        static Data_Identity_invariantIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_invariantIdentity.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                         &&&add(string("imap"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                                  &&&PureScript_Data_Identity::Data_Identity_functorIdentity()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Identity_extendIdentity() -> &dyn Any {
        static Data_Identity_extendIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_extendIdentity.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                      &&&add(string("extend"),
                                                                                             &&Func1::new(move
                                                                                                              |f|
                                                                                                              &Func1::new({
                                                                                                                              let f
                                                                                                                                  =
                                                                                                                                  f.clone();
                                                                                                                              move
                                                                                                                                  |m|
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                      m))
                                                                                                                          })),
                                                                                             add(string("Functor0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Identity::Data_Identity_functorIdentity()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Identity_euclideanRingIdentity() -> &dyn Any {
        static Data_Identity_euclideanRingIdentity: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Identity_euclideanRingIdentity.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictEuclideanRing|
                                                                            dictEuclideanRing.clone()))
    }
    pub fn Data_Identity_eqIdentity() -> &dyn Any {
        static Data_Identity_eqIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_eqIdentity.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 dictEq.clone()))
    }
    pub fn Data_Identity_eq1Identity() -> &dyn Any {
        static Data_Identity_eq1Identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_eq1Identity.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                   &&&add(string("eq1"),
                                                                                          &&Func1::new(move
                                                                                                           |dictEq|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_eqIdentity(),
                                                                                                                                                                               dictEq))),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Identity_ord1Identity() -> &dyn Any {
        static Data_Identity_ord1Identity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_ord1Identity.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                    &&&add(string("compare1"),
                                                                                           &&Func1::new(move
                                                                                                            |dictOrd|
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_ordIdentity(),
                                                                                                                                                                                dictOrd))),
                                                                                           add(string("Eq10"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Identity::Data_Identity_eq1Identity()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Identity_comonadIdentity() -> &dyn Any {
        static Data_Identity_comonadIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_comonadIdentity.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Comonad::Control_Comonad_Comonadusd_Dict(),
                                                                                       &&&add(string("extract"),
                                                                                              &&Func1::new(move
                                                                                                               |v|
                                                                                                               &Sharpurs_Prelude::unbox(v)),
                                                                                              add(string("Extend0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Identity::Data_Identity_extendIdentity()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Identity_commutativeRingIdentity() -> &dyn Any {
        static Data_Identity_commutativeRingIdentity:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_commutativeRingIdentity.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictCommutativeRing|
                                                                              dictCommutativeRing.clone()))
    }
    pub fn Data_Identity_boundedIdentity() -> &dyn Any {
        static Data_Identity_boundedIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_boundedIdentity.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictBounded|
                                                                      dictBounded.clone()))
    }
    pub fn Data_Identity_booleanAlgebraIdentity() -> &dyn Any {
        static Data_Identity_booleanAlgebraIdentity: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Identity_booleanAlgebraIdentity.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictBooleanAlgebra|
                                                                             dictBooleanAlgebra.clone()))
    }
    pub fn Data_Identity_applyIdentity() -> &dyn Any {
        static Data_Identity_applyIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_applyIdentity.get_or_init(||
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
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                         &&&matchValue_1))
                                                                                                                                 }
                                                                                                                         })),
                                                                                            add(string("Functor0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Identity::Data_Identity_functorIdentity()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Identity_bindIdentity() -> &dyn Any {
        static Data_Identity_bindIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_bindIdentity.get_or_init(||
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
                                                                                                                &PureScript_Data_Identity::Data_Identity_applyIdentity()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Identity_applicativeIdentity() -> &dyn Any {
        static Data_Identity_applicativeIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_applicativeIdentity.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                           &&&add(string("pure"),
                                                                                                  &&PureScript_Data_Identity::Data_Identity_Identity(),
                                                                                                  add(string("Apply0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_Identity::Data_Identity_applyIdentity()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))
    }
    pub fn Data_Identity_monadIdentity() -> &dyn Any {
        static Data_Identity_monadIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_monadIdentity.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                     &&&add(string("Applicative0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &PureScript_Data_Identity::Data_Identity_applicativeIdentity()),
                                                                                            add(string("Bind1"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused_1|
                                                                                                                 &PureScript_Data_Identity::Data_Identity_bindIdentity()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Identity_altIdentity() -> &dyn Any {
        static Data_Identity_altIdentity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Identity_altIdentity.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                   &&&add(string("alt"),
                                                                                          &&Func1::new(move
                                                                                                           |x|
                                                                                                           &Func1::new({
                                                                                                                           let x
                                                                                                                               =
                                                                                                                               x.clone();
                                                                                                                           move
                                                                                                                               |v|
                                                                                                                               &x
                                                                                                                       })),
                                                                                          add(string("Functor0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Identity::Data_Identity_functorIdentity()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
}

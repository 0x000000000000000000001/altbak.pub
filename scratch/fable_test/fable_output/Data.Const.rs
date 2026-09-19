pub mod PureScript_Data_Const {
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
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_ddf66a9c::PureScript_Data_Functor_Invariant;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Const_Const() -> &dyn Any {
        static Data_Const_Const: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_Const.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Const_showConst() -> &dyn Any {
        static Data_Const_showConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_showConst.get_or_init(||
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
                                                                                                                                                                                                  &&&string("(Const ")),
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
    pub fn Data_Const_semiringConst() -> &dyn Any {
        static Data_Const_semiringConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_semiringConst.get_or_init(||
                                                 &Func1::new(move
                                                                 |dictSemiring|
                                                                 dictSemiring.clone()))
    }
    pub fn Data_Const_semigroupoidConst() -> &dyn Any {
        static Data_Const_semigroupoidConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_semigroupoidConst.get_or_init(||
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
                                                                                                                                      let matchValue =
                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                      let matchValue_1 =
                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                       &&&matchValue_1)
                                                                                                                                  }
                                                                                                                          })),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Const_semigroupConst() -> &dyn Any {
        static Data_Const_semigroupConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_semigroupConst.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictSemigroup|
                                                                  dictSemigroup.clone()))
    }
    pub fn Data_Const_ringConst() -> &dyn Any {
        static Data_Const_ringConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_ringConst.get_or_init(||
                                             &Func1::new(move |dictRing|
                                                             dictRing.clone()))
    }
    pub fn Data_Const_ordConst() -> &dyn Any {
        static Data_Const_ordConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_ordConst.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            dictOrd.clone()))
    }
    pub fn Data_Const_newtypeConst() -> &dyn Any {
        static Data_Const_newtypeConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_newtypeConst.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                 &&&add(string("Coercible0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &Sharpurs_Prelude::Prim_undefined()),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Data_Const_monoidConst() -> &dyn Any {
        static Data_Const_monoidConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_monoidConst.get_or_init(||
                                               &Func1::new(move |dictMonoid|
                                                               dictMonoid.clone()))
    }
    pub fn Data_Const_heytingAlgebraConst() -> &dyn Any {
        static Data_Const_heytingAlgebraConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_heytingAlgebraConst.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictHeytingAlgebra|
                                                                       dictHeytingAlgebra.clone()))
    }
    pub fn Data_Const_functorConst() -> &dyn Any {
        static Data_Const_functorConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_functorConst.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                 &&&add(string("map"),
                                                                                        &&Func1::new(move
                                                                                                         |f|
                                                                                                         &Func1::new(move
                                                                                                                         |m|
                                                                                                                         {
                                                                                                                             let v =
                                                                                                                                 Sharpurs_Prelude::unbox(m);
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                              &&&v)
                                                                                                                         })),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Data_Const_invariantConst() -> &dyn Any {
        static Data_Const_invariantConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_invariantConst.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_Invariantusd_Dict(),
                                                                                   &&&add(string("imap"),
                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Invariant::Data_Functor_Invariant_imapF(),
                                                                                                                            &&&PureScript_Data_Const::Data_Const_functorConst()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Const_euclideanRingConst() -> &dyn Any {
        static Data_Const_euclideanRingConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_euclideanRingConst.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictEuclideanRing|
                                                                      dictEuclideanRing.clone()))
    }
    pub fn Data_Const_eqConst() -> &dyn Any {
        static Data_Const_eqConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_eqConst.get_or_init(||
                                           &Func1::new(move |dictEq|
                                                           dictEq.clone()))
    }
    pub fn Data_Const_eq1Const() -> &dyn Any {
        static Data_Const_eq1Const: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_eq1Const.get_or_init(||
                                            &Func1::new(move |dictEq|
                                                            {
                                                                let eq =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_eqConst(),
                                                                                                                                        dictEq));
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                                 &&&add(string("eq1"),
                                                                                                        &&Func1::new({
                                                                                                                         let eq
                                                                                                                             =
                                                                                                                             eq.clone();
                                                                                                                         move
                                                                                                                             |dictEq1|
                                                                                                                             &eq
                                                                                                                     }),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))
                                                            }))
    }
    pub fn Data_Const_ord1Const() -> &dyn Any {
        static Data_Const_ord1Const: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_ord1Const.get_or_init(||
                                             &Func1::new(move |dictOrd|
                                                             {
                                                                 let compare =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_ordConst(),
                                                                                                                                         dictOrd));
                                                                 let eq1Const1 =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_eq1Const(),
                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()));
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                                  &&&add(string("compare1"),
                                                                                                         &&Func1::new({
                                                                                                                          let compare
                                                                                                                              =
                                                                                                                              compare.clone();
                                                                                                                          move
                                                                                                                              |dictOrd1|
                                                                                                                              &compare
                                                                                                                      }),
                                                                                                         add(string("Eq10"),
                                                                                                             &&Func1::new({
                                                                                                                              let eq1Const1
                                                                                                                                  =
                                                                                                                                  eq1Const1.clone();
                                                                                                                              move
                                                                                                                                  |usd__unused|
                                                                                                                                  &eq1Const1
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
                                                             }))
    }
    pub fn Data_Const_commutativeRingConst() -> &dyn Any {
        static Data_Const_commutativeRingConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_commutativeRingConst.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictCommutativeRing|
                                                                        dictCommutativeRing.clone()))
    }
    pub fn Data_Const_boundedConst() -> &dyn Any {
        static Data_Const_boundedConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_boundedConst.get_or_init(||
                                                &Func1::new(move |dictBounded|
                                                                dictBounded.clone()))
    }
    pub fn Data_Const_booleanAlgebraConst() -> &dyn Any {
        static Data_Const_booleanAlgebraConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_booleanAlgebraConst.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictBooleanAlgebra|
                                                                       dictBooleanAlgebra.clone()))
    }
    pub fn Data_Const_applyConst() -> &dyn Any {
        static Data_Const_applyConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_applyConst.get_or_init(||
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
                                                                                                                                                   let matchValue =
                                                                                                                                                       Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                   let matchValue_1 =
                                                                                                                                                       Sharpurs_Prelude::unbox(v1);
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                             &&&dictSemigroup),
                                                                                                                                                                                                                                                          &&&matchValue),
                                                                                                                                                                                                                       &&&matchValue_1))
                                                                                                                                               }
                                                                                                                                       })
                                                                                                                   }),
                                                                                                      add(string("Functor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Const::Data_Const_functorConst()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))))
    }
    pub fn Data_Const_applicativeConst() -> &dyn Any {
        static Data_Const_applicativeConst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Const_applicativeConst.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictMonoid|
                                                                    {
                                                                        let applyConst1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_applyConst(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroup0"),
                                                                                                                                                       Sharpurs_Prelude::unbox(dictMonoid)),
                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                                         &&&add(string("pure"),
                                                                                                                &&Func1::new({
                                                                                                                                 let dictMonoid
                                                                                                                                     =
                                                                                                                                     dictMonoid.clone();
                                                                                                                                 move
                                                                                                                                     |v|
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Const::Data_Const_Const(),
                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                                                                                                         &&&dictMonoid))
                                                                                                                             }),
                                                                                                                add(string("Apply0"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let applyConst1
                                                                                                                                         =
                                                                                                                                         applyConst1.clone();
                                                                                                                                     move
                                                                                                                                         |usd__unused|
                                                                                                                                         &applyConst1
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>())))
                                                                    }))
    }
}

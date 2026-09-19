pub mod PureScript_Data_Equivalence {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_cf56105e::PureScript_Data_Functor_Contravariant;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Equivalence_Equivalence() -> &dyn Any {
        static Data_Equivalence_Equivalence: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Equivalence_Equivalence.get_or_init(||
                                                     &Func1::new(move |x|
                                                                     x.clone()))
    }
    pub fn Data_Equivalence_semigroupEquivalence() -> &dyn Any {
        static Data_Equivalence_semigroupEquivalence:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Equivalence_semigroupEquivalence.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                               &&&add(string("append"),
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
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence(),
                                                                                                                                                                                &&&Func1::new({
                                                                                                                                                                                                  let matchValue_1
                                                                                                                                                                                                      =
                                                                                                                                                                                                      matchValue_1.clone();
                                                                                                                                                                                                  move
                                                                                                                                                                                                      |a|
                                                                                                                                                                                                      &Func1::new({
                                                                                                                                                                                                                      let a
                                                                                                                                                                                                                          =
                                                                                                                                                                                                                          a.clone();
                                                                                                                                                                                                                      move
                                                                                                                                                                                                                          |b|
                                                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                    &&&a),
                                                                                                                                                                                                                                                                                                                                 b)),
                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                                                 &&&a),
                                                                                                                                                                                                                                                                                              b))
                                                                                                                                                                                                                  })
                                                                                                                                                                                              }))
                                                                                                                                           }
                                                                                                                                   })),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Data_Equivalence_newtypeEquivalence() -> &dyn Any {
        static Data_Equivalence_newtypeEquivalence: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Equivalence_newtypeEquivalence.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                             &&&add(string("Coercible0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &Sharpurs_Prelude::Prim_undefined()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Equivalence_monoidEquivalence() -> &dyn Any {
        static Data_Equivalence_monoidEquivalence: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Equivalence_monoidEquivalence.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                            &&&add(string("mempty"),
                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence(),
                                                                                                                                     &&&Func1::new(move
                                                                                                                                                       |v|
                                                                                                                                                       &Func1::new(move
                                                                                                                                                                       |v1|
                                                                                                                                                                       &true))),
                                                                                                   add(string("Semigroup0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Equivalence::Data_Equivalence_semigroupEquivalence()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Equivalence_defaultEquivalence() -> &dyn Any {
        static Data_Equivalence_defaultEquivalence: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Equivalence_defaultEquivalence.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictEq|
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                dictEq))))
    }
    pub fn Data_Equivalence_contravariantEquivalence() -> &dyn Any {
        static Data_Equivalence_contravariantEquivalence:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Equivalence_contravariantEquivalence.get_or_init(||
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor_Contravariant::Data_Functor_Contravariant_Contravariantusd_Dict(),
                                                                                                   &&&add(string("cmap"),
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
                                                                                                                                                   let matchValue_1 =
                                                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence(),
                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_on(),
                                                                                                                                                                                                                                                          &&&matchValue_1),
                                                                                                                                                                                                                       &&&matchValue))
                                                                                                                                               }
                                                                                                                                       })),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))
    }
    pub fn Data_Equivalence_comparisonEquivalence() -> &dyn Any {
        static Data_Equivalence_comparisonEquivalence:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Equivalence_comparisonEquivalence.get_or_init(||
                                                               &Func1::new(move
                                                                               |v|
                                                                               {
                                                                                   let p =
                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Equivalence::Data_Equivalence_Equivalence(),
                                                                                                                    &&&Func1::new({
                                                                                                                                      let p
                                                                                                                                          =
                                                                                                                                          p.clone();
                                                                                                                                      move
                                                                                                                                          |a|
                                                                                                                                          &Func1::new({
                                                                                                                                                          let a
                                                                                                                                                              =
                                                                                                                                                              a.clone();
                                                                                                                                                          move
                                                                                                                                                              |b|
                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ordering::Data_Ordering_eqOrdering()),
                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&p,
                                                                                                                                                                                                                                                                                                        &&&a),
                                                                                                                                                                                                                                                                     b)),
                                                                                                                                                                                               &&&LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor))
                                                                                                                                                      })
                                                                                                                                  }))
                                                                               }))
    }
}

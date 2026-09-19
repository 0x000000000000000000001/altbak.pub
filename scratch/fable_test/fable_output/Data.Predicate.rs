pub mod PureScript_Data_Predicate {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_13840e4f::PureScript_Data_BooleanAlgebra;
    use crate::module_cf56105e::PureScript_Data_Functor_Contravariant;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Predicate_Predicate() -> &dyn Any {
        static Data_Predicate_Predicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Predicate_Predicate.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Predicate_newtypePredicate() -> &dyn Any {
        static Data_Predicate_newtypePredicate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Predicate_newtypePredicate.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                         &&&add(string("Coercible0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &Sharpurs_Prelude::Prim_undefined()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Predicate_heytingAlgebraPredicate() -> &dyn Any {
        static Data_Predicate_heytingAlgebraPredicate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Predicate_heytingAlgebraPredicate.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraFunction(),
                                                                                                &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()))
    }
    pub fn Data_Predicate_contravariantPredicate() -> &dyn Any {
        static Data_Predicate_contravariantPredicate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Predicate_contravariantPredicate.get_or_init(||
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
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Predicate::Data_Predicate_Predicate(),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                                      &&&matchValue_1),
                                                                                                                                                                                                                   &&&matchValue))
                                                                                                                                           }
                                                                                                                                   })),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))
    }
    pub fn Data_Predicate_booleanAlgebraPredicate() -> &dyn Any {
        static Data_Predicate_booleanAlgebraPredicate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Predicate_booleanAlgebraPredicate.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_booleanAlgebraFn(),
                                                                                                &&&PureScript_Data_BooleanAlgebra::Data_BooleanAlgebra_booleanAlgebraBoolean()))
    }
}

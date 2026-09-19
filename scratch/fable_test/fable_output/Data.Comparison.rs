pub mod PureScript_Data_Comparison {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_cf56105e::PureScript_Data_Functor_Contravariant;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Comparison_semigroupFn() -> &dyn Any {
        static Data_Comparison_semigroupFn: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Comparison_semigroupFn.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_semigroupFn(),
                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_semigroupFn(),
                                                                                                                        &&&PureScript_Data_Ordering::Data_Ordering_semigroupOrdering())))
    }
    pub fn Data_Comparison_Comparison() -> &dyn Any {
        static Data_Comparison_Comparison: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Comparison_Comparison.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Comparison_semigroupComparison() -> &dyn Any {
        static Data_Comparison_semigroupComparison: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Comparison_semigroupComparison.get_or_init(||
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
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Comparison::Data_Comparison_Comparison(),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Comparison::Data_Comparison_semigroupFn()),
                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                 &&&matchValue_1))
                                                                                                                                         }
                                                                                                                                 })),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Comparison_newtypeComparison() -> &dyn Any {
        static Data_Comparison_newtypeComparison: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Comparison_newtypeComparison.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                           &&&add(string("Coercible0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &Sharpurs_Prelude::Prim_undefined()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Comparison_monoidComparison() -> &dyn Any {
        static Data_Comparison_monoidComparison: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Comparison_monoidComparison.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                          &&&add(string("mempty"),
                                                                                                 &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Comparison::Data_Comparison_Comparison(),
                                                                                                                                   &&&Func1::new(move
                                                                                                                                                     |v|
                                                                                                                                                     &Func1::new(move
                                                                                                                                                                     |v1|
                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)))),
                                                                                                 add(string("Semigroup0"),
                                                                                                     &&Func1::new(move
                                                                                                                      |usd__unused|
                                                                                                                      &PureScript_Data_Comparison::Data_Comparison_semigroupComparison()),
                                                                                                     empty::<string,
                                                                                                             &dyn Any>()))))
    }
    pub fn Data_Comparison_defaultComparison() -> &dyn Any {
        static Data_Comparison_defaultComparison: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Comparison_defaultComparison.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictOrd|
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Comparison::Data_Comparison_Comparison(),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                              dictOrd))))
    }
    pub fn Data_Comparison_contravariantComparison() -> &dyn Any {
        static Data_Comparison_contravariantComparison:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Comparison_contravariantComparison.get_or_init(||
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
                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Comparison::Data_Comparison_Comparison(),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_on(),
                                                                                                                                                                                                                                                        &&&matchValue_1),
                                                                                                                                                                                                                     &&&matchValue))
                                                                                                                                             }
                                                                                                                                     })),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))
    }
}

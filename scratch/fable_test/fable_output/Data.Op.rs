pub mod PureScript_Data_Op {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_cf56105e::PureScript_Data_Functor_Contravariant;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Op_Op() -> &dyn Any {
        static Data_Op_Op: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Op_Op.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Op_semigroupoidOp() -> &dyn Any {
        static Data_Op_semigroupoidOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Op_semigroupoidOp.get_or_init(||
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
                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Op::Data_Op_Op(),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                          &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                       &&&matchValue_1),
                                                                                                                                                                                                    &&&matchValue))
                                                                                                                            }
                                                                                                                    })),
                                                                                       empty::<string,
                                                                                               &dyn Any>())))
    }
    pub fn Data_Op_semigroupOp() -> &dyn Any {
        static Data_Op_semigroupOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Op_semigroupOp.get_or_init(||
                                            &Func1::new(move |dictSemigroup|
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_semigroupFn(),
                                                                                             dictSemigroup)))
    }
    pub fn Data_Op_newtypeOp() -> &dyn Any {
        static Data_Op_newtypeOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Op_newtypeOp.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                           &&&add(string("Coercible0"),
                                                                                  &&Func1::new(move
                                                                                                   |usd__unused|
                                                                                                   &Sharpurs_Prelude::Prim_undefined()),
                                                                                  empty::<string,
                                                                                          &dyn Any>())))
    }
    pub fn Data_Op_monoidOp() -> &dyn Any {
        static Data_Op_monoidOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Op_monoidOp.get_or_init(||
                                         &Func1::new(move |dictMonoid|
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_monoidFn(),
                                                                                          dictMonoid)))
    }
    pub fn Data_Op_contravariantOp() -> &dyn Any {
        static Data_Op_contravariantOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Op_contravariantOp.get_or_init(||
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
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Op::Data_Op_Op(),
                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                                                                        &&&matchValue_1),
                                                                                                                                                                                                     &&&matchValue))
                                                                                                                             }
                                                                                                                     })),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Data_Op_categoryOp() -> &dyn Any {
        static Data_Op_categoryOp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Op_categoryOp.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_Categoryusd_Dict(),
                                                                            &&&add(string("identity"),
                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Op::Data_Op_Op(),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                        &&&PureScript_Control_Category::Control_Category_categoryFn())),
                                                                                   add(string("Semigroupoid0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Op::Data_Op_semigroupoidOp()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
}

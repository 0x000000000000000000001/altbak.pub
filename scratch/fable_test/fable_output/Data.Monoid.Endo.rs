pub mod PureScript_Data_Monoid_Endo {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Monoid_Endo_Endo() -> &dyn Any {
        static Data_Monoid_Endo_Endo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Endo_Endo.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Monoid_Endo_showEndo() -> &dyn Any {
        static Data_Monoid_Endo_showEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Endo_showEndo.get_or_init(||
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
                                                                                                                                                                                                       &&&string("(Endo ")),
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
    pub fn Data_Monoid_Endo_semigroupEndo() -> &dyn Any {
        static Data_Monoid_Endo_semigroupEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Endo_semigroupEndo.get_or_init(||
                                                       &Func1::new(move
                                                                       |dictSemigroupoid|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                        &&&add(string("append"),
                                                                                                               &&Func1::new({
                                                                                                                                let dictSemigroupoid
                                                                                                                                    =
                                                                                                                                    dictSemigroupoid.clone();
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
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo(),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                                                                                                      &&&dictSemigroupoid),
                                                                                                                                                                                                                                                                   &&&matchValue),
                                                                                                                                                                                                                                &&&matchValue_1))
                                                                                                                                                        }
                                                                                                                                                })
                                                                                                                            }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Monoid_Endo_ordEndo() -> &dyn Any {
        static Data_Monoid_Endo_ordEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Endo_ordEndo.get_or_init(||
                                                 &Func1::new(move |dictOrd|
                                                                 dictOrd.clone()))
    }
    pub fn Data_Monoid_Endo_monoidEndo() -> &dyn Any {
        static Data_Monoid_Endo_monoidEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Endo_monoidEndo.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictCategory|
                                                                    {
                                                                        let semigroupEndo1 =
                                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_semigroupEndo(),
                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semigroupoid0"),
                                                                                                                                                       Sharpurs_Prelude::unbox(dictCategory)),
                                                                                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                         &&&add(string("mempty"),
                                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Endo::Data_Monoid_Endo_Endo(),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                                     dictCategory)),
                                                                                                                add(string("Semigroup0"),
                                                                                                                    &&Func1::new({
                                                                                                                                     let semigroupEndo1
                                                                                                                                         =
                                                                                                                                         semigroupEndo1.clone();
                                                                                                                                     move
                                                                                                                                         |usd__unused|
                                                                                                                                         &semigroupEndo1
                                                                                                                                 }),
                                                                                                                    empty::<string,
                                                                                                                            &dyn Any>())))
                                                                    }))
    }
    pub fn Data_Monoid_Endo_eqEndo() -> &dyn Any {
        static Data_Monoid_Endo_eqEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Endo_eqEndo.get_or_init(||
                                                &Func1::new(move |dictEq|
                                                                dictEq.clone()))
    }
    pub fn Data_Monoid_Endo_boundedEndo() -> &dyn Any {
        static Data_Monoid_Endo_boundedEndo: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Endo_boundedEndo.get_or_init(||
                                                     &Func1::new(move
                                                                     |dictBounded|
                                                                     dictBounded.clone()))
    }
}

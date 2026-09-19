pub mod PureScript_Data_Ord_Max {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Ord_Max_Max() -> &dyn Any {
        static Data_Ord_Max_Max: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Max_Max.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Ord_Max_showMax() -> &dyn Any {
        static Data_Ord_Max_showMax: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Max_showMax.get_or_init(||
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
                                                                                                                                                                                                  &&&string("(Max ")),
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
    pub fn Data_Ord_Max_semigroupMax() -> &dyn Any {
        static Data_Ord_Max_semigroupMax: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Max_semigroupMax.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                   &&&add(string("append"),
                                                                                                          &&Func1::new({
                                                                                                                           let dictOrd
                                                                                                                               =
                                                                                                                               dictOrd.clone();
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
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Max::Data_Ord_Max_Max(),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_max(),
                                                                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                                                           &&&matchValue_1))
                                                                                                                                                   }
                                                                                                                                           })
                                                                                                                       }),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Ord_Max_newtypeMax() -> &dyn Any {
        static Data_Ord_Max_newtypeMax: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Max_newtypeMax.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                 &&&add(string("Coercible0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &Sharpurs_Prelude::Prim_undefined()),
                                                                                        empty::<string,
                                                                                                &dyn Any>())))
    }
    pub fn Data_Ord_Max_monoidMax() -> &dyn Any {
        static Data_Ord_Max_monoidMax: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Max_monoidMax.get_or_init(||
                                               &Func1::new(move |dictBounded|
                                                               {
                                                                   let semigroupMax1 =
                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Max::Data_Ord_Max_semigroupMax(),
                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                  Sharpurs_Prelude::unbox(dictBounded)),
                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()));
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                    &&&add(string("mempty"),
                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Max::Data_Ord_Max_Max(),
                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                dictBounded)),
                                                                                                           add(string("Semigroup0"),
                                                                                                               &&Func1::new({
                                                                                                                                let semigroupMax1
                                                                                                                                    =
                                                                                                                                    semigroupMax1.clone();
                                                                                                                                move
                                                                                                                                    |usd__unused|
                                                                                                                                    &semigroupMax1
                                                                                                                            }),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>())))
                                                               }))
    }
    pub fn Data_Ord_Max_eqMax() -> &dyn Any {
        static Data_Ord_Max_eqMax: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Max_eqMax.get_or_init(||
                                           &Func1::new(move |dictEq|
                                                           dictEq.clone()))
    }
    pub fn Data_Ord_Max_ordMax() -> &dyn Any {
        static Data_Ord_Max_ordMax: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Max_ordMax.get_or_init(||
                                            &Func1::new(move |dictOrd|
                                                            {
                                                                let eqMax1 =
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Max::Data_Ord_Max_eqMax(),
                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                               Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                        &&&Sharpurs_Prelude::Prim_undefined()));
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                 &&&add(string("compare"),
                                                                                                        &&Func1::new({
                                                                                                                         let dictOrd
                                                                                                                             =
                                                                                                                             dictOrd.clone();
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
                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                            &&&dictOrd),
                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                      &&&matchValue_1)
                                                                                                                                                 }
                                                                                                                                         })
                                                                                                                     }),
                                                                                                        add(string("Eq0"),
                                                                                                            &&Func1::new({
                                                                                                                             let eqMax1
                                                                                                                                 =
                                                                                                                                 eqMax1.clone();
                                                                                                                             move
                                                                                                                                 |usd__unused|
                                                                                                                                 &eqMax1
                                                                                                                         }),
                                                                                                            empty::<string,
                                                                                                                    &dyn Any>())))
                                                            }))
    }
}

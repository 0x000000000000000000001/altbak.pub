pub mod PureScript_Data_Ord_Down {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Ord_Down_Down() -> &dyn Any {
        static Data_Ord_Down_Down: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Down_Down.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Ord_Down_showDown() -> &dyn Any {
        static Data_Ord_Down_showDown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Down_showDown.get_or_init(||
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
                                                                                                                                                                                                    &&&string("(Down ")),
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
    pub fn Data_Ord_Down_newtypeDown() -> &dyn Any {
        static Data_Ord_Down_newtypeDown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Down_newtypeDown.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                   &&&add(string("Coercible0"),
                                                                                          &&Func1::new(move
                                                                                                           |usd__unused|
                                                                                                           &Sharpurs_Prelude::Prim_undefined()),
                                                                                          empty::<string,
                                                                                                  &dyn Any>())))
    }
    pub fn Data_Ord_Down_eqDown() -> &dyn Any {
        static Data_Ord_Down_eqDown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Down_eqDown.get_or_init(||
                                             &Func1::new(move |dictEq|
                                                             dictEq.clone()))
    }
    pub fn Data_Ord_Down_ordDown() -> &dyn Any {
        static Data_Ord_Down_ordDown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Down_ordDown.get_or_init(||
                                              &Func1::new(move |dictOrd|
                                                              {
                                                                  let eqDown1 =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Down::Data_Ord_Down_eqDown(),
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
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ordering::Data_Ordering_invert(),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                 &&&dictOrd),
                                                                                                                                                                                                                                                              &&&matchValue),
                                                                                                                                                                                                                           &&&matchValue_1))
                                                                                                                                                   }
                                                                                                                                           })
                                                                                                                       }),
                                                                                                          add(string("Eq0"),
                                                                                                              &&Func1::new({
                                                                                                                               let eqDown1
                                                                                                                                   =
                                                                                                                                   eqDown1.clone();
                                                                                                                               move
                                                                                                                                   |usd__unused|
                                                                                                                                   &eqDown1
                                                                                                                           }),
                                                                                                              empty::<string,
                                                                                                                      &dyn Any>())))
                                                              }))
    }
    pub fn Data_Ord_Down_boundedDown() -> &dyn Any {
        static Data_Ord_Down_boundedDown: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Ord_Down_boundedDown.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictBounded|
                                                                  {
                                                                      let ordDown1 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Down::Data_Ord_Down_ordDown(),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ord0"),
                                                                                                                                                     Sharpurs_Prelude::unbox(dictBounded)),
                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                                       &&&add(string("top"),
                                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Down::Data_Ord_Down_Down(),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                   dictBounded)),
                                                                                                              add(string("bottom"),
                                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord_Down::Data_Ord_Down_Down(),
                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                       dictBounded)),
                                                                                                                  add(string("Ord0"),
                                                                                                                      &&Func1::new({
                                                                                                                                       let ordDown1
                                                                                                                                           =
                                                                                                                                           ordDown1.clone();
                                                                                                                                       move
                                                                                                                                           |usd__unused|
                                                                                                                                           &ordDown1
                                                                                                                                   }),
                                                                                                                      empty::<string,
                                                                                                                              &dyn Any>()))))
                                                                  }))
    }
}

pub mod PureScript_Data_Monoid_Alternate {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Monoid_Alternate_Alternate() -> &dyn Any {
        static Data_Monoid_Alternate_Alternate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_Alternate.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Data_Monoid_Alternate_showAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_showAlternate: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Alternate_showAlternate.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictShow|
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
                                                                                                                                                                                                                 &&&string("(Alternate ")),
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
    pub fn Data_Monoid_Alternate_semigroupAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_semigroupAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_semigroupAlternate.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |dictAlt|
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                                                  &&&add(string("append"),
                                                                                                                         &&Func1::new({
                                                                                                                                          let dictAlt
                                                                                                                                              =
                                                                                                                                              dictAlt.clone();
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
                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Alternate::Data_Monoid_Alternate_Alternate(),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_alt(),
                                                                                                                                                                                                                                                                                                                &&&dictAlt),
                                                                                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                                                                                          &&&matchValue_1))
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      }),
                                                                                                                         empty::<string,
                                                                                                                                 &dyn Any>()))))
    }
    pub fn Data_Monoid_Alternate_plusAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_plusAlternate: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Alternate_plusAlternate.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictPlus|
                                                                            dictPlus.clone()))
    }
    pub fn Data_Monoid_Alternate_ordAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_ordAlternate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_ordAlternate.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictOrd|
                                                                           dictOrd.clone()))
    }
    pub fn Data_Monoid_Alternate_ord1Alternate() -> &dyn Any {
        static Data_Monoid_Alternate_ord1Alternate: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Alternate_ord1Alternate.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictOrd1|
                                                                            dictOrd1.clone()))
    }
    pub fn Data_Monoid_Alternate_newtypeAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_newtypeAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_newtypeAlternate.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                &&&add(string("Coercible0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &Sharpurs_Prelude::Prim_undefined()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Data_Monoid_Alternate_monoidAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_monoidAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_monoidAlternate.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictPlus|
                                                                              {
                                                                                  let semigroupAlternate1 =
                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Alternate::Data_Monoid_Alternate_semigroupAlternate(),
                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Alt0"),
                                                                                                                                                                 Sharpurs_Prelude::unbox(dictPlus)),
                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                                                   &&&add(string("mempty"),
                                                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid_Alternate::Data_Monoid_Alternate_Alternate(),
                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_empty(),
                                                                                                                                                                                               dictPlus)),
                                                                                                                          add(string("Semigroup0"),
                                                                                                                              &&Func1::new({
                                                                                                                                               let semigroupAlternate1
                                                                                                                                                   =
                                                                                                                                                   semigroupAlternate1.clone();
                                                                                                                                               move
                                                                                                                                                   |usd__unused|
                                                                                                                                                   &semigroupAlternate1
                                                                                                                                           }),
                                                                                                                              empty::<string,
                                                                                                                                      &dyn Any>())))
                                                                              }))
    }
    pub fn Data_Monoid_Alternate_monadAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_monadAlternate: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Alternate_monadAlternate.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictMonad|
                                                                             dictMonad.clone()))
    }
    pub fn Data_Monoid_Alternate_functorAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_functorAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_functorAlternate.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictFunctor|
                                                                               dictFunctor.clone()))
    }
    pub fn Data_Monoid_Alternate_extendAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_extendAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_extendAlternate.get_or_init(||
                                                              &Func1::new(move
                                                                              |dictExtend|
                                                                              dictExtend.clone()))
    }
    pub fn Data_Monoid_Alternate_eqAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_eqAlternate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_eqAlternate.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictEq|
                                                                          dictEq.clone()))
    }
    pub fn Data_Monoid_Alternate_eq1Alternate() -> &dyn Any {
        static Data_Monoid_Alternate_eq1Alternate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_eq1Alternate.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictEq1|
                                                                           dictEq1.clone()))
    }
    pub fn Data_Monoid_Alternate_comonadAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_comonadAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_comonadAlternate.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictComonad|
                                                                               dictComonad.clone()))
    }
    pub fn Data_Monoid_Alternate_boundedAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_boundedAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_boundedAlternate.get_or_init(||
                                                               &Func1::new(move
                                                                               |dictBounded|
                                                                               dictBounded.clone()))
    }
    pub fn Data_Monoid_Alternate_bindAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_bindAlternate: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Alternate_bindAlternate.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictBind|
                                                                            dictBind.clone()))
    }
    pub fn Data_Monoid_Alternate_applyAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_applyAlternate: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Monoid_Alternate_applyAlternate.get_or_init(||
                                                             &Func1::new(move
                                                                             |dictApply|
                                                                             dictApply.clone()))
    }
    pub fn Data_Monoid_Alternate_applicativeAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_applicativeAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_applicativeAlternate.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictApplicative|
                                                                                   dictApplicative.clone()))
    }
    pub fn Data_Monoid_Alternate_alternativeAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_alternativeAlternate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_alternativeAlternate.get_or_init(||
                                                                   &Func1::new(move
                                                                                   |dictAlternative|
                                                                                   dictAlternative.clone()))
    }
    pub fn Data_Monoid_Alternate_altAlternate() -> &dyn Any {
        static Data_Monoid_Alternate_altAlternate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Monoid_Alternate_altAlternate.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictAlt|
                                                                           dictAlt.clone()))
    }
}

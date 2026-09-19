pub mod PureScript_Data_Maybe_Last {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_73921b5b::PureScript_Control_Alt;
    use crate::module_9699daad::PureScript_Control_Alternative;
    use crate::module_6afec8d8::PureScript_Control_Plus;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Maybe_Last_Last() -> &dyn Any {
        static Data_Maybe_Last_Last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_Last.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Maybe_Last_showLast() -> &dyn Any {
        static Data_Maybe_Last_showLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_showLast.get_or_init(||
                                                 &Func1::new(move |dictShow|
                                                                 {
                                                                     let showMaybe =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_showMaybe(),
                                                                                                          dictShow);
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                      &&&add(string("show"),
                                                                                                             &&Func1::new({
                                                                                                                              let showMaybe
                                                                                                                                  =
                                                                                                                                  showMaybe.clone();
                                                                                                                              move
                                                                                                                                  |v|
                                                                                                                                  {
                                                                                                                                      let a =
                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                             &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                          &&&string("(Last ")),
                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                   &&&showMaybe),
                                                                                                                                                                                                                                                                                &&&a)),
                                                                                                                                                                                                          &&&string(")")))
                                                                                                                                  }
                                                                                                                          }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>()))
                                                                 }))
    }
    pub fn Data_Maybe_Last_semigroupLast() -> &dyn Any {
        static Data_Maybe_Last_semigroupLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_semigroupLast.get_or_init(||
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
                                                                                                                                       let matchValue_1:
                                                                                                                                               LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                           Sharpurs_Prelude::unbox(v1);
                                                                                                                                       match matchValue_1.as_ref()
                                                                                                                                           {
                                                                                                                                           Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor
                                                                                                                                           =>
                                                                                                                                           &matchValue,
                                                                                                                                           _
                                                                                                                                           =>
                                                                                                                                           &matchValue_1,
                                                                                                                                       }
                                                                                                                                   }
                                                                                                                           })),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Maybe_Last_ordLast() -> &dyn Any {
        static Data_Maybe_Last_ordLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_ordLast.get_or_init(||
                                                &Func1::new(move |dictOrd|
                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_ordMaybe(),
                                                                                                 dictOrd)))
    }
    pub fn Data_Maybe_Last_ord1Last() -> &dyn Any {
        static Data_Maybe_Last_ord1Last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_ord1Last.get_or_init(||
                                                 &PureScript_Data_Maybe::Data_Maybe_ord1Maybe())
    }
    pub fn Data_Maybe_Last_newtypeLast() -> &dyn Any {
        static Data_Maybe_Last_newtypeLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_newtypeLast.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                     &&&add(string("Coercible0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &Sharpurs_Prelude::Prim_undefined()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_Maybe_Last_monoidLast() -> &dyn Any {
        static Data_Maybe_Last_monoidLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_monoidLast.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                    &&&add(string("mempty"),
                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe_Last::Data_Maybe_Last_Last(),
                                                                                                                             &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                           add(string("Semigroup0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Maybe_Last::Data_Maybe_Last_semigroupLast()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Maybe_Last_monadLast() -> &dyn Any {
        static Data_Maybe_Last_monadLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_monadLast.get_or_init(||
                                                  &PureScript_Data_Maybe::Data_Maybe_monadMaybe())
    }
    pub fn Data_Maybe_Last_invariantLast() -> &dyn Any {
        static Data_Maybe_Last_invariantLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_invariantLast.get_or_init(||
                                                      &PureScript_Data_Maybe::Data_Maybe_invariantMaybe())
    }
    pub fn Data_Maybe_Last_functorLast() -> &dyn Any {
        static Data_Maybe_Last_functorLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_functorLast.get_or_init(||
                                                    &PureScript_Data_Maybe::Data_Maybe_functorMaybe())
    }
    pub fn Data_Maybe_Last_extendLast() -> &dyn Any {
        static Data_Maybe_Last_extendLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_extendLast.get_or_init(||
                                                   &PureScript_Data_Maybe::Data_Maybe_extendMaybe())
    }
    pub fn Data_Maybe_Last_eqLast() -> &dyn Any {
        static Data_Maybe_Last_eqLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_eqLast.get_or_init(||
                                               &Func1::new(move |dictEq|
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_eqMaybe(),
                                                                                                dictEq)))
    }
    pub fn Data_Maybe_Last_eq1Last() -> &dyn Any {
        static Data_Maybe_Last_eq1Last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_eq1Last.get_or_init(||
                                                &PureScript_Data_Maybe::Data_Maybe_eq1Maybe())
    }
    pub fn Data_Maybe_Last_boundedLast() -> &dyn Any {
        static Data_Maybe_Last_boundedLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_boundedLast.get_or_init(||
                                                    &Func1::new(move
                                                                    |dictBounded|
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_boundedMaybe(),
                                                                                                     dictBounded)))
    }
    pub fn Data_Maybe_Last_bindLast() -> &dyn Any {
        static Data_Maybe_Last_bindLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_bindLast.get_or_init(||
                                                 &PureScript_Data_Maybe::Data_Maybe_bindMaybe())
    }
    pub fn Data_Maybe_Last_applyLast() -> &dyn Any {
        static Data_Maybe_Last_applyLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_applyLast.get_or_init(||
                                                  &PureScript_Data_Maybe::Data_Maybe_applyMaybe())
    }
    pub fn Data_Maybe_Last_applicativeLast() -> &dyn Any {
        static Data_Maybe_Last_applicativeLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_applicativeLast.get_or_init(||
                                                        &PureScript_Data_Maybe::Data_Maybe_applicativeMaybe())
    }
    pub fn Data_Maybe_Last_altLast() -> &dyn Any {
        static Data_Maybe_Last_altLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_altLast.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                 &&&add(string("alt"),
                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                          &&&PureScript_Data_Maybe_Last::Data_Maybe_Last_semigroupLast()),
                                                                                        add(string("Functor0"),
                                                                                            &&Func1::new(move
                                                                                                             |usd__unused|
                                                                                                             &PureScript_Data_Maybe_Last::Data_Maybe_Last_functorLast()),
                                                                                            empty::<string,
                                                                                                    &dyn Any>()))))
    }
    pub fn Data_Maybe_Last_plusLast() -> &dyn Any {
        static Data_Maybe_Last_plusLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_plusLast.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                  &&&add(string("empty"),
                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                           &&&PureScript_Data_Maybe_Last::Data_Maybe_Last_monoidLast()),
                                                                                         add(string("Alt0"),
                                                                                             &&Func1::new(move
                                                                                                              |usd__unused|
                                                                                                              &PureScript_Data_Maybe_Last::Data_Maybe_Last_altLast()),
                                                                                             empty::<string,
                                                                                                     &dyn Any>()))))
    }
    pub fn Data_Maybe_Last_alternativeLast() -> &dyn Any {
        static Data_Maybe_Last_alternativeLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_Last_alternativeLast.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                         &&&add(string("Applicative0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Maybe_Last::Data_Maybe_Last_applicativeLast()),
                                                                                                add(string("Plus1"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused_1|
                                                                                                                     &PureScript_Data_Maybe_Last::Data_Maybe_Last_plusLast()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
}

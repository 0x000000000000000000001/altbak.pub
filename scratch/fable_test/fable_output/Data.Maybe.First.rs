pub mod PureScript_Data_Maybe_First {
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
    pub fn Data_Maybe_First_First() -> &dyn Any {
        static Data_Maybe_First_First: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_First.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Maybe_First_showFirst() -> &dyn Any {
        static Data_Maybe_First_showFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_showFirst.get_or_init(||
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
                                                                                                                                                                                                            &&&string("First (")),
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
    pub fn Data_Maybe_First_semigroupFirst() -> &dyn Any {
        static Data_Maybe_First_semigroupFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_semigroupFirst.get_or_init(||
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
                                                                                                                                         let matchValue:
                                                                                                                                                 LrcPtr<Data_Maybe_Maybe> =
                                                                                                                                             Sharpurs_Prelude::unbox(&&v);
                                                                                                                                         if let Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(matchValue_1_0)
                                                                                                                                                =
                                                                                                                                                matchValue.as_ref()
                                                                                                                                            {
                                                                                                                                             &matchValue
                                                                                                                                         } else {
                                                                                                                                             &Sharpurs_Prelude::unbox(v1)
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                             })),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Maybe_First_ordFirst() -> &dyn Any {
        static Data_Maybe_First_ordFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_ordFirst.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_ordMaybe(),
                                                                                                   dictOrd)))
    }
    pub fn Data_Maybe_First_ord1First() -> &dyn Any {
        static Data_Maybe_First_ord1First: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_ord1First.get_or_init(||
                                                   &PureScript_Data_Maybe::Data_Maybe_ord1Maybe())
    }
    pub fn Data_Maybe_First_newtypeFirst() -> &dyn Any {
        static Data_Maybe_First_newtypeFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_newtypeFirst.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                       &&&add(string("Coercible0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &Sharpurs_Prelude::Prim_undefined()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Maybe_First_monoidFirst() -> &dyn Any {
        static Data_Maybe_First_monoidFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_monoidFirst.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                      &&&add(string("mempty"),
                                                                                             &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe_First::Data_Maybe_First_First(),
                                                                                                                               &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                             add(string("Semigroup0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Maybe_First::Data_Maybe_First_semigroupFirst()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Maybe_First_monadFirst() -> &dyn Any {
        static Data_Maybe_First_monadFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_monadFirst.get_or_init(||
                                                    &PureScript_Data_Maybe::Data_Maybe_monadMaybe())
    }
    pub fn Data_Maybe_First_invariantFirst() -> &dyn Any {
        static Data_Maybe_First_invariantFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_invariantFirst.get_or_init(||
                                                        &PureScript_Data_Maybe::Data_Maybe_invariantMaybe())
    }
    pub fn Data_Maybe_First_functorFirst() -> &dyn Any {
        static Data_Maybe_First_functorFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_functorFirst.get_or_init(||
                                                      &PureScript_Data_Maybe::Data_Maybe_functorMaybe())
    }
    pub fn Data_Maybe_First_extendFirst() -> &dyn Any {
        static Data_Maybe_First_extendFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_extendFirst.get_or_init(||
                                                     &PureScript_Data_Maybe::Data_Maybe_extendMaybe())
    }
    pub fn Data_Maybe_First_eqFirst() -> &dyn Any {
        static Data_Maybe_First_eqFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_eqFirst.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_eqMaybe(),
                                                                                                  dictEq)))
    }
    pub fn Data_Maybe_First_eq1First() -> &dyn Any {
        static Data_Maybe_First_eq1First: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_eq1First.get_or_init(||
                                                  &PureScript_Data_Maybe::Data_Maybe_eq1Maybe())
    }
    pub fn Data_Maybe_First_boundedFirst() -> &dyn Any {
        static Data_Maybe_First_boundedFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_boundedFirst.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictBounded|
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_boundedMaybe(),
                                                                                                       dictBounded)))
    }
    pub fn Data_Maybe_First_bindFirst() -> &dyn Any {
        static Data_Maybe_First_bindFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_bindFirst.get_or_init(||
                                                   &PureScript_Data_Maybe::Data_Maybe_bindMaybe())
    }
    pub fn Data_Maybe_First_applyFirst() -> &dyn Any {
        static Data_Maybe_First_applyFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_applyFirst.get_or_init(||
                                                    &PureScript_Data_Maybe::Data_Maybe_applyMaybe())
    }
    pub fn Data_Maybe_First_applicativeFirst() -> &dyn Any {
        static Data_Maybe_First_applicativeFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_applicativeFirst.get_or_init(||
                                                          &PureScript_Data_Maybe::Data_Maybe_applicativeMaybe())
    }
    pub fn Data_Maybe_First_altFirst() -> &dyn Any {
        static Data_Maybe_First_altFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_altFirst.get_or_init(||
                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alt::Control_Alt_Altusd_Dict(),
                                                                                   &&&add(string("alt"),
                                                                                          &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                            &&&PureScript_Data_Maybe_First::Data_Maybe_First_semigroupFirst()),
                                                                                          add(string("Functor0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Maybe_First::Data_Maybe_First_functorFirst()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_Maybe_First_plusFirst() -> &dyn Any {
        static Data_Maybe_First_plusFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_plusFirst.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Plus::Control_Plus_Plususd_Dict(),
                                                                                    &&&add(string("empty"),
                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_mempty(),
                                                                                                                             &&&PureScript_Data_Maybe_First::Data_Maybe_First_monoidFirst()),
                                                                                           add(string("Alt0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Maybe_First::Data_Maybe_First_altFirst()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>()))))
    }
    pub fn Data_Maybe_First_alternativeFirst() -> &dyn Any {
        static Data_Maybe_First_alternativeFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Maybe_First_alternativeFirst.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Alternative::Control_Alternative_Alternativeusd_Dict(),
                                                                                           &&&add(string("Applicative0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Maybe_First::Data_Maybe_First_applicativeFirst()),
                                                                                                  add(string("Plus1"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused_1|
                                                                                                                       &PureScript_Data_Maybe_First::Data_Maybe_First_plusFirst()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>()))))
    }
}

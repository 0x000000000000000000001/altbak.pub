pub mod PureScript_Data_Semigroup_Last {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_ed1bca0b::PureScript_Control_Monad;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Semigroup_Last_Last() -> &dyn Any {
        static Data_Semigroup_Last_Last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_Last.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Semigroup_Last_showLast() -> &dyn Any {
        static Data_Semigroup_Last_showLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_showLast.get_or_init(||
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
                                                                                                                                                                                                          &&&string("(Last ")),
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
    pub fn Data_Semigroup_Last_semigroupLast() -> &dyn Any {
        static Data_Semigroup_Last_semigroupLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_semigroupLast.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                           &&&add(string("append"),
                                                                                                  &&Func1::new(move
                                                                                                                   |v|
                                                                                                                   &Func1::new(move
                                                                                                                                   |x|
                                                                                                                                   x.clone())),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Semigroup_Last_ordLast() -> &dyn Any {
        static Data_Semigroup_Last_ordLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_ordLast.get_or_init(||
                                                    &Func1::new(move |dictOrd|
                                                                    dictOrd.clone()))
    }
    pub fn Data_Semigroup_Last_functorLast() -> &dyn Any {
        static Data_Semigroup_Last_functorLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_functorLast.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                         &&&add(string("map"),
                                                                                                &&Func1::new(move
                                                                                                                 |f|
                                                                                                                 &Func1::new({
                                                                                                                                 let f
                                                                                                                                     =
                                                                                                                                     f.clone();
                                                                                                                                 move
                                                                                                                                     |m|
                                                                                                                                     {
                                                                                                                                         let v =
                                                                                                                                             Sharpurs_Prelude::unbox(m);
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_Last(),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                             &&&v))
                                                                                                                                     }
                                                                                                                             })),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Semigroup_Last_eqLast() -> &dyn Any {
        static Data_Semigroup_Last_eqLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_eqLast.get_or_init(||
                                                   &Func1::new(move |dictEq|
                                                                   dictEq.clone()))
    }
    pub fn Data_Semigroup_Last_eq1Last() -> &dyn Any {
        static Data_Semigroup_Last_eq1Last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_eq1Last.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                     &&&add(string("eq1"),
                                                                                            &&Func1::new(move
                                                                                                             |dictEq|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_eqLast(),
                                                                                                                                                                                 dictEq))),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_Semigroup_Last_ord1Last() -> &dyn Any {
        static Data_Semigroup_Last_ord1Last: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_ord1Last.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                      &&&add(string("compare1"),
                                                                                             &&Func1::new(move
                                                                                                              |dictOrd|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_ordLast(),
                                                                                                                                                                                  dictOrd))),
                                                                                             add(string("Eq10"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Semigroup_Last::Data_Semigroup_Last_eq1Last()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Semigroup_Last_boundedLast() -> &dyn Any {
        static Data_Semigroup_Last_boundedLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_boundedLast.get_or_init(||
                                                        &Func1::new(move
                                                                        |dictBounded|
                                                                        dictBounded.clone()))
    }
    pub fn Data_Semigroup_Last_applyLast() -> &dyn Any {
        static Data_Semigroup_Last_applyLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_applyLast.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_Applyusd_Dict(),
                                                                                       &&&add(string("apply"),
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
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_Last(),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                           &&&matchValue_1))
                                                                                                                                   }
                                                                                                                           })),
                                                                                              add(string("Functor0"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused|
                                                                                                                   &PureScript_Data_Semigroup_Last::Data_Semigroup_Last_functorLast()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
    pub fn Data_Semigroup_Last_bindLast() -> &dyn Any {
        static Data_Semigroup_Last_bindLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_bindLast.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_Bindusd_Dict(),
                                                                                      &&&add(string("bind"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              &Func1::new({
                                                                                                                              let v
                                                                                                                                  =
                                                                                                                                  v.clone();
                                                                                                                              move
                                                                                                                                  |f|
                                                                                                                                  {
                                                                                                                                      let matchValue =
                                                                                                                                          Sharpurs_Prelude::unbox(&&v);
                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&Sharpurs_Prelude::unbox(f),
                                                                                                                                                                       &&&matchValue)
                                                                                                                                  }
                                                                                                                          })),
                                                                                             add(string("Apply0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Semigroup_Last::Data_Semigroup_Last_applyLast()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Semigroup_Last_applicativeLast() -> &dyn Any {
        static Data_Semigroup_Last_applicativeLast: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_Last_applicativeLast.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                             &&&add(string("pure"),
                                                                                                    &&PureScript_Data_Semigroup_Last::Data_Semigroup_Last_Last(),
                                                                                                    add(string("Apply0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Semigroup_Last::Data_Semigroup_Last_applyLast()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>()))))
    }
    pub fn Data_Semigroup_Last_monadLast() -> &dyn Any {
        static Data_Semigroup_Last_monadLast: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_Last_monadLast.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                       &&&add(string("Applicative0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Semigroup_Last::Data_Semigroup_Last_applicativeLast()),
                                                                                              add(string("Bind1"),
                                                                                                  &&Func1::new(move
                                                                                                                   |usd__unused_1|
                                                                                                                   &PureScript_Data_Semigroup_Last::Data_Semigroup_Last_bindLast()),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>()))))
    }
}

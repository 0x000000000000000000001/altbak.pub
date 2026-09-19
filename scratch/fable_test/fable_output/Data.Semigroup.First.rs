pub mod PureScript_Data_Semigroup_First {
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
    pub fn Data_Semigroup_First_First() -> &dyn Any {
        static Data_Semigroup_First_First: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_First.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Semigroup_First_showFirst() -> &dyn Any {
        static Data_Semigroup_First_showFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_showFirst.get_or_init(||
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
                                                                                                                                                                                                            &&&string("(First ")),
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
    pub fn Data_Semigroup_First_semigroupFirst() -> &dyn Any {
        static Data_Semigroup_First_semigroupFirst: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Semigroup_First_semigroupFirst.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_Semigroupusd_Dict(),
                                                                                             &&&add(string("append"),
                                                                                                    &&Func1::new(move
                                                                                                                     |x|
                                                                                                                     &Func1::new({
                                                                                                                                     let x
                                                                                                                                         =
                                                                                                                                         x.clone();
                                                                                                                                     move
                                                                                                                                         |v|
                                                                                                                                         &x
                                                                                                                                 })),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Semigroup_First_ordFirst() -> &dyn Any {
        static Data_Semigroup_First_ordFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_ordFirst.get_or_init(||
                                                      &Func1::new(move
                                                                      |dictOrd|
                                                                      dictOrd.clone()))
    }
    pub fn Data_Semigroup_First_functorFirst() -> &dyn Any {
        static Data_Semigroup_First_functorFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_functorFirst.get_or_init(||
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
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_First::Data_Semigroup_First_First(),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                               &&&v))
                                                                                                                                       }
                                                                                                                               })),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Semigroup_First_eqFirst() -> &dyn Any {
        static Data_Semigroup_First_eqFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_eqFirst.get_or_init(||
                                                     &Func1::new(move |dictEq|
                                                                     dictEq.clone()))
    }
    pub fn Data_Semigroup_First_eq1First() -> &dyn Any {
        static Data_Semigroup_First_eq1First: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_eq1First.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Eq1usd_Dict(),
                                                                                       &&&add(string("eq1"),
                                                                                              &&Func1::new(move
                                                                                                               |dictEq|
                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_First::Data_Semigroup_First_eqFirst(),
                                                                                                                                                                                   dictEq))),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Semigroup_First_ord1First() -> &dyn Any {
        static Data_Semigroup_First_ord1First: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_ord1First.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ord1usd_Dict(),
                                                                                        &&&add(string("compare1"),
                                                                                               &&Func1::new(move
                                                                                                                |dictOrd|
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_First::Data_Semigroup_First_ordFirst(),
                                                                                                                                                                                    dictOrd))),
                                                                                               add(string("Eq10"),
                                                                                                   &&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &PureScript_Data_Semigroup_First::Data_Semigroup_First_eq1First()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Semigroup_First_boundedFirst() -> &dyn Any {
        static Data_Semigroup_First_boundedFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_boundedFirst.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictBounded|
                                                                          dictBounded.clone()))
    }
    pub fn Data_Semigroup_First_applyFirst() -> &dyn Any {
        static Data_Semigroup_First_applyFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_applyFirst.get_or_init(||
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
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup_First::Data_Semigroup_First_First(),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                             &&&matchValue_1))
                                                                                                                                     }
                                                                                                                             })),
                                                                                                add(string("Functor0"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused|
                                                                                                                     &PureScript_Data_Semigroup_First::Data_Semigroup_First_functorFirst()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
    pub fn Data_Semigroup_First_bindFirst() -> &dyn Any {
        static Data_Semigroup_First_bindFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_bindFirst.get_or_init(||
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
                                                                                                                    &PureScript_Data_Semigroup_First::Data_Semigroup_First_applyFirst()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Semigroup_First_applicativeFirst() -> &dyn Any {
        static Data_Semigroup_First_applicativeFirst:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_applicativeFirst.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_Applicativeusd_Dict(),
                                                                                               &&&add(string("pure"),
                                                                                                      &&PureScript_Data_Semigroup_First::Data_Semigroup_First_First(),
                                                                                                      add(string("Apply0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Semigroup_First::Data_Semigroup_First_applyFirst()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Semigroup_First_monadFirst() -> &dyn Any {
        static Data_Semigroup_First_monadFirst: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Semigroup_First_monadFirst.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Monad::Control_Monad_Monadusd_Dict(),
                                                                                         &&&add(string("Applicative0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Semigroup_First::Data_Semigroup_First_applicativeFirst()),
                                                                                                add(string("Bind1"),
                                                                                                    &&Func1::new(move
                                                                                                                     |usd__unused_1|
                                                                                                                     &PureScript_Data_Semigroup_First::Data_Semigroup_First_bindFirst()),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))
    }
}

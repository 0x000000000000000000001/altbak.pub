pub mod PureScript_Data_Interval {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_32f29804::PureScript_Control_Extend;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_637e1ff5::PureScript_Data_Bifoldable;
    use crate::module_c4b10869::PureScript_Data_Bifunctor;
    use crate::module_bdb37be1::PureScript_Data_Bitraversable;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_419ece9e::PureScript_Data_Foldable;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_92875e2a::PureScript_Data_Traversable;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    #[derive(Clone, Debug,)]
    pub enum Data_Interval_Interval {
        Data_Interval_StartEndusd_Ctor(&dyn Any, &dyn Any),
        Data_Interval_DurationEndusd_Ctor(&dyn Any, &dyn Any),
        Data_Interval_StartDurationusd_Ctor(&dyn Any, &dyn Any),
        Data_Interval_DurationOnlyusd_Ctor(&dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Interval::Data_Interval_Interval {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug,)]
    pub enum Data_Interval_RecurringInterval {
        Data_Interval_RecurringIntervalusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Interval::Data_Interval_RecurringInterval {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Interval_showMaybe() -> &dyn Any {
        static Data_Interval_showMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_showMaybe.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_showMaybe(),
                                                                                 &&&PureScript_Data_Show::Data_Show_showInt()))
    }
    pub fn Data_Interval_eqMaybe() -> &dyn Any {
        static Data_Interval_eqMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_eqMaybe.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_eqMaybe(),
                                                                               &&&PureScript_Data_Eq::Data_Eq_eqInt()))
    }
    pub fn Data_Interval_ordMaybe() -> &dyn Any {
        static Data_Interval_ordMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_ordMaybe.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_ordMaybe(),
                                                                                &&&PureScript_Data_Ord::Data_Ord_ordInt()))
    }
    pub fn Data_Interval_StartEnd() -> &dyn Any {
        static Data_Interval_StartEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_StartEnd.get_or_init(||
                                               &Func1::new(move |usd__arg1|
                                                               Func1::new({
                                                                              let usd__arg1
                                                                                  =
                                                                                  usd__arg1.clone();
                                                                              move
                                                                                  |usd__arg2|
                                                                                  &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(usd__arg1,
                                                                                                                                                                                usd__arg2.clone()))
                                                                          })))
    }
    pub fn Data_Interval_DurationEnd() -> &dyn Any {
        static Data_Interval_DurationEnd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_DurationEnd.get_or_init(||
                                                  &Func1::new(move |usd__arg1|
                                                                  Func1::new({
                                                                                 let usd__arg1
                                                                                     =
                                                                                     usd__arg1.clone();
                                                                                 move
                                                                                     |usd__arg2|
                                                                                     &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(usd__arg1,
                                                                                                                                                                                      usd__arg2.clone()))
                                                                             })))
    }
    pub fn Data_Interval_StartDuration() -> &dyn Any {
        static Data_Interval_StartDuration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_StartDuration.get_or_init(||
                                                    &Func1::new(move
                                                                    |usd__arg1|
                                                                    Func1::new({
                                                                                   let usd__arg1
                                                                                       =
                                                                                       usd__arg1.clone();
                                                                                   move
                                                                                       |usd__arg2|
                                                                                       &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(usd__arg1,
                                                                                                                                                                                          usd__arg2.clone()))
                                                                               })))
    }
    pub fn Data_Interval_DurationOnly() -> &dyn Any {
        static Data_Interval_DurationOnly: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_DurationOnly.get_or_init(||
                                                   &Func1::new(move
                                                                   |usd__arg1|
                                                                   &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Interval_RecurringInterval() -> &dyn Any {
        static Data_Interval_RecurringInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_RecurringInterval.get_or_init(||
                                                        &Func1::new(move
                                                                        |usd__arg1|
                                                                        Func1::new({
                                                                                       let usd__arg1
                                                                                           =
                                                                                           usd__arg1.clone();
                                                                                       move
                                                                                           |usd__arg2|
                                                                                           &LrcPtr::new(PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(usd__arg1,
                                                                                                                                                                                                           usd__arg2.clone()))
                                                                                   })))
    }
    pub fn Data_Interval_showInterval() -> &dyn Any {
        static Data_Interval_showInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_showInterval.get_or_init(||
                                                   &Func1::new(move |dictShow|
                                                                   &Func1::new({
                                                                                   let dictShow
                                                                                       =
                                                                                       dictShow.clone();
                                                                                   move
                                                                                       |dictShow1|
                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                        &&&add(string("show"),
                                                                                                                               &&Func1::new({
                                                                                                                                                let dictShow1
                                                                                                                                                    =
                                                                                                                                                    dictShow1.clone();
                                                                                                                                                move
                                                                                                                                                    |v|
                                                                                                                                                    {
                                                                                                                                                        let matchValue:
                                                                                                                                                                LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                                                                        match matchValue.as_ref()
                                                                                                                                                            {
                                                                                                                                                            PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                                matchValue_1_1)
                                                                                                                                                            =>
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&&string("(DurationEnd ")),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                         &&&dictShow),
                                                                                                                                                                                                                                                                                                      &&matchValue_1_0)),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                      &&&string(" ")),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                               &&&dictShow1),
                                                                                                                                                                                                                                                                                                                                                                            &&matchValue_1_1)),
                                                                                                                                                                                                                                                                                                      &&&string(")"))))),
                                                                                                                                                            PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_2_0,
                                                                                                                                                                                                                                                  matchValue_2_1)
                                                                                                                                                            =>
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&&string("(StartDuration ")),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                         &&&dictShow1),
                                                                                                                                                                                                                                                                                                      &&matchValue_2_0)),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                      &&&string(" ")),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                               &&&dictShow),
                                                                                                                                                                                                                                                                                                                                                                            &&matchValue_2_1)),
                                                                                                                                                                                                                                                                                                      &&&string(")"))))),
                                                                                                                                                            PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_3_0)
                                                                                                                                                            =>
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&&string("(DurationOnly ")),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                         &&&dictShow),
                                                                                                                                                                                                                                                                                                      &&matchValue_3_0)),
                                                                                                                                                                                                                                &&&string(")"))),
                                                                                                                                                            PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_0_0,
                                                                                                                                                                                                                                             matchValue_0_1)
                                                                                                                                                            =>
                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&&string("(StartEnd ")),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                         &&&dictShow1),
                                                                                                                                                                                                                                                                                                      &&matchValue_0_0)),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                      &&&string(" ")),
                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                               &&&dictShow1),
                                                                                                                                                                                                                                                                                                                                                                            &&matchValue_0_1)),
                                                                                                                                                                                                                                                                                                      &&&string(")"))))),
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                            }),
                                                                                                                               empty::<string,
                                                                                                                                       &dyn Any>()))
                                                                               })))
    }
    pub fn Data_Interval_showRecurringInterval() -> &dyn Any {
        static Data_Interval_showRecurringInterval: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Interval_showRecurringInterval.get_or_init(||
                                                            &Func1::new(move
                                                                            |dictShow|
                                                                            {
                                                                                let showInterval1 =
                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval::Data_Interval_showInterval(),
                                                                                                                     dictShow);
                                                                                &Func1::new({
                                                                                                let showInterval1
                                                                                                    =
                                                                                                    showInterval1.clone();
                                                                                                move
                                                                                                    |dictShow1|
                                                                                                    {
                                                                                                        let showInterval2 =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&showInterval1,
                                                                                                                                             dictShow1);
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                                                         &&&add(string("show"),
                                                                                                                                                &&Func1::new({
                                                                                                                                                                 let showInterval2
                                                                                                                                                                     =
                                                                                                                                                                     showInterval2.clone();
                                                                                                                                                                 move
                                                                                                                                                                     |v|
                                                                                                                                                                     {
                                                                                                                                                                         let matchValue:
                                                                                                                                                                                 LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                             &&&string("(RecurringInterval ")),
                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Interval::Data_Interval_showMaybe()),
                                                                                                                                                                                                                                                                                                                   &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                          {
                                                                                                                                                                                                                                                                                                                          PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                                                                                                                      })),
                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                   &&&string(" ")),
                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                                                                                                                                            &&&showInterval2),
                                                                                                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                                                                                                PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                                                                                                                   &&&string(")")))))
                                                                                                                                                                     }
                                                                                                                                                             }),
                                                                                                                                                empty::<string,
                                                                                                                                                        &dyn Any>()))
                                                                                                    }
                                                                                            })
                                                                            }))
    }
    pub fn Data_Interval_over() -> &dyn Any {
        static Data_Interval_over: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_over.get_or_init(||
                                           &Func1::new(move |dictFunctor|
                                                           &Func1::new({
                                                                           let dictFunctor
                                                                               =
                                                                               dictFunctor.clone();
                                                                           move
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
                                                                                                       let matchValue_1:
                                                                                                               LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                              &&&dictFunctor),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                                                |usd__arg1|
                                                                                                                                                                                                                                Func1::new({
                                                                                                                                                                                                                                               let usd__arg1
                                                                                                                                                                                                                                                   =
                                                                                                                                                                                                                                                   usd__arg1.clone();
                                                                                                                                                                                                                                               move
                                                                                                                                                                                                                                                   |usd__arg2|
                                                                                                                                                                                                                                                   &LrcPtr::new(PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                   usd__arg2.clone()))
                                                                                                                                                                                                                                           })),
                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                     PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                 })),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                  PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                  =>
                                                                                                                                                                                  x.clone(),
                                                                                                                                                                              }))
                                                                                                   }
                                                                                           })
                                                                       })))
    }
    pub fn Data_Interval_interval() -> &dyn Any {
        static Data_Interval_interval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_interval.get_or_init(||
                                               &Func1::new(move |v|
                                                               &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                    {
                                                                    PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                       x)
                                                                    =>
                                                                    x.clone(),
                                                                }))
    }
    pub fn Data_Interval_foldableInterval_004040() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                         &&&add(string("foldl"),
                                                &&Func1::new(move |v|
                                                                 &Func1::new({
                                                                                 let v
                                                                                     =
                                                                                     v.clone();
                                                                                 move
                                                                                     |v1|
                                                                                     &Func1::new({
                                                                                                     let v1
                                                                                                         =
                                                                                                         v1.clone();
                                                                                                     move
                                                                                                         |v2|
                                                                                                         {
                                                                                                             let matchValue =
                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                             let matchValue_1 =
                                                                                                                 Sharpurs_Prelude::unbox(&&v1);
                                                                                                             let matchValue_2:
                                                                                                                     LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                 Sharpurs_Prelude::unbox(v2);
                                                                                                             match matchValue_2.as_ref()
                                                                                                                 {
                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_2_0_0,
                                                                                                                                                                                                  matchValue_2_0_1)
                                                                                                                 =>
                                                                                                                 {
                                                                                                                     let f =
                                                                                                                         matchValue;
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                               &&&matchValue_1),
                                                                                                                                                                                                                            &&matchValue_2_0_0)),
                                                                                                                                                      &&matchValue_2_0_1)
                                                                                                                 }
                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_2_1_0,
                                                                                                                                                                                                     matchValue_2_1_1)
                                                                                                                 =>
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                     &&&matchValue_1),
                                                                                                                                                  &&matchValue_2_1_1),
                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_2_2_0,
                                                                                                                                                                                                       matchValue_2_2_1)
                                                                                                                 =>
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                     &&&matchValue_1),
                                                                                                                                                  &&matchValue_2_2_0),
                                                                                                                 _
                                                                                                                 =>
                                                                                                                 &matchValue_1,
                                                                                                             }
                                                                                                         }
                                                                                                 })
                                                                             })),
                                                add(string("foldr"),
                                                    &&Func1::new({
                                                                     let Data_Interval_foldableInterval_004040_002d1
                                                                         =
                                                                         Data_Interval_foldableInterval_004040_002d1.clone();
                                                                     move
                                                                         |x_3|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldrDefault(),
                                                                                                                                             &&&Data_Interval_foldableInterval_004040_002d1.Value),
                                                                                                          x_3)
                                                                 }),
                                                    add(string("foldMap"),
                                                        &&Func1::new({
                                                                         let Data_Interval_foldableInterval_004040_002d1
                                                                             =
                                                                             Data_Interval_foldableInterval_004040_002d1.clone();
                                                                         move
                                                                             |dictMonoid|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMapDefaultL(),
                                                                                                                                                 &&&Data_Interval_foldableInterval_004040_002d1.Value),
                                                                                                              dictMonoid)
                                                                     }),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Interval_foldableInterval_004040_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_foldableInterval_004040_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_foldableInterval_004040_002d1.get_or_init(||
                                                                    Lazy(Data_Interval_foldableInterval_004040.clone()))
    }
    pub fn Data_Interval_foldableInterval() -> &dyn Any {
        static Data_Interval_foldableInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_foldableInterval.get_or_init(||
                                                       Data_Interval_foldableInterval_004040_002d1.Value)
    }
    pub fn Data_Interval_foldableRecurringInterval_004043() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_Foldableusd_Dict(),
                                         &&&add(string("foldl"),
                                                &&Func1::new(move |f|
                                                                 &Func1::new({
                                                                                 let f
                                                                                     =
                                                                                     f.clone();
                                                                                 move
                                                                                     |i|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldl(),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_Interval::Data_Interval_foldableInterval()),
                                                                                                                                                                                                                               &&&f),
                                                                                                                                                                                            i)),
                                                                                                                      &&&PureScript_Data_Interval::Data_Interval_interval())
                                                                             })),
                                                add(string("foldr"),
                                                    &&Func1::new(move |f_1|
                                                                     &Func1::new({
                                                                                     let f_1
                                                                                         =
                                                                                         f_1.clone();
                                                                                     move
                                                                                         |i_1|
                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldr(),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_Interval::Data_Interval_foldableInterval()),
                                                                                                                                                                                                                                   &&&f_1),
                                                                                                                                                                                                i_1)),
                                                                                                                          &&&PureScript_Data_Interval::Data_Interval_interval())
                                                                                 })),
                                                    add(string("foldMap"),
                                                        &&Func1::new({
                                                                         let Data_Interval_foldableRecurringInterval_004043_002d1
                                                                             =
                                                                             Data_Interval_foldableRecurringInterval_004043_002d1.clone();
                                                                         move
                                                                             |dictMonoid|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Foldable::Data_Foldable_foldMapDefaultL(),
                                                                                                                                                 &&&Data_Interval_foldableRecurringInterval_004043_002d1.Value),
                                                                                                              dictMonoid)
                                                                     }),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Interval_foldableRecurringInterval_004043_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_foldableRecurringInterval_004043_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_foldableRecurringInterval_004043_002d1.get_or_init(||
                                                                             Lazy(Data_Interval_foldableRecurringInterval_004043.clone()))
    }
    pub fn Data_Interval_foldableRecurringInterval() -> &dyn Any {
        static Data_Interval_foldableRecurringInterval:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_foldableRecurringInterval.get_or_init(||
                                                                Data_Interval_foldableRecurringInterval_004043_002d1.Value)
    }
    pub fn Data_Interval_eqInterval() -> &dyn Any {
        static Data_Interval_eqInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_eqInterval.get_or_init(||
                                                 &Func1::new(move |dictEq|
                                                                 &Func1::new({
                                                                                 let dictEq
                                                                                     =
                                                                                     dictEq.clone();
                                                                                 move
                                                                                     |dictEq1|
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                      &&&add(string("eq"),
                                                                                                                             &&Func1::new({
                                                                                                                                              let dictEq1
                                                                                                                                                  =
                                                                                                                                                  dictEq1.clone();
                                                                                                                                              move
                                                                                                                                                  |x|
                                                                                                                                                  &Func1::new({
                                                                                                                                                                  let x
                                                                                                                                                                      =
                                                                                                                                                                      x.clone();
                                                                                                                                                                  move
                                                                                                                                                                      |y|
                                                                                                                                                                      {
                                                                                                                                                                          let matchValue:
                                                                                                                                                                                  LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                                                              Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                          let matchValue_1:
                                                                                                                                                                                  LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                                                              Sharpurs_Prelude::unbox(y);
                                                                                                                                                                          if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                                                     matchValue_1_1)
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                                                         matchValue_1_1_1)
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                               &&&dictEq),
                                                                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                                                                   PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                            &&&dictEq1),
                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                    x)
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                      &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                             PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                         }))
                                                                                                                                                                              } else {
                                                                                                                                                                                  &false
                                                                                                                                                                              }
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_2_0,
                                                                                                                                                                                                                                                                           matchValue_2_1)
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_1_2_0,
                                                                                                                                                                                                                                                                               matchValue_1_2_1)
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                             &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                                   &&&dictEq1),
                                                                                                                                                                                                                                                                                                                                &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                       PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                       _
                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                       unreachable!(),
                                                                                                                                                                                                                                                                                                                                   }),
                                                                                                                                                                                                                                                                                             &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                    PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                                                })),
                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                &&&dictEq),
                                                                                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                                    PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                                                }),
                                                                                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                 _
                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                 unreachable!(),
                                                                                                                                                                                                                                                             }))
                                                                                                                                                                                  } else {
                                                                                                                                                                                      &false
                                                                                                                                                                                  }
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_3_0)
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_1_3_0)
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                 &&&dictEq),
                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(x)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                  PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(x)
                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                              })
                                                                                                                                                                                      } else {
                                                                                                                                                                                          &false
                                                                                                                                                                                      }
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_1_0_0,
                                                                                                                                                                                                                                                                              matchValue_1_0_1)
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                                       &&&dictEq1),
                                                                                                                                                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                                                           PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                        PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                                                    })),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                    &&&dictEq1),
                                                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                        PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                      x)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                                                                 }))
                                                                                                                                                                                      } else {
                                                                                                                                                                                          &false
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                              }
                                                                                                                                                                          }
                                                                                                                                                                      }
                                                                                                                                                              })
                                                                                                                                          }),
                                                                                                                             empty::<string,
                                                                                                                                     &dyn Any>()))
                                                                             })))
    }
    pub fn Data_Interval_eqRecurringInterval() -> &dyn Any {
        static Data_Interval_eqRecurringInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_eqRecurringInterval.get_or_init(||
                                                          &Func1::new(move
                                                                          |dictEq|
                                                                          {
                                                                              let eqInterval1 =
                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval::Data_Interval_eqInterval(),
                                                                                                                   dictEq);
                                                                              &Func1::new({
                                                                                              let eqInterval1
                                                                                                  =
                                                                                                  eqInterval1.clone();
                                                                                              move
                                                                                                  |dictEq1|
                                                                                                  {
                                                                                                      let eqInterval2 =
                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&&eqInterval1,
                                                                                                                                           dictEq1);
                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                                                       &&&add(string("eq"),
                                                                                                                                              &&Func1::new({
                                                                                                                                                               let eqInterval2
                                                                                                                                                                   =
                                                                                                                                                                   eqInterval2.clone();
                                                                                                                                                               move
                                                                                                                                                                   |x|
                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                   let x
                                                                                                                                                                                       =
                                                                                                                                                                                       x.clone();
                                                                                                                                                                                   move
                                                                                                                                                                                       |y|
                                                                                                                                                                                       {
                                                                                                                                                                                           let matchValue:
                                                                                                                                                                                                   LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                                           let matchValue_1:
                                                                                                                                                                                                   LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                                                               Sharpurs_Prelude::unbox(y);
                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Interval::Data_Interval_eqMaybe()),
                                                                                                                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                                                            PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                                               _)
                                                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                         PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                     })),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                     &&&eqInterval2),
                                                                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                                                         PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                                                     }),
                                                                                                                                                                                                                                                               &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                      PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                  }))
                                                                                                                                                                                       }
                                                                                                                                                                               })
                                                                                                                                                           }),
                                                                                                                                              empty::<string,
                                                                                                                                                      &dyn Any>()))
                                                                                                  }
                                                                                          })
                                                                          }))
    }
    pub fn Data_Interval_ordInterval() -> &dyn Any {
        static Data_Interval_ordInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_ordInterval.get_or_init(||
                                                  &Func1::new(move |dictOrd|
                                                                  {
                                                                      let eqInterval1 =
                                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval::Data_Interval_eqInterval(),
                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                     Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                              &&&Sharpurs_Prelude::Prim_undefined()));
                                                                      &Func1::new({
                                                                                      let dictOrd
                                                                                          =
                                                                                          dictOrd.clone();
                                                                                      let eqInterval1
                                                                                          =
                                                                                          eqInterval1.clone();
                                                                                      move
                                                                                          |dictOrd1|
                                                                                          {
                                                                                              let eqInterval2 =
                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&eqInterval1,
                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                             Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                                      &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                               &&&add(string("compare"),
                                                                                                                                      &&Func1::new({
                                                                                                                                                       let dictOrd1
                                                                                                                                                           =
                                                                                                                                                           dictOrd1.clone();
                                                                                                                                                       move
                                                                                                                                                           |x|
                                                                                                                                                           &Func1::new({
                                                                                                                                                                           let x
                                                                                                                                                                               =
                                                                                                                                                                               x.clone();
                                                                                                                                                                           move
                                                                                                                                                                               |y|
                                                                                                                                                                               {
                                                                                                                                                                                   let matchValue:
                                                                                                                                                                                           LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                                   let matchValue_1:
                                                                                                                                                                                           LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                                                                       Sharpurs_Prelude::unbox(y);
                                                                                                                                                                                   if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_0,
                                                                                                                                                                                                                                                                              matchValue_1_1)
                                                                                                                                                                                          =
                                                                                                                                                                                          matchValue.as_ref()
                                                                                                                                                                                      {
                                                                                                                                                                                       if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_1_0_0,
                                                                                                                                                                                                                                                                               matchValue_1_0_1)
                                                                                                                                                                                              =
                                                                                                                                                                                              matchValue_1.as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                           &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                       } else {
                                                                                                                                                                                           if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                                                                      matchValue_1_1_1)
                                                                                                                                                                                                  =
                                                                                                                                                                                                  matchValue_1.as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                               let matchValue_4:
                                                                                                                                                                                                       LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                   Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                                    &&&dictOrd),
                                                                                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                        PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                        _
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        unreachable!(),
                                                                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                     _
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     unreachable!(),
                                                                                                                                                                                                                                                                 }));
                                                                                                                                                                                               match matchValue_4.as_ref()
                                                                                                                                                                                                   {
                                                                                                                                                                                                   Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                                                                                                                   Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                                                                                                                                   _
                                                                                                                                                                                                   =>
                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                          &&&dictOrd1),
                                                                                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                              PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                              _
                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                              unreachable!(),
                                                                                                                                                                                                                                                                          }),
                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                           PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                           _
                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                           unreachable!(),
                                                                                                                                                                                                                                       }),
                                                                                                                                                                                               }
                                                                                                                                                                                           } else {
                                                                                                                                                                                               if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_1_2_0,
                                                                                                                                                                                                                                                                                            matchValue_1_2_1)
                                                                                                                                                                                                      =
                                                                                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                   &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                               } else {
                                                                                                                                                                                                   &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                               }
                                                                                                                                                                                           }
                                                                                                                                                                                       }
                                                                                                                                                                                   } else {
                                                                                                                                                                                       if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_2_0,
                                                                                                                                                                                                                                                                                    matchValue_2_1)
                                                                                                                                                                                              =
                                                                                                                                                                                              matchValue.as_ref()
                                                                                                                                                                                          {
                                                                                                                                                                                           if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_1_0_0,
                                                                                                                                                                                                                                                                                   matchValue_1_0_1)
                                                                                                                                                                                                  =
                                                                                                                                                                                                  matchValue_1.as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                               &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                           } else {
                                                                                                                                                                                               if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                                                                          matchValue_1_1_1)
                                                                                                                                                                                                      =
                                                                                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                   &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                               } else {
                                                                                                                                                                                                   if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_1_2_0,
                                                                                                                                                                                                                                                                                                matchValue_1_2_1)
                                                                                                                                                                                                          =
                                                                                                                                                                                                          matchValue_1.as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                       let matchValue_5:
                                                                                                                                                                                                               LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                           Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                                            &&&dictOrd1),
                                                                                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                                                                                PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                                                                                                _
                                                                                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                                                                                unreachable!(),
                                                                                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                                                                                      &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                             {
                                                                                                                                                                                                                                                                             PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                   _)
                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                                                                                             _
                                                                                                                                                                                                                                                                             =>
                                                                                                                                                                                                                                                                             unreachable!(),
                                                                                                                                                                                                                                                                         }));
                                                                                                                                                                                                       match matchValue_5.as_ref()
                                                                                                                                                                                                           {
                                                                                                                                                                                                           Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                                                                                                                           Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                                                                                                                                           _
                                                                                                                                                                                                           =>
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                  &&&dictOrd),
                                                                                                                                                                                                                                                                               &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                      PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                            x)
                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                      _
                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                            &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                   PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                         x)
                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                       }
                                                                                                                                                                                                   } else {
                                                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                   }
                                                                                                                                                                                               }
                                                                                                                                                                                           }
                                                                                                                                                                                       } else {
                                                                                                                                                                                           if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_3_0)
                                                                                                                                                                                                  =
                                                                                                                                                                                                  matchValue.as_ref()
                                                                                                                                                                                              {
                                                                                                                                                                                               if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                                                                          matchValue_1_1_1)
                                                                                                                                                                                                      =
                                                                                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                   &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                               } else {
                                                                                                                                                                                                   if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_1_2_0,
                                                                                                                                                                                                                                                                                                matchValue_1_2_1)
                                                                                                                                                                                                          =
                                                                                                                                                                                                          matchValue_1.as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                   } else {
                                                                                                                                                                                                       if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_1_3_0)
                                                                                                                                                                                                              =
                                                                                                                                                                                                              matchValue_1.as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                  &&&dictOrd),
                                                                                                                                                                                                                                                                               &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                      {
                                                                                                                                                                                                                                                                                      PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(x)
                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                      x.clone(),
                                                                                                                                                                                                                                                                                      _
                                                                                                                                                                                                                                                                                      =>
                                                                                                                                                                                                                                                                                      unreachable!(),
                                                                                                                                                                                                                                                                                  }),
                                                                                                                                                                                                                                            &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                   PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(x)
                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                   _
                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                   unreachable!(),
                                                                                                                                                                                                                                               })
                                                                                                                                                                                                       } else {
                                                                                                                                                                                                           &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                       }
                                                                                                                                                                                                   }
                                                                                                                                                                                               }
                                                                                                                                                                                           } else {
                                                                                                                                                                                               if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_1_0_0,
                                                                                                                                                                                                                                                                                       matchValue_1_0_1)
                                                                                                                                                                                                      =
                                                                                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                                                                                  {
                                                                                                                                                                                                   let matchValue_3:
                                                                                                                                                                                                           LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                                        &&&dictOrd1),
                                                                                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                                                                                            PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                                                                                            _
                                                                                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                                                                                            unreachable!(),
                                                                                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                                         PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                                                                                                         _
                                                                                                                                                                                                                                                                         =>
                                                                                                                                                                                                                                                                         unreachable!(),
                                                                                                                                                                                                                                                                     }));
                                                                                                                                                                                                   match matchValue_3.as_ref()
                                                                                                                                                                                                       {
                                                                                                                                                                                                       Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                       =>
                                                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                                                                                                                       Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                                                                                       =>
                                                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                                                                                                                                       _
                                                                                                                                                                                                       =>
                                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                              &&&dictOrd1),
                                                                                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                                  PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                                                  _
                                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                                  unreachable!(),
                                                                                                                                                                                                                                                                              }),
                                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                               PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                                               _
                                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                                               unreachable!(),
                                                                                                                                                                                                                                           }),
                                                                                                                                                                                                   }
                                                                                                                                                                                               } else {
                                                                                                                                                                                                   if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                                                                              matchValue_1_1_1)
                                                                                                                                                                                                          =
                                                                                                                                                                                                          matchValue_1.as_ref()
                                                                                                                                                                                                      {
                                                                                                                                                                                                       &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                   } else {
                                                                                                                                                                                                       if let PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_1_2_0,
                                                                                                                                                                                                                                                                                                    matchValue_1_2_1)
                                                                                                                                                                                                              =
                                                                                                                                                                                                              matchValue_1.as_ref()
                                                                                                                                                                                                          {
                                                                                                                                                                                                           &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                       } else {
                                                                                                                                                                                                           &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                       }
                                                                                                                                                                                                   }
                                                                                                                                                                                               }
                                                                                                                                                                                           }
                                                                                                                                                                                       }
                                                                                                                                                                                   }
                                                                                                                                                                               }
                                                                                                                                                                       })
                                                                                                                                                   }),
                                                                                                                                      add(string("Eq0"),
                                                                                                                                          &&Func1::new({
                                                                                                                                                           let eqInterval2
                                                                                                                                                               =
                                                                                                                                                               eqInterval2.clone();
                                                                                                                                                           move
                                                                                                                                                               |usd__unused|
                                                                                                                                                               &eqInterval2
                                                                                                                                                       }),
                                                                                                                                          empty::<string,
                                                                                                                                                  &dyn Any>())))
                                                                                          }
                                                                                  })
                                                                  }))
    }
    pub fn Data_Interval_ordRecurringInterval() -> &dyn Any {
        static Data_Interval_ordRecurringInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_ordRecurringInterval.get_or_init(||
                                                           &Func1::new(move
                                                                           |dictOrd|
                                                                           {
                                                                               let ordInterval1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval::Data_Interval_ordInterval(),
                                                                                                                    dictOrd);
                                                                               let eqRecurringInterval1 =
                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval::Data_Interval_eqRecurringInterval(),
                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                              Sharpurs_Prelude::unbox(dictOrd)),
                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined()));
                                                                               &Func1::new({
                                                                                               let eqRecurringInterval1
                                                                                                   =
                                                                                                   eqRecurringInterval1.clone();
                                                                                               let ordInterval1
                                                                                                   =
                                                                                                   ordInterval1.clone();
                                                                                               move
                                                                                                   |dictOrd1|
                                                                                                   {
                                                                                                       let ordInterval2 =
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&ordInterval1,
                                                                                                                                            dictOrd1);
                                                                                                       let eqRecurringInterval2 =
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&eqRecurringInterval1,
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Eq0"),
                                                                                                                                                                                      Sharpurs_Prelude::unbox(dictOrd1)),
                                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                                                        &&&add(string("compare"),
                                                                                                                                               &&Func1::new({
                                                                                                                                                                let ordInterval2
                                                                                                                                                                    =
                                                                                                                                                                    ordInterval2.clone();
                                                                                                                                                                move
                                                                                                                                                                    |x|
                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                    let x
                                                                                                                                                                                        =
                                                                                                                                                                                        x.clone();
                                                                                                                                                                                    move
                                                                                                                                                                                        |y|
                                                                                                                                                                                        {
                                                                                                                                                                                            let matchValue:
                                                                                                                                                                                                    LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                                                            let matchValue_1:
                                                                                                                                                                                                    LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(y);
                                                                                                                                                                                            let matchValue_3:
                                                                                                                                                                                                    LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                                                                                Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Interval::Data_Interval_ordMaybe()),
                                                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                     PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                 }),
                                                                                                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                  {
                                                                                                                                                                                                                                                                  PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                                                                              }));
                                                                                                                                                                                            match matchValue_3.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor
                                                                                                                                                                                                =>
                                                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor),
                                                                                                                                                                                                Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor
                                                                                                                                                                                                =>
                                                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor),
                                                                                                                                                                                                _
                                                                                                                                                                                                =>
                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                                                       &&&ordInterval2),
                                                                                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                           PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }),
                                                                                                                                                                                            }
                                                                                                                                                                                        }
                                                                                                                                                                                })
                                                                                                                                                            }),
                                                                                                                                               add(string("Eq0"),
                                                                                                                                                   &&Func1::new({
                                                                                                                                                                    let eqRecurringInterval2
                                                                                                                                                                        =
                                                                                                                                                                        eqRecurringInterval2.clone();
                                                                                                                                                                    move
                                                                                                                                                                        |usd__unused|
                                                                                                                                                                        &eqRecurringInterval2
                                                                                                                                                                }),
                                                                                                                                                   empty::<string,
                                                                                                                                                           &dyn Any>())))
                                                                                                   }
                                                                                           })
                                                                           }))
    }
    pub fn Data_Interval_bifunctorInterval() -> &dyn Any {
        static Data_Interval_bifunctorInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_bifunctorInterval.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                         &&&add(string("bimap"),
                                                                                                &&Func1::new(move
                                                                                                                 |v|
                                                                                                                 &Func1::new({
                                                                                                                                 let v
                                                                                                                                     =
                                                                                                                                     v.clone();
                                                                                                                                 move
                                                                                                                                     |v1|
                                                                                                                                     &Func1::new({
                                                                                                                                                     let v1
                                                                                                                                                         =
                                                                                                                                                         v1.clone();
                                                                                                                                                     move
                                                                                                                                                         |v2|
                                                                                                                                                         {
                                                                                                                                                             let matchValue =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                                                             let matchValue_1 =
                                                                                                                                                                 Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                                             let matchValue_2:
                                                                                                                                                                     LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                                                 Sharpurs_Prelude::unbox(v2);
                                                                                                                                                             match matchValue_2.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_2_1_0,
                                                                                                                                                                                                                                                     matchValue_2_1_1)
                                                                                                                                                                 =>
                                                                                                                                                                 &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                   &&matchValue_2_1_0),
                                                                                                                                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                   &&matchValue_2_1_1))),
                                                                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_2_2_0,
                                                                                                                                                                                                                                                       matchValue_2_2_1)
                                                                                                                                                                 =>
                                                                                                                                                                 &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                                     &&matchValue_2_2_0),
                                                                                                                                                                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                     &&matchValue_2_2_1))),
                                                                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_2_3_0)
                                                                                                                                                                 =>
                                                                                                                                                                 &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                    &&matchValue_2_3_0))),
                                                                                                                                                                 PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_2_0_0,
                                                                                                                                                                                                                                                  matchValue_2_0_1)
                                                                                                                                                                 =>
                                                                                                                                                                 {
                                                                                                                                                                     let f =
                                                                                                                                                                         matchValue_1;
                                                                                                                                                                     &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                    &&matchValue_2_0_0),
                                                                                                                                                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                                    &&matchValue_2_0_1)))
                                                                                                                                                                 }
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                 })
                                                                                                                             })),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Interval_bifunctorRecurringInterval() -> &dyn Any {
        static Data_Interval_bifunctorRecurringInterval:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_bifunctorRecurringInterval.get_or_init(||
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_Bifunctorusd_Dict(),
                                                                                                  &&&add(string("bimap"),
                                                                                                         &&Func1::new(move
                                                                                                                          |f|
                                                                                                                          &Func1::new({
                                                                                                                                          let f
                                                                                                                                              =
                                                                                                                                              f.clone();
                                                                                                                                          move
                                                                                                                                              |g|
                                                                                                                                              &Func1::new({
                                                                                                                                                              let g
                                                                                                                                                                  =
                                                                                                                                                                  g.clone();
                                                                                                                                                              move
                                                                                                                                                                  |v|
                                                                                                                                                                  {
                                                                                                                                                                      let matchValue =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&f);
                                                                                                                                                                      let matchValue_1 =
                                                                                                                                                                          Sharpurs_Prelude::unbox(&&g);
                                                                                                                                                                      let matchValue_2:
                                                                                                                                                                              LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                                          Sharpurs_Prelude::unbox(v);
                                                                                                                                                                      &LrcPtr::new(PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                           PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                       },
                                                                                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                                                                                                                                                                                                                                                                                &&&PureScript_Data_Interval::Data_Interval_bifunctorInterval()),
                                                                                                                                                                                                                                                                                                                                                                                             &&&matchValue),
                                                                                                                                                                                                                                                                                                                                                          &&&matchValue_1),
                                                                                                                                                                                                                                                                                                                       &&&match matchValue_2.as_ref()
                                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                                              PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                                          })))
                                                                                                                                                                  }
                                                                                                                                                          })
                                                                                                                                      })),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_Interval_functorInterval() -> &dyn Any {
        static Data_Interval_functorInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_functorInterval.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                       &&&add(string("map"),
                                                                                              &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifunctor::Data_Bifunctor_bimap(),
                                                                                                                                                                   &&&PureScript_Data_Interval::Data_Interval_bifunctorInterval()),
                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                                                                   &&&PureScript_Control_Category::Control_Category_categoryFn())),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Interval_extendInterval() -> &dyn Any {
        static Data_Interval_extendInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_extendInterval.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                      &&&add(string("extend"),
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
                                                                                                                                              LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                          Sharpurs_Prelude::unbox(v1);
                                                                                                                                      match matchValue_1.as_ref()
                                                                                                                                          {
                                                                                                                                          PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                              matchValue_1_1_1)
                                                                                                                                          =>
                                                                                                                                          &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                            matchValue_1))),
                                                                                                                                          PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_1_2_0,
                                                                                                                                                                                                                                matchValue_1_2_1)
                                                                                                                                          =>
                                                                                                                                          &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                              matchValue_1),
                                                                                                                                                                                                                                             matchValue_1_2_1)),
                                                                                                                                          PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_1_3_0)
                                                                                                                                          =>
                                                                                                                                          &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_1_3_0)),
                                                                                                                                          _
                                                                                                                                          =>
                                                                                                                                          {
                                                                                                                                              let f =
                                                                                                                                                  matchValue;
                                                                                                                                              let a:
                                                                                                                                                      LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                                  matchValue_1.clone();
                                                                                                                                              &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                             &&&a),
                                                                                                                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                             &&&a)))
                                                                                                                                          }
                                                                                                                                      }
                                                                                                                                  }
                                                                                                                          })),
                                                                                             add(string("Functor0"),
                                                                                                 &&Func1::new(move
                                                                                                                  |usd__unused|
                                                                                                                  &PureScript_Data_Interval::Data_Interval_functorInterval()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Interval_functorRecurringInterval() -> &dyn Any {
        static Data_Interval_functorRecurringInterval:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_functorRecurringInterval.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_Functorusd_Dict(),
                                                                                                &&&add(string("map"),
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
                                                                                                                                                let matchValue_1:
                                                                                                                                                        LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                                                &LrcPtr::new(PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                     PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                 },
                                                                                                                                                                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Interval::Data_Interval_functorInterval()),
                                                                                                                                                                                                                                                                                                                                    &&&matchValue),
                                                                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                                                                        PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                           x)
                                                                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                                                                    })))
                                                                                                                                            }
                                                                                                                                    })),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Data_Interval_extendRecurringInterval() -> &dyn Any {
        static Data_Interval_extendRecurringInterval:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_extendRecurringInterval.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_Extendusd_Dict(),
                                                                                               &&&add(string("extend"),
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
                                                                                                                                               let matchValue_1:
                                                                                                                                                       LrcPtr<PureScript_Data_Interval::Data_Interval_RecurringInterval> =
                                                                                                                                                   Sharpurs_Prelude::unbox(v);
                                                                                                                                               &LrcPtr::new(PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                    PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                       _)
                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                },
                                                                                                                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Extend::Control_Extend_extend(),
                                                                                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Interval::Data_Interval_extendInterval()),
                                                                                                                                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_const(),
                                                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                                                                                                                                         &&&matchValue_1))),
                                                                                                                                                                                                                                                                                                &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                       PureScript_Data_Interval::Data_Interval_RecurringInterval::Data_Interval_RecurringIntervalusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                   })))
                                                                                                                                           }
                                                                                                                                   })),
                                                                                                      add(string("Functor0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Interval::Data_Interval_functorRecurringInterval()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Interval_traversableInterval_004066() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                         &&&add(string("traverse"),
                                                &&Func1::new(move
                                                                 |dictApplicative|
                                                                 {
                                                                     let Apply0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                 Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     let Functor0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let Apply0
                                                                                         =
                                                                                         Apply0.clone();
                                                                                     let Functor0
                                                                                         =
                                                                                         Functor0.clone();
                                                                                     let dictApplicative
                                                                                         =
                                                                                         dictApplicative.clone();
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
                                                                                                                 let matchValue_1:
                                                                                                                         LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                     Sharpurs_Prelude::unbox(v1);
                                                                                                                 match matchValue_1.as_ref()
                                                                                                                     {
                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_1_1_0,
                                                                                                                                                                                                         matchValue_1_1_1)
                                                                                                                     =>
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_mapFlipped(),
                                                                                                                                                                                                                            &&&Functor0),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                            &&matchValue_1_1_1)),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&Func1::new(move
                                                                                                                                                                                                           |usd__arg1_1|
                                                                                                                                                                                                           Func1::new({
                                                                                                                                                                                                                          let usd__arg1_1
                                                                                                                                                                                                                              =
                                                                                                                                                                                                                              usd__arg1_1.clone();
                                                                                                                                                                                                                          move
                                                                                                                                                                                                                              |usd__arg2_1|
                                                                                                                                                                                                                              &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                               usd__arg2_1.clone()))
                                                                                                                                                                                                                      })),
                                                                                                                                                                                         &&matchValue_1_1_0)),
                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_1_2_0,
                                                                                                                                                                                                           matchValue_1_2_1)
                                                                                                                     =>
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_mapFlipped(),
                                                                                                                                                                                                                            &&&Functor0),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                            &&matchValue_1_2_0)),
                                                                                                                                                      &&&Func1::new({
                                                                                                                                                                        let matchValue_1
                                                                                                                                                                            =
                                                                                                                                                                            matchValue_1.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |v2|
                                                                                                                                                                            &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(v2.clone(),
                                                                                                                                                                                                                                                                               &match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                                                                    PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                                                                    _
                                                                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                                                                    unreachable!(),
                                                                                                                                                                                                                                                                                }))
                                                                                                                                                                    })),
                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_1_3_0)
                                                                                                                     =>
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                         &&&dictApplicative),
                                                                                                                                                      &&&LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_1_3_0))),
                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_1_0_0,
                                                                                                                                                                                                      matchValue_1_0_1)
                                                                                                                     =>
                                                                                                                     {
                                                                                                                         let f =
                                                                                                                             matchValue;
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                &&&Apply0),
                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                      &&&Functor0),
                                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                                     |usd__arg1|
                                                                                                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                                                                                                    let usd__arg1
                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                        usd__arg1.clone();
                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                        |usd__arg2|
                                                                                                                                                                                                                                                                                                        &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                      usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                }))),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                   &&matchValue_1_0_0))),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                             &&matchValue_1_0_1))
                                                                                                                     }
                                                                                                                 }
                                                                                                             }
                                                                                                     })
                                                                                 })
                                                                 }),
                                                add(string("sequence"),
                                                    &&Func1::new({
                                                                     let Data_Interval_traversableInterval_004066_002d1
                                                                         =
                                                                         Data_Interval_traversableInterval_004066_002d1.clone();
                                                                     move
                                                                         |dictApplicative_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequenceDefault(),
                                                                                                                                             &&&Data_Interval_traversableInterval_004066_002d1.Value),
                                                                                                          dictApplicative_1)
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Interval::Data_Interval_functorInterval()),
                                                        add(string("Foldable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Interval::Data_Interval_foldableInterval()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Interval_traversableInterval_004066_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_traversableInterval_004066_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_traversableInterval_004066_002d1.get_or_init(||
                                                                       Lazy(Data_Interval_traversableInterval_004066.clone()))
    }
    pub fn Data_Interval_traversableInterval() -> &dyn Any {
        static Data_Interval_traversableInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_traversableInterval.get_or_init(||
                                                          Data_Interval_traversableInterval_004066_002d1.Value)
    }
    pub fn Data_Interval_traversableRecurringInterval_004069() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_Traversableusd_Dict(),
                                         &&&add(string("traverse"),
                                                &&Func1::new(move
                                                                 |dictApplicative|
                                                                 {
                                                                     let Functor0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let Functor0
                                                                                         =
                                                                                         Functor0.clone();
                                                                                     let dictApplicative
                                                                                         =
                                                                                         dictApplicative.clone();
                                                                                     move
                                                                                         |f|
                                                                                         &Func1::new({
                                                                                                         let f
                                                                                                             =
                                                                                                             f.clone();
                                                                                                         move
                                                                                                             |i|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval::Data_Interval_over(),
                                                                                                                                                                                                                    &&&Functor0),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_traverse(),
                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Interval::Data_Interval_traversableInterval()),
                                                                                                                                                                                                                                                       &&&dictApplicative),
                                                                                                                                                                                                                    &&&f)),
                                                                                                                                              i)
                                                                                                     })
                                                                                 })
                                                                 }),
                                                add(string("sequence"),
                                                    &&Func1::new({
                                                                     let Data_Interval_traversableRecurringInterval_004069_002d1
                                                                         =
                                                                         Data_Interval_traversableRecurringInterval_004069_002d1.clone();
                                                                     move
                                                                         |dictApplicative_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Traversable::Data_Traversable_sequenceDefault(),
                                                                                                                                             &&&Data_Interval_traversableRecurringInterval_004069_002d1.Value),
                                                                                                          dictApplicative_1)
                                                                 }),
                                                    add(string("Functor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Interval::Data_Interval_functorRecurringInterval()),
                                                        add(string("Foldable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Interval::Data_Interval_foldableRecurringInterval()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Interval_traversableRecurringInterval_004069_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_traversableRecurringInterval_004069_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_traversableRecurringInterval_004069_002d1.get_or_init(||
                                                                                Lazy(Data_Interval_traversableRecurringInterval_004069.clone()))
    }
    pub fn Data_Interval_traversableRecurringInterval() -> &dyn Any {
        static Data_Interval_traversableRecurringInterval:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_traversableRecurringInterval.get_or_init(||
                                                                   Data_Interval_traversableRecurringInterval_004069_002d1.Value)
    }
    pub fn Data_Interval_bifoldableInterval_004072() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                         &&&add(string("bifoldl"),
                                                &&Func1::new(move |v|
                                                                 &Func1::new({
                                                                                 let v
                                                                                     =
                                                                                     v.clone();
                                                                                 move
                                                                                     |v1|
                                                                                     &Func1::new({
                                                                                                     let v1
                                                                                                         =
                                                                                                         v1.clone();
                                                                                                     move
                                                                                                         |v2|
                                                                                                         &Func1::new({
                                                                                                                         let v2
                                                                                                                             =
                                                                                                                             v2.clone();
                                                                                                                         move
                                                                                                                             |v3|
                                                                                                                             {
                                                                                                                                 let matchValue =
                                                                                                                                     Sharpurs_Prelude::unbox(&&v);
                                                                                                                                 let matchValue_1 =
                                                                                                                                     Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                 let matchValue_2 =
                                                                                                                                     Sharpurs_Prelude::unbox(&&v2);
                                                                                                                                 let matchValue_3:
                                                                                                                                         LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                     Sharpurs_Prelude::unbox(v3);
                                                                                                                                 match matchValue_3.as_ref()
                                                                                                                                     {
                                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_3_1_0,
                                                                                                                                                                                                                         matchValue_3_1_1)
                                                                                                                                     =>
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                               &&&matchValue_2),
                                                                                                                                                                                                                                            &&matchValue_3_1_0)),
                                                                                                                                                                      &&matchValue_3_1_1),
                                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_3_2_0,
                                                                                                                                                                                                                           matchValue_3_2_1)
                                                                                                                                     =>
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                               &&&matchValue_2),
                                                                                                                                                                                                                                            &&matchValue_3_2_1)),
                                                                                                                                                                      &&matchValue_3_2_0),
                                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_3_3_0)
                                                                                                                                     =>
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                         &&&matchValue_2),
                                                                                                                                                                      &&matchValue_3_3_0),
                                                                                                                                     PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_3_0_0,
                                                                                                                                                                                                                      matchValue_3_0_1)
                                                                                                                                     =>
                                                                                                                                     {
                                                                                                                                         let f =
                                                                                                                                             matchValue_1;
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&f,
                                                                                                                                                                                                                                                                                   &&&matchValue_2),
                                                                                                                                                                                                                                                &&matchValue_3_0_0)),
                                                                                                                                                                          &&matchValue_3_0_1)
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                             }
                                                                                                                     })
                                                                                                 })
                                                                             })),
                                                add(string("bifoldr"),
                                                    &&Func1::new({
                                                                     let Data_Interval_bifoldableInterval_004072_002d1
                                                                         =
                                                                         Data_Interval_bifoldableInterval_004072_002d1.clone();
                                                                     move
                                                                         |x_3|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldrDefault(),
                                                                                                                                             &&&Data_Interval_bifoldableInterval_004072_002d1.Value),
                                                                                                          x_3)
                                                                 }),
                                                    add(string("bifoldMap"),
                                                        &&Func1::new({
                                                                         let Data_Interval_bifoldableInterval_004072_002d1
                                                                             =
                                                                             Data_Interval_bifoldableInterval_004072_002d1.clone();
                                                                         move
                                                                             |dictMonoid|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMapDefaultL(),
                                                                                                                                                 &&&Data_Interval_bifoldableInterval_004072_002d1.Value),
                                                                                                              dictMonoid)
                                                                     }),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Interval_bifoldableInterval_004072_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_bifoldableInterval_004072_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_bifoldableInterval_004072_002d1.get_or_init(||
                                                                      Lazy(Data_Interval_bifoldableInterval_004072.clone()))
    }
    pub fn Data_Interval_bifoldableInterval() -> &dyn Any {
        static Data_Interval_bifoldableInterval: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_bifoldableInterval.get_or_init(||
                                                         Data_Interval_bifoldableInterval_004072_002d1.Value)
    }
    pub fn Data_Interval_bifoldableRecurringInterval_004075() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_Bifoldableusd_Dict(),
                                         &&&add(string("bifoldl"),
                                                &&Func1::new(move |f|
                                                                 &Func1::new({
                                                                                 let f
                                                                                     =
                                                                                     f.clone();
                                                                                 move
                                                                                     |g|
                                                                                     &Func1::new({
                                                                                                     let g
                                                                                                         =
                                                                                                         g.clone();
                                                                                                     move
                                                                                                         |i|
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldl(),
                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Interval::Data_Interval_bifoldableInterval()),
                                                                                                                                                                                                                                                                                      &&&f),
                                                                                                                                                                                                                                                   &&&g),
                                                                                                                                                                                                                i)),
                                                                                                                                          &&&PureScript_Data_Interval::Data_Interval_interval())
                                                                                                 })
                                                                             })),
                                                add(string("bifoldr"),
                                                    &&Func1::new(move |f_1|
                                                                     &Func1::new({
                                                                                     let f_1
                                                                                         =
                                                                                         f_1.clone();
                                                                                     move
                                                                                         |g_1|
                                                                                         &Func1::new({
                                                                                                         let g_1
                                                                                                             =
                                                                                                             g_1.clone();
                                                                                                         move
                                                                                                             |i_1|
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldr(),
                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Interval::Data_Interval_bifoldableInterval()),
                                                                                                                                                                                                                                                                                          &&&f_1),
                                                                                                                                                                                                                                                       &&&g_1),
                                                                                                                                                                                                                    i_1)),
                                                                                                                                              &&&PureScript_Data_Interval::Data_Interval_interval())
                                                                                                     })
                                                                                 })),
                                                    add(string("bifoldMap"),
                                                        &&Func1::new({
                                                                         let Data_Interval_bifoldableRecurringInterval_004075_002d1
                                                                             =
                                                                             Data_Interval_bifoldableRecurringInterval_004075_002d1.clone();
                                                                         move
                                                                             |dictMonoid|
                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bifoldable::Data_Bifoldable_bifoldMapDefaultL(),
                                                                                                                                                 &&&Data_Interval_bifoldableRecurringInterval_004075_002d1.Value),
                                                                                                              dictMonoid)
                                                                     }),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Interval_bifoldableRecurringInterval_004075_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_bifoldableRecurringInterval_004075_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_bifoldableRecurringInterval_004075_002d1.get_or_init(||
                                                                               Lazy(Data_Interval_bifoldableRecurringInterval_004075.clone()))
    }
    pub fn Data_Interval_bifoldableRecurringInterval() -> &dyn Any {
        static Data_Interval_bifoldableRecurringInterval:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_bifoldableRecurringInterval.get_or_init(||
                                                                  Data_Interval_bifoldableRecurringInterval_004075_002d1.Value)
    }
    pub fn Data_Interval_bitraversableInterval_004078() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                         &&&add(string("bitraverse"),
                                                &&Func1::new(move
                                                                 |dictApplicative|
                                                                 {
                                                                     let Apply0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                 Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     let Functor0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let Apply0
                                                                                         =
                                                                                         Apply0.clone();
                                                                                     let Functor0
                                                                                         =
                                                                                         Functor0.clone();
                                                                                     move
                                                                                         |v|
                                                                                         &Func1::new({
                                                                                                         let v
                                                                                                             =
                                                                                                             v.clone();
                                                                                                         move
                                                                                                             |v1|
                                                                                                             &Func1::new({
                                                                                                                             let v1
                                                                                                                                 =
                                                                                                                                 v1.clone();
                                                                                                                             move
                                                                                                                                 |v2|
                                                                                                                                 {
                                                                                                                                     let matchValue =
                                                                                                                                         Sharpurs_Prelude::unbox(&&v);
                                                                                                                                     let matchValue_1 =
                                                                                                                                         Sharpurs_Prelude::unbox(&&v1);
                                                                                                                                     let matchValue_2:
                                                                                                                                             LrcPtr<PureScript_Data_Interval::Data_Interval_Interval> =
                                                                                                                                         Sharpurs_Prelude::unbox(v2);
                                                                                                                                     match matchValue_2.as_ref()
                                                                                                                                         {
                                                                                                                                         PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(matchValue_2_1_0,
                                                                                                                                                                                                                             matchValue_2_1_1)
                                                                                                                                         =>
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                &&&Apply0),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                      &&&Functor0),
                                                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                                                     |usd__arg1_1|
                                                                                                                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                                                                                                                    let usd__arg1_1
                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                        usd__arg1_1.clone();
                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                        |usd__arg2_1|
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationEndusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                                                                                                         usd__arg2_1.clone()))
                                                                                                                                                                                                                                                                                                                }))),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                                                                                   &&matchValue_2_1_0))),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                             &&matchValue_2_1_1)),
                                                                                                                                         PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(matchValue_2_2_0,
                                                                                                                                                                                                                               matchValue_2_2_1)
                                                                                                                                         =>
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                &&&Apply0),
                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                      &&&Functor0),
                                                                                                                                                                                                                                                                                   &&&Func1::new(move
                                                                                                                                                                                                                                                                                                     |usd__arg1_2|
                                                                                                                                                                                                                                                                                                     Func1::new({
                                                                                                                                                                                                                                                                                                                    let usd__arg1_2
                                                                                                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                                                                                                        usd__arg1_2.clone();
                                                                                                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                                                                                                        |usd__arg2_2|
                                                                                                                                                                                                                                                                                                                        &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartDurationusd_Ctor(usd__arg1_2,
                                                                                                                                                                                                                                                                                                                                                                                                                           usd__arg2_2.clone()))
                                                                                                                                                                                                                                                                                                                }))),
                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue_1,
                                                                                                                                                                                                                                                                                   &&matchValue_2_2_0))),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                             &&matchValue_2_2_1)),
                                                                                                                                         PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(matchValue_2_3_0)
                                                                                                                                         =>
                                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                &&&Functor0),
                                                                                                                                                                                                             &&&Func1::new(move
                                                                                                                                                                                                                               |usd__arg1_3|
                                                                                                                                                                                                                               &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_DurationOnlyusd_Ctor(usd__arg1_3.clone())))),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                             &&matchValue_2_3_0)),
                                                                                                                                         PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(matchValue_2_0_0,
                                                                                                                                                                                                                          matchValue_2_0_1)
                                                                                                                                         =>
                                                                                                                                         {
                                                                                                                                             let r =
                                                                                                                                                 matchValue_1;
                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                    &&&Apply0),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                          &&&Functor0),
                                                                                                                                                                                                                                                                                       &&&Func1::new(move
                                                                                                                                                                                                                                                                                                         |usd__arg1|
                                                                                                                                                                                                                                                                                                         Func1::new({
                                                                                                                                                                                                                                                                                                                        let usd__arg1
                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                            usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                            |usd__arg2|
                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(PureScript_Data_Interval::Data_Interval_Interval::Data_Interval_StartEndusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                                          usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                    }))),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&r,
                                                                                                                                                                                                                                                                                       &&matchValue_2_0_0))),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&r,
                                                                                                                                                                                                                 &&matchValue_2_0_1))
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                         })
                                                                                                     })
                                                                                 })
                                                                 }),
                                                add(string("bisequence"),
                                                    &&Func1::new({
                                                                     let Data_Interval_bitraversableInterval_004078_002d1
                                                                         =
                                                                         Data_Interval_bitraversableInterval_004078_002d1.clone();
                                                                     move
                                                                         |dictApplicative_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bisequenceDefault(),
                                                                                                                                             &&&Data_Interval_bitraversableInterval_004078_002d1.Value),
                                                                                                          dictApplicative_1)
                                                                 }),
                                                    add(string("Bifunctor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Interval::Data_Interval_bifunctorInterval()),
                                                        add(string("Bifoldable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Interval::Data_Interval_bifoldableInterval()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Interval_bitraversableInterval_004078_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_bitraversableInterval_004078_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_bitraversableInterval_004078_002d1.get_or_init(||
                                                                         Lazy(Data_Interval_bitraversableInterval_004078.clone()))
    }
    pub fn Data_Interval_bitraversableInterval() -> &dyn Any {
        static Data_Interval_bitraversableInterval: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Interval_bitraversableInterval.get_or_init(||
                                                            Data_Interval_bitraversableInterval_004078_002d1.Value)
    }
    pub fn Data_Interval_bitraversableRecurringInterval_004081() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_Bitraversableusd_Dict(),
                                         &&&add(string("bitraverse"),
                                                &&Func1::new(move
                                                                 |dictApplicative|
                                                                 {
                                                                     let Functor0 =
                                                                         Sharpurs_Prelude::sharpurs_apply(&&find(string("Functor0"),
                                                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Apply0"),
                                                                                                                                                                                  Sharpurs_Prelude::unbox(dictApplicative)),
                                                                                                                                                                           &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                          &&&Sharpurs_Prelude::Prim_undefined());
                                                                     &Func1::new({
                                                                                     let Functor0
                                                                                         =
                                                                                         Functor0.clone();
                                                                                     let dictApplicative
                                                                                         =
                                                                                         dictApplicative.clone();
                                                                                     move
                                                                                         |l|
                                                                                         &Func1::new({
                                                                                                         let l
                                                                                                             =
                                                                                                             l.clone();
                                                                                                         move
                                                                                                             |r|
                                                                                                             &Func1::new({
                                                                                                                             let r
                                                                                                                                 =
                                                                                                                                 r.clone();
                                                                                                                             move
                                                                                                                                 |i|
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval::Data_Interval_over(),
                                                                                                                                                                                                                                        &&&Functor0),
                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bitraverse(),
                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Interval::Data_Interval_bitraversableInterval()),
                                                                                                                                                                                                                                                                                                              &&&dictApplicative),
                                                                                                                                                                                                                                                                           &&&l),
                                                                                                                                                                                                                                        &&&r)),
                                                                                                                                                                  i)
                                                                                                                         })
                                                                                                     })
                                                                                 })
                                                                 }),
                                                add(string("bisequence"),
                                                    &&Func1::new({
                                                                     let Data_Interval_bitraversableRecurringInterval_004081_002d1
                                                                         =
                                                                         Data_Interval_bitraversableRecurringInterval_004081_002d1.clone();
                                                                     move
                                                                         |dictApplicative_1|
                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bitraversable::Data_Bitraversable_bisequenceDefault(),
                                                                                                                                             &&&Data_Interval_bitraversableRecurringInterval_004081_002d1.Value),
                                                                                                          dictApplicative_1)
                                                                 }),
                                                    add(string("Bifunctor0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Interval::Data_Interval_bifunctorRecurringInterval()),
                                                        add(string("Bifoldable1"),
                                                            &&Func1::new(move
                                                                             |usd__unused_1|
                                                                             &PureScript_Data_Interval::Data_Interval_bifoldableRecurringInterval()),
                                                            empty::<string,
                                                                    &dyn Any>())))))
    }
    pub fn Data_Interval_bitraversableRecurringInterval_004081_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Interval_bitraversableRecurringInterval_004081_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Interval_bitraversableRecurringInterval_004081_002d1.get_or_init(||
                                                                                  Lazy(Data_Interval_bitraversableRecurringInterval_004081.clone()))
    }
    pub fn Data_Interval_bitraversableRecurringInterval() -> &dyn Any {
        static Data_Interval_bitraversableRecurringInterval:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_bitraversableRecurringInterval.get_or_init(||
                                                                     Data_Interval_bitraversableRecurringInterval_004081_002d1.Value)
    }
}

pub mod PureScript_Data_Interval_Duration {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_ed2bf3e0::PureScript_Data_Map_Internal;
    use crate::module_5e99fcdb::PureScript_Data_Monoid;
    use crate::module_146b7611::PureScript_Data_Newtype;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
    pub enum Data_Interval_Duration_DurationComponent {
        Data_Interval_Duration_Secondusd_Ctor,
        Data_Interval_Duration_Minuteusd_Ctor,
        Data_Interval_Duration_Hourusd_Ctor,
        Data_Interval_Duration_Dayusd_Ctor,
        Data_Interval_Duration_Weekusd_Ctor,
        Data_Interval_Duration_Monthusd_Ctor,
        Data_Interval_Duration_Yearusd_Ctor,
    }
    impl core::fmt::Display for
     PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent
     {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Interval_Duration_add() -> &dyn Any {
        static Data_Interval_Duration_add: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_add.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                    &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()))
    }
    pub fn Data_Interval_Duration_Second() -> &dyn Any {
        static Data_Interval_Duration_Second: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Second.get_or_init(||
                                                      &LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor))
    }
    pub fn Data_Interval_Duration_Minute() -> &dyn Any {
        static Data_Interval_Duration_Minute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Minute.get_or_init(||
                                                      &LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor))
    }
    pub fn Data_Interval_Duration_Hour() -> &dyn Any {
        static Data_Interval_Duration_Hour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Hour.get_or_init(||
                                                    &LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor))
    }
    pub fn Data_Interval_Duration_Day() -> &dyn Any {
        static Data_Interval_Duration_Day: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Day.get_or_init(||
                                                   &LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor))
    }
    pub fn Data_Interval_Duration_Week() -> &dyn Any {
        static Data_Interval_Duration_Week: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Week.get_or_init(||
                                                    &LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor))
    }
    pub fn Data_Interval_Duration_Month() -> &dyn Any {
        static Data_Interval_Duration_Month: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Month.get_or_init(||
                                                     &LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor))
    }
    pub fn Data_Interval_Duration_Year() -> &dyn Any {
        static Data_Interval_Duration_Year: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Year.get_or_init(||
                                                    &LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Yearusd_Ctor))
    }
    pub fn Data_Interval_Duration_Duration() -> &dyn Any {
        static Data_Interval_Duration_Duration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_Duration.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Data_Interval_Duration_showDurationComponent() -> &dyn Any {
        static Data_Interval_Duration_showDurationComponent:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_showDurationComponent.get_or_init(||
                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                                      &&&add(string("show"),
                                                                                                             &&Func1::new(move
                                                                                                                              |v|
                                                                                                                              {
                                                                                                                                  let matchValue:
                                                                                                                                          LrcPtr<PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent> =
                                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                                  match matchValue.as_ref()
                                                                                                                                      {
                                                                                                                                      PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                      =>
                                                                                                                                      &string("Second"),
                                                                                                                                      PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                      =>
                                                                                                                                      &string("Hour"),
                                                                                                                                      PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                      =>
                                                                                                                                      &string("Day"),
                                                                                                                                      PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                      =>
                                                                                                                                      &string("Week"),
                                                                                                                                      PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
                                                                                                                                      =>
                                                                                                                                      &string("Month"),
                                                                                                                                      PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Yearusd_Ctor
                                                                                                                                      =>
                                                                                                                                      &string("Year"),
                                                                                                                                      _
                                                                                                                                      =>
                                                                                                                                      &string("Minute"),
                                                                                                                                  }
                                                                                                                              }),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))
    }
    pub fn Data_Interval_Duration_showMap() -> &dyn Any {
        static Data_Interval_Duration_showMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_showMap.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_showMap(),
                                                                                                                           &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_showDurationComponent()),
                                                                                        &&&PureScript_Data_Show::Data_Show_showNumber()))
    }
    pub fn Data_Interval_Duration_showDuration() -> &dyn Any {
        static Data_Interval_Duration_showDuration: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Interval_Duration_showDuration.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                             &&&add(string("show"),
                                                                                                    &&Func1::new(move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let d =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                             &&&string("(Duration ")),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_showMap()),
                                                                                                                                                                                                                                                                   &&&d)),
                                                                                                                                                                                             &&&string(")")))
                                                                                                                     }),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Interval_Duration_newtypeDuration() -> &dyn Any {
        static Data_Interval_Duration_newtypeDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_newtypeDuration.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Newtype::Data_Newtype_Newtypeusd_Dict(),
                                                                                                &&&add(string("Coercible0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &Sharpurs_Prelude::Prim_undefined()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))
    }
    pub fn Data_Interval_Duration_eqDurationComponent() -> &dyn Any {
        static Data_Interval_Duration_eqDurationComponent:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_eqDurationComponent.get_or_init(||
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                                    &&&add(string("eq"),
                                                                                                           &&Func1::new(move
                                                                                                                            |x|
                                                                                                                            &Func1::new({
                                                                                                                                            let x
                                                                                                                                                =
                                                                                                                                                x.clone();
                                                                                                                                            move
                                                                                                                                                |y|
                                                                                                                                                {
                                                                                                                                                    let matchValue:
                                                                                                                                                            LrcPtr<PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent> =
                                                                                                                                                        Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                    let matchValue_1:
                                                                                                                                                            LrcPtr<PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent> =
                                                                                                                                                        Sharpurs_Prelude::unbox(y);
                                                                                                                                                    if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                           =
                                                                                                                                                           matchValue.as_ref()
                                                                                                                                                       {
                                                                                                                                                        if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                               =
                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                           {
                                                                                                                                                            &true
                                                                                                                                                        } else {
                                                                                                                                                            &false
                                                                                                                                                        }
                                                                                                                                                    } else {
                                                                                                                                                        if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                               =
                                                                                                                                                               matchValue.as_ref()
                                                                                                                                                           {
                                                                                                                                                            if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                               {
                                                                                                                                                                &true
                                                                                                                                                            } else {
                                                                                                                                                                &false
                                                                                                                                                            }
                                                                                                                                                        } else {
                                                                                                                                                            if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue.as_ref()
                                                                                                                                                               {
                                                                                                                                                                if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    &true
                                                                                                                                                                } else {
                                                                                                                                                                    &false
                                                                                                                                                                }
                                                                                                                                                            } else {
                                                                                                                                                                if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                           =
                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                        &true
                                                                                                                                                                    } else {
                                                                                                                                                                        &false
                                                                                                                                                                    }
                                                                                                                                                                } else {
                                                                                                                                                                    if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
                                                                                                                                                                           =
                                                                                                                                                                           matchValue.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                        if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
                                                                                                                                                                               =
                                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                            &true
                                                                                                                                                                        } else {
                                                                                                                                                                            &false
                                                                                                                                                                        }
                                                                                                                                                                    } else {
                                                                                                                                                                        if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Yearusd_Ctor
                                                                                                                                                                               =
                                                                                                                                                                               matchValue.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                            if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Yearusd_Ctor
                                                                                                                                                                                   =
                                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                                &true
                                                                                                                                                                            } else {
                                                                                                                                                                                &false
                                                                                                                                                                            }
                                                                                                                                                                        } else {
                                                                                                                                                                            if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                                                                   =
                                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                                &true
                                                                                                                                                                            } else {
                                                                                                                                                                                &false
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    }
                                                                                                                                                                }
                                                                                                                                                            }
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                }
                                                                                                                                        })),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))
    }
    pub fn Data_Interval_Duration_eqMap() -> &dyn Any {
        static Data_Interval_Duration_eqMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_eqMap.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_eqMap(),
                                                                                                                         &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_eqDurationComponent()),
                                                                                      &&&PureScript_Data_Eq::Data_Eq_eqNumber()))
    }
    pub fn Data_Interval_Duration_ordDurationComponent() -> &dyn Any {
        static Data_Interval_Duration_ordDurationComponent:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_ordDurationComponent.get_or_init(||
                                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                                     &&&add(string("compare"),
                                                                                                            &&Func1::new(move
                                                                                                                             |x|
                                                                                                                             &Func1::new({
                                                                                                                                             let x
                                                                                                                                                 =
                                                                                                                                                 x.clone();
                                                                                                                                             move
                                                                                                                                                 |y|
                                                                                                                                                 {
                                                                                                                                                     let matchValue:
                                                                                                                                                             LrcPtr<PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent> =
                                                                                                                                                         Sharpurs_Prelude::unbox(&&x);
                                                                                                                                                     let matchValue_1:
                                                                                                                                                             LrcPtr<PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent> =
                                                                                                                                                         Sharpurs_Prelude::unbox(y);
                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                            =
                                                                                                                                                            matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                         } else {
                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                             } else {
                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                        =
                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                 } else {
                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                            =
                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                     } else {
                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                         } else {
                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
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
                                                                                                                                                     } else {
                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue.as_ref()
                                                                                                                                                            {
                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                             } else {
                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                                        =
                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                 } else {
                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                            =
                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                     } else {
                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                         } else {
                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
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
                                                                                                                                                         } else {
                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    matchValue.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                                                        =
                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                 } else {
                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                                            =
                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                     } else {
                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                         } else {
                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                                        =
                                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                 } else {
                                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
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
                                                                                                                                                             } else {
                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                        =
                                                                                                                                                                        matchValue.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                                                            =
                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                     } else {
                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                         } else {
                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                                        =
                                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                 } else {
                                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                                            =
                                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                     } else {
                                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
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
                                                                                                                                                                 } else {
                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
                                                                                                                                                                            =
                                                                                                                                                                            matchValue.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                         } else {
                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                                        =
                                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                 } else {
                                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                                            =
                                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                     } else {
                                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                                                =
                                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                         } else {
                                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
                                                                                                                                                                                                    =
                                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                             } else {
                                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                             }
                                                                                                                                                                                         }
                                                                                                                                                                                     }
                                                                                                                                                                                 }
                                                                                                                                                                             }
                                                                                                                                                                         }
                                                                                                                                                                     } else {
                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Yearusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                                        =
                                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                 } else {
                                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                                            =
                                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                     } else {
                                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                                                =
                                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                         } else {
                                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
                                                                                                                                                                                                    =
                                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                             } else {
                                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Yearusd_Ctor
                                                                                                                                                                                                        =
                                                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                                                    {
                                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                                 } else {
                                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                 }
                                                                                                                                                                                             }
                                                                                                                                                                                         }
                                                                                                                                                                                     }
                                                                                                                                                                                 }
                                                                                                                                                                             }
                                                                                                                                                                         } else {
                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor
                                                                                                                                                                                        =
                                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                     &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                 } else {
                                                                                                                                                                                     if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor
                                                                                                                                                                                            =
                                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                                        {
                                                                                                                                                                                         &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                     } else {
                                                                                                                                                                                         if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor
                                                                                                                                                                                                =
                                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                             &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                         } else {
                                                                                                                                                                                             if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor
                                                                                                                                                                                                    =
                                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                                {
                                                                                                                                                                                                 &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                             } else {
                                                                                                                                                                                                 if let PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor
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
                                                                                                                                                                     }
                                                                                                                                                                 }
                                                                                                                                                             }
                                                                                                                                                         }
                                                                                                                                                     }
                                                                                                                                                 }
                                                                                                                                         })),
                                                                                                            add(string("Eq0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_Interval_Duration::Data_Interval_Duration_eqDurationComponent()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Interval_Duration_ordMap() -> &dyn Any {
        static Data_Interval_Duration_ordMap: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_ordMap.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_ordMap(),
                                                                                                                          &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_ordDurationComponent()),
                                                                                       &&&PureScript_Data_Ord::Data_Ord_ordNumber()))
    }
    pub fn Data_Interval_Duration_semigroupDuration() -> &dyn Any {
        static Data_Interval_Duration_semigroupDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_semigroupDuration.get_or_init(||
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
                                                                                                                                                  let matchValue_1 =
                                                                                                                                                      Sharpurs_Prelude::unbox(v1);
                                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_Duration(),
                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_unionWith(),
                                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_ordDurationComponent()),
                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_add()),
                                                                                                                                                                                                                                                         &&&matchValue),
                                                                                                                                                                                                                      &&&matchValue_1))
                                                                                                                                              }
                                                                                                                                      })),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))
    }
    pub fn Data_Interval_Duration_monoidDuration() -> &dyn Any {
        static Data_Interval_Duration_monoidDuration:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_monoidDuration.get_or_init(||
                                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Monoid::Data_Monoid_Monoidusd_Dict(),
                                                                                               &&&add(string("mempty"),
                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_Duration(),
                                                                                                                                        &&&PureScript_Data_Map_Internal::Data_Map_Internal_empty()),
                                                                                                      add(string("Semigroup0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Interval_Duration::Data_Interval_Duration_semigroupDuration()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>()))))
    }
    pub fn Data_Interval_Duration_eqDuration() -> &dyn Any {
        static Data_Interval_Duration_eqDuration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_eqDuration.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_Equsd_Dict(),
                                                                                           &&&add(string("eq"),
                                                                                                  &&Func1::new(move
                                                                                                                   |x|
                                                                                                                   &Func1::new({
                                                                                                                                   let x
                                                                                                                                       =
                                                                                                                                       x.clone();
                                                                                                                                   move
                                                                                                                                       |y|
                                                                                                                                       {
                                                                                                                                           let matchValue =
                                                                                                                                               Sharpurs_Prelude::unbox(&&x);
                                                                                                                                           let matchValue_1 =
                                                                                                                                               Sharpurs_Prelude::unbox(y);
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_eqMap()),
                                                                                                                                                                                                               &&&matchValue),
                                                                                                                                                                            &&&matchValue_1)
                                                                                                                                       }
                                                                                                                               })),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_Interval_Duration_ordDuration() -> &dyn Any {
        static Data_Interval_Duration_ordDuration: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_ordDuration.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_Ordusd_Dict(),
                                                                                            &&&add(string("compare"),
                                                                                                   &&Func1::new(move
                                                                                                                    |x|
                                                                                                                    &Func1::new({
                                                                                                                                    let x
                                                                                                                                        =
                                                                                                                                        x.clone();
                                                                                                                                    move
                                                                                                                                        |y|
                                                                                                                                        {
                                                                                                                                            let matchValue =
                                                                                                                                                Sharpurs_Prelude::unbox(&&x);
                                                                                                                                            let matchValue_1 =
                                                                                                                                                Sharpurs_Prelude::unbox(y);
                                                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Interval_Duration::Data_Interval_Duration_ordMap()),
                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                             &&&matchValue_1)
                                                                                                                                        }
                                                                                                                                })),
                                                                                                   add(string("Eq0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Interval_Duration::Data_Interval_Duration_eqDuration()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>()))))
    }
    pub fn Data_Interval_Duration_durationFromComponent() -> &dyn Any {
        static Data_Interval_Duration_durationFromComponent:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_durationFromComponent.get_or_init(||
                                                                     &Func1::new(move
                                                                                     |k|
                                                                                     &Func1::new({
                                                                                                     let k
                                                                                                         =
                                                                                                         k.clone();
                                                                                                     move
                                                                                                         |v|
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_Duration(),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Map_Internal::Data_Map_Internal_singleton(),
                                                                                                                                                                                                                &&&k),
                                                                                                                                                                             v))
                                                                                                 })))
    }
    pub fn Data_Interval_Duration_hour() -> &dyn Any {
        static Data_Interval_Duration_hour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_hour.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                     &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Hourusd_Ctor)))
    }
    pub fn Data_Interval_Duration_millisecond() -> &dyn Any {
        static Data_Interval_Duration_millisecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_millisecond.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                  &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                                                                                                  &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor))),
                                                                                            &&&Func1::new(move
                                                                                                              |v|
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                     &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingNumber()),
                                                                                                                                                                                  v),
                                                                                                                                               &&&1000.0_f64))))
    }
    pub fn Data_Interval_Duration_minute() -> &dyn Any {
        static Data_Interval_Duration_minute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_minute.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                       &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Minuteusd_Ctor)))
    }
    pub fn Data_Interval_Duration_month() -> &dyn Any {
        static Data_Interval_Duration_month: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_month.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                      &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Monthusd_Ctor)))
    }
    pub fn Data_Interval_Duration_second() -> &dyn Any {
        static Data_Interval_Duration_second: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_second.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                       &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Secondusd_Ctor)))
    }
    pub fn Data_Interval_Duration_week() -> &dyn Any {
        static Data_Interval_Duration_week: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_week.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                     &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Weekusd_Ctor)))
    }
    pub fn Data_Interval_Duration_year() -> &dyn Any {
        static Data_Interval_Duration_year: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_year.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                     &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Yearusd_Ctor)))
    }
    pub fn Data_Interval_Duration_day() -> &dyn Any {
        static Data_Interval_Duration_day: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Interval_Duration_day.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Interval_Duration::Data_Interval_Duration_durationFromComponent(),
                                                                                    &&&LrcPtr::new(PureScript_Data_Interval_Duration::Data_Interval_Duration_DurationComponent::Data_Interval_Duration_Dayusd_Ctor)))
    }
}

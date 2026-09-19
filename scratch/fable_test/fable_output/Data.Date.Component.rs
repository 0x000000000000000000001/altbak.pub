pub mod PureScript_Data_Date_Component {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_6a1c5ce6::PureScript_Data_Enum;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
    pub enum Data_Date_Component_Weekday {
        Data_Date_Component_Mondayusd_Ctor,
        Data_Date_Component_Tuesdayusd_Ctor,
        Data_Date_Component_Wednesdayusd_Ctor,
        Data_Date_Component_Thursdayusd_Ctor,
        Data_Date_Component_Fridayusd_Ctor,
        Data_Date_Component_Saturdayusd_Ctor,
        Data_Date_Component_Sundayusd_Ctor,
    }
    impl core::fmt::Display for
     PureScript_Data_Date_Component::Data_Date_Component_Weekday {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
    pub enum Data_Date_Component_Month {
        Data_Date_Component_Januaryusd_Ctor,
        Data_Date_Component_Februaryusd_Ctor,
        Data_Date_Component_Marchusd_Ctor,
        Data_Date_Component_Aprilusd_Ctor,
        Data_Date_Component_Mayusd_Ctor,
        Data_Date_Component_Juneusd_Ctor,
        Data_Date_Component_Julyusd_Ctor,
        Data_Date_Component_Augustusd_Ctor,
        Data_Date_Component_Septemberusd_Ctor,
        Data_Date_Component_Octoberusd_Ctor,
        Data_Date_Component_Novemberusd_Ctor,
        Data_Date_Component_Decemberusd_Ctor,
    }
    impl core::fmt::Display for
     PureScript_Data_Date_Component::Data_Date_Component_Month {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Date_Component_Year() -> &dyn Any {
        static Data_Date_Component_Year: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Year.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Date_Component_Monday() -> &dyn Any {
        static Data_Date_Component_Monday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Monday.get_or_init(||
                                                   &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor))
    }
    pub fn Data_Date_Component_Tuesday() -> &dyn Any {
        static Data_Date_Component_Tuesday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Tuesday.get_or_init(||
                                                    &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor))
    }
    pub fn Data_Date_Component_Wednesday() -> &dyn Any {
        static Data_Date_Component_Wednesday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Wednesday.get_or_init(||
                                                      &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor))
    }
    pub fn Data_Date_Component_Thursday() -> &dyn Any {
        static Data_Date_Component_Thursday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Thursday.get_or_init(||
                                                     &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor))
    }
    pub fn Data_Date_Component_Friday() -> &dyn Any {
        static Data_Date_Component_Friday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Friday.get_or_init(||
                                                   &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor))
    }
    pub fn Data_Date_Component_Saturday() -> &dyn Any {
        static Data_Date_Component_Saturday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Saturday.get_or_init(||
                                                     &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor))
    }
    pub fn Data_Date_Component_Sunday() -> &dyn Any {
        static Data_Date_Component_Sunday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Sunday.get_or_init(||
                                                   &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor))
    }
    pub fn Data_Date_Component_January() -> &dyn Any {
        static Data_Date_Component_January: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_January.get_or_init(||
                                                    &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor))
    }
    pub fn Data_Date_Component_February() -> &dyn Any {
        static Data_Date_Component_February: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_February.get_or_init(||
                                                     &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor))
    }
    pub fn Data_Date_Component_March() -> &dyn Any {
        static Data_Date_Component_March: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_March.get_or_init(||
                                                  &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor))
    }
    pub fn Data_Date_Component_April() -> &dyn Any {
        static Data_Date_Component_April: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_April.get_or_init(||
                                                  &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor))
    }
    pub fn Data_Date_Component_May() -> &dyn Any {
        static Data_Date_Component_May: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_May.get_or_init(||
                                                &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor))
    }
    pub fn Data_Date_Component_June() -> &dyn Any {
        static Data_Date_Component_June: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_June.get_or_init(||
                                                 &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor))
    }
    pub fn Data_Date_Component_July() -> &dyn Any {
        static Data_Date_Component_July: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_July.get_or_init(||
                                                 &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor))
    }
    pub fn Data_Date_Component_August() -> &dyn Any {
        static Data_Date_Component_August: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_August.get_or_init(||
                                                   &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor))
    }
    pub fn Data_Date_Component_September() -> &dyn Any {
        static Data_Date_Component_September: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_September.get_or_init(||
                                                      &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor))
    }
    pub fn Data_Date_Component_October() -> &dyn Any {
        static Data_Date_Component_October: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_October.get_or_init(||
                                                    &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor))
    }
    pub fn Data_Date_Component_November() -> &dyn Any {
        static Data_Date_Component_November: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_November.get_or_init(||
                                                     &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor))
    }
    pub fn Data_Date_Component_December() -> &dyn Any {
        static Data_Date_Component_December: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_December.get_or_init(||
                                                     &LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor))
    }
    pub fn Data_Date_Component_Day() -> &dyn Any {
        static Data_Date_Component_Day: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_Day.get_or_init(||
                                                &Func1::new(move |x|
                                                                x.clone()))
    }
    pub fn Data_Date_Component_showYear() -> &dyn Any {
        static Data_Date_Component_showYear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_showYear.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                      &&&add(string("show"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              {
                                                                                                                  let y =
                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                      &&&string("(Year ")),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                                                                                                            &&&y)),
                                                                                                                                                                                      &&&string(")")))
                                                                                                              }),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Date_Component_showWeekday() -> &dyn Any {
        static Data_Date_Component_showWeekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_showWeekday.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                         &&&add(string("show"),
                                                                                                &&Func1::new(move
                                                                                                                 |v|
                                                                                                                 {
                                                                                                                     let matchValue:
                                                                                                                             LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Weekday> =
                                                                                                                         Sharpurs_Prelude::unbox(v);
                                                                                                                     match matchValue.as_ref()
                                                                                                                         {
                                                                                                                         PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                         =>
                                                                                                                         &string("Tuesday"),
                                                                                                                         PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                         =>
                                                                                                                         &string("Wednesday"),
                                                                                                                         PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                         =>
                                                                                                                         &string("Thursday"),
                                                                                                                         PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                         =>
                                                                                                                         &string("Friday"),
                                                                                                                         PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
                                                                                                                         =>
                                                                                                                         &string("Saturday"),
                                                                                                                         PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor
                                                                                                                         =>
                                                                                                                         &string("Sunday"),
                                                                                                                         _
                                                                                                                         =>
                                                                                                                         &string("Monday"),
                                                                                                                     }
                                                                                                                 }),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))
    }
    pub fn Data_Date_Component_showMonth() -> &dyn Any {
        static Data_Date_Component_showMonth: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_showMonth.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                       &&&add(string("show"),
                                                                                              &&Func1::new(move
                                                                                                               |v|
                                                                                                               {
                                                                                                                   let matchValue:
                                                                                                                           LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Month> =
                                                                                                                       Sharpurs_Prelude::unbox(v);
                                                                                                                   match matchValue.as_ref()
                                                                                                                       {
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("February"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("March"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("April"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("May"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("June"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("July"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("August"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("September"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("October"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("November"),
                                                                                                                       PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor
                                                                                                                       =>
                                                                                                                       &string("December"),
                                                                                                                       _
                                                                                                                       =>
                                                                                                                       &string("January"),
                                                                                                                   }
                                                                                                               }),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))
    }
    pub fn Data_Date_Component_showDay() -> &dyn Any {
        static Data_Date_Component_showDay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_showDay.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                     &&&add(string("show"),
                                                                                            &&Func1::new(move
                                                                                                             |v|
                                                                                                             {
                                                                                                                 let d =
                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                        &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                     &&&string("(Day ")),
                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                                                                                                           &&&d)),
                                                                                                                                                                                     &&&string(")")))
                                                                                                             }),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_Date_Component_ordYear() -> &dyn Any {
        static Data_Date_Component_ordYear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_ordYear.get_or_init(||
                                                    &PureScript_Data_Ord::Data_Ord_ordInt())
    }
    pub fn Data_Date_Component_ordDay() -> &dyn Any {
        static Data_Date_Component_ordDay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_ordDay.get_or_init(||
                                                   &PureScript_Data_Ord::Data_Ord_ordInt())
    }
    pub fn Data_Date_Component_eqYear() -> &dyn Any {
        static Data_Date_Component_eqYear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_eqYear.get_or_init(||
                                                   &PureScript_Data_Eq::Data_Eq_eqInt())
    }
    pub fn Data_Date_Component_eqWeekday() -> &dyn Any {
        static Data_Date_Component_eqWeekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_eqWeekday.get_or_init(||
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
                                                                                                                                               LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Weekday> =
                                                                                                                                           Sharpurs_Prelude::unbox(&&x);
                                                                                                                                       let matchValue_1:
                                                                                                                                               LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Weekday> =
                                                                                                                                           Sharpurs_Prelude::unbox(y);
                                                                                                                                       if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                              =
                                                                                                                                              matchValue.as_ref()
                                                                                                                                          {
                                                                                                                                           if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                  =
                                                                                                                                                  matchValue_1.as_ref()
                                                                                                                                              {
                                                                                                                                               &true
                                                                                                                                           } else {
                                                                                                                                               &false
                                                                                                                                           }
                                                                                                                                       } else {
                                                                                                                                           if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                  =
                                                                                                                                                  matchValue.as_ref()
                                                                                                                                              {
                                                                                                                                               if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                      =
                                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                                  {
                                                                                                                                                   &true
                                                                                                                                               } else {
                                                                                                                                                   &false
                                                                                                                                               }
                                                                                                                                           } else {
                                                                                                                                               if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                      =
                                                                                                                                                      matchValue.as_ref()
                                                                                                                                                  {
                                                                                                                                                   if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                          =
                                                                                                                                                          matchValue_1.as_ref()
                                                                                                                                                      {
                                                                                                                                                       &true
                                                                                                                                                   } else {
                                                                                                                                                       &false
                                                                                                                                                   }
                                                                                                                                               } else {
                                                                                                                                                   if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                          =
                                                                                                                                                          matchValue.as_ref()
                                                                                                                                                      {
                                                                                                                                                       if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                              =
                                                                                                                                                              matchValue_1.as_ref()
                                                                                                                                                          {
                                                                                                                                                           &true
                                                                                                                                                       } else {
                                                                                                                                                           &false
                                                                                                                                                       }
                                                                                                                                                   } else {
                                                                                                                                                       if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
                                                                                                                                                              =
                                                                                                                                                              matchValue.as_ref()
                                                                                                                                                          {
                                                                                                                                                           if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
                                                                                                                                                                  =
                                                                                                                                                                  matchValue_1.as_ref()
                                                                                                                                                              {
                                                                                                                                                               &true
                                                                                                                                                           } else {
                                                                                                                                                               &false
                                                                                                                                                           }
                                                                                                                                                       } else {
                                                                                                                                                           if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor
                                                                                                                                                                  =
                                                                                                                                                                  matchValue.as_ref()
                                                                                                                                                              {
                                                                                                                                                               if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor
                                                                                                                                                                      =
                                                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                                                  {
                                                                                                                                                                   &true
                                                                                                                                                               } else {
                                                                                                                                                                   &false
                                                                                                                                                               }
                                                                                                                                                           } else {
                                                                                                                                                               if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor
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
    pub fn Data_Date_Component_ordWeekday() -> &dyn Any {
        static Data_Date_Component_ordWeekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_ordWeekday.get_or_init(||
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
                                                                                                                                                LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Weekday> =
                                                                                                                                            Sharpurs_Prelude::unbox(&&x);
                                                                                                                                        let matchValue_1:
                                                                                                                                                LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Weekday> =
                                                                                                                                            Sharpurs_Prelude::unbox(y);
                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                               =
                                                                                                                                               matchValue.as_ref()
                                                                                                                                           {
                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor
                                                                                                                                                   =
                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                               {
                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                            } else {
                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                       =
                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                } else {
                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                           =
                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                       {
                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                    } else {
                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                               =
                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                           {
                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                        } else {
                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                               {
                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                            } else {
                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
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
                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                   =
                                                                                                                                                   matchValue.as_ref()
                                                                                                                                               {
                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor
                                                                                                                                                       =
                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                   {
                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                } else {
                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                           =
                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                       {
                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                    } else {
                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                               =
                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                           {
                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                        } else {
                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                               {
                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                            } else {
                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                } else {
                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
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
                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                       =
                                                                                                                                                       matchValue.as_ref()
                                                                                                                                                   {
                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor
                                                                                                                                                           =
                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                       {
                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                    } else {
                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                               =
                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                           {
                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                        } else {
                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                               {
                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                            } else {
                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                } else {
                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                                           =
                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                    } else {
                                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
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
                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                           =
                                                                                                                                                           matchValue.as_ref()
                                                                                                                                                       {
                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor
                                                                                                                                                               =
                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                           {
                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                        } else {
                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                               {
                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                            } else {
                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                } else {
                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                                           =
                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                    } else {
                                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                                               =
                                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                        } else {
                                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
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
                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
                                                                                                                                                               =
                                                                                                                                                               matchValue.as_ref()
                                                                                                                                                           {
                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                               {
                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                            } else {
                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                } else {
                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                                           =
                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                    } else {
                                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                                               =
                                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                        } else {
                                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                                                   =
                                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                            } else {
                                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
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
                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor
                                                                                                                                                                   =
                                                                                                                                                                   matchValue.as_ref()
                                                                                                                                                               {
                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                } else {
                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                                           =
                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                    } else {
                                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                                               =
                                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                        } else {
                                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                                                   =
                                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                            } else {
                                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
                                                                                                                                                                                       =
                                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                } else {
                                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor
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
                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor
                                                                                                                                                                       =
                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                   {
                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                } else {
                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                                                                                                           =
                                                                                                                                                                           matchValue_1.as_ref()
                                                                                                                                                                       {
                                                                                                                                                                        &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                    } else {
                                                                                                                                                                        if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                                                                                                               =
                                                                                                                                                                               matchValue_1.as_ref()
                                                                                                                                                                           {
                                                                                                                                                                            &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                        } else {
                                                                                                                                                                            if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                                                                                                                   =
                                                                                                                                                                                   matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                                &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                            } else {
                                                                                                                                                                                if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                                                                                                                       =
                                                                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                                                                   {
                                                                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                } else {
                                                                                                                                                                                    if let PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
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
                                                                                                                    &PureScript_Data_Date_Component::Data_Date_Component_eqWeekday()),
                                                                                                   empty::<string,
                                                                                                           &dyn Any>()))))
    }
    pub fn Data_Date_Component_eqMonth() -> &dyn Any {
        static Data_Date_Component_eqMonth: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_eqMonth.get_or_init(||
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
                                                                                                                                             LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Month> =
                                                                                                                                         Sharpurs_Prelude::unbox(&&x);
                                                                                                                                     let matchValue_1:
                                                                                                                                             LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Month> =
                                                                                                                                         Sharpurs_Prelude::unbox(y);
                                                                                                                                     if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                            =
                                                                                                                                            matchValue.as_ref()
                                                                                                                                        {
                                                                                                                                         if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                =
                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                            {
                                                                                                                                             &true
                                                                                                                                         } else {
                                                                                                                                             &false
                                                                                                                                         }
                                                                                                                                     } else {
                                                                                                                                         if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                =
                                                                                                                                                matchValue.as_ref()
                                                                                                                                            {
                                                                                                                                             if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                    =
                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                {
                                                                                                                                                 &true
                                                                                                                                             } else {
                                                                                                                                                 &false
                                                                                                                                             }
                                                                                                                                         } else {
                                                                                                                                             if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                    =
                                                                                                                                                    matchValue.as_ref()
                                                                                                                                                {
                                                                                                                                                 if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                        =
                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                    {
                                                                                                                                                     &true
                                                                                                                                                 } else {
                                                                                                                                                     &false
                                                                                                                                                 }
                                                                                                                                             } else {
                                                                                                                                                 if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                        =
                                                                                                                                                        matchValue.as_ref()
                                                                                                                                                    {
                                                                                                                                                     if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                            =
                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                        {
                                                                                                                                                         &true
                                                                                                                                                     } else {
                                                                                                                                                         &false
                                                                                                                                                     }
                                                                                                                                                 } else {
                                                                                                                                                     if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                            =
                                                                                                                                                            matchValue.as_ref()
                                                                                                                                                        {
                                                                                                                                                         if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                            {
                                                                                                                                                             &true
                                                                                                                                                         } else {
                                                                                                                                                             &false
                                                                                                                                                         }
                                                                                                                                                     } else {
                                                                                                                                                         if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                =
                                                                                                                                                                matchValue.as_ref()
                                                                                                                                                            {
                                                                                                                                                             if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 &true
                                                                                                                                                             } else {
                                                                                                                                                                 &false
                                                                                                                                                             }
                                                                                                                                                         } else {
                                                                                                                                                             if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                    =
                                                                                                                                                                    matchValue.as_ref()
                                                                                                                                                                {
                                                                                                                                                                 if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                        =
                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     &true
                                                                                                                                                                 } else {
                                                                                                                                                                     &false
                                                                                                                                                                 }
                                                                                                                                                             } else {
                                                                                                                                                                 if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                        =
                                                                                                                                                                        matchValue.as_ref()
                                                                                                                                                                    {
                                                                                                                                                                     if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                            =
                                                                                                                                                                            matchValue_1.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                         &true
                                                                                                                                                                     } else {
                                                                                                                                                                         &false
                                                                                                                                                                     }
                                                                                                                                                                 } else {
                                                                                                                                                                     if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                            =
                                                                                                                                                                            matchValue.as_ref()
                                                                                                                                                                        {
                                                                                                                                                                         if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue_1.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             &true
                                                                                                                                                                         } else {
                                                                                                                                                                             &false
                                                                                                                                                                         }
                                                                                                                                                                     } else {
                                                                                                                                                                         if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
                                                                                                                                                                                =
                                                                                                                                                                                matchValue.as_ref()
                                                                                                                                                                            {
                                                                                                                                                                             if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 &true
                                                                                                                                                                             } else {
                                                                                                                                                                                 &false
                                                                                                                                                                             }
                                                                                                                                                                         } else {
                                                                                                                                                                             if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor
                                                                                                                                                                                    =
                                                                                                                                                                                    matchValue.as_ref()
                                                                                                                                                                                {
                                                                                                                                                                                 if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor
                                                                                                                                                                                        =
                                                                                                                                                                                        matchValue_1.as_ref()
                                                                                                                                                                                    {
                                                                                                                                                                                     &true
                                                                                                                                                                                 } else {
                                                                                                                                                                                     &false
                                                                                                                                                                                 }
                                                                                                                                                                             } else {
                                                                                                                                                                                 if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
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
                                                                                                                                                 }
                                                                                                                                             }
                                                                                                                                         }
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                         })),
                                                                                            empty::<string,
                                                                                                    &dyn Any>())))
    }
    pub fn Data_Date_Component_ordMonth() -> &dyn Any {
        static Data_Date_Component_ordMonth: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_ordMonth.get_or_init(||
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
                                                                                                                                              LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Month> =
                                                                                                                                          Sharpurs_Prelude::unbox(&&x);
                                                                                                                                      let matchValue_1:
                                                                                                                                              LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Month> =
                                                                                                                                          Sharpurs_Prelude::unbox(y);
                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                             =
                                                                                                                                             matchValue.as_ref()
                                                                                                                                         {
                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                 =
                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                             {
                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                          } else {
                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                     =
                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                 {
                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                              } else {
                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                         =
                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                  } else {
                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                             =
                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                         {
                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                      } else {
                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                 =
                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                             {
                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                          } else {
                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                     =
                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                              } else {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                      } else {
                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                 =
                                                                                                                                                 matchValue.as_ref()
                                                                                                                                             {
                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                     =
                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                 {
                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                              } else {
                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                         =
                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                  } else {
                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                             =
                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                         {
                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                      } else {
                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                 =
                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                             {
                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                          } else {
                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                     =
                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                              } else {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                          } else {
                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                     =
                                                                                                                                                     matchValue.as_ref()
                                                                                                                                                 {
                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                         =
                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                     {
                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                  } else {
                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                             =
                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                         {
                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                      } else {
                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                 =
                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                             {
                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                          } else {
                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                     =
                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                              } else {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                              } else {
                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                         =
                                                                                                                                                         matchValue.as_ref()
                                                                                                                                                     {
                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                             =
                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                         {
                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                      } else {
                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                 =
                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                             {
                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                          } else {
                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                     =
                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                              } else {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                                  } else {
                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                             =
                                                                                                                                                             matchValue.as_ref()
                                                                                                                                                         {
                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                                 =
                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                             {
                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                          } else {
                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                     =
                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                              } else {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                                      } else {
                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                 =
                                                                                                                                                                 matchValue.as_ref()
                                                                                                                                                             {
                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                                     =
                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                              } else {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                                          } else {
                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                     =
                                                                                                                                                                     matchValue.as_ref()
                                                                                                                                                                 {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                             =
                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                      } else {
                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                                              } else {
                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                         =
                                                                                                                                                                         matchValue.as_ref()
                                                                                                                                                                     {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                             =
                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                                      } else {
                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                          } else {
                                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                                                  } else {
                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                             =
                                                                                                                                                                             matchValue.as_ref()
                                                                                                                                                                         {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                                             =
                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                      } else {
                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                          } else {
                                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                                 {
                                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                                              } else {
                                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                                                      } else {
                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
                                                                                                                                                                                 =
                                                                                                                                                                                 matchValue.as_ref()
                                                                                                                                                                             {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                                             =
                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                      } else {
                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                          } else {
                                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                                 {
                                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                              } else {
                                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                                         =
                                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                                  } else {
                                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                                                                                              }
                                                                                                                                                                                          }
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                              }
                                                                                                                                                                          } else {
                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor
                                                                                                                                                                                     =
                                                                                                                                                                                     matchValue.as_ref()
                                                                                                                                                                                 {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                                             =
                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                      } else {
                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                          } else {
                                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                                 {
                                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                              } else {
                                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                                         =
                                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                                  } else {
                                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
                                                                                                                                                                                                                             =
                                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                                                                                                      } else {
                                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor
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
                                                                                                                                                                                                  }
                                                                                                                                                                                              }
                                                                                                                                                                                          }
                                                                                                                                                                                      }
                                                                                                                                                                                  }
                                                                                                                                                                              } else {
                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor
                                                                                                                                                                                         =
                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                     {
                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                                                                  } else {
                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                                                                                                             =
                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                      } else {
                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                                                                                                                 =
                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                             {
                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                          } else {
                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                                                                                                                     =
                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                              } else {
                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                                                                                                                         =
                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                  } else {
                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                                                                                                             =
                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                      } else {
                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                                                                                                                 =
                                                                                                                                                                                                                 matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                              &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                          } else {
                                                                                                                                                                                                              if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                                                                                                                     =
                                                                                                                                                                                                                     matchValue_1.as_ref()
                                                                                                                                                                                                                 {
                                                                                                                                                                                                                  &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                              } else {
                                                                                                                                                                                                                  if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                                                                                                                         =
                                                                                                                                                                                                                         matchValue_1.as_ref()
                                                                                                                                                                                                                     {
                                                                                                                                                                                                                      &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                                  } else {
                                                                                                                                                                                                                      if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                                                                                                             =
                                                                                                                                                                                                                             matchValue_1.as_ref()
                                                                                                                                                                                                                         {
                                                                                                                                                                                                                          &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                                                                                                      } else {
                                                                                                                                                                                                                          if let PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
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
                                                                                                                  &PureScript_Data_Date_Component::Data_Date_Component_eqMonth()),
                                                                                                 empty::<string,
                                                                                                         &dyn Any>()))))
    }
    pub fn Data_Date_Component_eqDay() -> &dyn Any {
        static Data_Date_Component_eqDay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_eqDay.get_or_init(||
                                                  &PureScript_Data_Eq::Data_Eq_eqInt())
    }
    pub fn Data_Date_Component_boundedYear() -> &dyn Any {
        static Data_Date_Component_boundedYear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_boundedYear.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                         &&&add(string("bottom"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Component::Data_Date_Component_Year(),
                                                                                                                                  &&&1_i32),
                                                                                                add(string("top"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Component::Data_Date_Component_Year(),
                                                                                                                                      &&&9999_i32),
                                                                                                    add(string("Ord0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Date_Component::Data_Date_Component_ordYear()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))))
    }
    pub fn Data_Date_Component_boundedWeekday() -> &dyn Any {
        static Data_Date_Component_boundedWeekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_boundedWeekday.get_or_init(||
                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                            &&&add(string("bottom"),
                                                                                                   &&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor),
                                                                                                   add(string("top"),
                                                                                                       &&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor),
                                                                                                       add(string("Ord0"),
                                                                                                           &&Func1::new(move
                                                                                                                            |usd__unused|
                                                                                                                            &PureScript_Data_Date_Component::Data_Date_Component_ordWeekday()),
                                                                                                           empty::<string,
                                                                                                                   &dyn Any>())))))
    }
    pub fn Data_Date_Component_boundedMonth() -> &dyn Any {
        static Data_Date_Component_boundedMonth: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_boundedMonth.get_or_init(||
                                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                          &&&add(string("bottom"),
                                                                                                 &&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor),
                                                                                                 add(string("top"),
                                                                                                     &&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor),
                                                                                                     add(string("Ord0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Date_Component::Data_Date_Component_ordMonth()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>())))))
    }
    pub fn Data_Date_Component_boundedEnumYear_0040102() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&9999_i32),
                                                add(string("toEnum"),
                                                    &&Func1::new(move |n|
                                                                     {
                                                                         let matchValue =
                                                                             Sharpurs_Prelude::unbox(n);
                                                                         if {
                                                                                let n1 =
                                                                                    matchValue;
                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                   &&&n1),
                                                                                                                                                                                                                &&&1_i32)),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                &&&n1),
                                                                                                                                                                             &&&9999_i32)))
                                                                            }
                                                                            {
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Component::Data_Date_Component_Year(),
                                                                                                                                                                     &&&matchValue)))
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                             } else {
                                                                                 panic!("{}",
                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Date.Component.fs"),
                                  Data1: 102_i32,
                                  Data2: 281_i32,}).get_Message(),)
                                                                             }
                                                                         }
                                                                     }),
                                                    add(string("fromEnum"),
                                                        &&Func1::new(move |v|
                                                                         &Sharpurs_Prelude::unbox(v)),
                                                        add(string("Bounded0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Date_Component::Data_Date_Component_boundedYear()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Date_Component_enumYear_0040103_002d1
                                                                                     =
                                                                                     Data_Date_Component_enumYear_0040103_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Date_Component_enumYear_0040103_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Date_Component_boundedEnumYear_0040102_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_boundedEnumYear_0040102_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_boundedEnumYear_0040102_002d1.get_or_init(||
                                                                          Lazy(Data_Date_Component_boundedEnumYear_0040102.clone()))
    }
    pub fn Data_Date_Component_enumYear_0040103() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumYear_0040102_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumYear_0040102_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumYear_0040102_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumYear_0040102_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Date_Component::Data_Date_Component_ordYear()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Date_Component_enumYear_0040103_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_enumYear_0040103_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_enumYear_0040103_002d1.get_or_init(||
                                                                   Lazy(Data_Date_Component_enumYear_0040103.clone()))
    }
    pub fn Data_Date_Component_boundedEnumYear() -> &dyn Any {
        static Data_Date_Component_boundedEnumYear: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Date_Component_boundedEnumYear.get_or_init(||
                                                            Data_Date_Component_boundedEnumYear_0040102_002d1.Value)
    }
    pub fn Data_Date_Component_enumYear() -> &dyn Any {
        static Data_Date_Component_enumYear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_enumYear.get_or_init(||
                                                     Data_Date_Component_enumYear_0040103_002d1.Value)
    }
    pub fn Data_Date_Component_boundedEnumWeekday_0040106() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&7_i32),
                                                add(string("toEnum"),
                                                    &&Func1::new(move |v|
                                                                     {
                                                                         let matchValue =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(1_i32,
                                                                                                                         &matchValue)
                                                                             {
                                                                             0_i32
                                                                             =>
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Mondayusd_Ctor))),
                                                                             _
                                                                             =>
                                                                             match &Sharpurs_Prelude::_007cLitInt_007c__007c(2_i32,
                                                                                                                             &matchValue)
                                                                                 {
                                                                                 0_i32
                                                                                 =>
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor))),
                                                                                 _
                                                                                 =>
                                                                                 match &Sharpurs_Prelude::_007cLitInt_007c__007c(3_i32,
                                                                                                                                 &matchValue)
                                                                                     {
                                                                                     0_i32
                                                                                     =>
                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor))),
                                                                                     _
                                                                                     =>
                                                                                     match &Sharpurs_Prelude::_007cLitInt_007c__007c(4_i32,
                                                                                                                                     &matchValue)
                                                                                         {
                                                                                         0_i32
                                                                                         =>
                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor))),
                                                                                         _
                                                                                         =>
                                                                                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(5_i32,
                                                                                                                                         &matchValue)
                                                                                             {
                                                                                             0_i32
                                                                                             =>
                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor))),
                                                                                             _
                                                                                             =>
                                                                                             match &Sharpurs_Prelude::_007cLitInt_007c__007c(6_i32,
                                                                                                                                             &matchValue)
                                                                                                 {
                                                                                                 0_i32
                                                                                                 =>
                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor))),
                                                                                                 _
                                                                                                 =>
                                                                                                 match &Sharpurs_Prelude::_007cLitInt_007c__007c(7_i32,
                                                                                                                                                 &matchValue)
                                                                                                     {
                                                                                                     0_i32
                                                                                                     =>
                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor))),
                                                                                                     _
                                                                                                     =>
                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                 },
                                                                                             },
                                                                                         },
                                                                                     },
                                                                                 },
                                                                             },
                                                                         }
                                                                     }),
                                                    add(string("fromEnum"),
                                                        &&Func1::new(move
                                                                         |v_1|
                                                                         {
                                                                             let matchValue_1:
                                                                                     LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Weekday> =
                                                                                 Sharpurs_Prelude::unbox(v_1);
                                                                             match matchValue_1.as_ref()
                                                                                 {
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Tuesdayusd_Ctor
                                                                                 =>
                                                                                 &2_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Wednesdayusd_Ctor
                                                                                 =>
                                                                                 &3_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Thursdayusd_Ctor
                                                                                 =>
                                                                                 &4_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Fridayusd_Ctor
                                                                                 =>
                                                                                 &5_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Saturdayusd_Ctor
                                                                                 =>
                                                                                 &6_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Weekday::Data_Date_Component_Sundayusd_Ctor
                                                                                 =>
                                                                                 &7_i32,
                                                                                 _
                                                                                 =>
                                                                                 &1_i32,
                                                                             }
                                                                         }),
                                                        add(string("Bounded0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Date_Component::Data_Date_Component_boundedWeekday()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Date_Component_enumWeekday_0040107_002d1
                                                                                     =
                                                                                     Data_Date_Component_enumWeekday_0040107_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Date_Component_enumWeekday_0040107_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Date_Component_boundedEnumWeekday_0040106_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_boundedEnumWeekday_0040106_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_boundedEnumWeekday_0040106_002d1.get_or_init(||
                                                                             Lazy(Data_Date_Component_boundedEnumWeekday_0040106.clone()))
    }
    pub fn Data_Date_Component_enumWeekday_0040107() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumWeekday_0040106_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumWeekday_0040106_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumWeekday_0040106_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumWeekday_0040106_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Date_Component::Data_Date_Component_ordWeekday()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Date_Component_enumWeekday_0040107_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_enumWeekday_0040107_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_enumWeekday_0040107_002d1.get_or_init(||
                                                                      Lazy(Data_Date_Component_enumWeekday_0040107.clone()))
    }
    pub fn Data_Date_Component_boundedEnumWeekday() -> &dyn Any {
        static Data_Date_Component_boundedEnumWeekday:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_boundedEnumWeekday.get_or_init(||
                                                               Data_Date_Component_boundedEnumWeekday_0040106_002d1.Value)
    }
    pub fn Data_Date_Component_enumWeekday() -> &dyn Any {
        static Data_Date_Component_enumWeekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_enumWeekday.get_or_init(||
                                                        Data_Date_Component_enumWeekday_0040107_002d1.Value)
    }
    pub fn Data_Date_Component_boundedEnumMonth_0040110() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&12_i32),
                                                add(string("toEnum"),
                                                    &&Func1::new(move |v|
                                                                     {
                                                                         let matchValue =
                                                                             Sharpurs_Prelude::unbox(v);
                                                                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(1_i32,
                                                                                                                         &matchValue)
                                                                             {
                                                                             0_i32
                                                                             =>
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor))),
                                                                             _
                                                                             =>
                                                                             match &Sharpurs_Prelude::_007cLitInt_007c__007c(2_i32,
                                                                                                                             &matchValue)
                                                                                 {
                                                                                 0_i32
                                                                                 =>
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor))),
                                                                                 _
                                                                                 =>
                                                                                 match &Sharpurs_Prelude::_007cLitInt_007c__007c(3_i32,
                                                                                                                                 &matchValue)
                                                                                     {
                                                                                     0_i32
                                                                                     =>
                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor))),
                                                                                     _
                                                                                     =>
                                                                                     match &Sharpurs_Prelude::_007cLitInt_007c__007c(4_i32,
                                                                                                                                     &matchValue)
                                                                                         {
                                                                                         0_i32
                                                                                         =>
                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor))),
                                                                                         _
                                                                                         =>
                                                                                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(5_i32,
                                                                                                                                         &matchValue)
                                                                                             {
                                                                                             0_i32
                                                                                             =>
                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor))),
                                                                                             _
                                                                                             =>
                                                                                             match &Sharpurs_Prelude::_007cLitInt_007c__007c(6_i32,
                                                                                                                                             &matchValue)
                                                                                                 {
                                                                                                 0_i32
                                                                                                 =>
                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor))),
                                                                                                 _
                                                                                                 =>
                                                                                                 match &Sharpurs_Prelude::_007cLitInt_007c__007c(7_i32,
                                                                                                                                                 &matchValue)
                                                                                                     {
                                                                                                     0_i32
                                                                                                     =>
                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor))),
                                                                                                     _
                                                                                                     =>
                                                                                                     match &Sharpurs_Prelude::_007cLitInt_007c__007c(8_i32,
                                                                                                                                                     &matchValue)
                                                                                                         {
                                                                                                         0_i32
                                                                                                         =>
                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor))),
                                                                                                         _
                                                                                                         =>
                                                                                                         match &Sharpurs_Prelude::_007cLitInt_007c__007c(9_i32,
                                                                                                                                                         &matchValue)
                                                                                                             {
                                                                                                             0_i32
                                                                                                             =>
                                                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor))),
                                                                                                             _
                                                                                                             =>
                                                                                                             match &Sharpurs_Prelude::_007cLitInt_007c__007c(10_i32,
                                                                                                                                                             &matchValue)
                                                                                                                 {
                                                                                                                 0_i32
                                                                                                                 =>
                                                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor))),
                                                                                                                 _
                                                                                                                 =>
                                                                                                                 match &Sharpurs_Prelude::_007cLitInt_007c__007c(11_i32,
                                                                                                                                                                 &matchValue)
                                                                                                                     {
                                                                                                                     0_i32
                                                                                                                     =>
                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor))),
                                                                                                                     _
                                                                                                                     =>
                                                                                                                     match &Sharpurs_Prelude::_007cLitInt_007c__007c(12_i32,
                                                                                                                                                                     &matchValue)
                                                                                                                         {
                                                                                                                         0_i32
                                                                                                                         =>
                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&LrcPtr::new(PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor))),
                                                                                                                         _
                                                                                                                         =>
                                                                                                                         &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                     },
                                                                                                                 },
                                                                                                             },
                                                                                                         },
                                                                                                     },
                                                                                                 },
                                                                                             },
                                                                                         },
                                                                                     },
                                                                                 },
                                                                             },
                                                                         }
                                                                     }),
                                                    add(string("fromEnum"),
                                                        &&Func1::new(move
                                                                         |v_1|
                                                                         {
                                                                             let matchValue_1:
                                                                                     LrcPtr<PureScript_Data_Date_Component::Data_Date_Component_Month> =
                                                                                 Sharpurs_Prelude::unbox(v_1);
                                                                             match matchValue_1.as_ref()
                                                                                 {
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                 =>
                                                                                 &2_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                 =>
                                                                                 &3_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                 =>
                                                                                 &4_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                 =>
                                                                                 &5_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                 =>
                                                                                 &6_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                 =>
                                                                                 &7_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                 =>
                                                                                 &8_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                 =>
                                                                                 &9_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                 =>
                                                                                 &10_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
                                                                                 =>
                                                                                 &11_i32,
                                                                                 PureScript_Data_Date_Component::Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor
                                                                                 =>
                                                                                 &12_i32,
                                                                                 _
                                                                                 =>
                                                                                 &1_i32,
                                                                             }
                                                                         }),
                                                        add(string("Bounded0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Date_Component::Data_Date_Component_boundedMonth()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Date_Component_enumMonth_0040111_002d1
                                                                                     =
                                                                                     Data_Date_Component_enumMonth_0040111_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Date_Component_enumMonth_0040111_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Date_Component_boundedEnumMonth_0040110_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_boundedEnumMonth_0040110_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_boundedEnumMonth_0040110_002d1.get_or_init(||
                                                                           Lazy(Data_Date_Component_boundedEnumMonth_0040110.clone()))
    }
    pub fn Data_Date_Component_enumMonth_0040111() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumMonth_0040110_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumMonth_0040110_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumMonth_0040110_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumMonth_0040110_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Date_Component::Data_Date_Component_ordMonth()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Date_Component_enumMonth_0040111_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_enumMonth_0040111_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_enumMonth_0040111_002d1.get_or_init(||
                                                                    Lazy(Data_Date_Component_enumMonth_0040111.clone()))
    }
    pub fn Data_Date_Component_boundedEnumMonth() -> &dyn Any {
        static Data_Date_Component_boundedEnumMonth: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Date_Component_boundedEnumMonth.get_or_init(||
                                                             Data_Date_Component_boundedEnumMonth_0040110_002d1.Value)
    }
    pub fn Data_Date_Component_enumMonth() -> &dyn Any {
        static Data_Date_Component_enumMonth: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_enumMonth.get_or_init(||
                                                      Data_Date_Component_enumMonth_0040111_002d1.Value)
    }
    pub fn Data_Date_Component_boundedDay() -> &dyn Any {
        static Data_Date_Component_boundedDay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_boundedDay.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                        &&&add(string("bottom"),
                                                                                               &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Component::Data_Date_Component_Day(),
                                                                                                                                 &&&1_i32),
                                                                                               add(string("top"),
                                                                                                   &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Component::Data_Date_Component_Day(),
                                                                                                                                     &&&31_i32),
                                                                                                   add(string("Ord0"),
                                                                                                       &&Func1::new(move
                                                                                                                        |usd__unused|
                                                                                                                        &PureScript_Data_Date_Component::Data_Date_Component_ordDay()),
                                                                                                       empty::<string,
                                                                                                               &dyn Any>())))))
    }
    pub fn Data_Date_Component_boundedEnumDay_0040116() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&31_i32),
                                                add(string("toEnum"),
                                                    &&Func1::new(move |n|
                                                                     {
                                                                         let matchValue =
                                                                             Sharpurs_Prelude::unbox(n);
                                                                         if {
                                                                                let n1 =
                                                                                    matchValue;
                                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                                                   &&&n1),
                                                                                                                                                                                                                &&&1_i32)),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                &&&n1),
                                                                                                                                                                             &&&31_i32)))
                                                                            }
                                                                            {
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date_Component::Data_Date_Component_Day(),
                                                                                                                                                                     &&&matchValue)))
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                             } else {
                                                                                 panic!("{}",
                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Date.Component.fs"),
                                  Data1: 116_i32,
                                  Data2: 278_i32,}).get_Message(),)
                                                                             }
                                                                         }
                                                                     }),
                                                    add(string("fromEnum"),
                                                        &&Func1::new(move |v|
                                                                         &Sharpurs_Prelude::unbox(v)),
                                                        add(string("Bounded0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Date_Component::Data_Date_Component_boundedDay()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Date_Component_enumDay_0040117_002d1
                                                                                     =
                                                                                     Data_Date_Component_enumDay_0040117_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Date_Component_enumDay_0040117_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Date_Component_boundedEnumDay_0040116_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_boundedEnumDay_0040116_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_boundedEnumDay_0040116_002d1.get_or_init(||
                                                                         Lazy(Data_Date_Component_boundedEnumDay_0040116.clone()))
    }
    pub fn Data_Date_Component_enumDay_0040117() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumDay_0040116_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Date_Component_boundedEnumDay_0040116_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumDay_0040116_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Date_Component_boundedEnumDay_0040116_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Date_Component::Data_Date_Component_ordDay()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Date_Component_enumDay_0040117_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Date_Component_enumDay_0040117_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Date_Component_enumDay_0040117_002d1.get_or_init(||
                                                                  Lazy(Data_Date_Component_enumDay_0040117.clone()))
    }
    pub fn Data_Date_Component_boundedEnumDay() -> &dyn Any {
        static Data_Date_Component_boundedEnumDay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_boundedEnumDay.get_or_init(||
                                                           Data_Date_Component_boundedEnumDay_0040116_002d1.Value)
    }
    pub fn Data_Date_Component_enumDay() -> &dyn Any {
        static Data_Date_Component_enumDay: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_Component_enumDay.get_or_init(||
                                                    Data_Date_Component_enumDay_0040117_002d1.Value)
    }
}

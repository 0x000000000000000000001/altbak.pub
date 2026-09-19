pub mod PureScript_Data_DateTime_Instant {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_b6aac8e0::PureScript_Data_Date_Component;
    use crate::module_d96ec2c1::PureScript_Data_Date;
    use crate::module_867835b4::PureScript_Data_DateTime::Data_DateTime_DateTime;
    use crate::module_6a1c5ce6::PureScript_Data_Enum;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_f62dae61::PureScript_Data_Time_Component;
    use crate::module_c4b34810::PureScript_Data_Time_Duration;
    use crate::module_f8599c00::PureScript_Data_Time;
    use crate::module_f8599c00::PureScript_Data_Time::Data_Time_Time;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_DateTime_Instant_fromDateTimeImpl() -> &dyn Any {
        static Data_DateTime_Instant_fromDateTimeImpl:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_fromDateTimeImpl.get_or_init(||
                                                               &Func1::new(move
                                                                               |arg0|
                                                                               &Func1::new(move
                                                                                               |arg1|
                                                                                               &Func1::new(move
                                                                                                               |arg2|
                                                                                                               &Func1::new(move
                                                                                                                               |arg3|
                                                                                                                               &Func1::new(move
                                                                                                                                               |arg4|
                                                                                                                                               &Func1::new(move
                                                                                                                                                               |arg5|
                                                                                                                                                               &Func1::new(move
                                                                                                                                                                               |arg6|
                                                                                                                                                                               panic!("{}",
                                                                                                                                                                                      1_i32.get_Message(),)))))))))
    }
    pub fn Data_DateTime_Instant_toDateTimeImpl() -> &dyn Any {
        static Data_DateTime_Instant_toDateTimeImpl: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_DateTime_Instant_toDateTimeImpl.get_or_init(||
                                                             &Func1::new(move
                                                                             |arg0|
                                                                             &Func1::new(move
                                                                                             |arg1|
                                                                                             panic!("{}",
                                                                                                    1_i32.get_Message(),))))
    }
    pub fn Data_DateTime_Instant_bottom() -> &dyn Any {
        static Data_DateTime_Instant_bottom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_bottom.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                      &&&PureScript_Data_Time_Component::Data_Time_Component_boundedHour()))
    }
    pub fn Data_DateTime_Instant_bottom1() -> &dyn Any {
        static Data_DateTime_Instant_bottom1: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_bottom1.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                       &&&PureScript_Data_Time_Component::Data_Time_Component_boundedMinute()))
    }
    pub fn Data_DateTime_Instant_bottom2() -> &dyn Any {
        static Data_DateTime_Instant_bottom2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_bottom2.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                       &&&PureScript_Data_Time_Component::Data_Time_Component_boundedSecond()))
    }
    pub fn Data_DateTime_Instant_bottom3() -> &dyn Any {
        static Data_DateTime_Instant_bottom3: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_bottom3.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                       &&&PureScript_Data_Time_Component::Data_Time_Component_boundedMillisecond()))
    }
    pub fn Data_DateTime_Instant_Instant() -> &dyn Any {
        static Data_DateTime_Instant_Instant: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_Instant.get_or_init(||
                                                      &Func1::new(move |x|
                                                                      x.clone()))
    }
    pub fn Data_DateTime_Instant_unInstant() -> &dyn Any {
        static Data_DateTime_Instant_unInstant: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_unInstant.get_or_init(||
                                                        &Func1::new(move |v|
                                                                        &Sharpurs_Prelude::unbox(v)))
    }
    pub fn Data_DateTime_Instant_toDateTime() -> &dyn Any {
        static Data_DateTime_Instant_toDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_toDateTime.get_or_init(||
                                                         {
                                                             let mkDateTime =
                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                  &&&Func1::new(move
                                                                                                                    |usd__unused|
                                                                                                                    &Func1::new(move
                                                                                                                                    |y|
                                                                                                                                    &Func1::new({
                                                                                                                                                    let y
                                                                                                                                                        =
                                                                                                                                                        y.clone();
                                                                                                                                                    move
                                                                                                                                                        |mo|
                                                                                                                                                        &Func1::new({
                                                                                                                                                                        let mo
                                                                                                                                                                            =
                                                                                                                                                                            mo.clone();
                                                                                                                                                                        move
                                                                                                                                                                            |d|
                                                                                                                                                                            &Func1::new({
                                                                                                                                                                                            let d
                                                                                                                                                                                                =
                                                                                                                                                                                                d.clone();
                                                                                                                                                                                            move
                                                                                                                                                                                                |h|
                                                                                                                                                                                                &Func1::new({
                                                                                                                                                                                                                let h
                                                                                                                                                                                                                    =
                                                                                                                                                                                                                    h.clone();
                                                                                                                                                                                                                move
                                                                                                                                                                                                                    |mi|
                                                                                                                                                                                                                    &Func1::new({
                                                                                                                                                                                                                                    let mi
                                                                                                                                                                                                                                        =
                                                                                                                                                                                                                                        mi.clone();
                                                                                                                                                                                                                                    move
                                                                                                                                                                                                                                        |s|
                                                                                                                                                                                                                                        &Func1::new({
                                                                                                                                                                                                                                                        let s
                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                            s.clone();
                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                            |ms|
                                                                                                                                                                                                                                                            &LrcPtr::new(Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_canonicalDate(),
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&y),
                                                                                                                                                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                          &&&mo))),
                                                                                                                                                                                                                                                                                                                                                                 &&&d),
                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(Data_Time_Time::Data_Time_Timeusd_Ctor(&h,
                                                                                                                                                                                                                                                                                                                                                                                    &mi,
                                                                                                                                                                                                                                                                                                                                                                                    &s,
                                                                                                                                                                                                                                                                                                                                                                                    ms.clone()))))
                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                })
                                                                                                                                                                                                            })
                                                                                                                                                                                        })
                                                                                                                                                                    })
                                                                                                                                                }))));
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_toDateTimeImpl(),
                                                                                              &&&mkDateTime)
                                                         })
    }
    pub fn Data_DateTime_Instant_showInstant() -> &dyn Any {
        static Data_DateTime_Instant_showInstant: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_showInstant.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                           &&&add(string("show"),
                                                                                                  &&Func1::new(move
                                                                                                                   |v|
                                                                                                                   {
                                                                                                                       let ms =
                                                                                                                           Sharpurs_Prelude::unbox(v);
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                           &&&string("(Instant ")),
                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Time_Duration::Data_Time_Duration_showMilliseconds()),
                                                                                                                                                                                                                                                                 &&&ms)),
                                                                                                                                                                                           &&&string(")")))
                                                                                                                   }),
                                                                                                  empty::<string,
                                                                                                          &dyn Any>())))
    }
    pub fn Data_DateTime_Instant_ordDateTime() -> &dyn Any {
        static Data_DateTime_Instant_ordDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_ordDateTime.get_or_init(||
                                                          &PureScript_Data_Time_Duration::Data_Time_Duration_ordMilliseconds())
    }
    pub fn Data_DateTime_Instant_instant() -> &dyn Any {
        static Data_DateTime_Instant_instant: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_instant.get_or_init(||
                                                      &Func1::new(move |v|
                                                                      {
                                                                          let matchValue =
                                                                              Sharpurs_Prelude::unbox(v);
                                                                          if {
                                                                                 let n =
                                                                                     matchValue;
                                                                                 Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                 &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                                                                    &&&n),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                                                                                                                                                                    &&&8639977881600000.0_f64))),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                    &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                                 &&&n),
                                                                                                                                                                              &&&8639977881599999.0_f64)))
                                                                             }
                                                                             {
                                                                              &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_Instant(),
                                                                                                                                                                      &&&matchValue)))
                                                                          } else {
                                                                              if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                 {
                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                              } else {
                                                                                  panic!("{}",
                                                                                         LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.DateTime.Instant.fs"),
                                  Data1: 29_i32,
                                  Data2: 66_i32,}).get_Message(),)
                                                                              }
                                                                          }
                                                                      }))
    }
    pub fn Data_DateTime_Instant_fromDateTime() -> &dyn Any {
        static Data_DateTime_Instant_fromDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_fromDateTime.get_or_init(||
                                                           &Func1::new(move
                                                                           |v|
                                                                           {
                                                                               let matchValue:
                                                                                       LrcPtr<Data_DateTime_DateTime> =
                                                                                   Sharpurs_Prelude::unbox(v);
                                                                               let t =
                                                                                   match matchValue.as_ref()
                                                                                       {
                                                                                       Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                              x)
                                                                                       =>
                                                                                       x.clone(),
                                                                                   };
                                                                               let d =
                                                                                   match matchValue.as_ref()
                                                                                       {
                                                                                       Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                              _)
                                                                                       =>
                                                                                       x.clone(),
                                                                                   };
                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn7(),
                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_fromDateTimeImpl()),
                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_year(),
                                                                                                                                                                                                                                                                                                                                                                     &&&d)),
                                                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_month(),
                                                                                                                                                                                                                                                                                                                                                                     &&&d))),
                                                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_day(),
                                                                                                                                                                                                                                                                                               &&&d)),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_hour(),
                                                                                                                                                                                                                                                            &&&t)),
                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_minute(),
                                                                                                                                                                                                                         &&&t)),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_second(),
                                                                                                                                                                                      &&&t)),
                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_millisecond(),
                                                                                                                                                   &&&t))
                                                                           }))
    }
    pub fn Data_DateTime_Instant_fromDate() -> &dyn Any {
        static Data_DateTime_Instant_fromDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_fromDate.get_or_init(||
                                                       &Func1::new(move |d|
                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn7(),
                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_fromDateTimeImpl()),
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_year(),
                                                                                                                                                                                                                                                                                                                                                             d)),
                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_month(),
                                                                                                                                                                                                                                                                                                                                                             d))),
                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_day(),
                                                                                                                                                                                                                                                                                       d)),
                                                                                                                                                                                                                 &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_bottom()),
                                                                                                                                                                              &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_bottom1()),
                                                                                                                                           &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_bottom2()),
                                                                                                        &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_bottom3())))
    }
    pub fn Data_DateTime_Instant_eqDateTime() -> &dyn Any {
        static Data_DateTime_Instant_eqDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_eqDateTime.get_or_init(||
                                                         &PureScript_Data_Time_Duration::Data_Time_Duration_eqMilliseconds())
    }
    pub fn Data_DateTime_Instant_diff() -> &dyn Any {
        static Data_DateTime_Instant_diff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_Instant_diff.get_or_init(||
                                                   &Func1::new(move
                                                                   |dictDuration|
                                                                   &Func1::new({
                                                                                   let dictDuration
                                                                                       =
                                                                                       dictDuration.clone();
                                                                                   move
                                                                                       |dt1|
                                                                                       &Func1::new({
                                                                                                       let dt1
                                                                                                           =
                                                                                                           dt1.clone();
                                                                                                       move
                                                                                                           |dt2|
                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_toDuration(),
                                                                                                                                                                               &&&dictDuration),
                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                     &&&PureScript_Data_Time_Duration::Data_Time_Duration_semigroupMilliseconds()),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_unInstant(),
                                                                                                                                                                                                                                                     &&&dt1)),
                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_negateDuration(),
                                                                                                                                                                                                                                                     &&&PureScript_Data_Time_Duration::Data_Time_Duration_durationMilliseconds()),
                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_unInstant(),
                                                                                                                                                                                                                                                     dt2))))
                                                                                                   })
                                                                               })))
    }
    pub fn Data_DateTime_Instant_boundedInstant() -> &dyn Any {
        static Data_DateTime_Instant_boundedInstant: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_DateTime_Instant_boundedInstant.get_or_init(||
                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                              &&&add(string("bottom"),
                                                                                                     &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_Instant(),
                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds(),
                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_negate(),
                                                                                                                                                                                                                                                &&&PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                                                                                                                             &&&8639977881600000.0_f64))),
                                                                                                     add(string("top"),
                                                                                                         &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_Instant(),
                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_Milliseconds(),
                                                                                                                                                                              &&&8639977881599999.0_f64)),
                                                                                                         add(string("Ord0"),
                                                                                                             &&Func1::new(move
                                                                                                                              |usd__unused|
                                                                                                                              &PureScript_Data_DateTime_Instant::Data_DateTime_Instant_ordDateTime()),
                                                                                                             empty::<string,
                                                                                                                     &dyn Any>())))))
    }
}

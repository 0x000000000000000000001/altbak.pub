pub mod PureScript_Data_DateTime {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_b6aac8e0::PureScript_Data_Date_Component;
    use crate::module_d96ec2c1::PureScript_Data_Date;
    use crate::module_6a1c5ce6::PureScript_Data_Enum;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_f62dae61::PureScript_Data_Time_Component;
    use crate::module_c4b34810::PureScript_Data_Time_Duration;
    use crate::module_f8599c00::PureScript_Data_Time;
    use crate::module_f8599c00::PureScript_Data_Time::Data_Time_Time;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_DateTime_adjustImpl() -> &dyn Any {
        static Data_DateTime_adjustImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_adjustImpl.get_or_init(||
                                                 &Func1::new(move |arg0|
                                                                 &Func1::new(move
                                                                                 |arg1|
                                                                                 &Func1::new(move
                                                                                                 |arg2|
                                                                                                 &Func1::new(move
                                                                                                                 |arg3|
                                                                                                                 &Func1::new(move
                                                                                                                                 |arg4|
                                                                                                                                 panic!("{}",
                                                                                                                                        1_i32.get_Message(),)))))))
    }
    pub fn Data_DateTime_calcDiff() -> &dyn Any {
        static Data_DateTime_calcDiff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_calcDiff.get_or_init(||
                                               &Func1::new(move |arg0|
                                                               &Func1::new(move
                                                                               |arg1|
                                                                               panic!("{}",
                                                                                      1_i32.get_Message(),))))
    }
    #[derive(Clone, Debug,)]
    pub enum Data_DateTime_DateTime {
        Data_DateTime_DateTimeusd_Ctor(&dyn Any, &dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_DateTime::Data_DateTime_DateTime {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_DateTime_DateTime() -> &dyn Any {
        static Data_DateTime_DateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_DateTime.get_or_init(||
                                               &Func1::new(move |usd__arg1|
                                                               Func1::new({
                                                                              let usd__arg1
                                                                                  =
                                                                                  usd__arg1.clone();
                                                                              move
                                                                                  |usd__arg2|
                                                                                  &LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(usd__arg1,
                                                                                                                                                                                usd__arg2.clone()))
                                                                          })))
    }
    pub fn Data_DateTime_toRecord() -> &dyn Any {
        static Data_DateTime_toRecord: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_toRecord.get_or_init(||
                                               &Func1::new(move |v|
                                                               {
                                                                   let matchValue:
                                                                           LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                       Sharpurs_Prelude::unbox(v);
                                                                   let t =
                                                                       match matchValue.as_ref()
                                                                           {
                                                                           PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                            x)
                                                                           =>
                                                                           x.clone(),
                                                                       };
                                                                   let d =
                                                                       match matchValue.as_ref()
                                                                           {
                                                                           PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                            _)
                                                                           =>
                                                                           x.clone(),
                                                                       };
                                                                   &add(string("year"),
                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                             &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumYear()),
                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_year(),
                                                                                                                                             &&&d)),
                                                                        add(string("month"),
                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                 &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_month(),
                                                                                                                                                 &&&d)),
                                                                            add(string("day"),
                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                     &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_day(),
                                                                                                                                                     &&&d)),
                                                                                add(string("hour"),
                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                         &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumHour()),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_hour(),
                                                                                                                                                         &&&t)),
                                                                                    add(string("minute"),
                                                                                        &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                             &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumMinute()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_minute(),
                                                                                                                                                             &&&t)),
                                                                                        add(string("second"),
                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                 &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumSecond()),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_second(),
                                                                                                                                                                 &&&t)),
                                                                                            add(string("millisecond"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                     &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumMillisecond()),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time::Data_Time_millisecond(),
                                                                                                                                                                     &&&t)),
                                                                                                empty::<string,
                                                                                                        &dyn Any>())))))))
                                                               }))
    }
    pub fn Data_DateTime_time() -> &dyn Any {
        static Data_DateTime_time: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_time.get_or_init(||
                                           &Func1::new(move |v|
                                                           &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                {
                                                                PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                 x)
                                                                => x.clone(),
                                                            }))
    }
    pub fn Data_DateTime_showDateTime() -> &dyn Any {
        static Data_DateTime_showDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_showDateTime.get_or_init(||
                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                    &&&add(string("show"),
                                                                                           &&Func1::new(move
                                                                                                            |v|
                                                                                                            {
                                                                                                                let matchValue:
                                                                                                                        LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                                                    Sharpurs_Prelude::unbox(v);
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                       &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                    &&&string("(DateTime ")),
                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                          &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Date::Data_Date_showDate()),
                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                 PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
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
                                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Time::Data_Time_showTime()),
                                                                                                                                                                                                                                                                                                                                &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                       {
                                                                                                                                                                                                                                                                                                                                       PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                                                                       =>
                                                                                                                                                                                                                                                                                                                                       x.clone(),
                                                                                                                                                                                                                                                                                                                                   })),
                                                                                                                                                                                                                                                          &&&string(")")))))
                                                                                                            }),
                                                                                           empty::<string,
                                                                                                   &dyn Any>())))
    }
    pub fn Data_DateTime_modifyTimeF() -> &dyn Any {
        static Data_DateTime_modifyTimeF: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_modifyTimeF.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictFunctor|
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
                                                                                                                      LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                     &&&dictFunctor),
                                                                                                                                                                                  &&&Func1::new({
                                                                                                                                                                                                    let matchValue_1
                                                                                                                                                                                                        =
                                                                                                                                                                                                        matchValue_1.clone();
                                                                                                                                                                                                    move
                                                                                                                                                                                                        |usd__arg1|
                                                                                                                                                                                                        &LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                                           PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                            _)
                                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                                       },
                                                                                                                                                                                                                                                                                                      usd__arg1.clone()))
                                                                                                                                                                                                })),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                         PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                          x)
                                                                                                                                                                                         =>
                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                     }))
                                                                                                          }
                                                                                                  })
                                                                              })))
    }
    pub fn Data_DateTime_modifyTime() -> &dyn Any {
        static Data_DateTime_modifyTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_modifyTime.get_or_init(||
                                                 &Func1::new(move |f|
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
                                                                                                 LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                         &LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(&match matchValue_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                        },
                                                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                x)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                           })))
                                                                                     }
                                                                             })))
    }
    pub fn Data_DateTime_modifyDateF() -> &dyn Any {
        static Data_DateTime_modifyDateF: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_modifyDateF.get_or_init(||
                                                  &Func1::new(move
                                                                  |dictFunctor|
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
                                                                                                                      LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                                                  Sharpurs_Prelude::unbox(v);
                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                     &&&dictFunctor),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                          |usd__arg1|
                                                                                                                                                                                                                                                                          Func1::new({
                                                                                                                                                                                                                                                                                         let usd__arg1
                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                             usd__arg1.clone();
                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                             |usd__arg2|
                                                                                                                                                                                                                                                                                             &LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                           usd__arg2.clone()))
                                                                                                                                                                                                                                                                                     }))),
                                                                                                                                                                                                                     &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                            {
                                                                                                                                                                                                                            PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                        })),
                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                         {
                                                                                                                                                                                         PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                                                                                                                                          _)
                                                                                                                                                                                         =>
                                                                                                                                                                                         x.clone(),
                                                                                                                                                                                     }))
                                                                                                          }
                                                                                                  })
                                                                              })))
    }
    pub fn Data_DateTime_modifyDate() -> &dyn Any {
        static Data_DateTime_modifyDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_modifyDate.get_or_init(||
                                                 &Func1::new(move |f|
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
                                                                                                 LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                         &LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&matchValue,
                                                                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                               {
                                                                                                                                                                                                                               PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                               =>
                                                                                                                                                                                                                               x.clone(),
                                                                                                                                                                                                                           }),
                                                                                                                                                                                       &match matchValue_1.as_ref()
                                                                                                                                                                                            {
                                                                                                                                                                                            PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                             x)
                                                                                                                                                                                            =>
                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                        }))
                                                                                     }
                                                                             })))
    }
    pub fn Data_DateTime_mkDateRec() -> &dyn Any {
        static Data_DateTime_mkDateRec: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_mkDateRec.get_or_init(||
                                                &Func1::new(move |y|
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
                                                                                                                                                                                        &add(string("year"),
                                                                                                                                                                                             &&y,
                                                                                                                                                                                             add(string("month"),
                                                                                                                                                                                                 &&mo,
                                                                                                                                                                                                 add(string("day"),
                                                                                                                                                                                                     &&d,
                                                                                                                                                                                                     add(string("hour"),
                                                                                                                                                                                                         &&h,
                                                                                                                                                                                                         add(string("minute"),
                                                                                                                                                                                                             &&mi,
                                                                                                                                                                                                             add(string("second"),
                                                                                                                                                                                                                 &&s,
                                                                                                                                                                                                                 add(string("millisecond"),
                                                                                                                                                                                                                     ms.clone(),
                                                                                                                                                                                                                     empty::<string,
                                                                                                                                                                                                                             &dyn Any>())))))))
                                                                                                                                                                                })
                                                                                                                                                            })
                                                                                                                                        })
                                                                                                                    })
                                                                                                })
                                                                            })))
    }
    pub fn Data_DateTime_eqDateTime() -> &dyn Any {
        static Data_DateTime_eqDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_eqDateTime.get_or_init(||
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
                                                                                                                                          LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                                                                      Sharpurs_Prelude::unbox(&&x);
                                                                                                                                  let matchValue_1:
                                                                                                                                          LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                                                                      Sharpurs_Prelude::unbox(y);
                                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                         &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Date::Data_Date_eqDate()),
                                                                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                                                                   PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                    _)
                                                                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                                                                               }),
                                                                                                                                                                                                                                         &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                            })),
                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                            &&&PureScript_Data_Time::Data_Time_eqTime()),
                                                                                                                                                                                                                                         &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                {
                                                                                                                                                                                                                                                PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                 x)
                                                                                                                                                                                                                                                =>
                                                                                                                                                                                                                                                x.clone(),
                                                                                                                                                                                                                                            }),
                                                                                                                                                                                                      &&&match matchValue_1.as_ref()
                                                                                                                                                                                                             {
                                                                                                                                                                                                             PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                                              x)
                                                                                                                                                                                                             =>
                                                                                                                                                                                                             x.clone(),
                                                                                                                                                                                                         }))
                                                                                                                              }
                                                                                                                      })),
                                                                                         empty::<string,
                                                                                                 &dyn Any>())))
    }
    pub fn Data_DateTime_ordDateTime() -> &dyn Any {
        static Data_DateTime_ordDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_ordDateTime.get_or_init(||
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
                                                                                                                                           LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                                                                       Sharpurs_Prelude::unbox(&&x);
                                                                                                                                   let matchValue_1:
                                                                                                                                           LrcPtr<PureScript_Data_DateTime::Data_DateTime_DateTime> =
                                                                                                                                       Sharpurs_Prelude::unbox(y);
                                                                                                                                   let matchValue_3:
                                                                                                                                           LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Date::Data_Date_ordDate()),
                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                             _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
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
                                                                                                                                                                                                                                              &&&PureScript_Data_Time::Data_Time_ordTime()),
                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                                                   x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                              }),
                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                               PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(_,
                                                                                                                                                                                                                                                                x)
                                                                                                                                                                               =>
                                                                                                                                                                               x.clone(),
                                                                                                                                                                           }),
                                                                                                                                   }
                                                                                                                               }
                                                                                                                       })),
                                                                                          add(string("Eq0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_DateTime::Data_DateTime_eqDateTime()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>()))))
    }
    pub fn Data_DateTime_diff() -> &dyn Any {
        static Data_DateTime_diff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_diff.get_or_init(||
                                           &Func1::new(move |dictDuration|
                                                           {
                                                               let toDuration =
                                                                   Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_toDuration(),
                                                                                                    dictDuration);
                                                               &Func1::new({
                                                                               let toDuration
                                                                                   =
                                                                                   toDuration.clone();
                                                                               move
                                                                                   |dt1|
                                                                                   &Func1::new({
                                                                                                   let dt1
                                                                                                       =
                                                                                                       dt1.clone();
                                                                                                   move
                                                                                                       |dt2|
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                           &&&toDuration),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn2(),
                                                                                                                                                                                                                                                 &&&PureScript_Data_DateTime::Data_DateTime_calcDiff()),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime::Data_DateTime_toRecord(),
                                                                                                                                                                                                                                                 &&&dt1)),
                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime::Data_DateTime_toRecord(),
                                                                                                                                                                                                              dt2)))
                                                                                               })
                                                                           })
                                                           }))
    }
    pub fn Data_DateTime_date() -> &dyn Any {
        static Data_DateTime_date: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_date.get_or_init(||
                                           &Func1::new(move |v|
                                                           &match Sharpurs_Prelude::unbox(v).as_ref()
                                                                {
                                                                PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(x,
                                                                                                                                                 _)
                                                                => x.clone(),
                                                            }))
    }
    pub fn Data_DateTime_boundedDateTime() -> &dyn Any {
        static Data_DateTime_boundedDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_boundedDateTime.get_or_init(||
                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                       &&&add(string("bottom"),
                                                                                              &&LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                              &&&PureScript_Data_Date::Data_Date_boundedDate()),
                                                                                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                                              &&&PureScript_Data_Time::Data_Time_boundedTime()))),
                                                                                              add(string("top"),
                                                                                                  &&LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                                  &&&PureScript_Data_Date::Data_Date_boundedDate()),
                                                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                                                  &&&PureScript_Data_Time::Data_Time_boundedTime()))),
                                                                                                  add(string("Ord0"),
                                                                                                      &&Func1::new(move
                                                                                                                       |usd__unused|
                                                                                                                       &PureScript_Data_DateTime::Data_DateTime_ordDateTime()),
                                                                                                      empty::<string,
                                                                                                              &dyn Any>())))))
    }
    pub fn Data_DateTime_adjust() -> &dyn Any {
        static Data_DateTime_adjust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_DateTime_adjust.get_or_init(||
                                             &Func1::new(move |dictDuration|
                                                             &Func1::new({
                                                                             let dictDuration
                                                                                 =
                                                                                 dictDuration.clone();
                                                                             move
                                                                                 |d|
                                                                                 &Func1::new({
                                                                                                 let d
                                                                                                     =
                                                                                                     d.clone();
                                                                                                 move
                                                                                                     |dt|
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                                            &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime::Data_DateTime_adjustImpl(),
                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_DateTime::Data_DateTime_mkDateRec()),
                                                                                                                                                                                                                                                                                                                     &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                       |usd__arg1|
                                                                                                                                                                                                                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                                                                                                                                                                                                  &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)),
                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Duration::Data_Time_Duration_fromDuration(),
                                                                                                                                                                                                                                                                                                                     &&&dictDuration),
                                                                                                                                                                                                                                                                                  &&&d)),
                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DateTime::Data_DateTime_toRecord(),
                                                                                                                                                                                                                                               dt))),
                                                                                                                                      &&&Func1::new(move
                                                                                                                                                        |rec_var|
                                                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                  &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                    |usd__arg1_1|
                                                                                                                                                                                                                                                                                                                    Func1::new({
                                                                                                                                                                                                                                                                                                                                   let usd__arg1_1
                                                                                                                                                                                                                                                                                                                                       =
                                                                                                                                                                                                                                                                                                                                       usd__arg1_1.clone();
                                                                                                                                                                                                                                                                                                                                   move
                                                                                                                                                                                                                                                                                                                                       |usd__arg2|
                                                                                                                                                                                                                                                                                                                                       &LrcPtr::new(PureScript_Data_DateTime::Data_DateTime_DateTime::Data_DateTime_DateTimeusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                                                                                                                     usd__arg2.clone()))
                                                                                                                                                                                                                                                                                                                               }))),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_join(),
                                                                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Date::Data_Date_exactDate()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumYear()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&find(string("year"),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           Sharpurs_Prelude::unbox(rec_var))))),
                                                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                                                                                                                                                                                                                                              &&find(string("month"),
                                                                                                                                                                                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::unbox(rec_var))))),
                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                                                                                                                                                                                                                                                                        &&find(string("day"),
                                                                                                                                                                                                                                                                                                                                                                               Sharpurs_Prelude::unbox(rec_var))))))),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |usd__arg1_2|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  let usd__arg1_2
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      =
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      usd__arg1_2.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |usd__arg2_1|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     let usd__arg2_1
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         usd__arg2_1.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |usd__arg3|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        let usd__arg3
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            =
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            usd__arg3.clone();
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        move
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            |usd__arg4|
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            &LrcPtr::new(Data_Time_Time::Data_Time_Timeusd_Ctor(usd__arg1_2,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                usd__arg2_1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                usd__arg3,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                usd__arg4.clone()))
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    })
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 })
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              }))),
                                                                                                                                                                                                                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumHour()),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 &&find(string("hour"),
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        Sharpurs_Prelude::unbox(rec_var))))),
                                                                                                                                                                                                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                                                                                              &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumMinute()),
                                                                                                                                                                                                                                                                                                                                                                                                           &&find(string("minute"),
                                                                                                                                                                                                                                                                                                                                                                                                                  Sharpurs_Prelude::unbox(rec_var))))),
                                                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumSecond()),
                                                                                                                                                                                                                                                                                                                                     &&find(string("second"),
                                                                                                                                                                                                                                                                                                                                            Sharpurs_Prelude::unbox(rec_var))))),
                                                                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                  &&&PureScript_Data_Time_Component::Data_Time_Component_boundedEnumMillisecond()),
                                                                                                                                                                                                                                                               &&find(string("millisecond"),
                                                                                                                                                                                                                                                                      Sharpurs_Prelude::unbox(rec_var)))))))
                                                                                             })
                                                                         })))
    }
}

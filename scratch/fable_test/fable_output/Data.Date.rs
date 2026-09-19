pub mod PureScript_Data_Date {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func0;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_ce6bdf8a::PureScript_Control_Applicative;
    use crate::module_bb42e836::PureScript_Control_Apply;
    use crate::module_6c4d46c3::PureScript_Control_Bind;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_b6aac8e0::PureScript_Data_Date_Component;
    use crate::module_b6aac8e0::PureScript_Data_Date_Component::Data_Date_Component_Month;
    use crate::module_6a1c5ce6::PureScript_Data_Enum;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_6e7709b7::PureScript_Data_Function;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_2eb6dec6::PureScript_Data_Int;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_c4b34810::PureScript_Data_Time_Duration;
    use crate::module_11800e3c::PureScript_Partial_Unsafe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub fn Data_Date_calcDiff() -> &dyn Any {
        static Data_Date_calcDiff: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_calcDiff.get_or_init(||
                                           &Func1::new(move |arg0|
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
                                                                                                                                           panic!("{}",
                                                                                                                                                  1_i32.get_Message(),))))))))
    }
    pub fn Data_Date_calcWeekday() -> &dyn Any {
        static Data_Date_calcWeekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_calcWeekday.get_or_init(||
                                              &Func1::new(move |arg0|
                                                              &Func1::new(move
                                                                              |arg1|
                                                                              &Func1::new(move
                                                                                              |arg2|
                                                                                              panic!("{}",
                                                                                                     1_i32.get_Message(),)))))
    }
    pub fn Data_Date_canonicalDateImpl() -> &dyn Any {
        static Data_Date_canonicalDateImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_canonicalDateImpl.get_or_init(||
                                                    &Func1::new(move |arg0|
                                                                    &Func1::new(move
                                                                                    |arg1|
                                                                                    &Func1::new(move
                                                                                                    |arg2|
                                                                                                    &Func1::new(move
                                                                                                                    |arg3|
                                                                                                                    panic!("{}",
                                                                                                                           1_i32.get_Message(),))))))
    }
    #[derive(Clone, Debug,)]
    pub enum Data_Date_Date {
        Data_Date_Dateusd_Ctor(&dyn Any, &dyn Any, &dyn Any),
    }
    impl core::fmt::Display for PureScript_Data_Date::Data_Date_Date {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Date_toEnum() -> &dyn Any {
        static Data_Date_toEnum: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_toEnum.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                          &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()))
    }
    pub fn Data_Date_ordMaybe() -> &dyn Any {
        static Data_Date_ordMaybe: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_ordMaybe.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_ordMaybe(),
                                                                            &&&PureScript_Data_Date_Component::Data_Date_Component_ordDay()))
    }
    pub fn Data_Date_Date() -> &dyn Any {
        static Data_Date_Date: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Date_Date.get_or_init(||
                                       &Func1::new(move |usd__arg1|
                                                       Func1::new({
                                                                      let usd__arg1
                                                                          =
                                                                          usd__arg1.clone();
                                                                      move
                                                                          |usd__arg2|
                                                                          Func1::new({
                                                                                         let usd__arg2
                                                                                             =
                                                                                             usd__arg2.clone();
                                                                                         move
                                                                                             |usd__arg3|
                                                                                             &LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(usd__arg1,
                                                                                                                                                                       usd__arg2,
                                                                                                                                                                       usd__arg3.clone()))
                                                                                     })
                                                                  })))
    }
    pub fn Data_Date_year() -> &dyn Any {
        static Data_Date_year: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Date_year.get_or_init(||
                                       &Func1::new(move |v|
                                                       &match Sharpurs_Prelude::unbox(v).as_ref()
                                                            {
                                                            PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                         _,
                                                                                                                         _)
                                                            => x.clone(),
                                                        }))
    }
    pub fn Data_Date_weekday() -> &dyn Any {
        static Data_Date_weekday: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_weekday.get_or_init(||
                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                           &&&Func1::new(move
                                                                                             |usd__unused|
                                                                                             &Func1::new(move
                                                                                                             |v|
                                                                                                             {
                                                                                                                 let matchValue:
                                                                                                                         LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                                     Sharpurs_Prelude::unbox(v);
                                                                                                                 let n =
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn3(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Date::Data_Date_calcWeekday()),
                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                   PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                                                                _,
                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                               }),
                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                               &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                            &&&match matchValue.as_ref()
                                                                                                                                                                                                                                   {
                                                                                                                                                                                                                                   PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                x,
                                                                                                                                                                                                                                                                                                _)
                                                                                                                                                                                                                                   =>
                                                                                                                                                                                                                                   x.clone(),
                                                                                                                                                                                                                               })),
                                                                                                                                                      &&&match matchValue.as_ref()
                                                                                                                                                             {
                                                                                                                                                             PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                          _,
                                                                                                                                                                                                                          x)
                                                                                                                                                             =>
                                                                                                                                                             x.clone(),
                                                                                                                                                         });
                                                                                                                 let matchValue_1 =
                                                                                                                     Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                     &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                  &&&n),
                                                                                                                                                                               &&&0_i32));
                                                                                                                 match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                  &matchValue_1)
                                                                                                                     {
                                                                                                                     0_i32
                                                                                                                     =>
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                            &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumWeekday()),
                                                                                                                                                                                         &&&7_i32)),
                                                                                                                     _
                                                                                                                     =>
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                            &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumWeekday()),
                                                                                                                                                                                         &&&n)),
                                                                                                                 }
                                                                                                             }))))
    }
    pub fn Data_Date_showDate() -> &dyn Any {
        static Data_Date_showDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_showDate.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                            &&&add(string("show"),
                                                                                   &&Func1::new(move
                                                                                                    |v|
                                                                                                    {
                                                                                                        let matchValue:
                                                                                                                LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                               &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                            &&&string("(Date ")),
                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Date_Component::Data_Date_Component_showYear()),
                                                                                                                                                                                                                                                  &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                         {
                                                                                                                                                                                                                                                         PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                      _,
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
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Date_Component::Data_Date_Component_showMonth()),
                                                                                                                                                                                                                                                                                                                        &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                               {
                                                                                                                                                                                                                                                                                                                               PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                            x,
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
                                                                                                                                                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Date_Component::Data_Date_Component_showDay()),
                                                                                                                                                                                                                                                                                                                                                                                              &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                                                     {
                                                                                                                                                                                                                                                                                                                                                                                                     PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                                                                                                                                                                                                                     =>
                                                                                                                                                                                                                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                                                                                                                                        &&&string(")")))))))
                                                                                                    }),
                                                                                   empty::<string,
                                                                                           &dyn Any>())))
    }
    pub fn Data_Date_month() -> &dyn Any {
        static Data_Date_month: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_month.get_or_init(||
                                        &Func1::new(move |v|
                                                        &match Sharpurs_Prelude::unbox(v).as_ref()
                                                             {
                                                             PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                          x,
                                                                                                                          _)
                                                             => x.clone(),
                                                         }))
    }
    pub fn Data_Date_isLeapYear() -> &dyn Any {
        static Data_Date_isLeapYear: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_isLeapYear.get_or_init(||
                                             &Func1::new(move |y|
                                                             {
                                                                 let y_prime =
                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                         &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumYear()),
                                                                                                      y);
                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                        &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                              &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                                                                                                                    &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                                                                                 &&&y_prime),
                                                                                                                                                                                                                                              &&&4_i32)),
                                                                                                                                                                        &&&0_i32)),
                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                           &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                                                                                                                    &&&y_prime),
                                                                                                                                                                                                                                                                                 &&&400_i32)),
                                                                                                                                                                                                           &&&0_i32)),
                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                                                           &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingInt()),
                                                                                                                                                                                                                                                                                                                    &&&y_prime),
                                                                                                                                                                                                                                                                                 &&&100_i32)),
                                                                                                                                                                                                           &&&0_i32))))
                                                             }))
    }
    pub fn Data_Date_lastDayOfMonth() -> &dyn Any {
        static Data_Date_lastDayOfMonth: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_lastDayOfMonth.get_or_init(||
                                                 &Func1::new(move |y|
                                                                 &Func1::new({
                                                                                 let y
                                                                                     =
                                                                                     y.clone();
                                                                                 move
                                                                                     |m|
                                                                                     {
                                                                                         let unsafeDay =
                                                                                             Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                    &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                                                                    &&&Func1::new(move
                                                                                                                                                                                                                      |usd__unused|
                                                                                                                                                                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                       &&&Sharpurs_Prelude::Prim_undefined())))),
                                                                                                                              &&&PureScript_Data_Date::Data_Date_toEnum());
                                                                                         let matchValue:
                                                                                                 LrcPtr<Data_Date_Component_Month> =
                                                                                             Sharpurs_Prelude::unbox(m);
                                                                                         if let Data_Date_Component_Month::Data_Date_Component_Februaryusd_Ctor
                                                                                                =
                                                                                                matchValue.as_ref()
                                                                                            {
                                                                                             if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_isLeapYear(),
                                                                                                                                                          &&&y))
                                                                                                {
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                  &&&29_i32)
                                                                                             } else {
                                                                                                 if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                    {
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                      &&&28_i32)
                                                                                                 } else {
                                                                                                     panic!("{}",
                                                                                                            string("Match failure: PureScript_Data_Date_Component.Data_Date_Component_Month"),)
                                                                                                 }
                                                                                             }
                                                                                         } else {
                                                                                             if let Data_Date_Component_Month::Data_Date_Component_Marchusd_Ctor
                                                                                                    =
                                                                                                    matchValue.as_ref()
                                                                                                {
                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                  &&&31_i32)
                                                                                             } else {
                                                                                                 if let Data_Date_Component_Month::Data_Date_Component_Aprilusd_Ctor
                                                                                                        =
                                                                                                        matchValue.as_ref()
                                                                                                    {
                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                      &&&30_i32)
                                                                                                 } else {
                                                                                                     if let Data_Date_Component_Month::Data_Date_Component_Mayusd_Ctor
                                                                                                            =
                                                                                                            matchValue.as_ref()
                                                                                                        {
                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                          &&&31_i32)
                                                                                                     } else {
                                                                                                         if let Data_Date_Component_Month::Data_Date_Component_Juneusd_Ctor
                                                                                                                =
                                                                                                                matchValue.as_ref()
                                                                                                            {
                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                              &&&30_i32)
                                                                                                         } else {
                                                                                                             if let Data_Date_Component_Month::Data_Date_Component_Julyusd_Ctor
                                                                                                                    =
                                                                                                                    matchValue.as_ref()
                                                                                                                {
                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                                  &&&31_i32)
                                                                                                             } else {
                                                                                                                 if let Data_Date_Component_Month::Data_Date_Component_Augustusd_Ctor
                                                                                                                        =
                                                                                                                        matchValue.as_ref()
                                                                                                                    {
                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                                      &&&31_i32)
                                                                                                                 } else {
                                                                                                                     if let Data_Date_Component_Month::Data_Date_Component_Septemberusd_Ctor
                                                                                                                            =
                                                                                                                            matchValue.as_ref()
                                                                                                                        {
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                                          &&&30_i32)
                                                                                                                     } else {
                                                                                                                         if let Data_Date_Component_Month::Data_Date_Component_Octoberusd_Ctor
                                                                                                                                =
                                                                                                                                matchValue.as_ref()
                                                                                                                            {
                                                                                                                             Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                                              &&&31_i32)
                                                                                                                         } else {
                                                                                                                             if let Data_Date_Component_Month::Data_Date_Component_Novemberusd_Ctor
                                                                                                                                    =
                                                                                                                                    matchValue.as_ref()
                                                                                                                                {
                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                                                  &&&30_i32)
                                                                                                                             } else {
                                                                                                                                 if let Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor
                                                                                                                                        =
                                                                                                                                        matchValue.as_ref()
                                                                                                                                    {
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                                                      &&&31_i32)
                                                                                                                                 } else {
                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&unsafeDay,
                                                                                                                                                                      &&&31_i32)
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
                                                                             })))
    }
    pub fn Data_Date_eqDate() -> &dyn Any {
        static Data_Date_eqDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_eqDate.get_or_init(||
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
                                                                                                                                  LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                                              Sharpurs_Prelude::unbox(&&x);
                                                                                                                          let matchValue_1:
                                                                                                                                  LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                                              Sharpurs_Prelude::unbox(y);
                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                 &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Date_Component::Data_Date_Component_eqYear()),
                                                                                                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                 PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                                                                                                       &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                              PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                           _,
                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                          })),
                                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                          &&&PureScript_Data_Date_Component::Data_Date_Component_eqMonth()),
                                                                                                                                                                                                                                                                                                       &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                              {
                                                                                                                                                                                                                                                                                                              PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                           x,
                                                                                                                                                                                                                                                                                                                                                                           _)
                                                                                                                                                                                                                                                                                                              =>
                                                                                                                                                                                                                                                                                                              x.clone(),
                                                                                                                                                                                                                                                                                                          }),
                                                                                                                                                                                                                                                                    &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                           PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                        x,
                                                                                                                                                                                                                                                                                                                                        _)
                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                       }))),
                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                    &&&PureScript_Data_Date_Component::Data_Date_Component_eqDay()),
                                                                                                                                                                                                                                 &&&match matchValue.as_ref()
                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                        PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                     _,
                                                                                                                                                                                                                                                                                                     x)
                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                    }),
                                                                                                                                                                                              &&&match matchValue_1.as_ref()
                                                                                                                                                                                                     {
                                                                                                                                                                                                     PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                  _,
                                                                                                                                                                                                                                                                  x)
                                                                                                                                                                                                     =>
                                                                                                                                                                                                     x.clone(),
                                                                                                                                                                                                 }))
                                                                                                                      }
                                                                                                              })),
                                                                                 empty::<string,
                                                                                         &dyn Any>())))
    }
    pub fn Data_Date_ordDate() -> &dyn Any {
        static Data_Date_ordDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_ordDate.get_or_init(||
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
                                                                                                                                   LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                                               Sharpurs_Prelude::unbox(&&x);
                                                                                                                           let matchValue_1:
                                                                                                                                   LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                                               Sharpurs_Prelude::unbox(y);
                                                                                                                           let matchValue_3:
                                                                                                                                   LrcPtr<Data_Ordering_Ordering> =
                                                                                                                               Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                &&&PureScript_Data_Date_Component::Data_Date_Component_ordYear()),
                                                                                                                                                                                                                             &&&match matchValue.as_ref()
                                                                                                                                                                                                                                    {
                                                                                                                                                                                                                                    PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                                                                 _,
                                                                                                                                                                                                                                                                                                 _)
                                                                                                                                                                                                                                    =>
                                                                                                                                                                                                                                    x.clone(),
                                                                                                                                                                                                                                }),
                                                                                                                                                                                          &&&match matchValue_1.as_ref()
                                                                                                                                                                                                 {
                                                                                                                                                                                                 PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                              _,
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
                                                                                                                               {
                                                                                                                                   let matchValue_4:
                                                                                                                                           LrcPtr<Data_Ordering_Ordering> =
                                                                                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_compare(),
                                                                                                                                                                                                                                                                        &&&PureScript_Data_Date_Component::Data_Date_Component_ordMonth()),
                                                                                                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                                                                                                            {
                                                                                                                                                                                                                                            PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                         x,
                                                                                                                                                                                                                                                                                                         _)
                                                                                                                                                                                                                                            =>
                                                                                                                                                                                                                                            x.clone(),
                                                                                                                                                                                                                                        }),
                                                                                                                                                                                                  &&&match matchValue_1.as_ref()
                                                                                                                                                                                                         {
                                                                                                                                                                                                         PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                      x,
                                                                                                                                                                                                                                                                      _)
                                                                                                                                                                                                         =>
                                                                                                                                                                                                         x.clone(),
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
                                                                                                                                                                                                                                              &&&PureScript_Data_Date_Component::Data_Date_Component_ordDay()),
                                                                                                                                                                                                           &&&match matchValue.as_ref()
                                                                                                                                                                                                                  {
                                                                                                                                                                                                                  PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                               _,
                                                                                                                                                                                                                                                                               x)
                                                                                                                                                                                                                  =>
                                                                                                                                                                                                                  x.clone(),
                                                                                                                                                                                                              }),
                                                                                                                                                                        &&&match matchValue_1.as_ref()
                                                                                                                                                                               {
                                                                                                                                                                               PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                            _,
                                                                                                                                                                                                                                            x)
                                                                                                                                                                               =>
                                                                                                                                                                               x.clone(),
                                                                                                                                                                           }),
                                                                                                                                   }
                                                                                                                               }
                                                                                                                           }
                                                                                                                       }
                                                                                                               })),
                                                                                  add(string("Eq0"),
                                                                                      &&Func1::new(move
                                                                                                       |usd__unused|
                                                                                                       &PureScript_Data_Date::Data_Date_eqDate()),
                                                                                      empty::<string,
                                                                                              &dyn Any>()))))
    }
    pub fn Data_Date_enumDate() -> &dyn Any {
        static Data_Date_enumDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_enumDate.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                                                            &&&add(string("succ"),
                                                                                   &&Func1::new(move
                                                                                                    |v|
                                                                                                    {
                                                                                                        let matchValue:
                                                                                                                LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                            Sharpurs_Prelude::unbox(v);
                                                                                                        let y =
                                                                                                            match matchValue.as_ref()
                                                                                                                {
                                                                                                                PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                             _,
                                                                                                                                                                             _)
                                                                                                                =>
                                                                                                                x.clone(),
                                                                                                            };
                                                                                                        let m =
                                                                                                            match matchValue.as_ref()
                                                                                                                {
                                                                                                                PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                             x,
                                                                                                                                                                             _)
                                                                                                                =>
                                                                                                                x.clone(),
                                                                                                            };
                                                                                                        let sm =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                &&&PureScript_Data_Date_Component::Data_Date_Component_enumMonth()),
                                                                                                                                             &&&m);
                                                                                                        let l =
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_lastDayOfMonth(),
                                                                                                                                                                                &&&y),
                                                                                                                                             &&&m);
                                                                                                        let sd =
                                                                                                            {
                                                                                                                let v1 =
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                        &&&PureScript_Data_Date_Component::Data_Date_Component_enumDay()),
                                                                                                                                                     &&&match matchValue.as_ref()
                                                                                                                                                            {
                                                                                                                                                            PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                         _,
                                                                                                                                                                                                                         x)
                                                                                                                                                            =>
                                                                                                                                                            x.clone(),
                                                                                                                                                        });
                                                                                                                let matchValue_1 =
                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                                    &&&PureScript_Data_Date::Data_Date_ordMaybe()),
                                                                                                                                                                                                                 &&&v1),
                                                                                                                                                                              &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&l))));
                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                 &matchValue_1)
                                                                                                                    {
                                                                                                                    0_i32
                                                                                                                    =>
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    &v1,
                                                                                                                }
                                                                                                            };
                                                                                                        let m_prime =
                                                                                                            {
                                                                                                                let matchValue_2 =
                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                              &&&sd));
                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                 &matchValue_2)
                                                                                                                    {
                                                                                                                    0_i32
                                                                                                                    =>
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromMaybe(),
                                                                                                                                                                                        &&&LrcPtr::new(Data_Date_Component_Month::Data_Date_Component_Januaryusd_Ctor)),
                                                                                                                                                     &&&sm),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    &m,
                                                                                                                }
                                                                                                            };
                                                                                                        let y_prime =
                                                                                                            {
                                                                                                                let matchValue_3 =
                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                    &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                                                                                                    &&&sd)),
                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                                                                 &&&sm)));
                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                 &matchValue_3)
                                                                                                                    {
                                                                                                                    0_i32
                                                                                                                    =>
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                        &&&PureScript_Data_Date_Component::Data_Date_Component_enumYear()),
                                                                                                                                                     &&&y),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&y)),
                                                                                                                }
                                                                                                            };
                                                                                                        let d_prime =
                                                                                                            {
                                                                                                                let matchValue_4 =
                                                                                                                    Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                              &&&sd));
                                                                                                                match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                 &matchValue_4)
                                                                                                                    {
                                                                                                                    0_i32
                                                                                                                    =>
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                        &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                                                     &&&1_i32),
                                                                                                                    _
                                                                                                                    =>
                                                                                                                    &sd,
                                                                                                                }
                                                                                                            };
                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                               &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                                                                           &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                                                                                                                                                                                          |usd__arg1|
                                                                                                                                                                                                                                                                                                                                          Func1::new({
                                                                                                                                                                                                                                                                                                                                                         let usd__arg1
                                                                                                                                                                                                                                                                                                                                                             =
                                                                                                                                                                                                                                                                                                                                                             usd__arg1.clone();
                                                                                                                                                                                                                                                                                                                                                         move
                                                                                                                                                                                                                                                                                                                                                             |usd__arg2|
                                                                                                                                                                                                                                                                                                                                                             Func1::new({
                                                                                                                                                                                                                                                                                                                                                                            let usd__arg2
                                                                                                                                                                                                                                                                                                                                                                                =
                                                                                                                                                                                                                                                                                                                                                                                usd__arg2.clone();
                                                                                                                                                                                                                                                                                                                                                                            move
                                                                                                                                                                                                                                                                                                                                                                                |usd__arg3|
                                                                                                                                                                                                                                                                                                                                                                                &LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(usd__arg1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                          usd__arg2,
                                                                                                                                                                                                                                                                                                                                                                                                                                                          usd__arg3.clone()))
                                                                                                                                                                                                                                                                                                                                                                        })
                                                                                                                                                                                                                                                                                                                                                     }))),
                                                                                                                                                                                                                                                                                     &&&y_prime)),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Maybe::Data_Maybe_applicativeMaybe()),
                                                                                                                                                                                                                                                  &&&m_prime))),
                                                                                                                                         &&&d_prime)
                                                                                                    }),
                                                                                   add(string("pred"),
                                                                                       &&Func1::new(move
                                                                                                        |v_1|
                                                                                                        {
                                                                                                            let matchValue_5:
                                                                                                                    LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                                Sharpurs_Prelude::unbox(v_1);
                                                                                                            let y_1 =
                                                                                                                match matchValue_5.as_ref()
                                                                                                                    {
                                                                                                                    PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                 _,
                                                                                                                                                                                 _)
                                                                                                                    =>
                                                                                                                    x.clone(),
                                                                                                                };
                                                                                                            let m_1 =
                                                                                                                match matchValue_5.as_ref()
                                                                                                                    {
                                                                                                                    PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                 x,
                                                                                                                                                                                 _)
                                                                                                                    =>
                                                                                                                    x.clone(),
                                                                                                                };
                                                                                                            let pm =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                    &&&PureScript_Data_Date_Component::Data_Date_Component_enumMonth()),
                                                                                                                                                 &&&m_1);
                                                                                                            let pd =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                    &&&PureScript_Data_Date_Component::Data_Date_Component_enumDay()),
                                                                                                                                                 &&&match matchValue_5.as_ref()
                                                                                                                                                        {
                                                                                                                                                        PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                     _,
                                                                                                                                                                                                                     x)
                                                                                                                                                        =>
                                                                                                                                                        x.clone(),
                                                                                                                                                    });
                                                                                                            let y_prime_1 =
                                                                                                                {
                                                                                                                    let matchValue_6 =
                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_conj(),
                                                                                                                                                                                                                                                        &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                                                                                                        &&&pd)),
                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                                                                     &&&pm)));
                                                                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                     &matchValue_6)
                                                                                                                        {
                                                                                                                        0_i32
                                                                                                                        =>
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                            &&&PureScript_Data_Date_Component::Data_Date_Component_enumYear()),
                                                                                                                                                         &&&y_1),
                                                                                                                        _
                                                                                                                        =>
                                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&y_1)),
                                                                                                                    }
                                                                                                                };
                                                                                                            let m_prime_1 =
                                                                                                                {
                                                                                                                    let matchValue_7 =
                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                                  &&&pd));
                                                                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                     &matchValue_7)
                                                                                                                        {
                                                                                                                        0_i32
                                                                                                                        =>
                                                                                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromMaybe(),
                                                                                                                                                                                            &&&LrcPtr::new(Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor)),
                                                                                                                                                         &&&pm),
                                                                                                                        _
                                                                                                                        =>
                                                                                                                        &m_1,
                                                                                                                    }
                                                                                                                };
                                                                                                            let l_1 =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_lastDayOfMonth(),
                                                                                                                                                                                    &&&y_1),
                                                                                                                                                 &&&m_prime_1);
                                                                                                            let d_prime_1 =
                                                                                                                {
                                                                                                                    let matchValue_8 =
                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_isNothing(),
                                                                                                                                                                                  &&&pd));
                                                                                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                     &matchValue_8)
                                                                                                                        {
                                                                                                                        0_i32
                                                                                                                        =>
                                                                                                                        &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&l_1)),
                                                                                                                        _
                                                                                                                        =>
                                                                                                                        &pd,
                                                                                                                    }
                                                                                                                };
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
                                                                                                                                                                                                                   &&&PureScript_Data_Maybe::Data_Maybe_applyMaybe()),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Apply::Control_Apply_apply(),
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
                                                                                                                                                                                                                                                                                                                                                                 |usd__arg2_1|
                                                                                                                                                                                                                                                                                                                                                                 Func1::new({
                                                                                                                                                                                                                                                                                                                                                                                let usd__arg2_1
                                                                                                                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                                                                                                                    usd__arg2_1.clone();
                                                                                                                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                                                                                                                    |usd__arg3_1|
                                                                                                                                                                                                                                                                                                                                                                                    &LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(usd__arg1_1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                              usd__arg2_1,
                                                                                                                                                                                                                                                                                                                                                                                                                                                              usd__arg3_1.clone()))
                                                                                                                                                                                                                                                                                                                                                                            })
                                                                                                                                                                                                                                                                                                                                                         }))),
                                                                                                                                                                                                                                                                                         &&&y_prime_1)),
                                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Applicative::Control_Applicative_pure(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Maybe::Data_Maybe_applicativeMaybe()),
                                                                                                                                                                                                                                                      &&&m_prime_1))),
                                                                                                                                             &&&d_prime_1)
                                                                                                        }),
                                                                                       add(string("Ord0"),
                                                                                           &&Func1::new(move
                                                                                                            |usd__unused|
                                                                                                            &PureScript_Data_Date::Data_Date_ordDate()),
                                                                                           empty::<string,
                                                                                                   &dyn Any>())))))
    }
    pub fn Data_Date_pred() -> &dyn Any {
        static Data_Date_pred: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Date_pred.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                        &&&PureScript_Data_Date::Data_Date_enumDate()))
    }
    pub fn Data_Date_diff() -> &dyn Any {
        static Data_Date_diff: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Date_diff.get_or_init(||
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
                                                                               |v|
                                                                               &Func1::new({
                                                                                               let v
                                                                                                   =
                                                                                                   v.clone();
                                                                                               move
                                                                                                   |v1|
                                                                                                   {
                                                                                                       let matchValue:
                                                                                                               LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                           Sharpurs_Prelude::unbox(&&v);
                                                                                                       let matchValue_1:
                                                                                                               LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                           Sharpurs_Prelude::unbox(v1);
                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_apply(),
                                                                                                                                                                           &&&toDuration),
                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn6(),
                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Date::Data_Date_calcDiff()),
                                                                                                                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                                 PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                                                                                                                              _,
                                                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                                             }),
                                                                                                                                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                                                                                                                             &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                                                                                                                                                          &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                                                                                                 {
                                                                                                                                                                                                                                                                                                                                                                 PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                                                                                              x,
                                                                                                                                                                                                                                                                                                                                                                                                                              _)
                                                                                                                                                                                                                                                                                                                                                                 =>
                                                                                                                                                                                                                                                                                                                                                                 x.clone(),
                                                                                                                                                                                                                                                                                                                                                             })),
                                                                                                                                                                                                                                                                                    &&&match matchValue.as_ref()
                                                                                                                                                                                                                                                                                           {
                                                                                                                                                                                                                                                                                           PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                                                        _,
                                                                                                                                                                                                                                                                                                                                                        x)
                                                                                                                                                                                                                                                                                           =>
                                                                                                                                                                                                                                                                                           x.clone(),
                                                                                                                                                                                                                                                                                       }),
                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                        PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                                                                                                                     _,
                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                    }),
                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                    &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                                                 &&&match matchValue_1.as_ref()
                                                                                                                                                                                                                                                        {
                                                                                                                                                                                                                                                        PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                                                     x,
                                                                                                                                                                                                                                                                                                                     _)
                                                                                                                                                                                                                                                        =>
                                                                                                                                                                                                                                                        x.clone(),
                                                                                                                                                                                                                                                    })),
                                                                                                                                                                           &&&match matchValue_1.as_ref()
                                                                                                                                                                                  {
                                                                                                                                                                                  PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                               _,
                                                                                                                                                                                                                                               x)
                                                                                                                                                                                  =>
                                                                                                                                                                                  x.clone(),
                                                                                                                                                                              }))
                                                                                                   }
                                                                                           })
                                                                       })
                                                       }))
    }
    pub fn Data_Date_day() -> &dyn Any {
        static Data_Date_day: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Date_day.get_or_init(||
                                      &Func1::new(move |v|
                                                      &match Sharpurs_Prelude::unbox(v).as_ref()
                                                           {
                                                           PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                        _,
                                                                                                                        x)
                                                           => x.clone(),
                                                       }))
    }
    pub fn Data_Date_canonicalDate() -> &dyn Any {
        static Data_Date_canonicalDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_canonicalDate.get_or_init(||
                                                &Func1::new(move |y|
                                                                &Func1::new({
                                                                                let y
                                                                                    =
                                                                                    y.clone();
                                                                                move
                                                                                    |m|
                                                                                    &Func1::new({
                                                                                                    let m
                                                                                                        =
                                                                                                        m.clone();
                                                                                                    move
                                                                                                        |d|
                                                                                                        {
                                                                                                            let mkDate =
                                                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Partial_Unsafe::Partial_Unsafe_unsafePartial(),
                                                                                                                                                 &&&Func1::new(move
                                                                                                                                                                   |usd__unused|
                                                                                                                                                                   &Func1::new(move
                                                                                                                                                                                   |y_prime|
                                                                                                                                                                                   &Func1::new({
                                                                                                                                                                                                   let y_prime
                                                                                                                                                                                                       =
                                                                                                                                                                                                       y_prime.clone();
                                                                                                                                                                                                   move
                                                                                                                                                                                                       |m_prime|
                                                                                                                                                                                                       &Func1::new({
                                                                                                                                                                                                                       let m_prime
                                                                                                                                                                                                                           =
                                                                                                                                                                                                                           m_prime.clone();
                                                                                                                                                                                                                       move
                                                                                                                                                                                                                           |d_prime|
                                                                                                                                                                                                                           &LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(&y_prime,
                                                                                                                                                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromJust(),
                                                                                                                                                                                                                                                                                                                                                                         &&&Sharpurs_Prelude::Prim_undefined()),
                                                                                                                                                                                                                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                                                                                                                            &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                                                                                                                                                                         &&&m_prime)),
                                                                                                                                                                                                                                                                                                     d_prime.clone()))
                                                                                                                                                                                                                   })
                                                                                                                                                                                               }))));
                                                                                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn4(),
                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Date::Data_Date_canonicalDateImpl()),
                                                                                                                                                                                                                                                      &&&mkDate),
                                                                                                                                                                                                                   &&&y),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumMonth()),
                                                                                                                                                                                                                   &&&m)),
                                                                                                                                             d)
                                                                                                        }
                                                                                                })
                                                                            })))
    }
    pub fn Data_Date_exactDate() -> &dyn Any {
        static Data_Date_exactDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_exactDate.get_or_init(||
                                            &Func1::new(move |y|
                                                            &Func1::new({
                                                                            let y
                                                                                =
                                                                                y.clone();
                                                                            move
                                                                                |m|
                                                                                &Func1::new({
                                                                                                let m
                                                                                                    =
                                                                                                    m.clone();
                                                                                                move
                                                                                                    |d|
                                                                                                    {
                                                                                                        let dt =
                                                                                                            &LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(&y,
                                                                                                                                                                                      &m,
                                                                                                                                                                                      d.clone()));
                                                                                                        let matchValue =
                                                                                                            Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                            &&&PureScript_Data_Date::Data_Date_eqDate()),
                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_canonicalDate(),
                                                                                                                                                                                                                                                                                                                  &&&y),
                                                                                                                                                                                                                                                                               &&&m),
                                                                                                                                                                                                                                            d)),
                                                                                                                                                                      &&&dt));
                                                                                                        match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                         &matchValue)
                                                                                                            {
                                                                                                            0_i32
                                                                                                            =>
                                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(&dt)),
                                                                                                            _
                                                                                                            =>
                                                                                                            &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor),
                                                                                                        }
                                                                                                    }
                                                                                            })
                                                                        })))
    }
    pub fn Data_Date_boundedDate() -> &dyn Any {
        static Data_Date_boundedDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_boundedDate.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                               &&&add(string("bottom"),
                                                                                      &&LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                  &&&PureScript_Data_Date_Component::Data_Date_Component_boundedYear()),
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                  &&&PureScript_Data_Date_Component::Data_Date_Component_boundedMonth()),
                                                                                                                                                                 Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                                                                                                                                                  &&&PureScript_Data_Date_Component::Data_Date_Component_boundedDay()))),
                                                                                      add(string("top"),
                                                                                          &&LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                      &&&PureScript_Data_Date_Component::Data_Date_Component_boundedYear()),
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                      &&&PureScript_Data_Date_Component::Data_Date_Component_boundedMonth()),
                                                                                                                                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                                                                                                                                                      &&&PureScript_Data_Date_Component::Data_Date_Component_boundedDay()))),
                                                                                          add(string("Ord0"),
                                                                                              &&Func1::new(move
                                                                                                               |usd__unused|
                                                                                                               &PureScript_Data_Date::Data_Date_ordDate()),
                                                                                              empty::<string,
                                                                                                      &dyn Any>())))))
    }
    pub fn Data_Date_adjust() -> &dyn Any {
        static Data_Date_adjust: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Date_adjust.get_or_init(||
                                         &Func1::new(move |v|
                                                         &Func1::new({
                                                                         let v
                                                                             =
                                                                             v.clone();
                                                                         move
                                                                             |date|
                                                                             {
                                                                                 let matchValue =
                                                                                     Sharpurs_Prelude::unbox(&&v);
                                                                                 let matchValue_1 =
                                                                                     Sharpurs_Prelude::unbox(date);
                                                                                 {
                                                                                     let adj_2 =
                                                                                         Func0::new({
                                                                                                        let adj_tco
                                                                                                            =
                                                                                                            adj_tco.clone();
                                                                                                        move
                                                                                                            ||
                                                                                                            &Func1::new({
                                                                                                                            let adj_tco
                                                                                                                                =
                                                                                                                                adj_tco.clone();
                                                                                                                            move
                                                                                                                                |v1|
                                                                                                                                Func1::new({
                                                                                                                                               let adj_tco
                                                                                                                                                   =
                                                                                                                                                   adj_tco.clone();
                                                                                                                                               let v1
                                                                                                                                                   =
                                                                                                                                                   v1.clone();
                                                                                                                                               move
                                                                                                                                                   |v2|
                                                                                                                                                   adj_tco(v1)(v2.clone())
                                                                                                                                           })
                                                                                                                        })
                                                                                                    });
                                                                                     let adj_1 =
                                                                                         Lazy(adj_2);
                                                                                     let adj_tco =
                                                                                         Func1::new({
                                                                                                        let adj_1
                                                                                                            =
                                                                                                            adj_1.clone();
                                                                                                        move
                                                                                                            |v1_1|
                                                                                                            Func1::new({
                                                                                                                           let adj_1
                                                                                                                               =
                                                                                                                               adj_1.clone();
                                                                                                                           let v1_1
                                                                                                                               =
                                                                                                                               v1_1.clone();
                                                                                                                           move
                                                                                                                               |v2_1|
                                                                                                                               {
                                                                                                                                   let matchValue_3 =
                                                                                                                                       Sharpurs_Prelude::unbox(&&v1_1);
                                                                                                                                   let matchValue_4:
                                                                                                                                           LrcPtr<PureScript_Data_Date::Data_Date_Date> =
                                                                                                                                       Sharpurs_Prelude::unbox(v2_1);
                                                                                                                                   match &Sharpurs_Prelude::_007cLitInt_007c__007c(0_i32,
                                                                                                                                                                                   &matchValue_3)
                                                                                                                                       {
                                                                                                                                       0_i32
                                                                                                                                       =>
                                                                                                                                       &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(dt.clone())),
                                                                                                                                       _
                                                                                                                                       =>
                                                                                                                                       {
                                                                                                                                           let y =
                                                                                                                                               match matchValue_4.as_ref()
                                                                                                                                                   {
                                                                                                                                                   PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(x,
                                                                                                                                                                                                                _,
                                                                                                                                                                                                                _)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                               };
                                                                                                                                           let m =
                                                                                                                                               match matchValue_4.as_ref()
                                                                                                                                                   {
                                                                                                                                                   PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                x,
                                                                                                                                                                                                                _)
                                                                                                                                                   =>
                                                                                                                                                   x.clone(),
                                                                                                                                               };
                                                                                                                                           let j =
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                   &&&matchValue_3),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                                                                                                                   &&&match matchValue_4.as_ref()
                                                                                                                                                                                                                          {
                                                                                                                                                                                                                          PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(_,
                                                                                                                                                                                                                                                                                       _,
                                                                                                                                                                                                                                                                                       x)
                                                                                                                                                                                                                          =>
                                                                                                                                                                                                                          x.clone(),
                                                                                                                                                                                                                      }));
                                                                                                                                           let low =
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThan(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                   &&&j),
                                                                                                                                                                                &&&1_i32);
                                                                                                                                           let l =
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Date::Data_Date_lastDayOfMonth(),
                                                                                                                                                                                                                   &&&y),
                                                                                                                                                                                &&{
                                                                                                                                                                                      let matchValue_6 =
                                                                                                                                                                                          Sharpurs_Prelude::unbox(&&low);
                                                                                                                                                                                      match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                                                                       &matchValue_6)
                                                                                                                                                                                          {
                                                                                                                                                                                          0_i32
                                                                                                                                                                                          =>
                                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromMaybe(),
                                                                                                                                                                                                                                                              &&&LrcPtr::new(Data_Date_Component_Month::Data_Date_Component_Decemberusd_Ctor)),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_pred(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Date_Component::Data_Date_Component_enumMonth()),
                                                                                                                                                                                                                                                              &&&m)),
                                                                                                                                                                                          _
                                                                                                                                                                                          =>
                                                                                                                                                                                          &m,
                                                                                                                                                                                      }
                                                                                                                                                                                  });
                                                                                                                                           let hi =
                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThan(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                   &&&j),
                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                      &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                                                                                                                   &&&l));
                                                                                                                                           let i_prime =
                                                                                                                                               {
                                                                                                                                                   let matchValue_7 =
                                                                                                                                                       &string("MissingExpr");
                                                                                                                                                   if Sharpurs_Prelude::unbox(&&low)
                                                                                                                                                      {
                                                                                                                                                       &j
                                                                                                                                                   } else {
                                                                                                                                                       if Sharpurs_Prelude::unbox(&&hi)
                                                                                                                                                          {
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                  &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                                                                     &&&j),
                                                                                                                                                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                                                                                                                                                                                                        &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                                                                                                                                                                                                     &&&l))),
                                                                                                                                                                                            &&&1_i32)
                                                                                                                                                       } else {
                                                                                                                                                           if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                              {
                                                                                                                                                               &0_i32
                                                                                                                                                           } else {
                                                                                                                                                               panic!("{}",
                                                                                                                                                                      LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Date.fs"),
                                  Data1: 53_i32,
                                  Data2: 1686_i32,}).get_Message(),)
                                                                                                                                                           }
                                                                                                                                                       }
                                                                                                                                                   }
                                                                                                                                               };
                                                                                                                                           let dt_prime =
                                                                                                                                               {
                                                                                                                                                   let matchValue_8 =
                                                                                                                                                       &string("MissingExpr");
                                                                                                                                                   if Sharpurs_Prelude::unbox(&&low)
                                                                                                                                                      {
                                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                                                                                           &&&PureScript_Data_Date::Data_Date_pred()),
                                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                                              &&&Func1::new({
                                                                                                                                                                                                                                                                                let m
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    m.clone();
                                                                                                                                                                                                                                                                                let y
                                                                                                                                                                                                                                                                                    =
                                                                                                                                                                                                                                                                                    y.clone();
                                                                                                                                                                                                                                                                                move
                                                                                                                                                                                                                                                                                    |usd__arg1|
                                                                                                                                                                                                                                                                                    &LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(&y,
                                                                                                                                                                                                                                                                                                                                                              &m,
                                                                                                                                                                                                                                                                                                                                                              usd__arg1.clone()))
                                                                                                                                                                                                                                                                            })),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                                                                                                                                                              &&&1_i32)))
                                                                                                                                                   } else {
                                                                                                                                                       if Sharpurs_Prelude::unbox(&&hi)
                                                                                                                                                          {
                                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_succ(),
                                                                                                                                                                                                                               &&&PureScript_Data_Date::Data_Date_enumDate()),
                                                                                                                                                                                            &&&LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(&y,
                                                                                                                                                                                                                                                                        &m,
                                                                                                                                                                                                                                                                        &l)))
                                                                                                                                                       } else {
                                                                                                                                                           if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                                                                                              {
                                                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_Maybe::Data_Maybe_functorMaybe()),
                                                                                                                                                                                                                                   &&&Func1::new({
                                                                                                                                                                                                                                                     let m
                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                         m.clone();
                                                                                                                                                                                                                                                     let y
                                                                                                                                                                                                                                                         =
                                                                                                                                                                                                                                                         y.clone();
                                                                                                                                                                                                                                                     move
                                                                                                                                                                                                                                                         |usd__arg1_1|
                                                                                                                                                                                                                                                         &LrcPtr::new(PureScript_Data_Date::Data_Date_Date::Data_Date_Dateusd_Ctor(&y,
                                                                                                                                                                                                                                                                                                                                   &m,
                                                                                                                                                                                                                                                                                                                                   usd__arg1_1.clone()))
                                                                                                                                                                                                                                                 })),
                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                                                                                                                                      &&&PureScript_Data_Date_Component::Data_Date_Component_boundedEnumDay()),
                                                                                                                                                                                                                                   &&&j))
                                                                                                                                                           } else {
                                                                                                                                                               panic!("{}",
                                                                                                                                                                      LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Date.fs"),
                                  Data1: 53_i32,
                                  Data2: 2311_i32,}).get_Message(),)
                                                                                                                                                           }
                                                                                                                                                       }
                                                                                                                                                   }
                                                                                                                                               };
                                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bindFlipped(),
                                                                                                                                                                                                                                                  &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&adj_1.Value,
                                                                                                                                                                                                                                                  &&&i_prime)),
                                                                                                                                                                            &&&dt_prime)
                                                                                                                                       }
                                                                                                                                   }
                                                                                                                               }
                                                                                                                       })
                                                                                                    });
                                                                                     let adj =
                                                                                         adj_1.Value;
                                                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Bind::Control_Bind_bind(),
                                                                                                                                                                                            &&&PureScript_Data_Maybe::Data_Maybe_bindMaybe()),
                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_fromNumber(),
                                                                                                                                                                                            &&&matchValue)),
                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function::Data_Function_flip(),
                                                                                                                                                                                            &&&adj),
                                                                                                                                                         &&&matchValue_1))
                                                                                 }
                                                                             }
                                                                     })))
    }
}

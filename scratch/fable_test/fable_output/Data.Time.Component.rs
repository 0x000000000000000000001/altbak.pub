pub mod PureScript_Data_Time_Component {
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
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_ab5378f8::PureScript_Data_Semigroup;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    use fable_library_rust::System::Lazy_1;
    pub fn Data_Time_Component_Second() -> &dyn Any {
        static Data_Time_Component_Second: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Second.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Time_Component_Minute() -> &dyn Any {
        static Data_Time_Component_Minute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Minute.get_or_init(||
                                                   &Func1::new(move |x|
                                                                   x.clone()))
    }
    pub fn Data_Time_Component_Millisecond() -> &dyn Any {
        static Data_Time_Component_Millisecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Millisecond.get_or_init(||
                                                        &Func1::new(move |x|
                                                                        x.clone()))
    }
    pub fn Data_Time_Component_Hour() -> &dyn Any {
        static Data_Time_Component_Hour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_Hour.get_or_init(||
                                                 &Func1::new(move |x|
                                                                 x.clone()))
    }
    pub fn Data_Time_Component_showSecond() -> &dyn Any {
        static Data_Time_Component_showSecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_showSecond.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                        &&&add(string("show"),
                                                                                               &&Func1::new(move
                                                                                                                |v|
                                                                                                                {
                                                                                                                    let m =
                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                        &&&string("(Second ")),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                                                                                                              &&&m)),
                                                                                                                                                                                        &&&string(")")))
                                                                                                                }),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Time_Component_showMinute() -> &dyn Any {
        static Data_Time_Component_showMinute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_showMinute.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                        &&&add(string("show"),
                                                                                               &&Func1::new(move
                                                                                                                |v|
                                                                                                                {
                                                                                                                    let m =
                                                                                                                        Sharpurs_Prelude::unbox(v);
                                                                                                                    Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                           &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                        &&&string("(Minute ")),
                                                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                              &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                                                                                                              &&&m)),
                                                                                                                                                                                        &&&string(")")))
                                                                                                                }),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Time_Component_showMillisecond() -> &dyn Any {
        static Data_Time_Component_showMillisecond: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Time_Component_showMillisecond.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                             &&&add(string("show"),
                                                                                                    &&Func1::new(move
                                                                                                                     |v|
                                                                                                                     {
                                                                                                                         let m =
                                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                                         Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                             &&&string("(Millisecond ")),
                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                                      &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                                                                                                                   &&&m)),
                                                                                                                                                                                             &&&string(")")))
                                                                                                                     }),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>())))
    }
    pub fn Data_Time_Component_showHour() -> &dyn Any {
        static Data_Time_Component_showHour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_showHour.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                                      &&&add(string("show"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              {
                                                                                                                  let h =
                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                  Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                         &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                      &&&string("(Hour ")),
                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semigroup::Data_Semigroup_append(),
                                                                                                                                                                                                                                                            &&&PureScript_Data_Semigroup::Data_Semigroup_semigroupString()),
                                                                                                                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_show(),
                                                                                                                                                                                                                                                                                               &&&PureScript_Data_Show::Data_Show_showInt()),
                                                                                                                                                                                                                                                            &&&h)),
                                                                                                                                                                                      &&&string(")")))
                                                                                                              }),
                                                                                             empty::<string,
                                                                                                     &dyn Any>())))
    }
    pub fn Data_Time_Component_ordSecond() -> &dyn Any {
        static Data_Time_Component_ordSecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_ordSecond.get_or_init(||
                                                      &PureScript_Data_Ord::Data_Ord_ordInt())
    }
    pub fn Data_Time_Component_ordMinute() -> &dyn Any {
        static Data_Time_Component_ordMinute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_ordMinute.get_or_init(||
                                                      &PureScript_Data_Ord::Data_Ord_ordInt())
    }
    pub fn Data_Time_Component_ordMillisecond() -> &dyn Any {
        static Data_Time_Component_ordMillisecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_ordMillisecond.get_or_init(||
                                                           &PureScript_Data_Ord::Data_Ord_ordInt())
    }
    pub fn Data_Time_Component_ordHour() -> &dyn Any {
        static Data_Time_Component_ordHour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_ordHour.get_or_init(||
                                                    &PureScript_Data_Ord::Data_Ord_ordInt())
    }
    pub fn Data_Time_Component_eqSecond() -> &dyn Any {
        static Data_Time_Component_eqSecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_eqSecond.get_or_init(||
                                                     &PureScript_Data_Eq::Data_Eq_eqInt())
    }
    pub fn Data_Time_Component_eqMinute() -> &dyn Any {
        static Data_Time_Component_eqMinute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_eqMinute.get_or_init(||
                                                     &PureScript_Data_Eq::Data_Eq_eqInt())
    }
    pub fn Data_Time_Component_eqMillisecond() -> &dyn Any {
        static Data_Time_Component_eqMillisecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_eqMillisecond.get_or_init(||
                                                          &PureScript_Data_Eq::Data_Eq_eqInt())
    }
    pub fn Data_Time_Component_eqHour() -> &dyn Any {
        static Data_Time_Component_eqHour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_eqHour.get_or_init(||
                                                   &PureScript_Data_Eq::Data_Eq_eqInt())
    }
    pub fn Data_Time_Component_boundedSecond() -> &dyn Any {
        static Data_Time_Component_boundedSecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_boundedSecond.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                           &&&add(string("bottom"),
                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Second(),
                                                                                                                                    &&&0_i32),
                                                                                                  add(string("top"),
                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Second(),
                                                                                                                                        &&&59_i32),
                                                                                                      add(string("Ord0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Time_Component::Data_Time_Component_ordSecond()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))))
    }
    pub fn Data_Time_Component_boundedMinute() -> &dyn Any {
        static Data_Time_Component_boundedMinute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_boundedMinute.get_or_init(||
                                                          Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                           &&&add(string("bottom"),
                                                                                                  &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Minute(),
                                                                                                                                    &&&0_i32),
                                                                                                  add(string("top"),
                                                                                                      &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Minute(),
                                                                                                                                        &&&59_i32),
                                                                                                      add(string("Ord0"),
                                                                                                          &&Func1::new(move
                                                                                                                           |usd__unused|
                                                                                                                           &PureScript_Data_Time_Component::Data_Time_Component_ordMinute()),
                                                                                                          empty::<string,
                                                                                                                  &dyn Any>())))))
    }
    pub fn Data_Time_Component_boundedMillisecond() -> &dyn Any {
        static Data_Time_Component_boundedMillisecond:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_boundedMillisecond.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                                &&&add(string("bottom"),
                                                                                                       &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Millisecond(),
                                                                                                                                         &&&0_i32),
                                                                                                       add(string("top"),
                                                                                                           &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Millisecond(),
                                                                                                                                             &&&999_i32),
                                                                                                           add(string("Ord0"),
                                                                                                               &&Func1::new(move
                                                                                                                                |usd__unused|
                                                                                                                                &PureScript_Data_Time_Component::Data_Time_Component_ordMillisecond()),
                                                                                                               empty::<string,
                                                                                                                       &dyn Any>())))))
    }
    pub fn Data_Time_Component_boundedHour() -> &dyn Any {
        static Data_Time_Component_boundedHour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_boundedHour.get_or_init(||
                                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                         &&&add(string("bottom"),
                                                                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Hour(),
                                                                                                                                  &&&0_i32),
                                                                                                add(string("top"),
                                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Hour(),
                                                                                                                                      &&&23_i32),
                                                                                                    add(string("Ord0"),
                                                                                                        &&Func1::new(move
                                                                                                                         |usd__unused|
                                                                                                                         &PureScript_Data_Time_Component::Data_Time_Component_ordHour()),
                                                                                                        empty::<string,
                                                                                                                &dyn Any>())))))
    }
    pub fn Data_Time_Component_boundedEnumSecond_004047() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&60_i32),
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
                                                                                                                                                                                                                &&&0_i32)),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                &&&n1),
                                                                                                                                                                             &&&59_i32)))
                                                                            }
                                                                            {
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Second(),
                                                                                                                                                                     &&&matchValue)))
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                             } else {
                                                                                 panic!("{}",
                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Time.Component.fs"),
                                  Data1: 47_i32,
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
                                                                             &PureScript_Data_Time_Component::Data_Time_Component_boundedSecond()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Time_Component_enumSecond_004048_002d1
                                                                                     =
                                                                                     Data_Time_Component_enumSecond_004048_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Time_Component_enumSecond_004048_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Time_Component_boundedEnumSecond_004047_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_boundedEnumSecond_004047_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_boundedEnumSecond_004047_002d1.get_or_init(||
                                                                           Lazy(Data_Time_Component_boundedEnumSecond_004047.clone()))
    }
    pub fn Data_Time_Component_enumSecond_004048() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumSecond_004047_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumSecond_004047_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumSecond_004047_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumSecond_004047_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Time_Component::Data_Time_Component_ordSecond()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Time_Component_enumSecond_004048_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_enumSecond_004048_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_enumSecond_004048_002d1.get_or_init(||
                                                                    Lazy(Data_Time_Component_enumSecond_004048.clone()))
    }
    pub fn Data_Time_Component_boundedEnumSecond() -> &dyn Any {
        static Data_Time_Component_boundedEnumSecond:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_boundedEnumSecond.get_or_init(||
                                                              Data_Time_Component_boundedEnumSecond_004047_002d1.Value)
    }
    pub fn Data_Time_Component_enumSecond() -> &dyn Any {
        static Data_Time_Component_enumSecond: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_enumSecond.get_or_init(||
                                                       Data_Time_Component_enumSecond_004048_002d1.Value)
    }
    pub fn Data_Time_Component_boundedEnumMinute_004051() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&60_i32),
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
                                                                                                                                                                                                                &&&0_i32)),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                &&&n1),
                                                                                                                                                                             &&&59_i32)))
                                                                            }
                                                                            {
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Minute(),
                                                                                                                                                                     &&&matchValue)))
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                             } else {
                                                                                 panic!("{}",
                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Time.Component.fs"),
                                  Data1: 51_i32,
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
                                                                             &PureScript_Data_Time_Component::Data_Time_Component_boundedMinute()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Time_Component_enumMinute_004052_002d1
                                                                                     =
                                                                                     Data_Time_Component_enumMinute_004052_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Time_Component_enumMinute_004052_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Time_Component_boundedEnumMinute_004051_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_boundedEnumMinute_004051_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_boundedEnumMinute_004051_002d1.get_or_init(||
                                                                           Lazy(Data_Time_Component_boundedEnumMinute_004051.clone()))
    }
    pub fn Data_Time_Component_enumMinute_004052() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumMinute_004051_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumMinute_004051_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumMinute_004051_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumMinute_004051_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Time_Component::Data_Time_Component_ordMinute()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Time_Component_enumMinute_004052_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_enumMinute_004052_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_enumMinute_004052_002d1.get_or_init(||
                                                                    Lazy(Data_Time_Component_enumMinute_004052.clone()))
    }
    pub fn Data_Time_Component_boundedEnumMinute() -> &dyn Any {
        static Data_Time_Component_boundedEnumMinute:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_boundedEnumMinute.get_or_init(||
                                                              Data_Time_Component_boundedEnumMinute_004051_002d1.Value)
    }
    pub fn Data_Time_Component_enumMinute() -> &dyn Any {
        static Data_Time_Component_enumMinute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_enumMinute.get_or_init(||
                                                       Data_Time_Component_enumMinute_004052_002d1.Value)
    }
    pub fn Data_Time_Component_boundedEnumMillisecond_004055() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&1000_i32),
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
                                                                                                                                                                                                                &&&0_i32)),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                &&&n1),
                                                                                                                                                                             &&&999_i32)))
                                                                            }
                                                                            {
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Millisecond(),
                                                                                                                                                                     &&&matchValue)))
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                             } else {
                                                                                 panic!("{}",
                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Time.Component.fs"),
                                  Data1: 55_i32,
                                  Data2: 288_i32,}).get_Message(),)
                                                                             }
                                                                         }
                                                                     }),
                                                    add(string("fromEnum"),
                                                        &&Func1::new(move |v|
                                                                         &Sharpurs_Prelude::unbox(v)),
                                                        add(string("Bounded0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Time_Component::Data_Time_Component_boundedMillisecond()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Time_Component_enumMillisecond_004056_002d1
                                                                                     =
                                                                                     Data_Time_Component_enumMillisecond_004056_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Time_Component_enumMillisecond_004056_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Time_Component_boundedEnumMillisecond_004055_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_boundedEnumMillisecond_004055_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_boundedEnumMillisecond_004055_002d1.get_or_init(||
                                                                                Lazy(Data_Time_Component_boundedEnumMillisecond_004055.clone()))
    }
    pub fn Data_Time_Component_enumMillisecond_004056() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumMillisecond_004055_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumMillisecond_004055_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumMillisecond_004055_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumMillisecond_004055_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Time_Component::Data_Time_Component_ordMillisecond()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Time_Component_enumMillisecond_004056_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_enumMillisecond_004056_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_enumMillisecond_004056_002d1.get_or_init(||
                                                                         Lazy(Data_Time_Component_enumMillisecond_004056.clone()))
    }
    pub fn Data_Time_Component_boundedEnumMillisecond() -> &dyn Any {
        static Data_Time_Component_boundedEnumMillisecond:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_boundedEnumMillisecond.get_or_init(||
                                                                   Data_Time_Component_boundedEnumMillisecond_004055_002d1.Value)
    }
    pub fn Data_Time_Component_enumMillisecond() -> &dyn Any {
        static Data_Time_Component_enumMillisecond: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Time_Component_enumMillisecond.get_or_init(||
                                                            Data_Time_Component_enumMillisecond_004056_002d1.Value)
    }
    pub fn Data_Time_Component_boundedEnumHour_004059() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_BoundedEnumusd_Dict(),
                                         &&&add(string("cardinality"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Cardinality(),
                                                                                  &&&24_i32),
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
                                                                                                                                                                                                                &&&0_i32)),
                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                                   &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                                &&&n1),
                                                                                                                                                                             &&&23_i32)))
                                                                            }
                                                                            {
                                                                             &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Time_Component::Data_Time_Component_Hour(),
                                                                                                                                                                     &&&matchValue)))
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                                             } else {
                                                                                 panic!("{}",
                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Time.Component.fs"),
                                  Data1: 59_i32,
                                  Data2: 279_i32,}).get_Message(),)
                                                                             }
                                                                         }
                                                                     }),
                                                    add(string("fromEnum"),
                                                        &&Func1::new(move |v|
                                                                         &Sharpurs_Prelude::unbox(v)),
                                                        add(string("Bounded0"),
                                                            &&Func1::new(move
                                                                             |usd__unused|
                                                                             &PureScript_Data_Time_Component::Data_Time_Component_boundedHour()),
                                                            add(string("Enum1"),
                                                                &&Func1::new({
                                                                                 let Data_Time_Component_enumHour_004060_002d1
                                                                                     =
                                                                                     Data_Time_Component_enumHour_004060_002d1.clone();
                                                                                 move
                                                                                     |usd__unused_1|
                                                                                     &Data_Time_Component_enumHour_004060_002d1.Value
                                                                             }),
                                                                empty::<string,
                                                                        &dyn Any>()))))))
    }
    pub fn Data_Time_Component_boundedEnumHour_004059_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_boundedEnumHour_004059_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_boundedEnumHour_004059_002d1.get_or_init(||
                                                                         Lazy(Data_Time_Component_boundedEnumHour_004059.clone()))
    }
    pub fn Data_Time_Component_enumHour_004060() -> &dyn Any {
        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_Enumusd_Dict(),
                                         &&&add(string("succ"),
                                                &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumHour_004059_002d1.Value)),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                           &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                        &&&Func1::new(move
                                                                                                                                                                          |v|
                                                                                                                                                                          Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                 &&&PureScript_Data_Semiring::Data_Semiring_semiringInt()),
                                                                                                                                                                                                                                              v),
                                                                                                                                                                                                           &&&1_i32))),
                                                                                                                     &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                        &&&Data_Time_Component_boundedEnumHour_004059_002d1.Value))),
                                                add(string("pred"),
                                                    &Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_toEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumHour_004059_002d1.Value)),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                               &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                            &&&Func1::new(move
                                                                                                                                                                              |v_1|
                                                                                                                                                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                     &&&PureScript_Data_Ring::Data_Ring_ringInt()),
                                                                                                                                                                                                                                                  v_1),
                                                                                                                                                                                                               &&&1_i32))),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Enum::Data_Enum_fromEnum(),
                                                                                                                                                            &&&Data_Time_Component_boundedEnumHour_004059_002d1.Value))),
                                                    add(string("Ord0"),
                                                        &&Func1::new(move
                                                                         |usd__unused|
                                                                         &PureScript_Data_Time_Component::Data_Time_Component_ordHour()),
                                                        empty::<string,
                                                                &dyn Any>()))))
    }
    pub fn Data_Time_Component_enumHour_004060_002d1()
     -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_Time_Component_enumHour_004060_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_Time_Component_enumHour_004060_002d1.get_or_init(||
                                                                  Lazy(Data_Time_Component_enumHour_004060.clone()))
    }
    pub fn Data_Time_Component_boundedEnumHour() -> &dyn Any {
        static Data_Time_Component_boundedEnumHour: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Time_Component_boundedEnumHour.get_or_init(||
                                                            Data_Time_Component_boundedEnumHour_004059_002d1.Value)
    }
    pub fn Data_Time_Component_enumHour() -> &dyn Any {
        static Data_Time_Component_enumHour: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Time_Component_enumHour.get_or_init(||
                                                     Data_Time_Component_enumHour_004060_002d1.Value)
    }
}

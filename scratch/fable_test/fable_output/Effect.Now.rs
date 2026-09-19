pub mod PureScript_Effect_Now {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_47c4f61::PureScript_Data_DateTime_Instant;
    use crate::module_867835b4::PureScript_Data_DateTime;
    use crate::module_b38d0a2::PureScript_Data_Functor;
    use crate::module_5706e1fc::PureScript_Effect;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Effect_Now_FFI {
        use super::*;
        use fable_library_rust::DateTimeOffset_::DateTimeOffset;
        use fable_library_rust::Native_::defaultOf;
        pub fn now(dummy: &dyn Any) -> &dyn Any {
            &({
                  let copyOfStruct: DateTimeOffset = DateTimeOffset::utcNow();
                  copyOfStruct.toUnixTimeMilliseconds()
              } as f64)
        }
        pub fn getTimezoneOffset(dummy: &dyn Any) -> &dyn Any {
            &-defaultOf::<&dyn Any>().total_minutes()
        }
    }
    pub fn Effect_Now_getTimezoneOffset() -> &dyn Any {
        static Effect_Now_getTimezoneOffset: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Now_getTimezoneOffset.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &PureScript_Effect_Now::Effect_Now_FFI::getTimezoneOffset(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Now_now() -> &dyn Any {
        static Effect_Now_now: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Effect_Now_now.get_or_init(||
                                       &Func1::new(move |arg0|
                                                       &PureScript_Effect_Now::Effect_Now_FFI::now(&Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Effect_Now_nowTime() -> &dyn Any {
        static Effect_Now_nowTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Now_nowTime.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Effect::Effect_functorEffect()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                     &&&PureScript_Data_DateTime::Data_DateTime_time()),
                                                                                                                                                  &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_toDateTime())),
                                                                            &&&PureScript_Effect_Now::Effect_Now_now()))
    }
    pub fn Effect_Now_nowDateTime() -> &dyn Any {
        static Effect_Now_nowDateTime: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Now_nowDateTime.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                      &&&PureScript_Effect::Effect_functorEffect()),
                                                                                                                   &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_toDateTime()),
                                                                                &&&PureScript_Effect_Now::Effect_Now_now()))
    }
    pub fn Effect_Now_nowDate() -> &dyn Any {
        static Effect_Now_nowDate: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Effect_Now_nowDate.get_or_init(||
                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Functor::Data_Functor_map(),
                                                                                                                                                  &&&PureScript_Effect::Effect_functorEffect()),
                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                                                                                     &&&PureScript_Data_DateTime::Data_DateTime_date()),
                                                                                                                                                  &&&PureScript_Data_DateTime_Instant::Data_DateTime_Instant_toDateTime())),
                                                                            &&&PureScript_Effect_Now::Effect_Now_now()))
    }
}

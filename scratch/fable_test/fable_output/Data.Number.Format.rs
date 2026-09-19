pub mod PureScript_Data_Number_Format {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_Number_Format_FFI {
        use super::*;
        pub fn toPrecisionNative(dVal: &dyn Any, numVal: &dyn Any)
         -> &dyn Any {
            &String_(numVal.clone())
        }
        pub fn toFixedNative(dVal: &dyn Any, numVal: &dyn Any) -> &dyn Any {
            &String_(numVal.clone())
        }
        pub fn toExponentialNative(dVal: &dyn Any, numVal: &dyn Any)
         -> &dyn Any {
            &String_(numVal.clone())
        }
        pub fn toString(numVal: &dyn Any) -> &dyn Any {
            &String_(numVal.clone())
        }
    }
    pub fn Data_Number_Format_toExponentialNative() -> &dyn Any {
        static Data_Number_Format_toExponentialNative:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_toExponentialNative.get_or_init(||
                                                               &Func1::new(move
                                                                               |dVal|
                                                                               Func1::new({
                                                                                              let dVal
                                                                                                  =
                                                                                                  dVal.clone();
                                                                                              move
                                                                                                  |numVal|
                                                                                                  PureScript_Data_Number_Format::Data_Number_Format_FFI::toExponentialNative(&dVal,
                                                                                                                                                                             numVal)
                                                                                          })))
    }
    pub fn Data_Number_Format_toFixedNative() -> &dyn Any {
        static Data_Number_Format_toFixedNative: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_toFixedNative.get_or_init(||
                                                         &Func1::new(move
                                                                         |dVal|
                                                                         Func1::new({
                                                                                        let dVal
                                                                                            =
                                                                                            dVal.clone();
                                                                                        move
                                                                                            |numVal|
                                                                                            PureScript_Data_Number_Format::Data_Number_Format_FFI::toFixedNative(&dVal,
                                                                                                                                                                 numVal)
                                                                                    })))
    }
    pub fn Data_Number_Format_toPrecisionNative() -> &dyn Any {
        static Data_Number_Format_toPrecisionNative: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_Number_Format_toPrecisionNative.get_or_init(||
                                                             &Func1::new(move
                                                                             |dVal|
                                                                             Func1::new({
                                                                                            let dVal
                                                                                                =
                                                                                                dVal.clone();
                                                                                            move
                                                                                                |numVal|
                                                                                                PureScript_Data_Number_Format::Data_Number_Format_FFI::toPrecisionNative(&dVal,
                                                                                                                                                                         numVal)
                                                                                        })))
    }
    pub fn Data_Number_Format_toString() -> &dyn Any {
        static Data_Number_Format_toString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_toString.get_or_init(||
                                                    &Func1::new(move |numVal|
                                                                    PureScript_Data_Number_Format::Data_Number_Format_FFI::toString(numVal)))
    }
    #[derive(Clone, Debug,)]
    pub enum Data_Number_Format_Format {
        Data_Number_Format_Precisionusd_Ctor(&dyn Any),
        Data_Number_Format_Fixedusd_Ctor(&dyn Any),
        Data_Number_Format_Exponentialusd_Ctor(&dyn Any),
    }
    impl core::fmt::Display for
     PureScript_Data_Number_Format::Data_Number_Format_Format {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Number_Format_Precision() -> &dyn Any {
        static Data_Number_Format_Precision: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_Precision.get_or_init(||
                                                     &Func1::new(move
                                                                     |usd__arg1|
                                                                     &LrcPtr::new(PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Precisionusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Number_Format_Fixed() -> &dyn Any {
        static Data_Number_Format_Fixed: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_Fixed.get_or_init(||
                                                 &Func1::new(move |usd__arg1|
                                                                 &LrcPtr::new(PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Fixedusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Number_Format_Exponential() -> &dyn Any {
        static Data_Number_Format_Exponential: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_Exponential.get_or_init(||
                                                       &Func1::new(move
                                                                       |usd__arg1|
                                                                       &LrcPtr::new(PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Exponentialusd_Ctor(usd__arg1.clone()))))
    }
    pub fn Data_Number_Format_toStringWith() -> &dyn Any {
        static Data_Number_Format_toStringWith: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_toStringWith.get_or_init(||
                                                        &Func1::new(move |v|
                                                                        {
                                                                            let matchValue:
                                                                                    LrcPtr<PureScript_Data_Number_Format::Data_Number_Format_Format> =
                                                                                Sharpurs_Prelude::unbox(v);
                                                                            match matchValue.as_ref()
                                                                                {
                                                                                PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Fixedusd_Ctor(matchValue_1_0)
                                                                                =>
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number_Format::Data_Number_Format_toFixedNative(),
                                                                                                                 &&matchValue_1_0),
                                                                                PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Exponentialusd_Ctor(matchValue_2_0)
                                                                                =>
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number_Format::Data_Number_Format_toExponentialNative(),
                                                                                                                 &&matchValue_2_0),
                                                                                PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Precisionusd_Ctor(matchValue_0_0)
                                                                                =>
                                                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number_Format::Data_Number_Format_toPrecisionNative(),
                                                                                                                 &&matchValue_0_0),
                                                                            }
                                                                        }))
    }
    pub fn Data_Number_Format_precision() -> &dyn Any {
        static Data_Number_Format_precision: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_precision.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                            &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                         &&&Func1::new(move
                                                                                                                                           |usd__arg1|
                                                                                                                                           &LrcPtr::new(PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Precisionusd_Ctor(usd__arg1.clone())))),
                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_clamp(),
                                                                                                                                                                                               &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                            &&&1_i32),
                                                                                                                         &&&21_i32)))
    }
    pub fn Data_Number_Format_fixed() -> &dyn Any {
        static Data_Number_Format_fixed: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_fixed.get_or_init(||
                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                        &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                     &&&Func1::new(move
                                                                                                                                       |usd__arg1|
                                                                                                                                       &LrcPtr::new(PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Fixedusd_Ctor(usd__arg1.clone())))),
                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_clamp(),
                                                                                                                                                                                           &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                        &&&0_i32),
                                                                                                                     &&&20_i32)))
    }
    pub fn Data_Number_Format_exponential() -> &dyn Any {
        static Data_Number_Format_exponential: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Format_exponential.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                                           &&&Func1::new(move
                                                                                                                                             |usd__arg1|
                                                                                                                                             &LrcPtr::new(PureScript_Data_Number_Format::Data_Number_Format_Format::Data_Number_Format_Exponentialusd_Ctor(usd__arg1.clone())))),
                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_clamp(),
                                                                                                                                                                                                 &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                              &&&0_i32),
                                                                                                                           &&&20_i32)))
    }
}

pub mod PureScript_Data_Number {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use crate::module_d2b7e5fc::PureScript_Data_Function_Uncurried;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub mod Data_Number_FFI {
        use super::*;
        use fable_library_rust::Convert_::tryParse;
        use fable_library_rust::Native_::defaultOf;
        use fable_library_rust::Native_::max as max_1;
        use fable_library_rust::Native_::min as min_1;
        pub fn nan() -> &dyn Any {
            static nan: MutCell<Option<&dyn Any>> = MutCell::new(None);
            nan.get_or_init(|| &f64::NAN)
        }
        pub fn isNaN(x: &dyn Any) -> &dyn Any { &x.is_nan() }
        pub fn infinity() -> &dyn Any {
            static infinity: MutCell<Option<&dyn Any>> = MutCell::new(None);
            infinity.get_or_init(|| &f64::INFINITY)
        }
        pub fn isFinite(x: &dyn Any) -> &dyn Any {
            &if !x.is_infinite() { !x.is_nan() } else { false }
        }
        pub fn fromStringImpl(strVal: &dyn Any, isFiniteFn: &dyn Any,
                              just: &dyn Any, nothing: &dyn Any) -> &dyn Any {
            let matchValue: LrcPtr<(bool, f64)> =
                {
                    let outArg: MutCell<f64> = MutCell::new(0.0_f64);
                    LrcPtr::new((tryParse(strVal.clone(), 167_i32, &outArg),
                                 outArg.get()))
                };
            if matchValue.0.clone() {
                let num: f64 = matchValue.1.clone();
                if isFiniteFn(&num) { just(&num) } else { nothing.clone() }
            } else { nothing.clone() }
        }
        pub fn abs(x: &dyn Any) -> &dyn Any { &x.abs() }
        pub fn acos(x: &dyn Any) -> &dyn Any { &x.acos() }
        pub fn asin(x: &dyn Any) -> &dyn Any { &x.asin() }
        pub fn atan(x: &dyn Any) -> &dyn Any { &x.atan() }
        pub fn atan2(y: &dyn Any, x: &dyn Any) -> &dyn Any {
            &y.atan2(x.clone())
        }
        pub fn ceil(x: &dyn Any) -> &dyn Any { &x.ceil() }
        pub fn cos(x: &dyn Any) -> &dyn Any { &x.cos() }
        pub fn exp(x: &dyn Any) -> &dyn Any { &x.exp() }
        pub fn floor(x: &dyn Any) -> &dyn Any { &x.floor() }
        pub fn log(x: &dyn Any) -> &dyn Any { &x.ln() }
        pub fn max(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &max_1(n1.clone(), n2.clone())
        }
        pub fn min(n1: &dyn Any, n2: &dyn Any) -> &dyn Any {
            &min_1(n1.clone(), n2.clone())
        }
        pub fn pow(n: &dyn Any, p: &dyn Any) -> &dyn Any { &n.pow(p.clone()) }
        pub fn remainder(n: &dyn Any, m: &dyn Any) -> &dyn Any {
            &(n.clone() % m.clone())
        }
        pub fn round(x: &dyn Any) -> &dyn Any { &defaultOf::<&dyn Any>() }
        pub fn sign(x: &dyn Any) -> &dyn Any { &x.signum() }
        pub fn sin(x: &dyn Any) -> &dyn Any { &x.sin() }
        pub fn sqrt(x: &dyn Any) -> &dyn Any { &x.sqrt() }
        pub fn tan(x: &dyn Any) -> &dyn Any { &x.tan() }
        pub fn trunc(x: &dyn Any) -> &dyn Any { &x.trunc() }
    }
    pub fn Data_Number_abs() -> &dyn Any {
        static Data_Number_abs: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_abs.get_or_init(||
                                        &Func1::new(move |x|
                                                        PureScript_Data_Number::Data_Number_FFI::abs(x)))
    }
    pub fn Data_Number_acos() -> &dyn Any {
        static Data_Number_acos: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_acos.get_or_init(||
                                         &Func1::new(move |x|
                                                         PureScript_Data_Number::Data_Number_FFI::acos(x)))
    }
    pub fn Data_Number_asin() -> &dyn Any {
        static Data_Number_asin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_asin.get_or_init(||
                                         &Func1::new(move |x|
                                                         PureScript_Data_Number::Data_Number_FFI::asin(x)))
    }
    pub fn Data_Number_atan() -> &dyn Any {
        static Data_Number_atan: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_atan.get_or_init(||
                                         &Func1::new(move |x|
                                                         PureScript_Data_Number::Data_Number_FFI::atan(x)))
    }
    pub fn Data_Number_atan2() -> &dyn Any {
        static Data_Number_atan2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_atan2.get_or_init(||
                                          &Func1::new(move |y|
                                                          Func1::new({
                                                                         let y
                                                                             =
                                                                             y.clone();
                                                                         move
                                                                             |x|
                                                                             PureScript_Data_Number::Data_Number_FFI::atan2(&y,
                                                                                                                            x)
                                                                     })))
    }
    pub fn Data_Number_ceil() -> &dyn Any {
        static Data_Number_ceil: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_ceil.get_or_init(||
                                         &Func1::new(move |x|
                                                         PureScript_Data_Number::Data_Number_FFI::ceil(x)))
    }
    pub fn Data_Number_cos() -> &dyn Any {
        static Data_Number_cos: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_cos.get_or_init(||
                                        &Func1::new(move |x|
                                                        PureScript_Data_Number::Data_Number_FFI::cos(x)))
    }
    pub fn Data_Number_exp() -> &dyn Any {
        static Data_Number_exp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_exp.get_or_init(||
                                        &Func1::new(move |x|
                                                        PureScript_Data_Number::Data_Number_FFI::exp(x)))
    }
    pub fn Data_Number_floor() -> &dyn Any {
        static Data_Number_floor: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_floor.get_or_init(||
                                          &Func1::new(move |x|
                                                          PureScript_Data_Number::Data_Number_FFI::floor(x)))
    }
    pub fn Data_Number_fromStringImpl() -> &dyn Any {
        static Data_Number_fromStringImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_fromStringImpl.get_or_init(||
                                                   &Func1::new(move |strVal|
                                                                   Func1::new({
                                                                                  let strVal
                                                                                      =
                                                                                      strVal.clone();
                                                                                  move
                                                                                      |isFiniteFn|
                                                                                      Func1::new({
                                                                                                     let isFiniteFn
                                                                                                         =
                                                                                                         isFiniteFn.clone();
                                                                                                     move
                                                                                                         |just|
                                                                                                         Func1::new({
                                                                                                                        let just
                                                                                                                            =
                                                                                                                            just.clone();
                                                                                                                        move
                                                                                                                            |nothing|
                                                                                                                            PureScript_Data_Number::Data_Number_FFI::fromStringImpl(&strVal,
                                                                                                                                                                                    &isFiniteFn,
                                                                                                                                                                                    &just,
                                                                                                                                                                                    nothing)
                                                                                                                    })
                                                                                                 })
                                                                              })))
    }
    pub fn Data_Number_infinity() -> &dyn Any {
        static Data_Number_infinity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_infinity.get_or_init(||
                                             &PureScript_Data_Number::Data_Number_FFI::infinity())
    }
    pub fn Data_Number_isFinite() -> &dyn Any {
        static Data_Number_isFinite: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_isFinite.get_or_init(||
                                             &Func1::new(move |x|
                                                             PureScript_Data_Number::Data_Number_FFI::isFinite(x)))
    }
    pub fn Data_Number_isNaN() -> &dyn Any {
        static Data_Number_isNaN: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_isNaN.get_or_init(||
                                          &Func1::new(move |x|
                                                          PureScript_Data_Number::Data_Number_FFI::isNaN(x)))
    }
    pub fn Data_Number_log() -> &dyn Any {
        static Data_Number_log: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_log.get_or_init(||
                                        &Func1::new(move |x|
                                                        PureScript_Data_Number::Data_Number_FFI::log(x)))
    }
    pub fn Data_Number_max() -> &dyn Any {
        static Data_Number_max: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_max.get_or_init(||
                                        &Func1::new(move |n|
                                                        Func1::new({
                                                                       let n =
                                                                           n.clone();
                                                                       move
                                                                           |n_1|
                                                                           PureScript_Data_Number::Data_Number_FFI::max(&n,
                                                                                                                        n_1)
                                                                   })))
    }
    pub fn Data_Number_min() -> &dyn Any {
        static Data_Number_min: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_min.get_or_init(||
                                        &Func1::new(move |n|
                                                        Func1::new({
                                                                       let n =
                                                                           n.clone();
                                                                       move
                                                                           |n_1|
                                                                           PureScript_Data_Number::Data_Number_FFI::min(&n,
                                                                                                                        n_1)
                                                                   })))
    }
    pub fn Data_Number_nan() -> &dyn Any {
        static Data_Number_nan: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_nan.get_or_init(||
                                        &PureScript_Data_Number::Data_Number_FFI::nan())
    }
    pub fn Data_Number_pow() -> &dyn Any {
        static Data_Number_pow: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_pow.get_or_init(||
                                        &Func1::new(move |n|
                                                        Func1::new({
                                                                       let n =
                                                                           n.clone();
                                                                       move
                                                                           |p|
                                                                           PureScript_Data_Number::Data_Number_FFI::pow(&n,
                                                                                                                        p)
                                                                   })))
    }
    pub fn Data_Number_remainder() -> &dyn Any {
        static Data_Number_remainder: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_remainder.get_or_init(||
                                              &Func1::new(move |n|
                                                              Func1::new({
                                                                             let n
                                                                                 =
                                                                                 n.clone();
                                                                             move
                                                                                 |m|
                                                                                 PureScript_Data_Number::Data_Number_FFI::remainder(&n,
                                                                                                                                    m)
                                                                         })))
    }
    pub fn Data_Number_round() -> &dyn Any {
        static Data_Number_round: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_round.get_or_init(||
                                          &Func1::new(move |x|
                                                          PureScript_Data_Number::Data_Number_FFI::round(x)))
    }
    pub fn Data_Number_sign() -> &dyn Any {
        static Data_Number_sign: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_sign.get_or_init(||
                                         &Func1::new(move |x|
                                                         PureScript_Data_Number::Data_Number_FFI::sign(x)))
    }
    pub fn Data_Number_sin() -> &dyn Any {
        static Data_Number_sin: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_sin.get_or_init(||
                                        &Func1::new(move |x|
                                                        PureScript_Data_Number::Data_Number_FFI::sin(x)))
    }
    pub fn Data_Number_sqrt() -> &dyn Any {
        static Data_Number_sqrt: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_sqrt.get_or_init(||
                                         &Func1::new(move |x|
                                                         PureScript_Data_Number::Data_Number_FFI::sqrt(x)))
    }
    pub fn Data_Number_tan() -> &dyn Any {
        static Data_Number_tan: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_tan.get_or_init(||
                                        &Func1::new(move |x|
                                                        PureScript_Data_Number::Data_Number_FFI::tan(x)))
    }
    pub fn Data_Number_trunc() -> &dyn Any {
        static Data_Number_trunc: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_trunc.get_or_init(||
                                          &Func1::new(move |x|
                                                          PureScript_Data_Number::Data_Number_FFI::trunc(x)))
    }
    pub fn Data_Number_tau() -> &dyn Any {
        static Data_Number_tau: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_tau.get_or_init(|| &6.283185307179586_f64)
    }
    pub fn Data_Number_sqrt2() -> &dyn Any {
        static Data_Number_sqrt2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_sqrt2.get_or_init(|| &1.4142135623730951_f64)
    }
    pub fn Data_Number_sqrt1_2() -> &dyn Any {
        static Data_Number_sqrt1_2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_sqrt1_2.get_or_init(|| &0.7071067811865476_f64)
    }
    pub fn Data_Number_pi() -> &dyn Any {
        static Data_Number_pi: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Number_pi.get_or_init(|| &3.141592653589793_f64)
    }
    pub fn Data_Number_log2e() -> &dyn Any {
        static Data_Number_log2e: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_log2e.get_or_init(|| &1.4426950408889634_f64)
    }
    pub fn Data_Number_log10e() -> &dyn Any {
        static Data_Number_log10e: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_log10e.get_or_init(|| &0.4342944819032518_f64)
    }
    pub fn Data_Number_ln2() -> &dyn Any {
        static Data_Number_ln2: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_ln2.get_or_init(|| &0.6931471805599453_f64)
    }
    pub fn Data_Number_ln10() -> &dyn Any {
        static Data_Number_ln10: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_ln10.get_or_init(|| &2.302585092994046_f64)
    }
    pub fn Data_Number_fromString() -> &dyn Any {
        static Data_Number_fromString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_fromString.get_or_init(||
                                               &Func1::new(move |str|
                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Function_Uncurried::Data_Function_Uncurried_runFn4(),
                                                                                                                                                                                                                                            &&&PureScript_Data_Number::Data_Number_fromStringImpl()),
                                                                                                                                                                                                         str),
                                                                                                                                                                      &&&PureScript_Data_Number::Data_Number_isFinite()),
                                                                                                                                   &&&Func1::new(move
                                                                                                                                                     |usd__arg1|
                                                                                                                                                     &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                                                &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor))))
    }
    pub fn Data_Number_e() -> &dyn Any {
        static Data_Number_e: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Number_e.get_or_init(|| &2.718281828459045_f64)
    }
}

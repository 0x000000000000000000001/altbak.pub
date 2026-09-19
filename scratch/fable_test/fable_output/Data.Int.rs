pub mod PureScript_Data_Int {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use crate::module_86d6df2::PureScript_Control_Category;
    use crate::module_655da3ed::PureScript_Control_Semigroupoid;
    use crate::module_a4848631::PureScript_Data_Boolean;
    use crate::module_d89c2f46::PureScript_Data_Bounded;
    use crate::module_cb047b05::PureScript_Data_CommutativeRing;
    use crate::module_5ba8cfce::PureScript_Data_DivisionRing;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_e83e2264::PureScript_Data_Int_Bits;
    use crate::module_f879ba47::PureScript_Data_Maybe;
    use crate::module_f879ba47::PureScript_Data_Maybe::Data_Maybe_Maybe;
    use crate::module_91782ab6::PureScript_Data_Number;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_5f769efb::PureScript_Data_Ordering::Data_Ordering_Ordering;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_baebcb76::PureScript_Data_Show;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::Microsoft::FSharp::Core::MatchFailureException;
    pub mod Data_Int_FFI {
        use super::*;
        use fable_library_rust::Convert_::toInt32_radix;
        use fable_library_rust::Convert_::toStringRadix;
        use fable_library_rust::Exception_::try_catch;
        use fable_library_rust::System::Exception;
        pub fn fromNumberImpl(just: &dyn Any, nothing: &dyn Any,
                              nVal: &dyn Any) -> &dyn Any {
            let n: f64 = nVal.clone();
            if n == n { just(&(n as i32)) } else { nothing.clone() }
        }
        pub fn toNumber(n: &dyn Any) -> &dyn Any { n.clone() }
        pub fn fromStringAsImpl(just: &dyn Any, nothing: &dyn Any,
                                radixVal: &dyn Any)
         -> Func1<&dyn Any, &dyn Any> {
            Func1::new({
                           let just = just.clone();
                           let nothing = nothing.clone();
                           let radixVal = radixVal.clone();
                           move |sVal|
                               try_catch(||
                                             just(&toInt32_radix(sVal.clone(),
                                                                 radixVal)),
                                         |matchValue: LrcPtr<Exception>|
                                             nothing)
                       })
        }
        pub fn toStringAs(radixVal: &dyn Any, iVal: &dyn Any) -> &dyn Any {
            &toStringRadix(iVal.clone(), radixVal.clone())
        }
        pub fn quot(xVal: &dyn Any, yVal: &dyn Any) -> &dyn Any {
            &(xVal.clone() / yVal.clone())
        }
        pub fn rem(xVal: &dyn Any, yVal: &dyn Any) -> &dyn Any {
            &(xVal.clone() % yVal.clone())
        }
        pub fn pow(xVal: &dyn Any, yVal: &dyn Any) -> &dyn Any {
            &(xVal.clone().powf(yVal.clone()) as i32)
        }
    }
    pub fn Data_Int_fromNumberImpl() -> &dyn Any {
        static Data_Int_fromNumberImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_fromNumberImpl.get_or_init(||
                                                &Func1::new(move |just|
                                                                Func1::new({
                                                                               let just
                                                                                   =
                                                                                   just.clone();
                                                                               move
                                                                                   |nothing|
                                                                                   Func1::new({
                                                                                                  let nothing
                                                                                                      =
                                                                                                      nothing.clone();
                                                                                                  move
                                                                                                      |nVal|
                                                                                                      PureScript_Data_Int::Data_Int_FFI::fromNumberImpl(&just,
                                                                                                                                                        &nothing,
                                                                                                                                                        nVal)
                                                                                              })
                                                                           })))
    }
    pub fn Data_Int_fromStringAsImpl() -> &dyn Any {
        static Data_Int_fromStringAsImpl: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_fromStringAsImpl.get_or_init(||
                                                  &Func1::new(move |just|
                                                                  Func1::new({
                                                                                 let just
                                                                                     =
                                                                                     just.clone();
                                                                                 move
                                                                                     |nothing|
                                                                                     Func1::new({
                                                                                                    let nothing
                                                                                                        =
                                                                                                        nothing.clone();
                                                                                                    move
                                                                                                        |radixVal|
                                                                                                        PureScript_Data_Int::Data_Int_FFI::fromStringAsImpl(&just,
                                                                                                                                                            &nothing,
                                                                                                                                                            radixVal)
                                                                                                })
                                                                             })))
    }
    pub fn Data_Int_pow() -> &dyn Any {
        static Data_Int_pow: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_pow.get_or_init(||
                                     &Func1::new(move |xVal|
                                                     Func1::new({
                                                                    let xVal =
                                                                        xVal.clone();
                                                                    move
                                                                        |yVal|
                                                                        PureScript_Data_Int::Data_Int_FFI::pow(&xVal,
                                                                                                               yVal)
                                                                })))
    }
    pub fn Data_Int_quot() -> &dyn Any {
        static Data_Int_quot: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_quot.get_or_init(||
                                      &Func1::new(move |xVal|
                                                      Func1::new({
                                                                     let xVal
                                                                         =
                                                                         xVal.clone();
                                                                     move
                                                                         |yVal|
                                                                         PureScript_Data_Int::Data_Int_FFI::quot(&xVal,
                                                                                                                 yVal)
                                                                 })))
    }
    pub fn Data_Int_rem() -> &dyn Any {
        static Data_Int_rem: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_rem.get_or_init(||
                                     &Func1::new(move |xVal|
                                                     Func1::new({
                                                                    let xVal =
                                                                        xVal.clone();
                                                                    move
                                                                        |yVal|
                                                                        PureScript_Data_Int::Data_Int_FFI::rem(&xVal,
                                                                                                               yVal)
                                                                })))
    }
    pub fn Data_Int_toNumber() -> &dyn Any {
        static Data_Int_toNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_toNumber.get_or_init(||
                                          &Func1::new(move |n|
                                                          PureScript_Data_Int::Data_Int_FFI::toNumber(n)))
    }
    pub fn Data_Int_toStringAs() -> &dyn Any {
        static Data_Int_toStringAs: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_toStringAs.get_or_init(||
                                            &Func1::new(move |radixVal|
                                                            Func1::new({
                                                                           let radixVal
                                                                               =
                                                                               radixVal.clone();
                                                                           move
                                                                               |iVal|
                                                                               PureScript_Data_Int::Data_Int_FFI::toStringAs(&radixVal,
                                                                                                                             iVal)
                                                                       })))
    }
    #[derive(Clone, Debug, PartialEq, PartialOrd, Hash, Eq, Ord,)]
    pub enum Data_Int_Parity { Data_Int_Evenusd_Ctor, Data_Int_Oddusd_Ctor, }
    impl core::fmt::Display for PureScript_Data_Int::Data_Int_Parity {
        fn fmt(&self, f: &mut core::fmt::Formatter) -> core::fmt::Result {
            write!(f, "{}", core::any::type_name::<Self>())
        }
    }
    pub fn Data_Int_top() -> &dyn Any {
        static Data_Int_top: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_top.get_or_init(||
                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_top(),
                                                                      &&&PureScript_Data_Bounded::Data_Bounded_boundedInt()))
    }
    pub fn Data_Int_bottom() -> &dyn Any {
        static Data_Int_bottom: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_bottom.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_bottom(),
                                                                         &&&PureScript_Data_Bounded::Data_Bounded_boundedInt()))
    }
    pub fn Data_Int_Radix() -> &dyn Any {
        static Data_Int_Radix: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_Radix.get_or_init(|| &Func1::new(move |x| x.clone()))
    }
    pub fn Data_Int_Even() -> &dyn Any {
        static Data_Int_Even: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_Even.get_or_init(||
                                      &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor))
    }
    pub fn Data_Int_Odd() -> &dyn Any {
        static Data_Int_Odd: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_Odd.get_or_init(||
                                     &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor))
    }
    pub fn Data_Int_showParity() -> &dyn Any {
        static Data_Int_showParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_showParity.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Show::Data_Show_Showusd_Dict(),
                                                                             &&&add(string("show"),
                                                                                    &&Func1::new(move
                                                                                                     |v|
                                                                                                     {
                                                                                                         let matchValue:
                                                                                                                 LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                             Sharpurs_Prelude::unbox(v);
                                                                                                         match matchValue.as_ref()
                                                                                                             {
                                                                                                             PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                             =>
                                                                                                             &string("Odd"),
                                                                                                             _
                                                                                                             =>
                                                                                                             &string("Even"),
                                                                                                         }
                                                                                                     }),
                                                                                    empty::<string,
                                                                                            &dyn Any>())))
    }
    pub fn Data_Int_radix() -> &dyn Any {
        static Data_Int_radix: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_radix.get_or_init(||
                                       &Func1::new(move |n|
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
                                                                                                                                                                                                  &&&2_i32)),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                     &&&PureScript_Data_Ord::Data_Ord_ordInt()),
                                                                                                                                                                                                  &&&n1),
                                                                                                                                                               &&&36_i32)))
                                                              } {
                                                               &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_Radix(),
                                                                                                                                                       &&&matchValue)))
                                                           } else {
                                                               if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                  {
                                                                   &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)
                                                               } else {
                                                                   panic!("{}",
                                                                          LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Int.fs"),
                                  Data1: 82_i32,
                                  Data2: 51_i32,}).get_Message(),)
                                                               }
                                                           }
                                                       }))
    }
    pub fn Data_Int_odd() -> &dyn Any {
        static Data_Int_odd: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_odd.get_or_init(||
                                     &Func1::new(move |x|
                                                     Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_notEq(),
                                                                                                                                                            &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int_Bits::Data_Int_Bits_and(),
                                                                                                                                                                                               x),
                                                                                                                                                            &&&1_i32)),
                                                                                      &&&0_i32)))
    }
    pub fn Data_Int_octal() -> &dyn Any {
        static Data_Int_octal: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_octal.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_Radix(),
                                                                        &&&8_i32))
    }
    pub fn Data_Int_hexadecimal() -> &dyn Any {
        static Data_Int_hexadecimal: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_hexadecimal.get_or_init(||
                                             Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_Radix(),
                                                                              &&&16_i32))
    }
    pub fn Data_Int_fromStringAs() -> &dyn Any {
        static Data_Int_fromStringAs: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_fromStringAs.get_or_init(||
                                              Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_fromStringAsImpl(),
                                                                                                                  &&&Func1::new(move
                                                                                                                                    |usd__arg1|
                                                                                                                                    &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                               &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Int_fromString() -> &dyn Any {
        static Data_Int_fromString: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_fromString.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_fromStringAs(),
                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_Radix(),
                                                                                                                &&&10_i32)))
    }
    pub fn Data_Int_fromNumber() -> &dyn Any {
        static Data_Int_fromNumber: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_fromNumber.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_fromNumberImpl(),
                                                                                                                &&&Func1::new(move
                                                                                                                                  |usd__arg1|
                                                                                                                                  &LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Justusd_Ctor(usd__arg1.clone())))),
                                                                             &&&LrcPtr::new(Data_Maybe_Maybe::Data_Maybe_Nothingusd_Ctor)))
    }
    pub fn Data_Int_unsafeClamp() -> &dyn Any {
        static Data_Int_unsafeClamp: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_unsafeClamp.get_or_init(||
                                             &Func1::new(move |x|
                                                             {
                                                                 let matchValue =
                                                                     Sharpurs_Prelude::unbox(x);
                                                                 if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                 &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number::Data_Number_isFinite(),
                                                                                                                                                                 &&&matchValue)))
                                                                    {
                                                                     &0_i32
                                                                 } else {
                                                                     if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_greaterThanOrEq(),
                                                                                                                                                                                                        &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                     &&&matchValue),
                                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_toNumber(),
                                                                                                                                                                     &&&PureScript_Data_Int::Data_Int_top())))
                                                                        {
                                                                         &PureScript_Data_Int::Data_Int_top()
                                                                     } else {
                                                                         if Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                            &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                         &&&matchValue),
                                                                                                                                      &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_toNumber(),
                                                                                                                                                                         &&&PureScript_Data_Int::Data_Int_bottom())))
                                                                            {
                                                                             &PureScript_Data_Int::Data_Int_bottom()
                                                                         } else {
                                                                             if Sharpurs_Prelude::unbox(&&PureScript_Data_Boolean::Data_Boolean_otherwise())
                                                                                {
                                                                                 Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Maybe::Data_Maybe_fromMaybe(),
                                                                                                                                                     &&&0_i32),
                                                                                                                  &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_fromNumber(),
                                                                                                                                                     &&&matchValue))
                                                                             } else {
                                                                                 panic!("{}",
                                                                                        LrcPtr::new(MatchFailureException{Data0:
                                      string("/Users/0x1/Documents/htdocs/altbak.pub/scratch/fable_test/Data.Int.fs"),
                                  Data1: 96_i32,
                                  Data2: 57_i32,}).get_Message(),)
                                                                             }
                                                                         }
                                                                     }
                                                                 }
                                                             }))
    }
    pub fn Data_Int_round() -> &dyn Any {
        static Data_Int_round: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_round.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                           &&&PureScript_Data_Int::Data_Int_unsafeClamp()),
                                                                        &&&PureScript_Data_Number::Data_Number_round()))
    }
    pub fn Data_Int_trunc() -> &dyn Any {
        static Data_Int_trunc: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_trunc.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                           &&&PureScript_Data_Int::Data_Int_unsafeClamp()),
                                                                        &&&PureScript_Data_Number::Data_Number_trunc()))
    }
    pub fn Data_Int_floor() -> &dyn Any {
        static Data_Int_floor: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_floor.get_or_init(||
                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                              &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                           &&&PureScript_Data_Int::Data_Int_unsafeClamp()),
                                                                        &&&PureScript_Data_Number::Data_Number_floor()))
    }
    pub fn Data_Int_even() -> &dyn Any {
        static Data_Int_even: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_even.get_or_init(||
                                      &Func1::new(move |x|
                                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                             &&&PureScript_Data_Eq::Data_Eq_eqInt()),
                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int_Bits::Data_Int_Bits_and(),
                                                                                                                                                                                                x),
                                                                                                                                                             &&&1_i32)),
                                                                                       &&&0_i32)))
    }
    pub fn Data_Int_parity() -> &dyn Any {
        static Data_Int_parity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_parity.get_or_init(||
                                        &Func1::new(move |n|
                                                        {
                                                            let matchValue =
                                                                Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_even(),
                                                                                                                          n));
                                                            match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                             &matchValue)
                                                                {
                                                                0_i32 =>
                                                                &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor),
                                                                _ =>
                                                                &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor),
                                                            }
                                                        }))
    }
    pub fn Data_Int_eqParity() -> &dyn Any {
        static Data_Int_eqParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_eqParity.get_or_init(||
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
                                                                                                                                   LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                                               Sharpurs_Prelude::unbox(&&x);
                                                                                                                           let matchValue_1:
                                                                                                                                   LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                                               Sharpurs_Prelude::unbox(y);
                                                                                                                           if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                                                  =
                                                                                                                                  matchValue.as_ref()
                                                                                                                              {
                                                                                                                               if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                                                      =
                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                  {
                                                                                                                                   &true
                                                                                                                               } else {
                                                                                                                                   &false
                                                                                                                               }
                                                                                                                           } else {
                                                                                                                               if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor
                                                                                                                                      =
                                                                                                                                      matchValue_1.as_ref()
                                                                                                                                  {
                                                                                                                                   &true
                                                                                                                               } else {
                                                                                                                                   &false
                                                                                                                               }
                                                                                                                           }
                                                                                                                       }
                                                                                                               })),
                                                                                  empty::<string,
                                                                                          &dyn Any>())))
    }
    pub fn Data_Int_ordParity() -> &dyn Any {
        static Data_Int_ordParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_ordParity.get_or_init(||
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
                                                                                                                                    LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                                                Sharpurs_Prelude::unbox(&&x);
                                                                                                                            let matchValue_1:
                                                                                                                                    LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                                                Sharpurs_Prelude::unbox(y);
                                                                                                                            if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                                                   =
                                                                                                                                   matchValue.as_ref()
                                                                                                                               {
                                                                                                                                if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                                                       =
                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                } else {
                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_GTusd_Ctor)
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor
                                                                                                                                       =
                                                                                                                                       matchValue_1.as_ref()
                                                                                                                                   {
                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_EQusd_Ctor)
                                                                                                                                } else {
                                                                                                                                    &LrcPtr::new(Data_Ordering_Ordering::Data_Ordering_LTusd_Ctor)
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                })),
                                                                                   add(string("Eq0"),
                                                                                       &&Func1::new(move
                                                                                                        |usd__unused|
                                                                                                        &PureScript_Data_Int::Data_Int_eqParity()),
                                                                                       empty::<string,
                                                                                               &dyn Any>()))))
    }
    pub fn Data_Int_semiringParity() -> &dyn Any {
        static Data_Int_semiringParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_semiringParity.get_or_init(||
                                                Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_Semiringusd_Dict(),
                                                                                 &&&add(string("zero"),
                                                                                        &&LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor),
                                                                                        add(string("add"),
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
                                                                                                                                         Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                         &&&PureScript_Data_Int::Data_Int_eqParity()),
                                                                                                                                                                                                                                      &&&x),
                                                                                                                                                                                                   y));
                                                                                                                                     match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                      &matchValue)
                                                                                                                                         {
                                                                                                                                         0_i32
                                                                                                                                         =>
                                                                                                                                         &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor),
                                                                                                                                         _
                                                                                                                                         =>
                                                                                                                                         &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor),
                                                                                                                                     }
                                                                                                                                 }
                                                                                                                         })),
                                                                                            add(string("one"),
                                                                                                &&LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor),
                                                                                                add(string("mul"),
                                                                                                    &&Func1::new(move
                                                                                                                     |v|
                                                                                                                     &Func1::new({
                                                                                                                                     let v
                                                                                                                                         =
                                                                                                                                         v.clone();
                                                                                                                                     move
                                                                                                                                         |v1|
                                                                                                                                         {
                                                                                                                                             let matchValue_1:
                                                                                                                                                     LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                                                                 Sharpurs_Prelude::unbox(&&v);
                                                                                                                                             let matchValue_2:
                                                                                                                                                     LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                                                                 Sharpurs_Prelude::unbox(v1);
                                                                                                                                             if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                                                                    =
                                                                                                                                                    matchValue_1.as_ref()
                                                                                                                                                {
                                                                                                                                                 if let PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                                                                        =
                                                                                                                                                        matchValue_2.as_ref()
                                                                                                                                                    {
                                                                                                                                                     &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor)
                                                                                                                                                 } else {
                                                                                                                                                     &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor)
                                                                                                                                                 }
                                                                                                                                             } else {
                                                                                                                                                 &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor)
                                                                                                                                             }
                                                                                                                                         }
                                                                                                                                 })),
                                                                                                    empty::<string,
                                                                                                            &dyn Any>()))))))
    }
    pub fn Data_Int_ringParity() -> &dyn Any {
        static Data_Int_ringParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_ringParity.get_or_init(||
                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_Ringusd_Dict(),
                                                                             &&&add(string("sub"),
                                                                                    &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                      &&&PureScript_Data_Int::Data_Int_semiringParity()),
                                                                                    add(string("Semiring0"),
                                                                                        &&Func1::new(move
                                                                                                         |usd__unused|
                                                                                                         &PureScript_Data_Int::Data_Int_semiringParity()),
                                                                                        empty::<string,
                                                                                                &dyn Any>()))))
    }
    pub fn Data_Int_divisionRingParity() -> &dyn Any {
        static Data_Int_divisionRingParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_divisionRingParity.get_or_init(||
                                                    Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_DivisionRing::Data_DivisionRing_DivisionRingusd_Dict(),
                                                                                     &&&add(string("recip"),
                                                                                            &Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Category::Control_Category_identity(),
                                                                                                                              &&&PureScript_Control_Category::Control_Category_categoryFn()),
                                                                                            add(string("Ring0"),
                                                                                                &&Func1::new(move
                                                                                                                 |usd__unused|
                                                                                                                 &PureScript_Data_Int::Data_Int_ringParity()),
                                                                                                empty::<string,
                                                                                                        &dyn Any>()))))
    }
    pub fn Data_Int_decimal() -> &dyn Any {
        static Data_Int_decimal: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_decimal.get_or_init(||
                                         Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_Radix(),
                                                                          &&&10_i32))
    }
    pub fn Data_Int_commutativeRingParity() -> &dyn Any {
        static Data_Int_commutativeRingParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_commutativeRingParity.get_or_init(||
                                                       Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_CommutativeRing::Data_CommutativeRing_CommutativeRingusd_Dict(),
                                                                                        &&&add(string("Ring0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Int::Data_Int_ringParity()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))
    }
    pub fn Data_Int_euclideanRingParity() -> &dyn Any {
        static Data_Int_euclideanRingParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_euclideanRingParity.get_or_init(||
                                                     Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_EuclideanRingusd_Dict(),
                                                                                      &&&add(string("degree"),
                                                                                             &&Func1::new(move
                                                                                                              |v|
                                                                                                              {
                                                                                                                  let matchValue:
                                                                                                                          LrcPtr<PureScript_Data_Int::Data_Int_Parity> =
                                                                                                                      Sharpurs_Prelude::unbox(v);
                                                                                                                  match matchValue.as_ref()
                                                                                                                      {
                                                                                                                      PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor
                                                                                                                      =>
                                                                                                                      &1_i32,
                                                                                                                      _
                                                                                                                      =>
                                                                                                                      &0_i32,
                                                                                                                  }
                                                                                                              }),
                                                                                             add(string("div"),
                                                                                                 &&Func1::new(move
                                                                                                                  |x|
                                                                                                                  &Func1::new({
                                                                                                                                  let x
                                                                                                                                      =
                                                                                                                                      x.clone();
                                                                                                                                  move
                                                                                                                                      |v_1|
                                                                                                                                      &x
                                                                                                                              })),
                                                                                                 add(string("mod"),
                                                                                                     &&Func1::new(move
                                                                                                                      |v_2|
                                                                                                                      &Func1::new(move
                                                                                                                                      |v1|
                                                                                                                                      &LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor))),
                                                                                                     add(string("CommutativeRing0"),
                                                                                                         &&Func1::new(move
                                                                                                                          |usd__unused|
                                                                                                                          &PureScript_Data_Int::Data_Int_commutativeRingParity()),
                                                                                                         empty::<string,
                                                                                                                 &dyn Any>()))))))
    }
    pub fn Data_Int_ceil() -> &dyn Any {
        static Data_Int_ceil: MutCell<Option<&dyn Any>> = MutCell::new(None);
        Data_Int_ceil.get_or_init(||
                                      Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Control_Semigroupoid::Control_Semigroupoid_compose(),
                                                                                                                                             &&&PureScript_Control_Semigroupoid::Control_Semigroupoid_semigroupoidFn()),
                                                                                                          &&&PureScript_Data_Int::Data_Int_unsafeClamp()),
                                                                       &&&PureScript_Data_Number::Data_Number_ceil()))
    }
    pub fn Data_Int_boundedParity() -> &dyn Any {
        static Data_Int_boundedParity: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_boundedParity.get_or_init(||
                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Bounded::Data_Bounded_Boundedusd_Dict(),
                                                                                &&&add(string("bottom"),
                                                                                       &&LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Evenusd_Ctor),
                                                                                       add(string("top"),
                                                                                           &&LrcPtr::new(PureScript_Data_Int::Data_Int_Parity::Data_Int_Oddusd_Ctor),
                                                                                           add(string("Ord0"),
                                                                                               &&Func1::new(move
                                                                                                                |usd__unused|
                                                                                                                &PureScript_Data_Int::Data_Int_ordParity()),
                                                                                               empty::<string,
                                                                                                       &dyn Any>())))))
    }
    pub fn Data_Int_binary() -> &dyn Any {
        static Data_Int_binary: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_binary.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_Radix(),
                                                                         &&&2_i32))
    }
    pub fn Data_Int_base36() -> &dyn Any {
        static Data_Int_base36: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Int_base36.get_or_init(||
                                        Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Int::Data_Int_Radix(),
                                                                         &&&36_i32))
    }
}

pub mod PureScript_Data_EuclideanRing {
    use super::*;
    use fable_library_rust::Map_::add;
    use fable_library_rust::Map_::empty;
    use fable_library_rust::Map_::find;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::LrcPtr;
    use fable_library_rust::Native_::MutCell;
    use fable_library_rust::String_::string;
    use fable_library_rust::Util_::Lazy;
    use crate::module_cb047b05::PureScript_Data_CommutativeRing;
    use crate::module_8617f961::PureScript_Data_Eq;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    use fable_library_rust::System::Lazy_1;
    pub mod Data_EuclideanRing_FFI {
        use super::*;
        pub fn absInt(n: i32) -> i32 {
            if n == i32::MIN { i32::MAX } else { n.abs() }
        }
        pub fn intDiv<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> i32 {
            let x: i32 = Sharpurs_Prelude::unbox(&&a);
            let y: i32 = Sharpurs_Prelude::unbox(&&b);
            if y == 0_i32 {
                0_i32
            } else {
                let yy: i32 =
                    PureScript_Data_EuclideanRing::Data_EuclideanRing_FFI::absInt(y);
                if if y == -1_i32 { x == i32::MIN } else { false } {
                    i32::MIN
                } else { (x - (x % yy + yy) % yy) / y }
            }
        }
        pub fn intDegree<a: Clone + 'static>(a: a) -> i32 {
            PureScript_Data_EuclideanRing::Data_EuclideanRing_FFI::absInt(Sharpurs_Prelude::unbox(&&a))
        }
        pub fn numDiv<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> f64 {
            Sharpurs_Prelude::unbox(&&a) / Sharpurs_Prelude::unbox(&&b)
        }
        pub fn intMod<a: Clone + 'static, b: Clone + 'static>(a: a, b: b)
         -> i32 {
            let x: i32 = Sharpurs_Prelude::unbox(&&a);
            let y: i32 = Sharpurs_Prelude::unbox(&&b);
            if y == 0_i32 {
                0_i32
            } else {
                let yy: i32 =
                    PureScript_Data_EuclideanRing::Data_EuclideanRing_FFI::absInt(y);
                (x % yy + yy) % yy
            }
        }
    }
    pub fn Data_EuclideanRing_intDegree() -> &dyn Any {
        static Data_EuclideanRing_intDegree: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_intDegree.get_or_init(||
                                                     &Func1::new(move |arg0|
                                                                     &PureScript_Data_EuclideanRing::Data_EuclideanRing_FFI::intDegree(Sharpurs_Prelude::unbox(arg0))))
    }
    pub fn Data_EuclideanRing_intDiv() -> &dyn Any {
        static Data_EuclideanRing_intDiv: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_intDiv.get_or_init(||
                                                  &Func1::new(move |arg0|
                                                                  &Func1::new({
                                                                                  let arg0
                                                                                      =
                                                                                      arg0.clone();
                                                                                  move
                                                                                      |arg1|
                                                                                      &PureScript_Data_EuclideanRing::Data_EuclideanRing_FFI::intDiv(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                     Sharpurs_Prelude::unbox(arg1))
                                                                              })))
    }
    pub fn Data_EuclideanRing_intMod() -> &dyn Any {
        static Data_EuclideanRing_intMod: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_intMod.get_or_init(||
                                                  &Func1::new(move |arg0|
                                                                  &Func1::new({
                                                                                  let arg0
                                                                                      =
                                                                                      arg0.clone();
                                                                                  move
                                                                                      |arg1|
                                                                                      &PureScript_Data_EuclideanRing::Data_EuclideanRing_FFI::intMod(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                     Sharpurs_Prelude::unbox(arg1))
                                                                              })))
    }
    pub fn Data_EuclideanRing_numDiv() -> &dyn Any {
        static Data_EuclideanRing_numDiv: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_numDiv.get_or_init(||
                                                  &Func1::new(move |arg0|
                                                                  &Func1::new({
                                                                                  let arg0
                                                                                      =
                                                                                      arg0.clone();
                                                                                  move
                                                                                      |arg1|
                                                                                      &PureScript_Data_EuclideanRing::Data_EuclideanRing_FFI::numDiv(Sharpurs_Prelude::unbox(&arg0),
                                                                                                                                                     Sharpurs_Prelude::unbox(arg1))
                                                                              })))
    }
    pub fn Data_EuclideanRing_EuclideanRingusd_Dict() -> &dyn Any {
        static Data_EuclideanRing_EuclideanRingusd_Dict:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_EuclideanRingusd_Dict.get_or_init(||
                                                                 &Func1::new(move
                                                                                 |x|
                                                                                 x.clone()))
    }
    pub fn Data_EuclideanRing_mod() -> &dyn Any {
        static Data_EuclideanRing_mod: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_mod.get_or_init(||
                                               &Func1::new(move |dict|
                                                               find(string("mod"),
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_EuclideanRing_gcd_004044() -> &dyn Any {
        &Func1::new(move |dictEq|
                        Func1::new({
                                       let dictEq = dictEq.clone();
                                       move |dictEuclideanRing|
                                           PureScript_Data_EuclideanRing::Data_EuclideanRing_gcd_tco(&dictEq,
                                                                                                     dictEuclideanRing)
                                   }))
    }
    pub fn Data_EuclideanRing_gcd_004044_002d1() -> LrcPtr<Lazy_1<&dyn Any>> {
        static Data_EuclideanRing_gcd_004044_002d1:
         MutCell<Option<LrcPtr<Lazy_1<&dyn Any>>>> =
            MutCell::new(None);
        Data_EuclideanRing_gcd_004044_002d1.get_or_init(||
                                                            Lazy(Data_EuclideanRing_gcd_004044.clone()))
    }
    pub fn Data_EuclideanRing_gcd_tco(dictEq: &dyn Any,
                                      dictEuclideanRing: &dyn Any)
     -> &dyn Any {
        let zero =
            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                             &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("CommutativeRing0"),
                                                                                                                                                                                                                         Sharpurs_Prelude::unbox(dictEuclideanRing)),
                                                                                                                                                                                                                  &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                                                 &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                &&&Sharpurs_Prelude::Prim_undefined()));
        &Func1::new({
                        let Data_EuclideanRing_gcd_004044_002d1 =
                            Data_EuclideanRing_gcd_004044_002d1.clone();
                        let dictEq = dictEq.clone();
                        let dictEuclideanRing = dictEuclideanRing.clone();
                        let zero = zero.clone();
                        move |a|
                            &Func1::new({
                                            let Data_EuclideanRing_gcd_004044_002d1
                                                =
                                                Data_EuclideanRing_gcd_004044_002d1.clone();
                                            let a = a.clone();
                                            move |b|
                                                {
                                                    let matchValue =
                                                        Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                        &&&dictEq),
                                                                                                                                                     b),
                                                                                                                  &&&zero));
                                                    match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                     &matchValue)
                                                        {
                                                        0_i32 => &a,
                                                        _ =>
                                                        Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&Data_EuclideanRing_gcd_004044_002d1.Value,
                                                                                                                                                                                                  &&&dictEq),
                                                                                                                                                               &&&dictEuclideanRing),
                                                                                                                            b),
                                                                                         &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_mod(),
                                                                                                                                                                                                  &&&dictEuclideanRing),
                                                                                                                                                               &&&a),
                                                                                                                            b)),
                                                    }
                                                }
                                        })
                    })
    }
    pub fn Data_EuclideanRing_gcd() -> &dyn Any {
        static Data_EuclideanRing_gcd: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_gcd.get_or_init(||
                                               Data_EuclideanRing_gcd_004044_002d1.Value)
    }
    pub fn Data_EuclideanRing_euclideanRingNumber() -> &dyn Any {
        static Data_EuclideanRing_euclideanRingNumber:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_euclideanRingNumber.get_or_init(||
                                                               Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_EuclideanRingusd_Dict(),
                                                                                                &&&add(string("degree"),
                                                                                                       &&Func1::new(move
                                                                                                                        |v|
                                                                                                                        &1_i32),
                                                                                                       add(string("div"),
                                                                                                           &&PureScript_Data_EuclideanRing::Data_EuclideanRing_numDiv(),
                                                                                                           add(string("mod"),
                                                                                                               &&Func1::new(move
                                                                                                                                |v_1|
                                                                                                                                &Func1::new(move
                                                                                                                                                |v1|
                                                                                                                                                &0.0_f64)),
                                                                                                               add(string("CommutativeRing0"),
                                                                                                                   &&Func1::new(move
                                                                                                                                    |usd__unused|
                                                                                                                                    &PureScript_Data_CommutativeRing::Data_CommutativeRing_commutativeRingNumber()),
                                                                                                                   empty::<string,
                                                                                                                           &dyn Any>()))))))
    }
    pub fn Data_EuclideanRing_euclideanRingInt() -> &dyn Any {
        static Data_EuclideanRing_euclideanRingInt: MutCell<Option<&dyn Any>>
         =
            MutCell::new(None);
        Data_EuclideanRing_euclideanRingInt.get_or_init(||
                                                            Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_EuclideanRingusd_Dict(),
                                                                                             &&&add(string("degree"),
                                                                                                    &&PureScript_Data_EuclideanRing::Data_EuclideanRing_intDegree(),
                                                                                                    add(string("div"),
                                                                                                        &&PureScript_Data_EuclideanRing::Data_EuclideanRing_intDiv(),
                                                                                                        add(string("mod"),
                                                                                                            &&PureScript_Data_EuclideanRing::Data_EuclideanRing_intMod(),
                                                                                                            add(string("CommutativeRing0"),
                                                                                                                &&Func1::new(move
                                                                                                                                 |usd__unused|
                                                                                                                                 &PureScript_Data_CommutativeRing::Data_CommutativeRing_commutativeRingInt()),
                                                                                                                empty::<string,
                                                                                                                        &dyn Any>()))))))
    }
    pub fn Data_EuclideanRing_div() -> &dyn Any {
        static Data_EuclideanRing_div: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_div.get_or_init(||
                                               &Func1::new(move |dict|
                                                               find(string("div"),
                                                                    Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
    pub fn Data_EuclideanRing_lcm() -> &dyn Any {
        static Data_EuclideanRing_lcm: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_lcm.get_or_init(||
                                               &Func1::new(move |dictEq|
                                                               &Func1::new({
                                                                               let dictEq
                                                                                   =
                                                                                   dictEq.clone();
                                                                               move
                                                                                   |dictEuclideanRing|
                                                                                   {
                                                                                       let Ring0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Ring0"),
                                                                                                                                   Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&find(string("CommutativeRing0"),
                                                                                                                                                                                                    Sharpurs_Prelude::unbox(dictEuclideanRing)),
                                                                                                                                                                                             &&&Sharpurs_Prelude::Prim_undefined()))),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       let zero =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_zero(),
                                                                                                                            &&Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                                                      Sharpurs_Prelude::unbox(&&Ring0)),
                                                                                                                                                               &&&Sharpurs_Prelude::Prim_undefined()));
                                                                                       let Semiring0 =
                                                                                           Sharpurs_Prelude::sharpurs_apply(&&find(string("Semiring0"),
                                                                                                                                   Sharpurs_Prelude::unbox(&&Ring0)),
                                                                                                                            &&&Sharpurs_Prelude::Prim_undefined());
                                                                                       &Func1::new({
                                                                                                       let Semiring0
                                                                                                           =
                                                                                                           Semiring0.clone();
                                                                                                       let dictEuclideanRing
                                                                                                           =
                                                                                                           dictEuclideanRing.clone();
                                                                                                       let zero
                                                                                                           =
                                                                                                           zero.clone();
                                                                                                       move
                                                                                                           |a|
                                                                                                           &Func1::new({
                                                                                                                           let a
                                                                                                                               =
                                                                                                                               a.clone();
                                                                                                                           move
                                                                                                                               |b|
                                                                                                                               {
                                                                                                                                   let matchValue =
                                                                                                                                       Sharpurs_Prelude::unbox(&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_disj(),
                                                                                                                                                                                                                                                                       &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                                                             &&&dictEq),
                                                                                                                                                                                                                                                                                                          &&&a),
                                                                                                                                                                                                                                                                       &&&zero)),
                                                                                                                                                                                                 &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Eq::Data_Eq_eq(),
                                                                                                                                                                                                                                                                                                          &&&dictEq),
                                                                                                                                                                                                                                                                       b),
                                                                                                                                                                                                                                    &&&zero)));
                                                                                                                                   match &Sharpurs_Prelude::_007cLitBool_007c__007c(true,
                                                                                                                                                                                    &matchValue)
                                                                                                                                       {
                                                                                                                                       0_i32
                                                                                                                                       =>
                                                                                                                                       &zero,
                                                                                                                                       _
                                                                                                                                       =>
                                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                              &&&dictEuclideanRing),
                                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                    &&&Semiring0),
                                                                                                                                                                                                                                                                                 &&&a),
                                                                                                                                                                                                                                              b)),
                                                                                                                                                                        &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_gcd(),
                                                                                                                                                                                                                                                                                                                    &&&dictEq),
                                                                                                                                                                                                                                                                                 &&&dictEuclideanRing),
                                                                                                                                                                                                                                              &&&a),
                                                                                                                                                                                                           b)),
                                                                                                                                   }
                                                                                                                               }
                                                                                                                       })
                                                                                                   })
                                                                                   }
                                                                           })))
    }
    pub fn Data_EuclideanRing_degree() -> &dyn Any {
        static Data_EuclideanRing_degree: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_EuclideanRing_degree.get_or_init(||
                                                  &Func1::new(move |dict|
                                                                  find(string("degree"),
                                                                       Sharpurs_Prelude::unbox(&&Sharpurs_Prelude::unbox(dict)))))
    }
}

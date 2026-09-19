pub mod PureScript_Data_Number_Approximate {
    use super::*;
    use fable_library_rust::Native_::Any;
    use fable_library_rust::Native_::Func1;
    use fable_library_rust::Native_::MutCell;
    use crate::module_26d9fe5f::PureScript_Data_EuclideanRing;
    use crate::module_c0687feb::PureScript_Data_HeytingAlgebra;
    use crate::module_91782ab6::PureScript_Data_Number;
    use crate::module_28ab9c8c::PureScript_Data_Ord;
    use crate::module_111ec07::PureScript_Data_Ring;
    use crate::module_be45b155::PureScript_Data_Semiring;
    use crate::module_aa21d1e7::Sharpurs_Prelude;
    pub fn Data_Number_Approximate_Tolerance() -> &dyn Any {
        static Data_Number_Approximate_Tolerance: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Approximate_Tolerance.get_or_init(||
                                                          &Func1::new(move |x|
                                                                          x.clone()))
    }
    pub fn Data_Number_Approximate_Fraction() -> &dyn Any {
        static Data_Number_Approximate_Fraction: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Approximate_Fraction.get_or_init(||
                                                         &Func1::new(move |x|
                                                                         x.clone()))
    }
    pub fn Data_Number_Approximate_eqRelative() -> &dyn Any {
        static Data_Number_Approximate_eqRelative: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Approximate_eqRelative.get_or_init(||
                                                           &Func1::new(move
                                                                           |v|
                                                                           &Func1::new({
                                                                                           let v
                                                                                               =
                                                                                               v.clone();
                                                                                           move
                                                                                               |v1|
                                                                                               &Func1::new({
                                                                                                               let v1
                                                                                                                   =
                                                                                                                   v1.clone();
                                                                                                               move
                                                                                                                   |v2|
                                                                                                                   {
                                                                                                                       let matchValue =
                                                                                                                           Sharpurs_Prelude::unbox(&&v);
                                                                                                                       let matchValue_1 =
                                                                                                                           Sharpurs_Prelude::unbox(&&v1);
                                                                                                                       let matchValue_2 =
                                                                                                                           Sharpurs_Prelude::unbox(v2);
                                                                                                                       match &Sharpurs_Prelude::_007cLitNumber_007c__007c(0.0_f64,
                                                                                                                                                                          &matchValue_1)
                                                                                                                           {
                                                                                                                           0_i32
                                                                                                                           =>
                                                                                                                           Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                  &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                               &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number::Data_Number_abs(),
                                                                                                                                                                                                                                  y)),
                                                                                                                                                            frac),
                                                                                                                           _
                                                                                                                           =>
                                                                                                                           match &Sharpurs_Prelude::_007cLitNumber_007c__007c(0.0_f64,
                                                                                                                                                                              &matchValue_2)
                                                                                                                               {
                                                                                                                               0_i32
                                                                                                                               =>
                                                                                                                               Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                      &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                   &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number::Data_Number_abs(),
                                                                                                                                                                                                                                      x)),
                                                                                                                                                                frac_1),
                                                                                                                               _
                                                                                                                               =>
                                                                                                                               {
                                                                                                                                   let y_1 =
                                                                                                                                       matchValue_2;
                                                                                                                                   let x_1 =
                                                                                                                                       matchValue_1;
                                                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                                          &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                                       &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number::Data_Number_abs(),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                                                                                                                                                                                                                                &&&x_1),
                                                                                                                                                                                                                                                                             &&&y_1))),
                                                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_EuclideanRing::Data_EuclideanRing_div(),
                                                                                                                                                                                                                                                                             &&&PureScript_Data_EuclideanRing::Data_EuclideanRing_euclideanRingNumber()),
                                                                                                                                                                                                                                          &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_mul(),
                                                                                                                                                                                                                                                                                                                                                   &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                                                                                                                &&&matchValue),
                                                                                                                                                                                                                                                                             &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number::Data_Number_abs(),
                                                                                                                                                                                                                                                                                                                &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Semiring::Data_Semiring_add(),
                                                                                                                                                                                                                                                                                                                                                                                                                         &&&PureScript_Data_Semiring::Data_Semiring_semiringNumber()),
                                                                                                                                                                                                                                                                                                                                                                                      &&&x_1),
                                                                                                                                                                                                                                                                                                                                                   &&&y_1)))),
                                                                                                                                                                                                       &&&2.0_f64))
                                                                                                                               }
                                                                                                                           },
                                                                                                                       }
                                                                                                                   }
                                                                                                           })
                                                                                       })))
    }
    pub fn Data_Number_Approximate_eqApproximate() -> &dyn Any {
        static Data_Number_Approximate_eqApproximate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Approximate_eqApproximate.get_or_init(||
                                                              {
                                                                  let onePPM =
                                                                      Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number_Approximate::Data_Number_Approximate_Fraction(),
                                                                                                       &&&1E-06_f64);
                                                                  Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number_Approximate::Data_Number_Approximate_eqRelative(),
                                                                                                   &&&onePPM)
                                                              })
    }
    pub fn Data_Number_Approximate_neqApproximate() -> &dyn Any {
        static Data_Number_Approximate_neqApproximate:
         MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Approximate_neqApproximate.get_or_init(||
                                                               &Func1::new(move
                                                                               |x|
                                                                               &Func1::new({
                                                                                               let x
                                                                                                   =
                                                                                                   x.clone();
                                                                                               move
                                                                                                   |y|
                                                                                                   Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_not(),
                                                                                                                                                                       &&&PureScript_Data_HeytingAlgebra::Data_HeytingAlgebra_heytingAlgebraBoolean()),
                                                                                                                                    &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number_Approximate::Data_Number_Approximate_eqApproximate(),
                                                                                                                                                                                                          &&&x),
                                                                                                                                                                       y))
                                                                                           })))
    }
    pub fn Data_Number_Approximate_eqAbsolute() -> &dyn Any {
        static Data_Number_Approximate_eqAbsolute: MutCell<Option<&dyn Any>> =
            MutCell::new(None);
        Data_Number_Approximate_eqAbsolute.get_or_init(||
                                                           &Func1::new(move
                                                                           |v|
                                                                           &Func1::new({
                                                                                           let v
                                                                                               =
                                                                                               v.clone();
                                                                                           move
                                                                                               |x|
                                                                                               &Func1::new({
                                                                                                               let x
                                                                                                                   =
                                                                                                                   x.clone();
                                                                                                               move
                                                                                                                   |y|
                                                                                                                   {
                                                                                                                       let matchValue =
                                                                                                                           Sharpurs_Prelude::unbox(&&v);
                                                                                                                       let matchValue_1 =
                                                                                                                           Sharpurs_Prelude::unbox(&&x);
                                                                                                                       let matchValue_2 =
                                                                                                                           Sharpurs_Prelude::unbox(y);
                                                                                                                       Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ord::Data_Ord_lessThanOrEq(),
                                                                                                                                                                                                                              &&&PureScript_Data_Ord::Data_Ord_ordNumber()),
                                                                                                                                                                                           &&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Number::Data_Number_abs(),
                                                                                                                                                                                                                              &&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&Sharpurs_Prelude::sharpurs_apply(&&&PureScript_Data_Ring::Data_Ring_sub(),
                                                                                                                                                                                                                                                                                                                                       &&&PureScript_Data_Ring::Data_Ring_ringNumber()),
                                                                                                                                                                                                                                                                                                    &&&matchValue_1),
                                                                                                                                                                                                                                                                 &&&matchValue_2))),
                                                                                                                                                        &&&matchValue)
                                                                                                                   }
                                                                                                           })
                                                                                       })))
    }
}

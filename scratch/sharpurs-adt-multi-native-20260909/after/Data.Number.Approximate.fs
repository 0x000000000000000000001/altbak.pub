[<AutoOpen>]
module PureScript_Data_Number_Approximate

open System
open System.Collections.Generic

let Data_Number_Approximate_Tolerance  = (box (fun (x: obj) -> (box x)))

let Data_Number_Approximate_Fraction  = (box (fun (x: obj) -> (box x)))

let Data_Number_Approximate_eqRelative  = (box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (match (((unbox ((box v))), (unbox ((box v1))), (unbox ((box v2))))) with | (frac, LitNumber 0.0 (), y) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((sharpurs_apply (box ((box Data_Number_abs))) (box ((box y))))))))) (box ((box frac))))) | (frac, x, LitNumber 0.0 ()) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((sharpurs_apply (box ((box Data_Number_abs))) (box ((box x))))))))) (box ((box frac))))) | (frac, x, y) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((sharpurs_apply (box ((box Data_Number_abs))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringNumber)))))) (box ((box x)))))) (box ((box y)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box Data_EuclideanRing_euclideanRingNumber)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box frac)))))) (box ((sharpurs_apply (box ((box Data_Number_abs))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))))))))) (box ((box 2.0)))))))))))))))

let Data_Number_Approximate_eqApproximate  = (let onePPM = (sharpurs_apply (box ((box Data_Number_Approximate_Fraction))) (box ((box 0.000001)))) in (sharpurs_apply (box ((box Data_Number_Approximate_eqRelative))) (box ((box onePPM)))))

let Data_Number_Approximate_neqApproximate  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_not))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Number_Approximate_eqApproximate))) (box ((box x)))))) (box ((box y)))))))))))

let Data_Number_Approximate_eqAbsolute  = (box (fun (v: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box v))), (unbox ((box x))), (unbox ((box y))))) with | (tolerance, x1, y1) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((sharpurs_apply (box ((box Data_Number_abs))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringNumber)))))) (box ((box x1)))))) (box ((box y1)))))))))))) (box ((box tolerance))))))))))))

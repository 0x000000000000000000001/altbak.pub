[<AutoOpen>]
module PureScript_Test_ArrayOps

open System
open System.Collections.Generic

let Test_ArrayOps_add  = (sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt))))

let Test_ArrayOps_range  = (box (fun (start: obj) -> (box (fun (end_var: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_range))) (box ((box start)))))) (box ((box end_var))))))))

let Test_ArrayOps_filterEvens  = (box (fun (arr: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_filter))) (box ((box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_mod))) (box ((box Data_EuclideanRing_euclideanRingInt)))))) (box ((box x)))))) (box ((box 2))))))))) (box ((box 0))))))))))) (box ((box arr))))))

let Test_ArrayOps_sumEvens  = (box (fun (n: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_foldl))) (box ((box Test_ArrayOps_add)))))) (box ((box 0)))))) (box ((sharpurs_apply (box ((box Test_ArrayOps_filterEvens))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_ArrayOps_range))) (box ((box 1)))))) (box ((box n))))))))))))

let Test_ArrayOps_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Array Processing (900 elements):"))))

let Test_ArrayOps_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 900))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ArrayOps_sumEvens))) (box ((box dummy)))))))))))))))

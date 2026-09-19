[<AutoOpen>]
module PureScript_Test_TCO

open System
open System.Collections.Generic

let Test_TCO_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Tail Call Optimization (100k calls):"))))

let rec Test_TCO_deepTailRec_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((Test_TCO_deepTailRec_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((box ((unbox<int> (box ((box acc)))) + (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_mod))) (box ((box Data_EuclideanRing_euclideanRingInt)))))) (box ((box n)))))) (box ((box 3))))))))))))))
and Test_TCO_deepTailRec = box (fun (v: obj) ->  (fun (v1: obj) -> Test_TCO_deepTailRec_tco v v1))


let Test_TCO_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_TCO_deepTailRec))) (box ((box dummy)))))) (box ((box 0)))))))))))))))

[<AutoOpen>]
module PureScript_Test_ListOps

open System
open System.Collections.Generic

type Test_ListOps_List =
  | Test_ListOps_Nilusd_Ctor
  | Test_ListOps_Consusd_Ctor of obj * obj

let Test_ListOps_add  = (sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt))))

let Test_ListOps_Nil  = (box Test_ListOps_Nilusd_Ctor)

let Test_ListOps_Cons  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Test_ListOps_Consusd_Ctor(usd__arg1, usd__arg2)))))))

let Test_ListOps_range  = (box (fun (start: obj) -> (box (fun (end_var: obj) -> (
                                                                                                                                                                                                        
                                                                                                                                                                                                        let rec go_tco (curr: obj) (acc: obj) : obj = ((match ((unbox ((box ((unbox<int> (box ((box curr)))) < (unbox<int> (box ((box start))))))))) with | LitBool true () -> ((box acc)) | _ -> ((go_tco ((box ((unbox<int> (box ((box curr)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Test_ListOps_Consusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box curr)))))) (box ((box acc))))))))) 
                                                                                                                                                                                                        and go = box ((fun (curr: obj) -> (fun (acc: obj) -> go_tco curr acc))) 
                                                                                                                                                                                                        in
                                                                                                                                                                                                        (go_tco ((box end_var)) ((box Test_ListOps_Nilusd_Ctor)))
                                                                                                                                                                                                        )))))

let rec Test_ListOps_foldl_tco (v: obj) (v1: obj) (v2: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))), (unbox ((box v2))))) with | (_, acc, Test_ListOps_Nilusd_Ctor) -> ((box acc)) | (f, acc, Test_ListOps_Consusd_Ctor(x, xs)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_ListOps_foldl))) (box ((box f)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((box acc)))))) (box ((box x))))))))) (box ((box xs)))))))
and Test_ListOps_foldl = box (fun (v: obj) ->  (fun (v1: obj) ->  (fun (v2: obj) -> Test_ListOps_foldl_tco v v1 v2)))


let Test_ListOps_filterEvens  = (box (fun (lst: obj) -> (
                                                                                                                                                                                                        
                                                                                                                                                                                                        let rec go_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (Test_ListOps_Nilusd_Ctor, acc) -> ((box acc)) | (Test_ListOps_Consusd_Ctor(x, xs), acc) -> ((match ((unbox ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_mod))) (box ((box Data_EuclideanRing_euclideanRingInt)))))) (box ((box x)))))) (box ((box 2))))))))) (box ((box 0))))))) with | LitBool true () -> ((go_tco ((box xs)) ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Test_ListOps_Consusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((box x)))))) (box ((box acc))))))) | _ -> ((go_tco ((box xs)) ((box acc)))))))) 
                                                                                                                                                                                                        and go = box ((fun (v: obj) -> (fun (v1: obj) -> go_tco v v1))) 
                                                                                                                                                                                                        in
                                                                                                                                                                                                        (go_tco ((box lst)) ((box Test_ListOps_Nilusd_Ctor)))
                                                                                                                                                                                                        )))

let Test_ListOps_sumEvens  = (box (fun (n: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_ListOps_foldl))) (box ((box Test_ListOps_add)))))) (box ((box 0)))))) (box ((sharpurs_apply (box ((box Test_ListOps_filterEvens))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_ListOps_range))) (box ((box 1)))))) (box ((box n))))))))))))

let Test_ListOps_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "List Processing (900 elements):"))))

let Test_ListOps_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 900))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ListOps_sumEvens))) (box ((box dummy)))))))))))))))

[<AutoOpen>]
module PureScript_Test_Polymorphism

open System
open System.Collections.Generic

let Test_Polymorphism_Monoidishusd_Dict  = (box (fun (x: obj) -> (box x)))

let Test_Polymorphism_mempty_  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "mempty_" (unbox<Map<string, obj>> ((box v))))))))

let Test_Polymorphism_mappend_  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "mappend_" (unbox<Map<string, obj>> ((box v))))))))

let Test_Polymorphism_polyLoop  = (box (fun (dictMonoidish: obj) -> (box (fun (n_init: obj) -> (box (fun (acc_init: obj) -> (
                                                                                                                                                                                                        
                                                                                                                                                                                                        let rec go_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((go_tco ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((box n)))))) (box ((box 1))))) ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_Polymorphism_mappend_))) (box ((box dictMonoidish)))))) (box ((box acc)))))) (box ((sharpurs_apply (box ((box Test_Polymorphism_mempty_))) (box ((box dictMonoidish)))))))))))) 
                                                                                                                                                                                                        and go = box ((fun (v: obj) -> (fun (v1: obj) -> go_tco v v1))) 
                                                                                                                                                                                                        in
                                                                                                                                                                                                        (go_tco ((box n_init)) ((box acc_init)))
                                                                                                                                                                                                        )))))))

let Test_Polymorphism_intMonoidish  = (sharpurs_apply (box ((box Test_Polymorphism_Monoidishusd_Dict))) (box ((box ((Map.add "mempty_" (box ((box 1))) (Map.add "mappend_" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringInt)))))) (box ((box x)))))) (box ((box y)))))))))) Map.empty)))))))

let Test_Polymorphism_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Polymorphism (10M Type Class Dict Lookups):"))))

let Test_Polymorphism_act  = (sharpurs_apply (box ((sharpurs_apply (box (Effect_bindE)) (box ((sharpurs_apply (box (Bench_opaque)) (box ((box (10000000)))))))))) (box ((box (fun (sharpurs_o_0: obj) -> (sharpurs_apply (box (Effect_pureE)) (box ((sharpurs_apply (box (Data_Show_showIntImpl)) (box ((let rec sharpurs_int_kernel (sharpurs_i_2: int) (sharpurs_i_3: int) : int = (if (sharpurs_i_2 = (0)) then sharpurs_i_3 else (sharpurs_int_kernel (sharpurs_i_2 - (1)) (sharpurs_i_3 + (1)))) in box (sharpurs_int_kernel (unbox<int> sharpurs_o_0) (0))))))))))))))

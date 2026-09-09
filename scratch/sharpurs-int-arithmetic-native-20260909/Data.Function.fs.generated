[<AutoOpen>]
module PureScript_Data_Function

open System
open System.Collections.Generic

let Data_Function_on  = (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((sharpurs_apply (box ((box g))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box g))) (box ((box y)))))))))))))))

let Data_Function_flip  = (box (fun (f: obj) -> (box (fun (b: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box f))) (box ((box a)))))) (box ((box b))))))))))

let Data_Function_const  = (box (fun (a: obj) -> (box (fun (v: obj) -> (box a)))))

let Data_Function_applyN  = (box (fun (f: obj) -> (
                                                                                                                                                                                                        
                                                                                                                                                                                                        let rec go_tco (n: obj) (acc: obj) : obj = ((match (((unbox ((box n))), (unbox ((box acc))))) with | (n1, acc1) when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordInt)))))) (box ((box n1)))))) (box ((box 0))))) -> ((box acc1)) | (n1, acc1) when (unbox (box Data_Boolean_otherwise)) -> ((go_tco ((box ((unbox<int> (box ((box n1)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((box f))) (box ((box acc1))))))))) 
                                                                                                                                                                                                        and go = box ((fun (n: obj) -> (fun (acc: obj) -> go_tco n acc))) 
                                                                                                                                                                                                        in
                                                                                                                                                                                                        (box go)
                                                                                                                                                                                                        )))

let Data_Function_applyFlipped  = (box (fun (x: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((box f))) (box ((box x))))))))

let Data_Function_apply  = (box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box f))) (box ((box x))))))))

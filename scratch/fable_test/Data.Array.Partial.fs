[<AutoOpen>]
module PureScript_Data_Array_Partial

open System
open System.Collections.Generic

let Data_Array_Partial_tail  = (box (fun (usd__unused: obj) -> (box (fun (xs: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_slice))) (box ((box 1)))))) (box ((sharpurs_apply (box ((box Data_Array_length))) (box ((box xs))))))))) (box ((box xs))))))))

let Data_Array_Partial_last  = (box (fun (usd__unused: obj) -> (box (fun (xs: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_unsafeIndex))) (box ((box Prim_undefined)))))) (box ((box xs)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((sharpurs_apply (box ((box Data_Array_length))) (box ((box xs))))))))) (box ((box 1)))))))))))

let Data_Array_Partial_init  = (box (fun (usd__unused: obj) -> (box (fun (xs: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_slice))) (box ((box 0)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_sub))) (box ((box Data_Ring_ringInt)))))) (box ((sharpurs_apply (box ((box Data_Array_length))) (box ((box xs))))))))) (box ((box 1))))))))) (box ((box xs))))))))

let Data_Array_Partial_head  = (box (fun (usd__unused: obj) -> (box (fun (xs: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Array_unsafeIndex))) (box ((box Prim_undefined)))))) (box ((box xs)))))) (box ((box 0))))))))

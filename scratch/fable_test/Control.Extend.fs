[<AutoOpen>]
module PureScript_Control_Extend

open System
open System.Collections.Generic

module Control_Extend_FFI =
    let arrayExtend (f: obj) =
        let f' = f :?> (obj -> obj)
        fun (xs: obj) ->
            let arr = xs :?> obj[]
            let res = Array.zeroCreate arr.Length
            for i = 0 to arr.Length - 1 do
                res.[i] <- f' (arr.[i..] :> obj)
            res :> obj
    

let Control_Extend_arrayExtend = box (fun (arg0: obj) -> box (Control_Extend_FFI.``arrayExtend`` (unbox arg0)))


let Control_Extend_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Control_Extend_Extendusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Extend_extendFn  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Control_Extend_Extendusd_Dict))) (box ((box ((Map.add "extend" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (w: obj) -> (sharpurs_apply (box ((box f))) (box ((box (fun (w_prime: obj) -> (sharpurs_apply (box ((box g))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))))) (box ((box w)))))) (box ((box w_prime)))))))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorFn))))) Map.empty)))))))))

let Control_Extend_extendArray  = (sharpurs_apply (box ((box Control_Extend_Extendusd_Dict))) (box ((box ((Map.add "extend" (box ((box Control_Extend_arrayExtend))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorArray))))) Map.empty)))))))

let Control_Extend_extend  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "extend" (unbox<Map<string, obj>> ((box v))))))))

let Control_Extend_extendFlipped  = (box (fun (dictExtend: obj) -> (box (fun (w: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Extend_extend))) (box ((box dictExtend)))))) (box ((box f)))))) (box ((box w))))))))))

let Control_Extend_duplicate  = (box (fun (dictExtend: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Extend_extend))) (box ((box dictExtend)))))) (box ((box Control_Extend_identity))))))

let Control_Extend_composeCoKleisliFlipped  = (box (fun (dictExtend: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (w: obj) -> (sharpurs_apply (box ((box f))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Extend_extend))) (box ((box dictExtend)))))) (box ((box g)))))) (box ((box w)))))))))))))))

let Control_Extend_composeCoKleisli  = (box (fun (dictExtend: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (w: obj) -> (sharpurs_apply (box ((box g))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Extend_extend))) (box ((box dictExtend)))))) (box ((box f)))))) (box ((box w)))))))))))))))

[<AutoOpen>]
module PureScript_Type_Equality

open System
open System.Collections.Generic

let Type_Equality_TypeEqualsusd_Dict  = (box (fun (x: obj) -> (box x)))

let Type_Equality_To  = (box (fun (x: obj) -> (box x)))

let Type_Equality_From  = (box (fun (x: obj) -> (box x)))

let Type_Equality_refl  = (sharpurs_apply (box ((box Type_Equality_TypeEqualsusd_Dict))) (box ((box ((Map.add "proof" (box ((box (fun (a: obj) -> (box a))))) (Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty)))))))

let Type_Equality_proof  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "proof" (unbox<Map<string, obj>> ((box v))))))))

let Type_Equality_to  = (box (fun (dictTypeEquals: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((box Type_Equality_proof))) (box ((box dictTypeEquals)))))) (box ((sharpurs_apply (box ((box Type_Equality_To))) (box ((box (fun (a: obj) -> (box a))))))))) in (match ((unbox ((box v)))) with | f -> ((box f))))))

let Type_Equality_from  = (box (fun (dictTypeEquals: obj) -> (let v = (sharpurs_apply (box ((sharpurs_apply (box ((box Type_Equality_proof))) (box ((box dictTypeEquals)))))) (box ((sharpurs_apply (box ((box Type_Equality_From))) (box ((box (fun (a: obj) -> (box a))))))))) in (match ((unbox ((box v)))) with | f -> ((box f))))))

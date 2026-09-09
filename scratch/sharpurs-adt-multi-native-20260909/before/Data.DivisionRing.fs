[<AutoOpen>]
module PureScript_Data_DivisionRing

open System
open System.Collections.Generic

let Data_DivisionRing_DivisionRingusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_DivisionRing_recip  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "recip" (unbox<Map<string, obj>> ((box v))))))))

let Data_DivisionRing_rightDiv  = (box (fun (dictDivisionRing: obj) -> (let Semiring0 = (sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Ring0" (unbox<Map<string, obj>> ((box dictDivisionRing)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined)))) in (box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Semiring0)))))) (box ((box a)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_DivisionRing_recip))) (box ((box dictDivisionRing)))))) (box ((box b))))))))))))))

let Data_DivisionRing_leftDiv  = (box (fun (dictDivisionRing: obj) -> (let Semiring0 = (sharpurs_apply (box ((Map.find "Semiring0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Ring0" (unbox<Map<string, obj>> ((box dictDivisionRing)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined)))) in (box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Semiring0)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_DivisionRing_recip))) (box ((box dictDivisionRing)))))) (box ((box b))))))))) (box ((box a)))))))))))

let Data_DivisionRing_divisionringNumber  = (sharpurs_apply (box ((box Data_DivisionRing_DivisionRingusd_Dict))) (box ((box ((Map.add "recip" (box ((box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box Data_EuclideanRing_euclideanRingNumber)))))) (box ((box 1.0)))))) (box ((box x)))))))) (Map.add "Ring0" (box ((box (fun (usd__unused: obj) -> (box Data_Ring_ringNumber))))) Map.empty)))))))

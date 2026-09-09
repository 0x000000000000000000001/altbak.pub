[<AutoOpen>]
module PureScript_Control_Alternative

open System
open System.Collections.Generic

let Control_Alternative_Alternativeusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Alternative_guard  = (box (fun (dictAlternative: obj) -> (let Applicative0 = (sharpurs_apply (box ((Map.find "Applicative0" (unbox<Map<string, obj>> ((box dictAlternative)))))) (box ((box Prim_undefined)))) in let empty = (sharpurs_apply (box ((box Control_Plus_empty))) (box ((sharpurs_apply (box ((Map.find "Plus1" (unbox<Map<string, obj>> ((box dictAlternative)))))) (box ((box Prim_undefined))))))) in (box (fun (v: obj) -> (match ((unbox ((box v)))) with | LitBool true () -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Applicative0)))))) (box ((box Data_Unit_unit))))) | LitBool false () -> ((box empty))))))))

let Control_Alternative_alternativeArray  = (sharpurs_apply (box ((box Control_Alternative_Alternativeusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Control_Applicative_applicativeArray))))) (Map.add "Plus1" (box ((box (fun (usd__unused: obj) -> (box Control_Plus_plusArray))))) Map.empty)))))))

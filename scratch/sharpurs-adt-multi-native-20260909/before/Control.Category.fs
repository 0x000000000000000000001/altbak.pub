[<AutoOpen>]
module PureScript_Control_Category

open System
open System.Collections.Generic

let Control_Category_Categoryusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Category_identity  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "identity" (unbox<Map<string, obj>> ((box v))))))))

let Control_Category_categoryFn  = (sharpurs_apply (box ((box Control_Category_Categoryusd_Dict))) (box ((box ((Map.add "identity" (box ((box (fun (x: obj) -> (box x))))) (Map.add "Semigroupoid0" (box ((box (fun (usd__unused: obj) -> (box Control_Semigroupoid_semigroupoidFn))))) Map.empty)))))))

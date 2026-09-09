[<AutoOpen>]
module PureScript_Control_Plus

open System
open System.Collections.Generic

let Control_Plus_Plususd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Plus_plusArray  = (sharpurs_apply (box ((box Control_Plus_Plususd_Dict))) (box ((box ((Map.add "empty" (box ((box [||]))) (Map.add "Alt0" (box ((box (fun (usd__unused: obj) -> (box Control_Alt_altArray))))) Map.empty)))))))

let Control_Plus_empty  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "empty" (unbox<Map<string, obj>> ((box v))))))))

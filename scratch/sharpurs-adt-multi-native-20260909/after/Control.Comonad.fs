[<AutoOpen>]
module PureScript_Control_Comonad

open System
open System.Collections.Generic

let Control_Comonad_Comonadusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Comonad_extract  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "extract" (unbox<Map<string, obj>> ((box v))))))))

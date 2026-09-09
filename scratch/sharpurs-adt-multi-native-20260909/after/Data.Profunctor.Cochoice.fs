[<AutoOpen>]
module PureScript_Data_Profunctor_Cochoice

open System
open System.Collections.Generic

let Data_Profunctor_Cochoice_Cochoiceusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_Cochoice_unright  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "unright" (unbox<Map<string, obj>> ((box v))))))))

let Data_Profunctor_Cochoice_unleft  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "unleft" (unbox<Map<string, obj>> ((box v))))))))

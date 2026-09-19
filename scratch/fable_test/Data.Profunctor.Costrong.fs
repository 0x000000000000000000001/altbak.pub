[<AutoOpen>]
module PureScript_Data_Profunctor_Costrong

open System
open System.Collections.Generic

let Data_Profunctor_Costrong_Costrongusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_Costrong_unsecond  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "unsecond" (unbox<Map<string, obj>> ((box v))))))))

let Data_Profunctor_Costrong_unfirst  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "unfirst" (unbox<Map<string, obj>> ((box v))))))))

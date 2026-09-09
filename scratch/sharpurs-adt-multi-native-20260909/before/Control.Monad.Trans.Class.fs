[<AutoOpen>]
module PureScript_Control_Monad_Trans_Class

open System
open System.Collections.Generic

let Control_Monad_Trans_Class_MonadTransusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_Trans_Class_lift  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "lift" (unbox<Map<string, obj>> ((box v))))))))

[<AutoOpen>]
module PureScript_Control_Monad_Cont_Class

open System
open System.Collections.Generic

let Control_Monad_Cont_Class_MonadContusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_Cont_Class_callCC  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "callCC" (unbox<Map<string, obj>> ((box v))))))))

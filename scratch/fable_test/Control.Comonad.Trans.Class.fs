[<AutoOpen>]
module PureScript_Control_Comonad_Trans_Class

open System
open System.Collections.Generic

let Control_Comonad_Trans_Class_ComonadTransusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Comonad_Trans_Class_lower  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "lower" (unbox<Map<string, obj>> ((box v))))))))

let Control_Comonad_Trans_Class_comonadTransIdentityT  = (sharpurs_apply (box ((box Control_Comonad_Trans_Class_ComonadTransusd_Dict))) (box ((box ((Map.add "lower" (box ((box (fun (dictComonad: obj) -> (box Control_Monad_Identity_Trans_runIdentityT))))) Map.empty))))))

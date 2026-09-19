[<AutoOpen>]
module PureScript_Effect_Class

open System
open System.Collections.Generic

let Effect_Class_MonadEffectusd_Dict  = (box (fun (x: obj) -> (box x)))

let Effect_Class_monadEffectEffect  = (sharpurs_apply (box ((box Effect_Class_MonadEffectusd_Dict))) (box ((box ((Map.add "liftEffect" (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn)))))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Effect_monadEffect))))) Map.empty)))))))

let Effect_Class_liftEffect  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "liftEffect" (unbox<Map<string, obj>> ((box v))))))))

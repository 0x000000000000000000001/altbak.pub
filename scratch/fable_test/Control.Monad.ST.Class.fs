[<AutoOpen>]
module PureScript_Control_Monad_ST_Class

open System
open System.Collections.Generic

let Control_Monad_ST_Class_MonadSTusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_ST_Class_monadSTST  = (sharpurs_apply (box ((box Control_Monad_ST_Class_MonadSTusd_Dict))) (box ((box ((Map.add "liftST" (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn)))))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_ST_Internal_monadST))))) Map.empty)))))))

let Control_Monad_ST_Class_monadSTEffect  = (sharpurs_apply (box ((box Control_Monad_ST_Class_MonadSTusd_Dict))) (box ((box ((Map.add "liftST" (box ((box Control_Monad_ST_Global_toEffect))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Effect_monadEffect))))) Map.empty)))))))

let Control_Monad_ST_Class_liftST  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "liftST" (unbox<Map<string, obj>> ((box v))))))))

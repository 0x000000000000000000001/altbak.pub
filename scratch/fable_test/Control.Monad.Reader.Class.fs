[<AutoOpen>]
module PureScript_Control_Monad_Reader_Class

open System
open System.Collections.Generic

let Control_Monad_Reader_Class_MonadAskusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_Reader_Class_MonadReaderusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_Reader_Class_monadAskFun  = (sharpurs_apply (box ((box Control_Monad_Reader_Class_MonadAskusd_Dict))) (box ((box ((Map.add "ask" (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn)))))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_monadFn))))) Map.empty)))))))

let Control_Monad_Reader_Class_monadReaderFun  = (sharpurs_apply (box ((box Control_Monad_Reader_Class_MonadReaderusd_Dict))) (box ((box ((Map.add "local" (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (Map.add "MonadAsk0" (box ((box (fun (usd__unused: obj) -> (box Control_Monad_Reader_Class_monadAskFun))))) Map.empty)))))))

let Control_Monad_Reader_Class_local  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "local" (unbox<Map<string, obj>> ((box v))))))))

let Control_Monad_Reader_Class_ask  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "ask" (unbox<Map<string, obj>> ((box v))))))))

let Control_Monad_Reader_Class_asks  = (box (fun (dictMonadAsk: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Apply0" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((sharpurs_apply (box ((Map.find "Monad0" (unbox<Map<string, obj>> ((box dictMonadAsk)))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined))))))))) (box ((box Prim_undefined)))) in let ask1 = (sharpurs_apply (box ((box Control_Monad_Reader_Class_ask))) (box ((box dictMonadAsk)))) in (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((box f)))))) (box ((box ask1)))))))))

[<AutoOpen>]
module PureScript_Control_Monad_Gen_Class

open System
open System.Collections.Generic

let Control_Monad_Gen_Class_MonadGenusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Monad_Gen_Class_sized  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "sized" (unbox<Map<string, obj>> ((box v))))))))

let Control_Monad_Gen_Class_resize  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "resize" (unbox<Map<string, obj>> ((box v))))))))

let Control_Monad_Gen_Class_chooseInt  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "chooseInt" (unbox<Map<string, obj>> ((box v))))))))

let Control_Monad_Gen_Class_chooseFloat  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "chooseFloat" (unbox<Map<string, obj>> ((box v))))))))

let Control_Monad_Gen_Class_chooseBool  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "chooseBool" (unbox<Map<string, obj>> ((box v))))))))

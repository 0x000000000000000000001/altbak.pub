[<AutoOpen>]
module PureScript_Control_Semigroupoid

open System
open System.Collections.Generic

let Control_Semigroupoid_Semigroupoidusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Semigroupoid_semigroupoidFn  = (sharpurs_apply (box ((box Control_Semigroupoid_Semigroupoidusd_Dict))) (box ((box ((Map.add "compose" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box f))) (box ((sharpurs_apply (box ((box g))) (box ((box x))))))))))))))) Map.empty))))))

let Control_Semigroupoid_compose  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "compose" (unbox<Map<string, obj>> ((box v))))))))

let Control_Semigroupoid_composeFlipped  = (box (fun (dictSemigroupoid: obj) -> (box (fun (f: obj) -> (box (fun (g: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box dictSemigroupoid)))))) (box ((box g)))))) (box ((box f))))))))))

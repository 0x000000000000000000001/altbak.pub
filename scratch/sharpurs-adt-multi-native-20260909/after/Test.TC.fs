[<AutoOpen>]
module PureScript_Test_TC

open System
open System.Collections.Generic

let Test_TC_MyClassusd_Dict  = (box (fun (x: obj) -> (box x)))

let Test_TC_myClassInt  = (sharpurs_apply (box ((box Test_TC_MyClassusd_Dict))) (box ((box ((Map.add "myMethod" (box ((box (fun (v: obj) -> (box "Int"))))) Map.empty))))))

let Test_TC_myMethod  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "myMethod" (unbox<Map<string, obj>> ((box v))))))))

let Test_TC_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "TC"))))

let Test_TC_act  = (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_TC_myMethod))) (box ((box Test_TC_myClassInt)))))) (box ((box 42)))))))

[<AutoOpen>]
module PureScript_Effect_Class_Console

open System
open System.Collections.Generic

let Effect_Class_Console_warnShow  = (box (fun (dictMonadEffect: obj) -> (let liftEffect = (sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect)))) in (box (fun (dictShow: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box liftEffect)))))) (box ((sharpurs_apply (box ((box Effect_Console_warnShow))) (box ((box dictShow))))))))))))

let Effect_Class_Console_warn  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_warn))))))

let Effect_Class_Console_timeLog  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_timeLog))))))

let Effect_Class_Console_timeEnd  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_timeEnd))))))

let Effect_Class_Console_time  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_time))))))

let Effect_Class_Console_logShow  = (box (fun (dictMonadEffect: obj) -> (let liftEffect = (sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect)))) in (box (fun (dictShow: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box liftEffect)))))) (box ((sharpurs_apply (box ((box Effect_Console_logShow))) (box ((box dictShow))))))))))))

let Effect_Class_Console_log  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_log))))))

let Effect_Class_Console_infoShow  = (box (fun (dictMonadEffect: obj) -> (let liftEffect = (sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect)))) in (box (fun (dictShow: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box liftEffect)))))) (box ((sharpurs_apply (box ((box Effect_Console_infoShow))) (box ((box dictShow))))))))))))

let Effect_Class_Console_info  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_info))))))

let Effect_Class_Console_groupEnd  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect)))))) (box ((box Effect_Console_groupEnd))))))

let Effect_Class_Console_groupCollapsed  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_groupCollapsed))))))

let Effect_Class_Console_group  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_group))))))

let Effect_Class_Console_grouped  = (box (fun (dictMonadEffect: obj) -> (let Monad0 = (sharpurs_apply (box ((Map.find "Monad0" (unbox<Map<string, obj>> ((box dictMonadEffect)))))) (box ((box Prim_undefined)))) in let Bind1 = (sharpurs_apply (box ((Map.find "Bind1" (unbox<Map<string, obj>> ((box Monad0)))))) (box ((box Prim_undefined)))) in let groupEnd1 = (sharpurs_apply (box ((box Effect_Class_Console_groupEnd))) (box ((box dictMonadEffect)))) in let Applicative0 = (sharpurs_apply (box ((Map.find "Applicative0" (unbox<Map<string, obj>> ((box Monad0)))))) (box ((box Prim_undefined)))) in (box (fun (name: obj) -> (box (fun (inner: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Bind1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Class_Console_group))) (box ((box dictMonadEffect)))))) (box ((box name))))))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Bind1)))))) (box ((box inner)))))) (box ((box (fun (result: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Bind1)))))) (box ((box groupEnd1)))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Applicative0)))))) (box ((box result))))))))))))))))))))))))))

let Effect_Class_Console_errorShow  = (box (fun (dictMonadEffect: obj) -> (let liftEffect = (sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect)))) in (box (fun (dictShow: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box liftEffect)))))) (box ((sharpurs_apply (box ((box Effect_Console_errorShow))) (box ((box dictShow))))))))))))

let Effect_Class_Console_error  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_error))))))

let Effect_Class_Console_debugShow  = (box (fun (dictMonadEffect: obj) -> (let liftEffect = (sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect)))) in (box (fun (dictShow: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box liftEffect)))))) (box ((sharpurs_apply (box ((box Effect_Console_debugShow))) (box ((box dictShow))))))))))))

let Effect_Class_Console_debug  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect))))))))) (box ((box Effect_Console_debug))))))

let Effect_Class_Console_clear  = (box (fun (dictMonadEffect: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box dictMonadEffect)))))) (box ((box Effect_Console_clear))))))

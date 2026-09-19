[<AutoOpen>]
module PureScript_Effect_Now

open System
open System.Collections.Generic

module Effect_Now_FFI =
    let now (dummy: obj) = box (float (System.DateTimeOffset.UtcNow.ToUnixTimeMilliseconds()))
    let getTimezoneOffset (dummy: obj) = 
        let offset = System.TimeZoneInfo.Local.GetUtcOffset(System.DateTime.Now).TotalMinutes
        box (-offset)
    

let Effect_Now_getTimezoneOffset = box (fun (arg0: obj) -> box (Effect_Now_FFI.``getTimezoneOffset`` (unbox arg0)))
let Effect_Now_now = box (fun (arg0: obj) -> box (Effect_Now_FFI.``now`` (unbox arg0)))


let Effect_Now_nowTime  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_functorEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_DateTime_time)))))) (box ((box Data_DateTime_Instant_toDateTime))))))))) (box ((box Effect_Now_now))))

let Effect_Now_nowDateTime  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_functorEffect)))))) (box ((box Data_DateTime_Instant_toDateTime)))))) (box ((box Effect_Now_now))))

let Effect_Now_nowDate  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_functorEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_DateTime_date)))))) (box ((box Data_DateTime_Instant_toDateTime))))))))) (box ((box Effect_Now_now))))

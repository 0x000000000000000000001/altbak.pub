[<AutoOpen>]
module PureScript_Data_DateTime_Instant

open System
open System.Collections.Generic

let Data_DateTime_Instant_fromDateTimeImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (fun (arg2: obj) -> box (fun (arg3: obj) -> box (fun (arg4: obj) -> box (fun (arg5: obj) -> box (fun (arg6: obj) -> box (Data.DateTime.Instant.FFI.FromDateTimeImpl(unbox arg0, unbox arg1, unbox arg2, unbox arg3, unbox arg4, unbox arg5, unbox arg6)))))))))
let Data_DateTime_Instant_toDateTimeImpl = box (fun (arg0: obj) -> box (fun (arg1: obj) -> box (Data.DateTime.Instant.FFI.ToDateTimeImpl(unbox arg0, unbox arg1))))


let Data_DateTime_Instant_bottom  = (sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box Data_Time_Component_boundedHour))))

let Data_DateTime_Instant_bottom1  = (sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box Data_Time_Component_boundedMinute))))

let Data_DateTime_Instant_bottom2  = (sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box Data_Time_Component_boundedSecond))))

let Data_DateTime_Instant_bottom3  = (sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box Data_Time_Component_boundedMillisecond))))

let Data_DateTime_Instant_Instant  = (box (fun (x: obj) -> (box x)))

let Data_DateTime_Instant_unInstant  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | ms -> ((box ms)))))

let Data_DateTime_Instant_toDateTime  = (let mkDateTime = (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (box (fun (y: obj) -> (box (fun (mo: obj) -> (box (fun (d: obj) -> (box (fun (h: obj) -> (box (fun (mi: obj) -> (box (fun (s: obj) -> (box (fun (ms: obj) -> (box (Data_DateTime_DateTimeusd_Ctor((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Date_canonicalDate))) (box ((box y)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Maybe_fromJust))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_toEnum))) (box ((box Data_Date_Component_boundedEnumMonth)))))) (box ((box mo)))))))))))) (box ((box d)))), (box (Data_Time_Timeusd_Ctor((box h), (box mi), (box s), (box ms)))))))))))))))))))))))))) in (sharpurs_apply (box ((box Data_DateTime_Instant_toDateTimeImpl))) (box ((box mkDateTime)))))

let Data_DateTime_Instant_showInstant  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | ms -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Instant ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Time_Duration_showMilliseconds)))))) (box ((box ms))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_DateTime_Instant_ordDateTime  = (box Data_Time_Duration_ordMilliseconds)

let Data_DateTime_Instant_instant  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | (n as ms) when (unbox (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_greaterThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((box n)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_negate))) (box ((box Data_Ring_ringNumber)))))) (box ((box 8639977881600000.0)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_lessThanOrEq))) (box ((box Data_Ord_ordNumber)))))) (box ((box n)))))) (box ((box 8639977881599999.0)))))))) -> ((box (Data_Maybe_Justusd_Ctor((sharpurs_apply (box ((box Data_DateTime_Instant_Instant))) (box ((box ms)))))))) | (n as ms) when (unbox (box Data_Boolean_otherwise)) -> ((box Data_Maybe_Nothingusd_Ctor)))))

let Data_DateTime_Instant_fromDateTime  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_DateTime_DateTimeusd_Ctor(d, t) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn7))) (box ((box Data_DateTime_Instant_fromDateTimeImpl)))))) (box ((sharpurs_apply (box ((box Data_Date_year))) (box ((box d))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_fromEnum))) (box ((box Data_Date_Component_boundedEnumMonth)))))) (box ((sharpurs_apply (box ((box Data_Date_month))) (box ((box d)))))))))))) (box ((sharpurs_apply (box ((box Data_Date_day))) (box ((box d))))))))) (box ((sharpurs_apply (box ((box Data_Time_hour))) (box ((box t))))))))) (box ((sharpurs_apply (box ((box Data_Time_minute))) (box ((box t))))))))) (box ((sharpurs_apply (box ((box Data_Time_second))) (box ((box t))))))))) (box ((sharpurs_apply (box ((box Data_Time_millisecond))) (box ((box t)))))))))))

let Data_DateTime_Instant_fromDate  = (box (fun (d: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn7))) (box ((box Data_DateTime_Instant_fromDateTimeImpl)))))) (box ((sharpurs_apply (box ((box Data_Date_year))) (box ((box d))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Enum_fromEnum))) (box ((box Data_Date_Component_boundedEnumMonth)))))) (box ((sharpurs_apply (box ((box Data_Date_month))) (box ((box d)))))))))))) (box ((sharpurs_apply (box ((box Data_Date_day))) (box ((box d))))))))) (box ((box Data_DateTime_Instant_bottom)))))) (box ((box Data_DateTime_Instant_bottom1)))))) (box ((box Data_DateTime_Instant_bottom2)))))) (box ((box Data_DateTime_Instant_bottom3))))))

let Data_DateTime_Instant_eqDateTime  = (box Data_Time_Duration_eqMilliseconds)

let Data_DateTime_Instant_diff  = (box (fun (dictDuration: obj) -> (box (fun (dt1: obj) -> (box (fun (dt2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Time_Duration_toDuration))) (box ((box dictDuration)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Time_Duration_semigroupMilliseconds)))))) (box ((sharpurs_apply (box ((box Data_DateTime_Instant_unInstant))) (box ((box dt1))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Time_Duration_negateDuration))) (box ((box Data_Time_Duration_durationMilliseconds)))))) (box ((sharpurs_apply (box ((box Data_DateTime_Instant_unInstant))) (box ((box dt2)))))))))))))))))))

let Data_DateTime_Instant_boundedInstant  = (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "bottom" (box ((sharpurs_apply (box ((box Data_DateTime_Instant_Instant))) (box ((sharpurs_apply (box ((box Data_Time_Duration_Milliseconds))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ring_negate))) (box ((box Data_Ring_ringNumber)))))) (box ((box 8639977881600000.0)))))))))))) (Map.add "top" (box ((sharpurs_apply (box ((box Data_DateTime_Instant_Instant))) (box ((sharpurs_apply (box ((box Data_Time_Duration_Milliseconds))) (box ((box 8639977881599999.0))))))))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box Data_DateTime_Instant_ordDateTime))))) Map.empty))))))))

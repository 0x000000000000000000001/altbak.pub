[<AutoOpen>]
module PureScript_Data_Interval_Duration

open System
open System.Collections.Generic

type Data_Interval_Duration_DurationComponent =
  | Data_Interval_Duration_Secondusd_Ctor
  | Data_Interval_Duration_Minuteusd_Ctor
  | Data_Interval_Duration_Hourusd_Ctor
  | Data_Interval_Duration_Dayusd_Ctor
  | Data_Interval_Duration_Weekusd_Ctor
  | Data_Interval_Duration_Monthusd_Ctor
  | Data_Interval_Duration_Yearusd_Ctor

let Data_Interval_Duration_add  = (sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber))))

let Data_Interval_Duration_Second  = (box Data_Interval_Duration_Secondusd_Ctor)

let Data_Interval_Duration_Minute  = (box Data_Interval_Duration_Minuteusd_Ctor)

let Data_Interval_Duration_Hour  = (box Data_Interval_Duration_Hourusd_Ctor)

let Data_Interval_Duration_Day  = (box Data_Interval_Duration_Dayusd_Ctor)

let Data_Interval_Duration_Week  = (box Data_Interval_Duration_Weekusd_Ctor)

let Data_Interval_Duration_Month  = (box Data_Interval_Duration_Monthusd_Ctor)

let Data_Interval_Duration_Year  = (box Data_Interval_Duration_Yearusd_Ctor)

let Data_Interval_Duration_Duration  = (box (fun (x: obj) -> (box x)))

let Data_Interval_Duration_showDurationComponent  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | Data_Interval_Duration_Minuteusd_Ctor -> ((box "Minute")) | Data_Interval_Duration_Secondusd_Ctor -> ((box "Second")) | Data_Interval_Duration_Hourusd_Ctor -> ((box "Hour")) | Data_Interval_Duration_Dayusd_Ctor -> ((box "Day")) | Data_Interval_Duration_Weekusd_Ctor -> ((box "Week")) | Data_Interval_Duration_Monthusd_Ctor -> ((box "Month")) | Data_Interval_Duration_Yearusd_Ctor -> ((box "Year"))))))) Map.empty))))))

let Data_Interval_Duration_showMap  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_showMap))) (box ((box Data_Interval_Duration_showDurationComponent)))))) (box ((box Data_Show_showNumber))))

let Data_Interval_Duration_showDuration  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | d -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Duration ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Interval_Duration_showMap)))))) (box ((box d))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_Interval_Duration_newtypeDuration  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Interval_Duration_eqDurationComponent  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (Data_Interval_Duration_Secondusd_Ctor, Data_Interval_Duration_Secondusd_Ctor) -> ((box true)) | (Data_Interval_Duration_Minuteusd_Ctor, Data_Interval_Duration_Minuteusd_Ctor) -> ((box true)) | (Data_Interval_Duration_Hourusd_Ctor, Data_Interval_Duration_Hourusd_Ctor) -> ((box true)) | (Data_Interval_Duration_Dayusd_Ctor, Data_Interval_Duration_Dayusd_Ctor) -> ((box true)) | (Data_Interval_Duration_Weekusd_Ctor, Data_Interval_Duration_Weekusd_Ctor) -> ((box true)) | (Data_Interval_Duration_Monthusd_Ctor, Data_Interval_Duration_Monthusd_Ctor) -> ((box true)) | (Data_Interval_Duration_Yearusd_Ctor, Data_Interval_Duration_Yearusd_Ctor) -> ((box true)) | (_, _) -> ((box false))))))))) Map.empty))))))

let Data_Interval_Duration_eqMap  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_eqMap))) (box ((box Data_Interval_Duration_eqDurationComponent)))))) (box ((box Data_Eq_eqNumber))))

let Data_Interval_Duration_ordDurationComponent  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (Data_Interval_Duration_Secondusd_Ctor, Data_Interval_Duration_Secondusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Interval_Duration_Secondusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (_, Data_Interval_Duration_Secondusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Interval_Duration_Minuteusd_Ctor, Data_Interval_Duration_Minuteusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Interval_Duration_Minuteusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (_, Data_Interval_Duration_Minuteusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Interval_Duration_Hourusd_Ctor, Data_Interval_Duration_Hourusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Interval_Duration_Hourusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (_, Data_Interval_Duration_Hourusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Interval_Duration_Dayusd_Ctor, Data_Interval_Duration_Dayusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Interval_Duration_Dayusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (_, Data_Interval_Duration_Dayusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Interval_Duration_Weekusd_Ctor, Data_Interval_Duration_Weekusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Interval_Duration_Weekusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (_, Data_Interval_Duration_Weekusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Interval_Duration_Monthusd_Ctor, Data_Interval_Duration_Monthusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor)) | (Data_Interval_Duration_Monthusd_Ctor, _) -> ((box Data_Ordering_LTusd_Ctor)) | (_, Data_Interval_Duration_Monthusd_Ctor) -> ((box Data_Ordering_GTusd_Ctor)) | (Data_Interval_Duration_Yearusd_Ctor, Data_Interval_Duration_Yearusd_Ctor) -> ((box Data_Ordering_EQusd_Ctor))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Interval_Duration_eqDurationComponent))))) Map.empty)))))))

let Data_Interval_Duration_ordMap  = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_ordMap))) (box ((box Data_Interval_Duration_ordDurationComponent)))))) (box ((box Data_Ord_ordNumber))))

let Data_Interval_Duration_semigroupDuration  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Interval_Duration_Duration))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_unionWith))) (box ((box Data_Interval_Duration_ordDurationComponent)))))) (box ((box Data_Interval_Duration_add)))))) (box ((box a)))))) (box ((box b))))))))))))))) Map.empty))))))

let Data_Interval_Duration_monoidDuration  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Interval_Duration_Duration))) (box ((box Data_Map_Internal_empty)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Interval_Duration_semigroupDuration))))) Map.empty)))))))

let Data_Interval_Duration_eqDuration  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Interval_Duration_eqMap)))))) (box ((box l)))))) (box ((box r)))))))))))) Map.empty))))))

let Data_Interval_Duration_ordDuration  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box Data_Interval_Duration_ordMap)))))) (box ((box l)))))) (box ((box r)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_Interval_Duration_eqDuration))))) Map.empty)))))))

let Data_Interval_Duration_durationFromComponent  = (box (fun (k: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((box Data_Interval_Duration_Duration))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_singleton))) (box ((box k)))))) (box ((box v)))))))))))

let Data_Interval_Duration_hour  = (sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Hourusd_Ctor))))

let Data_Interval_Duration_millisecond  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Secondusd_Ctor))))))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box Data_EuclideanRing_euclideanRingNumber)))))) (box ((box v)))))) (box ((box 1000.0)))))))))

let Data_Interval_Duration_minute  = (sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Minuteusd_Ctor))))

let Data_Interval_Duration_month  = (sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Monthusd_Ctor))))

let Data_Interval_Duration_second  = (sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Secondusd_Ctor))))

let Data_Interval_Duration_week  = (sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Weekusd_Ctor))))

let Data_Interval_Duration_year  = (sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Yearusd_Ctor))))

let Data_Interval_Duration_day  = (sharpurs_apply (box ((box Data_Interval_Duration_durationFromComponent))) (box ((box Data_Interval_Duration_Dayusd_Ctor))))

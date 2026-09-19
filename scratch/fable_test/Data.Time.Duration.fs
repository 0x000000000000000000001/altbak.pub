[<AutoOpen>]
module PureScript_Data_Time_Duration

open System
open System.Collections.Generic

let Data_Time_Duration_negate  = (sharpurs_apply (box ((box Data_Ring_negate))) (box ((box Data_Ring_ringNumber))))

let Data_Time_Duration_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Time_Duration_Seconds  = (box (fun (x: obj) -> (box x)))

let Data_Time_Duration_Minutes  = (box (fun (x: obj) -> (box x)))

let Data_Time_Duration_Milliseconds  = (box (fun (x: obj) -> (box x)))

let Data_Time_Duration_Hours  = (box (fun (x: obj) -> (box x)))

let Data_Time_Duration_Durationusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Time_Duration_Days  = (box (fun (x: obj) -> (box x)))

let Data_Time_Duration_toDuration  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "toDuration" (unbox<Map<string, obj>> ((box v))))))))

let Data_Time_Duration_showSeconds  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | n -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Seconds ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showNumber)))))) (box ((box n))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_Time_Duration_showMinutes  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | n -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Minutes ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showNumber)))))) (box ((box n))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_Time_Duration_showMilliseconds  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | n -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Milliseconds ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showNumber)))))) (box ((box n))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_Time_Duration_showHours  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | n -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Hours ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showNumber)))))) (box ((box n))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_Time_Duration_showDays  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | n -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Days ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showNumber)))))) (box ((box n))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_Time_Duration_semigroupSeconds  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Time_Duration_Seconds))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))

let Data_Time_Duration_semigroupMinutes  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Time_Duration_Minutes))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))

let Data_Time_Duration_semigroupMilliseconds  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Time_Duration_Milliseconds))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))

let Data_Time_Duration_semigroupHours  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Time_Duration_Hours))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))

let Data_Time_Duration_semigroupDays  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Time_Duration_Days))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))

let Data_Time_Duration_ordSeconds  = (box Data_Ord_ordNumber)

let Data_Time_Duration_ordMinutes  = (box Data_Ord_ordNumber)

let Data_Time_Duration_ordMilliseconds  = (box Data_Ord_ordNumber)

let Data_Time_Duration_ordHours  = (box Data_Ord_ordNumber)

let Data_Time_Duration_ordDays  = (box Data_Ord_ordNumber)

let Data_Time_Duration_newtypeSeconds  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Time_Duration_newtypeMinutes  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Time_Duration_newtypeMilliseconds  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Time_Duration_newtypeHours  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Time_Duration_newtypeDays  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Time_Duration_monoidSeconds  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Time_Duration_Seconds))) (box ((box 0.0)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Time_Duration_semigroupSeconds))))) Map.empty)))))))

let Data_Time_Duration_monoidMinutes  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Time_Duration_Minutes))) (box ((box 0.0)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Time_Duration_semigroupMinutes))))) Map.empty)))))))

let Data_Time_Duration_monoidMilliseconds  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Time_Duration_Milliseconds))) (box ((box 0.0)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Time_Duration_semigroupMilliseconds))))) Map.empty)))))))

let Data_Time_Duration_monoidHours  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Time_Duration_Hours))) (box ((box 0.0)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Time_Duration_semigroupHours))))) Map.empty)))))))

let Data_Time_Duration_monoidDays  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Time_Duration_Days))) (box ((box 0.0)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Time_Duration_semigroupDays))))) Map.empty)))))))

let Data_Time_Duration_fromDuration  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "fromDuration" (unbox<Map<string, obj>> ((box v))))))))

let Data_Time_Duration_negateDuration  = (box (fun (dictDuration: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Data_Time_Duration_toDuration))) (box ((box dictDuration))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Milliseconds)))))) (box ((box Data_Time_Duration_negate))))))))) (box ((sharpurs_apply (box ((box Data_Time_Duration_fromDuration))) (box ((box dictDuration))))))))))))

let Data_Time_Duration_eqSeconds  = (box Data_Eq_eqNumber)

let Data_Time_Duration_eqMinutes  = (box Data_Eq_eqNumber)

let Data_Time_Duration_eqMilliseconds  = (box Data_Eq_eqNumber)

let Data_Time_Duration_eqHours  = (box Data_Eq_eqNumber)

let Data_Time_Duration_eqDays  = (box Data_Eq_eqNumber)

let Data_Time_Duration_durationSeconds  = (sharpurs_apply (box ((box Data_Time_Duration_Durationusd_Dict))) (box ((box ((Map.add "fromDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Seconds)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box v)))))) (box ((box 1000.0))))))))))) (Map.add "toDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Milliseconds)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box Data_EuclideanRing_euclideanRingNumber)))))) (box ((box v)))))) (box ((box 1000.0))))))))))) Map.empty)))))))

let Data_Time_Duration_durationMinutes  = (sharpurs_apply (box ((box Data_Time_Duration_Durationusd_Dict))) (box ((box ((Map.add "fromDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Minutes)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box v)))))) (box ((box 60000.0))))))))))) (Map.add "toDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Milliseconds)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box Data_EuclideanRing_euclideanRingNumber)))))) (box ((box v)))))) (box ((box 60000.0))))))))))) Map.empty)))))))

let Data_Time_Duration_durationMilliseconds  = (sharpurs_apply (box ((box Data_Time_Duration_Durationusd_Dict))) (box ((box ((Map.add "fromDuration" (box ((box Data_Time_Duration_identity))) (Map.add "toDuration" (box ((box Data_Time_Duration_identity))) Map.empty)))))))

let Data_Time_Duration_durationHours  = (sharpurs_apply (box ((box Data_Time_Duration_Durationusd_Dict))) (box ((box ((Map.add "fromDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Hours)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box v)))))) (box ((box 3600000.0))))))))))) (Map.add "toDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Milliseconds)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box Data_EuclideanRing_euclideanRingNumber)))))) (box ((box v)))))) (box ((box 3600000.0))))))))))) Map.empty)))))))

let Data_Time_Duration_durationDays  = (sharpurs_apply (box ((box Data_Time_Duration_Durationusd_Dict))) (box ((box ((Map.add "fromDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Days)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Data_Semiring_semiringNumber)))))) (box ((box v)))))) (box ((box 86400000.0))))))))))) (Map.add "toDuration" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Newtype_over))) (box ((box Prim_undefined)))))) (box ((box Prim_undefined)))))) (box ((box Data_Time_Duration_Milliseconds)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_div))) (box ((box Data_EuclideanRing_euclideanRingNumber)))))) (box ((box v)))))) (box ((box 86400000.0))))))))))) Map.empty)))))))

let Data_Time_Duration_convertDuration  = (box (fun (dictDuration: obj) -> (let fromDuration1 = (sharpurs_apply (box ((box Data_Time_Duration_fromDuration))) (box ((box dictDuration)))) in (box (fun (dictDuration1: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Data_Time_Duration_toDuration))) (box ((box dictDuration1))))))))) (box ((box fromDuration1)))))))))

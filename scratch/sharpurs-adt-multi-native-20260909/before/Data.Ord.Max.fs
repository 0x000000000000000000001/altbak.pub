[<AutoOpen>]
module PureScript_Data_Ord_Max

open System
open System.Collections.Generic

let Data_Ord_Max_Max  = (box (fun (x: obj) -> (box x)))

let Data_Ord_Max_showMax  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Max ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Ord_Max_semigroupMax  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Ord_Max_Max))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_max))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))))

let Data_Ord_Max_newtypeMax  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Ord_Max_monoidMax  = (box (fun (dictBounded: obj) -> (let semigroupMax1 = (sharpurs_apply (box ((box Data_Ord_Max_semigroupMax))) (box ((sharpurs_apply (box ((Map.find "Ord0" (unbox<Map<string, obj>> ((box dictBounded)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Ord_Max_Max))) (box ((sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box dictBounded))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupMax1))))) Map.empty))))))))))

let Data_Ord_Max_eqMax  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Ord_Max_ordMax  = (box (fun (dictOrd: obj) -> (let eqMax1 = (sharpurs_apply (box ((box Data_Ord_Max_eqMax))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqMax1))))) Map.empty))))))))))

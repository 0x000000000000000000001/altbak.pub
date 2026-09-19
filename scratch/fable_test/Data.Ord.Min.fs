[<AutoOpen>]
module PureScript_Data_Ord_Min

open System
open System.Collections.Generic

let Data_Ord_Min_Min  = (box (fun (x: obj) -> (box x)))

let Data_Ord_Min_showMin  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Min ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Ord_Min_semigroupMin  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Ord_Min_Min))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_min))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y))))))))))))))) Map.empty))))))))

let Data_Ord_Min_newtypeMin  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Ord_Min_monoidMin  = (box (fun (dictBounded: obj) -> (let semigroupMin1 = (sharpurs_apply (box ((box Data_Ord_Min_semigroupMin))) (box ((sharpurs_apply (box ((Map.find "Ord0" (unbox<Map<string, obj>> ((box dictBounded)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Ord_Min_Min))) (box ((sharpurs_apply (box ((box Data_Bounded_top))) (box ((box dictBounded))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupMin1))))) Map.empty))))))))))

let Data_Ord_Min_eqMin  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Ord_Min_ordMin  = (box (fun (dictOrd: obj) -> (let eqMin1 = (sharpurs_apply (box ((box Data_Ord_Min_eqMin))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqMin1))))) Map.empty))))))))))

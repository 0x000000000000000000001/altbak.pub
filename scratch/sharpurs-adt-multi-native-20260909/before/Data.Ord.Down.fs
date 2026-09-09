[<AutoOpen>]
module PureScript_Data_Ord_Down

open System
open System.Collections.Generic

let Data_Ord_Down_Down  = (box (fun (x: obj) -> (box x)))

let Data_Ord_Down_showDown  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Down ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Ord_Down_newtypeDown  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Ord_Down_eqDown  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Ord_Down_ordDown  = (box (fun (dictOrd: obj) -> (let eqDown1 = (sharpurs_apply (box ((box Data_Ord_Down_eqDown))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Ordering_invert))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))) (box ((box x)))))) (box ((box y))))))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqDown1))))) Map.empty))))))))))

let Data_Ord_Down_boundedDown  = (box (fun (dictBounded: obj) -> (let ordDown1 = (sharpurs_apply (box ((box Data_Ord_Down_ordDown))) (box ((sharpurs_apply (box ((Map.find "Ord0" (unbox<Map<string, obj>> ((box dictBounded)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Bounded_Boundedusd_Dict))) (box ((box ((Map.add "top" (box ((sharpurs_apply (box ((box Data_Ord_Down_Down))) (box ((sharpurs_apply (box ((box Data_Bounded_bottom))) (box ((box dictBounded))))))))) (Map.add "bottom" (box ((sharpurs_apply (box ((box Data_Ord_Down_Down))) (box ((sharpurs_apply (box ((box Data_Bounded_top))) (box ((box dictBounded))))))))) (Map.add "Ord0" (box ((box (fun (usd__unused: obj) -> (box ordDown1))))) Map.empty)))))))))))

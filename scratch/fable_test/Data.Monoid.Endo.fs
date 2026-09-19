[<AutoOpen>]
module PureScript_Data_Monoid_Endo

open System
open System.Collections.Generic

let Data_Monoid_Endo_Endo  = (box (fun (x: obj) -> (box x)))

let Data_Monoid_Endo_showEndo  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Endo ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box x))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Monoid_Endo_semigroupEndo  = (box (fun (dictSemigroupoid: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Monoid_Endo_Endo))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box dictSemigroupoid)))))) (box ((box a)))))) (box ((box b))))))))))))))) Map.empty))))))))

let Data_Monoid_Endo_ordEndo  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Monoid_Endo_monoidEndo  = (box (fun (dictCategory: obj) -> (let semigroupEndo1 = (sharpurs_apply (box ((box Data_Monoid_Endo_semigroupEndo))) (box ((sharpurs_apply (box ((Map.find "Semigroupoid0" (unbox<Map<string, obj>> ((box dictCategory)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Monoid_Endo_Endo))) (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box dictCategory))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupEndo1))))) Map.empty))))))))))

let Data_Monoid_Endo_eqEndo  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Monoid_Endo_boundedEndo  = (box (fun (dictBounded: obj) -> (box dictBounded)))

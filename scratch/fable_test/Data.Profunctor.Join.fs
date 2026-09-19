[<AutoOpen>]
module PureScript_Data_Profunctor_Join

open System
open System.Collections.Generic

let Data_Profunctor_Join_Join  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_Join_showJoin  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Join ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box x))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Profunctor_Join_semigroupJoin  = (box (fun (dictSemigroupoid: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Profunctor_Join_Join))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box dictSemigroupoid)))))) (box ((box a)))))) (box ((box b))))))))))))))) Map.empty))))))))

let Data_Profunctor_Join_ordJoin  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Profunctor_Join_newtypeJoin  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Profunctor_Join_monoidJoin  = (box (fun (dictCategory: obj) -> (let semigroupJoin1 = (sharpurs_apply (box ((box Data_Profunctor_Join_semigroupJoin))) (box ((sharpurs_apply (box ((Map.find "Semigroupoid0" (unbox<Map<string, obj>> ((box dictCategory)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Profunctor_Join_Join))) (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box dictCategory))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupJoin1))))) Map.empty))))))))))

let Data_Profunctor_Join_invariantJoin  = (box (fun (dictProfunctor: obj) -> (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box g))), (unbox ((box v))))) with | (f1, g1, a) -> ((sharpurs_apply (box ((box Data_Profunctor_Join_Join))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_dimap))) (box ((box dictProfunctor)))))) (box ((box g1)))))) (box ((box f1)))))) (box ((box a))))))))))))))))) Map.empty))))))))

let Data_Profunctor_Join_eqJoin  = (box (fun (dictEq: obj) -> (box dictEq)))

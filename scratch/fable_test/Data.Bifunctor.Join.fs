[<AutoOpen>]
module PureScript_Data_Bifunctor_Join

open System
open System.Collections.Generic

let Data_Bifunctor_Join_Join  = (box (fun (x: obj) -> (box x)))

let Data_Bifunctor_Join_showJoin  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Join ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box x))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Bifunctor_Join_ordJoin  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Bifunctor_Join_newtypeJoin  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Bifunctor_Join_eqJoin  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Bifunctor_Join_bifunctorJoin  = (box (fun (dictBifunctor: obj) -> (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, a) -> ((sharpurs_apply (box ((box Data_Bifunctor_Join_Join))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bifunctor_bimap))) (box ((box dictBifunctor)))))) (box ((box f1)))))) (box ((box f1)))))) (box ((box a))))))))))))))) Map.empty))))))))

let Data_Bifunctor_Join_biapplyJoin  = (box (fun (dictBiapply: obj) -> (let bifunctorJoin1 = (sharpurs_apply (box ((box Data_Bifunctor_Join_bifunctorJoin))) (box ((sharpurs_apply (box ((Map.find "Bifunctor0" (unbox<Map<string, obj>> ((box dictBiapply)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (f, a) -> ((sharpurs_apply (box ((box Data_Bifunctor_Join_Join))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Biapply_biapply))) (box ((box dictBiapply)))))) (box ((box f)))))) (box ((box a))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box bifunctorJoin1))))) Map.empty))))))))))

let Data_Bifunctor_Join_biapplicativeJoin  = (box (fun (dictBiapplicative: obj) -> (let biapplyJoin1 = (sharpurs_apply (box ((box Data_Bifunctor_Join_biapplyJoin))) (box ((sharpurs_apply (box ((Map.find "Biapply0" (unbox<Map<string, obj>> ((box dictBiapplicative)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box Data_Bifunctor_Join_Join))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Biapplicative_bipure))) (box ((box dictBiapplicative)))))) (box ((box a)))))) (box ((box a))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box biapplyJoin1))))) Map.empty))))))))))

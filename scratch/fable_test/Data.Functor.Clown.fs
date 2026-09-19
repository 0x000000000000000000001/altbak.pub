[<AutoOpen>]
module PureScript_Data_Functor_Clown

open System
open System.Collections.Generic

let Data_Functor_Clown_Clown  = (box (fun (x: obj) -> (box x)))

let Data_Functor_Clown_showClown  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Clown ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box x))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Functor_Clown_profunctorClown  = (box (fun (dictContravariant: obj) -> (sharpurs_apply (box ((box Data_Profunctor_Profunctorusd_Dict))) (box ((box ((Map.add "dimap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, a) -> ((sharpurs_apply (box ((box Data_Functor_Clown_Clown))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_Contravariant_cmap))) (box ((box dictContravariant)))))) (box ((box f1)))))) (box ((box a))))))))))))))))) Map.empty))))))))

let Data_Functor_Clown_ordClown  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Functor_Clown_newtypeClown  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Functor_Clown_hoistClown  = (box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, a) -> ((sharpurs_apply (box ((box Data_Functor_Clown_Clown))) (box ((sharpurs_apply (box ((box f1))) (box ((box a)))))))))))))

let Data_Functor_Clown_functorClown  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (_, a) -> ((sharpurs_apply (box ((box Data_Functor_Clown_Clown))) (box ((box a)))))))))))) Map.empty))))))

let Data_Functor_Clown_eqClown  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Functor_Clown_bifunctorClown  = (box (fun (dictFunctor: obj) -> (sharpurs_apply (box ((box Data_Bifunctor_Bifunctorusd_Dict))) (box ((box ((Map.add "bimap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, a) -> ((sharpurs_apply (box ((box Data_Functor_Clown_Clown))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box f1)))))) (box ((box a))))))))))))))))) Map.empty))))))))

let Data_Functor_Clown_biapplyClown  = (box (fun (dictApply: obj) -> (let bifunctorClown1 = (sharpurs_apply (box ((box Data_Functor_Clown_bifunctorClown))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictApply)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Control_Biapply_Biapplyusd_Dict))) (box ((box ((Map.add "biapply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (fg, xy) -> ((sharpurs_apply (box ((box Data_Functor_Clown_Clown))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box dictApply)))))) (box ((box fg)))))) (box ((box xy))))))))))))))) (Map.add "Bifunctor0" (box ((box (fun (usd__unused: obj) -> (box bifunctorClown1))))) Map.empty))))))))))

let Data_Functor_Clown_biapplicativeClown  = (box (fun (dictApplicative: obj) -> (let biapplyClown1 = (sharpurs_apply (box ((box Data_Functor_Clown_biapplyClown))) (box ((sharpurs_apply (box ((Map.find "Apply0" (unbox<Map<string, obj>> ((box dictApplicative)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Control_Biapplicative_Biapplicativeusd_Dict))) (box ((box ((Map.add "bipure" (box ((box (fun (a: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((box Data_Functor_Clown_Clown))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box dictApplicative)))))) (box ((box a))))))))))))) (Map.add "Biapply0" (box ((box (fun (usd__unused: obj) -> (box biapplyClown1))))) Map.empty))))))))))

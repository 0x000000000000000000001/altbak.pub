[<AutoOpen>]
module PureScript_Data_Const

open System
open System.Collections.Generic

let Data_Const_Const  = (box (fun (x: obj) -> (box x)))

let Data_Const_showConst  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Const ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box x))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Const_semiringConst  = (box (fun (dictSemiring: obj) -> (box dictSemiring)))

let Data_Const_semigroupoidConst  = (sharpurs_apply (box ((box Control_Semigroupoid_Semigroupoidusd_Dict))) (box ((box ((Map.add "compose" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (_, x) -> ((sharpurs_apply (box ((box Data_Const_Const))) (box ((box x)))))))))))) Map.empty))))))

let Data_Const_semigroupConst  = (box (fun (dictSemigroup: obj) -> (box dictSemigroup)))

let Data_Const_ringConst  = (box (fun (dictRing: obj) -> (box dictRing)))

let Data_Const_ordConst  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Const_newtypeConst  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Const_monoidConst  = (box (fun (dictMonoid: obj) -> (box dictMonoid)))

let Data_Const_heytingAlgebraConst  = (box (fun (dictHeytingAlgebra: obj) -> (box dictHeytingAlgebra)))

let Data_Const_functorConst  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (m: obj) -> (match ((unbox ((box m)))) with | v -> ((sharpurs_apply (box ((box Data_Const_Const))) (box ((box v)))))))))))) Map.empty))))))

let Data_Const_invariantConst  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((sharpurs_apply (box ((box Data_Functor_Invariant_imapF))) (box ((box Data_Const_functorConst)))))) Map.empty))))))

let Data_Const_euclideanRingConst  = (box (fun (dictEuclideanRing: obj) -> (box dictEuclideanRing)))

let Data_Const_eqConst  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Const_eq1Const  = (box (fun (dictEq: obj) -> (let eq = (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Const_eqConst))) (box ((box dictEq))))))) in (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq1: obj) -> (box eq))))) Map.empty)))))))))

let Data_Const_ord1Const  = (box (fun (dictOrd: obj) -> (let compare = (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Const_ordConst))) (box ((box dictOrd))))))) in let eq1Const1 = (sharpurs_apply (box ((box Data_Const_eq1Const))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd1: obj) -> (box compare))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box eq1Const1))))) Map.empty))))))))))

let Data_Const_commutativeRingConst  = (box (fun (dictCommutativeRing: obj) -> (box dictCommutativeRing)))

let Data_Const_boundedConst  = (box (fun (dictBounded: obj) -> (box dictBounded)))

let Data_Const_booleanAlgebraConst  = (box (fun (dictBooleanAlgebra: obj) -> (box dictBooleanAlgebra)))

let Data_Const_applyConst  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Const_Const))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))))) (box ((box x)))))) (box ((box y))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Const_functorConst))))) Map.empty)))))))))

let Data_Const_applicativeConst  = (box (fun (dictMonoid: obj) -> (let applyConst1 = (sharpurs_apply (box ((box Data_Const_applyConst))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box Data_Const_Const))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box applyConst1))))) Map.empty))))))))))

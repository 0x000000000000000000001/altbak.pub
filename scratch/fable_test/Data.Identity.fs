[<AutoOpen>]
module PureScript_Data_Identity

open System
open System.Collections.Generic

let Data_Identity_Identity  = (box (fun (x: obj) -> (box x)))

let Data_Identity_showIdentity  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Identity ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box x))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Identity_semiringIdentity  = (box (fun (dictSemiring: obj) -> (box dictSemiring)))

let Data_Identity_semigroupIdentity  = (box (fun (dictSemigroup: obj) -> (box dictSemigroup)))

let Data_Identity_ringIdentity  = (box (fun (dictRing: obj) -> (box dictRing)))

let Data_Identity_ordIdentity  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Identity_newtypeIdentity  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Identity_monoidIdentity  = (box (fun (dictMonoid: obj) -> (box dictMonoid)))

let Data_Identity_lazyIdentity  = (box (fun (dictLazy: obj) -> (box dictLazy)))

let Data_Identity_heytingAlgebraIdentity  = (box (fun (dictHeytingAlgebra: obj) -> (box dictHeytingAlgebra)))

let Data_Identity_functorIdentity  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (m: obj) -> (match ((unbox ((box m)))) with | v -> ((sharpurs_apply (box ((box Data_Identity_Identity))) (box ((sharpurs_apply (box ((box f))) (box ((box v))))))))))))))) Map.empty))))))

let Data_Identity_invariantIdentity  = (sharpurs_apply (box ((box Data_Functor_Invariant_Invariantusd_Dict))) (box ((box ((Map.add "imap" (box ((sharpurs_apply (box ((box Data_Functor_Invariant_imapF))) (box ((box Data_Identity_functorIdentity)))))) Map.empty))))))

let Data_Identity_extendIdentity  = (sharpurs_apply (box ((box Control_Extend_Extendusd_Dict))) (box ((box ((Map.add "extend" (box ((box (fun (f: obj) -> (box (fun (m: obj) -> (sharpurs_apply (box ((box Data_Identity_Identity))) (box ((sharpurs_apply (box ((box f))) (box ((box m))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_functorIdentity))))) Map.empty)))))))

let Data_Identity_euclideanRingIdentity  = (box (fun (dictEuclideanRing: obj) -> (box dictEuclideanRing)))

let Data_Identity_eqIdentity  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Identity_eq1Identity  = (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Identity_eqIdentity))) (box ((box dictEq))))))))))) Map.empty))))))

let Data_Identity_ord1Identity  = (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Identity_ordIdentity))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_eq1Identity))))) Map.empty)))))))

let Data_Identity_comonadIdentity  = (sharpurs_apply (box ((box Control_Comonad_Comonadusd_Dict))) (box ((box ((Map.add "extract" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | x -> ((box x))))))) (Map.add "Extend0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_extendIdentity))))) Map.empty)))))))

let Data_Identity_commutativeRingIdentity  = (box (fun (dictCommutativeRing: obj) -> (box dictCommutativeRing)))

let Data_Identity_boundedIdentity  = (box (fun (dictBounded: obj) -> (box dictBounded)))

let Data_Identity_booleanAlgebraIdentity  = (box (fun (dictBooleanAlgebra: obj) -> (box dictBooleanAlgebra)))

let Data_Identity_applyIdentity  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (f, x) -> ((sharpurs_apply (box ((box Data_Identity_Identity))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_functorIdentity))))) Map.empty)))))))

let Data_Identity_bindIdentity  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box (fun (v: obj) -> (box (fun (f: obj) -> (match (((unbox ((box v))), (unbox ((box f))))) with | (m, f1) -> ((sharpurs_apply (box ((box f1))) (box ((box m)))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_applyIdentity))))) Map.empty)))))))

let Data_Identity_applicativeIdentity  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Data_Identity_Identity))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_applyIdentity))))) Map.empty)))))))

let Data_Identity_monadIdentity  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_applicativeIdentity))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_bindIdentity))))) Map.empty)))))))

let Data_Identity_altIdentity  = (sharpurs_apply (box ((box Control_Alt_Altusd_Dict))) (box ((box ((Map.add "alt" (box ((box (fun (x: obj) -> (box (fun (v: obj) -> (box x))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_functorIdentity))))) Map.empty)))))))

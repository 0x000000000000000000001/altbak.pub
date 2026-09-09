[<AutoOpen>]
module PureScript_Data_Equivalence

open System
open System.Collections.Generic

let Data_Equivalence_Equivalence  = (box (fun (x: obj) -> (box x)))

let Data_Equivalence_semigroupEquivalence  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (p, q) -> ((sharpurs_apply (box ((box Data_Equivalence_Equivalence))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box p))) (box ((box a)))))) (box ((box b))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box q))) (box ((box a)))))) (box ((box b)))))))))))))))))))))) Map.empty))))))

let Data_Equivalence_newtypeEquivalence  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Equivalence_monoidEquivalence  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Equivalence_Equivalence))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box true)))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Equivalence_semigroupEquivalence))))) Map.empty)))))))

let Data_Equivalence_defaultEquivalence  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Equivalence_Equivalence))) (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box dictEq)))))))))

let Data_Equivalence_contravariantEquivalence  = (sharpurs_apply (box ((box Data_Functor_Contravariant_Contravariantusd_Dict))) (box ((box ((Map.add "cmap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, g) -> ((sharpurs_apply (box ((box Data_Equivalence_Equivalence))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_on))) (box ((box g)))))) (box ((box f1))))))))))))))) Map.empty))))))

let Data_Equivalence_comparisonEquivalence  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | p -> ((sharpurs_apply (box ((box Data_Equivalence_Equivalence))) (box ((box (fun (a: obj) -> (box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Ordering_eqOrdering)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box p))) (box ((box a)))))) (box ((box b))))))))) (box ((box Data_Ordering_EQusd_Ctor)))))))))))))))

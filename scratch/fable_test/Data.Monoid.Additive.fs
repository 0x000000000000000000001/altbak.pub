[<AutoOpen>]
module PureScript_Data_Monoid_Additive

open System
open System.Collections.Generic

let Data_Monoid_Additive_Additive  = (box (fun (x: obj) -> (box x)))

let Data_Monoid_Additive_showAdditive  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Additive ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Monoid_Additive_semigroupAdditive  = (box (fun (dictSemiring: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Monoid_Additive_Additive))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box dictSemiring)))))) (box ((box a)))))) (box ((box b))))))))))))))) Map.empty))))))))

let Data_Monoid_Additive_ordAdditive  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Monoid_Additive_monoidAdditive  = (box (fun (dictSemiring: obj) -> (let semigroupAdditive1 = (sharpurs_apply (box ((box Data_Monoid_Additive_semigroupAdditive))) (box ((box dictSemiring)))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Monoid_Additive_Additive))) (box ((sharpurs_apply (box ((box Data_Semiring_zero))) (box ((box dictSemiring))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupAdditive1))))) Map.empty))))))))))

let Data_Monoid_Additive_functorAdditive  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (m: obj) -> (match ((unbox ((box m)))) with | v -> ((sharpurs_apply (box ((box Data_Monoid_Additive_Additive))) (box ((sharpurs_apply (box ((box f))) (box ((box v))))))))))))))) Map.empty))))))

let Data_Monoid_Additive_eqAdditive  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Monoid_Additive_eq1Additive  = (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Monoid_Additive_eqAdditive))) (box ((box dictEq))))))))))) Map.empty))))))

let Data_Monoid_Additive_ord1Additive  = (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Monoid_Additive_ordAdditive))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Additive_eq1Additive))))) Map.empty)))))))

let Data_Monoid_Additive_boundedAdditive  = (box (fun (dictBounded: obj) -> (box dictBounded)))

let Data_Monoid_Additive_applyAdditive  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (f, x) -> ((sharpurs_apply (box ((box Data_Monoid_Additive_Additive))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Additive_functorAdditive))))) Map.empty)))))))

let Data_Monoid_Additive_bindAdditive  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box (fun (v: obj) -> (box (fun (f: obj) -> (match (((unbox ((box v))), (unbox ((box f))))) with | (x, f1) -> ((sharpurs_apply (box ((box f1))) (box ((box x)))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Additive_applyAdditive))))) Map.empty)))))))

let Data_Monoid_Additive_applicativeAdditive  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Data_Monoid_Additive_Additive))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Additive_applyAdditive))))) Map.empty)))))))

let Data_Monoid_Additive_monadAdditive  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Additive_applicativeAdditive))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Additive_bindAdditive))))) Map.empty)))))))

[<AutoOpen>]
module PureScript_Data_Monoid_Conj

open System
open System.Collections.Generic

let Data_Monoid_Conj_Conj  = (box (fun (x: obj) -> (box x)))

let Data_Monoid_Conj_showConj  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Conj ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Monoid_Conj_semiringConj  = (box (fun (dictHeytingAlgebra: obj) -> (sharpurs_apply (box ((box Data_Semiring_Semiringusd_Dict))) (box ((box ((Map.add "zero" (box ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_tt))) (box ((box dictHeytingAlgebra))))))))) (Map.add "one" (box ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_ff))) (box ((box dictHeytingAlgebra))))))))) (Map.add "add" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box dictHeytingAlgebra)))))) (box ((box a)))))) (box ((box b))))))))))))))) (Map.add "mul" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_disj))) (box ((box dictHeytingAlgebra)))))) (box ((box a)))))) (box ((box b))))))))))))))) Map.empty)))))))))))

let Data_Monoid_Conj_semigroupConj  = (box (fun (dictHeytingAlgebra: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box dictHeytingAlgebra)))))) (box ((box a)))))) (box ((box b))))))))))))))) Map.empty))))))))

let Data_Monoid_Conj_ordConj  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Monoid_Conj_monoidConj  = (box (fun (dictHeytingAlgebra: obj) -> (let semigroupConj1 = (sharpurs_apply (box ((box Data_Monoid_Conj_semigroupConj))) (box ((box dictHeytingAlgebra)))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_tt))) (box ((box dictHeytingAlgebra))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupConj1))))) Map.empty))))))))))

let Data_Monoid_Conj_functorConj  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (m: obj) -> (match ((unbox ((box m)))) with | v -> ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((box f))) (box ((box v))))))))))))))) Map.empty))))))

let Data_Monoid_Conj_eqConj  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Monoid_Conj_eq1Conj  = (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Monoid_Conj_eqConj))) (box ((box dictEq))))))))))) Map.empty))))))

let Data_Monoid_Conj_ord1Conj  = (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Monoid_Conj_ordConj))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Conj_eq1Conj))))) Map.empty)))))))

let Data_Monoid_Conj_boundedConj  = (box (fun (dictBounded: obj) -> (box dictBounded)))

let Data_Monoid_Conj_applyConj  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (f, x) -> ((sharpurs_apply (box ((box Data_Monoid_Conj_Conj))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Conj_functorConj))))) Map.empty)))))))

let Data_Monoid_Conj_bindConj  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box (fun (v: obj) -> (box (fun (f: obj) -> (match (((unbox ((box v))), (unbox ((box f))))) with | (x, f1) -> ((sharpurs_apply (box ((box f1))) (box ((box x)))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Conj_applyConj))))) Map.empty)))))))

let Data_Monoid_Conj_applicativeConj  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Data_Monoid_Conj_Conj))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Conj_applyConj))))) Map.empty)))))))

let Data_Monoid_Conj_monadConj  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Conj_applicativeConj))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Conj_bindConj))))) Map.empty)))))))

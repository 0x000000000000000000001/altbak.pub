[<AutoOpen>]
module PureScript_Data_Monoid_Dual

open System
open System.Collections.Generic

let Data_Monoid_Dual_Dual  = (box (fun (x: obj) -> (box x)))

let Data_Monoid_Dual_showDual  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Dual ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Monoid_Dual_semigroupDual  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (x, y) -> ((sharpurs_apply (box ((box Data_Monoid_Dual_Dual))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))))) (box ((box y)))))) (box ((box x))))))))))))))) Map.empty))))))))

let Data_Monoid_Dual_ordDual  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Monoid_Dual_monoidDual  = (box (fun (dictMonoid: obj) -> (let semigroupDual1 = (sharpurs_apply (box ((box Data_Monoid_Dual_semigroupDual))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Monoid_Dual_Dual))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupDual1))))) Map.empty))))))))))

let Data_Monoid_Dual_functorDual  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (m: obj) -> (match ((unbox ((box m)))) with | v -> ((sharpurs_apply (box ((box Data_Monoid_Dual_Dual))) (box ((sharpurs_apply (box ((box f))) (box ((box v))))))))))))))) Map.empty))))))

let Data_Monoid_Dual_eqDual  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Monoid_Dual_eq1Dual  = (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Monoid_Dual_eqDual))) (box ((box dictEq))))))))))) Map.empty))))))

let Data_Monoid_Dual_ord1Dual  = (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Monoid_Dual_ordDual))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Dual_eq1Dual))))) Map.empty)))))))

let Data_Monoid_Dual_boundedDual  = (box (fun (dictBounded: obj) -> (box dictBounded)))

let Data_Monoid_Dual_applyDual  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (f, x) -> ((sharpurs_apply (box ((box Data_Monoid_Dual_Dual))) (box ((sharpurs_apply (box ((box f))) (box ((box x))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Dual_functorDual))))) Map.empty)))))))

let Data_Monoid_Dual_bindDual  = (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box (fun (v: obj) -> (box (fun (f: obj) -> (match (((unbox ((box v))), (unbox ((box f))))) with | (x, f1) -> ((sharpurs_apply (box ((box f1))) (box ((box x)))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Dual_applyDual))))) Map.empty)))))))

let Data_Monoid_Dual_applicativeDual  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Data_Monoid_Dual_Dual))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Dual_applyDual))))) Map.empty)))))))

let Data_Monoid_Dual_monadDual  = (sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Dual_applicativeDual))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Dual_bindDual))))) Map.empty)))))))

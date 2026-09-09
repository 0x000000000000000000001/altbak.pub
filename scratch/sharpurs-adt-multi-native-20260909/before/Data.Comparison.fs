[<AutoOpen>]
module PureScript_Data_Comparison

open System
open System.Collections.Generic

let Data_Comparison_semigroupFn  = (sharpurs_apply (box ((box Data_Semigroup_semigroupFn))) (box ((sharpurs_apply (box ((box Data_Semigroup_semigroupFn))) (box ((box Data_Ordering_semigroupOrdering)))))))

let Data_Comparison_Comparison  = (box (fun (x: obj) -> (box x)))

let Data_Comparison_semigroupComparison  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (p, q) -> ((sharpurs_apply (box ((box Data_Comparison_Comparison))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Comparison_semigroupFn)))))) (box ((box p)))))) (box ((box q))))))))))))))) Map.empty))))))

let Data_Comparison_newtypeComparison  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Comparison_monoidComparison  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Comparison_Comparison))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Ordering_EQusd_Ctor)))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Comparison_semigroupComparison))))) Map.empty)))))))

let Data_Comparison_defaultComparison  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Comparison_Comparison))) (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box dictOrd)))))))))

let Data_Comparison_contravariantComparison  = (sharpurs_apply (box ((box Data_Functor_Contravariant_Contravariantusd_Dict))) (box ((box ((Map.add "cmap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, g) -> ((sharpurs_apply (box ((box Data_Comparison_Comparison))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_on))) (box ((box g)))))) (box ((box f1))))))))))))))) Map.empty))))))

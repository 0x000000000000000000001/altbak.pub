[<AutoOpen>]
module PureScript_Data_Divisible

open System
open System.Collections.Generic

let Data_Divisible_Divisibleusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Divisible_divisiblePredicate  = (sharpurs_apply (box ((box Data_Divisible_Divisibleusd_Dict))) (box ((box ((Map.add "conquer" (box ((sharpurs_apply (box ((box Data_Predicate_Predicate))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box true))))))))) (Map.add "Divide0" (box ((box (fun (usd__unused: obj) -> (box Data_Divide_dividePredicate))))) Map.empty)))))))

let Data_Divisible_divisibleOp  = (box (fun (dictMonoid: obj) -> (let divideOp = (sharpurs_apply (box ((box Data_Divide_divideOp))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Divisible_Divisibleusd_Dict))) (box ((box ((Map.add "conquer" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_Op_Op)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid)))))))))))) (Map.add "Divide0" (box ((box (fun (usd__unused: obj) -> (box divideOp))))) Map.empty))))))))))

let Data_Divisible_divisibleEquivalence  = (sharpurs_apply (box ((box Data_Divisible_Divisibleusd_Dict))) (box ((box ((Map.add "conquer" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_Equivalence_Equivalence)))))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box true)))))))))) (Map.add "Divide0" (box ((box (fun (usd__unused: obj) -> (box Data_Divide_divideEquivalence))))) Map.empty)))))))

let Data_Divisible_divisibleComparison  = (sharpurs_apply (box ((box Data_Divisible_Divisibleusd_Dict))) (box ((box ((Map.add "conquer" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_Comparison_Comparison)))))) (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Data_Ordering_EQusd_Ctor)))))))))) (Map.add "Divide0" (box ((box (fun (usd__unused: obj) -> (box Data_Divide_divideComparison))))) Map.empty)))))))

let Data_Divisible_conquer  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "conquer" (unbox<Map<string, obj>> ((box v))))))))

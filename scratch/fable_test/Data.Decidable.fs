[<AutoOpen>]
module PureScript_Data_Decidable

open System
open System.Collections.Generic

let Data_Decidable_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Decidable_Decidableusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Decidable_lose  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "lose" (unbox<Map<string, obj>> ((box v))))))))

let Data_Decidable_lost  = (box (fun (dictDecidable: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Decidable_lose))) (box ((box dictDecidable)))))) (box ((box Data_Decidable_identity))))))

let Data_Decidable_decidablePredicate  = (sharpurs_apply (box ((box Data_Decidable_Decidableusd_Dict))) (box ((box ((Map.add "lose" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((box Data_Predicate_Predicate))) (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box Data_Void_absurd))) (box ((sharpurs_apply (box ((box f))) (box ((box a)))))))))))))))) (Map.add "Decide0" (box ((box (fun (usd__unused: obj) -> (box Data_Decide_choosePredicate))))) (Map.add "Divisible1" (box ((box (fun (usd__unused: obj) -> (box Data_Divisible_divisiblePredicate))))) Map.empty))))))))

let Data_Decidable_decidableOp  = (box (fun (dictMonoid: obj) -> (let chooseOp = (sharpurs_apply (box ((box Data_Decide_chooseOp))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in let divisibleOp = (sharpurs_apply (box ((box Data_Divisible_divisibleOp))) (box ((box dictMonoid)))) in (sharpurs_apply (box ((box Data_Decidable_Decidableusd_Dict))) (box ((box ((Map.add "lose" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((box Data_Op_Op))) (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box Data_Void_absurd))) (box ((sharpurs_apply (box ((box f))) (box ((box a)))))))))))))))) (Map.add "Decide0" (box ((box (fun (usd__unused: obj) -> (box chooseOp))))) (Map.add "Divisible1" (box ((box (fun (usd__unused: obj) -> (box divisibleOp))))) Map.empty)))))))))))

let Data_Decidable_decidableEquivalence  = (sharpurs_apply (box ((box Data_Decidable_Decidableusd_Dict))) (box ((box ((Map.add "lose" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((box Data_Equivalence_Equivalence))) (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box Data_Void_absurd))) (box ((sharpurs_apply (box ((box f))) (box ((box a)))))))))))))))) (Map.add "Decide0" (box ((box (fun (usd__unused: obj) -> (box Data_Decide_chooseEquivalence))))) (Map.add "Divisible1" (box ((box (fun (usd__unused: obj) -> (box Data_Divisible_divisibleEquivalence))))) Map.empty))))))))

let Data_Decidable_decidableComparison  = (sharpurs_apply (box ((box Data_Decidable_Decidableusd_Dict))) (box ((box ((Map.add "lose" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((box Data_Comparison_Comparison))) (box ((box (fun (a: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((box Data_Void_absurd))) (box ((sharpurs_apply (box ((box f))) (box ((box a)))))))))))))))))) (Map.add "Decide0" (box ((box (fun (usd__unused: obj) -> (box Data_Decide_chooseComparison))))) (Map.add "Divisible1" (box ((box (fun (usd__unused: obj) -> (box Data_Divisible_divisibleComparison))))) Map.empty))))))))

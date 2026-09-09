[<AutoOpen>]
module PureScript_Data_Predicate

open System
open System.Collections.Generic

let Data_Predicate_Predicate  = (box (fun (x: obj) -> (box x)))

let Data_Predicate_newtypePredicate  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Predicate_heytingAlgebraPredicate  = (sharpurs_apply (box ((box Data_HeytingAlgebra_heytingAlgebraFunction))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean))))

let Data_Predicate_contravariantPredicate  = (sharpurs_apply (box ((box Data_Functor_Contravariant_Contravariantusd_Dict))) (box ((box ((Map.add "cmap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, g) -> ((sharpurs_apply (box ((box Data_Predicate_Predicate))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box g)))))) (box ((box f1))))))))))))))) Map.empty))))))

let Data_Predicate_booleanAlgebraPredicate  = (sharpurs_apply (box ((box Data_BooleanAlgebra_booleanAlgebraFn))) (box ((box Data_BooleanAlgebra_booleanAlgebraBoolean))))

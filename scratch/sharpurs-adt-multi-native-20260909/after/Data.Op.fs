[<AutoOpen>]
module PureScript_Data_Op

open System
open System.Collections.Generic

let Data_Op_Op  = (box (fun (x: obj) -> (box x)))

let Data_Op_semigroupoidOp  = (sharpurs_apply (box ((box Control_Semigroupoid_Semigroupoidusd_Dict))) (box ((box ((Map.add "compose" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (f, g) -> ((sharpurs_apply (box ((box Data_Op_Op))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box g)))))) (box ((box f))))))))))))))) Map.empty))))))

let Data_Op_semigroupOp  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_semigroupFn))) (box ((box dictSemigroup))))))

let Data_Op_newtypeOp  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Op_monoidOp  = (box (fun (dictMonoid: obj) -> (sharpurs_apply (box ((box Data_Monoid_monoidFn))) (box ((box dictMonoid))))))

let Data_Op_contravariantOp  = (sharpurs_apply (box ((box Data_Functor_Contravariant_Contravariantusd_Dict))) (box ((box ((Map.add "cmap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, g) -> ((sharpurs_apply (box ((box Data_Op_Op))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box g)))))) (box ((box f1))))))))))))))) Map.empty))))))

let Data_Op_categoryOp  = (sharpurs_apply (box ((box Control_Category_Categoryusd_Dict))) (box ((box ((Map.add "identity" (box ((sharpurs_apply (box ((box Data_Op_Op))) (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))))))) (Map.add "Semigroupoid0" (box ((box (fun (usd__unused: obj) -> (box Data_Op_semigroupoidOp))))) Map.empty)))))))

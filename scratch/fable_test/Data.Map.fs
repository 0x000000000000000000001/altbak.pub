[<AutoOpen>]
module PureScript_Data_Map

open System
open System.Collections.Generic

let Data_Map_SemigroupMap  = (box (fun (x: obj) -> (box x)))

let Data_Map_traversableWithIndexSemigroupMap  = (box Data_Map_Internal_traversableWithIndexMap)

let Data_Map_traversableSemigroupMap  = (box Data_Map_Internal_traversableMap)

let Data_Map_showSemigroupMap  = (box (fun (dictShow: obj) -> (let showMap = (sharpurs_apply (box ((box Data_Map_Internal_showMap))) (box ((box dictShow)))) in (box (fun (dictShow1: obj) -> (sharpurs_apply (box ((box showMap))) (box ((box dictShow1)))))))))

let Data_Map_semigroupSemigroupMap  = (box (fun (dictOrd: obj) -> (box (fun (dictSemigroup: obj) -> (let append = (sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (l, r) -> ((sharpurs_apply (box ((box Data_Map_SemigroupMap))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_unionWith))) (box ((box dictOrd)))))) (box ((box append)))))) (box ((box l)))))) (box ((box r))))))))))))))) Map.empty)))))))))))

let Data_Map_plusSemigroupMap  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Map_Internal_plusMap))) (box ((box dictOrd))))))

let Data_Map_ordSemigroupMap  = (box (fun (dictOrd: obj) -> (let ordMap = (sharpurs_apply (box ((box Data_Map_Internal_ordMap))) (box ((box dictOrd)))) in (box (fun (dictOrd1: obj) -> (sharpurs_apply (box ((box ordMap))) (box ((box dictOrd1)))))))))

let Data_Map_ord1SemigroupMap  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Map_Internal_ord1Map))) (box ((box dictOrd))))))

let Data_Map_newtypeSemigroupMap  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Map_monoidSemigroupMap  = (box (fun (dictOrd: obj) -> (let semigroupSemigroupMap1 = (sharpurs_apply (box ((box Data_Map_semigroupSemigroupMap))) (box ((box dictOrd)))) in (box (fun (dictSemigroup: obj) -> (let semigroupSemigroupMap2 = (sharpurs_apply (box ((box semigroupSemigroupMap1))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Map_SemigroupMap))) (box ((box Data_Map_Internal_empty)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupSemigroupMap2))))) Map.empty)))))))))))))

let Data_Map_keys  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Set_fromMap)))))) (box ((sharpurs_apply (box ((box Data_Functor_void))) (box ((box Data_Map_Internal_functorMap)))))))

let Data_Map_functorWithIndexSemigroupMap  = (box Data_Map_Internal_functorWithIndexMap)

let Data_Map_functorSemigroupMap  = (box Data_Map_Internal_functorMap)

let Data_Map_foldableWithIndexSemigroupMap  = (box Data_Map_Internal_foldableWithIndexMap)

let Data_Map_foldableSemigroupMap  = (box Data_Map_Internal_foldableMap)

let Data_Map_eqSemigroupMap  = (box (fun (dictEq: obj) -> (let eqMap = (sharpurs_apply (box ((box Data_Map_Internal_eqMap))) (box ((box dictEq)))) in (box (fun (dictEq1: obj) -> (sharpurs_apply (box ((box eqMap))) (box ((box dictEq1)))))))))

let Data_Map_eq1SemigroupMap  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Map_Internal_eq1Map))) (box ((box dictEq))))))

let Data_Map_bindSemigroupMap  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Map_Internal_bindMap))) (box ((box dictOrd))))))

let Data_Map_applySemigroupMap  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Map_Internal_applyMap))) (box ((box dictOrd))))))

let Data_Map_altSemigroupMap  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Map_Internal_altMap))) (box ((box dictOrd))))))

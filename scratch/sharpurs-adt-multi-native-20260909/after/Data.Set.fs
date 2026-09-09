[<AutoOpen>]
module PureScript_Data_Set

open System
open System.Collections.Generic

let Data_Set_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Set_Set  = (box (fun (x: obj) -> (box x)))

let Data_Set_union  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_union))) (box ((box dictOrd)))))))))

let Data_Set_toggle  = (box (fun (dictOrd: obj) -> (box (fun (a: obj) -> (box (fun (v: obj) -> (match (((unbox ((box a))), (unbox ((box v))))) with | (a1, m) -> ((sharpurs_apply (box ((box Data_Set_Set))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_alter))) (box ((box dictOrd)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Maybe_maybe))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1)))))))) (box ((box Data_Unit_unit))))))))) (box ((box (fun (v1: obj) -> (box Data_Maybe_Nothingusd_Ctor))))))))))) (box ((box a1)))))) (box ((box m)))))))))))))))

let Data_Set_toMap  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((box s)))))

let Data_Set_toList  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | m -> ((sharpurs_apply (box ((box Data_Map_Internal_keys))) (box ((box m))))))))

let Data_Set_toUnfoldable  = (box (fun (dictUnfoldable: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Data_List_toUnfoldable))) (box ((box dictUnfoldable))))))))) (box ((box Data_Set_toList))))))

let Data_Set_size  = (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((box Data_Map_Internal_size))))

let Data_Set_singleton  = (box (fun (a: obj) -> (sharpurs_apply (box ((box Data_Set_Set))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_singleton))) (box ((box a)))))) (box ((box Data_Unit_unit)))))))))

let Data_Set_showSet  = (box (fun (dictShow: obj) -> (let showArray = (sharpurs_apply (box ((box Data_Show_showArray))) (box ((box dictShow)))) in (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(fromFoldable ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box showArray)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_toUnfoldable))) (box ((box Data_Unfoldable_unfoldableArray)))))) (box ((box s)))))))))))) (box ((box ")"))))))))))) Map.empty)))))))))

let Data_Set_semigroupSet  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((sharpurs_apply (box ((box Data_Set_union))) (box ((box dictOrd)))))) Map.empty))))))))

let Data_Set_member  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_member))) (box ((box dictOrd)))))))))

let Data_Set_isEmpty  = (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((box Data_Map_Internal_isEmpty))))

let Data_Set_intersection  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_intersection))) (box ((box dictOrd)))))))))

let Data_Set_insert  = (box (fun (dictOrd: obj) -> (box (fun (a: obj) -> (box (fun (v: obj) -> (match (((unbox ((box a))), (unbox ((box v))))) with | (a1, m) -> ((sharpurs_apply (box ((box Data_Set_Set))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_insert))) (box ((box dictOrd)))))) (box ((box a1)))))) (box ((box Data_Unit_unit)))))) (box ((box m)))))))))))))))

let Data_Set_fromMap  = (box Data_Set_Set)

let Data_Set_foldableSet  = (sharpurs_apply (box ((box Data_Foldable_Foldableusd_Dict))) (box ((box ((Map.add "foldMap" (box ((box (fun (dictMonoid: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Foldable_foldMap))) (box ((box Data_List_Types_foldableList)))))) (box ((box dictMonoid)))))) (box ((box f))))))))) (box ((box Data_Set_toList)))))))))) (Map.add "foldl" (box ((box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Foldable_foldl))) (box ((box Data_List_Types_foldableList)))))) (box ((box f)))))) (box ((box x))))))))) (box ((box Data_Set_toList)))))))))) (Map.add "foldr" (box ((box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Foldable_foldr))) (box ((box Data_List_Types_foldableList)))))) (box ((box f)))))) (box ((box x))))))))) (box ((box Data_Set_toList)))))))))) Map.empty))))))))

let Data_Set_findMin  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | m -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Maybe_functorMaybe)))))) (box ((box (fun (v1: obj) -> (Map.find "key" (unbox<Map<string, obj>> ((box v1))))))))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_findMin))) (box ((box m)))))))))))

let Data_Set_findMax  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | m -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Maybe_functorMaybe)))))) (box ((box (fun (v1: obj) -> (Map.find "key" (unbox<Map<string, obj>> ((box v1))))))))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_findMax))) (box ((box m)))))))))))

let Data_Set_filter  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_filterKeys))) (box ((box dictOrd)))))))))

let Data_Set_eqSet  = (box (fun (dictEq: obj) -> (let eqMap = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_eqMap))) (box ((box dictEq)))))) (box ((box Data_Eq_eqUnit)))) in (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (m1, m2) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box eqMap)))))) (box ((box m1)))))) (box ((box m2)))))))))))) Map.empty)))))))))

let Data_Set_ordSet  = (box (fun (dictOrd: obj) -> (let ordList = (sharpurs_apply (box ((box Data_List_Types_ordList))) (box ((box dictOrd)))) in let eqSet1 = (sharpurs_apply (box ((box Data_Set_eqSet))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (s1: obj) -> (box (fun (s2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box ordList)))))) (box ((sharpurs_apply (box ((box Data_Set_toList))) (box ((box s1))))))))) (box ((sharpurs_apply (box ((box Data_Set_toList))) (box ((box s2))))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqSet1))))) Map.empty))))))))))

let Data_Set_eq1Set  = (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box Data_Set_eqSet))) (box ((box dictEq))))))))))) Map.empty))))))

let Data_Set_ord1Set  = (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box Data_Set_ordSet))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box Data_Set_eq1Set))))) Map.empty)))))))

let Data_Set_empty  = (sharpurs_apply (box ((box Data_Set_Set))) (box ((box Data_Map_Internal_empty))))

let Data_Set_fromFoldable  = (box (fun (dictFoldable: obj) -> (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Foldable_foldl))) (box ((box dictFoldable)))))) (box ((box (fun (m: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_insert))) (box ((box dictOrd)))))) (box ((box a)))))) (box ((box m))))))))))))) (box ((box Data_Set_empty))))))))

let Data_Set_map  = (box (fun (dictOrd: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Foldable_foldl))) (box ((box Data_Set_foldableSet)))))) (box ((box (fun (m: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_insert))) (box ((box dictOrd)))))) (box ((sharpurs_apply (box ((box f))) (box ((box a))))))))) (box ((box m))))))))))))) (box ((box Data_Set_empty))))))))

let Data_Set_mapMaybe  = (box (fun (dictOrd: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Foldable_foldr))) (box ((box Data_Set_foldableSet)))))) (box ((box (fun (a: obj) -> (box (fun (acc: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Maybe_maybe))) (box ((box acc)))))) (box ((box (fun (b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_insert))) (box ((box dictOrd)))))) (box ((box b)))))) (box ((box acc))))))))))) (box ((sharpurs_apply (box ((box f))) (box ((box a)))))))))))))))) (box ((box Data_Set_empty))))))))

let Data_Set_monoidSet  = (box (fun (dictOrd: obj) -> (let semigroupSet1 = (sharpurs_apply (box ((box Data_Set_semigroupSet))) (box ((box dictOrd)))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((box Data_Set_empty))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupSet1))))) Map.empty))))))))))

let Data_Set_unions  = (box (fun (dictFoldable: obj) -> (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Foldable_foldl))) (box ((box dictFoldable)))))) (box ((sharpurs_apply (box ((box Data_Set_union))) (box ((box dictOrd))))))))) (box ((box Data_Set_empty))))))))

let Data_Set_difference  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_difference))) (box ((box dictOrd)))))))))

let Data_Set_subset  = (box (fun (dictOrd: obj) -> (box (fun (s1: obj) -> (box (fun (s2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_Set_isEmpty)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_difference))) (box ((box dictOrd)))))) (box ((box s1)))))) (box ((box s2)))))))))))))

let Data_Set_properSubset  = (box (fun (dictOrd: obj) -> (box (fun (s1: obj) -> (box (fun (s2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_HeytingAlgebra_conj))) (box ((box Data_HeytingAlgebra_heytingAlgebraBoolean)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_notEq))) (box ((box Data_Eq_eqInt)))))) (box ((sharpurs_apply (box ((box Data_Set_size))) (box ((box s1))))))))) (box ((sharpurs_apply (box ((box Data_Set_size))) (box ((box s2)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_subset))) (box ((box dictOrd)))))) (box ((box s1)))))) (box ((box s2)))))))))))))

let Data_Set_delete  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_delete))) (box ((box dictOrd)))))))))

let Data_Set_checkValid  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Map_Internal_checkValid))) (box ((box dictOrd)))))))))

let Data_Set_catMaybes  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_mapMaybe))) (box ((box dictOrd)))))) (box ((box Data_Set_identity))))))

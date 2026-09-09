[<AutoOpen>]
module PureScript_Data_Set_NonEmpty

open System
open System.Collections.Generic

let Data_Set_NonEmpty_coerce  = (sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined))))

let Data_Set_NonEmpty_NonEmptySet  = (box (fun (x: obj) -> (box x)))

let Data_Set_NonEmpty_unionSet  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((sharpurs_apply (box ((box Data_Set_semigroupSet))) (box ((box dictOrd))))))))))))

let Data_Set_NonEmpty_toUnfoldable1  = (box (fun (dictUnfoldable1: obj) -> (let stepNext = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_stepAscCps))) (box ((sharpurs_apply (box ((box Data_Function_Uncurried_mkFn3))) (box ((box (fun (k: obj) -> (box (fun (v: obj) -> (box (fun (next: obj) -> (box (Data_Maybe_Justusd_Ctor((box (Data_Tuple_Tupleusd_Ctor((box k), (box next))))))))))))))))))))) (box ((box (fun (v: obj) -> (box Data_Maybe_Nothingusd_Ctor)))))) in let stepHead = (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Map_Internal_stepAscCps))) (box ((sharpurs_apply (box ((box Data_Function_Uncurried_mkFn3))) (box ((box (fun (k: obj) -> (box (fun (v: obj) -> (box (fun (next: obj) -> (box (Data_Tuple_Tupleusd_Ctor((box k), (box next)))))))))))))))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box Partial_Unsafe_unsafeCrashWith))) (box ((box "toUnfoldable1: impossible"))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Unfoldable1_unfoldr1))) (box ((box dictUnfoldable1)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Tuple_functorTuple)))))) (box ((box stepNext)))))) (box ((box v)))))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box stepHead)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Map_Internal_toMapIter)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Set_toMap)))))) (box ((box Data_Set_NonEmpty_coerce))))))))))))))))

let Data_Set_NonEmpty_toUnfoldable11  = (sharpurs_apply (box ((box Data_Set_NonEmpty_toUnfoldable1))) (box ((box Data_List_Types_unfoldable1NonEmptyList))))

let Data_Set_NonEmpty_toUnfoldable  = (box (fun (dictUnfoldable: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_toUnfoldable))) (box ((box dictUnfoldable)))))))))

let Data_Set_NonEmpty_toSet  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((box s)))))

let Data_Set_NonEmpty_subset  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_subset))) (box ((box dictOrd)))))))))

let Data_Set_NonEmpty_size  = (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((box Data_Set_size))))

let Data_Set_NonEmpty_singleton  = (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((box Data_Set_singleton))))

let Data_Set_NonEmpty_showNonEmptySet  = (box (fun (dictShow: obj) -> (let showNonEmptyArray = (sharpurs_apply (box ((box Data_Array_NonEmpty_Internal_showNonEmptyArray))) (box ((box dictShow)))) in (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (s: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(fromFoldable1 ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box showNonEmptyArray)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_NonEmpty_toUnfoldable1))) (box ((box Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray)))))) (box ((box s)))))))))))) (box ((box ")"))))))))))) Map.empty)))))))))

let Data_Set_NonEmpty_semigroupNonEmptySet  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Set_semigroupSet))) (box ((box dictOrd))))))

let Data_Set_NonEmpty_properSubset  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_properSubset))) (box ((box dictOrd)))))))))

let Data_Set_NonEmpty_ordNonEmptySet  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Set_ordSet))) (box ((box dictOrd))))))

let Data_Set_NonEmpty_ord1NonEmptySet  = (box Data_Set_ord1Set)

let Data_Set_NonEmpty_min  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Maybe_fromJust))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_findMin))) (box ((box s))))))))))))))))

let Data_Set_NonEmpty_member  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_member))) (box ((box dictOrd)))))))))

let Data_Set_NonEmpty_max  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Maybe_fromJust))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_findMax))) (box ((box s))))))))))))))))

let Data_Set_NonEmpty_mapMaybe  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_mapMaybe))) (box ((box dictOrd)))))))))

let Data_Set_NonEmpty_map  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_map))) (box ((box dictOrd)))))))))

let Data_Set_NonEmpty_insert  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_insert))) (box ((box dictOrd)))))))))

let Data_Set_NonEmpty_fromSet  = (box (fun (s: obj) -> (match ((unbox ((sharpurs_apply (box ((box Data_Set_isEmpty))) (box ((box s))))))) with | LitBool true () -> ((box Data_Maybe_Nothingusd_Ctor)) | _ -> ((box (Data_Maybe_Justusd_Ctor((sharpurs_apply (box ((box Data_Set_NonEmpty_NonEmptySet))) (box ((box s)))))))))))

let Data_Set_NonEmpty_intersection  = (box (fun (dictOrd: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (s1, s2) -> ((sharpurs_apply (box ((box Data_Set_NonEmpty_fromSet))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_intersection))) (box ((box dictOrd)))))) (box ((box s1)))))) (box ((box s2)))))))))))))))

let Data_Set_NonEmpty_fromFoldable1  = (box (fun (dictFoldable1: obj) -> (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Foldable_foldMap1))) (box ((box dictFoldable1)))))) (box ((sharpurs_apply (box ((box Data_Set_NonEmpty_semigroupNonEmptySet))) (box ((box dictOrd))))))))) (box ((box Data_Set_NonEmpty_singleton))))))))

let Data_Set_NonEmpty_fromFoldable  = (box (fun (dictFoldable: obj) -> (let fromFoldable2 = (sharpurs_apply (box ((box Data_Set_fromFoldable))) (box ((box dictFoldable)))) in (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Set_NonEmpty_fromSet)))))) (box ((sharpurs_apply (box ((box fromFoldable2))) (box ((box dictOrd))))))))))))

let Data_Set_NonEmpty_foldableNonEmptySet  = (box Data_Set_foldableSet)

let Data_Set_NonEmpty_foldable1NonEmptySet  = (sharpurs_apply (box ((box Data_Semigroup_Foldable_Foldable1usd_Dict))) (box ((box ((Map.add "foldMap1" (box ((box (fun (dictSemigroup: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Foldable_foldMap1))) (box ((box Data_List_Types_foldable1NonEmptyList)))))) (box ((box dictSemigroup)))))) (box ((box f))))))))) (box ((box Data_Set_NonEmpty_toUnfoldable11)))))))))) (Map.add "foldr1" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Foldable_foldr1))) (box ((box Data_List_Types_foldable1NonEmptyList)))))) (box ((box f))))))))) (box ((box Data_Set_NonEmpty_toUnfoldable11)))))))) (Map.add "foldl1" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_Foldable_foldl1))) (box ((box Data_List_Types_foldable1NonEmptyList)))))) (box ((box f))))))))) (box ((box Data_Set_NonEmpty_toUnfoldable11)))))))) (Map.add "Foldable0" (box ((box (fun (usd__unused: obj) -> (box Data_Set_NonEmpty_foldableNonEmptySet))))) Map.empty)))))))))

let Data_Set_NonEmpty_filter  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_filter))) (box ((box dictOrd)))))))))

let Data_Set_NonEmpty_eqNonEmptySet  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Set_eqSet))) (box ((box dictEq))))))

let Data_Set_NonEmpty_eq1NonEmptySet  = (box Data_Set_eq1Set)

let Data_Set_NonEmpty_difference  = (box (fun (dictOrd: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (s1, s2) -> ((sharpurs_apply (box ((box Data_Set_NonEmpty_fromSet))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_difference))) (box ((box dictOrd)))))) (box ((box s1)))))) (box ((box s2)))))))))))))))

let Data_Set_NonEmpty_delete  = (box (fun (dictOrd: obj) -> (box (fun (a: obj) -> (box (fun (v: obj) -> (match (((unbox ((box a))), (unbox ((box v))))) with | (a1, s) -> ((sharpurs_apply (box ((box Data_Set_NonEmpty_fromSet))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Set_delete))) (box ((box dictOrd)))))) (box ((box a1)))))) (box ((box s)))))))))))))))

let Data_Set_NonEmpty_cons  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Safe_Coerce_coerce))) (box ((box Prim_undefined)))))) (box ((sharpurs_apply (box ((box Data_Set_insert))) (box ((box dictOrd)))))))))

[<AutoOpen>]
module PureScript_Data_List_ZipList

open System
open System.Collections.Generic

let Data_List_ZipList_ZipList  = (box (fun (x: obj) -> (box x)))

let Data_List_ZipList_traversableZipList  = (box Data_List_Lazy_Types_traversableList)

let Data_List_ZipList_showZipList  = (box (fun (dictShow: obj) -> (let showList = (sharpurs_apply (box ((box Data_List_Lazy_Types_showList))) (box ((box dictShow)))) in (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | xs -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(ZipList ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box showList)))))) (box ((box xs))))))))) (box ((box ")"))))))))))))) Map.empty)))))))))

let Data_List_ZipList_semigroupZipList  = (box Data_List_Lazy_Types_semigroupList)

let Data_List_ZipList_ordZipList  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_List_Lazy_Types_ordList))) (box ((box dictOrd))))))

let Data_List_ZipList_newtypeZipList  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_List_ZipList_monoidZipList  = (box Data_List_Lazy_Types_monoidList)

let Data_List_ZipList_functorZipList  = (box Data_List_Lazy_Types_functorList)

let Data_List_ZipList_foldableZipList  = (box Data_List_Lazy_Types_foldableList)

let Data_List_ZipList_eqZipList  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_List_Lazy_Types_eqList))) (box ((box dictEq))))))

let Data_List_ZipList_applyZipList  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (fs, xs) -> ((sharpurs_apply (box ((box Data_List_ZipList_ZipList))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_List_Lazy_zipWith))) (box ((box Data_Function_apply)))))) (box ((box fs)))))) (box ((box xs))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_List_ZipList_functorZipList))))) Map.empty)))))))

let Data_List_ZipList_zipListIsNotBind  = (box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((sharpurs_apply (box ((box Partial_Unsafe_unsafeCrashWith))) (box ((box "bind: unreachable")))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_List_ZipList_applyZipList))))) Map.empty)))))))))

let Data_List_ZipList_applicativeZipList  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_List_ZipList_ZipList)))))) (box ((box Data_List_Lazy_repeat)))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Data_List_ZipList_applyZipList))))) Map.empty)))))))

let Data_List_ZipList_altZipList  = (sharpurs_apply (box ((box Control_Alt_Altusd_Dict))) (box ((box ((Map.add "alt" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (xs, ys) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_List_ZipList_ZipList)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_List_Lazy_Types_semigroupList)))))) (box ((box xs)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_List_Lazy_drop))) (box ((sharpurs_apply (box ((box Data_List_Lazy_length))) (box ((box xs))))))))) (box ((box ys)))))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_List_ZipList_functorZipList))))) Map.empty)))))))

let Data_List_ZipList_plusZipList  = (sharpurs_apply (box ((box Control_Plus_Plususd_Dict))) (box ((box ((Map.add "empty" (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box Data_List_ZipList_monoidZipList)))))) (Map.add "Alt0" (box ((box (fun (usd__unused: obj) -> (box Data_List_ZipList_altZipList))))) Map.empty)))))))

let Data_List_ZipList_alternativeZipList  = (sharpurs_apply (box ((box Control_Alternative_Alternativeusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_List_ZipList_applicativeZipList))))) (Map.add "Plus1" (box ((box (fun (usd__unused: obj) -> (box Data_List_ZipList_plusZipList))))) Map.empty)))))))

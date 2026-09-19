[<AutoOpen>]
module PureScript_Data_Maybe_Last

open System
open System.Collections.Generic

let Data_Maybe_Last_Last  = (box (fun (x: obj) -> (box x)))

let Data_Maybe_Last_showLast  = (box (fun (dictShow: obj) -> (let showMaybe = (sharpurs_apply (box ((box Data_Maybe_showMaybe))) (box ((box dictShow)))) in (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Last ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box showMaybe)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty)))))))))

let Data_Maybe_Last_semigroupLast  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (_, (Data_Maybe_Justusd_Ctor(_) as last)) -> ((box last)) | (last, Data_Maybe_Nothingusd_Ctor) -> ((box last))))))))) Map.empty))))))

let Data_Maybe_Last_ordLast  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Maybe_ordMaybe))) (box ((box dictOrd))))))

let Data_Maybe_Last_ord1Last  = (box Data_Maybe_ord1Maybe)

let Data_Maybe_Last_newtypeLast  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Maybe_Last_monoidLast  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Maybe_Last_Last))) (box ((box Data_Maybe_Nothingusd_Ctor)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_Last_semigroupLast))))) Map.empty)))))))

let Data_Maybe_Last_monadLast  = (box Data_Maybe_monadMaybe)

let Data_Maybe_Last_invariantLast  = (box Data_Maybe_invariantMaybe)

let Data_Maybe_Last_functorLast  = (box Data_Maybe_functorMaybe)

let Data_Maybe_Last_extendLast  = (box Data_Maybe_extendMaybe)

let Data_Maybe_Last_eqLast  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Maybe_eqMaybe))) (box ((box dictEq))))))

let Data_Maybe_Last_eq1Last  = (box Data_Maybe_eq1Maybe)

let Data_Maybe_Last_boundedLast  = (box (fun (dictBounded: obj) -> (sharpurs_apply (box ((box Data_Maybe_boundedMaybe))) (box ((box dictBounded))))))

let Data_Maybe_Last_bindLast  = (box Data_Maybe_bindMaybe)

let Data_Maybe_Last_applyLast  = (box Data_Maybe_applyMaybe)

let Data_Maybe_Last_applicativeLast  = (box Data_Maybe_applicativeMaybe)

let Data_Maybe_Last_altLast  = (sharpurs_apply (box ((box Control_Alt_Altusd_Dict))) (box ((box ((Map.add "alt" (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Maybe_Last_semigroupLast)))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_Last_functorLast))))) Map.empty)))))))

let Data_Maybe_Last_plusLast  = (sharpurs_apply (box ((box Control_Plus_Plususd_Dict))) (box ((box ((Map.add "empty" (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box Data_Maybe_Last_monoidLast)))))) (Map.add "Alt0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_Last_altLast))))) Map.empty)))))))

let Data_Maybe_Last_alternativeLast  = (sharpurs_apply (box ((box Control_Alternative_Alternativeusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_Last_applicativeLast))))) (Map.add "Plus1" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_Last_plusLast))))) Map.empty)))))))

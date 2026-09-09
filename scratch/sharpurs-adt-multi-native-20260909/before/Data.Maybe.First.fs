[<AutoOpen>]
module PureScript_Data_Maybe_First

open System
open System.Collections.Generic

let Data_Maybe_First_First  = (box (fun (x: obj) -> (box x)))

let Data_Maybe_First_showFirst  = (box (fun (dictShow: obj) -> (let showMaybe = (sharpurs_apply (box ((box Data_Maybe_showMaybe))) (box ((box dictShow)))) in (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "First (")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box showMaybe)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty)))))))))

let Data_Maybe_First_semigroupFirst  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | ((Data_Maybe_Justusd_Ctor(_) as first), _) -> ((box first)) | (_, second) -> ((box second))))))))) Map.empty))))))

let Data_Maybe_First_ordFirst  = (box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Maybe_ordMaybe))) (box ((box dictOrd))))))

let Data_Maybe_First_ord1First  = (box Data_Maybe_ord1Maybe)

let Data_Maybe_First_newtypeFirst  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Maybe_First_monoidFirst  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Maybe_First_First))) (box ((box Data_Maybe_Nothingusd_Ctor)))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_First_semigroupFirst))))) Map.empty)))))))

let Data_Maybe_First_monadFirst  = (box Data_Maybe_monadMaybe)

let Data_Maybe_First_invariantFirst  = (box Data_Maybe_invariantMaybe)

let Data_Maybe_First_functorFirst  = (box Data_Maybe_functorMaybe)

let Data_Maybe_First_extendFirst  = (box Data_Maybe_extendMaybe)

let Data_Maybe_First_eqFirst  = (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Maybe_eqMaybe))) (box ((box dictEq))))))

let Data_Maybe_First_eq1First  = (box Data_Maybe_eq1Maybe)

let Data_Maybe_First_boundedFirst  = (box (fun (dictBounded: obj) -> (sharpurs_apply (box ((box Data_Maybe_boundedMaybe))) (box ((box dictBounded))))))

let Data_Maybe_First_bindFirst  = (box Data_Maybe_bindMaybe)

let Data_Maybe_First_applyFirst  = (box Data_Maybe_applyMaybe)

let Data_Maybe_First_applicativeFirst  = (box Data_Maybe_applicativeMaybe)

let Data_Maybe_First_altFirst  = (sharpurs_apply (box ((box Control_Alt_Altusd_Dict))) (box ((box ((Map.add "alt" (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Maybe_First_semigroupFirst)))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_First_functorFirst))))) Map.empty)))))))

let Data_Maybe_First_plusFirst  = (sharpurs_apply (box ((box Control_Plus_Plususd_Dict))) (box ((box ((Map.add "empty" (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box Data_Maybe_First_monoidFirst)))))) (Map.add "Alt0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_First_altFirst))))) Map.empty)))))))

let Data_Maybe_First_alternativeFirst  = (sharpurs_apply (box ((box Control_Alternative_Alternativeusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_First_applicativeFirst))))) (Map.add "Plus1" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_First_plusFirst))))) Map.empty)))))))

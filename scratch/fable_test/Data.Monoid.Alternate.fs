[<AutoOpen>]
module PureScript_Data_Monoid_Alternate

open System
open System.Collections.Generic

let Data_Monoid_Alternate_Alternate  = (box (fun (x: obj) -> (box x)))

let Data_Monoid_Alternate_showAlternate  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | a -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Alternate ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Monoid_Alternate_semigroupAlternate  = (box (fun (dictAlt: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a, b) -> ((sharpurs_apply (box ((box Data_Monoid_Alternate_Alternate))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Alt_alt))) (box ((box dictAlt)))))) (box ((box a)))))) (box ((box b))))))))))))))) Map.empty))))))))

let Data_Monoid_Alternate_plusAlternate  = (box (fun (dictPlus: obj) -> (box dictPlus)))

let Data_Monoid_Alternate_ordAlternate  = (box (fun (dictOrd: obj) -> (box dictOrd)))

let Data_Monoid_Alternate_ord1Alternate  = (box (fun (dictOrd1: obj) -> (box dictOrd1)))

let Data_Monoid_Alternate_newtypeAlternate  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Monoid_Alternate_monoidAlternate  = (box (fun (dictPlus: obj) -> (let semigroupAlternate1 = (sharpurs_apply (box ((box Data_Monoid_Alternate_semigroupAlternate))) (box ((sharpurs_apply (box ((Map.find "Alt0" (unbox<Map<string, obj>> ((box dictPlus)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Monoid_Alternate_Alternate))) (box ((sharpurs_apply (box ((box Control_Plus_empty))) (box ((box dictPlus))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupAlternate1))))) Map.empty))))))))))

let Data_Monoid_Alternate_monadAlternate  = (box (fun (dictMonad: obj) -> (box dictMonad)))

let Data_Monoid_Alternate_functorAlternate  = (box (fun (dictFunctor: obj) -> (box dictFunctor)))

let Data_Monoid_Alternate_extendAlternate  = (box (fun (dictExtend: obj) -> (box dictExtend)))

let Data_Monoid_Alternate_eqAlternate  = (box (fun (dictEq: obj) -> (box dictEq)))

let Data_Monoid_Alternate_eq1Alternate  = (box (fun (dictEq1: obj) -> (box dictEq1)))

let Data_Monoid_Alternate_comonadAlternate  = (box (fun (dictComonad: obj) -> (box dictComonad)))

let Data_Monoid_Alternate_boundedAlternate  = (box (fun (dictBounded: obj) -> (box dictBounded)))

let Data_Monoid_Alternate_bindAlternate  = (box (fun (dictBind: obj) -> (box dictBind)))

let Data_Monoid_Alternate_applyAlternate  = (box (fun (dictApply: obj) -> (box dictApply)))

let Data_Monoid_Alternate_applicativeAlternate  = (box (fun (dictApplicative: obj) -> (box dictApplicative)))

let Data_Monoid_Alternate_alternativeAlternate  = (box (fun (dictAlternative: obj) -> (box dictAlternative)))

let Data_Monoid_Alternate_altAlternate  = (box (fun (dictAlt: obj) -> (box dictAlt)))

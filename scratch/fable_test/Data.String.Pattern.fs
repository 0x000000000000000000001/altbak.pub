[<AutoOpen>]
module PureScript_Data_String_Pattern

open System
open System.Collections.Generic

let Data_String_Pattern_Replacement  = (box (fun (x: obj) -> (box x)))

let Data_String_Pattern_Pattern  = (box (fun (x: obj) -> (box x)))

let Data_String_Pattern_showReplacement  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Replacement ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showString)))))) (box ((box s))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_String_Pattern_showPattern  = (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | s -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(Pattern ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showString)))))) (box ((box s))))))))) (box ((box ")"))))))))))))) Map.empty))))))

let Data_String_Pattern_newtypeReplacement  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_String_Pattern_newtypePattern  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_String_Pattern_eqReplacement  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqString)))))) (box ((box l)))))) (box ((box r)))))))))))) Map.empty))))))

let Data_String_Pattern_ordReplacement  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box Data_Ord_ordString)))))) (box ((box l)))))) (box ((box r)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_String_Pattern_eqReplacement))))) Map.empty)))))))

let Data_String_Pattern_eqPattern  = (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq))) (box ((box Data_Eq_eqString)))))) (box ((box l)))))) (box ((box r)))))))))))) Map.empty))))))

let Data_String_Pattern_ordPattern  = (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare))) (box ((box Data_Ord_ordString)))))) (box ((box l)))))) (box ((box r)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box Data_String_Pattern_eqPattern))))) Map.empty)))))))

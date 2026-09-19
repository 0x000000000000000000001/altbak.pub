[<AutoOpen>]
module PureScript_Control_Monad_Identity_Trans

open System
open System.Collections.Generic

let Control_Monad_Identity_Trans_IdentityT  = (box (fun (x: obj) -> (box x)))

let Control_Monad_Identity_Trans_monadSTIdentityT  = (box (fun (dictMonadST: obj) -> (box dictMonadST)))

let Control_Monad_Identity_Trans_traversableIdentityT  = (box (fun (dictTraversable: obj) -> (box dictTraversable)))

let Control_Monad_Identity_Trans_runIdentityT  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | ma -> ((box ma)))))

let Control_Monad_Identity_Trans_plusIdentityT  = (box (fun (dictPlus: obj) -> (box dictPlus)))

let Control_Monad_Identity_Trans_newtypeIdentityT  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Control_Monad_Identity_Trans_monadWriterIdentityT  = (box (fun (dictMonadWriter: obj) -> (box dictMonadWriter)))

let Control_Monad_Identity_Trans_monadTransIdentityT  = (sharpurs_apply (box ((box Control_Monad_Trans_Class_MonadTransusd_Dict))) (box ((box ((Map.add "lift" (box ((box (fun (dictMonad: obj) -> (box Control_Monad_Identity_Trans_IdentityT))))) Map.empty))))))

let Control_Monad_Identity_Trans_monadThrowIdentityT  = (box (fun (dictMonadThrow: obj) -> (box dictMonadThrow)))

let Control_Monad_Identity_Trans_monadTellIdentityT  = (box (fun (dictMonadTell: obj) -> (box dictMonadTell)))

let Control_Monad_Identity_Trans_monadStateIdentityT  = (box (fun (dictMonadState: obj) -> (box dictMonadState)))

let Control_Monad_Identity_Trans_monadRecIdentityT  = (box (fun (dictMonadRec: obj) -> (box dictMonadRec)))

let Control_Monad_Identity_Trans_monadReaderIdentityT  = (box (fun (dictMonadReader: obj) -> (box dictMonadReader)))

let Control_Monad_Identity_Trans_monadPlusIdentityT  = (box (fun (dictMonadPlus: obj) -> (box dictMonadPlus)))

let Control_Monad_Identity_Trans_monadIdentityT  = (box (fun (dictMonad: obj) -> (box dictMonad)))

let Control_Monad_Identity_Trans_monadErrorIdentityT  = (box (fun (dictMonadError: obj) -> (box dictMonadError)))

let Control_Monad_Identity_Trans_monadEffectIdentityT  = (box (fun (dictMonadEffect: obj) -> (box dictMonadEffect)))

let Control_Monad_Identity_Trans_monadContIdentityT  = (box (fun (dictMonadCont: obj) -> (box dictMonadCont)))

let Control_Monad_Identity_Trans_monadAskIdentityT  = (box (fun (dictMonadAsk: obj) -> (box dictMonadAsk)))

let Control_Monad_Identity_Trans_mapIdentityT  = (box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, m) -> ((sharpurs_apply (box ((box Control_Monad_Identity_Trans_IdentityT))) (box ((sharpurs_apply (box ((box f1))) (box ((box m)))))))))))))

let Control_Monad_Identity_Trans_functorIdentityT  = (box (fun (dictFunctor: obj) -> (box dictFunctor)))

let Control_Monad_Identity_Trans_foldableIdentityT  = (box (fun (dictFoldable: obj) -> (box dictFoldable)))

let Control_Monad_Identity_Trans_extendIdentityI  = (box (fun (dictExtend: obj) -> (let functorIdentityT1 = (sharpurs_apply (box ((box Control_Monad_Identity_Trans_functorIdentityT))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictExtend)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Control_Extend_Extendusd_Dict))) (box ((box ((Map.add "extend" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, m) -> ((sharpurs_apply (box ((box Control_Monad_Identity_Trans_IdentityT))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Extend_extend))) (box ((box dictExtend)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f1)))))) (box ((box Control_Monad_Identity_Trans_IdentityT))))))))) (box ((box m))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box functorIdentityT1))))) Map.empty))))))))))

let Control_Monad_Identity_Trans_eqIdentityT  = (box (fun (dictEq1: obj) -> (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq1))) (box ((box dictEq1)))))) (box ((box dictEq)))))) (box ((box l)))))) (box ((box r)))))))))))) Map.empty))))))))))

let Control_Monad_Identity_Trans_ordIdentityT  = (box (fun (dictOrd1: obj) -> (let eqIdentityT1 = (sharpurs_apply (box ((box Control_Monad_Identity_Trans_eqIdentityT))) (box ((sharpurs_apply (box ((Map.find "Eq10" (unbox<Map<string, obj>> ((box dictOrd1)))))) (box ((box Prim_undefined))))))) in (box (fun (dictOrd: obj) -> (let eqIdentityT2 = (sharpurs_apply (box ((box eqIdentityT1))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare1))) (box ((box dictOrd1)))))) (box ((box dictOrd)))))) (box ((box l)))))) (box ((box r)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqIdentityT2))))) Map.empty)))))))))))))

let Control_Monad_Identity_Trans_eq1IdentityT  = (box (fun (dictEq1: obj) -> (let eqIdentityT1 = (sharpurs_apply (box ((box Control_Monad_Identity_Trans_eqIdentityT))) (box ((box dictEq1)))) in (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box eqIdentityT1))) (box ((box dictEq))))))))))) Map.empty)))))))))

let Control_Monad_Identity_Trans_ord1IdentityT  = (box (fun (dictOrd1: obj) -> (let ordIdentityT1 = (sharpurs_apply (box ((box Control_Monad_Identity_Trans_ordIdentityT))) (box ((box dictOrd1)))) in let eq1IdentityT1 = (sharpurs_apply (box ((box Control_Monad_Identity_Trans_eq1IdentityT))) (box ((sharpurs_apply (box ((Map.find "Eq10" (unbox<Map<string, obj>> ((box dictOrd1)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box ordIdentityT1))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box eq1IdentityT1))))) Map.empty))))))))))

let Control_Monad_Identity_Trans_comonadIdentityT  = (box (fun (dictComonad: obj) -> (let extendIdentityI1 = (sharpurs_apply (box ((box Control_Monad_Identity_Trans_extendIdentityI))) (box ((sharpurs_apply (box ((Map.find "Extend0" (unbox<Map<string, obj>> ((box dictComonad)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Control_Comonad_Comonadusd_Dict))) (box ((box ((Map.add "extract" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Control_Comonad_extract))) (box ((box dictComonad))))))))) (box ((box Control_Monad_Identity_Trans_runIdentityT)))))) (Map.add "Extend0" (box ((box (fun (usd__unused: obj) -> (box extendIdentityI1))))) Map.empty))))))))))

let Control_Monad_Identity_Trans_bindIdentityT  = (box (fun (dictBind: obj) -> (box dictBind)))

let Control_Monad_Identity_Trans_applyIdentityT  = (box (fun (dictApply: obj) -> (box dictApply)))

let Control_Monad_Identity_Trans_applicativeIdentityT  = (box (fun (dictApplicative: obj) -> (box dictApplicative)))

let Control_Monad_Identity_Trans_alternativeIdentityT  = (box (fun (dictAlternative: obj) -> (box dictAlternative)))

let Control_Monad_Identity_Trans_altIdentityT  = (box (fun (dictAlt: obj) -> (box dictAlt)))

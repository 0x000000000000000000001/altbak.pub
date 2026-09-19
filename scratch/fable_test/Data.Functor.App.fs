[<AutoOpen>]
module PureScript_Data_Functor_App

open System
open System.Collections.Generic

let Data_Functor_App_App  = (box (fun (x: obj) -> (box x)))

let Data_Functor_App_showApp  = (box (fun (dictShow: obj) -> (sharpurs_apply (box ((box Data_Show_Showusd_Dict))) (box ((box ((Map.add "show" (box ((box (fun (v: obj) -> (match ((unbox ((box v)))) with | fa -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((box "(App ")))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box Data_Semigroup_semigroupString)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box fa))))))))) (box ((box ")"))))))))))))) Map.empty))))))))

let Data_Functor_App_semigroupApp  = (box (fun (dictApply: obj) -> (box (fun (dictSemigroup: obj) -> (let append = (sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup)))) in (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (fa1, fa2) -> ((sharpurs_apply (box ((box Data_Functor_App_App))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_lift2))) (box ((box dictApply)))))) (box ((box append)))))) (box ((box fa1)))))) (box ((box fa2))))))))))))))) Map.empty)))))))))))

let Data_Functor_App_plusApp  = (box (fun (dictPlus: obj) -> (box dictPlus)))

let Data_Functor_App_newtypeApp  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Data_Functor_App_monoidApp  = (box (fun (dictApplicative: obj) -> (let semigroupApp1 = (sharpurs_apply (box ((box Data_Functor_App_semigroupApp))) (box ((sharpurs_apply (box ((Map.find "Apply0" (unbox<Map<string, obj>> ((box dictApplicative)))))) (box ((box Prim_undefined))))))) in (box (fun (dictMonoid: obj) -> (let semigroupApp2 = (sharpurs_apply (box ((box semigroupApp1))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((box Data_Functor_App_App))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box dictApplicative)))))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid)))))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupApp2))))) Map.empty)))))))))))))

let Data_Functor_App_monadPlusApp  = (box (fun (dictMonadPlus: obj) -> (box dictMonadPlus)))

let Data_Functor_App_monadApp  = (box (fun (dictMonad: obj) -> (box dictMonad)))

let Data_Functor_App_lazyApp  = (box (fun (dictLazy: obj) -> (box dictLazy)))

let Data_Functor_App_hoistLowerApp  = (box Unsafe_Coerce_unsafeCoerce)

let Data_Functor_App_hoistLiftApp  = (box Unsafe_Coerce_unsafeCoerce)

let Data_Functor_App_hoistApp  = (box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, fa) -> ((sharpurs_apply (box ((box Data_Functor_App_App))) (box ((sharpurs_apply (box ((box f1))) (box ((box fa)))))))))))))

let Data_Functor_App_functorApp  = (box (fun (dictFunctor: obj) -> (box dictFunctor)))

let Data_Functor_App_extendApp  = (box (fun (dictExtend: obj) -> (box dictExtend)))

let Data_Functor_App_eqApp  = (box (fun (dictEq1: obj) -> (box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_Equsd_Dict))) (box ((box ((Map.add "eq" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Eq_eq1))) (box ((box dictEq1)))))) (box ((box dictEq)))))) (box ((box l)))))) (box ((box r)))))))))))) Map.empty))))))))))

let Data_Functor_App_ordApp  = (box (fun (dictOrd1: obj) -> (let eqApp1 = (sharpurs_apply (box ((box Data_Functor_App_eqApp))) (box ((sharpurs_apply (box ((Map.find "Eq10" (unbox<Map<string, obj>> ((box dictOrd1)))))) (box ((box Prim_undefined))))))) in (box (fun (dictOrd: obj) -> (let eqApp2 = (sharpurs_apply (box ((box eqApp1))) (box ((sharpurs_apply (box ((Map.find "Eq0" (unbox<Map<string, obj>> ((box dictOrd)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ordusd_Dict))) (box ((box ((Map.add "compare" (box ((box (fun (x: obj) -> (box (fun (y: obj) -> (match (((unbox ((box x))), (unbox ((box y))))) with | (l, r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Ord_compare1))) (box ((box dictOrd1)))))) (box ((box dictOrd)))))) (box ((box l)))))) (box ((box r)))))))))))) (Map.add "Eq0" (box ((box (fun (usd__unused: obj) -> (box eqApp2))))) Map.empty)))))))))))))

let Data_Functor_App_eq1App  = (box (fun (dictEq1: obj) -> (let eqApp1 = (sharpurs_apply (box ((box Data_Functor_App_eqApp))) (box ((box dictEq1)))) in (sharpurs_apply (box ((box Data_Eq_Eq1usd_Dict))) (box ((box ((Map.add "eq1" (box ((box (fun (dictEq: obj) -> (sharpurs_apply (box ((box Data_Eq_eq))) (box ((sharpurs_apply (box ((box eqApp1))) (box ((box dictEq))))))))))) Map.empty)))))))))

let Data_Functor_App_ord1App  = (box (fun (dictOrd1: obj) -> (let ordApp1 = (sharpurs_apply (box ((box Data_Functor_App_ordApp))) (box ((box dictOrd1)))) in let eq1App1 = (sharpurs_apply (box ((box Data_Functor_App_eq1App))) (box ((sharpurs_apply (box ((Map.find "Eq10" (unbox<Map<string, obj>> ((box dictOrd1)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Ord_Ord1usd_Dict))) (box ((box ((Map.add "compare1" (box ((box (fun (dictOrd: obj) -> (sharpurs_apply (box ((box Data_Ord_compare))) (box ((sharpurs_apply (box ((box ordApp1))) (box ((box dictOrd))))))))))) (Map.add "Eq10" (box ((box (fun (usd__unused: obj) -> (box eq1App1))))) Map.empty))))))))))

let Data_Functor_App_comonadApp  = (box (fun (dictComonad: obj) -> (box dictComonad)))

let Data_Functor_App_bindApp  = (box (fun (dictBind: obj) -> (box dictBind)))

let Data_Functor_App_applyApp  = (box (fun (dictApply: obj) -> (box dictApply)))

let Data_Functor_App_applicativeApp  = (box (fun (dictApplicative: obj) -> (box dictApplicative)))

let Data_Functor_App_alternativeApp  = (box (fun (dictAlternative: obj) -> (box dictAlternative)))

let Data_Functor_App_altApp  = (box (fun (dictAlt: obj) -> (box dictAlt)))

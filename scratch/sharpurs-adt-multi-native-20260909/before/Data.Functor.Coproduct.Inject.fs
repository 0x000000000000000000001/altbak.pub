[<AutoOpen>]
module PureScript_Data_Functor_Coproduct_Inject

open System
open System.Collections.Generic

let Data_Functor_Coproduct_Inject_Injectusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Functor_Coproduct_Inject_prj  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "prj" (unbox<Map<string, obj>> ((box v))))))))

let Data_Functor_Coproduct_Inject_injectReflexive  = (sharpurs_apply (box ((box Data_Functor_Coproduct_Inject_Injectusd_Dict))) (box ((box ((Map.add "inj" (box ((sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn)))))) (Map.add "prj" (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1)))))))) Map.empty)))))))

let Data_Functor_Coproduct_Inject_injectLeft  = (sharpurs_apply (box ((box Data_Functor_Coproduct_Inject_Injectusd_Dict))) (box ((box ((Map.add "inj" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Functor_Coproduct_Coproduct)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1))))))))))) (Map.add "prj" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_Coproduct_coproduct))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Maybe_Justusd_Ctor(usd__arg1))))))))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box Data_Maybe_Nothingusd_Ctor))))))))) Map.empty)))))))

let Data_Functor_Coproduct_Inject_inj  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "inj" (unbox<Map<string, obj>> ((box v))))))))

let Data_Functor_Coproduct_Inject_injectRight  = (box (fun (dictInject: obj) -> (sharpurs_apply (box ((box Data_Functor_Coproduct_Inject_Injectusd_Dict))) (box ((box ((Map.add "inj" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Functor_Coproduct_Coproduct)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1))))))))))) (box ((sharpurs_apply (box ((box Data_Functor_Coproduct_Inject_inj))) (box ((box dictInject)))))))))))) (Map.add "prj" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_Coproduct_coproduct))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box Data_Maybe_Nothingusd_Ctor))))))))) (box ((sharpurs_apply (box ((box Data_Functor_Coproduct_Inject_prj))) (box ((box dictInject))))))))) Map.empty)))))))))

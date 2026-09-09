[<AutoOpen>]
module PureScript_Data_Profunctor_Choice

open System
open System.Collections.Generic

let Data_Profunctor_Choice_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Profunctor_Choice_Choiceusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_Choice_right  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "right" (unbox<Map<string, obj>> ((box v))))))))

let Data_Profunctor_Choice_left  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "left" (unbox<Map<string, obj>> ((box v))))))))

let Data_Profunctor_Choice_splitChoice  = (box (fun (dictSemigroupoid: obj) -> (box (fun (dictChoice: obj) -> (box (fun (l: obj) -> (box (fun (r: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box dictSemigroupoid)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Choice_left))) (box ((box dictChoice)))))) (box ((box l))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Choice_right))) (box ((box dictChoice)))))) (box ((box r)))))))))))))))

let Data_Profunctor_Choice_fanin  = (box (fun (dictSemigroupoid: obj) -> (box (fun (dictChoice: obj) -> (let Profunctor0 = (sharpurs_apply (box ((Map.find "Profunctor0" (unbox<Map<string, obj>> ((box dictChoice)))))) (box ((box Prim_undefined)))) in (box (fun (l: obj) -> (box (fun (r: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_rmap))) (box ((box Profunctor0)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Either_either))) (box ((box Data_Profunctor_Choice_identity)))))) (box ((box Data_Profunctor_Choice_identity))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Choice_splitChoice))) (box ((box dictSemigroupoid)))))) (box ((box dictChoice)))))) (box ((box l)))))) (box ((box r))))))))))))))))

let Data_Profunctor_Choice_choiceFn  = (sharpurs_apply (box ((box Data_Profunctor_Choice_Choiceusd_Dict))) (box ((box ((Map.add "left" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (a2b, Data_Either_Leftusd_Ctor(a)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1))))))))))) (box ((sharpurs_apply (box ((box a2b))) (box ((box a)))))))) | (_, Data_Either_Rightusd_Ctor(c)) -> ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1)))))))) (box ((box c)))))))))))) (Map.add "right" (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Either_functorEither)))))) (Map.add "Profunctor0" (box ((box (fun (usd__unused: obj) -> (box Data_Profunctor_profunctorFn))))) Map.empty))))))))

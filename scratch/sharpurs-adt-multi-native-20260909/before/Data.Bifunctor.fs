[<AutoOpen>]
module PureScript_Data_Bifunctor

open System
open System.Collections.Generic

let Data_Bifunctor_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Bifunctor_identity1  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Bifunctor_Bifunctorusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Bifunctor_bimap  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "bimap" (unbox<Map<string, obj>> ((box v))))))))

let Data_Bifunctor_bivoid  = (box (fun (dictBifunctor: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bifunctor_bimap))) (box ((box dictBifunctor)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box Data_Unit_unit))))))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box Data_Unit_unit)))))))))

let Data_Bifunctor_lmap  = (box (fun (dictBifunctor: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bifunctor_bimap))) (box ((box dictBifunctor)))))) (box ((box f)))))) (box ((box Data_Bifunctor_identity))))))))

let Data_Bifunctor_rmap  = (box (fun (dictBifunctor: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bifunctor_bimap))) (box ((box dictBifunctor)))))) (box ((box Data_Bifunctor_identity1))))))

let Data_Bifunctor_bifunctorTuple  = (sharpurs_apply (box ((box Data_Bifunctor_Bifunctorusd_Dict))) (box ((box ((Map.add "bimap" (box ((box (fun (f: obj) -> (box (fun (g: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box g))), (unbox ((box v))))) with | (f1, g1, Data_Tuple_Tupleusd_Ctor(x, y)) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((box f1))) (box ((box x))))))))) (box ((sharpurs_apply (box ((box g1))) (box ((box y))))))))))))))))) Map.empty))))))

let Data_Bifunctor_bifunctorEither  = (sharpurs_apply (box ((box Data_Bifunctor_Bifunctorusd_Dict))) (box ((box ((Map.add "bimap" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (match (((unbox ((box v))), (unbox ((box v1))), (unbox ((box v2))))) with | (f, _, Data_Either_Leftusd_Ctor(l)) -> ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1)))))))) (box ((sharpurs_apply (box ((box f))) (box ((box l)))))))) | (_, g, Data_Either_Rightusd_Ctor(r)) -> ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1)))))))) (box ((sharpurs_apply (box ((box g))) (box ((box r))))))))))))))))) Map.empty))))))

let Data_Bifunctor_bifunctorConst  = (sharpurs_apply (box ((box Data_Bifunctor_Bifunctorusd_Dict))) (box ((box ((Map.add "bimap" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box f))), (unbox ((box v))), (unbox ((box v1))))) with | (f1, _, a) -> ((sharpurs_apply (box ((box Data_Const_Const))) (box ((sharpurs_apply (box ((box f1))) (box ((box a))))))))))))))))) Map.empty))))))

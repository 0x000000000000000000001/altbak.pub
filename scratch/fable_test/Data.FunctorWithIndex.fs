[<AutoOpen>]
module PureScript_Data_FunctorWithIndex

open System
open System.Collections.Generic

module Data_FunctorWithIndex_FFI =
    let mapWithIndexArray = 
        fun (f: obj) -> fun (xs: obj) ->
            let arr = xs :?> obj[]
            let res = Array.zeroCreate arr.Length
            for i = 0 to arr.Length - 1 do
                let f1 = sharpurs_apply f (box i)
                res.[i] <- sharpurs_apply f1 arr.[i]
            res :> obj
    

let Data_FunctorWithIndex_mapWithIndexArray = box (Data_FunctorWithIndex_FFI.``mapWithIndexArray``)


let Data_FunctorWithIndex_map  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Tuple_functorTuple))))

let Data_FunctorWithIndex_map1  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Monoid_Multiplicative_functorMultiplicative))))

let Data_FunctorWithIndex_map2  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Maybe_functorMaybe))))

let Data_FunctorWithIndex_map3  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Maybe_Last_functorLast))))

let Data_FunctorWithIndex_map4  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Maybe_First_functorFirst))))

let Data_FunctorWithIndex_map5  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Either_functorEither))))

let Data_FunctorWithIndex_map6  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Monoid_Dual_functorDual))))

let Data_FunctorWithIndex_map7  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Monoid_Disj_functorDisj))))

let Data_FunctorWithIndex_map8  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Monoid_Conj_functorConj))))

let Data_FunctorWithIndex_map9  = (sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Monoid_Additive_functorAdditive))))

let Data_FunctorWithIndex_FunctorWithIndexusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_FunctorWithIndex_mapWithIndex  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "mapWithIndex" (unbox<Map<string, obj>> ((box v))))))))

let Data_FunctorWithIndex_mapDefault  = (box (fun (dictFunctorWithIndex: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box f)))))))))))

let Data_FunctorWithIndex_functorWithIndexTuple  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Tuple_functorTuple))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexProduct  = (box (fun (dictFunctorWithIndex: obj) -> (let functorProduct = (sharpurs_apply (box ((box Data_Functor_Product_functorProduct))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictFunctorWithIndex)))))) (box ((box Prim_undefined))))))) in (box (fun (dictFunctorWithIndex1: obj) -> (let functorProduct1 = (sharpurs_apply (box ((box functorProduct))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictFunctorWithIndex1)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, fga) -> ((sharpurs_apply (box ((box Data_Functor_Product_Product))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bifunctor_bimap))) (box ((box Data_Bifunctor_bifunctorTuple)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f1)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1))))))))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f1)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1))))))))))))))))) (box ((box fga))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box functorProduct1))))) Map.empty)))))))))))))

let Data_FunctorWithIndex_functorWithIndexMultiplicative  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map1)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Multiplicative_functorMultiplicative))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexMaybe  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map2)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_functorMaybe))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexLast  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map3)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_Last_functorLast))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexIdentity  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, a) -> ((sharpurs_apply (box ((box Data_Identity_Identity))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box f1))) (box ((box Data_Unit_unit)))))) (box ((box a))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_functorIdentity))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexFirst  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map4)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Maybe_First_functorFirst))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexEither  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map5)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Either_functorEither))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexDual  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map6)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Dual_functorDual))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexDisj  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map7)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Disj_functorDisj))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexCoproduct  = (box (fun (dictFunctorWithIndex: obj) -> (let functorCoproduct = (sharpurs_apply (box ((box Data_Functor_Coproduct_functorCoproduct))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictFunctorWithIndex)))))) (box ((box Prim_undefined))))))) in (box (fun (dictFunctorWithIndex1: obj) -> (let functorCoproduct1 = (sharpurs_apply (box ((box functorCoproduct))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictFunctorWithIndex1)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, e) -> ((sharpurs_apply (box ((box Data_Functor_Coproduct_Coproduct))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Bifunctor_bimap))) (box ((box Data_Bifunctor_bifunctorEither)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f1)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1))))))))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box f1)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1))))))))))))))))) (box ((box e))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box functorCoproduct1))))) Map.empty)))))))))))))

let Data_FunctorWithIndex_functorWithIndexConst  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (_, x) -> ((sharpurs_apply (box ((box Data_Const_Const))) (box ((box x)))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Const_functorConst))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexConj  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map8)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Conj_functorConj))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexCompose  = (box (fun (dictFunctorWithIndex: obj) -> (let functorCompose = (sharpurs_apply (box ((box Data_Functor_Compose_functorCompose))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictFunctorWithIndex)))))) (box ((box Prim_undefined))))))) in (box (fun (dictFunctorWithIndex1: obj) -> (let mapWithIndex1 = (sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex1)))) in let functorCompose1 = (sharpurs_apply (box ((box functorCompose))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictFunctorWithIndex1)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, fga) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_Functor_Compose_Compose)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box mapWithIndex1)))))) (box ((sharpurs_apply (box ((box Data_Tuple_curry))) (box ((box f1)))))))))))) (box ((box fga))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box functorCompose1))))) Map.empty)))))))))))))

let Data_FunctorWithIndex_functorWithIndexArray  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box Data_FunctorWithIndex_mapWithIndexArray))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorArray))))) Map.empty)))))))

let Data_FunctorWithIndex_functorWithIndexApp  = (box (fun (dictFunctorWithIndex: obj) -> (let functorApp = (sharpurs_apply (box ((box Data_Functor_App_functorApp))) (box ((sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictFunctorWithIndex)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (box (fun (v: obj) -> (match (((unbox ((box f))), (unbox ((box v))))) with | (f1, x) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_Functor_App_App)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_FunctorWithIndex_mapWithIndex))) (box ((box dictFunctorWithIndex)))))) (box ((box f1)))))) (box ((box x))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box functorApp))))) Map.empty))))))))))

let Data_FunctorWithIndex_functorWithIndexAdditive  = (sharpurs_apply (box ((box Data_FunctorWithIndex_FunctorWithIndexusd_Dict))) (box ((box ((Map.add "mapWithIndex" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Data_FunctorWithIndex_map9)))))) (box ((sharpurs_apply (box ((box f))) (box ((box Data_Unit_unit))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Monoid_Additive_functorAdditive))))) Map.empty)))))))

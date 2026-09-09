[<AutoOpen>]
module PureScript_Data_Functor

open System
open System.Collections.Generic

module Data_Functor_FFI =
    let arrayMap = box (fun (f: obj) -> box (fun (arr: obj) ->
        let a = unbox<obj[]> arr
        let l = a.Length
        let result = Array.zeroCreate<obj> l
        for i = 0 to l - 1 do
            result.[i] <- sharpurs_apply f a.[i]
        box result
    ))
    

let Data_Functor_arrayMap = box (Data_Functor_FFI.``arrayMap``)


let Data_Functor_Functorusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Functor_map  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "map" (unbox<Map<string, obj>> ((box v))))))))

let Data_Functor_mapFlipped  = (box (fun (dictFunctor: obj) -> (box (fun (fa: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box f)))))) (box ((box fa))))))))))

let Data_Functor_void  = (box (fun (dictFunctor: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box Data_Unit_unit)))))))))

let Data_Functor_voidLeft  = (box (fun (dictFunctor: obj) -> (box (fun (f: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box x))))))))) (box ((box f))))))))))

let Data_Functor_voidRight  = (box (fun (dictFunctor: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box x)))))))))))

let Data_Functor_functorProxy  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (box Type_Proxy_Proxyusd_Ctor))))))) Map.empty))))))

let Data_Functor_functorFn  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) Map.empty))))))

let Data_Functor_functorArray  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box Data_Functor_arrayMap))) Map.empty))))))

let Data_Functor_flap  = (box (fun (dictFunctor: obj) -> (box (fun (ff: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box (fun (f: obj) -> (sharpurs_apply (box ((box f))) (box ((box x))))))))))) (box ((box ff))))))))))

[<AutoOpen>]
module PureScript_Data_Distributive

open System
open System.Collections.Generic

let Data_Distributive_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Data_Distributive_unwrap1  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Data_Distributive_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Distributive_Distributiveusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Distributive_distributiveIdentity  = (sharpurs_apply (box ((box Data_Distributive_Distributiveusd_Dict))) (box ((box ((Map.add "distribute" (box ((box (fun (dictFunctor: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box Data_Distributive_unwrap))))))))))) (Map.add "collect" (box ((box (fun (dictFunctor: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Identity_Identity)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Distributive_unwrap1)))))) (box ((box f)))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Identity_functorIdentity))))) Map.empty))))))))

let Data_Distributive_distribute  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "distribute" (unbox<Map<string, obj>> ((box v))))))))

let rec Data_Distributive_distributiveFunction : obj = ((sharpurs_apply (box ((box Data_Distributive_Distributiveusd_Dict))) (box ((box ((Map.add "distribute" (box ((box (fun (dictFunctor: obj) -> (box (fun (a: obj) -> (box (fun (e: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box v)))))) (box ((box e))))))))))) (box ((box a)))))))))))) (Map.add "collect" (box ((box (fun (dictFunctor: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Distributive_distribute))) (box ((box Data_Distributive_distributiveFunction)))))) (box ((box dictFunctor))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box f))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Functor_functorFn))))) Map.empty)))))))))


let Data_Distributive_cotraverse  = (box (fun (dictDistributive: obj) -> (let Functor0 = (sharpurs_apply (box ((Map.find "Functor0" (unbox<Map<string, obj>> ((box dictDistributive)))))) (box ((box Prim_undefined)))) in let distribute1 = (sharpurs_apply (box ((box Data_Distributive_distribute))) (box ((box dictDistributive)))) in (box (fun (dictFunctor: obj) -> (let distribute2 = (sharpurs_apply (box ((box distribute1))) (box ((box dictFunctor)))) in (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Functor0)))))) (box ((box f))))))))) (box ((box distribute2))))))))))))

let Data_Distributive_collectDefault  = (box (fun (dictDistributive: obj) -> (let distribute1 = (sharpurs_apply (box ((box Data_Distributive_distribute))) (box ((box dictDistributive)))) in (box (fun (dictFunctor: obj) -> (let distribute2 = (sharpurs_apply (box ((box distribute1))) (box ((box dictFunctor)))) in (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box distribute2)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box f)))))))))))))))

let rec Data_Distributive_distributiveTuple_tco (dictTypeEquals: obj) : obj = ((sharpurs_apply (box ((box Data_Distributive_Distributiveusd_Dict))) (box ((box ((Map.add "collect" (box ((box (fun (dictFunctor: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Distributive_collectDefault))) (box ((Data_Distributive_distributiveTuple_tco ((box dictTypeEquals)))))))) (box ((box dictFunctor)))))))) (Map.add "distribute" (box ((box (fun (dictFunctor: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Data_Tuple_Tupleusd_Ctor(usd__arg1, usd__arg2))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Type_Equality_from))) (box ((box dictTypeEquals)))))) (box ((box Data_Unit_unit)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box dictFunctor)))))) (box ((box Data_Tuple_snd))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Data_Tuple_functorTuple))))) Map.empty)))))))))
and Data_Distributive_distributiveTuple = box (fun (dictTypeEquals: obj) -> Data_Distributive_distributiveTuple_tco dictTypeEquals)


let Data_Distributive_collect  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "collect" (unbox<Map<string, obj>> ((box v))))))))

let Data_Distributive_distributeDefault  = (box (fun (dictDistributive: obj) -> (box (fun (dictFunctor: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Distributive_collect))) (box ((box dictDistributive)))))) (box ((box dictFunctor)))))) (box ((box Data_Distributive_identity))))))))

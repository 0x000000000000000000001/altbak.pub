[<AutoOpen>]
module PureScript_Data_Profunctor_Strong

open System
open System.Collections.Generic

let Data_Profunctor_Strong_Strongusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_Strong_strongFn  = (sharpurs_apply (box ((box Data_Profunctor_Strong_Strongusd_Dict))) (box ((box ((Map.add "first" (box ((box (fun (a2b: obj) -> (box (fun (v: obj) -> (match (((unbox ((box a2b))), (unbox ((box v))))) with | (a2b1, Data_Tuple_Tupleusd_Ctor(a, c)) -> ((box (Data_Tuple_Tupleusd_Ctor((sharpurs_apply (box ((box a2b1))) (box ((box a)))), (box c)))))))))))) (Map.add "second" (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Data_Tuple_functorTuple)))))) (Map.add "Profunctor0" (box ((box (fun (usd__unused: obj) -> (box Data_Profunctor_profunctorFn))))) Map.empty))))))))

let Data_Profunctor_Strong_second  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "second" (unbox<Map<string, obj>> ((box v))))))))

let Data_Profunctor_Strong_first  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "first" (unbox<Map<string, obj>> ((box v))))))))

let Data_Profunctor_Strong_splitStrong  = (box (fun (dictSemigroupoid: obj) -> (box (fun (dictStrong: obj) -> (box (fun (l: obj) -> (box (fun (r: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box dictSemigroupoid)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Strong_first))) (box ((box dictStrong)))))) (box ((box l))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Strong_second))) (box ((box dictStrong)))))) (box ((box r)))))))))))))))

let Data_Profunctor_Strong_fanout  = (box (fun (dictSemigroupoid: obj) -> (box (fun (dictStrong: obj) -> (let Profunctor0 = (sharpurs_apply (box ((Map.find "Profunctor0" (unbox<Map<string, obj>> ((box dictStrong)))))) (box ((box Prim_undefined)))) in (box (fun (l: obj) -> (box (fun (r: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_lcmap))) (box ((box Profunctor0)))))) (box ((box (fun (a: obj) -> (box (Data_Tuple_Tupleusd_Ctor((box a), (box a))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_Strong_splitStrong))) (box ((box dictSemigroupoid)))))) (box ((box dictStrong)))))) (box ((box l)))))) (box ((box r))))))))))))))))

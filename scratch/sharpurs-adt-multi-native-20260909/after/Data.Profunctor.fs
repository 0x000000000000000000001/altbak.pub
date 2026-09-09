[<AutoOpen>]
module PureScript_Data_Profunctor

open System
open System.Collections.Generic

let Data_Profunctor_identity  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Profunctor_identity1  = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box Control_Category_categoryFn))))

let Data_Profunctor_wrap  = (sharpurs_apply (box ((box Data_Newtype_wrap))) (box ((box Prim_undefined))))

let Data_Profunctor_unwrap  = (sharpurs_apply (box ((box Data_Newtype_unwrap))) (box ((box Prim_undefined))))

let Data_Profunctor_Profunctorusd_Dict  = (box (fun (x: obj) -> (box x)))

let Data_Profunctor_profunctorFn  = (sharpurs_apply (box ((box Data_Profunctor_Profunctorusd_Dict))) (box ((box ((Map.add "dimap" (box ((box (fun (a2b: obj) -> (box (fun (c2d: obj) -> (box (fun (b2c: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box a2b)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box b2c)))))) (box ((box c2d))))))))))))))) Map.empty))))))

let Data_Profunctor_dimap  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "dimap" (unbox<Map<string, obj>> ((box v))))))))

let Data_Profunctor_lcmap  = (box (fun (dictProfunctor: obj) -> (box (fun (a2b: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_dimap))) (box ((box dictProfunctor)))))) (box ((box a2b)))))) (box ((box Data_Profunctor_identity))))))))

let Data_Profunctor_rmap  = (box (fun (dictProfunctor: obj) -> (box (fun (b2c: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_dimap))) (box ((box dictProfunctor)))))) (box ((box Data_Profunctor_identity1)))))) (box ((box b2c))))))))

let Data_Profunctor_unwrapIso  = (box (fun (dictProfunctor: obj) -> (box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_dimap))) (box ((box dictProfunctor)))))) (box ((box Data_Profunctor_wrap)))))) (box ((box Data_Profunctor_unwrap))))))))

let Data_Profunctor_wrapIso  = (box (fun (dictProfunctor: obj) -> (box (fun (usd__unused: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_dimap))) (box ((box dictProfunctor)))))) (box ((box Data_Profunctor_unwrap)))))) (box ((box Data_Profunctor_wrap))))))))))

let Data_Profunctor_arr  = (box (fun (dictCategory: obj) -> (let identity2 = (sharpurs_apply (box ((box Control_Category_identity))) (box ((box dictCategory)))) in (box (fun (dictProfunctor: obj) -> (box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Profunctor_rmap))) (box ((box dictProfunctor)))))) (box ((box f)))))) (box ((box identity2)))))))))))

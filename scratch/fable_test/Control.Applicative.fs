[<AutoOpen>]
module PureScript_Control_Applicative

open System
open System.Collections.Generic

let Control_Applicative_Applicativeusd_Dict  = (box (fun (x: obj) -> (box x)))

let Control_Applicative_pure  = (box (fun (dict: obj) -> (match ((unbox ((box dict)))) with | v -> ((Map.find "pure" (unbox<Map<string, obj>> ((box v))))))))

let Control_Applicative_unless  = (box (fun (dictApplicative: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (LitBool false (), m) -> ((box m)) | (LitBool true (), _) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box dictApplicative)))))) (box ((box Data_Unit_unit))))))))))))

let Control_Applicative_when  = (box (fun (dictApplicative: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (LitBool true (), m) -> ((box m)) | (LitBool false (), _) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box dictApplicative)))))) (box ((box Data_Unit_unit))))))))))))

let Control_Applicative_liftA1  = (box (fun (dictApplicative: obj) -> (let Apply0 = (sharpurs_apply (box ((Map.find "Apply0" (unbox<Map<string, obj>> ((box dictApplicative)))))) (box ((box Prim_undefined)))) in (box (fun (f: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box Apply0)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box dictApplicative)))))) (box ((box f))))))))) (box ((box a)))))))))))

let Control_Applicative_applicativeProxy  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (v: obj) -> (box Type_Proxy_Proxyusd_Ctor))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Apply_applyProxy))))) Map.empty)))))))

let Control_Applicative_applicativeFn  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (x: obj) -> (box (fun (v: obj) -> (box x))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Apply_applyFn))))) Map.empty)))))))

let Control_Applicative_applicativeArray  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (x: obj) -> (box [|(box x)|]))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Control_Apply_applyArray))))) Map.empty)))))))

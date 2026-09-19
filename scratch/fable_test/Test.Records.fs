[<AutoOpen>]
module PureScript_Test_Records

open System
open System.Collections.Generic

let rec Test_Records_updateRec_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), r) -> ((box r)) | (n, r) -> ((Test_Records_updateRec_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((let v2 = (box r) in (box ((Map.add "a" (box ((box ((unbox<int> (box ((Map.find "a" (unbox<Map<string, obj>> ((box r))))))) + (unbox<int> (box ((box 1)))))))) (Map.add "b" (box ((let v3 = (Map.find "b" (unbox<Map<string, obj>> ((box r)))) in (box ((Map.add "c" (box ((box ((unbox<int> (box ((Map.find "c" (unbox<Map<string, obj>> ((Map.find "b" (unbox<Map<string, obj>> ((box r)))))))))) + (unbox<int> (box ((box 2)))))))) (Map.add "d" (box ((let v4 = (Map.find "d" (unbox<Map<string, obj>> ((Map.find "b" (unbox<Map<string, obj>> ((box r))))))) in (box ((Map.add "e" (box ((box ((unbox<int> (box ((Map.find "e" (unbox<Map<string, obj>> ((Map.find "d" (unbox<Map<string, obj>> ((Map.find "b" (unbox<Map<string, obj>> ((box r))))))))))))) + (unbox<int> (box ((box 3)))))))) (Map.add "f" (box ((box ((unbox<int> (box ((Map.find "f" (unbox<Map<string, obj>> ((Map.find "d" (unbox<Map<string, obj>> ((Map.find "b" (unbox<Map<string, obj>> ((box r))))))))))))) + (unbox<int> (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_EuclideanRing_mod))) (box ((box Data_EuclideanRing_euclideanRingInt)))))) (box ((box n)))))) (box ((box 5))))))))))) (unbox<Map<string, obj>> (box v4))))))))) (unbox<Map<string, obj>> (box v3))))))))) (unbox<Map<string, obj>> (box v2))))))))))))
and Test_Records_updateRec = box (fun (v: obj) ->  (fun (v1: obj) -> Test_Records_updateRec_tco v v1))


let Test_Records_initial  = (box ((Map.add "a" (box ((box 0))) (Map.add "b" (box ((box ((Map.add "c" (box ((box 0))) (Map.add "d" (box ((box ((Map.add "e" (box ((box 0))) (Map.add "f" (box ((box 0))) Map.empty)))))) Map.empty)))))) Map.empty))))

let Test_Records_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Deep Record Updates (10k iterations):"))))

let Test_Records_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((Map.find "f" (unbox<Map<string, obj>> ((Map.find "d" (unbox<Map<string, obj>> ((Map.find "b" (unbox<Map<string, obj>> ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_Records_updateRec))) (box ((box dummy)))))) (box ((box Test_Records_initial))))))))))))))))))))))))

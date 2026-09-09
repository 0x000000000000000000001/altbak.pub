[<AutoOpen>]
module PureScript_Test_Ackermann

open System
open System.Collections.Generic

let Test_Ackermann_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Ackermann (3, 4):"))))

let rec Test_Ackermann_ackermann_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), n) -> ((box ((unbox<int> (box ((box n)))) + (unbox<int> (box ((box 1))))))) | (m, LitInt 0 ()) -> ((Test_Ackermann_ackermann_tco ((box ((unbox<int> (box ((box m)))) - (unbox<int> (box ((box 1))))))) ((box 1)))) | (m, n) -> ((Test_Ackermann_ackermann_tco ((box ((unbox<int> (box ((box m)))) - (unbox<int> (box ((box 1))))))) ((Test_Ackermann_ackermann_tco ((box m)) ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1)))))))))))))
and Test_Ackermann_ackermann = box (fun (v: obj) ->  (fun (v1: obj) -> Test_Ackermann_ackermann_tco v v1))


let Test_Ackermann_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 3))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_Ackermann_ackermann))) (box ((box dummy)))))) (box ((box 4)))))))))))))))

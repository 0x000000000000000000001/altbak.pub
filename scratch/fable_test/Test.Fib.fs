[<AutoOpen>]
module PureScript_Test_Fib

open System
open System.Collections.Generic

let rec Test_Fib_fib_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | LitInt 0 () -> ((box 0)) | LitInt 1 () -> ((box 1)) | n -> ((box ((unbox<int> (box ((Test_Fib_fib_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))))))) + (unbox<int> (box ((Test_Fib_fib_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 2))))))))))))))))
and Test_Fib_fib = box (fun (v: obj) -> Test_Fib_fib_tco v)


let Test_Fib_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Fibonacci:"))))

let Test_Fib_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_Fib_fib))) (box ((box dummy)))))))))))))))

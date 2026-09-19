[<AutoOpen>]
module PureScript_Test_FibFFI

open System
open System.Collections.Generic

module Test_FibFFI_FFI =
    let rec fib n = if n < 2 then n else fib (n - 1) + fib (n - 2)
    let runFibFFI (n: obj) = fib (unbox<int> n) :> obj
    

let Test_FibFFI_runFibFFI = box (fun (arg0: obj) -> box (Test_FibFFI_FFI.``runFibFFI`` (unbox arg0)))


let Test_FibFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Fibonacci FFI:"))))

let Test_FibFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_FibFFI_runFibFFI))) (box ((box dummy)))))))))))))))

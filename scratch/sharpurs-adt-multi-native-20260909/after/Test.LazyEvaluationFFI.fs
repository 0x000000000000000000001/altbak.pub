[<AutoOpen>]
module PureScript_Test_LazyEvaluationFFI

open System
open System.Collections.Generic

module Test_LazyEvaluationFFI_FFI =
    let runLazyEvaluationFFI (n: obj) =
        let n' = unbox<int> n
        let mutable res = 0
        for i in 1 .. n' do res <- res + 1
        res :> obj
    

let Test_LazyEvaluationFFI_runLazyEvaluationFFI = box (fun (arg0: obj) -> box (Test_LazyEvaluationFFI_FFI.``runLazyEvaluationFFI`` (unbox arg0)))


let Test_LazyEvaluationFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Lazy Evaluation FFI (1M Thunks Forced, 1k Depth):"))))

let Test_LazyEvaluationFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 1000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_LazyEvaluationFFI_runLazyEvaluationFFI))) (box ((box dummy)))))))))))))))

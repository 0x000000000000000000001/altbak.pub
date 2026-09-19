[<AutoOpen>]
module PureScript_Test_ArrayOpsFFI

open System
open System.Collections.Generic

module Test_ArrayOpsFFI_FFI =
    let runArrayOpsFFI (n: obj) =
        let finish = unbox<int> n
        let step = if finish >= 1 then 1 else -1
        let values = [| 1 .. step .. finish |]
        values |> Array.filter (fun value -> value % 2 = 0) |> Array.fold (+) 0 |> box
    

let Test_ArrayOpsFFI_runArrayOpsFFI = box (fun (arg0: obj) -> box (Test_ArrayOpsFFI_FFI.``runArrayOpsFFI`` (unbox arg0)))


let Test_ArrayOpsFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Array Processing FFI (900 elements):"))))

let Test_ArrayOpsFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 900))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ArrayOpsFFI_runArrayOpsFFI))) (box ((box dummy)))))))))))))))

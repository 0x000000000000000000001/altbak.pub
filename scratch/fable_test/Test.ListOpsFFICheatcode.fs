[<AutoOpen>]
module PureScript_Test_ListOpsFFICheatcode

open System
open System.Collections.Generic

module Test_ListOpsFFICheatcode_FFI =
    let runListOpsFFICheatcode (n: obj) =
        let mutable sum = 0
        for value in 1 .. unbox<int> n do
            if value % 2 = 0 then sum <- sum + value
        sum :> obj
    

let Test_ListOpsFFICheatcode_runListOpsFFICheatcode = box (fun (arg0: obj) -> box (Test_ListOpsFFICheatcode_FFI.``runListOpsFFICheatcode`` (unbox arg0)))


let Test_ListOpsFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "List Processing FFICheatcode (900 elements):"))))

let Test_ListOpsFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 900))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ListOpsFFICheatcode_runListOpsFFICheatcode))) (box ((box dummy)))))))))))))))

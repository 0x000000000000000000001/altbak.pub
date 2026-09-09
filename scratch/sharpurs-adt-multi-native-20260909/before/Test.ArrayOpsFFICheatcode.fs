[<AutoOpen>]
module PureScript_Test_ArrayOpsFFICheatcode

open System
open System.Collections.Generic

module Test_ArrayOpsFFICheatcode_FFI =
    let runArrayOpsFFICheatcode (n: obj) =
        let n' = unbox<int> n
        let mutable sum = 0
        for i in 1 .. n' do sum <- sum + i
        sum * 5 :> obj
    

let Test_ArrayOpsFFICheatcode_runArrayOpsFFICheatcode = box (fun (arg0: obj) -> box (Test_ArrayOpsFFICheatcode_FFI.``runArrayOpsFFICheatcode`` (unbox arg0)))


let Test_ArrayOpsFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Array Processing FFICheatcode (900 elements):"))))

let Test_ArrayOpsFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 900))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ArrayOpsFFICheatcode_runArrayOpsFFICheatcode))) (box ((box dummy)))))))))))))))

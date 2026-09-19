[<AutoOpen>]
module PureScript_Test_ChurchFFICheatcode

open System
open System.Collections.Generic

module Test_ChurchFFICheatcode_FFI =
    let runChurchFFICheatcode (input: obj) =
        let count = pown (unbox<int> input) 5
        let mutable result = 0
        for _ in 1 .. count do result <- result + 1
        result :> obj
    

let Test_ChurchFFICheatcode_runChurchFFICheatcode = box (fun (arg0: obj) -> box (Test_ChurchFFICheatcode_FFI.``runChurchFFICheatcode`` (unbox arg0)))


let Test_ChurchFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Church Numerals FFICheatcode (100k Closure Applications):"))))

let Test_ChurchFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ChurchFFICheatcode_runChurchFFICheatcode))) (box ((box dummy)))))))))))))))

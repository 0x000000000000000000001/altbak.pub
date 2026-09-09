[<AutoOpen>]
module PureScript_Test_TCOFFICheatcode

open System
open System.Collections.Generic

module Test_TCOFFICheatcode_FFI =
    let runTCOFFICheatcode (n: obj) =
        let mutable acc = 0
        for i in 1 .. (unbox<int> n) do acc <- acc + 1
        acc :> obj
    

let Test_TCOFFICheatcode_runTCOFFICheatcode = box (fun (arg0: obj) -> box (Test_TCOFFICheatcode_FFI.``runTCOFFICheatcode`` (unbox arg0)))


let Test_TCOFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Tail Call Optimization FFICheatcode (100k calls):"))))

let Test_TCOFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_TCOFFICheatcode_runTCOFFICheatcode))) (box ((box dummy)))))))))))))))

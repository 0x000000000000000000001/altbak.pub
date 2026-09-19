[<AutoOpen>]
module PureScript_Test_StateMonadFFICheatcode

open System
open System.Collections.Generic

module Test_StateMonadFFICheatcode_FFI =
    let runStateMonadFFICheatcode (input: obj) =
        let depth = unbox<int> input
        let mutable total = 0
        for _ in 1 .. 20 do
            let mutable state = 0
            for _ in 1 .. depth do state <- state + 1
            total <- total + state
        total :> obj
    

let Test_StateMonadFFICheatcode_runStateMonadFFICheatcode = box (fun (arg0: obj) -> box (Test_StateMonadFFICheatcode_FFI.``runStateMonadFFICheatcode`` (unbox arg0)))


let Test_StateMonadFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "State Monad FFICheatcode (1.2k Binds, 60 Stack Depth):"))))

let Test_StateMonadFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 60))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_StateMonadFFICheatcode_runStateMonadFFICheatcode))) (box ((box dummy)))))))))))))))

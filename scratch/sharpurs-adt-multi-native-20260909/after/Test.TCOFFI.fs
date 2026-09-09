[<AutoOpen>]
module PureScript_Test_TCOFFI

open System
open System.Collections.Generic

module Test_TCOFFI_FFI =
    let rec loop n acc = if n = 0 then acc else loop (n - 1) (acc + 1)
    let runTCOFFI (n: obj) = loop (unbox<int> n) 0 :> obj
    

let Test_TCOFFI_runTCOFFI = box (fun (arg0: obj) -> box (Test_TCOFFI_FFI.``runTCOFFI`` (unbox arg0)))


let Test_TCOFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Tail Call Optimization FFI (100k calls):"))))

let Test_TCOFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_TCOFFI_runTCOFFI))) (box ((box dummy)))))))))))))))

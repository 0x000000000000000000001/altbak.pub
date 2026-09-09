[<AutoOpen>]
module PureScript_Test_AckermannFFI

open System
open System.Collections.Generic

module Test_AckermannFFI_FFI =
    let rec ack m n =
        if m = 0 then n + 1
        elif n = 0 then ack (m - 1) 1
        else ack (m - 1) (ack m (n - 1))
    let runAckermannFFI (args: obj) =
        ack 3 4 :> obj
    

let Test_AckermannFFI_runAckermannFFI = box (fun (arg0: obj) -> box (Test_AckermannFFI_FFI.``runAckermannFFI`` (unbox arg0)))


let Test_AckermannFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Ackermann FFI (3, 4):"))))

let Test_AckermannFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 3))))))))) (box ((box (fun (m: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_AckermannFFI_runAckermannFFI))) (box ((box m)))))))))))))))

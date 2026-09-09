[<AutoOpen>]
module PureScript_Test_RecordsFFI

open System
open System.Collections.Generic

module Test_RecordsFFI_FFI =
    type Record = { value: int }
    let runRecordsFFI (n: obj) =
        let mutable r = { value = 0 }
        for i in 1 .. (unbox<int> n) do r <- { value = r.value + 2 }
        r.value :> obj
    

let Test_RecordsFFI_runRecordsFFI = box (fun (arg0: obj) -> box (Test_RecordsFFI_FFI.``runRecordsFFI`` (unbox arg0)))


let Test_RecordsFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Deep Record Updates FFI (10k iterations):"))))

let Test_RecordsFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RecordsFFI_runRecordsFFI))) (box ((box dummy)))))))))))))))

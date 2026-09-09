[<AutoOpen>]
module PureScript_Test_RecordsFFICheatcode

open System
open System.Collections.Generic

module Test_RecordsFFICheatcode_FFI =
    type Record = { value: int }
    let runRecordsFFICheatcode (n: obj) =
        let mutable r = { value = 0 }
        for i in 1 .. (unbox<int> n) do r <- { value = r.value + 2 }
        r.value :> obj
    

let Test_RecordsFFICheatcode_runRecordsFFICheatcode = box (fun (arg0: obj) -> box (Test_RecordsFFICheatcode_FFI.``runRecordsFFICheatcode`` (unbox arg0)))


let Test_RecordsFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Deep Record Updates FFICheatcode (10k iterations):"))))

let Test_RecordsFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RecordsFFICheatcode_runRecordsFFICheatcode))) (box ((box dummy)))))))))))))))

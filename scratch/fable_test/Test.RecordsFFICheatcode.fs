[<AutoOpen>]
module PureScript_Test_RecordsFFICheatcode

open System
open System.Collections.Generic

module Test_RecordsFFICheatcode_FFI =
    type Inner = { mutable e: int; mutable f: int }
    type Middle = { mutable c: int; d: Inner }
    type DeepRecord = { mutable a: int; b: Middle }
    let runRecordsFFICheatcode (n: obj) =
        let record = { a = 0; b = { c = 0; d = { e = 0; f = 0 } } }
        for remaining in unbox<int> n .. -1 .. 1 do
            record.a <- record.a + 1
            record.b.c <- record.b.c + 2
            record.b.d.e <- record.b.d.e + 3
            record.b.d.f <- record.b.d.f + remaining % 5
        record.b.d.f :> obj
    

let Test_RecordsFFICheatcode_runRecordsFFICheatcode = box (fun (arg0: obj) -> box (Test_RecordsFFICheatcode_FFI.``runRecordsFFICheatcode`` (unbox arg0)))


let Test_RecordsFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Deep Record Updates FFICheatcode (10k iterations):"))))

let Test_RecordsFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RecordsFFICheatcode_runRecordsFFICheatcode))) (box ((box dummy)))))))))))))))

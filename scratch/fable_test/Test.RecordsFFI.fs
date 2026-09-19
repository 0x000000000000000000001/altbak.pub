[<AutoOpen>]
module PureScript_Test_RecordsFFI

open System
open System.Collections.Generic

module Test_RecordsFFI_FFI =
    type Inner = { e: int; f: int }
    type Middle = { c: int; d: Inner }
    type DeepRecord = { a: int; b: Middle }
    let rec update n r =
        if n = 0 then r
        else update (n - 1)
                { a = r.a + 1
                  b = { c = r.b.c + 2; d = { e = r.b.d.e + 3; f = r.b.d.f + n % 5 } } }
    let runRecordsFFI (n: obj) =
        let initial = { a = 0; b = { c = 0; d = { e = 0; f = 0 } } }
        (update (unbox<int> n) initial).b.d.f :> obj
    

let Test_RecordsFFI_runRecordsFFI = box (fun (arg0: obj) -> box (Test_RecordsFFI_FFI.``runRecordsFFI`` (unbox arg0)))


let Test_RecordsFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Deep Record Updates FFI (10k iterations):"))))

let Test_RecordsFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 10000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RecordsFFI_runRecordsFFI))) (box ((box dummy)))))))))))))))

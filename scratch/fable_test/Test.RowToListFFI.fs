[<AutoOpen>]
module PureScript_Test_RowToListFFI

open System
open System.Collections.Generic

module Test_RowToListFFI_FFI =
    type RecordKeys = { keysImpl: unit -> int }
    let keysCons tail = { keysImpl = fun () -> 1 + tail.keysImpl () }
    let runRowToListFFI (_input: obj) =
        let nil = { keysImpl = fun () -> 0 }
        let dictionary = keysCons (keysCons (keysCons (keysCons (keysCons nil))))
        dictionary.keysImpl () :> obj
    

let Test_RowToListFFI_runRowToListFFI = box (fun (arg0: obj) -> box (Test_RowToListFFI_FFI.``runRowToListFFI`` (unbox arg0)))


let Test_RowToListFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "RowToList FFI (Keys Count):"))))

let Test_RowToListFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 0))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RowToListFFI_runRowToListFFI))) (box ((box dummy)))))))))))))))

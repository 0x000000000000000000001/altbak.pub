[<AutoOpen>]
module PureScript_Effect_Console

open System
open System.Collections.Generic

module Effect_Console_FFI =
    let log (s: obj) = box (fun _ -> printfn "%s" (unbox<string> s); box null)
    let time _ = box null
    let timeLog _ = box null
    let timeEnd _ = box null
    let info (s: obj) = box (fun _ -> printfn "%s" (unbox<string> s); box null)
    let groupEnd _ = box null
    let groupCollapsed _ = box null
    let group _ = box null
    let error (s: obj) = box (fun _ -> eprintfn "%s" (unbox<string> s); box null)
    let debug (s: obj) = box (fun _ -> printfn "%s" (unbox<string> s); box null)
    let clear _ = box null
    let warn (s: obj) = box (fun _ -> eprintfn "%s" (unbox<string> s); box null)
    

let Effect_Console_clear = box (fun (arg0: obj) -> box (Effect_Console_FFI.``clear`` (unbox arg0)))
let Effect_Console_debug = box (fun (arg0: obj) -> box (Effect_Console_FFI.``debug`` (unbox arg0)))
let Effect_Console_error = box (fun (arg0: obj) -> box (Effect_Console_FFI.``error`` (unbox arg0)))
let Effect_Console_group = box (fun (arg0: obj) -> box (Effect_Console_FFI.``group`` (unbox arg0)))
let Effect_Console_groupCollapsed = box (fun (arg0: obj) -> box (Effect_Console_FFI.``groupCollapsed`` (unbox arg0)))
let Effect_Console_groupEnd = box (fun (arg0: obj) -> box (Effect_Console_FFI.``groupEnd`` (unbox arg0)))
let Effect_Console_info = box (fun (arg0: obj) -> box (Effect_Console_FFI.``info`` (unbox arg0)))
let Effect_Console_log = box (fun (arg0: obj) -> box (Effect_Console_FFI.``log`` (unbox arg0)))
let Effect_Console_time = box (fun (arg0: obj) -> box (Effect_Console_FFI.``time`` (unbox arg0)))
let Effect_Console_timeEnd = box (fun (arg0: obj) -> box (Effect_Console_FFI.``timeEnd`` (unbox arg0)))
let Effect_Console_timeLog = box (fun (arg0: obj) -> box (Effect_Console_FFI.``timeLog`` (unbox arg0)))
let Effect_Console_warn = box (fun (arg0: obj) -> box (Effect_Console_FFI.``warn`` (unbox arg0)))


let Effect_Console_warnShow  = (box (fun (dictShow: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((box Effect_Console_warn))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a)))))))))))

let Effect_Console_logShow  = (box (fun (dictShow: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((box Effect_Console_log))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a)))))))))))

let Effect_Console_infoShow  = (box (fun (dictShow: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((box Effect_Console_info))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a)))))))))))

let Effect_Console_grouped  = (box (fun (name: obj) -> (box (fun (inner: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Effect_Console_group))) (box ((box name))))))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((box inner)))))) (box ((box (fun (result: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Effect_bindEffect)))))) (box ((box Effect_Console_groupEnd)))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((box result)))))))))))))))))))))))

let Effect_Console_errorShow  = (box (fun (dictShow: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((box Effect_Console_error))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a)))))))))))

let Effect_Console_debugShow  = (box (fun (dictShow: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((box Effect_Console_debug))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box dictShow)))))) (box ((box a)))))))))))

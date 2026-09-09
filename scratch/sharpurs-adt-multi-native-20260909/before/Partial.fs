[<AutoOpen>]
module PureScript_Partial

open System
open System.Collections.Generic

module Partial_FFI =
    let _crashWith msg = failwith (unbox<string> msg)
    

let Partial__crashWith = box (fun (arg0: obj) -> box (Partial_FFI.``_crashWith`` (unbox arg0)))


let Partial_crashWith  = (box (fun (usd__unused: obj) -> (box Partial__crashWith)))

let Partial_crash  = (box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Partial_crashWith))) (box ((box Prim_undefined)))))) (box ((box "Partial.crash: partial function"))))))

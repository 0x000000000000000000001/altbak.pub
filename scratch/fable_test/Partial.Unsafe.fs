[<AutoOpen>]
module PureScript_Partial_Unsafe

open System
open System.Collections.Generic

module Partial_Unsafe_FFI =
    let _unsafePartial f = sharpurs_apply f (box null)
    

let Partial_Unsafe__unsafePartial = box (fun (arg0: obj) -> box (Partial_Unsafe_FFI.``_unsafePartial`` (unbox arg0)))


let Partial_Unsafe_unsafePartial  = (box Partial_Unsafe__unsafePartial)

let Partial_Unsafe_unsafeCrashWith  = (box (fun (msg: obj) -> (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Partial_crashWith))) (box ((box Prim_undefined)))))) (box ((box msg)))))))))))

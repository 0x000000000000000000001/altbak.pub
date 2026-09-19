[<AutoOpen>]
module PureScript_Safe_Coerce

open System
open System.Collections.Generic

let Safe_Coerce_coerce  = (box (fun (usd__unused: obj) -> (box Unsafe_Coerce_unsafeCoerce)))

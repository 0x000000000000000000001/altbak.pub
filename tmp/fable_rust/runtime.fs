[<AutoOpen>]
module Sharpurs_Prelude
open System
open System.Collections.Generic
let (|LitBool|_|) (expected: bool) (value: obj) =
    match value with
    | :? bool as n -> if n = expected then Some() else None
    | _ -> None
let (|LitInt|_|) (expected: int) (value: obj) =
    match value with
    | :? int as n -> if n = expected then Some() else None
    | _ -> None
let undefined = Unchecked.defaultof<obj>
let Prim_undefined = undefined
let sharpurs_apply (func: obj) (arg: obj) : obj =
    match func with
    | :? (obj -> obj) as invoke -> invoke arg
    | :? (obj -> obj -> obj) as invoke -> box (fun (next: obj) -> invoke arg next)
    | :? (obj -> obj -> obj -> obj) as invoke -> box (fun (next: obj) -> box (fun (last: obj) -> invoke arg next last))
    | _ -> failwith "Unsupported function shape in pure sharpurs benchmark"

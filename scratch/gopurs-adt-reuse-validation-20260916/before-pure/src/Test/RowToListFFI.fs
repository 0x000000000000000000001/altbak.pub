module Test.RowToListFFI

type RecordKeys = { keysImpl: unit -> int }
let keysCons tail = { keysImpl = fun () -> 1 + tail.keysImpl () }
let runRowToListFFI (_input: obj) =
    let nil = { keysImpl = fun () -> 0 }
    let dictionary = keysCons (keysCons (keysCons (keysCons (keysCons nil))))
    dictionary.keysImpl () :> obj

module Test.RowToListFFI

// F# has no RowToList constraint. The row type indexes its recursive dictionary.
type RowNil = RowNil
type RowCons<'head, 'tail> = { Head: 'head; Tail: 'tail }
type RecordKeys<'row> = { KeysImpl: unit -> int }
let keysNil : RecordKeys<RowNil> = { KeysImpl = fun () -> 0 }
let keysCons (tail: RecordKeys<'tail>) : RecordKeys<RowCons<'head, 'tail>> =
    { KeysImpl = fun () -> 1 + tail.KeysImpl () }
let keys (dictionary: RecordKeys<'row>) (_record: 'row) = dictionary.KeysImpl ()

let runRowToListFFI (_input: obj) =
    let record = { Head = 1; Tail = { Head = "two"; Tail = { Head = true; Tail = { Head = 4.0; Tail = { Head = "five"; Tail = RowNil } } } } }
    keys (keysCons (keysCons (keysCons (keysCons (keysCons keysNil))))) record :> obj

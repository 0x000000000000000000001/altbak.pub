module Test.RecordsFFICheatcode

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

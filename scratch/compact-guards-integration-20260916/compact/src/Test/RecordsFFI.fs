module Test.RecordsFFI

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

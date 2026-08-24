module Test.RecordsFFI
type Record = { value: int }
let runRecordsFFI (n: obj) =
    let mutable r = { value = 0 }
    for i in 1 .. (unbox<int> n) do r <- { value = r.value + 2 }
    r.value :> obj

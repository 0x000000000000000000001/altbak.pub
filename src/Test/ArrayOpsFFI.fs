module Test.ArrayOpsFFI

let runArrayOpsFFI (n: obj) =
    let finish = unbox<int> n
    let step = if finish >= 1 then 1 else -1
    let values = [| 1 .. step .. finish |]
    values |> Array.filter (fun value -> value % 2 = 0) |> Array.fold (+) 0 |> box

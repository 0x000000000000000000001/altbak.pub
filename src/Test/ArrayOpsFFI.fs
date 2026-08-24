module Test.ArrayOpsFFI
let runArrayOpsFFI (n: obj) =
    let arr = Array.init (unbox<int> n) (fun i -> i + 1)
    let mapped = Array.map (fun x -> x * 5) arr
    Array.fold (+) 0 mapped :> obj

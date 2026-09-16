module Test.LazyEvaluationFFI

type Thunk = unit -> int
let buildThunks depth (initial: Thunk) =
    let mutable result = initial
    for _ in 1 .. depth do
        let previous = result
        result <- fun () -> previous () + 1
    result
let runLazyEvaluationFFI (input: obj) =
    let mutable result = 0
    for _ in 1 .. unbox<int> input do
        result <- result + (buildThunks 1000 (fun () -> 0)) ()
    result :> obj

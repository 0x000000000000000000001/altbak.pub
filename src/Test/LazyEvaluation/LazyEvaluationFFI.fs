module Test.LazyEvaluationFFI

type Lazy<'a> = Lazy of (unit -> 'a)
let defer thunk = Lazy thunk
let force (Lazy thunk) = thunk ()
let rec buildThunks depth initial =
    if depth = 0 then initial
    else buildThunks (depth - 1) (defer (fun () -> force initial + 1))
let runLazyEvaluationFFI (input: obj) =
    let rec go count acc =
        if count = 0 then acc
        else go (count - 1) (acc + force (buildThunks 1000 (defer (fun () -> 0))))
    go (unbox<int> input) 0 :> obj

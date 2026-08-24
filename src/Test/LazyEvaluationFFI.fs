module Test.LazyEvaluationFFI
let runLazyEvaluationFFI (n: obj) =
    let n' = unbox<int> n
    let mutable res = 0
    for i in 1 .. n' do res <- res + 1
    res :> obj

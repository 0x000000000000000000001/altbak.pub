module Test.LazyEvaluationFFICheatcode

let runLazyEvaluationFFICheatcode (input: obj) =
    let mutable result = 0
    for _ in 1 .. unbox<int> input do
        for _ in 1 .. 1000 do result <- result + 1
    result :> obj

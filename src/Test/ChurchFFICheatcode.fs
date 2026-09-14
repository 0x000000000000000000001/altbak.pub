module Test.ChurchFFICheatcode

let runChurchFFICheatcode (input: obj) =
    let count = pown (unbox<int> input) 5
    let mutable result = 0
    for _ in 1 .. count do result <- result + 1
    result :> obj

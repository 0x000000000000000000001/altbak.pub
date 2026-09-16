module Test.ArrayOpsFFICheatcode

let runArrayOpsFFICheatcode (n: obj) =
    let finish = unbox<int> n
    let step = if finish >= 1 then 1 else -1
    let mutable sum = 0
    for value in 1 .. step .. finish do
        if value % 2 = 0 then sum <- sum + value
    sum :> obj

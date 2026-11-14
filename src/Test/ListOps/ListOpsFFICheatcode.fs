module Test.ListOpsFFICheatcode

let runListOpsFFICheatcode (n: obj) =
    let mutable sum = 0
    for value in 1 .. unbox<int> n do
        if value % 2 = 0 then sum <- sum + value
    sum :> obj

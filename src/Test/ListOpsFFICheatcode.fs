module Test.ListOpsFFICheatcode
let runListOpsFFICheatcode (n: obj) =
    let n' = unbox<int> n
    let mutable sum = 0
    for i in 1 .. n' do sum <- sum + i
    sum * 5 :> obj

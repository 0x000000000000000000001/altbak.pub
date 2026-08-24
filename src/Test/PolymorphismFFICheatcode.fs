module Test.PolymorphismFFICheatcode
let runPolymorphismFFICheatcode (n: obj) =
    let n' = unbox<int> n
    let mutable sum = 0
    for i in 1 .. n' do sum <- sum + 1
    sum :> obj

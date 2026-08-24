module Test.ChurchFFICheatcode
let runChurchFFICheatcode (n: obj) =
    let mutable res = 0
    for i in 1 .. (unbox<int> n) do res <- res + 1
    res :> obj

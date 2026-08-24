module Test.TCOFFICheatcode
let runTCOFFICheatcode (n: obj) =
    let mutable acc = 0
    for i in 1 .. (unbox<int> n) do acc <- acc + 1
    acc :> obj

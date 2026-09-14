module Test.TCOFFICheatcode

let runTCOFFICheatcode (n: obj) =
    let mutable acc = 0
    for remaining in unbox<int> n .. -1 .. 1 do acc <- acc + remaining % 3
    acc :> obj

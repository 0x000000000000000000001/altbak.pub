module Test.StateMonadFFICheatcode
let runStateMonadFFICheatcode (n: obj) =
    let n' = unbox<int> n
    let mutable state = 0
    for i in 1 .. n' do
        for j in 1 .. 60 do state <- state + 1
    state :> obj

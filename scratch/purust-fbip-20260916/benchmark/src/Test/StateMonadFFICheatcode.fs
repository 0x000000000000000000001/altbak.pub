module Test.StateMonadFFICheatcode

let runStateMonadFFICheatcode (input: obj) =
    let depth = unbox<int> input
    let mutable total = 0
    for _ in 1 .. 20 do
        let mutable state = 0
        for _ in 1 .. depth do state <- state + 1
        total <- total + state
    total :> obj

module Test.PrimesFFICheatcode
let isPrime n =
    if n < 2 then false
    else
        let mutable p = true
        let mutable i = 2
        while p && i * i <= n do
            if n % i = 0 then p <- false
            i <- i + 1
        p
let runPrimesFFICheatcode (n: obj) =
    let n' = unbox<int> n
    let mutable sum = 0
    for i in 2 .. n' do
        if isPrime i then sum <- sum + i
    sum :> obj

module Test.ChurchFFI
let runChurchFFI (n: obj) = 
    let mutable res = 0
    let f x = res <- res + 1; x
    for i in 1 .. (unbox<int> n) do f i |> ignore
    res :> obj

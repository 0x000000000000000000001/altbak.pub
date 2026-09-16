module Test.ChurchFFI

type Church = (int -> int) -> int -> int
let rec fromInt n : Church =
    if n = 0 then fun _ value -> value
    else
        let previous = fromInt (n - 1)
        fun f value -> f (previous f value)
let multiply (m: Church) (n: Church) : Church = fun f value -> m (n f) value
let square n = multiply (fromInt n) (fromInt n)
let runChurchFFI (input: obj) =
    let n = unbox<int> input
    let fourthPower = multiply (square n) (square n)
    let fifthPower = multiply fourthPower (fromInt n)
    fifthPower ((+) 1) 0 :> obj

module Test.ChurchFFI

type Church<'a> = ('a -> 'a) -> 'a -> 'a
let zero : Church<'a> = fun _ value -> value
let successor (previous: Church<'a>) : Church<'a> = fun f value -> f (previous f value)
let rec fromInt n : Church<int> =
    if n = 0 then zero else successor (fromInt (n - 1))
let multiply (m: Church<'a>) (n: Church<'a>) : Church<'a> = fun f value -> m (n f) value
let square n = multiply (fromInt n) (fromInt n)
let runChurchFFI (input: obj) =
    let n = unbox<int> input
    let fourthPower = multiply (square n) (square n)
    let fifthPower = multiply fourthPower (fromInt n)
    fifthPower ((+) 1) 0 :> obj

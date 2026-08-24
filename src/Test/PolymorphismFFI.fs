module Test.PolymorphismFFI
type Show = { show: int -> string }
let showInt : Show = { show = fun x -> string x }
let print (dict: Show) x = dict.show x
let runPolymorphismFFI (n: obj) =
    let mutable res = 0
    for i in 1 .. (unbox<int> n) do res <- res + 1
    res :> obj

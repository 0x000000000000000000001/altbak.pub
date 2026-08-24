module Test.PrimesFFI
type List = Nil | Cons of int * List
let rec range s e = if s > e then Nil else Cons (s, range (s + 1) e)
let rec filter p = function Nil -> Nil | Cons (x, xs) -> if p x then Cons (x, filter p xs) else filter p xs
let rec sieve = function Nil -> Nil | Cons (p, xs) -> Cons (p, sieve (filter (fun x -> x % p <> 0) xs))
let rec sum acc = function Nil -> acc | Cons (x, xs) -> sum (acc + x) xs
let runPrimesFFI (n: obj) = sum 0 (sieve (range 2 (unbox<int> n))) :> obj

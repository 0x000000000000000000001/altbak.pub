module Test.ListOpsFFI
type List = Nil | Cons of int * List
let rec range s e acc = if s > e then acc else range s (e - 1) (Cons (e, acc))
let rec map f = function Nil -> Nil | Cons (x, xs) -> Cons (f x, map f xs)
let rec sum acc = function Nil -> acc | Cons (x, xs) -> sum (acc + x) xs
let runListOpsFFI (n: obj) = sum 0 (map (fun x -> x * 5) (range 1 (unbox<int> n) Nil)) :> obj

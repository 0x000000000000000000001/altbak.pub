module Test.PrimesFFI
type List<'a> = Nil | Cons of 'a * List<'a>
let range start finish =
    let rec go current acc =
        if current < start then acc else go (current - 1) (Cons (current, acc))
    go finish Nil
let reverse list =
    let rec go rest acc =
        match rest with
        | Nil -> acc
        | Cons (value, tail) -> go tail (Cons (value, acc))
    go list Nil
let filter predicate list =
    let rec go rest acc =
        match rest with
        | Nil -> reverse acc
        | Cons (value, tail) -> go tail (if predicate value then Cons (value, acc) else acc)
    go list Nil
let rec sieve = function Nil -> Nil | Cons (p, xs) -> Cons (p, sieve (filter (fun x -> x % p <> 0) xs))
let rec sum acc = function Nil -> acc | Cons (x, xs) -> sum (acc + x) xs
let runPrimesFFI (n: obj) = sum 0 (sieve (range 2 (unbox<int> n))) :> obj

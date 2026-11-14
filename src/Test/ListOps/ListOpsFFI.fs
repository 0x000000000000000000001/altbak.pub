module Test.ListOpsFFI

type List<'a> = Nil | Cons of 'a * List<'a>
let range start finish =
    let rec go current acc =
        if current < start then acc else go (current - 1) (Cons (current, acc))
    go finish Nil
let filterEvens list =
    let rec go rest acc =
        match rest with
        | Nil -> acc
        | Cons (value, tail) -> go tail (if value % 2 = 0 then Cons (value, acc) else acc)
    go list Nil
let rec foldl f acc = function Nil -> acc | Cons (x, xs) -> foldl f (f acc x) xs
let runListOpsFFI (n: obj) = foldl (+) 0 (filterEvens (range 1 (unbox<int> n))) :> obj

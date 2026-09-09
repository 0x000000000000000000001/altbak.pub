[<AutoOpen>]
module PureScript_Test_PrimesFFI

open System
open System.Collections.Generic

module Test_PrimesFFI_FFI =
    type List = Nil | Cons of int * List
    let rec range s e = if s > e then Nil else Cons (s, range (s + 1) e)
    let rec filter p = function Nil -> Nil | Cons (x, xs) -> if p x then Cons (x, filter p xs) else filter p xs
    let rec sieve = function Nil -> Nil | Cons (p, xs) -> Cons (p, sieve (filter (fun x -> x % p <> 0) xs))
    let rec sum acc = function Nil -> acc | Cons (x, xs) -> sum (acc + x) xs
    let runPrimesFFI (n: obj) = sum 0 (sieve (range 2 (unbox<int> n))) :> obj
    

let Test_PrimesFFI_runPrimesFFI = box (fun (arg0: obj) -> box (Test_PrimesFFI_FFI.``runPrimesFFI`` (unbox arg0)))


let Test_PrimesFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Prime Sieve FFI (sum primes up to 500):"))))

let Test_PrimesFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 500))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_PrimesFFI_runPrimesFFI))) (box ((box dummy)))))))))))))))

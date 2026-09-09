[<AutoOpen>]
module PureScript_Test_ListOpsFFI

open System
open System.Collections.Generic

module Test_ListOpsFFI_FFI =
    type List = Nil | Cons of int * List
    let rec range s e acc = if s > e then acc else range s (e - 1) (Cons (e, acc))
    let rec map f = function Nil -> Nil | Cons (x, xs) -> Cons (f x, map f xs)
    let rec sum acc = function Nil -> acc | Cons (x, xs) -> sum (acc + x) xs
    let runListOpsFFI (n: obj) = sum 0 (map (fun x -> x * 5) (range 1 (unbox<int> n) Nil)) :> obj
    

let Test_ListOpsFFI_runListOpsFFI = box (fun (arg0: obj) -> box (Test_ListOpsFFI_FFI.``runListOpsFFI`` (unbox arg0)))


let Test_ListOpsFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "List Processing FFI (900 elements):"))))

let Test_ListOpsFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 900))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ListOpsFFI_runListOpsFFI))) (box ((box dummy)))))))))))))))

[<AutoOpen>]
module PureScript_Test_ListOpsFFI

open System
open System.Collections.Generic

module Test_ListOpsFFI_FFI =
    type List = Nil | Cons of int * List
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
    

let Test_ListOpsFFI_runListOpsFFI = box (fun (arg0: obj) -> box (Test_ListOpsFFI_FFI.``runListOpsFFI`` (unbox arg0)))


let Test_ListOpsFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "List Processing FFI (900 elements):"))))

let Test_ListOpsFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 900))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_ListOpsFFI_runListOpsFFI))) (box ((box dummy)))))))))))))))

[<AutoOpen>]
module PureScript_Test_RBTreeFFI

open System
open System.Collections.Generic

module Test_RBTreeFFI_FFI =
    type Color = Red | Black
    type Tree = Empty | Node of Color * Tree * int * Tree
    let balance = function
        | Black, Node(Red, Node(Red, a, x, b), y, c), z, d
        | Black, Node(Red, a, x, Node(Red, b, y, c)), z, d
        | Black, a, x, Node(Red, Node(Red, b, y, c), z, d)
        | Black, a, x, Node(Red, b, y, Node(Red, c, z, d)) -> Node(Red, Node(Black, a, x, b), y, Node(Black, c, z, d))
        | c, a, x, b -> Node(c, a, x, b)
    let insert x t =
        let rec ins = function
            | Empty -> Node(Red, Empty, x, Empty)
            | Node(c, a, y, b) ->
                if x < y then balance(c, ins a, y, b)
                elif x > y then balance(c, a, y, ins b)
                else Node(c, a, y, b)
        match ins t with
        | Node(_, a, y, b) -> Node(Black, a, y, b)
        | Empty -> Empty
    let rec count = function Empty -> 0 | Node(_, a, _, b) -> 1 + count a + count b
    let runRBTreeFFI (n: obj) =
        let rec build i t = if i > (unbox<int> n) then t else build (i + 1) (insert i t)
        count (build 1 Empty) :> obj
    

let Test_RBTreeFFI_runRBTreeFFI = box (fun (arg0: obj) -> box (Test_RBTreeFFI_FFI.``runRBTreeFFI`` (unbox arg0)))


let Test_RBTreeFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Red-Black Tree FFI (100k Worst-Case Insertions):"))))

let Test_RBTreeFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RBTreeFFI_runRBTreeFFI))) (box ((box dummy)))))))))))))))

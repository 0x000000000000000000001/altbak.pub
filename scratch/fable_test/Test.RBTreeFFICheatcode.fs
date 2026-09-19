[<AutoOpen>]
module PureScript_Test_RBTreeFFICheatcode

open System
open System.Collections.Generic

module Test_RBTreeFFICheatcode_FFI =
    [<AllowNullLiteral>]
    type Node(value: int) =
        member val Value = value
        member val Black = false with get, set
        member val Left: Node = null with get, set
        member val Right: Node = null with get, set
    let red (tree: Node) = not (isNull tree) && not tree.Black
    let balance (tree: Node) =
        if not tree.Black then tree
        elif red tree.Left && red tree.Left.Left then
            let root = tree.Left
            tree.Left <- root.Right
            root.Right <- tree
            root.Left.Black <- true
            root.Black <- false
            root
        elif red tree.Left && red tree.Left.Right then
            let child = tree.Left
            let root = child.Right
            child.Right <- root.Left
            tree.Left <- root.Right
            root.Left <- child
            root.Right <- tree
            child.Black <- true
            root.Black <- false
            root
        elif red tree.Right && red tree.Right.Left then
            let child = tree.Right
            let root = child.Left
            tree.Right <- root.Left
            child.Left <- root.Right
            root.Left <- tree
            root.Right <- child
            child.Black <- true
            root.Black <- false
            root
        elif red tree.Right && red tree.Right.Right then
            let root = tree.Right
            tree.Right <- root.Left
            root.Left <- tree
            root.Right.Black <- true
            root.Black <- false
            root
        else tree
    let rec insertRed value (tree: Node) =
        if isNull tree then Node(value)
        elif value < tree.Value then
            tree.Left <- insertRed value tree.Left
            balance tree
        elif value > tree.Value then
            tree.Right <- insertRed value tree.Right
            balance tree
        else tree
    let insert value tree =
        let root = insertRed value tree
        root.Black <- true
        root
    let rec depth (tree: Node) = if isNull tree then 0 else 1 + max (depth tree.Left) (depth tree.Right)
    let runRBTreeFFICheatcode (input: obj) =
        let mutable tree: Node = null
        for value in unbox<int> input .. -1 .. 1 do tree <- insert value tree
        depth tree :> obj
    

let Test_RBTreeFFICheatcode_runRBTreeFFICheatcode = box (fun (arg0: obj) -> box (Test_RBTreeFFICheatcode_FFI.``runRBTreeFFICheatcode`` (unbox arg0)))


let Test_RBTreeFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Red-Black Tree FFICheatcode (100k Worst-Case Insertions):"))))

let Test_RBTreeFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RBTreeFFICheatcode_runRBTreeFFICheatcode))) (box ((box dummy)))))))))))))))

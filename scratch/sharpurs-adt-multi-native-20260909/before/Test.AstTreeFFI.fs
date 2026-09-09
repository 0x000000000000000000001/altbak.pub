[<AutoOpen>]
module PureScript_Test_AstTreeFFI

open System
open System.Collections.Generic

module Test_AstTreeFFI_FFI =
    type Expr = Value of int | Add of Expr * Expr
    let rec eval = function Value n -> n | Add (l, r) -> eval l + eval r
    let rec makeTree depth = if depth = 0 then Value 1 else Add (makeTree (depth - 1), makeTree (depth - 1))
    let runAstTreeFFI (d: obj) = eval (makeTree (unbox<int> d)) :> obj
    

let Test_AstTreeFFI_runAstTreeFFI = box (fun (arg0: obj) -> box (Test_AstTreeFFI_FFI.``runAstTreeFFI`` (unbox arg0)))


let Test_AstTreeFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "AST Evaluation FFI:"))))

let Test_AstTreeFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 3))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_AstTreeFFI_runAstTreeFFI))) (box ((box dummy)))))))))))))))

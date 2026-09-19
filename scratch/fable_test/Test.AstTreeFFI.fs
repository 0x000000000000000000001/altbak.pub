[<AutoOpen>]
module PureScript_Test_AstTreeFFI

open System
open System.Collections.Generic

module Test_AstTreeFFI_FFI =
    type Expr = Val of int | Add of Expr * Expr | Mul of Expr * Expr | Sub of Expr * Expr
    let rec build n =
        if n = 0 then Val 1
        else Add (Mul (Val n, build (n - 1)), Sub (build (n - 1), Val 1))
    let rec eval = function
        | Val n -> n
        | Add (a, b) -> eval a + eval b
        | Mul (a, b) -> eval a * eval b
        | Sub (a, b) -> eval a - eval b
    let runAstTreeFFI (n: obj) = eval (build (unbox<int> n)) :> obj
    

let Test_AstTreeFFI_runAstTreeFFI = box (fun (arg0: obj) -> box (Test_AstTreeFFI_FFI.``runAstTreeFFI`` (unbox arg0)))


let Test_AstTreeFFI_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "AST Evaluation FFI:"))))

let Test_AstTreeFFI_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 3))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_AstTreeFFI_runAstTreeFFI))) (box ((box dummy)))))))))))))))

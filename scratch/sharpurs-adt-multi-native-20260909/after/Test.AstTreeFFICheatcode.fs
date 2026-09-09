[<AutoOpen>]
module PureScript_Test_AstTreeFFICheatcode

open System
open System.Collections.Generic

module Test_AstTreeFFICheatcode_FFI =
    type Expr = Value of int | Add of Expr * Expr
    let rec eval = function Value n -> n | Add (l, r) -> eval l + eval r
    let rec makeTree depth = if depth = 0 then Value 1 else Add (makeTree (depth - 1), makeTree (depth - 1))
    let runAstTreeFFICheatcode (d: obj) = eval (makeTree (unbox<int> d)) :> obj
    

let Test_AstTreeFFICheatcode_runAstTreeFFICheatcode = box (fun (arg0: obj) -> box (Test_AstTreeFFICheatcode_FFI.``runAstTreeFFICheatcode`` (unbox arg0)))


let Test_AstTreeFFICheatcode_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "AST Evaluation FFICheatcode:"))))

let Test_AstTreeFFICheatcode_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 3))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_AstTreeFFICheatcode_runAstTreeFFICheatcode))) (box ((box dummy)))))))))))))))

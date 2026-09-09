[<AutoOpen>]
module PureScript_Test_AstTree

open System
open System.Collections.Generic

type Test_AstTree_Expr =
  | Test_AstTree_Valusd_Ctor of obj
  | Test_AstTree_Addusd_Ctor of obj * obj
  | Test_AstTree_Mulusd_Ctor of obj * obj
  | Test_AstTree_Subusd_Ctor of obj * obj

let Test_AstTree_Val  = (box ((fun (usd__arg1: obj) -> (box (Test_AstTree_Valusd_Ctor(usd__arg1))))))

let Test_AstTree_Add  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Test_AstTree_Addusd_Ctor(usd__arg1, usd__arg2)))))))

let Test_AstTree_Mul  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Test_AstTree_Mulusd_Ctor(usd__arg1, usd__arg2)))))))

let Test_AstTree_Sub  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (box (Test_AstTree_Subusd_Ctor(usd__arg1, usd__arg2)))))))

let rec Test_AstTree_eval_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | Test_AstTree_Valusd_Ctor(n) -> ((box n)) | Test_AstTree_Addusd_Ctor(a, b) -> ((box ((unbox<int> (box ((Test_AstTree_eval_tco ((box a)))))) + (unbox<int> (box ((Test_AstTree_eval_tco ((box b))))))))) | Test_AstTree_Mulusd_Ctor(a, b) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_mul))) (box ((box Data_Semiring_semiringInt)))))) (box ((Test_AstTree_eval_tco ((box a)))))))) (box ((Test_AstTree_eval_tco ((box b))))))) | Test_AstTree_Subusd_Ctor(a, b) -> ((box ((unbox<int> (box ((Test_AstTree_eval_tco ((box a)))))) - (unbox<int> (box ((Test_AstTree_eval_tco ((box b)))))))))))
and Test_AstTree_eval = box (fun (v: obj) -> Test_AstTree_eval_tco v)


let Test_AstTree_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "AST Evaluation:"))))

let rec Test_AstTree_buildTree_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | LitInt 0 () -> ((box (Test_AstTree_Valusd_Ctor((box 1))))) | n -> ((box (Test_AstTree_Addusd_Ctor((box (Test_AstTree_Mulusd_Ctor((box (Test_AstTree_Valusd_Ctor((box n)))), (Test_AstTree_buildTree_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))))))), (box (Test_AstTree_Subusd_Ctor((Test_AstTree_buildTree_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1)))))))), (box (Test_AstTree_Valusd_Ctor((box 1)))))))))))))
and Test_AstTree_buildTree = box (fun (v: obj) -> Test_AstTree_buildTree_tco v)


let Test_AstTree_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 3))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_AstTree_eval))) (box ((sharpurs_apply (box ((box Test_AstTree_buildTree))) (box ((box dummy))))))))))))))))))

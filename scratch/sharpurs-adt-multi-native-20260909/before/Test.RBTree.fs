[<AutoOpen>]
module PureScript_Test_RBTree

open System
open System.Collections.Generic

type Test_RBTree_Color =
    | Test_RBTree_Rusd_Ctor
    | Test_RBTree_Busd_Ctor
and Test_RBTree_Tree =
    | Test_RBTree_Eusd_Ctor
    | Test_RBTree_Tusd_Ctor of Test_RBTree_Color * Test_RBTree_Tree * int * Test_RBTree_Tree

let Test_RBTree_R_adt_native  : Test_RBTree_Color = Test_RBTree_Rusd_Ctor
let Test_RBTree_R : obj = box (Test_RBTree_R_adt_native)

let Test_RBTree_B_adt_native  : Test_RBTree_Color = Test_RBTree_Busd_Ctor
let Test_RBTree_B : obj = box (Test_RBTree_B_adt_native)

let Test_RBTree_E_adt_native  : Test_RBTree_Tree = Test_RBTree_Eusd_Ctor
let Test_RBTree_E : obj = box (Test_RBTree_E_adt_native)

let Test_RBTree_T_adt_native (sharpurs_adt_local_0: Test_RBTree_Color) (sharpurs_adt_local_1: Test_RBTree_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: Test_RBTree_Tree) : Test_RBTree_Tree = Test_RBTree_Tusd_Ctor(sharpurs_adt_local_0, sharpurs_adt_local_1, sharpurs_adt_local_2, sharpurs_adt_local_3)
let Test_RBTree_T : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (fun (sharpurs_adt_arg_2: obj) -> box (fun (sharpurs_adt_arg_3: obj) -> box (Test_RBTree_T_adt_native (unbox<Test_RBTree_Color> sharpurs_adt_arg_0) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_1) (unbox<int> sharpurs_adt_arg_2) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_3))))))

let Test_RBTree_max_direct (x: obj) (y: obj) : obj = (match ((unbox ((box ((unbox<int> (box ((box x)))) > (unbox<int> (box ((box y))))))))) with | LitBool true () -> ((box x)) | _ -> ((box y)))

let Test_RBTree_max_direct_apply (x: obj) (y: obj) : obj =
    try Test_RBTree_max_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let Test_RBTree_max = (box (fun (x: obj) -> (box (fun (y: obj) -> (Test_RBTree_max_direct x y)))))

let Test_RBTree_makeBlack_adt_native (sharpurs_adt_local_0: Test_RBTree_Tree) : Test_RBTree_Tree = (if (match sharpurs_adt_local_0 with | Test_RBTree_Tusd_Ctor(_, _, _, _) -> true | _ -> false) then Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, (match sharpurs_adt_local_0 with | Test_RBTree_Tusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"), (match sharpurs_adt_local_0 with | Test_RBTree_Tusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"), (match sharpurs_adt_local_0 with | Test_RBTree_Tusd_Ctor(_, _, _, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) else (if (match sharpurs_adt_local_0 with | Test_RBTree_Eusd_Ctor -> true | _ -> false) then Test_RBTree_Eusd_Ctor else (failwith "Failed pattern match")))
let Test_RBTree_makeBlack : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (Test_RBTree_makeBlack_adt_native (unbox<Test_RBTree_Tree> sharpurs_adt_arg_0)))

let Test_RBTree_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Red-Black Tree (100k Worst-Case Insertions):"))))

let rec Test_RBTree_depth_adt_native (sharpurs_adt_local_0: Test_RBTree_Tree) : int = (if (match sharpurs_adt_local_0 with | Test_RBTree_Eusd_Ctor -> true | _ -> false) then (0) else (if (match sharpurs_adt_local_0 with | Test_RBTree_Tusd_Ctor(_, _, _, _) -> true | _ -> false) then ((1) + (let sharpurs_adt_local_1: int = (Test_RBTree_depth_adt_native ((match sharpurs_adt_local_0 with | Test_RBTree_Tusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))) in (let sharpurs_adt_local_2: int = (Test_RBTree_depth_adt_native ((match sharpurs_adt_local_0 with | Test_RBTree_Tusd_Ctor(_, _, _, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))) in (if (sharpurs_adt_local_1 > sharpurs_adt_local_2) then sharpurs_adt_local_1 else sharpurs_adt_local_2)))) else (failwith "Failed pattern match")))
let Test_RBTree_depth : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (Test_RBTree_depth_adt_native (unbox<Test_RBTree_Tree> sharpurs_adt_arg_0)))
let Test_RBTree_depth_tco (sharpurs_adt_arg: obj) : obj = box (Test_RBTree_depth_adt_native (unbox<Test_RBTree_Tree> sharpurs_adt_arg))

let Test_RBTree_balance_direct (v: obj) (v1: obj) (v2: obj) (v3: obj) : obj = (match (((unbox ((box v))), (unbox ((box v1))), (unbox ((box v2))), (unbox ((box v3))))) with | (Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), a, x, b)), y, c), z, d) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), a, x, Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), b, y, c))), z, d) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (Test_RBTree_Busd_Ctor, a, x, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), b, y, c)), z, d)) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (Test_RBTree_Busd_Ctor, a, x, Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), b, y, Unbox(Test_RBTree_Tusd_Ctor(Unbox(Test_RBTree_Rusd_Ctor), c, z, d)))) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b))))))) (unbox ((box y))) (unbox ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_B_adt_native)))) (unbox ((box c))) (unbox ((box z))) (unbox ((box d)))))))))) | (color, a, x, b) -> ((box (Test_RBTree_T_adt_native (unbox ((box color))) (unbox ((box a))) (unbox ((box x))) (unbox ((box b)))))))

let Test_RBTree_balance_direct_apply (v: obj) (v1: obj) (v2: obj) (v3: obj) : obj =
    try Test_RBTree_balance_direct v v1 v2 v3
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let Test_RBTree_balance = (box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box (fun (v3: obj) -> (Test_RBTree_balance_direct v v1 v2 v3)))))))))

let rec Test_RBTree_ins_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (x, Test_RBTree_Eusd_Ctor) -> ((box (Test_RBTree_T_adt_native (unbox ((box (Test_RBTree_R_adt_native)))) (unbox ((box (Test_RBTree_E_adt_native)))) (unbox ((box x))) (unbox ((box (Test_RBTree_E_adt_native))))))) | (x, Test_RBTree_Tusd_Ctor(color, a, y, b)) -> ((match ((unbox ((box ((unbox<int> (box ((box x)))) < (unbox<int> (box ((box y))))))))) with | LitBool true () -> ((Test_RBTree_balance_direct_apply ((box ((box color)))) ((box ((Test_RBTree_ins_tco ((box x)) ((box a)))))) ((box ((box y)))) ((box ((box b)))))) | _ -> ((match ((unbox ((box ((unbox<int> (box ((box x)))) > (unbox<int> (box ((box y))))))))) with | LitBool true () -> ((Test_RBTree_balance_direct_apply ((box ((box color)))) ((box ((box a)))) ((box ((box y)))) ((box ((Test_RBTree_ins_tco ((box x)) ((box b)))))))) | _ -> ((box (Test_RBTree_T_adt_native (unbox ((box color))) (unbox ((box a))) (unbox ((box y))) (unbox ((box b))))))))))))
and Test_RBTree_ins = box (fun (v: obj) ->  (fun (v1: obj) -> Test_RBTree_ins_tco v v1))


let Test_RBTree_insert_direct (x: obj) (s: obj) : obj = (sharpurs_apply (box ((box Test_RBTree_makeBlack))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_RBTree_ins))) (box ((box x)))))) (box ((box s)))))))

let Test_RBTree_insert_direct_apply (x: obj) (s: obj) : obj =
    try Test_RBTree_insert_direct x s
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let Test_RBTree_insert = (box (fun (x: obj) -> (box (fun (s: obj) -> (Test_RBTree_insert_direct x s)))))

let rec Test_RBTree_buildTree_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((Test_RBTree_buildTree_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((Test_RBTree_insert_direct_apply ((box ((box n)))) ((box ((box acc))))))))))
and Test_RBTree_buildTree = box (fun (v: obj) ->  (fun (v1: obj) -> Test_RBTree_buildTree_tco v v1))


let Test_RBTree_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RBTree_depth))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_RBTree_buildTree))) (box ((box dummy)))))) (box ((box (Test_RBTree_E_adt_native)))))))))))))))))))

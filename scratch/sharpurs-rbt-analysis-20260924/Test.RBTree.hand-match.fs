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

let Test_RBTree_makeBlack_adt_native (sharpurs_adt_local_0: Test_RBTree_Tree) : Test_RBTree_Tree =
    match sharpurs_adt_local_0 with
    | Test_RBTree_Tusd_Ctor(_, a1, y1, b1) -> Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, a1, y1, b1)
    | Test_RBTree_Eusd_Ctor -> Test_RBTree_Eusd_Ctor
    | _ -> failwith "Failed pattern match"
let Test_RBTree_makeBlack_adt_native_apply (sharpurs_adt_local_0: Test_RBTree_Tree) : Test_RBTree_Tree =
    try Test_RBTree_makeBlack_adt_native sharpurs_adt_local_0
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let Test_RBTree_makeBlack : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (Test_RBTree_makeBlack_adt_native (unbox<Test_RBTree_Tree> sharpurs_adt_arg_0)))

let Test_RBTree_describe  = (sharpurs_apply (box ((box Effect_Console_log))) (box ((box "Red-Black Tree (100k Worst-Case Insertions):"))))

let rec Test_RBTree_depth_adt_native (sharpurs_adt_local_0: Test_RBTree_Tree) : int =
    match sharpurs_adt_local_0 with
    | Test_RBTree_Eusd_Ctor -> 0
    | Test_RBTree_Tusd_Ctor(_, a1, _, b1) ->
        let left = Test_RBTree_depth_adt_native a1
        let right = Test_RBTree_depth_adt_native b1
        if left > right then left else right
    | _ -> failwith "Failed pattern match"
let Test_RBTree_depth_adt_native_apply (sharpurs_adt_local_0: Test_RBTree_Tree) : int =
    try Test_RBTree_depth_adt_native sharpurs_adt_local_0
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let Test_RBTree_depth : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (Test_RBTree_depth_adt_native (unbox<Test_RBTree_Tree> sharpurs_adt_arg_0)))
let Test_RBTree_depth_tco (sharpurs_adt_arg_0: obj) : obj = box (Test_RBTree_depth_adt_native (unbox<Test_RBTree_Tree> sharpurs_adt_arg_0))

let rec Test_RBTree_balance_adt_native (sharpurs_adt_local_0: Test_RBTree_Color) (sharpurs_adt_local_1: Test_RBTree_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: Test_RBTree_Tree) : Test_RBTree_Tree =
    match sharpurs_adt_local_0, sharpurs_adt_local_1, sharpurs_adt_local_3 with
    | Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, a1, x1, b1), y1, c1), d1 ->
        Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, a1, x1, b1), y1, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, c1, sharpurs_adt_local_2, d1))
    | Test_RBTree_Busd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, a1, x1, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, b1, y1, c1)), d1 ->
        Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, a1, x1, b1), y1, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, c1, sharpurs_adt_local_2, d1))
    | Test_RBTree_Busd_Ctor, a1, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, b1, y1, c1), z1, d1) ->
        Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, a1, sharpurs_adt_local_2, b1), y1, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, c1, z1, d1))
    | Test_RBTree_Busd_Ctor, a1, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, b1, y1, Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, c1, z1, d1)) ->
        Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, a1, sharpurs_adt_local_2, b1), y1, Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, c1, z1, d1))
    | color, a1, d1 ->
        Test_RBTree_Tusd_Ctor(color, a1, sharpurs_adt_local_2, d1)
let Test_RBTree_balance_adt_native_apply (sharpurs_adt_local_0: Test_RBTree_Color) (sharpurs_adt_local_1: Test_RBTree_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: Test_RBTree_Tree) : Test_RBTree_Tree =
    try Test_RBTree_balance_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1 sharpurs_adt_local_2 sharpurs_adt_local_3
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let Test_RBTree_balance : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (fun (sharpurs_adt_arg_2: obj) -> box (fun (sharpurs_adt_arg_3: obj) -> box (Test_RBTree_balance_adt_native (unbox<Test_RBTree_Color> sharpurs_adt_arg_0) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_1) (unbox<int> sharpurs_adt_arg_2) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_3))))))

let rec Test_RBTree_ins_adt_native (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: Test_RBTree_Tree) : Test_RBTree_Tree =
    match sharpurs_adt_local_1 with
    | Test_RBTree_Eusd_Ctor ->
        Test_RBTree_Tusd_Ctor(Test_RBTree_Rusd_Ctor, Test_RBTree_Eusd_Ctor, sharpurs_adt_local_0, Test_RBTree_Eusd_Ctor)
    | Test_RBTree_Tusd_Ctor(color, a1, y1, b1) ->
        if sharpurs_adt_local_0 < y1 then Test_RBTree_balance_adt_native_apply color (Test_RBTree_ins_adt_native sharpurs_adt_local_0 a1) y1 b1
        elif sharpurs_adt_local_0 > y1 then Test_RBTree_balance_adt_native_apply color a1 y1 (Test_RBTree_ins_adt_native sharpurs_adt_local_0 b1)
        else Test_RBTree_Tusd_Ctor(color, a1, y1, b1)
    | _ -> failwith "Failed pattern match"
let Test_RBTree_ins_adt_native_apply (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: Test_RBTree_Tree) : Test_RBTree_Tree =
    try Test_RBTree_ins_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let Test_RBTree_ins : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (Test_RBTree_ins_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_1))))
let Test_RBTree_ins_tco (sharpurs_adt_arg_0: obj) (sharpurs_adt_arg_1: obj) : obj = box (Test_RBTree_ins_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_1))

let Test_RBTree_insert_adt_native (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: Test_RBTree_Tree) : Test_RBTree_Tree =
    match Test_RBTree_ins_adt_native_apply sharpurs_adt_local_0 sharpurs_adt_local_1 with
    | Test_RBTree_Tusd_Ctor(_, a1, y1, b1) -> Test_RBTree_Tusd_Ctor(Test_RBTree_Busd_Ctor, a1, y1, b1)
    | Test_RBTree_Eusd_Ctor -> Test_RBTree_Eusd_Ctor
    | _ -> failwith "Failed pattern match"
let Test_RBTree_insert_adt_native_apply (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: Test_RBTree_Tree) : Test_RBTree_Tree =
    try Test_RBTree_insert_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let Test_RBTree_insert : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (Test_RBTree_insert_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_1))))

let rec Test_RBTree_buildTree_adt_native (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: Test_RBTree_Tree) : Test_RBTree_Tree = (if (sharpurs_adt_local_0 = (0)) then sharpurs_adt_local_1 else (Test_RBTree_buildTree_adt_native ((sharpurs_adt_local_0 - (1))) ((Test_RBTree_insert_adt_native_apply (sharpurs_adt_local_0) (sharpurs_adt_local_1)))))
let Test_RBTree_buildTree_adt_native_apply (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: Test_RBTree_Tree) : Test_RBTree_Tree =
    try Test_RBTree_buildTree_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let Test_RBTree_buildTree : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (Test_RBTree_buildTree_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_1))))
let Test_RBTree_buildTree_tco (sharpurs_adt_arg_0: obj) (sharpurs_adt_arg_1: obj) : obj = box (Test_RBTree_buildTree_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<Test_RBTree_Tree> sharpurs_adt_arg_1))

let Test_RBTree_act  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Bench_opaque))) (box ((box 100000))))))))) (box ((box (fun (dummy: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Show_show))) (box ((box Data_Show_showInt)))))) (box ((sharpurs_apply (box ((box Test_RBTree_depth))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Test_RBTree_buildTree))) (box ((box dummy)))))) (box ((box (Test_RBTree_E_adt_native)))))))))))))))))))

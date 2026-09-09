type AdtMulti_Color =
    | AdtMulti_Rubyusd_Ctor
    | AdtMulti_Onyxusd_Ctor
and AdtMulti_Tree =
    | AdtMulti_Tipusd_Ctor
    | AdtMulti_Branchusd_Ctor of AdtMulti_Color * AdtMulti_Tree * int * AdtMulti_Tree

let AdtMulti_Ruby_adt_native  : AdtMulti_Color = AdtMulti_Rubyusd_Ctor
let AdtMulti_Ruby : obj = box (AdtMulti_Ruby_adt_native)

let AdtMulti_Onyx_adt_native  : AdtMulti_Color = AdtMulti_Onyxusd_Ctor
let AdtMulti_Onyx : obj = box (AdtMulti_Onyx_adt_native)

let AdtMulti_Tip_adt_native  : AdtMulti_Tree = AdtMulti_Tipusd_Ctor
let AdtMulti_Tip : obj = box (AdtMulti_Tip_adt_native)

let AdtMulti_Branch_adt_native (sharpurs_adt_local_0: AdtMulti_Color) (sharpurs_adt_local_1: AdtMulti_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: AdtMulti_Tree) : AdtMulti_Tree = AdtMulti_Branchusd_Ctor(sharpurs_adt_local_0, sharpurs_adt_local_1, sharpurs_adt_local_2, sharpurs_adt_local_3)
let AdtMulti_Branch : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (fun (sharpurs_adt_arg_2: obj) -> box (fun (sharpurs_adt_arg_3: obj) -> box (AdtMulti_Branch_adt_native (unbox<AdtMulti_Color> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1) (unbox<int> sharpurs_adt_arg_2) (unbox<AdtMulti_Tree> sharpurs_adt_arg_3))))))

let AdtMulti_unused_adt_native (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: int) : AdtMulti_Tree = sharpurs_adt_local_0
let AdtMulti_unused_adt_native_apply (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: int) : AdtMulti_Tree =
    try AdtMulti_unused_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_unused : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_unused_adt_native (unbox<AdtMulti_Tree> sharpurs_adt_arg_0) (unbox<int> sharpurs_adt_arg_1))))

let AdtMulti_score_adt_native (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: int) : int = (if (match sharpurs_adt_local_0 with | AdtMulti_Tipusd_Ctor -> true | _ -> false) then sharpurs_adt_local_1 else (if (match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, _, _, _) -> true | _ -> false) then ((match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") + sharpurs_adt_local_1) else (failwith "Failed pattern match")))
let AdtMulti_score_adt_native_apply (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: int) : int =
    try AdtMulti_score_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_score : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_score_adt_native (unbox<AdtMulti_Tree> sharpurs_adt_arg_0) (unbox<int> sharpurs_adt_arg_1))))

let AdtMulti_shortAndSafe_adt_native (sharpurs_adt_local_0: bool) (sharpurs_adt_local_1: AdtMulti_Tree) : bool = (sharpurs_adt_local_0 && ((AdtMulti_score_adt_native_apply (sharpurs_adt_local_1) ((0))) > (0)))
let AdtMulti_shortAndSafe_adt_native_apply (sharpurs_adt_local_0: bool) (sharpurs_adt_local_1: AdtMulti_Tree) : bool =
    try AdtMulti_shortAndSafe_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_shortAndSafe : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_shortAndSafe_adt_native (unbox<bool> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1))))

let AdtMulti_shortOrSafe_adt_native (sharpurs_adt_local_0: bool) (sharpurs_adt_local_1: AdtMulti_Tree) : bool = (sharpurs_adt_local_0 || ((AdtMulti_score_adt_native_apply (sharpurs_adt_local_1) ((0))) > (0)))
let AdtMulti_shortOrSafe_adt_native_apply (sharpurs_adt_local_0: bool) (sharpurs_adt_local_1: AdtMulti_Tree) : bool =
    try AdtMulti_shortOrSafe_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_shortOrSafe : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_shortOrSafe_adt_native (unbox<bool> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1))))

let AdtMulti_readNonEmptyInline_direct (tree: obj) (increment: obj) : obj = (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box (fun (usd__unused: obj) -> (match ((unbox ((box tree)))) with | AdtMulti_Branchusd_Ctor(_, _, value, _) -> ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment)))))))))))) (box ((box Prim_undefined)))))))))

let AdtMulti_readNonEmptyInline_direct_apply (tree: obj) (increment: obj) : obj =
    try AdtMulti_readNonEmptyInline_direct tree increment
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_readNonEmptyInline = (box (fun (tree: obj) -> (box (fun (increment: obj) -> (AdtMulti_readNonEmptyInline_direct tree increment)))))

let AdtMulti_readNonEmpty_direct (tree: obj) (increment: obj) : obj = (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box (fun (usd__unused: obj) -> (match ((unbox ((box tree)))) with | AdtMulti_Branchusd_Ctor(_, _, value, _) -> ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment)))))))))))) (box ((box Prim_undefined)))))))))

let AdtMulti_readNonEmpty_direct_apply (tree: obj) (increment: obj) : obj =
    try AdtMulti_readNonEmpty_direct tree increment
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_readNonEmpty = (box (fun (tree: obj) -> (box (fun (increment: obj) -> (AdtMulti_readNonEmpty_direct tree increment)))))

let AdtMulti_returned  = (box (fun (tree: obj) -> (let value = (sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_readNonEmpty))) (box ((box tree)))))) (box ((box 0)))) in (box (fun (increment: obj) -> (box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment)))))))))))

let AdtMulti_shortAnd_direct (flag: obj) (tree: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((box ((unbox<int> (box ((AdtMulti_readNonEmpty_direct_apply ((box ((box tree)))) ((box ((box 0)))))))) > (unbox<int> (box ((box 0))))))) | _ -> ((box false)))

let AdtMulti_shortAnd_direct_apply (flag: obj) (tree: obj) : obj =
    try AdtMulti_shortAnd_direct flag tree
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_shortAnd = (box (fun (flag: obj) -> (box (fun (tree: obj) -> (AdtMulti_shortAnd_direct flag tree)))))

let AdtMulti_shortOr_direct (flag: obj) (tree: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((box true)) | _ -> ((box ((unbox<int> (box ((AdtMulti_readNonEmpty_direct_apply ((box ((box tree)))) ((box ((box 0)))))))) > (unbox<int> (box ((box 0))))))))

let AdtMulti_shortOr_direct_apply (flag: obj) (tree: obj) : obj =
    try AdtMulti_shortOr_direct flag tree
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_shortOr = (box (fun (flag: obj) -> (box (fun (tree: obj) -> (AdtMulti_shortOr_direct flag tree)))))

let AdtMulti_nested_adt_native (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: int) : int = (if ((match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, _, _, _) -> true | _ -> false) && ((match (match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(sharpurs_adt_field, _, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") with | AdtMulti_Onyxusd_Ctor -> true | _ -> false) && ((match (match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") with | AdtMulti_Branchusd_Ctor(_, _, _, _) -> true | _ -> false) && (match (match (match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") with | AdtMulti_Branchusd_Ctor(sharpurs_adt_field, _, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") with | AdtMulti_Rubyusd_Ctor -> true | _ -> false)))) then ((match (match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") with | AdtMulti_Branchusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") + sharpurs_adt_local_1) else sharpurs_adt_local_1)
let AdtMulti_nested_adt_native_apply (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: int) : int =
    try AdtMulti_nested_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_nested : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_nested_adt_native (unbox<AdtMulti_Tree> sharpurs_adt_arg_0) (unbox<int> sharpurs_adt_arg_1))))

let AdtMulti_importedTotal_direct (v: obj) (v1: obj) : obj = (match (((unbox ((box v))), (unbox ((box v1))))) with | (AdtMulti_Tipusd_Ctor, increment) -> ((sharpurs_apply (box ((box AdtMultiExternal_offset))) (box ((box increment))))) | (AdtMulti_Branchusd_Ctor(_, _, value, _), increment) -> ((sharpurs_apply (box ((box AdtMultiExternal_offset))) (box ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment)))))))))))

let AdtMulti_importedTotal_direct_apply (v: obj) (v1: obj) : obj =
    try AdtMulti_importedTotal_direct v v1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_importedTotal = (box (fun (v: obj) -> (box (fun (v1: obj) -> (AdtMulti_importedTotal_direct v v1)))))

let AdtMulti_importedFailure_direct (v: obj) (v1: obj) : obj = (match (((unbox ((box v))), (unbox ((box v1))))) with | (AdtMulti_Tipusd_Ctor, increment) -> ((sharpurs_apply (box ((box AdtMultiExternal_readZero))) (box ((box increment))))) | (AdtMulti_Branchusd_Ctor(_, _, value, _), increment) -> ((sharpurs_apply (box ((box AdtMultiExternal_readZero))) (box ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment)))))))))))

let AdtMulti_importedFailure_direct_apply (v: obj) (v1: obj) : obj =
    try AdtMulti_importedFailure_direct v v1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_importedFailure = (box (fun (v: obj) -> (box (fun (v1: obj) -> (AdtMulti_importedFailure_direct v v1)))))

let AdtMulti_throughImported_direct (tree: obj) (increment: obj) : obj = (box ((unbox<int> (box ((AdtMulti_importedFailure_direct_apply ((box ((box tree)))) ((box ((box increment)))))))) + (unbox<int> (box ((box 1))))))

let AdtMulti_throughImported_direct_apply (tree: obj) (increment: obj) : obj =
    try AdtMulti_throughImported_direct tree increment
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_throughImported = (box (fun (tree: obj) -> (box (fun (increment: obj) -> (AdtMulti_throughImported_direct tree increment)))))

let AdtMulti_generic  = (box (fun (v: obj) -> (box (fun (value: obj) -> (box value)))))

let rec AdtMulti_depth_adt_native (sharpurs_adt_local_0: AdtMulti_Tree) : int = (if (match sharpurs_adt_local_0 with | AdtMulti_Tipusd_Ctor -> true | _ -> false) then (0) else (if (match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, _, _, _) -> true | _ -> false) then (let sharpurs_adt_local_1: int = (AdtMulti_depth_adt_native ((match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, _, _, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))) in (let sharpurs_adt_local_2: int = (AdtMulti_depth_adt_native ((match sharpurs_adt_local_0 with | AdtMulti_Branchusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))) in (if (sharpurs_adt_local_2 > sharpurs_adt_local_1) then ((1) + sharpurs_adt_local_2) else ((1) + sharpurs_adt_local_1)))) else (failwith "Failed pattern match")))
let AdtMulti_depth_adt_native_apply (sharpurs_adt_local_0: AdtMulti_Tree) : int =
    try AdtMulti_depth_adt_native sharpurs_adt_local_0
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_depth : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (AdtMulti_depth_adt_native (unbox<AdtMulti_Tree> sharpurs_adt_arg_0)))
let AdtMulti_depth_tco (sharpurs_adt_arg_0: obj) : obj = box (AdtMulti_depth_adt_native (unbox<AdtMulti_Tree> sharpurs_adt_arg_0))

let AdtMulti_combine_direct (first: obj) (a: obj) (second: obj) (b: obj) : obj = (box ((unbox<int> (box ((AdtMulti_readNonEmpty_direct_apply ((box ((box first)))) ((box ((box a)))))))) + (unbox<int> (box ((AdtMulti_readNonEmpty_direct_apply ((box ((box second)))) ((box ((box b))))))))))

let AdtMulti_combine_direct_apply (first: obj) (a: obj) (second: obj) (b: obj) : obj =
    try AdtMulti_combine_direct first a second b
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_combine = (box (fun (first: obj) -> (box (fun (a: obj) -> (box (fun (second: obj) -> (box (fun (b: obj) -> (AdtMulti_combine_direct first a second b)))))))))

let AdtMulti_booleanOnly_direct (flag: obj) (value: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((box value)) | _ -> ((box 0)))

let AdtMulti_booleanOnly_direct_apply (flag: obj) (value: obj) : obj =
    try AdtMulti_booleanOnly_direct flag value
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_booleanOnly = (box (fun (flag: obj) -> (box (fun (value: obj) -> (AdtMulti_booleanOnly_direct flag value)))))

let AdtMulti_bodyFailureInline_direct (tree: obj) (increment: obj) : obj = (box ((unbox<int> (box ((AdtMulti_readNonEmptyInline_direct_apply ((box ((box tree)))) ((box ((box increment)))))))) + (unbox<int> (box ((box 1))))))

let AdtMulti_bodyFailureInline_direct_apply (tree: obj) (increment: obj) : obj =
    try AdtMulti_bodyFailureInline_direct tree increment
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_bodyFailureInline = (box (fun (tree: obj) -> (box (fun (increment: obj) -> (AdtMulti_bodyFailureInline_direct tree increment)))))

let AdtMulti_bodyFailure_direct (tree: obj) (increment: obj) : obj = (box ((unbox<int> (box ((AdtMulti_readNonEmpty_direct_apply ((box ((box tree)))) ((box ((box increment)))))))) + (unbox<int> (box ((box 1))))))

let AdtMulti_bodyFailure_direct_apply (tree: obj) (increment: obj) : obj =
    try AdtMulti_bodyFailure_direct tree increment
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_bodyFailure = (box (fun (tree: obj) -> (box (fun (increment: obj) -> (AdtMulti_bodyFailure_direct tree increment)))))

let AdtMulti_assemble_adt_native (sharpurs_adt_local_0: AdtMulti_Color) (sharpurs_adt_local_1: AdtMulti_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: AdtMulti_Tree) : AdtMulti_Tree = AdtMulti_Branchusd_Ctor(sharpurs_adt_local_0, sharpurs_adt_local_1, sharpurs_adt_local_2, sharpurs_adt_local_3)
let AdtMulti_assemble_adt_native_apply (sharpurs_adt_local_0: AdtMulti_Color) (sharpurs_adt_local_1: AdtMulti_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: AdtMulti_Tree) : AdtMulti_Tree =
    try AdtMulti_assemble_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1 sharpurs_adt_local_2 sharpurs_adt_local_3
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_assemble : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (fun (sharpurs_adt_arg_2: obj) -> box (fun (sharpurs_adt_arg_3: obj) -> box (AdtMulti_assemble_adt_native (unbox<AdtMulti_Color> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1) (unbox<int> sharpurs_adt_arg_2) (unbox<AdtMulti_Tree> sharpurs_adt_arg_3))))))

let AdtMulti_choose_adt_native (sharpurs_adt_local_0: bool) (sharpurs_adt_local_1: AdtMulti_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: AdtMulti_Tree) : AdtMulti_Tree = (if sharpurs_adt_local_0 then (AdtMulti_assemble_adt_native_apply (AdtMulti_Onyxusd_Ctor) (sharpurs_adt_local_1) ((sharpurs_adt_local_2 + (1))) (sharpurs_adt_local_3)) else (AdtMulti_assemble_adt_native_apply (AdtMulti_Rubyusd_Ctor) (sharpurs_adt_local_3) ((sharpurs_adt_local_2 - (1))) (sharpurs_adt_local_1)))
let AdtMulti_choose_adt_native_apply (sharpurs_adt_local_0: bool) (sharpurs_adt_local_1: AdtMulti_Tree) (sharpurs_adt_local_2: int) (sharpurs_adt_local_3: AdtMulti_Tree) : AdtMulti_Tree =
    try AdtMulti_choose_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1 sharpurs_adt_local_2 sharpurs_adt_local_3
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_choose : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (fun (sharpurs_adt_arg_2: obj) -> box (fun (sharpurs_adt_arg_3: obj) -> box (AdtMulti_choose_adt_native (unbox<bool> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1) (unbox<int> sharpurs_adt_arg_2) (unbox<AdtMulti_Tree> sharpurs_adt_arg_3))))))

let rec AdtMulti_insert_adt_native (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: AdtMulti_Tree) : AdtMulti_Tree = (if (match sharpurs_adt_local_1 with | AdtMulti_Tipusd_Ctor -> true | _ -> false) then (AdtMulti_assemble_adt_native_apply (AdtMulti_Rubyusd_Ctor) (AdtMulti_Tipusd_Ctor) (sharpurs_adt_local_0) (AdtMulti_Tipusd_Ctor)) else (if (match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, _, _) -> true | _ -> false) then (if (sharpurs_adt_local_0 < (match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) then (AdtMulti_assemble_adt_native_apply ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(sharpurs_adt_field, _, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) ((AdtMulti_insert_adt_native (sharpurs_adt_local_0) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")))) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, _, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))) else (if (sharpurs_adt_local_0 > (match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) then (AdtMulti_assemble_adt_native_apply ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(sharpurs_adt_field, _, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) ((AdtMulti_insert_adt_native (sharpurs_adt_local_0) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, _, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))))) else sharpurs_adt_local_1)) else (failwith "Failed pattern match")))
let AdtMulti_insert_adt_native_apply (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: AdtMulti_Tree) : AdtMulti_Tree =
    try AdtMulti_insert_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_insert : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_insert_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1))))
let AdtMulti_insert_tco (sharpurs_adt_arg_0: obj) (sharpurs_adt_arg_1: obj) : obj = box (AdtMulti_insert_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1))

let AdtMulti_partial  = (sharpurs_apply (box ((box AdtMulti_assemble))) (box ((box (AdtMulti_Onyx_adt_native)))))

let rec AdtMulti_shift_adt_native (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: AdtMulti_Tree) : AdtMulti_Tree = (if (match sharpurs_adt_local_1 with | AdtMulti_Tipusd_Ctor -> true | _ -> false) then AdtMulti_Tipusd_Ctor else (if (match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, _, _) -> true | _ -> false) then (AdtMulti_assemble_adt_native_apply ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(sharpurs_adt_field, _, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")) ((AdtMulti_shift_adt_native (sharpurs_adt_local_0) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, sharpurs_adt_field, _, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor")))) (((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, sharpurs_adt_field, _) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor") + sharpurs_adt_local_0)) ((AdtMulti_shift_adt_native (sharpurs_adt_local_0) ((match sharpurs_adt_local_1 with | AdtMulti_Branchusd_Ctor(_, _, _, sharpurs_adt_field) -> sharpurs_adt_field | _ -> failwith "Invalid ADT constructor"))))) else (failwith "Failed pattern match")))
let AdtMulti_shift_adt_native_apply (sharpurs_adt_local_0: int) (sharpurs_adt_local_1: AdtMulti_Tree) : AdtMulti_Tree =
    try AdtMulti_shift_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_shift : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_shift_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1))))
let AdtMulti_shift_tco (sharpurs_adt_arg_0: obj) (sharpurs_adt_arg_1: obj) : obj = box (AdtMulti_shift_adt_native (unbox<int> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1))

let AdtMulti_argumentOrder_direct (first: obj) (second: obj) : obj = (AdtMulti_combine_direct_apply ((box ((box first)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box first)))) ((box ((box 0)))))))) ((box ((box second)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box second)))) ((box ((box 0)))))))))

let AdtMulti_argumentOrder_direct_apply (first: obj) (second: obj) : obj =
    try AdtMulti_argumentOrder_direct first second
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_argumentOrder = (box (fun (first: obj) -> (box (fun (second: obj) -> (AdtMulti_argumentOrder_direct first second)))))

let AdtMulti_argumentNative_adt_native (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: AdtMulti_Tree) : int = (AdtMulti_score_adt_native_apply (sharpurs_adt_local_1) ((AdtMulti_score_adt_native_apply (sharpurs_adt_local_0) ((0)))))
let AdtMulti_argumentNative_adt_native_apply (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: AdtMulti_Tree) : int =
    try AdtMulti_argumentNative_adt_native sharpurs_adt_local_0 sharpurs_adt_local_1
    with ex -> raise (System.Reflection.TargetInvocationException(ex))
let AdtMulti_argumentNative : obj = box (fun (sharpurs_adt_arg_0: obj) -> box (fun (sharpurs_adt_arg_1: obj) -> box (AdtMulti_argumentNative_adt_native (unbox<AdtMulti_Tree> sharpurs_adt_arg_0) (unbox<AdtMulti_Tree> sharpurs_adt_arg_1))))

let AdtMulti_argumentFailure_direct (outer: obj) (inner: obj) : obj = (AdtMulti_readNonEmpty_direct_apply ((box ((box outer)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box inner)))) ((box ((box 0)))))))))

let AdtMulti_argumentFailure_direct_apply (outer: obj) (inner: obj) : obj =
    try AdtMulti_argumentFailure_direct outer inner
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let AdtMulti_argumentFailure = (box (fun (outer: obj) -> (box (fun (inner: obj) -> (AdtMulti_argumentFailure_direct outer inner)))))
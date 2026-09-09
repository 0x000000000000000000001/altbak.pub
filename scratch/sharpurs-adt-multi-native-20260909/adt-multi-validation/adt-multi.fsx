let objMap = Map.empty<string, obj>
let unbox<'a> (x: obj) : 'a = unbox x
let (|Unbox|) (x: obj) = unbox x

let undefined = Unchecked.defaultof<obj>
let Prim_undefined = undefined
let intMod a b = unbox<int> a % unbox<int> b
let semiringInt = 0

// PureScript's Euclidean Int modulo, including zero and Int32.MinValue / -1.
let sharpurs_int_mod (left: int) (right: int) : int =
    if right = 0 || right = -1 then 0
    else
        let remainder = left % right
        if remainder < 0 then
            if right > 0 then remainder + right else remainder - right
        else remainder

let sharpurs_apply (func: obj) (arg: obj) : obj =
    if isNull func then failwith "sharpurs_apply: func is null!"
    match func with
    | :? (obj -> obj) as invoke ->
        try invoke arg
        // Keep the exception boundary of MethodInfo.Invoke for callers and FFI.
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    | _ ->
        let method = func.GetType().GetMethods() |> Array.find (fun m -> m.Name = "Invoke" && m.GetParameters().Length = 1)
        method.Invoke(func, [| arg |])



let (|LitInt|_|) (expected: int) (value: obj) = if value :? int && unbox<int> value = expected then Some() else None
let (|LitBool|_|) (expected: bool) (value: obj) = if value :? bool && unbox<bool> value = expected then Some() else None
let events = ResizeArray<int>()
let injectedFailure = System.InvalidOperationException("adt-multi injected typed failure")
let track : obj = box (fun (label: obj) -> box (fun (value: obj) -> events.Add(unbox<int> label); value))
let AdtMultiConsumer_trackColor = track
let AdtMultiConsumer_trackTree = track
let AdtMultiConsumer_trackInt = track
let Partial_Unsafe_unsafePartial : obj = box (fun (value: obj) -> sharpurs_apply value (box ()))
let Partial_Unsafe__unsafePartial = Partial_Unsafe_unsafePartial
let Data_Boolean_otherwise = box true


let AdtMultiExternal_readZero  = (box (fun (value: obj) -> (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box (fun (usd__unused: obj) -> (match ((unbox ((box value)))) with | LitInt 0 () -> ((box 0))))))) (box ((box Prim_undefined)))))))))))

let AdtMultiExternal_offset  = (box (fun (value: obj) -> (box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box 1))))))))

module Native =
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
    
    let AdtMultiConsumer_viaReturned  = (box AdtMulti_returned)
    
    let AdtMultiConsumer_orderedPartial  = (box (fun (tree: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_assemble))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackColor))) (box ((box 1)))))) (box ((box (AdtMulti_Onyx_adt_native)))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 2)))))) (box ((box tree)))))))))
    
    let AdtMultiConsumer_ordered  = (box (fun (tree: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_assemble))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackColor))) (box ((box 1)))))) (box ((box (AdtMulti_Onyx_adt_native)))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 2)))))) (box ((box tree))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackInt))) (box ((box 3)))))) (box ((box 7))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 4)))))) (box ((box tree)))))))))
    
    let AdtMultiConsumer_lastArgumentFailure  = (box (fun (tree: obj) -> (box (fun (bad: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_readNonEmpty))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 1)))))) (box ((box tree))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackInt))) (box ((box 2)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_readNonEmpty))) (box ((box bad)))))) (box ((box 0))))))))))))))
    
    let AdtMultiConsumer_firstArgumentFailure  = (box (fun (bad: obj) -> (box (fun (tree: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_choose))) (box ((box true)))))) (box ((box tree)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackInt))) (box ((box 1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_readNonEmpty))) (box ((box bad)))))) (box ((box 0)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 2)))))) (box ((box tree)))))))))))
    
    let AdtMultiConsumer_captured  = (box AdtMulti_insert)


module Oracle =
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
    
    let AdtMulti_unused_direct (tree: obj) (v: obj) : obj = (box tree)
    
    let AdtMulti_unused_direct_apply (tree: obj) (v: obj) : obj =
        try AdtMulti_unused_direct tree v
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_unused = (box (fun (tree: obj) -> (box (fun (v: obj) -> (AdtMulti_unused_direct tree v)))))
    
    let AdtMulti_score_direct (v: obj) (v1: obj) : obj = (match (((unbox ((box v))), (unbox ((box v1))))) with | (AdtMulti_Tipusd_Ctor, increment) -> ((box increment)) | (AdtMulti_Branchusd_Ctor(_, _, value, _), increment) -> ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment))))))))
    
    let AdtMulti_score_direct_apply (v: obj) (v1: obj) : obj =
        try AdtMulti_score_direct v v1
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_score = (box (fun (v: obj) -> (box (fun (v1: obj) -> (AdtMulti_score_direct v v1)))))
    
    let AdtMulti_shortAndSafe_direct (flag: obj) (tree: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((box ((unbox<int> (box ((AdtMulti_score_direct_apply ((box ((box tree)))) ((box ((box 0)))))))) > (unbox<int> (box ((box 0))))))) | _ -> ((box false)))
    
    let AdtMulti_shortAndSafe_direct_apply (flag: obj) (tree: obj) : obj =
        try AdtMulti_shortAndSafe_direct flag tree
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_shortAndSafe = (box (fun (flag: obj) -> (box (fun (tree: obj) -> (AdtMulti_shortAndSafe_direct flag tree)))))
    
    let AdtMulti_shortOrSafe_direct (flag: obj) (tree: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((box true)) | _ -> ((box ((unbox<int> (box ((AdtMulti_score_direct_apply ((box ((box tree)))) ((box ((box 0)))))))) > (unbox<int> (box ((box 0))))))))
    
    let AdtMulti_shortOrSafe_direct_apply (flag: obj) (tree: obj) : obj =
        try AdtMulti_shortOrSafe_direct flag tree
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_shortOrSafe = (box (fun (flag: obj) -> (box (fun (tree: obj) -> (AdtMulti_shortOrSafe_direct flag tree)))))
    
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
    
    let AdtMulti_nested_direct (v: obj) (v1: obj) : obj = (match (((unbox ((box v))), (unbox ((box v1))))) with | (AdtMulti_Branchusd_Ctor(Unbox(AdtMulti_Onyxusd_Ctor), Unbox(AdtMulti_Branchusd_Ctor(Unbox(AdtMulti_Rubyusd_Ctor), _, value, _)), _, _), increment) -> ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment))))))) | (_, increment) -> ((box increment)))
    
    let AdtMulti_nested_direct_apply (v: obj) (v1: obj) : obj =
        try AdtMulti_nested_direct v v1
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_nested = (box (fun (v: obj) -> (box (fun (v1: obj) -> (AdtMulti_nested_direct v v1)))))
    
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
    
    let rec AdtMulti_depth_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | AdtMulti_Tipusd_Ctor -> ((box 0)) | AdtMulti_Branchusd_Ctor(_, left, _, right) -> ((let b = (AdtMulti_depth_tco ((box right))) in let a = (AdtMulti_depth_tco ((box left))) in (box ((unbox<int> (box ((box 1)))) + (unbox<int> (box ((match ((unbox ((box ((unbox<int> (box ((box a)))) > (unbox<int> (box ((box b))))))))) with | LitBool true () -> ((box a)) | _ -> ((box b))))))))))))
    and AdtMulti_depth = box (fun (v: obj) -> AdtMulti_depth_tco v)
    
    
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
    
    let AdtMulti_assemble_direct (color: obj) (left: obj) (value: obj) (right: obj) : obj = (box (AdtMulti_Branch_adt_native (unbox ((box color))) (unbox ((box left))) (unbox ((box value))) (unbox ((box right)))))
    
    let AdtMulti_assemble_direct_apply (color: obj) (left: obj) (value: obj) (right: obj) : obj =
        try AdtMulti_assemble_direct color left value right
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_assemble = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (AdtMulti_assemble_direct color left value right)))))))))
    
    let AdtMulti_choose_direct (flag: obj) (left: obj) (value: obj) (right: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((AdtMulti_assemble_direct_apply ((box ((box (AdtMulti_Onyx_adt_native))))) ((box ((box left)))) ((box ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box 1))))))))) ((box ((box right)))))) | _ -> ((AdtMulti_assemble_direct_apply ((box ((box (AdtMulti_Ruby_adt_native))))) ((box ((box right)))) ((box ((box ((unbox<int> (box ((box value)))) - (unbox<int> (box ((box 1))))))))) ((box ((box left)))))))
    
    let AdtMulti_choose_direct_apply (flag: obj) (left: obj) (value: obj) (right: obj) : obj =
        try AdtMulti_choose_direct flag left value right
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_choose = (box (fun (flag: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (AdtMulti_choose_direct flag left value right)))))))))
    
    let rec AdtMulti_insert_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (value, AdtMulti_Tipusd_Ctor) -> ((AdtMulti_assemble_direct_apply ((box ((box (AdtMulti_Ruby_adt_native))))) ((box ((box (AdtMulti_Tip_adt_native))))) ((box ((box value)))) ((box ((box (AdtMulti_Tip_adt_native))))))) | (value, (AdtMulti_Branchusd_Ctor(color, left, old, right) as tree)) when (unbox (box ((unbox<int> (box ((box value)))) < (unbox<int> (box ((box old))))))) -> ((AdtMulti_assemble_direct_apply ((box ((box color)))) ((box ((AdtMulti_insert_tco ((box value)) ((box left)))))) ((box ((box old)))) ((box ((box right)))))) | (value, (AdtMulti_Branchusd_Ctor(color, left, old, right) as tree)) when (unbox (box ((unbox<int> (box ((box value)))) > (unbox<int> (box ((box old))))))) -> ((AdtMulti_assemble_direct_apply ((box ((box color)))) ((box ((box left)))) ((box ((box old)))) ((box ((AdtMulti_insert_tco ((box value)) ((box right)))))))) | (value, (AdtMulti_Branchusd_Ctor(color, left, old, right) as tree)) when (unbox (box Data_Boolean_otherwise)) -> ((box tree))))
    and AdtMulti_insert = box (fun (v: obj) ->  (fun (v1: obj) -> AdtMulti_insert_tco v v1))
    
    
    let AdtMulti_partial  = (sharpurs_apply (box ((box AdtMulti_assemble))) (box ((box (AdtMulti_Onyx_adt_native)))))
    
    let rec AdtMulti_shift_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (_, AdtMulti_Tipusd_Ctor) -> ((box (AdtMulti_Tip_adt_native))) | (amount, AdtMulti_Branchusd_Ctor(color, left, value, right)) -> ((AdtMulti_assemble_direct_apply ((box ((box color)))) ((box ((AdtMulti_shift_tco ((box amount)) ((box left)))))) ((box ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box amount))))))))) ((box ((AdtMulti_shift_tco ((box amount)) ((box right))))))))))
    and AdtMulti_shift = box (fun (v: obj) ->  (fun (v1: obj) -> AdtMulti_shift_tco v v1))
    
    
    let AdtMulti_argumentOrder_direct (first: obj) (second: obj) : obj = (AdtMulti_combine_direct_apply ((box ((box first)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box first)))) ((box ((box 0)))))))) ((box ((box second)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box second)))) ((box ((box 0)))))))))
    
    let AdtMulti_argumentOrder_direct_apply (first: obj) (second: obj) : obj =
        try AdtMulti_argumentOrder_direct first second
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_argumentOrder = (box (fun (first: obj) -> (box (fun (second: obj) -> (AdtMulti_argumentOrder_direct first second)))))
    
    let AdtMulti_argumentNative_direct (first: obj) (second: obj) : obj = (AdtMulti_score_direct_apply ((box ((box second)))) ((box ((AdtMulti_score_direct_apply ((box ((box first)))) ((box ((box 0)))))))))
    
    let AdtMulti_argumentNative_direct_apply (first: obj) (second: obj) : obj =
        try AdtMulti_argumentNative_direct first second
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_argumentNative = (box (fun (first: obj) -> (box (fun (second: obj) -> (AdtMulti_argumentNative_direct first second)))))
    
    let AdtMulti_argumentFailure_direct (outer: obj) (inner: obj) : obj = (AdtMulti_readNonEmpty_direct_apply ((box ((box outer)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box inner)))) ((box ((box 0)))))))))
    
    let AdtMulti_argumentFailure_direct_apply (outer: obj) (inner: obj) : obj =
        try AdtMulti_argumentFailure_direct outer inner
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_argumentFailure = (box (fun (outer: obj) -> (box (fun (inner: obj) -> (AdtMulti_argumentFailure_direct outer inner)))))
    
    let AdtMultiConsumer_viaReturned  = (box AdtMulti_returned)
    
    let AdtMultiConsumer_orderedPartial  = (box (fun (tree: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_assemble))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackColor))) (box ((box 1)))))) (box ((box (AdtMulti_Onyx_adt_native)))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 2)))))) (box ((box tree)))))))))
    
    let AdtMultiConsumer_ordered  = (box (fun (tree: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_assemble))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackColor))) (box ((box 1)))))) (box ((box (AdtMulti_Onyx_adt_native)))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 2)))))) (box ((box tree))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackInt))) (box ((box 3)))))) (box ((box 7))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 4)))))) (box ((box tree)))))))))
    
    let AdtMultiConsumer_lastArgumentFailure  = (box (fun (tree: obj) -> (box (fun (bad: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_readNonEmpty))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 1)))))) (box ((box tree))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackInt))) (box ((box 2)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_readNonEmpty))) (box ((box bad)))))) (box ((box 0))))))))))))))
    
    let AdtMultiConsumer_firstArgumentFailure  = (box (fun (bad: obj) -> (box (fun (tree: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_choose))) (box ((box true)))))) (box ((box tree)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackInt))) (box ((box 1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMulti_readNonEmpty))) (box ((box bad)))))) (box ((box 0)))))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box AdtMultiConsumer_trackTree))) (box ((box 2)))))) (box ((box tree)))))))))))
    
    let AdtMultiConsumer_captured  = (box AdtMulti_insert)


module InjectedNative =
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
    
    let AdtMulti_score_adt_native (sharpurs_adt_local_0: AdtMulti_Tree) (sharpurs_adt_local_1: int) : int = events.Add(77); raise injectedFailure
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


module InjectedOracle =
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
    
    let AdtMulti_unused_direct (tree: obj) (v: obj) : obj = (box tree)
    
    let AdtMulti_unused_direct_apply (tree: obj) (v: obj) : obj =
        try AdtMulti_unused_direct tree v
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_unused = (box (fun (tree: obj) -> (box (fun (v: obj) -> (AdtMulti_unused_direct tree v)))))
    
    let AdtMulti_score_direct (v: obj) (v1: obj) : obj = events.Add(77); raise injectedFailure
    
    let AdtMulti_score_direct_apply (v: obj) (v1: obj) : obj =
        try AdtMulti_score_direct v v1
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_score = (box (fun (v: obj) -> (box (fun (v1: obj) -> (AdtMulti_score_direct v v1)))))
    
    let AdtMulti_shortAndSafe_direct (flag: obj) (tree: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((box ((unbox<int> (box ((AdtMulti_score_direct_apply ((box ((box tree)))) ((box ((box 0)))))))) > (unbox<int> (box ((box 0))))))) | _ -> ((box false)))
    
    let AdtMulti_shortAndSafe_direct_apply (flag: obj) (tree: obj) : obj =
        try AdtMulti_shortAndSafe_direct flag tree
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_shortAndSafe = (box (fun (flag: obj) -> (box (fun (tree: obj) -> (AdtMulti_shortAndSafe_direct flag tree)))))
    
    let AdtMulti_shortOrSafe_direct (flag: obj) (tree: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((box true)) | _ -> ((box ((unbox<int> (box ((AdtMulti_score_direct_apply ((box ((box tree)))) ((box ((box 0)))))))) > (unbox<int> (box ((box 0))))))))
    
    let AdtMulti_shortOrSafe_direct_apply (flag: obj) (tree: obj) : obj =
        try AdtMulti_shortOrSafe_direct flag tree
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_shortOrSafe = (box (fun (flag: obj) -> (box (fun (tree: obj) -> (AdtMulti_shortOrSafe_direct flag tree)))))
    
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
    
    let AdtMulti_nested_direct (v: obj) (v1: obj) : obj = (match (((unbox ((box v))), (unbox ((box v1))))) with | (AdtMulti_Branchusd_Ctor(Unbox(AdtMulti_Onyxusd_Ctor), Unbox(AdtMulti_Branchusd_Ctor(Unbox(AdtMulti_Rubyusd_Ctor), _, value, _)), _, _), increment) -> ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box increment))))))) | (_, increment) -> ((box increment)))
    
    let AdtMulti_nested_direct_apply (v: obj) (v1: obj) : obj =
        try AdtMulti_nested_direct v v1
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_nested = (box (fun (v: obj) -> (box (fun (v1: obj) -> (AdtMulti_nested_direct v v1)))))
    
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
    
    let rec AdtMulti_depth_tco (v: obj) : obj = ((match ((unbox ((box v)))) with | AdtMulti_Tipusd_Ctor -> ((box 0)) | AdtMulti_Branchusd_Ctor(_, left, _, right) -> ((let b = (AdtMulti_depth_tco ((box right))) in let a = (AdtMulti_depth_tco ((box left))) in (box ((unbox<int> (box ((box 1)))) + (unbox<int> (box ((match ((unbox ((box ((unbox<int> (box ((box a)))) > (unbox<int> (box ((box b))))))))) with | LitBool true () -> ((box a)) | _ -> ((box b))))))))))))
    and AdtMulti_depth = box (fun (v: obj) -> AdtMulti_depth_tco v)
    
    
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
    
    let AdtMulti_assemble_direct (color: obj) (left: obj) (value: obj) (right: obj) : obj = (box (AdtMulti_Branch_adt_native (unbox ((box color))) (unbox ((box left))) (unbox ((box value))) (unbox ((box right)))))
    
    let AdtMulti_assemble_direct_apply (color: obj) (left: obj) (value: obj) (right: obj) : obj =
        try AdtMulti_assemble_direct color left value right
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_assemble = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (AdtMulti_assemble_direct color left value right)))))))))
    
    let AdtMulti_choose_direct (flag: obj) (left: obj) (value: obj) (right: obj) : obj = (match ((unbox ((box flag)))) with | LitBool true () -> ((AdtMulti_assemble_direct_apply ((box ((box (AdtMulti_Onyx_adt_native))))) ((box ((box left)))) ((box ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box 1))))))))) ((box ((box right)))))) | _ -> ((AdtMulti_assemble_direct_apply ((box ((box (AdtMulti_Ruby_adt_native))))) ((box ((box right)))) ((box ((box ((unbox<int> (box ((box value)))) - (unbox<int> (box ((box 1))))))))) ((box ((box left)))))))
    
    let AdtMulti_choose_direct_apply (flag: obj) (left: obj) (value: obj) (right: obj) : obj =
        try AdtMulti_choose_direct flag left value right
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_choose = (box (fun (flag: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (AdtMulti_choose_direct flag left value right)))))))))
    
    let rec AdtMulti_insert_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (value, AdtMulti_Tipusd_Ctor) -> ((AdtMulti_assemble_direct_apply ((box ((box (AdtMulti_Ruby_adt_native))))) ((box ((box (AdtMulti_Tip_adt_native))))) ((box ((box value)))) ((box ((box (AdtMulti_Tip_adt_native))))))) | (value, (AdtMulti_Branchusd_Ctor(color, left, old, right) as tree)) when (unbox (box ((unbox<int> (box ((box value)))) < (unbox<int> (box ((box old))))))) -> ((AdtMulti_assemble_direct_apply ((box ((box color)))) ((box ((AdtMulti_insert_tco ((box value)) ((box left)))))) ((box ((box old)))) ((box ((box right)))))) | (value, (AdtMulti_Branchusd_Ctor(color, left, old, right) as tree)) when (unbox (box ((unbox<int> (box ((box value)))) > (unbox<int> (box ((box old))))))) -> ((AdtMulti_assemble_direct_apply ((box ((box color)))) ((box ((box left)))) ((box ((box old)))) ((box ((AdtMulti_insert_tco ((box value)) ((box right)))))))) | (value, (AdtMulti_Branchusd_Ctor(color, left, old, right) as tree)) when (unbox (box Data_Boolean_otherwise)) -> ((box tree))))
    and AdtMulti_insert = box (fun (v: obj) ->  (fun (v1: obj) -> AdtMulti_insert_tco v v1))
    
    
    let AdtMulti_partial  = (sharpurs_apply (box ((box AdtMulti_assemble))) (box ((box (AdtMulti_Onyx_adt_native)))))
    
    let rec AdtMulti_shift_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (_, AdtMulti_Tipusd_Ctor) -> ((box (AdtMulti_Tip_adt_native))) | (amount, AdtMulti_Branchusd_Ctor(color, left, value, right)) -> ((AdtMulti_assemble_direct_apply ((box ((box color)))) ((box ((AdtMulti_shift_tco ((box amount)) ((box left)))))) ((box ((box ((unbox<int> (box ((box value)))) + (unbox<int> (box ((box amount))))))))) ((box ((AdtMulti_shift_tco ((box amount)) ((box right))))))))))
    and AdtMulti_shift = box (fun (v: obj) ->  (fun (v1: obj) -> AdtMulti_shift_tco v v1))
    
    
    let AdtMulti_argumentOrder_direct (first: obj) (second: obj) : obj = (AdtMulti_combine_direct_apply ((box ((box first)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box first)))) ((box ((box 0)))))))) ((box ((box second)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box second)))) ((box ((box 0)))))))))
    
    let AdtMulti_argumentOrder_direct_apply (first: obj) (second: obj) : obj =
        try AdtMulti_argumentOrder_direct first second
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_argumentOrder = (box (fun (first: obj) -> (box (fun (second: obj) -> (AdtMulti_argumentOrder_direct first second)))))
    
    let AdtMulti_argumentNative_direct (first: obj) (second: obj) : obj = (AdtMulti_score_direct_apply ((box ((box second)))) ((box ((AdtMulti_score_direct_apply ((box ((box first)))) ((box ((box 0)))))))))
    
    let AdtMulti_argumentNative_direct_apply (first: obj) (second: obj) : obj =
        try AdtMulti_argumentNative_direct first second
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_argumentNative = (box (fun (first: obj) -> (box (fun (second: obj) -> (AdtMulti_argumentNative_direct first second)))))
    
    let AdtMulti_argumentFailure_direct (outer: obj) (inner: obj) : obj = (AdtMulti_readNonEmpty_direct_apply ((box ((box outer)))) ((box ((AdtMulti_readNonEmpty_direct_apply ((box ((box inner)))) ((box ((box 0)))))))))
    
    let AdtMulti_argumentFailure_direct_apply (outer: obj) (inner: obj) : obj =
        try AdtMulti_argumentFailure_direct outer inner
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let AdtMulti_argumentFailure = (box (fun (outer: obj) -> (box (fun (inner: obj) -> (AdtMulti_argumentFailure_direct outer inner)))))



let mutable checks = 0
let check label condition = if not condition then failwith label else checks <- checks + 1
let apply = sharpurs_apply
let call2 fn a b = apply (apply fn a) b
let call4 fn a b c d = apply (apply (apply (apply fn a) b) c) d
let node n = call4 Native.AdtMulti_assemble Native.AdtMulti_Ruby Native.AdtMulti_Tip (box n) Native.AdtMulti_Tip
let reference n = call4 Oracle.AdtMulti_assemble Oracle.AdtMulti_Ruby Oracle.AdtMulti_Tip (box n) Oracle.AdtMulti_Tip
let rec nativeShape (tree: Native.AdtMulti_Tree) =
    match tree with
    | Native.AdtMulti_Tipusd_Ctor -> "."
    | Native.AdtMulti_Branchusd_Ctor(color, left, value, right) -> sprintf "(%A,%s,%d,%s)" color (nativeShape left) value (nativeShape right)
let rec oracleShape (tree: Oracle.AdtMulti_Tree) =
    match tree with
    | Oracle.AdtMulti_Tipusd_Ctor -> "."
    | Oracle.AdtMulti_Branchusd_Ctor(color, left, value, right) -> sprintf "(%A,%s,%d,%s)" color (oracleShape left) value (oracleShape right)
let sameTree actual expected = nativeShape (unbox actual) = oracleShape (unbox expected)
let build fn empty values = List.fold (fun tree value -> call2 fn (box value) tree) empty values
let samples = [[1..40]; [40..-1..1]; [7; 2; 11; 0; 4; 9; 13; 2; 7]; [System.Int32.MinValue; 0; System.Int32.MaxValue; -1; 1]]
for values in samples do
    let tree = build Native.AdtMulti_insert Native.AdtMulti_Tip values
    let old = build Oracle.AdtMulti_insert Oracle.AdtMulti_Tip values
    check "recursive insert entire structure matches oracle" (sameTree tree old)
    check "native depth agrees" (apply Native.AdtMulti_depth tree = apply Oracle.AdtMulti_depth old)
    for value in values do
        check "duplicate insertion preserves complete tree" (sameTree (call2 Native.AdtMulti_insert (box value) tree) old)
    check "duplicate root insertion preserves root identity" (System.Object.ReferenceEquals(call2 Native.AdtMulti_insert (box values.Head) tree, tree))
    for amount in [System.Int32.MinValue; -1; 0; 1; System.Int32.MaxValue] do
        check "recursive shift, Int wrap and dependency calls agree" (sameTree (call2 Native.AdtMulti_shift (box amount) tree) (call2 Oracle.AdtMulti_shift (box amount) old))
    check "input remains immutable after transformations" (sameTree tree old)
for value in [System.Int32.MinValue; -1; 0; 1; System.Int32.MaxValue] do
    let tree = node value
    let old = reference value
    for increment in [System.Int32.MinValue; -1; 0; 1; System.Int32.MaxValue] do
        for fn, baseline in [Native.AdtMulti_readNonEmpty, Oracle.AdtMulti_readNonEmpty; Native.AdtMulti_bodyFailure, Oracle.AdtMulti_bodyFailure;
                             Native.AdtMulti_importedTotal, Oracle.AdtMulti_importedTotal] do
            check "two argument native Int result and overflow agree" (call2 fn tree (box increment) = call2 baseline old (box increment))
    for flag in [false; true] do
        check "four argument Boolean branch agrees" (sameTree (call4 Native.AdtMulti_choose (box flag) tree (box value) tree) (call4 Oracle.AdtMulti_choose (box flag) old (box value) old))
    check "unused argument keeps original ADT identity" (System.Object.ReferenceEquals(call2 Native.AdtMulti_unused tree (box 99), tree))
let child = node 7
let original = reference 7
for fn, empty, leaf, encode in [Native.AdtMultiConsumer_ordered, Native.AdtMulti_Tip, child, (fun value -> nativeShape (unbox value));
                               Oracle.AdtMultiConsumer_ordered, Oracle.AdtMulti_Tip, original, (fun value -> oracleShape (unbox value))] do
    events.Clear()
    let value = apply fn leaf
    check "consumer arguments evaluated once in order" (List.ofSeq events = [1; 2; 3; 4])
    check "ordered construction actually creates node" (encode value <> encode empty)
for fn, leaf, empty in [Native.AdtMultiConsumer_orderedPartial, child, Native.AdtMulti_Tip;
                        Oracle.AdtMultiConsumer_orderedPartial, original, Oracle.AdtMulti_Tip] do
    events.Clear()
    let partial = apply fn leaf
    check "partial supplied arguments evaluate eagerly" (List.ofSeq events = [1; 2])
    let captured = apply partial (box 9)
    apply captured leaf |> ignore
    apply captured empty |> ignore
    check "reused partial captures arguments without reevaluation" (List.ofSeq events = [1; 2])
let first = apply Native.AdtMulti_assemble Native.AdtMulti_Onyx
let second = apply first child
let third = apply second (box 9)
for right in [child; Native.AdtMulti_Tip] do
    let result = apply third right |> unbox<Native.AdtMulti_Tree>
    match result with
    | Native.AdtMulti_Branchusd_Ctor(_, left, value, actualRight) ->
        check "partial wrapper captures native child identity" (System.Object.ReferenceEquals(left, child))
        check "partial wrapper captures Int" (value = 9)
        check "partial wrapper can be reused with new argument" (System.Object.ReferenceEquals(actualRight, right))
    | _ -> failwith "expected branch"
let capturedInsert = apply Native.AdtMultiConsumer_captured (box 13)
check "function value preserves first use" (sameTree (apply capturedInsert Native.AdtMulti_Tip) (call2 Oracle.AdtMulti_insert (box 13) Oracle.AdtMulti_Tip))
check "function value preserves reuse" (sameTree (apply capturedInsert child) (call2 Oracle.AdtMulti_insert (box 13) original))
let capture action = try action() |> ignore; None with ex -> Some ex
let rec signature (ex: System.Exception) =
    match ex with
    | :? System.Reflection.TargetInvocationException as wrapper -> let depth, cause = signature wrapper.InnerException in depth + 1, cause
    | cause -> 0, cause.GetType().FullName + ": " + cause.Message
let failure action = match capture action with Some ex -> signature ex | None -> failwith "expected fixture failure"
for fn, baseline, bypass, evaluate in [Native.AdtMulti_shortAnd, Oracle.AdtMulti_shortAnd, false, true;
                                      Native.AdtMulti_shortOr, Oracle.AdtMulti_shortOr, true, false] do
    check "short-circuit skips failing right hand side" (call2 fn (box bypass) Native.AdtMulti_Tip = call2 baseline (box bypass) Oracle.AdtMulti_Tip)
    let actual = failure (fun () -> call2 fn (box evaluate) Native.AdtMulti_Tip)
    let expected = failure (fun () -> call2 baseline (box evaluate) Oracle.AdtMulti_Tip)
    if actual <> expected then printfn "short-circuit exception mismatch: native=%A oracle=%A" actual expected
    check "short-circuit evaluates necessary right hand side with preserved envelope" (actual = expected)
    check "right hand side success" (call2 fn (box evaluate) child = call2 baseline (box evaluate) original)
for fn, baseline in [Native.AdtMulti_shortAndSafe, Oracle.AdtMulti_shortAndSafe; Native.AdtMulti_shortOrSafe, Oracle.AdtMulti_shortOrSafe] do
    for flag in [false; true] do
        for tree, old in [Native.AdtMulti_Tip, Oracle.AdtMulti_Tip; child, original] do
            check "native lazy Boolean primitive matches source branches" (call2 fn (box flag) tree = call2 baseline (box flag) old)
for tree, old in [Native.AdtMulti_Tip, Oracle.AdtMulti_Tip; child, original;
                  call4 Native.AdtMulti_assemble Native.AdtMulti_Onyx child (box 11) Native.AdtMulti_Tip,
                  call4 Oracle.AdtMulti_assemble Oracle.AdtMulti_Onyx original (box 11) Oracle.AdtMulti_Tip] do
    check "nested tags short-circuit constructor accessors" (call2 Native.AdtMulti_nested tree (box 3) = call2 Oracle.AdtMulti_nested old (box 3))
for native, generic, nativeArgs, genericArgs in [
    Native.AdtMulti_readNonEmpty, Oracle.AdtMulti_readNonEmpty, [Native.AdtMulti_Tip; box 0], [Oracle.AdtMulti_Tip; box 0]
    Native.AdtMulti_bodyFailure, Oracle.AdtMulti_bodyFailure, [Native.AdtMulti_Tip; box 0], [Oracle.AdtMulti_Tip; box 0]
    Native.AdtMulti_readNonEmptyInline, Oracle.AdtMulti_readNonEmptyInline, [Native.AdtMulti_Tip; box 0], [Oracle.AdtMulti_Tip; box 0]
    Native.AdtMulti_bodyFailureInline, Oracle.AdtMulti_bodyFailureInline, [Native.AdtMulti_Tip; box 0], [Oracle.AdtMulti_Tip; box 0]
    Native.AdtMulti_importedFailure, Oracle.AdtMulti_importedFailure, [Native.AdtMulti_Tip; box 1], [Oracle.AdtMulti_Tip; box 1]
    Native.AdtMulti_throughImported, Oracle.AdtMulti_throughImported, [Native.AdtMulti_Tip; box 1], [Oracle.AdtMulti_Tip; box 1]
    Native.AdtMulti_argumentFailure, Oracle.AdtMulti_argumentFailure, [child; Native.AdtMulti_Tip], [original; Oracle.AdtMulti_Tip]
    Native.AdtMulti_argumentFailure, Oracle.AdtMulti_argumentFailure, [Native.AdtMulti_Tip; child], [Oracle.AdtMulti_Tip; original]
    Native.AdtMulti_argumentOrder, Oracle.AdtMulti_argumentOrder, [Native.AdtMulti_Tip; child], [Oracle.AdtMulti_Tip; original]
    Native.AdtMulti_argumentOrder, Oracle.AdtMulti_argumentOrder, [child; Native.AdtMulti_Tip], [original; Oracle.AdtMulti_Tip]
] do
    let invoke fn args () = List.fold apply fn args
    let actual = failure (invoke native nativeArgs)
    let expected = failure (invoke generic genericArgs)
    check "typed body or argument failure keeps complete generic exception envelope" (actual = expected)
    check "exception test reached actual pattern match failure" (snd actual |> fun message -> message.Contains("MatchFailureException"))
check "external partial helper successful branch" (call2 Native.AdtMulti_importedFailure Native.AdtMulti_Tip (box 0) = call2 Oracle.AdtMulti_importedFailure Oracle.AdtMulti_Tip (box 0))
for fn, baseline, expectedEvents in [Native.AdtMultiConsumer_firstArgumentFailure, Oracle.AdtMultiConsumer_firstArgumentFailure, [];
                                      Native.AdtMultiConsumer_lastArgumentFailure, Oracle.AdtMultiConsumer_lastArgumentFailure, [1]] do
    let left, right, oldLeft, oldRight =
        if System.Object.ReferenceEquals(fn, Native.AdtMultiConsumer_firstArgumentFailure) then Native.AdtMulti_Tip, child, Oracle.AdtMulti_Tip, original
        else child, Native.AdtMulti_Tip, original, Oracle.AdtMulti_Tip
    events.Clear()
    let actual = failure (fun () -> call2 fn left right)
    check "argument failure stops later argument effects" (List.ofSeq events = expectedEvents)
    events.Clear()
    let expected = failure (fun () -> call2 baseline oldLeft oldRight)
    check "consumer exception chain matches generic oracle" (actual = expected)
    check "consumer oracle has same event order" (List.ofSeq events = expectedEvents)
for fn, leaf, empty in [Native.AdtMulti_returned, child, Native.AdtMulti_Tip; Oracle.AdtMulti_returned, original, Oracle.AdtMulti_Tip] do
    let returned = apply fn leaf
    for increment in [0; 1; 9] do
        check "returned closure retains captured value" (unbox<int> (apply returned (box increment)) = 7 + increment)
    check "returned closure preserves eager source let failure" (capture (fun () -> apply fn empty) |> Option.isSome)
for fn, baseline, bypass, evaluate in [InjectedNative.AdtMulti_shortAndSafe, InjectedOracle.AdtMulti_shortAndSafe, false, true;
                                      InjectedNative.AdtMulti_shortOrSafe, InjectedOracle.AdtMulti_shortOrSafe, true, false] do
    events.Clear()
    check "injected native lazy RHS bypass matches generic source" (call2 fn (box bypass) InjectedNative.AdtMulti_Tip = call2 baseline (box bypass) InjectedOracle.AdtMulti_Tip)
    check "bypassed RHS was not evaluated in either version" (events.Count = 0)
    let actual = failure (fun () -> call2 fn (box evaluate) InjectedNative.AdtMulti_Tip)
    check "native RHS evaluated exactly once" (List.ofSeq events = [77])
    events.Clear()
    let expected = failure (fun () -> call2 baseline (box evaluate) InjectedOracle.AdtMulti_Tip)
    check "generic RHS evaluated exactly once" (List.ofSeq events = [77])
    check "native cross-call wraps same sentinel exactly as generic call" (actual = expected && fst actual = 2)
events.Clear()
let argumentException = failure (fun () -> call2 InjectedNative.AdtMulti_argumentNative InjectedNative.AdtMulti_Tip InjectedNative.AdtMulti_Tip)
check "failing native argument prevents outer callee evaluation" (List.ofSeq events = [77])
events.Clear()
let oracleArgumentException = failure (fun () -> call2 InjectedOracle.AdtMulti_argumentNative InjectedOracle.AdtMulti_Tip InjectedOracle.AdtMulti_Tip)
check "generic failing argument prevents outer callee evaluation" (List.ofSeq events = [77])
check "argument evaluates before outer invocation guard" (argumentException = oracleArgumentException && fst argumentException = 2)
printfn "adt-multi runtime: %d checks passed" checks

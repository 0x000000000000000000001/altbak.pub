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
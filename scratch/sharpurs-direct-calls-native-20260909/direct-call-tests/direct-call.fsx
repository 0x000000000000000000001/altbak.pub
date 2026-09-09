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



let (|LitBool|_|) (expected: bool) (value: obj) = if value :? bool && unbox<bool> value = expected then Some() else None
let events = ResizeArray<int>()
let failure = System.InvalidOperationException("direct-call fixture failure")
let DirectCall_track : obj = box (fun (label: obj) -> box (fun (value: obj) -> events.Add(unbox<int> label); value))
let DirectCall_explode : obj = box (fun (_: obj) -> events.Add(99); raise failure : obj)
let DirectFallback_track = DirectCall_track
let DirectFallback_explode = DirectCall_explode
let DirectRemote_remote : obj = box (fun (x: obj) -> box (fun (_: obj) -> x))


type DirectCall_Color =
  | DirectCall_Warmusd_Ctor
  | DirectCall_Coolusd_Ctor

type DirectCall_Tree =
  | DirectCall_Tipusd_Ctor
  | DirectCall_Forkusd_Ctor of obj * obj * obj * obj

let DirectCall_Warm  = (box DirectCall_Warmusd_Ctor)

let DirectCall_Cool  = (box DirectCall_Coolusd_Ctor)

let DirectCall_Tip  = (box DirectCall_Tipusd_Ctor)

let DirectCall_Fork  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (fun (usd__arg4: obj) -> (box (DirectCall_Forkusd_Ctor(usd__arg1, usd__arg2, usd__arg3, usd__arg4)))))))))

let DirectCall_unary  = (box (fun (x: obj) -> (box x)))

let DirectCall_stringPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_shadowed  = (box (fun (ordinal1: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box 4)))))) (box ((box 3)))))) (box ((box 2)))))) (box ((box 1))))))

let DirectCall_seed  = (box 17)

let DirectCall_returned  = (box (fun (flag: obj) -> (box (fun (value: obj) -> (let saved = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 90)))))) (box ((box value)))) in (box (fun (other: obj) -> (match ((unbox ((box flag)))) with | LitBool true () -> ((box saved)) | _ -> ((box other))))))))))

let DirectCall_return_direct  = (box 7)

let DirectCall_return  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_ordinal_direct (a: obj) (b: obj) (c: obj) (d: obj) : obj = (match ((unbox ((box ((unbox<int> (box ((box a)))) < (unbox<int> (box ((box b))))))))) with | LitBool true () -> ((box c)) | _ -> ((box d)))

let DirectCall_ordinal_direct_apply (a: obj) (b: obj) (c: obj) (d: obj) : obj =
    try DirectCall_ordinal_direct a b c d
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_ordinal = (box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (DirectCall_ordinal_direct a b c d)))))))))

let DirectCall_partial1  = (sharpurs_apply (box ((box DirectCall_ordinal))) (box ((box 1))))

let DirectCall_partial2  = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_ordinal))) (box ((box 1)))))) (box ((box 2))))

let DirectCall_partial3  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_ordinal))) (box ((box 1)))))) (box ((box 2)))))) (box ((box 3))))

let DirectCall_saturated  = (box (fun (x: obj) -> (DirectCall_ordinal_direct_apply ((box ((box x)))) ((box ((box 10)))) ((box ((box 20)))) ((box ((box 30)))))))

let DirectCall_orderedPartial  = (box (fun (x: obj) -> (sharpurs_apply (box ((box DirectCall_ordinal))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 1)))))) (box ((box x)))))))))

let DirectCall_ordered  = (box (fun (x: obj) -> (DirectCall_ordinal_direct_apply ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 1)))))) (box ((box x))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 2)))))) (box ((box 10))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 3)))))) (box ((box 20))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 4)))))) (box ((box 30))))))))))

let DirectCall_numberPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_localShadow  = (box (fun (x: obj) -> (let ordinal1 = (box (fun (a: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box a))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box x)))))) (box ((box 2)))))) (box ((box 3)))))) (box ((box 4)))))))

let DirectCall_imported_direct (x: obj) (y: obj) : obj = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectRemote_remote))) (box ((box x)))))) (box ((box y))))

let DirectCall_imported_direct_apply (x: obj) (y: obj) : obj =
    try DirectCall_imported_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_imported = (box (fun (x: obj) -> (box (fun (y: obj) -> (DirectCall_imported_direct x y)))))

let DirectCall_higher  = (box (fun (fn: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box fn))) (box ((box x))))))))

let DirectCall_guarded_direct (first: obj) (x: obj) (second: obj) (y: obj) : obj = (match ((unbox ((box first)))) with | LitBool true () -> ((box x)) | _ -> ((match ((unbox ((box second)))) with | LitBool true () -> ((box y)) | _ -> ((box DirectCall_seed)))))

let DirectCall_guarded_direct_apply (first: obj) (x: obj) (second: obj) (y: obj) : obj =
    try DirectCall_guarded_direct first x second y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_guarded = (box (fun (first: obj) -> (box (fun (x: obj) -> (box (fun (second: obj) -> (box (fun (y: obj) -> (DirectCall_guarded_direct first x second y)))))))))

let DirectCall_generic  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectCall_failBody_direct (x: obj) (v: obj) : obj = (sharpurs_apply (box ((box DirectCall_explode))) (box ((box x))))

let DirectCall_failBody_direct_apply (x: obj) (v: obj) : obj =
    try DirectCall_failBody_direct x v
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_failBody = (box (fun (x: obj) -> (box (fun (v: obj) -> (DirectCall_failBody_direct x v)))))

let DirectCall_captured_direct (x: obj) (y: obj) : obj = (DirectCall_ordinal_direct_apply ((box ((box x)))) ((box ((box y)))) ((box ((box DirectCall_seed)))) ((box ((box 29)))))

let DirectCall_captured_direct_apply (x: obj) (y: obj) : obj =
    try DirectCall_captured_direct x y
    with ex -> raise (System.Reflection.TargetInvocationException(ex))

let DirectCall_captured = (box (fun (x: obj) -> (box (fun (y: obj) -> (DirectCall_captured_direct x y)))))

let DirectCall_bodyCall  = (box (fun (x: obj) -> (DirectCall_failBody_direct_apply ((box ((box x)))) ((box ((box 0)))))))

let DirectCall_asValue  = (box DirectCall_ordinal)

let DirectCall_arrange  = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (box (DirectCall_Forkusd_Ctor((box color), (box left), (box value), (box right))))))))))))

let DirectCall_argumentCall  = (box (fun (x: obj) -> (DirectCall_ordinal_direct_apply ((box ((sharpurs_apply (box ((box DirectCall_explode))) (box ((box x))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 2)))))) (box ((box 10))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 3)))))) (box ((box 20))))))) ((box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectCall_track))) (box ((box 4)))))) (box ((box 30))))))))))

type DirectFallback_Color =
  | DirectFallback_Warmusd_Ctor
  | DirectFallback_Coolusd_Ctor

type DirectFallback_Tree =
  | DirectFallback_Tipusd_Ctor
  | DirectFallback_Forkusd_Ctor of obj * obj * obj * obj

let DirectFallback_Warm  = (box DirectFallback_Warmusd_Ctor)

let DirectFallback_Cool  = (box DirectFallback_Coolusd_Ctor)

let DirectFallback_Tip  = (box DirectFallback_Tipusd_Ctor)

let DirectFallback_Fork  = (box ((fun (usd__arg1: obj) -> (fun (usd__arg2: obj) -> (fun (usd__arg3: obj) -> (fun (usd__arg4: obj) -> (box (DirectFallback_Forkusd_Ctor(usd__arg1, usd__arg2, usd__arg3, usd__arg4)))))))))

let DirectFallback_unary  = (box (fun (x: obj) -> (box x)))

let DirectFallback_stringPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_shadowed  = (box (fun (ordinal1: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box 4)))))) (box ((box 3)))))) (box ((box 2)))))) (box ((box 1))))))

let DirectFallback_seed  = (box 17)

let DirectFallback_returned  = (box (fun (flag: obj) -> (box (fun (value: obj) -> (let saved = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 90)))))) (box ((box value)))) in (box (fun (other: obj) -> (match ((unbox ((box flag)))) with | LitBool true () -> ((box saved)) | _ -> ((box other))))))))))

let DirectFallback_return_direct  = (box 7)

let DirectFallback_return  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_ordinal  = (box (fun (a: obj) -> (box (fun (b: obj) -> (box (fun (c: obj) -> (box (fun (d: obj) -> (match ((unbox ((box ((unbox<int> (box ((box a)))) < (unbox<int> (box ((box b))))))))) with | LitBool true () -> ((box c)) | _ -> ((box d)))))))))))

let DirectFallback_partial1  = (sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box 1))))

let DirectFallback_partial2  = (sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box 1)))))) (box ((box 2))))

let DirectFallback_partial3  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box 1)))))) (box ((box 2)))))) (box ((box 3))))

let DirectFallback_saturated  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box x)))))) (box ((box 10)))))) (box ((box 20)))))) (box ((box 30))))))

let DirectFallback_orderedPartial  = (box (fun (x: obj) -> (sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 1)))))) (box ((box x)))))))))

let DirectFallback_ordered  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 1)))))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 2)))))) (box ((box 10))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 3)))))) (box ((box 20))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 4)))))) (box ((box 30)))))))))

let DirectFallback_numberPair  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_localShadow  = (box (fun (x: obj) -> (let ordinal1 = (box (fun (a: obj) -> (box (fun (v: obj) -> (box (fun (v1: obj) -> (box (fun (v2: obj) -> (box a))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box ordinal1))) (box ((box x)))))) (box ((box 2)))))) (box ((box 3)))))) (box ((box 4)))))))

let DirectFallback_imported  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box DirectRemote_remote))) (box ((box x)))))) (box ((box y))))))))

let DirectFallback_higher  = (box (fun (fn: obj) -> (box (fun (x: obj) -> (sharpurs_apply (box ((box fn))) (box ((box x))))))))

let DirectFallback_guarded  = (box (fun (first: obj) -> (box (fun (x: obj) -> (box (fun (second: obj) -> (box (fun (y: obj) -> (match ((unbox ((box first)))) with | LitBool true () -> ((box x)) | _ -> ((match ((unbox ((box second)))) with | LitBool true () -> ((box y)) | _ -> ((box DirectFallback_seed)))))))))))))

let DirectFallback_generic  = (box (fun (x: obj) -> (box (fun (v: obj) -> (box x)))))

let DirectFallback_failBody  = (box (fun (x: obj) -> (box (fun (v: obj) -> (sharpurs_apply (box ((box DirectFallback_explode))) (box ((box x))))))))

let DirectFallback_captured  = (box (fun (x: obj) -> (box (fun (y: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((box x)))))) (box ((box y)))))) (box ((box DirectFallback_seed)))))) (box ((box 29))))))))

let DirectFallback_bodyCall  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_failBody))) (box ((box x)))))) (box ((box 0))))))

let DirectFallback_asValue  = (box DirectFallback_ordinal)

let DirectFallback_arrange  = (box (fun (color: obj) -> (box (fun (left: obj) -> (box (fun (value: obj) -> (box (fun (right: obj) -> (box (DirectFallback_Forkusd_Ctor((box color), (box left), (box value), (box right))))))))))))

let DirectFallback_argumentCall  = (box (fun (x: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_ordinal))) (box ((sharpurs_apply (box ((box DirectFallback_explode))) (box ((box x))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 2)))))) (box ((box 10))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 3)))))) (box ((box 20))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box DirectFallback_track))) (box ((box 4)))))) (box ((box 30)))))))))


let mutable checks = 0
let check label ok = if not ok then failwith label else checks <- checks + 1
let apply = sharpurs_apply
let call2 fn x y = apply (apply fn (box x)) (box y)
let call4 fn a b c d = apply (apply (apply (apply fn (box a)) (box b)) (box c)) (box d)
let asInt value = unbox<int> value
for a in [System.Int32.MinValue; -1; 0; 1; 9; 10; System.Int32.MaxValue] do
    for b in [System.Int32.MinValue; 0; 10; System.Int32.MaxValue] do
        check "public wrapper matches generic oracle" (call4 DirectCall_ordinal a b 20 30 = call4 DirectFallback_ordinal a b 20 30)
        check "raw direct method matches public wrapper" (DirectCall_ordinal_direct (box a) (box b) (box 20) (box 30) = call4 DirectCall_ordinal a b 20 30)
        check "captured top-level constant" (call2 DirectCall_captured a b = call2 DirectFallback_captured a b)
    check "saturated caller matches generic oracle" (apply DirectCall_saturated (box a) = apply DirectFallback_saturated (box a))
for first in [false; true] do
    for second in [false; true] do
        check "Boolean arguments and captured value" (call4 DirectCall_guarded first 11 second 22 = call4 DirectFallback_guarded first 11 second 22)
for fn in [DirectCall_ordered; DirectFallback_ordered] do
    events.Clear()
    check "ordered result" (asInt (apply fn (box 3)) = 20)
    check "all arguments evaluated once in source order" (List.ofSeq events = [1; 2; 3; 4])
let one = apply DirectCall_ordinal (box 1)
let two = apply one (box 2)
let three = apply two (box 3)
check "one captured argument reusable first" (asInt (call2 (apply one (box 2)) 3 4) = 3)
check "one captured argument reusable second" (asInt (call2 (apply one (box 0)) 5 6) = 6)
check "two captured arguments reusable first" (asInt (call2 two 7 8) = 7)
check "two captured arguments reusable second" (asInt (call2 two 9 10) = 9)
check "three captured arguments reusable first" (asInt (apply three (box 40)) = 3)
check "three captured arguments reusable second" (asInt (apply three (box 50)) = 3)
check "top-level one argument partial" (asInt (call2 (apply DirectCall_partial1 (box 2)) 3 4) = 3)
check "top-level two argument partial" (asInt (call2 DirectCall_partial2 7 8) = 7)
check "top-level three argument partial" (asInt (apply DirectCall_partial3 (box 8)) = 3)
check "function as a value" (asInt (call4 DirectCall_asValue 1 2 3 4) = 3)
events.Clear()
let orderedOne = apply DirectCall_orderedPartial (box 1)
check "partial supplied argument evaluates eagerly" (List.ofSeq events = [1])
check "ordered partial first use" (asInt (call2 (apply orderedOne (box 2)) 3 4) = 3)
check "ordered partial second use" (asInt (call2 (apply orderedOne (box 0)) 5 6) = 6)
check "partial capture does not repeat argument effects" (List.ofSeq events = [1])
let firstOfFour : obj = box (fun (a: obj) -> box (fun (_: obj) -> box (fun (_: obj) -> box (fun (_: obj) -> a))))
check "shadowed parameter remains dynamic" (asInt (apply DirectCall_shadowed firstOfFour) = 4)
check "local shadow remains local" (asInt (apply DirectCall_localShadow (box 42)) = 42)
check "imported function retains object ABI" (asInt (call2 DirectCall_imported 41 99) = 41)
for fn in [DirectCall_returned; DirectFallback_returned] do
    events.Clear()
    let returned = call2 fn true 41
    check "returned function boundary evaluates let eagerly" (List.ofSeq events = [90])
    check "returned function first use" (asInt (apply returned (box 1)) = 41)
    check "returned function reused" (asInt (apply returned (box 2)) = 41)
    check "returned function retains capture" (List.ofSeq events = [90])
let captured action = try action() |> ignore; failwith "expected failure" with ex -> ex
let rec chain (ex: System.Exception) =
    match ex with
    | :? System.Reflection.TargetInvocationException as wrapper -> 1 + chain wrapper.InnerException
    | cause when System.Object.ReferenceEquals(cause, failure) -> 0
    | cause -> failwithf "Unexpected exception: %A" cause
let publicFailure = chain (captured (fun () -> call2 DirectCall_failBody 1 2))
let fallbackFailure = chain (captured (fun () -> call2 DirectFallback_failBody 1 2))
check "public body exception chain matches generic oracle" (publicFailure = fallbackFailure && publicFailure = 2)
let directFailure = chain (captured (fun () -> apply DirectCall_bodyCall (box 1)))
let fallbackCallerFailure = chain (captured (fun () -> apply DirectFallback_bodyCall (box 1)))
check "direct body exception chain matches generic caller" (directFailure = fallbackCallerFailure && directFailure = 3)
for fn in [DirectCall_argumentCall; DirectFallback_argumentCall] do
    events.Clear()
    check "argument exception stays outside direct invocation envelope" (chain (captured (fun () -> apply fn (box 1))) = 2)
    check "argument failure stops later evaluations" (List.ofSeq events = [99])
printfn "direct-call runtime: %d checks passed" checks

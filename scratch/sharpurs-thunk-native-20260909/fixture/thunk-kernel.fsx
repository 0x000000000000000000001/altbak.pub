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



open System
let (|LitInt|_|) (expected: int) (value: obj) = if value :? int && unbox<int> value = expected then Some() else None
let (|LitBool|_|) (expected: bool) (value: obj) = if value :? bool && unbox<bool> value = expected then Some() else None
let Data_Unit_unit = box ()
let events = ResizeArray<int>()
let ThunkExternal_track : obj = box (fun (value: obj) -> events.Add(unbox<int> value); value)
let Partial_Unsafe_unsafePartial : obj = box (fun (value: obj) -> sharpurs_apply value (box ()))
let Partial_Unsafe__unsafePartial = Partial_Unsafe_unsafePartial
let binary operation : obj = box (fun (x: obj) -> box (fun (y: obj) -> operation x y))
let intAdd = binary (fun x y -> box (unbox<int> x + unbox<int> y))
let intSub = binary (fun x y -> box (unbox<int> x - unbox<int> y))
let Data_Semiring_semiringInt : obj = box (Map.ofList ["add", intAdd])
let Data_Ring_ringInt : obj = box (Map.ofList ["sub", intSub])
let Data_Semiring_semiringNumber : obj = box (Map.ofList ["add", binary (fun x y -> box (unbox<float> x + unbox<float> y))])
let Data_Semiring_add : obj = box (fun (dict: obj) -> Map.find "add" (unbox<Map<string,obj>> dict))
let Data_Ring_sub : obj = box (fun (dict: obj) -> Map.find "sub" (unbox<Map<string,obj>> dict))


let ThunkExternal_partialSeed  = (box (fun (value: obj) -> (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box (fun (usd__unused: obj) -> (match ((unbox ((box value)))) with | LitInt 0 () -> ((box 7))))))) (box ((box Prim_undefined)))))))))))

module Native =
    let rec private TypedThunks_chainBy_thunk_native (sharpurs_thunk_local_0: int) (sharpurs_thunk_local_1: int) (sharpurs_thunk_local_2: (unit -> int)) : (unit -> int) = (if (sharpurs_thunk_local_1 = (0)) then sharpurs_thunk_local_2 else (TypedThunks_chainBy_thunk_native (sharpurs_thunk_local_0) ((sharpurs_thunk_local_1 - (1))) ((fun (sharpurs_thunk_local_3: unit) -> ((sharpurs_thunk_local_2 ()) + sharpurs_thunk_local_0)))))
    
    let rec private TypedThunks_chain_thunk_native (sharpurs_thunk_local_0: int) (sharpurs_thunk_local_1: (unit -> int)) : (unit -> int) = (if (sharpurs_thunk_local_0 = (0)) then sharpurs_thunk_local_1 else (TypedThunks_chain_thunk_native ((sharpurs_thunk_local_0 - (1))) ((fun (sharpurs_thunk_local_2: unit) -> ((sharpurs_thunk_local_1 ()) + (1))))))
    
    let TypedThunks_Susp  = (box (fun (x: obj) -> (box x)))
    
    let TypedThunks_partialSeed  = (box (fun (value: obj) -> (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box (fun (usd__unused: obj) -> (match ((unbox ((box value)))) with | LitInt 0 () -> ((box 7))))))) (box ((box Prim_undefined)))))))))))
    
    let TypedThunks_force  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | callback -> ((sharpurs_apply (box ((box callback))) (box ((box Data_Unit_unit))))))))
    
    let TypedThunks_delay  = (box TypedThunks_Susp)
    
    let rec TypedThunks_numberChain_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((TypedThunks_numberChain_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((box acc))))))))) (box ((box 0.5))))))))))))))
    and TypedThunks_numberChain = box (fun (v: obj) ->  (fun (v1: obj) -> TypedThunks_numberChain_tco v v1))
    
    
    let TypedThunks_numberRun  = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_numberChain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box seed))))))))))))))))
    
    let rec TypedThunks_chainBy_tco (v: obj) (v1: obj) (v2: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))), (unbox ((box v2))))) with | (_, LitInt 0 (), acc) -> ((box acc)) | (step, n, acc) -> ((TypedThunks_chainBy_tco ((box step)) ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v3: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((box acc))))))) + (unbox<int> (box ((box step))))))))))))))))
    and TypedThunks_chainBy = box (fun (v: obj) ->  (fun (v1: obj) ->  (fun (v2: obj) -> TypedThunks_chainBy_tco v v1 v2)))
    
    
    let TypedThunks_runBy_direct (step: obj) (depth: obj) (seed: obj) : obj = (box ((TypedThunks_chainBy_thunk_native ((unbox<int> (box step))) ((unbox<int> (box depth))) ((let sharpurs_thunk_capture_0: int = (unbox<int> (box seed)) in (fun () -> sharpurs_thunk_capture_0)))) ()))
    
    let TypedThunks_runBy_direct_apply (step: obj) (depth: obj) (seed: obj) : obj =
        try TypedThunks_runBy_direct step depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_runBy = (box (fun (step: obj) -> (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_runBy_direct step depth seed)))))))
    
    let rec TypedThunks_chain_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((TypedThunks_chain_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v2: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((box acc))))))) + (unbox<int> (box ((box 1))))))))))))))))
    and TypedThunks_chain = box (fun (v: obj) ->  (fun (v1: obj) -> TypedThunks_chain_tco v v1))
    
    
    let TypedThunks_escaped  = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box seed)))))))))))))
    
    let TypedThunks_importedRun_direct (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box ThunkExternal_partialSeed))) (box ((box seed)))))))))))))))
    
    let TypedThunks_importedRun_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_importedRun_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_importedRun = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_importedRun_direct depth seed)))))
    
    let TypedThunks_opaqueRun_direct (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box ThunkExternal_track))) (box ((box seed)))))))))))))))
    
    let TypedThunks_opaqueRun_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_opaqueRun_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_opaqueRun = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_opaqueRun_direct depth seed)))))
    
    let TypedThunks_partialChain  = (sharpurs_apply (box ((box TypedThunks_chain))) (box ((box 3))))
    
    let TypedThunks_partialRun_direct (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box TypedThunks_partialSeed))) (box ((box seed)))))))))))))))
    
    let TypedThunks_partialRun_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_partialRun_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_partialRun = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_partialRun_direct depth seed)))))
    
    let TypedThunks_run_direct (depth: obj) (seed: obj) : obj = (box ((TypedThunks_chain_thunk_native ((unbox<int> (box depth))) ((let sharpurs_thunk_capture_0: int = (unbox<int> (box seed)) in (fun () -> sharpurs_thunk_capture_0)))) ()))
    
    let TypedThunks_run_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_run_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_run = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_run_direct depth seed)))))
    
    let TypedThunks_runLiteral  = (box (fun (depth: obj) -> (box ((TypedThunks_chain_thunk_native ((unbox<int> (box depth))) ((fun () -> (7)))) ()))))
    
    let TypedThunks_runTwo_direct (depth: obj) (left: obj) (right: obj) : obj = (box ((unbox<int> (box ((box ((TypedThunks_chain_thunk_native ((unbox<int> (box depth))) ((let sharpurs_thunk_capture_0: int = (unbox<int> (box left)) in (fun () -> sharpurs_thunk_capture_0)))) ()))))) + (unbox<int> (box ((box ((TypedThunks_chain_thunk_native ((unbox<int> (box depth))) ((let sharpurs_thunk_capture_0: int = (unbox<int> (box right)) in (fun () -> sharpurs_thunk_capture_0)))) ())))))))
    
    let TypedThunks_runTwo_direct_apply (depth: obj) (left: obj) (right: obj) : obj =
        try TypedThunks_runTwo_direct depth left right
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_runTwo = (box (fun (depth: obj) -> (box (fun (left: obj) -> (box (fun (right: obj) -> (TypedThunks_runTwo_direct depth left right)))))))
    
    let TypedThunks_unknown  = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((box seed)))))))))))
    
    let TypedThunks_unknownCallback  = (box (fun (depth: obj) -> (box (fun (callback: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box callback))))))))))))))
    
    let TypedThunks_visibleDelay  = (box (fun (depth: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box 4))))))))))))))


module Oracle =
    let TypedThunks_Susp  = (box (fun (x: obj) -> (box x)))
    
    let TypedThunks_partialSeed  = (box (fun (value: obj) -> (sharpurs_apply (box ((box Partial_Unsafe_unsafePartial))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((box (fun (usd__unused: obj) -> (match ((unbox ((box value)))) with | LitInt 0 () -> ((box 7))))))) (box ((box Prim_undefined)))))))))))
    
    let TypedThunks_force  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | callback -> ((sharpurs_apply (box ((box callback))) (box ((box Data_Unit_unit))))))))
    
    let TypedThunks_delay  = (box TypedThunks_Susp)
    
    let rec TypedThunks_numberChain_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((TypedThunks_numberChain_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Semiring_add))) (box ((box Data_Semiring_semiringNumber)))))) (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((box acc))))))))) (box ((box 0.5))))))))))))))
    and TypedThunks_numberChain = box (fun (v: obj) ->  (fun (v1: obj) -> TypedThunks_numberChain_tco v v1))
    
    
    let TypedThunks_numberRun  = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_numberChain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box seed))))))))))))))))
    
    let rec TypedThunks_chainBy_tco (v: obj) (v1: obj) (v2: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))), (unbox ((box v2))))) with | (_, LitInt 0 (), acc) -> ((box acc)) | (step, n, acc) -> ((TypedThunks_chainBy_tco ((box step)) ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v3: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((box acc))))))) + (unbox<int> (box ((box step))))))))))))))))
    and TypedThunks_chainBy = box (fun (v: obj) ->  (fun (v1: obj) ->  (fun (v2: obj) -> TypedThunks_chainBy_tco v v1 v2)))
    
    
    let TypedThunks_runBy_direct (step: obj) (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chainBy))) (box ((box step)))))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box seed))))))))))))
    
    let TypedThunks_runBy_direct_apply (step: obj) (depth: obj) (seed: obj) : obj =
        try TypedThunks_runBy_direct step depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_runBy = (box (fun (step: obj) -> (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_runBy_direct step depth seed)))))))
    
    let rec TypedThunks_chain_tco (v: obj) (v1: obj) : obj = ((match (((unbox ((box v))), (unbox ((box v1))))) with | (LitInt 0 (), acc) -> ((box acc)) | (n, acc) -> ((TypedThunks_chain_tco ((box ((unbox<int> (box ((box n)))) - (unbox<int> (box ((box 1))))))) ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v2: obj) -> (box ((unbox<int> (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((box acc))))))) + (unbox<int> (box ((box 1))))))))))))))))
    and TypedThunks_chain = box (fun (v: obj) ->  (fun (v1: obj) -> TypedThunks_chain_tco v v1))
    
    
    let TypedThunks_escaped  = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box seed)))))))))))))
    
    let TypedThunks_importedRun_direct (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box ThunkExternal_partialSeed))) (box ((box seed)))))))))))))))
    
    let TypedThunks_importedRun_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_importedRun_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_importedRun = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_importedRun_direct depth seed)))))
    
    let TypedThunks_opaqueRun_direct (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box ThunkExternal_track))) (box ((box seed)))))))))))))))
    
    let TypedThunks_opaqueRun_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_opaqueRun_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_opaqueRun = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_opaqueRun_direct depth seed)))))
    
    let TypedThunks_partialChain  = (sharpurs_apply (box ((box TypedThunks_chain))) (box ((box 3))))
    
    let TypedThunks_partialRun_direct (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((box TypedThunks_partialSeed))) (box ((box seed)))))))))))))))
    
    let TypedThunks_partialRun_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_partialRun_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_partialRun = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_partialRun_direct depth seed)))))
    
    let TypedThunks_run_direct (depth: obj) (seed: obj) : obj = (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box seed))))))))))))
    
    let TypedThunks_run_direct_apply (depth: obj) (seed: obj) : obj =
        try TypedThunks_run_direct depth seed
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_run = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (TypedThunks_run_direct depth seed)))))
    
    let TypedThunks_runLiteral  = (box (fun (depth: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box 7))))))))))))))
    
    let TypedThunks_runTwo_direct (depth: obj) (left: obj) (right: obj) : obj = (box ((unbox<int> (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box left))))))))))))))) + (unbox<int> (box ((sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box right)))))))))))))))))
    
    let TypedThunks_runTwo_direct_apply (depth: obj) (left: obj) (right: obj) : obj =
        try TypedThunks_runTwo_direct depth left right
        with ex -> raise (System.Reflection.TargetInvocationException(ex))
    
    let TypedThunks_runTwo = (box (fun (depth: obj) -> (box (fun (left: obj) -> (box (fun (right: obj) -> (TypedThunks_runTwo_direct depth left right)))))))
    
    let TypedThunks_unknown  = (box (fun (depth: obj) -> (box (fun (seed: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((box seed)))))))))))
    
    let TypedThunks_unknownCallback  = (box (fun (depth: obj) -> (box (fun (callback: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box callback))))))))))))))
    
    let TypedThunks_visibleDelay  = (box (fun (depth: obj) -> (sharpurs_apply (box ((box TypedThunks_force))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box TypedThunks_chain))) (box ((box depth)))))) (box ((sharpurs_apply (box ((box TypedThunks_delay))) (box ((box (fun (v: obj) -> (box 4))))))))))))))



let mutable checks = 0
let check label condition = if not condition then failwith label else checks <- checks + 1
let apply = sharpurs_apply
let call fn n seed = apply (apply fn (box n)) (box seed) |> unbox<int>
let cases = [(0, -2147483648, -2147483648); (0, -1, -1); (0, 0, 0); (0, 17, 17); (0, 2147483647, 2147483647); (1, -2147483648, -2147483647); (1, -1, 0); (1, 0, 1); (1, 17, 18); (1, 2147483647, -2147483648); (3, -2147483648, -2147483645); (3, -1, 2); (3, 0, 3); (3, 17, 20); (3, 2147483647, -2147483646); (37, -2147483648, -2147483611); (37, -1, 36); (37, 0, 37); (37, 17, 54); (37, 2147483647, -2147483612); (1000, -2147483648, -2147482648); (1000, -1, 999); (1000, 0, 1000); (1000, 17, 1017); (1000, 2147483647, -2147482649)]
for depth, seed, expected in cases do
    check "native route agrees with JS" (call Native.TypedThunks_run depth seed = expected)
    check "generic oracle agrees with JS" (call Oracle.TypedThunks_run depth seed = expected)
    check "literal seed result" (apply Native.TypedThunks_runLiteral (box depth) |> unbox<int> = depth + 7)
    let left = apply (apply Native.TypedThunks_escaped (box depth)) (box seed)
    check "escaping thunk stays callable" (apply Native.TypedThunks_force left |> unbox<int> = expected)
    check "escaping thunk reusable" (apply Native.TypedThunks_force left |> unbox<int> = expected)
    let partial = apply Native.TypedThunks_run (box depth)
    check "partial Int capture reused" (apply partial (box seed) |> unbox<int> = expected)
    check "partial captures independent" (apply partial (box 9) |> unbox<int> = depth + 9)
for step, depth, seed in [(2, 5, 7); (-3, 10, 29); (System.Int32.MaxValue, 3, 7)] do
    let callBy fn = apply (apply (apply fn (box step)) (box depth)) (box seed) |> unbox<int>
    check "additional captured step" (callBy Native.TypedThunks_runBy = callBy Oracle.TypedThunks_runBy)
for depth in [0; 1; 3; 1000] do
    let seed : obj = box (fun (_: obj) -> events.Add(31); box 11)
    events.Clear()
    check "unknown callback retains result" (call Native.TypedThunks_unknownCallback depth seed = depth + 11)
    check "unknown callback executes once" (List.ofSeq events = [31])
    let partial = apply Native.TypedThunks_chain (box depth)
    let left = apply partial seed
    check "construction is delayed" (List.ofSeq events = [31])
    for _ in [1..2] do apply Native.TypedThunks_force left |> ignore
    check "no memoization added" (List.ofSeq events = [31; 31; 31])
    events.Clear()
    check "opaque nested seed result" (call Native.TypedThunks_opaqueRun depth 12 = depth + 12)
    check "opaque seed observes exactly one call" (List.ofSeq events = [12])
let signature action =
    let rec shape (error: exn) wrappers =
        match error with
        | :? System.Reflection.TargetInvocationException as invocation when not (isNull invocation.InnerException) -> shape invocation.InnerException (wrappers + 1)
        | cause -> cause.GetType().FullName, wrappers
    let caught = try action() |> ignore; None with ex -> Some (shape ex 0)
    match caught with Some value -> value | None -> failwith "Expected exception"
for depth in [0; 1; 3] do
    check "local partial exception boundary" (signature (fun () -> call Native.TypedThunks_partialRun depth 1) = signature (fun () -> call Oracle.TypedThunks_partialRun depth 1))
    check "imported partial exception boundary" (signature (fun () -> call Native.TypedThunks_importedRun depth 1) = signature (fun () -> call Oracle.TypedThunks_importedRun depth 1))
let sentinel = InvalidOperationException("delayed callback")
for depth in [0; 1; 3; 1000] do
    let mutable throws = 0
    let seed : obj = box (fun (_: obj) -> throws <- throws + 1; raise sentinel : obj)
    let chain = apply (apply Native.TypedThunks_chain (box depth)) seed
    check "throwing seed stays delayed" (throws = 0)
    for repeat in [1..2] do
        let mutable caught = None
        try apply Native.TypedThunks_force chain |> ignore with ex -> caught <- Some ex
        let rec cause (ex: exn) count =
            match ex with
            | :? System.Reflection.TargetInvocationException as e when not (isNull e.InnerException) -> cause e.InnerException (count + 1)
            | e -> e, count
        let exceptionCause, wrappers = cause caught.Value 0
        check "deferred cause identity" (Object.ReferenceEquals(exceptionCause, sentinel))
        check "original force boundaries" (wrappers = 2 * (depth + 1))
        check "repeat force calls again" (throws = repeat)
let mutable total = 0
for _ in [1..1000] do total <- total + call Native.TypedThunks_run 1000 0
check "full million-thunk workload" (total = 1000000)
printfn "thunk-kernel runtime: %d checks passed" checks

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
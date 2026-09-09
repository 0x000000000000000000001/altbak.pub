[<AutoOpen>]
module PureScript_Effect_Aff

open System
open System.Collections.Generic

module Effect_Aff_FFI =
    open System
    open System.Threading
    
    let rec _unwrapException (e: Exception) =
        match e with
        | :? AggregateException as ae when ae.InnerExceptions.Count = 1 -> _unwrapException ae.InnerException
        | _ -> e
    
    type AffState = { Token: CancellationToken; KillError: Exception option ref; Supervisor: CancellationTokenSource option }
    type AffFn = AffState -> Async<obj>
    
    type NativeFiber(aff: AffFn) =
        let tcs = new System.Threading.Tasks.TaskCompletionSource<Result<obj, Exception>>()
        let cts = new CancellationTokenSource()
        let killError = ref None
        let mutable started = false
        
        member this.Cts = cts
        member this.Aff = aff
        member this.KillError = killError.Value
        member this.Cancel(ex: Exception) =
            killError.Value <- Some ex
            cts.Cancel()
        member this.Start() =
            lock this (fun () ->
                if not started then
                    started <- true
                    Sharpurs_Prelude.SharpursRuntime.EventLoopAdd(1)
                    Async.StartWithContinuations(
                        aff { Token = cts.Token; KillError = killError; Supervisor = None },
                        (fun res -> 
                            tcs.TrySetResult(Ok res) |> ignore
                            Sharpurs_Prelude.SharpursRuntime.EventLoopDone()
                        ),
                        (fun ex -> 
                            tcs.TrySetResult(Error (_unwrapException ex)) |> ignore
                            Sharpurs_Prelude.SharpursRuntime.EventLoopDone()
                        ),
                        (fun ex -> 
                            tcs.TrySetResult(Error (new Exception("Cancelled", _unwrapException ex))) |> ignore
                            Sharpurs_Prelude.SharpursRuntime.EventLoopDone()
                        ),
                        CancellationToken.None
                    )
            )
            
        member this.WaitAsync() = Async.AwaitTask tcs.Task
        member this.IsSuspended = not started
    
    let _applyFn (func: obj) (arg: obj) : obj =
        let method = func.GetType().GetMethods() |> Array.find (fun m -> m.Name = "Invoke" && m.GetParameters().Length = 1)
        method.Invoke(func, [| arg |])
    
    let _pure = fun (a: obj) -> 
        box (fun (ctx: AffState) -> async.Return(a))
    
    let _throwError = fun (eVal: obj) -> 
        box (fun (ctx: AffState) -> async {
            let e = eVal :?> Exception
            return raise e
        })
    
    let _catchError = fun (affVal: obj) -> fun (kVal: obj) ->
        box (fun (ctx: AffState) -> async {
            try
                let aff = affVal :?> AffFn
                return! aff ctx
            with 
            | :? Exception as e ->
                let handler = _applyFn kVal (box e) |> unbox<AffFn>
                return! handler ctx
        })
    
    
    let _map = fun (fVal: obj) -> fun (affVal: obj) ->
        box (fun (ctx: AffState) -> async {
            let aff = affVal :?> AffFn
            let! res = aff ctx
            return _applyFn fVal res
        })
    
    let _bind = fun (affVal: obj) -> fun (kVal: obj) ->
        box (fun (ctx: AffState) -> async {
            if ctx.Token.IsCancellationRequested then
                return raise (defaultArg ctx.KillError.Value (new Exception("Cancelled")))
            else
                let aff = affVal :?> AffFn
                let! res = aff ctx
                if ctx.Token.IsCancellationRequested then
                    return raise (defaultArg ctx.KillError.Value (new Exception("Cancelled")))
                else
                    let nextAff = _applyFn kVal res |> unbox<AffFn>
                    return! nextAff ctx
        })
    
    let _delay = fun (rightVal: obj) -> fun (msVal: obj) ->
        box (fun (ctx: AffState) -> async {
            let ms = msVal :?> float
            try
                let! res = Async.AwaitTask(System.Threading.Tasks.Task.Delay(int ms, ctx.Token))
                return box null
            with
            | :? System.Threading.Tasks.TaskCanceledException ->
                return raise (defaultArg ctx.KillError.Value (new Exception("Cancelled")))
        })
    
    let _liftEffect = fun (effVal: obj) ->
        box (fun (ctx: AffState) -> async {
            return _applyFn effVal null
        })
    
    let _makeFiberNative = fun (affVal: obj) ->
        box (fun (_dummy: obj) ->
            let aff = affVal :?> AffFn
            box (new NativeFiber(aff))
        )
    
    let _runFiber = fun (nfVal: obj) ->
        box (fun (_dummy: obj) ->
            let nf = nfVal :?> NativeFiber
            nf.Start()
            box null
        )
    
    let _killFiber = fun (nfVal: obj) -> fun (errVal: obj) -> fun (onErrorVal: obj) -> fun (onSuccessVal: obj) ->
        box (fun (_dummy: obj) ->
            let nf = nfVal :?> NativeFiber
            let ex = errVal :?> Exception
            nf.Cancel(ex)
            nf.Start() // In case it wasn't started
            
            Async.Start(async {
                let! res = nf.WaitAsync()
                let effect = _applyFn onSuccessVal (box null)
                _applyFn effect null |> ignore
            })
    
            box (fun (_dummy: obj) -> box null)
        )
    
    let _joinFiber = fun (nfVal: obj) -> fun (onErrorVal: obj) -> fun (onSuccessVal: obj) ->
        box (fun (_dummy: obj) ->
            let nf = nfVal :?> NativeFiber
            nf.Start() // In case it wasn't started
            
            Async.Start(async {
                let! res = nf.WaitAsync()
                match res with
                | Ok v ->
                    let effect = _applyFn onSuccessVal v
                    _applyFn effect null |> ignore
                | Error ex ->
                    let effect = _applyFn onErrorVal (box ex)
                    _applyFn effect null |> ignore
            })
    
            box (fun (_dummy: obj) -> box null)
        )
    
    let _onCompleteFiber = fun (nfVal: obj) -> fun (onCompleteVal: obj) ->
        box (fun (_dummy: obj) ->
            box (fun (_dummy: obj) -> box null) // Stub
        )
    
    let _isSuspendedFiber = fun (nfVal: obj) ->
        box (fun (_dummy: obj) ->
            let nf = nfVal :?> NativeFiber
            box nf.IsSuspended
        )
    
    let _makeAffImpl = fun (buildVal: obj) ->
        box (fun (ctx: AffState) -> async {
            let tcs = new System.Threading.Tasks.TaskCompletionSource<obj>()
            let mutable completed = false
            let lockObj = obj()
            
            let successCb (res: obj) = 
                let lockTaken = lock lockObj (fun () -> if completed then false else completed <- true; true)
                if lockTaken then tcs.TrySetResult(res) |> ignore
                box (fun (_: obj) -> box null)
    
            let errorCb (errVal: obj) =
                let lockTaken = lock lockObj (fun () -> if completed then false else completed <- true; true)
                let err = errVal :?> Exception
                if lockTaken then tcs.TrySetException(err) |> ignore
                box (fun (_: obj) -> box null)
    
            let buildFn = _applyFn buildVal (box errorCb)
            let effCanceler = _applyFn buildFn (box successCb)
            let canceler = _applyFn effCanceler null
            
            let reg = ctx.Token.Register(fun () ->
                let actualEx = defaultArg ctx.KillError.Value (new Exception("Cancelled"))
                let cancelAffObj = _applyFn canceler (box actualEx)
                printfn "cancelAffObj type: %s" (cancelAffObj.GetType().FullName)
                let cancelAff = cancelAffObj |> unbox<AffFn>
                let emptyCtx = { Token = CancellationToken.None; KillError = ref None; Supervisor = None }
                let t = Async.StartAsTask(cancelAff emptyCtx, cancellationToken = CancellationToken.None)
                t.ContinueWith(fun (t2: System.Threading.Tasks.Task<obj>) -> 
                    let lockTaken = lock lockObj (fun () -> if completed then false else completed <- true; true)
                    if lockTaken then tcs.TrySetException(actualEx) |> ignore
                ) |> ignore
            )
            
            try
                try
                    let! res = Async.AwaitTask(tcs.Task)
                    return res
                with
                | ex ->
                    return raise (_unwrapException ex)
            finally
                reg.Dispose()
        })
    
    let _forkAffNative = fun (affVal: obj) ->
        box (fun (ctx: AffState) -> async {
            let aff = affVal :?> AffFn
            let nf = new NativeFiber(aff)
            match ctx.Supervisor with
            | Some supCts -> supCts.Token.Register(fun () -> nf.Cts.Cancel()) |> ignore
            | None -> ()
            return box nf
        })
    
    let _makeSupervisedFiber = fun (affVal: obj) ->
        let f : obj -> obj = fun _dummy ->
            let aff = affVal :?> AffFn
            let supCts = new CancellationTokenSource()
            let supervisedAff (ctx: AffState) = async {
                return! aff { ctx with Supervisor = Some supCts }
            }
            let nf = new NativeFiber(supervisedAff)
            
            supCts.Token.Register(fun () -> nf.Cts.Cancel()) |> ignore
            
            let recd = Map.empty |> Map.add "supervisor" (box supCts) |> Map.add "fiber" (box nf)
            box recd
        box f
    
    let _killAll = fun (errVal: obj) -> fun (supVal: obj) -> fun (cbVal: obj) ->
        let f : obj -> obj = fun _dummy ->
            let supCts = supVal :?> CancellationTokenSource
            supCts.Cancel()
            _applyFn cbVal null |> ignore
            let canceler : obj -> obj = fun _dummy2 -> box null
            box canceler
        box f
    
    let _sequential = fun (affVal: obj) -> affVal
    
    let awaitUninterruptibly (t: System.Threading.Tasks.Task<'T>) =
        Async.FromContinuations(fun (onSuccess, onError, _) ->
            let mutable completed = false
            let lockObj = obj()
            t.ContinueWith(fun (t2: System.Threading.Tasks.Task<'T>) ->
                let lockTaken = lock lockObj (fun () -> 
                    if completed then false else completed <- true; true)
                if lockTaken then
                    if t2.IsFaulted then onError (_unwrapException t2.Exception)
                    elif t2.IsCanceled then onError (new OperationCanceledException())
                    else onSuccess t2.Result
            ) |> ignore
        )
    
    let generalBracket = fun (acquireVal: obj) -> fun (optionsVal: obj) -> fun (useVal: obj) ->
        box (fun (ctx: AffState) -> async {
            let acquireFn = acquireVal :?> AffFn
            // acquire is uninterruptible
            let emptyCtx = { Token = CancellationToken.None; KillError = ref None; Supervisor = None }
            let! resourceResult = async {
                try
                    let t = Async.StartAsTask(acquireFn emptyCtx, cancellationToken = CancellationToken.None)
                    let! res = awaitUninterruptibly t
                    return Ok res
                with
                | :? Exception as e -> 
                    return Error e
            }
            
            match resourceResult with
            | Error err -> return raise err
            | Ok resource ->
                let options = optionsVal :?> Map<string, obj>
                
                // Check if killed during acquire
                if ctx.Token.IsCancellationRequested then
                    let ex = defaultArg ctx.KillError.Value (new Exception("Cancelled"))
                    let cleanupFn = _applyFn (_applyFn options.["killed"] (box ex)) resource |> unbox<AffFn>
                    let t = Async.StartAsTask(cleanupFn emptyCtx, cancellationToken = CancellationToken.None)
                    let! _ = awaitUninterruptibly t
                    return raise ex
                else
                    let useFn = _applyFn useVal resource |> unbox<AffFn>
                    let! useResult = async {
                        try 
                            let! res = useFn ctx
                            return Ok res
                        with 
                        | :? Exception as e -> return Error e
                    }
                    
                    let cleanupAffFn = 
                        match useResult with
                        | Ok v ->
                            let cleanupFn = _applyFn (_applyFn options.["completed"] v) resource |> unbox<AffFn>
                            cleanupFn
                        | Error e ->
                            if e.Message = "Cancelled" || e :? OperationCanceledException then
                                let actualEx = defaultArg ctx.KillError.Value e
                                let cleanupFn = _applyFn (_applyFn options.["killed"] (box actualEx)) resource |> unbox<AffFn>
                                cleanupFn
                            else
                                let cleanupFn = _applyFn (_applyFn options.["failed"] (box e)) resource |> unbox<AffFn>
                                cleanupFn
                            
                    let t = Async.StartAsTask(cleanupAffFn emptyCtx, cancellationToken = CancellationToken.None)
                    let! _ = awaitUninterruptibly t
                    
                    match useResult with
                    | Ok v -> return v
                    | Error e -> return raise e
        })
    
    let _parAffMap = fun (fVal: obj) -> fun (affVal: obj) ->
        box (fun (ctx: AffState) -> async {
            let aff = affVal :?> AffFn
            let! res = aff ctx
            return _applyFn fVal res
        })
    
    let _parAffApply = fun (aff1Val: obj) -> fun (aff2Val: obj) ->
        box (fun (ctx: AffState) -> async {
            let aff1 = aff1Val :?> AffFn
            let aff2 = aff2Val :?> AffFn
            
            let! child1 = Async.StartChild(aff1 ctx)
            let! child2 = Async.StartChild(aff2 ctx)
            let! res1 = child1
            let! res2 = child2
            
            return _applyFn res1 res2
        })
    
    let _parAffAlt = fun (aff1Val: obj) -> fun (aff2Val: obj) ->
        box (fun (ctx: AffState) -> async {
            let cts = CancellationTokenSource.CreateLinkedTokenSource(ctx.Token)
            let token = { Token = cts.Token; KillError = ctx.KillError; Supervisor = None }
            let aff1 = aff1Val :?> AffFn
            let aff2 = aff2Val :?> AffFn
            
            let t1 = Async.StartAsTask(aff1 token, cancellationToken = CancellationToken.None)
            let t2 = Async.StartAsTask(aff2 token, cancellationToken = CancellationToken.None)
            
            let! firstCompleted = System.Threading.Tasks.Task.WhenAny(t1, t2) |> Async.AwaitTask
            let otherTask = if firstCompleted = t1 then t2 else t1
            
            if firstCompleted.Status = System.Threading.Tasks.TaskStatus.RanToCompletion then
                if not otherTask.IsCompleted then cts.Cancel()
                try
                    let! _ = Async.AwaitTask(otherTask)
                    ()
                with | _ -> ()
                return firstCompleted.Result
            else
                try
                    let! res = Async.AwaitTask(otherTask)
                    return res
                with
                | _ -> 
                    return raise (_unwrapException firstCompleted.Exception)
        })
    

let Effect_Aff__bind = box (Effect_Aff_FFI.``_bind``)
let Effect_Aff__catchError = box (Effect_Aff_FFI.``_catchError``)
let Effect_Aff__delay = box (Effect_Aff_FFI.``_delay``)
let Effect_Aff__forkAffNative = box (Effect_Aff_FFI.``_forkAffNative``)
let Effect_Aff__isSuspendedFiber = box (Effect_Aff_FFI.``_isSuspendedFiber``)
let Effect_Aff__joinFiber = box (Effect_Aff_FFI.``_joinFiber``)
let Effect_Aff__killAll = box (Effect_Aff_FFI.``_killAll``)
let Effect_Aff__killFiber = box (Effect_Aff_FFI.``_killFiber``)
let Effect_Aff__liftEffect = box (Effect_Aff_FFI.``_liftEffect``)
let Effect_Aff__makeAffImpl = box (Effect_Aff_FFI.``_makeAffImpl``)
let Effect_Aff__makeFiberNative = box (Effect_Aff_FFI.``_makeFiberNative``)
let Effect_Aff__makeSupervisedFiber = box (Effect_Aff_FFI.``_makeSupervisedFiber``)
let Effect_Aff__map = box (Effect_Aff_FFI.``_map``)
let Effect_Aff__onCompleteFiber = box (Effect_Aff_FFI.``_onCompleteFiber``)
let Effect_Aff__parAffAlt = box (Effect_Aff_FFI.``_parAffAlt``)
let Effect_Aff__parAffApply = box (Effect_Aff_FFI.``_parAffApply``)
let Effect_Aff__parAffMap = box (Effect_Aff_FFI.``_parAffMap``)
let Effect_Aff__pure = box (Effect_Aff_FFI.``_pure``)
let Effect_Aff__runFiber = box (Effect_Aff_FFI.``_runFiber``)
let Effect_Aff__sequential = box (Effect_Aff_FFI.``_sequential``)
let Effect_Aff__throwError = box (Effect_Aff_FFI.``_throwError``)
let Effect_Aff_generalBracket = box (Effect_Aff_FFI.``generalBracket``)


let Effect_Aff_pure  = (sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect))))

let Effect_Aff_void  = (sharpurs_apply (box ((box Data_Functor_void))) (box ((box Effect_functorEffect))))

let Effect_Aff_void1  = (sharpurs_apply (box ((box Data_Functor_void))) (box ((box Effect_functorEffect))))

let Effect_Aff_Fiber  = (box (fun (x: obj) -> (box x)))

let Effect_Aff_Canceler  = (box (fun (x: obj) -> (box x)))

let Effect_Aff_newtypeCanceler  = (sharpurs_apply (box ((box Data_Newtype_Newtypeusd_Dict))) (box ((box ((Map.add "Coercible0" (box ((box (fun (usd__unused: obj) -> (box Prim_undefined))))) Map.empty))))))

let Effect_Aff_makeFiber  = (box (fun (aff: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Effect_Aff__makeFiberNative))) (box ((box aff))))))))) (box ((box (fun (nf: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Effect_Aff_pure)))))) (box ((sharpurs_apply (box ((box Effect_Aff_Fiber))) (box ((box ((Map.add "run" (box ((sharpurs_apply (box ((box Effect_Aff__runFiber))) (box ((box nf)))))) (Map.add "kill" (box ((sharpurs_apply (box ((box Effect_Aff__killFiber))) (box ((box nf)))))) (Map.add "join" (box ((sharpurs_apply (box ((box Effect_Aff__joinFiber))) (box ((box nf)))))) (Map.add "onComplete" (box ((sharpurs_apply (box ((box Effect_Aff__onCompleteFiber))) (box ((box nf)))))) (Map.add "isSuspended" (box ((sharpurs_apply (box ((box Effect_Aff__isSuspendedFiber))) (box ((box nf)))))) Map.empty))))))))))))))))))))

let Effect_Aff_makeAff  = (box (fun (build: obj) -> (sharpurs_apply (box ((box Effect_Aff__makeAffImpl))) (box ((box (fun (onError: obj) -> (box (fun (onSuccess: obj) -> (sharpurs_apply (box ((box build))) (box ((box (fun (either: obj) -> (match ((unbox ((box either)))) with | Data_Either_Leftusd_Ctor(err) -> ((sharpurs_apply (box ((box onError))) (box ((box err))))) | Data_Either_Rightusd_Ctor(val_var) -> ((sharpurs_apply (box ((box onSuccess))) (box ((box val_var))))))))))))))))))))

let Effect_Aff_launchSuspendedAff  = (box (fun (aff: obj) -> (sharpurs_apply (box ((box Effect_Aff_makeFiber))) (box ((box aff))))))

let Effect_Aff_launchAff  = (box (fun (aff: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Effect_Aff_makeFiber))) (box ((box aff))))))))) (box ((box (fun (fiber: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Effect_bindEffect)))))) (box ((match ((unbox ((box fiber)))) with | f -> ((Map.find "run" (unbox<Map<string, obj>> ((box f))))))))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((box fiber))))))))))))))))

let Effect_Aff_launchAff_  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Aff_void)))))) (box ((box Effect_Aff_launchAff))))

let Effect_Aff_functorParAff  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box Effect_Aff__parAffMap))) Map.empty))))))

let Effect_Aff_functorAff  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box Effect_Aff__map))) Map.empty))))))

let Effect_Aff_delay  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | n -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn2))) (box ((box Effect_Aff__delay)))))) (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1))))))))))) (box ((box n))))))))

let Effect_Aff_bracket  = (box (fun (acquire: obj) -> (box (fun (completed: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Aff_generalBracket))) (box ((box acquire)))))) (box ((box ((Map.add "killed" (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box completed)))))) (Map.add "failed" (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box completed)))))) (Map.add "completed" (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box completed)))))) Map.empty))))))))))))

let Effect_Aff_applyParAff  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box Effect_Aff__parAffApply))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_functorParAff))))) Map.empty)))))))

let Effect_Aff_semigroupParAff  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_lift2))) (box ((box Effect_Aff_applyParAff)))))) (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup))))))))) Map.empty))))))))

let rec Effect_Aff_monadAff : obj = ((sharpurs_apply (box ((box Control_Monad_Monadusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applicativeAff))))) (Map.add "Bind1" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_bindAff))))) Map.empty))))))))
 and Effect_Aff_bindAff : obj = ((sharpurs_apply (box ((box Control_Bind_Bindusd_Dict))) (box ((box ((Map.add "bind" (box ((box Effect_Aff__bind))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applyAff))))) Map.empty))))))))
 and Effect_Aff_applyAff : obj = ((sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((sharpurs_apply (box ((box Control_Monad_ap))) (box ((box Effect_Aff_monadAff)))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_functorAff))))) Map.empty))))))))
 and Effect_Aff_applicativeAff : obj = ((sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box Effect_Aff__pure))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applyAff))))) Map.empty))))))))


let Effect_Aff_pure1  = (sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff))))

let Effect_Aff_cancelWith  = (box (fun (aff: obj) -> (box (fun (v: obj) -> (match (((unbox ((box aff))), (unbox ((box v))))) with | (aff1, cancel) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Aff_generalBracket))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box Data_Unit_unit))))))))) (box ((box ((Map.add "killed" (box ((box (fun (e: obj) -> (box (fun (v1: obj) -> (sharpurs_apply (box ((box cancel))) (box ((box e)))))))))) (Map.add "failed" (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff))))))))) (Map.add "completed" (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff))))))))) Map.empty)))))))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box aff1)))))))))))))

let Effect_Aff_finally  = (box (fun (fin: obj) -> (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Aff_bracket))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box Data_Unit_unit))))))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box fin))))))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box a)))))))))))

let Effect_Aff_invincible  = (box (fun (a: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Aff_bracket))) (box ((box a)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box Data_Unit_unit)))))))))))) (box ((box Effect_Aff_pure1))))))

let Effect_Aff_lazyAff  = (sharpurs_apply (box ((box Control_Lazy_Lazyusd_Dict))) (box ((box ((Map.add "defer" (box ((box (fun (f: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_Aff_bindAff)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box Data_Unit_unit))))))))) (box ((box f)))))))) Map.empty))))))

let Effect_Aff_parallelAff  = (sharpurs_apply (box ((box Control_Parallel_Class_Parallelusd_Dict))) (box ((box ((Map.add "parallel" (box ((box Unsafe_Coerce_unsafeCoerce))) (Map.add "sequential" (box ((box Effect_Aff__sequential))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applyAff))))) (Map.add "Apply1" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applyParAff))))) Map.empty)))))))))

let Effect_Aff_applicativeParAff  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Control_Parallel_Class_parallel))) (box ((box Effect_Aff_parallelAff))))))))) (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applyParAff))))) Map.empty)))))))

let Effect_Aff_monoidParAff  = (box (fun (dictMonoid: obj) -> (let semigroupParAff1 = (sharpurs_apply (box ((box Effect_Aff_semigroupParAff))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeParAff)))))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupParAff1))))) Map.empty))))))))))

let Effect_Aff_semigroupCanceler  = (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((box (fun (v: obj) -> (box (fun (v1: obj) -> (match (((unbox ((box v))), (unbox ((box v1))))) with | (c1, c2) -> ((sharpurs_apply (box ((box Effect_Aff_Canceler))) (box ((box (fun (err: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Parallel_parSequence_))) (box ((box Effect_Aff_parallelAff)))))) (box ((box Effect_Aff_applicativeParAff)))))) (box ((box Data_Foldable_foldableArray)))))) (box ((box [|(sharpurs_apply (box ((box c1))) (box ((box err)))); (sharpurs_apply (box ((box c2))) (box ((box err))))|]))))))))))))))))) Map.empty))))))

let Effect_Aff_semigroupAff  = (box (fun (dictSemigroup: obj) -> (sharpurs_apply (box ((box Data_Semigroup_Semigroupusd_Dict))) (box ((box ((Map.add "append" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_lift2))) (box ((box Effect_Aff_applyAff)))))) (box ((sharpurs_apply (box ((box Data_Semigroup_append))) (box ((box dictSemigroup))))))))) Map.empty))))))))

let Effect_Aff_monadEffectAff  = (sharpurs_apply (box ((box Effect_Class_MonadEffectusd_Dict))) (box ((box ((Map.add "liftEffect" (box ((box Effect_Aff__liftEffect))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_monadAff))))) Map.empty)))))))

let Effect_Aff_liftEffect  = (sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box Effect_Aff_monadEffectAff))))

let Effect_Aff_effectCanceler  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Aff_Canceler)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Data_Function_const)))))) (box ((box Effect_Aff_liftEffect)))))))

let Effect_Aff_joinFiber  = (box (fun (v: obj) -> (match ((unbox ((box v)))) with | t -> ((sharpurs_apply (box ((box Effect_Aff_makeAff))) (box ((box (fun (k: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_functorEffect)))))) (box ((box Effect_Aff_effectCanceler)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((Map.find "join" (unbox<Map<string, obj>> ((box t)))))) (box ((box (fun (err: obj) -> (sharpurs_apply (box ((box k))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1)))))))) (box ((box err)))))))))))))) (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box k))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1)))))))) (box ((box a))))))))))))))))))))))))

let Effect_Aff_functorFiber  = (sharpurs_apply (box ((box Data_Functor_Functorusd_Dict))) (box ((box ((Map.add "map" (box ((box (fun (f: obj) -> (box (fun (t: obj) -> (sharpurs_apply (box ((box Effect_Unsafe_unsafePerformEffect))) (box ((sharpurs_apply (box ((box Effect_Aff_makeFiber))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_Aff_functorAff)))))) (box ((box f)))))) (box ((sharpurs_apply (box ((box Effect_Aff_joinFiber))) (box ((box t))))))))))))))))))) Map.empty))))))

let Effect_Aff_applyFiber  = (sharpurs_apply (box ((box Control_Apply_Applyusd_Dict))) (box ((box ((Map.add "apply" (box ((box (fun (t1: obj) -> (box (fun (t2: obj) -> (sharpurs_apply (box ((box Effect_Unsafe_unsafePerformEffect))) (box ((sharpurs_apply (box ((box Effect_Aff_makeFiber))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Apply_apply))) (box ((box Effect_Aff_applyAff)))))) (box ((sharpurs_apply (box ((box Effect_Aff_joinFiber))) (box ((box t1))))))))) (box ((sharpurs_apply (box ((box Effect_Aff_joinFiber))) (box ((box t2))))))))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_functorFiber))))) Map.empty)))))))

let Effect_Aff_applicativeFiber  = (sharpurs_apply (box ((box Control_Applicative_Applicativeusd_Dict))) (box ((box ((Map.add "pure" (box ((box (fun (a: obj) -> (sharpurs_apply (box ((box Effect_Unsafe_unsafePerformEffect))) (box ((sharpurs_apply (box ((box Effect_Aff_makeFiber))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box a)))))))))))))) (Map.add "Apply0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applyFiber))))) Map.empty)))))))

let Effect_Aff_forkAff  = (box (fun (aff: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_Aff_bindAff)))))) (box ((sharpurs_apply (box ((box Effect_Aff__forkAffNative))) (box ((box aff))))))))) (box ((box (fun (nf: obj) -> (let fiber = (sharpurs_apply (box ((box Effect_Aff_Fiber))) (box ((box ((Map.add "run" (box ((sharpurs_apply (box ((box Effect_Aff__runFiber))) (box ((box nf)))))) (Map.add "kill" (box ((sharpurs_apply (box ((box Effect_Aff__killFiber))) (box ((box nf)))))) (Map.add "join" (box ((sharpurs_apply (box ((box Effect_Aff__joinFiber))) (box ((box nf)))))) (Map.add "onComplete" (box ((sharpurs_apply (box ((box Effect_Aff__onCompleteFiber))) (box ((box nf)))))) (Map.add "isSuspended" (box ((sharpurs_apply (box ((box Effect_Aff__isSuspendedFiber))) (box ((box nf)))))) Map.empty)))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Effect_Aff_bindAff)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box Effect_Aff_monadEffectAff)))))) (box ((sharpurs_apply (box ((box Effect_Aff__runFiber))) (box ((box nf)))))))))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box fiber)))))))))))))))))

let Effect_Aff_killFiber  = (box (fun (e: obj) -> (box (fun (v: obj) -> (match (((unbox ((box e))), (unbox ((box v))))) with | (e1, t) -> ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_Aff_bindAff)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box Effect_Aff_monadEffectAff)))))) (box ((Map.find "isSuspended" (unbox<Map<string, obj>> ((box t)))))))))))) (box ((box (fun (suspended: obj) -> (match ((unbox ((box suspended)))) with | LitBool true () -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Effect_Aff_liftEffect)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Effect_Aff_void1)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((Map.find "kill" (unbox<Map<string, obj>> ((box t)))))) (box ((box e1)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((box Data_Unit_unit)))))))))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((box Data_Unit_unit))))))))))))))))) | _ -> ((sharpurs_apply (box ((box Effect_Aff_makeAff))) (box ((box (fun (k: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_functorEffect)))))) (box ((box Effect_Aff_effectCanceler)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((Map.find "kill" (unbox<Map<string, obj>> ((box t)))))) (box ((box e1)))))) (box ((box (fun (err: obj) -> (sharpurs_apply (box ((box k))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Leftusd_Ctor(usd__arg1)))))))) (box ((box err)))))))))))))) (box ((box (fun (v1: obj) -> (sharpurs_apply (box ((box k))) (box ((sharpurs_apply (box ((box ((fun (usd__arg1: obj) -> (box (Data_Either_Rightusd_Ctor(usd__arg1)))))))) (box ((box Data_Unit_unit)))))))))))))))))))))))))))))))))

let Effect_Aff_fiberCanceler  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Aff_Canceler)))))) (box ((sharpurs_apply (box ((box Data_Function_flip))) (box ((box Effect_Aff_killFiber)))))))

let Effect_Aff_supervise  = (box (fun (aff: obj) -> (let killError = (sharpurs_apply (box ((box Effect_Exception_error))) (box ((box "[Aff] Child fiber outlived parent")))) in let killAll = (box (fun (err: obj) -> (box (fun (sup: obj) -> (sharpurs_apply (box ((box Effect_Aff_makeAff))) (box ((box (fun (k: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_Uncurried_runFn3))) (box ((box Effect_Aff__killAll)))))) (box ((box err)))))) (box ((Map.find "supervisor" (unbox<Map<string, obj>> ((box sup))))))))) (box ((sharpurs_apply (box ((box k))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Data_Either_applicativeEither)))))) (box ((box Data_Unit_unit))))))))))))))))))) in let acquire = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_bindEffect)))))) (box ((sharpurs_apply (box ((box Effect_Aff__makeSupervisedFiber))) (box ((box aff))))))))) (box ((box (fun (sup: obj) -> (let fiber = (sharpurs_apply (box ((box Effect_Aff_Fiber))) (box ((box ((Map.add "run" (box ((sharpurs_apply (box ((box Effect_Aff__runFiber))) (box ((Map.find "fiber" (unbox<Map<string, obj>> ((box sup))))))))) (Map.add "kill" (box ((sharpurs_apply (box ((box Effect_Aff__killFiber))) (box ((Map.find "fiber" (unbox<Map<string, obj>> ((box sup))))))))) (Map.add "join" (box ((sharpurs_apply (box ((box Effect_Aff__joinFiber))) (box ((Map.find "fiber" (unbox<Map<string, obj>> ((box sup))))))))) (Map.add "onComplete" (box ((sharpurs_apply (box ((box Effect_Aff__onCompleteFiber))) (box ((Map.find "fiber" (unbox<Map<string, obj>> ((box sup))))))))) (Map.add "isSuspended" (box ((sharpurs_apply (box ((box Effect_Aff__isSuspendedFiber))) (box ((Map.find "fiber" (unbox<Map<string, obj>> ((box sup))))))))) Map.empty)))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_discard))) (box ((box Control_Bind_discardUnit)))))) (box ((box Effect_bindEffect)))))) (box ((match ((unbox ((box fiber)))) with | f -> ((Map.find "run" (unbox<Map<string, obj>> ((box f))))))))))) (box ((box (fun (usd__unused: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((box ((Map.add "supervisor" (box ((Map.find "supervisor" (unbox<Map<string, obj>> ((box sup)))))) (Map.add "fiber" (box ((box fiber))) Map.empty)))))))))))))))))) in (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Aff_generalBracket))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box Effect_Aff_monadEffectAff)))))) (box ((box acquire))))))))) (box ((box ((Map.add "killed" (box ((box (fun (err: obj) -> (box (fun (sup: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Parallel_parSequence_))) (box ((box Effect_Aff_parallelAff)))))) (box ((box Effect_Aff_applicativeParAff)))))) (box ((box Data_Foldable_foldableArray)))))) (box ((box [|(sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Aff_killFiber))) (box ((box err)))))) (box ((Map.find "fiber" (unbox<Map<string, obj>> ((box sup))))))); (sharpurs_apply (box ((sharpurs_apply (box ((box killAll))) (box ((box err)))))) (box ((box sup))))|])))))))))) (Map.add "failed" (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((box killAll))) (box ((box killError))))))))) (Map.add "completed" (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((box killAll))) (box ((box killError))))))))) Map.empty)))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Aff_joinFiber)))))) (box ((box (fun (v: obj) -> (Map.find "fiber" (unbox<Map<string, obj>> ((box v)))))))))))))))

let Effect_Aff_suspendAff  = (box (fun (aff: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box Effect_Aff_monadEffectAff)))))) (box ((sharpurs_apply (box ((box Effect_Aff_makeFiber))) (box ((box aff)))))))))

let Effect_Aff_monadSTAff  = (sharpurs_apply (box ((box Control_Monad_ST_Class_MonadSTusd_Dict))) (box ((box ((Map.add "liftST" (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((sharpurs_apply (box ((box Control_Monad_ST_Class_liftST))) (box ((box Control_Monad_ST_Class_monadSTEffect))))))))) (box ((sharpurs_apply (box ((box Effect_Class_liftEffect))) (box ((box Effect_Aff_monadEffectAff))))))))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_monadAff))))) Map.empty)))))))

let Effect_Aff_monadThrowAff  = (sharpurs_apply (box ((box Control_Monad_Error_Class_MonadThrowusd_Dict))) (box ((box ((Map.add "throwError" (box ((box Effect_Aff__throwError))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_monadAff))))) Map.empty)))))))

let Effect_Aff_monadErrorAff  = (sharpurs_apply (box ((box Control_Monad_Error_Class_MonadErrorusd_Dict))) (box ((box ((Map.add "catchError" (box ((box Effect_Aff__catchError))) (Map.add "MonadThrow0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_monadThrowAff))))) Map.empty)))))))

let Effect_Aff_attempt  = (sharpurs_apply (box ((box Control_Monad_Error_Class_try))) (box ((box Effect_Aff_monadErrorAff))))

let Effect_Aff_runAff  = (box (fun (k: obj) -> (box (fun (aff: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Effect_Aff_launchAff)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bindFlipped))) (box ((box Effect_Aff_bindAff)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Aff_liftEffect)))))) (box ((box k))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_Error_Class_try))) (box ((box Effect_Aff_monadErrorAff)))))) (box ((box aff))))))))))))))

let Effect_Aff_runAff_  = (box (fun (k: obj) -> (box (fun (aff: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Effect_Aff_void)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Effect_Aff_runAff))) (box ((box k)))))) (box ((box aff)))))))))))

let Effect_Aff_runSuspendedAff  = (box (fun (k: obj) -> (box (fun (aff: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Data_Function_apply))) (box ((box Effect_Aff_launchSuspendedAff)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bindFlipped))) (box ((box Effect_Aff_bindAff)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_compose))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Aff_liftEffect)))))) (box ((box k))))))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_Error_Class_try))) (box ((box Effect_Aff_monadErrorAff)))))) (box ((box aff))))))))))))))

let Effect_Aff_monadRecAff  = (sharpurs_apply (box ((box Control_Monad_Rec_Class_MonadRecusd_Dict))) (box ((box ((Map.add "tailRecM" (box ((box (fun (k: obj) -> (
                                                                                                                                                                                                        
                                                                                                                                                                                                        let rec go_tco (a: obj) : obj = ((sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Bind_bind))) (box ((box Effect_Aff_bindAff)))))) (box ((sharpurs_apply (box ((box k))) (box ((box a))))))))) (box ((box (fun (res: obj) -> (match ((unbox ((box res)))) with | Control_Monad_Rec_Class_Doneusd_Ctor(r) -> ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box r))))) | Control_Monad_Rec_Class_Loopusd_Ctor(b) -> ((go_tco ((box b))))))))))) 
                                                                                                                                                                                                        and go = box ((fun (a: obj) -> go_tco a)) 
                                                                                                                                                                                                        in
                                                                                                                                                                                                        (box go)
                                                                                                                                                                                                        ))))) (Map.add "Monad0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_monadAff))))) Map.empty)))))))

let Effect_Aff_monoidAff  = (box (fun (dictMonoid: obj) -> (let semigroupAff1 = (sharpurs_apply (box ((box Effect_Aff_semigroupAff))) (box ((sharpurs_apply (box ((Map.find "Semigroup0" (unbox<Map<string, obj>> ((box dictMonoid)))))) (box ((box Prim_undefined))))))) in (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box dictMonoid))))))))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box semigroupAff1))))) Map.empty))))))))))

let Effect_Aff_nonCanceler  = (sharpurs_apply (box ((box Effect_Aff_Canceler))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_Aff_applicativeAff)))))) (box ((box Data_Unit_unit))))))))))

let Effect_Aff_monoidCanceler  = (sharpurs_apply (box ((box Data_Monoid_Monoidusd_Dict))) (box ((box ((Map.add "mempty" (box ((box Effect_Aff_nonCanceler))) (Map.add "Semigroup0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_semigroupCanceler))))) Map.empty)))))))

let Effect_Aff_never  = (sharpurs_apply (box ((box Effect_Aff_makeAff))) (box ((box (fun (v: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((box Control_Applicative_pure))) (box ((box Effect_applicativeEffect)))))) (box ((sharpurs_apply (box ((box Data_Monoid_mempty))) (box ((box Effect_Aff_monoidCanceler))))))))))))

let Effect_Aff_apathize  = (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Semigroupoid_composeFlipped))) (box ((box Control_Semigroupoid_semigroupoidFn)))))) (box ((box Effect_Aff_attempt)))))) (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Data_Functor_map))) (box ((box Effect_Aff_functorAff)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box Data_Unit_unit))))))))))

let Effect_Aff_altParAff  = (sharpurs_apply (box ((box Control_Alt_Altusd_Dict))) (box ((box ((Map.add "alt" (box ((box Effect_Aff__parAffAlt))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_functorParAff))))) Map.empty)))))))

let Effect_Aff_altAff  = (sharpurs_apply (box ((box Control_Alt_Altusd_Dict))) (box ((box ((Map.add "alt" (box ((box (fun (a1: obj) -> (box (fun (a2: obj) -> (sharpurs_apply (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_Error_Class_catchError))) (box ((box Effect_Aff_monadErrorAff)))))) (box ((box a1)))))) (box ((sharpurs_apply (box ((box Data_Function_const))) (box ((box a2))))))))))))) (Map.add "Functor0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_functorAff))))) Map.empty)))))))

let Effect_Aff_plusAff  = (sharpurs_apply (box ((box Control_Plus_Plususd_Dict))) (box ((box ((Map.add "empty" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Monad_Error_Class_throwError))) (box ((box Effect_Aff_monadThrowAff)))))) (box ((sharpurs_apply (box ((box Effect_Exception_error))) (box ((box "Always fails"))))))))) (Map.add "Alt0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_altAff))))) Map.empty)))))))

let Effect_Aff_plusParAff  = (sharpurs_apply (box ((box Control_Plus_Plususd_Dict))) (box ((box ((Map.add "empty" (box ((sharpurs_apply (box ((sharpurs_apply (box ((box Control_Parallel_Class_parallel))) (box ((box Effect_Aff_parallelAff)))))) (box ((sharpurs_apply (box ((box Control_Plus_empty))) (box ((box Effect_Aff_plusAff))))))))) (Map.add "Alt0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_altParAff))))) Map.empty)))))))

let Effect_Aff_alternativeParAff  = (sharpurs_apply (box ((box Control_Alternative_Alternativeusd_Dict))) (box ((box ((Map.add "Applicative0" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_applicativeParAff))))) (Map.add "Plus1" (box ((box (fun (usd__unused: obj) -> (box Effect_Aff_plusParAff))))) Map.empty)))))))

package Effect_Aff


import (
	import_atomic "sync/atomic"
	"context"
	"fmt"
	"sync"
	"time"
	"gopurs/output/gopurs_runtime"
)

type AffFn = func(context.Context) (any, error)

type BindNode struct {
	Aff any
	K   func(any) AffFn
}

func runAffSync(aff AffFn, ctx context.Context) (any, error) {
	var current = aff
	var stack []func(any) AffFn

	for {
		val, err := current(ctx)
		if err != nil {
			return nil, err
		}

		if node, ok := val.(BindNode); ok {
			stack = append(stack, node.K)
			current = node.Aff.(AffFn)
		} else {
			if len(stack) > 0 {
				k := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				current = k(val)
			} else {
				return val, nil
			}
		}
	}
}

func _Pure(val any) any {
	return func(ctx context.Context) (any, error) {
		select {
		case <-ctx.Done():
			return nil, context.Cause(ctx)
		default:
			return val, nil
		}
	}
}

func _Bind(aff AffFn, k func(any) AffFn) any {
	return func(ctx context.Context) (any, error) {
		return BindNode{Aff: aff, K: k}, nil
	}
}

func _Delay(right any, ms float64) any {
	return func(ctx context.Context) (any, error) {
		duration := time.Duration(ms) * time.Millisecond
		timer := time.NewTimer(duration)
		defer timer.Stop()

		select {
		case <-timer.C:
			return nil, nil
		case <-ctx.Done():
			return nil, context.Cause(ctx)
		}
	}
}

func _LiftEffect(eff func(any) any) any {
	return func(ctx context.Context) (any, error) {
		select {
		case <-ctx.Done():
			return nil, context.Cause(ctx)
		default:
			return eff(nil), nil
		}
	}
}

func _MakeAffImpl(build func(func(error) func(any) any) func(func(any) func(any) any) func(any) func(any) AffFn) any {
	return func(ctx context.Context) (any, error) {
		resultChan := make(chan struct {
			val any
			err error
		}, 1)

		onError := func(err error) func(any) any {
			return func(_ any) any {
				select {
				case resultChan <- struct{val any; err error}{nil, err}:
				default:
				}
				return nil
			}
		}

		onSuccess := func(val any) func(any) any {
			return func(_ any) any {
				select {
				case resultChan <- struct{val any; err error}{val, nil}:
				default:
				}
				return nil
			}
		}

		cancelerEffect := build(onError)(onSuccess)
		canceler := cancelerEffect(nil)

		select {
		case res := <-resultChan:
			return res.val, res.err
		case <-ctx.Done():
			cancelFnAff := canceler(context.Cause(ctx))
			_, _ = runAffSync(cancelFnAff, context.Background())
			return nil, context.Cause(ctx)
		}
	}
}




type key int
const killErrKey key = 0
const supervisorKey key = 1
type KillState struct {
	Err error
}

type Supervisor struct {
	Ctx    context.Context
	Cancel context.CancelCauseFunc
	Wg     *sync.WaitGroup
}

type NativeFiber struct {
	Aff AffFn
	Ctx        context.Context
	Done       chan struct{}
	Start      chan struct{}
	Val        any
	Err        error
	Cancel     context.CancelCauseFunc
	Id         int64
	mu         sync.Mutex
	IsComplete int32
}

func _MakeFiberNative(aff AffFn) any {
	return func(_ any) any {
		ctx, cancel := context.WithCancelCause(context.Background())
		done := make(chan struct{})
		start := make(chan struct{})
		
		fiberId := time.Now().UnixNano()
		nf := &NativeFiber{
			Aff:        aff,
			Ctx:        ctx,
			Done:       done,
			Start:      start,
			Cancel:     cancel,
			Id:         fiberId,
			IsComplete: 0,
		}
		return nf
	}
}


func _ForkAffNative(aff_ any) any {
	aff := gopurs_runtime.Unbox[AffFn](aff_.(gopurs_runtime.Value))
	return func(ctx context.Context) (any, error) {
		var childCtx context.Context
		var cancel context.CancelCauseFunc
		supAny := ctx.Value(supervisorKey)
		if supAny != nil {
			sup := supAny.(*Supervisor)
			childCtx, cancel = context.WithCancelCause(sup.Ctx)
		} else {
			childCtx, cancel = context.WithCancelCause(context.Background())
		}
		done := make(chan struct{})
		start := make(chan struct{})
		
		fiberId := time.Now().UnixNano()
		nf := &NativeFiber{
			Aff:        aff,
			Ctx:        childCtx,
			Done:       done,
			Start:      start,
			Cancel:     cancel,
			Id:         fiberId,
			IsComplete: 0,
		}
		close(start)

		if supAny != nil {
			sup := supAny.(*Supervisor)
			sup.Wg.Add(1)
			
			gopurs_runtime.Retain()
			go func() {
				defer gopurs_runtime.Release()
				defer sup.Wg.Done()
				<-nf.Start
				
				// Keep the supervisor in the child's context
				ctxWithSup := context.WithValue(childCtx, supervisorKey, sup)
				val, err := runAffSync(aff, ctxWithSup)
				
				nf.mu.Lock()
				nf.Val = val
				nf.Err = err
				import_atomic.StoreInt32(&nf.IsComplete, 1)
				nf.mu.Unlock()
				close(nf.Done)
			}()
		} else {
			gopurs_runtime.Retain()
			go func() {
				defer gopurs_runtime.Release()
				<-nf.Start
				
				val, err := runAffSync(aff, childCtx)
				
				nf.mu.Lock()
				nf.Val = val
				nf.Err = err
				import_atomic.StoreInt32(&nf.IsComplete, 1)
				nf.mu.Unlock()
				close(nf.Done)
			}()
		}

		return nf, nil
	}
}

func _RunFiber(nf *NativeFiber, _ interface{}) any {
	select {
	case <-nf.Start:
	default:
		gopurs_runtime.Retain()
		go func() {
			defer gopurs_runtime.Release()
			val, err := runAffSync(nf.Aff, nf.Ctx)
			nf.mu.Lock()
			nf.Val = val
			nf.Err = err
			import_atomic.StoreInt32(&nf.IsComplete, 1)
			nf.mu.Unlock()
			close(nf.Done)
		}()
		close(nf.Start)
	}
	return nil
}

func _KillFiber(nf *NativeFiber, errAny error, onError func(any) func(any) any, onSuccess func(any) func(any) any) any {
	return func(_ any) any {
		nf.Cancel(errAny)
	select {
	case <-nf.Start:
	default:
		close(nf.Start)
	}
		go func() {
			<-nf.Done
			onSuccess(nil)(nil)
		}()
		return func(_ any) any {
			return nil
		}
	}
}

func _JoinFiber(nf *NativeFiber, onError func(any) func(any) any, onSuccess func(any) func(any) any) any {
	return func(_ any) any {
		_RunFiber(nf, nil)
		go func() {
			<-nf.Done
			
			if nf.Err != nil {
				onError(nf.Err)(nil)
			} else {
				onSuccess(nf.Val)(nil)
			}
		}()
		return func(_ any) any {
			return nil
		}
	}
}


func _OnCompleteFiber(nf *NativeFiber, onCompleteAny any) any {
	return func(_ any) any {
		return func(_ any) any {
			return nil
		}
	}
}

func _IsSuspendedFiber(nf *NativeFiber) func(any) any {
	return func(any) any {
		return false
	}
}

func _ThrowError(err error) any {
	return func(ctx context.Context) (any, error) {
		return nil, err
	}
}

func _CatchError(aff AffFn, handler func(any) AffFn) any {
	return func(ctx context.Context) (any, error) {
		val, err := runAffSync(aff, ctx)
		if err != nil {
			if context.Cause(ctx) != nil && context.Cause(ctx) == err {
				return nil, err
			}
			return runAffSync(handler(err), ctx)
		}
		return val, nil
	}
}

func _Map(f func(any) any, aff AffFn) any {
	return func(ctx context.Context) (any, error) {
		val, err := runAffSync(aff, ctx)
		if err != nil {
			return nil, err
		}
		return f(val), nil
	}
}

func _ParAffMap(f func(any) any, aff AffFn) any {
	return _Map(f, aff)
}

func _ParAffApply(aff1 AffFn, aff2 AffFn) any {
	return func(ctx context.Context) (any, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		var wg sync.WaitGroup
		wg.Add(2)

		var res1 any
		var res2 any

		var firstErr error
		var mu sync.Mutex

		go func() {
			defer wg.Done()
			var err1 error
			res1, err1 = runAffSync(aff1, ctx)
			if err1 != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err1
				}
				mu.Unlock()
				cancel()
			}
		}()

		go func() {
			defer wg.Done()
			var err2 error
			res2, err2 = runAffSync(aff2, ctx)
			if err2 != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = err2
				}
				mu.Unlock()
				cancel()
			}
		}()

		wg.Wait()

		if firstErr != nil {
			return nil, firstErr
		}

		if val, ok := res1.(gopurs_runtime.Value); ok {
			return gopurs_runtime.Apply(val, gopurs_runtime.Box(res2)), nil
		}
		if res1 == nil {
			return nil, nil
		}
		f := res1.(func(any) any)
		return f(res2), nil
	}
}
func _ParAffAlt(aff1 AffFn, aff2 AffFn) any {
	return func(ctx context.Context) (any, error) {
		fn1 := aff1
		fn2 := aff2

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		type Result struct {
			val any
			err error
		}
		resCh := make(chan Result, 2)

		go func() {
			val, err := runAffSync(fn1, ctx)
			resCh <- Result{val, err}
		}()
		go func() {
			val, err := runAffSync(fn2, ctx)
			resCh <- Result{val, err}
		}()

		var firstErr error
		for i := 0; i < 2; i++ {
			res := <-resCh
			if res.err == nil {
				cancel()
				if i == 0 {
					<-resCh
				}
				return res.val, nil
			}
			if firstErr == nil {
				firstErr = res.err
			}
		}
		return nil, firstErr
	}
}
func _KillAll(err_ any, sup_ any, cb_ any) any {
	return func(_ any) any {
		sup := gopurs_runtime.Unbox[*Supervisor](sup_.(gopurs_runtime.Value))
		cb := gopurs_runtime.Unbox[func(any) any](cb_.(gopurs_runtime.Value))
		
		go func() {
			errGo := fmt.Errorf("Supervised fiber canceled")
			sup.Cancel(errGo)
			sup.Wg.Wait()
			cb(nil)
		}()
		
		// Return empty Canceler: Error -> Aff Unit
		return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			aff := func(ctx context.Context) (any, error) {
				return nil, nil
			}
			return gopurs_runtime.Box(aff)
		})
	}
}
func _Sequential(aff AffFn) any { return aff }
func GeneralBracket(acquireBox any, optionsBox any, useBox any) any {
	return func(ctx context.Context) (any, error) {
		acquireVal := acquireBox.(gopurs_runtime.Value)
		acquireFn := gopurs_runtime.Unbox[AffFn](acquireVal)
		
		acquireCtx := context.WithoutCancel(ctx)
		resource, err := runAffSync(acquireFn, acquireCtx)
		if err != nil {
			return nil, err
		}
		
		optionsVal := optionsBox.(gopurs_runtime.Value)
		
		if ctx.Err() != nil {
			err = context.Cause(ctx)
			if err == nil {
				err = ctx.Err()
			}
			killedBox := gopurs_runtime.RecordGet(optionsVal, "killed")
			cleanupBox := gopurs_runtime.Apply2(killedBox, gopurs_runtime.Box(err), gopurs_runtime.Box(resource))
			cleanupFn := gopurs_runtime.Unbox[AffFn](cleanupBox)
			cleanupCtx := context.WithoutCancel(ctx)
			_, _ = runAffSync(cleanupFn, cleanupCtx)
			return nil, err
		}
		
		useVal := useBox.(gopurs_runtime.Value)
		useResultBox := gopurs_runtime.Apply(useVal, gopurs_runtime.Box(resource))
		useFn := gopurs_runtime.Unbox[AffFn](useResultBox)
		
		val, err := runAffSync(useFn, ctx)
		
		if err != nil {
			// Check if it was canceled
			cause := context.Cause(ctx)
			var cleanupBox any
			if cause != nil && (err == cause || err.Error() == cause.Error()) {
				killedBox := gopurs_runtime.RecordGet(optionsVal, "killed")
				errBox := gopurs_runtime.Box(err)
				cleanupBox = gopurs_runtime.Apply2(killedBox, errBox, gopurs_runtime.Box(resource))
			} else {
				failedBox := gopurs_runtime.RecordGet(optionsVal, "failed")
				errBox := gopurs_runtime.Box(err)
				cleanupBox = gopurs_runtime.Apply2(failedBox, errBox, gopurs_runtime.Box(resource))
			}
			cleanupFn := gopurs_runtime.Unbox[AffFn](cleanupBox)
			cleanupCtx := context.WithoutCancel(ctx)
			_, _ = runAffSync(cleanupFn, cleanupCtx)
			return nil, err
		} else {
			completedBox := gopurs_runtime.RecordGet(optionsVal, "completed")
			cleanupBox := gopurs_runtime.Apply2(completedBox, gopurs_runtime.Box(val), gopurs_runtime.Box(resource))
			cleanupFn := gopurs_runtime.Unbox[AffFn](cleanupBox)
			cleanupCtx := context.WithoutCancel(ctx)
			_, _ = runAffSync(cleanupFn, cleanupCtx)
			return val, nil
		}
	}
}

func _MakeSupervisedFiber(aff AffFn) any {
	return func(_ any) any {
		supCtx, cancel := context.WithCancelCause(context.Background())
		sup := &Supervisor{
			Ctx:    supCtx,
			Cancel: cancel,
			Wg:     &sync.WaitGroup{},
		}
		
		ctxWithSup := context.WithValue(supCtx, supervisorKey, sup)
		
		fiberId := time.Now().UnixNano()
		nf := &NativeFiber{
			Aff:        aff,
			Ctx:        ctxWithSup, // supervised
			Done:       make(chan struct{}),
			Start:      make(chan struct{}),
			Cancel:     cancel,
			Id:         fiberId,
		}
		
		rec := gopurs_runtime.RecordDict2("fiber", "supervisor", gopurs_runtime.Box(nf), gopurs_runtime.Box(sup))
		
		return rec
	}
}


// --- Auto-generated FFI wrappers ---
var _Gopurs__Bind = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)]), (Func [(TypeVar a)] (ADT ["Effect","Aff","Aff"] [(TypeVar b)]))] (ADT ["Effect","Aff","Aff"] [(TypeVar b)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_arg1 := func(p0_0 any) AffFn {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[AffFn](inner_res0)
		}
	go_res := _Bind(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__CatchError = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)]), (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Aff","Aff"] [(TypeVar a)]))] (ADT ["Effect","Aff","Aff"] [(TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_arg1 := func(p0_0 any) AffFn {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[AffFn](inner_res0)
		}
	go_res := _CatchError(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Delay = // TAST: (ADT ["Data","Function","Uncurried","Fn2"] [(Func [(ADT ["Data","Unit","Unit"] [])] (ADT ["Data","Either","Either"] [(TypeVar a), (ADT ["Data","Unit","Unit"] [])])), Number, (ADT ["Effect","Aff","Aff"] [(ADT ["Data","Unit","Unit"] [])])])
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := _Delay(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ForkAffNative = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)])] (ADT ["Effect","Aff","Aff"] [(ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _ForkAffNative(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__IsSuspendedFiber = // TAST: (Func [(ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)])] (ADT ["Effect","Effect"] [Boolean]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*NativeFiber](arg0)
	go_res := _IsSuspendedFiber(go_arg0)
	return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res(arg)
				return gopurs_runtime.Box(inner_res)
			})
})
var _Gopurs__JoinFiber = // TAST: (Func [(ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)]), (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])), (Func [(TypeVar a)] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*NativeFiber](arg0)
	go_arg1 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg2 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_res := _JoinFiber(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__KillAll = // TAST: (ADT ["Data","Function","Uncurried","Fn3"] [(ADT ["Effect","Exception","Error"] []), (ADT ["Effect","Aff","Supervisor"] []), (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]), (ADT ["Effect","Effect"] [(Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Aff","Aff"] [(ADT ["Data","Unit","Unit"] [])]))])])
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := _KillAll(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__KillFiber = // TAST: (Func [(ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)]), (ADT ["Effect","Exception","Error"] []), (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])), (Func [(ADT ["Data","Unit","Unit"] [])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*NativeFiber](arg0)
	go_arg1 := gopurs_runtime.Unbox[error](arg1)
	go_arg2 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg3 := func(p0_0 any) func(any) any {
			inner_res0 := gopurs_runtime.Apply(arg3, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) any {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_res := _KillFiber(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__LiftEffect = // TAST: (Func [(ADT ["Effect","Effect"] [(TypeVar a)])] (ADT ["Effect","Aff","Aff"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := _LiftEffect(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__MakeAffImpl = // TAST: (Func [(Func [(Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])), (Func [(TypeVar a)] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))] (ADT ["Effect","Effect"] [(Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Aff","Aff"] [(ADT ["Data","Unit","Unit"] [])]))]))] (ADT ["Effect","Aff","Aff"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 func(error) func(any) any) func(func(any) func(any) any) func(any) func(any) AffFn {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
						inner_res := p0_0(gopurs_runtime.Unbox[error](arg))
						return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
						inner_res := inner_res(arg)
						return gopurs_runtime.Box(inner_res)
					})
					}))
			return func(p1_0 func(any) func(any) any) func(any) func(any) AffFn {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
						inner_res := p1_0(arg)
						return gopurs_runtime.Func(func(arg gopurs_runtime.Value) gopurs_runtime.Value {
						inner_res := inner_res(arg)
						return gopurs_runtime.Box(inner_res)
					})
					}))
			return func(p2_0 any) func(any) AffFn {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 any) AffFn {
			inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
			return gopurs_runtime.Unbox[AffFn](inner_res3)
		}
		}
		}
		}
	go_res := _MakeAffImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__MakeFiberNative = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)])] (ADT ["Effect","Effect"] [(ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_res := _MakeFiberNative(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__MakeSupervisedFiber = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)])] (ADT ["Effect","Effect"] [(Record [fiber: (ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)]), supervisor: (ADT ["Effect","Aff","Supervisor"] [])])]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_res := _MakeSupervisedFiber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Map = // TAST: (Func [(Func [(TypeVar a)] (TypeVar b)), (ADT ["Effect","Aff","Aff"] [(TypeVar a)])] (ADT ["Effect","Aff","Aff"] [(TypeVar b)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := gopurs_runtime.Unbox[AffFn](arg1)
	go_res := _Map(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__OnCompleteFiber = // TAST: (Func [(ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)]), Any] (ADT ["Effect","Effect"] [(ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*NativeFiber](arg0)
	go_arg1 := arg1
	go_res := _OnCompleteFiber(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ParAffAlt = // TAST: (Func [(ADT ["Effect","Aff","ParAff"] [(TypeVar a)]), (ADT ["Effect","Aff","ParAff"] [(TypeVar a)])] (ADT ["Effect","Aff","ParAff"] [(TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_arg1 := gopurs_runtime.Unbox[AffFn](arg1)
	go_res := _ParAffAlt(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ParAffApply = // TAST: (Func [(ADT ["Effect","Aff","ParAff"] [(Func [(TypeVar a)] (TypeVar b))]), (ADT ["Effect","Aff","ParAff"] [(TypeVar a)])] (ADT ["Effect","Aff","ParAff"] [(TypeVar b)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_arg1 := gopurs_runtime.Unbox[AffFn](arg1)
	go_res := _ParAffApply(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ParAffMap = // TAST: (Func [(Func [(TypeVar a)] (TypeVar b)), (ADT ["Effect","Aff","ParAff"] [(TypeVar a)])] (ADT ["Effect","Aff","ParAff"] [(TypeVar b)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := gopurs_runtime.Unbox[AffFn](arg1)
	go_res := _ParAffMap(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Pure = // TAST: (Func [(TypeVar a)] (ADT ["Effect","Aff","Aff"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Pure(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__RunFiber = // TAST: (Func [(ADT ["Effect","Aff","NativeFiber"] [(TypeVar a)])] (ADT ["Effect","Effect"] [(ADT ["Data","Unit","Unit"] [])]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*NativeFiber](arg0)
	go_arg1 := arg1
	go_res := _RunFiber(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Sequential = // TAST: Any
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_res := _Sequential(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ThrowError = // TAST: (Func [(ADT ["Effect","Exception","Error"] [])] (ADT ["Effect","Aff","Aff"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[error](arg0)
	go_res := _ThrowError(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_GeneralBracket = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)]), Any, (Func [(TypeVar a)] (ADT ["Effect","Aff","Aff"] [(TypeVar b)]))] (ADT ["Effect","Aff","Aff"] [(TypeVar b)]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := GeneralBracket(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
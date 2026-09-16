import (
	import_atomic "sync/atomic"
	"context"
	"fmt"
	"sync"
	"time"
	"gopurs/output/gopurs_runtime"
)

type AffFn = func(context.Context) (any, error)

type AffBindNode struct {
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

		if node, ok := val.(AffBindNode); ok {
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
		return AffBindNode{Aff: aff, K: k}, nil
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

func _RunFiber(nf *NativeFiber, x interface{}) any {
	return internalRunFiber(nf, x)
}

func internalRunFiber(nf *NativeFiber, _ interface{}) any {
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
		internalRunFiber(nf, nil)
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
	return internalMap(f, aff)
}

func internalMap(f func(any) any, aff AffFn) any {
	return func(ctx context.Context) (any, error) {
		val, err := runAffSync(aff, ctx)
		if err != nil {
			return nil, err
		}
		return f(val), nil
	}
}

func _ParAffMap(f func(any) any, aff AffFn) any {
	return internalMap(f, aff)
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

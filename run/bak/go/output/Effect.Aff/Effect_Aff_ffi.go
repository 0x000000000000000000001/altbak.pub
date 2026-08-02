package Effect_Aff

import "gopurs/output/gopurs_runtime"

import (
	"context"
	"fmt"
	"time"
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
			return nil, ctx.Err()
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
			return nil, ctx.Err()
		}
	}
}

func _LiftEffect(eff func(any) any) any {
	return func(ctx context.Context) (any, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
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
			cancelFnAff := canceler(fmt.Errorf("context canceled"))
			go runAffSync(cancelFnAff, context.Background())
			return nil, ctx.Err()
		}
	}
}

func makeFiberNative(aff AffFn) map[string]any {
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan struct {
		val any
		err error
	}, 1)

	go func() {
		val, err := runAffSync(aff, ctx)
		resultChan <- struct {
			val any
			err error
		}{val, err}
	}()

	fiber := map[string]any{
		"run": func(_ any) any { return nil },
		"kill": func(errAny any, onError func(any) any, onSuccess func(any) any) any {
			return func(_ any) any {
				cancel()
				return func(_ any) any {
					res := <-resultChan
					if res.err != nil {
						return onError(res.err)
					}
					return onSuccess(res.val)
				}
			}
		},
		"join": func(onError func(any) any, onSuccess func(any) any) any {
			return func(_ any) any {
				return func(_ any) any {
					res := <-resultChan
					if res.err != nil {
						return onError(res.err)
					}
					return onSuccess(res.val)
				}
			}
		},
		"onComplete": func(onComplete func(any) any) any {
			return func(_ any) any {
				return func(_ any) any {
					return nil
				}
			}
		},
		"isSuspended": func(_ any) bool { return false },
	}
	return fiber
}

func _MakeFiber(aff AffFn) any {
	return func(_ any) any {
		return makeFiberNative(aff)
	}
}

func _Fork(isSuspended any, aff AffFn) any {
    return func(ctx context.Context) (any, error) {
        fiber := makeFiberNative(aff)
        return fiber, nil
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

func _ParAffMap(_ any, _ any) any { panic("Not implemented") }
func _ParAffApply(_ any, _ any) any { panic("Not implemented") }
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
			select {
			case res := <-resCh:
				if res.err == nil {
					return res.val, nil
				}
				if firstErr == nil {
					firstErr = res.err
				} else {
					return nil, firstErr
				}
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		return nil, firstErr
	}
}
func _MakeSupervisedFiber(aff AffFn) any {
	return func(_ any) any {
		panic("Not implemented")
	}
}
func _KillAll(_ any, _ any, _ any) any { panic("Not implemented") }
func _Sequential(aff AffFn) any { return aff }
func GeneralBracket(_ any, _ any, _ any) any { panic("Not implemented") }


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
var _Gopurs__Fork = // TAST: (Func [Boolean, (ADT ["Effect","Aff","Aff"] [(TypeVar a)])] (ADT ["Effect","Aff","Aff"] [Any]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[AffFn](arg1)
	go_res := _Fork(go_arg0, go_arg1)
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
var _Gopurs__MakeFiber = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)])] (ADT ["Effect","Effect"] [Any]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_res := _MakeFiber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__MakeSupervisedFiber = // TAST: (Func [(ADT ["Effect","Aff","Aff"] [(TypeVar a)])] (ADT ["Effect","Effect"] [Any]))
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
var _Gopurs__ParAffAlt = // TAST: (Func [(ADT ["Effect","Aff","ParAff"] [(TypeVar a)]), (ADT ["Effect","Aff","ParAff"] [(TypeVar a)])] (ADT ["Effect","Aff","ParAff"] [(TypeVar a)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[AffFn](arg0)
	go_arg1 := gopurs_runtime.Unbox[AffFn](arg1)
	go_res := _ParAffAlt(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ParAffApply = // TAST: (Func [(ADT ["Effect","Aff","ParAff"] [(Func [(TypeVar a)] (TypeVar b))]), (ADT ["Effect","Aff","ParAff"] [(TypeVar a)])] (ADT ["Effect","Aff","ParAff"] [(TypeVar b)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffApply(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ParAffMap = // TAST: (Func [(Func [(TypeVar a)] (TypeVar b)), (ADT ["Effect","Aff","ParAff"] [(TypeVar a)])] (ADT ["Effect","Aff","ParAff"] [(TypeVar b)]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffMap(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Pure = // TAST: (Func [(TypeVar a)] (ADT ["Effect","Aff","Aff"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Pure(go_arg0)
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
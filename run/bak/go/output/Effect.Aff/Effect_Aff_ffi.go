package Effect_Aff

import "gopurs/output/gopurs_runtime"

import (
	"context"
	"fmt"
	"time"
)

func unboxAff(aff interface{}) func(context.Context) (interface{}, error) {
	if fn, ok := aff.(func(context.Context) (interface{}, error)); ok {
		return fn
	}
	val := aff.(gopurs_runtime.Value)
	return (*(*interface{})(val.UnsafePtr)).(func(context.Context) (interface{}, error))
}

type BindNode struct {
	Aff interface{}
	K   func(interface{}) interface{}
}

func runAffSync(aff func(context.Context) (interface{}, error), ctx context.Context) (interface{}, error) {
	var current = aff
	var stack []func(interface{}) interface{}

	for {
		val, err := current(ctx)
		if err != nil {
			return nil, err
		}

		if node, ok := val.(BindNode); ok {
			stack = append(stack, node.K)
			current = unboxAff(node.Aff)
		} else {
			if len(stack) > 0 {
				k := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				current = unboxAff(k(val))
			} else {
				return val, nil
			}
		}
	}
}

type Aff func(ctx context.Context) (interface{}, error)

func _Pure(val interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return val, nil
		}
	}
}

func _Bind(aff interface{}, k func(interface{}) interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		return BindNode{Aff: aff, K: k}, nil
	}
}

func _Delay(right interface{}, ms float64) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		duration := time.Duration(ms) * time.Millisecond
		timer := time.NewTimer(duration)
		defer timer.Stop()

		select {
		case <-timer.C:
			return nil, nil // PureScript Unit is usually represented as nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func _LiftEffect(eff func(interface{}) interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return eff(nil), nil
		}
	}
}

func MakeAff(build interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resultChan := make(chan struct {
			val interface{}
			err error
		}, 1)

		callback := gopurs_runtime.Func(func(either gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				if either.IntVal == 3711209382 { // Left
					errVal := (*struct{Rc uint32; Value0 gopurs_runtime.Value})(either.UnsafePtr).Value0
					resultChan <- struct{val interface{}; err error}{nil, fmt.Errorf("Aff Error: %+v", errVal)}
				} else { // Right
					val := (*struct{Rc uint32; Value0 gopurs_runtime.Value})(either.UnsafePtr).Value0
					resultChan <- struct{val interface{}; err error}{val, nil}
				}
				return gopurs_runtime.Value{}
			})
		})

		buildVal := build.(gopurs_runtime.Value)
		cancelerEffectVal := gopurs_runtime.Apply(buildVal, callback)
		cancelerVal := gopurs_runtime.Apply(cancelerEffectVal, gopurs_runtime.Value{})

		select {
		case res := <-resultChan:
			return res.val, res.err
		case <-ctx.Done():
			cancelFn := gopurs_runtime.Apply(cancelerVal, gopurs_runtime.Box(fmt.Errorf("context canceled")))
			gopurs_runtime.Apply(cancelFn, gopurs_runtime.Value{})
			return nil, ctx.Err()
		}
	}
}

func _MakeFiber(ffiUtil interface{}, aff interface{}, _ interface{}) interface{} {
	ctx, cancel := context.WithCancel(context.Background())
	resultChan := make(chan struct {
		val interface{}
		err error
	}, 1)

	go func() {
		val, err := runAffSync(unboxAff(aff), ctx)
		resultChan <- struct {
			val interface{}
			err error
		}{val, err}
	}()

	// Return the Fiber record as expected by PureScript
	fiber := map[string]interface{}{
		"run": func(_ interface{}) interface{} {
			return nil
		},
		"kill": func(err interface{}) interface{} {
			return func(k interface{}) interface{} {
				return func(_ interface{}) interface{} {
					cancel()
					return func(_ interface{}) interface{} {
						res := <-resultChan
						return k.(func(interface{}) interface{})(res.val).(func(interface{}) interface{})(nil)
					}
				}
			}
		},
		"join": func(k interface{}) interface{} {
			return func(_ interface{}) interface{} {
				return func(_ interface{}) interface{} {
					res := <-resultChan
					return k.(func(interface{}) interface{})(res.val).(func(interface{}) interface{})(nil)
				}
			}
		},
		"onComplete": func(onComplete interface{}) interface{} {
			return func(_ interface{}) interface{} {
				return func(_ interface{}) interface{} {
					return nil
				}
			}
		},
		"isSuspended": func(_ interface{}) interface{} {
			return false
		},
	}
	return fiber
}

func _Fork(isSuspended interface{}, aff interface{}) interface{} {
    // forkAff :: forall a. Aff a -> Aff (Fiber a)
    // _fork uses _MakeFiber internally in purescript, but the FFI signature for _fork is:
    // foreign import _fork :: forall a. Boolean -> Aff a -> Aff (Fiber a)
    return func(ctx context.Context) (interface{}, error) {
        // Just call _MakeFiber with nil ffiUtil
        fiber := _MakeFiber(nil, aff, nil)
        return fiber, nil
    }
}

func _ThrowError(err interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		if val, ok := err.(gopurs_runtime.Value); ok {
			if val.Type == 13 {
				if e, ok := (*(*interface{})(val.UnsafePtr)).(error); ok {
					return nil, e
				}
			}
		}
		if e, ok := err.(error); ok {
			return nil, e
		}
		return nil, fmt.Errorf("%v", err)
	}
}

func _CatchError(aff interface{}, handler func(interface{}) interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		val, err := runAffSync(unboxAff(aff), ctx)
		if err != nil {
			return runAffSync(unboxAff(handler(err)), ctx)
		}
		return val, nil
	}
}

func _Map(f func(interface{}) interface{}, aff interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		val, err := runAffSync(unboxAff(aff), ctx)
		if err != nil {
			return nil, err
		}
		return f(val), nil
	}
}

func _ParAffMap(_ interface{}, _ interface{}) interface{} { panic("Not implemented: _parAffMap") }
func _ParAffApply(_ interface{}, _ interface{}) interface{} { panic("Not implemented: _parAffApply") }
func _ParAffAlt(aff1 interface{}, aff2 interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		fn1 := unboxAff(aff1)
		fn2 := unboxAff(aff2)

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		type Result struct {
			val interface{}
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
func _MakeSupervisedFiber(_ interface{}, _ interface{}) interface{} { panic("Not implemented: _makeSupervisedFiber") }
func _KillAll(_ interface{}, _ interface{}, _ interface{}) interface{} { panic("Not implemented: _killAll") }
func _Sequential(aff interface{}) interface{} { return aff }
func GeneralBracket(_ interface{}, _ interface{}, _ interface{}) interface{} { panic("Not implemented: generalBracket") }


// --- Auto-generated FFI wrappers ---
func Call__Pure(arg0 interface{}) interface{} {
	return _Pure(arg0)
}
var _Gopurs__Pure = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Pure(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__Bind(arg0 interface{}, arg1 func(interface{}) interface{}) interface{} {
	return _Bind(arg0, arg1)
}
var _Gopurs__Bind = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_res := _Bind(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__Delay(arg0 interface{}, arg1 float64) interface{} {
	return _Delay(arg0, arg1)
}
var _Gopurs__Delay = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := _Delay(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__LiftEffect(arg0 func(interface{}) interface{}) interface{} {
	return _LiftEffect(arg0)
}
var _Gopurs__LiftEffect = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := _LiftEffect(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_makeAff(arg0 interface{}) interface{} {
	return MakeAff(arg0)
}
var _Gopurs_MakeAff = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MakeAff(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__MakeFiber(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return _MakeFiber(arg0, arg1, arg2)
}
var _Gopurs__MakeFiber = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := _MakeFiber(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call__Fork(arg0 interface{}, arg1 interface{}) interface{} {
	return _Fork(arg0, arg1)
}
var _Gopurs__Fork = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Fork(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__ThrowError(arg0 interface{}) interface{} {
	return _ThrowError(arg0)
}
var _Gopurs__ThrowError = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _ThrowError(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__CatchError(arg0 interface{}, arg1 func(interface{}) interface{}) interface{} {
	return _CatchError(arg0, arg1)
}
var _Gopurs__CatchError = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_res := _CatchError(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__Map(arg0 func(interface{}) interface{}, arg1 interface{}) interface{} {
	return _Map(arg0, arg1)
}
var _Gopurs__Map = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_res := _Map(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__ParAffMap(arg0 interface{}, arg1 interface{}) interface{} {
	return _ParAffMap(arg0, arg1)
}
var _Gopurs__ParAffMap = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffMap(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__ParAffApply(arg0 interface{}, arg1 interface{}) interface{} {
	return _ParAffApply(arg0, arg1)
}
var _Gopurs__ParAffApply = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffApply(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__ParAffAlt(arg0 interface{}, arg1 interface{}) interface{} {
	return _ParAffAlt(arg0, arg1)
}
var _Gopurs__ParAffAlt = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffAlt(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__MakeSupervisedFiber(arg0 interface{}, arg1 interface{}) interface{} {
	return _MakeSupervisedFiber(arg0, arg1)
}
var _Gopurs__MakeSupervisedFiber = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _MakeSupervisedFiber(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__KillAll(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return _KillAll(arg0, arg1, arg2)
}
var _Gopurs__KillAll = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := _KillAll(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call__Sequential(arg0 interface{}) interface{} {
	return _Sequential(arg0)
}
var _Gopurs__Sequential = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Sequential(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_generalBracket(arg0 interface{}, arg1 interface{}, arg2 interface{}) interface{} {
	return GeneralBracket(arg0, arg1, arg2)
}
var _Gopurs_GeneralBracket = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := GeneralBracket(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})

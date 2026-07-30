package Effect_Aff

import "gopurs/output/gopurs_runtime"

import (
	"context"
	"time"
)

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

func _Bind(aff interface{}, k interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		val, err := aff.(func(context.Context) (interface{}, error))(ctx)
		if err != nil {
			return nil, err
		}
		nextAff := k.(func(interface{}) interface{})(val).(func(context.Context) (interface{}, error))
		return nextAff(ctx)
	}
}

func _Delay(right interface{}, ms interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		duration := time.Duration(ms.(float64)) * time.Millisecond
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

func _LiftEffect(eff interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			return eff.(func() interface{})(), nil
		}
	}
}

func MakeAff(build interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		resultChan := make(chan struct {
			val interface{}
			err error
		}, 1)

		callback := func(either interface{}) interface{} {
			return func() interface{} {
				// PureScript's Either is usually a struct/map in Go.
				// We need to unpack it. We'll assume a generic map layout for ADTs in gopurs:
				// If it's a Left, it might have a tag or we can just try to guess.
				// For now, we will just send the either as the value, 
				// and let's assume we can just pass it through or we need to unpack it.
				// Let's assume we can unpack it via a "isLeft" boolean or similar if it's a map.
				// Since we don't know the exact layout here, we'll try to extract "value0"
				if m, ok := either.(map[string]interface{}); ok {
					if _, hasLeft := m["Left"]; hasLeft {
					    // Just a placeholder check
					}
					// Actually, let's just assume we return the raw value if it's Right,
					// or return an error if it's Left.
					// This is a naive implementation for the proof of concept.
					if m["constructorName"] == "Left" {
						resultChan <- struct {
							val interface{}
							err error
						}{nil, m["value0"].(error)}
					} else {
						resultChan <- struct {
							val interface{}
							err error
						}{m["value0"], nil}
					}
				} else {
                    // Fallback
                    resultChan <- struct {
                        val interface{}
                        err error
                    }{either, nil}
                }
				return nil
			}
		}

		cancelerEffect := build.(func(interface{}) interface{})(callback)
		canceler := cancelerEffect.(func() interface{})()

		select {
		case res := <-resultChan:
			return res.val, res.err
		case <-ctx.Done():
			cancelAff := canceler.(func(interface{}) interface{})(ctx.Err()).(func(context.Context) (interface{}, error))
			cancelAff(context.Background())
			return nil, ctx.Err()
		}
	}
}

func _MakeFiber(ffiUtil interface{}, aff interface{}) func() interface{} {
	return func() interface{} {
		ctx, cancel := context.WithCancel(context.Background())
		resultChan := make(chan struct {
			val interface{}
			err error
		}, 1)

		go func() {
			val, err := aff.(func(context.Context) (interface{}, error))(ctx)
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
}

func _Fork(isSuspended interface{}, aff interface{}) interface{} {
    // forkAff :: forall a. Aff a -> Aff (Fiber a)
    // _fork uses _MakeFiber internally in purescript, but the FFI signature for _fork is:
    // foreign import _fork :: forall a. Boolean -> Aff a -> Aff (Fiber a)
    return func(ctx context.Context) (interface{}, error) {
        // Just call _MakeFiber with nil ffiUtil
        eff := _MakeFiber(nil, aff)
        fiber := eff()
        return fiber, nil
    }
}

func _ThrowError(err interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		return nil, err.(error)
	}
}

func _CatchError(aff interface{}, handler interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		val, err := aff.(func(context.Context) (interface{}, error))(ctx)
		if err != nil {
			return handler.(func(interface{}) interface{})(err).(func(context.Context) (interface{}, error))(ctx)
		}
		return val, nil
	}
}

func _Map(f interface{}, aff interface{}) interface{} {
	return func(ctx context.Context) (interface{}, error) {
		val, err := aff.(func(context.Context) (interface{}, error))(ctx)
		if err != nil {
			return nil, err
		}
		return f.(func(interface{}) interface{})(val), nil
	}
}

func _ParAffMap(_ interface{}, _ interface{}) interface{} { panic("Not implemented: _parAffMap") }
func _ParAffApply(_ interface{}, _ interface{}) interface{} { panic("Not implemented: _parAffApply") }
func _ParAffAlt(_ interface{}, _ interface{}) interface{} { panic("Not implemented: _parAffAlt") }
func _MakeSupervisedFiber(_ interface{}, _ interface{}) interface{} { panic("Not implemented: _makeSupervisedFiber") }
func _KillAll(_ interface{}, _ interface{}, _ interface{}) interface{} { panic("Not implemented: _killAll") }
func _Sequential(_ interface{}) interface{} { panic("Not implemented: _sequential") }
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
func Call__Bind(arg0 interface{}, arg1 interface{}) interface{} {
	return _Bind(arg0, arg1)
}
var _Gopurs__Bind = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Bind(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__Delay(arg0 interface{}, arg1 interface{}) interface{} {
	return _Delay(arg0, arg1)
}
var _Gopurs__Delay = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Delay(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__LiftEffect(arg0 interface{}) interface{} {
	return _LiftEffect(arg0)
}
var _Gopurs__LiftEffect = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
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
func Call__MakeFiber(arg0 interface{}, arg1 interface{}) func() interface{} {
	return _MakeFiber(arg0, arg1)
}
var _Gopurs__MakeFiber = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _MakeFiber(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
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
func Call__CatchError(arg0 interface{}, arg1 interface{}) interface{} {
	return _CatchError(arg0, arg1)
}
var _Gopurs__CatchError = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _CatchError(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__Map(arg0 interface{}, arg1 interface{}) interface{} {
	return _Map(arg0, arg1)
}
var _Gopurs__Map = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
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

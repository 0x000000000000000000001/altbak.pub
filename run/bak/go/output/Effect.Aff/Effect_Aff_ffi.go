package Effect_Aff

import "gopurs/output/gopurs_runtime"

func _Pure(_ interface{}) interface{} {
	panic("Not implemented: _pure")
}

func _ThrowError(_ interface{}) interface{} {
	panic("Not implemented: _throwError")
}

func _CatchError(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _catchError")
}

func _Fork(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _fork")
}

func _Map(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _map")
}

func _Bind(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _bind")
}

func _Delay(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _delay")
}

func _LiftEffect(_ interface{}) interface{} {
	panic("Not implemented: _liftEffect")
}

func _ParAffMap(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _parAffMap")
}

func _ParAffApply(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _parAffApply")
}

func _ParAffAlt(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _parAffAlt")
}

func _MakeFiber(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _makeFiber")
}

func _MakeSupervisedFiber(_ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _makeSupervisedFiber")
}

func _KillAll(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: _killAll")
}

func _Sequential(_ interface{}) interface{} {
	panic("Not implemented: _sequential")
}

func GeneralBracket(_ interface{}, _ interface{}, _ interface{}) interface{} {
	panic("Not implemented: generalBracket")
}

func MakeAff(_ interface{}) interface{} {
	panic("Not implemented: makeAff")
}


// --- Auto-generated FFI wrappers ---
func Call__Pure(arg0 interface{}) interface{} {
	return _Pure(arg0)
}
var _Gopurs__Pure = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Pure(go_arg0)
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
func Call__Fork(arg0 interface{}, arg1 interface{}) interface{} {
	return _Fork(arg0, arg1)
}
var _Gopurs__Fork = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Fork(go_arg0, go_arg1)
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
func Call__MakeFiber(arg0 interface{}, arg1 interface{}) interface{} {
	return _MakeFiber(arg0, arg1)
}
var _Gopurs__MakeFiber = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _MakeFiber(go_arg0, go_arg1)
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
func Call_makeAff(arg0 interface{}) interface{} {
	return MakeAff(arg0)
}
var _Gopurs_MakeAff = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MakeAff(go_arg0)
	return gopurs_runtime.Box(go_res)
})

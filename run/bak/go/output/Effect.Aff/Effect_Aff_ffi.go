package Effect_Aff

import "gopurs/output/gopurs_runtime"

func _Pure(_ any) any {
	panic("Not implemented: _pure")
}

func _ThrowError(_ any) any {
	panic("Not implemented: _throwError")
}

func _CatchError(_ any, _ any) any {
	panic("Not implemented: _catchError")
}

func _Fork(_ any, _ any) any {
	panic("Not implemented: _fork")
}

func _Map(_ any, _ any) any {
	panic("Not implemented: _map")
}

func _Bind(_ any, _ any) any {
	panic("Not implemented: _bind")
}

func _Delay(_ any, _ any) any {
	panic("Not implemented: _delay")
}

func _LiftEffect(_ any) any {
	panic("Not implemented: _liftEffect")
}

func _ParAffMap(_ any, _ any) any {
	panic("Not implemented: _parAffMap")
}

func _ParAffApply(_ any, _ any) any {
	panic("Not implemented: _parAffApply")
}

func _ParAffAlt(_ any, _ any) any {
	panic("Not implemented: _parAffAlt")
}

func _MakeFiber(_ any, _ any) any {
	panic("Not implemented: _makeFiber")
}

func _MakeSupervisedFiber(_ any, _ any) any {
	panic("Not implemented: _makeSupervisedFiber")
}

func _KillAll(_ any, _ any, _ any) any {
	panic("Not implemented: _killAll")
}

func _Sequential(_ any) any {
	panic("Not implemented: _sequential")
}

func GeneralBracket(_ any, _ any, _ any) any {
	panic("Not implemented: generalBracket")
}

func MakeAff(_ any) any {
	panic("Not implemented: makeAff")
}


// --- Auto-generated FFI wrappers ---
func Call__Pure(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Pure(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__Pure = gopurs_runtime.Func(Call__Pure)
func Call__ThrowError(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _ThrowError(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__ThrowError = gopurs_runtime.Func(Call__ThrowError)
func Call__CatchError(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _CatchError(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__CatchError = gopurs_runtime.Func2(Call__CatchError)
func Call__Fork(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Fork(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__Fork = gopurs_runtime.Func2(Call__Fork)
func Call__Map(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Map(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__Map = gopurs_runtime.Func2(Call__Map)
func Call__Bind(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Bind(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__Bind = gopurs_runtime.Func2(Call__Bind)
func Call__Delay(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _Delay(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__Delay = gopurs_runtime.Func2(Call__Delay)
func Call__LiftEffect(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _LiftEffect(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__LiftEffect = gopurs_runtime.Func(Call__LiftEffect)
func Call__ParAffMap(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffMap(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__ParAffMap = gopurs_runtime.Func2(Call__ParAffMap)
func Call__ParAffApply(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffApply(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__ParAffApply = gopurs_runtime.Func2(Call__ParAffApply)
func Call__ParAffAlt(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _ParAffAlt(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__ParAffAlt = gopurs_runtime.Func2(Call__ParAffAlt)
func Call__MakeFiber(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _MakeFiber(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__MakeFiber = gopurs_runtime.Func2(Call__MakeFiber)
func Call__MakeSupervisedFiber(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_res := _MakeSupervisedFiber(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__MakeSupervisedFiber = gopurs_runtime.Func2(Call__MakeSupervisedFiber)
func Call__KillAll(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := _KillAll(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__KillAll = gopurs_runtime.Func3(Call__KillAll)
func Call__Sequential(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Sequential(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__Sequential = gopurs_runtime.Func(Call__Sequential)
func Call_generalBracket(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_res := GeneralBracket(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_GeneralBracket = gopurs_runtime.Func3(Call_generalBracket)
func Call_makeAff(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := MakeAff(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_MakeAff = gopurs_runtime.Func(Call_makeAff)

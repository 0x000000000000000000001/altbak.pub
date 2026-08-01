package Data_String_CodePoints

import "gopurs/output/gopurs_runtime"

func _UnsafeCodePointAt0(_ interface{}) interface{} {
	panic("Not implemented: X_unsafeCodePointAt0")
}

func _CodePointAt(_ interface{}) interface{} {
	panic("Not implemented: X_codePointAt")
}

func _CountPrefix(_ interface{}) interface{} {
	panic("Not implemented: X_countPrefix")
}

func _FromCodePointArray(_ interface{}) interface{} {
	panic("Not implemented: X_fromCodePointArray")
}

func _Singleton(_ interface{}) interface{} {
	panic("Not implemented: X_singleton")
}

func _Take(_ interface{}) interface{} {
	panic("Not implemented: X_take")
}

func _ToCodePointArray(_ interface{}) interface{} {
	panic("Not implemented: X_toCodePointArray")
}


// --- Auto-generated FFI wrappers ---
func Call__UnsafeCodePointAt0(arg0 interface{}) interface{} {
	return _UnsafeCodePointAt0(arg0)
}
var _Gopurs__UnsafeCodePointAt0 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _UnsafeCodePointAt0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__CodePointAt(arg0 interface{}) interface{} {
	return _CodePointAt(arg0)
}
var _Gopurs__CodePointAt = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _CodePointAt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__CountPrefix(arg0 interface{}) interface{} {
	return _CountPrefix(arg0)
}
var _Gopurs__CountPrefix = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _CountPrefix(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__FromCodePointArray(arg0 interface{}) interface{} {
	return _FromCodePointArray(arg0)
}
var _Gopurs__FromCodePointArray = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _FromCodePointArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__Singleton(arg0 interface{}) interface{} {
	return _Singleton(arg0)
}
var _Gopurs__Singleton = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Singleton(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__Take(arg0 interface{}) interface{} {
	return _Take(arg0)
}
var _Gopurs__Take = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _Take(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call__ToCodePointArray(arg0 interface{}) interface{} {
	return _ToCodePointArray(arg0)
}
var _Gopurs__ToCodePointArray = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := _ToCodePointArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})

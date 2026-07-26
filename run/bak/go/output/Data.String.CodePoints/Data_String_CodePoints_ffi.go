package Data_String_CodePoints

import "gopurs/output/gopurs_runtime"

func X_unsafeCodePointAt0(_ interface{}) interface{} {
	panic("Not implemented: X_unsafeCodePointAt0")
}

func X_codePointAt(_ interface{}) interface{} {
	panic("Not implemented: X_codePointAt")
}

func X_countPrefix(_ interface{}) interface{} {
	panic("Not implemented: X_countPrefix")
}

func X_fromCodePointArray(_ interface{}) interface{} {
	panic("Not implemented: X_fromCodePointArray")
}

func X_singleton(_ interface{}) interface{} {
	panic("Not implemented: X_singleton")
}

func X_take(_ interface{}) interface{} {
	panic("Not implemented: X_take")
}

func X_toCodePointArray(_ interface{}) interface{} {
	panic("Not implemented: X_toCodePointArray")
}


// --- Auto-generated FFI wrappers ---
func Call_x_unsafeCodePointAt0(arg0 interface{}) interface{} {
	return X_unsafeCodePointAt0(arg0)
}
var _Gopurs_X_unsafeCodePointAt0 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_unsafeCodePointAt0(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_x_codePointAt(arg0 interface{}) interface{} {
	return X_codePointAt(arg0)
}
var _Gopurs_X_codePointAt = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_codePointAt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_x_countPrefix(arg0 interface{}) interface{} {
	return X_countPrefix(arg0)
}
var _Gopurs_X_countPrefix = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_countPrefix(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_x_fromCodePointArray(arg0 interface{}) interface{} {
	return X_fromCodePointArray(arg0)
}
var _Gopurs_X_fromCodePointArray = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_fromCodePointArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_x_singleton(arg0 interface{}) interface{} {
	return X_singleton(arg0)
}
var _Gopurs_X_singleton = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_singleton(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_x_take(arg0 interface{}) interface{} {
	return X_take(arg0)
}
var _Gopurs_X_take = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_take(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_x_toCodePointArray(arg0 interface{}) interface{} {
	return X_toCodePointArray(arg0)
}
var _Gopurs_X_toCodePointArray = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_toCodePointArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})

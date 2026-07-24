package Data_String_CodePoints

import "gopurs/output/gopurs_runtime"

func X_unsafeCodePointAt0(_ any) any {
	panic("Not implemented: X_unsafeCodePointAt0")
}

func X_codePointAt(_ any) any {
	panic("Not implemented: X_codePointAt")
}

func X_countPrefix(_ any) any {
	panic("Not implemented: X_countPrefix")
}

func X_fromCodePointArray(_ any) any {
	panic("Not implemented: X_fromCodePointArray")
}

func X_singleton(_ any) any {
	panic("Not implemented: X_singleton")
}

func X_take(_ any) any {
	panic("Not implemented: X_take")
}

func X_toCodePointArray(_ any) any {
	panic("Not implemented: X_toCodePointArray")
}


// --- Auto-generated FFI wrappers ---
func Call_x_unsafeCodePointAt0(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_unsafeCodePointAt0(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_X_unsafeCodePointAt0 = gopurs_runtime.Func(Call_x_unsafeCodePointAt0)
func Call_x_codePointAt(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_codePointAt(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_X_codePointAt = gopurs_runtime.Func(Call_x_codePointAt)
func Call_x_countPrefix(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_countPrefix(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_X_countPrefix = gopurs_runtime.Func(Call_x_countPrefix)
func Call_x_fromCodePointArray(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_fromCodePointArray(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_X_fromCodePointArray = gopurs_runtime.Func(Call_x_fromCodePointArray)
func Call_x_singleton(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_singleton(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_X_singleton = gopurs_runtime.Func(Call_x_singleton)
func Call_x_take(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_take(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_X_take = gopurs_runtime.Func(Call_x_take)
func Call_x_toCodePointArray(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := X_toCodePointArray(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_X_toCodePointArray = gopurs_runtime.Func(Call_x_toCodePointArray)

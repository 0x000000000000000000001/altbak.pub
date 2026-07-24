package Data_Show

import "gopurs/output/gopurs_runtime"

import "fmt"
func ShowIntImpl(n int) string {
	return fmt.Sprintf("%v", n)
}
func ShowNumberImpl(n float64) string {
	return fmt.Sprintf("%f", n)
}
func ShowCharImpl(c string) string {
	return fmt.Sprintf("'%s'", c)
}
func ShowStringImpl(s string) string {
	return fmt.Sprintf("%q", s)
}
func ShowArrayImpl(f func(any) string, arr []any) string {
	res := "["
	for i, v := range arr {
		if i > 0 {
			res += ","
		}
		res += f(v)
	}
	res += "]"
	return res
}


// --- Auto-generated FFI wrappers ---
func Call_showIntImpl(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := ShowIntImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_ShowIntImpl = gopurs_runtime.Func(Call_showIntImpl)
func Call_showNumberImpl(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := ShowNumberImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_ShowNumberImpl = gopurs_runtime.Func(Call_showNumberImpl)
func Call_showCharImpl(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ShowCharImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_ShowCharImpl = gopurs_runtime.Func(Call_showCharImpl)
func Call_showStringImpl(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ShowStringImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_ShowStringImpl = gopurs_runtime.Func(Call_showStringImpl)
func Call_showArrayImpl(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) string {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[string](inner_res0)
		}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := ShowArrayImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_ShowArrayImpl = gopurs_runtime.Func2(Call_showArrayImpl)

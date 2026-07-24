package Data_Number_Format

import "gopurs/output/gopurs_runtime"

import "fmt"
func ToExponentialNative(fractionDigits int, n float64) string { return fmt.Sprintf("%.*e", fractionDigits, n) }
func ToFixedNative(fractionDigits int, n float64) string { return fmt.Sprintf("%.*f", fractionDigits, n) }
func ToPrecisionNative(precision int, n float64) string { return fmt.Sprintf("%.*g", precision, n) }
func ToString(n float64) string { return fmt.Sprintf("%v", n) }


// --- Auto-generated FFI wrappers ---
func Call_toExponentialNative(arg0 int, arg1 float64) string { return fmt.Sprintf("%.*e", fractionDigits, n) } {
	return ToExponentialNative(arg0, arg1)
}
var _Gopurs_ToExponentialNative = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := ToExponentialNative(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_toFixedNative(arg0 int, arg1 float64) string { return fmt.Sprintf("%.*f", fractionDigits, n) } {
	return ToFixedNative(arg0, arg1)
}
var _Gopurs_ToFixedNative = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := ToFixedNative(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_toPrecisionNative(arg0 int, arg1 float64) string { return fmt.Sprintf("%.*g", precision, n) } {
	return ToPrecisionNative(arg0, arg1)
}
var _Gopurs_ToPrecisionNative = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := ToPrecisionNative(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_toString(arg0 float64) string { return fmt.Sprintf("%v", n) } {
	return ToString(arg0)
}
var _Gopurs_ToString = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := ToString(go_arg0)
	return gopurs_runtime.Box(go_res)
})

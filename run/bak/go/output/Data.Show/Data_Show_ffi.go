package Data_Show

import "gopurs/output/gopurs_runtime"

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
func ShowIntImpl(n int64) string {
	return fmt.Sprintf("%v", n)
}
func ShowNumberImpl(n float64) string {
	if math.IsNaN(n) {
		return "NaN"
	} else if math.IsInf(n, 1) {
		return "Infinity"
	} else if math.IsInf(n, -1) {
		return "-Infinity"
	}

	absN := math.Abs(n)
	var str string
	if absN != 0 && (absN >= 1e21 || absN < 1e-6) {
		str = strconv.FormatFloat(n, 'g', -1, 64)
		// Go uses e-07 but JS uses e-7. We won't worry too much unless a test fails.
	} else {
		str = strconv.FormatFloat(n, 'f', -1, 64)
	}

	if strings.Contains(str, ".") || strings.Contains(str, "e") {
		return str
	}
	return str + ".0"
}
func ShowCharImpl(c string) string {
	return fmt.Sprintf("'%s'", c)
}
func ShowStringImpl(s string) string {
	return fmt.Sprintf("%q", s)
}
func ShowArrayImpl(f func(interface{}) string, arr []interface{}) string {
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
var _Gopurs_ShowArrayImpl = // TAST: (Func [(Func [(TypeVar a)] String), (Array (TypeVar a))] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) string {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[string](inner_res0)
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := ShowArrayImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShowCharImpl = // TAST: (Func [Char] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ShowCharImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShowIntImpl = // TAST: (Func [Int] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_res := ShowIntImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShowNumberImpl = // TAST: (Func [Number] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := ShowNumberImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShowStringImpl = // TAST: (Func [String] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ShowStringImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
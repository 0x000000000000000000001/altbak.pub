package Data_String_CodeUnits

import "gopurs/output/gopurs_runtime"

import (
	"strings"
)

func FromCharArray(a []string) string {
	var b strings.Builder
	for _, v := range a {
		b.WriteString(v)
	}
	return b.String()
}

func ToCharArray(str string) []string {
	arr := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = string(str[i])
	}
	return arr
}

func Singleton(c string) string {
	return c
}

func _CharAt(just func(string) interface{}, nothing interface{}, idx int, str string) interface{} {
	if idx >= 0 && idx < len(str) {
		return just(string(str[idx]))
	}
	return nothing
}

func _ToChar(just func(string) interface{}, nothing interface{}, str string) interface{} {
	if len(str) == 1 {
		return just(str)
	}
	return nothing
}

func Length(s string) int {
	return len(s)
}

func CountPrefix(p func(string) bool, str string) int {
	i := 0
	for i < len(str) {
		if p(string(str[i])) {
			i++
		} else {
			break
		}
	}
	return i
}

func _IndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	idx := strings.Index(s, x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func _IndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	if startIdx < 0 || startIdx > len(str) {
		return nothing
	}
	idx := strings.Index(str[startIdx:], x)
	if idx == -1 {
		return nothing
	}
	return just(idx + startIdx)
}

func _LastIndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	idx := strings.LastIndex(s, x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func _LastIndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	if startIdx < 0 || startIdx >= len(str) {
		return nothing
	}
	idx := strings.LastIndex(str[:startIdx+len(x)], x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func Take(idx int, str string) string {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	return str[:idx]
}

func Drop(idx int, str string) string {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	return str[idx:]
}

func Slice(start int, end int, str string) string {
	if start < 0 {
		start = len(str) + start
	}
	if end < 0 {
		end = len(str) + end
	}
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > len(str) {
		start = len(str)
	}
	if end > len(str) {
		end = len(str)
	}
	if start > end {
		return ""
	}
	return str[start:end]
}

func SplitAt(idx int, str string) map[string]interface{} {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	rec := make(map[string]interface{})
	rec["before"] = str[:idx]
	rec["after"] = str[idx:]
	return rec
}


// --- Auto-generated FFI wrappers ---
var _Gopurs__CharAt = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), Int, String] (ADT ["Data","Maybe","Maybe"] [Char]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _CharAt(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__IndexOf = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), String, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _IndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__IndexOfStartingAt = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), String, Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := _IndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__LastIndexOf = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), String, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _LastIndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__LastIndexOfStartingAt = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), String, Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := _LastIndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ToChar = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), String] (ADT ["Data","Maybe","Maybe"] [Char]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := _ToChar(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_CountPrefix = // TAST: (Func [(Func [Char] Boolean), String] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := CountPrefix(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Drop = // TAST: (Func [Int, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Drop(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FromCharArray = // TAST: (Func [(Array Char)] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]string, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[string](v) }
	go_res := FromCharArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Length = // TAST: (Func [String] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Length(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Singleton = // TAST: (Func [Char] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Singleton(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Slice = // TAST: (Func [Int, Int, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Slice(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_SplitAt = // TAST: (Func [Int, String] (Record [after: String, before: String]))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := SplitAt(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				_raw := go_res
				res_map := make(map[string]gopurs_runtime.Value)
					res_map["after"] = gopurs_runtime.Box(_raw["after"])
					res_map["before"] = gopurs_runtime.Box(_raw["before"])
				return gopurs_runtime.Record(res_map)
			}()
})
var _Gopurs_Take = // TAST: (Func [Int, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Take(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToCharArray = // TAST: (Func [String] (Array Char))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ToCharArray(go_arg0)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
package purescript

import "gopurs/output/gopurs_runtime"

import (
	"strings"
)

func Data_String_CodeUnits_FromCharArray(a []string) string {
	var b strings.Builder
	for _, v := range a {
		b.WriteString(v)
	}
	return b.String()
}

func Data_String_CodeUnits_ToCharArray(str string) []string {
	arr := make([]string, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = string(str[i])
	}
	return arr
}

func Data_String_CodeUnits_Singleton(c string) string {
	return c
}

func Data_String_CodeUnits__CharAt(just func(string) interface{}, nothing interface{}, idx int, str string) interface{} {
	if idx >= 0 && idx < len(str) {
		return just(string(str[idx]))
	}
	return nothing
}

func Data_String_CodeUnits__ToChar(just func(string) interface{}, nothing interface{}, str string) interface{} {
	if len(str) == 1 {
		return just(str)
	}
	return nothing
}

func Data_String_CodeUnits_Length(s string) int {
	return len(s)
}

func Data_String_CodeUnits_CountPrefix(p func(string) bool, str string) int {
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

func Data_String_CodeUnits__IndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	idx := strings.Index(s, x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func Data_String_CodeUnits__IndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	if startIdx < 0 || startIdx > len(str) {
		return nothing
	}
	idx := strings.Index(str[startIdx:], x)
	if idx == -1 {
		return nothing
	}
	return just(idx + startIdx)
}

func Data_String_CodeUnits__LastIndexOf(just func(int) interface{}, nothing interface{}, x string, s string) interface{} {
	idx := strings.LastIndex(s, x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func Data_String_CodeUnits__LastIndexOfStartingAt(just func(int) interface{}, nothing interface{}, x string, startIdx int, str string) interface{} {
	if startIdx < 0 {
		startIdx = 0
	} else if startIdx > len(str) {
		startIdx = len(str)
	}
	end := startIdx + len(x)
	if end > len(str) {
		end = len(str)
	}
	idx := strings.LastIndex(str[:end], x)
	if idx == -1 {
		return nothing
	}
	return just(idx)
}

func Data_String_CodeUnits_Take(idx int, str string) string {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	return str[:idx]
}

func Data_String_CodeUnits_Drop(idx int, str string) string {
	if idx < 0 {
		idx = 0
	}
	if idx > len(str) {
		idx = len(str)
	}
	return str[idx:]
}

func Data_String_CodeUnits_Slice(start int, end int, str string) string {
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

func Data_String_CodeUnits_SplitAt(idx int, str string) map[string]interface{} {
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
var _Gopurs_Data_String_CodeUnits__CharAt = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), Int, String] (ADT ["Data","Maybe","Maybe"] [Char]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_CodeUnits__CharAt(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__IndexOf = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_CodeUnits__IndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__IndexOfStartingAt = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_String_CodeUnits__IndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__LastIndexOf = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := Data_String_CodeUnits__LastIndexOf(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__LastIndexOfStartingAt = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String, Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_String_CodeUnits__LastIndexOfStartingAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits__ToChar = // TAST: (Func [(ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), String] (ADT ["Data","Maybe","Maybe"] [Char]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_CodeUnits__ToChar(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_CountPrefix = // TAST: (Func [(Func [Char] Boolean), String] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 string) bool {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_CountPrefix(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_Drop = // TAST: (Func [Int, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_Drop(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_FromCharArray = // TAST: (Func [(Array Char)] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]string, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[string](v) }
	go_res := Data_String_CodeUnits_FromCharArray(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_Length = // TAST: (Func [String] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_CodeUnits_Length(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_Singleton = // TAST: (Func [Char] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_CodeUnits_Singleton(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_Slice = // TAST: (Func [Int, Int, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_CodeUnits_Slice(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_SplitAt = // TAST: (Func [Int, String] (Record (Row [before: String, after: String] Empty)))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_SplitAt(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				_raw := go_res
				res_map := make(map[string]gopurs_runtime.Value)
					res_map["before"] = gopurs_runtime.Box(_raw["before"])
					res_map["after"] = gopurs_runtime.Box(_raw["after"])
				return gopurs_runtime.Record(res_map)
			}()
})
var _Gopurs_Data_String_CodeUnits_Take = // TAST: (Func [Int, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_CodeUnits_Take(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_CodeUnits_ToCharArray = // TAST: (Func [String] (Array Char))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_CodeUnits_ToCharArray(go_arg0)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
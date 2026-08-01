package Data_String_CodePoints

import "gopurs/output/gopurs_runtime"

func _UnsafeCodePointAt0(fallback interface{}, str string) int64 {
	for _, r := range str {
		return int64(r)
	}
	return 0
}

func _CodePointAt(fallback interface{}, just func(interface{}) interface{}, nothing interface{}, unsafeCodePointAt0 interface{}, index int64, str string) interface{} {
	runes := []rune(str)
	if index < 0 || index >= int64(len(runes)) {
		return nothing
	}
	return just(int64(runes[index]))
}

func _CountPrefix(fallback interface{}, unsafeCodePointAt0 interface{}, pred func(int64) bool, str string) int64 {
	runes := []rune(str)
	count := int64(0)
	for _, r := range runes {
		if !pred(int64(r)) {
			break
		}
		count++
	}
	return count
}

func _FromCodePointArray(singleton interface{}, cps []interface{}) string {
	runes := make([]rune, len(cps))
	for i, cp := range cps {
		switch v := cp.(type) {
		case int64:
			runes[i] = rune(v)
		case float64:
			runes[i] = rune(v)
		case int:
			runes[i] = rune(v)
		}
	}
	return string(runes)
}

func _Singleton(fallback interface{}, cp int64) string {
	return string(rune(cp))
}

func _Take(fallback interface{}, n int64, str string) string {
	runes := []rune(str)
	if n < 0 {
		n = 0
	}
	if n >= int64(len(runes)) {
		return str
	}
	return string(runes[:n])
}

func _ToCodePointArray(fallback interface{}, unsafeCodePointAt0 interface{}, str string) []interface{} {
	runes := []rune(str)
	res := make([]interface{}, len(runes))
	for i, r := range runes {
		res[i] = int64(r)
	}
	return res
}


// --- Auto-generated FFI wrappers ---
func Call__UnsafeCodePointAt0(arg0 interface{}, arg1 string) int64 {
	return _UnsafeCodePointAt0(arg0, arg1)
}
var _Gopurs__UnsafeCodePointAt0 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := _UnsafeCodePointAt0(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__CodePointAt(arg0 interface{}, arg1 func(interface{}) interface{}, arg2 interface{}, arg3 interface{}, arg4 int64, arg5 string) interface{} {
	return _CodePointAt(arg0, arg1, arg2, arg3, arg4, arg5)
}
var _Gopurs__CodePointAt = gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
		}
	go_arg2 := arg2
	go_arg3 := arg3
	go_arg4 := gopurs_runtime.Unbox[int64](arg4)
	go_arg5 := gopurs_runtime.Unbox[string](arg5)
	go_res := _CodePointAt(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})
func Call__CountPrefix(arg0 interface{}, arg1 interface{}, arg2 func(int64) bool, arg3 string) int64 {
	return _CountPrefix(arg0, arg1, arg2, arg3)
}
var _Gopurs__CountPrefix = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := func(p0_0 int64) bool {
			inner_res0 := gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _CountPrefix(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call__FromCodePointArray(arg0 interface{}, arg1 []interface{}) string {
	return _FromCodePointArray(arg0, arg1)
}
var _Gopurs__FromCodePointArray = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := _FromCodePointArray(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__Singleton(arg0 interface{}, arg1 int64) string {
	return _Singleton(arg0, arg1)
}
var _Gopurs__Singleton = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := _Singleton(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call__Take(arg0 interface{}, arg1 int64, arg2 string) string {
	return _Take(arg0, arg1, arg2)
}
var _Gopurs__Take = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := _Take(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call__ToCodePointArray(arg0 interface{}, arg1 interface{}, arg2 string) []interface{} {
	return _ToCodePointArray(arg0, arg1, arg2)
}
var _Gopurs__ToCodePointArray = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := _ToCodePointArray(go_arg0, go_arg1, go_arg2)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})

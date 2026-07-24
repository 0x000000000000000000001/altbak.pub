package Data_String_Common

import "gopurs/output/gopurs_runtime"

import (
	"strings"
)

func _LocaleCompare(lt any, eq any, gt any, s1 string, s2 string) any {
	cmp := strings.Compare(s1, s2)
	if cmp < 0 {
		return lt
	} else if cmp > 0 {
		return gt
	}
	return eq
}

func Replace(s1 string, s2 string, s3 string) string {
	return strings.Replace(s3, s1, s2, 1)
}

func ReplaceAll(s1 string, s2 string, s3 string) string {
	return strings.ReplaceAll(s3, s1, s2)
}

func Split(sep string, s string) []string {
	return strings.Split(s, sep)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func JoinWith(s string, xs []string) string {
	return strings.Join(xs, s)
}


// --- Auto-generated FFI wrappers ---
func Call__LocaleCompare(arg0 any, arg1 any, arg2 any, arg3 string, arg4 string) any {
	return _LocaleCompare(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs__LocaleCompare = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := _LocaleCompare(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_replace(arg0 string, arg1 string, arg2 string) string {
	return Replace(arg0, arg1, arg2)
}
var _Gopurs_Replace = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Replace(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_replaceAll(arg0 string, arg1 string, arg2 string) string {
	return ReplaceAll(arg0, arg1, arg2)
}
var _Gopurs_ReplaceAll = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := ReplaceAll(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_split(arg0 string, arg1 string) []string {
	return Split(arg0, arg1)
}
var _Gopurs_Split = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Split(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
func Call_toLower(arg0 string) string {
	return ToLower(arg0)
}
var _Gopurs_ToLower = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ToLower(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_toUpper(arg0 string) string {
	return ToUpper(arg0)
}
var _Gopurs_ToUpper = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ToUpper(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_trim(arg0 string) string {
	return Trim(arg0)
}
var _Gopurs_Trim = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Trim(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_joinWith(arg0 string, arg1 []string) string {
	return JoinWith(arg0, arg1)
}
var _Gopurs_JoinWith = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	arg1_arr := arg1.PtrVal().([]gopurs_runtime.Value)
	go_arg1 := make([]string, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[string](v) }
	go_res := JoinWith(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

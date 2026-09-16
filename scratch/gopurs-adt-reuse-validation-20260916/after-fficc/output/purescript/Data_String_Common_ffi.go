package purescript

import "gopurs/output/gopurs_runtime"

import (
	"strings"
)

func Data_String_Common__LocaleCompare(lt interface{}, eq interface{}, gt interface{}, s1 string, s2 string) interface{} {
	cmp := strings.Compare(s1, s2)
	if cmp < 0 {
		return lt
	} else if cmp > 0 {
		return gt
	}
	return eq
}

func Data_String_Common_Replace(s1 string, s2 string, s3 string) string {
	return strings.Replace(s3, s1, s2, 1)
}

func Data_String_Common_ReplaceAll(s1 string, s2 string, s3 string) string {
	return strings.ReplaceAll(s3, s1, s2)
}

func Data_String_Common_Split(sep string, s string) []string {
	return strings.Split(s, sep)
}

func Data_String_Common_ToLower(s string) string {
	return strings.ToLower(s)
}

func Data_String_Common_ToUpper(s string) string {
	return strings.ToUpper(s)
}

func Data_String_Common_Trim(s string) string {
	return strings.TrimSpace(s)
}

func Data_String_Common_JoinWith(s string, xs []string) string {
	return strings.Join(xs, s)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_String_Common__LocaleCompare = // TAST: (Func [(ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), (ADT ["Data","Ordering","Ordering"] []), String, String] (ADT ["Data","Ordering","Ordering"] []))
gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := Data_String_Common__LocaleCompare(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Common_JoinWith = // TAST: (Func [String, (Array String)] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]string, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[string](v) }
	go_res := Data_String_Common_JoinWith(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Common_Replace = // TAST: (Func [String, String, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_Common_Replace(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Common_ReplaceAll = // TAST: (Func [String, String, String] String)
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Data_String_Common_ReplaceAll(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Common_Split = // TAST: (Func [String, String] (Array String))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Data_String_Common_Split(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_Data_String_Common_ToLower = // TAST: (Func [String] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_Common_ToLower(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Common_ToUpper = // TAST: (Func [String] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_Common_ToUpper(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Common_Trim = // TAST: (Func [String] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := Data_String_Common_Trim(go_arg0)
	return gopurs_runtime.Box(go_res)
})
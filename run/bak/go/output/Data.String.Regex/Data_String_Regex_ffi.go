package Data_String_Regex

import "gopurs/output/gopurs_runtime"

import "regexp"
func _Match(just func(any) any, nothing any, r *regexp.Regexp, s string) any { return nothing }
func _ReplaceBy(r *regexp.Regexp, f func(string, []any) string, s string) string { return s }
func _Search(just func(any) any, nothing any, r *regexp.Regexp, s string) any { return nothing }

func FlagsImpl(r any) string { return "" }
func RegexImpl(left any, right any, s1 string, s2 string) any { return left }
func Replace(r any, s1 string, s2 string) string { return s2 }
func ShowRegexImpl(r any) string { return "" }
func Source(r any) string { return "" }
func Split(r any, s string) []string { return []string{s} }
func Test(r any, s string) bool { return false }


// --- Auto-generated FFI wrappers ---
var _Gopurs__Match = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*regexp.Regexp](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _Match(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__ReplaceBy = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[*regexp.Regexp](arg0)
	go_arg1 := func(p0_0 string, p0_1 []any) string {
			inner_res0 := gopurs_runtime.Apply2(arg1, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
			return gopurs_runtime.Unbox[string](inner_res0)
		}
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := _ReplaceBy(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs__Search = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[*regexp.Regexp](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := _Search(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FlagsImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := FlagsImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_RegexImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := RegexImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Replace = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_arg2 := gopurs_runtime.Unbox[string](arg2)
	go_res := Replace(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ShowRegexImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := ShowRegexImpl(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Source = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Source(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Split = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Split(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})
var _Gopurs_Test = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := Test(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

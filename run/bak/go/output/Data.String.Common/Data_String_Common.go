package Data_String_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null(s_0_box.StrVal()))
})
	})
	return cache_null
}

var cache_localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		cache_localeCompare = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return func(inner_arg0 string, inner_arg1 string) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Apply3(Get__localeCompare(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})), gopurs_runtime.Str(inner_arg0), gopurs_runtime.Str(inner_arg1))
}(arg0.StrVal(), arg1.StrVal())
})
	})
	return cache_localeCompare
}

var cache__localeCompare gopurs_runtime.Value
var once__localeCompare sync.Once
func Get__localeCompare() gopurs_runtime.Value {
	once__localeCompare.Do(func() {
		cache__localeCompare = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
return _LocaleCompare(arg0, arg1, arg2, arg3.StrVal(), arg4.StrVal())
})
	})
	return cache__localeCompare
}

var cache_joinWith gopurs_runtime.Value
var once_joinWith sync.Once
func Get_joinWith() gopurs_runtime.Value {
	once_joinWith.Do(func() {
		cache_joinWith = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(JoinWith(arg0.StrVal(), func() []string {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))
})
	})
	return cache_joinWith
}

var cache_replace gopurs_runtime.Value
var once_replace sync.Once
func Get_replace() gopurs_runtime.Value {
	once_replace.Do(func() {
		cache_replace = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Replace(arg0.StrVal(), arg1.StrVal(), arg2.StrVal()))
})
	})
	return cache_replace
}

var cache_replaceAll gopurs_runtime.Value
var once_replaceAll sync.Once
func Get_replaceAll() gopurs_runtime.Value {
	once_replaceAll.Do(func() {
		cache_replaceAll = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ReplaceAll(arg0.StrVal(), arg1.StrVal(), arg2.StrVal()))
})
	})
	return cache_replaceAll
}

var cache_split gopurs_runtime.Value
var once_split sync.Once
func Get_split() gopurs_runtime.Value {
	once_split.Do(func() {
		cache_split = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Split(arg0.StrVal(), arg1.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_split
}

var cache_toLower gopurs_runtime.Value
var once_toLower sync.Once
func Get_toLower() gopurs_runtime.Value {
	once_toLower.Do(func() {
		cache_toLower = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ToLower(arg0.StrVal()))
})
	})
	return cache_toLower
}

var cache_toUpper gopurs_runtime.Value
var once_toUpper sync.Once
func Get_toUpper() gopurs_runtime.Value {
	once_toUpper.Do(func() {
		cache_toUpper = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ToUpper(arg0.StrVal()))
})
	})
	return cache_toUpper
}

var cache_trim gopurs_runtime.Value
var once_trim sync.Once
func Get_trim() gopurs_runtime.Value {
	once_trim.Do(func() {
		cache_trim = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Trim(arg0.StrVal()))
})
	})
	return cache_trim
}

func Call_null(s_0_loop string) bool {
var s_0 string = s_0_loop
_ = s_0
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"), gopurs_runtime.Str(s_0), gopurs_runtime.Str("")).IntVal) != (0)
}

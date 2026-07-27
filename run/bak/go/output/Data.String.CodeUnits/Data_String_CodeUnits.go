package Data_String_CodeUnits

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	unsafe "unsafe"
)

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(v_0_box.StrVal()))}
})
	})
	return cache_uncons
}

var cache_toChar gopurs_runtime.Value
var once_toChar sync.Once
func Get_toChar() gopurs_runtime.Value {
	once_toChar.Do(func() {
		cache_toChar = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 string) *pkg_Data_Maybe.Constructor_Just[string] {
return (*pkg_Data_Maybe.Constructor_Just[string])(gopurs_runtime.Apply(gopurs_runtime.Apply2(Get__toChar(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), gopurs_runtime.Str(inner_arg0)).UnsafePtr)
}(arg0.StrVal()))}
})
	})
	return cache_toChar
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_takeWhile(func(inner_arg0 string) bool {
return (gopurs_runtime.Apply(p_0_box, gopurs_runtime.Str(inner_arg0)).IntVal) != (0)
}, s_1_box.StrVal()))
})
	})
	return cache_takeWhile
}

var cache_takeRight gopurs_runtime.Value
var once_takeRight sync.Once
func Get_takeRight() gopurs_runtime.Value {
	once_takeRight.Do(func() {
		cache_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_takeRight(i_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_takeRight
}

var cache_stripSuffix gopurs_runtime.Value
var once_stripSuffix sync.Once
func Get_stripSuffix() gopurs_runtime.Value {
	once_stripSuffix.Do(func() {
		cache_stripSuffix = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, str_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_stripSuffix(v_0_box.StrVal(), str_1_box.StrVal()))}
})
	})
	return cache_stripSuffix
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, str_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_stripPrefix(v_0_box.StrVal(), str_1_box.StrVal()))}
})
	})
	return cache_stripPrefix
}

var cache_startsWith gopurs_runtime.Value
var once_startsWith sync.Once
func Get_startsWith() gopurs_runtime.Value {
	once_startsWith.Do(func() {
		cache_startsWith = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_startsWith(pat_0_box.StrVal(), x_1_box.StrVal()))
})
	})
	return cache_startsWith
}

var cache_lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		cache_lastIndexOf_prime = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 string, inner_arg1 int64, inner_arg2 string) *pkg_Data_Maybe.Constructor_Just[int64] {
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply3(gopurs_runtime.Apply2(Get__lastIndexOfStartingAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), gopurs_runtime.Str(inner_arg0), gopurs_runtime.Int(inner_arg1), gopurs_runtime.Str(inner_arg2)).UnsafePtr)
}(arg0.StrVal(), arg1.IntVal, arg2.StrVal()))}
})
	})
	return cache_lastIndexOf_prime
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 string, inner_arg1 string) *pkg_Data_Maybe.Constructor_Just[int64] {
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.Apply2(Get__lastIndexOf(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), gopurs_runtime.Str(inner_arg0), gopurs_runtime.Str(inner_arg1)).UnsafePtr)
}(arg0.StrVal(), arg1.StrVal()))}
})
	})
	return cache_lastIndexOf
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 string, inner_arg1 int64, inner_arg2 string) *pkg_Data_Maybe.Constructor_Just[int64] {
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply3(gopurs_runtime.Apply2(Get__indexOfStartingAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), gopurs_runtime.Str(inner_arg0), gopurs_runtime.Int(inner_arg1), gopurs_runtime.Str(inner_arg2)).UnsafePtr)
}(arg0.StrVal(), arg1.IntVal, arg2.StrVal()))}
})
	})
	return cache_indexOf_prime
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 string, inner_arg1 string) *pkg_Data_Maybe.Constructor_Just[int64] {
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.Apply2(Get__indexOf(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), gopurs_runtime.Str(inner_arg0), gopurs_runtime.Str(inner_arg1)).UnsafePtr)
}(arg0.StrVal(), arg1.StrVal()))}
})
	})
	return cache_indexOf
}

var cache_endsWith gopurs_runtime.Value
var once_endsWith sync.Once
func Get_endsWith() gopurs_runtime.Value {
	once_endsWith.Do(func() {
		cache_endsWith = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_endsWith(pat_0_box.StrVal(), x_1_box.StrVal()))
})
	})
	return cache_endsWith
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_dropWhile(func(inner_arg0 string) bool {
return (gopurs_runtime.Apply(p_0_box, gopurs_runtime.Str(inner_arg0)).IntVal) != (0)
}, s_1_box.StrVal()))
})
	})
	return cache_dropWhile
}

var cache_dropRight gopurs_runtime.Value
var once_dropRight sync.Once
func Get_dropRight() gopurs_runtime.Value {
	once_dropRight.Do(func() {
		cache_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_dropRight(i_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_dropRight
}

var cache_contains gopurs_runtime.Value
var once_contains sync.Once
func Get_contains() gopurs_runtime.Value {
	once_contains.Do(func() {
		cache_contains = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_contains(pat_0_box.StrVal())
})
	})
	return cache_contains
}

var cache_charAt gopurs_runtime.Value
var once_charAt sync.Once
func Get_charAt() gopurs_runtime.Value {
	once_charAt.Do(func() {
		cache_charAt = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 int64, inner_arg1 string) *pkg_Data_Maybe.Constructor_Just[string] {
return (*pkg_Data_Maybe.Constructor_Just[string])(gopurs_runtime.Apply2(gopurs_runtime.Apply2(Get__charAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), gopurs_runtime.Int(inner_arg0), gopurs_runtime.Str(inner_arg1)).UnsafePtr)
}(arg0.IntVal, arg1.StrVal()))}
})
	})
	return cache_charAt
}

var cache__charAt gopurs_runtime.Value
var once__charAt sync.Once
func Get__charAt() gopurs_runtime.Value {
	once__charAt.Do(func() {
		cache__charAt = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_CharAt(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2.IntVal, arg3.StrVal()))}
})
	})
	return cache__charAt
}

var cache__indexOf gopurs_runtime.Value
var once__indexOf sync.Once
func Get__indexOf() gopurs_runtime.Value {
	once__indexOf.Do(func() {
		cache__indexOf = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_IndexOf(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2.StrVal(), arg3.StrVal()))}
})
	})
	return cache__indexOf
}

var cache__indexOfStartingAt gopurs_runtime.Value
var once__indexOfStartingAt sync.Once
func Get__indexOfStartingAt() gopurs_runtime.Value {
	once__indexOfStartingAt.Do(func() {
		cache__indexOfStartingAt = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_IndexOfStartingAt(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2.StrVal(), arg3.IntVal, arg4.StrVal()))}
})
	})
	return cache__indexOfStartingAt
}

var cache__lastIndexOf gopurs_runtime.Value
var once__lastIndexOf sync.Once
func Get__lastIndexOf() gopurs_runtime.Value {
	once__lastIndexOf.Do(func() {
		cache__lastIndexOf = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_LastIndexOf(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2.StrVal(), arg3.StrVal()))}
})
	})
	return cache__lastIndexOf
}

var cache__lastIndexOfStartingAt gopurs_runtime.Value
var once__lastIndexOfStartingAt sync.Once
func Get__lastIndexOfStartingAt() gopurs_runtime.Value {
	once__lastIndexOfStartingAt.Do(func() {
		cache__lastIndexOfStartingAt = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_LastIndexOfStartingAt(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2.StrVal(), arg3.IntVal, arg4.StrVal()))}
})
	})
	return cache__lastIndexOfStartingAt
}

var cache__toChar gopurs_runtime.Value
var once__toChar sync.Once
func Get__toChar() gopurs_runtime.Value {
	once__toChar.Do(func() {
		cache__toChar = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_ToChar(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2.StrVal()))}
})
	})
	return cache__toChar
}

var cache_countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		cache_countPrefix = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(CountPrefix(func(inner_arg0 string) bool {
return (gopurs_runtime.Apply(arg0, gopurs_runtime.Str(inner_arg0)).IntVal) != (0)
}, arg1.StrVal()))
})
	})
	return cache_countPrefix
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Drop(arg0.IntVal, arg1.StrVal()))
})
	})
	return cache_drop
}

var cache_fromCharArray gopurs_runtime.Value
var once_fromCharArray sync.Once
func Get_fromCharArray() gopurs_runtime.Value {
	once_fromCharArray.Do(func() {
		cache_fromCharArray = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(FromCharArray(func() []string {
					arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))
})
	})
	return cache_fromCharArray
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Length(arg0.StrVal()))
})
	})
	return cache_length
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Singleton(arg0.StrVal()))
})
	})
	return cache_singleton
}

var cache_slice gopurs_runtime.Value
var once_slice sync.Once
func Get_slice() gopurs_runtime.Value {
	once_slice.Do(func() {
		cache_slice = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Slice(arg0.IntVal, arg1.IntVal, arg2.StrVal()))
})
	})
	return cache_slice
}

var cache_splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		cache_splitAt = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return SplitAt(arg0.IntVal, arg1.StrVal())
})
	})
	return cache_splitAt
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Take(arg0.IntVal, arg1.StrVal()))
})
	})
	return cache_take
}

var cache_toCharArray gopurs_runtime.Value
var once_toCharArray sync.Once
func Get_toCharArray() gopurs_runtime.Value {
	once_toCharArray.Do(func() {
		cache_toCharArray = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := ToCharArray(arg0.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toCharArray
}

func Call_uncons(v_0_loop string) *pkg_Data_Maybe.Constructor_Just[interface{}] {
var v_0 string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == ("") {
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(v_0)), gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Int(1), gopurs_runtime.Str(v_0))))})})
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(__t0.UnsafePtr)
}

func Call_takeWhile(p_0_loop func(string) bool, s_1_loop string) string {
var p_0 func(string) bool = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_countPrefix(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(p_0(arg0.StrVal()))
}), gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_takeRight(i_0_loop int64, s_1_loop string) string {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(s_1)).IntVal) - (i_0)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_stripSuffix(v_0_loop string, str_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 string = v_0_loop
_ = v_0
var str_1 string = str_1_loop
_ = str_1
v1_2_0 := gopurs_runtime.Apply2(Get_splitAt(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(str_1)).IntVal) - (gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(v_0)).IntVal)), gopurs_runtime.Str(str_1))
_ = v1_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"), gopurs_runtime.RecordGet(v1_2_0, "after"), gopurs_runtime.Str(v_0)).IntVal) != (0) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(v1_2_0, "before"))})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t1.UnsafePtr)
}

func Call_stripPrefix(v_0_loop string, str_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 string = v_0_loop
_ = v_0
var str_1 string = str_1_loop
_ = str_1
v1_2_0 := gopurs_runtime.Apply2(Get_splitAt(), gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(v_0)), gopurs_runtime.Str(str_1))
_ = v1_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"), gopurs_runtime.RecordGet(v1_2_0, "before"), gopurs_runtime.Str(v_0)).IntVal) != (0) {
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(v1_2_0, "after"))})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[string])(__t1.UnsafePtr)
}

func Call_startsWith(pat_0_loop string, x_1_loop string) bool {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 string = x_1_loop
_ = x_1
__local_var_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_stripPrefix(pat_0, x_1))}
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_endsWith(pat_0_loop string, x_1_loop string) bool {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 string = x_1_loop
_ = x_1
__local_var_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_stripSuffix(pat_0, x_1))}
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_dropWhile(p_0_loop func(string) bool, s_1_loop string) string {
var p_0 func(string) bool = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Apply2(Get_countPrefix(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(p_0(arg0.StrVal()))
}), gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_dropRight(i_0_loop int64, s_1_loop string) string {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(s_1)).IntVal) - (i_0)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_contains(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
__local_var_1_0 := gopurs_runtime.Apply(Get_indexOf(), gopurs_runtime.Str(pat_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136 && __local_var_3_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136 && __local_var_3_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
}

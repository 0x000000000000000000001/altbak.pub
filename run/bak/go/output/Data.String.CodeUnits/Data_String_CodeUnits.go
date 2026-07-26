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
return Call_uncons(v_0_box.StrVal())
})
	})
	return cache_uncons
}

var cache_toChar gopurs_runtime.Value
var once_toChar sync.Once
func Get_toChar() gopurs_runtime.Value {
	once_toChar.Do(func() {
		cache_toChar = gopurs_runtime.Apply2(Get__toChar(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil})
	})
	return cache_toChar
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(p_0_box, s_1_box.StrVal())
})
	})
	return cache_takeWhile
}

var cache_takeRight gopurs_runtime.Value
var once_takeRight sync.Once
func Get_takeRight() gopurs_runtime.Value {
	once_takeRight.Do(func() {
		cache_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeRight(i_0_box.IntVal, s_1_box.StrVal())
})
	})
	return cache_takeRight
}

var cache_stripSuffix gopurs_runtime.Value
var once_stripSuffix sync.Once
func Get_stripSuffix() gopurs_runtime.Value {
	once_stripSuffix.Do(func() {
		cache_stripSuffix = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, str_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripSuffix(v_0_box, str_1_box.StrVal())
})
	})
	return cache_stripSuffix
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, str_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripPrefix(v_0_box, str_1_box.StrVal())
})
	})
	return cache_stripPrefix
}

var cache_startsWith gopurs_runtime.Value
var once_startsWith sync.Once
func Get_startsWith() gopurs_runtime.Value {
	once_startsWith.Do(func() {
		cache_startsWith = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_startsWith(pat_0_box, x_1_box.StrVal())
})
	})
	return cache_startsWith
}

var cache_lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		cache_lastIndexOf_prime = gopurs_runtime.Apply2(Get__lastIndexOfStartingAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil})
	})
	return cache_lastIndexOf_prime
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Apply2(Get__lastIndexOf(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil})
	})
	return cache_lastIndexOf
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Apply2(Get__indexOfStartingAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil})
	})
	return cache_indexOf_prime
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Apply2(Get__indexOf(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil})
	})
	return cache_indexOf
}

var cache_endsWith gopurs_runtime.Value
var once_endsWith sync.Once
func Get_endsWith() gopurs_runtime.Value {
	once_endsWith.Do(func() {
		cache_endsWith = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_endsWith(pat_0_box, x_1_box.StrVal())
})
	})
	return cache_endsWith
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(p_0_box, s_1_box.StrVal())
})
	})
	return cache_dropWhile
}

var cache_dropRight gopurs_runtime.Value
var once_dropRight sync.Once
func Get_dropRight() gopurs_runtime.Value {
	once_dropRight.Do(func() {
		cache_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropRight(i_0_box.IntVal, s_1_box.StrVal())
})
	})
	return cache_dropRight
}

var cache_contains gopurs_runtime.Value
var once_contains sync.Once
func Get_contains() gopurs_runtime.Value {
	once_contains.Do(func() {
		cache_contains = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_contains(pat_0_box)
})
	})
	return cache_contains
}

var cache_charAt gopurs_runtime.Value
var once_charAt sync.Once
func Get_charAt() gopurs_runtime.Value {
	once_charAt.Do(func() {
		cache_charAt = gopurs_runtime.Apply2(Get__charAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil})
	})
	return cache_charAt
}

func Call_uncons(v_0_loop string) gopurs_runtime.Value {
var v_0 string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == ("") {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just{gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(v_0)), gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Int(1), gopurs_runtime.Str(v_0)))})}
}
end_branch_0:
return __t0
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1))
}

func Call_takeRight(i_0_loop int64, s_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(s_1)).IntVal) - (i_0)), gopurs_runtime.Str(s_1))
}

func Call_stripSuffix(v_0_loop gopurs_runtime.Value, str_1_loop string) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var str_1 string = str_1_loop
_ = str_1
v1_2_0 := gopurs_runtime.Apply2(Get_splitAt(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(str_1)).IntVal) - (gopurs_runtime.Apply(Get_length(), v_0).IntVal)), gopurs_runtime.Str(str_1))
_ = v1_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"), gopurs_runtime.RecordGet(v1_2_0, "after"), v_0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just{gopurs_runtime.RecordGet(v1_2_0, "before")})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_1:
return __t1
}

func Call_stripPrefix(v_0_loop gopurs_runtime.Value, str_1_loop string) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var str_1 string = str_1_loop
_ = str_1
v1_2_0 := gopurs_runtime.Apply2(Get_splitAt(), gopurs_runtime.Apply(Get_length(), v_0), gopurs_runtime.Str(str_1))
_ = v1_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"), gopurs_runtime.RecordGet(v1_2_0, "before"), v_0).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just{gopurs_runtime.RecordGet(v1_2_0, "after")})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
}
end_branch_1:
return __t1
}

func Call_startsWith(pat_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var x_1 string = x_1_loop
_ = x_1
__local_var_2_0 := Call_stripPrefix(pat_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3589588149) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_endsWith(pat_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var x_1 string = x_1_loop
_ = x_1
__local_var_2_0 := Call_stripSuffix(pat_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3589588149) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_dropWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1))
}

func Call_dropRight(i_0_loop int64, s_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(s_1)).IntVal) - (i_0)), gopurs_runtime.Str(s_1))
}

func Call_contains(pat_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
__local_var_1_0 := gopurs_runtime.Apply(Get_indexOf(), pat_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 3589588149) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1.Type == 9 && __local_var_3_1.IntVal == 930809136) {
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

func Get__charAt() gopurs_runtime.Value {
	return _Gopurs__CharAt
}

func Get__indexOf() gopurs_runtime.Value {
	return _Gopurs__IndexOf
}

func Get__indexOfStartingAt() gopurs_runtime.Value {
	return _Gopurs__IndexOfStartingAt
}

func Get__lastIndexOf() gopurs_runtime.Value {
	return _Gopurs__LastIndexOf
}

func Get__lastIndexOfStartingAt() gopurs_runtime.Value {
	return _Gopurs__LastIndexOfStartingAt
}

func Get__toChar() gopurs_runtime.Value {
	return _Gopurs__ToChar
}

func Get_countPrefix() gopurs_runtime.Value {
	return _Gopurs_CountPrefix
}

func Get_drop() gopurs_runtime.Value {
	return _Gopurs_Drop
}

func Get_fromCharArray() gopurs_runtime.Value {
	return _Gopurs_FromCharArray
}

func Get_length() gopurs_runtime.Value {
	return _Gopurs_Length
}

func Get_singleton() gopurs_runtime.Value {
	return _Gopurs_Singleton
}

func Get_slice() gopurs_runtime.Value {
	return _Gopurs_Slice
}

func Get_splitAt() gopurs_runtime.Value {
	return _Gopurs_SplitAt
}

func Get_take() gopurs_runtime.Value {
	return _Gopurs_Take
}

func Get_toCharArray() gopurs_runtime.Value {
	return _Gopurs_ToCharArray
}

package Data_String_CodeUnits

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_toChar = gopurs_runtime.Apply2(Get__toChar(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_toChar
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_takeWhile(p_0_box, s_1_box.StrVal()))
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
		cache_lastIndexOf_prime = gopurs_runtime.Apply2(Get__lastIndexOfStartingAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_lastIndexOf_prime
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Apply2(Get__lastIndexOf(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_lastIndexOf
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Apply2(Get__indexOfStartingAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_indexOf_prime
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Apply2(Get__indexOf(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
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
return gopurs_runtime.Str(Call_dropWhile(p_0_box, s_1_box.StrVal()))
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
		cache_charAt = gopurs_runtime.Apply2(Get__charAt(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_charAt
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_eq__472317769 gopurs_runtime.Value
var once_eq__472317769 sync.Once
func Get_eq__472317769() gopurs_runtime.Value {
	once_eq__472317769.Do(func() {
		cache_eq__472317769 = gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq")
	})
	return cache_eq__472317769
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_const__1243414737 gopurs_runtime.Value
var once_const__1243414737 sync.Once
func Get_const__1243414737() gopurs_runtime.Value {
	once_const__1243414737.Do(func() {
		cache_const__1243414737 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1243414737(a_0_box, v_1_box)
})
	})
	return cache_const__1243414737
}

var cache_const__2082174484 gopurs_runtime.Value
var once_const__2082174484 sync.Once
func Get_const__2082174484() gopurs_runtime.Value {
	once_const__2082174484.Do(func() {
		cache_const__2082174484 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2082174484(a_0_box, v_1_box)
})
	})
	return cache_const__2082174484
}

var cache_const__4157258135 gopurs_runtime.Value
var once_const__4157258135 sync.Once
func Get_const__4157258135() gopurs_runtime.Value {
	once_const__4157258135.Do(func() {
		cache_const__4157258135 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4157258135(a_0_box, v_1_box)
})
	})
	return cache_const__4157258135
}

var cache_const__1562253172 gopurs_runtime.Value
var once_const__1562253172 sync.Once
func Get_const__1562253172() gopurs_runtime.Value {
	once_const__1562253172.Do(func() {
		cache_const__1562253172 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1562253172(a_0_box, v_1_box)
})
	})
	return cache_const__1562253172
}

var cache_isJust__2514352589 gopurs_runtime.Value
var once_isJust__2514352589 sync.Once
func Get_isJust__2514352589() gopurs_runtime.Value {
	once_isJust__2514352589.Do(func() {
		cache_isJust__2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__2514352589(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v2_0_box)))
})
	})
	return cache_isJust__2514352589
}

var cache_isJust__2475527019 gopurs_runtime.Value
var once_isJust__2475527019 sync.Once
func Get_isJust__2475527019() gopurs_runtime.Value {
	once_isJust__2475527019.Do(func() {
		cache_isJust__2475527019 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__2475527019(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](v2_0_box)))
})
	})
	return cache_isJust__2475527019
}

var cache_maybe__3078346790 gopurs_runtime.Value
var once_maybe__3078346790 sync.Once
func Get_maybe__3078346790() gopurs_runtime.Value {
	once_maybe__3078346790.Do(func() {
		cache_maybe__3078346790 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3078346790(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3078346790
}

var cache_maybe__1510464358 gopurs_runtime.Value
var once_maybe__1510464358 sync.Once
func Get_maybe__1510464358() gopurs_runtime.Value {
	once_maybe__1510464358.Do(func() {
		cache_maybe__1510464358 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1510464358(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1510464358
}

var cache_maybe__3718989812 gopurs_runtime.Value
var once_maybe__3718989812 sync.Once
func Get_maybe__3718989812() gopurs_runtime.Value {
	once_maybe__3718989812.Do(func() {
		cache_maybe__3718989812 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3718989812(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3718989812
}

var cache_maybe__1647364852 gopurs_runtime.Value
var once_maybe__1647364852 sync.Once
func Get_maybe__1647364852() gopurs_runtime.Value {
	once_maybe__1647364852.Do(func() {
		cache_maybe__1647364852 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1647364852(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1647364852
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

func Call_uncons(v_0_loop string) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == ("") {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(v_0)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Int(1), gopurs_runtime.Str(v_0)).StrVal()))})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_takeRight(i_0_loop int64, s_1_loop string) string {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(s_1)), gopurs_runtime.Int(i_0)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_stripSuffix(v_0_loop string, str_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 string = v_0_loop
_ = v_0
var str_1 string = str_1_loop
_ = str_1
v1_2_0 := gopurs_runtime.Apply2(Get_splitAt(), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(str_1)), gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(v_0))), gopurs_runtime.Str(str_1))
_ = v1_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__472317769(), gopurs_runtime.RecordGet(v1_2_0, "after"), gopurs_runtime.Str(v_0)).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(v1_2_0, "before")})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1)
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
if (gopurs_runtime.Apply2(Get_eq__472317769(), gopurs_runtime.RecordGet(v1_2_0, "before"), gopurs_runtime.Str(v_0)).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(v1_2_0, "after")})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1)
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

func Call_dropWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_drop(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_dropRight(i_0_loop int64, s_1_loop string) string {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply(Get_length(), gopurs_runtime.Str(s_1)), gopurs_runtime.Int(i_0)), gopurs_runtime.Str(s_1)).StrVal()
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

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_const__1243414737(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2082174484(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__4157258135(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__1562253172(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_isJust__2514352589(v2_0_loop *pkg_Data_Maybe.Constructor_Just[int64]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[int64] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isJust__2475527019(v2_0_loop *pkg_Data_Maybe.Constructor_Just[string]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[string] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__3078346790(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__1510464358(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3718989812(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__1647364852(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

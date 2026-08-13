package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Char_Gen_foldable1NonEmpty gopurs_runtime.Value
var once_Data_Char_Gen_foldable1NonEmpty sync.Once
func Get_Data_Char_Gen_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_Char_Gen_foldable1NonEmpty.Do(func() {
		cache_Data_Char_Gen_foldable1NonEmpty = func() gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_0_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_0_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_1.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), gopurs_runtime.Apply(f_2, x_5), acc_6)
})
}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray5 := (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1
_ = arr_val_foldlArray5
res_go_foldlArray5 := gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)
_ = res_go_foldlArray5
arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
_ = arr_go_foldlArray5
for _, v_foldlArray5 := range *arr_go_foldlArray5 {
res_go_foldlArray5 = gopurs_runtime.Apply2(f_0, res_go_foldlArray5, v_foldlArray5)
}
return res_go_foldlArray5
}()
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray5 := (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1
_ = arr_val_foldlArray5
res_go_foldlArray5 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)
_ = res_go_foldlArray5
arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
_ = arr_go_foldlArray5
for _, v_foldlArray5 := range *arr_go_foldlArray5 {
res_go_foldlArray5 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), s_4, gopurs_runtime.Apply(f_2, a1_5))
})
}), res_go_foldlArray5, v_foldlArray5)
}
return res_go_foldlArray5
}()
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray4 := (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1
_ = arr_val_foldlArray4
res_go_foldlArray4 := (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
_ = res_go_foldlArray4
arr_go_foldlArray4 := (*[]gopurs_runtime.Value)(arr_val_foldlArray4.UnsafePtr)
_ = arr_go_foldlArray4
for _, v_foldlArray4 := range *arr_go_foldlArray4 {
res_go_foldlArray4 = gopurs_runtime.Apply2(f_1, res_go_foldlArray4, v_foldlArray4)
}
return res_go_foldlArray4
}()
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)
_ = __local_var_3_3
// TAST (Let): __local_var_4_4 -> *Constructor_Data_Maybe_Just
__local_var_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(f_1, a1_4)
_ = __local_var_5_6
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 930809136 && v2_6.UnsafePtr == nil) {
__t7 = a1_4
goto end_branch_7
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 930809136 && v2_6.UnsafePtr != nil) {
__t7 = gopurs_runtime.Apply(__local_var_5_6, (*Constructor_Data_Maybe_Just)(v2_6.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
_ = __local_var_5_5
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_5_5, x_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
_ = __local_var_4_4
var __t8 gopurs_runtime.Value
{
if (__local_var_4_4 == nil) {
__t8 = (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
goto end_branch_8
} else {

}
}
{
if (__local_var_4_4 != nil) {
__t8 = gopurs_runtime.Apply(__local_var_3_3, (__local_var_4_4).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})})}
}()
	})
	return cache_Data_Char_Gen_foldable1NonEmpty
}

var cache_Data_Char_Gen_genUnicodeChar gopurs_runtime.Value
var once_Data_Char_Gen_genUnicodeChar sync.Once
func Get_Data_Char_Gen_genUnicodeChar() gopurs_runtime.Value {
	once_Data_Char_Gen_genUnicodeChar.Do(func() {
		cache_Data_Char_Gen_genUnicodeChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genUnicodeChar(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genUnicodeChar
}

var cache_Data_Char_Gen_genDigitChar gopurs_runtime.Value
var once_Data_Char_Gen_genDigitChar sync.Once
func Get_Data_Char_Gen_genDigitChar() gopurs_runtime.Value {
	once_Data_Char_Gen_genDigitChar.Do(func() {
		cache_Data_Char_Gen_genDigitChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genDigitChar(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genDigitChar
}

var cache_Data_Char_Gen_genAsciiChar_prime gopurs_runtime.Value
var once_Data_Char_Gen_genAsciiChar_prime sync.Once
func Get_Data_Char_Gen_genAsciiChar_prime() gopurs_runtime.Value {
	once_Data_Char_Gen_genAsciiChar_prime.Do(func() {
		cache_Data_Char_Gen_genAsciiChar_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genAsciiChar_prime(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genAsciiChar_prime
}

var cache_Data_Char_Gen_genAsciiChar gopurs_runtime.Value
var once_Data_Char_Gen_genAsciiChar sync.Once
func Get_Data_Char_Gen_genAsciiChar() gopurs_runtime.Value {
	once_Data_Char_Gen_genAsciiChar.Do(func() {
		cache_Data_Char_Gen_genAsciiChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genAsciiChar(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genAsciiChar
}

var cache_Data_Char_Gen_genAlphaUppercase gopurs_runtime.Value
var once_Data_Char_Gen_genAlphaUppercase sync.Once
func Get_Data_Char_Gen_genAlphaUppercase() gopurs_runtime.Value {
	once_Data_Char_Gen_genAlphaUppercase.Do(func() {
		cache_Data_Char_Gen_genAlphaUppercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genAlphaUppercase
}

var cache_Data_Char_Gen_genAlphaLowercase gopurs_runtime.Value
var once_Data_Char_Gen_genAlphaLowercase sync.Once
func Get_Data_Char_Gen_genAlphaLowercase() gopurs_runtime.Value {
	once_Data_Char_Gen_genAlphaLowercase.Do(func() {
		cache_Data_Char_Gen_genAlphaLowercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genAlphaLowercase(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genAlphaLowercase
}

var cache_Data_Char_Gen_genAlpha gopurs_runtime.Value
var once_Data_Char_Gen_genAlpha sync.Once
func Get_Data_Char_Gen_genAlpha() gopurs_runtime.Value {
	once_Data_Char_Gen_genAlpha.Do(func() {
		cache_Data_Char_Gen_genAlpha = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genAlpha(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genAlpha
}

var cache_Data_Char_Gen_genAlphaLowercase__4294897069 gopurs_runtime.Value
var once_Data_Char_Gen_genAlphaLowercase__4294897069 sync.Once
func Get_Data_Char_Gen_genAlphaLowercase__4294897069() gopurs_runtime.Value {
	once_Data_Char_Gen_genAlphaLowercase__4294897069.Do(func() {
		cache_Data_Char_Gen_genAlphaLowercase__4294897069 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genAlphaLowercase__4294897069(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genAlphaLowercase__4294897069
}

func Call_Data_Char_Gen_genUnicodeChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(65536)))
}

func Call_Data_Char_Gen_genDigitChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(48), gopurs_runtime.Int(57)))
}

func Call_Data_Char_Gen_genAsciiChar_prime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(127)))
}

func Call_Data_Char_Gen_genAsciiChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(32), gopurs_runtime.Int(127)))
}

func Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))
}

func Call_Data_Char_Gen_genAlphaLowercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}

func Call_Data_Char_Gen_genAlpha(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): foldableNonEmpty1_1_1 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_1_1 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_2_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_2
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_5_3 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_5_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_2.V0), gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_5_3.V0), gopurs_runtime.Apply(f_3, x_6), acc_7)
})
}), gopurs_runtime.RecordGet(dictMonoid_1, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray6 := (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1
_ = arr_val_foldlArray6
res_go_foldlArray6 := gopurs_runtime.Apply2(f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)
_ = res_go_foldlArray6
arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
_ = arr_go_foldlArray6
for _, v_foldlArray6 := range *arr_go_foldlArray6 {
res_go_foldlArray6 = gopurs_runtime.Apply2(f_1, res_go_foldlArray6, v_foldlArray6)
}
return res_go_foldlArray6
}()
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_1, b_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_1_1
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_1_0 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_1_1)}
}), gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray6 := (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V1
_ = arr_val_foldlArray6
res_go_foldlArray6 := gopurs_runtime.Apply(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_4.UnsafePtr).V0)
_ = res_go_foldlArray6
arr_go_foldlArray6 := (*[]gopurs_runtime.Value)(arr_val_foldlArray6.UnsafePtr)
_ = arr_go_foldlArray6
for _, v_foldlArray6 := range *arr_go_foldlArray6 {
res_go_foldlArray6 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), s_5, gopurs_runtime.Apply(f_3, a1_6))
})
}), res_go_foldlArray6, v_foldlArray6)
}
return res_go_foldlArray6
}()
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray5 := (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1
_ = arr_val_foldlArray5
res_go_foldlArray5 := (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0
_ = res_go_foldlArray5
arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
_ = arr_go_foldlArray5
for _, v_foldlArray5 := range *arr_go_foldlArray5 {
res_go_foldlArray5 = gopurs_runtime.Apply2(f_2, res_go_foldlArray5, v_foldlArray5)
}
return res_go_foldlArray5
}()
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)
_ = __local_var_4_4
// TAST (Let): __local_var_5_5 -> *Constructor_Data_Maybe_Just
__local_var_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(f_2, a1_5)
_ = __local_var_6_7
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr == nil) {
__t8 = a1_5
goto end_branch_8
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr != nil) {
__t8 = gopurs_runtime.Apply(__local_var_6_7, (*Constructor_Data_Maybe_Just)(v2_7.UnsafePtr).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
_ = __local_var_6_6
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_6_6, x_7)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
_ = __local_var_5_5
var __t9 gopurs_runtime.Value
{
if (__local_var_5_5 == nil) {
__t9 = (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0
goto end_branch_9
} else {

}
}
{
if (__local_var_5_5 != nil) {
__t9 = gopurs_runtime.Apply(__local_var_4_4, (__local_var_5_5).V0)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
})
})}
_ = __local_var_1_0
// TAST (Let): __local_var_2_10 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_2_10 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_0.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_3.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_2_10)}).IntVal) - (1))), gopurs_runtime.Func(func(n_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_11_0 gopurs_runtime.Value
go__go_4_11_0 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop int64 = v_5_loop_val.IntVal
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_11_0:
for {
if false { continue go__go_4_11_0 }
var v_5 int64 = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t15 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr != nil) {
var __t14 gopurs_runtime.Value
{
var __t_tag_12 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1
if (__t_tag_12 == nil) {
__t14 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_14
} else {

}
}
{
var __t13 bool
{
if (v_5) > (0) {
__t13 = false
goto end_branch_13
} else {

}
}
{
__t13 = true
}
end_branch_13:
if __t13 {
__t14 = (*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V0
goto end_branch_14
} else {

}
}
{
v_5_loop = (v_5) - (1)
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_11_0
__t14 = gopurs_runtime.Value{}
}
end_branch_14:
__t15 = __t14
goto end_branch_15
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 759514854 && v1_6.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_2_10)})
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_11_0, gopurs_runtime.Int(n_3.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_0.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_2_10)})))})
}))
}

func Call_Data_Char_Gen_genAlphaLowercase__4294897069(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}



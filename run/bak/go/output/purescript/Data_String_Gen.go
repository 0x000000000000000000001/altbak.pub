package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Gen_genString gopurs_runtime.Value
var once_Data_String_Gen_genString sync.Once
func Get_Data_String_Gen_genString() gopurs_runtime.Value {
	once_Data_String_Gen_genString.Do(func() {
		cache_Data_String_Gen_genString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Data_String_Gen_genString
}

var cache_Data_String_Gen_genUnicodeString gopurs_runtime.Value
var once_Data_String_Gen_genUnicodeString sync.Once
func Get_Data_String_Gen_genUnicodeString() gopurs_runtime.Value {
	once_Data_String_Gen_genUnicodeString.Do(func() {
		cache_Data_String_Gen_genUnicodeString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genUnicodeString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_Data_String_Gen_genUnicodeString
}

var cache_Data_String_Gen_genDigitString gopurs_runtime.Value
var once_Data_String_Gen_genDigitString sync.Once
func Get_Data_String_Gen_genDigitString() gopurs_runtime.Value {
	once_Data_String_Gen_genDigitString.Do(func() {
		cache_Data_String_Gen_genDigitString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genDigitString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_Data_String_Gen_genDigitString
}

var cache_Data_String_Gen_genAsciiString_prime gopurs_runtime.Value
var once_Data_String_Gen_genAsciiString_prime sync.Once
func Get_Data_String_Gen_genAsciiString_prime() gopurs_runtime.Value {
	once_Data_String_Gen_genAsciiString_prime.Do(func() {
		cache_Data_String_Gen_genAsciiString_prime = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genAsciiString_prime(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_Data_String_Gen_genAsciiString_prime
}

var cache_Data_String_Gen_genAsciiString gopurs_runtime.Value
var once_Data_String_Gen_genAsciiString sync.Once
func Get_Data_String_Gen_genAsciiString() gopurs_runtime.Value {
	once_Data_String_Gen_genAsciiString.Do(func() {
		cache_Data_String_Gen_genAsciiString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genAsciiString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_Data_String_Gen_genAsciiString
}

var cache_Data_String_Gen_genAlphaUppercaseString gopurs_runtime.Value
var once_Data_String_Gen_genAlphaUppercaseString sync.Once
func Get_Data_String_Gen_genAlphaUppercaseString() gopurs_runtime.Value {
	once_Data_String_Gen_genAlphaUppercaseString.Do(func() {
		cache_Data_String_Gen_genAlphaUppercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genAlphaUppercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_Data_String_Gen_genAlphaUppercaseString
}

var cache_Data_String_Gen_genAlphaString gopurs_runtime.Value
var once_Data_String_Gen_genAlphaString sync.Once
func Get_Data_String_Gen_genAlphaString() gopurs_runtime.Value {
	once_Data_String_Gen_genAlphaString.Do(func() {
		cache_Data_String_Gen_genAlphaString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genAlphaString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_Data_String_Gen_genAlphaString
}

var cache_Data_String_Gen_genAlphaLowercaseString gopurs_runtime.Value
var once_Data_String_Gen_genAlphaLowercaseString sync.Once
func Get_Data_String_Gen_genAlphaLowercaseString() gopurs_runtime.Value {
	once_Data_String_Gen_genAlphaLowercaseString.Do(func() {
		cache_Data_String_Gen_genAlphaLowercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genAlphaLowercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_Data_String_Gen_genAlphaLowercaseString
}

var cache_Data_String_Gen_genString__3531165174 gopurs_runtime.Value
var once_Data_String_Gen_genString__3531165174 sync.Once
func Get_Data_String_Gen_genString__3531165174() gopurs_runtime.Value {
	once_Data_String_Gen_genString__3531165174.Do(func() {
		cache_Data_String_Gen_genString__3531165174 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Gen_genString__3531165174(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Data_String_Gen_genString__3531165174
}

func Call_Data_String_Gen_genString(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V3), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_3 -> gopurs_runtime.Value
Monad0_8_3 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_8_3
// TAST (Let): pure_9_4 -> gopurs_runtime.Value
pure_9_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_4
// TAST (Let): Bind1_10_5 -> *Constructor_Control_Bind_Bind
Bind1_10_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_5
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
var __t11 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t11 = false
goto end_branch_11
} else {

}
}
{
__t11 = true
}
end_branch_11:
if __t11 {
__t12 = gopurs_runtime.Apply(pure_9_4, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_12
} else {

}
}
{
// TAST (Let): __local_var_12_9 -> gopurs_runtime.Value
__local_var_12_9 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_9
// TAST (Let): __local_var_13_10 -> gopurs_runtime.Value
__local_var_13_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_10
__t12 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_5.V1), genChar_5, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_4, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_9)})}, gopurs_runtime.Int((__local_var_13_10.IntVal) - (1))})}})})
}))
}
end_branch_12:
return __t12
}))
_ = __local_var_11_8
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_13
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t6 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t7 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_7
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t7 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_8, gopurs_runtime.Apply(__local_var_12_13, x_13))
})))))
}))
}))
})
}

func Call_Data_String_Gen_genUnicodeString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(65536)))
_ = __local_var_5_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_4 -> gopurs_runtime.Value
Monad0_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_8_4
// TAST (Let): pure_9_5 -> gopurs_runtime.Value
pure_9_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_5
// TAST (Let): Bind1_10_6 -> *Constructor_Control_Bind_Bind
Bind1_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_6
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t12 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t12 = false
goto end_branch_12
} else {

}
}
{
__t12 = true
}
end_branch_12:
if __t12 {
__t13 = gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_13
} else {

}
}
{
// TAST (Let): __local_var_12_10 -> gopurs_runtime.Value
__local_var_12_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_10
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_11
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), __local_var_5_3, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_10)})}, gopurs_runtime.Int((__local_var_13_11.IntVal) - (1))})}})})
}))
}
end_branch_13:
return __t13
}))
_ = __local_var_11_9
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t7 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_9, gopurs_runtime.Apply(__local_var_12_14, x_13))
})))))
}))
}))
}

func Call_Data_String_Gen_genDigitString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(48), gopurs_runtime.Int(57)))
_ = __local_var_5_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_4 -> gopurs_runtime.Value
Monad0_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_8_4
// TAST (Let): pure_9_5 -> gopurs_runtime.Value
pure_9_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_5
// TAST (Let): Bind1_10_6 -> *Constructor_Control_Bind_Bind
Bind1_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_6
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t12 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t12 = false
goto end_branch_12
} else {

}
}
{
__t12 = true
}
end_branch_12:
if __t12 {
__t13 = gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_13
} else {

}
}
{
// TAST (Let): __local_var_12_10 -> gopurs_runtime.Value
__local_var_12_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_10
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_11
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), __local_var_5_3, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_10)})}, gopurs_runtime.Int((__local_var_13_11.IntVal) - (1))})}})})
}))
}
end_branch_13:
return __t13
}))
_ = __local_var_11_9
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t7 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_9, gopurs_runtime.Apply(__local_var_12_14, x_13))
})))))
}))
}))
}

func Call_Data_String_Gen_genAsciiString_prime(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(127)))
_ = __local_var_5_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_4 -> gopurs_runtime.Value
Monad0_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_8_4
// TAST (Let): pure_9_5 -> gopurs_runtime.Value
pure_9_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_5
// TAST (Let): Bind1_10_6 -> *Constructor_Control_Bind_Bind
Bind1_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_6
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t12 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t12 = false
goto end_branch_12
} else {

}
}
{
__t12 = true
}
end_branch_12:
if __t12 {
__t13 = gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_13
} else {

}
}
{
// TAST (Let): __local_var_12_10 -> gopurs_runtime.Value
__local_var_12_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_10
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_11
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), __local_var_5_3, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_10)})}, gopurs_runtime.Int((__local_var_13_11.IntVal) - (1))})}})})
}))
}
end_branch_13:
return __t13
}))
_ = __local_var_11_9
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t7 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_9, gopurs_runtime.Apply(__local_var_12_14, x_13))
})))))
}))
}))
}

func Call_Data_String_Gen_genAsciiString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(32), gopurs_runtime.Int(127)))
_ = __local_var_5_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_4 -> gopurs_runtime.Value
Monad0_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_8_4
// TAST (Let): pure_9_5 -> gopurs_runtime.Value
pure_9_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_5
// TAST (Let): Bind1_10_6 -> *Constructor_Control_Bind_Bind
Bind1_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_6
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t12 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t12 = false
goto end_branch_12
} else {

}
}
{
__t12 = true
}
end_branch_12:
if __t12 {
__t13 = gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_13
} else {

}
}
{
// TAST (Let): __local_var_12_10 -> gopurs_runtime.Value
__local_var_12_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_10
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_11
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), __local_var_5_3, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_10)})}, gopurs_runtime.Int((__local_var_13_11.IntVal) - (1))})}})})
}))
}
end_branch_13:
return __t13
}))
_ = __local_var_11_9
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t7 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_9, gopurs_runtime.Apply(__local_var_12_14, x_13))
})))))
}))
}))
}

func Call_Data_String_Gen_genAlphaUppercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))
_ = __local_var_5_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_4 -> gopurs_runtime.Value
Monad0_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_8_4
// TAST (Let): pure_9_5 -> gopurs_runtime.Value
pure_9_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_5
// TAST (Let): Bind1_10_6 -> *Constructor_Control_Bind_Bind
Bind1_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_6
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t12 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t12 = false
goto end_branch_12
} else {

}
}
{
__t12 = true
}
end_branch_12:
if __t12 {
__t13 = gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_13
} else {

}
}
{
// TAST (Let): __local_var_12_10 -> gopurs_runtime.Value
__local_var_12_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_10
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_11
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), __local_var_5_3, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_10)})}, gopurs_runtime.Int((__local_var_13_11.IntVal) - (1))})}})})
}))
}
end_branch_13:
return __t13
}))
_ = __local_var_11_9
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t7 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_9, gopurs_runtime.Apply(__local_var_12_14, x_13))
})))))
}))
}))
}

func Call_Data_String_Gen_genAlphaString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
// TAST (Let): foldableNonEmpty1_5_5 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_5_5 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_6 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_6
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_7
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_6.V0), gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_7.V0), gopurs_runtime.Apply(f_7, x_10), acc_11)
})
}), gopurs_runtime.RecordGet(dictMonoid_5, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray10 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray10
res_go_foldlArray10 := gopurs_runtime.Apply2(f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = res_go_foldlArray10
arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
_ = arr_go_foldlArray10
for _, v_foldlArray10 := range *arr_go_foldlArray10 {
res_go_foldlArray10 = gopurs_runtime.Apply2(f_5, res_go_foldlArray10, v_foldlArray10)
}
return res_go_foldlArray10
}()
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_5_5
// TAST (Let): __local_var_5_4 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_5_4 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_5_5)}
}), gopurs_runtime.Func(func(dictSemigroup_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray10 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray10
res_go_foldlArray10 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray10
arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
_ = arr_go_foldlArray10
for _, v_foldlArray10 := range *arr_go_foldlArray10 {
res_go_foldlArray10 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_6, "append"), s_9, gopurs_runtime.Apply(f_7, a1_10))
})
}), res_go_foldlArray10, v_foldlArray10)
}
return res_go_foldlArray10
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray9 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray9
res_go_foldlArray9 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
_ = res_go_foldlArray9
arr_go_foldlArray9 := (*[]gopurs_runtime.Value)(arr_val_foldlArray9.UnsafePtr)
_ = arr_go_foldlArray9
for _, v_foldlArray9 := range *arr_go_foldlArray9 {
res_go_foldlArray9 = gopurs_runtime.Apply2(f_6, res_go_foldlArray9, v_foldlArray9)
}
return res_go_foldlArray9
}()
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = __local_var_8_8
// TAST (Let): __local_var_9_9 -> *Constructor_Data_Maybe_Just
__local_var_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(f_6, a1_9)
_ = __local_var_10_11
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr == nil) {
__t12 = a1_9
goto end_branch_12
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr != nil) {
__t12 = gopurs_runtime.Apply(__local_var_10_11, (*Constructor_Data_Maybe_Just)(v2_11.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
_ = __local_var_10_10
return gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_10_10, x_11)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
_ = __local_var_9_9
var __t13 gopurs_runtime.Value
{
if (__local_var_9_9 == nil) {
__t13 = (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
goto end_branch_13
} else {

}
}
{
if (__local_var_9_9 != nil) {
__t13 = gopurs_runtime.Apply(__local_var_8_8, (__local_var_9_9).V0)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
})}
_ = __local_var_5_4
// TAST (Let): __local_var_6_14 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_6_14 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_4.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_7.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_14)}).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_8_15_0 gopurs_runtime.Value
go__go_8_15_0 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop int64 = v_9_loop_val.IntVal
var v1_10_loop gopurs_runtime.Value = v1_10_loop_val
go__go_8_15_0:
for {
if false { continue go__go_8_15_0 }
var v_9 int64 = v_9_loop
_ = v_9
var v1_10 gopurs_runtime.Value = v1_10_loop
_ = v1_10
var __t19 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr != nil) {
var __t18 gopurs_runtime.Value
{
var __t_tag_16 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1
if (__t_tag_16 == nil) {
__t18 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_18
} else {

}
}
{
var __t17 bool
{
if (v_9) > (0) {
__t17 = false
goto end_branch_17
} else {

}
}
{
__t17 = true
}
end_branch_17:
if __t17 {
__t18 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_18
} else {

}
}
{
v_9_loop = (v_9) - (1)
v1_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1)}
continue go__go_8_15_0
__t18 = gopurs_runtime.Value{}
}
end_branch_18:
__t19 = __t18
goto end_branch_19
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr == nil) {
__t19 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return x_11
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_14)})
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}
}()
})
})
return gopurs_runtime.Apply2(go__go_8_15_0, gopurs_runtime.Int(n_7.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_4.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_14)})))})
}))
_ = __local_var_5_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_20 -> gopurs_runtime.Value
Monad0_8_20 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_8_20
// TAST (Let): pure_9_21 -> gopurs_runtime.Value
pure_9_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_20, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_21
// TAST (Let): Bind1_10_22 -> *Constructor_Control_Bind_Bind
Bind1_10_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_20, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_22
// TAST (Let): __local_var_11_25 -> gopurs_runtime.Value
__local_var_11_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
var __t28 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t28 = false
goto end_branch_28
} else {

}
}
{
__t28 = true
}
end_branch_28:
if __t28 {
__t29 = gopurs_runtime.Apply(pure_9_21, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_29
} else {

}
}
{
// TAST (Let): __local_var_12_26 -> gopurs_runtime.Value
__local_var_12_26 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_26
// TAST (Let): __local_var_13_27 -> gopurs_runtime.Value
__local_var_13_27 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_27
__t29 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_22.V1), __local_var_5_3, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_21, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_26)})}, gopurs_runtime.Int((__local_var_13_27.IntVal) - (1))})}})})
}))
}
end_branch_29:
return __t29
}))
_ = __local_var_11_25
// TAST (Let): __local_var_12_30 -> gopurs_runtime.Value
__local_var_12_30 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_30
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_20, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t23 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t24 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_24
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t24 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_24
} else {

}
}
{
__t24 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_24:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t24)}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_25, gopurs_runtime.Apply(__local_var_12_30, x_13))
})))))
}))
}))
}

func Call_Data_String_Gen_genAlphaLowercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
_ = __local_var_5_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_4 -> gopurs_runtime.Value
Monad0_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_8_4
// TAST (Let): pure_9_5 -> gopurs_runtime.Value
pure_9_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_5
// TAST (Let): Bind1_10_6 -> *Constructor_Control_Bind_Bind
Bind1_10_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_6
// TAST (Let): __local_var_11_9 -> gopurs_runtime.Value
__local_var_11_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
var __t12 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t12 = false
goto end_branch_12
} else {

}
}
{
__t12 = true
}
end_branch_12:
if __t12 {
__t13 = gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_13
} else {

}
}
{
// TAST (Let): __local_var_12_10 -> gopurs_runtime.Value
__local_var_12_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_10
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_11
__t13 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_6.V1), __local_var_5_3, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_5, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_10)})}, gopurs_runtime.Int((__local_var_13_11.IntVal) - (1))})}})})
}))
}
end_branch_13:
return __t13
}))
_ = __local_var_11_9
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_14
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t7 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t8 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_8
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t8 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t8)}
})), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_9, gopurs_runtime.Apply(__local_var_12_14, x_13))
})))))
}))
}))
}

func Call_Data_String_Gen_genString__3531165174(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V3), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_8_3 -> gopurs_runtime.Value
Monad0_8_3 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_8_3
// TAST (Let): pure_9_4 -> gopurs_runtime.Value
pure_9_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_4
// TAST (Let): Bind1_10_5 -> *Constructor_Control_Bind_Bind
Bind1_10_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_5
// TAST (Let): __local_var_11_8 -> gopurs_runtime.Value
__local_var_11_8 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
var __t11 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1.IntVal) > (0) {
__t11 = false
goto end_branch_11
} else {

}
}
{
__t11 = true
}
end_branch_11:
if __t11 {
__t12 = gopurs_runtime.Apply(pure_9_4, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0))}})})
goto end_branch_12
} else {

}
}
{
// TAST (Let): __local_var_12_9 -> gopurs_runtime.Value
__local_var_12_9 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V0
_ = __local_var_12_9
// TAST (Let): __local_var_13_10 -> gopurs_runtime.Value
__local_var_13_10 := (*Constructor_Data_Tuple_Tuple)(v_11.UnsafePtr).V1
_ = __local_var_13_10
__t12 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_5.V1), genChar_5, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_9_4, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_14, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_12_9)})}, gopurs_runtime.Int((__local_var_13_10.IntVal) - (1))})}})})
}))
}
end_branch_12:
return __t12
}))
_ = __local_var_11_8
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_12_13
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_8_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply5(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t6 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr == nil) {
__t7 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_7
} else {

}
}
{
if (v_11.Type == 9 && v_11.IntVal == 759514854 && v_11.UnsafePtr != nil) {
__t7 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v_11.UnsafePtr).V1)}})}}
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
})), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_8, gopurs_runtime.Apply(__local_var_12_13, x_13))
})))))
}))
}))
})
}



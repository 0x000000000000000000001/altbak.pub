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
return Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
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
return Call_Data_String_Gen_genString__3531165174(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_Data_String_Gen_genString__3531165174
}

func Call_Data_String_Gen_genString(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V3), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply4(Get_Control_Monad_Gen_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()))}, genChar_5)))
}))
}))
})
}

func Call_Data_String_Gen_genUnicodeString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(Get_Data_Char_Gen_genUnicodeChar(), dictMonadGen_1))
}

func Call_Data_String_Gen_genDigitString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(Get_Data_Char_Gen_genDigitChar(), dictMonadGen_1))
}

func Call_Data_String_Gen_genAsciiString_prime(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(Get_Data_Char_Gen_genAsciiChar_prime(), dictMonadGen_1))
}

func Call_Data_String_Gen_genAsciiString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(Get_Data_Char_Gen_genAsciiChar(), dictMonadGen_1))
}

func Call_Data_String_Gen_genAlphaUppercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(Get_Data_Char_Gen_genAlphaUppercase(), dictMonadGen_1))
}

func Call_Data_String_Gen_genAlphaString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(Get_Data_Char_Gen_genAlpha(), dictMonadGen_1))
}

func Call_Data_String_Gen_genAlphaLowercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_Data_String_Gen_genString(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(Get_Data_Char_Gen_genAlphaLowercase(), dictMonadGen_1))
}

func Call_Data_String_Gen_genString__3531165174(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Functor0_4_2 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V3), gopurs_runtime.Int(1), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_Ord_max__2538992856(), gopurs_runtime.Int(1), gopurs_runtime.Int(size_6.IntVal)).IntVal)), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_7.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_2.V0), Get_Data_String_CodeUnits_fromCharArray(), gopurs_runtime.Apply4(Get_Control_Monad_Gen_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](Get_Data_Unfoldable_unfoldableArray()))}, genChar_5)))
}))
}))
})
}



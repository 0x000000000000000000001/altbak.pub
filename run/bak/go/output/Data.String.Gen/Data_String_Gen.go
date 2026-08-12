package Data_String_Gen

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Control_Monad_Gen_Class "gopurs/output/Control.Monad.Gen.Class"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Char_Gen "gopurs/output/Data.Char.Gen"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_genString gopurs_runtime.Value
var once_genString sync.Once
func Get_genString() gopurs_runtime.Value {
	once_genString.Do(func() {
		cache_genString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_genString
}

var cache_genString__gopurs_runtime_Value_3531165174 gopurs_runtime.Value
var once_genString__gopurs_runtime_Value_3531165174 sync.Once
func Get_genString__gopurs_runtime_Value_3531165174() gopurs_runtime.Value {
	once_genString__gopurs_runtime_Value_3531165174.Do(func() {
		cache_genString__gopurs_runtime_Value_3531165174 = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genString__gopurs_runtime_Value_3531165174(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1_box))
})
	})
	return cache_genString__gopurs_runtime_Value_3531165174
}

var cache_genUnicodeString gopurs_runtime.Value
var once_genUnicodeString sync.Once
func Get_genUnicodeString() gopurs_runtime.Value {
	once_genUnicodeString.Do(func() {
		cache_genUnicodeString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genUnicodeString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genUnicodeString
}

var cache_genDigitString gopurs_runtime.Value
var once_genDigitString sync.Once
func Get_genDigitString() gopurs_runtime.Value {
	once_genDigitString.Do(func() {
		cache_genDigitString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDigitString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genDigitString
}

var cache_genAsciiString_prime gopurs_runtime.Value
var once_genAsciiString_prime sync.Once
func Get_genAsciiString_prime() gopurs_runtime.Value {
	once_genAsciiString_prime.Do(func() {
		cache_genAsciiString_prime = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiString_prime(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAsciiString_prime
}

var cache_genAsciiString gopurs_runtime.Value
var once_genAsciiString sync.Once
func Get_genAsciiString() gopurs_runtime.Value {
	once_genAsciiString.Do(func() {
		cache_genAsciiString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAsciiString
}

var cache_genAlphaUppercaseString gopurs_runtime.Value
var once_genAlphaUppercaseString sync.Once
func Get_genAlphaUppercaseString() gopurs_runtime.Value {
	once_genAlphaUppercaseString.Do(func() {
		cache_genAlphaUppercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaUppercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAlphaUppercaseString
}

var cache_genAlphaString gopurs_runtime.Value
var once_genAlphaString sync.Once
func Get_genAlphaString() gopurs_runtime.Value {
	once_genAlphaString.Do(func() {
		cache_genAlphaString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAlphaString
}

var cache_genAlphaLowercaseString gopurs_runtime.Value
var once_genAlphaLowercaseString sync.Once
func Get_genAlphaLowercaseString() gopurs_runtime.Value {
	once_genAlphaLowercaseString.Do(func() {
		cache_genAlphaLowercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaLowercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return cache_genAlphaLowercaseString
}

func Call_genString(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
Bind1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
Functor0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Int(1), size_6)
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 1527465420) {
__t4 = size_6
goto end_branch_4
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 902936544) {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 380165415) {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Apply2(Bind1_3_1.V1, gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(1), __t4), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return newSize_7
}), gopurs_runtime.Apply2(Functor0_4_2.V0, pkg_Data_String_CodeUnits.Get_fromCharArray(), gopurs_runtime.Apply4(pkg_Control_Monad_Gen.Get_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, pkg_Data_Unfoldable.Get_unfoldableArray(), genChar_5)))
}))
}))
})
}

func Call_genString__gopurs_runtime_Value_3531165174(dictMonadRec_0_loop *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value], dictMonadGen_1_loop *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value] = dictMonadGen_1_loop
_ = dictMonadGen_1
Monad0_2_0 := gopurs_runtime.Apply(dictMonadGen_1.V0, gopurs_runtime.Value{})
_ = Monad0_2_0
Bind1_3_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
Functor0_4_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_2
return gopurs_runtime.Func(func(genChar_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadGen_1.V5, gopurs_runtime.Func(func(size_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Int(1), size_6)
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 1527465420) {
__t4 = size_6
goto end_branch_4
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 902936544) {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 380165415) {
__t4 = gopurs_runtime.Int(1)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Apply2(Bind1_3_1.V1, gopurs_runtime.Apply2(dictMonadGen_1.V3, gopurs_runtime.Int(1), __t4), gopurs_runtime.Func(func(newSize_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictMonadGen_1.V4, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return newSize_7
}), gopurs_runtime.Apply2(Functor0_4_2.V0, pkg_Data_String_CodeUnits.Get_fromCharArray(), gopurs_runtime.Apply4(pkg_Control_Monad_Gen.Get_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, pkg_Data_Unfoldable.Get_unfoldableArray(), genChar_5)))
}))
}))
})
}

func Call_genUnicodeString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genUnicodeChar(), dictMonadGen_1))
}

func Call_genDigitString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genDigitChar(), dictMonadGen_1))
}

func Call_genAsciiString_prime(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar_prime(), dictMonadGen_1))
}

func Call_genAsciiString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar(), dictMonadGen_1))
}

func Call_genAlphaUppercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaUppercase(), dictMonadGen_1))
}

func Call_genAlphaString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlpha(), dictMonadGen_1))
}

func Call_genAlphaLowercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Rec_Class.Constructor_MonadRec[gopurs_runtime.Value]](dictMonadRec_0), gopurs_runtime.CoerceToStruct[pkg_Control_Monad_Gen_Class.Constructor_MonadGen[gopurs_runtime.Value]](dictMonadGen_1)), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaLowercase(), dictMonadGen_1))
}



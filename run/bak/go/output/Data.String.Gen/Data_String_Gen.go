package Data_String_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Char_Gen "gopurs/output/Data.Char.Gen"
)

var max gopurs_runtime.Value
var once_max sync.Once
func Get_max() gopurs_runtime.Value {
	once_max.Do(func() {
		max = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_max(x_0_box, y_1_box)
})
	})
	return max
}

var genString gopurs_runtime.Value
var once_genString sync.Once
func Get_genString() gopurs_runtime.Value {
	once_genString.Do(func() {
		genString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genString
}

var genUnicodeString gopurs_runtime.Value
var once_genUnicodeString sync.Once
func Get_genUnicodeString() gopurs_runtime.Value {
	once_genUnicodeString.Do(func() {
		genUnicodeString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genUnicodeString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genUnicodeString
}

var genDigitString gopurs_runtime.Value
var once_genDigitString sync.Once
func Get_genDigitString() gopurs_runtime.Value {
	once_genDigitString.Do(func() {
		genDigitString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDigitString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genDigitString
}

var genAsciiString_prime gopurs_runtime.Value
var once_genAsciiString_prime sync.Once
func Get_genAsciiString_prime() gopurs_runtime.Value {
	once_genAsciiString_prime.Do(func() {
		genAsciiString_prime = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiString_prime(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genAsciiString_prime
}

var genAsciiString gopurs_runtime.Value
var once_genAsciiString sync.Once
func Get_genAsciiString() gopurs_runtime.Value {
	once_genAsciiString.Do(func() {
		genAsciiString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genAsciiString
}

var genAlphaUppercaseString gopurs_runtime.Value
var once_genAlphaUppercaseString sync.Once
func Get_genAlphaUppercaseString() gopurs_runtime.Value {
	once_genAlphaUppercaseString.Do(func() {
		genAlphaUppercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaUppercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genAlphaUppercaseString
}

var genAlphaString gopurs_runtime.Value
var once_genAlphaString sync.Once
func Get_genAlphaString() gopurs_runtime.Value {
	once_genAlphaString.Do(func() {
		genAlphaString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genAlphaString
}

var genAlphaLowercaseString gopurs_runtime.Value
var once_genAlphaLowercaseString sync.Once
func Get_genAlphaLowercaseString() gopurs_runtime.Value {
	once_genAlphaLowercaseString.Do(func() {
		genAlphaLowercaseString = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaLowercaseString(dictMonadRec_0_box, dictMonadGen_1_box)
})
	})
	return genAlphaLowercaseString
}

func Call_max(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
v_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), x_0, y_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 3866105248) {
__t1 = y_1
goto end_branch_1
} else {

}
}
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 1111389260) {
__t1 = x_0
goto end_branch_1
} else {

}
}
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 2098047435) {
__t1 = x_0
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

func Call_genString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
Bind1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_0
unfoldable1_3_1 := gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_unfoldable(), dictMonadRec_0, dictMonadGen_1, pkg_Data_Unfoldable.Get_unfoldableArray())
_ = unfoldable1_3_1
return gopurs_runtime.Func(func(genChar_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_1, "sized"), gopurs_runtime.Func(func(size_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_0, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "chooseInt"), gopurs_runtime.Int(1), Call_max(gopurs_runtime.Int(1), size_5)), gopurs_runtime.Func(func(newSize_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_1, "resize"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return newSize_6
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_String_CodeUnits.Get_fromCharArray(), gopurs_runtime.Apply(unfoldable1_3_1, genChar_4)))
}))
}))
})
}

func Call_genUnicodeString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genUnicodeChar(), dictMonadGen_1))
}

func Call_genDigitString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genDigitChar(), dictMonadGen_1))
}

func Call_genAsciiString_prime(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar_prime(), dictMonadGen_1))
}

func Call_genAsciiString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAsciiChar(), dictMonadGen_1))
}

func Call_genAlphaUppercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaUppercase(), dictMonadGen_1))
}

func Call_genAlphaString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlpha(), dictMonadGen_1))
}

func Call_genAlphaLowercaseString(dictMonadRec_0_loop gopurs_runtime.Value, dictMonadGen_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 gopurs_runtime.Value = dictMonadGen_1_loop
_ = dictMonadGen_1
return gopurs_runtime.Apply(Call_genString(dictMonadRec_0, dictMonadGen_1), gopurs_runtime.Apply(pkg_Data_Char_Gen.Get_genAlphaLowercase(), dictMonadGen_1))
}



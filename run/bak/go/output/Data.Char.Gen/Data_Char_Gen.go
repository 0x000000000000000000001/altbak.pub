package Data_Char_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
)

var toEnumWithDefaults gopurs_runtime.Value
var once_toEnumWithDefaults sync.Once
func Get_toEnumWithDefaults() gopurs_runtime.Value {
	once_toEnumWithDefaults.Do(func() {
		toEnumWithDefaults = gopurs_runtime.Func3(func(low_0 gopurs_runtime.Value, high_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_0 := gopurs_runtime.Apply(pkg_Data_Enum.Get_charToEnum(), x_2)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_3_0.StrVal == "Just")).IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(v_3_0.UnsafePtr)[0]
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_3_0.StrVal == "Nothing")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_2.IntVal < gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), pkg_Data_Bounded.Get_bottomChar()).IntVal)).IntVal != 0 {
__t2 = low_0
goto end_branch_2
} else {

}
}
{
__t2 = high_1
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
})
	})
	return toEnumWithDefaults
}

var foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		foldable1NonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())
	})
	return foldable1NonEmpty
}

var genUnicodeChar gopurs_runtime.Value
var once_genUnicodeChar sync.Once
func Get_genUnicodeChar() gopurs_runtime.Value {
	once_genUnicodeChar.Do(func() {
		genUnicodeChar = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar(), pkg_Data_Bounded.Get_topChar()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(65536)))
})
	})
	return genUnicodeChar
}

var genDigitChar gopurs_runtime.Value
var once_genDigitChar sync.Once
func Get_genDigitChar() gopurs_runtime.Value {
	once_genDigitChar.Do(func() {
		genDigitChar = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar(), pkg_Data_Bounded.Get_topChar()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(48), gopurs_runtime.Int(57)))
})
	})
	return genDigitChar
}

var genAsciiChar_prime gopurs_runtime.Value
var once_genAsciiChar_prime sync.Once
func Get_genAsciiChar_prime() gopurs_runtime.Value {
	once_genAsciiChar_prime.Do(func() {
		genAsciiChar_prime = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar(), pkg_Data_Bounded.Get_topChar()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(127)))
})
	})
	return genAsciiChar_prime
}

var genAsciiChar gopurs_runtime.Value
var once_genAsciiChar sync.Once
func Get_genAsciiChar() gopurs_runtime.Value {
	once_genAsciiChar.Do(func() {
		genAsciiChar = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar(), pkg_Data_Bounded.Get_topChar()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(32), gopurs_runtime.Int(127)))
})
	})
	return genAsciiChar
}

var genAlphaUppercase gopurs_runtime.Value
var once_genAlphaUppercase sync.Once
func Get_genAlphaUppercase() gopurs_runtime.Value {
	once_genAlphaUppercase.Do(func() {
		genAlphaUppercase = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar(), pkg_Data_Bounded.Get_topChar()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))
})
	})
	return genAlphaUppercase
}

var genAlphaLowercase gopurs_runtime.Value
var once_genAlphaLowercase sync.Once
func Get_genAlphaLowercase() gopurs_runtime.Value {
	once_genAlphaLowercase.Do(func() {
		genAlphaLowercase = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), pkg_Data_Bounded.Get_bottomChar(), pkg_Data_Bounded.Get_topChar()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
})
	})
	return genAlphaLowercase
}

var genAlpha gopurs_runtime.Value
var once_genAlpha sync.Once
func Get_genAlpha() gopurs_runtime.Value {
	once_genAlpha.Do(func() {
		genAlpha = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_oneOf(), dictMonadGen_0, Get_foldable1NonEmpty(), gopurs_runtime.Constructor2("NonEmpty", gopurs_runtime.Apply(Get_genAlphaLowercase(), dictMonadGen_0), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(Get_genAlphaUppercase(), dictMonadGen_0)})))
})
	})
	return genAlpha
}



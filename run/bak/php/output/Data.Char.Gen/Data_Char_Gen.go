package Data_Char_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	unsafe "unsafe"
)

var cache_toEnumWithDefaults gopurs_runtime.Value
var once_toEnumWithDefaults sync.Once
func Get_toEnumWithDefaults() gopurs_runtime.Value {
	once_toEnumWithDefaults.Do(func() {
		cache_toEnumWithDefaults = gopurs_runtime.Apply(pkg_Data_Enum.Get_toEnumWithDefaults(), pkg_Data_Enum.Get_boundedEnumChar())
	})
	return cache_toEnumWithDefaults
}

var cache_foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		cache_foldable1NonEmpty = gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())
	})
	return cache_foldable1NonEmpty
}

var cache_genUnicodeChar gopurs_runtime.Value
var once_genUnicodeChar sync.Once
func Get_genUnicodeChar() gopurs_runtime.Value {
	once_genUnicodeChar.Do(func() {
		cache_genUnicodeChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genUnicodeChar(dictMonadGen_0_box)
})
	})
	return cache_genUnicodeChar
}

var cache_genDigitChar gopurs_runtime.Value
var once_genDigitChar sync.Once
func Get_genDigitChar() gopurs_runtime.Value {
	once_genDigitChar.Do(func() {
		cache_genDigitChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDigitChar(dictMonadGen_0_box)
})
	})
	return cache_genDigitChar
}

var cache_genAsciiChar_prime gopurs_runtime.Value
var once_genAsciiChar_prime sync.Once
func Get_genAsciiChar_prime() gopurs_runtime.Value {
	once_genAsciiChar_prime.Do(func() {
		cache_genAsciiChar_prime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiChar_prime(dictMonadGen_0_box)
})
	})
	return cache_genAsciiChar_prime
}

var cache_genAsciiChar gopurs_runtime.Value
var once_genAsciiChar sync.Once
func Get_genAsciiChar() gopurs_runtime.Value {
	once_genAsciiChar.Do(func() {
		cache_genAsciiChar = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAsciiChar(dictMonadGen_0_box)
})
	})
	return cache_genAsciiChar
}

var cache_genAlphaUppercase gopurs_runtime.Value
var once_genAlphaUppercase sync.Once
func Get_genAlphaUppercase() gopurs_runtime.Value {
	once_genAlphaUppercase.Do(func() {
		cache_genAlphaUppercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaUppercase(dictMonadGen_0_box)
})
	})
	return cache_genAlphaUppercase
}

var cache_genAlphaLowercase gopurs_runtime.Value
var once_genAlphaLowercase sync.Once
func Get_genAlphaLowercase() gopurs_runtime.Value {
	once_genAlphaLowercase.Do(func() {
		cache_genAlphaLowercase = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaLowercase(dictMonadGen_0_box)
})
	})
	return cache_genAlphaLowercase
}

var cache_genAlpha gopurs_runtime.Value
var once_genAlpha sync.Once
func Get_genAlpha() gopurs_runtime.Value {
	once_genAlpha.Do(func() {
		cache_genAlpha = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlpha(dictMonadGen_0_box)
})
	})
	return cache_genAlpha
}

func Call_genUnicodeChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(0), gopurs_runtime.Int(65536)))
}

func Call_genDigitChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(48), gopurs_runtime.Int(57)))
}

func Call_genAsciiChar_prime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(0), gopurs_runtime.Int(127)))
}

func Call_genAsciiChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(32), gopurs_runtime.Int(127)))
}

func Call_genAlphaUppercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(65), gopurs_runtime.Int(90)))
}

func Call_genAlphaLowercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_toEnumWithDefaults(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}

func Call_genAlpha(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_oneOf(), dictMonadGen_0, Get_foldable1NonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(Get_genAlphaLowercase(), dictMonadGen_0), gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply(Get_genAlphaUppercase(), dictMonadGen_0)})})})
}



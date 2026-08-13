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
		cache_Data_Char_Gen_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), Get_Data_Foldable_foldableArray())))}
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
return gopurs_runtime.Apply3(Get_Control_Monad_Gen_oneOf(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_0))}, gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), Get_Data_Foldable_foldableArray())))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122))), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
}

func Call_Data_Char_Gen_genAlphaLowercase__4294897069(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}



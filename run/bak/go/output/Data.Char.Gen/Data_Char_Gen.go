package Data_Char_Gen

import (
	pkg_Control_Monad_Gen "gopurs/output/Control.Monad.Gen"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		cache_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(pkg_Data_NonEmpty.Get_foldable1NonEmpty(), pkg_Data_Foldable.Get_foldableArray())))}
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

var cache_genAlphaLowercase__gopurs_runtime_Value_4294897069 gopurs_runtime.Value
var once_genAlphaLowercase__gopurs_runtime_Value_4294897069 sync.Once
func Get_genAlphaLowercase__gopurs_runtime_Value_4294897069() gopurs_runtime.Value {
	once_genAlphaLowercase__gopurs_runtime_Value_4294897069.Do(func() {
		cache_genAlphaLowercase__gopurs_runtime_Value_4294897069 = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genAlphaLowercase__gopurs_runtime_Value_4294897069(dictMonadGen_0_box)
})
	})
	return cache_genAlphaLowercase__gopurs_runtime_Value_4294897069
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
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (x_2.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_1_0).IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(65536)))
}

func Call_genDigitChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (x_2.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_1_0).IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(48), gopurs_runtime.Int(57)))
}

func Call_genAsciiChar_prime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (x_2.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_1_0).IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(127)))
}

func Call_genAsciiChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (x_2.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_1_0).IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(32), gopurs_runtime.Int(127)))
}

func Call_genAlphaUppercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (x_2.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_1_0).IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(65), gopurs_runtime.Int(90)))
}

func Call_genAlphaLowercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (x_2.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_1_0).IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}

func Call_genAlphaLowercase__gopurs_runtime_Value_4294897069(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_1.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_1.Type == 9 && v_3_1.IntVal == 930809136 && v_3_1.UnsafePtr == nil) {
var __t3 gopurs_runtime.Value
{
var __t4 gopurs_runtime.Value
{
if (x_2.IntVal) < (gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_1_0).IntVal) {
__t4 = gopurs_runtime.Bool(true)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Bool(false)
}
end_branch_4:
if (__t4.IntVal) != (0) {
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top")
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(97), gopurs_runtime.Int(122)))
}

func Call_genAlpha(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply3(pkg_Control_Monad_Gen.Get_oneOf(), dictMonadGen_0, gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[*pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_foldable1NonEmpty()))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_genAlphaLowercase(dictMonadGen_0), gopurs_runtime.Array([]gopurs_runtime.Value{Call_genAlphaUppercase(dictMonadGen_0)})})})
}



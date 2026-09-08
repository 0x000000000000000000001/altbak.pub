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
		cache_Data_Char_Gen_foldable1NonEmpty = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_4217626592_4151366573(Rebox_Data_Char_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))))}
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

var cache_Data_Char_Gen_genAsciiChar_prime_ gopurs_runtime.Value
var once_Data_Char_Gen_genAsciiChar_prime_ sync.Once
func Get_Data_Char_Gen_genAsciiChar_prime_() gopurs_runtime.Value {
	once_Data_Char_Gen_genAsciiChar_prime_.Do(func() {
		cache_Data_Char_Gen_genAsciiChar_prime_ = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Char_Gen_genAsciiChar_prime_(dictMonadGen_0_box)
})
	})
	return cache_Data_Char_Gen_genAsciiChar_prime_
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

func Call_Data_Char_Gen_genUnicodeChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
var v_2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_742090555_3094389156(Rebox_Data_Char_Gen_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_charToEnum(), gopurs_runtime.Int(x_1.IntVal))))))})
var __t2 gopurs_runtime.Value
{
if (v_2_0 != nil) {
__t2 = (v_2_0).V0
goto end_branch_2
} else {

}
}
{
if (v_2_0 == nil) {
var __t1 gopurs_runtime.Value
{
if (x_1.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) {
__t1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(65536))))
}

func Call_Data_Char_Gen_genDigitChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
var v_2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_742090555_3094389156(Rebox_Data_Char_Gen_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_charToEnum(), gopurs_runtime.Int(x_1.IntVal))))))})
var __t2 gopurs_runtime.Value
{
if (v_2_0 != nil) {
__t2 = (v_2_0).V0
goto end_branch_2
} else {

}
}
{
if (v_2_0 == nil) {
var __t1 gopurs_runtime.Value
{
if (x_1.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) {
__t1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(48)), gopurs_runtime.Int(int64(57))))
}

func Call_Data_Char_Gen_genAsciiChar_prime_(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
var v_2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_742090555_3094389156(Rebox_Data_Char_Gen_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_charToEnum(), gopurs_runtime.Int(x_1.IntVal))))))})
var __t2 gopurs_runtime.Value
{
if (v_2_0 != nil) {
__t2 = (v_2_0).V0
goto end_branch_2
} else {

}
}
{
if (v_2_0 == nil) {
var __t1 gopurs_runtime.Value
{
if (x_1.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) {
__t1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(int64(127))))
}

func Call_Data_Char_Gen_genAsciiChar(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
var v_2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_742090555_3094389156(Rebox_Data_Char_Gen_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_charToEnum(), gopurs_runtime.Int(x_1.IntVal))))))})
var __t2 gopurs_runtime.Value
{
if (v_2_0 != nil) {
__t2 = (v_2_0).V0
goto end_branch_2
} else {

}
}
{
if (v_2_0 == nil) {
var __t1 gopurs_runtime.Value
{
if (x_1.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) {
__t1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(32)), gopurs_runtime.Int(int64(127))))
}

func Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
var v_2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_742090555_3094389156(Rebox_Data_Char_Gen_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_charToEnum(), gopurs_runtime.Int(x_1.IntVal))))))})
var __t2 gopurs_runtime.Value
{
if (v_2_0 != nil) {
__t2 = (v_2_0).V0
goto end_branch_2
} else {

}
}
{
if (v_2_0 == nil) {
var __t1 gopurs_runtime.Value
{
if (x_1.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) {
__t1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(65)), gopurs_runtime.Int(int64(90))))
}

func Call_Data_Char_Gen_genAlphaLowercase(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_2_0 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
var v_2_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_742090555_3094389156(Rebox_Data_Char_Gen_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Enum_charToEnum(), gopurs_runtime.Int(x_1.IntVal))))))})
var __t2 gopurs_runtime.Value
{
if (v_2_0 != nil) {
__t2 = (v_2_0).V0
goto end_branch_2
} else {

}
}
{
if (v_2_0 == nil) {
var __t1 gopurs_runtime.Value
{
if (x_1.IntVal) < (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal) {
__t1 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(int64(97)), gopurs_runtime.Int(int64(122))))
}

func Call_Data_Char_Gen_genAlpha(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply3(Get_Control_Monad_Gen_oneOf(), gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))}, gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Char_Gen_4217626592_4151366573(Rebox_Data_Char_Gen_4151366573_4217626592(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_NonEmpty_foldable1NonEmpty(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})))))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, Call_Data_Char_Gen_genAlphaLowercase(gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dictMonadGen_0))}), gopurs_runtime.Array([]gopurs_runtime.Value{Call_Data_Char_Gen_genAlphaUppercase(dictMonadGen_0)})}))})
}

func Rebox_Data_Char_Gen_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_Char_Gen_4151366573_4217626592(in *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Char_Gen_4217626592_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[*Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Char_Gen_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}



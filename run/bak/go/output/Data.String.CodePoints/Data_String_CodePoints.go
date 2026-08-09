package Data_String_CodePoints

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	unsafe "unsafe"
)

var cache_lessThanOrEq gopurs_runtime.Value
var once_lessThanOrEq sync.Once
func Get_lessThanOrEq() gopurs_runtime.Value {
	once_lessThanOrEq.Do(func() {
		cache_lessThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThanOrEq
}

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 380165415) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_greaterThan
}

var cache_lessThan gopurs_runtime.Value
var once_lessThan sync.Once
func Get_lessThan() gopurs_runtime.Value {
	once_lessThan.Do(func() {
		cache_lessThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_lessThan
}

var cache_compare gopurs_runtime.Value
var once_compare sync.Once
func Get_compare() gopurs_runtime.Value {
	once_compare.Do(func() {
		cache_compare = gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
	})
	return cache_compare
}

var cache_greaterThanOrEq gopurs_runtime.Value
var once_greaterThanOrEq sync.Once
func Get_greaterThanOrEq() gopurs_runtime.Value {
	once_greaterThanOrEq.Do(func() {
		cache_greaterThanOrEq = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func(func(a1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1527465420) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
})
})
}()
	})
	return cache_greaterThanOrEq
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(s_0_box.StrVal()))}
})
	})
	return cache_uncons
}

var cache_unconsButWithTuple gopurs_runtime.Value
var once_unconsButWithTuple sync.Once
func Get_unconsButWithTuple() gopurs_runtime.Value {
	once_unconsButWithTuple.Do(func() {
		cache_unconsButWithTuple = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_unconsButWithTuple(s_0_box.StrVal()))}
})
	})
	return cache_unconsButWithTuple
}

var cache_toCodePointArrayFallback gopurs_runtime.Value
var once_toCodePointArrayFallback sync.Once
func Get_toCodePointArrayFallback() gopurs_runtime.Value {
	once_toCodePointArrayFallback.Do(func() {
		cache_toCodePointArrayFallback = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toCodePointArrayFallback(s_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toCodePointArrayFallback
}

var cache_showCodePoint gopurs_runtime.Value
var once_showCodePoint sync.Once
func Get_showCodePoint() gopurs_runtime.Value {
	once_showCodePoint.Do(func() {
		cache_showCodePoint = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(CodePoint 0x"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), gopurs_runtime.Apply2(pkg_Data_Int.Get_toStringAs(), gopurs_runtime.Int(16), v_0)), gopurs_runtime.Str(")"))).StrVal())
}))
	})
	return cache_showCodePoint
}

var cache_unsafeCodePointAt0Fallback gopurs_runtime.Value
var once_unsafeCodePointAt0Fallback sync.Once
func Get_unsafeCodePointAt0Fallback() gopurs_runtime.Value {
	once_unsafeCodePointAt0Fallback.Do(func() {
		cache_unsafeCodePointAt0Fallback = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_unsafeCodePointAt0Fallback(s_0_box.StrVal()))
})
	})
	return cache_unsafeCodePointAt0Fallback
}

var cache_unsafeCodePointAt0 gopurs_runtime.Value
var once_unsafeCodePointAt0 sync.Once
func Get_unsafeCodePointAt0() gopurs_runtime.Value {
	once_unsafeCodePointAt0.Do(func() {
		cache_unsafeCodePointAt0 = gopurs_runtime.Apply(Get__unsafeCodePointAt0(), Get_unsafeCodePointAt0Fallback())
	})
	return cache_unsafeCodePointAt0
}

var cache_toCodePointArray gopurs_runtime.Value
var once_toCodePointArray sync.Once
func Get_toCodePointArray() gopurs_runtime.Value {
	once_toCodePointArray.Do(func() {
		cache_toCodePointArray = gopurs_runtime.Apply2(Get__toCodePointArray(), Get_toCodePointArrayFallback(), Get_unsafeCodePointAt0())
	})
	return cache_toCodePointArray
}

var cache_length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		cache_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_length(x_0_box.StrVal()))
})
	})
	return cache_length
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_lastIndexOf(p_0_box.StrVal(), s_1_box.StrVal()))}
})
	})
	return cache_lastIndexOf
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_indexOf(p_0_box.StrVal(), s_1_box.StrVal()))}
})
	})
	return cache_indexOf
}

var cache_fromCharCode gopurs_runtime.Value
var once_fromCharCode sync.Once
func Get_fromCharCode() gopurs_runtime.Value {
	once_fromCharCode.Do(func() {
		cache_fromCharCode = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Enum.Get_toEnumWithDefaults(), pkg_Data_Enum.Get_boundedEnumChar(), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top"))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), gopurs_runtime.Apply(__local_var_0_0, x_1))
})
}()
	})
	return cache_fromCharCode
}

var cache_singletonFallback gopurs_runtime.Value
var once_singletonFallback sync.Once
func Get_singletonFallback() gopurs_runtime.Value {
	once_singletonFallback.Do(func() {
		cache_singletonFallback = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_singletonFallback(v_0_box.IntVal))
})
	})
	return cache_singletonFallback
}

var cache_fromCodePointArray gopurs_runtime.Value
var once_fromCodePointArray sync.Once
func Get_fromCodePointArray() gopurs_runtime.Value {
	once_fromCodePointArray.Do(func() {
		cache_fromCodePointArray = gopurs_runtime.Apply(Get__fromCodePointArray(), Get_singletonFallback())
	})
	return cache_fromCodePointArray
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Apply(Get__singleton(), Get_singletonFallback())
	})
	return cache_singleton
}

var cache_takeFallback gopurs_runtime.Value
var once_takeFallback sync.Once
func Get_takeFallback() gopurs_runtime.Value {
	once_takeFallback.Do(func() {
		cache_takeFallback = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_takeFallback(v_0_box.IntVal, v1_1_box.StrVal()))
})
	})
	return cache_takeFallback
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Apply(Get__take(), Get_takeFallback())
	})
	return cache_take
}

var cache_lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		cache_lastIndexOf_prime = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_lastIndexOf_prime(p_0_box.StrVal(), i_1_box.IntVal, s_2_box.StrVal()))}
})
	})
	return cache_lastIndexOf_prime
}

var cache_splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		cache_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitAt(i_0_box.IntVal, s_1_box.StrVal())
})
	})
	return cache_splitAt
}

var cache_eqCodePoint gopurs_runtime.Value
var once_eqCodePoint sync.Once
func Get_eqCodePoint() gopurs_runtime.Value {
	once_eqCodePoint.Do(func() {
		cache_eqCodePoint = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.IntVal) == (y_1.IntVal))
})
}))
	})
	return cache_eqCodePoint
}

var cache_ordCodePoint gopurs_runtime.Value
var once_ordCodePoint sync.Once
func Get_ordCodePoint() gopurs_runtime.Value {
	once_ordCodePoint.Do(func() {
		cache_ordCodePoint = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCodePoint()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_compare(), x_0, y_1)
})
}))
	})
	return cache_ordCodePoint
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_drop(n_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_drop
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_indexOf_prime(p_0_box.StrVal(), i_1_box.IntVal, s_2_box.StrVal()))}
})
	})
	return cache_indexOf_prime
}

var cache_countTail gopurs_runtime.Value
var once_countTail sync.Once
func Get_countTail() gopurs_runtime.Value {
	once_countTail.Do(func() {
		cache_countTail = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value, accum_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_countTail(p_0_box, s_1_box.StrVal(), accum_2_box.IntVal))
})
	})
	return cache_countTail
}

var cache_countFallback gopurs_runtime.Value
var once_countFallback sync.Once
func Get_countFallback() gopurs_runtime.Value {
	once_countFallback.Do(func() {
		cache_countFallback = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_countFallback(p_0_box, s_1_box.StrVal()))
})
	})
	return cache_countFallback
}

var cache_countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		cache_countPrefix = gopurs_runtime.Apply2(Get__countPrefix(), Get_countFallback(), Get_unsafeCodePointAt0())
	})
	return cache_countPrefix
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_dropWhile(p_0_box, s_1_box.StrVal()))
})
	})
	return cache_dropWhile
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_takeWhile(p_0_box, s_1_box.StrVal()))
})
	})
	return cache_takeWhile
}

var cache_codePointFromChar gopurs_runtime.Value
var once_codePointFromChar sync.Once
func Get_codePointFromChar() gopurs_runtime.Value {
	once_codePointFromChar.Do(func() {
		cache_codePointFromChar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_codePointFromChar(x_0_box.StrVal()))
})
	})
	return cache_codePointFromChar
}

var cache_codePointAtFallback gopurs_runtime.Value
var once_codePointAtFallback sync.Once
func Get_codePointAtFallback() gopurs_runtime.Value {
	once_codePointAtFallback.Do(func() {
		cache_codePointAtFallback = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_codePointAtFallback(n_0_box.IntVal, s_1_box.StrVal()))}
})
	})
	return cache_codePointAtFallback
}

var cache_codePointAt gopurs_runtime.Value
var once_codePointAt sync.Once
func Get_codePointAt() gopurs_runtime.Value {
	once_codePointAt.Do(func() {
		cache_codePointAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_codePointAt(v_0_box.IntVal, v1_1_box.StrVal()))}
})
	})
	return cache_codePointAt
}

var cache_boundedCodePoint gopurs_runtime.Value
var once_boundedCodePoint sync.Once
func Get_boundedCodePoint() gopurs_runtime.Value {
	once_boundedCodePoint.Do(func() {
		cache_boundedCodePoint = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordCodePoint()
}), gopurs_runtime.Int(0), gopurs_runtime.Int(1114111))
	})
	return cache_boundedCodePoint
}

var cache_boundedEnumCodePoint gopurs_runtime.Value
var once_boundedEnumCodePoint sync.Once
func Get_boundedEnumCodePoint() gopurs_runtime.Value {
	once_boundedEnumCodePoint.Do(func() {
		cache_boundedEnumCodePoint = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedCodePoint()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumCodePoint()
}), gopurs_runtime.Int(1114112), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_greaterThanOrEq(), n_0, gopurs_runtime.Int(0)), gopurs_runtime.Apply2(Get_lessThanOrEq(), n_0, gopurs_runtime.Int(1114111))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, n_0})}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(__t0.UnsafePtr))}
}))
	})
	return cache_boundedEnumCodePoint
}

var cache_enumCodePoint gopurs_runtime.Value
var once_enumCodePoint sync.Once
func Get_enumCodePoint() gopurs_runtime.Value {
	once_enumCodePoint.Do(func() {
		cache_enumCodePoint = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordCodePoint()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "fromEnum"), a_0).IntVal) - (1))).UnsafePtr))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "fromEnum"), a_0).IntVal) + (1))).UnsafePtr))}
}))
	})
	return cache_enumCodePoint
}

func Call_uncons(s_0_loop string) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var s_0 string = s_0_loop
_ = s_0
__local_var_1_1 := gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0))
_ = __local_var_1_1
v_1_0 := gopurs_runtime.Apply6(Get__codePointAt(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
})
}), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), gopurs_runtime.Apply(__local_var_1_1, x_2))
}), gopurs_runtime.Int(0), gopurs_runtime.Str(s_0))
_ = v_1_0
var __t2 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0.IntVal), gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply3(Get__take(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})
}), gopurs_runtime.Int(1), gopurs_runtime.Str(s_0))), gopurs_runtime.Str(s_0)).StrVal()))})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__t2.UnsafePtr)
}

func Call_unconsButWithTuple(s_0_loop string) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var s_0 string = s_0_loop
_ = s_0
return (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(v_1, "head"), gopurs_runtime.RecordGet(v_1, "tail")})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(s_0))}).UnsafePtr)
}

func Call_toCodePointArrayFallback(s_0_loop string) []int64 {
var s_0 string = s_0_loop
_ = s_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr"), Get_unconsButWithTuple(), gopurs_runtime.Str(s_0)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_unsafeCodePointAt0Fallback(s_0_loop string) int64 {
var s_0 string = s_0_loop
_ = s_0
cu0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(s_0)))
_ = cu0_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(55296), cu0_1_0), gopurs_runtime.Apply2(Get_lessThanOrEq(), cu0_1_0, gopurs_runtime.Int(56319))), gopurs_runtime.Apply2(Get_greaterThan(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(s_0)), gopurs_runtime.Int(1))).IntVal) != (0) {
cu1_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(1), gopurs_runtime.Str(s_0)))
_ = cu1_2_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(56320), cu1_2_2), gopurs_runtime.Apply2(Get_lessThanOrEq(), cu1_2_2, gopurs_runtime.Int(57343))).IntVal) != (0) {
__t3 = gopurs_runtime.Int(((((cu0_1_0.IntVal) - (55296)) * (1024)) + ((cu1_2_2.IntVal) - (56320))) + (65536))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int(cu0_1_0.IntVal)
}
end_branch_3:
__t1 = gopurs_runtime.Int(__t3.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Int(cu0_1_0.IntVal)
}
end_branch_1:
return __t1.IntVal
}

func Call_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Str(x_0))))).IntVal
}

func Call_lastIndexOf(p_0_loop string, s_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 string = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), i_2, gopurs_runtime.Str(s_1)))))).IntVal)
}), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_lastIndexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1))).UnsafePtr)
}

func Call_indexOf(p_0_loop string, s_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 string = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), i_2, gopurs_runtime.Str(s_1)))))).IntVal)
}), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_indexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1))).UnsafePtr)
}

func Call_singletonFallback(v_0_loop int64) string {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThanOrEq(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(65535)).IntVal) != (0) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(v_0)).StrVal())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "div"), gopurs_runtime.Int((v_0) - (65536)), gopurs_runtime.Int(1024)).IntVal) + (55296))), gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod"), gopurs_runtime.Int((v_0) - (65536)), gopurs_runtime.Int(1024)).IntVal) + (56320)))).StrVal())
}
end_branch_0:
return __t0.StrVal()
}

func Call_takeFallback(v_0_loop int64, v1_1_loop string) string {
takeFallback:
for {
if false { continue takeFallback }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal) != (0) {
__t2 = gopurs_runtime.Str("")
goto end_branch_2
} else {

}
}
{
v2_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(v1_1))}
_ = v2_2_0
var __t1 gopurs_runtime.Value
{
if (v2_2_0.Type == 9 && v2_2_0.IntVal == 930809136 && v2_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, "head")), gopurs_runtime.Str(Call_takeFallback((v_0) - (1), gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_2_0.UnsafePtr).V0, "tail").StrVal()))).StrVal())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str(v1_1)
}
end_branch_1:
__t2 = gopurs_runtime.Str(__t1.StrVal())
}
end_branch_2:
return __t2.StrVal()
}
}

func Call_lastIndexOf_prime(p_0_loop string, i_1_loop int64, s_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 string = p_0_loop
_ = p_0
var i_1 int64 = i_1_loop
_ = i_1
var s_2 string = s_2_loop
_ = s_2
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), k_3, gopurs_runtime.Str(s_2)))))).IntVal)
}), gopurs_runtime.Apply3(pkg_Data_String_CodeUnits.Get_lastIndexOf_prime(), gopurs_runtime.Str(p_0), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2))), gopurs_runtime.Str(s_2))).UnsafePtr)
}

func Call_splitAt(i_0_loop int64, s_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
before_2_0 := gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(s_1))
_ = before_2_0
return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), before_2_0), gopurs_runtime.Str(s_1)).StrVal()), gopurs_runtime.Str(before_2_0.StrVal()))
}

func Call_drop(n_0_loop int64, s_1_loop string) string {
var n_0 int64 = n_0_loop
_ = n_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(n_0), gopurs_runtime.Str(s_1))), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_indexOf_prime(p_0_loop string, i_1_loop int64, s_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 string = p_0_loop
_ = p_0
var i_1 int64 = i_1_loop
_ = i_1
var s_2 string = s_2_loop
_ = s_2
s_prime_3_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2))), gopurs_runtime.Str(s_2))
_ = s_prime_3_0
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((i_1) + (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), k_4, s_prime_3_0))))).IntVal))
}), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_indexOf(), gopurs_runtime.Str(p_0), s_prime_3_0)).UnsafePtr)
}

func Call_countTail(p_0_loop gopurs_runtime.Value, s_1_loop string, accum_2_loop int64) int64 {
countTail:
for {
if false { continue countTail }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
var accum_2 int64 = accum_2_loop
_ = accum_2
v_3_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(s_1))}
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 930809136 && v_3_0.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_0.UnsafePtr).V0, "head")).IntVal) != (0) {
p_0_loop = p_0
s_1_loop = gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_0.UnsafePtr).V0, "tail").StrVal()
accum_2_loop = (accum_2) + (1)
continue countTail
__t2 = gopurs_runtime.Int(gopurs_runtime.Value{}.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Int(accum_2)
}
end_branch_2:
__t1 = gopurs_runtime.Int(__t2.IntVal)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Int(accum_2)
}
end_branch_1:
return __t1.IntVal
}
}

func Call_countFallback(p_0_loop gopurs_runtime.Value, s_1_loop string) int64 {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Int(Call_countTail(p_0, s_1, 0)).IntVal
}

func Call_dropWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1))), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_codePointFromChar(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), gopurs_runtime.Str(x_0)).IntVal
}

func Call_codePointAtFallback(n_0_loop int64, s_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
codePointAtFallback:
for {
if false { continue codePointAtFallback }
var n_0 int64 = n_0_loop
_ = n_0
var s_1 string = s_1_loop
_ = s_1
v_2_0 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(s_1))}
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (v_2_0.Type == 9 && v_2_0.IntVal == 930809136 && v_2_0.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
if (n_0) == (0) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "head")})}.UnsafePtr))}
goto end_branch_2
} else {

}
}
{
n_0_loop = (n_0) - (1)
s_1_loop = gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_2_0.UnsafePtr).V0, "tail").StrVal()
continue codePointAtFallback
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{}.UnsafePtr))}
}
end_branch_2:
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(__t2.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
}
end_branch_1:
return (*pkg_Data_Maybe.Constructor_Just[int64])(__t1.UnsafePtr)
}
}

func Call_codePointAt(v_0_loop int64, v1_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_lessThan(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_0
} else {

}
}
{
if (v_0) == (0) {
var __t1 gopurs_runtime.Value
{
if (v1_1) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}.UnsafePtr))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_unsafeCodePointAt0(), gopurs_runtime.Str(v1_1))})}.UnsafePtr))}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply6(Get__codePointAt(), Get_codePointAtFallback(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, Get_unsafeCodePointAt0(), gopurs_runtime.Int(v_0), gopurs_runtime.Str(v1_1)).UnsafePtr))}
}
end_branch_0:
return (*pkg_Data_Maybe.Constructor_Just[int64])(__t0.UnsafePtr)
}

func Get__codePointAt() gopurs_runtime.Value {
	return _Gopurs__CodePointAt
}

func Get__countPrefix() gopurs_runtime.Value {
	return _Gopurs__CountPrefix
}

func Get__fromCodePointArray() gopurs_runtime.Value {
	return _Gopurs__FromCodePointArray
}

func Get__singleton() gopurs_runtime.Value {
	return _Gopurs__Singleton
}

func Get__take() gopurs_runtime.Value {
	return _Gopurs__Take
}

func Get__toCodePointArray() gopurs_runtime.Value {
	return _Gopurs__ToCodePointArray
}

func Get__unsafeCodePointAt0() gopurs_runtime.Value {
	return _Gopurs__UnsafeCodePointAt0
}

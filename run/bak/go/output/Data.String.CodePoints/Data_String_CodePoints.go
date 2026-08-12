package Data_String_CodePoints

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_unsurrogate gopurs_runtime.Value
var once_unsurrogate sync.Once
func Get_unsurrogate() gopurs_runtime.Value {
	once_unsurrogate.Do(func() {
		cache_unsurrogate = gopurs_runtime.Func2(func(lead_0_box gopurs_runtime.Value, trail_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_unsurrogate(lead_0_box.IntVal, trail_1_box.IntVal))
})
	})
	return cache_unsurrogate
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
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(CodePoint 0x"), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_Int.Get_toStringAs(), gopurs_runtime.Int(16), gopurs_runtime.Int(v_0.IntVal)).StrVal())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
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
__local_var_0_0 := gopurs_runtime.Apply2(Get_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "top").StrVal()))
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
return gopurs_runtime.Bool((Call_eq__2843686287(gopurs_runtime.Int(x_0.IntVal), gopurs_runtime.Int(y_1.IntVal)).IntVal) != (0))
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_compare__372254389(gopurs_runtime.Int(x_0.IntVal), gopurs_runtime.Int(y_1.IntVal)).IntVal)), UnsafePtr: nil}
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
}), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(1114111), gopurs_runtime.Int(1)).IntVal), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(gopurs_runtime.Int(n_0.IntVal), gopurs_runtime.Int(0))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(n_0.IntVal), gopurs_runtime.Int(1114111))).IntVal) != (0))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(n_0.IntVal)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0))}
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
}), gopurs_runtime.Apply2(Get_defaultPred__2391565248(), gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "toEnum"), gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "fromEnum")), gopurs_runtime.Apply2(Get_defaultSucc__2391565248(), gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "toEnum"), gopurs_runtime.RecordGet(Get_boundedEnumCodePoint(), "fromEnum")))
	})
	return cache_enumCodePoint
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_bottom__338427193 gopurs_runtime.Value
var once_bottom__338427193 sync.Once
func Get_bottom__338427193() gopurs_runtime.Value {
	once_bottom__338427193.Do(func() {
		cache_bottom__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom__338427193(dict_0_box)
})
	})
	return cache_bottom__338427193
}

var cache_top__338427193 gopurs_runtime.Value
var once_top__338427193 sync.Once
func Get_top__338427193() gopurs_runtime.Value {
	once_top__338427193.Do(func() {
		cache_top__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top__338427193(dict_0_box)
})
	})
	return cache_top__338427193
}

var cache_defaultPred__2391565248 gopurs_runtime.Value
var once_defaultPred__2391565248 sync.Once
func Get_defaultPred__2391565248() gopurs_runtime.Value {
	once_defaultPred__2391565248.Do(func() {
		cache_defaultPred__2391565248 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__2391565248(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultPred__2391565248
}

var cache_defaultPred__1581620096 gopurs_runtime.Value
var once_defaultPred__1581620096 sync.Once
func Get_defaultPred__1581620096() gopurs_runtime.Value {
	once_defaultPred__1581620096.Do(func() {
		cache_defaultPred__1581620096 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__1581620096(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultPred__1581620096
}

var cache_defaultPred__2204581824 gopurs_runtime.Value
var once_defaultPred__2204581824 sync.Once
func Get_defaultPred__2204581824() gopurs_runtime.Value {
	once_defaultPred__2204581824.Do(func() {
		cache_defaultPred__2204581824 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultPred__2204581824(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultPred__2204581824
}

var cache_defaultSucc__2391565248 gopurs_runtime.Value
var once_defaultSucc__2391565248 sync.Once
func Get_defaultSucc__2391565248() gopurs_runtime.Value {
	once_defaultSucc__2391565248.Do(func() {
		cache_defaultSucc__2391565248 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__2391565248(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultSucc__2391565248
}

var cache_defaultSucc__1581620096 gopurs_runtime.Value
var once_defaultSucc__1581620096 sync.Once
func Get_defaultSucc__1581620096() gopurs_runtime.Value {
	once_defaultSucc__1581620096.Do(func() {
		cache_defaultSucc__1581620096 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__1581620096(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultSucc__1581620096
}

var cache_defaultSucc__2204581824 gopurs_runtime.Value
var once_defaultSucc__2204581824 sync.Once
func Get_defaultSucc__2204581824() gopurs_runtime.Value {
	once_defaultSucc__2204581824.Do(func() {
		cache_defaultSucc__2204581824 = gopurs_runtime.Func3(func(toEnum_prime_0_box gopurs_runtime.Value, fromEnum_prime_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_defaultSucc__2204581824(toEnum_prime_0_box, fromEnum_prime_1_box, a_2_box))}
})
	})
	return cache_defaultSucc__2204581824
}

var cache_fromEnum__1649438469 gopurs_runtime.Value
var once_fromEnum__1649438469 sync.Once
func Get_fromEnum__1649438469() gopurs_runtime.Value {
	once_fromEnum__1649438469.Do(func() {
		cache_fromEnum__1649438469 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__1649438469(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromEnum__1649438469
}

var cache_fromEnum__679972887 gopurs_runtime.Value
var once_fromEnum__679972887 sync.Once
func Get_fromEnum__679972887() gopurs_runtime.Value {
	once_fromEnum__679972887.Do(func() {
		cache_fromEnum__679972887 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__679972887(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[*pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_fromEnum__679972887
}

var cache_fromEnum__1606852103 gopurs_runtime.Value
var once_fromEnum__1606852103 sync.Once
func Get_fromEnum__1606852103() gopurs_runtime.Value {
	once_fromEnum__1606852103.Do(func() {
		cache_fromEnum__1606852103 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__1606852103(__eta0_0_box)
})
	})
	return cache_fromEnum__1606852103
}

var cache_fromEnum__1637084359 gopurs_runtime.Value
var once_fromEnum__1637084359 sync.Once
func Get_fromEnum__1637084359() gopurs_runtime.Value {
	once_fromEnum__1637084359.Do(func() {
		cache_fromEnum__1637084359 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromEnum__1637084359(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_fromEnum__1637084359
}

var cache_toEnum__4261336164 gopurs_runtime.Value
var once_toEnum__4261336164 sync.Once
func Get_toEnum__4261336164() gopurs_runtime.Value {
	once_toEnum__4261336164.Do(func() {
		cache_toEnum__4261336164 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__4261336164(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toEnum__4261336164
}

var cache_toEnum__3317293286 gopurs_runtime.Value
var once_toEnum__3317293286 sync.Once
func Get_toEnum__3317293286() gopurs_runtime.Value {
	once_toEnum__3317293286.Do(func() {
		cache_toEnum__3317293286 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnum__3317293286(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_toEnum__3317293286
}

var cache_toEnumWithDefaults__3941305703 gopurs_runtime.Value
var once_toEnumWithDefaults__3941305703 sync.Once
func Get_toEnumWithDefaults__3941305703() gopurs_runtime.Value {
	once_toEnumWithDefaults__3941305703.Do(func() {
		cache_toEnumWithDefaults__3941305703 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, __eta2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnumWithDefaults__3941305703(__eta0_0_box, __eta1_1_box, __eta2_2_box)
})
	})
	return cache_toEnumWithDefaults__3941305703
}

var cache_toEnumWithDefaults__3558602759 gopurs_runtime.Value
var once_toEnumWithDefaults__3558602759 sync.Once
func Get_toEnumWithDefaults__3558602759() gopurs_runtime.Value {
	once_toEnumWithDefaults__3558602759.Do(func() {
		cache_toEnumWithDefaults__3558602759 = gopurs_runtime.Func(func(dictBoundedEnum_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toEnumWithDefaults__3558602759(gopurs_runtime.CoerceToStruct[pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]](dictBoundedEnum_0_box))
})
	})
	return cache_toEnumWithDefaults__3558602759
}

var cache_eq__2843686287 gopurs_runtime.Value
var once_eq__2843686287 sync.Once
func Get_eq__2843686287() gopurs_runtime.Value {
	once_eq__2843686287.Do(func() {
		cache_eq__2843686287 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2843686287(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2843686287
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_div__2185172824 gopurs_runtime.Value
var once_div__2185172824 sync.Once
func Get_div__2185172824() gopurs_runtime.Value {
	once_div__2185172824.Do(func() {
		cache_div__2185172824 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2185172824(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_div__2185172824
}

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_mod__2185172824 gopurs_runtime.Value
var once_mod__2185172824 sync.Once
func Get_mod__2185172824() gopurs_runtime.Value {
	once_mod__2185172824.Do(func() {
		cache_mod__2185172824 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2185172824(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mod__2185172824
}

var cache_mod__2579358968 gopurs_runtime.Value
var once_mod__2579358968 sync.Once
func Get_mod__2579358968() gopurs_runtime.Value {
	once_mod__2579358968.Do(func() {
		cache_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__2579358968
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__291265340 gopurs_runtime.Value
var once_map__291265340 sync.Once
func Get_map__291265340() gopurs_runtime.Value {
	once_map__291265340.Do(func() {
		cache_map__291265340 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__291265340(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v1_1_box)))}
})
	})
	return cache_map__291265340
}

var cache_map__2116777468 gopurs_runtime.Value
var once_map__2116777468 sync.Once
func Get_map__2116777468() gopurs_runtime.Value {
	once_map__2116777468.Do(func() {
		cache_map__2116777468 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__2116777468(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__2116777468
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_fromJust__2181618881 gopurs_runtime.Value
var once_fromJust__2181618881 sync.Once
func Get_fromJust__2181618881() gopurs_runtime.Value {
	once_fromJust__2181618881.Do(func() {
		cache_fromJust__2181618881 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__2181618881(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__2181618881
}

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_isNothing__2591355336 gopurs_runtime.Value
var once_isNothing__2591355336 sync.Once
func Get_isNothing__2591355336() gopurs_runtime.Value {
	once_isNothing__2591355336.Do(func() {
		cache_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2591355336
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
}

var cache_compare__372254389 gopurs_runtime.Value
var once_compare__372254389 sync.Once
func Get_compare__372254389() gopurs_runtime.Value {
	once_compare__372254389.Do(func() {
		cache_compare__372254389 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__372254389(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_compare__372254389
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_greaterThanOrEq__4087042607 gopurs_runtime.Value
var once_greaterThanOrEq__4087042607 sync.Once
func Get_greaterThanOrEq__4087042607() gopurs_runtime.Value {
	once_greaterThanOrEq__4087042607.Do(func() {
		cache_greaterThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__4087042607
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_lessThan__1710332219 gopurs_runtime.Value
var once_lessThan__1710332219 sync.Once
func Get_lessThan__1710332219() gopurs_runtime.Value {
	once_lessThan__1710332219.Do(func() {
		cache_lessThan__1710332219 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1710332219(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1710332219
}

var cache_lessThan__1697837627 gopurs_runtime.Value
var once_lessThan__1697837627 sync.Once
func Get_lessThan__1697837627() gopurs_runtime.Value {
	once_lessThan__1697837627.Do(func() {
		cache_lessThan__1697837627 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1697837627(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1697837627
}

var cache_lessThan__4087042607 gopurs_runtime.Value
var once_lessThan__4087042607 sync.Once
func Get_lessThan__4087042607() gopurs_runtime.Value {
	once_lessThan__4087042607.Do(func() {
		cache_lessThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThan__4087042607
}

var cache_lessThan__1409282474 gopurs_runtime.Value
var once_lessThan__1409282474 sync.Once
func Get_lessThan__1409282474() gopurs_runtime.Value {
	once_lessThan__1409282474.Do(func() {
		cache_lessThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThan__1409282474
}

var cache_lessThanOrEq__4087042607 gopurs_runtime.Value
var once_lessThanOrEq__4087042607 sync.Once
func Get_lessThanOrEq__4087042607() gopurs_runtime.Value {
	once_lessThanOrEq__4087042607.Do(func() {
		cache_lessThanOrEq__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_lessThanOrEq__4087042607
}

var cache_lessThanOrEq__1409282474 gopurs_runtime.Value
var once_lessThanOrEq__1409282474 sync.Once
func Get_lessThanOrEq__1409282474() gopurs_runtime.Value {
	once_lessThanOrEq__1409282474.Do(func() {
		cache_lessThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_lessThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_lessThanOrEq__1409282474
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_add__560788792 gopurs_runtime.Value
var once_add__560788792 sync.Once
func Get_add__560788792() gopurs_runtime.Value {
	once_add__560788792.Do(func() {
		cache_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_add__560788792
}

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__560788792 gopurs_runtime.Value
var once_mul__560788792 sync.Once
func Get_mul__560788792() gopurs_runtime.Value {
	once_mul__560788792.Do(func() {
		cache_mul__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_mul__560788792
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

var cache_fst__20422131 gopurs_runtime.Value
var once_fst__20422131 sync.Once
func Get_fst__20422131() gopurs_runtime.Value {
	once_fst__20422131.Do(func() {
		cache_fst__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__20422131
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_unfoldableArray__644327338 gopurs_runtime.Value
var once_unfoldableArray__644327338 sync.Once
func Get_unfoldableArray__644327338() gopurs_runtime.Value {
	once_unfoldableArray__644327338.Do(func() {
		cache_unfoldableArray__644327338 = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unfoldable1.Get_unfoldable1Array()
}), gopurs_runtime.Apply4(pkg_Data_Unfoldable.Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldableArray__644327338
}

var cache_unfoldr__1842498883 gopurs_runtime.Value
var once_unfoldr__1842498883 sync.Once
func Get_unfoldr__1842498883() gopurs_runtime.Value {
	once_unfoldr__1842498883.Do(func() {
		cache_unfoldr__1842498883 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1842498883(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_unfoldr__1842498883
}

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unfoldable1Array__4196906331 gopurs_runtime.Value
var once_unfoldable1Array__4196906331 sync.Once
func Get_unfoldable1Array__4196906331() gopurs_runtime.Value {
	once_unfoldable1Array__4196906331.Do(func() {
		cache_unfoldable1Array__4196906331 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(pkg_Data_Unfoldable1.Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldable1Array__4196906331
}

var cache_unsafePartial__905556445 gopurs_runtime.Value
var once_unsafePartial__905556445 sync.Once
func Get_unsafePartial__905556445() gopurs_runtime.Value {
	once_unsafePartial__905556445.Do(func() {
		cache_unsafePartial__905556445 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__905556445
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1090765981 gopurs_runtime.Value
var once_unsafePartial__1090765981 sync.Once
func Get_unsafePartial__1090765981() gopurs_runtime.Value {
	once_unsafePartial__1090765981.Do(func() {
		cache_unsafePartial__1090765981 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1090765981
}

func Call_unsurrogate(lead_0_loop int64, trail_1_loop int64) int64 {
var lead_0 int64 = lead_0_loop
_ = lead_0
var trail_1 int64 = trail_1_loop
_ = trail_1
return Call_add__560788792(gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(Call_mul__560788792(gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(lead_0), gopurs_runtime.Int(55296)).IntVal), gopurs_runtime.Int(1024)).IntVal), gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(trail_1), gopurs_runtime.Int(56320)).IntVal)).IntVal), gopurs_runtime.Int(65536)).IntVal
}

func Call_uncons(s_0_loop string) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var s_0 string = s_0_loop
_ = s_0
__local_var_1_2 := gopurs_runtime.Apply(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0))
_ = __local_var_1_2
__local_var_1_1 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), gopurs_runtime.Apply(__local_var_1_2, x_2))
})
_ = __local_var_1_1
v_1_0 := gopurs_runtime.Apply6(Get__codePointAt(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
})
}), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
}), gopurs_runtime.Int(0), gopurs_runtime.Str(s_0))
_ = v_1_0
var __t3 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1_0.UnsafePtr).V0.IntVal), gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(gopurs_runtime.Apply3(Get__take(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})
}), gopurs_runtime.Int(1), gopurs_runtime.Str(s_0)).StrVal())).IntVal), gopurs_runtime.Str(s_0)).StrVal()))})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3)
}

func Call_unconsButWithTuple(s_0_loop string) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, string]] {
var s_0 string = s_0_loop
_ = s_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, string]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__2116777468(gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.RecordGet(v_1, "head").IntVal), gopurs_runtime.Str(gopurs_runtime.RecordGet(v_1, "tail").StrVal())})}
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(s_0))})))})
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
cu0_1_0 := Call_fromEnum__1606852103(gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(s_0)).StrVal())).IntVal
_ = cu0_1_0
var __t3 int64
{
if (Call_conj__3676519832(gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(55296), gopurs_runtime.Int(cu0_1_0))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(cu0_1_0), gopurs_runtime.Int(56319))).IntVal) != (0))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(s_0)).IntVal), gopurs_runtime.Int(1))).IntVal) != (0))).IntVal) != (0) {
cu1_2_1 := Call_fromEnum__1606852103(gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(1), gopurs_runtime.Str(s_0)).StrVal())).IntVal
_ = cu1_2_1
var __t2 int64
{
if (Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(56320), gopurs_runtime.Int(cu1_2_1))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(cu1_2_1), gopurs_runtime.Int(57343))).IntVal) != (0))).IntVal) != (0) {
__t2 = gopurs_runtime.Int(Call_unsurrogate(cu0_1_0, cu1_2_1)).IntVal
goto end_branch_2
} else {

}
}
{
__t2 = cu0_1_0
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = cu0_1_0
}
end_branch_3:
return __t3
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
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__291265340(gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Str(s_1)).StrVal()))))).IntVal)
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_lastIndexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1)))))})
}

func Call_indexOf(p_0_loop string, s_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 string = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__291265340(gopurs_runtime.Func(func(i_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int(i_2.IntVal), gopurs_runtime.Str(s_1)).StrVal()))))).IntVal)
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_indexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1)))))})
}

func Call_singletonFallback(v_0_loop int64) string {
var v_0 int64 = v_0_loop
_ = v_0
var __t0 string
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(65535))).IntVal) != (0) {
__t0 = gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(v_0)).StrVal()
goto end_branch_0
} else {

}
}
{
__t0 = Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(Call_div__2185172824(gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(65536)).IntVal), gopurs_runtime.Int(1024)).IntVal), gopurs_runtime.Int(55296)).IntVal)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(Call_mod__2185172824(gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(65536)).IntVal), gopurs_runtime.Int(1024)).IntVal), gopurs_runtime.Int(56320)).IntVal)).StrVal())).StrVal()
}
end_branch_0:
return __t0
}

func Call_takeFallback(v_0_loop int64, v1_1_loop string) string {
takeFallback:
for {
if false { continue takeFallback }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
var __t2 string
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t2 = ""
goto end_branch_2
} else {

}
}
{
v2_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(v1_1))})
_ = v2_2_0
var __t1 string
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2_0)}.UnsafePtr != nil) {
__t1 = Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.Int(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2_0)}.UnsafePtr).V0, "head").IntVal)).StrVal()), gopurs_runtime.Str(gopurs_runtime.Str(Call_takeFallback(Call_sub__1043827704(gopurs_runtime.Int(v_0), gopurs_runtime.Int(1)).IntVal, gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2_0)}.UnsafePtr).V0, "tail").StrVal())).StrVal())).StrVal()
goto end_branch_1
} else {

}
}
{
__t1 = v1_1
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return __t2
}
}

func Call_lastIndexOf_prime(p_0_loop string, i_1_loop int64, s_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 string = p_0_loop
_ = p_0
var i_1 int64 = i_1_loop
_ = i_1
var s_2 string = s_2_loop
_ = s_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__291265340(gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int(k_3.IntVal), gopurs_runtime.Str(s_2)).StrVal()))))).IntVal)
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply3(pkg_Data_String_CodeUnits.Get_lastIndexOf_prime(), gopurs_runtime.Str(p_0), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2)).StrVal())).IntVal), gopurs_runtime.Str(s_2)))))})
}

func Call_splitAt(i_0_loop int64, s_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
before_2_0 := gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(s_1))
_ = before_2_0
return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(before_2_0.StrVal())).IntVal), gopurs_runtime.Str(s_1)).StrVal()), gopurs_runtime.Str(before_2_0.StrVal()))
}

func Call_drop(n_0_loop int64, s_1_loop string) string {
var n_0 int64 = n_0_loop
_ = n_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(n_0), gopurs_runtime.Str(s_1)).StrVal())).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_indexOf_prime(p_0_loop string, i_1_loop int64, s_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var p_0 string = p_0_loop
_ = p_0
var i_1 int64 = i_1_loop
_ = i_1
var s_2 string = s_2_loop
_ = s_2
s_prime_3_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2)).StrVal())).IntVal), gopurs_runtime.Str(s_2)).StrVal()
_ = s_prime_3_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__291265340(gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(i_1), gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int(k_4.IntVal), gopurs_runtime.Str(s_prime_3_0)).StrVal()))))).IntVal)).IntVal)
}), gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_indexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_prime_3_0)))))})
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
v_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(s_1))})
_ = v_3_0
var __t2 int64
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr != nil) {
var __t1 int64
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0, "head").IntVal)).IntVal) != (0) {
p_0_loop = p_0
s_1_loop = gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_3_0)}.UnsafePtr).V0, "tail").StrVal()
accum_2_loop = Call_add__560788792(gopurs_runtime.Int(accum_2), gopurs_runtime.Int(1)).IntVal
continue countTail
__t1 = gopurs_runtime.Value{}.IntVal
goto end_branch_1
} else {

}
}
{
__t1 = accum_2
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = accum_2
}
end_branch_2:
return __t2
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
return gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)).IntVal), gopurs_runtime.Str(s_1)).StrVal())).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_takeWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_countPrefix(), p_0, gopurs_runtime.Str(s_1)).IntVal), gopurs_runtime.Str(s_1)).StrVal()
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
v_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_uncons(s_1))})
_ = v_2_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if (Call_eq__2843686287(gopurs_runtime.Int(n_0), gopurs_runtime.Int(0)).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "head").IntVal)})}
goto end_branch_1
} else {

}
}
{
n_0_loop = Call_sub__1043827704(gopurs_runtime.Int(n_0), gopurs_runtime.Int(1)).IntVal
s_1_loop = gopurs_runtime.RecordGet((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_2_0)}.UnsafePtr).V0, "tail").StrVal()
continue codePointAtFallback
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t2)
}
}

func Call_codePointAt(v_0_loop int64, v1_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(v_0), gopurs_runtime.Int(0))).IntVal) != (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (v_0) == (0) {
var __t0 gopurs_runtime.Value
{
if (v1_1) == ("") {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.Apply(Get_unsafeCodePointAt0(), gopurs_runtime.Str(v1_1)).IntVal)})}
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply6(Get__codePointAt(), Get_codePointAtFallback(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, Get_unsafeCodePointAt0(), gopurs_runtime.Int(v_0), gopurs_runtime.Str(v1_1))))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t1)
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}

func Call_bottom__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_top__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_defaultPred__2391565248(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultPred__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultPred__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(Call_sub__1043827704(gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultSucc__2391565248(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultSucc__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_defaultSucc__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Int(Call_add__560788792(gopurs_runtime.Int(gopurs_runtime.Apply(fromEnum_prime_1, a_2).IntVal), gopurs_runtime.Int(1)).IntVal)))
}

func Call_fromEnum__1649438469(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_fromEnum__679972887(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[*pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[*pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_fromEnum__1606852103(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), __eta0_0).IntVal)
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_toEnum__4261336164(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnumWithDefaults__3941305703(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, __eta2_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var __eta2_2 gopurs_runtime.Value = __eta2_2_loop
_ = __eta2_2
v_3_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), gopurs_runtime.Int(__eta2_2.IntVal))
_ = v_3_0
var __t2 gopurs_runtime.Value
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 930809136 && v_3_0.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3_0.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
if (v_3_0.Type == 9 && v_3_0.IntVal == 930809136 && v_3_0.UnsafePtr == nil) {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(__eta2_2.IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), gopurs_runtime.RecordGet(pkg_Data_Bounded.Get_boundedChar(), "bottom")).IntVal))).IntVal) != (0) {
__t1 = __eta0_0
goto end_branch_1
} else {

}
}
{
__t1 = __eta1_1
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
}

func Call_toEnumWithDefaults__3558602759(dictBoundedEnum_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBoundedEnum_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictBoundedEnum_0.V0, gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(dictBoundedEnum_0.V4, gopurs_runtime.Int(x_4.IntVal))
_ = v_5_1
var __t3 gopurs_runtime.Value
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_5_1.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (v_5_1.Type == 9 && v_5_1.IntVal == 930809136 && v_5_1.UnsafePtr == nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(x_4.IntVal), gopurs_runtime.Int(gopurs_runtime.Apply(dictBoundedEnum_0.V3, bottom2_1_0).IntVal))).IntVal) != (0) {
__t2 = low_2
goto end_branch_2
} else {

}
}
{
__t2 = high_3
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
})
})
})
}

func Call_eq__2843686287(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.IntVal) == (__eta1_1.IntVal))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_div__2185172824(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) / (__eta1_1.IntVal))
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_mod__2185172824(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), __eta0_0, __eta1_1)
}

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__291265340(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[int64]) *pkg_Data_Maybe.Constructor_Just[int64] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[int64] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Int(gopurs_runtime.Apply(v_0, gopurs_runtime.Int((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0.IntVal)).IntVal)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0)
}

func Call_map__2116777468(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, string]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[int64, string]](gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)))}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[int64, string]]](__t0)
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_fromJust__2181618881(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_isNothing__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__372254389(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply5(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, __eta0_0, __eta1_1)
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_greaterThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_lessThan__1710332219(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThan__1697837627(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) < (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_lessThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_lessThanOrEq__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_lessThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) * (__eta1_1.IntVal))
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_fst__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_unfoldr__1842498883(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply6(pkg_Data_Unfoldable.Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 930809136 && v_3.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_3.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd(), __eta0_0, __eta1_1)
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

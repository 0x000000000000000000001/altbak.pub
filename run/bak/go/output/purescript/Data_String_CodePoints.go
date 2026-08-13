package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_CodePoints_fromEnum gopurs_runtime.Value
var once_Data_String_CodePoints_fromEnum sync.Once
func Get_Data_String_CodePoints_fromEnum() gopurs_runtime.Value {
	once_Data_String_CodePoints_fromEnum.Do(func() {
		cache_Data_String_CodePoints_fromEnum = gopurs_runtime.RecordGet(Get_Data_Enum_boundedEnumChar(), "fromEnum")
	})
	return cache_Data_String_CodePoints_fromEnum
}

var cache_Data_String_CodePoints_CodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_CodePoint sync.Once
func Get_Data_String_CodePoints_CodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_CodePoint.Do(func() {
		cache_Data_String_CodePoints_CodePoint = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_CodePoints_CodePoint(x_0_box)
})
	})
	return cache_Data_String_CodePoints_CodePoint
}

var cache_Data_String_CodePoints_unsurrogate gopurs_runtime.Value
var once_Data_String_CodePoints_unsurrogate sync.Once
func Get_Data_String_CodePoints_unsurrogate() gopurs_runtime.Value {
	once_Data_String_CodePoints_unsurrogate.Do(func() {
		cache_Data_String_CodePoints_unsurrogate = gopurs_runtime.Func2(func(lead_0_box gopurs_runtime.Value, trail_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_CodePoints_unsurrogate(lead_0_box.IntVal, trail_1_box.IntVal))
})
	})
	return cache_Data_String_CodePoints_unsurrogate
}

var cache_Data_String_CodePoints_uncons gopurs_runtime.Value
var once_Data_String_CodePoints_uncons sync.Once
func Get_Data_String_CodePoints_uncons() gopurs_runtime.Value {
	once_Data_String_CodePoints_uncons.Do(func() {
		cache_Data_String_CodePoints_uncons = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_uncons(s_0_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_uncons
}

var cache_Data_String_CodePoints_unconsButWithTuple gopurs_runtime.Value
var once_Data_String_CodePoints_unconsButWithTuple sync.Once
func Get_Data_String_CodePoints_unconsButWithTuple() gopurs_runtime.Value {
	once_Data_String_CodePoints_unconsButWithTuple.Do(func() {
		cache_Data_String_CodePoints_unconsButWithTuple = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_unconsButWithTuple(s_0_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_unconsButWithTuple
}

var cache_Data_String_CodePoints_toCodePointArrayFallback gopurs_runtime.Value
var once_Data_String_CodePoints_toCodePointArrayFallback sync.Once
func Get_Data_String_CodePoints_toCodePointArrayFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_toCodePointArrayFallback.Do(func() {
		cache_Data_String_CodePoints_toCodePointArrayFallback = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_String_CodePoints_toCodePointArrayFallback(s_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_String_CodePoints_toCodePointArrayFallback
}

var cache_Data_String_CodePoints_showCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_showCodePoint sync.Once
func Get_Data_String_CodePoints_showCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_showCodePoint.Do(func() {
		cache_Data_String_CodePoints_showCodePoint = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(CodePoint 0x") + (gopurs_runtime.Apply(Get_Data_String_Common_toUpper(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_Int_toStringAs(), gopurs_runtime.Int(16), gopurs_runtime.Int(v_0.IntVal)).StrVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_String_CodePoints_showCodePoint
}

var cache_Data_String_CodePoints_isTrail gopurs_runtime.Value
var once_Data_String_CodePoints_isTrail sync.Once
func Get_Data_String_CodePoints_isTrail() gopurs_runtime.Value {
	once_Data_String_CodePoints_isTrail.Do(func() {
		cache_Data_String_CodePoints_isTrail = gopurs_runtime.Func(func(cu_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_String_CodePoints_isTrail(cu_0_box.IntVal))
})
	})
	return cache_Data_String_CodePoints_isTrail
}

var cache_Data_String_CodePoints_isLead gopurs_runtime.Value
var once_Data_String_CodePoints_isLead sync.Once
func Get_Data_String_CodePoints_isLead() gopurs_runtime.Value {
	once_Data_String_CodePoints_isLead.Do(func() {
		cache_Data_String_CodePoints_isLead = gopurs_runtime.Func(func(cu_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_String_CodePoints_isLead(cu_0_box.IntVal))
})
	})
	return cache_Data_String_CodePoints_isLead
}

var cache_Data_String_CodePoints_unsafeCodePointAt0Fallback gopurs_runtime.Value
var once_Data_String_CodePoints_unsafeCodePointAt0Fallback sync.Once
func Get_Data_String_CodePoints_unsafeCodePointAt0Fallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_unsafeCodePointAt0Fallback.Do(func() {
		cache_Data_String_CodePoints_unsafeCodePointAt0Fallback = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_CodePoints_unsafeCodePointAt0Fallback(s_0_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_unsafeCodePointAt0Fallback
}

var cache_Data_String_CodePoints_unsafeCodePointAt0 gopurs_runtime.Value
var once_Data_String_CodePoints_unsafeCodePointAt0 sync.Once
func Get_Data_String_CodePoints_unsafeCodePointAt0() gopurs_runtime.Value {
	once_Data_String_CodePoints_unsafeCodePointAt0.Do(func() {
		cache_Data_String_CodePoints_unsafeCodePointAt0 = gopurs_runtime.Apply(Get_Data_String_CodePoints__unsafeCodePointAt0(), Get_Data_String_CodePoints_unsafeCodePointAt0Fallback())
	})
	return cache_Data_String_CodePoints_unsafeCodePointAt0
}

var cache_Data_String_CodePoints_toCodePointArray gopurs_runtime.Value
var once_Data_String_CodePoints_toCodePointArray sync.Once
func Get_Data_String_CodePoints_toCodePointArray() gopurs_runtime.Value {
	once_Data_String_CodePoints_toCodePointArray.Do(func() {
		cache_Data_String_CodePoints_toCodePointArray = gopurs_runtime.Apply2(Get_Data_String_CodePoints__toCodePointArray(), Get_Data_String_CodePoints_toCodePointArrayFallback(), Get_Data_String_CodePoints_unsafeCodePointAt0())
	})
	return cache_Data_String_CodePoints_toCodePointArray
}

var cache_Data_String_CodePoints_length gopurs_runtime.Value
var once_Data_String_CodePoints_length sync.Once
func Get_Data_String_CodePoints_length() gopurs_runtime.Value {
	once_Data_String_CodePoints_length.Do(func() {
		cache_Data_String_CodePoints_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_CodePoints_length(x_0_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_length
}

var cache_Data_String_CodePoints_lastIndexOf gopurs_runtime.Value
var once_Data_String_CodePoints_lastIndexOf sync.Once
func Get_Data_String_CodePoints_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_CodePoints_lastIndexOf.Do(func() {
		cache_Data_String_CodePoints_lastIndexOf = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_lastIndexOf(p_0_box.StrVal(), s_1_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_lastIndexOf
}

var cache_Data_String_CodePoints_indexOf gopurs_runtime.Value
var once_Data_String_CodePoints_indexOf sync.Once
func Get_Data_String_CodePoints_indexOf() gopurs_runtime.Value {
	once_Data_String_CodePoints_indexOf.Do(func() {
		cache_Data_String_CodePoints_indexOf = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_indexOf(p_0_box.StrVal(), s_1_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_indexOf
}

var cache_Data_String_CodePoints_fromCharCode gopurs_runtime.Value
var once_Data_String_CodePoints_fromCharCode sync.Once
func Get_Data_String_CodePoints_fromCharCode() gopurs_runtime.Value {
	once_Data_String_CodePoints_fromCharCode.Do(func() {
		cache_Data_String_CodePoints_fromCharCode = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_CodePoints_fromCharCode(x_0_box.IntVal))
})
	})
	return cache_Data_String_CodePoints_fromCharCode
}

var cache_Data_String_CodePoints_singletonFallback gopurs_runtime.Value
var once_Data_String_CodePoints_singletonFallback sync.Once
func Get_Data_String_CodePoints_singletonFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_singletonFallback.Do(func() {
		cache_Data_String_CodePoints_singletonFallback = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_CodePoints_singletonFallback(v_0_box.IntVal))
})
	})
	return cache_Data_String_CodePoints_singletonFallback
}

var cache_Data_String_CodePoints_fromCodePointArray gopurs_runtime.Value
var once_Data_String_CodePoints_fromCodePointArray sync.Once
func Get_Data_String_CodePoints_fromCodePointArray() gopurs_runtime.Value {
	once_Data_String_CodePoints_fromCodePointArray.Do(func() {
		cache_Data_String_CodePoints_fromCodePointArray = gopurs_runtime.Apply(Get_Data_String_CodePoints__fromCodePointArray(), Get_Data_String_CodePoints_singletonFallback())
	})
	return cache_Data_String_CodePoints_fromCodePointArray
}

var cache_Data_String_CodePoints_singleton gopurs_runtime.Value
var once_Data_String_CodePoints_singleton sync.Once
func Get_Data_String_CodePoints_singleton() gopurs_runtime.Value {
	once_Data_String_CodePoints_singleton.Do(func() {
		cache_Data_String_CodePoints_singleton = gopurs_runtime.Apply(Get_Data_String_CodePoints__singleton(), Get_Data_String_CodePoints_singletonFallback())
	})
	return cache_Data_String_CodePoints_singleton
}

var cache_Data_String_CodePoints_takeFallback gopurs_runtime.Value
var once_Data_String_CodePoints_takeFallback sync.Once
func Get_Data_String_CodePoints_takeFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_takeFallback.Do(func() {
		cache_Data_String_CodePoints_takeFallback = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_CodePoints_takeFallback(v_0_box.IntVal, v1_1_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_takeFallback
}

var cache_Data_String_CodePoints_take gopurs_runtime.Value
var once_Data_String_CodePoints_take sync.Once
func Get_Data_String_CodePoints_take() gopurs_runtime.Value {
	once_Data_String_CodePoints_take.Do(func() {
		cache_Data_String_CodePoints_take = gopurs_runtime.Apply(Get_Data_String_CodePoints__take(), Get_Data_String_CodePoints_takeFallback())
	})
	return cache_Data_String_CodePoints_take
}

var cache_Data_String_CodePoints_lastIndexOf_prime gopurs_runtime.Value
var once_Data_String_CodePoints_lastIndexOf_prime sync.Once
func Get_Data_String_CodePoints_lastIndexOf_prime() gopurs_runtime.Value {
	once_Data_String_CodePoints_lastIndexOf_prime.Do(func() {
		cache_Data_String_CodePoints_lastIndexOf_prime = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_lastIndexOf_prime(p_0_box.StrVal(), i_1_box.IntVal, s_2_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_lastIndexOf_prime
}

var cache_Data_String_CodePoints_splitAt gopurs_runtime.Value
var once_Data_String_CodePoints_splitAt sync.Once
func Get_Data_String_CodePoints_splitAt() gopurs_runtime.Value {
	once_Data_String_CodePoints_splitAt.Do(func() {
		cache_Data_String_CodePoints_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_CodePoints_splitAt(i_0_box.IntVal, s_1_box.StrVal())
})
	})
	return cache_Data_String_CodePoints_splitAt
}

var cache_Data_String_CodePoints_eqCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_eqCodePoint sync.Once
func Get_Data_String_CodePoints_eqCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_eqCodePoint.Do(func() {
		cache_Data_String_CodePoints_eqCodePoint = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.IntVal) == (y_1.IntVal))
})
}))
	})
	return cache_Data_String_CodePoints_eqCodePoint
}

var cache_Data_String_CodePoints_ordCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_ordCodePoint sync.Once
func Get_Data_String_CodePoints_ordCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_ordCodePoint.Do(func() {
		cache_Data_String_CodePoints_ordCodePoint = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_String_CodePoints_eqCodePoint()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(x_0.IntVal), gopurs_runtime.Int(y_1.IntVal)).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_Data_String_CodePoints_ordCodePoint
}

var cache_Data_String_CodePoints_drop gopurs_runtime.Value
var once_Data_String_CodePoints_drop sync.Once
func Get_Data_String_CodePoints_drop() gopurs_runtime.Value {
	once_Data_String_CodePoints_drop.Do(func() {
		cache_Data_String_CodePoints_drop = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_CodePoints_drop(n_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_drop
}

var cache_Data_String_CodePoints_indexOf_prime gopurs_runtime.Value
var once_Data_String_CodePoints_indexOf_prime sync.Once
func Get_Data_String_CodePoints_indexOf_prime() gopurs_runtime.Value {
	once_Data_String_CodePoints_indexOf_prime.Do(func() {
		cache_Data_String_CodePoints_indexOf_prime = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, i_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_indexOf_prime(p_0_box.StrVal(), i_1_box.IntVal, s_2_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_indexOf_prime
}

var cache_Data_String_CodePoints_countTail gopurs_runtime.Value
var once_Data_String_CodePoints_countTail sync.Once
func Get_Data_String_CodePoints_countTail() gopurs_runtime.Value {
	once_Data_String_CodePoints_countTail.Do(func() {
		cache_Data_String_CodePoints_countTail = gopurs_runtime.Func3(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value, accum_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_CodePoints_countTail(p_0_box, s_1_box.StrVal(), accum_2_box.IntVal))
})
	})
	return cache_Data_String_CodePoints_countTail
}

var cache_Data_String_CodePoints_countFallback gopurs_runtime.Value
var once_Data_String_CodePoints_countFallback sync.Once
func Get_Data_String_CodePoints_countFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_countFallback.Do(func() {
		cache_Data_String_CodePoints_countFallback = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_CodePoints_countFallback(p_0_box, s_1_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_countFallback
}

var cache_Data_String_CodePoints_countPrefix gopurs_runtime.Value
var once_Data_String_CodePoints_countPrefix sync.Once
func Get_Data_String_CodePoints_countPrefix() gopurs_runtime.Value {
	once_Data_String_CodePoints_countPrefix.Do(func() {
		cache_Data_String_CodePoints_countPrefix = gopurs_runtime.Apply2(Get_Data_String_CodePoints__countPrefix(), Get_Data_String_CodePoints_countFallback(), Get_Data_String_CodePoints_unsafeCodePointAt0())
	})
	return cache_Data_String_CodePoints_countPrefix
}

var cache_Data_String_CodePoints_dropWhile gopurs_runtime.Value
var once_Data_String_CodePoints_dropWhile sync.Once
func Get_Data_String_CodePoints_dropWhile() gopurs_runtime.Value {
	once_Data_String_CodePoints_dropWhile.Do(func() {
		cache_Data_String_CodePoints_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_CodePoints_dropWhile(p_0_box, s_1_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_dropWhile
}

var cache_Data_String_CodePoints_takeWhile gopurs_runtime.Value
var once_Data_String_CodePoints_takeWhile sync.Once
func Get_Data_String_CodePoints_takeWhile() gopurs_runtime.Value {
	once_Data_String_CodePoints_takeWhile.Do(func() {
		cache_Data_String_CodePoints_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_CodePoints_takeWhile(p_0_box, s_1_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_takeWhile
}

var cache_Data_String_CodePoints_codePointFromChar gopurs_runtime.Value
var once_Data_String_CodePoints_codePointFromChar sync.Once
func Get_Data_String_CodePoints_codePointFromChar() gopurs_runtime.Value {
	once_Data_String_CodePoints_codePointFromChar.Do(func() {
		cache_Data_String_CodePoints_codePointFromChar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_CodePoints_codePointFromChar(x_0_box.StrVal()))
})
	})
	return cache_Data_String_CodePoints_codePointFromChar
}

var cache_Data_String_CodePoints_codePointAtFallback gopurs_runtime.Value
var once_Data_String_CodePoints_codePointAtFallback sync.Once
func Get_Data_String_CodePoints_codePointAtFallback() gopurs_runtime.Value {
	once_Data_String_CodePoints_codePointAtFallback.Do(func() {
		cache_Data_String_CodePoints_codePointAtFallback = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_codePointAtFallback(n_0_box.IntVal, s_1_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_codePointAtFallback
}

var cache_Data_String_CodePoints_codePointAt gopurs_runtime.Value
var once_Data_String_CodePoints_codePointAt sync.Once
func Get_Data_String_CodePoints_codePointAt() gopurs_runtime.Value {
	once_Data_String_CodePoints_codePointAt.Do(func() {
		cache_Data_String_CodePoints_codePointAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_CodePoints_codePointAt(v_0_box.IntVal, v1_1_box.StrVal()))}
})
	})
	return cache_Data_String_CodePoints_codePointAt
}

var cache_Data_String_CodePoints_boundedCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_boundedCodePoint sync.Once
func Get_Data_String_CodePoints_boundedCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_boundedCodePoint.Do(func() {
		cache_Data_String_CodePoints_boundedCodePoint = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_String_CodePoints_ordCodePoint()
}), gopurs_runtime.Int(0), gopurs_runtime.Int(1114111))
	})
	return cache_Data_String_CodePoints_boundedCodePoint
}

var cache_Data_String_CodePoints_boundedEnumCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_boundedEnumCodePoint sync.Once
func Get_Data_String_CodePoints_boundedEnumCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_boundedEnumCodePoint.Do(func() {
		cache_Data_String_CodePoints_boundedEnumCodePoint = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_String_CodePoints_boundedCodePoint()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_String_CodePoints_enumCodePoint()
}), gopurs_runtime.Int(1114112), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
var __t0 bool
{
if (n_0.IntVal) < (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t_and_2 bool = false
if __t0 {

var __t1 bool
{
if (n_0.IntVal) > (1114111) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
__t_and_2 = __t1
}
if __t_and_2 {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(n_0.IntVal)})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t3))}
}))
	})
	return cache_Data_String_CodePoints_boundedEnumCodePoint
}

var cache_Data_String_CodePoints_enumCodePoint gopurs_runtime.Value
var once_Data_String_CodePoints_enumCodePoint sync.Once
func Get_Data_String_CodePoints_enumCodePoint() gopurs_runtime.Value {
	once_Data_String_CodePoints_enumCodePoint.Do(func() {
		cache_Data_String_CodePoints_enumCodePoint = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_String_CodePoints_ordCodePoint()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_String_CodePoints_boundedEnumCodePoint(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_String_CodePoints_boundedEnumCodePoint(), "fromEnum"), gopurs_runtime.Int(a_0.IntVal)).IntVal) - (1)))))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_String_CodePoints_boundedEnumCodePoint(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_String_CodePoints_boundedEnumCodePoint(), "fromEnum"), gopurs_runtime.Int(a_0.IntVal)).IntVal) + (1)))))}
}))
	})
	return cache_Data_String_CodePoints_enumCodePoint
}

func Call_Data_String_CodePoints_CodePoint(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_CodePoints_unsurrogate(lead_0_loop int64, trail_1_loop int64) int64 {
var lead_0 int64 = lead_0_loop
_ = lead_0
var trail_1 int64 = trail_1_loop
_ = trail_1
return ((((lead_0) - (55296)) * (1024)) + ((trail_1) - (56320))) + (65536)
}

func Call_Data_String_CodePoints_uncons(s_0_loop string) *Constructor_Data_Maybe_Just {
var s_0 string = s_0_loop
_ = s_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(Get_Data_String_Unsafe_charAt(), gopurs_runtime.Int(0))
_ = __local_var_1_2
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Enum_boundedEnumChar(), "fromEnum"), gopurs_runtime.Apply(__local_var_1_2, x_2))
})
_ = __local_var_1_1
// TAST (Let): v_1_0 -> gopurs_runtime.Value
v_1_0 := gopurs_runtime.Apply6(Get_Data_String_CodePoints__codePointAt(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
})
}), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, x_2)
}), gopurs_runtime.Int(0), gopurs_runtime.Str(s_0))
_ = v_1_0
var __t3 gopurs_runtime.Value
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_3
} else {

}
}
{
if (v_1_0.Type == 9 && v_1_0.IntVal == 930809136 && v_1_0.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int((*Constructor_Data_Maybe_Just)(v_1_0.UnsafePtr).V0.IntVal), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply3(Get_Data_String_CodePoints__take(), gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t3)
}

func Call_Data_String_CodePoints_unconsButWithTuple(s_0_loop string) *Constructor_Data_Maybe_Just {
var s_0 string = s_0_loop
_ = s_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := Call_Data_String_CodePoints_uncons(s_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((__local_var_1_0).V0, "head").IntVal), gopurs_runtime.Str(gopurs_runtime.RecordGet((__local_var_1_0).V0, "tail").StrVal())})}})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_CodePoints_toCodePointArrayFallback(s_0_loop string) []int64 {
var s_0 string = s_0_loop
_ = s_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Unfoldable_unfoldableArray(), "unfoldr"), Get_Data_String_CodePoints_unconsButWithTuple(), gopurs_runtime.Str(s_0)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_Data_String_CodePoints_isTrail(cu_0_loop int64) bool {
var cu_0 int64 = cu_0_loop
_ = cu_0
var __t0 bool
{
if (56320) > (cu_0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t_and_2 bool = false
if __t0 {

var __t1 bool
{
if (cu_0) > (57343) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
__t_and_2 = __t1
}
return __t_and_2
}

func Call_Data_String_CodePoints_isLead(cu_0_loop int64) bool {
var cu_0 int64 = cu_0_loop
_ = cu_0
var __t0 bool
{
if (55296) > (cu_0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
var __t_and_2 bool = false
if __t0 {

var __t1 bool
{
if (cu_0) > (56319) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
__t_and_2 = __t1
}
return __t_and_2
}

func Call_Data_String_CodePoints_unsafeCodePointAt0Fallback(s_0_loop string) int64 {
var s_0 string = s_0_loop
_ = s_0
// TAST (Let): cu0_1_0 -> int64
cu0_1_0 := gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_Unsafe_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(s_0)).StrVal())).IntVal
_ = cu0_1_0
var __t11 int64
{
var __t1 bool
{
if (55296) > (cu0_1_0) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
var __t_and_3 bool = false
if __t1 {

var __t2 bool
{
if (cu0_1_0) > (56319) {
__t2 = false
goto end_branch_2
} else {

}
}
{
__t2 = true
}
end_branch_2:
__t_and_3 = __t2
}
var __t_and_5 bool = false
if __t_and_3 {

var __t4 bool
{
if (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(s_0)).IntVal) > (1) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t_and_5 = __t4
}
if __t_and_5 {
// TAST (Let): cu1_2_6 -> int64
cu1_2_6 := gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_Unsafe_charAt(), gopurs_runtime.Int(1), gopurs_runtime.Str(s_0)).StrVal())).IntVal
_ = cu1_2_6
var __t10 int64
{
var __t7 bool
{
if (56320) > (cu1_2_6) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = true
}
end_branch_7:
var __t_and_9 bool = false
if __t7 {

var __t8 bool
{
if (cu1_2_6) > (57343) {
__t8 = false
goto end_branch_8
} else {

}
}
{
__t8 = true
}
end_branch_8:
__t_and_9 = __t8
}
if __t_and_9 {
__t10 = ((((cu0_1_0) - (55296)) * (1024)) + ((cu1_2_6) - (56320))) + (65536)
goto end_branch_10
} else {

}
}
{
__t10 = cu0_1_0
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
__t11 = cu0_1_0
}
end_branch_11:
return __t11
}

func Call_Data_String_CodePoints_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(x_0))))).IntVal
}

func Call_Data_String_CodePoints_lastIndexOf(p_0_loop string, s_1_loop string) *Constructor_Data_Maybe_Just {
var p_0 string = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_String_CodeUnits_lastIndexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1)))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((__local_var_2_0).V0.IntVal), gopurs_runtime.Str(s_1)).StrVal()))))).IntVal)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_CodePoints_indexOf(p_0_loop string, s_1_loop string) *Constructor_Data_Maybe_Just {
var p_0 string = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
// TAST (Let): __local_var_2_0 -> *Constructor_Data_Maybe_Just
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_1)))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((__local_var_2_0).V0.IntVal), gopurs_runtime.Str(s_1)).StrVal()))))).IntVal)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_CodePoints_fromCharCode(x_0_loop int64) string {
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Apply3(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "top").StrVal()), gopurs_runtime.Int(x_0))).StrVal()
}

func Call_Data_String_CodePoints_singletonFallback(v_0_loop int64) string {
var v_0 int64 = v_0_loop
_ = v_0
var __t1 string
{
var __t0 bool
{
if (v_0) > (65535) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Apply3(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "top").StrVal()), gopurs_runtime.Int(v_0))).StrVal()
goto end_branch_1
} else {

}
}
{
__t1 = (gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Apply3(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "top").StrVal()), gopurs_runtime.Int((((v_0) - (65536)) / (1024)) + (55296)))).StrVal()) + (gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Apply3(Get_Data_Enum_toEnumWithDefaults__3941305703(), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "bottom").StrVal()), gopurs_runtime.Str(gopurs_runtime.RecordGet(Get_Data_Bounded_boundedChar(), "top").StrVal()), gopurs_runtime.Int((gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int((v_0) - (65536)), gopurs_runtime.Int(1024)).IntVal) + (56320)))).StrVal())
}
end_branch_1:
return __t1
}

func Call_Data_String_CodePoints_takeFallback(v_0_loop int64, v1_1_loop string) string {
takeFallback:
for {
if false { continue takeFallback }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
var __t3 string
{
var __t2 bool
{
if (v_0) < (1) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
if __t2 {
__t3 = ""
goto end_branch_3
} else {

}
}
{
// TAST (Let): v2_2_0 -> *Constructor_Data_Maybe_Just
v2_2_0 := Call_Data_String_CodePoints_uncons(v1_1)
_ = v2_2_0
var __t1 string
{
if (v2_2_0 != nil) {
__t1 = (gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(gopurs_runtime.RecordGet((v2_2_0).V0, "head").IntVal)).StrVal()) + (Call_Data_String_CodePoints_takeFallback((v_0) - (1), gopurs_runtime.RecordGet((v2_2_0).V0, "tail").StrVal()))
goto end_branch_1
} else {

}
}
{
__t1 = v1_1
}
end_branch_1:
__t3 = __t1
}
end_branch_3:
return __t3
}
}

func Call_Data_String_CodePoints_lastIndexOf_prime(p_0_loop string, i_1_loop int64, s_2_loop string) *Constructor_Data_Maybe_Just {
var p_0 string = p_0_loop
_ = p_0
var i_1 int64 = i_1_loop
_ = i_1
var s_2 string = s_2_loop
_ = s_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_String_CodeUnits_lastIndexOf_prime(), gopurs_runtime.Str(p_0), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2)).StrVal())).IntVal), gopurs_runtime.Str(s_2)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((__local_var_3_0).V0.IntVal), gopurs_runtime.Str(s_2)).StrVal()))))).IntVal)})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_CodePoints_splitAt(i_0_loop int64, s_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var s_1 string = s_1_loop
_ = s_1
// TAST (Let): before_2_0 -> gopurs_runtime.Value
before_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(s_1))
_ = before_2_0
return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(before_2_0.StrVal())).IntVal), gopurs_runtime.Str(s_1)).StrVal()), gopurs_runtime.Str(before_2_0.StrVal()))
}

func Call_Data_String_CodePoints_drop(n_0_loop int64, s_1_loop string) string {
var n_0 int64 = n_0_loop
_ = n_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(n_0), gopurs_runtime.Str(s_1)).StrVal())).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodePoints_indexOf_prime(p_0_loop string, i_1_loop int64, s_2_loop string) *Constructor_Data_Maybe_Just {
var p_0 string = p_0_loop
_ = p_0
var i_1 int64 = i_1_loop
_ = i_1
var s_2 string = s_2_loop
_ = s_2
// TAST (Let): s_prime_3_0 -> string
s_prime_3_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_1), gopurs_runtime.Str(s_2)).StrVal())).IntVal), gopurs_runtime.Str(s_2)).StrVal()
_ = s_prime_3_0
// TAST (Let): __local_var_4_1 -> *Constructor_Data_Maybe_Just
__local_var_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str(p_0), gopurs_runtime.Str(s_prime_3_0)))
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1 != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_1) + (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((__local_var_4_1).V0.IntVal), gopurs_runtime.Str(s_prime_3_0)).StrVal()))))).IntVal))})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2)
}

func Call_Data_String_CodePoints_countTail(p_0_loop gopurs_runtime.Value, s_1_loop string, accum_2_loop int64) int64 {
countTail:
for {
if false { continue countTail }
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
var accum_2 int64 = accum_2_loop
_ = accum_2
// TAST (Let): v_3_0 -> *Constructor_Data_Maybe_Just
v_3_0 := Call_Data_String_CodePoints_uncons(s_1)
_ = v_3_0
var __t2 int64
{
if (v_3_0 != nil) {
var __t1 int64
{
if (gopurs_runtime.Apply(p_0, gopurs_runtime.Int(gopurs_runtime.RecordGet((v_3_0).V0, "head").IntVal)).IntVal) != (0) {
p_0_loop = p_0
s_1_loop = gopurs_runtime.RecordGet((v_3_0).V0, "tail").StrVal()
accum_2_loop = (accum_2) + (1)
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

func Call_Data_String_CodePoints_countFallback(p_0_loop gopurs_runtime.Value, s_1_loop string) int64 {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return Call_Data_String_CodePoints_countTail(p_0, s_1, 0)
}

func Call_Data_String_CodePoints_dropWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodePoints_countPrefix(), p_0, gopurs_runtime.Str(s_1)).IntVal), gopurs_runtime.Str(s_1)).StrVal())).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodePoints_takeWhile(p_0_loop gopurs_runtime.Value, s_1_loop string) string {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Data_String_CodePoints_countPrefix(), p_0, gopurs_runtime.Str(s_1)).IntVal), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_Data_String_CodePoints_codePointFromChar(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Enum_boundedEnumChar(), "fromEnum"), gopurs_runtime.Str(x_0)).IntVal
}

func Call_Data_String_CodePoints_codePointAtFallback(n_0_loop int64, s_1_loop string) *Constructor_Data_Maybe_Just {
codePointAtFallback:
for {
if false { continue codePointAtFallback }
var n_0 int64 = n_0_loop
_ = n_0
var s_1 string = s_1_loop
_ = s_1
// TAST (Let): v_2_0 -> *Constructor_Data_Maybe_Just
v_2_0 := Call_Data_String_CodePoints_uncons(s_1)
_ = v_2_0
var __t2 gopurs_runtime.Value
{
if (v_2_0 != nil) {
var __t1 gopurs_runtime.Value
{
if (n_0) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.RecordGet((v_2_0).V0, "head").IntVal)})}
goto end_branch_1
} else {

}
}
{
n_0_loop = (n_0) - (1)
s_1_loop = gopurs_runtime.RecordGet((v_2_0).V0, "tail").StrVal()
continue codePointAtFallback
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{}))}
}
end_branch_1:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2)
}
}

func Call_Data_String_CodePoints_codePointAt(v_0_loop int64, v1_1_loop string) *Constructor_Data_Maybe_Just {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
var __t2 gopurs_runtime.Value
{
var __t0 bool
{
if (v_0) < (0) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_2
} else {

}
}
{
if (v_0) == (0) {
var __t1 gopurs_runtime.Value
{
if (v1_1) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodePoints_unsafeCodePointAt0(), gopurs_runtime.Str(v1_1)).IntVal)})}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply6(Get_Data_String_CodePoints__codePointAt(), Get_Data_String_CodePoints_codePointAtFallback(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, Get_Data_String_CodePoints_unsafeCodePointAt0(), gopurs_runtime.Int(v_0), gopurs_runtime.Str(v1_1))))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2)
}

func Get_Data_String_CodePoints__codePointAt() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__CodePointAt
}

func Get_Data_String_CodePoints__countPrefix() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__CountPrefix
}

func Get_Data_String_CodePoints__fromCodePointArray() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__FromCodePointArray
}

func Get_Data_String_CodePoints__singleton() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__Singleton
}

func Get_Data_String_CodePoints__take() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__Take
}

func Get_Data_String_CodePoints__toCodePointArray() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__ToCodePointArray
}

func Get_Data_String_CodePoints__unsafeCodePointAt0() gopurs_runtime.Value {
	return _Gopurs_Data_String_CodePoints__UnsafeCodePointAt0
}

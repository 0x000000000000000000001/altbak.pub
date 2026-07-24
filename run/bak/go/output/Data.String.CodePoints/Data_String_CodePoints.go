package Data_String_CodePoints

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Bounded "gopurs/output/Data.Bounded"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
)

var showCodePoint gopurs_runtime.Value
var once_showCodePoint sync.Once
func Get_showCodePoint() gopurs_runtime.Value {
	once_showCodePoint.Do(func() {
		showCodePoint = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("(CodePoint 0x").StrVal + gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), gopurs_runtime.Apply2(pkg_Data_Int.Get_toStringAs(), gopurs_runtime.Int(16), v_0)).StrVal).StrVal + gopurs_runtime.Str(")").StrVal)
}))
	})
	return showCodePoint
}

var uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		uncons = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
v_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), s_0)
_ = v_1_0
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t4 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_4
} else {

}
}
{
if (gopurs_runtime.Bool(v_1_0.IntVal == gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t4 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), s_0)), gopurs_runtime.Str("")))
goto end_branch_4
} else {

}
}
{
cu1_2_1 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(1), s_0))
_ = cu1_2_1
cu0_3_2 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), s_0))
_ = cu0_3_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Int(55296).IntVal <= cu0_3_2.IntVal).IntVal != 0 && gopurs_runtime.Bool(cu0_3_2.IntVal <= gopurs_runtime.Int(56319).IntVal).IntVal != 0).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Int(56320).IntVal <= cu1_2_1.IntVal).IntVal != 0 && gopurs_runtime.Bool(cu1_2_1.IntVal <= gopurs_runtime.Int(57343).IntVal).IntVal != 0).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(cu0_3_2.IntVal - gopurs_runtime.Int(55296).IntVal * gopurs_runtime.Int(1024).IntVal + cu1_2_1.IntVal - gopurs_runtime.Int(56320).IntVal + gopurs_runtime.Int(65536).IntVal), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(2), s_0)))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordDict2("head", "tail", cu0_3_2, gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(1), s_0)))
}
end_branch_3:
__t4 = __t3
}
end_branch_4:
return __t4
})
	})
	return uncons
}

var unconsButWithTuple gopurs_runtime.Value
var once_unconsButWithTuple sync.Once
func Get_unconsButWithTuple() gopurs_runtime.Value {
	once_unconsButWithTuple.Do(func() {
		unconsButWithTuple = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_uncons(), s_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_1_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Constructor2("Tuple", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)[0], "head"), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(__local_var_1_0.UnsafePtr)[0], "tail")))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
	})
	return unconsButWithTuple
}

var toCodePointArrayFallback gopurs_runtime.Value
var once_toCodePointArrayFallback sync.Once
func Get_toCodePointArrayFallback() gopurs_runtime.Value {
	once_toCodePointArrayFallback.Do(func() {
		toCodePointArrayFallback = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr"), Get_unconsButWithTuple(), s_0)
})
	})
	return toCodePointArrayFallback
}

var unsafeCodePointAt0Fallback gopurs_runtime.Value
var once_unsafeCodePointAt0Fallback sync.Once
func Get_unsafeCodePointAt0Fallback() gopurs_runtime.Value {
	once_unsafeCodePointAt0Fallback.Do(func() {
		unsafeCodePointAt0Fallback = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
cu0_1_0 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), s_0))
_ = cu0_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Int(55296).IntVal <= cu0_1_0.IntVal).IntVal != 0 && gopurs_runtime.Bool(cu0_1_0.IntVal <= gopurs_runtime.Int(56319).IntVal).IntVal != 0).IntVal != 0 && gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), s_0).IntVal > gopurs_runtime.Int(1).IntVal).IntVal != 0)).IntVal != 0 {
cu1_2_2 := gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(1), s_0))
_ = cu1_2_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(gopurs_runtime.Int(56320).IntVal <= cu1_2_2.IntVal).IntVal != 0 && gopurs_runtime.Bool(cu1_2_2.IntVal <= gopurs_runtime.Int(57343).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Int(cu0_1_0.IntVal - gopurs_runtime.Int(55296).IntVal * gopurs_runtime.Int(1024).IntVal + cu1_2_2.IntVal - gopurs_runtime.Int(56320).IntVal + gopurs_runtime.Int(65536).IntVal)
goto end_branch_3
} else {

}
}
{
__t3 = cu0_1_0
}
end_branch_3:
__t1 = __t3
goto end_branch_1
} else {

}
}
{
__t1 = cu0_1_0
}
end_branch_1:
return __t1
})
	})
	return unsafeCodePointAt0Fallback
}

var unsafeCodePointAt0 gopurs_runtime.Value
var once_unsafeCodePointAt0 sync.Once
func Get_unsafeCodePointAt0() gopurs_runtime.Value {
	once_unsafeCodePointAt0.Do(func() {
		unsafeCodePointAt0 = gopurs_runtime.Apply(Get__unsafeCodePointAt0(), Get_unsafeCodePointAt0Fallback())
	})
	return unsafeCodePointAt0
}

var toCodePointArray gopurs_runtime.Value
var once_toCodePointArray sync.Once
func Get_toCodePointArray() gopurs_runtime.Value {
	once_toCodePointArray.Do(func() {
		toCodePointArray = gopurs_runtime.Apply2(Get__toCodePointArray(), Get_toCodePointArrayFallback(), Get_unsafeCodePointAt0())
	})
	return toCodePointArray
}

var length gopurs_runtime.Value
var once_length sync.Once
func Get_length() gopurs_runtime.Value {
	once_length.Do(func() {
		length = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(int64(len(gopurs_runtime.Apply(Get_toCodePointArray(), x_0).PtrVal.([]gopurs_runtime.Value))))
})
	})
	return length
}

var lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		lastIndexOf = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_lastIndexOf(), p_0, s_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(int64(len(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0], s_1)).PtrVal.([]gopurs_runtime.Value)))))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
	})
	return lastIndexOf
}

var indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		indexOf = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_indexOf(), p_0, s_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_2_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(int64(len(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0], s_1)).PtrVal.([]gopurs_runtime.Value)))))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
	})
	return indexOf
}

var fromCharCode gopurs_runtime.Value
var once_fromCharCode sync.Once
func Get_fromCharCode() gopurs_runtime.Value {
	once_fromCharCode.Do(func() {
		fromCharCode = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
v_1_0 := gopurs_runtime.Apply(pkg_Data_Enum.Get_charToEnum(), x_0)
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_1_0.StrVal == "Just")).IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(v_1_0.UnsafePtr)[0]
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Bool(v_1_0.StrVal == "Nothing")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(x_0.IntVal < gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), pkg_Data_Bounded.Get_bottomChar()).IntVal)).IntVal != 0 {
__t2 = pkg_Data_Bounded.Get_bottomChar()
goto end_branch_2
} else {

}
}
{
__t2 = pkg_Data_Bounded.Get_topChar()
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
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), __t1)
})
	})
	return fromCharCode
}

var singletonFallback gopurs_runtime.Value
var once_singletonFallback sync.Once
func Get_singletonFallback() gopurs_runtime.Value {
	once_singletonFallback.Do(func() {
		singletonFallback = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal <= gopurs_runtime.Int(65535).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Apply(Get_fromCharCode(), v_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(65536).IntVal / gopurs_runtime.Int(1024).IntVal + gopurs_runtime.Int(55296).IntVal)).StrVal + gopurs_runtime.Apply(Get_fromCharCode(), gopurs_runtime.Int(gopurs_runtime.Apply2(pkg_Data_EuclideanRing.Get_intMod(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(65536).IntVal), gopurs_runtime.Int(1024)).IntVal + gopurs_runtime.Int(56320).IntVal)).StrVal)
}
end_branch_0:
return __t0
})
	})
	return singletonFallback
}

var fromCodePointArray gopurs_runtime.Value
var once_fromCodePointArray sync.Once
func Get_fromCodePointArray() gopurs_runtime.Value {
	once_fromCodePointArray.Do(func() {
		fromCodePointArray = gopurs_runtime.Apply(Get__fromCodePointArray(), Get_singletonFallback())
	})
	return fromCodePointArray
}

var singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		singleton = gopurs_runtime.Apply(Get__singleton(), Get_singletonFallback())
	})
	return singleton
}

var takeFallback gopurs_runtime.Value
var once_takeFallback sync.Once
func Get_takeFallback() gopurs_runtime.Value {
	once_takeFallback.Do(func() {
		takeFallback = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
takeFallback:
for {
if false { continue takeFallback }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal < gopurs_runtime.Int(1).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Str("")
goto end_branch_2
} else {

}
}
{
v2_2_0 := gopurs_runtime.Apply(Get_uncons(), v1_1)
_ = v2_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_2_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Str(gopurs_runtime.Apply(Get_singleton(), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v2_2_0.UnsafePtr)[0], "head")).StrVal + gopurs_runtime.Apply2(Get_takeFallback(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v2_2_0.UnsafePtr)[0], "tail")).StrVal)
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
}()
})
})
	})
	return takeFallback
}

var take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		take = gopurs_runtime.Apply(Get__take(), Get_takeFallback())
	})
	return take
}

var lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		lastIndexOf_prime = gopurs_runtime.Func3(func(p_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply3(pkg_Data_String_CodeUnits.Get_lastIndexOf_prime(), p_0, gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), i_1, s_2)), s_2)
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_3_0.StrVal == "Just")).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(int64(len(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), (*[1024]gopurs_runtime.Value)(__local_var_3_0.UnsafePtr)[0], s_2)).PtrVal.([]gopurs_runtime.Value)))))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
})
	})
	return lastIndexOf_prime
}

var splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		splitAt = gopurs_runtime.Func2(func(i_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
before_2_0 := gopurs_runtime.Apply2(Get_take(), i_0, s_1)
_ = before_2_0
return gopurs_runtime.RecordDict2("before", "after", before_2_0, gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), before_2_0), s_1))
})
	})
	return splitAt
}

var eqCodePoint gopurs_runtime.Value
var once_eqCodePoint sync.Once
func Get_eqCodePoint() gopurs_runtime.Value {
	once_eqCodePoint.Do(func() {
		eqCodePoint = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(x_0.IntVal == y_1.IntVal)
}))
	})
	return eqCodePoint
}

var ordCodePoint gopurs_runtime.Value
var once_ordCodePoint sync.Once
func Get_ordCodePoint() gopurs_runtime.Value {
	once_ordCodePoint.Do(func() {
		ordCodePoint = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordInt(), "compare"), x_0, y_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCodePoint()
}))
	})
	return ordCodePoint
}

var drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		drop = gopurs_runtime.Func2(func(n_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), n_0, s_1)), s_1)
})
	})
	return drop
}

var indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		indexOf_prime = gopurs_runtime.Func3(func(p_0 gopurs_runtime.Value, i_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
s_prime_3_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), i_1, s_2)), s_2)
_ = s_prime_3_0
__local_var_4_1 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_indexOf(), p_0, s_prime_3_0)
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_4_1.StrVal == "Just")).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Int(i_1.IntVal + gopurs_runtime.Int(int64(len(gopurs_runtime.Apply(Get_toCodePointArray(), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), (*[1024]gopurs_runtime.Value)(__local_var_4_1.UnsafePtr)[0], s_prime_3_0)).PtrVal.([]gopurs_runtime.Value)))).IntVal))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_2:
return __t2
})
	})
	return indexOf_prime
}

var countTail gopurs_runtime.Value
var once_countTail sync.Once
func Get_countTail() gopurs_runtime.Value {
	once_countTail.Do(func() {
		countTail = gopurs_runtime.Func(func(p_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(accum_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
countTail:
for {
if false { continue countTail }
var p_0 = p_0_loop
_ = p_0
var s_1 = s_1_loop
_ = s_1
var accum_2 = accum_2_loop
_ = accum_2
v_3_0 := gopurs_runtime.Apply(Get_uncons(), s_1)
_ = v_3_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(v_3_0.StrVal == "Just").IntVal != 0 && gopurs_runtime.Apply(p_0, gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v_3_0.UnsafePtr)[0], "head")).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Apply3(Get_countTail(), p_0, gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v_3_0.UnsafePtr)[0], "tail"), gopurs_runtime.Int(accum_2.IntVal + gopurs_runtime.Int(1).IntVal))
goto end_branch_1
} else {

}
}
{
__t1 = accum_2
}
end_branch_1:
return __t1
}
}()
})
})
})
	})
	return countTail
}

var countFallback gopurs_runtime.Value
var once_countFallback sync.Once
func Get_countFallback() gopurs_runtime.Value {
	once_countFallback.Do(func() {
		countFallback = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_countTail(), p_0, s_1, gopurs_runtime.Int(0))
})
	})
	return countFallback
}

var countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		countPrefix = gopurs_runtime.Apply2(Get__countPrefix(), Get_countFallback(), Get_unsafeCodePointAt0())
	})
	return countPrefix
}

var dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		dropWhile = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, s_1), s_1)), s_1)
})
	})
	return dropWhile
}

var takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		takeWhile = gopurs_runtime.Func2(func(p_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_take(), gopurs_runtime.Apply2(Get_countPrefix(), p_0, s_1), s_1)
})
	})
	return takeWhile
}

var codePointFromChar gopurs_runtime.Value
var once_codePointFromChar sync.Once
func Get_codePointFromChar() gopurs_runtime.Value {
	once_codePointFromChar.Do(func() {
		codePointFromChar = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Enum.Get_toCharCode(), x_0)
})
	})
	return codePointFromChar
}

var codePointAtFallback gopurs_runtime.Value
var once_codePointAtFallback sync.Once
func Get_codePointAtFallback() gopurs_runtime.Value {
	once_codePointAtFallback.Do(func() {
		codePointAtFallback = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
codePointAtFallback:
for {
if false { continue codePointAtFallback }
var n_0 = n_0_loop
_ = n_0
var s_1 = s_1_loop
_ = s_1
v_2_0 := gopurs_runtime.Apply(Get_uncons(), s_1)
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_2_0.StrVal == "Just")).IntVal != 0 {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(n_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t2 = gopurs_runtime.Constructor1("Just", gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v_2_0.UnsafePtr)[0], "head"))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply2(Get_codePointAtFallback(), gopurs_runtime.Int(n_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.RecordGet((*[1024]gopurs_runtime.Value)(v_2_0.UnsafePtr)[0], "tail"))
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}
}()
})
})
	})
	return codePointAtFallback
}

var codePointAt gopurs_runtime.Value
var once_codePointAt sync.Once
func Get_codePointAt() gopurs_runtime.Value {
	once_codePointAt.Do(func() {
		codePointAt = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal < gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v1_1.StrVal == gopurs_runtime.Str("").StrVal)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor0("Nothing")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor1("Just", gopurs_runtime.Apply(Get_unsafeCodePointAt0(), v1_1))
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.Apply5(Get__codePointAt(), Get_codePointAtFallback(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Constructor0("Nothing"), Get_unsafeCodePointAt0(), v_0), v1_1)
}
end_branch_0:
return __t0
})
	})
	return codePointAt
}

var boundedCodePoint gopurs_runtime.Value
var once_boundedCodePoint sync.Once
func Get_boundedCodePoint() gopurs_runtime.Value {
	once_boundedCodePoint.Do(func() {
		boundedCodePoint = gopurs_runtime.RecordDict3("bottom", "top", "Ord0", gopurs_runtime.Int(0), gopurs_runtime.Int(1114111), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordCodePoint()
}))
	})
	return boundedCodePoint
}

var boundedEnumCodePoint gopurs_runtime.Value
var once_boundedEnumCodePoint sync.Once
func Get_boundedEnumCodePoint() gopurs_runtime.Value {
	once_boundedEnumCodePoint.Do(func() {
		boundedEnumCodePoint = gopurs_runtime.RecordDict5("cardinality", "fromEnum", "toEnum", "Bounded0", "Enum1", gopurs_runtime.Int(1114112), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(n_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(n_0.IntVal <= gopurs_runtime.Int(1114111).IntVal).IntVal != 0)).IntVal != 0 {
__t0 = gopurs_runtime.Constructor1("Just", n_0)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_0:
return __t0
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedCodePoint()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumCodePoint()
}))
	})
	return boundedEnumCodePoint
}

var enumCodePoint gopurs_runtime.Value
var once_enumCodePoint sync.Once
func Get_enumCodePoint() gopurs_runtime.Value {
	once_enumCodePoint.Do(func() {
		enumCodePoint = gopurs_runtime.RecordDict3("succ", "pred", "Ord0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := a_0.IntVal + gopurs_runtime.Int(1).IntVal
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_0.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_0.IntVal <= gopurs_runtime.Int(1114111).IntVal).IntVal != 0)).IntVal != 0 {
__t1 = gopurs_runtime.Constructor1("Just", __local_var_1_0)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_1:
return __t1
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_2 := a_0.IntVal - gopurs_runtime.Int(1).IntVal
_ = __local_var_1_2
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(__local_var_1_2.IntVal >= gopurs_runtime.Int(0).IntVal).IntVal != 0 && gopurs_runtime.Bool(__local_var_1_2.IntVal <= gopurs_runtime.Int(1114111).IntVal).IntVal != 0)).IntVal != 0 {
__t3 = gopurs_runtime.Constructor1("Just", __local_var_1_2)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Constructor0("Nothing")
}
end_branch_3:
return __t3
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordCodePoint()
}))
	})
	return enumCodePoint
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

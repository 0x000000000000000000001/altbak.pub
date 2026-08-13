package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString = Get_Data_String_NonEmpty_Internal_NonEmptyString()
	})
	return cache_Data_String_NonEmpty_CodeUnits_toNonEmptyString
}

var cache_Data_String_NonEmpty_CodeUnits_snoc gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_snoc sync.Once
func Get_Data_String_NonEmpty_CodeUnits_snoc() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_snoc.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_snoc(c_0_box.StrVal(), s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_snoc
}

var cache_Data_String_NonEmpty_CodeUnits_singleton gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_singleton sync.Once
func Get_Data_String_NonEmpty_CodeUnits_singleton() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_singleton.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_singleton(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_singleton
}

var cache_Data_String_NonEmpty_CodeUnits_liftS gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_liftS(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS
}

var cache_Data_String_NonEmpty_CodeUnits_takeWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_takeWhile sync.Once
func Get_Data_String_NonEmpty_CodeUnits_takeWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_takeWhile.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_takeWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_takeWhile(f_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_takeWhile
}

var cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime sync.Once
func Get_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime(pat_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime
}

var cache_Data_String_NonEmpty_CodeUnits_lastIndexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_lastIndexOf sync.Once
func Get_Data_String_NonEmpty_CodeUnits_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_lastIndexOf.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_lastIndexOf = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_lastIndexOf(x_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_lastIndexOf
}

var cache_Data_String_NonEmpty_CodeUnits_indexOf_prime gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_indexOf_prime sync.Once
func Get_Data_String_NonEmpty_CodeUnits_indexOf_prime() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_indexOf_prime.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_indexOf_prime = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_indexOf_prime(pat_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_indexOf_prime
}

var cache_Data_String_NonEmpty_CodeUnits_indexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_indexOf sync.Once
func Get_Data_String_NonEmpty_CodeUnits_indexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_indexOf.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_indexOf = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_indexOf(x_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_indexOf
}

var cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(v_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyString
}

var cache_Data_String_NonEmpty_CodeUnits_length gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_length sync.Once
func Get_Data_String_NonEmpty_CodeUnits_length() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_length.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_NonEmpty_CodeUnits_length(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_length
}

var cache_Data_String_NonEmpty_CodeUnits_splitAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_splitAt sync.Once
func Get_Data_String_NonEmpty_CodeUnits_splitAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_splitAt.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_splitAt
}

var cache_Data_String_NonEmpty_CodeUnits_take gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_take sync.Once
func Get_Data_String_NonEmpty_CodeUnits_take() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_take.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_take(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_take
}

var cache_Data_String_NonEmpty_CodeUnits_takeRight gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_takeRight sync.Once
func Get_Data_String_NonEmpty_CodeUnits_takeRight() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_takeRight.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_takeRight(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_takeRight
}

var cache_Data_String_NonEmpty_CodeUnits_toChar gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toChar sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toChar() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toChar.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toChar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_toChar(x_0_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_toChar
}

var cache_Data_String_NonEmpty_CodeUnits_toCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toCharArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_String_NonEmpty_CodeUnits_toCharArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_toCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__3897574428()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_toCharArray(), gopurs_runtime.Str(x_1.StrVal()))
_ = __local_var_2_1
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal) > (0) {
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}
end_branch_3:
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t3))})
})
}()
	})
	return cache_Data_String_NonEmpty_CodeUnits_toNonEmptyCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_uncons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_uncons sync.Once
func Get_Data_String_NonEmpty_CodeUnits_uncons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_uncons.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_uncons(nes_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_uncons
}

var cache_Data_String_NonEmpty_CodeUnits_fromFoldable1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromFoldable1 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromFoldable1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromFoldable1.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromFoldable1
}

var cache_Data_String_NonEmpty_CodeUnits_fromCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromCharArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_fromCharArray(func() []string {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray sync.Once
func Get_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__4121089788()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))).IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodeUnits_fromCharArray(), func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())})}
}
end_branch_1:
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))})
})
}()
	})
	return cache_Data_String_NonEmpty_CodeUnits_fromNonEmptyCharArray
}

var cache_Data_String_NonEmpty_CodeUnits_dropWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_dropWhile sync.Once
func Get_Data_String_NonEmpty_CodeUnits_dropWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_dropWhile.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_dropWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_dropWhile(f_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_dropWhile
}

var cache_Data_String_NonEmpty_CodeUnits_dropRight gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_dropRight sync.Once
func Get_Data_String_NonEmpty_CodeUnits_dropRight() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_dropRight.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_dropRight(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_dropRight
}

var cache_Data_String_NonEmpty_CodeUnits_drop gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_drop sync.Once
func Get_Data_String_NonEmpty_CodeUnits_drop() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_drop.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_drop(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_drop
}

var cache_Data_String_NonEmpty_CodeUnits_countPrefix gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_countPrefix sync.Once
func Get_Data_String_NonEmpty_CodeUnits_countPrefix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_countPrefix.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_countPrefix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_countPrefix(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_countPrefix
}

var cache_Data_String_NonEmpty_CodeUnits_cons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_cons sync.Once
func Get_Data_String_NonEmpty_CodeUnits_cons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_cons.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_cons(c_0_box.StrVal(), s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_cons
}

var cache_Data_String_NonEmpty_CodeUnits_charAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_charAt sync.Once
func Get_Data_String_NonEmpty_CodeUnits_charAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_charAt.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_charAt = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_charAt(x_0_box.IntVal)
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_charAt
}

var cache_Data_String_NonEmpty_CodeUnits_liftS__1768125498 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS__1768125498 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS__1768125498() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS__1768125498.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS__1768125498 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_NonEmpty_CodeUnits_liftS__1768125498(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS__1768125498
}

var cache_Data_String_NonEmpty_CodeUnits_liftS__1220682938 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS__1220682938 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS__1220682938() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS__1220682938.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS__1220682938 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_liftS__1220682938(f_0_box, v_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS__1220682938
}

var cache_Data_String_NonEmpty_CodeUnits_liftS__4059757050 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS__4059757050 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS__4059757050() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS__4059757050.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS__4059757050 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodeUnits_liftS__4059757050(f_0_box, v_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS__4059757050
}

var cache_Data_String_NonEmpty_CodeUnits_liftS__549717202 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS__549717202 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS__549717202() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS__549717202.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS__549717202 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodeUnits_liftS__549717202(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS__549717202
}

var cache_Data_String_NonEmpty_CodeUnits_liftS__3230749042 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodeUnits_liftS__3230749042 sync.Once
func Get_Data_String_NonEmpty_CodeUnits_liftS__3230749042() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodeUnits_liftS__3230749042.Do(func() {
		cache_Data_String_NonEmpty_CodeUnits_liftS__3230749042 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodeUnits_liftS__3230749042(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodeUnits_liftS__3230749042
}

func Call_Data_String_NonEmpty_CodeUnits_snoc(c_0_loop string, s_1_loop string) string {
var c_0 string = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (s_1) + (gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Str(c_0)).StrVal())
}

func Call_Data_String_NonEmpty_CodeUnits_singleton(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Str(x_0)).StrVal()
}

func Call_Data_String_NonEmpty_CodeUnits_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_CodeUnits_takeWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_liftS__549717202(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_takeWhile(), f_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(__local_var_3_1.StrVal())})}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2))}
})
}

func Call_Data_String_NonEmpty_CodeUnits_lastIndexOf_prime(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_lastIndexOf_prime(), gopurs_runtime.Str(pat_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Str(v_4.StrVal()))
})
})
}

func Call_Data_String_NonEmpty_CodeUnits_lastIndexOf(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_lastIndexOf(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_CodeUnits_indexOf_prime(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf_prime(), gopurs_runtime.Str(pat_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Str(v_4.StrVal()))
})
})
}

func Call_Data_String_NonEmpty_CodeUnits_indexOf(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_indexOf(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_CodeUnits_fromNonEmptyString(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return v_0
}

func Call_Data_String_NonEmpty_CodeUnits_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal())).IntVal
}

func Call_Data_String_NonEmpty_CodeUnits_splitAt(i_0_loop int64, nes_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
// TAST (Let): v_2_0 -> gopurs_runtime.Value
v_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "after").StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.RecordGet(v_2_0, "after").StrVal())})}
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "before").StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.RecordGet(v_2_0, "before").StrVal())})}
}
end_branch_2:
return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2))})
}

func Call_Data_String_NonEmpty_CodeUnits_take(i_0_loop int64, nes_1_loop string) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (i_0) < (1) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_NonEmpty_CodeUnits_takeRight(i_0_loop int64, nes_1_loop string) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (i_0) < (1) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) - (i_0)), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_NonEmpty_CodeUnits_toChar(x_0_loop string) *Constructor_Data_Maybe_Just {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(Get_Data_String_CodeUnits_toChar(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal())))
}

func Call_Data_String_NonEmpty_CodeUnits_toCharArray(x_0_loop string) []string {
var x_0 string = x_0_loop
_ = x_0
return func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_String_CodeUnits_toCharArray(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal())).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
}

func Call_Data_String_NonEmpty_CodeUnits_uncons(nes_0_loop string) gopurs_runtime.Value {
var nes_0 string = nes_0_loop
_ = nes_0
// TAST (Let): __local_var_1_0 -> string
__local_var_1_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(1), gopurs_runtime.Str(nes_0)).StrVal()
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(__local_var_1_0)})}
}
end_branch_1:
return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_Unsafe_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(nes_0)).StrVal()), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))})
}

func Call_Data_String_NonEmpty_CodeUnits_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_semigroupString()))}, Get_Data_String_NonEmpty_CodeUnits_singleton())
}

func Call_Data_String_NonEmpty_CodeUnits_fromCharArray(v_0_loop []string) *Constructor_Data_Maybe_Just {
var v_0 []string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodeUnits_fromCharArray(), func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t0)
}

func Call_Data_String_NonEmpty_CodeUnits_dropWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodeUnits_liftS__549717202(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_dropWhile(), f_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(__local_var_3_1.StrVal())})}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t2))}
})
}

func Call_Data_String_NonEmpty_CodeUnits_dropRight(i_0_loop int64, nes_1_loop string) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (i_0) < (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_take(), gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) - (i_0)), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_NonEmpty_CodeUnits_drop(i_0_loop int64, nes_1_loop string) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (i_0) < (gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(nes_1)).IntVal) {
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_NonEmpty_CodeUnits_countPrefix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_countPrefix(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_CodeUnits_cons(c_0_loop string, s_1_loop string) string {
var c_0 string = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (gopurs_runtime.Apply(Get_Data_String_CodeUnits_singleton(), gopurs_runtime.Str(c_0)).StrVal()) + (s_1)
}

func Call_Data_String_NonEmpty_CodeUnits_charAt(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_charAt(), gopurs_runtime.Int(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_CodeUnits_liftS__1768125498(f_0_loop gopurs_runtime.Value, v_1_loop string) int64 {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).IntVal
}

func Call_Data_String_NonEmpty_CodeUnits_liftS__1220682938(f_0_loop gopurs_runtime.Value, v_1_loop string) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)))
}

func Call_Data_String_NonEmpty_CodeUnits_liftS__4059757050(f_0_loop gopurs_runtime.Value, v_1_loop string) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)))
}

func Call_Data_String_NonEmpty_CodeUnits_liftS__549717202(f_0_loop gopurs_runtime.Value, v_1_loop string) string {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).StrVal()
}

func Call_Data_String_NonEmpty_CodeUnits_liftS__3230749042(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}



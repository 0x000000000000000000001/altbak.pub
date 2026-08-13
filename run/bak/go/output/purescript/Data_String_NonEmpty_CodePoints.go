package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_NonEmpty_CodePoints_toNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodePoints_toNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toNonEmptyString = Get_Data_String_NonEmpty_Internal_NonEmptyString()
	})
	return cache_Data_String_NonEmpty_CodePoints_toNonEmptyString
}

var cache_Data_String_NonEmpty_CodePoints_snoc gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_snoc sync.Once
func Get_Data_String_NonEmpty_CodePoints_snoc() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_snoc.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_snoc(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_snoc
}

var cache_Data_String_NonEmpty_CodePoints_singleton gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_singleton sync.Once
func Get_Data_String_NonEmpty_CodePoints_singleton() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_singleton.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_singleton(x_0_box.IntVal))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_singleton
}

var cache_Data_String_NonEmpty_CodePoints_liftS gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS sync.Once
func Get_Data_String_NonEmpty_CodePoints_liftS() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_liftS(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS
}

var cache_Data_String_NonEmpty_CodePoints_takeWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_takeWhile sync.Once
func Get_Data_String_NonEmpty_CodePoints_takeWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_takeWhile.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_takeWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_takeWhile(f_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodePoints_takeWhile
}

var cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_lastIndexOf_prime sync.Once
func Get_Data_String_NonEmpty_CodePoints_lastIndexOf_prime() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_lastIndexOf_prime.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_lastIndexOf_prime(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_lastIndexOf_prime
}

var cache_Data_String_NonEmpty_CodePoints_lastIndexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_lastIndexOf sync.Once
func Get_Data_String_NonEmpty_CodePoints_lastIndexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_lastIndexOf.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_lastIndexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_lastIndexOf(x_0_box.StrVal(), v_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_lastIndexOf
}

var cache_Data_String_NonEmpty_CodePoints_indexOf_prime gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_indexOf_prime sync.Once
func Get_Data_String_NonEmpty_CodePoints_indexOf_prime() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_indexOf_prime.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_indexOf_prime = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_indexOf_prime(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_indexOf_prime
}

var cache_Data_String_NonEmpty_CodePoints_indexOf gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_indexOf sync.Once
func Get_Data_String_NonEmpty_CodePoints_indexOf() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_indexOf.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_indexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_indexOf(x_0_box.StrVal(), v_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_indexOf
}

var cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_fromNonEmptyString(v_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromNonEmptyString
}

var cache_Data_String_NonEmpty_CodePoints_length gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_length sync.Once
func Get_Data_String_NonEmpty_CodePoints_length() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_length.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_length = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_NonEmpty_CodePoints_length(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_length
}

var cache_Data_String_NonEmpty_CodePoints_splitAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_splitAt sync.Once
func Get_Data_String_NonEmpty_CodePoints_splitAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_splitAt.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodePoints_splitAt
}

var cache_Data_String_NonEmpty_CodePoints_take gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_take sync.Once
func Get_Data_String_NonEmpty_CodePoints_take() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_take.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_take(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_take
}

var cache_Data_String_NonEmpty_CodePoints_toCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_toCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_String_NonEmpty_CodePoints_toCodePointArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_String_NonEmpty_CodePoints_toCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__911089788()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(x_1.StrVal()))
_ = __local_var_2_1
var __t3 *Constructor_Data_Maybe_Just
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
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
})
}()
	})
	return cache_Data_String_NonEmpty_CodePoints_toNonEmptyCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_uncons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_uncons sync.Once
func Get_Data_String_NonEmpty_CodePoints_uncons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_uncons.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_uncons(nes_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodePoints_uncons
}

var cache_Data_String_NonEmpty_CodePoints_fromFoldable1 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromFoldable1 sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromFoldable1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromFoldable1.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_fromFoldable1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromFoldable1
}

var cache_Data_String_NonEmpty_CodePoints_fromCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromCodePointArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_fromCodePointArray(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_fromCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray sync.Once
func Get_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__4121089788()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))).IntVal) == (0) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCodePointArray(), func() gopurs_runtime.Value {
					arr := func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())}
}
end_branch_1:
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
})
}()
	})
	return cache_Data_String_NonEmpty_CodePoints_fromNonEmptyCodePointArray
}

var cache_Data_String_NonEmpty_CodePoints_dropWhile gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_dropWhile sync.Once
func Get_Data_String_NonEmpty_CodePoints_dropWhile() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_dropWhile.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_dropWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_dropWhile(f_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodePoints_dropWhile
}

var cache_Data_String_NonEmpty_CodePoints_drop gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_drop sync.Once
func Get_Data_String_NonEmpty_CodePoints_drop() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_drop.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_drop(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_drop
}

var cache_Data_String_NonEmpty_CodePoints_countPrefix gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_countPrefix sync.Once
func Get_Data_String_NonEmpty_CodePoints_countPrefix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_countPrefix.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_countPrefix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_countPrefix(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_CodePoints_countPrefix
}

var cache_Data_String_NonEmpty_CodePoints_cons gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_cons sync.Once
func Get_Data_String_NonEmpty_CodePoints_cons() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_cons.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_cons(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_cons
}

var cache_Data_String_NonEmpty_CodePoints_codePointAt gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_codePointAt sync.Once
func Get_Data_String_NonEmpty_CodePoints_codePointAt() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_codePointAt.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_codePointAt = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_codePointAt(x_0_box.IntVal, v_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_codePointAt
}

var cache_Data_String_NonEmpty_CodePoints_liftS__1768125498 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS__1768125498 sync.Once
func Get_Data_String_NonEmpty_CodePoints_liftS__1768125498() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS__1768125498.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS__1768125498 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_String_NonEmpty_CodePoints_liftS__1768125498(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS__1768125498
}

var cache_Data_String_NonEmpty_CodePoints_liftS__1220682938 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS__1220682938 sync.Once
func Get_Data_String_NonEmpty_CodePoints_liftS__1220682938() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS__1220682938.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS__1220682938 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_CodePoints_liftS__1220682938(f_0_box, v_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS__1220682938
}

var cache_Data_String_NonEmpty_CodePoints_liftS__549717202 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS__549717202 sync.Once
func Get_Data_String_NonEmpty_CodePoints_liftS__549717202() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS__549717202.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS__549717202 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_CodePoints_liftS__549717202(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS__549717202
}

var cache_Data_String_NonEmpty_CodePoints_liftS__3230749042 gopurs_runtime.Value
var once_Data_String_NonEmpty_CodePoints_liftS__3230749042 sync.Once
func Get_Data_String_NonEmpty_CodePoints_liftS__3230749042() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CodePoints_liftS__3230749042.Do(func() {
		cache_Data_String_NonEmpty_CodePoints_liftS__3230749042 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CodePoints_liftS__3230749042(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_CodePoints_liftS__3230749042
}

func Call_Data_String_NonEmpty_CodePoints_snoc(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (s_1) + (gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(c_0)).StrVal())
}

func Call_Data_String_NonEmpty_CodePoints_singleton(x_0_loop int64) string {
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(x_0)).StrVal()
}

func Call_Data_String_NonEmpty_CodePoints_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_CodePoints_takeWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_liftS__549717202(), gopurs_runtime.Apply(Get_Data_String_CodePoints_takeWhile(), f_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 *Constructor_Data_Maybe_Just
{
if (__local_var_3_1.StrVal()) == ("") {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(__local_var_3_1.StrVal())}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
})
}

func Call_Data_String_NonEmpty_CodePoints_lastIndexOf_prime(pat_0_loop string, x_1_loop int64, v_2_loop string) *Constructor_Data_Maybe_Just {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_String_CodePoints_lastIndexOf_prime(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2)))
}

func Call_Data_String_NonEmpty_CodePoints_lastIndexOf(x_0_loop string, v_1_loop string) *Constructor_Data_Maybe_Just {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_String_CodePoints_lastIndexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1)))
}

func Call_Data_String_NonEmpty_CodePoints_indexOf_prime(pat_0_loop string, x_1_loop int64, v_2_loop string) *Constructor_Data_Maybe_Just {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_String_CodePoints_indexOf_prime(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2)))
}

func Call_Data_String_NonEmpty_CodePoints_indexOf(x_0_loop string, v_1_loop string) *Constructor_Data_Maybe_Just {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_String_CodePoints_indexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1)))
}

func Call_Data_String_NonEmpty_CodePoints_fromNonEmptyString(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return v_0
}

func Call_Data_String_NonEmpty_CodePoints_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal()))))).IntVal
}

func Call_Data_String_NonEmpty_CodePoints_splitAt(i_0_loop int64, nes_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
// TAST (Let): v_2_0 -> gopurs_runtime.Value
v_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodePoints_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
_ = v_2_0
var __t1 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.RecordGet(v_2_0, "after").StrVal()) == ("") {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.RecordGet(v_2_0, "after").StrVal())}
}
end_branch_1:
var __t2 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.RecordGet(v_2_0, "before").StrVal()) == ("") {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.RecordGet(v_2_0, "before").StrVal())}
}
end_branch_2:
return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
}

func Call_Data_String_NonEmpty_CodePoints_take(i_0_loop int64, nes_1_loop string) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t1 *Constructor_Data_Maybe_Just
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
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_String_NonEmpty_CodePoints_toCodePointArray(x_0_loop string) []int64 {
var x_0 string = x_0_loop
_ = x_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal())).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_Data_String_NonEmpty_CodePoints_uncons(nes_0_loop string) gopurs_runtime.Value {
var nes_0 string = nes_0_loop
_ = nes_0
// TAST (Let): __local_var_1_0 -> string
__local_var_1_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(1), gopurs_runtime.Str(nes_0)).StrVal())).IntVal), gopurs_runtime.Str(nes_0)).StrVal()
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_1_0) == ("") {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(__local_var_1_0)}
}
end_branch_1:
return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__1577979644()
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_String_CodePoints_codePointAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(nes_0))))}).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_String_NonEmpty_CodePoints_fromFoldable1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFoldable1_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_semigroupString()))}, Get_Data_String_NonEmpty_CodePoints_singleton())
}

func Call_Data_String_NonEmpty_CodePoints_fromCodePointArray(v_0_loop []int64) *Constructor_Data_Maybe_Just {
var v_0 []int64 = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (0) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_CodePoints_fromCodePointArray(), func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_String_NonEmpty_CodePoints_dropWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_NonEmpty_CodePoints_liftS__549717202(), gopurs_runtime.Apply(Get_Data_String_CodePoints_dropWhile(), f_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 *Constructor_Data_Maybe_Just
{
if (__local_var_3_1.StrVal()) == ("") {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(__local_var_3_1.StrVal())}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
})
}

func Call_Data_String_NonEmpty_CodePoints_drop(i_0_loop int64, nes_1_loop string) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t1 *Constructor_Data_Maybe_Just
{
var __t0 bool
{
if (i_0) < (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(Get_Data_String_CodePoints_toCodePointArray(), gopurs_runtime.Str(nes_1))))).IntVal) {
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
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodeUnits_drop(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Data_String_CodeUnits_length(), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())).IntVal), gopurs_runtime.Str(nes_1)).StrVal())}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_String_NonEmpty_CodePoints_countPrefix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodePoints_countPrefix(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_CodePoints_cons(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return (gopurs_runtime.Apply(Get_Data_String_CodePoints_singleton(), gopurs_runtime.Int(c_0)).StrVal()) + (s_1)
}

func Call_Data_String_NonEmpty_CodePoints_codePointAt(x_0_loop int64, v_1_loop string) *Constructor_Data_Maybe_Just {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_String_CodePoints_codePointAt(), gopurs_runtime.Int(x_0), gopurs_runtime.Str(v_1)))
}

func Call_Data_String_NonEmpty_CodePoints_liftS__1768125498(f_0_loop gopurs_runtime.Value, v_1_loop string) int64 {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).IntVal
}

func Call_Data_String_NonEmpty_CodePoints_liftS__1220682938(f_0_loop gopurs_runtime.Value, v_1_loop string) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)))
}

func Call_Data_String_NonEmpty_CodePoints_liftS__549717202(f_0_loop gopurs_runtime.Value, v_1_loop string) string {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).StrVal()
}

func Call_Data_String_NonEmpty_CodePoints_liftS__3230749042(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}



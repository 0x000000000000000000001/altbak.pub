package Data_String_NonEmpty_CodeUnits

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Array_NonEmpty_Internal "gopurs/output/Data.Array.NonEmpty.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_String_Unsafe "gopurs/output/Data.String.Unsafe"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_snoc gopurs_runtime.Value
var once_snoc sync.Once
func Get_snoc() gopurs_runtime.Value {
	once_snoc.Do(func() {
		cache_snoc = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_snoc(c_0_box.StrVal(), s_1_box.StrVal()))
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_singleton(x_0_box.StrVal()))
})
	})
	return cache_singleton
}

var cache_takeWhile gopurs_runtime.Value
var once_takeWhile sync.Once
func Get_takeWhile() gopurs_runtime.Value {
	once_takeWhile.Do(func() {
		cache_takeWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_takeWhile(f_0_box)
})
	})
	return cache_takeWhile
}

var cache_lastIndexOf_prime gopurs_runtime.Value
var once_lastIndexOf_prime sync.Once
func Get_lastIndexOf_prime() gopurs_runtime.Value {
	once_lastIndexOf_prime.Do(func() {
		cache_lastIndexOf_prime = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lastIndexOf_prime(pat_0_box.StrVal())
})
	})
	return cache_lastIndexOf_prime
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lastIndexOf(x_0_box.StrVal())
})
	})
	return cache_lastIndexOf
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_indexOf_prime(pat_0_box.StrVal())
})
	})
	return cache_indexOf_prime
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_indexOf(x_0_box.StrVal())
})
	})
	return cache_indexOf
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

var cache_splitAt gopurs_runtime.Value
var once_splitAt sync.Once
func Get_splitAt() gopurs_runtime.Value {
	once_splitAt.Do(func() {
		cache_splitAt = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_splitAt(i_0_box.IntVal, nes_1_box.StrVal())
})
	})
	return cache_splitAt
}

var cache_take gopurs_runtime.Value
var once_take sync.Once
func Get_take() gopurs_runtime.Value {
	once_take.Do(func() {
		cache_take = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_take(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_take
}

var cache_takeRight gopurs_runtime.Value
var once_takeRight sync.Once
func Get_takeRight() gopurs_runtime.Value {
	once_takeRight.Do(func() {
		cache_takeRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_takeRight(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_takeRight
}

var cache_toChar gopurs_runtime.Value
var once_toChar sync.Once
func Get_toChar() gopurs_runtime.Value {
	once_toChar.Do(func() {
		cache_toChar = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_toChar(x_0_box.StrVal()))}
})
	})
	return cache_toChar
}

var cache_toCharArray gopurs_runtime.Value
var once_toCharArray sync.Once
func Get_toCharArray() gopurs_runtime.Value {
	once_toCharArray.Do(func() {
		cache_toCharArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toCharArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toCharArray
}

var cache_toNonEmptyCharArray gopurs_runtime.Value
var once_toNonEmptyCharArray sync.Once
func Get_toNonEmptyCharArray() gopurs_runtime.Value {
	once_toNonEmptyCharArray.Do(func() {
		cache_toNonEmptyCharArray = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(Get_unsafePartial__3550303069(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust__3897574428()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toCharArray(), gopurs_runtime.Str(x_1.StrVal()))
_ = __local_var_2_1
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (gopurs_runtime.Int(gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal).IntVal) > (gopurs_runtime.Int(0).IntVal) {
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_3:
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]gopurs_runtime.Value]](__t3))})
})
}()
	})
	return cache_toNonEmptyCharArray
}

var cache_uncons gopurs_runtime.Value
var once_uncons sync.Once
func Get_uncons() gopurs_runtime.Value {
	once_uncons.Do(func() {
		cache_uncons = gopurs_runtime.Func(func(nes_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncons(nes_0_box.StrVal())
})
	})
	return cache_uncons
}

var cache_fromFoldable1 gopurs_runtime.Value
var once_fromFoldable1 sync.Once
func Get_fromFoldable1() gopurs_runtime.Value {
	once_fromFoldable1.Do(func() {
		cache_fromFoldable1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromFoldable1(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_fromFoldable1
}

var cache_fromCharArray gopurs_runtime.Value
var once_fromCharArray sync.Once
func Get_fromCharArray() gopurs_runtime.Value {
	once_fromCharArray.Do(func() {
		cache_fromCharArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromCharArray(func() []string {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))}
})
	})
	return cache_fromCharArray
}

var cache_fromNonEmptyCharArray gopurs_runtime.Value
var once_fromNonEmptyCharArray sync.Once
func Get_fromNonEmptyCharArray() gopurs_runtime.Value {
	once_fromNonEmptyCharArray.Do(func() {
		cache_fromNonEmptyCharArray = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(Get_unsafePartial__1849950365(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust__4121089788()
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_fromCharArray(), func() gopurs_runtime.Value {
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
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1))})
})
}()
	})
	return cache_fromNonEmptyCharArray
}

var cache_dropWhile gopurs_runtime.Value
var once_dropWhile sync.Once
func Get_dropWhile() gopurs_runtime.Value {
	once_dropWhile.Do(func() {
		cache_dropWhile = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_dropWhile(f_0_box)
})
	})
	return cache_dropWhile
}

var cache_dropRight gopurs_runtime.Value
var once_dropRight sync.Once
func Get_dropRight() gopurs_runtime.Value {
	once_dropRight.Do(func() {
		cache_dropRight = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_dropRight(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_dropRight
}

var cache_drop gopurs_runtime.Value
var once_drop sync.Once
func Get_drop() gopurs_runtime.Value {
	once_drop.Do(func() {
		cache_drop = gopurs_runtime.Func2(func(i_0_box gopurs_runtime.Value, nes_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_drop(i_0_box.IntVal, nes_1_box.StrVal()))}
})
	})
	return cache_drop
}

var cache_countPrefix gopurs_runtime.Value
var once_countPrefix sync.Once
func Get_countPrefix() gopurs_runtime.Value {
	once_countPrefix.Do(func() {
		cache_countPrefix = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_countPrefix(x_0_box)
})
	})
	return cache_countPrefix
}

var cache_cons gopurs_runtime.Value
var once_cons sync.Once
func Get_cons() gopurs_runtime.Value {
	once_cons.Do(func() {
		cache_cons = gopurs_runtime.Func2(func(c_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_cons(c_0_box.StrVal(), s_1_box.StrVal()))
})
	})
	return cache_cons
}

var cache_charAt gopurs_runtime.Value
var once_charAt sync.Once
func Get_charAt() gopurs_runtime.Value {
	once_charAt.Do(func() {
		cache_charAt = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_charAt(x_0_box.IntVal)
})
	})
	return cache_charAt
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

var cache_fromArray__260997498 gopurs_runtime.Value
var once_fromArray__260997498 sync.Once
func Get_fromArray__260997498() gopurs_runtime.Value {
	once_fromArray__260997498.Do(func() {
		cache_fromArray__260997498 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromArray__260997498(func() []string {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()))}
})
	})
	return cache_fromArray__260997498
}

var cache_toArray__1949224283 gopurs_runtime.Value
var once_toArray__1949224283 sync.Once
func Get_toArray__1949224283() gopurs_runtime.Value {
	once_toArray__1949224283.Do(func() {
		cache_toArray__1949224283 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray__1949224283(func() []string {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toArray__1949224283
}

var cache_unsafeFromArray__3238140064 gopurs_runtime.Value
var once_unsafeFromArray__3238140064 sync.Once
func Get_unsafeFromArray__3238140064() gopurs_runtime.Value {
	once_unsafeFromArray__3238140064.Do(func() {
		cache_unsafeFromArray__3238140064 = pkg_Data_Array_NonEmpty_Internal.Get_NonEmptyArray()
	})
	return cache_unsafeFromArray__3238140064
}

var cache_fromJust__4121089788 gopurs_runtime.Value
var once_fromJust__4121089788 sync.Once
func Get_fromJust__4121089788() gopurs_runtime.Value {
	once_fromJust__4121089788.Do(func() {
		cache_fromJust__4121089788 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__4121089788(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__4121089788
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

var cache_fromJust__3897574428 gopurs_runtime.Value
var once_fromJust__3897574428 sync.Once
func Get_fromJust__3897574428() gopurs_runtime.Value {
	once_fromJust__3897574428.Do(func() {
		cache_fromJust__3897574428 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__3897574428(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__3897574428
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

var cache_greaterThan__3448835524 gopurs_runtime.Value
var once_greaterThan__3448835524 sync.Once
func Get_greaterThan__3448835524() gopurs_runtime.Value {
	once_greaterThan__3448835524.Do(func() {
		cache_greaterThan__3448835524 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__3448835524(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__3448835524
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

var cache_foldMap1__3675913824 gopurs_runtime.Value
var once_foldMap1__3675913824 sync.Once
func Get_foldMap1__3675913824() gopurs_runtime.Value {
	once_foldMap1__3675913824.Do(func() {
		cache_foldMap1__3675913824 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3675913824(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3675913824
}

var cache_foldMap1__3342855683 gopurs_runtime.Value
var once_foldMap1__3342855683 sync.Once
func Get_foldMap1__3342855683() gopurs_runtime.Value {
	once_foldMap1__3342855683.Do(func() {
		cache_foldMap1__3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3342855683(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3342855683
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

var cache_liftS__1768125498 gopurs_runtime.Value
var once_liftS__1768125498 sync.Once
func Get_liftS__1768125498() gopurs_runtime.Value {
	once_liftS__1768125498.Do(func() {
		cache_liftS__1768125498 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_liftS__1768125498(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_liftS__1768125498
}

var cache_liftS__1220682938 gopurs_runtime.Value
var once_liftS__1220682938 sync.Once
func Get_liftS__1220682938() gopurs_runtime.Value {
	once_liftS__1220682938.Do(func() {
		cache_liftS__1220682938 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_liftS__1220682938(f_0_box, v_1_box.StrVal()))}
})
	})
	return cache_liftS__1220682938
}

var cache_liftS__4059757050 gopurs_runtime.Value
var once_liftS__4059757050 sync.Once
func Get_liftS__4059757050() gopurs_runtime.Value {
	once_liftS__4059757050.Do(func() {
		cache_liftS__4059757050 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_liftS__4059757050(f_0_box, v_1_box.StrVal()))}
})
	})
	return cache_liftS__4059757050
}

var cache_liftS__549717202 gopurs_runtime.Value
var once_liftS__549717202 sync.Once
func Get_liftS__549717202() gopurs_runtime.Value {
	once_liftS__549717202.Do(func() {
		cache_liftS__549717202 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftS__549717202(f_0_box, v_1_box.StrVal())
})
	})
	return cache_liftS__549717202
}

var cache_liftS__3230749042 gopurs_runtime.Value
var once_liftS__3230749042 sync.Once
func Get_liftS__3230749042() gopurs_runtime.Value {
	once_liftS__3230749042.Do(func() {
		cache_liftS__3230749042 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftS__3230749042(f_0_box, v_1_box.StrVal())
})
	})
	return cache_liftS__3230749042
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_unsafePartial__1849950365 gopurs_runtime.Value
var once_unsafePartial__1849950365 sync.Once
func Get_unsafePartial__1849950365() gopurs_runtime.Value {
	once_unsafePartial__1849950365.Do(func() {
		cache_unsafePartial__1849950365 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1849950365
}

var cache_unsafePartial__3550303069 gopurs_runtime.Value
var once_unsafePartial__3550303069 sync.Once
func Get_unsafePartial__3550303069() gopurs_runtime.Value {
	once_unsafePartial__3550303069.Do(func() {
		cache_unsafePartial__3550303069 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__3550303069
}

func Call_snoc(c_0_loop string, s_1_loop string) string {
var c_0 string = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return Call_append__493084344(gopurs_runtime.Str(s_1), gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), gopurs_runtime.Str(c_0)).StrVal())).StrVal()
}

func Call_singleton(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), gopurs_runtime.Str(x_0)).StrVal()
}

func Call_takeWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply(Get_liftS__549717202(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_takeWhile(), f_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_3_1.StrVal())})}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t2))}
})
}

func Call_lastIndexOf_prime(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf_prime(), gopurs_runtime.Str(pat_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Str(v_4.StrVal()))
})
})
}

func Call_lastIndexOf(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_lastIndexOf(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_indexOf_prime(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf_prime(), gopurs_runtime.Str(pat_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Str(v_4.StrVal()))
})
})
}

func Call_indexOf(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_indexOf(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal())).IntVal
}

func Call_splitAt(i_0_loop int64, nes_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
v_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
_ = v_2_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "after").StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.RecordGet(v_2_0, "after").StrVal())})}
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_2_0, "before").StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.RecordGet(v_2_0, "before").StrVal())})}
}
end_branch_2:
return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t2))})
}

func Call_take(i_0_loop int64, nes_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(i_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_takeRight(i_0_loop int64, nes_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(gopurs_runtime.Int(i_0), gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int((gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(nes_1)).IntVal) - (i_0)), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_toChar(x_0_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toChar(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal())))
}

func Call_toCharArray(x_0_loop string) []string {
var x_0 string = x_0_loop
_ = x_0
return func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_toCharArray(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal())).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
}

func Call_uncons(nes_0_loop string) gopurs_runtime.Value {
var nes_0 string = nes_0_loop
_ = nes_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(1), gopurs_runtime.Str(nes_0)).StrVal()
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_1_0)})}
}
end_branch_1:
return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_Unsafe.Get_charAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(nes_0)).StrVal()), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1))})
}

func Call_fromFoldable1(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[string]](pkg_Data_Semigroup.Get_semigroupString()))}, Get_singleton())
}

func Call_fromCharArray(v_0_loop []string) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 []string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(v_0))).IntVal) == (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_fromCharArray(), func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_dropWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply(Get_liftS__549717202(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_dropWhile(), f_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_1_0, x_2)
_ = __local_var_3_1
var __t2 gopurs_runtime.Value
{
if (__local_var_3_1.StrVal()) == ("") {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_3_1.StrVal())})}
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t2))}
})
}

func Call_dropRight(i_0_loop int64, nes_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(gopurs_runtime.Int(i_0), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(nes_1)).IntVal))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_take(), gopurs_runtime.Int((gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(nes_1)).IntVal) - (i_0)), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_drop(i_0_loop int64, nes_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(gopurs_runtime.Int(i_0), gopurs_runtime.Int(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Str(nes_1)).IntVal))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1)).StrVal())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_countPrefix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_countPrefix(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_cons(c_0_loop string, s_1_loop string) string {
var c_0 string = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_singleton(), gopurs_runtime.Str(c_0)).StrVal()), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_charAt(x_0_loop int64) gopurs_runtime.Value {
var x_0 int64 = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_charAt(), gopurs_runtime.Int(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromArray__260997498(xs_0_loop []string) *pkg_Data_Maybe.Constructor_Just[[]string] {
var xs_0 []string = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_greaterThan__4087042607(), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_0))).IntVal), gopurs_runtime.Int(0)).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]string]](__t0)
}

func Call_toArray__1949224283(v_0_loop []string) []string {
var v_0 []string = v_0_loop
_ = v_0
return v_0
}

func Call_fromJust__4121089788(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
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

func Call_fromJust__3897574428(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
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

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_greaterThan__3448835524(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
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

func Call_foldMap1__3675913824(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(dict_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[string]](pkg_Data_Semigroup.Get_semigroupString()))})
}

func Call_foldMap1__3342855683(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
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

func Call_liftS__1768125498(f_0_loop gopurs_runtime.Value, v_1_loop string) int64 {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).IntVal
}

func Call_liftS__1220682938(f_0_loop gopurs_runtime.Value, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)))
}

func Call_liftS__4059757050(f_0_loop gopurs_runtime.Value, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)))
}

func Call_liftS__549717202(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_liftS__3230749042(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}



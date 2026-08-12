package Data_String_NonEmpty_CodePoints

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Array_NonEmpty_Internal "gopurs/output/Data.Array.NonEmpty.Internal"
	pkg_Data_Enum "gopurs/output/Data.Enum"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_String_CodePoints "gopurs/output/Data.String.CodePoints"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
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
return gopurs_runtime.Str(Call_snoc(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_snoc
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_singleton(x_0_box.IntVal))
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
		cache_lastIndexOf_prime = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_lastIndexOf_prime(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal()))}
})
	})
	return cache_lastIndexOf_prime
}

var cache_lastIndexOf gopurs_runtime.Value
var once_lastIndexOf sync.Once
func Get_lastIndexOf() gopurs_runtime.Value {
	once_lastIndexOf.Do(func() {
		cache_lastIndexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_lastIndexOf(x_0_box.StrVal(), v_1_box.StrVal()))}
})
	})
	return cache_lastIndexOf
}

var cache_indexOf_prime gopurs_runtime.Value
var once_indexOf_prime sync.Once
func Get_indexOf_prime() gopurs_runtime.Value {
	once_indexOf_prime.Do(func() {
		cache_indexOf_prime = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_indexOf_prime(pat_0_box.StrVal(), x_1_box.IntVal, v_2_box.StrVal()))}
})
	})
	return cache_indexOf_prime
}

var cache_indexOf gopurs_runtime.Value
var once_indexOf sync.Once
func Get_indexOf() gopurs_runtime.Value {
	once_indexOf.Do(func() {
		cache_indexOf = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_indexOf(x_0_box.StrVal(), v_1_box.StrVal()))}
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

var cache_toCodePointArray gopurs_runtime.Value
var once_toCodePointArray sync.Once
func Get_toCodePointArray() gopurs_runtime.Value {
	once_toCodePointArray.Do(func() {
		cache_toCodePointArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toCodePointArray(x_0_box.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toCodePointArray
}

var cache_toNonEmptyCodePointArray gopurs_runtime.Value
var once_toNonEmptyCodePointArray sync.Once
func Get_toNonEmptyCodePointArray() gopurs_runtime.Value {
	once_toNonEmptyCodePointArray.Do(func() {
		cache_toNonEmptyCodePointArray = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(Get_unsafePartial__1954781085(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust__911089788()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_1 := gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), x_1)
_ = __local_var_2_1
var __t3 gopurs_runtime.Value
{
var __t2 bool
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_2_1))).IntVal) > (0) {
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_2_1})}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_3:
return gopurs_runtime.Apply(__local_var_0_0, __t3)
})
}()
	})
	return cache_toNonEmptyCodePointArray
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

var cache_fromCodePointArray gopurs_runtime.Value
var once_fromCodePointArray sync.Once
func Get_fromCodePointArray() gopurs_runtime.Value {
	once_fromCodePointArray.Do(func() {
		cache_fromCodePointArray = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromCodePointArray(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))}
})
	})
	return cache_fromCodePointArray
}

var cache_fromNonEmptyCodePointArray gopurs_runtime.Value
var once_fromNonEmptyCodePointArray sync.Once
func Get_fromNonEmptyCodePointArray() gopurs_runtime.Value {
	once_fromNonEmptyCodePointArray.Do(func() {
		cache_fromNonEmptyCodePointArray = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(Get_unsafePartial__1849950365(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust__4121089788()
}))
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(x_1))).IntVal) == (0) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_fromCodePointArray(), x_1)})}
}
end_branch_1:
return gopurs_runtime.Apply(__local_var_0_0, __t1)
})
}()
	})
	return cache_fromNonEmptyCodePointArray
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
return gopurs_runtime.Str(Call_cons(c_0_box.IntVal, s_1_box.StrVal()))
})
	})
	return cache_cons
}

var cache_codePointAt gopurs_runtime.Value
var once_codePointAt sync.Once
func Get_codePointAt() gopurs_runtime.Value {
	once_codePointAt.Do(func() {
		cache_codePointAt = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_codePointAt(x_0_box.IntVal, v_1_box.StrVal()))}
})
	})
	return cache_codePointAt
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

var cache_fromArray__2195001498 gopurs_runtime.Value
var once_fromArray__2195001498 sync.Once
func Get_fromArray__2195001498() gopurs_runtime.Value {
	once_fromArray__2195001498.Do(func() {
		cache_fromArray__2195001498 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromArray__2195001498(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))}
})
	})
	return cache_fromArray__2195001498
}

var cache_toArray__2781090619 gopurs_runtime.Value
var once_toArray__2781090619 sync.Once
func Get_toArray__2781090619() gopurs_runtime.Value {
	once_toArray__2781090619.Do(func() {
		cache_toArray__2781090619 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_toArray__2781090619(func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(v_0_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_toArray__2781090619
}

var cache_unsafeFromArray__3704033600 gopurs_runtime.Value
var once_unsafeFromArray__3704033600 sync.Once
func Get_unsafeFromArray__3704033600() gopurs_runtime.Value {
	once_unsafeFromArray__3704033600.Do(func() {
		cache_unsafeFromArray__3704033600 = pkg_Data_Array_NonEmpty_Internal.Get_NonEmptyArray()
	})
	return cache_unsafeFromArray__3704033600
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

var cache_fromEnum__1606852103 gopurs_runtime.Value
var once_fromEnum__1606852103 sync.Once
func Get_fromEnum__1606852103() gopurs_runtime.Value {
	once_fromEnum__1606852103.Do(func() {
		cache_fromEnum__1606852103 = gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum")
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
		cache_toEnumWithDefaults__3941305703 = func() gopurs_runtime.Value {
bottom2_0_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "Bounded0"), gopurs_runtime.Value{}), "bottom")
_ = bottom2_0_0
return gopurs_runtime.Func(func(low_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
v_4_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "toEnum"), x_3)
_ = v_4_1
var __t3 gopurs_runtime.Value
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr != nil) {
__t3 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_4_1.UnsafePtr).V0
goto end_branch_3
} else {

}
}
{
if (v_4_1.Type == 9 && v_4_1.IntVal == 930809136 && v_4_1.UnsafePtr == nil) {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThan__4087042607(x_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Enum.Get_boundedEnumChar(), "fromEnum"), bottom2_0_0))).IntVal) != (0) {
__t2 = low_1
goto end_branch_2
} else {

}
}
{
__t2 = high_2
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
}()
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
		cache_eq__2843686287 = pkg_Data_Eq.Get_eqIntImpl()
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
		cache_div__2185172824 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "div")
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
		cache_mod__2185172824 = gopurs_runtime.RecordGet(pkg_Data_EuclideanRing.Get_euclideanRingInt(), "mod")
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
		cache_map__291265340 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__291265340
}

var cache_map__2116777468 gopurs_runtime.Value
var once_map__2116777468 sync.Once
func Get_map__2116777468() gopurs_runtime.Value {
	once_map__2116777468.Do(func() {
		cache_map__2116777468 = gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map")
	})
	return cache_map__2116777468
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj")
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
		cache_disj__3676519832 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj")
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
		cache_not__3201284355 = gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not")
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

var cache_fromJust__1577979644 gopurs_runtime.Value
var once_fromJust__1577979644 sync.Once
func Get_fromJust__1577979644() gopurs_runtime.Value {
	once_fromJust__1577979644.Do(func() {
		cache_fromJust__1577979644 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1577979644(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__1577979644
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

var cache_fromJust__911089788 gopurs_runtime.Value
var once_fromJust__911089788 sync.Once
func Get_fromJust__911089788() gopurs_runtime.Value {
	once_fromJust__911089788.Do(func() {
		cache_fromJust__911089788 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__911089788(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fromJust__911089788
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
		cache_sub__1043827704 = pkg_Data_Ring.Get_intSub()
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

var cache_foldMap1__1646952192 gopurs_runtime.Value
var once_foldMap1__1646952192 sync.Once
func Get_foldMap1__1646952192() gopurs_runtime.Value {
	once_foldMap1__1646952192.Do(func() {
		cache_foldMap1__1646952192 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__1646952192(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__1646952192
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
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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
		cache_add__560788792 = pkg_Data_Semiring.Get_intAdd()
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
		cache_mul__560788792 = pkg_Data_Semiring.Get_intMul()
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
}), gopurs_runtime.Apply4(pkg_Data_Unfoldable.Get_unfoldrArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldableArray__644327338
}

var cache_unfoldr__1842498883 gopurs_runtime.Value
var once_unfoldr__1842498883 sync.Once
func Get_unfoldr__1842498883() gopurs_runtime.Value {
	once_unfoldr__1842498883.Do(func() {
		cache_unfoldr__1842498883 = gopurs_runtime.RecordGet(pkg_Data_Unfoldable.Get_unfoldableArray(), "unfoldr")
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
		cache_unfoldable1Array__4196906331 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(pkg_Data_Unfoldable1.Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
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

var cache_unsafePartial__1849950365 gopurs_runtime.Value
var once_unsafePartial__1849950365 sync.Once
func Get_unsafePartial__1849950365() gopurs_runtime.Value {
	once_unsafePartial__1849950365.Do(func() {
		cache_unsafePartial__1849950365 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1849950365
}

var cache_unsafePartial__1090765981 gopurs_runtime.Value
var once_unsafePartial__1090765981 sync.Once
func Get_unsafePartial__1090765981() gopurs_runtime.Value {
	once_unsafePartial__1090765981.Do(func() {
		cache_unsafePartial__1090765981 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1090765981
}

var cache_unsafePartial__1954781085 gopurs_runtime.Value
var once_unsafePartial__1954781085 sync.Once
func Get_unsafePartial__1954781085() gopurs_runtime.Value {
	once_unsafePartial__1954781085.Do(func() {
		cache_unsafePartial__1954781085 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1954781085
}

func Call_snoc(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(s_1), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), gopurs_runtime.Int(c_0))).StrVal()
}

func Call_singleton(x_0_loop int64) string {
var x_0 int64 = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), gopurs_runtime.Int(x_0)).StrVal()
}

func Call_takeWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply(Get_liftS__549717202(), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_takeWhile(), f_0))
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_3_1})}
}
end_branch_2:
return __t2
})
}

func Call_lastIndexOf_prime(pat_0_loop string, x_1_loop int64, v_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply3(pkg_Data_String_CodePoints.Get_lastIndexOf_prime(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2)))
}

func Call_lastIndexOf(x_0_loop string, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_lastIndexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1)))
}

func Call_indexOf_prime(pat_0_loop string, x_1_loop int64, v_2_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var pat_0 string = pat_0_loop
_ = pat_0
var x_1 int64 = x_1_loop
_ = x_1
var v_2 string = v_2_loop
_ = v_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply3(pkg_Data_String_CodePoints.Get_indexOf_prime(), gopurs_runtime.Str(pat_0), gopurs_runtime.Int(x_1), gopurs_runtime.Str(v_2)))
}

func Call_indexOf(x_0_loop string, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_indexOf(), gopurs_runtime.Str(x_0), gopurs_runtime.Str(v_1)))
}

func Call_length(x_0_loop string) int64 {
var x_0 string = x_0_loop
_ = x_0
return gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), gopurs_runtime.Str(x_0))))).IntVal
}

func Call_splitAt(i_0_loop int64, nes_1_loop string) gopurs_runtime.Value {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
v_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_splitAt(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))
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
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(v_2_0, "after")})}
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(v_2_0, "before")})}
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_toCodePointArray(x_0_loop string) []int64 {
var x_0 string = x_0_loop
_ = x_0
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), gopurs_runtime.Str(x_0)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_uncons(nes_0_loop string) gopurs_runtime.Value {
var nes_0 string = nes_0_loop
_ = nes_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(1), gopurs_runtime.Str(nes_0))), gopurs_runtime.Str(nes_0))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_1_0})}
}
end_branch_1:
return gopurs_runtime.RecordDict2("head", "tail", gopurs_runtime.Int(gopurs_runtime.Apply2(Get_unsafePartial__905556445(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust__1577979644()
}), gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_codePointAt(), gopurs_runtime.Int(0), gopurs_runtime.Str(nes_0))).IntVal), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1))})
}

func Call_fromFoldable1(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply2(dictFoldable1_0.V1, pkg_Data_Semigroup.Get_semigroupString(), Get_singleton())
}

func Call_fromCodePointArray(v_0_loop []int64) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 []int64 = v_0_loop
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
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_fromCodePointArray(), func() gopurs_runtime.Value {
					arr := v_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}())})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_dropWhile(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
__local_var_1_0 := gopurs_runtime.Apply(Get_liftS__549717202(), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_dropWhile(), f_0))
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
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, __local_var_3_1})}
}
end_branch_2:
return __t2
})
}

func Call_drop(i_0_loop int64, nes_1_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var i_0 int64 = i_0_loop
_ = i_0
var nes_1 string = nes_1_loop
_ = nes_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__4087042607(gopurs_runtime.Int(i_0), gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_toCodePointArray(), gopurs_runtime.Str(nes_1))))))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_drop(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_length(), gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(i_0), gopurs_runtime.Str(nes_1))), gopurs_runtime.Str(nes_1))})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_countPrefix(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_countPrefix(), x_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, v_2)
})
}

func Call_cons(c_0_loop int64, s_1_loop string) string {
var c_0 int64 = c_0_loop
_ = c_0
var s_1 string = s_1_loop
_ = s_1
return gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(pkg_Data_String_CodePoints.Get_singleton(), gopurs_runtime.Int(c_0)), gopurs_runtime.Str(s_1)).StrVal()
}

func Call_codePointAt(x_0_loop int64, v_1_loop string) *pkg_Data_Maybe.Constructor_Just[int64] {
var x_0 int64 = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_codePointAt(), gopurs_runtime.Int(x_0), gopurs_runtime.Str(v_1)))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fromArray__2195001498(xs_0_loop []int64) *pkg_Data_Maybe.Constructor_Just[[]int64] {
var xs_0 []int64 = xs_0_loop
_ = xs_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int(int64(len(xs_0))), gopurs_runtime.Int(0))).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
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
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[[]int64]](__t0)
}

func Call_toArray__2781090619(v_0_loop []int64) []int64 {
var v_0 []int64 = v_0_loop
_ = v_0
return v_0
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

func Call_defaultPred__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply(fromEnum_prime_1, a_2), gopurs_runtime.Int(1))))
}

func Call_defaultPred__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Apply2(Get_sub__1043827704(), gopurs_runtime.Apply(fromEnum_prime_1, a_2), gopurs_runtime.Int(1))))
}

func Call_defaultSucc__1581620096(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Apply(fromEnum_prime_1, a_2), gopurs_runtime.Int(1))))
}

func Call_defaultSucc__2204581824(toEnum_prime_0_loop gopurs_runtime.Value, fromEnum_prime_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var toEnum_prime_0 gopurs_runtime.Value = toEnum_prime_0_loop
_ = toEnum_prime_0
var fromEnum_prime_1 gopurs_runtime.Value = fromEnum_prime_1_loop
_ = fromEnum_prime_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(toEnum_prime_0, gopurs_runtime.Apply2(Get_add__560788792(), gopurs_runtime.Apply(fromEnum_prime_1, a_2), gopurs_runtime.Int(1))))
}

func Call_fromEnum__1637084359(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_toEnum__3317293286(dict_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_toEnumWithDefaults__3558602759(dictBoundedEnum_0_loop *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictBoundedEnum_0 *pkg_Data_Enum.Constructor_BoundedEnum[gopurs_runtime.Value] = dictBoundedEnum_0_loop
_ = dictBoundedEnum_0
bottom2_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictBoundedEnum_0.V0, gopurs_runtime.Value{}), "bottom")
_ = bottom2_1_0
return gopurs_runtime.Func(func(low_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(high_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
v_5_1 := gopurs_runtime.Apply(dictBoundedEnum_0.V4, x_4)
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
if (gopurs_runtime.Bool(Call_lessThan__4087042607(x_4, gopurs_runtime.Apply(dictBoundedEnum_0.V3, bottom2_1_0))).IntVal) != (0) {
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

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
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

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_fromJust__1577979644(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_fromJust__911089788(v_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__1646952192(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(dict_0.V1, pkg_Data_Semigroup.Get_semigroupString())
}

func Call_foldMap1__3342855683(dict_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



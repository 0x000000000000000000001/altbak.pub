package Data_Monoid

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Symbol "gopurs/output/Data.Symbol"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monoidUnit gopurs_runtime.Value
var once_monoidUnit sync.Once
func Get_monoidUnit() gopurs_runtime.Value {
	once_monoidUnit.Do(func() {
		cache_monoidUnit = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupUnit()
}), pkg_Data_Unit.Get_unit())
	})
	return cache_monoidUnit
}

var cache_monoidString gopurs_runtime.Value
var once_monoidString sync.Once
func Get_monoidString() gopurs_runtime.Value {
	once_monoidString.Do(func() {
		cache_monoidString = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupString()
}), gopurs_runtime.Str(""))
	})
	return cache_monoidString
}

var cache_monoidRecordNil gopurs_runtime.Value
var once_monoidRecordNil sync.Once
func Get_monoidRecordNil() gopurs_runtime.Value {
	once_monoidRecordNil.Do(func() {
		cache_monoidRecordNil = gopurs_runtime.RecordDict2("SemigroupRecord0", "memptyRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupRecordNil()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}))
	})
	return cache_monoidRecordNil
}

var cache_monoidOrdering gopurs_runtime.Value
var once_monoidOrdering sync.Once
func Get_monoidOrdering() gopurs_runtime.Value {
	once_monoidOrdering.Do(func() {
		cache_monoidOrdering = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ordering.Get_semigroupOrdering()
}), gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil})
	})
	return cache_monoidOrdering
}

var cache_monoidArray gopurs_runtime.Value
var once_monoidArray sync.Once
func Get_monoidArray() gopurs_runtime.Value {
	once_monoidArray.Do(func() {
		cache_monoidArray = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semigroup.Get_semigroupArray()
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
	})
	return cache_monoidArray
}

var cache_memptyRecord gopurs_runtime.Value
var once_memptyRecord sync.Once
func Get_memptyRecord() gopurs_runtime.Value {
	once_memptyRecord.Do(func() {
		cache_memptyRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_memptyRecord(gopurs_runtime.CoerceToStruct[Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_memptyRecord
}

var cache_monoidRecord gopurs_runtime.Value
var once_monoidRecord sync.Once
func Get_monoidRecord() gopurs_runtime.Value {
	once_monoidRecord.Do(func() {
		cache_monoidRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictMonoidRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidRecord(_dollar__unused_0_box, dictMonoidRecord_1_box)
})
	})
	return cache_monoidRecord
}

var cache_mempty gopurs_runtime.Value
var once_mempty sync.Once
func Get_mempty() gopurs_runtime.Value {
	once_mempty.Do(func() {
		cache_mempty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty(dict_0_box)
})
	})
	return cache_mempty
}

var cache_monoidFn gopurs_runtime.Value
var once_monoidFn sync.Once
func Get_monoidFn() gopurs_runtime.Value {
	once_monoidFn.Do(func() {
		cache_monoidFn = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidFn(dictMonoid_0_box)
})
	})
	return cache_monoidFn
}

var cache_monoidRecordCons gopurs_runtime.Value
var once_monoidRecordCons sync.Once
func Get_monoidRecordCons() gopurs_runtime.Value {
	once_monoidRecordCons.Do(func() {
		cache_monoidRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidRecordCons(dictIsSymbol_0_box, dictMonoid_1_box)
})
	})
	return cache_monoidRecordCons
}

var cache_power gopurs_runtime.Value
var once_power sync.Once
func Get_power() gopurs_runtime.Value {
	once_power.Do(func() {
		cache_power = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_power(gopurs_runtime.CoerceToStruct[Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_power
}

var cache_guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		cache_guard = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_guard(gopurs_runtime.CoerceToStruct[Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_guard
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

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_memptyRecord__2391219712 gopurs_runtime.Value
var once_memptyRecord__2391219712 sync.Once
func Get_memptyRecord__2391219712() gopurs_runtime.Value {
	once_memptyRecord__2391219712.Do(func() {
		cache_memptyRecord__2391219712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_memptyRecord__2391219712(gopurs_runtime.CoerceToStruct[Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_memptyRecord__2391219712
}

var cache_memptyRecord__4056812038 gopurs_runtime.Value
var once_memptyRecord__4056812038 sync.Once
func Get_memptyRecord__4056812038() gopurs_runtime.Value {
	once_memptyRecord__4056812038.Do(func() {
		cache_memptyRecord__4056812038 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_memptyRecord__4056812038(gopurs_runtime.CoerceToStruct[Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_memptyRecord__4056812038
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

var cache_semigroupArray__1728406699 gopurs_runtime.Value
var once_semigroupArray__1728406699 sync.Once
func Get_semigroupArray__1728406699() gopurs_runtime.Value {
	once_semigroupArray__1728406699.Do(func() {
		cache_semigroupArray__1728406699 = gopurs_runtime.RecordDict1("append", pkg_Data_Semigroup.Get_concatArray())
	})
	return cache_semigroupArray__1728406699
}

var cache_semigroupRecordNil__2406047365 gopurs_runtime.Value
var once_semigroupRecordNil__2406047365 sync.Once
func Get_semigroupRecordNil__2406047365() gopurs_runtime.Value {
	once_semigroupRecordNil__2406047365.Do(func() {
		cache_semigroupRecordNil__2406047365 = gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_semigroupRecordNil__2406047365
}

var cache_reflectSymbol__3416619207 gopurs_runtime.Value
var once_reflectSymbol__3416619207 sync.Once
func Get_reflectSymbol__3416619207() gopurs_runtime.Value {
	once_reflectSymbol__3416619207.Do(func() {
		cache_reflectSymbol__3416619207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__3416619207(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__3416619207
}

var cache_reflectSymbol__1166932993 gopurs_runtime.Value
var once_reflectSymbol__1166932993 sync.Once
func Get_reflectSymbol__1166932993() gopurs_runtime.Value {
	once_reflectSymbol__1166932993.Do(func() {
		cache_reflectSymbol__1166932993 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__1166932993(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__1166932993
}

type Constructor_MonoidRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2415148183] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "SemigroupRecord0": return c.V0
		case "memptyRecord": return c.V1
		default: panic("Key not found in dictionary Constructor_MonoidRecord: " + key)
		}
	}
}


type Constructor_Monoid[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_m
}


func init() {
	gopurs_runtime.StructGetters[1722653594] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Monoid[gopurs_runtime.Value])(ptr)
		switch key {
		case "Semigroup0": return c.V0
		case "mempty": return c.V1
		default: panic("Key not found in dictionary Constructor_Monoid: " + key)
		}
	}
}


func Call_memptyRecord(dict_0_loop *Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_monoidRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictMonoidRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictMonoidRecord_1 gopurs_runtime.Value = dictMonoidRecord_1_loop
_ = dictMonoidRecord_1
semigroupRecord1_2_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "SemigroupRecord0"), gopurs_runtime.Value{}), "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = semigroupRecord1_2_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRecord1_2_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_mempty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_monoidFn(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semigroupFn_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}))
_ = semigroupFn_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
}))
}

func Call_monoidRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
mempty1_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty1_2_0
Semigroup0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_3_1
return gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoidRecord_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "SemigroupRecord0"), gopurs_runtime.Value{})
_ = __local_var_6_3
semigroupRecordCons1_6_2 := gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
key_10_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_10_4
get_11_5 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_10_4)
_ = get_11_5
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_10_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Semigroup0_3_1, "append"), gopurs_runtime.Apply(get_11_5, ra_8), gopurs_runtime.Apply(get_11_5, rb_9)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_3, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9))
})
})
}))
_ = semigroupRecordCons1_6_2
return gopurs_runtime.RecordDict2("SemigroupRecord0", "memptyRecord", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRecordCons1_6_2
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), mempty1_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}))
})
})
}

func Call_power(dictMonoid_0_loop *Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
mempty1_1_0 := dictMonoid_0.V1
_ = mempty1_1_0
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_0.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_0 gopurs_runtime.Value
_ = go__go_4_2_0
go__go_4_2_0 = gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(Call_lessThanOrEq__4087042607(p_5, gopurs_runtime.Int(0))).IntVal) != (0) {
__t5 = mempty1_1_0
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), p_5, gopurs_runtime.Int(1)).IntVal) != (0) {
__t5 = x_3
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_eq__2843686287(), gopurs_runtime.Apply2(Get_mod__2185172824(), p_5, gopurs_runtime.Int(2)), gopurs_runtime.Int(0)).IntVal) != (0) {
x_prime_6_4 := gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Apply2(Get_div__2185172824(), p_5, gopurs_runtime.Int(2)))
_ = x_prime_6_4
__t5 = gopurs_runtime.Apply2(Semigroup0_2_1.V0, x_prime_6_4, x_prime_6_4)
goto end_branch_5
} else {

}
}
{
x_prime_6_3 := gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Apply2(Get_div__2185172824(), p_5, gopurs_runtime.Int(2)))
_ = x_prime_6_3
__t5 = gopurs_runtime.Apply2(Semigroup0_2_1.V0, x_prime_6_3, gopurs_runtime.Apply2(Semigroup0_2_1.V0, x_prime_6_3, x_3))
}
end_branch_5:
return __t5
})
return go__go_4_2_0
})
}

func Call_guard(dictMonoid_0_loop *Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
mempty1_1_0 := dictMonoid_0.V1
_ = mempty1_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (v_2.IntVal) != (0) {
__t1 = v1_3
goto end_branch_1
} else {

}
}
{
__t1 = mempty1_1_0
}
end_branch_1:
return __t1
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

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_memptyRecord__2391219712(dict_0_loop *Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_memptyRecord__4056812038(dict_0_loop *Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__3416619207(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__1166932993(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



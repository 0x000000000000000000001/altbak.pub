package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Monoid_semigroupRecord gopurs_runtime.Value
var once_Data_Monoid_semigroupRecord sync.Once
func Get_Data_Monoid_semigroupRecord() gopurs_runtime.Value {
	once_Data_Monoid_semigroupRecord.Do(func() {
		cache_Data_Monoid_semigroupRecord = gopurs_runtime.Func(func(dictSemigroupRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_semigroupRecord(dictSemigroupRecord_0_box)
})
	})
	return cache_Data_Monoid_semigroupRecord
}

var cache_Data_Monoid_MonoidRecord_dollarDict gopurs_runtime.Value
var once_Data_Monoid_MonoidRecord_dollarDict sync.Once
func Get_Data_Monoid_MonoidRecord_dollarDict() gopurs_runtime.Value {
	once_Data_Monoid_MonoidRecord_dollarDict.Do(func() {
		cache_Data_Monoid_MonoidRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_MonoidRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_Monoid_MonoidRecord_dollarDict
}

var cache_Data_Monoid_Monoid_dollarDict gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollarDict sync.Once
func Get_Data_Monoid_Monoid_dollarDict() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollarDict.Do(func() {
		cache_Data_Monoid_Monoid_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_Monoid_dollarDict(x_0_box)
})
	})
	return cache_Data_Monoid_Monoid_dollarDict
}

var cache_Data_Monoid_monoidUnit gopurs_runtime.Value
var once_Data_Monoid_monoidUnit sync.Once
func Get_Data_Monoid_monoidUnit() gopurs_runtime.Value {
	once_Data_Monoid_monoidUnit.Do(func() {
		cache_Data_Monoid_monoidUnit = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_semigroupUnit()))}
}), Get_Data_Unit_unit()})}
	})
	return cache_Data_Monoid_monoidUnit
}

var cache_Data_Monoid_monoidString gopurs_runtime.Value
var once_Data_Monoid_monoidString sync.Once
func Get_Data_Monoid_monoidString() gopurs_runtime.Value {
	once_Data_Monoid_monoidString.Do(func() {
		cache_Data_Monoid_monoidString = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_semigroupString()))}
}), gopurs_runtime.Str("")})}
	})
	return cache_Data_Monoid_monoidString
}

var cache_Data_Monoid_monoidRecordNil gopurs_runtime.Value
var once_Data_Monoid_monoidRecordNil sync.Once
func Get_Data_Monoid_monoidRecordNil() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecordNil.Do(func() {
		cache_Data_Monoid_monoidRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_MonoidRecord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord](Get_Data_Semigroup_semigroupRecordNil()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})})}
	})
	return cache_Data_Monoid_monoidRecordNil
}

var cache_Data_Monoid_monoidOrdering gopurs_runtime.Value
var once_Data_Monoid_monoidOrdering sync.Once
func Get_Data_Monoid_monoidOrdering() gopurs_runtime.Value {
	once_Data_Monoid_monoidOrdering.Do(func() {
		cache_Data_Monoid_monoidOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Ordering_semigroupOrdering()))}
}), gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}})}
	})
	return cache_Data_Monoid_monoidOrdering
}

var cache_Data_Monoid_monoidArray gopurs_runtime.Value
var once_Data_Monoid_monoidArray sync.Once
func Get_Data_Monoid_monoidArray() gopurs_runtime.Value {
	once_Data_Monoid_monoidArray.Do(func() {
		cache_Data_Monoid_monoidArray = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_semigroupArray()))}
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
	})
	return cache_Data_Monoid_monoidArray
}

var cache_Data_Monoid_memptyRecord gopurs_runtime.Value
var once_Data_Monoid_memptyRecord sync.Once
func Get_Data_Monoid_memptyRecord() gopurs_runtime.Value {
	once_Data_Monoid_memptyRecord.Do(func() {
		cache_Data_Monoid_memptyRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_memptyRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_MonoidRecord](dict_0_box))
})
	})
	return cache_Data_Monoid_memptyRecord
}

var cache_Data_Monoid_monoidRecord gopurs_runtime.Value
var once_Data_Monoid_monoidRecord sync.Once
func Get_Data_Monoid_monoidRecord() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecord.Do(func() {
		cache_Data_Monoid_monoidRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictMonoidRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_monoidRecord(_dollar__unused_0_box, dictMonoidRecord_1_box)
})
	})
	return cache_Data_Monoid_monoidRecord
}

var cache_Data_Monoid_mempty gopurs_runtime.Value
var once_Data_Monoid_mempty sync.Once
func Get_Data_Monoid_mempty() gopurs_runtime.Value {
	once_Data_Monoid_mempty.Do(func() {
		cache_Data_Monoid_mempty = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_mempty(dict_0_box)
})
	})
	return cache_Data_Monoid_mempty
}

var cache_Data_Monoid_monoidFn gopurs_runtime.Value
var once_Data_Monoid_monoidFn sync.Once
func Get_Data_Monoid_monoidFn() gopurs_runtime.Value {
	once_Data_Monoid_monoidFn.Do(func() {
		cache_Data_Monoid_monoidFn = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_monoidFn(dictMonoid_0_box)
})
	})
	return cache_Data_Monoid_monoidFn
}

var cache_Data_Monoid_monoidRecordCons gopurs_runtime.Value
var once_Data_Monoid_monoidRecordCons sync.Once
func Get_Data_Monoid_monoidRecordCons() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecordCons.Do(func() {
		cache_Data_Monoid_monoidRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_monoidRecordCons(dictIsSymbol_0_box, dictMonoid_1_box)
})
	})
	return cache_Data_Monoid_monoidRecordCons
}

var cache_Data_Monoid_power gopurs_runtime.Value
var once_Data_Monoid_power sync.Once
func Get_Data_Monoid_power() gopurs_runtime.Value {
	once_Data_Monoid_power.Do(func() {
		cache_Data_Monoid_power = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_power(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
})
	})
	return cache_Data_Monoid_power
}

var cache_Data_Monoid_guard gopurs_runtime.Value
var once_Data_Monoid_guard sync.Once
func Get_Data_Monoid_guard() gopurs_runtime.Value {
	once_Data_Monoid_guard.Do(func() {
		cache_Data_Monoid_guard = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_guard(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
})
	})
	return cache_Data_Monoid_guard
}

var cache_Data_Monoid_mempty__2312420373 gopurs_runtime.Value
var once_Data_Monoid_mempty__2312420373 sync.Once
func Get_Data_Monoid_mempty__2312420373() gopurs_runtime.Value {
	once_Data_Monoid_mempty__2312420373.Do(func() {
		cache_Data_Monoid_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_mempty__2312420373(dict_0_box)
})
	})
	return cache_Data_Monoid_mempty__2312420373
}

var cache_Data_Monoid_memptyRecord__2391219712 gopurs_runtime.Value
var once_Data_Monoid_memptyRecord__2391219712 sync.Once
func Get_Data_Monoid_memptyRecord__2391219712() gopurs_runtime.Value {
	once_Data_Monoid_memptyRecord__2391219712.Do(func() {
		cache_Data_Monoid_memptyRecord__2391219712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_memptyRecord__2391219712(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_MonoidRecord](dict_0_box))
})
	})
	return cache_Data_Monoid_memptyRecord__2391219712
}

var cache_Data_Monoid_memptyRecord__4056812038 gopurs_runtime.Value
var once_Data_Monoid_memptyRecord__4056812038 sync.Once
func Get_Data_Monoid_memptyRecord__4056812038() gopurs_runtime.Value {
	once_Data_Monoid_memptyRecord__4056812038.Do(func() {
		cache_Data_Monoid_memptyRecord__4056812038 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_memptyRecord__4056812038(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_MonoidRecord](dict_0_box))
})
	})
	return cache_Data_Monoid_memptyRecord__4056812038
}

type Constructor_Data_Monoid_MonoidRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2415148183] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_MonoidRecord)(ptr)
		_ = c
		switch key {
		case "SemigroupRecord0": return gopurs_runtime.Box(c.V0)
		case "memptyRecord": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Monoid_MonoidRecord: " + key)
		}
	}
}


type Constructor_Data_Monoid_Monoid struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1722653594] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_Monoid)(ptr)
		_ = c
		switch key {
		case "Semigroup0": return gopurs_runtime.Box(c.V0)
		case "mempty": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Monoid_Monoid: " + key)
		}
	}
}


func Call_Data_Monoid_semigroupRecord(dictSemigroupRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupRecord_0 gopurs_runtime.Value = dictSemigroupRecord_0_loop
_ = dictSemigroupRecord_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroupRecord_0, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})}
}

func Call_Data_Monoid_MonoidRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_Monoid_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Monoid_memptyRecord(dict_0_loop *Constructor_Data_Monoid_MonoidRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Monoid_MonoidRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Monoid_monoidRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictMonoidRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictMonoidRecord_1 gopurs_runtime.Value = dictMonoidRecord_1_loop
_ = dictMonoidRecord_1
// TAST (Let): semigroupRecord1_2_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupRecord1_2_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "SemigroupRecord0"), gopurs_runtime.Value{}), "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}
_ = semigroupRecord1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupRecord1_2_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})}
}

func Call_Data_Monoid_mempty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_Data_Monoid_monoidFn(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupFn_1_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupFn_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
})))
_ = semigroupFn_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
})})}
}

func Call_Data_Monoid_monoidRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): mempty1_2_0 -> gopurs_runtime.Value
mempty1_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty1_2_0
// TAST (Let): Semigroup0_3_1 -> gopurs_runtime.Value
Semigroup0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_3_1
return gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictMonoidRecord_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_3 -> gopurs_runtime.Value
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "SemigroupRecord0"), gopurs_runtime.Value{})
_ = __local_var_6_3
// TAST (Let): semigroupRecordCons1_6_2 -> *Constructor_Data_Semigroup_SemigroupRecord
semigroupRecordCons1_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord](gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_10_4 -> gopurs_runtime.Value
key_10_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_10_4
// TAST (Let): get_11_5 -> gopurs_runtime.Value
get_11_5 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_10_4.StrVal()))
_ = get_11_5
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_10_4.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Semigroup0_3_1, "append"), gopurs_runtime.Apply(get_11_5, ra_8), gopurs_runtime.Apply(get_11_5, rb_9)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_6_3, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9))
})
})
})))
_ = semigroupRecordCons1_6_2
return gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_MonoidRecord{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer(semigroupRecordCons1_6_2)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), mempty1_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
})})}
})
})
}

func Call_Data_Monoid_power(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty1_1_0 -> gopurs_runtime.Value
mempty1_1_0 := gopurs_runtime.Box(dictMonoid_0.V1)
_ = mempty1_1_0
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_0 gopurs_runtime.Value
_ = go__go_4_2_0
go__go_4_2_0 = gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
var __t4 bool
{
if (p_5.IntVal) > (0) {
__t4 = false
goto end_branch_4
} else {

}
}
{
__t4 = true
}
end_branch_4:
if __t4 {
__t6 = mempty1_1_0
goto end_branch_6
} else {

}
}
{
if (p_5.IntVal) == (1) {
__t6 = x_3
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(p_5.IntVal), gopurs_runtime.Int(2)).IntVal) == (0) {
// TAST (Let): x_prime_6_5 -> gopurs_runtime.Value
x_prime_6_5 := gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Int((p_5.IntVal) / (2)))
_ = x_prime_6_5
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), x_prime_6_5, x_prime_6_5)
goto end_branch_6
} else {

}
}
{
// TAST (Let): x_prime_6_3 -> gopurs_runtime.Value
x_prime_6_3 := gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Int((p_5.IntVal) / (2)))
_ = x_prime_6_3
__t6 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), x_prime_6_3, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), x_prime_6_3, x_3))
}
end_branch_6:
return __t6
})
return go__go_4_2_0
})
}

func Call_Data_Monoid_guard(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty1_1_0 -> gopurs_runtime.Value
mempty1_1_0 := gopurs_runtime.Box(dictMonoid_0.V1)
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

func Call_Data_Monoid_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_Data_Monoid_memptyRecord__2391219712(dict_0_loop *Constructor_Data_Monoid_MonoidRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Monoid_MonoidRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Monoid_memptyRecord__4056812038(dict_0_loop *Constructor_Data_Monoid_MonoidRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Monoid_MonoidRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



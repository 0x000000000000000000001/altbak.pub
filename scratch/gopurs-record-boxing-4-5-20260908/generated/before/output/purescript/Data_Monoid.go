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

var cache_Data_Monoid_MonoidRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Monoid_MonoidRecord_dollar_Dict sync.Once
func Get_Data_Monoid_MonoidRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Monoid_MonoidRecord_dollar_Dict.Do(func() {
		cache_Data_Monoid_MonoidRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_MonoidRecord_dollar_Dict(func() struct{
	SemigroupRecord0 gopurs_runtime.Value
	memptyRecord gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	SemigroupRecord0 gopurs_runtime.Value
	memptyRecord gopurs_runtime.Value
}{}
					clone.SemigroupRecord0 = gopurs_runtime.RecordGet(orig, "SemigroupRecord0")
					clone.memptyRecord = gopurs_runtime.RecordGet(orig, "memptyRecord")
					return clone
				}()))}
})
	})
	return cache_Data_Monoid_MonoidRecord_dollar_Dict
}

var cache_Data_Monoid_Monoid_dollar_Dict gopurs_runtime.Value
var once_Data_Monoid_Monoid_dollar_Dict sync.Once
func Get_Data_Monoid_Monoid_dollar_Dict() gopurs_runtime.Value {
	once_Data_Monoid_Monoid_dollar_Dict.Do(func() {
		cache_Data_Monoid_Monoid_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Call_Data_Monoid_Monoid_dollar_Dict(func() struct{
	Semigroup0 gopurs_runtime.Value
	mempty gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Semigroup0 gopurs_runtime.Value
	mempty gopurs_runtime.Value
}{}
					clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
					clone.mempty = gopurs_runtime.RecordGet(orig, "mempty")
					return clone
				}()))}
})
	})
	return cache_Data_Monoid_Monoid_dollar_Dict
}

var cache_Data_Monoid_monoidUnit gopurs_runtime.Value
var once_Data_Monoid_monoidUnit sync.Once
func Get_Data_Monoid_monoidUnit() gopurs_runtime.Value {
	once_Data_Monoid_monoidUnit.Do(func() {
		cache_Data_Monoid_monoidUnit = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupUnit()))}
}), Get_Data_Unit_unit()}))}
	})
	return cache_Data_Monoid_monoidUnit
}

var cache_Data_Monoid_monoidString gopurs_runtime.Value
var once_Data_Monoid_monoidString sync.Once
func Get_Data_Monoid_monoidString() gopurs_runtime.Value {
	once_Data_Monoid_monoidString.Do(func() {
		cache_Data_Monoid_monoidString = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_1950344881_1201789390((&Constructor_Data_Monoid_Monoid[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_443971153_4179793454(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[string]](Get_Data_Semigroup_semigroupString())))}
}), ""})))}
	})
	return cache_Data_Monoid_monoidString
}

var cache_Data_Monoid_monoidRecordNil gopurs_runtime.Value
var once_Data_Monoid_monoidRecordNil sync.Once
func Get_Data_Monoid_monoidRecordNil() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecordNil.Do(func() {
		cache_Data_Monoid_monoidRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Semigroup_semigroupRecordNil()))}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}()
})}))}
	})
	return cache_Data_Monoid_monoidRecordNil
}

var cache_Data_Monoid_monoidOrdering gopurs_runtime.Value
var once_Data_Monoid_monoidOrdering sync.Once
func Get_Data_Monoid_monoidOrdering() gopurs_runtime.Value {
	once_Data_Monoid_monoidOrdering.Do(func() {
		cache_Data_Monoid_monoidOrdering = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_2443930307_1201789390((&Constructor_Data_Monoid_Monoid[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_1625289059_4179793454(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[uint32]](Get_Data_Ordering_semigroupOrdering())))}
}), 902936544})))}
	})
	return cache_Data_Monoid_monoidOrdering
}

var cache_Data_Monoid_monoidArray gopurs_runtime.Value
var once_Data_Monoid_monoidArray sync.Once
func Get_Data_Monoid_monoidArray() gopurs_runtime.Value {
	once_Data_Monoid_monoidArray.Do(func() {
		cache_Data_Monoid_monoidArray = gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_67307368_1201789390((&Constructor_Data_Monoid_Monoid[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Monoid_986430664_4179793454(Rebox_Data_Monoid_4179793454_986430664(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupArray()))))}
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()})))}
	})
	return cache_Data_Monoid_monoidArray
}

var cache_Data_Monoid_memptyRecord gopurs_runtime.Value
var once_Data_Monoid_memptyRecord sync.Once
func Get_Data_Monoid_memptyRecord() gopurs_runtime.Value {
	once_Data_Monoid_memptyRecord.Do(func() {
		cache_Data_Monoid_memptyRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_memptyRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Monoid_memptyRecord
}

var cache_Data_Monoid_monoidRecord gopurs_runtime.Value
var once_Data_Monoid_monoidRecord sync.Once
func Get_Data_Monoid_monoidRecord() gopurs_runtime.Value {
	once_Data_Monoid_monoidRecord.Do(func() {
		cache_Data_Monoid_monoidRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictMonoidRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_monoidRecord(_dollar___unused_0_box, dictMonoidRecord_1_box)
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

var cache_Data_Monoid_mempty__1029283410 gopurs_runtime.Value
var once_Data_Monoid_mempty__1029283410 sync.Once
func Get_Data_Monoid_mempty__1029283410() gopurs_runtime.Value {
	once_Data_Monoid_mempty__1029283410.Do(func() {
		cache_Data_Monoid_mempty__1029283410 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_mempty__1029283410(__eta0_0_box)
})
	})
	return cache_Data_Monoid_mempty__1029283410
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
return Call_Data_Monoid_power(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_Data_Monoid_power
}

var cache_Data_Monoid_guard gopurs_runtime.Value
var once_Data_Monoid_guard sync.Once
func Get_Data_Monoid_guard() gopurs_runtime.Value {
	once_Data_Monoid_guard.Do(func() {
		cache_Data_Monoid_guard = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Monoid_guard(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_Data_Monoid_guard
}

type Constructor_Data_Monoid_MonoidRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2415148183] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_MonoidRecord[any, any, any])(ptr)
		_ = c
		switch key {
		case "SemigroupRecord0": return gopurs_runtime.Box(c.V0)
		case "memptyRecord": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Monoid_MonoidRecord: " + key)
		}
	}
}


type Constructor_Data_Monoid_Monoid[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_m
}


func init() {
	gopurs_runtime.StructGetters[1722653594] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Monoid_Monoid[any])(ptr)
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
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroupRecord_0, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Monoid_MonoidRecord_dollar_Dict(x_0_loop struct{
	SemigroupRecord0 gopurs_runtime.Value
	memptyRecord gopurs_runtime.Value
}) *Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	SemigroupRecord0 gopurs_runtime.Value
	memptyRecord gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"SemigroupRecord0", "memptyRecord"}, []gopurs_runtime.Value{orig.SemigroupRecord0, orig.memptyRecord})
				}())
}

func Call_Data_Monoid_Monoid_dollar_Dict(x_0_loop struct{
	Semigroup0 gopurs_runtime.Value
	mempty gopurs_runtime.Value
}) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
var x_0 struct{
	Semigroup0 gopurs_runtime.Value
	mempty gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Semigroup0", "mempty"}, []gopurs_runtime.Value{orig.Semigroup0, orig.mempty})
				}())
}

func Call_Data_Monoid_memptyRecord(dict_0_loop *Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Monoid_monoidRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictMonoidRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictMonoidRecord_1 gopurs_runtime.Value = dictMonoidRecord_1_loop
_ = dictMonoidRecord_1
// TAST (Let): semigroupRecord1_2_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Record (Row [] (TypeVar row)))])
semigroupRecord1_2_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "SemigroupRecord0"), gopurs_runtime.Value{}), "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = semigroupRecord1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupRecord1_2_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_1, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Monoid_mempty(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_Data_Monoid_mempty__1029283410(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
mempty__1029283410:
for {
if false { continue mempty__1029283410 }
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Effect_Aff_monoidCanceler()).V1), __eta0_0)
}
}

func Call_Data_Monoid_monoidFn(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupFn_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar a)] (TypeVar b))])
semigroupFn_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})})
_ = semigroupFn_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
})}))}
}

func Call_Data_Monoid_monoidRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictMonoid_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictMonoid_1 gopurs_runtime.Value = dictMonoid_1_loop
_ = dictMonoid_1
// TAST (Let): mempty1_2_0 shape=Other bindingType=(TypeVar focus)
mempty1_2_0 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = mempty1_2_0
// TAST (Let): Semigroup0_3_1 shape=App(Other) bindingType=Any
Semigroup0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_3_1
return gopurs_runtime.Func2(func(_dollar___unused_4 gopurs_runtime.Value, dictMonoidRecord_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupRecordCons1_6_2 shape=App(Var) bindingType=(TypeApp (ADT ["Data","Semigroup","SemigroupRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)]), (TypeVar row), (TypeVar subrow)])
semigroupRecordCons1_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Get_Data_Semigroup_semigroupRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "SemigroupRecord0"), gopurs_runtime.Value{}), Semigroup0_3_1))
_ = semigroupRecordCons1_6_2
return gopurs_runtime.Value{Type: 9, IntVal: 2415148183, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_MonoidRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3847494007, UnsafePtr: unsafe.Pointer(semigroupRecordCons1_6_2)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), mempty1_2_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoidRecord_5, "memptyRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
})}))}
})
}

func Call_Data_Monoid_power(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty1_1_0 shape=Other bindingType=(TypeVar m)
mempty1_1_0 := gopurs_runtime.Box(dictMonoid_0.V1)
_ = mempty1_1_0
// TAST (Let): Semigroup0_2_1 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_0 gopurs_runtime.Value
_ = go__go_4_2_0
// FALLBACK TCO: isLoop=false len=1
go__go_4_2_0 = gopurs_runtime.Func(func(p_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (p_5.IntVal) <= (int64(0)) {
__t5 = mempty1_1_0
goto end_branch_5
} else {

}
}
{
if (p_5.IntVal) == (int64(1)) {
__t5 = x_3
goto end_branch_5
} else {

}
}
{
if ((p_5.IntVal) % (int64(2))) == (int64(0)) {
// TAST (Let): x_prime__6_4 shape=App(Other) bindingType=(TypeVar m)
x_prime__6_4 := gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Int((p_5.IntVal) / (int64(2))))
_ = x_prime__6_4
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), x_prime__6_4, x_prime__6_4)
goto end_branch_5
} else {

}
}
{
// TAST (Let): x_prime__6_3 shape=App(Other) bindingType=(TypeVar m)
x_prime__6_3 := gopurs_runtime.Apply(go__go_4_2_0, gopurs_runtime.Int((p_5.IntVal) / (int64(2))))
_ = x_prime__6_3
__t5 = gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), x_prime__6_3, gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), x_prime__6_3, x_3))
}
end_branch_5:
return __t5
})
return go__go_4_2_0
})
}

func Call_Data_Monoid_guard(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty1_1_0 shape=Other bindingType=(TypeVar m)
mempty1_1_0 := gopurs_runtime.Box(dictMonoid_0.V1)
_ = mempty1_1_0
return gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
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
}

func Rebox_Data_Monoid_1625289059_4179793454(in *Constructor_Data_Semigroup_Semigroup[uint32]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Monoid_1950344881_1201789390(in *Constructor_Data_Monoid_Monoid[string]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Data_Monoid_2443930307_1201789390(in *Constructor_Data_Monoid_Monoid[uint32]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
	return out
}

func Rebox_Data_Monoid_4179793454_986430664(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Monoid_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Monoid_67307368_1201789390(in *Constructor_Data_Monoid_Monoid[[]gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Array(in.V1)
	return out
}

func Rebox_Data_Monoid_986430664_4179793454(in *Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}



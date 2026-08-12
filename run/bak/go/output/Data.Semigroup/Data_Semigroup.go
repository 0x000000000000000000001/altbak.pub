package Data_Semigroup

import (
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Void "gopurs/output/Data.Void"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_semigroupVoid gopurs_runtime.Value
var once_semigroupVoid sync.Once
func Get_semigroupVoid() gopurs_runtime.Value {
	once_semigroupVoid.Do(func() {
		cache_semigroupVoid = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Void.Get_absurd__gopurs_runtime_Value_331654555()
}))
	})
	return cache_semigroupVoid
}

var cache_semigroupUnit gopurs_runtime.Value
var once_semigroupUnit sync.Once
func Get_semigroupUnit() gopurs_runtime.Value {
	once_semigroupUnit.Do(func() {
		cache_semigroupUnit = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}))
	})
	return cache_semigroupUnit
}

var cache_semigroupString gopurs_runtime.Value
var once_semigroupString sync.Once
func Get_semigroupString() gopurs_runtime.Value {
	once_semigroupString.Do(func() {
		cache_semigroupString = gopurs_runtime.RecordDict1("append", Get_concatString())
	})
	return cache_semigroupString
}

var cache_semigroupRecordNil gopurs_runtime.Value
var once_semigroupRecordNil sync.Once
func Get_semigroupRecordNil() gopurs_runtime.Value {
	once_semigroupRecordNil.Do(func() {
		cache_semigroupRecordNil = gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_semigroupRecordNil
}

var cache_semigroupRecordNil__gopurs_runtime_Value_2406047365 gopurs_runtime.Value
var once_semigroupRecordNil__gopurs_runtime_Value_2406047365 sync.Once
func Get_semigroupRecordNil__gopurs_runtime_Value_2406047365() gopurs_runtime.Value {
	once_semigroupRecordNil__gopurs_runtime_Value_2406047365.Do(func() {
		cache_semigroupRecordNil__gopurs_runtime_Value_2406047365 = gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_semigroupRecordNil__gopurs_runtime_Value_2406047365
}

var cache_semigroupProxy gopurs_runtime.Value
var once_semigroupProxy sync.Once
func Get_semigroupProxy() gopurs_runtime.Value {
	once_semigroupProxy.Do(func() {
		cache_semigroupProxy = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_semigroupProxy
}

var cache_semigroupArray gopurs_runtime.Value
var once_semigroupArray sync.Once
func Get_semigroupArray() gopurs_runtime.Value {
	once_semigroupArray.Do(func() {
		cache_semigroupArray = gopurs_runtime.RecordDict1("append", Get_concatArray())
	})
	return cache_semigroupArray
}

var cache_semigroupArray__ptrConstructor_Semigroup_arrstring__4207347319 gopurs_runtime.Value
var once_semigroupArray__ptrConstructor_Semigroup_arrstring__4207347319 sync.Once
func Get_semigroupArray__ptrConstructor_Semigroup_arrstring__4207347319() gopurs_runtime.Value {
	once_semigroupArray__ptrConstructor_Semigroup_arrstring__4207347319.Do(func() {
		cache_semigroupArray__ptrConstructor_Semigroup_arrstring__4207347319 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Semigroup[[]string]{1, Get_concatArray()})}
	})
	return cache_semigroupArray__ptrConstructor_Semigroup_arrstring__4207347319
}

var cache_semigroupArray__ptrConstructor_Semigroup_arrgopurs_runtime_Value__777842900 gopurs_runtime.Value
var once_semigroupArray__ptrConstructor_Semigroup_arrgopurs_runtime_Value__777842900 sync.Once
func Get_semigroupArray__ptrConstructor_Semigroup_arrgopurs_runtime_Value__777842900() gopurs_runtime.Value {
	once_semigroupArray__ptrConstructor_Semigroup_arrgopurs_runtime_Value__777842900.Do(func() {
		cache_semigroupArray__ptrConstructor_Semigroup_arrgopurs_runtime_Value__777842900 = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Semigroup[[]gopurs_runtime.Value]{1, Get_concatArray()})}
	})
	return cache_semigroupArray__ptrConstructor_Semigroup_arrgopurs_runtime_Value__777842900
}

var cache_semigroupArray__gopurs_runtime_Value_1728406699 gopurs_runtime.Value
var once_semigroupArray__gopurs_runtime_Value_1728406699 sync.Once
func Get_semigroupArray__gopurs_runtime_Value_1728406699() gopurs_runtime.Value {
	once_semigroupArray__gopurs_runtime_Value_1728406699.Do(func() {
		cache_semigroupArray__gopurs_runtime_Value_1728406699 = gopurs_runtime.RecordDict1("append", Get_concatArray())
	})
	return cache_semigroupArray__gopurs_runtime_Value_1728406699
}

var cache_appendRecord gopurs_runtime.Value
var once_appendRecord sync.Once
func Get_appendRecord() gopurs_runtime.Value {
	once_appendRecord.Do(func() {
		cache_appendRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_appendRecord(gopurs_runtime.CoerceToStruct[Constructor_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_appendRecord
}

var cache_appendRecord__gopurs_runtime_Value_1546996774 gopurs_runtime.Value
var once_appendRecord__gopurs_runtime_Value_1546996774 sync.Once
func Get_appendRecord__gopurs_runtime_Value_1546996774() gopurs_runtime.Value {
	once_appendRecord__gopurs_runtime_Value_1546996774.Do(func() {
		cache_appendRecord__gopurs_runtime_Value_1546996774 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_appendRecord__gopurs_runtime_Value_1546996774(gopurs_runtime.CoerceToStruct[Constructor_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_appendRecord__gopurs_runtime_Value_1546996774
}

var cache_semigroupRecord gopurs_runtime.Value
var once_semigroupRecord sync.Once
func Get_semigroupRecord() gopurs_runtime.Value {
	once_semigroupRecord.Do(func() {
		cache_semigroupRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictSemigroupRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupRecord(_dollar__unused_0_box, dictSemigroupRecord_1_box)
})
	})
	return cache_semigroupRecord
}

var cache_append gopurs_runtime.Value
var once_append sync.Once
func Get_append() gopurs_runtime.Value {
	once_append.Do(func() {
		cache_append = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append(gopurs_runtime.CoerceToStruct[Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append
}

var cache_append__gopurs_runtime_Value_1230318264 gopurs_runtime.Value
var once_append__gopurs_runtime_Value_1230318264 sync.Once
func Get_append__gopurs_runtime_Value_1230318264() gopurs_runtime.Value {
	once_append__gopurs_runtime_Value_1230318264.Do(func() {
		cache_append__gopurs_runtime_Value_1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__gopurs_runtime_Value_1230318264(gopurs_runtime.CoerceToStruct[Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__gopurs_runtime_Value_1230318264
}

var cache_semigroupFn gopurs_runtime.Value
var once_semigroupFn sync.Once
func Get_semigroupFn() gopurs_runtime.Value {
	once_semigroupFn.Do(func() {
		cache_semigroupFn = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupFn(dictSemigroup_0_box)
})
	})
	return cache_semigroupFn
}

var cache_semigroupRecordCons gopurs_runtime.Value
var once_semigroupRecordCons sync.Once
func Get_semigroupRecordCons() gopurs_runtime.Value {
	once_semigroupRecordCons.Do(func() {
		cache_semigroupRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictSemigroupRecord_2_box gopurs_runtime.Value, dictSemigroup_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictSemigroupRecord_2_box, dictSemigroup_3_box)
})
	})
	return cache_semigroupRecordCons
}

type Constructor_SemigroupRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3847494007] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "appendRecord": return c.V0
		default: panic("Key not found in dictionary Constructor_SemigroupRecord: " + key)
		}
	}
}


type Constructor_Semigroup[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2053112122] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Semigroup[gopurs_runtime.Value])(ptr)
		switch key {
		case "append": return c.V0
		default: panic("Key not found in dictionary Constructor_Semigroup: " + key)
		}
	}
}


func Call_appendRecord(dict_0_loop *Constructor_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_appendRecord__gopurs_runtime_Value_1546996774(dict_0_loop *Constructor_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemigroupRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_semigroupRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictSemigroupRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictSemigroupRecord_1 gopurs_runtime.Value = dictSemigroupRecord_1_loop
_ = dictSemigroupRecord_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroupRecord_1, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_append(dict_0_loop *Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__gopurs_runtime_Value_1230318264(dict_0_loop *Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_semigroupFn(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}))
}

func Call_semigroupRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictSemigroupRecord_2_loop gopurs_runtime.Value, dictSemigroup_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictSemigroupRecord_2 gopurs_runtime.Value = dictSemigroupRecord_2_loop
_ = dictSemigroupRecord_2
var dictSemigroup_3 gopurs_runtime.Value = dictSemigroup_3_loop
_ = dictSemigroup_3
return gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
key_7_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_7_0
get_8_1 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_7_0)
_ = get_8_1
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_7_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_3, "append"), gopurs_runtime.Apply(get_8_1, ra_5), gopurs_runtime.Apply(get_8_1, rb_6)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemigroupRecord_2, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6))
})
})
}))
}

func Get_concatArray() gopurs_runtime.Value {
	return _Gopurs_ConcatArray
}

func Get_concatString() gopurs_runtime.Value {
	return _Gopurs_ConcatString
}

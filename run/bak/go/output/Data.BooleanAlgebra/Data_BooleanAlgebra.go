package Data_BooleanAlgebra

import (
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_booleanAlgebraUnit gopurs_runtime.Value
var once_booleanAlgebraUnit sync.Once
func Get_booleanAlgebraUnit() gopurs_runtime.Value {
	once_booleanAlgebraUnit.Do(func() {
		cache_booleanAlgebraUnit = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraUnit()
}))
	})
	return cache_booleanAlgebraUnit
}

var cache_booleanAlgebraRecordNil gopurs_runtime.Value
var once_booleanAlgebraRecordNil sync.Once
func Get_booleanAlgebraRecordNil() gopurs_runtime.Value {
	once_booleanAlgebraRecordNil.Do(func() {
		cache_booleanAlgebraRecordNil = gopurs_runtime.RecordDict1("HeytingAlgebraRecord0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraRecordNil()
}))
	})
	return cache_booleanAlgebraRecordNil
}

var cache_booleanAlgebraRecordCons gopurs_runtime.Value
var once_booleanAlgebraRecordCons sync.Once
func Get_booleanAlgebraRecordCons() gopurs_runtime.Value {
	once_booleanAlgebraRecordCons.Do(func() {
		cache_booleanAlgebraRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictBooleanAlgebraRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_booleanAlgebraRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictBooleanAlgebraRecord_2_box)
})
	})
	return cache_booleanAlgebraRecordCons
}

var cache_booleanAlgebraRecord gopurs_runtime.Value
var once_booleanAlgebraRecord sync.Once
func Get_booleanAlgebraRecord() gopurs_runtime.Value {
	once_booleanAlgebraRecord.Do(func() {
		cache_booleanAlgebraRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictBooleanAlgebraRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_booleanAlgebraRecord(_dollar__unused_0_box, dictBooleanAlgebraRecord_1_box)
})
	})
	return cache_booleanAlgebraRecord
}

var cache_booleanAlgebraProxy gopurs_runtime.Value
var once_booleanAlgebraProxy sync.Once
func Get_booleanAlgebraProxy() gopurs_runtime.Value {
	once_booleanAlgebraProxy.Do(func() {
		cache_booleanAlgebraProxy = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraProxy()
}))
	})
	return cache_booleanAlgebraProxy
}

var cache_booleanAlgebraFn gopurs_runtime.Value
var once_booleanAlgebraFn sync.Once
func Get_booleanAlgebraFn() gopurs_runtime.Value {
	once_booleanAlgebraFn.Do(func() {
		cache_booleanAlgebraFn = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_booleanAlgebraFn(dictBooleanAlgebra_0_box)
})
	})
	return cache_booleanAlgebraFn
}

var cache_booleanAlgebraBoolean gopurs_runtime.Value
var once_booleanAlgebraBoolean sync.Once
func Get_booleanAlgebraBoolean() gopurs_runtime.Value {
	once_booleanAlgebraBoolean.Do(func() {
		cache_booleanAlgebraBoolean = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean()
}))
	})
	return cache_booleanAlgebraBoolean
}

type Constructor_BooleanAlgebraRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[700691287] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_BooleanAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "HeytingAlgebraRecord0": return c.V0
		default: panic("Key not found in dictionary Constructor_BooleanAlgebraRecord: " + key)
		}
	}
}


type Constructor_BooleanAlgebra[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3257204378] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_BooleanAlgebra[gopurs_runtime.Value])(ptr)
		switch key {
		case "HeytingAlgebra0": return c.V0
		default: panic("Key not found in dictionary Constructor_BooleanAlgebra: " + key)
		}
	}
}


func Call_booleanAlgebraRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictBooleanAlgebraRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictBooleanAlgebraRecord_2 gopurs_runtime.Value = dictBooleanAlgebraRecord_2_loop
_ = dictBooleanAlgebraRecord_2
heytingAlgebraRecordCons1_3_0 := gopurs_runtime.Apply3(pkg_Data_HeytingAlgebra.Get_heytingAlgebraRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebraRecord_2, "HeytingAlgebraRecord0"), gopurs_runtime.Value{}))
_ = heytingAlgebraRecordCons1_3_0
return gopurs_runtime.Func(func(dictBooleanAlgebra_4 gopurs_runtime.Value) gopurs_runtime.Value {
heytingAlgebraRecordCons2_5_1 := gopurs_runtime.Apply(heytingAlgebraRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_4, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraRecordCons2_5_1
return gopurs_runtime.RecordDict1("HeytingAlgebraRecord0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecordCons2_5_1
}))
})
}

func Call_booleanAlgebraRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictBooleanAlgebraRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictBooleanAlgebraRecord_1 gopurs_runtime.Value = dictBooleanAlgebraRecord_1_loop
_ = dictBooleanAlgebraRecord_1
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebraRecord_1, "HeytingAlgebraRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
heytingAlgebraRecord1_2_0 := gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "conjRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "disjRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "ffRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "impliesRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "notRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "ttRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})})
_ = heytingAlgebraRecord1_2_0
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecord1_2_0
}))
}

func Call_booleanAlgebraFn(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{})
_ = __local_var_1_1
heytingAlgebraFunction_1_0 := gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "conj"), gopurs_runtime.Apply(f_2, a_4), gopurs_runtime.Apply(g_3, a_4))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "disj"), gopurs_runtime.Apply(f_2, a_4), gopurs_runtime.Apply(g_3, a_4))
})
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "ff")
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "implies"), gopurs_runtime.Apply(f_2, a_4), gopurs_runtime.Apply(g_3, a_4))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "not"), gopurs_runtime.Apply(f_2, a_3))
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "tt")
})})
_ = heytingAlgebraFunction_1_0
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraFunction_1_0
}))
}



package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_BooleanAlgebra_heytingAlgebraRecord gopurs_runtime.Value
var once_Data_BooleanAlgebra_heytingAlgebraRecord sync.Once
func Get_Data_BooleanAlgebra_heytingAlgebraRecord() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_heytingAlgebraRecord.Do(func() {
		cache_Data_BooleanAlgebra_heytingAlgebraRecord = gopurs_runtime.Apply(Get_Data_HeytingAlgebra_heytingAlgebraRecord(), gopurs_runtime.Value{})
	})
	return cache_Data_BooleanAlgebra_heytingAlgebraRecord
}

var cache_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict gopurs_runtime.Value
var once_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict sync.Once
func Get_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict.Do(func() {
		cache_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict
}

var cache_Data_BooleanAlgebra_BooleanAlgebra_dollarDict gopurs_runtime.Value
var once_Data_BooleanAlgebra_BooleanAlgebra_dollarDict sync.Once
func Get_Data_BooleanAlgebra_BooleanAlgebra_dollarDict() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_BooleanAlgebra_dollarDict.Do(func() {
		cache_Data_BooleanAlgebra_BooleanAlgebra_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_BooleanAlgebra_BooleanAlgebra_dollarDict(x_0_box)
})
	})
	return cache_Data_BooleanAlgebra_BooleanAlgebra_dollarDict
}

var cache_Data_BooleanAlgebra_booleanAlgebraUnit gopurs_runtime.Value
var once_Data_BooleanAlgebra_booleanAlgebraUnit sync.Once
func Get_Data_BooleanAlgebra_booleanAlgebraUnit() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_booleanAlgebraUnit.Do(func() {
		cache_Data_BooleanAlgebra_booleanAlgebraUnit = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_HeytingAlgebra_heytingAlgebraUnit()
}))
	})
	return cache_Data_BooleanAlgebra_booleanAlgebraUnit
}

var cache_Data_BooleanAlgebra_booleanAlgebraRecordNil gopurs_runtime.Value
var once_Data_BooleanAlgebra_booleanAlgebraRecordNil sync.Once
func Get_Data_BooleanAlgebra_booleanAlgebraRecordNil() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_booleanAlgebraRecordNil.Do(func() {
		cache_Data_BooleanAlgebra_booleanAlgebraRecordNil = gopurs_runtime.RecordDict1("HeytingAlgebraRecord0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_HeytingAlgebra_heytingAlgebraRecordNil()
}))
	})
	return cache_Data_BooleanAlgebra_booleanAlgebraRecordNil
}

var cache_Data_BooleanAlgebra_booleanAlgebraRecordCons gopurs_runtime.Value
var once_Data_BooleanAlgebra_booleanAlgebraRecordCons sync.Once
func Get_Data_BooleanAlgebra_booleanAlgebraRecordCons() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_booleanAlgebraRecordCons.Do(func() {
		cache_Data_BooleanAlgebra_booleanAlgebraRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictBooleanAlgebraRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_BooleanAlgebra_booleanAlgebraRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictBooleanAlgebraRecord_2_box)
})
	})
	return cache_Data_BooleanAlgebra_booleanAlgebraRecordCons
}

var cache_Data_BooleanAlgebra_booleanAlgebraRecord gopurs_runtime.Value
var once_Data_BooleanAlgebra_booleanAlgebraRecord sync.Once
func Get_Data_BooleanAlgebra_booleanAlgebraRecord() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_booleanAlgebraRecord.Do(func() {
		cache_Data_BooleanAlgebra_booleanAlgebraRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictBooleanAlgebraRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_BooleanAlgebra_booleanAlgebraRecord(_dollar__unused_0_box, dictBooleanAlgebraRecord_1_box)
})
	})
	return cache_Data_BooleanAlgebra_booleanAlgebraRecord
}

var cache_Data_BooleanAlgebra_booleanAlgebraProxy gopurs_runtime.Value
var once_Data_BooleanAlgebra_booleanAlgebraProxy sync.Once
func Get_Data_BooleanAlgebra_booleanAlgebraProxy() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_booleanAlgebraProxy.Do(func() {
		cache_Data_BooleanAlgebra_booleanAlgebraProxy = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_HeytingAlgebra_heytingAlgebraProxy()
}))
	})
	return cache_Data_BooleanAlgebra_booleanAlgebraProxy
}

var cache_Data_BooleanAlgebra_booleanAlgebraFn gopurs_runtime.Value
var once_Data_BooleanAlgebra_booleanAlgebraFn sync.Once
func Get_Data_BooleanAlgebra_booleanAlgebraFn() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_booleanAlgebraFn.Do(func() {
		cache_Data_BooleanAlgebra_booleanAlgebraFn = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_BooleanAlgebra_booleanAlgebraFn(dictBooleanAlgebra_0_box)
})
	})
	return cache_Data_BooleanAlgebra_booleanAlgebraFn
}

var cache_Data_BooleanAlgebra_booleanAlgebraBoolean gopurs_runtime.Value
var once_Data_BooleanAlgebra_booleanAlgebraBoolean sync.Once
func Get_Data_BooleanAlgebra_booleanAlgebraBoolean() gopurs_runtime.Value {
	once_Data_BooleanAlgebra_booleanAlgebraBoolean.Do(func() {
		cache_Data_BooleanAlgebra_booleanAlgebraBoolean = gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_HeytingAlgebra_heytingAlgebraBoolean()
}))
	})
	return cache_Data_BooleanAlgebra_booleanAlgebraBoolean
}

type Constructor_Data_BooleanAlgebra_BooleanAlgebraRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[700691287] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_BooleanAlgebra_BooleanAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "HeytingAlgebraRecord0": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_BooleanAlgebra_BooleanAlgebraRecord: " + key)
		}
	}
}


type Constructor_Data_BooleanAlgebra_BooleanAlgebra[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3257204378] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "HeytingAlgebra0": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_BooleanAlgebra_BooleanAlgebra: " + key)
		}
	}
}


func Call_Data_BooleanAlgebra_BooleanAlgebraRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_BooleanAlgebra_BooleanAlgebra_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_BooleanAlgebra_booleanAlgebraRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictBooleanAlgebraRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictBooleanAlgebraRecord_2 gopurs_runtime.Value = dictBooleanAlgebraRecord_2_loop
_ = dictBooleanAlgebraRecord_2
// TAST (Let): heytingAlgebraRecordCons1_3_0 -> gopurs_runtime.Value
heytingAlgebraRecordCons1_3_0 := gopurs_runtime.Apply3(Get_Data_HeytingAlgebra_heytingAlgebraRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebraRecord_2, "HeytingAlgebraRecord0"), gopurs_runtime.Value{}))
_ = heytingAlgebraRecordCons1_3_0
return gopurs_runtime.Func(func(dictBooleanAlgebra_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): heytingAlgebraRecordCons2_5_1 -> gopurs_runtime.Value
heytingAlgebraRecordCons2_5_1 := gopurs_runtime.Apply(heytingAlgebraRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_4, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraRecordCons2_5_1
return gopurs_runtime.RecordDict1("HeytingAlgebraRecord0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecordCons2_5_1
}))
})
}

func Call_Data_BooleanAlgebra_booleanAlgebraRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictBooleanAlgebraRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictBooleanAlgebraRecord_1 gopurs_runtime.Value = dictBooleanAlgebraRecord_1_loop
_ = dictBooleanAlgebraRecord_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebraRecord_1, "HeytingAlgebraRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): heytingAlgebraRecord1_2_0 -> gopurs_runtime.Value
heytingAlgebraRecord1_2_0 := gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "conjRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "disjRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "ffRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "impliesRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "notRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "ttRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = heytingAlgebraRecord1_2_0
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraRecord1_2_0
}))
}

func Call_Data_BooleanAlgebra_booleanAlgebraFn(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): heytingAlgebraFunction_1_0 -> gopurs_runtime.Value
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



package Data_Semiring

import (
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_zeroRecord gopurs_runtime.Value
var once_zeroRecord sync.Once
func Get_zeroRecord() gopurs_runtime.Value {
	once_zeroRecord.Do(func() {
		cache_zeroRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zeroRecord(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_zeroRecord
}

var cache_zeroRecord__gopurs_runtime_Value_2618241440 gopurs_runtime.Value
var once_zeroRecord__gopurs_runtime_Value_2618241440 sync.Once
func Get_zeroRecord__gopurs_runtime_Value_2618241440() gopurs_runtime.Value {
	once_zeroRecord__gopurs_runtime_Value_2618241440.Do(func() {
		cache_zeroRecord__gopurs_runtime_Value_2618241440 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zeroRecord__gopurs_runtime_Value_2618241440(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_zeroRecord__gopurs_runtime_Value_2618241440
}

var cache_zero gopurs_runtime.Value
var once_zero sync.Once
func Get_zero() gopurs_runtime.Value {
	once_zero.Do(func() {
		cache_zero = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero(dict_0_box)
})
	})
	return cache_zero
}

var cache_zero__gopurs_runtime_Value_1204848985 gopurs_runtime.Value
var once_zero__gopurs_runtime_Value_1204848985 sync.Once
func Get_zero__gopurs_runtime_Value_1204848985() gopurs_runtime.Value {
	once_zero__gopurs_runtime_Value_1204848985.Do(func() {
		cache_zero__gopurs_runtime_Value_1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__gopurs_runtime_Value_1204848985(dict_0_box)
})
	})
	return cache_zero__gopurs_runtime_Value_1204848985
}

var cache_semiringUnit gopurs_runtime.Value
var once_semiringUnit sync.Once
func Get_semiringUnit() gopurs_runtime.Value {
	once_semiringUnit.Do(func() {
		cache_semiringUnit = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), pkg_Data_Unit.Get_unit(), pkg_Data_Unit.Get_unit())
	})
	return cache_semiringUnit
}

var cache_semiringRecordNil gopurs_runtime.Value
var once_semiringRecordNil sync.Once
func Get_semiringRecordNil() gopurs_runtime.Value {
	once_semiringRecordNil.Do(func() {
		cache_semiringRecordNil = gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
}))
	})
	return cache_semiringRecordNil
}

var cache_semiringRecordNil__gopurs_runtime_Value_3131124892 gopurs_runtime.Value
var once_semiringRecordNil__gopurs_runtime_Value_3131124892 sync.Once
func Get_semiringRecordNil__gopurs_runtime_Value_3131124892() gopurs_runtime.Value {
	once_semiringRecordNil__gopurs_runtime_Value_3131124892.Do(func() {
		cache_semiringRecordNil__gopurs_runtime_Value_3131124892 = gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
}))
	})
	return cache_semiringRecordNil__gopurs_runtime_Value_3131124892
}

var cache_semiringProxy gopurs_runtime.Value
var once_semiringProxy sync.Once
func Get_semiringProxy() gopurs_runtime.Value {
	once_semiringProxy.Do(func() {
		cache_semiringProxy = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_semiringProxy
}

var cache_semiringProxy__gopurs_runtime_Value_786350887 gopurs_runtime.Value
var once_semiringProxy__gopurs_runtime_Value_786350887 sync.Once
func Get_semiringProxy__gopurs_runtime_Value_786350887() gopurs_runtime.Value {
	once_semiringProxy__gopurs_runtime_Value_786350887.Do(func() {
		cache_semiringProxy__gopurs_runtime_Value_786350887 = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})
	})
	return cache_semiringProxy__gopurs_runtime_Value_786350887
}

var cache_semiringNumber gopurs_runtime.Value
var once_semiringNumber sync.Once
func Get_semiringNumber() gopurs_runtime.Value {
	once_semiringNumber.Do(func() {
		cache_semiringNumber = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_numAdd(), Get_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
	})
	return cache_semiringNumber
}

var cache_semiringInt gopurs_runtime.Value
var once_semiringInt sync.Once
func Get_semiringInt() gopurs_runtime.Value {
	once_semiringInt.Do(func() {
		cache_semiringInt = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_intAdd(), Get_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0))
	})
	return cache_semiringInt
}

var cache_oneRecord gopurs_runtime.Value
var once_oneRecord sync.Once
func Get_oneRecord() gopurs_runtime.Value {
	once_oneRecord.Do(func() {
		cache_oneRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneRecord(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_oneRecord
}

var cache_oneRecord__gopurs_runtime_Value_2618241440 gopurs_runtime.Value
var once_oneRecord__gopurs_runtime_Value_2618241440 sync.Once
func Get_oneRecord__gopurs_runtime_Value_2618241440() gopurs_runtime.Value {
	once_oneRecord__gopurs_runtime_Value_2618241440.Do(func() {
		cache_oneRecord__gopurs_runtime_Value_2618241440 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneRecord__gopurs_runtime_Value_2618241440(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_oneRecord__gopurs_runtime_Value_2618241440
}

var cache_one gopurs_runtime.Value
var once_one sync.Once
func Get_one() gopurs_runtime.Value {
	once_one.Do(func() {
		cache_one = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one(dict_0_box)
})
	})
	return cache_one
}

var cache_one__gopurs_runtime_Value_1204848985 gopurs_runtime.Value
var once_one__gopurs_runtime_Value_1204848985 sync.Once
func Get_one__gopurs_runtime_Value_1204848985() gopurs_runtime.Value {
	once_one__gopurs_runtime_Value_1204848985.Do(func() {
		cache_one__gopurs_runtime_Value_1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__gopurs_runtime_Value_1204848985(dict_0_box)
})
	})
	return cache_one__gopurs_runtime_Value_1204848985
}

var cache_mulRecord gopurs_runtime.Value
var once_mulRecord sync.Once
func Get_mulRecord() gopurs_runtime.Value {
	once_mulRecord.Do(func() {
		cache_mulRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mulRecord(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mulRecord
}

var cache_mulRecord__gopurs_runtime_Value_2348817094 gopurs_runtime.Value
var once_mulRecord__gopurs_runtime_Value_2348817094 sync.Once
func Get_mulRecord__gopurs_runtime_Value_2348817094() gopurs_runtime.Value {
	once_mulRecord__gopurs_runtime_Value_2348817094.Do(func() {
		cache_mulRecord__gopurs_runtime_Value_2348817094 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mulRecord__gopurs_runtime_Value_2348817094(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mulRecord__gopurs_runtime_Value_2348817094
}

var cache_mul gopurs_runtime.Value
var once_mul sync.Once
func Get_mul() gopurs_runtime.Value {
	once_mul.Do(func() {
		cache_mul = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul(gopurs_runtime.CoerceToStruct[Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul
}

var cache_mul__gopurs_runtime_Value_1614463960 gopurs_runtime.Value
var once_mul__gopurs_runtime_Value_1614463960 sync.Once
func Get_mul__gopurs_runtime_Value_1614463960() gopurs_runtime.Value {
	once_mul__gopurs_runtime_Value_1614463960.Do(func() {
		cache_mul__gopurs_runtime_Value_1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__gopurs_runtime_Value_1614463960(gopurs_runtime.CoerceToStruct[Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__gopurs_runtime_Value_1614463960
}

var cache_addRecord gopurs_runtime.Value
var once_addRecord sync.Once
func Get_addRecord() gopurs_runtime.Value {
	once_addRecord.Do(func() {
		cache_addRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_addRecord(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_addRecord
}

var cache_addRecord__gopurs_runtime_Value_2348817094 gopurs_runtime.Value
var once_addRecord__gopurs_runtime_Value_2348817094 sync.Once
func Get_addRecord__gopurs_runtime_Value_2348817094() gopurs_runtime.Value {
	once_addRecord__gopurs_runtime_Value_2348817094.Do(func() {
		cache_addRecord__gopurs_runtime_Value_2348817094 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_addRecord__gopurs_runtime_Value_2348817094(gopurs_runtime.CoerceToStruct[Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_addRecord__gopurs_runtime_Value_2348817094
}

var cache_semiringRecord gopurs_runtime.Value
var once_semiringRecord sync.Once
func Get_semiringRecord() gopurs_runtime.Value {
	once_semiringRecord.Do(func() {
		cache_semiringRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictSemiringRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringRecord(_dollar__unused_0_box, dictSemiringRecord_1_box)
})
	})
	return cache_semiringRecord
}

var cache_add gopurs_runtime.Value
var once_add sync.Once
func Get_add() gopurs_runtime.Value {
	once_add.Do(func() {
		cache_add = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add(gopurs_runtime.CoerceToStruct[Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add
}

var cache_add__gopurs_runtime_Value_1614463960 gopurs_runtime.Value
var once_add__gopurs_runtime_Value_1614463960 sync.Once
func Get_add__gopurs_runtime_Value_1614463960() gopurs_runtime.Value {
	once_add__gopurs_runtime_Value_1614463960.Do(func() {
		cache_add__gopurs_runtime_Value_1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__gopurs_runtime_Value_1614463960(gopurs_runtime.CoerceToStruct[Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__gopurs_runtime_Value_1614463960
}

var cache_semiringFn gopurs_runtime.Value
var once_semiringFn sync.Once
func Get_semiringFn() gopurs_runtime.Value {
	once_semiringFn.Do(func() {
		cache_semiringFn = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringFn(dictSemiring_0_box)
})
	})
	return cache_semiringFn
}

var cache_semiringRecordCons gopurs_runtime.Value
var once_semiringRecordCons sync.Once
func Get_semiringRecordCons() gopurs_runtime.Value {
	once_semiringRecordCons.Do(func() {
		cache_semiringRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictSemiringRecord_2_box gopurs_runtime.Value, dictSemiring_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictSemiringRecord_2_box, dictSemiring_3_box)
})
	})
	return cache_semiringRecordCons
}

type Constructor_SemiringRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3914418263] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "addRecord": return c.V0
		case "mulRecord": return c.V1
		case "oneRecord": return c.V2
		case "zeroRecord": return c.V3
		default: panic("Key not found in dictionary Constructor_SemiringRecord: " + key)
		}
	}
}


type Constructor_Semiring[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 T_a
	V3 T_a
}


func init() {
	gopurs_runtime.StructGetters[134961754] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Semiring[gopurs_runtime.Value])(ptr)
		switch key {
		case "add": return c.V0
		case "mul": return c.V1
		case "one": return c.V2
		case "zero": return c.V3
		default: panic("Key not found in dictionary Constructor_Semiring: " + key)
		}
	}
}


func Call_zeroRecord(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_zeroRecord__gopurs_runtime_Value_2618241440(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_zero(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_zero__gopurs_runtime_Value_1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_oneRecord(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_oneRecord__gopurs_runtime_Value_2618241440(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_one(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_one__gopurs_runtime_Value_1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_mulRecord(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_mulRecord__gopurs_runtime_Value_2348817094(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_mul(dict_0_loop *Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_mul__gopurs_runtime_Value_1614463960(dict_0_loop *Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_addRecord(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_addRecord__gopurs_runtime_Value_2348817094(dict_0_loop *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_semiringRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictSemiringRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictSemiringRecord_1 gopurs_runtime.Value = dictSemiringRecord_1_loop
_ = dictSemiringRecord_1
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}))
}

func Call_add(dict_0_loop *Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__gopurs_runtime_Value_1614463960(dict_0_loop *Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_semiringFn(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "one")
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "zero")
}))
}

func Call_semiringRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictSemiringRecord_2_loop gopurs_runtime.Value, dictSemiring_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictSemiringRecord_2 gopurs_runtime.Value = dictSemiringRecord_2_loop
_ = dictSemiringRecord_2
var dictSemiring_3 gopurs_runtime.Value = dictSemiring_3_loop
_ = dictSemiring_3
one1_4_0 := gopurs_runtime.RecordGet(dictSemiring_3, "one")
_ = one1_4_0
zero1_5_1 := gopurs_runtime.RecordGet(dictSemiring_3, "zero")
_ = zero1_5_1
return gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})
_ = key_9_2
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_2)
_ = get_10_3
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_3, "add"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemiringRecord_2, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})
_ = key_9_4
get_10_5 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_4)
_ = get_10_5
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_3, "mul"), gopurs_runtime.Apply(get_10_5, ra_7), gopurs_runtime.Apply(get_10_5, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemiringRecord_2, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), one1_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_2, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}))
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), zero1_5_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_2, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}))
})
}))
}

func Get_intAdd() gopurs_runtime.Value {
	return _Gopurs_IntAdd
}

func Get_intMul() gopurs_runtime.Value {
	return _Gopurs_IntMul
}

func Get_numAdd() gopurs_runtime.Value {
	return _Gopurs_NumAdd
}

func Get_numMul() gopurs_runtime.Value {
	return _Gopurs_NumMul
}

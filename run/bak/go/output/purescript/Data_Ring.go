package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ring_semiringRecord gopurs_runtime.Value
var once_Data_Ring_semiringRecord sync.Once
func Get_Data_Ring_semiringRecord() gopurs_runtime.Value {
	once_Data_Ring_semiringRecord.Do(func() {
		cache_Data_Ring_semiringRecord = gopurs_runtime.Func(func(dictSemiringRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_semiringRecord(dictSemiringRecord_0_box)
})
	})
	return cache_Data_Ring_semiringRecord
}

var cache_Data_Ring_RingRecord_dollarDict gopurs_runtime.Value
var once_Data_Ring_RingRecord_dollarDict sync.Once
func Get_Data_Ring_RingRecord_dollarDict() gopurs_runtime.Value {
	once_Data_Ring_RingRecord_dollarDict.Do(func() {
		cache_Data_Ring_RingRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_RingRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_Ring_RingRecord_dollarDict
}

var cache_Data_Ring_Ring_dollarDict gopurs_runtime.Value
var once_Data_Ring_Ring_dollarDict sync.Once
func Get_Data_Ring_Ring_dollarDict() gopurs_runtime.Value {
	once_Data_Ring_Ring_dollarDict.Do(func() {
		cache_Data_Ring_Ring_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_Ring_dollarDict(x_0_box)
})
	})
	return cache_Data_Ring_Ring_dollarDict
}

var cache_Data_Ring_subRecord gopurs_runtime.Value
var once_Data_Ring_subRecord sync.Once
func Get_Data_Ring_subRecord() gopurs_runtime.Value {
	once_Data_Ring_subRecord.Do(func() {
		cache_Data_Ring_subRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_subRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord](dict_0_box))
})
	})
	return cache_Data_Ring_subRecord
}

var cache_Data_Ring_sub gopurs_runtime.Value
var once_Data_Ring_sub sync.Once
func Get_Data_Ring_sub() gopurs_runtime.Value {
	once_Data_Ring_sub.Do(func() {
		cache_Data_Ring_sub = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub
}

var cache_Data_Ring_ringUnit gopurs_runtime.Value
var once_Data_Ring_ringUnit sync.Once
func Get_Data_Ring_ringUnit() gopurs_runtime.Value {
	once_Data_Ring_ringUnit.Do(func() {
		cache_Data_Ring_ringUnit = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semiring_semiringUnit()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}))
	})
	return cache_Data_Ring_ringUnit
}

var cache_Data_Ring_ringRecordNil gopurs_runtime.Value
var once_Data_Ring_ringRecordNil sync.Once
func Get_Data_Ring_ringRecordNil() gopurs_runtime.Value {
	once_Data_Ring_ringRecordNil.Do(func() {
		cache_Data_Ring_ringRecordNil = gopurs_runtime.RecordDict2("SemiringRecord0", "subRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semiring_semiringRecordNil()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_Data_Ring_ringRecordNil
}

var cache_Data_Ring_ringRecordCons gopurs_runtime.Value
var once_Data_Ring_ringRecordCons sync.Once
func Get_Data_Ring_ringRecordCons() gopurs_runtime.Value {
	once_Data_Ring_ringRecordCons.Do(func() {
		cache_Data_Ring_ringRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_ringRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictRingRecord_2_box)
})
	})
	return cache_Data_Ring_ringRecordCons
}

var cache_Data_Ring_ringRecord gopurs_runtime.Value
var once_Data_Ring_ringRecord sync.Once
func Get_Data_Ring_ringRecord() gopurs_runtime.Value {
	once_Data_Ring_ringRecord.Do(func() {
		cache_Data_Ring_ringRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_ringRecord(_dollar__unused_0_box, dictRingRecord_1_box)
})
	})
	return cache_Data_Ring_ringRecord
}

var cache_Data_Ring_ringProxy gopurs_runtime.Value
var once_Data_Ring_ringProxy sync.Once
func Get_Data_Ring_ringProxy() gopurs_runtime.Value {
	once_Data_Ring_ringProxy.Do(func() {
		cache_Data_Ring_ringProxy = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semiring_semiringProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Ring_ringProxy
}

var cache_Data_Ring_ringNumber gopurs_runtime.Value
var once_Data_Ring_ringNumber sync.Once
func Get_Data_Ring_ringNumber() gopurs_runtime.Value {
	once_Data_Ring_ringNumber.Do(func() {
		cache_Data_Ring_ringNumber = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), Get_Data_Ring_numSub())
	})
	return cache_Data_Ring_ringNumber
}

var cache_Data_Ring_ringInt gopurs_runtime.Value
var once_Data_Ring_ringInt sync.Once
func Get_Data_Ring_ringInt() gopurs_runtime.Value {
	once_Data_Ring_ringInt.Do(func() {
		cache_Data_Ring_ringInt = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0))
}), Get_Data_Ring_intSub())
	})
	return cache_Data_Ring_ringInt
}

var cache_Data_Ring_ringFn gopurs_runtime.Value
var once_Data_Ring_ringFn sync.Once
func Get_Data_Ring_ringFn() gopurs_runtime.Value {
	once_Data_Ring_ringFn.Do(func() {
		cache_Data_Ring_ringFn = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_ringFn(dictRing_0_box)
})
	})
	return cache_Data_Ring_ringFn
}

var cache_Data_Ring_negate gopurs_runtime.Value
var once_Data_Ring_negate sync.Once
func Get_Data_Ring_negate() gopurs_runtime.Value {
	once_Data_Ring_negate.Do(func() {
		cache_Data_Ring_negate = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_0_box))
})
	})
	return cache_Data_Ring_negate
}

var cache_Data_Ring_negate__1541574592 gopurs_runtime.Value
var once_Data_Ring_negate__1541574592 sync.Once
func Get_Data_Ring_negate__1541574592() gopurs_runtime.Value {
	once_Data_Ring_negate__1541574592.Do(func() {
		cache_Data_Ring_negate__1541574592 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate__1541574592(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_0_box))
})
	})
	return cache_Data_Ring_negate__1541574592
}

var cache_Data_Ring_negate__2635823316 gopurs_runtime.Value
var once_Data_Ring_negate__2635823316 sync.Once
func Get_Data_Ring_negate__2635823316() gopurs_runtime.Value {
	once_Data_Ring_negate__2635823316.Do(func() {
		cache_Data_Ring_negate__2635823316 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate__2635823316(__eta0_0_box)
})
	})
	return cache_Data_Ring_negate__2635823316
}

var cache_Data_Ring_negate__2151342916 gopurs_runtime.Value
var once_Data_Ring_negate__2151342916 sync.Once
func Get_Data_Ring_negate__2151342916() gopurs_runtime.Value {
	once_Data_Ring_negate__2151342916.Do(func() {
		cache_Data_Ring_negate__2151342916 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate__2151342916(__eta0_0_box)
})
	})
	return cache_Data_Ring_negate__2151342916
}

var cache_Data_Ring_negate__1364373265 gopurs_runtime.Value
var once_Data_Ring_negate__1364373265 sync.Once
func Get_Data_Ring_negate__1364373265() gopurs_runtime.Value {
	once_Data_Ring_negate__1364373265.Do(func() {
		cache_Data_Ring_negate__1364373265 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate__1364373265(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_0_box))
})
	})
	return cache_Data_Ring_negate__1364373265
}

var cache_Data_Ring_negate__753141756 gopurs_runtime.Value
var once_Data_Ring_negate__753141756 sync.Once
func Get_Data_Ring_negate__753141756() gopurs_runtime.Value {
	once_Data_Ring_negate__753141756.Do(func() {
		cache_Data_Ring_negate__753141756 = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate__753141756(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dictRing_0_box))
})
	})
	return cache_Data_Ring_negate__753141756
}

var cache_Data_Ring_ringProxy__3154238922 gopurs_runtime.Value
var once_Data_Ring_ringProxy__3154238922 sync.Once
func Get_Data_Ring_ringProxy__3154238922() gopurs_runtime.Value {
	once_Data_Ring_ringProxy__3154238922.Do(func() {
		cache_Data_Ring_ringProxy__3154238922 = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semiring_semiringProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Ring_ringProxy__3154238922
}

var cache_Data_Ring_ringRecordNil__3706727823 gopurs_runtime.Value
var once_Data_Ring_ringRecordNil__3706727823 sync.Once
func Get_Data_Ring_ringRecordNil__3706727823() gopurs_runtime.Value {
	once_Data_Ring_ringRecordNil__3706727823.Do(func() {
		cache_Data_Ring_ringRecordNil__3706727823 = gopurs_runtime.RecordDict2("SemiringRecord0", "subRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Semiring_semiringRecordNil()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_Data_Ring_ringRecordNil__3706727823
}

var cache_Data_Ring_sub__2927892844 gopurs_runtime.Value
var once_Data_Ring_sub__2927892844 sync.Once
func Get_Data_Ring_sub__2927892844() gopurs_runtime.Value {
	once_Data_Ring_sub__2927892844.Do(func() {
		cache_Data_Ring_sub__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__2927892844(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__2927892844
}

var cache_Data_Ring_sub__101133084 gopurs_runtime.Value
var once_Data_Ring_sub__101133084 sync.Once
func Get_Data_Ring_sub__101133084() gopurs_runtime.Value {
	once_Data_Ring_sub__101133084.Do(func() {
		cache_Data_Ring_sub__101133084 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__101133084(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__101133084
}

var cache_Data_Ring_sub__1124926121 gopurs_runtime.Value
var once_Data_Ring_sub__1124926121 sync.Once
func Get_Data_Ring_sub__1124926121() gopurs_runtime.Value {
	once_Data_Ring_sub__1124926121.Do(func() {
		cache_Data_Ring_sub__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__1124926121(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__1124926121
}

var cache_Data_Ring_sub__1841809173 gopurs_runtime.Value
var once_Data_Ring_sub__1841809173 sync.Once
func Get_Data_Ring_sub__1841809173() gopurs_runtime.Value {
	once_Data_Ring_sub__1841809173.Do(func() {
		cache_Data_Ring_sub__1841809173 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__1841809173(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__1841809173
}

var cache_Data_Ring_sub__190951261 gopurs_runtime.Value
var once_Data_Ring_sub__190951261 sync.Once
func Get_Data_Ring_sub__190951261() gopurs_runtime.Value {
	once_Data_Ring_sub__190951261.Do(func() {
		cache_Data_Ring_sub__190951261 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__190951261(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__190951261
}

var cache_Data_Ring_sub__1043827704 gopurs_runtime.Value
var once_Data_Ring_sub__1043827704 sync.Once
func Get_Data_Ring_sub__1043827704() gopurs_runtime.Value {
	once_Data_Ring_sub__1043827704.Do(func() {
		cache_Data_Ring_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Ring_sub__1043827704
}

var cache_Data_Ring_sub__1135378904 gopurs_runtime.Value
var once_Data_Ring_sub__1135378904 sync.Once
func Get_Data_Ring_sub__1135378904() gopurs_runtime.Value {
	once_Data_Ring_sub__1135378904.Do(func() {
		cache_Data_Ring_sub__1135378904 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__1135378904(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Ring_sub__1135378904
}

var cache_Data_Ring_sub__1023426360 gopurs_runtime.Value
var once_Data_Ring_sub__1023426360 sync.Once
func Get_Data_Ring_sub__1023426360() gopurs_runtime.Value {
	once_Data_Ring_sub__1023426360.Do(func() {
		cache_Data_Ring_sub__1023426360 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__1023426360(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__1023426360
}

var cache_Data_Ring_sub__3675938712 gopurs_runtime.Value
var once_Data_Ring_sub__3675938712 sync.Once
func Get_Data_Ring_sub__3675938712() gopurs_runtime.Value {
	once_Data_Ring_sub__3675938712.Do(func() {
		cache_Data_Ring_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__3675938712(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__3675938712
}

var cache_Data_Ring_sub__2345699288 gopurs_runtime.Value
var once_Data_Ring_sub__2345699288 sync.Once
func Get_Data_Ring_sub__2345699288() gopurs_runtime.Value {
	once_Data_Ring_sub__2345699288.Do(func() {
		cache_Data_Ring_sub__2345699288 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__2345699288(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__2345699288
}

var cache_Data_Ring_sub__3885659384 gopurs_runtime.Value
var once_Data_Ring_sub__3885659384 sync.Once
func Get_Data_Ring_sub__3885659384() gopurs_runtime.Value {
	once_Data_Ring_sub__3885659384.Do(func() {
		cache_Data_Ring_sub__3885659384 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__3885659384(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__3885659384
}

var cache_Data_Ring_sub__871462840 gopurs_runtime.Value
var once_Data_Ring_sub__871462840 sync.Once
func Get_Data_Ring_sub__871462840() gopurs_runtime.Value {
	once_Data_Ring_sub__871462840.Do(func() {
		cache_Data_Ring_sub__871462840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__871462840(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring](dict_0_box))
})
	})
	return cache_Data_Ring_sub__871462840
}

var cache_Data_Ring_subRecord__3826282112 gopurs_runtime.Value
var once_Data_Ring_subRecord__3826282112 sync.Once
func Get_Data_Ring_subRecord__3826282112() gopurs_runtime.Value {
	once_Data_Ring_subRecord__3826282112.Do(func() {
		cache_Data_Ring_subRecord__3826282112 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_subRecord__3826282112(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord](dict_0_box))
})
	})
	return cache_Data_Ring_subRecord__3826282112
}

var cache_Data_Ring_subRecord__2650724742 gopurs_runtime.Value
var once_Data_Ring_subRecord__2650724742 sync.Once
func Get_Data_Ring_subRecord__2650724742() gopurs_runtime.Value {
	once_Data_Ring_subRecord__2650724742.Do(func() {
		cache_Data_Ring_subRecord__2650724742 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_subRecord__2650724742(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord](dict_0_box))
})
	})
	return cache_Data_Ring_subRecord__2650724742
}

type Constructor_Data_Ring_RingRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4246880791] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_RingRecord)(ptr)
		_ = c
		switch key {
		case "SemiringRecord0": return gopurs_runtime.Box(c.V0)
		case "subRecord": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ring_RingRecord: " + key)
		}
	}
}


type Constructor_Data_Ring_Ring struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3955491866] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_Ring)(ptr)
		_ = c
		switch key {
		case "Semiring0": return gopurs_runtime.Box(c.V0)
		case "sub": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ring_Ring: " + key)
		}
	}
}


func Call_Data_Ring_semiringRecord(dictSemiringRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiringRecord_0 gopurs_runtime.Value = dictSemiringRecord_0_loop
_ = dictSemiringRecord_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_0, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_0, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_0, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_0, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_Ring_RingRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ring_Ring_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Ring_subRecord(dict_0_loop *Constructor_Data_Ring_RingRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_RingRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_ringRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictRingRecord_2 gopurs_runtime.Value = dictRingRecord_2_loop
_ = dictRingRecord_2
// TAST (Let): semiringRecordCons1_3_0 -> gopurs_runtime.Value
semiringRecordCons1_3_0 := gopurs_runtime.Apply3(Get_Data_Semiring_semiringRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_2, "SemiringRecord0"), gopurs_runtime.Value{}))
_ = semiringRecordCons1_3_0
return gopurs_runtime.Func(func(dictRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semiringRecordCons2_5_1 -> gopurs_runtime.Value
semiringRecordCons2_5_1 := gopurs_runtime.Apply(semiringRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_4, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringRecordCons2_5_1
return gopurs_runtime.RecordDict2("SemiringRecord0", "subRecord", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecordCons2_5_1
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_9_2 -> gopurs_runtime.Value
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_9_2
// TAST (Let): get_10_3 -> gopurs_runtime.Value
get_10_3 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_9_2.StrVal()))
_ = get_10_3
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_2.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_4, "sub"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictRingRecord_2, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
})
})
}))
})
}

func Call_Data_Ring_ringRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictRingRecord_1 gopurs_runtime.Value = dictRingRecord_1_loop
_ = dictRingRecord_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): semiringRecord1_2_0 -> gopurs_runtime.Value
semiringRecord1_2_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = semiringRecord1_2_0
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_2_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_Ring_ringFn(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semiringFn_1_0 -> gopurs_runtime.Value
semiringFn_1_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "add"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mul"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "one")
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "zero")
}))
_ = semiringFn_1_0
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringFn_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}))
}

func Call_Data_Ring_negate(dictRing_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dictRing_0 *Constructor_Data_Ring_Ring = dictRing_0_loop
_ = dictRing_0
// TAST (Let): Semiring0_1_0 -> *Constructor_Data_Semiring_Semiring
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_0.V0), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_0.V1), gopurs_runtime.Box(Semiring0_1_0.V3), a_2)
})
}

func Call_Data_Ring_negate__1541574592(dictRing_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dictRing_0 *Constructor_Data_Ring_Ring = dictRing_0_loop
_ = dictRing_0
// TAST (Let): Semiring0_1_0 -> *Constructor_Data_Semiring_Semiring
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_0.V0), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_0.V1), gopurs_runtime.Int(gopurs_runtime.Box(Semiring0_1_0.V3).IntVal), gopurs_runtime.Int(a_2.IntVal)).IntVal)
})
}

func Call_Data_Ring_negate__2635823316(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Int(-(__eta0_0.IntVal))
}

func Call_Data_Ring_negate__2151342916(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Float(-(__eta0_0.FloatVal()))
}

func Call_Data_Ring_negate__1364373265(dictRing_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dictRing_0 *Constructor_Data_Ring_Ring = dictRing_0_loop
_ = dictRing_0
// TAST (Let): Semiring0_1_0 -> *Constructor_Data_Semiring_Semiring
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_0.V0), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_0.V1), gopurs_runtime.Box(Semiring0_1_0.V3), a_2)
})
}

func Call_Data_Ring_negate__753141756(dictRing_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dictRing_0 *Constructor_Data_Ring_Ring = dictRing_0_loop
_ = dictRing_0
// TAST (Let): Semiring0_1_0 -> *Constructor_Data_Semiring_Semiring
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_0.V0), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Box(Semiring0_1_0.V3)))}, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a_2))})))}
})
}

func Call_Data_Ring_sub__2927892844(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__101133084(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__1124926121(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__1841809173(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__190951261(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_Data_Ring_sub__1135378904(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) - (__eta1_1.FloatVal()))
}

func Call_Data_Ring_sub__1023426360(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__3675938712(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__2345699288(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__3885659384(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__871462840(dict_0_loop *Constructor_Data_Ring_Ring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_subRecord__3826282112(dict_0_loop *Constructor_Data_Ring_RingRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_RingRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_subRecord__2650724742(dict_0_loop *Constructor_Data_Ring_RingRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_RingRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Get_Data_Ring_intSub() gopurs_runtime.Value {
	return _Gopurs_Data_Ring_IntSub
}

func Get_Data_Ring_numSub() gopurs_runtime.Value {
	return _Gopurs_Data_Ring_NumSub
}

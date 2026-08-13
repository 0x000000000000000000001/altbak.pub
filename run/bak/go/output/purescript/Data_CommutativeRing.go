package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_CommutativeRing_ringRecord gopurs_runtime.Value
var once_Data_CommutativeRing_ringRecord sync.Once
func Get_Data_CommutativeRing_ringRecord() gopurs_runtime.Value {
	once_Data_CommutativeRing_ringRecord.Do(func() {
		cache_Data_CommutativeRing_ringRecord = gopurs_runtime.Func(func(dictRingRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_ringRecord(dictRingRecord_0_box)
})
	})
	return cache_Data_CommutativeRing_ringRecord
}

var cache_Data_CommutativeRing_CommutativeRingRecord_dollarDict gopurs_runtime.Value
var once_Data_CommutativeRing_CommutativeRingRecord_dollarDict sync.Once
func Get_Data_CommutativeRing_CommutativeRingRecord_dollarDict() gopurs_runtime.Value {
	once_Data_CommutativeRing_CommutativeRingRecord_dollarDict.Do(func() {
		cache_Data_CommutativeRing_CommutativeRingRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_CommutativeRingRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_CommutativeRing_CommutativeRingRecord_dollarDict
}

var cache_Data_CommutativeRing_CommutativeRing_dollarDict gopurs_runtime.Value
var once_Data_CommutativeRing_CommutativeRing_dollarDict sync.Once
func Get_Data_CommutativeRing_CommutativeRing_dollarDict() gopurs_runtime.Value {
	once_Data_CommutativeRing_CommutativeRing_dollarDict.Do(func() {
		cache_Data_CommutativeRing_CommutativeRing_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_CommutativeRing_dollarDict(x_0_box)
})
	})
	return cache_Data_CommutativeRing_CommutativeRing_dollarDict
}

var cache_Data_CommutativeRing_commutativeRingUnit gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingUnit sync.Once
func Get_Data_CommutativeRing_commutativeRingUnit() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingUnit.Do(func() {
		cache_Data_CommutativeRing_commutativeRingUnit = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ring_ringUnit()
}))
	})
	return cache_Data_CommutativeRing_commutativeRingUnit
}

var cache_Data_CommutativeRing_commutativeRingRecordNil gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingRecordNil sync.Once
func Get_Data_CommutativeRing_commutativeRingRecordNil() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingRecordNil.Do(func() {
		cache_Data_CommutativeRing_commutativeRingRecordNil = gopurs_runtime.RecordDict1("RingRecord0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ring_ringRecordNil()
}))
	})
	return cache_Data_CommutativeRing_commutativeRingRecordNil
}

var cache_Data_CommutativeRing_commutativeRingRecordCons gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingRecordCons sync.Once
func Get_Data_CommutativeRing_commutativeRingRecordCons() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingRecordCons.Do(func() {
		cache_Data_CommutativeRing_commutativeRingRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictCommutativeRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_commutativeRingRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictCommutativeRingRecord_2_box)
})
	})
	return cache_Data_CommutativeRing_commutativeRingRecordCons
}

var cache_Data_CommutativeRing_commutativeRingRecord gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingRecord sync.Once
func Get_Data_CommutativeRing_commutativeRingRecord() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingRecord.Do(func() {
		cache_Data_CommutativeRing_commutativeRingRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictCommutativeRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_commutativeRingRecord(_dollar__unused_0_box, dictCommutativeRingRecord_1_box)
})
	})
	return cache_Data_CommutativeRing_commutativeRingRecord
}

var cache_Data_CommutativeRing_commutativeRingProxy gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingProxy sync.Once
func Get_Data_CommutativeRing_commutativeRingProxy() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingProxy.Do(func() {
		cache_Data_CommutativeRing_commutativeRingProxy = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ring_ringProxy()
}))
	})
	return cache_Data_CommutativeRing_commutativeRingProxy
}

var cache_Data_CommutativeRing_commutativeRingNumber gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingNumber sync.Once
func Get_Data_CommutativeRing_commutativeRingNumber() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingNumber.Do(func() {
		cache_Data_CommutativeRing_commutativeRingNumber = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), Get_Data_Ring_numSub())
}))
	})
	return cache_Data_CommutativeRing_commutativeRingNumber
}

var cache_Data_CommutativeRing_commutativeRingInt gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingInt sync.Once
func Get_Data_CommutativeRing_commutativeRingInt() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingInt.Do(func() {
		cache_Data_CommutativeRing_commutativeRingInt = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0))
}), Get_Data_Ring_intSub())
}))
	})
	return cache_Data_CommutativeRing_commutativeRingInt
}

var cache_Data_CommutativeRing_commutativeRingFn gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingFn sync.Once
func Get_Data_CommutativeRing_commutativeRingFn() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingFn.Do(func() {
		cache_Data_CommutativeRing_commutativeRingFn = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_commutativeRingFn(dictCommutativeRing_0_box)
})
	})
	return cache_Data_CommutativeRing_commutativeRingFn
}

type Constructor_Data_CommutativeRing_CommutativeRingRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[292222263] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_CommutativeRing_CommutativeRingRecord)(ptr)
		_ = c
		switch key {
		case "RingRecord0": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_CommutativeRing_CommutativeRingRecord: " + key)
		}
	}
}


type Constructor_Data_CommutativeRing_CommutativeRing struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1775085946] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_CommutativeRing_CommutativeRing)(ptr)
		_ = c
		switch key {
		case "Ring0": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_CommutativeRing_CommutativeRing: " + key)
		}
	}
}


func Call_Data_CommutativeRing_ringRecord(dictRingRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRingRecord_0 gopurs_runtime.Value = dictRingRecord_0_loop
_ = dictRingRecord_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_0, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semiringRecord1_1_0 -> gopurs_runtime.Value
semiringRecord1_1_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = semiringRecord1_1_0
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_0, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_CommutativeRing_CommutativeRingRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_CommutativeRing_CommutativeRing_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_CommutativeRing_commutativeRingRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictCommutativeRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictCommutativeRingRecord_2 gopurs_runtime.Value = dictCommutativeRingRecord_2_loop
_ = dictCommutativeRingRecord_2
// TAST (Let): ringRecordCons1_3_0 -> gopurs_runtime.Value
ringRecordCons1_3_0 := gopurs_runtime.Apply3(Get_Data_Ring_ringRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRingRecord_2, "RingRecord0"), gopurs_runtime.Value{}))
_ = ringRecordCons1_3_0
return gopurs_runtime.Func(func(dictCommutativeRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ringRecordCons2_5_1 -> gopurs_runtime.Value
ringRecordCons2_5_1 := gopurs_runtime.Apply(ringRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_4, "Ring0"), gopurs_runtime.Value{}))
_ = ringRecordCons2_5_1
return gopurs_runtime.RecordDict1("RingRecord0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return ringRecordCons2_5_1
}))
})
}

func Call_Data_CommutativeRing_commutativeRingRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictCommutativeRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictCommutativeRingRecord_1 gopurs_runtime.Value = dictCommutativeRingRecord_1_loop
_ = dictCommutativeRingRecord_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRingRecord_1, "RingRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): semiringRecord1_3_2 -> gopurs_runtime.Value
semiringRecord1_3_2 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = semiringRecord1_3_2
// TAST (Let): ringRecord1_2_0 -> gopurs_runtime.Value
ringRecord1_2_0 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_3_2
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = ringRecord1_2_0
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ringRecord1_2_0
}))
}

func Call_Data_CommutativeRing_commutativeRingFn(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): semiringFn_2_2 -> gopurs_runtime.Value
semiringFn_2_2 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "add"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "mul"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
})
})
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_2_3, "one")
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_2_3, "zero")
}))
_ = semiringFn_2_2
// TAST (Let): ringFn_1_0 -> gopurs_runtime.Value
ringFn_1_0 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringFn_2_2
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "sub"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
})
})
}))
_ = ringFn_1_0
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringFn_1_0
}))
}



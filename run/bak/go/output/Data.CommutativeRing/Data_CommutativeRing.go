package Data_CommutativeRing

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
)

var cache_commutativeRingUnit gopurs_runtime.Value
var once_commutativeRingUnit sync.Once
func Get_commutativeRingUnit() gopurs_runtime.Value {
	once_commutativeRingUnit.Do(func() {
		cache_commutativeRingUnit = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringUnit()
}))
	})
	return cache_commutativeRingUnit
}

var cache_commutativeRingRecordNil gopurs_runtime.Value
var once_commutativeRingRecordNil sync.Once
func Get_commutativeRingRecordNil() gopurs_runtime.Value {
	once_commutativeRingRecordNil.Do(func() {
		cache_commutativeRingRecordNil = gopurs_runtime.RecordDict1("RingRecord0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringRecordNil()
}))
	})
	return cache_commutativeRingRecordNil
}

var cache_commutativeRingRecordCons gopurs_runtime.Value
var once_commutativeRingRecordCons sync.Once
func Get_commutativeRingRecordCons() gopurs_runtime.Value {
	once_commutativeRingRecordCons.Do(func() {
		cache_commutativeRingRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictCommutativeRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_commutativeRingRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictCommutativeRingRecord_2_box)
})
	})
	return cache_commutativeRingRecordCons
}

var cache_commutativeRingRecord gopurs_runtime.Value
var once_commutativeRingRecord sync.Once
func Get_commutativeRingRecord() gopurs_runtime.Value {
	once_commutativeRingRecord.Do(func() {
		cache_commutativeRingRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictCommutativeRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_commutativeRingRecord(_dollar__unused_0_box, dictCommutativeRingRecord_1_box)
})
	})
	return cache_commutativeRingRecord
}

var cache_commutativeRingProxy gopurs_runtime.Value
var once_commutativeRingProxy sync.Once
func Get_commutativeRingProxy() gopurs_runtime.Value {
	once_commutativeRingProxy.Do(func() {
		cache_commutativeRingProxy = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringProxy()
}))
	})
	return cache_commutativeRingProxy
}

var cache_commutativeRingNumber gopurs_runtime.Value
var once_commutativeRingNumber sync.Once
func Get_commutativeRingNumber() gopurs_runtime.Value {
	once_commutativeRingNumber.Do(func() {
		cache_commutativeRingNumber = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ring.Get_ringNumber()
}))
	})
	return cache_commutativeRingNumber
}

var cache_commutativeRingInt gopurs_runtime.Value
var once_commutativeRingInt sync.Once
func Get_commutativeRingInt() gopurs_runtime.Value {
	once_commutativeRingInt.Do(func() {
		cache_commutativeRingInt = gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", pkg_Data_Semiring.Get_intAdd(), pkg_Data_Semiring.Get_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0))
}), pkg_Data_Ring.Get_intSub())
}))
	})
	return cache_commutativeRingInt
}

var cache_commutativeRingFn gopurs_runtime.Value
var once_commutativeRingFn sync.Once
func Get_commutativeRingFn() gopurs_runtime.Value {
	once_commutativeRingFn.Do(func() {
		cache_commutativeRingFn = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_commutativeRingFn(dictCommutativeRing_0_box)
})
	})
	return cache_commutativeRingFn
}

func Call_commutativeRingRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictCommutativeRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictCommutativeRingRecord_2 gopurs_runtime.Value = dictCommutativeRingRecord_2_loop
_ = dictCommutativeRingRecord_2
ringRecordCons1_3_0 := gopurs_runtime.Apply3(pkg_Data_Ring.Get_ringRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRingRecord_2, "RingRecord0"), gopurs_runtime.Value{}))
_ = ringRecordCons1_3_0
return gopurs_runtime.Func(func(dictCommutativeRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
ringRecordCons2_5_1 := gopurs_runtime.Apply(ringRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_4, "Ring0"), gopurs_runtime.Value{}))
_ = ringRecordCons2_5_1
return gopurs_runtime.RecordDict1("RingRecord0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return ringRecordCons2_5_1
}))
})
}

func Call_commutativeRingRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictCommutativeRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictCommutativeRingRecord_1 gopurs_runtime.Value = dictCommutativeRingRecord_1_loop
_ = dictCommutativeRingRecord_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRingRecord_1, "RingRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_3_1
semiringRecord1_4_3 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = semiringRecord1_4_3
ringRecord1_4_2 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_4_3
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = ringRecord1_4_2
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ringRecord1_4_2
}))
}

func Call_commutativeRingFn(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_2_1
zero1_3_3 := gopurs_runtime.RecordGet(__local_var_2_1, "zero")
_ = zero1_3_3
one1_4_4 := gopurs_runtime.RecordGet(__local_var_2_1, "one")
_ = one1_4_4
semiringFn_5_5 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, g_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "add"), gopurs_runtime.Apply(f_5, x_7), gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, g_6 gopurs_runtime.Value, x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "mul"), gopurs_runtime.Apply(f_5, x_7), gopurs_runtime.Apply(g_6, x_7))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return one1_4_4
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return zero1_3_3
}))
_ = semiringFn_5_5
ringFn_3_2 := gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringFn_5_5
}), gopurs_runtime.Func3(func(f_6 gopurs_runtime.Value, g_7 gopurs_runtime.Value, x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "sub"), gopurs_runtime.Apply(f_6, x_8), gopurs_runtime.Apply(g_7, x_8))
}))
_ = ringFn_3_2
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ringFn_3_2
}))
}



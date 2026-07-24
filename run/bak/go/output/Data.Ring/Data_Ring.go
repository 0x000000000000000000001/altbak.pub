package Data_Ring

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var subRecord gopurs_runtime.Value
var once_subRecord sync.Once
func Get_subRecord() gopurs_runtime.Value {
	once_subRecord.Do(func() {
		subRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "subRecord")
}()
})
	})
	return subRecord
}

var sub gopurs_runtime.Value
var once_sub sync.Once
func Get_sub() gopurs_runtime.Value {
	once_sub.Do(func() {
		sub = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "sub")
}()
})
	})
	return sub
}

var ringUnit gopurs_runtime.Value
var once_ringUnit sync.Once
func Get_ringUnit() gopurs_runtime.Value {
	once_ringUnit.Do(func() {
		ringUnit = gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringUnit()
}))
	})
	return ringUnit
}

var ringRecordNil gopurs_runtime.Value
var once_ringRecordNil sync.Once
func Get_ringRecordNil() gopurs_runtime.Value {
	once_ringRecordNil.Do(func() {
		ringRecordNil = gopurs_runtime.RecordDict2("subRecord", "SemiringRecord0", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringRecordNil()
}))
	})
	return ringRecordNil
}

var ringRecordCons gopurs_runtime.Value
var once_ringRecordCons sync.Once
func Get_ringRecordCons() gopurs_runtime.Value {
	once_ringRecordCons.Do(func() {
		ringRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ringRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictRingRecord_2_box)
})
	})
	return ringRecordCons
}

var ringRecord gopurs_runtime.Value
var once_ringRecord sync.Once
func Get_ringRecord() gopurs_runtime.Value {
	once_ringRecord.Do(func() {
		ringRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ringRecord(_dollar__unused_0_box, dictRingRecord_1_box)
})
	})
	return ringRecord
}

var ringProxy gopurs_runtime.Value
var once_ringProxy sync.Once
func Get_ringProxy() gopurs_runtime.Value {
	once_ringProxy.Do(func() {
		ringProxy = gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringProxy()
}))
	})
	return ringProxy
}

var ringNumber gopurs_runtime.Value
var once_ringNumber sync.Once
func Get_ringNumber() gopurs_runtime.Value {
	once_ringNumber.Do(func() {
		ringNumber = gopurs_runtime.RecordDict2("sub", "Semiring0", Get_numSub(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringNumber()
}))
	})
	return ringNumber
}

var ringInt gopurs_runtime.Value
var once_ringInt sync.Once
func Get_ringInt() gopurs_runtime.Value {
	once_ringInt.Do(func() {
		ringInt = gopurs_runtime.RecordDict2("sub", "Semiring0", Get_intSub(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringInt()
}))
	})
	return ringInt
}

var ringFn gopurs_runtime.Value
var once_ringFn sync.Once
func Get_ringFn() gopurs_runtime.Value {
	once_ringFn.Do(func() {
		ringFn = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_0
zero1_2_1 := gopurs_runtime.RecordGet(__local_var_1_0, "zero")
_ = zero1_2_1
one1_3_3 := gopurs_runtime.RecordGet(__local_var_1_0, "one")
_ = one1_3_3
semiringFn_3_2 := gopurs_runtime.RecordDict4("add", "zero", "mul", "one", gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "add"), gopurs_runtime.Apply(f_4, x_6), gopurs_runtime.Apply(g_5, x_6))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return zero1_2_1
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mul"), gopurs_runtime.Apply(f_4, x_6), gopurs_runtime.Apply(g_5, x_6))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return one1_3_3
}))
_ = semiringFn_3_2
return gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, g_5 gopurs_runtime.Value, x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(f_4, x_6), gopurs_runtime.Apply(g_5, x_6))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringFn_3_2
}))
}()
})
	})
	return ringFn
}

var negate gopurs_runtime.Value
var once_negate sync.Once
func Get_negate() gopurs_runtime.Value {
	once_negate.Do(func() {
		negate = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
zero_1_0 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{}), "zero")
_ = zero_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), zero_1_0, a_2)
})
}()
})
	})
	return negate
}

func Call_ringRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictRingRecord_2 gopurs_runtime.Value = dictRingRecord_2_loop
_ = dictRingRecord_2
semiringRecordCons1_3_0 := gopurs_runtime.Apply3(pkg_Data_Semiring.Get_semiringRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_2, "SemiringRecord0"), gopurs_runtime.Value{}))
_ = semiringRecordCons1_3_0
return gopurs_runtime.Func(func(dictRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
semiringRecordCons2_5_1 := gopurs_runtime.Apply(semiringRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_4, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringRecordCons2_5_1
return gopurs_runtime.RecordDict2("subRecord", "SemiringRecord0", gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_9_2
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_2)
_ = get_10_3
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_4, "sub"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictRingRecord_2, "subRecord"), gopurs_runtime.Constructor0("Proxy"), ra_7, rb_8))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecordCons2_5_1
}))
})
}

func Call_ringRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictRingRecord_1 gopurs_runtime.Value = dictRingRecord_1_loop
_ = dictRingRecord_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_0
semiringRecord1_3_1 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "addRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "mulRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "oneRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_0, "zeroRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")))
_ = semiringRecord1_3_1
return gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "subRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_3_1
}))
}

func Get_intSub() gopurs_runtime.Value {
	return _Gopurs_IntSub
}

func Get_numSub() gopurs_runtime.Value {
	return _Gopurs_NumSub
}

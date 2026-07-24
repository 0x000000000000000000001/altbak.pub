package Data_Semiring

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var zeroRecord gopurs_runtime.Value
var once_zeroRecord sync.Once
func Get_zeroRecord() gopurs_runtime.Value {
	once_zeroRecord.Do(func() {
		zeroRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "zeroRecord")
})
	})
	return zeroRecord
}

var zero gopurs_runtime.Value
var once_zero sync.Once
func Get_zero() gopurs_runtime.Value {
	once_zero.Do(func() {
		zero = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "zero")
})
	})
	return zero
}

var semiringUnit gopurs_runtime.Value
var once_semiringUnit sync.Once
func Get_semiringUnit() gopurs_runtime.Value {
	once_semiringUnit.Do(func() {
		semiringUnit = gopurs_runtime.RecordDict4("add", "zero", "mul", "one", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), pkg_Data_Unit.Get_unit(), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), pkg_Data_Unit.Get_unit())
	})
	return semiringUnit
}

var semiringRecordNil gopurs_runtime.Value
var once_semiringRecordNil sync.Once
func Get_semiringRecordNil() gopurs_runtime.Value {
	once_semiringRecordNil.Do(func() {
		semiringRecordNil = gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}))
	})
	return semiringRecordNil
}

var semiringProxy gopurs_runtime.Value
var once_semiringProxy sync.Once
func Get_semiringProxy() gopurs_runtime.Value {
	once_semiringProxy.Do(func() {
		semiringProxy = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy"))
	})
	return semiringProxy
}

var semiringNumber gopurs_runtime.Value
var once_semiringNumber sync.Once
func Get_semiringNumber() gopurs_runtime.Value {
	once_semiringNumber.Do(func() {
		semiringNumber = gopurs_runtime.RecordDict4("add", "zero", "mul", "one", Get_numAdd(), gopurs_runtime.Float(0.0), Get_numMul(), gopurs_runtime.Float(1.0))
	})
	return semiringNumber
}

var semiringInt gopurs_runtime.Value
var once_semiringInt sync.Once
func Get_semiringInt() gopurs_runtime.Value {
	once_semiringInt.Do(func() {
		semiringInt = gopurs_runtime.RecordDict4("add", "zero", "mul", "one", Get_intAdd(), gopurs_runtime.Int(0), Get_intMul(), gopurs_runtime.Int(1))
	})
	return semiringInt
}

var oneRecord gopurs_runtime.Value
var once_oneRecord sync.Once
func Get_oneRecord() gopurs_runtime.Value {
	once_oneRecord.Do(func() {
		oneRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "oneRecord")
})
	})
	return oneRecord
}

var one gopurs_runtime.Value
var once_one sync.Once
func Get_one() gopurs_runtime.Value {
	once_one.Do(func() {
		one = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "one")
})
	})
	return one
}

var mulRecord gopurs_runtime.Value
var once_mulRecord sync.Once
func Get_mulRecord() gopurs_runtime.Value {
	once_mulRecord.Do(func() {
		mulRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "mulRecord")
})
	})
	return mulRecord
}

var mul gopurs_runtime.Value
var once_mul sync.Once
func Get_mul() gopurs_runtime.Value {
	once_mul.Do(func() {
		mul = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "mul")
})
	})
	return mul
}

var addRecord gopurs_runtime.Value
var once_addRecord sync.Once
func Get_addRecord() gopurs_runtime.Value {
	once_addRecord.Do(func() {
		addRecord = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "addRecord")
})
	})
	return addRecord
}

var semiringRecord gopurs_runtime.Value
var once_semiringRecord sync.Once
func Get_semiringRecord() gopurs_runtime.Value {
	once_semiringRecord.Do(func() {
		semiringRecord = gopurs_runtime.Func2(func(_dollar__unused_0 gopurs_runtime.Value, dictSemiringRecord_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_1, "addRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_1, "mulRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_1, "oneRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_1, "zeroRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")))
})
	})
	return semiringRecord
}

var add gopurs_runtime.Value
var once_add sync.Once
func Get_add() gopurs_runtime.Value {
	once_add.Do(func() {
		add = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "add")
})
	})
	return add
}

var semiringFn gopurs_runtime.Value
var once_semiringFn sync.Once
func Get_semiringFn() gopurs_runtime.Value {
	once_semiringFn.Do(func() {
		semiringFn = gopurs_runtime.Func(func(dictSemiring_0 gopurs_runtime.Value) gopurs_runtime.Value {
zero1_1_0 := gopurs_runtime.RecordGet(dictSemiring_0, "zero")
_ = zero1_1_0
one1_2_1 := gopurs_runtime.RecordGet(dictSemiring_0, "one")
_ = one1_2_1
return gopurs_runtime.RecordDict4("add", "zero", "mul", "one", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return zero1_1_0
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return one1_2_1
}))
})
	})
	return semiringFn
}

var semiringRecordCons gopurs_runtime.Value
var once_semiringRecordCons sync.Once
func Get_semiringRecordCons() gopurs_runtime.Value {
	once_semiringRecordCons.Do(func() {
		semiringRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0 gopurs_runtime.Value, _dollar__unused_1 gopurs_runtime.Value, dictSemiringRecord_2 gopurs_runtime.Value, dictSemiring_3 gopurs_runtime.Value) gopurs_runtime.Value {
one1_4_0 := gopurs_runtime.RecordGet(dictSemiring_3, "one")
_ = one1_4_0
zero1_5_1 := gopurs_runtime.RecordGet(dictSemiring_3, "zero")
_ = zero1_5_1
return gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_9_2
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_2)
_ = get_10_3
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_3, "add"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemiringRecord_2, "addRecord"), gopurs_runtime.Constructor0("Proxy"), ra_7, rb_8))
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_9_4
get_10_5 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_4)
_ = get_10_5
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_3, "mul"), gopurs_runtime.Apply(get_10_5, ra_7), gopurs_runtime.Apply(get_10_5, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemiringRecord_2, "mulRecord"), gopurs_runtime.Constructor0("Proxy"), ra_7, rb_8))
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy")), one1_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_2, "oneRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")))
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy")), zero1_5_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_2, "zeroRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")))
}))
})
	})
	return semiringRecordCons
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

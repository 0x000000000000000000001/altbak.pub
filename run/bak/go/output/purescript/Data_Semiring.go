package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semiring_SemiringRecord_dollarDict gopurs_runtime.Value
var once_Data_Semiring_SemiringRecord_dollarDict sync.Once
func Get_Data_Semiring_SemiringRecord_dollarDict() gopurs_runtime.Value {
	once_Data_Semiring_SemiringRecord_dollarDict.Do(func() {
		cache_Data_Semiring_SemiringRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_SemiringRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_Semiring_SemiringRecord_dollarDict
}

var cache_Data_Semiring_Semiring_dollarDict gopurs_runtime.Value
var once_Data_Semiring_Semiring_dollarDict sync.Once
func Get_Data_Semiring_Semiring_dollarDict() gopurs_runtime.Value {
	once_Data_Semiring_Semiring_dollarDict.Do(func() {
		cache_Data_Semiring_Semiring_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_Semiring_dollarDict(x_0_box)
})
	})
	return cache_Data_Semiring_Semiring_dollarDict
}

var cache_Data_Semiring_zeroRecord gopurs_runtime.Value
var once_Data_Semiring_zeroRecord sync.Once
func Get_Data_Semiring_zeroRecord() gopurs_runtime.Value {
	once_Data_Semiring_zeroRecord.Do(func() {
		cache_Data_Semiring_zeroRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zeroRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_zeroRecord
}

var cache_Data_Semiring_zero gopurs_runtime.Value
var once_Data_Semiring_zero sync.Once
func Get_Data_Semiring_zero() gopurs_runtime.Value {
	once_Data_Semiring_zero.Do(func() {
		cache_Data_Semiring_zero = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero(dict_0_box)
})
	})
	return cache_Data_Semiring_zero
}

var cache_Data_Semiring_semiringUnit gopurs_runtime.Value
var once_Data_Semiring_semiringUnit sync.Once
func Get_Data_Semiring_semiringUnit() gopurs_runtime.Value {
	once_Data_Semiring_semiringUnit.Do(func() {
		cache_Data_Semiring_semiringUnit = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), Get_Data_Unit_unit(), Get_Data_Unit_unit())
	})
	return cache_Data_Semiring_semiringUnit
}

var cache_Data_Semiring_semiringRecordNil gopurs_runtime.Value
var once_Data_Semiring_semiringRecordNil sync.Once
func Get_Data_Semiring_semiringRecordNil() gopurs_runtime.Value {
	once_Data_Semiring_semiringRecordNil.Do(func() {
		cache_Data_Semiring_semiringRecordNil = gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Semiring_semiringRecordNil
}

var cache_Data_Semiring_semiringProxy gopurs_runtime.Value
var once_Data_Semiring_semiringProxy sync.Once
func Get_Data_Semiring_semiringProxy() gopurs_runtime.Value {
	once_Data_Semiring_semiringProxy.Do(func() {
		cache_Data_Semiring_semiringProxy = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
	})
	return cache_Data_Semiring_semiringProxy
}

var cache_Data_Semiring_semiringNumber gopurs_runtime.Value
var once_Data_Semiring_semiringNumber sync.Once
func Get_Data_Semiring_semiringNumber() gopurs_runtime.Value {
	once_Data_Semiring_semiringNumber.Do(func() {
		cache_Data_Semiring_semiringNumber = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
	})
	return cache_Data_Semiring_semiringNumber
}

var cache_Data_Semiring_semiringInt gopurs_runtime.Value
var once_Data_Semiring_semiringInt sync.Once
func Get_Data_Semiring_semiringInt() gopurs_runtime.Value {
	once_Data_Semiring_semiringInt.Do(func() {
		cache_Data_Semiring_semiringInt = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_intAdd(), Get_Data_Semiring_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0))
	})
	return cache_Data_Semiring_semiringInt
}

var cache_Data_Semiring_oneRecord gopurs_runtime.Value
var once_Data_Semiring_oneRecord sync.Once
func Get_Data_Semiring_oneRecord() gopurs_runtime.Value {
	once_Data_Semiring_oneRecord.Do(func() {
		cache_Data_Semiring_oneRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_oneRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_oneRecord
}

var cache_Data_Semiring_one gopurs_runtime.Value
var once_Data_Semiring_one sync.Once
func Get_Data_Semiring_one() gopurs_runtime.Value {
	once_Data_Semiring_one.Do(func() {
		cache_Data_Semiring_one = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_one(dict_0_box)
})
	})
	return cache_Data_Semiring_one
}

var cache_Data_Semiring_mulRecord gopurs_runtime.Value
var once_Data_Semiring_mulRecord sync.Once
func Get_Data_Semiring_mulRecord() gopurs_runtime.Value {
	once_Data_Semiring_mulRecord.Do(func() {
		cache_Data_Semiring_mulRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mulRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_mulRecord
}

var cache_Data_Semiring_mul gopurs_runtime.Value
var once_Data_Semiring_mul sync.Once
func Get_Data_Semiring_mul() gopurs_runtime.Value {
	once_Data_Semiring_mul.Do(func() {
		cache_Data_Semiring_mul = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mul(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_mul
}

var cache_Data_Semiring_addRecord gopurs_runtime.Value
var once_Data_Semiring_addRecord sync.Once
func Get_Data_Semiring_addRecord() gopurs_runtime.Value {
	once_Data_Semiring_addRecord.Do(func() {
		cache_Data_Semiring_addRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_addRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_addRecord
}

var cache_Data_Semiring_semiringRecord gopurs_runtime.Value
var once_Data_Semiring_semiringRecord sync.Once
func Get_Data_Semiring_semiringRecord() gopurs_runtime.Value {
	once_Data_Semiring_semiringRecord.Do(func() {
		cache_Data_Semiring_semiringRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictSemiringRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_semiringRecord(_dollar__unused_0_box, dictSemiringRecord_1_box)
})
	})
	return cache_Data_Semiring_semiringRecord
}

var cache_Data_Semiring_add gopurs_runtime.Value
var once_Data_Semiring_add sync.Once
func Get_Data_Semiring_add() gopurs_runtime.Value {
	once_Data_Semiring_add.Do(func() {
		cache_Data_Semiring_add = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add
}

var cache_Data_Semiring_semiringFn gopurs_runtime.Value
var once_Data_Semiring_semiringFn sync.Once
func Get_Data_Semiring_semiringFn() gopurs_runtime.Value {
	once_Data_Semiring_semiringFn.Do(func() {
		cache_Data_Semiring_semiringFn = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_semiringFn(dictSemiring_0_box)
})
	})
	return cache_Data_Semiring_semiringFn
}

var cache_Data_Semiring_semiringRecordCons gopurs_runtime.Value
var once_Data_Semiring_semiringRecordCons sync.Once
func Get_Data_Semiring_semiringRecordCons() gopurs_runtime.Value {
	once_Data_Semiring_semiringRecordCons.Do(func() {
		cache_Data_Semiring_semiringRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictSemiringRecord_2_box gopurs_runtime.Value, dictSemiring_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_semiringRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictSemiringRecord_2_box, dictSemiring_3_box)
})
	})
	return cache_Data_Semiring_semiringRecordCons
}

var cache_Data_Semiring_add__2927892844 gopurs_runtime.Value
var once_Data_Semiring_add__2927892844 sync.Once
func Get_Data_Semiring_add__2927892844() gopurs_runtime.Value {
	once_Data_Semiring_add__2927892844.Do(func() {
		cache_Data_Semiring_add__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__2927892844(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__2927892844
}

var cache_Data_Semiring_add__101133084 gopurs_runtime.Value
var once_Data_Semiring_add__101133084 sync.Once
func Get_Data_Semiring_add__101133084() gopurs_runtime.Value {
	once_Data_Semiring_add__101133084.Do(func() {
		cache_Data_Semiring_add__101133084 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__101133084(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__101133084
}

var cache_Data_Semiring_add__1124926121 gopurs_runtime.Value
var once_Data_Semiring_add__1124926121 sync.Once
func Get_Data_Semiring_add__1124926121() gopurs_runtime.Value {
	once_Data_Semiring_add__1124926121.Do(func() {
		cache_Data_Semiring_add__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__1124926121(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__1124926121
}

var cache_Data_Semiring_add__1841809173 gopurs_runtime.Value
var once_Data_Semiring_add__1841809173 sync.Once
func Get_Data_Semiring_add__1841809173() gopurs_runtime.Value {
	once_Data_Semiring_add__1841809173.Do(func() {
		cache_Data_Semiring_add__1841809173 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__1841809173(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__1841809173
}

var cache_Data_Semiring_add__190951261 gopurs_runtime.Value
var once_Data_Semiring_add__190951261 sync.Once
func Get_Data_Semiring_add__190951261() gopurs_runtime.Value {
	once_Data_Semiring_add__190951261.Do(func() {
		cache_Data_Semiring_add__190951261 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__190951261(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__190951261
}

var cache_Data_Semiring_add__560788792 gopurs_runtime.Value
var once_Data_Semiring_add__560788792 sync.Once
func Get_Data_Semiring_add__560788792() gopurs_runtime.Value {
	once_Data_Semiring_add__560788792.Do(func() {
		cache_Data_Semiring_add__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semiring_add__560788792
}

var cache_Data_Semiring_add__137136408 gopurs_runtime.Value
var once_Data_Semiring_add__137136408 sync.Once
func Get_Data_Semiring_add__137136408() gopurs_runtime.Value {
	once_Data_Semiring_add__137136408.Do(func() {
		cache_Data_Semiring_add__137136408 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__137136408(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semiring_add__137136408
}

var cache_Data_Semiring_add__113410424 gopurs_runtime.Value
var once_Data_Semiring_add__113410424 sync.Once
func Get_Data_Semiring_add__113410424() gopurs_runtime.Value {
	once_Data_Semiring_add__113410424.Do(func() {
		cache_Data_Semiring_add__113410424 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__113410424(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__113410424
}

var cache_Data_Semiring_add__1614463960 gopurs_runtime.Value
var once_Data_Semiring_add__1614463960 sync.Once
func Get_Data_Semiring_add__1614463960() gopurs_runtime.Value {
	once_Data_Semiring_add__1614463960.Do(func() {
		cache_Data_Semiring_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__1614463960(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__1614463960
}

var cache_Data_Semiring_add__1336467000 gopurs_runtime.Value
var once_Data_Semiring_add__1336467000 sync.Once
func Get_Data_Semiring_add__1336467000() gopurs_runtime.Value {
	once_Data_Semiring_add__1336467000.Do(func() {
		cache_Data_Semiring_add__1336467000 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__1336467000(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__1336467000
}

var cache_Data_Semiring_add__3584309752 gopurs_runtime.Value
var once_Data_Semiring_add__3584309752 sync.Once
func Get_Data_Semiring_add__3584309752() gopurs_runtime.Value {
	once_Data_Semiring_add__3584309752.Do(func() {
		cache_Data_Semiring_add__3584309752 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_add__3584309752(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_add__3584309752
}

var cache_Data_Semiring_addRecord__177134144 gopurs_runtime.Value
var once_Data_Semiring_addRecord__177134144 sync.Once
func Get_Data_Semiring_addRecord__177134144() gopurs_runtime.Value {
	once_Data_Semiring_addRecord__177134144.Do(func() {
		cache_Data_Semiring_addRecord__177134144 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_addRecord__177134144(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_addRecord__177134144
}

var cache_Data_Semiring_addRecord__2348817094 gopurs_runtime.Value
var once_Data_Semiring_addRecord__2348817094 sync.Once
func Get_Data_Semiring_addRecord__2348817094() gopurs_runtime.Value {
	once_Data_Semiring_addRecord__2348817094.Do(func() {
		cache_Data_Semiring_addRecord__2348817094 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_addRecord__2348817094(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_addRecord__2348817094
}

var cache_Data_Semiring_mul__101133084 gopurs_runtime.Value
var once_Data_Semiring_mul__101133084 sync.Once
func Get_Data_Semiring_mul__101133084() gopurs_runtime.Value {
	once_Data_Semiring_mul__101133084.Do(func() {
		cache_Data_Semiring_mul__101133084 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mul__101133084(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_mul__101133084
}

var cache_Data_Semiring_mul__560788792 gopurs_runtime.Value
var once_Data_Semiring_mul__560788792 sync.Once
func Get_Data_Semiring_mul__560788792() gopurs_runtime.Value {
	once_Data_Semiring_mul__560788792.Do(func() {
		cache_Data_Semiring_mul__560788792 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mul__560788792(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semiring_mul__560788792
}

var cache_Data_Semiring_mul__137136408 gopurs_runtime.Value
var once_Data_Semiring_mul__137136408 sync.Once
func Get_Data_Semiring_mul__137136408() gopurs_runtime.Value {
	once_Data_Semiring_mul__137136408.Do(func() {
		cache_Data_Semiring_mul__137136408 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mul__137136408(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semiring_mul__137136408
}

var cache_Data_Semiring_mul__1614463960 gopurs_runtime.Value
var once_Data_Semiring_mul__1614463960 sync.Once
func Get_Data_Semiring_mul__1614463960() gopurs_runtime.Value {
	once_Data_Semiring_mul__1614463960.Do(func() {
		cache_Data_Semiring_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mul__1614463960(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](dict_0_box))
})
	})
	return cache_Data_Semiring_mul__1614463960
}

var cache_Data_Semiring_mulRecord__177134144 gopurs_runtime.Value
var once_Data_Semiring_mulRecord__177134144 sync.Once
func Get_Data_Semiring_mulRecord__177134144() gopurs_runtime.Value {
	once_Data_Semiring_mulRecord__177134144.Do(func() {
		cache_Data_Semiring_mulRecord__177134144 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mulRecord__177134144(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_mulRecord__177134144
}

var cache_Data_Semiring_mulRecord__2348817094 gopurs_runtime.Value
var once_Data_Semiring_mulRecord__2348817094 sync.Once
func Get_Data_Semiring_mulRecord__2348817094() gopurs_runtime.Value {
	once_Data_Semiring_mulRecord__2348817094.Do(func() {
		cache_Data_Semiring_mulRecord__2348817094 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_mulRecord__2348817094(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_mulRecord__2348817094
}

var cache_Data_Semiring_one__1469952505 gopurs_runtime.Value
var once_Data_Semiring_one__1469952505 sync.Once
func Get_Data_Semiring_one__1469952505() gopurs_runtime.Value {
	once_Data_Semiring_one__1469952505.Do(func() {
		cache_Data_Semiring_one__1469952505 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_one__1469952505(dict_0_box)
})
	})
	return cache_Data_Semiring_one__1469952505
}

var cache_Data_Semiring_one__1204848985 gopurs_runtime.Value
var once_Data_Semiring_one__1204848985 sync.Once
func Get_Data_Semiring_one__1204848985() gopurs_runtime.Value {
	once_Data_Semiring_one__1204848985.Do(func() {
		cache_Data_Semiring_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_one__1204848985(dict_0_box)
})
	})
	return cache_Data_Semiring_one__1204848985
}

var cache_Data_Semiring_one__3289155481 gopurs_runtime.Value
var once_Data_Semiring_one__3289155481 sync.Once
func Get_Data_Semiring_one__3289155481() gopurs_runtime.Value {
	once_Data_Semiring_one__3289155481.Do(func() {
		cache_Data_Semiring_one__3289155481 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_one__3289155481(dict_0_box)
})
	})
	return cache_Data_Semiring_one__3289155481
}

var cache_Data_Semiring_oneRecord__283722528 gopurs_runtime.Value
var once_Data_Semiring_oneRecord__283722528 sync.Once
func Get_Data_Semiring_oneRecord__283722528() gopurs_runtime.Value {
	once_Data_Semiring_oneRecord__283722528.Do(func() {
		cache_Data_Semiring_oneRecord__283722528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_oneRecord__283722528(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_oneRecord__283722528
}

var cache_Data_Semiring_oneRecord__2618241440 gopurs_runtime.Value
var once_Data_Semiring_oneRecord__2618241440 sync.Once
func Get_Data_Semiring_oneRecord__2618241440() gopurs_runtime.Value {
	once_Data_Semiring_oneRecord__2618241440.Do(func() {
		cache_Data_Semiring_oneRecord__2618241440 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_oneRecord__2618241440(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_oneRecord__2618241440
}

var cache_Data_Semiring_semiringProxy__3507564775 gopurs_runtime.Value
var once_Data_Semiring_semiringProxy__3507564775 sync.Once
func Get_Data_Semiring_semiringProxy__3507564775() gopurs_runtime.Value {
	once_Data_Semiring_semiringProxy__3507564775.Do(func() {
		cache_Data_Semiring_semiringProxy__3507564775 = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
	})
	return cache_Data_Semiring_semiringProxy__3507564775
}

var cache_Data_Semiring_semiringProxy__786350887 gopurs_runtime.Value
var once_Data_Semiring_semiringProxy__786350887 sync.Once
func Get_Data_Semiring_semiringProxy__786350887() gopurs_runtime.Value {
	once_Data_Semiring_semiringProxy__786350887.Do(func() {
		cache_Data_Semiring_semiringProxy__786350887 = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
	})
	return cache_Data_Semiring_semiringProxy__786350887
}

var cache_Data_Semiring_semiringRecordNil__1274248348 gopurs_runtime.Value
var once_Data_Semiring_semiringRecordNil__1274248348 sync.Once
func Get_Data_Semiring_semiringRecordNil__1274248348() gopurs_runtime.Value {
	once_Data_Semiring_semiringRecordNil__1274248348.Do(func() {
		cache_Data_Semiring_semiringRecordNil__1274248348 = gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Semiring_semiringRecordNil__1274248348
}

var cache_Data_Semiring_semiringRecordNil__3131124892 gopurs_runtime.Value
var once_Data_Semiring_semiringRecordNil__3131124892 sync.Once
func Get_Data_Semiring_semiringRecordNil__3131124892() gopurs_runtime.Value {
	once_Data_Semiring_semiringRecordNil__3131124892.Do(func() {
		cache_Data_Semiring_semiringRecordNil__3131124892 = gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_Data_Semiring_semiringRecordNil__3131124892
}

var cache_Data_Semiring_zero__3961231853 gopurs_runtime.Value
var once_Data_Semiring_zero__3961231853 sync.Once
func Get_Data_Semiring_zero__3961231853() gopurs_runtime.Value {
	once_Data_Semiring_zero__3961231853.Do(func() {
		cache_Data_Semiring_zero__3961231853 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero__3961231853(dict_0_box)
})
	})
	return cache_Data_Semiring_zero__3961231853
}

var cache_Data_Semiring_zero__2621037725 gopurs_runtime.Value
var once_Data_Semiring_zero__2621037725 sync.Once
func Get_Data_Semiring_zero__2621037725() gopurs_runtime.Value {
	once_Data_Semiring_zero__2621037725.Do(func() {
		cache_Data_Semiring_zero__2621037725 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero__2621037725(dict_0_box)
})
	})
	return cache_Data_Semiring_zero__2621037725
}

var cache_Data_Semiring_zero__60838105 gopurs_runtime.Value
var once_Data_Semiring_zero__60838105 sync.Once
func Get_Data_Semiring_zero__60838105() gopurs_runtime.Value {
	once_Data_Semiring_zero__60838105.Do(func() {
		cache_Data_Semiring_zero__60838105 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero__60838105(dict_0_box)
})
	})
	return cache_Data_Semiring_zero__60838105
}

var cache_Data_Semiring_zero__1469952505 gopurs_runtime.Value
var once_Data_Semiring_zero__1469952505 sync.Once
func Get_Data_Semiring_zero__1469952505() gopurs_runtime.Value {
	once_Data_Semiring_zero__1469952505.Do(func() {
		cache_Data_Semiring_zero__1469952505 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero__1469952505(dict_0_box)
})
	})
	return cache_Data_Semiring_zero__1469952505
}

var cache_Data_Semiring_zero__1204848985 gopurs_runtime.Value
var once_Data_Semiring_zero__1204848985 sync.Once
func Get_Data_Semiring_zero__1204848985() gopurs_runtime.Value {
	once_Data_Semiring_zero__1204848985.Do(func() {
		cache_Data_Semiring_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero__1204848985(dict_0_box)
})
	})
	return cache_Data_Semiring_zero__1204848985
}

var cache_Data_Semiring_zero__2603823513 gopurs_runtime.Value
var once_Data_Semiring_zero__2603823513 sync.Once
func Get_Data_Semiring_zero__2603823513() gopurs_runtime.Value {
	once_Data_Semiring_zero__2603823513.Do(func() {
		cache_Data_Semiring_zero__2603823513 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero__2603823513(dict_0_box)
})
	})
	return cache_Data_Semiring_zero__2603823513
}

var cache_Data_Semiring_zero__3289155481 gopurs_runtime.Value
var once_Data_Semiring_zero__3289155481 sync.Once
func Get_Data_Semiring_zero__3289155481() gopurs_runtime.Value {
	once_Data_Semiring_zero__3289155481.Do(func() {
		cache_Data_Semiring_zero__3289155481 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zero__3289155481(dict_0_box)
})
	})
	return cache_Data_Semiring_zero__3289155481
}

var cache_Data_Semiring_zeroRecord__283722528 gopurs_runtime.Value
var once_Data_Semiring_zeroRecord__283722528 sync.Once
func Get_Data_Semiring_zeroRecord__283722528() gopurs_runtime.Value {
	once_Data_Semiring_zeroRecord__283722528.Do(func() {
		cache_Data_Semiring_zeroRecord__283722528 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zeroRecord__283722528(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_zeroRecord__283722528
}

var cache_Data_Semiring_zeroRecord__2618241440 gopurs_runtime.Value
var once_Data_Semiring_zeroRecord__2618241440 sync.Once
func Get_Data_Semiring_zeroRecord__2618241440() gopurs_runtime.Value {
	once_Data_Semiring_zeroRecord__2618241440.Do(func() {
		cache_Data_Semiring_zeroRecord__2618241440 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semiring_zeroRecord__2618241440(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord](dict_0_box))
})
	})
	return cache_Data_Semiring_zeroRecord__2618241440
}

type Constructor_Data_Semiring_SemiringRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3914418263] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semiring_SemiringRecord)(ptr)
		_ = c
		switch key {
		case "addRecord": return gopurs_runtime.Box(c.V0)
		case "mulRecord": return gopurs_runtime.Box(c.V1)
		case "oneRecord": return gopurs_runtime.Box(c.V2)
		case "zeroRecord": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Semiring_SemiringRecord: " + key)
		}
	}
}


type Constructor_Data_Semiring_Semiring struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[134961754] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semiring_Semiring)(ptr)
		_ = c
		switch key {
		case "add": return gopurs_runtime.Box(c.V0)
		case "mul": return gopurs_runtime.Box(c.V1)
		case "one": return gopurs_runtime.Box(c.V2)
		case "zero": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Data_Semiring_Semiring: " + key)
		}
	}
}


func Call_Data_Semiring_SemiringRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semiring_Semiring_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semiring_zeroRecord(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semiring_zero(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_oneRecord(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semiring_one(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_Data_Semiring_mulRecord(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_mul(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_addRecord(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_semiringRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictSemiringRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictSemiringRecord_1 gopurs_runtime.Value = dictSemiringRecord_1_loop
_ = dictSemiringRecord_1
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_Semiring_add(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_semiringFn(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_Data_Semiring_semiringRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictSemiringRecord_2_loop gopurs_runtime.Value, dictSemiring_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictSemiringRecord_2 gopurs_runtime.Value = dictSemiringRecord_2_loop
_ = dictSemiringRecord_2
var dictSemiring_3 gopurs_runtime.Value = dictSemiring_3_loop
_ = dictSemiring_3
// TAST (Let): one1_4_0 -> gopurs_runtime.Value
one1_4_0 := gopurs_runtime.RecordGet(dictSemiring_3, "one")
_ = one1_4_0
// TAST (Let): zero1_5_1 -> gopurs_runtime.Value
zero1_5_1 := gopurs_runtime.RecordGet(dictSemiring_3, "zero")
_ = zero1_5_1
return gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_9_2 -> gopurs_runtime.Value
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_9_2
// TAST (Let): get_10_3 -> gopurs_runtime.Value
get_10_3 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_9_2.StrVal()))
_ = get_10_3
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_2.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_3, "add"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemiringRecord_2, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_9_4 -> gopurs_runtime.Value
key_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_9_4
// TAST (Let): get_10_5 -> gopurs_runtime.Value
get_10_5 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_9_4.StrVal()))
_ = get_10_5
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_4.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_3, "mul"), gopurs_runtime.Apply(get_10_5, ra_7), gopurs_runtime.Apply(get_10_5, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemiringRecord_2, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), one1_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_2, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), zero1_5_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_2, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
})
}))
}

func Call_Data_Semiring_add__2927892844(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__101133084(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__1124926121(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__1841809173(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__190951261(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) + (__eta1_1.IntVal))
}

func Call_Data_Semiring_add__137136408(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) + (__eta1_1.FloatVal()))
}

func Call_Data_Semiring_add__113410424(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__1614463960(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__1336467000(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_add__3584309752(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_addRecord__177134144(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_addRecord__2348817094(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semiring_mul__101133084(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_mul__560788792(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) * (__eta1_1.IntVal))
}

func Call_Data_Semiring_mul__137136408(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Float((__eta0_0.FloatVal()) * (__eta1_1.FloatVal()))
}

func Call_Data_Semiring_mul__1614463960(dict_0_loop *Constructor_Data_Semiring_Semiring) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_Semiring = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_mulRecord__177134144(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_mulRecord__2348817094(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Semiring_one__1469952505(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_Data_Semiring_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_Data_Semiring_one__3289155481(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_Data_Semiring_oneRecord__283722528(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semiring_oneRecord__2618241440(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Semiring_zero__3961231853(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_zero__2621037725(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_zero__60838105(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_zero__1469952505(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_zero__2603823513(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_zero__3289155481(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_Data_Semiring_zeroRecord__283722528(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_Semiring_zeroRecord__2618241440(dict_0_loop *Constructor_Data_Semiring_SemiringRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semiring_SemiringRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Get_Data_Semiring_intAdd() gopurs_runtime.Value {
	return _Gopurs_Data_Semiring_IntAdd
}

func Get_Data_Semiring_intMul() gopurs_runtime.Value {
	return _Gopurs_Data_Semiring_IntMul
}

func Get_Data_Semiring_numAdd() gopurs_runtime.Value {
	return _Gopurs_Data_Semiring_NumAdd
}

func Get_Data_Semiring_numMul() gopurs_runtime.Value {
	return _Gopurs_Data_Semiring_NumMul
}

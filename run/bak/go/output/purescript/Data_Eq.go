package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Eq_EqRecord_dollarDict gopurs_runtime.Value
var once_Data_Eq_EqRecord_dollarDict sync.Once
func Get_Data_Eq_EqRecord_dollarDict() gopurs_runtime.Value {
	once_Data_Eq_EqRecord_dollarDict.Do(func() {
		cache_Data_Eq_EqRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_EqRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_Eq_EqRecord_dollarDict
}

var cache_Data_Eq_Eq_dollarDict gopurs_runtime.Value
var once_Data_Eq_Eq_dollarDict sync.Once
func Get_Data_Eq_Eq_dollarDict() gopurs_runtime.Value {
	once_Data_Eq_Eq_dollarDict.Do(func() {
		cache_Data_Eq_Eq_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Eq_dollarDict(x_0_box)
})
	})
	return cache_Data_Eq_Eq_dollarDict
}

var cache_Data_Eq_Eq1_dollarDict gopurs_runtime.Value
var once_Data_Eq_Eq1_dollarDict sync.Once
func Get_Data_Eq_Eq1_dollarDict() gopurs_runtime.Value {
	once_Data_Eq_Eq1_dollarDict.Do(func() {
		cache_Data_Eq_Eq1_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_Eq1_dollarDict(x_0_box)
})
	})
	return cache_Data_Eq_Eq1_dollarDict
}

var cache_Data_Eq_eqVoid gopurs_runtime.Value
var once_Data_Eq_eqVoid sync.Once
func Get_Data_Eq_eqVoid() gopurs_runtime.Value {
	once_Data_Eq_eqVoid.Do(func() {
		cache_Data_Eq_eqVoid = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Eq_eqVoid
}

var cache_Data_Eq_eqUnit gopurs_runtime.Value
var once_Data_Eq_eqUnit sync.Once
func Get_Data_Eq_eqUnit() gopurs_runtime.Value {
	once_Data_Eq_eqUnit.Do(func() {
		cache_Data_Eq_eqUnit = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Eq_eqUnit
}

var cache_Data_Eq_eqString gopurs_runtime.Value
var once_Data_Eq_eqString sync.Once
func Get_Data_Eq_eqString() gopurs_runtime.Value {
	once_Data_Eq_eqString.Do(func() {
		cache_Data_Eq_eqString = gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqStringImpl())
	})
	return cache_Data_Eq_eqString
}

var cache_Data_Eq_eqRowNil gopurs_runtime.Value
var once_Data_Eq_eqRowNil sync.Once
func Get_Data_Eq_eqRowNil() gopurs_runtime.Value {
	once_Data_Eq_eqRowNil.Do(func() {
		cache_Data_Eq_eqRowNil = gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})
}))
	})
	return cache_Data_Eq_eqRowNil
}

var cache_Data_Eq_eqRecord gopurs_runtime.Value
var once_Data_Eq_eqRecord sync.Once
func Get_Data_Eq_eqRecord() gopurs_runtime.Value {
	once_Data_Eq_eqRecord.Do(func() {
		cache_Data_Eq_eqRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord](dict_0_box))
})
	})
	return cache_Data_Eq_eqRecord
}

var cache_Data_Eq_eqRec gopurs_runtime.Value
var once_Data_Eq_eqRec sync.Once
func Get_Data_Eq_eqRec() gopurs_runtime.Value {
	once_Data_Eq_eqRec.Do(func() {
		cache_Data_Eq_eqRec = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictEqRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRec(_dollar__unused_0_box, dictEqRecord_1_box)
})
	})
	return cache_Data_Eq_eqRec
}

var cache_Data_Eq_eqProxy gopurs_runtime.Value
var once_Data_Eq_eqProxy sync.Once
func Get_Data_Eq_eqProxy() gopurs_runtime.Value {
	once_Data_Eq_eqProxy.Do(func() {
		cache_Data_Eq_eqProxy = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Eq_eqProxy
}

var cache_Data_Eq_eqNumber gopurs_runtime.Value
var once_Data_Eq_eqNumber sync.Once
func Get_Data_Eq_eqNumber() gopurs_runtime.Value {
	once_Data_Eq_eqNumber.Do(func() {
		cache_Data_Eq_eqNumber = gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqNumberImpl())
	})
	return cache_Data_Eq_eqNumber
}

var cache_Data_Eq_eqInt gopurs_runtime.Value
var once_Data_Eq_eqInt sync.Once
func Get_Data_Eq_eqInt() gopurs_runtime.Value {
	once_Data_Eq_eqInt.Do(func() {
		cache_Data_Eq_eqInt = gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
	})
	return cache_Data_Eq_eqInt
}

var cache_Data_Eq_eqChar gopurs_runtime.Value
var once_Data_Eq_eqChar sync.Once
func Get_Data_Eq_eqChar() gopurs_runtime.Value {
	once_Data_Eq_eqChar.Do(func() {
		cache_Data_Eq_eqChar = gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqCharImpl())
	})
	return cache_Data_Eq_eqChar
}

var cache_Data_Eq_eqBoolean gopurs_runtime.Value
var once_Data_Eq_eqBoolean sync.Once
func Get_Data_Eq_eqBoolean() gopurs_runtime.Value {
	once_Data_Eq_eqBoolean.Do(func() {
		cache_Data_Eq_eqBoolean = gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqBooleanImpl())
	})
	return cache_Data_Eq_eqBoolean
}

var cache_Data_Eq_eq1 gopurs_runtime.Value
var once_Data_Eq_eq1 sync.Once
func Get_Data_Eq_eq1() gopurs_runtime.Value {
	once_Data_Eq_eq1.Do(func() {
		cache_Data_Eq_eq1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](dict_0_box))
})
	})
	return cache_Data_Eq_eq1
}

var cache_Data_Eq_eq gopurs_runtime.Value
var once_Data_Eq_eq sync.Once
func Get_Data_Eq_eq() gopurs_runtime.Value {
	once_Data_Eq_eq.Do(func() {
		cache_Data_Eq_eq = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq
}

var cache_Data_Eq_eqArray gopurs_runtime.Value
var once_Data_Eq_eqArray sync.Once
func Get_Data_Eq_eqArray() gopurs_runtime.Value {
	once_Data_Eq_eqArray.Do(func() {
		cache_Data_Eq_eqArray = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqArray(dictEq_0_box)
})
	})
	return cache_Data_Eq_eqArray
}

var cache_Data_Eq_eq1Array gopurs_runtime.Value
var once_Data_Eq_eq1Array sync.Once
func Get_Data_Eq_eq1Array() gopurs_runtime.Value {
	once_Data_Eq_eq1Array.Do(func() {
		cache_Data_Eq_eq1Array = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}))
	})
	return cache_Data_Eq_eq1Array
}

var cache_Data_Eq_eqRowCons gopurs_runtime.Value
var once_Data_Eq_eqRowCons sync.Once
func Get_Data_Eq_eqRowCons() gopurs_runtime.Value {
	once_Data_Eq_eqRowCons.Do(func() {
		cache_Data_Eq_eqRowCons = gopurs_runtime.Func4(func(dictEqRecord_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictIsSymbol_2_box gopurs_runtime.Value, dictEq_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRowCons(dictEqRecord_0_box, _dollar__unused_1_box, dictIsSymbol_2_box, dictEq_3_box)
})
	})
	return cache_Data_Eq_eqRowCons
}

var cache_Data_Eq_notEq gopurs_runtime.Value
var once_Data_Eq_notEq sync.Once
func Get_Data_Eq_notEq() gopurs_runtime.Value {
	once_Data_Eq_notEq.Do(func() {
		cache_Data_Eq_notEq = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_Data_Eq_notEq
}

var cache_Data_Eq_notEq1 gopurs_runtime.Value
var once_Data_Eq_notEq1 sync.Once
func Get_Data_Eq_notEq1() gopurs_runtime.Value {
	once_Data_Eq_notEq1.Do(func() {
		cache_Data_Eq_notEq1 = gopurs_runtime.Func4(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](dictEq1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_1_box), x_2_box, y_3_box))
})
	})
	return cache_Data_Eq_notEq1
}

var cache_Data_Eq_eq__1710332219 gopurs_runtime.Value
var once_Data_Eq_eq__1710332219 sync.Once
func Get_Data_Eq_eq__1710332219() gopurs_runtime.Value {
	once_Data_Eq_eq__1710332219.Do(func() {
		cache_Data_Eq_eq__1710332219 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__1710332219(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__1710332219
}

var cache_Data_Eq_eq__1559309819 gopurs_runtime.Value
var once_Data_Eq_eq__1559309819 sync.Once
func Get_Data_Eq_eq__1559309819() gopurs_runtime.Value {
	once_Data_Eq_eq__1559309819.Do(func() {
		cache_Data_Eq_eq__1559309819 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__1559309819(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__1559309819
}

var cache_Data_Eq_eq__3259097883 gopurs_runtime.Value
var once_Data_Eq_eq__3259097883 sync.Once
func Get_Data_Eq_eq__3259097883() gopurs_runtime.Value {
	once_Data_Eq_eq__3259097883.Do(func() {
		cache_Data_Eq_eq__3259097883 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__3259097883(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__3259097883
}

var cache_Data_Eq_eq__874329115 gopurs_runtime.Value
var once_Data_Eq_eq__874329115 sync.Once
func Get_Data_Eq_eq__874329115() gopurs_runtime.Value {
	once_Data_Eq_eq__874329115.Do(func() {
		cache_Data_Eq_eq__874329115 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__874329115(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__874329115
}

var cache_Data_Eq_eq__3621906651 gopurs_runtime.Value
var once_Data_Eq_eq__3621906651 sync.Once
func Get_Data_Eq_eq__3621906651() gopurs_runtime.Value {
	once_Data_Eq_eq__3621906651.Do(func() {
		cache_Data_Eq_eq__3621906651 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__3621906651(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__3621906651
}

var cache_Data_Eq_eq__2843686287 gopurs_runtime.Value
var once_Data_Eq_eq__2843686287 sync.Once
func Get_Data_Eq_eq__2843686287() gopurs_runtime.Value {
	once_Data_Eq_eq__2843686287.Do(func() {
		cache_Data_Eq_eq__2843686287 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2843686287(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Eq_eq__2843686287
}

var cache_Data_Eq_eq__472317769 gopurs_runtime.Value
var once_Data_Eq_eq__472317769 sync.Once
func Get_Data_Eq_eq__472317769() gopurs_runtime.Value {
	once_Data_Eq_eq__472317769.Do(func() {
		cache_Data_Eq_eq__472317769 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__472317769(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Eq_eq__472317769
}

var cache_Data_Eq_eq__2276491096 gopurs_runtime.Value
var once_Data_Eq_eq__2276491096 sync.Once
func Get_Data_Eq_eq__2276491096() gopurs_runtime.Value {
	once_Data_Eq_eq__2276491096.Do(func() {
		cache_Data_Eq_eq__2276491096 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2276491096(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Eq_eq__2276491096
}

var cache_Data_Eq_eq__2384498378 gopurs_runtime.Value
var once_Data_Eq_eq__2384498378 sync.Once
func Get_Data_Eq_eq__2384498378() gopurs_runtime.Value {
	once_Data_Eq_eq__2384498378.Do(func() {
		cache_Data_Eq_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2384498378(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__2384498378
}

var cache_Data_Eq_eq__3977211983 gopurs_runtime.Value
var once_Data_Eq_eq__3977211983 sync.Once
func Get_Data_Eq_eq__3977211983() gopurs_runtime.Value {
	once_Data_Eq_eq__3977211983.Do(func() {
		cache_Data_Eq_eq__3977211983 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__3977211983(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Eq_eq__3977211983
}

var cache_Data_Eq_eq__3887832182 gopurs_runtime.Value
var once_Data_Eq_eq__3887832182 sync.Once
func Get_Data_Eq_eq__3887832182() gopurs_runtime.Value {
	once_Data_Eq_eq__3887832182.Do(func() {
		cache_Data_Eq_eq__3887832182 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_eq__3887832182(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal)))
})
	})
	return cache_Data_Eq_eq__3887832182
}

var cache_Data_Eq_eq__1204755874 gopurs_runtime.Value
var once_Data_Eq_eq__1204755874 sync.Once
func Get_Data_Eq_eq__1204755874() gopurs_runtime.Value {
	once_Data_Eq_eq__1204755874.Do(func() {
		cache_Data_Eq_eq__1204755874 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_eq__1204755874(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](x_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](y_1_box)))
})
	})
	return cache_Data_Eq_eq__1204755874
}

var cache_Data_Eq_eq__196302102 gopurs_runtime.Value
var once_Data_Eq_eq__196302102 sync.Once
func Get_Data_Eq_eq__196302102() gopurs_runtime.Value {
	once_Data_Eq_eq__196302102.Do(func() {
		cache_Data_Eq_eq__196302102 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_eq__196302102(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal)))
})
	})
	return cache_Data_Eq_eq__196302102
}

var cache_Data_Eq_eq__1241439021 gopurs_runtime.Value
var once_Data_Eq_eq__1241439021 sync.Once
func Get_Data_Eq_eq__1241439021() gopurs_runtime.Value {
	once_Data_Eq_eq__1241439021.Do(func() {
		cache_Data_Eq_eq__1241439021 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_eq__1241439021(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal)))
})
	})
	return cache_Data_Eq_eq__1241439021
}

var cache_Data_Eq_eq__501078914 gopurs_runtime.Value
var once_Data_Eq_eq__501078914 sync.Once
func Get_Data_Eq_eq__501078914() gopurs_runtime.Value {
	once_Data_Eq_eq__501078914.Do(func() {
		cache_Data_Eq_eq__501078914 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__501078914(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__501078914
}

var cache_Data_Eq_eq__2484408063 gopurs_runtime.Value
var once_Data_Eq_eq__2484408063 sync.Once
func Get_Data_Eq_eq__2484408063() gopurs_runtime.Value {
	once_Data_Eq_eq__2484408063.Do(func() {
		cache_Data_Eq_eq__2484408063 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2484408063(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__2484408063
}

var cache_Data_Eq_eq__1425708671 gopurs_runtime.Value
var once_Data_Eq_eq__1425708671 sync.Once
func Get_Data_Eq_eq__1425708671() gopurs_runtime.Value {
	once_Data_Eq_eq__1425708671.Do(func() {
		cache_Data_Eq_eq__1425708671 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__1425708671(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__1425708671
}

var cache_Data_Eq_eq__3293889322 gopurs_runtime.Value
var once_Data_Eq_eq__3293889322 sync.Once
func Get_Data_Eq_eq__3293889322() gopurs_runtime.Value {
	once_Data_Eq_eq__3293889322.Do(func() {
		cache_Data_Eq_eq__3293889322 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__3293889322(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__3293889322
}

var cache_Data_Eq_eq__2224314568 gopurs_runtime.Value
var once_Data_Eq_eq__2224314568 sync.Once
func Get_Data_Eq_eq__2224314568() gopurs_runtime.Value {
	once_Data_Eq_eq__2224314568.Do(func() {
		cache_Data_Eq_eq__2224314568 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2224314568(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Eq_eq__2224314568
}

var cache_Data_Eq_eq__2541178864 gopurs_runtime.Value
var once_Data_Eq_eq__2541178864 sync.Once
func Get_Data_Eq_eq__2541178864() gopurs_runtime.Value {
	once_Data_Eq_eq__2541178864.Do(func() {
		cache_Data_Eq_eq__2541178864 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2541178864(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__2541178864
}

var cache_Data_Eq_eq__3433516078 gopurs_runtime.Value
var once_Data_Eq_eq__3433516078 sync.Once
func Get_Data_Eq_eq__3433516078() gopurs_runtime.Value {
	once_Data_Eq_eq__3433516078.Do(func() {
		cache_Data_Eq_eq__3433516078 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__3433516078(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Eq_eq__3433516078
}

var cache_Data_Eq_eq__1272715810 gopurs_runtime.Value
var once_Data_Eq_eq__1272715810 sync.Once
func Get_Data_Eq_eq__1272715810() gopurs_runtime.Value {
	once_Data_Eq_eq__1272715810.Do(func() {
		cache_Data_Eq_eq__1272715810 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_eq__1272715810(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal)))
})
	})
	return cache_Data_Eq_eq__1272715810
}

var cache_Data_Eq_eq__1287514754 gopurs_runtime.Value
var once_Data_Eq_eq__1287514754 sync.Once
func Get_Data_Eq_eq__1287514754() gopurs_runtime.Value {
	once_Data_Eq_eq__1287514754.Do(func() {
		cache_Data_Eq_eq__1287514754 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_eq__1287514754(gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](x_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Time_Time](y_1_box)))
})
	})
	return cache_Data_Eq_eq__1287514754
}

var cache_Data_Eq_eq__163522700 gopurs_runtime.Value
var once_Data_Eq_eq__163522700 sync.Once
func Get_Data_Eq_eq__163522700() gopurs_runtime.Value {
	once_Data_Eq_eq__163522700.Do(func() {
		cache_Data_Eq_eq__163522700 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__163522700(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dict_0_box))
})
	})
	return cache_Data_Eq_eq__163522700
}

var cache_Data_Eq_eq1__1773593252 gopurs_runtime.Value
var once_Data_Eq_eq1__1773593252 sync.Once
func Get_Data_Eq_eq1__1773593252() gopurs_runtime.Value {
	once_Data_Eq_eq1__1773593252.Do(func() {
		cache_Data_Eq_eq1__1773593252 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq1__1773593252(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](dict_0_box))
})
	})
	return cache_Data_Eq_eq1__1773593252
}

var cache_Data_Eq_eq1__3199040333 gopurs_runtime.Value
var once_Data_Eq_eq1__3199040333 sync.Once
func Get_Data_Eq_eq1__3199040333() gopurs_runtime.Value {
	once_Data_Eq_eq1__3199040333.Do(func() {
		cache_Data_Eq_eq1__3199040333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq1__3199040333(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](dict_0_box))
})
	})
	return cache_Data_Eq_eq1__3199040333
}

var cache_Data_Eq_eq1__2184765036 gopurs_runtime.Value
var once_Data_Eq_eq1__2184765036 sync.Once
func Get_Data_Eq_eq1__2184765036() gopurs_runtime.Value {
	once_Data_Eq_eq1__2184765036.Do(func() {
		cache_Data_Eq_eq1__2184765036 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq1__2184765036(dictEq_0_box)
})
	})
	return cache_Data_Eq_eq1__2184765036
}

var cache_Data_Eq_eq1__3587165073 gopurs_runtime.Value
var once_Data_Eq_eq1__3587165073 sync.Once
func Get_Data_Eq_eq1__3587165073() gopurs_runtime.Value {
	once_Data_Eq_eq1__3587165073.Do(func() {
		cache_Data_Eq_eq1__3587165073 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq1__3587165073(dictEq_0_box)
})
	})
	return cache_Data_Eq_eq1__3587165073
}

var cache_Data_Eq_eq1Array__2389734302 gopurs_runtime.Value
var once_Data_Eq_eq1Array__2389734302 sync.Once
func Get_Data_Eq_eq1Array__2389734302() gopurs_runtime.Value {
	once_Data_Eq_eq1Array__2389734302.Do(func() {
		cache_Data_Eq_eq1Array__2389734302 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}))
	})
	return cache_Data_Eq_eq1Array__2389734302
}

var cache_Data_Eq_eqProxy__2316826405 gopurs_runtime.Value
var once_Data_Eq_eqProxy__2316826405 sync.Once
func Get_Data_Eq_eqProxy__2316826405() gopurs_runtime.Value {
	once_Data_Eq_eqProxy__2316826405.Do(func() {
		cache_Data_Eq_eqProxy__2316826405 = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Eq_eqProxy__2316826405
}

var cache_Data_Eq_eqProxy__2077298405 gopurs_runtime.Value
var once_Data_Eq_eqProxy__2077298405 sync.Once
func Get_Data_Eq_eqProxy__2077298405() gopurs_runtime.Value {
	once_Data_Eq_eqProxy__2077298405.Do(func() {
		cache_Data_Eq_eqProxy__2077298405 = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
}))
	})
	return cache_Data_Eq_eqProxy__2077298405
}

var cache_Data_Eq_eqRecord__1610867122 gopurs_runtime.Value
var once_Data_Eq_eqRecord__1610867122 sync.Once
func Get_Data_Eq_eqRecord__1610867122() gopurs_runtime.Value {
	once_Data_Eq_eqRecord__1610867122.Do(func() {
		cache_Data_Eq_eqRecord__1610867122 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRecord__1610867122(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord](dict_0_box))
})
	})
	return cache_Data_Eq_eqRecord__1610867122
}

var cache_Data_Eq_eqRecord__1747372340 gopurs_runtime.Value
var once_Data_Eq_eqRecord__1747372340 sync.Once
func Get_Data_Eq_eqRecord__1747372340() gopurs_runtime.Value {
	once_Data_Eq_eqRecord__1747372340.Do(func() {
		cache_Data_Eq_eqRecord__1747372340 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRecord__1747372340(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord](dict_0_box))
})
	})
	return cache_Data_Eq_eqRecord__1747372340
}

var cache_Data_Eq_eqRowNil__3458192115 gopurs_runtime.Value
var once_Data_Eq_eqRowNil__3458192115 sync.Once
func Get_Data_Eq_eqRowNil__3458192115() gopurs_runtime.Value {
	once_Data_Eq_eqRowNil__3458192115.Do(func() {
		cache_Data_Eq_eqRowNil__3458192115 = gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})
}))
	})
	return cache_Data_Eq_eqRowNil__3458192115
}

var cache_Data_Eq_eqRowNil__2234696885 gopurs_runtime.Value
var once_Data_Eq_eqRowNil__2234696885 sync.Once
func Get_Data_Eq_eqRowNil__2234696885() gopurs_runtime.Value {
	once_Data_Eq_eqRowNil__2234696885.Do(func() {
		cache_Data_Eq_eqRowNil__2234696885 = gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})
})
}))
	})
	return cache_Data_Eq_eqRowNil__2234696885
}

var cache_Data_Eq_notEq__2843686287 gopurs_runtime.Value
var once_Data_Eq_notEq__2843686287 sync.Once
func Get_Data_Eq_notEq__2843686287() gopurs_runtime.Value {
	once_Data_Eq_notEq__2843686287.Do(func() {
		cache_Data_Eq_notEq__2843686287 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__2843686287(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_Data_Eq_notEq__2843686287
}

var cache_Data_Eq_notEq__2334967935 gopurs_runtime.Value
var once_Data_Eq_notEq__2334967935 sync.Once
func Get_Data_Eq_notEq__2334967935() gopurs_runtime.Value {
	once_Data_Eq_notEq__2334967935.Do(func() {
		cache_Data_Eq_notEq__2334967935 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__2334967935(x_0_box.FloatVal(), y_1_box.FloatVal()))
})
	})
	return cache_Data_Eq_notEq__2334967935
}

var cache_Data_Eq_notEq__2384498378 gopurs_runtime.Value
var once_Data_Eq_notEq__2384498378 sync.Once
func Get_Data_Eq_notEq__2384498378() gopurs_runtime.Value {
	once_Data_Eq_notEq__2384498378.Do(func() {
		cache_Data_Eq_notEq__2384498378 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__2384498378(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_Data_Eq_notEq__2384498378
}

var cache_Data_Eq_notEq__1272715810 gopurs_runtime.Value
var once_Data_Eq_notEq__1272715810 sync.Once
func Get_Data_Eq_notEq__1272715810() gopurs_runtime.Value {
	once_Data_Eq_notEq__1272715810.Do(func() {
		cache_Data_Eq_notEq__1272715810 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__1272715810(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal)))
})
	})
	return cache_Data_Eq_notEq__1272715810
}

type Constructor_Data_Eq_EqRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1311326743] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Eq_EqRecord)(ptr)
		_ = c
		switch key {
		case "eqRecord": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Eq_EqRecord: " + key)
		}
	}
}


type Constructor_Data_Eq_Eq struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1012063514] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Eq_Eq)(ptr)
		_ = c
		switch key {
		case "eq": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Eq_Eq: " + key)
		}
	}
}


type Constructor_Data_Eq_Eq1 struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1715248107] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Eq_Eq1)(ptr)
		_ = c
		switch key {
		case "eq1": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Eq_Eq1: " + key)
		}
	}
}


func Call_Data_Eq_EqRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Eq_Eq_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Eq_Eq1_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Eq_eqRecord(dict_0_loop *Constructor_Data_Eq_EqRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_EqRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eqRec(_dollar__unused_0_loop gopurs_runtime.Value, dictEqRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictEqRecord_1 gopurs_runtime.Value = dictEqRecord_1_loop
_ = dictEqRecord_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEqRecord_1, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_Eq_eq1(dict_0_loop *Constructor_Data_Eq_Eq1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eqArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq")))
}

func Call_Data_Eq_eqRowCons(dictEqRecord_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictIsSymbol_2_loop gopurs_runtime.Value, dictEq_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEqRecord_0 gopurs_runtime.Value = dictEqRecord_0_loop
_ = dictEqRecord_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictIsSymbol_2 gopurs_runtime.Value = dictIsSymbol_2_loop
_ = dictIsSymbol_2
var dictEq_3 gopurs_runtime.Value = dictEq_3_loop
_ = dictEq_3
return gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_7_0 -> gopurs_runtime.Value
get_7_0 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_2, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()))
_ = get_7_0
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_3, "eq"), gopurs_runtime.Apply(get_7_0, ra_5), gopurs_runtime.Apply(get_7_0, rb_6)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEqRecord_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
})
})
}))
}

func Call_Data_Eq_notEq(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return (Call_Data_Eq_eq__2276491096(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), x_1, y_2).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_Data_Eq_notEq1(dictEq1_0_loop *Constructor_Data_Eq_Eq1, dictEq_1_loop *Constructor_Data_Eq_Eq, x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var dictEq1_0 *Constructor_Data_Eq_Eq1 = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 *Constructor_Data_Eq_Eq = dictEq_1_loop
_ = dictEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return (Call_Data_Eq_eq__2276491096(gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.Box(dictEq1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_1)}, x_2, y_3).IntVal) != (0)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_Data_Eq_eq__1710332219(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__1559309819(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__3259097883(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__874329115(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__3621906651(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__2843686287(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.IntVal) == (__eta1_1.IntVal))
}

func Call_Data_Eq_eq__472317769(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.StrVal()) == (__eta1_1.StrVal()))
}

func Call_Data_Eq_eq__2276491096(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) == ((__eta1_1.IntVal) != (0)))
}

func Call_Data_Eq_eq__2384498378(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__3977211983(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_String_Regex_Flags_eqArray()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Eq_eq__3887832182(x_0_loop uint32, y_1_loop uint32) bool {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t11 bool
{
if (x_0 == 1908470532) {
var __t0 bool
{
if (y_1 == 1908470532) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t11 = __t0
goto end_branch_11
} else {

}
}
{
if (x_0 == 2455627378) {
var __t1 bool
{
if (y_1 == 2455627378) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t11 = __t1
goto end_branch_11
} else {

}
}
{
if (x_0 == 4162469099) {
var __t2 bool
{
if (y_1 == 4162469099) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t11 = __t2
goto end_branch_11
} else {

}
}
{
if (x_0 == 1692989816) {
var __t3 bool
{
if (y_1 == 1692989816) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t11 = __t3
goto end_branch_11
} else {

}
}
{
if (x_0 == 330658827) {
var __t4 bool
{
if (y_1 == 330658827) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t11 = __t4
goto end_branch_11
} else {

}
}
{
if (x_0 == 4067355978) {
var __t5 bool
{
if (y_1 == 4067355978) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t11 = __t5
goto end_branch_11
} else {

}
}
{
if (x_0 == 2276710548) {
var __t6 bool
{
if (y_1 == 2276710548) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t11 = __t6
goto end_branch_11
} else {

}
}
{
if (x_0 == 243771071) {
var __t7 bool
{
if (y_1 == 243771071) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t11 = __t7
goto end_branch_11
} else {

}
}
{
if (x_0 == 215731793) {
var __t8 bool
{
if (y_1 == 215731793) {
__t8 = true
goto end_branch_8
} else {

}
}
{
__t8 = false
}
end_branch_8:
__t11 = __t8
goto end_branch_11
} else {

}
}
{
if (x_0 == 8639228) {
var __t9 bool
{
if (y_1 == 8639228) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
__t11 = __t9
goto end_branch_11
} else {

}
}
{
if (x_0 == 49471444) {
var __t10 bool
{
if (y_1 == 49471444) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t11 = __t10
goto end_branch_11
} else {

}
}
{
if ((x_0 == 3889233761)) && ((y_1 == 3889233761)) {
__t11 = true
goto end_branch_11
} else {

}
}
{
__t11 = false
}
end_branch_11:
return __t11
}

func Call_Data_Eq_eq__1204755874(x_0_loop *Constructor_Data_Date_Date, y_1_loop *Constructor_Data_Date_Date) bool {
var x_0 *Constructor_Data_Date_Date = x_0_loop
_ = x_0
var y_1 *Constructor_Data_Date_Date = y_1_loop
_ = y_1
return (((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Date_Component_eqYear(), "eq"), gopurs_runtime.Int((x_0).V0), gopurs_runtime.Int((y_1).V0)).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Date_Component_eqMonth(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64((x_0).V1), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((y_1).V1), UnsafePtr: nil}).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Date_Component_eqDay(), "eq"), gopurs_runtime.Int((x_0).V2), gopurs_runtime.Int((y_1).V2)).IntVal) != (0))
}

func Call_Data_Eq_eq__196302102(x_0_loop uint32, y_1_loop uint32) bool {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t1 bool
{
if (x_0 == 2591059121) {
var __t0 bool
{
if (y_1 == 2591059121) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t1 = __t0
goto end_branch_1
} else {

}
}
{
if ((x_0 == 658452902)) && ((y_1 == 658452902)) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_Data_Eq_eq__1241439021(x_0_loop uint32, y_1_loop uint32) bool {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t6 bool
{
if (x_0 == 3908053364) {
var __t0 bool
{
if (y_1 == 3908053364) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t6 = __t0
goto end_branch_6
} else {

}
}
{
if (x_0 == 217821258) {
var __t1 bool
{
if (y_1 == 217821258) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t6 = __t1
goto end_branch_6
} else {

}
}
{
if (x_0 == 1292308612) {
var __t2 bool
{
if (y_1 == 1292308612) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
if (x_0 == 2311060696) {
var __t3 bool
{
if (y_1 == 2311060696) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (x_0 == 401302776) {
var __t4 bool
{
if (y_1 == 401302776) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
if (x_0 == 3327533908) {
var __t5 bool
{
if (y_1 == 3327533908) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if ((x_0 == 3631736139)) && ((y_1 == 3631736139)) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
return __t6
}

func Call_Data_Eq_eq__501078914(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__2484408063(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__1425708671(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__3293889322(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__2224314568(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_Duration_eqMap()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Eq_eq__2541178864(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__3433516078(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_eqMaybe()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Eq_eq__1272715810(v_0_loop uint32, v1_1_loop uint32) bool {
var v_0 uint32 = v_0_loop
_ = v_0
var v1_1 uint32 = v1_1_loop
_ = v1_1
var __t2 bool
{
if (v_0 == 1527465420) {
var __t0 bool
{
if (v1_1 == 1527465420) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t2 = __t0
goto end_branch_2
} else {

}
}
{
if (v_0 == 380165415) {
var __t1 bool
{
if (v1_1 == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if ((v_0 == 902936544)) && ((v1_1 == 902936544)) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return __t2
}

func Call_Data_Eq_eq__1287514754(x_0_loop *Constructor_Data_Time_Time, y_1_loop *Constructor_Data_Time_Time) bool {
var x_0 *Constructor_Data_Time_Time = x_0_loop
_ = x_0
var y_1 *Constructor_Data_Time_Time = y_1_loop
_ = y_1
return ((((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Time_Component_eqHour(), "eq"), gopurs_runtime.Int((x_0).V0), gopurs_runtime.Int((y_1).V0)).IntVal) != (0)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Time_Component_eqMinute(), "eq"), gopurs_runtime.Int((x_0).V1), gopurs_runtime.Int((y_1).V1)).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Time_Component_eqSecond(), "eq"), gopurs_runtime.Int((x_0).V2), gopurs_runtime.Int((y_1).V2)).IntVal) != (0))) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Time_Component_eqMillisecond(), "eq"), gopurs_runtime.Int((x_0).V3), gopurs_runtime.Int((y_1).V3)).IntVal) != (0))
}

func Call_Data_Eq_eq__163522700(dict_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq1__1773593252(dict_0_loop *Constructor_Data_Eq_Eq1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq1__3199040333(dict_0_loop *Constructor_Data_Eq_Eq1) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq1 = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq1__2184765036(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_Lazy_eqLazy(), dictEq_0), "eq")
}

func Call_Data_Eq_eq1__3587165073(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_Data_Lazy_eqLazy(), dictEq_0), "eq")
}

func Call_Data_Eq_eqRecord__1610867122(dict_0_loop *Constructor_Data_Eq_EqRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_EqRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eqRecord__1747372340(dict_0_loop *Constructor_Data_Eq_EqRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_EqRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_notEq__2843686287(x_0_loop int64, y_1_loop int64) bool {
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
return ((x_0) == (y_1)) != (true)
}

func Call_Data_Eq_notEq__2334967935(x_0_loop float64, y_1_loop float64) bool {
var x_0 float64 = x_0_loop
_ = x_0
var y_1 float64 = y_1_loop
_ = y_1
return ((x_0) == (y_1)) != (true)
}

func Call_Data_Eq_notEq__2384498378(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), x_1, y_2).IntVal) != (0)) != (true)
}

func Call_Data_Eq_notEq__1272715810(x_0_loop uint32, y_1_loop uint32) bool {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
return ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ordering_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(x_0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(y_1), UnsafePtr: nil}).IntVal) != (0)) != (true)
}

func Get_Data_Eq_eqArrayImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Eq_EqArrayImpl
}

func Get_Data_Eq_eqBooleanImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Eq_EqBooleanImpl
}

func Get_Data_Eq_eqCharImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Eq_EqCharImpl
}

func Get_Data_Eq_eqIntImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Eq_EqIntImpl
}

func Get_Data_Eq_eqNumberImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Eq_EqNumberImpl
}

func Get_Data_Eq_eqStringImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Eq_EqStringImpl
}

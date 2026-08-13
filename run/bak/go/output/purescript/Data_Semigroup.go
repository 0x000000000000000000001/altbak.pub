package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Semigroup_SemigroupRecord_dollarDict gopurs_runtime.Value
var once_Data_Semigroup_SemigroupRecord_dollarDict sync.Once
func Get_Data_Semigroup_SemigroupRecord_dollarDict() gopurs_runtime.Value {
	once_Data_Semigroup_SemigroupRecord_dollarDict.Do(func() {
		cache_Data_Semigroup_SemigroupRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_SemigroupRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_Semigroup_SemigroupRecord_dollarDict
}

var cache_Data_Semigroup_Semigroup_dollarDict gopurs_runtime.Value
var once_Data_Semigroup_Semigroup_dollarDict sync.Once
func Get_Data_Semigroup_Semigroup_dollarDict() gopurs_runtime.Value {
	once_Data_Semigroup_Semigroup_dollarDict.Do(func() {
		cache_Data_Semigroup_Semigroup_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_Semigroup_dollarDict(x_0_box)
})
	})
	return cache_Data_Semigroup_Semigroup_dollarDict
}

var cache_Data_Semigroup_semigroupVoid gopurs_runtime.Value
var once_Data_Semigroup_semigroupVoid sync.Once
func Get_Data_Semigroup_semigroupVoid() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupVoid.Do(func() {
		cache_Data_Semigroup_semigroupVoid = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Void_absurd()
}))
	})
	return cache_Data_Semigroup_semigroupVoid
}

var cache_Data_Semigroup_semigroupUnit gopurs_runtime.Value
var once_Data_Semigroup_semigroupUnit sync.Once
func Get_Data_Semigroup_semigroupUnit() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupUnit.Do(func() {
		cache_Data_Semigroup_semigroupUnit = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}))
	})
	return cache_Data_Semigroup_semigroupUnit
}

var cache_Data_Semigroup_semigroupString gopurs_runtime.Value
var once_Data_Semigroup_semigroupString sync.Once
func Get_Data_Semigroup_semigroupString() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupString.Do(func() {
		cache_Data_Semigroup_semigroupString = gopurs_runtime.RecordDict1("append", Get_Data_Semigroup_concatString())
	})
	return cache_Data_Semigroup_semigroupString
}

var cache_Data_Semigroup_semigroupRecordNil gopurs_runtime.Value
var once_Data_Semigroup_semigroupRecordNil sync.Once
func Get_Data_Semigroup_semigroupRecordNil() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupRecordNil.Do(func() {
		cache_Data_Semigroup_semigroupRecordNil = gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_Data_Semigroup_semigroupRecordNil
}

var cache_Data_Semigroup_semigroupProxy gopurs_runtime.Value
var once_Data_Semigroup_semigroupProxy sync.Once
func Get_Data_Semigroup_semigroupProxy() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupProxy.Do(func() {
		cache_Data_Semigroup_semigroupProxy = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Semigroup_semigroupProxy
}

var cache_Data_Semigroup_semigroupArray gopurs_runtime.Value
var once_Data_Semigroup_semigroupArray sync.Once
func Get_Data_Semigroup_semigroupArray() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupArray.Do(func() {
		cache_Data_Semigroup_semigroupArray = gopurs_runtime.RecordDict1("append", Get_Data_Semigroup_concatArray())
	})
	return cache_Data_Semigroup_semigroupArray
}

var cache_Data_Semigroup_appendRecord gopurs_runtime.Value
var once_Data_Semigroup_appendRecord sync.Once
func Get_Data_Semigroup_appendRecord() gopurs_runtime.Value {
	once_Data_Semigroup_appendRecord.Do(func() {
		cache_Data_Semigroup_appendRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_appendRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord](dict_0_box))
})
	})
	return cache_Data_Semigroup_appendRecord
}

var cache_Data_Semigroup_semigroupRecord gopurs_runtime.Value
var once_Data_Semigroup_semigroupRecord sync.Once
func Get_Data_Semigroup_semigroupRecord() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupRecord.Do(func() {
		cache_Data_Semigroup_semigroupRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictSemigroupRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_semigroupRecord(_dollar__unused_0_box, dictSemigroupRecord_1_box)
})
	})
	return cache_Data_Semigroup_semigroupRecord
}

var cache_Data_Semigroup_append gopurs_runtime.Value
var once_Data_Semigroup_append sync.Once
func Get_Data_Semigroup_append() gopurs_runtime.Value {
	once_Data_Semigroup_append.Do(func() {
		cache_Data_Semigroup_append = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append
}

var cache_Data_Semigroup_semigroupFn gopurs_runtime.Value
var once_Data_Semigroup_semigroupFn sync.Once
func Get_Data_Semigroup_semigroupFn() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupFn.Do(func() {
		cache_Data_Semigroup_semigroupFn = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_semigroupFn(dictSemigroup_0_box)
})
	})
	return cache_Data_Semigroup_semigroupFn
}

var cache_Data_Semigroup_semigroupRecordCons gopurs_runtime.Value
var once_Data_Semigroup_semigroupRecordCons sync.Once
func Get_Data_Semigroup_semigroupRecordCons() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupRecordCons.Do(func() {
		cache_Data_Semigroup_semigroupRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictSemigroupRecord_2_box gopurs_runtime.Value, dictSemigroup_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_semigroupRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictSemigroupRecord_2_box, dictSemigroup_3_box)
})
	})
	return cache_Data_Semigroup_semigroupRecordCons
}

var cache_Data_Semigroup_append__2637663146 gopurs_runtime.Value
var once_Data_Semigroup_append__2637663146 sync.Once
func Get_Data_Semigroup_append__2637663146() gopurs_runtime.Value {
	once_Data_Semigroup_append__2637663146.Do(func() {
		cache_Data_Semigroup_append__2637663146 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2637663146(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__2637663146
}

var cache_Data_Semigroup_append__1124926121 gopurs_runtime.Value
var once_Data_Semigroup_append__1124926121 sync.Once
func Get_Data_Semigroup_append__1124926121() gopurs_runtime.Value {
	once_Data_Semigroup_append__1124926121.Do(func() {
		cache_Data_Semigroup_append__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__1124926121(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__1124926121
}

var cache_Data_Semigroup_append__4093645121 gopurs_runtime.Value
var once_Data_Semigroup_append__4093645121 sync.Once
func Get_Data_Semigroup_append__4093645121() gopurs_runtime.Value {
	once_Data_Semigroup_append__4093645121.Do(func() {
		cache_Data_Semigroup_append__4093645121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__4093645121(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__4093645121
}

var cache_Data_Semigroup_append__2462288412 gopurs_runtime.Value
var once_Data_Semigroup_append__2462288412 sync.Once
func Get_Data_Semigroup_append__2462288412() gopurs_runtime.Value {
	once_Data_Semigroup_append__2462288412.Do(func() {
		cache_Data_Semigroup_append__2462288412 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2462288412(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__2462288412
}

var cache_Data_Semigroup_append__1442818457 gopurs_runtime.Value
var once_Data_Semigroup_append__1442818457 sync.Once
func Get_Data_Semigroup_append__1442818457() gopurs_runtime.Value {
	once_Data_Semigroup_append__1442818457.Do(func() {
		cache_Data_Semigroup_append__1442818457 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__1442818457(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__1442818457
}

var cache_Data_Semigroup_append__755695413 gopurs_runtime.Value
var once_Data_Semigroup_append__755695413 sync.Once
func Get_Data_Semigroup_append__755695413() gopurs_runtime.Value {
	once_Data_Semigroup_append__755695413.Do(func() {
		cache_Data_Semigroup_append__755695413 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__755695413(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__755695413
}

var cache_Data_Semigroup_append__2832914972 gopurs_runtime.Value
var once_Data_Semigroup_append__2832914972 sync.Once
func Get_Data_Semigroup_append__2832914972() gopurs_runtime.Value {
	once_Data_Semigroup_append__2832914972.Do(func() {
		cache_Data_Semigroup_append__2832914972 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2832914972(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__2832914972
}

var cache_Data_Semigroup_append__204561377 gopurs_runtime.Value
var once_Data_Semigroup_append__204561377 sync.Once
func Get_Data_Semigroup_append__204561377() gopurs_runtime.Value {
	once_Data_Semigroup_append__204561377.Do(func() {
		cache_Data_Semigroup_append__204561377 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__204561377(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__204561377
}

var cache_Data_Semigroup_append__3641242355 gopurs_runtime.Value
var once_Data_Semigroup_append__3641242355 sync.Once
func Get_Data_Semigroup_append__3641242355() gopurs_runtime.Value {
	once_Data_Semigroup_append__3641242355.Do(func() {
		cache_Data_Semigroup_append__3641242355 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__3641242355(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__3641242355
}

var cache_Data_Semigroup_append__3678571768 gopurs_runtime.Value
var once_Data_Semigroup_append__3678571768 sync.Once
func Get_Data_Semigroup_append__3678571768() gopurs_runtime.Value {
	once_Data_Semigroup_append__3678571768.Do(func() {
		cache_Data_Semigroup_append__3678571768 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_Data_Semigroup_append__3678571768(v_0_box.FloatVal(), v1_1_box.FloatVal()))
})
	})
	return cache_Data_Semigroup_append__3678571768
}

var cache_Data_Semigroup_append__493084344 gopurs_runtime.Value
var once_Data_Semigroup_append__493084344 sync.Once
func Get_Data_Semigroup_append__493084344() gopurs_runtime.Value {
	once_Data_Semigroup_append__493084344.Do(func() {
		cache_Data_Semigroup_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semigroup_append__493084344
}

var cache_Data_Semigroup_append__1230318264 gopurs_runtime.Value
var once_Data_Semigroup_append__1230318264 sync.Once
func Get_Data_Semigroup_append__1230318264() gopurs_runtime.Value {
	once_Data_Semigroup_append__1230318264.Do(func() {
		cache_Data_Semigroup_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__1230318264(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__1230318264
}

var cache_Data_Semigroup_append__2285093048 gopurs_runtime.Value
var once_Data_Semigroup_append__2285093048 sync.Once
func Get_Data_Semigroup_append__2285093048() gopurs_runtime.Value {
	once_Data_Semigroup_append__2285093048.Do(func() {
		cache_Data_Semigroup_append__2285093048 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2285093048(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semigroup_append__2285093048
}

var cache_Data_Semigroup_append__365446200 gopurs_runtime.Value
var once_Data_Semigroup_append__365446200 sync.Once
func Get_Data_Semigroup_append__365446200() gopurs_runtime.Value {
	once_Data_Semigroup_append__365446200.Do(func() {
		cache_Data_Semigroup_append__365446200 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__365446200(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semigroup_append__365446200
}

var cache_Data_Semigroup_append__2734706680 gopurs_runtime.Value
var once_Data_Semigroup_append__2734706680 sync.Once
func Get_Data_Semigroup_append__2734706680() gopurs_runtime.Value {
	once_Data_Semigroup_append__2734706680.Do(func() {
		cache_Data_Semigroup_append__2734706680 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2734706680(xs_0_box, ys_1_box)
})
	})
	return cache_Data_Semigroup_append__2734706680
}

var cache_Data_Semigroup_append__2013893496 gopurs_runtime.Value
var once_Data_Semigroup_append__2013893496 sync.Once
func Get_Data_Semigroup_append__2013893496() gopurs_runtime.Value {
	once_Data_Semigroup_append__2013893496.Do(func() {
		cache_Data_Semigroup_append__2013893496 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Semigroup_append__2013893496(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](xs_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](ys_1_box)))}
})
	})
	return cache_Data_Semigroup_append__2013893496
}

var cache_Data_Semigroup_append__868515608 gopurs_runtime.Value
var once_Data_Semigroup_append__868515608 sync.Once
func Get_Data_Semigroup_append__868515608() gopurs_runtime.Value {
	once_Data_Semigroup_append__868515608.Do(func() {
		cache_Data_Semigroup_append__868515608 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_Semigroup_append__868515608(uint32(v_0_box.IntVal), uint32(v1_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_Data_Semigroup_append__868515608
}

var cache_Data_Semigroup_append__2331703800 gopurs_runtime.Value
var once_Data_Semigroup_append__2331703800 sync.Once
func Get_Data_Semigroup_append__2331703800() gopurs_runtime.Value {
	once_Data_Semigroup_append__2331703800.Do(func() {
		cache_Data_Semigroup_append__2331703800 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2331703800(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__2331703800
}

var cache_Data_Semigroup_append__4016378200 gopurs_runtime.Value
var once_Data_Semigroup_append__4016378200 sync.Once
func Get_Data_Semigroup_append__4016378200() gopurs_runtime.Value {
	once_Data_Semigroup_append__4016378200.Do(func() {
		cache_Data_Semigroup_append__4016378200 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__4016378200(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__4016378200
}

var cache_Data_Semigroup_append__988370296 gopurs_runtime.Value
var once_Data_Semigroup_append__988370296 sync.Once
func Get_Data_Semigroup_append__988370296() gopurs_runtime.Value {
	once_Data_Semigroup_append__988370296.Do(func() {
		cache_Data_Semigroup_append__988370296 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__988370296(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](dict_0_box))
})
	})
	return cache_Data_Semigroup_append__988370296
}

var cache_Data_Semigroup_append__2611908184 gopurs_runtime.Value
var once_Data_Semigroup_append__2611908184 sync.Once
func Get_Data_Semigroup_append__2611908184() gopurs_runtime.Value {
	once_Data_Semigroup_append__2611908184.Do(func() {
		cache_Data_Semigroup_append__2611908184 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_append__2611908184(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_Semigroup_append__2611908184
}

var cache_Data_Semigroup_appendRecord__2378130976 gopurs_runtime.Value
var once_Data_Semigroup_appendRecord__2378130976 sync.Once
func Get_Data_Semigroup_appendRecord__2378130976() gopurs_runtime.Value {
	once_Data_Semigroup_appendRecord__2378130976.Do(func() {
		cache_Data_Semigroup_appendRecord__2378130976 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_appendRecord__2378130976(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord](dict_0_box))
})
	})
	return cache_Data_Semigroup_appendRecord__2378130976
}

var cache_Data_Semigroup_appendRecord__1546996774 gopurs_runtime.Value
var once_Data_Semigroup_appendRecord__1546996774 sync.Once
func Get_Data_Semigroup_appendRecord__1546996774() gopurs_runtime.Value {
	once_Data_Semigroup_appendRecord__1546996774.Do(func() {
		cache_Data_Semigroup_appendRecord__1546996774 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Semigroup_appendRecord__1546996774(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_SemigroupRecord](dict_0_box))
})
	})
	return cache_Data_Semigroup_appendRecord__1546996774
}

var cache_Data_Semigroup_semigroupArray__4207347319 gopurs_runtime.Value
var once_Data_Semigroup_semigroupArray__4207347319 sync.Once
func Get_Data_Semigroup_semigroupArray__4207347319() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupArray__4207347319.Do(func() {
		cache_Data_Semigroup_semigroupArray__4207347319 = gopurs_runtime.RecordDict1("append", Get_Data_Semigroup_concatArray())
	})
	return cache_Data_Semigroup_semigroupArray__4207347319
}

var cache_Data_Semigroup_semigroupArray__777842900 gopurs_runtime.Value
var once_Data_Semigroup_semigroupArray__777842900 sync.Once
func Get_Data_Semigroup_semigroupArray__777842900() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupArray__777842900.Do(func() {
		cache_Data_Semigroup_semigroupArray__777842900 = gopurs_runtime.RecordDict1("append", Get_Data_Semigroup_concatArray())
	})
	return cache_Data_Semigroup_semigroupArray__777842900
}

var cache_Data_Semigroup_semigroupArray__1728406699 gopurs_runtime.Value
var once_Data_Semigroup_semigroupArray__1728406699 sync.Once
func Get_Data_Semigroup_semigroupArray__1728406699() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupArray__1728406699.Do(func() {
		cache_Data_Semigroup_semigroupArray__1728406699 = gopurs_runtime.RecordDict1("append", Get_Data_Semigroup_concatArray())
	})
	return cache_Data_Semigroup_semigroupArray__1728406699
}

var cache_Data_Semigroup_semigroupRecordNil__2406047365 gopurs_runtime.Value
var once_Data_Semigroup_semigroupRecordNil__2406047365 sync.Once
func Get_Data_Semigroup_semigroupRecordNil__2406047365() gopurs_runtime.Value {
	once_Data_Semigroup_semigroupRecordNil__2406047365.Do(func() {
		cache_Data_Semigroup_semigroupRecordNil__2406047365 = gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_Data_Semigroup_semigroupRecordNil__2406047365
}

type Constructor_Data_Semigroup_SemigroupRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3847494007] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_SemigroupRecord)(ptr)
		_ = c
		switch key {
		case "appendRecord": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Semigroup_SemigroupRecord: " + key)
		}
	}
}


type Constructor_Data_Semigroup_Semigroup struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2053112122] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Semigroup_Semigroup)(ptr)
		_ = c
		switch key {
		case "append": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Semigroup_Semigroup: " + key)
		}
	}
}


func Call_Data_Semigroup_SemigroupRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_Semigroup_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Semigroup_appendRecord(dict_0_loop *Constructor_Data_Semigroup_SemigroupRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_SemigroupRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_semigroupRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictSemigroupRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictSemigroupRecord_1 gopurs_runtime.Value = dictSemigroupRecord_1_loop
_ = dictSemigroupRecord_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroupRecord_1, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_Semigroup_append(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_semigroupFn(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}))
}

func Call_Data_Semigroup_semigroupRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictSemigroupRecord_2_loop gopurs_runtime.Value, dictSemigroup_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictSemigroupRecord_2 gopurs_runtime.Value = dictSemigroupRecord_2_loop
_ = dictSemigroupRecord_2
var dictSemigroup_3 gopurs_runtime.Value = dictSemigroup_3_loop
_ = dictSemigroup_3
return gopurs_runtime.RecordDict1("appendRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_7_0 -> gopurs_runtime.Value
key_7_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_7_0
// TAST (Let): get_8_1 -> gopurs_runtime.Value
get_8_1 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_7_0.StrVal()))
_ = get_8_1
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_7_0.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_3, "append"), gopurs_runtime.Apply(get_8_1, ra_5), gopurs_runtime.Apply(get_8_1, rb_6)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictSemigroupRecord_2, "appendRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6))
})
})
}))
}

func Call_Data_Semigroup_append__2637663146(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__1124926121(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__4093645121(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__2462288412(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__1442818457(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__755695413(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__2832914972(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__204561377(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__3641242355(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__3678571768(v_0_loop float64, v1_1_loop float64) float64 {
var v_0 float64 = v_0_loop
_ = v_0
var v1_1 float64 = v1_1_loop
_ = v1_1
return gopurs_runtime.Apply3(Get_Data_Semiring_add(), gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring](Get_Data_Semiring_semiringNumber()))}, gopurs_runtime.Float(v_0), gopurs_runtime.Float(v1_1)).FloatVal()
}

func Call_Data_Semigroup_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_Data_Semigroup_append__1230318264(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__2285093048(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), __eta0_0, __eta1_1)
}

func Call_Data_Semigroup_append__365446200(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), __eta0_0, __eta1_1)
}

func Call_Data_Semigroup_append__2734706680(xs_0_loop gopurs_runtime.Value, ys_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var xs_0 gopurs_runtime.Value = xs_0_loop
_ = xs_0
var ys_1 gopurs_runtime.Value = ys_1_loop
_ = ys_1
return gopurs_runtime.Apply3(Get_Data_Functor_go__map(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Lazy_functorLazy()))}, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), ys_1)))}
goto end_branch_0
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 218341868 && v_2.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](gopurs_runtime.Apply2(Get_Data_List_Lazy_Types_Cons(), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_semigroupList(), "append"), (*Constructor_Data_List_Lazy_Types_Cons)(v_2.UnsafePtr).V1, ys_1))))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 218341868, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Lazy_Types_Cons](__t0))}
}), gopurs_runtime.Apply2(Get_Data_Newtype_unwrap(), gopurs_runtime.Value{}, xs_0))
}

func Call_Data_Semigroup_append__2013893496(xs_0_loop *Constructor_Data_List_Types_Cons, ys_1_loop *Constructor_Data_List_Types_Cons) *Constructor_Data_List_Types_Cons {
var xs_0 *Constructor_Data_List_Types_Cons = xs_0_loop
_ = xs_0
var ys_1 *Constructor_Data_List_Types_Cons = ys_1_loop
_ = ys_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply4(Get_Data_Foldable_foldr(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_List_Types_foldableList()))}, Get_Data_List_Types_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_Data_Semigroup_append__868515608(v_0_loop uint32, v1_1_loop uint32) uint32 {
var v_0 uint32 = v_0_loop
_ = v_0
var v1_1 uint32 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0 == 1527465420) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Ordering_LT().IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 == 380165415) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Get_Data_Ordering_GT().IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (v_0 == 902936544) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(v1_1), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return uint32(__t0.IntVal)
}

func Call_Data_Semigroup_append__2331703800(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__4016378200(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__988370296(dict_0_loop *Constructor_Data_Semigroup_Semigroup) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_Semigroup = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_append__2611908184(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Comparison_semigroupFn()).V0), __eta0_0, __eta1_1)
}

func Call_Data_Semigroup_appendRecord__2378130976(dict_0_loop *Constructor_Data_Semigroup_SemigroupRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_SemigroupRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Semigroup_appendRecord__1546996774(dict_0_loop *Constructor_Data_Semigroup_SemigroupRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Semigroup_SemigroupRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Get_Data_Semigroup_concatArray() gopurs_runtime.Value {
	return _Gopurs_Data_Semigroup_ConcatArray
}

func Get_Data_Semigroup_concatString() gopurs_runtime.Value {
	return _Gopurs_Data_Semigroup_ConcatString
}

package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Eq_EqRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Eq_EqRecord_dollar_Dict sync.Once
func Get_Data_Eq_EqRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Eq_EqRecord_dollar_Dict.Do(func() {
		cache_Data_Eq_EqRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer(Call_Data_Eq_EqRecord_dollar_Dict(func() struct{
	eqRecord gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	eqRecord gopurs_runtime.Value
}{}
					clone.eqRecord = gopurs_runtime.RecordGet(orig, "eqRecord")
					return clone
				}()))}
})
	})
	return cache_Data_Eq_EqRecord_dollar_Dict
}

var cache_Data_Eq_Eq_dollar_Dict gopurs_runtime.Value
var once_Data_Eq_Eq_dollar_Dict sync.Once
func Get_Data_Eq_Eq_dollar_Dict() gopurs_runtime.Value {
	once_Data_Eq_Eq_dollar_Dict.Do(func() {
		cache_Data_Eq_Eq_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Call_Data_Eq_Eq_dollar_Dict(func() struct{
	eq gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	eq gopurs_runtime.Value
}{}
					clone.eq = gopurs_runtime.RecordGet(orig, "eq")
					return clone
				}()))}
})
	})
	return cache_Data_Eq_Eq_dollar_Dict
}

var cache_Data_Eq_Eq1_dollar_Dict gopurs_runtime.Value
var once_Data_Eq_Eq1_dollar_Dict sync.Once
func Get_Data_Eq_Eq1_dollar_Dict() gopurs_runtime.Value {
	once_Data_Eq_Eq1_dollar_Dict.Do(func() {
		cache_Data_Eq_Eq1_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(Call_Data_Eq_Eq1_dollar_Dict(func() struct{
	eq1 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	eq1 gopurs_runtime.Value
}{}
					clone.eq1 = gopurs_runtime.RecordGet(orig, "eq1")
					return clone
				}()))}
})
	})
	return cache_Data_Eq_Eq1_dollar_Dict
}

var cache_Data_Eq_eqVoid gopurs_runtime.Value
var once_Data_Eq_eqVoid sync.Once
func Get_Data_Eq_eqVoid() gopurs_runtime.Value {
	once_Data_Eq_eqVoid.Do(func() {
		cache_Data_Eq_eqVoid = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})}))}
	})
	return cache_Data_Eq_eqVoid
}

var cache_Data_Eq_eqUnit gopurs_runtime.Value
var once_Data_Eq_eqUnit sync.Once
func Get_Data_Eq_eqUnit() gopurs_runtime.Value {
	once_Data_Eq_eqUnit.Do(func() {
		cache_Data_Eq_eqUnit = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})}))}
	})
	return cache_Data_Eq_eqUnit
}

var cache_Data_Eq_eqString gopurs_runtime.Value
var once_Data_Eq_eqString sync.Once
func Get_Data_Eq_eqString() gopurs_runtime.Value {
	once_Data_Eq_eqString.Do(func() {
		cache_Data_Eq_eqString = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Eq_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, Get_Data_Eq_eqStringImpl()})))}
	})
	return cache_Data_Eq_eqString
}

var cache_Data_Eq_eqRowNil gopurs_runtime.Value
var once_Data_Eq_eqRowNil sync.Once
func Get_Data_Eq_eqRowNil() gopurs_runtime.Value {
	once_Data_Eq_eqRowNil.Do(func() {
		cache_Data_Eq_eqRowNil = gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})}))}
	})
	return cache_Data_Eq_eqRowNil
}

var cache_Data_Eq_eqRecord gopurs_runtime.Value
var once_Data_Eq_eqRecord sync.Once
func Get_Data_Eq_eqRecord() gopurs_runtime.Value {
	once_Data_Eq_eqRecord.Do(func() {
		cache_Data_Eq_eqRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Eq_eqRecord
}

var cache_Data_Eq_eqRec gopurs_runtime.Value
var once_Data_Eq_eqRec sync.Once
func Get_Data_Eq_eqRec() gopurs_runtime.Value {
	once_Data_Eq_eqRec.Do(func() {
		cache_Data_Eq_eqRec = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictEqRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRec(_dollar___unused_0_box, dictEqRecord_1_box)
})
	})
	return cache_Data_Eq_eqRec
}

var cache_Data_Eq_eqProxy gopurs_runtime.Value
var once_Data_Eq_eqProxy sync.Once
func Get_Data_Eq_eqProxy() gopurs_runtime.Value {
	once_Data_Eq_eqProxy.Do(func() {
		cache_Data_Eq_eqProxy = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Eq_3768443459_3790796878((&Constructor_Data_Eq_Eq[uint32]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})})))}
	})
	return cache_Data_Eq_eqProxy
}

var cache_Data_Eq_eqNumber gopurs_runtime.Value
var once_Data_Eq_eqNumber sync.Once
func Get_Data_Eq_eqNumber() gopurs_runtime.Value {
	once_Data_Eq_eqNumber.Do(func() {
		cache_Data_Eq_eqNumber = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Eq_687527510_3790796878((&Constructor_Data_Eq_Eq[float64]{1, Get_Data_Eq_eqNumberImpl()})))}
	})
	return cache_Data_Eq_eqNumber
}

var cache_Data_Eq_eqInt gopurs_runtime.Value
var once_Data_Eq_eqInt sync.Once
func Get_Data_Eq_eqInt() gopurs_runtime.Value {
	once_Data_Eq_eqInt.Do(func() {
		cache_Data_Eq_eqInt = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Eq_1053099733_3790796878((&Constructor_Data_Eq_Eq[int64]{1, Get_Data_Eq_eqIntImpl()})))}
	})
	return cache_Data_Eq_eqInt
}

var cache_Data_Eq_eqChar gopurs_runtime.Value
var once_Data_Eq_eqChar sync.Once
func Get_Data_Eq_eqChar() gopurs_runtime.Value {
	once_Data_Eq_eqChar.Do(func() {
		cache_Data_Eq_eqChar = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Eq_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, Get_Data_Eq_eqCharImpl()})))}
	})
	return cache_Data_Eq_eqChar
}

var cache_Data_Eq_eqBoolean gopurs_runtime.Value
var once_Data_Eq_eqBoolean sync.Once
func Get_Data_Eq_eqBoolean() gopurs_runtime.Value {
	once_Data_Eq_eqBoolean.Do(func() {
		cache_Data_Eq_eqBoolean = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Eq_2737952170_3790796878((&Constructor_Data_Eq_Eq[bool]{1, Get_Data_Eq_eqBooleanImpl()})))}
	})
	return cache_Data_Eq_eqBoolean
}

var cache_Data_Eq_eq1 gopurs_runtime.Value
var once_Data_Eq_eq1 sync.Once
func Get_Data_Eq_eq1() gopurs_runtime.Value {
	once_Data_Eq_eq1.Do(func() {
		cache_Data_Eq_eq1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Eq_eq1
}

var cache_Data_Eq_eq gopurs_runtime.Value
var once_Data_Eq_eq sync.Once
func Get_Data_Eq_eq() gopurs_runtime.Value {
	once_Data_Eq_eq.Do(func() {
		cache_Data_Eq_eq = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Eq_eq
}

var cache_Data_Eq_eq__720879455 gopurs_runtime.Value
var once_Data_Eq_eq__720879455 sync.Once
func Get_Data_Eq_eq__720879455() gopurs_runtime.Value {
	once_Data_Eq_eq__720879455.Do(func() {
		cache_Data_Eq_eq__720879455 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__720879455(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__720879455
}

var cache_Data_Eq_eq__2791951487 gopurs_runtime.Value
var once_Data_Eq_eq__2791951487 sync.Once
func Get_Data_Eq_eq__2791951487() gopurs_runtime.Value {
	once_Data_Eq_eq__2791951487.Do(func() {
		cache_Data_Eq_eq__2791951487 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2791951487(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__2791951487
}

var cache_Data_Eq_eq__3889410207 gopurs_runtime.Value
var once_Data_Eq_eq__3889410207 sync.Once
func Get_Data_Eq_eq__3889410207() gopurs_runtime.Value {
	once_Data_Eq_eq__3889410207.Do(func() {
		cache_Data_Eq_eq__3889410207 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__3889410207(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__3889410207
}

var cache_Data_Eq_eq__1136572319 gopurs_runtime.Value
var once_Data_Eq_eq__1136572319 sync.Once
func Get_Data_Eq_eq__1136572319() gopurs_runtime.Value {
	once_Data_Eq_eq__1136572319.Do(func() {
		cache_Data_Eq_eq__1136572319 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__1136572319(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__1136572319
}

var cache_Data_Eq_eq__1624178527 gopurs_runtime.Value
var once_Data_Eq_eq__1624178527 sync.Once
func Get_Data_Eq_eq__1624178527() gopurs_runtime.Value {
	once_Data_Eq_eq__1624178527.Do(func() {
		cache_Data_Eq_eq__1624178527 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__1624178527(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__1624178527
}

var cache_Data_Eq_eq__777679551 gopurs_runtime.Value
var once_Data_Eq_eq__777679551 sync.Once
func Get_Data_Eq_eq__777679551() gopurs_runtime.Value {
	once_Data_Eq_eq__777679551.Do(func() {
		cache_Data_Eq_eq__777679551 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__777679551(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__777679551
}

var cache_Data_Eq_eq__666462110 gopurs_runtime.Value
var once_Data_Eq_eq__666462110 sync.Once
func Get_Data_Eq_eq__666462110() gopurs_runtime.Value {
	once_Data_Eq_eq__666462110.Do(func() {
		cache_Data_Eq_eq__666462110 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__666462110(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Eq_eq__666462110
}

var cache_Data_Eq_eq__2250289689 gopurs_runtime.Value
var once_Data_Eq_eq__2250289689 sync.Once
func Get_Data_Eq_eq__2250289689() gopurs_runtime.Value {
	once_Data_Eq_eq__2250289689.Do(func() {
		cache_Data_Eq_eq__2250289689 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2250289689(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Eq_eq__2250289689
}

var cache_Data_Eq_eq__2235094066 gopurs_runtime.Value
var once_Data_Eq_eq__2235094066 sync.Once
func Get_Data_Eq_eq__2235094066() gopurs_runtime.Value {
	once_Data_Eq_eq__2235094066.Do(func() {
		cache_Data_Eq_eq__2235094066 = gopurs_runtime.Func2(func(__eta_norm_0_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__2235094066(__eta_norm_0_unused_0_box, __eta_norm_1_1_box)
})
	})
	return cache_Data_Eq_eq__2235094066
}

var cache_Data_Eq_eq__4087304255 gopurs_runtime.Value
var once_Data_Eq_eq__4087304255 sync.Once
func Get_Data_Eq_eq__4087304255() gopurs_runtime.Value {
	once_Data_Eq_eq__4087304255.Do(func() {
		cache_Data_Eq_eq__4087304255 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__4087304255(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__4087304255
}

var cache_Data_Eq_eq__3216698335 gopurs_runtime.Value
var once_Data_Eq_eq__3216698335 sync.Once
func Get_Data_Eq_eq__3216698335() gopurs_runtime.Value {
	once_Data_Eq_eq__3216698335.Do(func() {
		cache_Data_Eq_eq__3216698335 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__3216698335(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__3216698335
}

var cache_Data_Eq_eq__801415071 gopurs_runtime.Value
var once_Data_Eq_eq__801415071 sync.Once
func Get_Data_Eq_eq__801415071() gopurs_runtime.Value {
	once_Data_Eq_eq__801415071.Do(func() {
		cache_Data_Eq_eq__801415071 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__801415071(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__801415071
}

var cache_Data_Eq_eq__1265619615 gopurs_runtime.Value
var once_Data_Eq_eq__1265619615 sync.Once
func Get_Data_Eq_eq__1265619615() gopurs_runtime.Value {
	once_Data_Eq_eq__1265619615.Do(func() {
		cache_Data_Eq_eq__1265619615 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__1265619615(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__1265619615
}

var cache_Data_Eq_eq__954999263 gopurs_runtime.Value
var once_Data_Eq_eq__954999263 sync.Once
func Get_Data_Eq_eq__954999263() gopurs_runtime.Value {
	once_Data_Eq_eq__954999263.Do(func() {
		cache_Data_Eq_eq__954999263 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__954999263(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__954999263
}

var cache_Data_Eq_eq__1530357663 gopurs_runtime.Value
var once_Data_Eq_eq__1530357663 sync.Once
func Get_Data_Eq_eq__1530357663() gopurs_runtime.Value {
	once_Data_Eq_eq__1530357663.Do(func() {
		cache_Data_Eq_eq__1530357663 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq__1530357663(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Eq_eq__1530357663
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
		cache_Data_Eq_eq1Array = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
})}))}
	})
	return cache_Data_Eq_eq1Array
}

var cache_Data_Eq_eqRowCons gopurs_runtime.Value
var once_Data_Eq_eqRowCons sync.Once
func Get_Data_Eq_eqRowCons() gopurs_runtime.Value {
	once_Data_Eq_eqRowCons.Do(func() {
		cache_Data_Eq_eqRowCons = gopurs_runtime.Func4(func(dictEqRecord_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictIsSymbol_2_box gopurs_runtime.Value, dictEq_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eqRowCons(dictEqRecord_0_box, _dollar___unused_1_box, dictIsSymbol_2_box, dictEq_3_box)
})
	})
	return cache_Data_Eq_eqRowCons
}

var cache_Data_Eq_notEq gopurs_runtime.Value
var once_Data_Eq_notEq sync.Once
func Get_Data_Eq_notEq() gopurs_runtime.Value {
	once_Data_Eq_notEq.Do(func() {
		cache_Data_Eq_notEq = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_Data_Eq_notEq
}

var cache_Data_Eq_notEq__666462110 gopurs_runtime.Value
var once_Data_Eq_notEq__666462110 sync.Once
func Get_Data_Eq_notEq__666462110() gopurs_runtime.Value {
	once_Data_Eq_notEq__666462110.Do(func() {
		cache_Data_Eq_notEq__666462110 = gopurs_runtime.Func2(func(y_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__666462110(uint32(y_unused_0_box.IntVal), x_1_box))
})
	})
	return cache_Data_Eq_notEq__666462110
}

var cache_Data_Eq_notEq__2250289689 gopurs_runtime.Value
var once_Data_Eq_notEq__2250289689 sync.Once
func Get_Data_Eq_notEq__2250289689() gopurs_runtime.Value {
	once_Data_Eq_notEq__2250289689.Do(func() {
		cache_Data_Eq_notEq__2250289689 = gopurs_runtime.Func2(func(y_unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__2250289689(uint32(y_unused_0_box.IntVal), x_1_box))
})
	})
	return cache_Data_Eq_notEq__2250289689
}

var cache_Data_Eq_notEq__1265619615 gopurs_runtime.Value
var once_Data_Eq_notEq__1265619615 sync.Once
func Get_Data_Eq_notEq__1265619615() gopurs_runtime.Value {
	once_Data_Eq_notEq__1265619615.Do(func() {
		cache_Data_Eq_notEq__1265619615 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__1265619615(x_0_box.IntVal, y_1_box.IntVal))
})
	})
	return cache_Data_Eq_notEq__1265619615
}

var cache_Data_Eq_notEq__954999263 gopurs_runtime.Value
var once_Data_Eq_notEq__954999263 sync.Once
func Get_Data_Eq_notEq__954999263() gopurs_runtime.Value {
	once_Data_Eq_notEq__954999263.Do(func() {
		cache_Data_Eq_notEq__954999263 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq__954999263(x_0_box.FloatVal(), y_1_box.FloatVal()))
})
	})
	return cache_Data_Eq_notEq__954999263
}

var cache_Data_Eq_notEq1 gopurs_runtime.Value
var once_Data_Eq_notEq1 sync.Once
func Get_Data_Eq_notEq1() gopurs_runtime.Value {
	once_Data_Eq_notEq1.Do(func() {
		cache_Data_Eq_notEq1 = gopurs_runtime.Func4(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value, y_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Eq_notEq1(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](dictEq1_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_1_box), x_2_box, y_3_box))
})
	})
	return cache_Data_Eq_notEq1
}

type Constructor_Data_Eq_EqRecord[T_rowlist any, T_row any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1311326743] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Eq_EqRecord[any, any])(ptr)
		_ = c
		switch key {
		case "eqRecord": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Eq_EqRecord: " + key)
		}
	}
}


type Constructor_Data_Eq_Eq[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1012063514] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Eq_Eq[any])(ptr)
		_ = c
		switch key {
		case "eq": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Eq_Eq: " + key)
		}
	}
}


type Constructor_Data_Eq_Eq1[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1715248107] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Eq_Eq1[any])(ptr)
		_ = c
		switch key {
		case "eq1": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Eq_Eq1: " + key)
		}
	}
}


func Call_Data_Eq_EqRecord_dollar_Dict(x_0_loop struct{
	eqRecord gopurs_runtime.Value
}) *Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	eqRecord gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("eqRecord", orig.eqRecord)
				}())
}

func Call_Data_Eq_Eq_dollar_Dict(x_0_loop struct{
	eq gopurs_runtime.Value
}) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
var x_0 struct{
	eq gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("eq", orig.eq)
				}())
}

func Call_Data_Eq_Eq1_dollar_Dict(x_0_loop struct{
	eq1 gopurs_runtime.Value
}) *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] {
var x_0 struct{
	eq1 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("eq1", orig.eq1)
				}())
}

func Call_Data_Eq_eqRecord(dict_0_loop *Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eqRec(_dollar___unused_0_loop gopurs_runtime.Value, dictEqRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictEqRecord_1 gopurs_runtime.Value = dictEqRecord_1_loop
_ = dictEqRecord_1
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEqRecord_1, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Eq_eq1(dict_0_loop *Constructor_Data_Eq_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq(dict_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Eq_eq__720879455(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__720879455:
for {
if false { continue eq__720879455 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t25 bool
{
var __t_tag_3 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_3) == 1908470532) {
var __t_tag_4 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_4) == 1908470532)
goto end_branch_25
} else {

}
}
{
var __t_tag_5 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_5) == 2455627378) {
var __t_tag_6 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_6) == 2455627378)
goto end_branch_25
} else {

}
}
{
var __t_tag_7 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_7) == 4162469099) {
var __t_tag_8 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_8) == 4162469099)
goto end_branch_25
} else {

}
}
{
var __t_tag_9 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_9) == 1692989816) {
var __t_tag_10 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_10) == 1692989816)
goto end_branch_25
} else {

}
}
{
var __t_tag_11 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_11) == 330658827) {
var __t_tag_12 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_12) == 330658827)
goto end_branch_25
} else {

}
}
{
var __t_tag_13 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_13) == 4067355978) {
var __t_tag_14 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_14) == 4067355978)
goto end_branch_25
} else {

}
}
{
var __t_tag_15 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_15) == 2276710548) {
var __t_tag_16 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_16) == 2276710548)
goto end_branch_25
} else {

}
}
{
var __t_tag_17 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_17) == 243771071) {
var __t_tag_18 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_18) == 243771071)
goto end_branch_25
} else {

}
}
{
var __t_tag_19 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_19) == 215731793) {
var __t_tag_20 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_20) == 215731793)
goto end_branch_25
} else {

}
}
{
var __t_tag_21 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_21) == 8639228) {
var __t_tag_22 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_22) == 8639228)
goto end_branch_25
} else {

}
}
{
var __t_tag_23 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_23) == 49471444) {
var __t_tag_24 uint32 = uint32(__eta_norm_0_1.IntVal)
__t25 = (uint32(__t_tag_24) == 49471444)
goto end_branch_25
} else {

}
}
{
var __t_tag_0 uint32 = uint32(__eta_norm_1_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 3889233761) {

var __t_tag_1 uint32 = uint32(__eta_norm_0_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 3889233761)
}
__t25 = __t_and_2
}
end_branch_25:
return gopurs_runtime.Bool(__t25)
}
}

func Call_Data_Eq_eq__2791951487(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__2791951487:
for {
if false { continue eq__2791951487 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_Date_Date]](Get_Data_Date_eqDate()).V0), __eta_norm_1_0, __eta_norm_0_1).IntVal) != (0))
}
}

func Call_Data_Eq_eq__3889410207(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__3889410207:
for {
if false { continue eq__3889410207 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t5 bool
{
var __t_tag_3 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_3) == 2591059121) {
var __t_tag_4 uint32 = uint32(__eta_norm_0_1.IntVal)
__t5 = (uint32(__t_tag_4) == 2591059121)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 uint32 = uint32(__eta_norm_1_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 658452902) {

var __t_tag_1 uint32 = uint32(__eta_norm_0_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 658452902)
}
__t5 = __t_and_2
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
}
}

func Call_Data_Eq_eq__1136572319(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__1136572319:
for {
if false { continue eq__1136572319 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t15 bool
{
var __t_tag_3 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_3) == 3908053364) {
var __t_tag_4 uint32 = uint32(__eta_norm_0_1.IntVal)
__t15 = (uint32(__t_tag_4) == 3908053364)
goto end_branch_15
} else {

}
}
{
var __t_tag_5 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_5) == 217821258) {
var __t_tag_6 uint32 = uint32(__eta_norm_0_1.IntVal)
__t15 = (uint32(__t_tag_6) == 217821258)
goto end_branch_15
} else {

}
}
{
var __t_tag_7 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_7) == 1292308612) {
var __t_tag_8 uint32 = uint32(__eta_norm_0_1.IntVal)
__t15 = (uint32(__t_tag_8) == 1292308612)
goto end_branch_15
} else {

}
}
{
var __t_tag_9 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_9) == 2311060696) {
var __t_tag_10 uint32 = uint32(__eta_norm_0_1.IntVal)
__t15 = (uint32(__t_tag_10) == 2311060696)
goto end_branch_15
} else {

}
}
{
var __t_tag_11 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_11) == 401302776) {
var __t_tag_12 uint32 = uint32(__eta_norm_0_1.IntVal)
__t15 = (uint32(__t_tag_12) == 401302776)
goto end_branch_15
} else {

}
}
{
var __t_tag_13 uint32 = uint32(__eta_norm_1_0.IntVal)
if (uint32(__t_tag_13) == 3327533908) {
var __t_tag_14 uint32 = uint32(__eta_norm_0_1.IntVal)
__t15 = (uint32(__t_tag_14) == 3327533908)
goto end_branch_15
} else {

}
}
{
var __t_tag_0 uint32 = uint32(__eta_norm_1_0.IntVal)
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 3631736139) {

var __t_tag_1 uint32 = uint32(__eta_norm_0_1.IntVal)
__t_and_2 = (uint32(__t_tag_1) == 3631736139)
}
__t15 = __t_and_2
}
end_branch_15:
return gopurs_runtime.Bool(__t15)
}
}

func Call_Data_Eq_eq__1624178527(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__1624178527:
for {
if false { continue eq__1624178527 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_eqMap()).V0), __eta_norm_1_0, __eta_norm_0_1).IntVal) != (0))
}
}

func Call_Data_Eq_eq__777679551(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__777679551:
for {
if false { continue eq__777679551 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
var __t5 bool
{
var __t_tag_3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
if (__t_tag_3 == nil) {
var __t_tag_4 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1)
__t5 = (__t_tag_4 == nil)
goto end_branch_5
} else {

}
}
{
var __t_tag_0 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_1_0)
var __t_and_2 bool = false
if (__t_tag_0 != nil) {

var __t_tag_1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__eta_norm_0_1)
__t_and_2 = ((__t_tag_1 != nil)) && (((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_1_0.UnsafePtr).V0.IntVal) == ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(__eta_norm_0_1.UnsafePtr).V0.IntVal))
}
__t5 = __t_and_2
}
end_branch_5:
return gopurs_runtime.Bool(__t5)
}
}

func Call_Data_Eq_eq__666462110(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__666462110:
for {
if false { continue eq__666462110 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
var __t_tag_0 uint32 = uint32(__eta_norm_1_1.IntVal)
return gopurs_runtime.Bool((uint32(__t_tag_0) == 902936544))
}
}

func Call_Data_Eq_eq__2250289689(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__2250289689:
for {
if false { continue eq__2250289689 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
var __t_tag_0 uint32 = uint32(__eta_norm_1_1.IntVal)
return gopurs_runtime.Bool((uint32(__t_tag_0) == 380165415))
}
}

func Call_Data_Eq_eq__2235094066(__eta_norm_0_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__2235094066:
for {
if false { continue eq__2235094066 }
var __eta_norm_0_unused_0 gopurs_runtime.Value = __eta_norm_0_unused_0_loop
_ = __eta_norm_0_unused_0
var __eta_norm_1_1 gopurs_runtime.Value = __eta_norm_1_1_loop
_ = __eta_norm_1_1
var __t_tag_0 uint32 = uint32(__eta_norm_1_1.IntVal)
return gopurs_runtime.Bool((uint32(__t_tag_0) == 1527465420))
}
}

func Call_Data_Eq_eq__4087304255(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__4087304255:
for {
if false { continue eq__4087304255 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool((((((*Constructor_Data_Time_Time)(__eta_norm_1_0.UnsafePtr).V0) == ((*Constructor_Data_Time_Time)(__eta_norm_0_1.UnsafePtr).V0)) && (((*Constructor_Data_Time_Time)(__eta_norm_1_0.UnsafePtr).V1) == ((*Constructor_Data_Time_Time)(__eta_norm_0_1.UnsafePtr).V1))) && (((*Constructor_Data_Time_Time)(__eta_norm_1_0.UnsafePtr).V2) == ((*Constructor_Data_Time_Time)(__eta_norm_0_1.UnsafePtr).V2))) && (((*Constructor_Data_Time_Time)(__eta_norm_1_0.UnsafePtr).V3) == ((*Constructor_Data_Time_Time)(__eta_norm_0_1.UnsafePtr).V3)))
}
}

func Call_Data_Eq_eq__3216698335(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__3216698335:
for {
if false { continue eq__3216698335 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]string]](Get_Data_String_Regex_Flags_eqArray()).V0), __eta_norm_1_0, __eta_norm_0_1).IntVal) != (0))
}
}

func Call_Data_Eq_eq__801415071(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__801415071:
for {
if false { continue eq__801415071 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool(((__eta_norm_1_0.IntVal) != (0)) == ((__eta_norm_0_1.IntVal) != (0)))
}
}

func Call_Data_Eq_eq__1265619615(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__1265619615:
for {
if false { continue eq__1265619615 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool((__eta_norm_1_0.IntVal) == (__eta_norm_0_1.IntVal))
}
}

func Call_Data_Eq_eq__954999263(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__954999263:
for {
if false { continue eq__954999263 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool((__eta_norm_1_0.FloatVal()) == (__eta_norm_0_1.FloatVal()))
}
}

func Call_Data_Eq_eq__1530357663(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
eq__1530357663:
for {
if false { continue eq__1530357663 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Bool((__eta_norm_1_0.StrVal()) == (__eta_norm_0_1.StrVal()))
}
}

func Call_Data_Eq_eqArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Eq_1939691112_3790796878((&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))})))}
}

func Call_Data_Eq_eqRowCons(dictEqRecord_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictIsSymbol_2_loop gopurs_runtime.Value, dictEq_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEqRecord_0 gopurs_runtime.Value = dictEqRecord_0_loop
_ = dictEqRecord_0
var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
_ = _dollar___unused_1
var dictIsSymbol_2 gopurs_runtime.Value = dictIsSymbol_2_loop
_ = dictIsSymbol_2
var dictEq_3 gopurs_runtime.Value = dictEq_3_loop
_ = dictEq_3
return gopurs_runtime.Value{Type: 9, IntVal: 1311326743, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_EqRecord[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, ra_5 gopurs_runtime.Value, rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get__3905987523_7_0 shape=App(Var) bindingType=(Func [(Record (Row [] (TypeVar row)))] (TypeVar focus))
get__3905987523_7_0 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_2, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()))
_ = get__3905987523_7_0
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_3, "eq"), gopurs_runtime.Apply(get__3905987523_7_0, ra_5), gopurs_runtime.Apply(get__3905987523_7_0, rb_6)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEqRecord_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
})}))}
}

func Call_Data_Eq_notEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return ((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), x_1, y_2).IntVal) != (0)) != (true)
}

func Call_Data_Eq_notEq__666462110(y_unused_0_loop uint32, x_1_loop gopurs_runtime.Value) bool {
notEq__666462110:
for {
if false { continue notEq__666462110 }
var y_unused_0 uint32 = y_unused_0_loop
_ = y_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return false
}
}

func Call_Data_Eq_notEq__2250289689(y_unused_0_loop uint32, x_1_loop gopurs_runtime.Value) bool {
notEq__2250289689:
for {
if false { continue notEq__2250289689 }
var y_unused_0 uint32 = y_unused_0_loop
_ = y_unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return false
}
}

func Call_Data_Eq_notEq__1265619615(x_0_loop int64, y_1_loop int64) bool {
notEq__1265619615:
for {
if false { continue notEq__1265619615 }
var x_0 int64 = x_0_loop
_ = x_0
var y_1 int64 = y_1_loop
_ = y_1
return (gopurs_runtime.Bool((x_0) == (y_1)).IntVal) == (gopurs_runtime.Bool(false).IntVal)
}
}

func Call_Data_Eq_notEq__954999263(x_0_loop float64, y_1_loop float64) bool {
notEq__954999263:
for {
if false { continue notEq__954999263 }
var x_0 float64 = x_0_loop
_ = x_0
var y_1 float64 = y_1_loop
_ = y_1
return (gopurs_runtime.Bool((x_0) == (y_1)).FloatVal()) == (gopurs_runtime.Bool(false).FloatVal())
}
}

func Call_Data_Eq_notEq1(dictEq1_0_loop *Constructor_Data_Eq_Eq1[gopurs_runtime.Value], dictEq_1_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_2_loop gopurs_runtime.Value, y_3_loop gopurs_runtime.Value) bool {
var dictEq1_0 *Constructor_Data_Eq_Eq1[gopurs_runtime.Value] = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_1_loop
_ = dictEq_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
var y_3 gopurs_runtime.Value = y_3_loop
_ = y_3
return ((gopurs_runtime.Apply3(gopurs_runtime.Box(dictEq1_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_1)}, x_2, y_3).IntVal) != (0)) != (true)
}

func Rebox_Data_Eq_1053099733_3790796878(in *Constructor_Data_Eq_Eq[int64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Eq_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Eq_1939691112_3790796878(in *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Eq_2737952170_3790796878(in *Constructor_Data_Eq_Eq[bool]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Eq_3768443459_3790796878(in *Constructor_Data_Eq_Eq[uint32]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Eq_687527510_3790796878(in *Constructor_Data_Eq_Eq[float64]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
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

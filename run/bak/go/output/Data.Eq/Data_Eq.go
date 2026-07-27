package Data_Eq

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var cache_eqVoid gopurs_runtime.Value
var once_eqVoid sync.Once
func Get_eqVoid() gopurs_runtime.Value {
	once_eqVoid.Do(func() {
		cache_eqVoid = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))))
	})
	return cache_eqVoid
}

var cache_eqUnit gopurs_runtime.Value
var once_eqUnit sync.Once
func Get_eqUnit() gopurs_runtime.Value {
	once_eqUnit.Do(func() {
		cache_eqUnit = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))))
	})
	return cache_eqUnit
}

var cache_eqString gopurs_runtime.Value
var once_eqString sync.Once
func Get_eqString() gopurs_runtime.Value {
	once_eqString.Do(func() {
		cache_eqString = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", Get_eqStringImpl())))
	})
	return cache_eqString
}

var cache_eqRowNil gopurs_runtime.Value
var once_eqRowNil sync.Once
func Get_eqRowNil() gopurs_runtime.Value {
	once_eqRowNil.Do(func() {
		cache_eqRowNil = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))))
	})
	return cache_eqRowNil
}

var cache_eqRecord gopurs_runtime.Value
var once_eqRecord sync.Once
func Get_eqRecord() gopurs_runtime.Value {
	once_eqRecord.Do(func() {
		cache_eqRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqRecord(dict_0_box)
})
	})
	return cache_eqRecord
}

var cache_eqRec gopurs_runtime.Value
var once_eqRec sync.Once
func Get_eqRec() gopurs_runtime.Value {
	once_eqRec.Do(func() {
		cache_eqRec = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictEqRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_eqRec(_dollar__unused_0_box, dictEqRecord_1_box))
})
	})
	return cache_eqRec
}

var cache_eqProxy gopurs_runtime.Value
var once_eqProxy sync.Once
func Get_eqProxy() gopurs_runtime.Value {
	once_eqProxy.Do(func() {
		cache_eqProxy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))))
	})
	return cache_eqProxy
}

var cache_eqNumber gopurs_runtime.Value
var once_eqNumber sync.Once
func Get_eqNumber() gopurs_runtime.Value {
	once_eqNumber.Do(func() {
		cache_eqNumber = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", Get_eqNumberImpl())))
	})
	return cache_eqNumber
}

var cache_eqInt gopurs_runtime.Value
var once_eqInt sync.Once
func Get_eqInt() gopurs_runtime.Value {
	once_eqInt.Do(func() {
		cache_eqInt = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", Get_eqIntImpl())))
	})
	return cache_eqInt
}

var cache_eqChar gopurs_runtime.Value
var once_eqChar sync.Once
func Get_eqChar() gopurs_runtime.Value {
	once_eqChar.Do(func() {
		cache_eqChar = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", Get_eqCharImpl())))
	})
	return cache_eqChar
}

var cache_eqBoolean gopurs_runtime.Value
var once_eqBoolean sync.Once
func Get_eqBoolean() gopurs_runtime.Value {
	once_eqBoolean.Do(func() {
		cache_eqBoolean = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", Get_eqBooleanImpl())))
	})
	return cache_eqBoolean
}

var cache_eq1 gopurs_runtime.Value
var once_eq1 sync.Once
func Get_eq1() gopurs_runtime.Value {
	once_eq1.Do(func() {
		cache_eq1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1(dict_0_box)
})
	})
	return cache_eq1
}

var cache_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882 gopurs_runtime.Value
var once_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882 sync.Once
func Get_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882() gopurs_runtime.Value {
	once_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882.Do(func() {
		cache_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882(dict_0_box)
})
	})
	return cache_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882
}

var cache_eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		cache_eq = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq(dict_0_box)
})
	})
	return cache_eq
}

var cache_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619 gopurs_runtime.Value
var once_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619 sync.Once
func Get_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619() gopurs_runtime.Value {
	once_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619.Do(func() {
		cache_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619(dict_0_box)
})
	})
	return cache_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619
}

var cache_eqArray gopurs_runtime.Value
var once_eqArray sync.Once
func Get_eqArray() gopurs_runtime.Value {
	once_eqArray.Do(func() {
		cache_eqArray = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_eqArray(dictEq_0_box))
})
	})
	return cache_eqArray
}

var cache_eq1Array gopurs_runtime.Value
var once_eq1Array sync.Once
func Get_eq1Array() gopurs_runtime.Value {
	once_eq1Array.Do(func() {
		cache_eq1Array = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}))))
	})
	return cache_eq1Array
}

var cache_eqRowCons gopurs_runtime.Value
var once_eqRowCons sync.Once
func Get_eqRowCons() gopurs_runtime.Value {
	once_eqRowCons.Do(func() {
		cache_eqRowCons = gopurs_runtime.Func4(func(dictEqRecord_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictIsSymbol_2_box gopurs_runtime.Value, dictEq_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_eqRowCons(dictEqRecord_0_box, _dollar__unused_1_box, dictIsSymbol_2_box, dictEq_3_box))
})
	})
	return cache_eqRowCons
}

var cache_notEq gopurs_runtime.Value
var once_notEq sync.Once
func Get_notEq() gopurs_runtime.Value {
	once_notEq.Do(func() {
		cache_notEq = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq(dictEq_0_box, gopurs_runtime.UnboxAny(x_1_box), gopurs_runtime.UnboxAny(y_2_box)))
})
	})
	return cache_notEq
}

var cache_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619 gopurs_runtime.Value
var once_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619 sync.Once
func Get_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619() gopurs_runtime.Value {
	once_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619.Do(func() {
		cache_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619(dictEq_0_box, gopurs_runtime.UnboxAny(x_1_box), gopurs_runtime.UnboxAny(y_2_box)))
})
	})
	return cache_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619
}

var cache_notEq1 gopurs_runtime.Value
var once_notEq1 sync.Once
func Get_notEq1() gopurs_runtime.Value {
	once_notEq1.Do(func() {
		cache_notEq1 = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_notEq1(dictEq1_0_box, dictEq_1_box)
})
	})
	return cache_notEq1
}

var cache_eqArrayImpl gopurs_runtime.Value
var once_eqArrayImpl sync.Once
func Get_eqArrayImpl() gopurs_runtime.Value {
	once_eqArrayImpl.Do(func() {
		cache_eqArrayImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(EqArrayImpl(func(inner_arg0 interface{}, inner_arg1 interface{}) bool {
return (gopurs_runtime.Apply2(arg0, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)).IntVal) != (0)
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}(), func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_eqArrayImpl
}

var cache_eqBooleanImpl gopurs_runtime.Value
var once_eqBooleanImpl sync.Once
func Get_eqBooleanImpl() gopurs_runtime.Value {
	once_eqBooleanImpl.Do(func() {
		cache_eqBooleanImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(EqBooleanImpl((arg0.IntVal) != (0), (arg1.IntVal) != (0)))
})
	})
	return cache_eqBooleanImpl
}

var cache_eqCharImpl gopurs_runtime.Value
var once_eqCharImpl sync.Once
func Get_eqCharImpl() gopurs_runtime.Value {
	once_eqCharImpl.Do(func() {
		cache_eqCharImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(EqCharImpl(arg0.StrVal(), arg1.StrVal()))
})
	})
	return cache_eqCharImpl
}

var cache_eqIntImpl gopurs_runtime.Value
var once_eqIntImpl sync.Once
func Get_eqIntImpl() gopurs_runtime.Value {
	once_eqIntImpl.Do(func() {
		cache_eqIntImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(EqIntImpl(arg0.IntVal, arg1.IntVal))
})
	})
	return cache_eqIntImpl
}

var cache_eqNumberImpl gopurs_runtime.Value
var once_eqNumberImpl sync.Once
func Get_eqNumberImpl() gopurs_runtime.Value {
	once_eqNumberImpl.Do(func() {
		cache_eqNumberImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(EqNumberImpl(arg0.FloatVal(), arg1.FloatVal()))
})
	})
	return cache_eqNumberImpl
}

var cache_eqStringImpl gopurs_runtime.Value
var once_eqStringImpl sync.Once
func Get_eqStringImpl() gopurs_runtime.Value {
	once_eqStringImpl.Do(func() {
		cache_eqStringImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(EqStringImpl(arg0.StrVal(), arg1.StrVal()))
})
	})
	return cache_eqStringImpl
}

func Call_eqRecord(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "eqRecord")
}

func Call_eqRec(_dollar__unused_0_loop gopurs_runtime.Value, dictEqRecord_1_loop gopurs_runtime.Value) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictEqRecord_1 gopurs_runtime.Value = dictEqRecord_1_loop
_ = dictEqRecord_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEqRecord_1, "eqRecord"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))))
}

func Call_eq1(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "eq1")
}

func Call_eq1__func_gopurs_runtime_Value__gopurs_runtime_Value__interface____interface____bool_1105292882(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "eq1")
}

func Call_eq(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "eq")
}

func Call_eq__func_gopurs_runtime_Value__interface____interface____bool_1930883619(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "eq")
}

func Call_eqArray(dictEq_0_loop gopurs_runtime.Value) interface{} {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))))
}

func Call_eqRowCons(dictEqRecord_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictIsSymbol_2_loop gopurs_runtime.Value, dictEq_3_loop gopurs_runtime.Value) interface{} {
var dictEqRecord_0 gopurs_runtime.Value = dictEqRecord_0_loop
_ = dictEqRecord_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictIsSymbol_2 gopurs_runtime.Value = dictIsSymbol_2_loop
_ = dictIsSymbol_2
var dictEq_3 gopurs_runtime.Value = dictEq_3_loop
_ = dictEq_3
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, ra_5 gopurs_runtime.Value, rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
get_7_0 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_2, "reflectSymbol"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})))
_ = get_7_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_3, "eq"), gopurs_runtime.Apply(get_7_0, ra_5), gopurs_runtime.Apply(get_7_0, rb_6)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEqRecord_0, "eqRecord"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), ra_5, rb_6))
})))
}

func Call_notEq(dictEq_0_loop gopurs_runtime.Value, x_1_loop interface{}, y_2_loop interface{}) bool {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 interface{} = x_1_loop
_ = x_1
var y_2 interface{} = y_2_loop
_ = y_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Any(x_1), gopurs_runtime.Any(y_2)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_notEq__func_gopurs_runtime_Value__interface____interface____bool_1930883619(dictEq_0_loop gopurs_runtime.Value, x_1_loop interface{}, y_2_loop interface{}) bool {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 interface{} = x_1_loop
_ = x_1
var y_2 interface{} = y_2_loop
_ = y_2
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Any(x_1), gopurs_runtime.Any(y_2)), gopurs_runtime.Bool(false)).IntVal) != (0)
}

func Call_notEq1(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
eq12_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
_ = eq12_2_0
return gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(eq12_2_0, x_3, y_4), gopurs_runtime.Bool(false))
})
}

package Data_Eq

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var eqVoid gopurs_runtime.Value
var once_eqVoid sync.Once
func Get_eqVoid() gopurs_runtime.Value {
	once_eqVoid.Do(func() {
		eqVoid = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return eqVoid
}

var eqUnit gopurs_runtime.Value
var once_eqUnit sync.Once
func Get_eqUnit() gopurs_runtime.Value {
	once_eqUnit.Do(func() {
		eqUnit = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return eqUnit
}

var eqString gopurs_runtime.Value
var once_eqString sync.Once
func Get_eqString() gopurs_runtime.Value {
	once_eqString.Do(func() {
		eqString = gopurs_runtime.RecordDict1("eq", Get_eqStringImpl())
	})
	return eqString
}

var eqRowNil gopurs_runtime.Value
var once_eqRowNil sync.Once
func Get_eqRowNil() gopurs_runtime.Value {
	once_eqRowNil.Do(func() {
		eqRowNil = gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return eqRowNil
}

var eqRecord gopurs_runtime.Value
var once_eqRecord sync.Once
func Get_eqRecord() gopurs_runtime.Value {
	once_eqRecord.Do(func() {
		eqRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "eqRecord")
}()
})
	})
	return eqRecord
}

var eqRec gopurs_runtime.Value
var once_eqRec sync.Once
func Get_eqRec() gopurs_runtime.Value {
	once_eqRec.Do(func() {
		eqRec = gopurs_runtime.Func2(Call_eqRec)
	})
	return eqRec
}

var eqProxy gopurs_runtime.Value
var once_eqProxy sync.Once
func Get_eqProxy() gopurs_runtime.Value {
	once_eqProxy.Do(func() {
		eqProxy = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}))
	})
	return eqProxy
}

var eqNumber gopurs_runtime.Value
var once_eqNumber sync.Once
func Get_eqNumber() gopurs_runtime.Value {
	once_eqNumber.Do(func() {
		eqNumber = gopurs_runtime.RecordDict1("eq", Get_eqNumberImpl())
	})
	return eqNumber
}

var eqInt gopurs_runtime.Value
var once_eqInt sync.Once
func Get_eqInt() gopurs_runtime.Value {
	once_eqInt.Do(func() {
		eqInt = gopurs_runtime.RecordDict1("eq", Get_eqIntImpl())
	})
	return eqInt
}

var eqChar gopurs_runtime.Value
var once_eqChar sync.Once
func Get_eqChar() gopurs_runtime.Value {
	once_eqChar.Do(func() {
		eqChar = gopurs_runtime.RecordDict1("eq", Get_eqCharImpl())
	})
	return eqChar
}

var eqBoolean gopurs_runtime.Value
var once_eqBoolean sync.Once
func Get_eqBoolean() gopurs_runtime.Value {
	once_eqBoolean.Do(func() {
		eqBoolean = gopurs_runtime.RecordDict1("eq", Get_eqBooleanImpl())
	})
	return eqBoolean
}

var eq1 gopurs_runtime.Value
var once_eq1 sync.Once
func Get_eq1() gopurs_runtime.Value {
	once_eq1.Do(func() {
		eq1 = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "eq1")
}()
})
	})
	return eq1
}

var eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		eq = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "eq")
}()
})
	})
	return eq
}

var eqArray gopurs_runtime.Value
var once_eqArray sync.Once
func Get_eqArray() gopurs_runtime.Value {
	once_eqArray.Do(func() {
		eqArray = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0_loop, "eq")))
}()
})
	})
	return eqArray
}

var eq1Array gopurs_runtime.Value
var once_eq1Array sync.Once
func Get_eq1Array() gopurs_runtime.Value {
	once_eq1Array.Do(func() {
		eq1Array = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}))
	})
	return eq1Array
}

var eqRowCons gopurs_runtime.Value
var once_eqRowCons sync.Once
func Get_eqRowCons() gopurs_runtime.Value {
	once_eqRowCons.Do(func() {
		eqRowCons = gopurs_runtime.Func4(Call_eqRowCons)
	})
	return eqRowCons
}

var notEq gopurs_runtime.Value
var once_notEq sync.Once
func Get_notEq() gopurs_runtime.Value {
	once_notEq.Do(func() {
		notEq = gopurs_runtime.Func3(Call_notEq)
	})
	return notEq
}

var notEq1 gopurs_runtime.Value
var once_notEq1 sync.Once
func Get_notEq1() gopurs_runtime.Value {
	once_notEq1.Do(func() {
		notEq1 = gopurs_runtime.Func2(Call_notEq1)
	})
	return notEq1
}

func Call_eqRec(_dollar__unused_0_loop gopurs_runtime.Value, dictEqRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictEqRecord_1 gopurs_runtime.Value = dictEqRecord_1_loop
_ = dictEqRecord_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEqRecord_1_loop, "eqRecord"), gopurs_runtime.Constructor0("Proxy")))
}

func Call_eqRowCons(dictEqRecord_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictIsSymbol_2_loop gopurs_runtime.Value, dictEq_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEqRecord_0 gopurs_runtime.Value = dictEqRecord_0_loop
_ = dictEqRecord_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictIsSymbol_2 gopurs_runtime.Value = dictIsSymbol_2_loop
_ = dictIsSymbol_2
var dictEq_3 gopurs_runtime.Value = dictEq_3_loop
_ = dictEq_3
return gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, ra_5 gopurs_runtime.Value, rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
get_7_0 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_2_loop, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy")))
_ = get_7_0
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_3_loop, "eq"), gopurs_runtime.Apply(get_7_0, ra_5), gopurs_runtime.Apply(get_7_0, rb_6)).IntVal != 0 && gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEqRecord_0_loop, "eqRecord"), gopurs_runtime.Constructor0("Proxy"), ra_5, rb_6).IntVal != 0)
}))
}

func Call_notEq(dictEq_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return gopurs_runtime.Bool(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0_loop, "eq"), x_1_loop, y_2_loop).IntVal != 0 != true)
}

func Call_notEq1(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
eq12_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0_loop, "eq1"), dictEq_1_loop)
_ = eq12_2_0
return gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply2(eq12_2_0, x_3, y_4).IntVal != 0 != true)
})
}

func Get_eqArrayImpl() gopurs_runtime.Value {
	return _Gopurs_EqArrayImpl
}

func Get_eqBooleanImpl() gopurs_runtime.Value {
	return _Gopurs_EqBooleanImpl
}

func Get_eqCharImpl() gopurs_runtime.Value {
	return _Gopurs_EqCharImpl
}

func Get_eqIntImpl() gopurs_runtime.Value {
	return _Gopurs_EqIntImpl
}

func Get_eqNumberImpl() gopurs_runtime.Value {
	return _Gopurs_EqNumberImpl
}

func Get_eqStringImpl() gopurs_runtime.Value {
	return _Gopurs_EqStringImpl
}

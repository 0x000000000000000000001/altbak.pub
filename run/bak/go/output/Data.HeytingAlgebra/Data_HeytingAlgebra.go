package Data_HeytingAlgebra

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
)

var ttRecord gopurs_runtime.Value
var once_ttRecord sync.Once
func Get_ttRecord() gopurs_runtime.Value {
	once_ttRecord.Do(func() {
		ttRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ttRecord")
}()
})
	})
	return ttRecord
}

var tt gopurs_runtime.Value
var once_tt sync.Once
func Get_tt() gopurs_runtime.Value {
	once_tt.Do(func() {
		tt = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}()
})
	})
	return tt
}

var notRecord gopurs_runtime.Value
var once_notRecord sync.Once
func Get_notRecord() gopurs_runtime.Value {
	once_notRecord.Do(func() {
		notRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "notRecord")
}()
})
	})
	return notRecord
}

var not gopurs_runtime.Value
var once_not sync.Once
func Get_not() gopurs_runtime.Value {
	once_not.Do(func() {
		not = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "not")
}()
})
	})
	return not
}

var impliesRecord gopurs_runtime.Value
var once_impliesRecord sync.Once
func Get_impliesRecord() gopurs_runtime.Value {
	once_impliesRecord.Do(func() {
		impliesRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "impliesRecord")
}()
})
	})
	return impliesRecord
}

var implies gopurs_runtime.Value
var once_implies sync.Once
func Get_implies() gopurs_runtime.Value {
	once_implies.Do(func() {
		implies = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "implies")
}()
})
	})
	return implies
}

var heytingAlgebraUnit gopurs_runtime.Value
var once_heytingAlgebraUnit sync.Once
func Get_heytingAlgebraUnit() gopurs_runtime.Value {
	once_heytingAlgebraUnit.Do(func() {
		heytingAlgebraUnit = gopurs_runtime.RecordDict([]string{"ff", "tt", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{pkg_Data_Unit.Get_unit(), pkg_Data_Unit.Get_unit(), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})})
	})
	return heytingAlgebraUnit
}

var heytingAlgebraRecordNil gopurs_runtime.Value
var once_heytingAlgebraRecordNil sync.Once
func Get_heytingAlgebraRecordNil() gopurs_runtime.Value {
	once_heytingAlgebraRecordNil.Do(func() {
		heytingAlgebraRecordNil = gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "ffRecord", "impliesRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})})
	})
	return heytingAlgebraRecordNil
}

var heytingAlgebraProxy gopurs_runtime.Value
var once_heytingAlgebraProxy sync.Once
func Get_heytingAlgebraProxy() gopurs_runtime.Value {
	once_heytingAlgebraProxy.Do(func() {
		heytingAlgebraProxy = gopurs_runtime.RecordDict([]string{"conj", "disj", "implies", "ff", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor0("Proxy")
}), gopurs_runtime.Constructor0("Proxy")})
	})
	return heytingAlgebraProxy
}

var ffRecord gopurs_runtime.Value
var once_ffRecord sync.Once
func Get_ffRecord() gopurs_runtime.Value {
	once_ffRecord.Do(func() {
		ffRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ffRecord")
}()
})
	})
	return ffRecord
}

var ff gopurs_runtime.Value
var once_ff sync.Once
func Get_ff() gopurs_runtime.Value {
	once_ff.Do(func() {
		ff = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}()
})
	})
	return ff
}

var disjRecord gopurs_runtime.Value
var once_disjRecord sync.Once
func Get_disjRecord() gopurs_runtime.Value {
	once_disjRecord.Do(func() {
		disjRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "disjRecord")
}()
})
	})
	return disjRecord
}

var disj gopurs_runtime.Value
var once_disj sync.Once
func Get_disj() gopurs_runtime.Value {
	once_disj.Do(func() {
		disj = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "disj")
}()
})
	})
	return disj
}

var heytingAlgebraBoolean gopurs_runtime.Value
var once_heytingAlgebraBoolean sync.Once
func Get_heytingAlgebraBoolean() gopurs_runtime.Value {
	once_heytingAlgebraBoolean.Do(func() {
		heytingAlgebraBoolean = gopurs_runtime.RecordDict([]string{"ff", "tt", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_heytingAlgebraBoolean(), "not"), a_0), b_1)
}), Get_boolConj(), Get_boolDisj(), Get_boolNot()})
	})
	return heytingAlgebraBoolean
}

var conjRecord gopurs_runtime.Value
var once_conjRecord sync.Once
func Get_conjRecord() gopurs_runtime.Value {
	once_conjRecord.Do(func() {
		conjRecord = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "conjRecord")
}()
})
	})
	return conjRecord
}

var heytingAlgebraRecord gopurs_runtime.Value
var once_heytingAlgebraRecord sync.Once
func Get_heytingAlgebraRecord() gopurs_runtime.Value {
	once_heytingAlgebraRecord.Do(func() {
		heytingAlgebraRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictHeytingAlgebraRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heytingAlgebraRecord(_dollar__unused_0_box, dictHeytingAlgebraRecord_1_box)
})
	})
	return heytingAlgebraRecord
}

var conj gopurs_runtime.Value
var once_conj sync.Once
func Get_conj() gopurs_runtime.Value {
	once_conj.Do(func() {
		conj = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "conj")
}()
})
	})
	return conj
}

var heytingAlgebraFunction gopurs_runtime.Value
var once_heytingAlgebraFunction sync.Once
func Get_heytingAlgebraFunction() gopurs_runtime.Value {
	once_heytingAlgebraFunction.Do(func() {
		heytingAlgebraFunction = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
ff1_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
_ = ff1_1_0
tt1_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
_ = tt1_2_1
return gopurs_runtime.RecordDict([]string{"ff", "tt", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ff1_1_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return tt1_2_1
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), gopurs_runtime.Apply(f_3, a_5), gopurs_runtime.Apply(g_4, a_5))
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), gopurs_runtime.Apply(f_3, a_5), gopurs_runtime.Apply(g_4, a_5))
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), gopurs_runtime.Apply(f_3, a_5), gopurs_runtime.Apply(g_4, a_5))
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), gopurs_runtime.Apply(f_3, a_4))
})})
}()
})
	})
	return heytingAlgebraFunction
}

var heytingAlgebraRecordCons gopurs_runtime.Value
var once_heytingAlgebraRecordCons sync.Once
func Get_heytingAlgebraRecordCons() gopurs_runtime.Value {
	once_heytingAlgebraRecordCons.Do(func() {
		heytingAlgebraRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictHeytingAlgebraRecord_2_box gopurs_runtime.Value, dictHeytingAlgebra_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heytingAlgebraRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictHeytingAlgebraRecord_2_box, dictHeytingAlgebra_3_box)
})
	})
	return heytingAlgebraRecordCons
}

func Call_heytingAlgebraRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictHeytingAlgebraRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictHeytingAlgebraRecord_1 gopurs_runtime.Value = dictHeytingAlgebraRecord_1_loop
_ = dictHeytingAlgebraRecord_1
return gopurs_runtime.RecordDict([]string{"ff", "tt", "conj", "disj", "implies", "not"}, []gopurs_runtime.Value{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "ffRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "ttRecord"), gopurs_runtime.Constructor0("Proxy"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "conjRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "disjRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "impliesRecord"), gopurs_runtime.Constructor0("Proxy")), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "notRecord"), gopurs_runtime.Constructor0("Proxy"))})
}

func Call_heytingAlgebraRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictHeytingAlgebraRecord_2_loop gopurs_runtime.Value, dictHeytingAlgebra_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictHeytingAlgebraRecord_2 gopurs_runtime.Value = dictHeytingAlgebraRecord_2_loop
_ = dictHeytingAlgebraRecord_2
var dictHeytingAlgebra_3 gopurs_runtime.Value = dictHeytingAlgebra_3_loop
_ = dictHeytingAlgebra_3
ff1_4_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "ff")
_ = ff1_4_0
tt1_5_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "tt")
_ = tt1_5_1
return gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "impliesRecord", "ffRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_9_2
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_2)
_ = get_10_3
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "conj"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "conjRecord"), gopurs_runtime.Constructor0("Proxy"), ra_7, rb_8))
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_9_4
get_10_5 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_4)
_ = get_10_5
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "disj"), gopurs_runtime.Apply(get_10_5, ra_7), gopurs_runtime.Apply(get_10_5, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "disjRecord"), gopurs_runtime.Constructor0("Proxy"), ra_7, rb_8))
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_9_6
get_10_7 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), key_9_6)
_ = get_10_7
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_9_6, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "implies"), gopurs_runtime.Apply(get_10_7, ra_7), gopurs_runtime.Apply(get_10_7, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "impliesRecord"), gopurs_runtime.Constructor0("Proxy"), ra_7, rb_8))
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy")), ff1_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "ffRecord"), gopurs_runtime.Constructor0("Proxy"), row_7))
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, row_7 gopurs_runtime.Value) gopurs_runtime.Value {
key_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy"))
_ = key_8_8
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), key_8_8, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "not"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_8_8, row_7)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "notRecord"), gopurs_runtime.Constructor0("Proxy"), row_7))
}), gopurs_runtime.Func2(func(v_6 gopurs_runtime.Value, row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Constructor0("Proxy")), tt1_5_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "ttRecord"), gopurs_runtime.Constructor0("Proxy"), row_7))
})})
}

func Get_boolConj() gopurs_runtime.Value {
	return _Gopurs_BoolConj
}

func Get_boolDisj() gopurs_runtime.Value {
	return _Gopurs_BoolDisj
}

func Get_boolNot() gopurs_runtime.Value {
	return _Gopurs_BoolNot
}

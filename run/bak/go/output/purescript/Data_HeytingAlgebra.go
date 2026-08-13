package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict gopurs_runtime.Value
var once_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict sync.Once
func Get_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict.Do(func() {
		cache_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict
}

var cache_Data_HeytingAlgebra_HeytingAlgebra_dollarDict gopurs_runtime.Value
var once_Data_HeytingAlgebra_HeytingAlgebra_dollarDict sync.Once
func Get_Data_HeytingAlgebra_HeytingAlgebra_dollarDict() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_HeytingAlgebra_dollarDict.Do(func() {
		cache_Data_HeytingAlgebra_HeytingAlgebra_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_HeytingAlgebra_dollarDict(x_0_box)
})
	})
	return cache_Data_HeytingAlgebra_HeytingAlgebra_dollarDict
}

var cache_Data_HeytingAlgebra_ttRecord gopurs_runtime.Value
var once_Data_HeytingAlgebra_ttRecord sync.Once
func Get_Data_HeytingAlgebra_ttRecord() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ttRecord.Do(func() {
		cache_Data_HeytingAlgebra_ttRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ttRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_ttRecord
}

var cache_Data_HeytingAlgebra_tt gopurs_runtime.Value
var once_Data_HeytingAlgebra_tt sync.Once
func Get_Data_HeytingAlgebra_tt() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_tt.Do(func() {
		cache_Data_HeytingAlgebra_tt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_tt(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_tt
}

var cache_Data_HeytingAlgebra_notRecord gopurs_runtime.Value
var once_Data_HeytingAlgebra_notRecord sync.Once
func Get_Data_HeytingAlgebra_notRecord() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_notRecord.Do(func() {
		cache_Data_HeytingAlgebra_notRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_notRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_notRecord
}

var cache_Data_HeytingAlgebra_not gopurs_runtime.Value
var once_Data_HeytingAlgebra_not sync.Once
func Get_Data_HeytingAlgebra_not() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_not.Do(func() {
		cache_Data_HeytingAlgebra_not = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_not(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_not
}

var cache_Data_HeytingAlgebra_impliesRecord gopurs_runtime.Value
var once_Data_HeytingAlgebra_impliesRecord sync.Once
func Get_Data_HeytingAlgebra_impliesRecord() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_impliesRecord.Do(func() {
		cache_Data_HeytingAlgebra_impliesRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_impliesRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_impliesRecord
}

var cache_Data_HeytingAlgebra_implies gopurs_runtime.Value
var once_Data_HeytingAlgebra_implies sync.Once
func Get_Data_HeytingAlgebra_implies() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_implies.Do(func() {
		cache_Data_HeytingAlgebra_implies = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_implies(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_implies
}

var cache_Data_HeytingAlgebra_heytingAlgebraUnit gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraUnit sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraUnit() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraUnit.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraUnit = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), Get_Data_Unit_unit(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Data_Unit_unit()})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraUnit
}

var cache_Data_HeytingAlgebra_heytingAlgebraRecordNil gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraRecordNil sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraRecordNil() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraRecordNil.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraRecordNil = gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "ffRecord", "impliesRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraRecordNil
}

var cache_Data_HeytingAlgebra_heytingAlgebraProxy gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraProxy sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraProxy() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraProxy.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraProxy = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraProxy
}

var cache_Data_HeytingAlgebra_ffRecord gopurs_runtime.Value
var once_Data_HeytingAlgebra_ffRecord sync.Once
func Get_Data_HeytingAlgebra_ffRecord() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ffRecord.Do(func() {
		cache_Data_HeytingAlgebra_ffRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ffRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_ffRecord
}

var cache_Data_HeytingAlgebra_ff gopurs_runtime.Value
var once_Data_HeytingAlgebra_ff sync.Once
func Get_Data_HeytingAlgebra_ff() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ff.Do(func() {
		cache_Data_HeytingAlgebra_ff = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ff(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_ff
}

var cache_Data_HeytingAlgebra_disjRecord gopurs_runtime.Value
var once_Data_HeytingAlgebra_disjRecord sync.Once
func Get_Data_HeytingAlgebra_disjRecord() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_disjRecord.Do(func() {
		cache_Data_HeytingAlgebra_disjRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_disjRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_disjRecord
}

var cache_Data_HeytingAlgebra_disj gopurs_runtime.Value
var once_Data_HeytingAlgebra_disj sync.Once
func Get_Data_HeytingAlgebra_disj() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_disj.Do(func() {
		cache_Data_HeytingAlgebra_disj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_disj(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_disj
}

var cache_Data_HeytingAlgebra_heytingAlgebraBoolean gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraBoolean sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraBoolean() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraBoolean.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraBoolean = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{Get_Data_HeytingAlgebra_boolConj(), Get_Data_HeytingAlgebra_boolDisj(), gopurs_runtime.Bool(false), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_Data_HeytingAlgebra_disj__3676519832(gopurs_runtime.Bool((Call_Data_HeytingAlgebra_not__3201284355(gopurs_runtime.Bool((a_0.IntVal) != (0))).IntVal) != (0)), gopurs_runtime.Bool((b_1.IntVal) != (0))).IntVal) != (0))
})
}), Get_Data_HeytingAlgebra_boolNot(), gopurs_runtime.Bool(true)})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraBoolean
}

var cache_Data_HeytingAlgebra_conjRecord gopurs_runtime.Value
var once_Data_HeytingAlgebra_conjRecord sync.Once
func Get_Data_HeytingAlgebra_conjRecord() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conjRecord.Do(func() {
		cache_Data_HeytingAlgebra_conjRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conjRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conjRecord
}

var cache_Data_HeytingAlgebra_heytingAlgebraRecord gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraRecord sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraRecord() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraRecord.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictHeytingAlgebraRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_heytingAlgebraRecord(_dollar__unused_0_box, dictHeytingAlgebraRecord_1_box)
})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraRecord
}

var cache_Data_HeytingAlgebra_conj gopurs_runtime.Value
var once_Data_HeytingAlgebra_conj sync.Once
func Get_Data_HeytingAlgebra_conj() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conj.Do(func() {
		cache_Data_HeytingAlgebra_conj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conj(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conj
}

var cache_Data_HeytingAlgebra_heytingAlgebraFunction gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraFunction sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraFunction() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraFunction.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraFunction = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_heytingAlgebraFunction(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraFunction
}

var cache_Data_HeytingAlgebra_heytingAlgebraRecordCons gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraRecordCons sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraRecordCons() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraRecordCons.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictHeytingAlgebraRecord_2_box gopurs_runtime.Value, dictHeytingAlgebra_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_heytingAlgebraRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictHeytingAlgebraRecord_2_box, dictHeytingAlgebra_3_box)
})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraRecordCons
}

var cache_Data_HeytingAlgebra_conj__2927892844 gopurs_runtime.Value
var once_Data_HeytingAlgebra_conj__2927892844 sync.Once
func Get_Data_HeytingAlgebra_conj__2927892844() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conj__2927892844.Do(func() {
		cache_Data_HeytingAlgebra_conj__2927892844 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conj__2927892844(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conj__2927892844
}

var cache_Data_HeytingAlgebra_conj__4093645121 gopurs_runtime.Value
var once_Data_HeytingAlgebra_conj__4093645121 sync.Once
func Get_Data_HeytingAlgebra_conj__4093645121() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conj__4093645121.Do(func() {
		cache_Data_HeytingAlgebra_conj__4093645121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conj__4093645121(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conj__4093645121
}

var cache_Data_HeytingAlgebra_conj__204561377 gopurs_runtime.Value
var once_Data_HeytingAlgebra_conj__204561377 sync.Once
func Get_Data_HeytingAlgebra_conj__204561377() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conj__204561377.Do(func() {
		cache_Data_HeytingAlgebra_conj__204561377 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conj__204561377(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conj__204561377
}

var cache_Data_HeytingAlgebra_conj__3676519832 gopurs_runtime.Value
var once_Data_HeytingAlgebra_conj__3676519832 sync.Once
func Get_Data_HeytingAlgebra_conj__3676519832() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conj__3676519832.Do(func() {
		cache_Data_HeytingAlgebra_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_HeytingAlgebra_conj__3676519832
}

var cache_Data_HeytingAlgebra_conj__3472268504 gopurs_runtime.Value
var once_Data_HeytingAlgebra_conj__3472268504 sync.Once
func Get_Data_HeytingAlgebra_conj__3472268504() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conj__3472268504.Do(func() {
		cache_Data_HeytingAlgebra_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conj__3472268504(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conj__3472268504
}

var cache_Data_HeytingAlgebra_conjRecord__2439193216 gopurs_runtime.Value
var once_Data_HeytingAlgebra_conjRecord__2439193216 sync.Once
func Get_Data_HeytingAlgebra_conjRecord__2439193216() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conjRecord__2439193216.Do(func() {
		cache_Data_HeytingAlgebra_conjRecord__2439193216 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conjRecord__2439193216(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conjRecord__2439193216
}

var cache_Data_HeytingAlgebra_conjRecord__497482630 gopurs_runtime.Value
var once_Data_HeytingAlgebra_conjRecord__497482630 sync.Once
func Get_Data_HeytingAlgebra_conjRecord__497482630() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_conjRecord__497482630.Do(func() {
		cache_Data_HeytingAlgebra_conjRecord__497482630 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_conjRecord__497482630(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_conjRecord__497482630
}

var cache_Data_HeytingAlgebra_disj__3676519832 gopurs_runtime.Value
var once_Data_HeytingAlgebra_disj__3676519832 sync.Once
func Get_Data_HeytingAlgebra_disj__3676519832() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_disj__3676519832.Do(func() {
		cache_Data_HeytingAlgebra_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_Data_HeytingAlgebra_disj__3676519832
}

var cache_Data_HeytingAlgebra_disj__3472268504 gopurs_runtime.Value
var once_Data_HeytingAlgebra_disj__3472268504 sync.Once
func Get_Data_HeytingAlgebra_disj__3472268504() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_disj__3472268504.Do(func() {
		cache_Data_HeytingAlgebra_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_disj__3472268504(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_disj__3472268504
}

var cache_Data_HeytingAlgebra_disjRecord__2439193216 gopurs_runtime.Value
var once_Data_HeytingAlgebra_disjRecord__2439193216 sync.Once
func Get_Data_HeytingAlgebra_disjRecord__2439193216() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_disjRecord__2439193216.Do(func() {
		cache_Data_HeytingAlgebra_disjRecord__2439193216 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_disjRecord__2439193216(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_disjRecord__2439193216
}

var cache_Data_HeytingAlgebra_disjRecord__497482630 gopurs_runtime.Value
var once_Data_HeytingAlgebra_disjRecord__497482630 sync.Once
func Get_Data_HeytingAlgebra_disjRecord__497482630() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_disjRecord__497482630.Do(func() {
		cache_Data_HeytingAlgebra_disjRecord__497482630 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_disjRecord__497482630(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_disjRecord__497482630
}

var cache_Data_HeytingAlgebra_ff__2527024921 gopurs_runtime.Value
var once_Data_HeytingAlgebra_ff__2527024921 sync.Once
func Get_Data_HeytingAlgebra_ff__2527024921() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ff__2527024921.Do(func() {
		cache_Data_HeytingAlgebra_ff__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ff__2527024921(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_ff__2527024921
}

var cache_Data_HeytingAlgebra_ffRecord__2798009952 gopurs_runtime.Value
var once_Data_HeytingAlgebra_ffRecord__2798009952 sync.Once
func Get_Data_HeytingAlgebra_ffRecord__2798009952() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ffRecord__2798009952.Do(func() {
		cache_Data_HeytingAlgebra_ffRecord__2798009952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ffRecord__2798009952(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_ffRecord__2798009952
}

var cache_Data_HeytingAlgebra_ffRecord__465956064 gopurs_runtime.Value
var once_Data_HeytingAlgebra_ffRecord__465956064 sync.Once
func Get_Data_HeytingAlgebra_ffRecord__465956064() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ffRecord__465956064.Do(func() {
		cache_Data_HeytingAlgebra_ffRecord__465956064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ffRecord__465956064(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_ffRecord__465956064
}

var cache_Data_HeytingAlgebra_heytingAlgebraProxy__549257500 gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraProxy__549257500 sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraProxy__549257500() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraProxy__549257500.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraProxy__549257500 = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraProxy__549257500
}

var cache_Data_HeytingAlgebra_heytingAlgebraRecordNil__28213168 gopurs_runtime.Value
var once_Data_HeytingAlgebra_heytingAlgebraRecordNil__28213168 sync.Once
func Get_Data_HeytingAlgebra_heytingAlgebraRecordNil__28213168() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_heytingAlgebraRecordNil__28213168.Do(func() {
		cache_Data_HeytingAlgebra_heytingAlgebraRecordNil__28213168 = gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "ffRecord", "impliesRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})
	})
	return cache_Data_HeytingAlgebra_heytingAlgebraRecordNil__28213168
}

var cache_Data_HeytingAlgebra_implies__3472268504 gopurs_runtime.Value
var once_Data_HeytingAlgebra_implies__3472268504 sync.Once
func Get_Data_HeytingAlgebra_implies__3472268504() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_implies__3472268504.Do(func() {
		cache_Data_HeytingAlgebra_implies__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_implies__3472268504(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_implies__3472268504
}

var cache_Data_HeytingAlgebra_impliesRecord__2439193216 gopurs_runtime.Value
var once_Data_HeytingAlgebra_impliesRecord__2439193216 sync.Once
func Get_Data_HeytingAlgebra_impliesRecord__2439193216() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_impliesRecord__2439193216.Do(func() {
		cache_Data_HeytingAlgebra_impliesRecord__2439193216 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_impliesRecord__2439193216(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_impliesRecord__2439193216
}

var cache_Data_HeytingAlgebra_impliesRecord__497482630 gopurs_runtime.Value
var once_Data_HeytingAlgebra_impliesRecord__497482630 sync.Once
func Get_Data_HeytingAlgebra_impliesRecord__497482630() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_impliesRecord__497482630.Do(func() {
		cache_Data_HeytingAlgebra_impliesRecord__497482630 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_impliesRecord__497482630(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_impliesRecord__497482630
}

var cache_Data_HeytingAlgebra_not__3201284355 gopurs_runtime.Value
var once_Data_HeytingAlgebra_not__3201284355 sync.Once
func Get_Data_HeytingAlgebra_not__3201284355() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_not__3201284355.Do(func() {
		cache_Data_HeytingAlgebra_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_not__3201284355(__eta0_0_box)
})
	})
	return cache_Data_HeytingAlgebra_not__3201284355
}

var cache_Data_HeytingAlgebra_not__1505204753 gopurs_runtime.Value
var once_Data_HeytingAlgebra_not__1505204753 sync.Once
func Get_Data_HeytingAlgebra_not__1505204753() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_not__1505204753.Do(func() {
		cache_Data_HeytingAlgebra_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_not__1505204753(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_not__1505204753
}

var cache_Data_HeytingAlgebra_not__2235433470 gopurs_runtime.Value
var once_Data_HeytingAlgebra_not__2235433470 sync.Once
func Get_Data_HeytingAlgebra_not__2235433470() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_not__2235433470.Do(func() {
		cache_Data_HeytingAlgebra_not__2235433470 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_not__2235433470(__eta0_0_box)
})
	})
	return cache_Data_HeytingAlgebra_not__2235433470
}

var cache_Data_HeytingAlgebra_notRecord__3181681457 gopurs_runtime.Value
var once_Data_HeytingAlgebra_notRecord__3181681457 sync.Once
func Get_Data_HeytingAlgebra_notRecord__3181681457() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_notRecord__3181681457.Do(func() {
		cache_Data_HeytingAlgebra_notRecord__3181681457 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_notRecord__3181681457(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_notRecord__3181681457
}

var cache_Data_HeytingAlgebra_notRecord__726562039 gopurs_runtime.Value
var once_Data_HeytingAlgebra_notRecord__726562039 sync.Once
func Get_Data_HeytingAlgebra_notRecord__726562039() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_notRecord__726562039.Do(func() {
		cache_Data_HeytingAlgebra_notRecord__726562039 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_notRecord__726562039(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_notRecord__726562039
}

var cache_Data_HeytingAlgebra_tt__2527024921 gopurs_runtime.Value
var once_Data_HeytingAlgebra_tt__2527024921 sync.Once
func Get_Data_HeytingAlgebra_tt__2527024921() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_tt__2527024921.Do(func() {
		cache_Data_HeytingAlgebra_tt__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_tt__2527024921(dict_0_box)
})
	})
	return cache_Data_HeytingAlgebra_tt__2527024921
}

var cache_Data_HeytingAlgebra_ttRecord__2798009952 gopurs_runtime.Value
var once_Data_HeytingAlgebra_ttRecord__2798009952 sync.Once
func Get_Data_HeytingAlgebra_ttRecord__2798009952() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ttRecord__2798009952.Do(func() {
		cache_Data_HeytingAlgebra_ttRecord__2798009952 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ttRecord__2798009952(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_ttRecord__2798009952
}

var cache_Data_HeytingAlgebra_ttRecord__465956064 gopurs_runtime.Value
var once_Data_HeytingAlgebra_ttRecord__465956064 sync.Once
func Get_Data_HeytingAlgebra_ttRecord__465956064() gopurs_runtime.Value {
	once_Data_HeytingAlgebra_ttRecord__465956064.Do(func() {
		cache_Data_HeytingAlgebra_ttRecord__465956064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_HeytingAlgebra_ttRecord__465956064(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord](dict_0_box))
})
	})
	return cache_Data_HeytingAlgebra_ttRecord__465956064
}

type Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3558753879] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord)(ptr)
		_ = c
		switch key {
		case "conjRecord": return gopurs_runtime.Box(c.V0)
		case "disjRecord": return gopurs_runtime.Box(c.V1)
		case "ffRecord": return gopurs_runtime.Box(c.V2)
		case "impliesRecord": return gopurs_runtime.Box(c.V3)
		case "notRecord": return gopurs_runtime.Box(c.V4)
		case "ttRecord": return gopurs_runtime.Box(c.V5)
		default: panic("Key not found in dictionary Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord: " + key)
		}
	}
}


type Constructor_Data_HeytingAlgebra_HeytingAlgebra struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[926771738] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_HeytingAlgebra_HeytingAlgebra)(ptr)
		_ = c
		switch key {
		case "conj": return gopurs_runtime.Box(c.V0)
		case "disj": return gopurs_runtime.Box(c.V1)
		case "ff": return gopurs_runtime.Box(c.V2)
		case "implies": return gopurs_runtime.Box(c.V3)
		case "not": return gopurs_runtime.Box(c.V4)
		case "tt": return gopurs_runtime.Box(c.V5)
		default: panic("Key not found in dictionary Constructor_Data_HeytingAlgebra_HeytingAlgebra: " + key)
		}
	}
}


func Call_Data_HeytingAlgebra_HeytingAlgebraRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_HeytingAlgebra_HeytingAlgebra_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_HeytingAlgebra_ttRecord(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Data_HeytingAlgebra_tt(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}

func Call_Data_HeytingAlgebra_notRecord(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_not(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_impliesRecord(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_implies(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_ffRecord(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_HeytingAlgebra_ff(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}

func Call_Data_HeytingAlgebra_disjRecord(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_disj(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_conjRecord(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_heytingAlgebraRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictHeytingAlgebraRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictHeytingAlgebraRecord_1 gopurs_runtime.Value = dictHeytingAlgebraRecord_1_loop
_ = dictHeytingAlgebraRecord_1
return gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "conjRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "disjRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "ffRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "impliesRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "notRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "ttRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
}

func Call_Data_HeytingAlgebra_conj(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_heytingAlgebraFunction(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
return gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), gopurs_runtime.Apply(f_1, a_3), gopurs_runtime.Apply(g_2, a_3))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), gopurs_runtime.Apply(f_1, a_3), gopurs_runtime.Apply(g_2, a_3))
})
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), gopurs_runtime.Apply(f_1, a_3), gopurs_runtime.Apply(g_2, a_3))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), gopurs_runtime.Apply(f_1, a_2))
})
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
})})
}

func Call_Data_HeytingAlgebra_heytingAlgebraRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictHeytingAlgebraRecord_2_loop gopurs_runtime.Value, dictHeytingAlgebra_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictHeytingAlgebraRecord_2 gopurs_runtime.Value = dictHeytingAlgebraRecord_2_loop
_ = dictHeytingAlgebraRecord_2
var dictHeytingAlgebra_3 gopurs_runtime.Value = dictHeytingAlgebra_3_loop
_ = dictHeytingAlgebra_3
// TAST (Let): ff1_4_0 -> gopurs_runtime.Value
ff1_4_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "ff")
_ = ff1_4_0
// TAST (Let): tt1_5_1 -> gopurs_runtime.Value
tt1_5_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "tt")
_ = tt1_5_1
return gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "ffRecord", "impliesRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_9_2 -> gopurs_runtime.Value
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_9_2
// TAST (Let): get_10_3 -> gopurs_runtime.Value
get_10_3 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_9_2.StrVal()))
_ = get_10_3
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_2.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "conj"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "conjRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
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
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_4.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "disj"), gopurs_runtime.Apply(get_10_5, ra_7), gopurs_runtime.Apply(get_10_5, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "disjRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), ff1_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "ffRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(row_7.IntVal)), UnsafePtr: nil}))
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_9_6 -> gopurs_runtime.Value
key_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_9_6
// TAST (Let): get_10_7 -> gopurs_runtime.Value
get_10_7 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_9_6.StrVal()))
_ = get_10_7
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_6.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "implies"), gopurs_runtime.Apply(get_10_7, ra_7), gopurs_runtime.Apply(get_10_7, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "impliesRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_8_8 -> gopurs_runtime.Value
key_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_8_8
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_8_8.StrVal()), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "not"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_8_8.StrVal()), row_7)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "notRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, row_7))
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), tt1_5_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "ttRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(row_7.IntVal)), UnsafePtr: nil}))
})
})})
}

func Call_Data_HeytingAlgebra_conj__2927892844(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_conj__4093645121(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_conj__204561377(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_Data_HeytingAlgebra_conj__3472268504(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_conjRecord__2439193216(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_conjRecord__497482630(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_HeytingAlgebra_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_Data_HeytingAlgebra_disj__3472268504(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_disjRecord__2439193216(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_disjRecord__497482630(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_HeytingAlgebra_ff__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}

func Call_Data_HeytingAlgebra_ffRecord__2798009952(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_HeytingAlgebra_ffRecord__465956064(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_HeytingAlgebra_implies__3472268504(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_impliesRecord__2439193216(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_impliesRecord__497482630(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Data_HeytingAlgebra_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_Data_HeytingAlgebra_not__1505204753(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebra) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebra = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_not__2235433470(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra](Get_Data_Interval_Duration_Iso_heytingAlgebraFunction()).V4), __eta0_0)
}

func Call_Data_HeytingAlgebra_notRecord__3181681457(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_notRecord__726562039(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Data_HeytingAlgebra_tt__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}

func Call_Data_HeytingAlgebra_ttRecord__2798009952(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Data_HeytingAlgebra_ttRecord__465956064(dict_0_loop *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_HeytingAlgebra_HeytingAlgebraRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Get_Data_HeytingAlgebra_boolConj() gopurs_runtime.Value {
	return _Gopurs_Data_HeytingAlgebra_BoolConj
}

func Get_Data_HeytingAlgebra_boolDisj() gopurs_runtime.Value {
	return _Gopurs_Data_HeytingAlgebra_BoolDisj
}

func Get_Data_HeytingAlgebra_boolNot() gopurs_runtime.Value {
	return _Gopurs_Data_HeytingAlgebra_BoolNot
}

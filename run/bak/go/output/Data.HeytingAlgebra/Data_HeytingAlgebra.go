package Data_HeytingAlgebra

import (
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_ttRecord gopurs_runtime.Value
var once_ttRecord sync.Once
func Get_ttRecord() gopurs_runtime.Value {
	once_ttRecord.Do(func() {
		cache_ttRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ttRecord(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_ttRecord
}

var cache_ttRecord__gopurs_runtime_Value_465956064 gopurs_runtime.Value
var once_ttRecord__gopurs_runtime_Value_465956064 sync.Once
func Get_ttRecord__gopurs_runtime_Value_465956064() gopurs_runtime.Value {
	once_ttRecord__gopurs_runtime_Value_465956064.Do(func() {
		cache_ttRecord__gopurs_runtime_Value_465956064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ttRecord__gopurs_runtime_Value_465956064(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_ttRecord__gopurs_runtime_Value_465956064
}

var cache_tt gopurs_runtime.Value
var once_tt sync.Once
func Get_tt() gopurs_runtime.Value {
	once_tt.Do(func() {
		cache_tt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tt(dict_0_box)
})
	})
	return cache_tt
}

var cache_tt__gopurs_runtime_Value_2527024921 gopurs_runtime.Value
var once_tt__gopurs_runtime_Value_2527024921 sync.Once
func Get_tt__gopurs_runtime_Value_2527024921() gopurs_runtime.Value {
	once_tt__gopurs_runtime_Value_2527024921.Do(func() {
		cache_tt__gopurs_runtime_Value_2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tt__gopurs_runtime_Value_2527024921(dict_0_box)
})
	})
	return cache_tt__gopurs_runtime_Value_2527024921
}

var cache_notRecord gopurs_runtime.Value
var once_notRecord sync.Once
func Get_notRecord() gopurs_runtime.Value {
	once_notRecord.Do(func() {
		cache_notRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_notRecord(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_notRecord
}

var cache_notRecord__gopurs_runtime_Value_726562039 gopurs_runtime.Value
var once_notRecord__gopurs_runtime_Value_726562039 sync.Once
func Get_notRecord__gopurs_runtime_Value_726562039() gopurs_runtime.Value {
	once_notRecord__gopurs_runtime_Value_726562039.Do(func() {
		cache_notRecord__gopurs_runtime_Value_726562039 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_notRecord__gopurs_runtime_Value_726562039(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_notRecord__gopurs_runtime_Value_726562039
}

var cache_not gopurs_runtime.Value
var once_not sync.Once
func Get_not() gopurs_runtime.Value {
	once_not.Do(func() {
		cache_not = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not
}

var cache_not__gopurs_runtime_Value_1505204753 gopurs_runtime.Value
var once_not__gopurs_runtime_Value_1505204753 sync.Once
func Get_not__gopurs_runtime_Value_1505204753() gopurs_runtime.Value {
	once_not__gopurs_runtime_Value_1505204753.Do(func() {
		cache_not__gopurs_runtime_Value_1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__gopurs_runtime_Value_1505204753(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__gopurs_runtime_Value_1505204753
}

var cache_impliesRecord gopurs_runtime.Value
var once_impliesRecord sync.Once
func Get_impliesRecord() gopurs_runtime.Value {
	once_impliesRecord.Do(func() {
		cache_impliesRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_impliesRecord(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_impliesRecord
}

var cache_impliesRecord__gopurs_runtime_Value_497482630 gopurs_runtime.Value
var once_impliesRecord__gopurs_runtime_Value_497482630 sync.Once
func Get_impliesRecord__gopurs_runtime_Value_497482630() gopurs_runtime.Value {
	once_impliesRecord__gopurs_runtime_Value_497482630.Do(func() {
		cache_impliesRecord__gopurs_runtime_Value_497482630 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_impliesRecord__gopurs_runtime_Value_497482630(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_impliesRecord__gopurs_runtime_Value_497482630
}

var cache_implies gopurs_runtime.Value
var once_implies sync.Once
func Get_implies() gopurs_runtime.Value {
	once_implies.Do(func() {
		cache_implies = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_implies(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_implies
}

var cache_implies__gopurs_runtime_Value_3472268504 gopurs_runtime.Value
var once_implies__gopurs_runtime_Value_3472268504 sync.Once
func Get_implies__gopurs_runtime_Value_3472268504() gopurs_runtime.Value {
	once_implies__gopurs_runtime_Value_3472268504.Do(func() {
		cache_implies__gopurs_runtime_Value_3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_implies__gopurs_runtime_Value_3472268504(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_implies__gopurs_runtime_Value_3472268504
}

var cache_heytingAlgebraUnit gopurs_runtime.Value
var once_heytingAlgebraUnit sync.Once
func Get_heytingAlgebraUnit() gopurs_runtime.Value {
	once_heytingAlgebraUnit.Do(func() {
		cache_heytingAlgebraUnit = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), pkg_Data_Unit.Get_unit(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), pkg_Data_Unit.Get_unit()})
	})
	return cache_heytingAlgebraUnit
}

var cache_heytingAlgebraRecordNil gopurs_runtime.Value
var once_heytingAlgebraRecordNil sync.Once
func Get_heytingAlgebraRecordNil() gopurs_runtime.Value {
	once_heytingAlgebraRecordNil.Do(func() {
		cache_heytingAlgebraRecordNil = gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "ffRecord", "impliesRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_heytingAlgebraRecordNil
}

var cache_heytingAlgebraRecordNil__gopurs_runtime_Value_28213168 gopurs_runtime.Value
var once_heytingAlgebraRecordNil__gopurs_runtime_Value_28213168 sync.Once
func Get_heytingAlgebraRecordNil__gopurs_runtime_Value_28213168() gopurs_runtime.Value {
	once_heytingAlgebraRecordNil__gopurs_runtime_Value_28213168.Do(func() {
		cache_heytingAlgebraRecordNil__gopurs_runtime_Value_28213168 = gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "ffRecord", "impliesRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
	return cache_heytingAlgebraRecordNil__gopurs_runtime_Value_28213168
}

var cache_heytingAlgebraProxy gopurs_runtime.Value
var once_heytingAlgebraProxy sync.Once
func Get_heytingAlgebraProxy() gopurs_runtime.Value {
	once_heytingAlgebraProxy.Do(func() {
		cache_heytingAlgebraProxy = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}})
	})
	return cache_heytingAlgebraProxy
}

var cache_heytingAlgebraProxy__gopurs_runtime_Value_549257500 gopurs_runtime.Value
var once_heytingAlgebraProxy__gopurs_runtime_Value_549257500 sync.Once
func Get_heytingAlgebraProxy__gopurs_runtime_Value_549257500() gopurs_runtime.Value {
	once_heytingAlgebraProxy__gopurs_runtime_Value_549257500.Do(func() {
		cache_heytingAlgebraProxy__gopurs_runtime_Value_549257500 = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}})
	})
	return cache_heytingAlgebraProxy__gopurs_runtime_Value_549257500
}

var cache_ffRecord gopurs_runtime.Value
var once_ffRecord sync.Once
func Get_ffRecord() gopurs_runtime.Value {
	once_ffRecord.Do(func() {
		cache_ffRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ffRecord(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_ffRecord
}

var cache_ffRecord__gopurs_runtime_Value_465956064 gopurs_runtime.Value
var once_ffRecord__gopurs_runtime_Value_465956064 sync.Once
func Get_ffRecord__gopurs_runtime_Value_465956064() gopurs_runtime.Value {
	once_ffRecord__gopurs_runtime_Value_465956064.Do(func() {
		cache_ffRecord__gopurs_runtime_Value_465956064 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ffRecord__gopurs_runtime_Value_465956064(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_ffRecord__gopurs_runtime_Value_465956064
}

var cache_ff gopurs_runtime.Value
var once_ff sync.Once
func Get_ff() gopurs_runtime.Value {
	once_ff.Do(func() {
		cache_ff = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ff(dict_0_box)
})
	})
	return cache_ff
}

var cache_ff__gopurs_runtime_Value_2527024921 gopurs_runtime.Value
var once_ff__gopurs_runtime_Value_2527024921 sync.Once
func Get_ff__gopurs_runtime_Value_2527024921() gopurs_runtime.Value {
	once_ff__gopurs_runtime_Value_2527024921.Do(func() {
		cache_ff__gopurs_runtime_Value_2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ff__gopurs_runtime_Value_2527024921(dict_0_box)
})
	})
	return cache_ff__gopurs_runtime_Value_2527024921
}

var cache_disjRecord gopurs_runtime.Value
var once_disjRecord sync.Once
func Get_disjRecord() gopurs_runtime.Value {
	once_disjRecord.Do(func() {
		cache_disjRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disjRecord(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disjRecord
}

var cache_disjRecord__gopurs_runtime_Value_497482630 gopurs_runtime.Value
var once_disjRecord__gopurs_runtime_Value_497482630 sync.Once
func Get_disjRecord__gopurs_runtime_Value_497482630() gopurs_runtime.Value {
	once_disjRecord__gopurs_runtime_Value_497482630.Do(func() {
		cache_disjRecord__gopurs_runtime_Value_497482630 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disjRecord__gopurs_runtime_Value_497482630(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disjRecord__gopurs_runtime_Value_497482630
}

var cache_disj gopurs_runtime.Value
var once_disj sync.Once
func Get_disj() gopurs_runtime.Value {
	once_disj.Do(func() {
		cache_disj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj
}

var cache_disj__gopurs_runtime_Value_3472268504 gopurs_runtime.Value
var once_disj__gopurs_runtime_Value_3472268504 sync.Once
func Get_disj__gopurs_runtime_Value_3472268504() gopurs_runtime.Value {
	once_disj__gopurs_runtime_Value_3472268504.Do(func() {
		cache_disj__gopurs_runtime_Value_3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__gopurs_runtime_Value_3472268504(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__gopurs_runtime_Value_3472268504
}

var cache_heytingAlgebraBoolean gopurs_runtime.Value
var once_heytingAlgebraBoolean sync.Once
func Get_heytingAlgebraBoolean() gopurs_runtime.Value {
	once_heytingAlgebraBoolean.Do(func() {
		cache_heytingAlgebraBoolean = gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{Get_boolConj(), Get_boolDisj(), gopurs_runtime.Bool(false), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_heytingAlgebraBoolean(), "not"), a_0), b_1).IntVal) != (0))
})
}), Get_boolNot(), gopurs_runtime.Bool(true)})
	})
	return cache_heytingAlgebraBoolean
}

var cache_conjRecord gopurs_runtime.Value
var once_conjRecord sync.Once
func Get_conjRecord() gopurs_runtime.Value {
	once_conjRecord.Do(func() {
		cache_conjRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conjRecord(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conjRecord
}

var cache_conjRecord__gopurs_runtime_Value_497482630 gopurs_runtime.Value
var once_conjRecord__gopurs_runtime_Value_497482630 sync.Once
func Get_conjRecord__gopurs_runtime_Value_497482630() gopurs_runtime.Value {
	once_conjRecord__gopurs_runtime_Value_497482630.Do(func() {
		cache_conjRecord__gopurs_runtime_Value_497482630 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conjRecord__gopurs_runtime_Value_497482630(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conjRecord__gopurs_runtime_Value_497482630
}

var cache_heytingAlgebraRecord gopurs_runtime.Value
var once_heytingAlgebraRecord sync.Once
func Get_heytingAlgebraRecord() gopurs_runtime.Value {
	once_heytingAlgebraRecord.Do(func() {
		cache_heytingAlgebraRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictHeytingAlgebraRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heytingAlgebraRecord(_dollar__unused_0_box, dictHeytingAlgebraRecord_1_box)
})
	})
	return cache_heytingAlgebraRecord
}

var cache_conj gopurs_runtime.Value
var once_conj sync.Once
func Get_conj() gopurs_runtime.Value {
	once_conj.Do(func() {
		cache_conj = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj
}

var cache_conj__gopurs_runtime_Value_3472268504 gopurs_runtime.Value
var once_conj__gopurs_runtime_Value_3472268504 sync.Once
func Get_conj__gopurs_runtime_Value_3472268504() gopurs_runtime.Value {
	once_conj__gopurs_runtime_Value_3472268504.Do(func() {
		cache_conj__gopurs_runtime_Value_3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__gopurs_runtime_Value_3472268504(gopurs_runtime.CoerceToStruct[Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__gopurs_runtime_Value_3472268504
}

var cache_heytingAlgebraFunction gopurs_runtime.Value
var once_heytingAlgebraFunction sync.Once
func Get_heytingAlgebraFunction() gopurs_runtime.Value {
	once_heytingAlgebraFunction.Do(func() {
		cache_heytingAlgebraFunction = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heytingAlgebraFunction(dictHeytingAlgebra_0_box)
})
	})
	return cache_heytingAlgebraFunction
}

var cache_heytingAlgebraRecordCons gopurs_runtime.Value
var once_heytingAlgebraRecordCons sync.Once
func Get_heytingAlgebraRecordCons() gopurs_runtime.Value {
	once_heytingAlgebraRecordCons.Do(func() {
		cache_heytingAlgebraRecordCons = gopurs_runtime.Func4(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictHeytingAlgebraRecord_2_box gopurs_runtime.Value, dictHeytingAlgebra_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heytingAlgebraRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictHeytingAlgebraRecord_2_box, dictHeytingAlgebra_3_box)
})
	})
	return cache_heytingAlgebraRecordCons
}

type Constructor_HeytingAlgebraRecord[T_rowlist any, T_row any, T_subrow any] struct {
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
		c := (*Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "conjRecord": return c.V0
		case "disjRecord": return c.V1
		case "ffRecord": return c.V2
		case "impliesRecord": return c.V3
		case "notRecord": return c.V4
		case "ttRecord": return c.V5
		default: panic("Key not found in dictionary Constructor_HeytingAlgebraRecord: " + key)
		}
	}
}


type Constructor_HeytingAlgebra[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 T_a
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 T_a
}


func init() {
	gopurs_runtime.StructGetters[926771738] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_HeytingAlgebra[gopurs_runtime.Value])(ptr)
		switch key {
		case "conj": return c.V0
		case "disj": return c.V1
		case "ff": return c.V2
		case "implies": return c.V3
		case "not": return c.V4
		case "tt": return c.V5
		default: panic("Key not found in dictionary Constructor_HeytingAlgebra: " + key)
		}
	}
}


func Call_ttRecord(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_ttRecord__gopurs_runtime_Value_465956064(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_tt(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}

func Call_tt__gopurs_runtime_Value_2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}

func Call_notRecord(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_notRecord__gopurs_runtime_Value_726562039(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_not(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_not__gopurs_runtime_Value_1505204753(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_impliesRecord(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_impliesRecord__gopurs_runtime_Value_497482630(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_implies(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_implies__gopurs_runtime_Value_3472268504(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_ffRecord(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_ffRecord__gopurs_runtime_Value_465956064(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_ff(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}

func Call_ff__gopurs_runtime_Value_2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}

func Call_disjRecord(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_disjRecord__gopurs_runtime_Value_497482630(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_disj(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_disj__gopurs_runtime_Value_3472268504(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_conjRecord(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conjRecord__gopurs_runtime_Value_497482630(dict_0_loop *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebraRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_heytingAlgebraRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictHeytingAlgebraRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictHeytingAlgebraRecord_1 gopurs_runtime.Value = dictHeytingAlgebraRecord_1_loop
_ = dictHeytingAlgebraRecord_1
return gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "conjRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "disjRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "ffRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "impliesRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "notRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_1, "ttRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)})})
}

func Call_conj(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_conj__gopurs_runtime_Value_3472268504(dict_0_loop *Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_heytingAlgebraFunction(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.RecordDict([]string{"conjRecord", "disjRecord", "ffRecord", "impliesRecord", "notRecord", "ttRecord"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}).StrVal()
_ = key_9_2
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Str(key_9_2))
_ = get_10_3
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Str(key_9_2), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "conj"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "conjRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}).StrVal()
_ = key_9_4
get_10_5 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Str(key_9_4))
_ = get_10_5
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Str(key_9_4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "disj"), gopurs_runtime.Apply(get_10_5, ra_7), gopurs_runtime.Apply(get_10_5, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "disjRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), ff1_4_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "ffRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, row_7))
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}).StrVal()
_ = key_9_6
get_10_7 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Str(key_9_6))
_ = get_10_7
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Str(key_9_6), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "implies"), gopurs_runtime.Apply(get_10_7, ra_7), gopurs_runtime.Apply(get_10_7, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "impliesRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, ra_7, rb_8))
})
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
key_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}).StrVal()
_ = key_8_8
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Str(key_8_8), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_3, "not"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Str(key_8_8), row_7)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "notRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, row_7))
})
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(row_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}), tt1_5_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictHeytingAlgebraRecord_2, "ttRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(nil)}, row_7))
})
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

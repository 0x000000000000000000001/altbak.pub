package Data_Bounded

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_topRecord gopurs_runtime.Value
var once_topRecord sync.Once
func Get_topRecord() gopurs_runtime.Value {
	once_topRecord.Do(func() {
		cache_topRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_topRecord(dict_0_box)
})
	})
	return cache_topRecord
}

var cache_topRecord__gopurs_runtime_Value_2610646464 gopurs_runtime.Value
var once_topRecord__gopurs_runtime_Value_2610646464 sync.Once
func Get_topRecord__gopurs_runtime_Value_2610646464() gopurs_runtime.Value {
	once_topRecord__gopurs_runtime_Value_2610646464.Do(func() {
		cache_topRecord__gopurs_runtime_Value_2610646464 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_topRecord__gopurs_runtime_Value_2610646464(dict_0_box)
})
	})
	return cache_topRecord__gopurs_runtime_Value_2610646464
}

var cache_top gopurs_runtime.Value
var once_top sync.Once
func Get_top() gopurs_runtime.Value {
	once_top.Do(func() {
		cache_top = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top(dict_0_box)
})
	})
	return cache_top
}

var cache_top__gopurs_runtime_Value_338427193 gopurs_runtime.Value
var once_top__gopurs_runtime_Value_338427193 sync.Once
func Get_top__gopurs_runtime_Value_338427193() gopurs_runtime.Value {
	once_top__gopurs_runtime_Value_338427193.Do(func() {
		cache_top__gopurs_runtime_Value_338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top__gopurs_runtime_Value_338427193(dict_0_box)
})
	})
	return cache_top__gopurs_runtime_Value_338427193
}

var cache_boundedUnit gopurs_runtime.Value
var once_boundedUnit sync.Once
func Get_boundedUnit() gopurs_runtime.Value {
	once_boundedUnit.Do(func() {
		cache_boundedUnit = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordUnit()
}), pkg_Data_Unit.Get_unit(), pkg_Data_Unit.Get_unit())
	})
	return cache_boundedUnit
}

var cache_boundedRecordNil gopurs_runtime.Value
var once_boundedRecordNil sync.Once
func Get_boundedRecordNil() gopurs_runtime.Value {
	once_boundedRecordNil.Do(func() {
		cache_boundedRecordNil = gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordRecordNil()
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
	return cache_boundedRecordNil
}

var cache_boundedProxy gopurs_runtime.Value
var once_boundedProxy sync.Once
func Get_boundedProxy() gopurs_runtime.Value {
	once_boundedProxy.Do(func() {
		cache_boundedProxy = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordProxy()
}), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
	})
	return cache_boundedProxy
}

var cache_boundedOrdering gopurs_runtime.Value
var once_boundedOrdering sync.Once
func Get_boundedOrdering() gopurs_runtime.Value {
	once_boundedOrdering.Do(func() {
		cache_boundedOrdering = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordOrdering()
}), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
	})
	return cache_boundedOrdering
}

var cache_boundedNumber gopurs_runtime.Value
var once_boundedNumber sync.Once
func Get_boundedNumber() gopurs_runtime.Value {
	once_boundedNumber.Do(func() {
		cache_boundedNumber = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordNumber()
}), gopurs_runtime.Float(Get_bottomNumber().FloatVal()), gopurs_runtime.Float(Get_topNumber().FloatVal()))
	})
	return cache_boundedNumber
}

var cache_boundedInt gopurs_runtime.Value
var once_boundedInt sync.Once
func Get_boundedInt() gopurs_runtime.Value {
	once_boundedInt.Do(func() {
		cache_boundedInt = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}))
}), gopurs_runtime.Int(Get_bottomInt().IntVal), gopurs_runtime.Int(Get_topInt().IntVal))
	})
	return cache_boundedInt
}

var cache_boundedChar gopurs_runtime.Value
var once_boundedChar sync.Once
func Get_boundedChar() gopurs_runtime.Value {
	once_boundedChar.Do(func() {
		cache_boundedChar = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordChar()
}), gopurs_runtime.Str(Get_bottomChar().StrVal()), gopurs_runtime.Str(Get_topChar().StrVal()))
	})
	return cache_boundedChar
}

var cache_boundedBoolean gopurs_runtime.Value
var once_boundedBoolean sync.Once
func Get_boundedBoolean() gopurs_runtime.Value {
	once_boundedBoolean.Do(func() {
		cache_boundedBoolean = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordBoolean()
}), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true))
	})
	return cache_boundedBoolean
}

var cache_bottomRecord gopurs_runtime.Value
var once_bottomRecord sync.Once
func Get_bottomRecord() gopurs_runtime.Value {
	once_bottomRecord.Do(func() {
		cache_bottomRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottomRecord(dict_0_box)
})
	})
	return cache_bottomRecord
}

var cache_bottomRecord__gopurs_runtime_Value_2610646464 gopurs_runtime.Value
var once_bottomRecord__gopurs_runtime_Value_2610646464 sync.Once
func Get_bottomRecord__gopurs_runtime_Value_2610646464() gopurs_runtime.Value {
	once_bottomRecord__gopurs_runtime_Value_2610646464.Do(func() {
		cache_bottomRecord__gopurs_runtime_Value_2610646464 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottomRecord__gopurs_runtime_Value_2610646464(dict_0_box)
})
	})
	return cache_bottomRecord__gopurs_runtime_Value_2610646464
}

var cache_boundedRecord gopurs_runtime.Value
var once_boundedRecord sync.Once
func Get_boundedRecord() gopurs_runtime.Value {
	once_boundedRecord.Do(func() {
		cache_boundedRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictBoundedRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedRecord(_dollar__unused_0_box, dictBoundedRecord_1_box)
})
	})
	return cache_boundedRecord
}

var cache_bottom gopurs_runtime.Value
var once_bottom sync.Once
func Get_bottom() gopurs_runtime.Value {
	once_bottom.Do(func() {
		cache_bottom = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom(dict_0_box)
})
	})
	return cache_bottom
}

var cache_bottom__gopurs_runtime_Value_338427193 gopurs_runtime.Value
var once_bottom__gopurs_runtime_Value_338427193 sync.Once
func Get_bottom__gopurs_runtime_Value_338427193() gopurs_runtime.Value {
	once_bottom__gopurs_runtime_Value_338427193.Do(func() {
		cache_bottom__gopurs_runtime_Value_338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom__gopurs_runtime_Value_338427193(dict_0_box)
})
	})
	return cache_bottom__gopurs_runtime_Value_338427193
}

var cache_boundedRecordCons gopurs_runtime.Value
var once_boundedRecordCons sync.Once
func Get_boundedRecordCons() gopurs_runtime.Value {
	once_boundedRecordCons.Do(func() {
		cache_boundedRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictBounded_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedRecordCons(dictIsSymbol_0_box, dictBounded_1_box)
})
	})
	return cache_boundedRecordCons
}

func Call_topRecord(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "topRecord")
}

func Call_topRecord__gopurs_runtime_Value_2610646464(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "topRecord")
}

func Call_top(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_top__gopurs_runtime_Value_338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_bottomRecord(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottomRecord")
}

func Call_bottomRecord__gopurs_runtime_Value_2610646464(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottomRecord")
}

func Call_boundedRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictBoundedRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictBoundedRecord_1 gopurs_runtime.Value = dictBoundedRecord_1_loop
_ = dictBoundedRecord_1
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_1, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
eqRec1_3_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = eqRec1_3_2
ordRecord1_2_0 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_3_2
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = ordRecord1_2_0
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecord1_2_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
}

func Call_bottom(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_bottom__gopurs_runtime_Value_338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_boundedRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictBounded_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictBounded_1 gopurs_runtime.Value = dictBounded_1_loop
_ = dictBounded_1
top1_2_0 := gopurs_runtime.RecordGet(dictBounded_1, "top")
_ = top1_2_0
bottom1_3_1 := gopurs_runtime.RecordGet(dictBounded_1, "bottom")
_ = bottom1_3_1
Ord0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_1, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_4_2
return gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBoundedRecord_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_7, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_8_4
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_4, "EqRecord0"), gopurs_runtime.Value{})
_ = __local_var_9_6
__local_var_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_4_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_10_7
eqRowCons2_9_5 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
get_14_8 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = get_14_8
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_7, "eq"), gopurs_runtime.Apply(get_14_8, ra_12), gopurs_runtime.Apply(get_14_8, rb_13)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_6, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, ra_12, rb_13)).IntVal) != (0))
})
})
}))
_ = eqRowCons2_9_5
ordRecordCons_8_3 := gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_9_5
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
key_13_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
_ = key_13_9
left_14_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_4_2, "compare"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_13_9, ra_11), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_13_9, rb_12))
_ = left_14_10
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), left_14_10, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Bool(false)).IntVal) != (0) {
__t11 = left_14_10
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_4, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, ra_11, rb_12)
}
end_branch_11:
return __t11
})
})
}))
_ = ordRecordCons_8_3
return gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecordCons_8_3
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), bottom1_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, rowProxy_10))
})
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), top1_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, rowProxy_10))
})
}))
})
})
})
}

func Get_bottomChar() gopurs_runtime.Value {
	return _Gopurs_BottomChar
}

func Get_bottomInt() gopurs_runtime.Value {
	return _Gopurs_BottomInt
}

func Get_bottomNumber() gopurs_runtime.Value {
	return _Gopurs_BottomNumber
}

func Get_topChar() gopurs_runtime.Value {
	return _Gopurs_TopChar
}

func Get_topInt() gopurs_runtime.Value {
	return _Gopurs_TopInt
}

func Get_topNumber() gopurs_runtime.Value {
	return _Gopurs_TopNumber
}

package Data_Bounded

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
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
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
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
}), Get_bottomNumber(), Get_topNumber())
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
}), Get_bottomInt(), Get_topInt())
	})
	return cache_boundedInt
}

var cache_boundedChar gopurs_runtime.Value
var once_boundedChar sync.Once
func Get_boundedChar() gopurs_runtime.Value {
	once_boundedChar.Do(func() {
		cache_boundedChar = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Ord.Get_ordChar()
}), Get_bottomChar(), Get_topChar())
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
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V1
}

func Call_top(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V1
}

func Call_bottomRecord(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V0
}

func Call_boundedRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictBoundedRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictBoundedRecord_1 gopurs_runtime.Value = dictBoundedRecord_1_loop
_ = dictBoundedRecord_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_1, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_0
eqRec1_3_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = eqRec1_3_1
ordRecord1_4_2 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_3_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = ordRecord1_4_2
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecord1_4_2
}), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData2)(dictBoundedRecord_1.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData2)(dictBoundedRecord_1.UnsafePtr)).V1, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
}

func Call_bottom(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V0
}

func Call_boundedRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictBounded_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictBounded_1 gopurs_runtime.Value = dictBounded_1_loop
_ = dictBounded_1
top1_2_0 := ((*gopurs_runtime.RecordData2)(dictBounded_1.UnsafePtr)).V1
_ = top1_2_0
bottom1_3_1 := ((*gopurs_runtime.RecordData2)(dictBounded_1.UnsafePtr)).V0
_ = bottom1_3_1
Ord0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_1, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_4_2
return gopurs_runtime.Func3(func(_dollar__unused_5 gopurs_runtime.Value, _dollar__unused_6 gopurs_runtime.Value, dictBoundedRecord_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_7, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_8_3
__local_var_9_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_3, "EqRecord0"), gopurs_runtime.Value{})
_ = __local_var_9_4
__local_var_10_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_4_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_10_5
eqRowCons2_11_7 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func3(func(v_11 gopurs_runtime.Value, ra_12 gopurs_runtime.Value, rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
get_14_8 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictIsSymbol_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}))
_ = get_14_8
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_5, "eq"), gopurs_runtime.Apply(get_14_8, ra_12), gopurs_runtime.Apply(get_14_8, rb_13)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_4, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, ra_12, rb_13))
}))
_ = eqRowCons2_11_7
ordRecordCons_11_6 := gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_11_7
}), gopurs_runtime.Func3(func(v_12 gopurs_runtime.Value, ra_13 gopurs_runtime.Value, rb_14 gopurs_runtime.Value) gopurs_runtime.Value {
key_15_9 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictIsSymbol_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
_ = key_15_9
left_16_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_4_2, "compare"), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_15_9, ra_13), gopurs_runtime.Apply2(pkg_Record_Unsafe.Get_unsafeGet(), key_15_9, rb_14))
_ = left_16_10
var __t11 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq"), left_16_10, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}), gopurs_runtime.Bool(false)).IntVal) != (0) {
__t11 = left_16_10
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_3, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, ra_13, rb_14)
}
end_branch_11:
return __t11
}))
_ = ordRecordCons_11_6
return gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecordCons_11_6
}), gopurs_runtime.Func2(func(v_12 gopurs_runtime.Value, rowProxy_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictIsSymbol_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), bottom1_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, rowProxy_13))
}), gopurs_runtime.Func2(func(v_12 gopurs_runtime.Value, rowProxy_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictIsSymbol_0.UnsafePtr)).V0, gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), top1_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}, rowProxy_13))
}))
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

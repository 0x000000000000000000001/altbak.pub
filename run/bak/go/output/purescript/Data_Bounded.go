package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Bounded_ordRecord gopurs_runtime.Value
var once_Data_Bounded_ordRecord sync.Once
func Get_Data_Bounded_ordRecord() gopurs_runtime.Value {
	once_Data_Bounded_ordRecord.Do(func() {
		cache_Data_Bounded_ordRecord = gopurs_runtime.Func(func(dictOrdRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_ordRecord(dictOrdRecord_0_box)
})
	})
	return cache_Data_Bounded_ordRecord
}

var cache_Data_Bounded_BoundedRecord_dollarDict gopurs_runtime.Value
var once_Data_Bounded_BoundedRecord_dollarDict sync.Once
func Get_Data_Bounded_BoundedRecord_dollarDict() gopurs_runtime.Value {
	once_Data_Bounded_BoundedRecord_dollarDict.Do(func() {
		cache_Data_Bounded_BoundedRecord_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_BoundedRecord_dollarDict(x_0_box)
})
	})
	return cache_Data_Bounded_BoundedRecord_dollarDict
}

var cache_Data_Bounded_Bounded_dollarDict gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollarDict sync.Once
func Get_Data_Bounded_Bounded_dollarDict() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollarDict.Do(func() {
		cache_Data_Bounded_Bounded_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_Bounded_dollarDict(x_0_box)
})
	})
	return cache_Data_Bounded_Bounded_dollarDict
}

var cache_Data_Bounded_topRecord gopurs_runtime.Value
var once_Data_Bounded_topRecord sync.Once
func Get_Data_Bounded_topRecord() gopurs_runtime.Value {
	once_Data_Bounded_topRecord.Do(func() {
		cache_Data_Bounded_topRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_topRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord](dict_0_box))
})
	})
	return cache_Data_Bounded_topRecord
}

var cache_Data_Bounded_top gopurs_runtime.Value
var once_Data_Bounded_top sync.Once
func Get_Data_Bounded_top() gopurs_runtime.Value {
	once_Data_Bounded_top.Do(func() {
		cache_Data_Bounded_top = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_top(dict_0_box)
})
	})
	return cache_Data_Bounded_top
}

var cache_Data_Bounded_boundedUnit gopurs_runtime.Value
var once_Data_Bounded_boundedUnit sync.Once
func Get_Data_Bounded_boundedUnit() gopurs_runtime.Value {
	once_Data_Bounded_boundedUnit.Do(func() {
		cache_Data_Bounded_boundedUnit = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordUnit()
}), Get_Data_Unit_unit(), Get_Data_Unit_unit())
	})
	return cache_Data_Bounded_boundedUnit
}

var cache_Data_Bounded_boundedRecordNil gopurs_runtime.Value
var once_Data_Bounded_boundedRecordNil sync.Once
func Get_Data_Bounded_boundedRecordNil() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecordNil.Do(func() {
		cache_Data_Bounded_boundedRecordNil = gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordRecordNil()
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
	return cache_Data_Bounded_boundedRecordNil
}

var cache_Data_Bounded_boundedProxy gopurs_runtime.Value
var once_Data_Bounded_boundedProxy sync.Once
func Get_Data_Bounded_boundedProxy() gopurs_runtime.Value {
	once_Data_Bounded_boundedProxy.Do(func() {
		cache_Data_Bounded_boundedProxy = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordProxy()
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
	})
	return cache_Data_Bounded_boundedProxy
}

var cache_Data_Bounded_boundedOrdering gopurs_runtime.Value
var once_Data_Bounded_boundedOrdering sync.Once
func Get_Data_Bounded_boundedOrdering() gopurs_runtime.Value {
	once_Data_Bounded_boundedOrdering.Do(func() {
		cache_Data_Bounded_boundedOrdering = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordOrdering()
}), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})
	})
	return cache_Data_Bounded_boundedOrdering
}

var cache_Data_Bounded_boundedNumber gopurs_runtime.Value
var once_Data_Bounded_boundedNumber sync.Once
func Get_Data_Bounded_boundedNumber() gopurs_runtime.Value {
	once_Data_Bounded_boundedNumber.Do(func() {
		cache_Data_Bounded_boundedNumber = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordNumber()
}), gopurs_runtime.Float(Get_Data_Bounded_bottomNumber().FloatVal()), gopurs_runtime.Float(Get_Data_Bounded_topNumber().FloatVal()))
	})
	return cache_Data_Bounded_boundedNumber
}

var cache_Data_Bounded_boundedInt gopurs_runtime.Value
var once_Data_Bounded_boundedInt sync.Once
func Get_Data_Bounded_boundedInt() gopurs_runtime.Value {
	once_Data_Bounded_boundedInt.Do(func() {
		cache_Data_Bounded_boundedInt = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordInt()
}), gopurs_runtime.Int(Get_Data_Bounded_bottomInt().IntVal), gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal))
	})
	return cache_Data_Bounded_boundedInt
}

var cache_Data_Bounded_boundedChar gopurs_runtime.Value
var once_Data_Bounded_boundedChar sync.Once
func Get_Data_Bounded_boundedChar() gopurs_runtime.Value {
	once_Data_Bounded_boundedChar.Do(func() {
		cache_Data_Bounded_boundedChar = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordChar()
}), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal()), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal()))
	})
	return cache_Data_Bounded_boundedChar
}

var cache_Data_Bounded_boundedBoolean gopurs_runtime.Value
var once_Data_Bounded_boundedBoolean sync.Once
func Get_Data_Bounded_boundedBoolean() gopurs_runtime.Value {
	once_Data_Bounded_boundedBoolean.Do(func() {
		cache_Data_Bounded_boundedBoolean = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Ord_ordBoolean()
}), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true))
	})
	return cache_Data_Bounded_boundedBoolean
}

var cache_Data_Bounded_bottomRecord gopurs_runtime.Value
var once_Data_Bounded_bottomRecord sync.Once
func Get_Data_Bounded_bottomRecord() gopurs_runtime.Value {
	once_Data_Bounded_bottomRecord.Do(func() {
		cache_Data_Bounded_bottomRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_bottomRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord](dict_0_box))
})
	})
	return cache_Data_Bounded_bottomRecord
}

var cache_Data_Bounded_boundedRecord gopurs_runtime.Value
var once_Data_Bounded_boundedRecord sync.Once
func Get_Data_Bounded_boundedRecord() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecord.Do(func() {
		cache_Data_Bounded_boundedRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictBoundedRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_boundedRecord(_dollar__unused_0_box, dictBoundedRecord_1_box)
})
	})
	return cache_Data_Bounded_boundedRecord
}

var cache_Data_Bounded_bottom gopurs_runtime.Value
var once_Data_Bounded_bottom sync.Once
func Get_Data_Bounded_bottom() gopurs_runtime.Value {
	once_Data_Bounded_bottom.Do(func() {
		cache_Data_Bounded_bottom = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_bottom(dict_0_box)
})
	})
	return cache_Data_Bounded_bottom
}

var cache_Data_Bounded_boundedRecordCons gopurs_runtime.Value
var once_Data_Bounded_boundedRecordCons sync.Once
func Get_Data_Bounded_boundedRecordCons() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecordCons.Do(func() {
		cache_Data_Bounded_boundedRecordCons = gopurs_runtime.Func2(func(dictIsSymbol_0_box gopurs_runtime.Value, dictBounded_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_boundedRecordCons(dictIsSymbol_0_box, dictBounded_1_box)
})
	})
	return cache_Data_Bounded_boundedRecordCons
}

var cache_Data_Bounded_bottom__338427193 gopurs_runtime.Value
var once_Data_Bounded_bottom__338427193 sync.Once
func Get_Data_Bounded_bottom__338427193() gopurs_runtime.Value {
	once_Data_Bounded_bottom__338427193.Do(func() {
		cache_Data_Bounded_bottom__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_bottom__338427193(dict_0_box)
})
	})
	return cache_Data_Bounded_bottom__338427193
}

var cache_Data_Bounded_bottomRecord__1740832576 gopurs_runtime.Value
var once_Data_Bounded_bottomRecord__1740832576 sync.Once
func Get_Data_Bounded_bottomRecord__1740832576() gopurs_runtime.Value {
	once_Data_Bounded_bottomRecord__1740832576.Do(func() {
		cache_Data_Bounded_bottomRecord__1740832576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_bottomRecord__1740832576(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord](dict_0_box))
})
	})
	return cache_Data_Bounded_bottomRecord__1740832576
}

var cache_Data_Bounded_bottomRecord__2610646464 gopurs_runtime.Value
var once_Data_Bounded_bottomRecord__2610646464 sync.Once
func Get_Data_Bounded_bottomRecord__2610646464() gopurs_runtime.Value {
	once_Data_Bounded_bottomRecord__2610646464.Do(func() {
		cache_Data_Bounded_bottomRecord__2610646464 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_bottomRecord__2610646464(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord](dict_0_box))
})
	})
	return cache_Data_Bounded_bottomRecord__2610646464
}

var cache_Data_Bounded_top__338427193 gopurs_runtime.Value
var once_Data_Bounded_top__338427193 sync.Once
func Get_Data_Bounded_top__338427193() gopurs_runtime.Value {
	once_Data_Bounded_top__338427193.Do(func() {
		cache_Data_Bounded_top__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_top__338427193(dict_0_box)
})
	})
	return cache_Data_Bounded_top__338427193
}

var cache_Data_Bounded_topRecord__1740832576 gopurs_runtime.Value
var once_Data_Bounded_topRecord__1740832576 sync.Once
func Get_Data_Bounded_topRecord__1740832576() gopurs_runtime.Value {
	once_Data_Bounded_topRecord__1740832576.Do(func() {
		cache_Data_Bounded_topRecord__1740832576 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_topRecord__1740832576(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord](dict_0_box))
})
	})
	return cache_Data_Bounded_topRecord__1740832576
}

var cache_Data_Bounded_topRecord__2610646464 gopurs_runtime.Value
var once_Data_Bounded_topRecord__2610646464 sync.Once
func Get_Data_Bounded_topRecord__2610646464() gopurs_runtime.Value {
	once_Data_Bounded_topRecord__2610646464.Do(func() {
		cache_Data_Bounded_topRecord__2610646464 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_topRecord__2610646464(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord](dict_0_box))
})
	})
	return cache_Data_Bounded_topRecord__2610646464
}

type Constructor_Data_Bounded_BoundedRecord struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4260658871] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bounded_BoundedRecord)(ptr)
		_ = c
		switch key {
		case "OrdRecord0": return gopurs_runtime.Box(c.V0)
		case "bottomRecord": return gopurs_runtime.Box(c.V1)
		case "topRecord": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Bounded_BoundedRecord: " + key)
		}
	}
}


type Constructor_Data_Bounded_Bounded struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3510799738] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bounded_Bounded)(ptr)
		_ = c
		switch key {
		case "Ord0": return gopurs_runtime.Box(c.V0)
		case "bottom": return gopurs_runtime.Box(c.V1)
		case "top": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Bounded_Bounded: " + key)
		}
	}
}


func Call_Data_Bounded_ordRecord(dictOrdRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrdRecord_0 gopurs_runtime.Value = dictOrdRecord_0_loop
_ = dictOrdRecord_0
// TAST (Let): eqRec1_1_0 -> gopurs_runtime.Value
eqRec1_1_0 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = eqRec1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_Bounded_BoundedRecord_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bounded_Bounded_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Bounded_topRecord(dict_0_loop *Constructor_Data_Bounded_BoundedRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bounded_top(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_Data_Bounded_bottomRecord(dict_0_loop *Constructor_Data_Bounded_BoundedRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Bounded_boundedRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictBoundedRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictBoundedRecord_1 gopurs_runtime.Value = dictBoundedRecord_1_loop
_ = dictBoundedRecord_1
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_1, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): eqRec1_3_2 -> gopurs_runtime.Value
eqRec1_3_2 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = eqRec1_3_2
// TAST (Let): ordRecord1_2_0 -> gopurs_runtime.Value
ordRecord1_2_0 := gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRec1_3_2
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = ordRecord1_2_0
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecord1_2_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_Data_Bounded_bottom(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_Data_Bounded_boundedRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, dictBounded_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var dictBounded_1 gopurs_runtime.Value = dictBounded_1_loop
_ = dictBounded_1
// TAST (Let): top1_2_0 -> gopurs_runtime.Value
top1_2_0 := gopurs_runtime.RecordGet(dictBounded_1, "top")
_ = top1_2_0
// TAST (Let): bottom1_3_1 -> gopurs_runtime.Value
bottom1_3_1 := gopurs_runtime.RecordGet(dictBounded_1, "bottom")
_ = bottom1_3_1
// TAST (Let): Ord0_4_2 -> gopurs_runtime.Value
Ord0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_1, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_4_2
return gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBoundedRecord_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_4 -> gopurs_runtime.Value
__local_var_8_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_7, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_8_4
// TAST (Let): __local_var_9_6 -> gopurs_runtime.Value
__local_var_9_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_4, "EqRecord0"), gopurs_runtime.Value{})
_ = __local_var_9_6
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_4_2, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_10_7
// TAST (Let): eqRowCons2_9_5 -> gopurs_runtime.Value
eqRowCons2_9_5 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_14_8 -> gopurs_runtime.Value
get_14_8 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()))
_ = get_14_8
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_7, "eq"), gopurs_runtime.Apply(get_14_8, ra_12), gopurs_runtime.Apply(get_14_8, rb_13)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_9_6, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_12, rb_13).IntVal) != (0)))
})
})
}))
_ = eqRowCons2_9_5
// TAST (Let): ordRecordCons_8_3 -> gopurs_runtime.Value
ordRecordCons_8_3 := gopurs_runtime.RecordDict2("EqRecord0", "compareRecord", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return eqRowCons2_9_5
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_13_9 -> gopurs_runtime.Value
key_13_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_13_9
// TAST (Let): left_14_10 -> gopurs_runtime.Value
left_14_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_4_2, "compare"), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_13_9.StrVal()), ra_11), gopurs_runtime.Apply2(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_13_9.StrVal()), rb_12))
_ = left_14_10
var __t15 uint32
{
var __t14 bool
{
var __t_tag_11 uint32 = uint32(left_14_10.IntVal)
if (uint32(__t_tag_11) == 1527465420) {
__t14 = false
goto end_branch_14
} else {

}
}
{
var __t_tag_12 uint32 = uint32(left_14_10.IntVal)
if (uint32(__t_tag_12) == 380165415) {
__t14 = false
goto end_branch_14
} else {

}
}
{
var __t_tag_13 uint32 = uint32(left_14_10.IntVal)
if (uint32(__t_tag_13) == 902936544) {
__t14 = true
goto end_branch_14
} else {

}
}
{
__t14 = false
}
end_branch_14:
if (__t14) != (true) {
__t15 = uint32(left_14_10.IntVal)
goto end_branch_15
} else {

}
}
{
__t15 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_8_4, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_11, rb_12).IntVal)
}
end_branch_15:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t15), UnsafePtr: nil}
})
})
}))
_ = ordRecordCons_8_3
return gopurs_runtime.RecordDict3("OrdRecord0", "bottomRecord", "topRecord", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return ordRecordCons_8_3
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), bottom1_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(rowProxy_10.IntVal)), UnsafePtr: nil}))
})
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), top1_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(rowProxy_10.IntVal)), UnsafePtr: nil}))
})
}))
})
})
})
}

func Call_Data_Bounded_bottom__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_Data_Bounded_bottomRecord__1740832576(dict_0_loop *Constructor_Data_Bounded_BoundedRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Bounded_bottomRecord__2610646464(dict_0_loop *Constructor_Data_Bounded_BoundedRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Bounded_top__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_Data_Bounded_topRecord__1740832576(dict_0_loop *Constructor_Data_Bounded_BoundedRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bounded_topRecord__2610646464(dict_0_loop *Constructor_Data_Bounded_BoundedRecord) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Get_Data_Bounded_bottomChar() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_BottomChar
}

func Get_Data_Bounded_bottomInt() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_BottomInt
}

func Get_Data_Bounded_bottomNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_BottomNumber
}

func Get_Data_Bounded_topChar() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_TopChar
}

func Get_Data_Bounded_topInt() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_TopInt
}

func Get_Data_Bounded_topNumber() gopurs_runtime.Value {
	return _Gopurs_Data_Bounded_TopNumber
}

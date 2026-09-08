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

var cache_Data_Bounded_BoundedRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Bounded_BoundedRecord_dollar_Dict sync.Once
func Get_Data_Bounded_BoundedRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Bounded_BoundedRecord_dollar_Dict.Do(func() {
		cache_Data_Bounded_BoundedRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4260658871, UnsafePtr: unsafe.Pointer(Call_Data_Bounded_BoundedRecord_dollar_Dict(func() struct{
	OrdRecord0 gopurs_runtime.Value
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	OrdRecord0 gopurs_runtime.Value
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}{}
					clone.OrdRecord0 = gopurs_runtime.RecordGet(orig, "OrdRecord0")
					clone.bottomRecord = gopurs_runtime.RecordGet(orig, "bottomRecord")
					clone.topRecord = gopurs_runtime.RecordGet(orig, "topRecord")
					return clone
				}()))}
})
	})
	return cache_Data_Bounded_BoundedRecord_dollar_Dict
}

var cache_Data_Bounded_Bounded_dollar_Dict gopurs_runtime.Value
var once_Data_Bounded_Bounded_dollar_Dict sync.Once
func Get_Data_Bounded_Bounded_dollar_Dict() gopurs_runtime.Value {
	once_Data_Bounded_Bounded_dollar_Dict.Do(func() {
		cache_Data_Bounded_Bounded_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Call_Data_Bounded_Bounded_dollar_Dict(func() struct{
	Ord0 gopurs_runtime.Value
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ord0 gopurs_runtime.Value
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}{}
					clone.Ord0 = gopurs_runtime.RecordGet(orig, "Ord0")
					clone.bottom = gopurs_runtime.RecordGet(orig, "bottom")
					clone.top = gopurs_runtime.RecordGet(orig, "top")
					return clone
				}()))}
})
	})
	return cache_Data_Bounded_Bounded_dollar_Dict
}

var cache_Data_Bounded_topRecord gopurs_runtime.Value
var once_Data_Bounded_topRecord sync.Once
func Get_Data_Bounded_topRecord() gopurs_runtime.Value {
	once_Data_Bounded_topRecord.Do(func() {
		cache_Data_Bounded_topRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_topRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
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

var cache_Data_Bounded_top__235423490 gopurs_runtime.Value
var once_Data_Bounded_top__235423490 sync.Once
func Get_Data_Bounded_top__235423490() gopurs_runtime.Value {
	once_Data_Bounded_top__235423490.Do(func() {
		cache_Data_Bounded_top__235423490 = gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())
	})
	return cache_Data_Bounded_top__235423490
}

var cache_Data_Bounded_top__2089552681 gopurs_runtime.Value
var once_Data_Bounded_top__2089552681 sync.Once
func Get_Data_Bounded_top__2089552681() gopurs_runtime.Value {
	once_Data_Bounded_top__2089552681.Do(func() {
		cache_Data_Bounded_top__2089552681 = gopurs_runtime.Int(Get_Data_Bounded_topInt().IntVal)
	})
	return cache_Data_Bounded_top__2089552681
}

var cache_Data_Bounded_boundedUnit gopurs_runtime.Value
var once_Data_Bounded_boundedUnit sync.Once
func Get_Data_Bounded_boundedUnit() gopurs_runtime.Value {
	once_Data_Bounded_boundedUnit.Do(func() {
		cache_Data_Bounded_boundedUnit = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordUnit()))}
}), Get_Data_Unit_unit(), Get_Data_Unit_unit()}))}
	})
	return cache_Data_Bounded_boundedUnit
}

var cache_Data_Bounded_boundedRecordNil gopurs_runtime.Value
var once_Data_Bounded_boundedRecordNil sync.Once
func Get_Data_Bounded_boundedRecordNil() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecordNil.Do(func() {
		cache_Data_Bounded_boundedRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 4260658871, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Ord_ordRecordNil()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}()
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict([]string{}, []gopurs_runtime.Value{})
				}()
})}))}
	})
	return cache_Data_Bounded_boundedRecordNil
}

var cache_Data_Bounded_boundedProxy gopurs_runtime.Value
var once_Data_Bounded_boundedProxy sync.Once
func Get_Data_Bounded_boundedProxy() gopurs_runtime.Value {
	once_Data_Bounded_boundedProxy.Do(func() {
		cache_Data_Bounded_boundedProxy = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566((&Constructor_Data_Bounded_Bounded[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3730953251_4177771502(Rebox_Data_Bounded_4177771502_3730953251(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordProxy()))))}
}), uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal), uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)})))}
	})
	return cache_Data_Bounded_boundedProxy
}

var cache_Data_Bounded_boundedOrdering gopurs_runtime.Value
var once_Data_Bounded_boundedOrdering sync.Once
func Get_Data_Bounded_boundedOrdering() gopurs_runtime.Value {
	once_Data_Bounded_boundedOrdering.Do(func() {
		cache_Data_Bounded_boundedOrdering = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_832288803_2094947566((&Constructor_Data_Bounded_Bounded[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3730953251_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[uint32]](Get_Data_Ord_ordOrdering())))}
}), 1527465420, 380165415})))}
	})
	return cache_Data_Bounded_boundedOrdering
}

var cache_Data_Bounded_boundedNumber gopurs_runtime.Value
var once_Data_Bounded_boundedNumber sync.Once
func Get_Data_Bounded_boundedNumber() gopurs_runtime.Value {
	once_Data_Bounded_boundedNumber.Do(func() {
		cache_Data_Bounded_boundedNumber = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_4294747894_2094947566((&Constructor_Data_Bounded_Bounded[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3047586294_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[float64]](Get_Data_Ord_ordNumber())))}
}), Get_Data_Bounded_bottomNumber().FloatVal(), Get_Data_Bounded_topNumber().FloatVal()})))}
	})
	return cache_Data_Bounded_boundedNumber
}

var cache_Data_Bounded_boundedInt gopurs_runtime.Value
var once_Data_Bounded_boundedInt sync.Once
func Get_Data_Bounded_boundedInt() gopurs_runtime.Value {
	once_Data_Bounded_boundedInt.Do(func() {
		cache_Data_Bounded_boundedInt = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3764732725_2094947566((&Constructor_Data_Bounded_Bounded[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())))}
}), Get_Data_Bounded_bottomInt().IntVal, Get_Data_Bounded_topInt().IntVal})))}
	})
	return cache_Data_Bounded_boundedInt
}

var cache_Data_Bounded_boundedChar gopurs_runtime.Value
var once_Data_Bounded_boundedChar sync.Once
func Get_Data_Bounded_boundedChar() gopurs_runtime.Value {
	once_Data_Bounded_boundedChar.Do(func() {
		cache_Data_Bounded_boundedChar = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_2420955921_2094947566((&Constructor_Data_Bounded_Bounded[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_2406510097_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[string]](Get_Data_Ord_ordChar())))}
}), Get_Data_Bounded_bottomChar().StrVal(), Get_Data_Bounded_topChar().StrVal()})))}
	})
	return cache_Data_Bounded_boundedChar
}

var cache_Data_Bounded_boundedBoolean gopurs_runtime.Value
var once_Data_Bounded_boundedBoolean sync.Once
func Get_Data_Bounded_boundedBoolean() gopurs_runtime.Value {
	once_Data_Bounded_boundedBoolean.Do(func() {
		cache_Data_Bounded_boundedBoolean = gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_3728870730_2094947566((&Constructor_Data_Bounded_Bounded[bool]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Bounded_219188042_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[bool]](Get_Data_Ord_ordBoolean())))}
}), false, true})))}
	})
	return cache_Data_Bounded_boundedBoolean
}

var cache_Data_Bounded_bottomRecord gopurs_runtime.Value
var once_Data_Bounded_bottomRecord sync.Once
func Get_Data_Bounded_bottomRecord() gopurs_runtime.Value {
	once_Data_Bounded_bottomRecord.Do(func() {
		cache_Data_Bounded_bottomRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_bottomRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Bounded_bottomRecord
}

var cache_Data_Bounded_boundedRecord gopurs_runtime.Value
var once_Data_Bounded_boundedRecord sync.Once
func Get_Data_Bounded_boundedRecord() gopurs_runtime.Value {
	once_Data_Bounded_boundedRecord.Do(func() {
		cache_Data_Bounded_boundedRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictBoundedRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Bounded_boundedRecord(_dollar___unused_0_box, dictBoundedRecord_1_box)
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

var cache_Data_Bounded_bottom__235423490 gopurs_runtime.Value
var once_Data_Bounded_bottom__235423490 sync.Once
func Get_Data_Bounded_bottom__235423490() gopurs_runtime.Value {
	once_Data_Bounded_bottom__235423490.Do(func() {
		cache_Data_Bounded_bottom__235423490 = gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())
	})
	return cache_Data_Bounded_bottom__235423490
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

type Constructor_Data_Bounded_BoundedRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4260658871] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bounded_BoundedRecord[any, any, any])(ptr)
		_ = c
		switch key {
		case "OrdRecord0": return gopurs_runtime.Box(c.V0)
		case "bottomRecord": return gopurs_runtime.Box(c.V1)
		case "topRecord": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Bounded_BoundedRecord: " + key)
		}
	}
}


type Constructor_Data_Bounded_Bounded[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
	V2 T_a
}


func init() {
	gopurs_runtime.StructGetters[3510799738] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Bounded_Bounded[any])(ptr)
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
// TAST (Let): eqRec1_1_0 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Record (Row [] (TypeVar row)))])
eqRec1_1_0 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = eqRec1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqRec1_1_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrdRecord_0, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Bounded_BoundedRecord_dollar_Dict(x_0_loop struct{
	OrdRecord0 gopurs_runtime.Value
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
}) *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	OrdRecord0 gopurs_runtime.Value
	bottomRecord gopurs_runtime.Value
	topRecord gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"OrdRecord0", "bottomRecord", "topRecord"}, []gopurs_runtime.Value{orig.OrdRecord0, orig.bottomRecord, orig.topRecord})
				}())
}

func Call_Data_Bounded_Bounded_dollar_Dict(x_0_loop struct{
	Ord0 gopurs_runtime.Value
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
}) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
var x_0 struct{
	Ord0 gopurs_runtime.Value
	bottom gopurs_runtime.Value
	top gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Ord0", "bottom", "top"}, []gopurs_runtime.Value{orig.Ord0, orig.bottom, orig.top})
				}())
}

func Call_Data_Bounded_topRecord(dict_0_loop *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Bounded_top(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_Data_Bounded_bottomRecord(dict_0_loop *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Bounded_boundedRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictBoundedRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictBoundedRecord_1 gopurs_runtime.Value = dictBoundedRecord_1_loop
_ = dictBoundedRecord_1
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_1, "OrdRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): eqRec1_3_2 shape=LitRecord bindingType=(ADT ["Data","Eq","Eq"] [(Record (Row [] (TypeVar row)))])
eqRec1_3_2 := (&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "EqRecord0"), gopurs_runtime.Value{}), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = eqRec1_3_2
// TAST (Let): ordRecord1_2_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Ord","Ord"] [(Record (Row [] (TypeVar row)))])
ordRecord1_2_0 := (&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqRec1_3_2)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "compareRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = ordRecord1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordRecord1_2_0)}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_1, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
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
// TAST (Let): top1_2_0 shape=Other bindingType=(TypeVar focus)
top1_2_0 := gopurs_runtime.RecordGet(dictBounded_1, "top")
_ = top1_2_0
// TAST (Let): bottom1_3_1 shape=Other bindingType=(TypeVar focus)
bottom1_3_1 := gopurs_runtime.RecordGet(dictBounded_1, "bottom")
_ = bottom1_3_1
// TAST (Let): Ord0_4_2 shape=App(Other) bindingType=Any
Ord0_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_1, "Ord0"), gopurs_runtime.Value{})
_ = Ord0_4_2
return gopurs_runtime.Func3(func(_dollar___unused_5 gopurs_runtime.Value, _dollar___unused_6 gopurs_runtime.Value, dictBoundedRecord_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ordRecordCons_8_3 shape=App(Var) bindingType=(TypeApp (ADT ["Data","Ord","OrdRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)]), (TypeVar row)])
ordRecordCons_8_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_OrdRecord[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply4(Get_Data_Ord_ordRecordCons(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedRecord_7, "OrdRecord0"), gopurs_runtime.Value{}), gopurs_runtime.Value{}, dictIsSymbol_0, Ord0_4_2))
_ = ordRecordCons_8_3
return gopurs_runtime.Value{Type: 9, IntVal: 4260658871, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_BoundedRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4162894775, UnsafePtr: unsafe.Pointer(ordRecordCons_8_3)}
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), bottom1_3_1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "bottomRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(rowProxy_10.IntVal)), UnsafePtr: nil}))
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, rowProxy_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), top1_2_0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBoundedRecord_7, "topRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(rowProxy_10.IntVal)), UnsafePtr: nil}))
})}))}
})
}

func Rebox_Data_Bounded_219188042_4177771502(in *Constructor_Data_Ord_Ord[bool]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Bounded_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Bounded_2420955921_2094947566(in *Constructor_Data_Bounded_Bounded[string]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Str(in.V1)
		out.V2 = gopurs_runtime.Str(in.V2)
	return out
}

func Rebox_Data_Bounded_3047586294_4177771502(in *Constructor_Data_Ord_Ord[float64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Bounded_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Bounded_3728870730_2094947566(in *Constructor_Data_Bounded_Bounded[bool]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Bool(in.V1)
		out.V2 = gopurs_runtime.Bool(in.V2)
	return out
}

func Rebox_Data_Bounded_3730953251_4177771502(in *Constructor_Data_Ord_Ord[uint32]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Bounded_3764732725_2094947566(in *Constructor_Data_Bounded_Bounded[int64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
		out.V2 = gopurs_runtime.Int(in.V2)
	return out
}

func Rebox_Data_Bounded_4177771502_3730953251(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Bounded_4294747894_2094947566(in *Constructor_Data_Bounded_Bounded[float64]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Float(in.V1)
		out.V2 = gopurs_runtime.Float(in.V2)
	return out
}

func Rebox_Data_Bounded_832288803_2094947566(in *Constructor_Data_Bounded_Bounded[uint32]) *Constructor_Data_Bounded_Bounded[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V1), UnsafePtr: nil}
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
	return out
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

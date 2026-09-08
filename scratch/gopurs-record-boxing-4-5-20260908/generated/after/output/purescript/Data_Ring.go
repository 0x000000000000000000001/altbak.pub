package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Ring_semiringRecord gopurs_runtime.Value
var once_Data_Ring_semiringRecord sync.Once
func Get_Data_Ring_semiringRecord() gopurs_runtime.Value {
	once_Data_Ring_semiringRecord.Do(func() {
		cache_Data_Ring_semiringRecord = gopurs_runtime.Func(func(dictSemiringRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_semiringRecord(dictSemiringRecord_0_box)
})
	})
	return cache_Data_Ring_semiringRecord
}

var cache_Data_Ring_RingRecord_dollar_Dict gopurs_runtime.Value
var once_Data_Ring_RingRecord_dollar_Dict sync.Once
func Get_Data_Ring_RingRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ring_RingRecord_dollar_Dict.Do(func() {
		cache_Data_Ring_RingRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer(Call_Data_Ring_RingRecord_dollar_Dict(func() struct{
	SemiringRecord0 gopurs_runtime.Value
	subRecord gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	SemiringRecord0 gopurs_runtime.Value
	subRecord gopurs_runtime.Value
}{}
					clone.SemiringRecord0 = gopurs_runtime.RecordGet(orig, "SemiringRecord0")
					clone.subRecord = gopurs_runtime.RecordGet(orig, "subRecord")
					return clone
				}()))}
})
	})
	return cache_Data_Ring_RingRecord_dollar_Dict
}

var cache_Data_Ring_Ring_dollar_Dict gopurs_runtime.Value
var once_Data_Ring_Ring_dollar_Dict sync.Once
func Get_Data_Ring_Ring_dollar_Dict() gopurs_runtime.Value {
	once_Data_Ring_Ring_dollar_Dict.Do(func() {
		cache_Data_Ring_Ring_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Call_Data_Ring_Ring_dollar_Dict(func() struct{
	Semiring0 gopurs_runtime.Value
	sub gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Semiring0 gopurs_runtime.Value
	sub gopurs_runtime.Value
}{}
					clone.Semiring0 = gopurs_runtime.RecordGet(orig, "Semiring0")
					clone.sub = gopurs_runtime.RecordGet(orig, "sub")
					return clone
				}()))}
})
	})
	return cache_Data_Ring_Ring_dollar_Dict
}

var cache_Data_Ring_subRecord gopurs_runtime.Value
var once_Data_Ring_subRecord sync.Once
func Get_Data_Ring_subRecord() gopurs_runtime.Value {
	once_Data_Ring_subRecord.Do(func() {
		cache_Data_Ring_subRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_subRecord(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Ring_subRecord
}

var cache_Data_Ring_sub gopurs_runtime.Value
var once_Data_Ring_sub sync.Once
func Get_Data_Ring_sub() gopurs_runtime.Value {
	once_Data_Ring_sub.Do(func() {
		cache_Data_Ring_sub = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Ring_sub
}

var cache_Data_Ring_sub__3052583336 gopurs_runtime.Value
var once_Data_Ring_sub__3052583336 sync.Once
func Get_Data_Ring_sub__3052583336() gopurs_runtime.Value {
	once_Data_Ring_sub__3052583336.Do(func() {
		cache_Data_Ring_sub__3052583336 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__3052583336(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ring_sub__3052583336
}

var cache_Data_Ring_sub__3277610658 gopurs_runtime.Value
var once_Data_Ring_sub__3277610658 sync.Once
func Get_Data_Ring_sub__3277610658() gopurs_runtime.Value {
	once_Data_Ring_sub__3277610658.Do(func() {
		cache_Data_Ring_sub__3277610658 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_unused_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__3277610658(__eta_norm_1_unused_0_box, __eta_norm_0_unused_1_box)
})
	})
	return cache_Data_Ring_sub__3277610658
}

var cache_Data_Ring_sub__782417183 gopurs_runtime.Value
var once_Data_Ring_sub__782417183 sync.Once
func Get_Data_Ring_sub__782417183() gopurs_runtime.Value {
	once_Data_Ring_sub__782417183.Do(func() {
		cache_Data_Ring_sub__782417183 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__782417183(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ring_sub__782417183
}

var cache_Data_Ring_sub__294757560 gopurs_runtime.Value
var once_Data_Ring_sub__294757560 sync.Once
func Get_Data_Ring_sub__294757560() gopurs_runtime.Value {
	once_Data_Ring_sub__294757560.Do(func() {
		cache_Data_Ring_sub__294757560 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__294757560(__eta_norm_1_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ring_sub__294757560
}

var cache_Data_Ring_sub__1123525868 gopurs_runtime.Value
var once_Data_Ring_sub__1123525868 sync.Once
func Get_Data_Ring_sub__1123525868() gopurs_runtime.Value {
	once_Data_Ring_sub__1123525868.Do(func() {
		cache_Data_Ring_sub__1123525868 = gopurs_runtime.Func2(func(__eta_norm_1_unused_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_sub__1123525868(__eta_norm_1_unused_0_box, __eta_norm_0_1_box)
})
	})
	return cache_Data_Ring_sub__1123525868
}

var cache_Data_Ring_ringUnit gopurs_runtime.Value
var once_Data_Ring_ringUnit sync.Once
func Get_Data_Ring_ringUnit() gopurs_runtime.Value {
	once_Data_Ring_ringUnit.Do(func() {
		cache_Data_Ring_ringUnit = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringUnit()))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})}))}
	})
	return cache_Data_Ring_ringUnit
}

var cache_Data_Ring_ringRecordNil gopurs_runtime.Value
var once_Data_Ring_ringRecordNil sync.Once
func Get_Data_Ring_ringRecordNil() gopurs_runtime.Value {
	once_Data_Ring_ringRecordNil.Do(func() {
		cache_Data_Ring_ringRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3914418263, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Semiring_semiringRecordNil()))}
}), gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value, v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := struct{

}{}
				_ = orig
				return gopurs_runtime.RecordDict0()
				}()
})}))}
	})
	return cache_Data_Ring_ringRecordNil
}

var cache_Data_Ring_ringRecordCons gopurs_runtime.Value
var once_Data_Ring_ringRecordCons sync.Once
func Get_Data_Ring_ringRecordCons() gopurs_runtime.Value {
	once_Data_Ring_ringRecordCons.Do(func() {
		cache_Data_Ring_ringRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_ringRecordCons(dictIsSymbol_0_box, _dollar___unused_1_box, dictRingRecord_2_box)
})
	})
	return cache_Data_Ring_ringRecordCons
}

var cache_Data_Ring_ringRecord gopurs_runtime.Value
var once_Data_Ring_ringRecord sync.Once
func Get_Data_Ring_ringRecord() gopurs_runtime.Value {
	once_Data_Ring_ringRecord.Do(func() {
		cache_Data_Ring_ringRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_ringRecord(_dollar___unused_0_box, dictRingRecord_1_box)
})
	})
	return cache_Data_Ring_ringRecord
}

var cache_Data_Ring_ringProxy gopurs_runtime.Value
var once_Data_Ring_ringProxy sync.Once
func Get_Data_Ring_ringProxy() gopurs_runtime.Value {
	once_Data_Ring_ringProxy.Do(func() {
		cache_Data_Ring_ringProxy = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2370240579_3060961102((&Constructor_Data_Ring_Ring[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2722024963_2826095630(Rebox_Data_Ring_2826095630_2722024963(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Get_Data_Semiring_semiringProxy()))))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Ring_ringProxy
}

var cache_Data_Ring_ringNumber gopurs_runtime.Value
var once_Data_Ring_ringNumber sync.Once
func Get_Data_Ring_ringNumber() gopurs_runtime.Value {
	once_Data_Ring_ringNumber.Do(func() {
		cache_Data_Ring_ringNumber = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2259598678_3060961102((&Constructor_Data_Ring_Ring[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_602713622_2826095630(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[float64]](Get_Data_Semiring_semiringNumber())))}
}), Get_Data_Ring_numSub()})))}
	})
	return cache_Data_Ring_ringNumber
}

var cache_Data_Ring_ringInt gopurs_runtime.Value
var once_Data_Ring_ringInt sync.Once
func Get_Data_Ring_ringInt() gopurs_runtime.Value {
	once_Data_Ring_ringInt.Do(func() {
		cache_Data_Ring_ringInt = gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_2534126933_3060961102((&Constructor_Data_Ring_Ring[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(Rebox_Data_Ring_348932501_2826095630(gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[int64]](Get_Data_Semiring_semiringInt())))}
}), Get_Data_Ring_intSub()})))}
	})
	return cache_Data_Ring_ringInt
}

var cache_Data_Ring_ringFn gopurs_runtime.Value
var once_Data_Ring_ringFn sync.Once
func Get_Data_Ring_ringFn() gopurs_runtime.Value {
	once_Data_Ring_ringFn.Do(func() {
		cache_Data_Ring_ringFn = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_ringFn(dictRing_0_box)
})
	})
	return cache_Data_Ring_ringFn
}

var cache_Data_Ring_negate gopurs_runtime.Value
var once_Data_Ring_negate sync.Once
func Get_Data_Ring_negate() gopurs_runtime.Value {
	once_Data_Ring_negate.Do(func() {
		cache_Data_Ring_negate = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_Data_Ring_negate
}

var cache_Data_Ring_negate__3883318852 gopurs_runtime.Value
var once_Data_Ring_negate__3883318852 sync.Once
func Get_Data_Ring_negate__3883318852() gopurs_runtime.Value {
	once_Data_Ring_negate__3883318852.Do(func() {
		cache_Data_Ring_negate__3883318852 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate__3883318852(__eta_norm_0_0_box)
})
	})
	return cache_Data_Ring_negate__3883318852
}

var cache_Data_Ring_negate__3140807556 gopurs_runtime.Value
var once_Data_Ring_negate__3140807556 sync.Once
func Get_Data_Ring_negate__3140807556() gopurs_runtime.Value {
	once_Data_Ring_negate__3140807556.Do(func() {
		cache_Data_Ring_negate__3140807556 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ring_negate__3140807556(__eta_norm_0_0_box)
})
	})
	return cache_Data_Ring_negate__3140807556
}

type Constructor_Data_Ring_RingRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4246880791] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_RingRecord[any, any, any])(ptr)
		_ = c
		switch key {
		case "SemiringRecord0": return gopurs_runtime.Box(c.V0)
		case "subRecord": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ring_RingRecord: " + key)
		}
	}
}


type Constructor_Data_Ring_Ring[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3955491866] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Ring_Ring[any])(ptr)
		_ = c
		switch key {
		case "Semiring0": return gopurs_runtime.Box(c.V0)
		case "sub": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Ring_Ring: " + key)
		}
	}
}


func Call_Data_Ring_semiringRecord(dictSemiringRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiringRecord_0 gopurs_runtime.Value = dictSemiringRecord_0_loop
_ = dictSemiringRecord_0
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_0, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemiringRecord_0, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_0, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiringRecord_0, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Ring_RingRecord_dollar_Dict(x_0_loop struct{
	SemiringRecord0 gopurs_runtime.Value
	subRecord gopurs_runtime.Value
}) *Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	SemiringRecord0 gopurs_runtime.Value
	subRecord gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("SemiringRecord0", "subRecord", orig.SemiringRecord0, orig.subRecord)
				}())
}

func Call_Data_Ring_Ring_dollar_Dict(x_0_loop struct{
	Semiring0 gopurs_runtime.Value
	sub gopurs_runtime.Value
}) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
var x_0 struct{
	Semiring0 gopurs_runtime.Value
	sub gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Semiring0", "sub", orig.Semiring0, orig.sub)
				}())
}

func Call_Data_Ring_subRecord(dict_0_loop *Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub(dict_0_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Data_Ring_sub__3052583336(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
sub__3052583336:
for {
if false { continue sub__3052583336 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Int((__eta_norm_1_0.IntVal) - (__eta_norm_0_1.IntVal))
}
}

func Call_Data_Ring_sub__3277610658(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_unused_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
sub__3277610658:
for {
if false { continue sub__3277610658 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_unused_1 gopurs_runtime.Value = __eta_norm_0_unused_1_loop
_ = __eta_norm_0_unused_1
return gopurs_runtime.Int((gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_topChar().StrVal())).IntVal) - (gopurs_runtime.Apply(Get_Data_Enum_toCharCode(), gopurs_runtime.Str(Get_Data_Bounded_bottomChar().StrVal())).IntVal))
}
}

func Call_Data_Ring_sub__782417183(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
sub__782417183:
for {
if false { continue sub__782417183 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Int(-(__eta_norm_0_1.IntVal))
}
}

func Call_Data_Ring_sub__294757560(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
sub__294757560:
for {
if false { continue sub__294757560 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Float((__eta_norm_1_0.FloatVal()) - (__eta_norm_0_1.FloatVal()))
}
}

func Call_Data_Ring_sub__1123525868(__eta_norm_1_unused_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
sub__1123525868:
for {
if false { continue sub__1123525868 }
var __eta_norm_1_unused_0 gopurs_runtime.Value = __eta_norm_1_unused_0_loop
_ = __eta_norm_1_unused_0
var __eta_norm_0_1 gopurs_runtime.Value = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return gopurs_runtime.Float(-(__eta_norm_0_1.FloatVal()))
}
}

func Call_Data_Ring_ringRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
_ = _dollar___unused_1
var dictRingRecord_2 gopurs_runtime.Value = dictRingRecord_2_loop
_ = dictRingRecord_2
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_2, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_3_1
// TAST (Let): semiringRecordCons1__193435443_3_0 shape=Let(Abs(Let(Let(LitRecord)))) bindingType=Any
semiringRecordCons1__193435443_3_0 := gopurs_runtime.Func(func(dictSemiring_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): one1_5_2 shape=Other bindingType=(TypeVar focus)
one1_5_2 := gopurs_runtime.RecordGet(dictSemiring_4, "one")
_ = one1_5_2
// TAST (Let): zero1_6_3 shape=Other bindingType=(TypeVar focus)
zero1_6_3 := gopurs_runtime.RecordGet(dictSemiring_4, "zero")
_ = zero1_6_3
return gopurs_runtime.Value{Type: 9, IntVal: 3914418263, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, ra_8 gopurs_runtime.Value, rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_10_4 shape=App(Other) bindingType=String
key_10_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
_ = key_10_4
// TAST (Let): get__3905987523_11_5 shape=App(Var) bindingType=(Func [(Record (Row [] (TypeVar row)))] (TypeVar focus))
get__3905987523_11_5 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_10_4))
_ = get__3905987523_11_5
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_10_4), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_4, "add"), gopurs_runtime.Apply(get__3905987523_11_5, ra_8), gopurs_runtime.Apply(get__3905987523_11_5, rb_9)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9))
}), gopurs_runtime.Func3(func(v_7 gopurs_runtime.Value, ra_8 gopurs_runtime.Value, rb_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_10_6 shape=App(Other) bindingType=String
key_10_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
_ = key_10_6
// TAST (Let): get__3905987523_11_7 shape=App(Var) bindingType=(Func [(Record (Row [] (TypeVar row)))] (TypeVar focus))
get__3905987523_11_7 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_10_6))
_ = get__3905987523_11_7
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_10_6), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_4, "mul"), gopurs_runtime.Apply(get__3905987523_11_7, ra_8), gopurs_runtime.Apply(get__3905987523_11_7, rb_9)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_8, rb_9))
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), one1_5_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}), gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()), zero1_6_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
})}))}
})
_ = semiringRecordCons1__193435443_3_0
return gopurs_runtime.Func(func(dictRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semiringRecordCons2_5_8 shape=App(Other) bindingType=(TypeApp (ADT ["Data","Semiring","SemiringRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)]), (TypeVar row), (TypeVar subrow)])
semiringRecordCons2_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_SemiringRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(semiringRecordCons1__193435443_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_4, "Semiring0"), gopurs_runtime.Value{})))
_ = semiringRecordCons2_5_8
return gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3914418263, UnsafePtr: unsafe.Pointer(semiringRecordCons2_5_8)}
}), gopurs_runtime.Func3(func(v_6 gopurs_runtime.Value, ra_7 gopurs_runtime.Value, rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): key_9_9 shape=App(Other) bindingType=String
key_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}).StrVal()
_ = key_9_9
// TAST (Let): get__3905987523_10_10 shape=App(Var) bindingType=(Func [(Record (Row [] (TypeVar row)))] (TypeVar focus))
get__3905987523_10_10 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str(key_9_9))
_ = get__3905987523_10_10
return gopurs_runtime.Apply3(Get_Record_Unsafe_unsafeSet(), gopurs_runtime.Str(key_9_9), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_4, "sub"), gopurs_runtime.Apply(get__3905987523_10_10, ra_7), gopurs_runtime.Apply(get__3905987523_10_10, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictRingRecord_2, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
})}))}
})
}

func Call_Data_Ring_ringRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictRingRecord_1 gopurs_runtime.Value = dictRingRecord_1_loop
_ = dictRingRecord_1
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): semiringRecord1_2_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Semiring","Semiring"] [(Record (Row [] (TypeVar row)))])
semiringRecord1_2_0 := (&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = semiringRecord1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringRecord1_2_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_Ring_ringFn(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semiringFn_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Semiring","Semiring"] [(Func [(TypeVar a)] (TypeVar b))])
semiringFn_1_0 := (&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "add"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mul"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "one")
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "zero")
})})
_ = semiringFn_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringFn_1_0)}
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})}))}
}

func Call_Data_Ring_negate(dictRing_0_loop *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *Constructor_Data_Ring_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
// TAST (Let): Semiring0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semiring","Semiring"] [(TypeVar a)])
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictRing_0.V0), gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictRing_0.V1), gopurs_runtime.Box(Semiring0_1_0.V3), a_2)
})
}

func Call_Data_Ring_negate__3883318852(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
negate__3883318852:
for {
if false { continue negate__3883318852 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Int(-(__eta_norm_0_0.IntVal))
}
}

func Call_Data_Ring_negate__3140807556(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
negate__3140807556:
for {
if false { continue negate__3140807556 }
var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
_ = __eta_norm_0_0
return gopurs_runtime.Float(-(__eta_norm_0_0.FloatVal()))
}
}

func Rebox_Data_Ring_2259598678_3060961102(in *Constructor_Data_Ring_Ring[float64]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ring_2370240579_3060961102(in *Constructor_Data_Ring_Ring[uint32]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ring_2534126933_3060961102(in *Constructor_Data_Ring_Ring[int64]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Ring_2722024963_2826095630(in *Constructor_Data_Semiring_Semiring[uint32]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V3), UnsafePtr: nil}
	return out
}

func Rebox_Data_Ring_2826095630_2722024963(in *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]) *Constructor_Data_Semiring_Semiring[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = uint32(in.V2.IntVal)
		out.V3 = uint32(in.V3.IntVal)
	return out
}

func Rebox_Data_Ring_348932501_2826095630(in *Constructor_Data_Semiring_Semiring[int64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Int(in.V2)
		out.V3 = gopurs_runtime.Int(in.V3)
	return out
}

func Rebox_Data_Ring_602713622_2826095630(in *Constructor_Data_Semiring_Semiring[float64]) *Constructor_Data_Semiring_Semiring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Float(in.V2)
		out.V3 = gopurs_runtime.Float(in.V3)
	return out
}

func Get_Data_Ring_intSub() gopurs_runtime.Value {
	return _Gopurs_Data_Ring_IntSub
}

func Get_Data_Ring_numSub() gopurs_runtime.Value {
	return _Gopurs_Data_Ring_NumSub
}

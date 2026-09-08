package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_CommutativeRing_ringRecord gopurs_runtime.Value
var once_Data_CommutativeRing_ringRecord sync.Once
func Get_Data_CommutativeRing_ringRecord() gopurs_runtime.Value {
	once_Data_CommutativeRing_ringRecord.Do(func() {
		cache_Data_CommutativeRing_ringRecord = gopurs_runtime.Func(func(dictRingRecord_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_ringRecord(dictRingRecord_0_box)
})
	})
	return cache_Data_CommutativeRing_ringRecord
}

var cache_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict gopurs_runtime.Value
var once_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict sync.Once
func Get_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict() gopurs_runtime.Value {
	once_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict.Do(func() {
		cache_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 292222263, UnsafePtr: unsafe.Pointer(Call_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict(func() struct{
	RingRecord0 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	RingRecord0 gopurs_runtime.Value
}{}
					clone.RingRecord0 = gopurs_runtime.RecordGet(orig, "RingRecord0")
					return clone
				}()))}
})
	})
	return cache_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict
}

var cache_Data_CommutativeRing_CommutativeRing_dollar_Dict gopurs_runtime.Value
var once_Data_CommutativeRing_CommutativeRing_dollar_Dict sync.Once
func Get_Data_CommutativeRing_CommutativeRing_dollar_Dict() gopurs_runtime.Value {
	once_Data_CommutativeRing_CommutativeRing_dollar_Dict.Do(func() {
		cache_Data_CommutativeRing_CommutativeRing_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Call_Data_CommutativeRing_CommutativeRing_dollar_Dict(func() struct{
	Ring0 gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Ring0 gopurs_runtime.Value
}{}
					clone.Ring0 = gopurs_runtime.RecordGet(orig, "Ring0")
					return clone
				}()))}
})
	})
	return cache_Data_CommutativeRing_CommutativeRing_dollar_Dict
}

var cache_Data_CommutativeRing_commutativeRingUnit gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingUnit sync.Once
func Get_Data_CommutativeRing_commutativeRingUnit() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingUnit.Do(func() {
		cache_Data_CommutativeRing_commutativeRingUnit = gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer((&Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](Get_Data_Ring_ringUnit()))}
})}))}
	})
	return cache_Data_CommutativeRing_commutativeRingUnit
}

var cache_Data_CommutativeRing_commutativeRingRecordNil gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingRecordNil sync.Once
func Get_Data_CommutativeRing_commutativeRingRecordNil() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingRecordNil.Do(func() {
		cache_Data_CommutativeRing_commutativeRingRecordNil = gopurs_runtime.Value{Type: 9, IntVal: 292222263, UnsafePtr: unsafe.Pointer((&Constructor_Data_CommutativeRing_CommutativeRingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Ring_ringRecordNil()))}
})}))}
	})
	return cache_Data_CommutativeRing_commutativeRingRecordNil
}

var cache_Data_CommutativeRing_commutativeRingRecordCons gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingRecordCons sync.Once
func Get_Data_CommutativeRing_commutativeRingRecordCons() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingRecordCons.Do(func() {
		cache_Data_CommutativeRing_commutativeRingRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar___unused_1_box gopurs_runtime.Value, dictCommutativeRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_commutativeRingRecordCons(dictIsSymbol_0_box, _dollar___unused_1_box, dictCommutativeRingRecord_2_box)
})
	})
	return cache_Data_CommutativeRing_commutativeRingRecordCons
}

var cache_Data_CommutativeRing_commutativeRingRecord gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingRecord sync.Once
func Get_Data_CommutativeRing_commutativeRingRecord() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingRecord.Do(func() {
		cache_Data_CommutativeRing_commutativeRingRecord = gopurs_runtime.Func2(func(_dollar___unused_0_box gopurs_runtime.Value, dictCommutativeRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_commutativeRingRecord(_dollar___unused_0_box, dictCommutativeRingRecord_1_box)
})
	})
	return cache_Data_CommutativeRing_commutativeRingRecord
}

var cache_Data_CommutativeRing_commutativeRingProxy gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingProxy sync.Once
func Get_Data_CommutativeRing_commutativeRingProxy() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingProxy.Do(func() {
		cache_Data_CommutativeRing_commutativeRingProxy = gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_CommutativeRing_2537586851_1073849710((&Constructor_Data_CommutativeRing_CommutativeRing[uint32]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_CommutativeRing_2370240579_3060961102(Rebox_Data_CommutativeRing_3060961102_2370240579(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](Get_Data_Ring_ringProxy()))))}
})})))}
	})
	return cache_Data_CommutativeRing_commutativeRingProxy
}

var cache_Data_CommutativeRing_commutativeRingNumber gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingNumber sync.Once
func Get_Data_CommutativeRing_commutativeRingNumber() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingNumber.Do(func() {
		cache_Data_CommutativeRing_commutativeRingNumber = gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_CommutativeRing_2766724982_1073849710((&Constructor_Data_CommutativeRing_CommutativeRing[float64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_CommutativeRing_2259598678_3060961102(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[float64]](Get_Data_Ring_ringNumber())))}
})})))}
	})
	return cache_Data_CommutativeRing_commutativeRingNumber
}

var cache_Data_CommutativeRing_commutativeRingInt gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingInt sync.Once
func Get_Data_CommutativeRing_commutativeRingInt() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingInt.Do(func() {
		cache_Data_CommutativeRing_commutativeRingInt = gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(Rebox_Data_CommutativeRing_357988789_1073849710((&Constructor_Data_CommutativeRing_CommutativeRing[int64]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(Rebox_Data_CommutativeRing_2534126933_3060961102(gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[int64]](Get_Data_Ring_ringInt())))}
})})))}
	})
	return cache_Data_CommutativeRing_commutativeRingInt
}

var cache_Data_CommutativeRing_commutativeRingFn gopurs_runtime.Value
var once_Data_CommutativeRing_commutativeRingFn sync.Once
func Get_Data_CommutativeRing_commutativeRingFn() gopurs_runtime.Value {
	once_Data_CommutativeRing_commutativeRingFn.Do(func() {
		cache_Data_CommutativeRing_commutativeRingFn = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_CommutativeRing_commutativeRingFn(dictCommutativeRing_0_box)
})
	})
	return cache_Data_CommutativeRing_commutativeRingFn
}

type Constructor_Data_CommutativeRing_CommutativeRingRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[292222263] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_CommutativeRing_CommutativeRingRecord[any, any, any])(ptr)
		_ = c
		switch key {
		case "RingRecord0": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_CommutativeRing_CommutativeRingRecord: " + key)
		}
	}
}


type Constructor_Data_CommutativeRing_CommutativeRing[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1775085946] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_CommutativeRing_CommutativeRing[any])(ptr)
		_ = c
		switch key {
		case "Ring0": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_CommutativeRing_CommutativeRing: " + key)
		}
	}
}


func Call_Data_CommutativeRing_ringRecord(dictRingRecord_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRingRecord_0 gopurs_runtime.Value = dictRingRecord_0_loop
_ = dictRingRecord_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_0, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semiringRecord1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Data","Semiring","Semiring"] [(Record (Row [] (TypeVar row)))])
semiringRecord1_1_0 := (&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = semiringRecord1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringRecord1_1_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_0, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})}))}
}

func Call_Data_CommutativeRing_CommutativeRingRecord_dollar_Dict(x_0_loop struct{
	RingRecord0 gopurs_runtime.Value
}) *Constructor_Data_CommutativeRing_CommutativeRingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	RingRecord0 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"RingRecord0"}, []gopurs_runtime.Value{orig.RingRecord0})
				}())
}

func Call_Data_CommutativeRing_CommutativeRing_dollar_Dict(x_0_loop struct{
	Ring0 gopurs_runtime.Value
}) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
var x_0 struct{
	Ring0 gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Ring0"}, []gopurs_runtime.Value{orig.Ring0})
				}())
}

func Call_Data_CommutativeRing_commutativeRingRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar___unused_1_loop gopurs_runtime.Value, dictCommutativeRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar___unused_1 gopurs_runtime.Value = _dollar___unused_1_loop
_ = _dollar___unused_1
var dictCommutativeRingRecord_2 gopurs_runtime.Value = dictCommutativeRingRecord_2_loop
_ = dictCommutativeRingRecord_2
// TAST (Let): ringRecordCons1__193435443_3_0 shape=App(Var) bindingType=Any
ringRecordCons1__193435443_3_0 := gopurs_runtime.Apply3(Get_Data_Ring_ringRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRingRecord_2, "RingRecord0"), gopurs_runtime.Value{}))
_ = ringRecordCons1__193435443_3_0
return gopurs_runtime.Func(func(dictCommutativeRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): ringRecordCons2_5_1 shape=App(Other) bindingType=(TypeApp (ADT ["Data","Ring","RingRecord"] []) [(TypeApp (ADT ["Prim","RowList","Cons"] []) [(TypeVar key), (TypeVar focus), (TypeVar rowlistTail)]), (TypeVar row), (TypeVar subrow)])
ringRecordCons2_5_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(ringRecordCons1__193435443_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_4, "Ring0"), gopurs_runtime.Value{})))
_ = ringRecordCons2_5_1
return gopurs_runtime.Value{Type: 9, IntVal: 292222263, UnsafePtr: unsafe.Pointer((&Constructor_Data_CommutativeRing_CommutativeRingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4246880791, UnsafePtr: unsafe.Pointer(ringRecordCons2_5_1)}
})}))}
})
}

func Call_Data_CommutativeRing_commutativeRingRecord(_dollar___unused_0_loop gopurs_runtime.Value, dictCommutativeRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var dictCommutativeRingRecord_1 gopurs_runtime.Value = dictCommutativeRingRecord_1_loop
_ = dictCommutativeRingRecord_1
// TAST (Let): __local_var_2_1 shape=App(Other) bindingType=Any
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRingRecord_1, "RingRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): __local_var_3_3 shape=App(Other) bindingType=Any
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): semiringRecord1_3_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Semiring","Semiring"] [(Record (Row [] (TypeVar row)))])
semiringRecord1_3_2 := (&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = semiringRecord1_3_2
// TAST (Let): ringRecord1_2_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Ring","Ring"] [(Record (Row [] (TypeVar row)))])
ringRecord1_2_0 := (&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringRecord1_3_2)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})})
_ = ringRecord1_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer((&Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(ringRecord1_2_0)}
})}))}
}

func Call_Data_CommutativeRing_commutativeRingFn(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): semiringFn_2_2 shape=Let(LitRecord) bindingType=(ADT ["Data","Semiring","Semiring"] [(Func [(TypeVar a)] (TypeVar b))])
semiringFn_2_2 := (&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "add"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "mul"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_2_3, "one")
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_2_3, "zero")
})})
_ = semiringFn_2_2
// TAST (Let): ringFn_1_0 shape=Let(Let(LitRecord)) bindingType=(ADT ["Data","Ring","Ring"] [(Func [(TypeVar a)] (TypeVar b))])
ringFn_1_0 := (&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringFn_2_2)}
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "sub"), gopurs_runtime.Apply(f_3, x_5), gopurs_runtime.Apply(g_4, x_5))
})})
_ = ringFn_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer((&Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(ringFn_1_0)}
})}))}
}

func Rebox_Data_CommutativeRing_2259598678_3060961102(in *Constructor_Data_Ring_Ring[float64]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_CommutativeRing_2370240579_3060961102(in *Constructor_Data_Ring_Ring[uint32]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_CommutativeRing_2534126933_3060961102(in *Constructor_Data_Ring_Ring[int64]) *Constructor_Data_Ring_Ring[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_CommutativeRing_2537586851_1073849710(in *Constructor_Data_CommutativeRing_CommutativeRing[uint32]) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_CommutativeRing_2766724982_1073849710(in *Constructor_Data_CommutativeRing_CommutativeRing[float64]) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_CommutativeRing_3060961102_2370240579(in *Constructor_Data_Ring_Ring[gopurs_runtime.Value]) *Constructor_Data_Ring_Ring[uint32] {
	if in == nil { return nil }
	out := &Constructor_Data_Ring_Ring[uint32]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_CommutativeRing_357988789_1073849710(in *Constructor_Data_CommutativeRing_CommutativeRing[int64]) *Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}



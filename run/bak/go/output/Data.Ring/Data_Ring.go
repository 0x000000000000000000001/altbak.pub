package Data_Ring

import (
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Symbol "gopurs/output/Data.Symbol"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Record_Unsafe "gopurs/output/Record.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_subRecord gopurs_runtime.Value
var once_subRecord sync.Once
func Get_subRecord() gopurs_runtime.Value {
	once_subRecord.Do(func() {
		cache_subRecord = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_subRecord(gopurs_runtime.CoerceToStruct[Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_subRecord
}

var cache_sub gopurs_runtime.Value
var once_sub sync.Once
func Get_sub() gopurs_runtime.Value {
	once_sub.Do(func() {
		cache_sub = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub(gopurs_runtime.CoerceToStruct[Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub
}

var cache_ringUnit gopurs_runtime.Value
var once_ringUnit sync.Once
func Get_ringUnit() gopurs_runtime.Value {
	once_ringUnit.Do(func() {
		cache_ringUnit = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringUnit()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
})
}))
	})
	return cache_ringUnit
}

var cache_ringRecordNil gopurs_runtime.Value
var once_ringRecordNil sync.Once
func Get_ringRecordNil() gopurs_runtime.Value {
	once_ringRecordNil.Do(func() {
		cache_ringRecordNil = gopurs_runtime.RecordDict2("SemiringRecord0", "subRecord", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringRecordNil()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict0()
})
})
}))
	})
	return cache_ringRecordNil
}

var cache_ringRecordCons gopurs_runtime.Value
var once_ringRecordCons sync.Once
func Get_ringRecordCons() gopurs_runtime.Value {
	once_ringRecordCons.Do(func() {
		cache_ringRecordCons = gopurs_runtime.Func3(func(dictIsSymbol_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, dictRingRecord_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ringRecordCons(dictIsSymbol_0_box, _dollar__unused_1_box, dictRingRecord_2_box)
})
	})
	return cache_ringRecordCons
}

var cache_ringRecord gopurs_runtime.Value
var once_ringRecord sync.Once
func Get_ringRecord() gopurs_runtime.Value {
	once_ringRecord.Do(func() {
		cache_ringRecord = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, dictRingRecord_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ringRecord(_dollar__unused_0_box, dictRingRecord_1_box)
})
	})
	return cache_ringRecord
}

var cache_ringProxy gopurs_runtime.Value
var once_ringProxy sync.Once
func Get_ringProxy() gopurs_runtime.Value {
	once_ringProxy.Do(func() {
		cache_ringProxy = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Semiring.Get_semiringProxy()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}))
	})
	return cache_ringProxy
}

var cache_ringNumber gopurs_runtime.Value
var once_ringNumber sync.Once
func Get_ringNumber() gopurs_runtime.Value {
	once_ringNumber.Do(func() {
		cache_ringNumber = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", pkg_Data_Semiring.Get_numAdd(), pkg_Data_Semiring.Get_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
}), Get_numSub())
	})
	return cache_ringNumber
}

var cache_ringInt gopurs_runtime.Value
var once_ringInt sync.Once
func Get_ringInt() gopurs_runtime.Value {
	once_ringInt.Do(func() {
		cache_ringInt = gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", pkg_Data_Semiring.Get_intAdd(), pkg_Data_Semiring.Get_intMul(), gopurs_runtime.Int(1), gopurs_runtime.Int(0))
}), Get_intSub())
	})
	return cache_ringInt
}

var cache_ringFn gopurs_runtime.Value
var once_ringFn sync.Once
func Get_ringFn() gopurs_runtime.Value {
	once_ringFn.Do(func() {
		cache_ringFn = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ringFn(dictRing_0_box)
})
	})
	return cache_ringFn
}

var cache_negate gopurs_runtime.Value
var once_negate sync.Once
func Get_negate() gopurs_runtime.Value {
	once_negate.Do(func() {
		cache_negate = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_negate(gopurs_runtime.CoerceToStruct[Constructor_Ring[gopurs_runtime.Value]](dictRing_0_box))
})
	})
	return cache_negate
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_subRecord__3826282112 gopurs_runtime.Value
var once_subRecord__3826282112 sync.Once
func Get_subRecord__3826282112() gopurs_runtime.Value {
	once_subRecord__3826282112.Do(func() {
		cache_subRecord__3826282112 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_subRecord__3826282112(gopurs_runtime.CoerceToStruct[Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_subRecord__3826282112
}

var cache_subRecord__2650724742 gopurs_runtime.Value
var once_subRecord__2650724742 sync.Once
func Get_subRecord__2650724742() gopurs_runtime.Value {
	once_subRecord__2650724742.Do(func() {
		cache_subRecord__2650724742 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_subRecord__2650724742(gopurs_runtime.CoerceToStruct[Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_subRecord__2650724742
}

var cache_semiringProxy__786350887 gopurs_runtime.Value
var once_semiringProxy__786350887 sync.Once
func Get_semiringProxy__786350887() gopurs_runtime.Value {
	once_semiringProxy__786350887.Do(func() {
		cache_semiringProxy__786350887 = gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}
})
}), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
	})
	return cache_semiringProxy__786350887
}

var cache_semiringRecordNil__3131124892 gopurs_runtime.Value
var once_semiringRecordNil__3131124892 sync.Once
func Get_semiringRecordNil__3131124892() gopurs_runtime.Value {
	once_semiringRecordNil__3131124892.Do(func() {
		cache_semiringRecordNil__3131124892 = gopurs_runtime.RecordDict4("addRecord", "mulRecord", "oneRecord", "zeroRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.RecordDict0()
})
}))
	})
	return cache_semiringRecordNil__3131124892
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
}

var cache_reflectSymbol__3416619207 gopurs_runtime.Value
var once_reflectSymbol__3416619207 sync.Once
func Get_reflectSymbol__3416619207() gopurs_runtime.Value {
	once_reflectSymbol__3416619207.Do(func() {
		cache_reflectSymbol__3416619207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__3416619207(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__3416619207
}

var cache_reflectSymbol__1166932993 gopurs_runtime.Value
var once_reflectSymbol__1166932993 sync.Once
func Get_reflectSymbol__1166932993() gopurs_runtime.Value {
	once_reflectSymbol__1166932993.Do(func() {
		cache_reflectSymbol__1166932993 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__1166932993(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__1166932993
}

type Constructor_RingRecord[T_rowlist any, T_row any, T_subrow any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4246880791] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "SemiringRecord0": return c.V0
		case "subRecord": return c.V1
		default: panic("Key not found in dictionary Constructor_RingRecord: " + key)
		}
	}
}


type Constructor_Ring[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3955491866] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Ring[gopurs_runtime.Value])(ptr)
		switch key {
		case "Semiring0": return c.V0
		case "sub": return c.V1
		default: panic("Key not found in dictionary Constructor_Ring: " + key)
		}
	}
}


func Call_subRecord(dict_0_loop *Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub(dict_0_loop *Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_ringRecordCons(dictIsSymbol_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, dictRingRecord_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var dictRingRecord_2 gopurs_runtime.Value = dictRingRecord_2_loop
_ = dictRingRecord_2
semiringRecordCons1_3_0 := gopurs_runtime.Apply3(pkg_Data_Semiring.Get_semiringRecordCons(), dictIsSymbol_0, gopurs_runtime.Value{}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_2, "SemiringRecord0"), gopurs_runtime.Value{}))
_ = semiringRecordCons1_3_0
return gopurs_runtime.Func(func(dictRing_4 gopurs_runtime.Value) gopurs_runtime.Value {
semiringRecordCons2_5_1 := gopurs_runtime.Apply(semiringRecordCons1_3_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_4, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringRecordCons2_5_1
return gopurs_runtime.RecordDict2("SemiringRecord0", "subRecord", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecordCons2_5_1
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_8 gopurs_runtime.Value) gopurs_runtime.Value {
key_9_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
_ = key_9_2
get_10_3 := gopurs_runtime.Apply(pkg_Record_Unsafe.Get_unsafeGet(), gopurs_runtime.Str(key_9_2.StrVal()))
_ = get_10_3
return gopurs_runtime.Apply3(pkg_Record_Unsafe.Get_unsafeSet(), gopurs_runtime.Str(key_9_2.StrVal()), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_4, "sub"), gopurs_runtime.Apply(get_10_3, ra_7), gopurs_runtime.Apply(get_10_3, rb_8)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictRingRecord_2, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_7, rb_8))
})
})
}))
})
}

func Call_ringRecord(_dollar__unused_0_loop gopurs_runtime.Value, dictRingRecord_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var dictRingRecord_1 gopurs_runtime.Value = dictRingRecord_1_loop
_ = dictRingRecord_1
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "SemiringRecord0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semiringRecord1_2_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "addRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "mulRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "oneRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "zeroRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
_ = semiringRecord1_2_0
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringRecord1_2_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRingRecord_1, "subRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}))
}

func Call_ringFn(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})
_ = __local_var_1_1
semiringFn_1_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "add"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mul"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "one")
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(__local_var_1_1, "zero")
}))
_ = semiringFn_1_0
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringFn_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}))
}

func Call_negate(dictRing_0_loop *Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictRing_0 *Constructor_Ring[gopurs_runtime.Value] = dictRing_0_loop
_ = dictRing_0
Semiring0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](gopurs_runtime.Apply(dictRing_0.V0, gopurs_runtime.Value{}))
_ = Semiring0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictRing_0.V1, Semiring0_1_0.V3, a_2)
})
}

func Call_sub__3675938712(dict_0_loop *Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_subRecord__3826282112(dict_0_loop *Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_subRecord__2650724742(dict_0_loop *Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_RingRecord[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_reflectSymbol__3416619207(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__1166932993(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_intSub() gopurs_runtime.Value {
	return _Gopurs_IntSub
}

func Get_numSub() gopurs_runtime.Value {
	return _Gopurs_NumSub
}
